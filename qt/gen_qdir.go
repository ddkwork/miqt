package qt

import (
	"unsafe"
)

type QDir__Filter int

const (
	QDir__Dirs           QDir__Filter = 1
	QDir__Files          QDir__Filter = 2
	QDir__Drives         QDir__Filter = 4
	QDir__NoSymLinks     QDir__Filter = 8
	QDir__AllEntries     QDir__Filter = 7
	QDir__TypeMask       QDir__Filter = 15
	QDir__Readable       QDir__Filter = 16
	QDir__Writable       QDir__Filter = 32
	QDir__Executable     QDir__Filter = 64
	QDir__PermissionMask QDir__Filter = 112
	QDir__Modified       QDir__Filter = 128
	QDir__Hidden         QDir__Filter = 256
	QDir__System         QDir__Filter = 512
	QDir__AccessMask     QDir__Filter = 1008
	QDir__AllDirs        QDir__Filter = 1024
	QDir__CaseSensitive  QDir__Filter = 2048
	QDir__NoDot          QDir__Filter = 8192
	QDir__NoDotDot       QDir__Filter = 16384
	QDir__NoDotAndDotDot QDir__Filter = 24576
	QDir__NoFilter       QDir__Filter = -1
)

type QDir__SortFlag int

const (
	QDir__Name        QDir__SortFlag = 0
	QDir__Time        QDir__SortFlag = 1
	QDir__Size        QDir__SortFlag = 2
	QDir__Unsorted    QDir__SortFlag = 3
	QDir__SortByMask  QDir__SortFlag = 3
	QDir__DirsFirst   QDir__SortFlag = 4
	QDir__Reversed    QDir__SortFlag = 8
	QDir__IgnoreCase  QDir__SortFlag = 16
	QDir__DirsLast    QDir__SortFlag = 32
	QDir__LocaleAware QDir__SortFlag = 64
	QDir__Type        QDir__SortFlag = 128
	QDir__NoSort      QDir__SortFlag = -1
)

type QDir struct {
	h          uintptr
	isSubclass bool
}

// NewQDir constructs a new QDir object.
func NewQDir(param1 *QDir) *QDir {

	ret := newQDir(QDir_new(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDir2 constructs a new QDir object.
func NewQDir2() *QDir {

	ret := newQDir(QDir_new2())
	ret.isSubclass = true
	return ret
}

// NewQDir3 constructs a new QDir object.
func NewQDir3(path string, nameFilter string) *QDir {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	nameFilter_ms := struct_miqt_string{}
	nameFilter_ms.data = CString(nameFilter)
	nameFilter_ms.len = size_t(len(nameFilter))
	defer free(unsafe.Pointer(nameFilter_ms.data))

	ret := newQDir(QDir_new3(path_ms, nameFilter_ms))
	ret.isSubclass = true
	return ret
}

// NewQDir4 constructs a new QDir object.
func NewQDir4(path string) *QDir {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))

	ret := newQDir(QDir_new4(path_ms))
	ret.isSubclass = true
	return ret
}

// NewQDir5 constructs a new QDir object.
func NewQDir5(path string, nameFilter string, sort SortFlags) *QDir {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	nameFilter_ms := struct_miqt_string{}
	nameFilter_ms.data = CString(nameFilter)
	nameFilter_ms.len = size_t(len(nameFilter))
	defer free(unsafe.Pointer(nameFilter_ms.data))

	ret := newQDir(QDir_new5(path_ms, nameFilter_ms, sort))
	ret.isSubclass = true
	return ret
}

// NewQDir6 constructs a new QDir object.
func NewQDir6(path string, nameFilter string, sort SortFlags, filter Filters) *QDir {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	nameFilter_ms := struct_miqt_string{}
	nameFilter_ms.data = CString(nameFilter)
	nameFilter_ms.len = size_t(len(nameFilter))
	defer free(unsafe.Pointer(nameFilter_ms.data))

	ret := newQDir(QDir_new6(path_ms, nameFilter_ms, sort, filter))
	ret.isSubclass = true
	return ret
}

func (this *QDir) OperatorAssign(param1 *QDir) {
	QDir_OperatorAssign(this.h, param1.cPointer())
}

func (this *QDir) Swap(other *QDir) {
	QDir_Swap(this.h, other.cPointer())
}

func (this *QDir) SetPath(path string) {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	QDir_SetPath(this.h, path_ms)
}

func (this *QDir) Path() string {
	var _ms struct_miqt_string = QDir_Path(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDir) AbsolutePath() string {
	var _ms struct_miqt_string = QDir_AbsolutePath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDir) CanonicalPath() string {
	var _ms struct_miqt_string = QDir_CanonicalPath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDir_SetSearchPaths(prefix string, searchPaths []string) {
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	searchPaths_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(searchPaths))))
	defer free(unsafe.Pointer(searchPaths_CArray))
	for i := range searchPaths {
		searchPaths_i_ms := struct_miqt_string{}
		searchPaths_i_ms.data = CString(searchPaths[i])
		searchPaths_i_ms.len = size_t(len(searchPaths[i]))
		defer free(unsafe.Pointer(searchPaths_i_ms.data))
		searchPaths_CArray[i] = searchPaths_i_ms
	}
	searchPaths_ma := struct_miqt_array{len: size_t(len(searchPaths)), data: unsafe.Pointer(searchPaths_CArray)}
	QDir_SetSearchPaths(prefix_ms, searchPaths_ma)
}

func QDir_AddSearchPath(prefix string, path string) {
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	QDir_AddSearchPath(prefix_ms, path_ms)
}

func QDir_SearchPaths(prefix string) []string {
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	var _ma struct_miqt_array = QDir_SearchPaths(prefix_ms)
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

func (this *QDir) DirName() string {
	var _ms struct_miqt_string = QDir_DirName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDir) FilePath(fileName string) string {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	var _ms struct_miqt_string = QDir_FilePath(this.h, fileName_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDir) AbsoluteFilePath(fileName string) string {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	var _ms struct_miqt_string = QDir_AbsoluteFilePath(this.h, fileName_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDir) RelativeFilePath(fileName string) string {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	var _ms struct_miqt_string = QDir_RelativeFilePath(this.h, fileName_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDir_ToNativeSeparators(pathName string) string {
	pathName_ms := struct_miqt_string{}
	pathName_ms.data = CString(pathName)
	pathName_ms.len = size_t(len(pathName))
	defer free(unsafe.Pointer(pathName_ms.data))
	var _ms struct_miqt_string = QDir_ToNativeSeparators(pathName_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDir_FromNativeSeparators(pathName string) string {
	pathName_ms := struct_miqt_string{}
	pathName_ms.data = CString(pathName)
	pathName_ms.len = size_t(len(pathName))
	defer free(unsafe.Pointer(pathName_ms.data))
	var _ms struct_miqt_string = QDir_FromNativeSeparators(pathName_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDir) Cd(dirName string) bool {
	dirName_ms := struct_miqt_string{}
	dirName_ms.data = CString(dirName)
	dirName_ms.len = size_t(len(dirName))
	defer free(unsafe.Pointer(dirName_ms.data))
	return (bool)(QDir_Cd(this.h, dirName_ms))
}

func (this *QDir) CdUp() bool {
	return (bool)(QDir_CdUp(this.h))
}

func (this *QDir) NameFilters() []string {
	var _ma struct_miqt_array = QDir_NameFilters(this.h)
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

func (this *QDir) SetNameFilters(nameFilters []string) {
	nameFilters_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(nameFilters))))
	defer free(unsafe.Pointer(nameFilters_CArray))
	for i := range nameFilters {
		nameFilters_i_ms := struct_miqt_string{}
		nameFilters_i_ms.data = CString(nameFilters[i])
		nameFilters_i_ms.len = size_t(len(nameFilters[i]))
		defer free(unsafe.Pointer(nameFilters_i_ms.data))
		nameFilters_CArray[i] = nameFilters_i_ms
	}
	nameFilters_ma := struct_miqt_array{len: size_t(len(nameFilters)), data: unsafe.Pointer(nameFilters_CArray)}
	QDir_SetNameFilters(this.h, nameFilters_ma)
}

func (this *QDir) Filter() Filters {
	xxxxxxxxx
}

func (this *QDir) SetFilter(filter Filters) {
	QDir_SetFilter(this.h, filter)
}

func (this *QDir) Sorting() SortFlags {
	xxxxxxxxx
}

func (this *QDir) SetSorting(sort SortFlags) {
	QDir_SetSorting(this.h, sort)
}

func (this *QDir) Count() int64 {
	return (int64)(QDir_Count(this.h))
}

func (this *QDir) IsEmpty() bool {
	return (bool)(QDir_IsEmpty(this.h))
}

func (this *QDir) OperatorSubscript(param1 int64) string {
	var _ms struct_miqt_string = QDir_OperatorSubscript(this.h, (ptrdiff_t)(param1))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDir_NameFiltersFromString(nameFilter string) []string {
	nameFilter_ms := struct_miqt_string{}
	nameFilter_ms.data = CString(nameFilter)
	nameFilter_ms.len = size_t(len(nameFilter))
	defer free(unsafe.Pointer(nameFilter_ms.data))
	var _ma struct_miqt_array = QDir_NameFiltersFromString(nameFilter_ms)
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

func (this *QDir) EntryList() []string {
	var _ma struct_miqt_array = QDir_EntryList(this.h)
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

func (this *QDir) EntryListWithNameFilters(nameFilters []string) []string {
	nameFilters_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(nameFilters))))
	defer free(unsafe.Pointer(nameFilters_CArray))
	for i := range nameFilters {
		nameFilters_i_ms := struct_miqt_string{}
		nameFilters_i_ms.data = CString(nameFilters[i])
		nameFilters_i_ms.len = size_t(len(nameFilters[i]))
		defer free(unsafe.Pointer(nameFilters_i_ms.data))
		nameFilters_CArray[i] = nameFilters_i_ms
	}
	nameFilters_ma := struct_miqt_array{len: size_t(len(nameFilters)), data: unsafe.Pointer(nameFilters_CArray)}
	var _ma struct_miqt_array = QDir_EntryListWithNameFilters(this.h, nameFilters_ma)
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

func (this *QDir) EntryInfoList() []QFileInfo {
	var _ma struct_miqt_array = QDir_EntryInfoList(this.h)
	_ret := make([]QFileInfo, int(_ma.len))
	_outCast := (*[0xffff]*QFileInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQFileInfo(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDir) EntryInfoListWithNameFilters(nameFilters []string) []QFileInfo {
	nameFilters_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(nameFilters))))
	defer free(unsafe.Pointer(nameFilters_CArray))
	for i := range nameFilters {
		nameFilters_i_ms := struct_miqt_string{}
		nameFilters_i_ms.data = CString(nameFilters[i])
		nameFilters_i_ms.len = size_t(len(nameFilters[i]))
		defer free(unsafe.Pointer(nameFilters_i_ms.data))
		nameFilters_CArray[i] = nameFilters_i_ms
	}
	nameFilters_ma := struct_miqt_array{len: size_t(len(nameFilters)), data: unsafe.Pointer(nameFilters_CArray)}
	var _ma struct_miqt_array = QDir_EntryInfoListWithNameFilters(this.h, nameFilters_ma)
	_ret := make([]QFileInfo, int(_ma.len))
	_outCast := (*[0xffff]*QFileInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQFileInfo(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDir) Mkdir(dirName string) bool {
	dirName_ms := struct_miqt_string{}
	dirName_ms.data = CString(dirName)
	dirName_ms.len = size_t(len(dirName))
	defer free(unsafe.Pointer(dirName_ms.data))
	return (bool)(QDir_Mkdir(this.h, dirName_ms))
}

func (this *QDir) Mkdir2(dirName string, permissions Permission) bool {
	dirName_ms := struct_miqt_string{}
	dirName_ms.data = CString(dirName)
	dirName_ms.len = size_t(len(dirName))
	defer free(unsafe.Pointer(dirName_ms.data))
	return (bool)(QDir_Mkdir2(this.h, dirName_ms, (int)(permissions)))
}

func (this *QDir) Rmdir(dirName string) bool {
	dirName_ms := struct_miqt_string{}
	dirName_ms.data = CString(dirName)
	dirName_ms.len = size_t(len(dirName))
	defer free(unsafe.Pointer(dirName_ms.data))
	return (bool)(QDir_Rmdir(this.h, dirName_ms))
}

func (this *QDir) Mkpath(dirPath string) bool {
	dirPath_ms := struct_miqt_string{}
	dirPath_ms.data = CString(dirPath)
	dirPath_ms.len = size_t(len(dirPath))
	defer free(unsafe.Pointer(dirPath_ms.data))
	return (bool)(QDir_Mkpath(this.h, dirPath_ms))
}

func (this *QDir) Rmpath(dirPath string) bool {
	dirPath_ms := struct_miqt_string{}
	dirPath_ms.data = CString(dirPath)
	dirPath_ms.len = size_t(len(dirPath))
	defer free(unsafe.Pointer(dirPath_ms.data))
	return (bool)(QDir_Rmpath(this.h, dirPath_ms))
}

func (this *QDir) RemoveRecursively() bool {
	return (bool)(QDir_RemoveRecursively(this.h))
}

func (this *QDir) IsReadable() bool {
	return (bool)(QDir_IsReadable(this.h))
}

func (this *QDir) Exists() bool {
	return (bool)(QDir_Exists(this.h))
}

func (this *QDir) IsRoot() bool {
	return (bool)(QDir_IsRoot(this.h))
}

func QDir_IsRelativePath(path string) bool {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	return (bool)(QDir_IsRelativePath(path_ms))
}

func QDir_IsAbsolutePath(path string) bool {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	return (bool)(QDir_IsAbsolutePath(path_ms))
}

func (this *QDir) IsRelative() bool {
	return (bool)(QDir_IsRelative(this.h))
}

func (this *QDir) IsAbsolute() bool {
	return (bool)(QDir_IsAbsolute(this.h))
}

func (this *QDir) MakeAbsolute() bool {
	return (bool)(QDir_MakeAbsolute(this.h))
}

func (this *QDir) Remove(fileName string) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	return (bool)(QDir_Remove(this.h, fileName_ms))
}

func (this *QDir) Rename(oldName string, newName string) bool {
	oldName_ms := struct_miqt_string{}
	oldName_ms.data = CString(oldName)
	oldName_ms.len = size_t(len(oldName))
	defer free(unsafe.Pointer(oldName_ms.data))
	newName_ms := struct_miqt_string{}
	newName_ms.data = CString(newName)
	newName_ms.len = size_t(len(newName))
	defer free(unsafe.Pointer(newName_ms.data))
	return (bool)(QDir_Rename(this.h, oldName_ms, newName_ms))
}

func (this *QDir) ExistsWithName(name string) bool {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	return (bool)(QDir_ExistsWithName(this.h, name_ms))
}

func QDir_Drives() []QFileInfo {
	var _ma struct_miqt_array = QDir_Drives()
	_ret := make([]QFileInfo, int(_ma.len))
	_outCast := (*[0xffff]*QFileInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQFileInfo(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QDir_ListSeparator() *QChar {
	_goptr := newQChar(QDir_ListSeparator())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDir_Separator() *QChar {
	_goptr := newQChar(QDir_Separator())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDir_SetCurrent(path string) bool {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	return (bool)(QDir_SetCurrent(path_ms))
}

func QDir_Current() *QDir {
	_goptr := newQDir(QDir_Current())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDir_CurrentPath() string {
	var _ms struct_miqt_string = QDir_CurrentPath()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDir_Home() *QDir {
	_goptr := newQDir(QDir_Home())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDir_HomePath() string {
	var _ms struct_miqt_string = QDir_HomePath()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDir_Root() *QDir {
	_goptr := newQDir(QDir_Root())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDir_RootPath() string {
	var _ms struct_miqt_string = QDir_RootPath()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDir_Temp() *QDir {
	_goptr := newQDir(QDir_Temp())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDir_TempPath() string {
	var _ms struct_miqt_string = QDir_TempPath()
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDir_Match(filters []string, fileName string) bool {
	filters_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(filters))))
	defer free(unsafe.Pointer(filters_CArray))
	for i := range filters {
		filters_i_ms := struct_miqt_string{}
		filters_i_ms.data = CString(filters[i])
		filters_i_ms.len = size_t(len(filters[i]))
		defer free(unsafe.Pointer(filters_i_ms.data))
		filters_CArray[i] = filters_i_ms
	}
	filters_ma := struct_miqt_array{len: size_t(len(filters)), data: unsafe.Pointer(filters_CArray)}
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	return (bool)(QDir_Match(filters_ma, fileName_ms))
}

func QDir_Match2(filter string, fileName string) bool {
	filter_ms := struct_miqt_string{}
	filter_ms.data = CString(filter)
	filter_ms.len = size_t(len(filter))
	defer free(unsafe.Pointer(filter_ms.data))
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	return (bool)(QDir_Match2(filter_ms, fileName_ms))
}

func QDir_CleanPath(path string) string {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	var _ms struct_miqt_string = QDir_CleanPath(path_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDir) Refresh() {
	QDir_Refresh(this.h)
}

func (this *QDir) Count1(param1 Disambiguated_t) int64 {
	return (int64)(QDir_Count1(this.h, param1.cPointer()))
}

func (this *QDir) IsEmpty1(filters Filters) bool {
	return (bool)(QDir_IsEmpty1(this.h, filters))
}

func (this *QDir) EntryList1(filters Filters) []string {
	var _ma struct_miqt_array = QDir_EntryList1(this.h, filters)
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

func (this *QDir) EntryList2(filters Filters, sort SortFlags) []string {
	var _ma struct_miqt_array = QDir_EntryList2(this.h, filters, sort)
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

func (this *QDir) EntryList22(nameFilters []string, filters Filters) []string {
	nameFilters_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(nameFilters))))
	defer free(unsafe.Pointer(nameFilters_CArray))
	for i := range nameFilters {
		nameFilters_i_ms := struct_miqt_string{}
		nameFilters_i_ms.data = CString(nameFilters[i])
		nameFilters_i_ms.len = size_t(len(nameFilters[i]))
		defer free(unsafe.Pointer(nameFilters_i_ms.data))
		nameFilters_CArray[i] = nameFilters_i_ms
	}
	nameFilters_ma := struct_miqt_array{len: size_t(len(nameFilters)), data: unsafe.Pointer(nameFilters_CArray)}
	var _ma struct_miqt_array = QDir_EntryList22(this.h, nameFilters_ma, filters)
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

func (this *QDir) EntryList3(nameFilters []string, filters Filters, sort SortFlags) []string {
	nameFilters_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(nameFilters))))
	defer free(unsafe.Pointer(nameFilters_CArray))
	for i := range nameFilters {
		nameFilters_i_ms := struct_miqt_string{}
		nameFilters_i_ms.data = CString(nameFilters[i])
		nameFilters_i_ms.len = size_t(len(nameFilters[i]))
		defer free(unsafe.Pointer(nameFilters_i_ms.data))
		nameFilters_CArray[i] = nameFilters_i_ms
	}
	nameFilters_ma := struct_miqt_array{len: size_t(len(nameFilters)), data: unsafe.Pointer(nameFilters_CArray)}
	var _ma struct_miqt_array = QDir_EntryList3(this.h, nameFilters_ma, filters, sort)
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

func (this *QDir) EntryInfoList1(filters Filters) []QFileInfo {
	var _ma struct_miqt_array = QDir_EntryInfoList1(this.h, filters)
	_ret := make([]QFileInfo, int(_ma.len))
	_outCast := (*[0xffff]*QFileInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQFileInfo(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDir) EntryInfoList2(filters Filters, sort SortFlags) []QFileInfo {
	var _ma struct_miqt_array = QDir_EntryInfoList2(this.h, filters, sort)
	_ret := make([]QFileInfo, int(_ma.len))
	_outCast := (*[0xffff]*QFileInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQFileInfo(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDir) EntryInfoList22(nameFilters []string, filters Filters) []QFileInfo {
	nameFilters_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(nameFilters))))
	defer free(unsafe.Pointer(nameFilters_CArray))
	for i := range nameFilters {
		nameFilters_i_ms := struct_miqt_string{}
		nameFilters_i_ms.data = CString(nameFilters[i])
		nameFilters_i_ms.len = size_t(len(nameFilters[i]))
		defer free(unsafe.Pointer(nameFilters_i_ms.data))
		nameFilters_CArray[i] = nameFilters_i_ms
	}
	nameFilters_ma := struct_miqt_array{len: size_t(len(nameFilters)), data: unsafe.Pointer(nameFilters_CArray)}
	var _ma struct_miqt_array = QDir_EntryInfoList22(this.h, nameFilters_ma, filters)
	_ret := make([]QFileInfo, int(_ma.len))
	_outCast := (*[0xffff]*QFileInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQFileInfo(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDir) EntryInfoList3(nameFilters []string, filters Filters, sort SortFlags) []QFileInfo {
	nameFilters_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(nameFilters))))
	defer free(unsafe.Pointer(nameFilters_CArray))
	for i := range nameFilters {
		nameFilters_i_ms := struct_miqt_string{}
		nameFilters_i_ms.data = CString(nameFilters[i])
		nameFilters_i_ms.len = size_t(len(nameFilters[i]))
		defer free(unsafe.Pointer(nameFilters_i_ms.data))
		nameFilters_CArray[i] = nameFilters_i_ms
	}
	nameFilters_ma := struct_miqt_array{len: size_t(len(nameFilters)), data: unsafe.Pointer(nameFilters_CArray)}
	var _ma struct_miqt_array = QDir_EntryInfoList3(this.h, nameFilters_ma, filters, sort)
	_ret := make([]QFileInfo, int(_ma.len))
	_outCast := (*[0xffff]*QFileInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQFileInfo(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}
