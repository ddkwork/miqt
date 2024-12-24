package main

import (
	"io/fs"
	"path/filepath"
	"testing"
)

func TestName(t *testing.T) {
	main()
}

func TestAddCMAKE_PREFIX_PATH(t *testing.T) {
	// project(dombookmarks LANGUAGES CXX)
	// set(CMAKE_PREFIX_PATH "D:/qt6/qt_static/lib/cmake")
	filepath.Walk("Qt-6.8.1", func(path string, info fs.FileInfo, err error) error {
		if filepath.Base(path) == "CMakeLists.txt" {
			println(path)
		}
		return err
	})
}
