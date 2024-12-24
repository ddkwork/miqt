package qt

import (
	"unsafe"
)

type QDirIterator__IteratorFlag int

const (
	QDirIterator__NoIteratorFlags QDirIterator__IteratorFlag = 0
	QDirIterator__FollowSymlinks  QDirIterator__IteratorFlag = 1
	QDirIterator__Subdirectories  QDirIterator__IteratorFlag = 2
)

type QDirIterator struct {
	h          uintptr
	isSubclass bool
}

// NewQDirIterator constructs a new QDirIterator object.
func NewQDirIterator(dir *QDir) *QDirIterator {
	g := newQDirIterator(QDirIterator_new(dir.cPointer()))
	g.isSubclass = true
	return g
}

// NewQDirIterator2 constructs a new QDirIterator object.
func NewQDirIterator2(path string) *QDirIterator {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	g := newQDirIterator(QDirIterator_new2(path_ms))
	g.isSubclass = true
	return g
}

// NewQDirIterator3 constructs a new QDirIterator object.
func NewQDirIterator3(path string, filter Filter) *QDirIterator {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	g := newQDirIterator(QDirIterator_new3(path_ms, (int)(filter)))
	g.isSubclass = true
	return g
}

// NewQDirIterator4 constructs a new QDirIterator object.
func NewQDirIterator4(path string, nameFilters []string) *QDirIterator {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
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
	g := newQDirIterator(QDirIterator_new4(path_ms, nameFilters_ma))
	g.isSubclass = true
	return g
}

// NewQDirIterator5 constructs a new QDirIterator object.
func NewQDirIterator5(dir *QDir, flags IteratorFlags) *QDirIterator {
	g := newQDirIterator(QDirIterator_new5(dir.cPointer(), flags))
	g.isSubclass = true
	return g
}

// NewQDirIterator6 constructs a new QDirIterator object.
func NewQDirIterator6(path string, flags IteratorFlags) *QDirIterator {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	g := newQDirIterator(QDirIterator_new6(path_ms, flags))
	g.isSubclass = true
	return g
}

// NewQDirIterator7 constructs a new QDirIterator object.
func NewQDirIterator7(path string, filter Filter, flags IteratorFlags) *QDirIterator {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	g := newQDirIterator(QDirIterator_new7(path_ms, (int)(filter), flags))
	g.isSubclass = true
	return g
}

// NewQDirIterator8 constructs a new QDirIterator object.
func NewQDirIterator8(path string, nameFilters []string, filters Filter) *QDirIterator {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
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
	g := newQDirIterator(QDirIterator_new8(path_ms, nameFilters_ma, (int)(filters)))
	g.isSubclass = true
	return g
}

// NewQDirIterator9 constructs a new QDirIterator object.
func NewQDirIterator9(path string, nameFilters []string, filters Filter, flags IteratorFlags) *QDirIterator {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
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
	g := newQDirIterator(QDirIterator_new9(path_ms, nameFilters_ma, (int)(filters), flags))
	g.isSubclass = true
	return g
}

func (this *QDirIterator) Next() string {
	var _ms struct_miqt_string = QDirIterator_Next(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDirIterator) NextFileInfo() *QFileInfo {
	_goptr := newQFileInfo(QDirIterator_NextFileInfo(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDirIterator) HasNext() bool {
	return (bool)(QDirIterator_HasNext(this.h))
}

func (this *QDirIterator) FileName() string {
	var _ms struct_miqt_string = QDirIterator_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDirIterator) FilePath() string {
	var _ms struct_miqt_string = QDirIterator_FilePath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDirIterator) FileInfo() *QFileInfo {
	_goptr := newQFileInfo(QDirIterator_FileInfo(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDirIterator) Path() string {
	var _ms struct_miqt_string = QDirIterator_Path(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
