package qt

import (
	"unsafe"
)

type QNtfsPermissionCheckGuard struct {
	h          uintptr
	isSubclass bool
}

// NewQNtfsPermissionCheckGuard constructs a new QNtfsPermissionCheckGuard object.
func NewQNtfsPermissionCheckGuard() *QNtfsPermissionCheckGuard {

	ret := newQNtfsPermissionCheckGuard(QNtfsPermissionCheckGuard_new())
	ret.isSubclass = true
	return ret
}

type QFile struct {
	h          uintptr
	isSubclass bool
}

// NewQFile constructs a new QFile object.
func NewQFile() *QFile {

	ret := newQFile(QFile_new())
	ret.isSubclass = true
	return ret
}

// NewQFile2 constructs a new QFile object.
func NewQFile2(name string) *QFile {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))

	ret := newQFile(QFile_new2(name_ms))
	ret.isSubclass = true
	return ret
}

// NewQFile3 constructs a new QFile object.
func NewQFile3(parent *QObject) *QFile {

	ret := newQFile(QFile_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFile4 constructs a new QFile object.
func NewQFile4(name string, parent *QObject) *QFile {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))

	ret := newQFile(QFile_new4(name_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QFile) MetaObject() *QMetaObject {
	return newQMetaObject(QFile_MetaObject(this.h))
}

func (this *QFile) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QFile_Metacast(this.h, param1_Cstring))
}

func QFile_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QFile_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFile) FileName() string {
	var _ms struct_miqt_string = QFile_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFile) SetFileName(name string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QFile_SetFileName(this.h, name_ms)
}

func QFile_EncodeName(fileName string) []byte {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	var _bytearray struct_miqt_string = QFile_EncodeName(fileName_ms)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QFile_DecodeName(localFileName []byte) string {
	localFileName_alias := struct_miqt_string{}
	localFileName_alias.data = (char)(unsafe.Pointer(&localFileName[0]))
	localFileName_alias.len = size_t(len(localFileName))
	var _ms struct_miqt_string = QFile_DecodeName(localFileName_alias)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFile_DecodeNameWithLocalFileName(localFileName string) string {
	localFileName_Cstring := CString(localFileName)
	defer free(unsafe.Pointer(localFileName_Cstring))
	var _ms struct_miqt_string = QFile_DecodeNameWithLocalFileName(localFileName_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFile) Exists() bool {
	return (bool)(QFile_Exists(this.h))
}

func QFile_ExistsWithFileName(fileName string) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	return (bool)(QFile_ExistsWithFileName(fileName_ms))
}

func (this *QFile) SymLinkTarget() string {
	var _ms struct_miqt_string = QFile_SymLinkTarget(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFile_SymLinkTargetWithFileName(fileName string) string {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	var _ms struct_miqt_string = QFile_SymLinkTargetWithFileName(fileName_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFile) Remove() bool {
	return (bool)(QFile_Remove(this.h))
}

func QFile_RemoveWithFileName(fileName string) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	return (bool)(QFile_RemoveWithFileName(fileName_ms))
}

func QFile_SupportsMoveToTrash() bool {
	return (bool)(QFile_SupportsMoveToTrash())
}

func (this *QFile) MoveToTrash() bool {
	return (bool)(QFile_MoveToTrash(this.h))
}

func QFile_MoveToTrashWithFileName(fileName string) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	return (bool)(QFile_MoveToTrashWithFileName(fileName_ms))
}

func (this *QFile) Rename(newName string) bool {
	newName_ms := struct_miqt_string{}
	newName_ms.data = CString(newName)
	newName_ms.len = size_t(len(newName))
	defer free(unsafe.Pointer(newName_ms.data))
	return (bool)(QFile_Rename(this.h, newName_ms))
}

func QFile_Rename2(oldName string, newName string) bool {
	oldName_ms := struct_miqt_string{}
	oldName_ms.data = CString(oldName)
	oldName_ms.len = size_t(len(oldName))
	defer free(unsafe.Pointer(oldName_ms.data))
	newName_ms := struct_miqt_string{}
	newName_ms.data = CString(newName)
	newName_ms.len = size_t(len(newName))
	defer free(unsafe.Pointer(newName_ms.data))
	return (bool)(QFile_Rename2(oldName_ms, newName_ms))
}

func (this *QFile) Link(newName string) bool {
	newName_ms := struct_miqt_string{}
	newName_ms.data = CString(newName)
	newName_ms.len = size_t(len(newName))
	defer free(unsafe.Pointer(newName_ms.data))
	return (bool)(QFile_Link(this.h, newName_ms))
}

func QFile_Link2(fileName string, newName string) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	newName_ms := struct_miqt_string{}
	newName_ms.data = CString(newName)
	newName_ms.len = size_t(len(newName))
	defer free(unsafe.Pointer(newName_ms.data))
	return (bool)(QFile_Link2(fileName_ms, newName_ms))
}

func (this *QFile) Copy(newName string) bool {
	newName_ms := struct_miqt_string{}
	newName_ms.data = CString(newName)
	newName_ms.len = size_t(len(newName))
	defer free(unsafe.Pointer(newName_ms.data))
	return (bool)(QFile_Copy(this.h, newName_ms))
}

func QFile_Copy2(fileName string, newName string) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	newName_ms := struct_miqt_string{}
	newName_ms.data = CString(newName)
	newName_ms.len = size_t(len(newName))
	defer free(unsafe.Pointer(newName_ms.data))
	return (bool)(QFile_Copy2(fileName_ms, newName_ms))
}

func (this *QFile) Open(flags OpenMode) bool {
	return (bool)(QFile_Open(this.h, flags))
}

func (this *QFile) Open2(flags OpenMode, permissions Permissions) bool {
	return (bool)(QFile_Open2(this.h, flags, permissions))
}

func (this *QFile) Open4(fd int, ioFlags OpenMode) bool {
	return (bool)(QFile_Open4(this.h, (int)(fd), ioFlags))
}

func (this *QFile) Size() int64 {
	return (int64)(QFile_Size(this.h))
}

func (this *QFile) Resize(sz int64) bool {
	return (bool)(QFile_Resize(this.h, (longlong)(sz)))
}

func QFile_Resize2(filename string, sz int64) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	return (bool)(QFile_Resize2(filename_ms, (longlong)(sz)))
}

func (this *QFile) Permissions() Permissions {
	xxxxxxxxx
}

func QFile_PermissionsWithFilename(filename string) Permissions {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	xxxxxxxxx
}

func (this *QFile) SetPermissions(permissionSpec Permissions) bool {
	return (bool)(QFile_SetPermissions(this.h, permissionSpec))
}

func QFile_SetPermissions2(filename string, permissionSpec Permissions) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	return (bool)(QFile_SetPermissions2(filename_ms, permissionSpec))
}

func QFile_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFile_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFile_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFile_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFile) Open33(fd int, ioFlags OpenMode, handleFlags FileHandleFlags) bool {
	return (bool)(QFile_Open33(this.h, (int)(fd), ioFlags, handleFlags))
}

func (this *QFile) callVirtualBase_FileName() string {

	var _ms struct_miqt_string = QFile_virtualbase_FileName(unsafe.Pointer(this.h))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QFile) OnFileName(slot func(super func() string) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFile_override_virtual_FileName(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFile_FileName
func miqt_exec_callback_QFile_FileName(self QFile, cb intptr_t) struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() string) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFile{h: self}).callVirtualBase_FileName)
	virtualReturn_ms := struct_miqt_string{}
	virtualReturn_ms.data = CString(virtualReturn)
	virtualReturn_ms.len = size_t(len(virtualReturn))
	defer free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}

func (this *QFile) callVirtualBase_Open(flags OpenMode) bool {

	return (bool)(QFile_virtualbase_Open(unsafe.Pointer(this.h), flags))

}
func (this *QFile) OnOpen(slot func(super func(flags OpenMode) bool, flags OpenMode) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFile_override_virtual_Open(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFile_Open
func miqt_exec_callback_QFile_Open(self QFile, cb intptr_t, flags OpenMode) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(flags OpenMode) bool, flags OpenMode) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QFile{h: self}).callVirtualBase_Open, slotval1)

	return (bool)(virtualReturn)

}

func (this *QFile) callVirtualBase_Size() int64 {

	return (int64)(QFile_virtualbase_Size(unsafe.Pointer(this.h)))

}
func (this *QFile) OnSize(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFile_override_virtual_Size(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFile_Size
func miqt_exec_callback_QFile_Size(self QFile, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFile{h: self}).callVirtualBase_Size)

	return (longlong)(virtualReturn)

}

func (this *QFile) callVirtualBase_Resize(sz int64) bool {

	return (bool)(QFile_virtualbase_Resize(unsafe.Pointer(this.h), (longlong)(sz)))

}
func (this *QFile) OnResize(slot func(super func(sz int64) bool, sz int64) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFile_override_virtual_Resize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFile_Resize
func miqt_exec_callback_QFile_Resize(self QFile, cb intptr_t, sz longlong) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(sz int64) bool, sz int64) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(sz)

	virtualReturn := gofunc((&QFile{h: self}).callVirtualBase_Resize, slotval1)

	return (bool)(virtualReturn)

}

func (this *QFile) callVirtualBase_Permissions() Permissions {

	xxxxxxxxx
}
func (this *QFile) OnPermissions(slot func(super func() Permissions) Permissions) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFile_override_virtual_Permissions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFile_Permissions
func miqt_exec_callback_QFile_Permissions(self QFile, cb intptr_t) Permissions {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() Permissions) Permissions)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFile{h: self}).callVirtualBase_Permissions)

	return virtualReturn

}

func (this *QFile) callVirtualBase_SetPermissions(permissionSpec Permissions) bool {

	return (bool)(QFile_virtualbase_SetPermissions(unsafe.Pointer(this.h), permissionSpec))

}
func (this *QFile) OnSetPermissions(slot func(super func(permissionSpec Permissions) bool, permissionSpec Permissions) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFile_override_virtual_SetPermissions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFile_SetPermissions
func miqt_exec_callback_QFile_SetPermissions(self QFile, cb intptr_t, permissionSpec Permissions) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(permissionSpec Permissions) bool, permissionSpec Permissions) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QFile{h: self}).callVirtualBase_SetPermissions, slotval1)

	return (bool)(virtualReturn)

}

func (this *QFile) callVirtualBase_Close() {

	QFile_virtualbase_Close(unsafe.Pointer(this.h))

}
func (this *QFile) OnClose(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFile_override_virtual_Close(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFile_Close
func miqt_exec_callback_QFile_Close(self QFile, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QFile{h: self}).callVirtualBase_Close)

}

func (this *QFile) callVirtualBase_IsSequential() bool {

	return (bool)(QFile_virtualbase_IsSequential(unsafe.Pointer(this.h)))

}
func (this *QFile) OnIsSequential(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFile_override_virtual_IsSequential(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFile_IsSequential
func miqt_exec_callback_QFile_IsSequential(self QFile, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFile{h: self}).callVirtualBase_IsSequential)

	return (bool)(virtualReturn)

}

func (this *QFile) callVirtualBase_Pos() int64 {

	return (int64)(QFile_virtualbase_Pos(unsafe.Pointer(this.h)))

}
func (this *QFile) OnPos(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFile_override_virtual_Pos(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFile_Pos
func miqt_exec_callback_QFile_Pos(self QFile, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFile{h: self}).callVirtualBase_Pos)

	return (longlong)(virtualReturn)

}

func (this *QFile) callVirtualBase_Seek(offset int64) bool {

	return (bool)(QFile_virtualbase_Seek(unsafe.Pointer(this.h), (longlong)(offset)))

}
func (this *QFile) OnSeek(slot func(super func(offset int64) bool, offset int64) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFile_override_virtual_Seek(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFile_Seek
func miqt_exec_callback_QFile_Seek(self QFile, cb intptr_t, offset longlong) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset int64) bool, offset int64) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(offset)

	virtualReturn := gofunc((&QFile{h: self}).callVirtualBase_Seek, slotval1)

	return (bool)(virtualReturn)

}

func (this *QFile) callVirtualBase_AtEnd() bool {

	return (bool)(QFile_virtualbase_AtEnd(unsafe.Pointer(this.h)))

}
func (this *QFile) OnAtEnd(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFile_override_virtual_AtEnd(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFile_AtEnd
func miqt_exec_callback_QFile_AtEnd(self QFile, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFile{h: self}).callVirtualBase_AtEnd)

	return (bool)(virtualReturn)

}

func (this *QFile) callVirtualBase_ReadData(data string, maxlen int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QFile_virtualbase_ReadData(unsafe.Pointer(this.h), data_Cstring, (longlong)(maxlen)))

}
func (this *QFile) OnReadData(slot func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFile_override_virtual_ReadData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFile_ReadData
func miqt_exec_callback_QFile_ReadData(self QFile, cb intptr_t, data *char, maxlen longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(maxlen)

	virtualReturn := gofunc((&QFile{h: self}).callVirtualBase_ReadData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}

func (this *QFile) callVirtualBase_WriteData(data string, lenVal int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QFile_virtualbase_WriteData(unsafe.Pointer(this.h), data_Cstring, (longlong)(lenVal)))

}
func (this *QFile) OnWriteData(slot func(super func(data string, lenVal int64) int64, data string, lenVal int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFile_override_virtual_WriteData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFile_WriteData
func miqt_exec_callback_QFile_WriteData(self QFile, cb intptr_t, data *const_char, lenVal longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, lenVal int64) int64, data string, lenVal int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(lenVal)

	virtualReturn := gofunc((&QFile{h: self}).callVirtualBase_WriteData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}

func (this *QFile) callVirtualBase_ReadLineData(data string, maxlen int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QFile_virtualbase_ReadLineData(unsafe.Pointer(this.h), data_Cstring, (longlong)(maxlen)))

}
func (this *QFile) OnReadLineData(slot func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFile_override_virtual_ReadLineData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFile_ReadLineData
func miqt_exec_callback_QFile_ReadLineData(self QFile, cb intptr_t, data *char, maxlen longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(maxlen)

	virtualReturn := gofunc((&QFile{h: self}).callVirtualBase_ReadLineData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}
