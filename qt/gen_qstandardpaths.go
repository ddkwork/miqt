package qt

import (
	"unsafe"
)

type QStandardPaths__StandardLocation int

const (
	QStandardPaths__DesktopLocation       QStandardPaths__StandardLocation = 0
	QStandardPaths__DocumentsLocation     QStandardPaths__StandardLocation = 1
	QStandardPaths__FontsLocation         QStandardPaths__StandardLocation = 2
	QStandardPaths__ApplicationsLocation  QStandardPaths__StandardLocation = 3
	QStandardPaths__MusicLocation         QStandardPaths__StandardLocation = 4
	QStandardPaths__MoviesLocation        QStandardPaths__StandardLocation = 5
	QStandardPaths__PicturesLocation      QStandardPaths__StandardLocation = 6
	QStandardPaths__TempLocation          QStandardPaths__StandardLocation = 7
	QStandardPaths__HomeLocation          QStandardPaths__StandardLocation = 8
	QStandardPaths__AppLocalDataLocation  QStandardPaths__StandardLocation = 9
	QStandardPaths__CacheLocation         QStandardPaths__StandardLocation = 10
	QStandardPaths__GenericDataLocation   QStandardPaths__StandardLocation = 11
	QStandardPaths__RuntimeLocation       QStandardPaths__StandardLocation = 12
	QStandardPaths__ConfigLocation        QStandardPaths__StandardLocation = 13
	QStandardPaths__DownloadLocation      QStandardPaths__StandardLocation = 14
	QStandardPaths__GenericCacheLocation  QStandardPaths__StandardLocation = 15
	QStandardPaths__GenericConfigLocation QStandardPaths__StandardLocation = 16
	QStandardPaths__AppDataLocation       QStandardPaths__StandardLocation = 17
	QStandardPaths__AppConfigLocation     QStandardPaths__StandardLocation = 18
	QStandardPaths__PublicShareLocation   QStandardPaths__StandardLocation = 19
	QStandardPaths__TemplatesLocation     QStandardPaths__StandardLocation = 20
	QStandardPaths__StateLocation         QStandardPaths__StandardLocation = 21
	QStandardPaths__GenericStateLocation  QStandardPaths__StandardLocation = 22
)

type QStandardPaths__LocateOption int

const (
	QStandardPaths__LocateFile      QStandardPaths__LocateOption = 0
	QStandardPaths__LocateDirectory QStandardPaths__LocateOption = 1
)

type QStandardPaths struct {
	h          uintptr
	isSubclass bool
}

func QStandardPaths_WritableLocation(typeVal StandardLocation) string {
	var _ms struct_miqt_string = QStandardPaths_WritableLocation(typeVal)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStandardPaths_StandardLocations(typeVal StandardLocation) []string {
	var _ma struct_miqt_array = QStandardPaths_StandardLocations(typeVal)
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

func QStandardPaths_Locate(typeVal StandardLocation, fileName string) string {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	var _ms struct_miqt_string = QStandardPaths_Locate(typeVal, fileName_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStandardPaths_LocateAll(typeVal StandardLocation, fileName string) []string {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	var _ma struct_miqt_array = QStandardPaths_LocateAll(typeVal, fileName_ms)
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

func QStandardPaths_DisplayName(typeVal StandardLocation) string {
	var _ms struct_miqt_string = QStandardPaths_DisplayName(typeVal)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStandardPaths_FindExecutable(executableName string) string {
	executableName_ms := struct_miqt_string{}
	executableName_ms.data = CString(executableName)
	executableName_ms.len = size_t(len(executableName))
	defer free(unsafe.Pointer(executableName_ms.data))
	var _ms struct_miqt_string = QStandardPaths_FindExecutable(executableName_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStandardPaths_SetTestModeEnabled(testMode bool) {
	QStandardPaths_SetTestModeEnabled((bool)(testMode))
}

func QStandardPaths_IsTestModeEnabled() bool {
	return (bool)(QStandardPaths_IsTestModeEnabled())
}

func QStandardPaths_Locate3(typeVal StandardLocation, fileName string, options LocateOptions) string {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	var _ms struct_miqt_string = QStandardPaths_Locate3(typeVal, fileName_ms, options)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStandardPaths_LocateAll3(typeVal StandardLocation, fileName string, options LocateOptions) []string {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	var _ma struct_miqt_array = QStandardPaths_LocateAll3(typeVal, fileName_ms, options)
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

func QStandardPaths_FindExecutable2(executableName string, paths []string) string {
	executableName_ms := struct_miqt_string{}
	executableName_ms.data = CString(executableName)
	executableName_ms.len = size_t(len(executableName))
	defer free(unsafe.Pointer(executableName_ms.data))
	paths_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(paths))))
	defer free(unsafe.Pointer(paths_CArray))
	for i := range paths {
		paths_i_ms := struct_miqt_string{}
		paths_i_ms.data = CString(paths[i])
		paths_i_ms.len = size_t(len(paths[i]))
		defer free(unsafe.Pointer(paths_i_ms.data))
		paths_CArray[i] = paths_i_ms
	}
	paths_ma := struct_miqt_array{len: size_t(len(paths)), data: unsafe.Pointer(paths_CArray)}
	var _ms struct_miqt_string = QStandardPaths_FindExecutable2(executableName_ms, paths_ma)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
