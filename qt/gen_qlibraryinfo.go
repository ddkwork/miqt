package qt

import (
	"unsafe"
)

type QLibraryInfo__LibraryPath int

const (
	QLibraryInfo__PrefixPath             QLibraryInfo__LibraryPath = 0
	QLibraryInfo__DocumentationPath      QLibraryInfo__LibraryPath = 1
	QLibraryInfo__HeadersPath            QLibraryInfo__LibraryPath = 2
	QLibraryInfo__LibrariesPath          QLibraryInfo__LibraryPath = 3
	QLibraryInfo__LibraryExecutablesPath QLibraryInfo__LibraryPath = 4
	QLibraryInfo__BinariesPath           QLibraryInfo__LibraryPath = 5
	QLibraryInfo__PluginsPath            QLibraryInfo__LibraryPath = 6
	QLibraryInfo__QmlImportsPath         QLibraryInfo__LibraryPath = 7
	QLibraryInfo__Qml2ImportsPath        QLibraryInfo__LibraryPath = 7
	QLibraryInfo__ArchDataPath           QLibraryInfo__LibraryPath = 8
	QLibraryInfo__DataPath               QLibraryInfo__LibraryPath = 9
	QLibraryInfo__TranslationsPath       QLibraryInfo__LibraryPath = 10
	QLibraryInfo__ExamplesPath           QLibraryInfo__LibraryPath = 11
	QLibraryInfo__TestsPath              QLibraryInfo__LibraryPath = 12
	QLibraryInfo__SettingsPath           QLibraryInfo__LibraryPath = 100
)

type QLibraryInfo struct {
	h          uintptr
	isSubclass bool
}

func QLibraryInfo_Build() string {
	_ret := QLibraryInfo_Build()
	return GoString(_ret)
}

func QLibraryInfo_IsDebugBuild() bool {
	return (bool)(QLibraryInfo_IsDebugBuild())
}

func QLibraryInfo_IsSharedBuild() bool {
	return (bool)(QLibraryInfo_IsSharedBuild())
}

func QLibraryInfo_Version() *QVersionNumber {
	_goptr := newQVersionNumber(QLibraryInfo_Version())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QLibraryInfo_Path(p LibraryPath) string {
	var _ms struct_miqt_string = QLibraryInfo_Path(p)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QLibraryInfo_Paths(p LibraryPath) []string {
	var _ma struct_miqt_array = QLibraryInfo_Paths(p)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QLibraryInfo_Location(location LibraryLocation) string {
	var _ms struct_miqt_string = QLibraryInfo_Location(location)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QLibraryInfo_PlatformPluginArguments(platformName string) []string {
	platformName_ms := struct_miqt_string{}
	platformName_ms.data = CString(platformName)
	platformName_ms.len = size_t(len(platformName))
	defer free(unsafe.Pointer(platformName_ms.data))
	var _ma struct_miqt_array = QLibraryInfo_PlatformPluginArguments(platformName_ms)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}
