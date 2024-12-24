package qt

import (
	"unsafe"
)

type QFileSystemModel__Roles int

const (
	QFileSystemModel__FileIconRole    QFileSystemModel__Roles = 1
	QFileSystemModel__FileInfoRole    QFileSystemModel__Roles = 252
	QFileSystemModel__FilePathRole    QFileSystemModel__Roles = 257
	QFileSystemModel__FileNameRole    QFileSystemModel__Roles = 258
	QFileSystemModel__FilePermissions QFileSystemModel__Roles = 259
)

type QFileSystemModel__Option int

const (
	QFileSystemModel__DontWatchForChanges         QFileSystemModel__Option = 1
	QFileSystemModel__DontResolveSymlinks         QFileSystemModel__Option = 2
	QFileSystemModel__DontUseCustomDirectoryIcons QFileSystemModel__Option = 4
)

type QFileSystemModel struct {
	h          uintptr
	isSubclass bool
}

// NewQFileSystemModel constructs a new QFileSystemModel object.
func NewQFileSystemModel() *QFileSystemModel {
	ret := newQFileSystemModel(QFileSystemModel_new())
	ret.isSubclass = true
	return ret
}

// NewQFileSystemModel2 constructs a new QFileSystemModel object.
func NewQFileSystemModel2(parent *QObject) *QFileSystemModel {
	ret := newQFileSystemModel(QFileSystemModel_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QFileSystemModel) MetaObject() *QMetaObject {
	return newQMetaObject(QFileSystemModel_MetaObject(this.h))
}

func (this *QFileSystemModel) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QFileSystemModel_Metacast(this.h, param1_Cstring))
}

func QFileSystemModel_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QFileSystemModel_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileSystemModel) RootPathChanged(newPath string) {
	newPath_ms := struct_miqt_string{}
	newPath_ms.data = CString(newPath)
	newPath_ms.len = size_t(len(newPath))
	defer free(unsafe.Pointer(newPath_ms.data))
	QFileSystemModel_RootPathChanged(this.h, newPath_ms)
}

func (this *QFileSystemModel) OnRootPathChanged(slot func(newPath string)) {
	QFileSystemModel_connect_RootPathChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_RootPathChanged
func miqt_exec_callback_QFileSystemModel_RootPathChanged(cb intptr_t, newPath struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(newPath string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var newPath_ms struct_miqt_string = newPath
	newPath_ret := GoStringN(newPath_ms.data, int(int64(newPath_ms.len)))
	free(unsafe.Pointer(newPath_ms.data))
	slotval1 := newPath_ret

	gofunc(slotval1)
}

func (this *QFileSystemModel) FileRenamed(path string, oldName string, newName string) {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	oldName_ms := struct_miqt_string{}
	oldName_ms.data = CString(oldName)
	oldName_ms.len = size_t(len(oldName))
	defer free(unsafe.Pointer(oldName_ms.data))
	newName_ms := struct_miqt_string{}
	newName_ms.data = CString(newName)
	newName_ms.len = size_t(len(newName))
	defer free(unsafe.Pointer(newName_ms.data))
	QFileSystemModel_FileRenamed(this.h, path_ms, oldName_ms, newName_ms)
}

func (this *QFileSystemModel) OnFileRenamed(slot func(path string, oldName string, newName string)) {
	QFileSystemModel_connect_FileRenamed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_FileRenamed
func miqt_exec_callback_QFileSystemModel_FileRenamed(cb intptr_t, path struct_miqt_string, oldName struct_miqt_string, newName struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(path string, oldName string, newName string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var path_ms struct_miqt_string = path
	path_ret := GoStringN(path_ms.data, int(int64(path_ms.len)))
	free(unsafe.Pointer(path_ms.data))
	slotval1 := path_ret
	var oldName_ms struct_miqt_string = oldName
	oldName_ret := GoStringN(oldName_ms.data, int(int64(oldName_ms.len)))
	free(unsafe.Pointer(oldName_ms.data))
	slotval2 := oldName_ret
	var newName_ms struct_miqt_string = newName
	newName_ret := GoStringN(newName_ms.data, int(int64(newName_ms.len)))
	free(unsafe.Pointer(newName_ms.data))
	slotval3 := newName_ret

	gofunc(slotval1, slotval2, slotval3)
}

func (this *QFileSystemModel) DirectoryLoaded(path string) {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	QFileSystemModel_DirectoryLoaded(this.h, path_ms)
}

func (this *QFileSystemModel) OnDirectoryLoaded(slot func(path string)) {
	QFileSystemModel_connect_DirectoryLoaded(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_DirectoryLoaded
func miqt_exec_callback_QFileSystemModel_DirectoryLoaded(cb intptr_t, path struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(path string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var path_ms struct_miqt_string = path
	path_ret := GoStringN(path_ms.data, int(int64(path_ms.len)))
	free(unsafe.Pointer(path_ms.data))
	slotval1 := path_ret

	gofunc(slotval1)
}

func (this *QFileSystemModel) Index(row int, column int, parent *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QFileSystemModel_Index(this.h, (int)(row), (int)(column), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSystemModel) IndexWithPath(path string) *QModelIndex {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	_goptr := newQModelIndex(QFileSystemModel_IndexWithPath(this.h, path_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSystemModel) Parent(child *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QFileSystemModel_Parent(this.h, child.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSystemModel) Sibling(row int, column int, idx *QModelIndex) *QModelIndex {
	_goptr := newQModelIndex(QFileSystemModel_Sibling(this.h, (int)(row), (int)(column), idx.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSystemModel) HasChildren(parent *QModelIndex) bool {
	return (bool)(QFileSystemModel_HasChildren(this.h, parent.cPointer()))
}

func (this *QFileSystemModel) CanFetchMore(parent *QModelIndex) bool {
	return (bool)(QFileSystemModel_CanFetchMore(this.h, parent.cPointer()))
}

func (this *QFileSystemModel) FetchMore(parent *QModelIndex) {
	QFileSystemModel_FetchMore(this.h, parent.cPointer())
}

func (this *QFileSystemModel) RowCount(parent *QModelIndex) int {
	return (int)(QFileSystemModel_RowCount(this.h, parent.cPointer()))
}

func (this *QFileSystemModel) ColumnCount(parent *QModelIndex) int {
	return (int)(QFileSystemModel_ColumnCount(this.h, parent.cPointer()))
}

func (this *QFileSystemModel) MyComputer() *QVariant {
	_goptr := newQVariant(QFileSystemModel_MyComputer(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSystemModel) Data(index *QModelIndex, role int) *QVariant {
	_goptr := newQVariant(QFileSystemModel_Data(this.h, index.cPointer(), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSystemModel) SetData(index *QModelIndex, value *QVariant, role int) bool {
	return (bool)(QFileSystemModel_SetData(this.h, index.cPointer(), value.cPointer(), (int)(role)))
}

func (this *QFileSystemModel) HeaderData(section int, orientation Orientation, role int) *QVariant {
	_goptr := newQVariant(QFileSystemModel_HeaderData(this.h, (int)(section), (int)(orientation), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSystemModel) Flags(index *QModelIndex) ItemFlag {
	return (ItemFlag)(QFileSystemModel_Flags(this.h, index.cPointer()))
}

func (this *QFileSystemModel) Sort(column int, order SortOrder) {
	QFileSystemModel_Sort(this.h, (int)(column), (int)(order))
}

func (this *QFileSystemModel) MimeTypes() []string {
	var _ma struct_miqt_array = QFileSystemModel_MimeTypes(this.h)
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

func (this *QFileSystemModel) MimeData(indexes []QModelIndex) *QMimeData {
	indexes_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(indexes))))
	defer free(unsafe.Pointer(indexes_CArray))
	for i := range indexes {
		indexes_CArray[i] = indexes[i].cPointer()
	}
	indexes_ma := struct_miqt_array{len: size_t(len(indexes)), data: unsafe.Pointer(indexes_CArray)}
	return newQMimeData(QFileSystemModel_MimeData(this.h, indexes_ma))
}

func (this *QFileSystemModel) DropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	return (bool)(QFileSystemModel_DropMimeData(this.h, data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))
}

func (this *QFileSystemModel) SupportedDropActions() DropAction {
	return (DropAction)(QFileSystemModel_SupportedDropActions(this.h))
}

func (this *QFileSystemModel) RoleNames() map[int][]byte {
	var _mm struct_miqt_map = QFileSystemModel_RoleNames(this.h)
	_ret := make(map[int][]byte, int(_mm.len))
	_Keys := (*[0xffff]int)(unsafe.Pointer(_mm.keys))
	_Values := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_mm.values))
	for i := 0; i < int(_mm.len); i++ {
		_entry_Key := (int)(_Keys[i])

		var _hashval_bytearray struct_miqt_string = _Values[i]
		_hashval_ret := GoBytes(unsafe.Pointer(_hashval_bytearray.data), int(int64(_hashval_bytearray.len)))
		free(unsafe.Pointer(_hashval_bytearray.data))
		_entry_Value := _hashval_ret
		_ret[_entry_Key] = _entry_Value
	}
	return _ret
}

func (this *QFileSystemModel) SetRootPath(path string) *QModelIndex {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	_goptr := newQModelIndex(QFileSystemModel_SetRootPath(this.h, path_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSystemModel) RootPath() string {
	var _ms struct_miqt_string = QFileSystemModel_RootPath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileSystemModel) RootDirectory() *QDir {
	_goptr := newQDir(QFileSystemModel_RootDirectory(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSystemModel) SetIconProvider(provider *QAbstractFileIconProvider) {
	QFileSystemModel_SetIconProvider(this.h, provider.cPointer())
}

func (this *QFileSystemModel) IconProvider() *QAbstractFileIconProvider {
	return newQAbstractFileIconProvider(QFileSystemModel_IconProvider(this.h))
}

func (this *QFileSystemModel) SetFilter(filters Filter) {
	QFileSystemModel_SetFilter(this.h, (int)(filters))
}

func (this *QFileSystemModel) Filter() Filter {
	return (Filter)(QFileSystemModel_Filter(this.h))
}

func (this *QFileSystemModel) SetResolveSymlinks(enable bool) {
	QFileSystemModel_SetResolveSymlinks(this.h, (bool)(enable))
}

func (this *QFileSystemModel) ResolveSymlinks() bool {
	return (bool)(QFileSystemModel_ResolveSymlinks(this.h))
}

func (this *QFileSystemModel) SetReadOnly(enable bool) {
	QFileSystemModel_SetReadOnly(this.h, (bool)(enable))
}

func (this *QFileSystemModel) IsReadOnly() bool {
	return (bool)(QFileSystemModel_IsReadOnly(this.h))
}

func (this *QFileSystemModel) SetNameFilterDisables(enable bool) {
	QFileSystemModel_SetNameFilterDisables(this.h, (bool)(enable))
}

func (this *QFileSystemModel) NameFilterDisables() bool {
	return (bool)(QFileSystemModel_NameFilterDisables(this.h))
}

func (this *QFileSystemModel) SetNameFilters(filters []string) {
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
	QFileSystemModel_SetNameFilters(this.h, filters_ma)
}

func (this *QFileSystemModel) NameFilters() []string {
	var _ma struct_miqt_array = QFileSystemModel_NameFilters(this.h)
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

func (this *QFileSystemModel) SetOption(option Option) {
	QFileSystemModel_SetOption(this.h, option)
}

func (this *QFileSystemModel) TestOption(option Option) bool {
	return (bool)(QFileSystemModel_TestOption(this.h, option))
}

func (this *QFileSystemModel) SetOptions(options Options) {
	QFileSystemModel_SetOptions(this.h, options)
}

func (this *QFileSystemModel) Options() Options {
	xxxxxxxxx
}

func (this *QFileSystemModel) FilePath(index *QModelIndex) string {
	var _ms struct_miqt_string = QFileSystemModel_FilePath(this.h, index.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileSystemModel) IsDir(index *QModelIndex) bool {
	return (bool)(QFileSystemModel_IsDir(this.h, index.cPointer()))
}

func (this *QFileSystemModel) Size(index *QModelIndex) int64 {
	return (int64)(QFileSystemModel_Size(this.h, index.cPointer()))
}

func (this *QFileSystemModel) Type(index *QModelIndex) string {
	var _ms struct_miqt_string = QFileSystemModel_Type(this.h, index.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileSystemModel) LastModified(index *QModelIndex) *QDateTime {
	_goptr := newQDateTime(QFileSystemModel_LastModified(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSystemModel) LastModified2(index *QModelIndex, tz *QTimeZone) *QDateTime {
	_goptr := newQDateTime(QFileSystemModel_LastModified2(this.h, index.cPointer(), tz.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSystemModel) Mkdir(parent *QModelIndex, name string) *QModelIndex {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	_goptr := newQModelIndex(QFileSystemModel_Mkdir(this.h, parent.cPointer(), name_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSystemModel) Rmdir(index *QModelIndex) bool {
	return (bool)(QFileSystemModel_Rmdir(this.h, index.cPointer()))
}

func (this *QFileSystemModel) FileName(index *QModelIndex) string {
	var _ms struct_miqt_string = QFileSystemModel_FileName(this.h, index.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileSystemModel) FileIcon(index *QModelIndex) *QIcon {
	_goptr := newQIcon(QFileSystemModel_FileIcon(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSystemModel) Permissions(index *QModelIndex) Permission {
	return (Permission)(QFileSystemModel_Permissions(this.h, index.cPointer()))
}

func (this *QFileSystemModel) FileInfo(index *QModelIndex) *QFileInfo {
	_goptr := newQFileInfo(QFileSystemModel_FileInfo(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSystemModel) Remove(index *QModelIndex) bool {
	return (bool)(QFileSystemModel_Remove(this.h, index.cPointer()))
}

func QFileSystemModel_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFileSystemModel_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileSystemModel_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFileSystemModel_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileSystemModel) Index2(path string, column int) *QModelIndex {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	_goptr := newQModelIndex(QFileSystemModel_Index2(this.h, path_ms, (int)(column)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSystemModel) MyComputer1(role int) *QVariant {
	_goptr := newQVariant(QFileSystemModel_MyComputer1(this.h, (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSystemModel) SetOption2(option Option, on bool) {
	QFileSystemModel_SetOption2(this.h, option, (bool)(on))
}

func (this *QFileSystemModel) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QFileSystemModel_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QFileSystemModel) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_MetaObject
func miqt_exec_callback_QFileSystemModel_MetaObject(self QFileSystemModel, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QFileSystemModel) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QFileSystemModel_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QFileSystemModel) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_Metacast
func miqt_exec_callback_QFileSystemModel_Metacast(self QFileSystemModel, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
