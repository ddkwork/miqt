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

func (this *QFile) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QFile_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QFile) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFile_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFile_MetaObject
func miqt_exec_callback_QFile_MetaObject(self QFile, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFile{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QFile) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QFile_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QFile) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFile_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFile_Metacast
func miqt_exec_callback_QFile_Metacast(self QFile, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QFile{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
