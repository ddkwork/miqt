package main

import (
	"os"
	"testing"
)

func TestName(t *testing.T) {
	os.Setenv("PKG_CONFIG_PATH", qt6Dir+"lib\\pkgconfig")
	main()
}
