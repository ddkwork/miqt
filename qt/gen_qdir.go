package qt

/*

#include "gen_qdir.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QDir__Filter int

const (
	QDir__Filter__Dirs           QDir__Filter = 1
	QDir__Filter__Files          QDir__Filter = 2
	QDir__Filter__Drives         QDir__Filter = 4
	QDir__Filter__NoSymLinks     QDir__Filter = 8
	QDir__Filter__AllEntries     QDir__Filter = 7
	QDir__Filter__TypeMask       QDir__Filter = 15
	QDir__Filter__Readable       QDir__Filter = 16
	QDir__Filter__Writable       QDir__Filter = 32
	QDir__Filter__Executable     QDir__Filter = 64
	QDir__Filter__PermissionMask QDir__Filter = 112
	QDir__Filter__Modified       QDir__Filter = 128
	QDir__Filter__Hidden         QDir__Filter = 256
	QDir__Filter__System         QDir__Filter = 512
	QDir__Filter__AccessMask     QDir__Filter = 1008
	QDir__Filter__AllDirs        QDir__Filter = 1024
	QDir__Filter__CaseSensitive  QDir__Filter = 2048
	QDir__Filter__NoDot          QDir__Filter = 8192
	QDir__Filter__NoDotDot       QDir__Filter = 16384
	QDir__Filter__NoDotAndDotDot QDir__Filter = 24576
	QDir__Filter__NoFilter       QDir__Filter = -1
)

type QDir__SortFlag int

const (
	QDir__SortFlag__Name        QDir__SortFlag = 0
	QDir__SortFlag__Time        QDir__SortFlag = 1
	QDir__SortFlag__Size        QDir__SortFlag = 2
	QDir__SortFlag__Unsorted    QDir__SortFlag = 3
	QDir__SortFlag__SortByMask  QDir__SortFlag = 3
	QDir__SortFlag__DirsFirst   QDir__SortFlag = 4
	QDir__SortFlag__Reversed    QDir__SortFlag = 8
	QDir__SortFlag__IgnoreCase  QDir__SortFlag = 16
	QDir__SortFlag__DirsLast    QDir__SortFlag = 32
	QDir__SortFlag__LocaleAware QDir__SortFlag = 64
	QDir__SortFlag__Type        QDir__SortFlag = 128
	QDir__SortFlag__NoSort      QDir__SortFlag = -1
)

type QDir struct {
	h *C.QDir
}

func (this *QDir) cPointer() *C.QDir {
	if this == nil {
		return nil
	}
	return this.h
}

func newQDir(h *C.QDir) *QDir {
	if h == nil {
		return nil
	}
	return &QDir{h: h}
}

func newQDir_U(h unsafe.Pointer) *QDir {
	return newQDir((*C.QDir)(h))
}

// NewQDir constructs a new QDir object.
func NewQDir(param1 *QDir) *QDir {
	ret := C.QDir_new(param1.cPointer())
	return newQDir(ret)
}

// NewQDir2 constructs a new QDir object.
func NewQDir2() *QDir {
	ret := C.QDir_new2()
	return newQDir(ret)
}

// NewQDir3 constructs a new QDir object.
func NewQDir3(path string, nameFilter string) *QDir {
	path_Cstring := C.CString(path)
	defer C.free(unsafe.Pointer(path_Cstring))
	nameFilter_Cstring := C.CString(nameFilter)
	defer C.free(unsafe.Pointer(nameFilter_Cstring))
	ret := C.QDir_new3(path_Cstring, C.size_t(len(path)), nameFilter_Cstring, C.size_t(len(nameFilter)))
	return newQDir(ret)
}

// NewQDir4 constructs a new QDir object.
func NewQDir4(path string) *QDir {
	path_Cstring := C.CString(path)
	defer C.free(unsafe.Pointer(path_Cstring))
	ret := C.QDir_new4(path_Cstring, C.size_t(len(path)))
	return newQDir(ret)
}

// NewQDir5 constructs a new QDir object.
func NewQDir5(path string, nameFilter string, sort int) *QDir {
	path_Cstring := C.CString(path)
	defer C.free(unsafe.Pointer(path_Cstring))
	nameFilter_Cstring := C.CString(nameFilter)
	defer C.free(unsafe.Pointer(nameFilter_Cstring))
	ret := C.QDir_new5(path_Cstring, C.size_t(len(path)), nameFilter_Cstring, C.size_t(len(nameFilter)), (C.int)(sort))
	return newQDir(ret)
}

// NewQDir6 constructs a new QDir object.
func NewQDir6(path string, nameFilter string, sort int, filter int) *QDir {
	path_Cstring := C.CString(path)
	defer C.free(unsafe.Pointer(path_Cstring))
	nameFilter_Cstring := C.CString(nameFilter)
	defer C.free(unsafe.Pointer(nameFilter_Cstring))
	ret := C.QDir_new6(path_Cstring, C.size_t(len(path)), nameFilter_Cstring, C.size_t(len(nameFilter)), (C.int)(sort), (C.int)(filter))
	return newQDir(ret)
}

func (this *QDir) OperatorAssign(param1 *QDir) {
	C.QDir_OperatorAssign(this.h, param1.cPointer())
}

func (this *QDir) OperatorAssignWithPath(path string) {
	path_Cstring := C.CString(path)
	defer C.free(unsafe.Pointer(path_Cstring))
	C.QDir_OperatorAssignWithPath(this.h, path_Cstring, C.size_t(len(path)))
}

func (this *QDir) Swap(other *QDir) {
	C.QDir_Swap(this.h, other.cPointer())
}

func (this *QDir) SetPath(path string) {
	path_Cstring := C.CString(path)
	defer C.free(unsafe.Pointer(path_Cstring))
	C.QDir_SetPath(this.h, path_Cstring, C.size_t(len(path)))
}

func (this *QDir) Path() string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QDir_Path(this.h, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) AbsolutePath() string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QDir_AbsolutePath(this.h, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) CanonicalPath() string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QDir_CanonicalPath(this.h, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QDir_AddResourceSearchPath(path string) {
	path_Cstring := C.CString(path)
	defer C.free(unsafe.Pointer(path_Cstring))
	C.QDir_AddResourceSearchPath(path_Cstring, C.size_t(len(path)))
}

func QDir_SetSearchPaths(prefix string, searchPaths []string) {
	prefix_Cstring := C.CString(prefix)
	defer C.free(unsafe.Pointer(prefix_Cstring))
	// For the C ABI, malloc two C arrays; raw char* pointers and their lengths
	searchPaths_CArray := (*[0xffff]*C.char)(C.malloc(C.size_t(8 * len(searchPaths))))
	searchPaths_Lengths := (*[0xffff]C.uint64_t)(C.malloc(C.size_t(8 * len(searchPaths))))
	defer C.free(unsafe.Pointer(searchPaths_CArray))
	defer C.free(unsafe.Pointer(searchPaths_Lengths))
	for i := range searchPaths {
		single_cstring := C.CString(searchPaths[i])
		defer C.free(unsafe.Pointer(single_cstring))
		searchPaths_CArray[i] = single_cstring
		searchPaths_Lengths[i] = (C.uint64_t)(len(searchPaths[i]))
	}
	C.QDir_SetSearchPaths(prefix_Cstring, C.size_t(len(prefix)), &searchPaths_CArray[0], &searchPaths_Lengths[0], C.size_t(len(searchPaths)))
}

func QDir_AddSearchPath(prefix string, path string) {
	prefix_Cstring := C.CString(prefix)
	defer C.free(unsafe.Pointer(prefix_Cstring))
	path_Cstring := C.CString(path)
	defer C.free(unsafe.Pointer(path_Cstring))
	C.QDir_AddSearchPath(prefix_Cstring, C.size_t(len(prefix)), path_Cstring, C.size_t(len(path)))
}

func QDir_SearchPaths(prefix string) []string {
	prefix_Cstring := C.CString(prefix)
	defer C.free(unsafe.Pointer(prefix_Cstring))
	var _out **C.char = nil
	var _out_Lengths *C.int = nil
	var _out_len C.size_t = 0
	C.QDir_SearchPaths(prefix_Cstring, C.size_t(len(prefix)), &_out, &_out_Lengths, &_out_len)
	ret := make([]string, int(_out_len))
	_outCast := (*[0xffff]*C.char)(unsafe.Pointer(_out)) // hey ya
	_out_LengthsCast := (*[0xffff]C.int)(unsafe.Pointer(_out_Lengths))
	for i := 0; i < int(_out_len); i++ {
		ret[i] = C.GoStringN(_outCast[i], _out_LengthsCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) DirName() string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QDir_DirName(this.h, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) FilePath(fileName string) string {
	fileName_Cstring := C.CString(fileName)
	defer C.free(unsafe.Pointer(fileName_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QDir_FilePath(this.h, fileName_Cstring, C.size_t(len(fileName)), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) AbsoluteFilePath(fileName string) string {
	fileName_Cstring := C.CString(fileName)
	defer C.free(unsafe.Pointer(fileName_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QDir_AbsoluteFilePath(this.h, fileName_Cstring, C.size_t(len(fileName)), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) RelativeFilePath(fileName string) string {
	fileName_Cstring := C.CString(fileName)
	defer C.free(unsafe.Pointer(fileName_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QDir_RelativeFilePath(this.h, fileName_Cstring, C.size_t(len(fileName)), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QDir_ToNativeSeparators(pathName string) string {
	pathName_Cstring := C.CString(pathName)
	defer C.free(unsafe.Pointer(pathName_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QDir_ToNativeSeparators(pathName_Cstring, C.size_t(len(pathName)), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QDir_FromNativeSeparators(pathName string) string {
	pathName_Cstring := C.CString(pathName)
	defer C.free(unsafe.Pointer(pathName_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QDir_FromNativeSeparators(pathName_Cstring, C.size_t(len(pathName)), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) Cd(dirName string) bool {
	dirName_Cstring := C.CString(dirName)
	defer C.free(unsafe.Pointer(dirName_Cstring))
	ret := C.QDir_Cd(this.h, dirName_Cstring, C.size_t(len(dirName)))
	return (bool)(ret)
}

func (this *QDir) CdUp() bool {
	ret := C.QDir_CdUp(this.h)
	return (bool)(ret)
}

func (this *QDir) NameFilters() []string {
	var _out **C.char = nil
	var _out_Lengths *C.int = nil
	var _out_len C.size_t = 0
	C.QDir_NameFilters(this.h, &_out, &_out_Lengths, &_out_len)
	ret := make([]string, int(_out_len))
	_outCast := (*[0xffff]*C.char)(unsafe.Pointer(_out)) // hey ya
	_out_LengthsCast := (*[0xffff]C.int)(unsafe.Pointer(_out_Lengths))
	for i := 0; i < int(_out_len); i++ {
		ret[i] = C.GoStringN(_outCast[i], _out_LengthsCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) SetNameFilters(nameFilters []string) {
	// For the C ABI, malloc two C arrays; raw char* pointers and their lengths
	nameFilters_CArray := (*[0xffff]*C.char)(C.malloc(C.size_t(8 * len(nameFilters))))
	nameFilters_Lengths := (*[0xffff]C.uint64_t)(C.malloc(C.size_t(8 * len(nameFilters))))
	defer C.free(unsafe.Pointer(nameFilters_CArray))
	defer C.free(unsafe.Pointer(nameFilters_Lengths))
	for i := range nameFilters {
		single_cstring := C.CString(nameFilters[i])
		defer C.free(unsafe.Pointer(single_cstring))
		nameFilters_CArray[i] = single_cstring
		nameFilters_Lengths[i] = (C.uint64_t)(len(nameFilters[i]))
	}
	C.QDir_SetNameFilters(this.h, &nameFilters_CArray[0], &nameFilters_Lengths[0], C.size_t(len(nameFilters)))
}

func (this *QDir) Filter() int {
	ret := C.QDir_Filter(this.h)
	return (int)(ret)
}

func (this *QDir) SetFilter(filter int) {
	C.QDir_SetFilter(this.h, (C.int)(filter))
}

func (this *QDir) Sorting() int {
	ret := C.QDir_Sorting(this.h)
	return (int)(ret)
}

func (this *QDir) SetSorting(sort int) {
	C.QDir_SetSorting(this.h, (C.int)(sort))
}

func (this *QDir) Count() uint {
	ret := C.QDir_Count(this.h)
	return (uint)(ret)
}

func (this *QDir) IsEmpty() bool {
	ret := C.QDir_IsEmpty(this.h)
	return (bool)(ret)
}

func (this *QDir) OperatorSubscript(param1 int) string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QDir_OperatorSubscript(this.h, (C.int)(param1), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QDir_NameFiltersFromString(nameFilter string) []string {
	nameFilter_Cstring := C.CString(nameFilter)
	defer C.free(unsafe.Pointer(nameFilter_Cstring))
	var _out **C.char = nil
	var _out_Lengths *C.int = nil
	var _out_len C.size_t = 0
	C.QDir_NameFiltersFromString(nameFilter_Cstring, C.size_t(len(nameFilter)), &_out, &_out_Lengths, &_out_len)
	ret := make([]string, int(_out_len))
	_outCast := (*[0xffff]*C.char)(unsafe.Pointer(_out)) // hey ya
	_out_LengthsCast := (*[0xffff]C.int)(unsafe.Pointer(_out_Lengths))
	for i := 0; i < int(_out_len); i++ {
		ret[i] = C.GoStringN(_outCast[i], _out_LengthsCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) EntryList() []string {
	var _out **C.char = nil
	var _out_Lengths *C.int = nil
	var _out_len C.size_t = 0
	C.QDir_EntryList(this.h, &_out, &_out_Lengths, &_out_len)
	ret := make([]string, int(_out_len))
	_outCast := (*[0xffff]*C.char)(unsafe.Pointer(_out)) // hey ya
	_out_LengthsCast := (*[0xffff]C.int)(unsafe.Pointer(_out_Lengths))
	for i := 0; i < int(_out_len); i++ {
		ret[i] = C.GoStringN(_outCast[i], _out_LengthsCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) EntryListWithNameFilters(nameFilters []string) []string {
	// For the C ABI, malloc two C arrays; raw char* pointers and their lengths
	nameFilters_CArray := (*[0xffff]*C.char)(C.malloc(C.size_t(8 * len(nameFilters))))
	nameFilters_Lengths := (*[0xffff]C.uint64_t)(C.malloc(C.size_t(8 * len(nameFilters))))
	defer C.free(unsafe.Pointer(nameFilters_CArray))
	defer C.free(unsafe.Pointer(nameFilters_Lengths))
	for i := range nameFilters {
		single_cstring := C.CString(nameFilters[i])
		defer C.free(unsafe.Pointer(single_cstring))
		nameFilters_CArray[i] = single_cstring
		nameFilters_Lengths[i] = (C.uint64_t)(len(nameFilters[i]))
	}
	var _out **C.char = nil
	var _out_Lengths *C.int = nil
	var _out_len C.size_t = 0
	C.QDir_EntryListWithNameFilters(this.h, &nameFilters_CArray[0], &nameFilters_Lengths[0], C.size_t(len(nameFilters)), &_out, &_out_Lengths, &_out_len)
	ret := make([]string, int(_out_len))
	_outCast := (*[0xffff]*C.char)(unsafe.Pointer(_out)) // hey ya
	_out_LengthsCast := (*[0xffff]C.int)(unsafe.Pointer(_out_Lengths))
	for i := 0; i < int(_out_len); i++ {
		ret[i] = C.GoStringN(_outCast[i], _out_LengthsCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) EntryInfoList() []QFileInfo {
	var _out **C.QFileInfo = nil
	var _out_len C.size_t = 0
	C.QDir_EntryInfoList(this.h, &_out, &_out_len)
	ret := make([]QFileInfo, int(_out_len))
	_outCast := (*[0xffff]*C.QFileInfo)(unsafe.Pointer(_out)) // so fresh so clean
	for i := 0; i < int(_out_len); i++ {
		ret[i] = *newQFileInfo(_outCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) EntryInfoListWithNameFilters(nameFilters []string) []QFileInfo {
	// For the C ABI, malloc two C arrays; raw char* pointers and their lengths
	nameFilters_CArray := (*[0xffff]*C.char)(C.malloc(C.size_t(8 * len(nameFilters))))
	nameFilters_Lengths := (*[0xffff]C.uint64_t)(C.malloc(C.size_t(8 * len(nameFilters))))
	defer C.free(unsafe.Pointer(nameFilters_CArray))
	defer C.free(unsafe.Pointer(nameFilters_Lengths))
	for i := range nameFilters {
		single_cstring := C.CString(nameFilters[i])
		defer C.free(unsafe.Pointer(single_cstring))
		nameFilters_CArray[i] = single_cstring
		nameFilters_Lengths[i] = (C.uint64_t)(len(nameFilters[i]))
	}
	var _out **C.QFileInfo = nil
	var _out_len C.size_t = 0
	C.QDir_EntryInfoListWithNameFilters(this.h, &nameFilters_CArray[0], &nameFilters_Lengths[0], C.size_t(len(nameFilters)), &_out, &_out_len)
	ret := make([]QFileInfo, int(_out_len))
	_outCast := (*[0xffff]*C.QFileInfo)(unsafe.Pointer(_out)) // so fresh so clean
	for i := 0; i < int(_out_len); i++ {
		ret[i] = *newQFileInfo(_outCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) Mkdir(dirName string) bool {
	dirName_Cstring := C.CString(dirName)
	defer C.free(unsafe.Pointer(dirName_Cstring))
	ret := C.QDir_Mkdir(this.h, dirName_Cstring, C.size_t(len(dirName)))
	return (bool)(ret)
}

func (this *QDir) Rmdir(dirName string) bool {
	dirName_Cstring := C.CString(dirName)
	defer C.free(unsafe.Pointer(dirName_Cstring))
	ret := C.QDir_Rmdir(this.h, dirName_Cstring, C.size_t(len(dirName)))
	return (bool)(ret)
}

func (this *QDir) Mkpath(dirPath string) bool {
	dirPath_Cstring := C.CString(dirPath)
	defer C.free(unsafe.Pointer(dirPath_Cstring))
	ret := C.QDir_Mkpath(this.h, dirPath_Cstring, C.size_t(len(dirPath)))
	return (bool)(ret)
}

func (this *QDir) Rmpath(dirPath string) bool {
	dirPath_Cstring := C.CString(dirPath)
	defer C.free(unsafe.Pointer(dirPath_Cstring))
	ret := C.QDir_Rmpath(this.h, dirPath_Cstring, C.size_t(len(dirPath)))
	return (bool)(ret)
}

func (this *QDir) RemoveRecursively() bool {
	ret := C.QDir_RemoveRecursively(this.h)
	return (bool)(ret)
}

func (this *QDir) IsReadable() bool {
	ret := C.QDir_IsReadable(this.h)
	return (bool)(ret)
}

func (this *QDir) Exists() bool {
	ret := C.QDir_Exists(this.h)
	return (bool)(ret)
}

func (this *QDir) IsRoot() bool {
	ret := C.QDir_IsRoot(this.h)
	return (bool)(ret)
}

func QDir_IsRelativePath(path string) bool {
	path_Cstring := C.CString(path)
	defer C.free(unsafe.Pointer(path_Cstring))
	ret := C.QDir_IsRelativePath(path_Cstring, C.size_t(len(path)))
	return (bool)(ret)
}

func QDir_IsAbsolutePath(path string) bool {
	path_Cstring := C.CString(path)
	defer C.free(unsafe.Pointer(path_Cstring))
	ret := C.QDir_IsAbsolutePath(path_Cstring, C.size_t(len(path)))
	return (bool)(ret)
}

func (this *QDir) IsRelative() bool {
	ret := C.QDir_IsRelative(this.h)
	return (bool)(ret)
}

func (this *QDir) IsAbsolute() bool {
	ret := C.QDir_IsAbsolute(this.h)
	return (bool)(ret)
}

func (this *QDir) MakeAbsolute() bool {
	ret := C.QDir_MakeAbsolute(this.h)
	return (bool)(ret)
}

func (this *QDir) OperatorEqual(dir *QDir) bool {
	ret := C.QDir_OperatorEqual(this.h, dir.cPointer())
	return (bool)(ret)
}

func (this *QDir) OperatorNotEqual(dir *QDir) bool {
	ret := C.QDir_OperatorNotEqual(this.h, dir.cPointer())
	return (bool)(ret)
}

func (this *QDir) Remove(fileName string) bool {
	fileName_Cstring := C.CString(fileName)
	defer C.free(unsafe.Pointer(fileName_Cstring))
	ret := C.QDir_Remove(this.h, fileName_Cstring, C.size_t(len(fileName)))
	return (bool)(ret)
}

func (this *QDir) Rename(oldName string, newName string) bool {
	oldName_Cstring := C.CString(oldName)
	defer C.free(unsafe.Pointer(oldName_Cstring))
	newName_Cstring := C.CString(newName)
	defer C.free(unsafe.Pointer(newName_Cstring))
	ret := C.QDir_Rename(this.h, oldName_Cstring, C.size_t(len(oldName)), newName_Cstring, C.size_t(len(newName)))
	return (bool)(ret)
}

func (this *QDir) ExistsWithName(name string) bool {
	name_Cstring := C.CString(name)
	defer C.free(unsafe.Pointer(name_Cstring))
	ret := C.QDir_ExistsWithName(this.h, name_Cstring, C.size_t(len(name)))
	return (bool)(ret)
}

func QDir_Drives() []QFileInfo {
	var _out **C.QFileInfo = nil
	var _out_len C.size_t = 0
	C.QDir_Drives(&_out, &_out_len)
	ret := make([]QFileInfo, int(_out_len))
	_outCast := (*[0xffff]*C.QFileInfo)(unsafe.Pointer(_out)) // so fresh so clean
	for i := 0; i < int(_out_len); i++ {
		ret[i] = *newQFileInfo(_outCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func QDir_ListSeparator() *QChar {
	ret := C.QDir_ListSeparator()
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQChar(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QChar) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func QDir_Separator() *QChar {
	ret := C.QDir_Separator()
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQChar(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QChar) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func QDir_SetCurrent(path string) bool {
	path_Cstring := C.CString(path)
	defer C.free(unsafe.Pointer(path_Cstring))
	ret := C.QDir_SetCurrent(path_Cstring, C.size_t(len(path)))
	return (bool)(ret)
}

func QDir_Current() *QDir {
	ret := C.QDir_Current()
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQDir(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QDir) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func QDir_CurrentPath() string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QDir_CurrentPath(&_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QDir_Home() *QDir {
	ret := C.QDir_Home()
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQDir(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QDir) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func QDir_HomePath() string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QDir_HomePath(&_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QDir_Root() *QDir {
	ret := C.QDir_Root()
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQDir(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QDir) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func QDir_RootPath() string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QDir_RootPath(&_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QDir_Temp() *QDir {
	ret := C.QDir_Temp()
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQDir(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QDir) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func QDir_TempPath() string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QDir_TempPath(&_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QDir_Match(filters []string, fileName string) bool {
	// For the C ABI, malloc two C arrays; raw char* pointers and their lengths
	filters_CArray := (*[0xffff]*C.char)(C.malloc(C.size_t(8 * len(filters))))
	filters_Lengths := (*[0xffff]C.uint64_t)(C.malloc(C.size_t(8 * len(filters))))
	defer C.free(unsafe.Pointer(filters_CArray))
	defer C.free(unsafe.Pointer(filters_Lengths))
	for i := range filters {
		single_cstring := C.CString(filters[i])
		defer C.free(unsafe.Pointer(single_cstring))
		filters_CArray[i] = single_cstring
		filters_Lengths[i] = (C.uint64_t)(len(filters[i]))
	}
	fileName_Cstring := C.CString(fileName)
	defer C.free(unsafe.Pointer(fileName_Cstring))
	ret := C.QDir_Match(&filters_CArray[0], &filters_Lengths[0], C.size_t(len(filters)), fileName_Cstring, C.size_t(len(fileName)))
	return (bool)(ret)
}

func QDir_Match2(filter string, fileName string) bool {
	filter_Cstring := C.CString(filter)
	defer C.free(unsafe.Pointer(filter_Cstring))
	fileName_Cstring := C.CString(fileName)
	defer C.free(unsafe.Pointer(fileName_Cstring))
	ret := C.QDir_Match2(filter_Cstring, C.size_t(len(filter)), fileName_Cstring, C.size_t(len(fileName)))
	return (bool)(ret)
}

func QDir_CleanPath(path string) string {
	path_Cstring := C.CString(path)
	defer C.free(unsafe.Pointer(path_Cstring))
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QDir_CleanPath(path_Cstring, C.size_t(len(path)), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) Refresh() {
	C.QDir_Refresh(this.h)
}

func (this *QDir) IsEmpty1(filters int) bool {
	ret := C.QDir_IsEmpty1(this.h, (C.int)(filters))
	return (bool)(ret)
}

func (this *QDir) EntryList1(filters int) []string {
	var _out **C.char = nil
	var _out_Lengths *C.int = nil
	var _out_len C.size_t = 0
	C.QDir_EntryList1(this.h, (C.int)(filters), &_out, &_out_Lengths, &_out_len)
	ret := make([]string, int(_out_len))
	_outCast := (*[0xffff]*C.char)(unsafe.Pointer(_out)) // hey ya
	_out_LengthsCast := (*[0xffff]C.int)(unsafe.Pointer(_out_Lengths))
	for i := 0; i < int(_out_len); i++ {
		ret[i] = C.GoStringN(_outCast[i], _out_LengthsCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) EntryList2(filters int, sort int) []string {
	var _out **C.char = nil
	var _out_Lengths *C.int = nil
	var _out_len C.size_t = 0
	C.QDir_EntryList2(this.h, (C.int)(filters), (C.int)(sort), &_out, &_out_Lengths, &_out_len)
	ret := make([]string, int(_out_len))
	_outCast := (*[0xffff]*C.char)(unsafe.Pointer(_out)) // hey ya
	_out_LengthsCast := (*[0xffff]C.int)(unsafe.Pointer(_out_Lengths))
	for i := 0; i < int(_out_len); i++ {
		ret[i] = C.GoStringN(_outCast[i], _out_LengthsCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) EntryList22(nameFilters []string, filters int) []string {
	// For the C ABI, malloc two C arrays; raw char* pointers and their lengths
	nameFilters_CArray := (*[0xffff]*C.char)(C.malloc(C.size_t(8 * len(nameFilters))))
	nameFilters_Lengths := (*[0xffff]C.uint64_t)(C.malloc(C.size_t(8 * len(nameFilters))))
	defer C.free(unsafe.Pointer(nameFilters_CArray))
	defer C.free(unsafe.Pointer(nameFilters_Lengths))
	for i := range nameFilters {
		single_cstring := C.CString(nameFilters[i])
		defer C.free(unsafe.Pointer(single_cstring))
		nameFilters_CArray[i] = single_cstring
		nameFilters_Lengths[i] = (C.uint64_t)(len(nameFilters[i]))
	}
	var _out **C.char = nil
	var _out_Lengths *C.int = nil
	var _out_len C.size_t = 0
	C.QDir_EntryList22(this.h, &nameFilters_CArray[0], &nameFilters_Lengths[0], C.size_t(len(nameFilters)), (C.int)(filters), &_out, &_out_Lengths, &_out_len)
	ret := make([]string, int(_out_len))
	_outCast := (*[0xffff]*C.char)(unsafe.Pointer(_out)) // hey ya
	_out_LengthsCast := (*[0xffff]C.int)(unsafe.Pointer(_out_Lengths))
	for i := 0; i < int(_out_len); i++ {
		ret[i] = C.GoStringN(_outCast[i], _out_LengthsCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) EntryList3(nameFilters []string, filters int, sort int) []string {
	// For the C ABI, malloc two C arrays; raw char* pointers and their lengths
	nameFilters_CArray := (*[0xffff]*C.char)(C.malloc(C.size_t(8 * len(nameFilters))))
	nameFilters_Lengths := (*[0xffff]C.uint64_t)(C.malloc(C.size_t(8 * len(nameFilters))))
	defer C.free(unsafe.Pointer(nameFilters_CArray))
	defer C.free(unsafe.Pointer(nameFilters_Lengths))
	for i := range nameFilters {
		single_cstring := C.CString(nameFilters[i])
		defer C.free(unsafe.Pointer(single_cstring))
		nameFilters_CArray[i] = single_cstring
		nameFilters_Lengths[i] = (C.uint64_t)(len(nameFilters[i]))
	}
	var _out **C.char = nil
	var _out_Lengths *C.int = nil
	var _out_len C.size_t = 0
	C.QDir_EntryList3(this.h, &nameFilters_CArray[0], &nameFilters_Lengths[0], C.size_t(len(nameFilters)), (C.int)(filters), (C.int)(sort), &_out, &_out_Lengths, &_out_len)
	ret := make([]string, int(_out_len))
	_outCast := (*[0xffff]*C.char)(unsafe.Pointer(_out)) // hey ya
	_out_LengthsCast := (*[0xffff]C.int)(unsafe.Pointer(_out_Lengths))
	for i := 0; i < int(_out_len); i++ {
		ret[i] = C.GoStringN(_outCast[i], _out_LengthsCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) EntryInfoList1(filters int) []QFileInfo {
	var _out **C.QFileInfo = nil
	var _out_len C.size_t = 0
	C.QDir_EntryInfoList1(this.h, (C.int)(filters), &_out, &_out_len)
	ret := make([]QFileInfo, int(_out_len))
	_outCast := (*[0xffff]*C.QFileInfo)(unsafe.Pointer(_out)) // so fresh so clean
	for i := 0; i < int(_out_len); i++ {
		ret[i] = *newQFileInfo(_outCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) EntryInfoList2(filters int, sort int) []QFileInfo {
	var _out **C.QFileInfo = nil
	var _out_len C.size_t = 0
	C.QDir_EntryInfoList2(this.h, (C.int)(filters), (C.int)(sort), &_out, &_out_len)
	ret := make([]QFileInfo, int(_out_len))
	_outCast := (*[0xffff]*C.QFileInfo)(unsafe.Pointer(_out)) // so fresh so clean
	for i := 0; i < int(_out_len); i++ {
		ret[i] = *newQFileInfo(_outCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) EntryInfoList22(nameFilters []string, filters int) []QFileInfo {
	// For the C ABI, malloc two C arrays; raw char* pointers and their lengths
	nameFilters_CArray := (*[0xffff]*C.char)(C.malloc(C.size_t(8 * len(nameFilters))))
	nameFilters_Lengths := (*[0xffff]C.uint64_t)(C.malloc(C.size_t(8 * len(nameFilters))))
	defer C.free(unsafe.Pointer(nameFilters_CArray))
	defer C.free(unsafe.Pointer(nameFilters_Lengths))
	for i := range nameFilters {
		single_cstring := C.CString(nameFilters[i])
		defer C.free(unsafe.Pointer(single_cstring))
		nameFilters_CArray[i] = single_cstring
		nameFilters_Lengths[i] = (C.uint64_t)(len(nameFilters[i]))
	}
	var _out **C.QFileInfo = nil
	var _out_len C.size_t = 0
	C.QDir_EntryInfoList22(this.h, &nameFilters_CArray[0], &nameFilters_Lengths[0], C.size_t(len(nameFilters)), (C.int)(filters), &_out, &_out_len)
	ret := make([]QFileInfo, int(_out_len))
	_outCast := (*[0xffff]*C.QFileInfo)(unsafe.Pointer(_out)) // so fresh so clean
	for i := 0; i < int(_out_len); i++ {
		ret[i] = *newQFileInfo(_outCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) EntryInfoList3(nameFilters []string, filters int, sort int) []QFileInfo {
	// For the C ABI, malloc two C arrays; raw char* pointers and their lengths
	nameFilters_CArray := (*[0xffff]*C.char)(C.malloc(C.size_t(8 * len(nameFilters))))
	nameFilters_Lengths := (*[0xffff]C.uint64_t)(C.malloc(C.size_t(8 * len(nameFilters))))
	defer C.free(unsafe.Pointer(nameFilters_CArray))
	defer C.free(unsafe.Pointer(nameFilters_Lengths))
	for i := range nameFilters {
		single_cstring := C.CString(nameFilters[i])
		defer C.free(unsafe.Pointer(single_cstring))
		nameFilters_CArray[i] = single_cstring
		nameFilters_Lengths[i] = (C.uint64_t)(len(nameFilters[i]))
	}
	var _out **C.QFileInfo = nil
	var _out_len C.size_t = 0
	C.QDir_EntryInfoList3(this.h, &nameFilters_CArray[0], &nameFilters_Lengths[0], C.size_t(len(nameFilters)), (C.int)(filters), (C.int)(sort), &_out, &_out_len)
	ret := make([]QFileInfo, int(_out_len))
	_outCast := (*[0xffff]*C.QFileInfo)(unsafe.Pointer(_out)) // so fresh so clean
	for i := 0; i < int(_out_len); i++ {
		ret[i] = *newQFileInfo(_outCast[i])
	}
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QDir) Delete() {
	C.QDir_Delete(this.h)
}
