package qt

import (
	"unsafe"
)

type QTemporaryFile struct {
	h          uintptr
	isSubclass bool
}

// NewQTemporaryFile constructs a new QTemporaryFile object.
func NewQTemporaryFile() *QTemporaryFile {

	ret := newQTemporaryFile(QTemporaryFile_new())
	ret.isSubclass = true
	return ret
}

// NewQTemporaryFile2 constructs a new QTemporaryFile object.
func NewQTemporaryFile2(templateName string) *QTemporaryFile {
	templateName_ms := struct_miqt_string{}
	templateName_ms.data = CString(templateName)
	templateName_ms.len = size_t(len(templateName))
	defer free(unsafe.Pointer(templateName_ms.data))

	ret := newQTemporaryFile(QTemporaryFile_new2(templateName_ms))
	ret.isSubclass = true
	return ret
}

// NewQTemporaryFile3 constructs a new QTemporaryFile object.
func NewQTemporaryFile3(parent *QObject) *QTemporaryFile {

	ret := newQTemporaryFile(QTemporaryFile_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTemporaryFile4 constructs a new QTemporaryFile object.
func NewQTemporaryFile4(templateName string, parent *QObject) *QTemporaryFile {
	templateName_ms := struct_miqt_string{}
	templateName_ms.data = CString(templateName)
	templateName_ms.len = size_t(len(templateName))
	defer free(unsafe.Pointer(templateName_ms.data))

	ret := newQTemporaryFile(QTemporaryFile_new4(templateName_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTemporaryFile) MetaObject() *QMetaObject {
	return newQMetaObject(QTemporaryFile_MetaObject(this.h))
}

func (this *QTemporaryFile) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTemporaryFile_Metacast(this.h, param1_Cstring))
}

func QTemporaryFile_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTemporaryFile_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTemporaryFile) AutoRemove() bool {
	return (bool)(QTemporaryFile_AutoRemove(this.h))
}

func (this *QTemporaryFile) SetAutoRemove(b bool) {
	QTemporaryFile_SetAutoRemove(this.h, (bool)(b))
}

func (this *QTemporaryFile) Open() bool {
	return (bool)(QTemporaryFile_Open(this.h))
}

func (this *QTemporaryFile) FileName() string {
	var _ms struct_miqt_string = QTemporaryFile_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTemporaryFile) FileTemplate() string {
	var _ms struct_miqt_string = QTemporaryFile_FileTemplate(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTemporaryFile) SetFileTemplate(name string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QTemporaryFile_SetFileTemplate(this.h, name_ms)
}

func (this *QTemporaryFile) Rename(newName string) bool {
	newName_ms := struct_miqt_string{}
	newName_ms.data = CString(newName)
	newName_ms.len = size_t(len(newName))
	defer free(unsafe.Pointer(newName_ms.data))
	return (bool)(QTemporaryFile_Rename(this.h, newName_ms))
}

func QTemporaryFile_CreateNativeFile(fileName string) *QTemporaryFile {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	return newQTemporaryFile(QTemporaryFile_CreateNativeFile(fileName_ms))
}

func QTemporaryFile_CreateNativeFileWithFile(file *QFile) *QTemporaryFile {
	return newQTemporaryFile(QTemporaryFile_CreateNativeFileWithFile(file.cPointer()))
}

func QTemporaryFile_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTemporaryFile_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTemporaryFile_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTemporaryFile_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTemporaryFile) callVirtualBase_FileName() string {

	var _ms struct_miqt_string = QTemporaryFile_virtualbase_FileName(unsafe.Pointer(this.h))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QTemporaryFile) OnFileName(slot func(super func() string) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTemporaryFile_override_virtual_FileName(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTemporaryFile_FileName
func miqt_exec_callback_QTemporaryFile_FileName(self QTemporaryFile, cb intptr_t) struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() string) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTemporaryFile{h: self}).callVirtualBase_FileName)
	virtualReturn_ms := struct_miqt_string{}
	virtualReturn_ms.data = CString(virtualReturn)
	virtualReturn_ms.len = size_t(len(virtualReturn))
	defer free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}

func (this *QTemporaryFile) callVirtualBase_OpenWithFlags(flags OpenMode) bool {

	return (bool)(QTemporaryFile_virtualbase_OpenWithFlags(unsafe.Pointer(this.h), flags))

}
func (this *QTemporaryFile) OnOpenWithFlags(slot func(super func(flags OpenMode) bool, flags OpenMode) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTemporaryFile_override_virtual_OpenWithFlags(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTemporaryFile_OpenWithFlags
func miqt_exec_callback_QTemporaryFile_OpenWithFlags(self QTemporaryFile, cb intptr_t, flags OpenMode) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(flags OpenMode) bool, flags OpenMode) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QTemporaryFile{h: self}).callVirtualBase_OpenWithFlags, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTemporaryFile) callVirtualBase_Size() int64 {

	return (int64)(QTemporaryFile_virtualbase_Size(unsafe.Pointer(this.h)))

}
func (this *QTemporaryFile) OnSize(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTemporaryFile_override_virtual_Size(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTemporaryFile_Size
func miqt_exec_callback_QTemporaryFile_Size(self QTemporaryFile, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTemporaryFile{h: self}).callVirtualBase_Size)

	return (longlong)(virtualReturn)

}

func (this *QTemporaryFile) callVirtualBase_Resize(sz int64) bool {

	return (bool)(QTemporaryFile_virtualbase_Resize(unsafe.Pointer(this.h), (longlong)(sz)))

}
func (this *QTemporaryFile) OnResize(slot func(super func(sz int64) bool, sz int64) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTemporaryFile_override_virtual_Resize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTemporaryFile_Resize
func miqt_exec_callback_QTemporaryFile_Resize(self QTemporaryFile, cb intptr_t, sz longlong) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(sz int64) bool, sz int64) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(sz)

	virtualReturn := gofunc((&QTemporaryFile{h: self}).callVirtualBase_Resize, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTemporaryFile) callVirtualBase_Permissions() Permissions {

	xxxxxxxxx
}
func (this *QTemporaryFile) OnPermissions(slot func(super func() Permissions) Permissions) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTemporaryFile_override_virtual_Permissions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTemporaryFile_Permissions
func miqt_exec_callback_QTemporaryFile_Permissions(self QTemporaryFile, cb intptr_t) Permissions {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() Permissions) Permissions)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTemporaryFile{h: self}).callVirtualBase_Permissions)

	return virtualReturn

}

func (this *QTemporaryFile) callVirtualBase_SetPermissions(permissionSpec Permissions) bool {

	return (bool)(QTemporaryFile_virtualbase_SetPermissions(unsafe.Pointer(this.h), permissionSpec))

}
func (this *QTemporaryFile) OnSetPermissions(slot func(super func(permissionSpec Permissions) bool, permissionSpec Permissions) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTemporaryFile_override_virtual_SetPermissions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTemporaryFile_SetPermissions
func miqt_exec_callback_QTemporaryFile_SetPermissions(self QTemporaryFile, cb intptr_t, permissionSpec Permissions) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(permissionSpec Permissions) bool, permissionSpec Permissions) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QTemporaryFile{h: self}).callVirtualBase_SetPermissions, slotval1)

	return (bool)(virtualReturn)

}
