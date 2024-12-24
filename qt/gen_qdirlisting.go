package qt

import (
	"unsafe"
)

type QDirListing__IteratorFlag int

const (
	QDirListing__Default             QDirListing__IteratorFlag = 0
	QDirListing__ExcludeFiles        QDirListing__IteratorFlag = 4
	QDirListing__ExcludeDirs         QDirListing__IteratorFlag = 8
	QDirListing__ExcludeSpecial      QDirListing__IteratorFlag = 16
	QDirListing__ResolveSymlinks     QDirListing__IteratorFlag = 32
	QDirListing__FilesOnly           QDirListing__IteratorFlag = 24
	QDirListing__DirsOnly            QDirListing__IteratorFlag = 20
	QDirListing__IncludeHidden       QDirListing__IteratorFlag = 64
	QDirListing__IncludeDotAndDotDot QDirListing__IteratorFlag = 128
	QDirListing__CaseSensitive       QDirListing__IteratorFlag = 256
	QDirListing__Recursive           QDirListing__IteratorFlag = 1024
	QDirListing__FollowDirSymlinks   QDirListing__IteratorFlag = 2048
)

type QDirListing struct {
	h          uintptr
	isSubclass bool
}

// NewQDirListing constructs a new QDirListing object.
func NewQDirListing(path string) *QDirListing {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	g := newQDirListing(QDirListing_new(path_ms))
	g.isSubclass = true
	return g
}

// NewQDirListing2 constructs a new QDirListing object.
func NewQDirListing2(path string, nameFilters []string) *QDirListing {
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
	g := newQDirListing(QDirListing_new2(path_ms, nameFilters_ma))
	g.isSubclass = true
	return g
}

// NewQDirListing3 constructs a new QDirListing object.
func NewQDirListing3(path string, flags IteratorFlags) *QDirListing {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	g := newQDirListing(QDirListing_new3(path_ms, flags))
	g.isSubclass = true
	return g
}

// NewQDirListing4 constructs a new QDirListing object.
func NewQDirListing4(path string, nameFilters []string, flags IteratorFlags) *QDirListing {
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
	g := newQDirListing(QDirListing_new4(path_ms, nameFilters_ma, flags))
	g.isSubclass = true
	return g
}

func (this *QDirListing) Swap(other *QDirListing) {
	QDirListing_Swap(this.h, other.cPointer())
}

func (this *QDirListing) IteratorPath() string {
	var _ms struct_miqt_string = QDirListing_IteratorPath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDirListing) IteratorFlags() IteratorFlags {
	xxxxxxxxx
}

func (this *QDirListing) NameFilters() []string {
	var _ma struct_miqt_array = QDirListing_NameFilters(this.h)
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

func (this *QDirListing) Begin() const_iterator {
	xxxxxxxxx
}

func (this *QDirListing) Cbegin() const_iterator {
	xxxxxxxxx
}

func (this *QDirListing) End() sentinel {
	xxxxxxxxx
}

func (this *QDirListing) Cend() sentinel {
	xxxxxxxxx
}

func (this *QDirListing) ConstBegin() const_iterator {
	xxxxxxxxx
}

func (this *QDirListing) ConstEnd() sentinel {
	xxxxxxxxx
}

type QDirListing__DirEntry struct {
	h          uintptr
	isSubclass bool
}

// NewQDirListing__DirEntry constructs a new QDirListing::DirEntry object.
func NewQDirListing__DirEntry(param1 *DirEntry) *QDirListing__DirEntry {
	g := newQDirListing__DirEntry(QDirListing__DirEntry_new(param1))
	g.isSubclass = true
	return g
}

// NewQDirListing__DirEntry2 constructs a new QDirListing::DirEntry object.
func NewQDirListing__DirEntry2() *QDirListing__DirEntry {
	g := newQDirListing__DirEntry(QDirListing__DirEntry_new2())
	g.isSubclass = true
	return g
}

func (this *QDirListing__DirEntry) FileName() string {
	var _ms struct_miqt_string = QDirListing__DirEntry_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDirListing__DirEntry) BaseName() string {
	var _ms struct_miqt_string = QDirListing__DirEntry_BaseName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDirListing__DirEntry) CompleteBaseName() string {
	var _ms struct_miqt_string = QDirListing__DirEntry_CompleteBaseName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDirListing__DirEntry) Suffix() string {
	var _ms struct_miqt_string = QDirListing__DirEntry_Suffix(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDirListing__DirEntry) BundleName() string {
	var _ms struct_miqt_string = QDirListing__DirEntry_BundleName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDirListing__DirEntry) CompleteSuffix() string {
	var _ms struct_miqt_string = QDirListing__DirEntry_CompleteSuffix(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDirListing__DirEntry) FilePath() string {
	var _ms struct_miqt_string = QDirListing__DirEntry_FilePath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDirListing__DirEntry) IsDir() bool {
	return (bool)(QDirListing__DirEntry_IsDir(this.h))
}

func (this *QDirListing__DirEntry) IsFile() bool {
	return (bool)(QDirListing__DirEntry_IsFile(this.h))
}

func (this *QDirListing__DirEntry) IsSymLink() bool {
	return (bool)(QDirListing__DirEntry_IsSymLink(this.h))
}

func (this *QDirListing__DirEntry) Exists() bool {
	return (bool)(QDirListing__DirEntry_Exists(this.h))
}

func (this *QDirListing__DirEntry) IsHidden() bool {
	return (bool)(QDirListing__DirEntry_IsHidden(this.h))
}

func (this *QDirListing__DirEntry) IsReadable() bool {
	return (bool)(QDirListing__DirEntry_IsReadable(this.h))
}

func (this *QDirListing__DirEntry) IsWritable() bool {
	return (bool)(QDirListing__DirEntry_IsWritable(this.h))
}

func (this *QDirListing__DirEntry) IsExecutable() bool {
	return (bool)(QDirListing__DirEntry_IsExecutable(this.h))
}

func (this *QDirListing__DirEntry) FileInfo() *QFileInfo {
	_goptr := newQFileInfo(QDirListing__DirEntry_FileInfo(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDirListing__DirEntry) CanonicalFilePath() string {
	var _ms struct_miqt_string = QDirListing__DirEntry_CanonicalFilePath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDirListing__DirEntry) AbsoluteFilePath() string {
	var _ms struct_miqt_string = QDirListing__DirEntry_AbsoluteFilePath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDirListing__DirEntry) AbsolutePath() string {
	var _ms struct_miqt_string = QDirListing__DirEntry_AbsolutePath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDirListing__DirEntry) Size() int64 {
	return (int64)(QDirListing__DirEntry_Size(this.h))
}

func (this *QDirListing__DirEntry) BirthTime(tz *QTimeZone) *QDateTime {
	_goptr := newQDateTime(QDirListing__DirEntry_BirthTime(this.h, tz.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDirListing__DirEntry) MetadataChangeTime(tz *QTimeZone) *QDateTime {
	_goptr := newQDateTime(QDirListing__DirEntry_MetadataChangeTime(this.h, tz.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDirListing__DirEntry) LastModified(tz *QTimeZone) *QDateTime {
	_goptr := newQDateTime(QDirListing__DirEntry_LastModified(this.h, tz.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDirListing__DirEntry) LastRead(tz *QTimeZone) *QDateTime {
	_goptr := newQDateTime(QDirListing__DirEntry_LastRead(this.h, tz.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDirListing__DirEntry) FileTime(typeVal QFileDevice__FileTime, tz *QTimeZone) *QDateTime {
	_goptr := newQDateTime(QDirListing__DirEntry_FileTime(this.h, (int)(typeVal), tz.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDirListing__DirEntry) OperatorAssign(param1 *DirEntry) {
	QDirListing__DirEntry_OperatorAssign(this.h, param1)
}

type QDirListing__sentinel struct {
	h          uintptr
	isSubclass bool
}

// NewQDirListing__sentinel constructs a new QDirListing::sentinel object.
func NewQDirListing__sentinel() *QDirListing__sentinel {
	g := newQDirListing__sentinel(QDirListing__sentinel_new())
	g.isSubclass = true
	return g
}

// NewQDirListing__sentinel2 constructs a new QDirListing::sentinel object.
func NewQDirListing__sentinel2(param1 *sentinel) *QDirListing__sentinel {
	g := newQDirListing__sentinel(QDirListing__sentinel_new2(param1))
	g.isSubclass = true
	return g
}

type QDirListing__const_iterator struct {
	h          uintptr
	isSubclass bool
}

// NewQDirListing__const_iterator constructs a new QDirListing::const_iterator object.
func NewQDirListing__const_iterator() *QDirListing__const_iterator {
	g := newQDirListing__const_iterator(QDirListing__const_iterator_new())
	g.isSubclass = true
	return g
}

func (this *QDirListing__const_iterator) OperatorMultiply() reference {
	xxxxxxxxx
}

func (this *QDirListing__const_iterator) OperatorMinusGreater() pointer {
	xxxxxxxxx
}

func (this *QDirListing__const_iterator) OperatorPlusPlus() *const_iterator {
	xxxxxxxxx
}

func (this *QDirListing__const_iterator) OperatorPlusPlusWithInt(param1 int) {
	QDirListing__const_iterator_OperatorPlusPlusWithInt(this.h, (int)(param1))
}
