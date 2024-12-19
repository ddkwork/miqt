package main

import (
	"os"
	"testing"
)

func TestName(t *testing.T) {
	os.Setenv("PKG_CONFIG_PATH", "D:\\qt6\\qt_static\\lib\\pkgconfig")
	main()
}
