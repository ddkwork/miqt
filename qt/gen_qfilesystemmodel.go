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

func (this *QFileSystemModel) callVirtualBase_Index(row int, column int, parent *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QFileSystemModel_virtualbase_Index(unsafe.Pointer(this.h), (int)(row), (int)(column), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFileSystemModel) OnIndex(slot func(super func(row int, column int, parent *QModelIndex) *QModelIndex, row int, column int, parent *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_Index(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_Index
func miqt_exec_callback_QFileSystemModel_Index(self QFileSystemModel, cb intptr_t, row int, column int, parent *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, column int, parent *QModelIndex) *QModelIndex, row int, column int, parent *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_Index, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QFileSystemModel) callVirtualBase_Parent(child *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QFileSystemModel_virtualbase_Parent(unsafe.Pointer(this.h), child.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFileSystemModel) OnParent(slot func(super func(child *QModelIndex) *QModelIndex, child *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_Parent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_Parent
func miqt_exec_callback_QFileSystemModel_Parent(self QFileSystemModel, cb intptr_t, child *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(child *QModelIndex) *QModelIndex, child *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(child)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_Parent, slotval1)

	return virtualReturn.cPointer()

}

func (this *QFileSystemModel) callVirtualBase_Sibling(row int, column int, idx *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QFileSystemModel_virtualbase_Sibling(unsafe.Pointer(this.h), (int)(row), (int)(column), idx.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFileSystemModel) OnSibling(slot func(super func(row int, column int, idx *QModelIndex) *QModelIndex, row int, column int, idx *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_Sibling(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_Sibling
func miqt_exec_callback_QFileSystemModel_Sibling(self QFileSystemModel, cb intptr_t, row int, column int, idx *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, column int, idx *QModelIndex) *QModelIndex, row int, column int, idx *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(column)

	slotval3 := newQModelIndex(idx)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_Sibling, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QFileSystemModel) callVirtualBase_HasChildren(parent *QModelIndex) bool {

	return (bool)(QFileSystemModel_virtualbase_HasChildren(unsafe.Pointer(this.h), parent.cPointer()))

}
func (this *QFileSystemModel) OnHasChildren(slot func(super func(parent *QModelIndex) bool, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_HasChildren(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_HasChildren
func miqt_exec_callback_QFileSystemModel_HasChildren(self QFileSystemModel, cb intptr_t, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex) bool, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_HasChildren, slotval1)

	return (bool)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_CanFetchMore(parent *QModelIndex) bool {

	return (bool)(QFileSystemModel_virtualbase_CanFetchMore(unsafe.Pointer(this.h), parent.cPointer()))

}
func (this *QFileSystemModel) OnCanFetchMore(slot func(super func(parent *QModelIndex) bool, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_CanFetchMore(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_CanFetchMore
func miqt_exec_callback_QFileSystemModel_CanFetchMore(self QFileSystemModel, cb intptr_t, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex) bool, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_CanFetchMore, slotval1)

	return (bool)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_FetchMore(parent *QModelIndex) {

	QFileSystemModel_virtualbase_FetchMore(unsafe.Pointer(this.h), parent.cPointer())

}
func (this *QFileSystemModel) OnFetchMore(slot func(super func(parent *QModelIndex), parent *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_FetchMore(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_FetchMore
func miqt_exec_callback_QFileSystemModel_FetchMore(self QFileSystemModel, cb intptr_t, parent *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex), parent *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	gofunc((&QFileSystemModel{h: self}).callVirtualBase_FetchMore, slotval1)

}

func (this *QFileSystemModel) callVirtualBase_RowCount(parent *QModelIndex) int {

	return (int)(QFileSystemModel_virtualbase_RowCount(unsafe.Pointer(this.h), parent.cPointer()))

}
func (this *QFileSystemModel) OnRowCount(slot func(super func(parent *QModelIndex) int, parent *QModelIndex) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_RowCount(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_RowCount
func miqt_exec_callback_QFileSystemModel_RowCount(self QFileSystemModel, cb intptr_t, parent *QModelIndex) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex) int, parent *QModelIndex) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_RowCount, slotval1)

	return (int)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_ColumnCount(parent *QModelIndex) int {

	return (int)(QFileSystemModel_virtualbase_ColumnCount(unsafe.Pointer(this.h), parent.cPointer()))

}
func (this *QFileSystemModel) OnColumnCount(slot func(super func(parent *QModelIndex) int, parent *QModelIndex) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_ColumnCount(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_ColumnCount
func miqt_exec_callback_QFileSystemModel_ColumnCount(self QFileSystemModel, cb intptr_t, parent *QModelIndex) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex) int, parent *QModelIndex) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_ColumnCount, slotval1)

	return (int)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_Data(index *QModelIndex, role int) *QVariant {

	_goptr := newQVariant(QFileSystemModel_virtualbase_Data(unsafe.Pointer(this.h), index.cPointer(), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFileSystemModel) OnData(slot func(super func(index *QModelIndex, role int) *QVariant, index *QModelIndex, role int) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_Data(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_Data
func miqt_exec_callback_QFileSystemModel_Data(self QFileSystemModel, cb intptr_t, index *QModelIndex, role int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, role int) *QVariant, index *QModelIndex, role int) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := (int)(role)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_Data, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QFileSystemModel) callVirtualBase_SetData(index *QModelIndex, value *QVariant, role int) bool {

	return (bool)(QFileSystemModel_virtualbase_SetData(unsafe.Pointer(this.h), index.cPointer(), value.cPointer(), (int)(role)))

}
func (this *QFileSystemModel) OnSetData(slot func(super func(index *QModelIndex, value *QVariant, role int) bool, index *QModelIndex, value *QVariant, role int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_SetData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_SetData
func miqt_exec_callback_QFileSystemModel_SetData(self QFileSystemModel, cb intptr_t, index *QModelIndex, value *QVariant, role int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, value *QVariant, role int) bool, index *QModelIndex, value *QVariant, role int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := newQVariant(value)

	slotval3 := (int)(role)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_SetData, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_HeaderData(section int, orientation Orientation, role int) *QVariant {

	_goptr := newQVariant(QFileSystemModel_virtualbase_HeaderData(unsafe.Pointer(this.h), (int)(section), (int)(orientation), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFileSystemModel) OnHeaderData(slot func(super func(section int, orientation Orientation, role int) *QVariant, section int, orientation Orientation, role int) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_HeaderData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_HeaderData
func miqt_exec_callback_QFileSystemModel_HeaderData(self QFileSystemModel, cb intptr_t, section int, orientation int, role int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(section int, orientation Orientation, role int) *QVariant, section int, orientation Orientation, role int) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(section)

	slotval2 := (Orientation)(orientation)

	slotval3 := (int)(role)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_HeaderData, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QFileSystemModel) callVirtualBase_Flags(index *QModelIndex) ItemFlag {

	return (ItemFlag)(QFileSystemModel_virtualbase_Flags(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QFileSystemModel) OnFlags(slot func(super func(index *QModelIndex) ItemFlag, index *QModelIndex) ItemFlag) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_Flags(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_Flags
func miqt_exec_callback_QFileSystemModel_Flags(self QFileSystemModel, cb intptr_t, index *QModelIndex) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) ItemFlag, index *QModelIndex) ItemFlag)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_Flags, slotval1)

	return (int)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_Sort(column int, order SortOrder) {

	QFileSystemModel_virtualbase_Sort(unsafe.Pointer(this.h), (int)(column), (int)(order))

}
func (this *QFileSystemModel) OnSort(slot func(super func(column int, order SortOrder), column int, order SortOrder)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_Sort(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_Sort
func miqt_exec_callback_QFileSystemModel_Sort(self QFileSystemModel, cb intptr_t, column int, order int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, order SortOrder), column int, order SortOrder))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (SortOrder)(order)

	gofunc((&QFileSystemModel{h: self}).callVirtualBase_Sort, slotval1, slotval2)

}

func (this *QFileSystemModel) callVirtualBase_MimeTypes() []string {

	var _ma struct_miqt_array = QFileSystemModel_virtualbase_MimeTypes(unsafe.Pointer(this.h))
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
func (this *QFileSystemModel) OnMimeTypes(slot func(super func() []string) []string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_MimeTypes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_MimeTypes
func miqt_exec_callback_QFileSystemModel_MimeTypes(self QFileSystemModel, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []string) []string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_MimeTypes)
	virtualReturn_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_i_ms := struct_miqt_string{}
		virtualReturn_i_ms.data = CString(virtualReturn[i])
		virtualReturn_i_ms.len = size_t(len(virtualReturn[i]))
		defer free(unsafe.Pointer(virtualReturn_i_ms.data))
		virtualReturn_CArray[i] = virtualReturn_i_ms
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QFileSystemModel) callVirtualBase_MimeData(indexes []QModelIndex) *QMimeData {
	indexes_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(indexes))))
	defer free(unsafe.Pointer(indexes_CArray))
	for i := range indexes {
		indexes_CArray[i] = indexes[i].cPointer()
	}
	indexes_ma := struct_miqt_array{len: size_t(len(indexes)), data: unsafe.Pointer(indexes_CArray)}

	return newQMimeData(QFileSystemModel_virtualbase_MimeData(unsafe.Pointer(this.h), indexes_ma))

}
func (this *QFileSystemModel) OnMimeData(slot func(super func(indexes []QModelIndex) *QMimeData, indexes []QModelIndex) *QMimeData) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_MimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_MimeData
func miqt_exec_callback_QFileSystemModel_MimeData(self QFileSystemModel, cb intptr_t, indexes struct_miqt_array) *QMimeData {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(indexes []QModelIndex) *QMimeData, indexes []QModelIndex) *QMimeData)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var indexes_ma struct_miqt_array = indexes
	indexes_ret := make([]QModelIndex, int(indexes_ma.len))
	indexes_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(indexes_ma.data)) // hey ya
	for i := 0; i < int(indexes_ma.len); i++ {
		indexes_lv_goptr := newQModelIndex(indexes_outCast[i])
		indexes_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		indexes_ret[i] = *indexes_lv_goptr
	}
	slotval1 := indexes_ret

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_MimeData, slotval1)

	return virtualReturn.cPointer()

}

func (this *QFileSystemModel) callVirtualBase_DropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {

	return (bool)(QFileSystemModel_virtualbase_DropMimeData(unsafe.Pointer(this.h), data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))

}
func (this *QFileSystemModel) OnDropMimeData(slot func(super func(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool, data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_DropMimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_DropMimeData
func miqt_exec_callback_QFileSystemModel_DropMimeData(self QFileSystemModel, cb intptr_t, data *QMimeData, action int, row int, column int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool, data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMimeData(data)

	slotval2 := (DropAction)(action)

	slotval3 := (int)(row)

	slotval4 := (int)(column)

	slotval5 := newQModelIndex(parent)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_DropMimeData, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_SupportedDropActions() DropAction {

	return (DropAction)(QFileSystemModel_virtualbase_SupportedDropActions(unsafe.Pointer(this.h)))

}
func (this *QFileSystemModel) OnSupportedDropActions(slot func(super func() DropAction) DropAction) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_SupportedDropActions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_SupportedDropActions
func miqt_exec_callback_QFileSystemModel_SupportedDropActions(self QFileSystemModel, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() DropAction) DropAction)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_SupportedDropActions)

	return (int)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_RoleNames() map[int][]byte {

	var _mm struct_miqt_map = QFileSystemModel_virtualbase_RoleNames(unsafe.Pointer(this.h))
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
func (this *QFileSystemModel) OnRoleNames(slot func(super func() map[int][]byte) map[int][]byte) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_RoleNames(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_RoleNames
func miqt_exec_callback_QFileSystemModel_RoleNames(self QFileSystemModel, cb intptr_t) struct_miqt_map {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() map[int][]byte) map[int][]byte)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_RoleNames)
	virtualReturn_Keys_CArray := (*[0xffff]int)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_Keys_CArray))
	virtualReturn_Values_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_Values_CArray))
	virtualReturn_ctr := 0
	for virtualReturn_k, virtualReturn_v := range virtualReturn {
		virtualReturn_Keys_CArray[virtualReturn_ctr] = (int)(virtualReturn_k)
		virtualReturn_v_alias := struct_miqt_string{}
		virtualReturn_v_alias.data = (char)(unsafe.Pointer(&virtualReturn_v[0]))
		virtualReturn_v_alias.len = size_t(len(virtualReturn_v))
		virtualReturn_Values_CArray[virtualReturn_ctr] = virtualReturn_v_alias
		virtualReturn_ctr++
	}
	virtualReturn_mm := struct_miqt_map{
		len:    size_t(len(virtualReturn)),
		keys:   unsafe.Pointer(virtualReturn_Keys_CArray),
		values: unsafe.Pointer(virtualReturn_Values_CArray),
	}

	return virtualReturn_mm

}

func (this *QFileSystemModel) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QFileSystemModel_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFileSystemModel) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_TimerEvent
func miqt_exec_callback_QFileSystemModel_TimerEvent(self QFileSystemModel, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QFileSystemModel{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QFileSystemModel) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QFileSystemModel_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QFileSystemModel) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_Event
func miqt_exec_callback_QFileSystemModel_Event(self QFileSystemModel, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_SetHeaderData(section int, orientation Orientation, value *QVariant, role int) bool {

	return (bool)(QFileSystemModel_virtualbase_SetHeaderData(unsafe.Pointer(this.h), (int)(section), (int)(orientation), value.cPointer(), (int)(role)))

}
func (this *QFileSystemModel) OnSetHeaderData(slot func(super func(section int, orientation Orientation, value *QVariant, role int) bool, section int, orientation Orientation, value *QVariant, role int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_SetHeaderData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_SetHeaderData
func miqt_exec_callback_QFileSystemModel_SetHeaderData(self QFileSystemModel, cb intptr_t, section int, orientation int, value *QVariant, role int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(section int, orientation Orientation, value *QVariant, role int) bool, section int, orientation Orientation, value *QVariant, role int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(section)

	slotval2 := (Orientation)(orientation)

	slotval3 := newQVariant(value)

	slotval4 := (int)(role)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_SetHeaderData, slotval1, slotval2, slotval3, slotval4)

	return (bool)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_ItemData(index *QModelIndex) map[int]QVariant {

	var _mm struct_miqt_map = QFileSystemModel_virtualbase_ItemData(unsafe.Pointer(this.h), index.cPointer())
	_ret := make(map[int]QVariant, int(_mm.len))
	_Keys := (*[0xffff]int)(unsafe.Pointer(_mm.keys))
	_Values := (*[0xffff]*QVariant)(unsafe.Pointer(_mm.values))
	for i := 0; i < int(_mm.len); i++ {
		_entry_Key := (int)(_Keys[i])

		_mapval_goptr := newQVariant(_Values[i])
		_mapval_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_entry_Value := *_mapval_goptr

		_ret[_entry_Key] = _entry_Value
	}
	return _ret

}
func (this *QFileSystemModel) OnItemData(slot func(super func(index *QModelIndex) map[int]QVariant, index *QModelIndex) map[int]QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_ItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_ItemData
func miqt_exec_callback_QFileSystemModel_ItemData(self QFileSystemModel, cb intptr_t, index *QModelIndex) struct_miqt_map {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) map[int]QVariant, index *QModelIndex) map[int]QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_ItemData, slotval1)
	virtualReturn_Keys_CArray := (*[0xffff]int)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_Keys_CArray))
	virtualReturn_Values_CArray := (*[0xffff]*QVariant)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_Values_CArray))
	virtualReturn_ctr := 0
	for virtualReturn_k, virtualReturn_v := range virtualReturn {
		virtualReturn_Keys_CArray[virtualReturn_ctr] = (int)(virtualReturn_k)
		virtualReturn_Values_CArray[virtualReturn_ctr] = virtualReturn_v.cPointer()
		virtualReturn_ctr++
	}
	virtualReturn_mm := struct_miqt_map{
		len:    size_t(len(virtualReturn)),
		keys:   unsafe.Pointer(virtualReturn_Keys_CArray),
		values: unsafe.Pointer(virtualReturn_Values_CArray),
	}

	return virtualReturn_mm

}

func (this *QFileSystemModel) callVirtualBase_SetItemData(index *QModelIndex, roles map[int]QVariant) bool {
	roles_Keys_CArray := (*[0xffff]int)(malloc(size_t(8 * len(roles))))
	defer free(unsafe.Pointer(roles_Keys_CArray))
	roles_Values_CArray := (*[0xffff]*QVariant)(malloc(size_t(8 * len(roles))))
	defer free(unsafe.Pointer(roles_Values_CArray))
	roles_ctr := 0
	for roles_k, roles_v := range roles {
		roles_Keys_CArray[roles_ctr] = (int)(roles_k)
		roles_Values_CArray[roles_ctr] = roles_v.cPointer()
		roles_ctr++
	}
	roles_mm := struct_miqt_map{
		len:    size_t(len(roles)),
		keys:   unsafe.Pointer(roles_Keys_CArray),
		values: unsafe.Pointer(roles_Values_CArray),
	}

	return (bool)(QFileSystemModel_virtualbase_SetItemData(unsafe.Pointer(this.h), index.cPointer(), roles_mm))

}
func (this *QFileSystemModel) OnSetItemData(slot func(super func(index *QModelIndex, roles map[int]QVariant) bool, index *QModelIndex, roles map[int]QVariant) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_SetItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_SetItemData
func miqt_exec_callback_QFileSystemModel_SetItemData(self QFileSystemModel, cb intptr_t, index *QModelIndex, roles struct_miqt_map) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, roles map[int]QVariant) bool, index *QModelIndex, roles map[int]QVariant) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	var roles_mm struct_miqt_map = roles
	roles_ret := make(map[int]QVariant, int(roles_mm.len))
	roles_Keys := (*[0xffff]int)(unsafe.Pointer(roles_mm.keys))
	roles_Values := (*[0xffff]*QVariant)(unsafe.Pointer(roles_mm.values))
	for i := 0; i < int(roles_mm.len); i++ {
		roles_entry_Key := (int)(roles_Keys[i])

		roles_mapval_goptr := newQVariant(roles_Values[i])
		roles_mapval_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		roles_entry_Value := *roles_mapval_goptr

		roles_ret[roles_entry_Key] = roles_entry_Value
	}
	slotval2 := roles_ret

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_SetItemData, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_ClearItemData(index *QModelIndex) bool {

	return (bool)(QFileSystemModel_virtualbase_ClearItemData(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QFileSystemModel) OnClearItemData(slot func(super func(index *QModelIndex) bool, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_ClearItemData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_ClearItemData
func miqt_exec_callback_QFileSystemModel_ClearItemData(self QFileSystemModel, cb intptr_t, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) bool, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_ClearItemData, slotval1)

	return (bool)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_CanDropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {

	return (bool)(QFileSystemModel_virtualbase_CanDropMimeData(unsafe.Pointer(this.h), data.cPointer(), (int)(action), (int)(row), (int)(column), parent.cPointer()))

}
func (this *QFileSystemModel) OnCanDropMimeData(slot func(super func(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool, data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_CanDropMimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_CanDropMimeData
func miqt_exec_callback_QFileSystemModel_CanDropMimeData(self QFileSystemModel, cb intptr_t, data *QMimeData, action int, row int, column int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool, data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMimeData(data)

	slotval2 := (DropAction)(action)

	slotval3 := (int)(row)

	slotval4 := (int)(column)

	slotval5 := newQModelIndex(parent)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_CanDropMimeData, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_SupportedDragActions() DropAction {

	return (DropAction)(QFileSystemModel_virtualbase_SupportedDragActions(unsafe.Pointer(this.h)))

}
func (this *QFileSystemModel) OnSupportedDragActions(slot func(super func() DropAction) DropAction) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_SupportedDragActions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_SupportedDragActions
func miqt_exec_callback_QFileSystemModel_SupportedDragActions(self QFileSystemModel, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() DropAction) DropAction)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_SupportedDragActions)

	return (int)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_InsertRows(row int, count int, parent *QModelIndex) bool {

	return (bool)(QFileSystemModel_virtualbase_InsertRows(unsafe.Pointer(this.h), (int)(row), (int)(count), parent.cPointer()))

}
func (this *QFileSystemModel) OnInsertRows(slot func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_InsertRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_InsertRows
func miqt_exec_callback_QFileSystemModel_InsertRows(self QFileSystemModel, cb intptr_t, row int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_InsertRows, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_InsertColumns(column int, count int, parent *QModelIndex) bool {

	return (bool)(QFileSystemModel_virtualbase_InsertColumns(unsafe.Pointer(this.h), (int)(column), (int)(count), parent.cPointer()))

}
func (this *QFileSystemModel) OnInsertColumns(slot func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_InsertColumns(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_InsertColumns
func miqt_exec_callback_QFileSystemModel_InsertColumns(self QFileSystemModel, cb intptr_t, column int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_InsertColumns, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_RemoveRows(row int, count int, parent *QModelIndex) bool {

	return (bool)(QFileSystemModel_virtualbase_RemoveRows(unsafe.Pointer(this.h), (int)(row), (int)(count), parent.cPointer()))

}
func (this *QFileSystemModel) OnRemoveRows(slot func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_RemoveRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_RemoveRows
func miqt_exec_callback_QFileSystemModel_RemoveRows(self QFileSystemModel, cb intptr_t, row int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int, count int, parent *QModelIndex) bool, row int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_RemoveRows, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_RemoveColumns(column int, count int, parent *QModelIndex) bool {

	return (bool)(QFileSystemModel_virtualbase_RemoveColumns(unsafe.Pointer(this.h), (int)(column), (int)(count), parent.cPointer()))

}
func (this *QFileSystemModel) OnRemoveColumns(slot func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_RemoveColumns(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_RemoveColumns
func miqt_exec_callback_QFileSystemModel_RemoveColumns(self QFileSystemModel, cb intptr_t, column int, count int, parent *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int, count int, parent *QModelIndex) bool, column int, count int, parent *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	slotval2 := (int)(count)

	slotval3 := newQModelIndex(parent)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_RemoveColumns, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_MoveRows(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {

	return (bool)(QFileSystemModel_virtualbase_MoveRows(unsafe.Pointer(this.h), sourceParent.cPointer(), (int)(sourceRow), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))

}
func (this *QFileSystemModel) OnMoveRows(slot func(super func(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool, sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_MoveRows(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_MoveRows
func miqt_exec_callback_QFileSystemModel_MoveRows(self QFileSystemModel, cb intptr_t, sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool, sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(sourceParent)

	slotval2 := (int)(sourceRow)

	slotval3 := (int)(count)

	slotval4 := newQModelIndex(destinationParent)

	slotval5 := (int)(destinationChild)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_MoveRows, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_MoveColumns(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {

	return (bool)(QFileSystemModel_virtualbase_MoveColumns(unsafe.Pointer(this.h), sourceParent.cPointer(), (int)(sourceColumn), (int)(count), destinationParent.cPointer(), (int)(destinationChild)))

}
func (this *QFileSystemModel) OnMoveColumns(slot func(super func(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool, sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_MoveColumns(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_MoveColumns
func miqt_exec_callback_QFileSystemModel_MoveColumns(self QFileSystemModel, cb intptr_t, sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool, sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(sourceParent)

	slotval2 := (int)(sourceColumn)

	slotval3 := (int)(count)

	slotval4 := newQModelIndex(destinationParent)

	slotval5 := (int)(destinationChild)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_MoveColumns, slotval1, slotval2, slotval3, slotval4, slotval5)

	return (bool)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_Buddy(index *QModelIndex) *QModelIndex {

	_goptr := newQModelIndex(QFileSystemModel_virtualbase_Buddy(unsafe.Pointer(this.h), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFileSystemModel) OnBuddy(slot func(super func(index *QModelIndex) *QModelIndex, index *QModelIndex) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_Buddy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_Buddy
func miqt_exec_callback_QFileSystemModel_Buddy(self QFileSystemModel, cb intptr_t, index *QModelIndex) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QModelIndex, index *QModelIndex) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_Buddy, slotval1)

	return virtualReturn.cPointer()

}

func (this *QFileSystemModel) callVirtualBase_Match(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex {

	var _ma struct_miqt_array = QFileSystemModel_virtualbase_Match(unsafe.Pointer(this.h), start.cPointer(), (int)(role), value.cPointer(), (int)(hits), (int)(flags))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret

}
func (this *QFileSystemModel) OnMatch(slot func(super func(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex, start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_Match(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_Match
func miqt_exec_callback_QFileSystemModel_Match(self QFileSystemModel, cb intptr_t, start *QModelIndex, role int, value *QVariant, hits int, flags int) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex, start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(start)

	slotval2 := (int)(role)

	slotval3 := newQVariant(value)

	slotval4 := (int)(hits)

	slotval5 := (MatchFlag)(flags)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_Match, slotval1, slotval2, slotval3, slotval4, slotval5)
	virtualReturn_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = virtualReturn[i].cPointer()
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QFileSystemModel) callVirtualBase_Span(index *QModelIndex) *QSize {

	_goptr := newQSize(QFileSystemModel_virtualbase_Span(unsafe.Pointer(this.h), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFileSystemModel) OnSpan(slot func(super func(index *QModelIndex) *QSize, index *QModelIndex) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_Span(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_Span
func miqt_exec_callback_QFileSystemModel_Span(self QFileSystemModel, cb intptr_t, index *QModelIndex) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QSize, index *QModelIndex) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_Span, slotval1)

	return virtualReturn.cPointer()

}

func (this *QFileSystemModel) callVirtualBase_MultiData(index *QModelIndex, roleDataSpan QModelRoleDataSpan) {

	QFileSystemModel_virtualbase_MultiData(unsafe.Pointer(this.h), index.cPointer(), roleDataSpan.cPointer())

}
func (this *QFileSystemModel) OnMultiData(slot func(super func(index *QModelIndex, roleDataSpan QModelRoleDataSpan), index *QModelIndex, roleDataSpan QModelRoleDataSpan)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_MultiData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_MultiData
func miqt_exec_callback_QFileSystemModel_MultiData(self QFileSystemModel, cb intptr_t, index *QModelIndex, roleDataSpan *QModelRoleDataSpan) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, roleDataSpan QModelRoleDataSpan), index *QModelIndex, roleDataSpan QModelRoleDataSpan))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	roleDataSpan_goptr := newQModelRoleDataSpan(roleDataSpan)
	roleDataSpan_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval2 := *roleDataSpan_goptr

	gofunc((&QFileSystemModel{h: self}).callVirtualBase_MultiData, slotval1, slotval2)

}

func (this *QFileSystemModel) callVirtualBase_Submit() bool {

	return (bool)(QFileSystemModel_virtualbase_Submit(unsafe.Pointer(this.h)))

}
func (this *QFileSystemModel) OnSubmit(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_Submit(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_Submit
func miqt_exec_callback_QFileSystemModel_Submit(self QFileSystemModel, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFileSystemModel{h: self}).callVirtualBase_Submit)

	return (bool)(virtualReturn)

}

func (this *QFileSystemModel) callVirtualBase_Revert() {

	QFileSystemModel_virtualbase_Revert(unsafe.Pointer(this.h))

}
func (this *QFileSystemModel) OnRevert(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_Revert(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_Revert
func miqt_exec_callback_QFileSystemModel_Revert(self QFileSystemModel, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QFileSystemModel{h: self}).callVirtualBase_Revert)

}

func (this *QFileSystemModel) callVirtualBase_ResetInternalData() {

	QFileSystemModel_virtualbase_ResetInternalData(unsafe.Pointer(this.h))

}
func (this *QFileSystemModel) OnResetInternalData(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSystemModel_override_virtual_ResetInternalData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSystemModel_ResetInternalData
func miqt_exec_callback_QFileSystemModel_ResetInternalData(self QFileSystemModel, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QFileSystemModel{h: self}).callVirtualBase_ResetInternalData)

}
