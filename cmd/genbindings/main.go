package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/ddkwork/golibrary/mylog"
)

const (
	qt6Dir               = "D:\\qt6\\qt_static\\" // Qt6_DIR use environment variable instead ?
	ClangSubprocessCount = 2
	BaseModule           = "github.com/mappu/miqt"
	clangBin             = "clang" // from environment variable instead,need check it in main function
)

func main() {
	os.Getenv("Qt6_DIR")
	exec.Command("clang -v") // todo panic it if not found
	os.RemoveAll(qt6Dir + "lib\\pkgconfig")
	mylog.Check(os.CopyFS(qt6Dir+"lib\\pkgconfig", os.DirFS("pkgconfig")))

	outDir := flag.String("outdir", "../../", "Output directory for generated gen_** files")
	flag.Parse()
	ProcessLibraries(*outDir)
}

func exportPrefix() string {
	switch runtime.GOOS {
	case "windows":
		return "extern __declspec(dllexport) "
		// case "darwin":
		// case "linux":
		// case "freebsd":
		// case "openbsd":
		// case "android":
	}
	panic("not supported os") // todo
}

func cacheFilePath(inputHeader string) string {
	//rel, err := filepath.Rel(inputHeader, "D:\\qt6\\qt_static")
	//if err != nil {
	//	panic(err)
	//}
	inputHeader = strings.TrimPrefix(inputHeader, qt6Dir+"include") // todo make uint test
	s := filepath.Join("cachedir", strings.Replace(inputHeader, `/`, `__`, -1)+".json")
	os.MkdirAll(filepath.Dir(s), 0o777) //? why
	return s
}

func importPathForQtPackage(packageName string) string {
	return BaseModule + "/" + packageName
}

func findHeadersInDir(srcDir string, allowHeader func(string) bool) []string {
	content := mylog.Check2(os.ReadDir(srcDir))

	var ret []string
	for _, includeFile := range content {
		if includeFile.IsDir() {
			continue
		}
		if !strings.HasSuffix(includeFile.Name(), `.h`) {
			continue
		}
		fullPath := filepath.Join(srcDir, includeFile.Name())
		if !allowHeader(fullPath) {
			continue
		}
		ret = append(ret, fullPath)
	}
	return ret
}

func cleanGeneratedFilesInDir(dirpath string) {
	log.Printf("Cleaning up output directory %q...", dirpath)
	_ = os.MkdirAll(dirpath, 0o755)
	existing := mylog.Check2(os.ReadDir(dirpath))

	cleaned := 0
	for _, e := range existing {
		if e.IsDir() {
			continue
		}
		if !strings.HasPrefix(e.Name(), `gen_`) {
			continue
		}
		mylog.
			// One of ours, clean up
			Check(os.Remove(filepath.Join(dirpath, e.Name())))

		cleaned++
	}
	log.Printf("Removed %d file(s).", cleaned)
}

func pkgConfigCflags(packageName string) string {
	cmd := exec.Command(`pkg-config`, `--cflags`, packageName)
	stdout := mylog.Check2(cmd.Output())

	return string(stdout)
}

func generate(packageName string, srcDirs []string, allowHeaderFn func(string) bool, cflagsCombined, outDir string, matcher ClangMatcher) {
	var includeFiles []string
	for _, srcDir := range srcDirs {
		if strings.HasSuffix(srcDir, `.h`) {
			includeFiles = append(includeFiles, srcDir) // single .h
		} else {
			includeFiles = append(includeFiles, findHeadersInDir(srcDir, allowHeaderFn)...)
		}
	}
	log.Printf("Found %d header files to process.", len(includeFiles))
	cflags := strings.Fields(cflagsCombined)
	outDir = filepath.Join(outDir, packageName)
	cleanGeneratedFilesInDir(outDir)
	var processHeaders []*CppParsedHeader
	atr := astTransformRedundant{
		preserve: make(map[string]*CppEnum),
	}
	// PASS 0 (Fill clang cache)
	generateClangCaches(includeFiles, cflags, matcher)
	// The cache should now be fully populated.
	// PASS 1 (clang2il)
	// todo add test
	for _, inputHeader := range includeFiles {
		cacheFile := cacheFilePath(inputHeader)
		astJson := mylog.Check2(os.ReadFile(cacheFile))

		// Json decode
		var astInner []any = nil
		mylog.Check(json.Unmarshal(astJson, &astInner))

		if astInner == nil {
			panic("Null in cache file for " + inputHeader)
		}
		// Convert it to our intermediate format
		parsed := mylog.Check2(parseHeader(astInner, ""))

		// todo fix me

		// panic(err)

		parsed.Filename = inputHeader // Stash
		// AST transforms on our IL
		astTransformChildClasses(parsed) // must be first
		astTransformOptional(parsed)
		astTransformOverloads(parsed)
		astTransformConstructorOrder(parsed)
		atr.Process(parsed)
		// Update global state tracker (AFTER astTransformChildClasses)
		addKnownTypes(packageName, parsed)
		processHeaders = append(processHeaders, parsed)
	}

	// PASS 2
	for _, parsed := range processHeaders {
		log.Printf("Processing %q...", parsed.Filename)
		// More AST transforms on our IL
		astTransformTypedefs(parsed)
		astTransformBlocklist(parsed) // Must happen after typedef transformation
		{
			// Save the IL file for debug inspection
			jb := mylog.Check2(json.MarshalIndent(parsed, "", "\t"))
			mylog.Check(os.WriteFile(cacheFilePath(parsed.Filename)+".ours.json", jb, 0o644))

		}
		// Breakout if there is nothing bindable
		if parsed.Empty() {
			log.Printf("Nothing in this header was bindable.")
			continue
		}
		// todo move .h .cpp into another dir,because we only need build c abi dll once
		// Emit 3 code files from the intermediate format
		outputName := filepath.Join(outDir, "gen_"+strings.TrimSuffix(filepath.Base(parsed.Filename), `.h`))
		// For packages where we scan multiple directories, it's possible that
		// there are filename collisions (e.g. Qt 6 has QtWidgets/qaction.h include
		// QtGui/qaction.h as a compatibility measure).
		// If the path exists, disambiguate it
		counter := 0
		for {
			testName := outputName
			if counter > 0 {
				testName += fmt.Sprintf(".%d", counter)
			}
			if _, e := os.Stat(testName + ".go"); e != nil && os.IsNotExist(e) {
				outputName = testName // Safe
				break
			}
			counter++
		}
		goSrc := mylog.Check2(emitGo(parsed, filepath.Base(parsed.Filename), packageName))
		mylog.Check(os.WriteFile(outputName+".go", []byte(goSrc), 0o644))
		bindingCppSrc := mylog.Check2(emitBindingCpp(parsed, filepath.Base(parsed.Filename)))
		mylog.Check(os.WriteFile(outputName+".cpp", []byte(bindingCppSrc), 0o644))
		bindingHSrc := mylog.Check2(emitBindingHeader(parsed, filepath.Base(parsed.Filename), packageName))
		mylog.Check(os.WriteFile(outputName+".h", []byte(bindingHSrc), 0o644))
		// Done
	}
	log.Printf("Processing %d file(s) completed", len(includeFiles))
}

func generateClangCaches(includeFiles []string, cflags []string, matcher ClangMatcher) {
	clangChan := make(chan string)
	var clangWg sync.WaitGroup
	ctx := context.Background()
	for i := 0; i < ClangSubprocessCount; i++ {
		clangWg.Add(1)
		go func() {
			defer clangWg.Done()
			log.Printf("Clang worker: starting")
			for {
				inputHeader, ok := <-clangChan
				if !ok {
					return // Done
				}
				log.Printf("Clang worker got message for file %q", inputHeader)
				// Parse the file
				// This seems to intermittently fail, so allow retrying
				astInner := mustClangExec(ctx, inputHeader, cflags, matcher)
				// Write to cache
				jb := mylog.Check2(json.MarshalIndent(astInner, "", "\t"))
				mylog.Check(os.WriteFile(cacheFilePath(inputHeader), jb, 0o644))

				astInner = nil
				jb = nil
				runtime.GC()
			}
			log.Printf("Clang worker: exiting")
		}()
	}
	for _, inputHeader := range includeFiles {
		// Check if there is a matching cache hit
		cacheFile := cacheFilePath(inputHeader)
		if _, e := os.Stat(cacheFile); e != nil && os.IsNotExist(e) {
			// Nonexistent cache file, regenerate from clang
			log.Printf("No AST cache for file %q, running clang...", filepath.Base(inputHeader))
			clangChan <- inputHeader
		}
	}
	// Done with all clang workers
	close(clangChan)
	clangWg.Wait()
}
