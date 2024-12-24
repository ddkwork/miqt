package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/ddkwork/golibrary/mylog"
)

const (
	ClangMaxRetries = 1
	ClangRetryDelay = time.Second
)

type ClangMatcher func(astNodeFilename string) bool

func ClangMatchSameHeaderDefinitionOnly(astNodeFilename string) bool {
	return astNodeFilename == ""
}

type clangMatchUnderPath struct {
	basePath string
}

func (c *clangMatchUnderPath) Match(astNodeFilename string) bool {
	if astNodeFilename == "" {
		return true
	}
	return strings.HasPrefix(astNodeFilename, c.basePath)
}

//

func clangExec(ctx context.Context, inputHeader string, cflags []string, matcher ClangMatcher) ([]any, error) {
	clangArgs := []string{`-x`, `c++`}
	clangArgs = append(clangArgs, cflags...)
	clangArgs = append(clangArgs, `-Xclang`, `-ast-dump=json`, `-fsyntax-only`, inputHeader)
	log.Println("clang " + strings.Join(clangArgs, " "))
	cmd := exec.CommandContext(ctx, clangBin, clangArgs...)
	pr := mylog.Check2(cmd.StdoutPipe())

	cmd.Stderr = os.Stderr
	mylog.Check(cmd.Start())

	var wg sync.WaitGroup
	var inner []any
	var innerErr error

	wg.Add(1)
	go func() {
		defer wg.Done()
		inner, innerErr = clangStripUpToFile(pr, matcher)
	}()
	mylog.Check(cmd.Wait())

	wg.Wait()

	return inner, innerErr
}

func mustClangExec(ctx context.Context, inputHeader string, cflags []string, matcher ClangMatcher) []any {
	for i := 0; i < ClangMaxRetries; i++ {
		astInner := mylog.Check2(clangExec(ctx, inputHeader, cflags, matcher))

		// Log and continue with next retry

		// time.Sleep(ClangRetryDelay)

		// Success
		return astInner
	}

	// Failed 5x
	// Panic
	panic("Clang failed 5x parsing file " + inputHeader)
}

// clangStripUpToFile strips all AST nodes from the clang output until we find
// one that really originated in the source file.
// This cleans out everything in the translation unit that came from an
// #included file.
// @ref https://stackoverflow.com/a/71128654
func clangStripUpToFile(stdout io.Reader, matcher ClangMatcher) ([]any, error) {
	obj := map[string]any{}
	mylog.Check(json.NewDecoder(stdout).Decode(&obj))

	inner, ok := obj["inner"].([]any)
	if !ok {
		return nil, errors.New("no inner")
	}

	// This can't be done by matching the first position only, since it's possible
	// that there are more #include<>s further down the file

	ret := make([]any, 0, len(inner))

	for _, entry := range inner {

		entry, ok := entry.(map[string]any)
		if !ok {
			return nil, errors.New("entry is not a map")
		}

		// Check where this AST node came from, if it was directly written
		// in this header or if it as part of an #include

		match_filename := ""

		if loc, ok := entry["loc"].(map[string]any); ok {
			if includedFrom, ok := loc["includedFrom"].(map[string]any); ok {
				if filename, ok := includedFrom["file"].(string); ok {
					match_filename = filename
				}
			}

			if match_filename == "" {
				if expansionloc, ok := loc["expansionLoc"].(map[string]any); ok {
					if filename, ok := expansionloc["file"].(string); ok {
						match_filename = filename
					} else if includedFrom, ok := expansionloc["includedFrom"].(map[string]any); ok {
						if filename, ok := includedFrom["file"].(string); ok {
							match_filename = filename
						}
					}
				}
			}
		} else {
			return nil, errors.New("no loc")
		}

		// log.Printf("# name=%v kind=%v filename=%q\n", entry["name"], entry["kind"], match_filename)

		if matcher(match_filename) {
			// Keep
			ret = append(ret, entry)
		}

		// Otherwise, discard this AST node, it comes from some imported file
		// that we will likely scan separately
	}

	return ret, nil
}
