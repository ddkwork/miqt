package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func parseHeader(inner []interface{}) (*CppParsedHeader, error) {

	var ret CppParsedHeader

	for _, node := range inner {

		node, ok := node.(map[string]interface{})
		if !ok {
			return nil, errors.New("inner[] element not an object")
		}

		kind, ok := node["kind"].(string)
		if !ok {
			return nil, errors.New("node has no kind")
		}

		switch kind {

		case "CXXRecordDecl":
			// Must have a name
			nodename, ok := node["name"].(string)
			if !ok {
				return nil, errors.New("node has no name")
			}

			fmt.Printf("-> %q name=%q\n", kind, nodename)
			if classInner, ok := node["inner"].([]interface{}); ok {

				// Check if this was 'struct' (default visible) or 'class' (default invisible)
				visible := true
				if tagUsed, ok := node["tagUsed"].(string); ok && tagUsed == "class" {
					visible = false
				}

				// Process the inner class definition
				obj, err := processType(classInner, nodename, visible)
				if err != nil {
					panic(err)
				}

				ret.Classes = append(ret.Classes, obj)
			}

		case "StaticAssertDecl":
			// ignore

		default:
			return nil, fmt.Errorf("missing handling for clang ast node type %q", kind)
		}
	}

	return &ret, nil // done
}

func processType(inner []interface{}, className string, visibility bool) (CppClass, error) {
	var ret CppClass
	ret.ClassName = className

nextMethod:
	for _, node := range inner {
		node, ok := node.(map[string]interface{})
		if !ok {
			return CppClass{}, errors.New("inner[] element not an object")
		}

		kind, ok := node["kind"]
		if !ok {
			panic("inner element has no kind")
		}

		switch kind {
		case "AccessSpecDecl":
			// Swap between visible/invisible
			access, ok := node["access"].(string)
			if !ok {
				panic("AccessSpecDecl missing `access` field")
			}

			switch access {
			case "public":
				visibility = true
			case "private", "protected":
				visibility = false
			default:
				panic("unexpected access visibility '" + access + "'")
			}

		case "FriendDecl":
			// Safe to ignore

		case "CXXConstructorDecl":
			// panic("TODO")

		case "CXXDestructorDecl":
			// panic("TODO")

		case "CXXMethodDecl":
			if !visibility {
				continue // Skip private/protected
			}

			// Method
			methodName, ok := node["name"].(string)
			if !ok {
				return CppClass{}, errors.New("method has no name")
			}

			var mm CppMethod
			mm.MethodName = methodName

			if typobj, ok := node["type"].(map[string]interface{}); ok {
				if qualType, ok := typobj["qualType"].(string); ok {
					// The qualType is the whole type of the method, including its parameter types
					// If anything here is too complicated, skip the whole method

					var err error = nil
					mm.ReturnType, mm.Parameters, err = parseTypeString(qualType)
					if err != nil {
						if errors.Is(err, ErrTooComplex) {
							log.Printf("Skipping method %q with complex type %q", mm.MethodName, qualType)
							continue nextMethod
						}
						// Real error
						return CppClass{}, err
					}

				}
			}

			if methodInner, ok := node["inner"].([]interface{}); ok {
				paramCounter := 0
				for _, methodObj := range methodInner {
					methodObj, ok := methodObj.(map[string]interface{})
					if !ok {
						return CppClass{}, errors.New("inner[] element not an object")
					}

					switch methodObj["kind"] {
					case "ParmVarDecl":
						// Parameter variable
						parmName, _ := methodObj["name"].(string) // n.b. may be unnamed
						if parmName == "" {

							// Generate a default parameter name
							// Super nice autogen names if this is a Q_PROPERTY setter:
							if len(mm.Parameters) == 1 && strings.HasPrefix(mm.MethodName, "set") {
								parmName = strings.ToLower(string(mm.MethodName[3])) + mm.MethodName[4:]

							} else {
								// Otherwise - default
								parmName = fmt.Sprintf("param%d", paramCounter+1)
							}
						}

						// Block reserved Go words, replace with generic parameters
						if parmName == "default" || parmName == "const" || parmName == "func" {
							parmName += "Val"
						}

						// Update the name for the existing nth parameter
						mm.Parameters[paramCounter].ParameterName = parmName

						// If this parameter has any internal AST nodes of its
						// own, assume it means it's an optional parameter
						if _, ok := methodObj["inner"]; ok {
							mm.Parameters[paramCounter].Optional = true
						}

						// Next
						paramCounter++

					default:
						// Something else inside a declaration??
						fmt.Printf("==> NOT IMPLEMENTED CXXMethodDecl->%q\n", kind)
					}
				}
			}

			ret.Methods = append(ret.Methods, mm)

		default:
			fmt.Printf("==> NOT IMPLEMENTED %q\n", kind)
		}
	}

	return ret, nil // done
}

var ErrTooComplex error = errors.New("Type declaration is too complex to parse")

// parseTypeString converts a string like
// - `QString (const char *, const char *, int)`
// - `void (const QKeySequence \u0026)`
// into its (A) return type and (B) separate parameter types.
// These clang strings never contain the parameter's name, so the names here are
// not filled in.
func parseTypeString(typeString string) (CppParameter, []CppParameter, error) {

	if strings.Contains(typeString, `::`) {
		return CppParameter{}, nil, ErrTooComplex
	}

	// Cut to exterior-most (, ) pair
	opos := strings.Index(typeString, `(`)
	epos := strings.LastIndex(typeString, `)`)

	if opos == -1 || epos == -1 {
		return CppParameter{}, nil, fmt.Errorf("Type string %q missing brackets")
	}

	returnType := parseSingleTypeString(strings.TrimSpace(typeString[0:opos]))

	inner := typeString[opos+1 : epos]

	// Should be no more brackets
	if strings.ContainsAny(inner, `()`) {
		return CppParameter{}, nil, ErrTooComplex
	}

	// Parameters are separated by commas and nesting can not be possible
	params := strings.Split(inner, `,`)

	ret := make([]CppParameter, 0, len(params))
	for _, p := range params {

		insert := parseSingleTypeString(p)

		if insert.ParameterType != "" {
			ret = append(ret, insert)
		}
	}

	return returnType, ret, nil
}

func parseSingleTypeString(p string) CppParameter {

	tokens := strings.Split(strings.TrimSpace(p), " ")
	insert := CppParameter{}
	for _, tok := range tokens {
		if tok == "const" {
			insert.Const = true
		} else if tok == "&" { // U+0026
			insert.ByRef = true
		} else if tok == "*" {
			insert.Pointer = true
		} else {
			// Valid part of the type name
			insert.ParameterType += " " + tok
		}
	}
	insert.ParameterType = strings.TrimSpace(insert.ParameterType)

	return insert
}