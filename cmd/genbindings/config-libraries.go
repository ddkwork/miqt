package main

import (
	"path/filepath"
	"strings"
)

func ProcessLibraries(outDir string) {
	AllowAllHeaders := func(string) bool { return true }
	flushKnownTypes()
	InsertTypedefs()
	generate(
		"qt",
		[]string{
			qt6Dir + "include\\QtCore",
			qt6Dir + "include\\QtGui",
			qt6Dir + "include\\QtWidgets",
		},
		func(fullpath string) bool {
			// Block cbor and generate it separately
			fname := filepath.Base(fullpath)
			if strings.HasPrefix(fname, "qcbor") {
				return false
			}
			if fname == "qcomparehelpers.h" {
				return false
			}
			return Widgets_AllowHeader(fullpath)
		},
		"--std=c++17 "+pkgConfigCflags("Qt6Widgets"),
		outDir,
		ClangMatchSameHeaderDefinitionOnly,
	)
	generate(
		"qt/cbor",
		[]string{
			qt6Dir + "include\\QtCore",
		},
		func(fullpath string) bool {
			// Only include the same json, xml, cbor files excluded above
			fname := filepath.Base(fullpath)
			return strings.HasPrefix(fname, "qcbor")
		},
		"--std=c++20 "+pkgConfigCflags("Qt6Core"),
		outDir,
		ClangMatchSameHeaderDefinitionOnly,
	)
	generate(
		"qt/printsupport",
		[]string{
			qt6Dir + "include\\QtPrintSupport",
		},
		AllowAllHeaders,
		"--std=c++17 "+pkgConfigCflags("Qt6PrintSupport"),
		outDir,
		ClangMatchSameHeaderDefinitionOnly,
	)
	generate(
		"qt/svg",
		[]string{
			qt6Dir + "include\\QtSvg",
			qt6Dir + "include\\QtSvgWidgets",
		},
		AllowAllHeaders,
		"--std=c++17 "+pkgConfigCflags("Qt6SvgWidgets"),
		outDir,
		ClangMatchSameHeaderDefinitionOnly,
	)
	generate(
		"qt/network",
		[]string{
			qt6Dir + "include\\QtNetwork",
		},
		func(fullpath string) bool {
			fname := filepath.Base(fullpath)
			return fname != "qtnetwork-config.h"
		},
		"--std=c++17 "+pkgConfigCflags("Qt6Network"),
		outDir,
		ClangMatchSameHeaderDefinitionOnly,
	)
	generate(
		"qt/multimedia",
		[]string{
			qt6Dir + "include\\QtMultimedia",
			qt6Dir + "include\\QtMultimediaWidgets",
		},
		AllowAllHeaders,
		"--std=c++17 "+pkgConfigCflags("Qt6MultimediaWidgets"),
		outDir,
		ClangMatchSameHeaderDefinitionOnly,
	)
	generate(
		"qt/spatialaudio", // Qt 6 Spatial Audio (on Debian this is a dependency of Qt6Multimedia)
		[]string{
			qt6Dir + "include\\QtSpatialAudio",
		},
		AllowAllHeaders,
		"--std=c++17 "+pkgConfigCflags("Qt6SpatialAudio"),
		outDir,
		ClangMatchSameHeaderDefinitionOnly,
	)
	generate(
		"qt/webchannel",
		[]string{
			qt6Dir + "include\\QtWebChannel",
		},
		AllowAllHeaders,
		"--std=c++17 "+pkgConfigCflags("Qt6WebChannel"), // todo bug Qt6Qml.pc not has
		outDir,
		ClangMatchSameHeaderDefinitionOnly,
	)
	//generate(
	//	"qt/webengine", //todo lib not found,need rebuild in qt6.9 repo
	//	[]string{
	//		qt6Dir + "include\\QtWebEngineCore",
	//		qt6Dir + "include\\QtWebEngineWidgets",
	//	},
	//	func(fullpath string) bool {
	//		baseName := filepath.Base(fullpath)
	//		if baseName == "qtwebenginewidgets-config.h" {
	//			return false
	//		}
	//		return true
	//	},
	//	"--std=c++17 "+pkgConfigCflags("Qt6WebEngineWidgets"),
	//	outDir,
	//	ClangMatchSameHeaderDefinitionOnly,
	//)
	//generate(//todo copy dockerfile
	//	"qt-restricted-extras/qscintilla6", // Depends on QtCore/Gui/Widgets, QPrintSupport
	//	[]string{
	//		qt6Dir + "include\\Qsci",
	//	},
	//	AllowAllHeaders,
	//	"--std=c++17 "+pkgConfigCflags("Qt6PrintSupport"),
	//	outDir,
	//	ClangMatchSameHeaderDefinitionOnly,
	//)
}
