package qt

import (
	"unsafe"
)

type QSaveFile struct {
	h          uintptr
	isSubclass bool
}

// NewQSaveFile constructs a new QSaveFile object.
func NewQSaveFile(name string) *QSaveFile {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))

	ret := newQSaveFile(QSaveFile_new(name_ms))
	ret.isSubclass = true
	return ret
}

// NewQSaveFile2 constructs a new QSaveFile object.
func NewQSaveFile2() *QSaveFile {

	ret := newQSaveFile(QSaveFile_new2())
	ret.isSubclass = true
	return ret
}

// NewQSaveFile3 constructs a new QSaveFile object.
func NewQSaveFile3(name string, parent *QObject) *QSaveFile {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))

	ret := newQSaveFile(QSaveFile_new3(name_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSaveFile4 constructs a new QSaveFile object.
func NewQSaveFile4(parent *QObject) *QSaveFile {

	ret := newQSaveFile(QSaveFile_new4(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSaveFile) MetaObject() *QMetaObject {
	return newQMetaObject(QSaveFile_MetaObject(this.h))
}

func (this *QSaveFile) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSaveFile_Metacast(this.h, param1_Cstring))
}

func QSaveFile_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSaveFile_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSaveFile) FileName() string {
	var _ms struct_miqt_string = QSaveFile_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSaveFile) SetFileName(name string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QSaveFile_SetFileName(this.h, name_ms)
}

func (this *QSaveFile) Open(flags OpenMode) bool {
	return (bool)(QSaveFile_Open(this.h, flags))
}

func (this *QSaveFile) Commit() bool {
	return (bool)(QSaveFile_Commit(this.h))
}

func (this *QSaveFile) CancelWriting() {
	QSaveFile_CancelWriting(this.h)
}

func (this *QSaveFile) SetDirectWriteFallback(enabled bool) {
	QSaveFile_SetDirectWriteFallback(this.h, (bool)(enabled))
}

func (this *QSaveFile) DirectWriteFallback() bool {
	return (bool)(QSaveFile_DirectWriteFallback(this.h))
}

func QSaveFile_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSaveFile_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSaveFile_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSaveFile_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSaveFile) callVirtualBase_FileName() string {

	var _ms struct_miqt_string = QSaveFile_virtualbase_FileName(unsafe.Pointer(this.h))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QSaveFile) OnFileName(slot func(super func() string) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSaveFile_override_virtual_FileName(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSaveFile_FileName
func miqt_exec_callback_QSaveFile_FileName(self QSaveFile, cb intptr_t) struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() string) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSaveFile{h: self}).callVirtualBase_FileName)
	virtualReturn_ms := struct_miqt_string{}
	virtualReturn_ms.data = CString(virtualReturn)
	virtualReturn_ms.len = size_t(len(virtualReturn))
	defer free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}

func (this *QSaveFile) callVirtualBase_Open(flags OpenMode) bool {

	return (bool)(QSaveFile_virtualbase_Open(unsafe.Pointer(this.h), flags))

}
func (this *QSaveFile) OnOpen(slot func(super func(flags OpenMode) bool, flags OpenMode) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSaveFile_override_virtual_Open(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSaveFile_Open
func miqt_exec_callback_QSaveFile_Open(self QSaveFile, cb intptr_t, flags OpenMode) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(flags OpenMode) bool, flags OpenMode) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QSaveFile{h: self}).callVirtualBase_Open, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSaveFile) callVirtualBase_WriteData(data string, lenVal int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QSaveFile_virtualbase_WriteData(unsafe.Pointer(this.h), data_Cstring, (longlong)(lenVal)))

}
func (this *QSaveFile) OnWriteData(slot func(super func(data string, lenVal int64) int64, data string, lenVal int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSaveFile_override_virtual_WriteData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSaveFile_WriteData
func miqt_exec_callback_QSaveFile_WriteData(self QSaveFile, cb intptr_t, data *const_char, lenVal longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, lenVal int64) int64, data string, lenVal int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(lenVal)

	virtualReturn := gofunc((&QSaveFile{h: self}).callVirtualBase_WriteData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}

func (this *QSaveFile) callVirtualBase_IsSequential() bool {

	return (bool)(QSaveFile_virtualbase_IsSequential(unsafe.Pointer(this.h)))

}
func (this *QSaveFile) OnIsSequential(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSaveFile_override_virtual_IsSequential(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSaveFile_IsSequential
func miqt_exec_callback_QSaveFile_IsSequential(self QSaveFile, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSaveFile{h: self}).callVirtualBase_IsSequential)

	return (bool)(virtualReturn)

}

func (this *QSaveFile) callVirtualBase_Pos() int64 {

	return (int64)(QSaveFile_virtualbase_Pos(unsafe.Pointer(this.h)))

}
func (this *QSaveFile) OnPos(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSaveFile_override_virtual_Pos(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSaveFile_Pos
func miqt_exec_callback_QSaveFile_Pos(self QSaveFile, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSaveFile{h: self}).callVirtualBase_Pos)

	return (longlong)(virtualReturn)

}

func (this *QSaveFile) callVirtualBase_Seek(offset int64) bool {

	return (bool)(QSaveFile_virtualbase_Seek(unsafe.Pointer(this.h), (longlong)(offset)))

}
func (this *QSaveFile) OnSeek(slot func(super func(offset int64) bool, offset int64) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSaveFile_override_virtual_Seek(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSaveFile_Seek
func miqt_exec_callback_QSaveFile_Seek(self QSaveFile, cb intptr_t, offset longlong) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset int64) bool, offset int64) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(offset)

	virtualReturn := gofunc((&QSaveFile{h: self}).callVirtualBase_Seek, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSaveFile) callVirtualBase_AtEnd() bool {

	return (bool)(QSaveFile_virtualbase_AtEnd(unsafe.Pointer(this.h)))

}
func (this *QSaveFile) OnAtEnd(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSaveFile_override_virtual_AtEnd(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSaveFile_AtEnd
func miqt_exec_callback_QSaveFile_AtEnd(self QSaveFile, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSaveFile{h: self}).callVirtualBase_AtEnd)

	return (bool)(virtualReturn)

}

func (this *QSaveFile) callVirtualBase_Size() int64 {

	return (int64)(QSaveFile_virtualbase_Size(unsafe.Pointer(this.h)))

}
func (this *QSaveFile) OnSize(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSaveFile_override_virtual_Size(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSaveFile_Size
func miqt_exec_callback_QSaveFile_Size(self QSaveFile, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSaveFile{h: self}).callVirtualBase_Size)

	return (longlong)(virtualReturn)

}

func (this *QSaveFile) callVirtualBase_Resize(sz int64) bool {

	return (bool)(QSaveFile_virtualbase_Resize(unsafe.Pointer(this.h), (longlong)(sz)))

}
func (this *QSaveFile) OnResize(slot func(super func(sz int64) bool, sz int64) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSaveFile_override_virtual_Resize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSaveFile_Resize
func miqt_exec_callback_QSaveFile_Resize(self QSaveFile, cb intptr_t, sz longlong) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(sz int64) bool, sz int64) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int64)(sz)

	virtualReturn := gofunc((&QSaveFile{h: self}).callVirtualBase_Resize, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSaveFile) callVirtualBase_Permissions() Permissions {

	xxxxxxxxx
}
func (this *QSaveFile) OnPermissions(slot func(super func() Permissions) Permissions) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSaveFile_override_virtual_Permissions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSaveFile_Permissions
func miqt_exec_callback_QSaveFile_Permissions(self QSaveFile, cb intptr_t) Permissions {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() Permissions) Permissions)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSaveFile{h: self}).callVirtualBase_Permissions)

	return virtualReturn

}

func (this *QSaveFile) callVirtualBase_SetPermissions(permissionSpec Permissions) bool {

	return (bool)(QSaveFile_virtualbase_SetPermissions(unsafe.Pointer(this.h), permissionSpec))

}
func (this *QSaveFile) OnSetPermissions(slot func(super func(permissionSpec Permissions) bool, permissionSpec Permissions) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSaveFile_override_virtual_SetPermissions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSaveFile_SetPermissions
func miqt_exec_callback_QSaveFile_SetPermissions(self QSaveFile, cb intptr_t, permissionSpec Permissions) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(permissionSpec Permissions) bool, permissionSpec Permissions) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QSaveFile{h: self}).callVirtualBase_SetPermissions, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSaveFile) callVirtualBase_ReadData(data string, maxlen int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QSaveFile_virtualbase_ReadData(unsafe.Pointer(this.h), data_Cstring, (longlong)(maxlen)))

}
func (this *QSaveFile) OnReadData(slot func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSaveFile_override_virtual_ReadData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSaveFile_ReadData
func miqt_exec_callback_QSaveFile_ReadData(self QSaveFile, cb intptr_t, data *char, maxlen longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(maxlen)

	virtualReturn := gofunc((&QSaveFile{h: self}).callVirtualBase_ReadData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}

func (this *QSaveFile) callVirtualBase_ReadLineData(data string, maxlen int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	return (int64)(QSaveFile_virtualbase_ReadLineData(unsafe.Pointer(this.h), data_Cstring, (longlong)(maxlen)))

}
func (this *QSaveFile) OnReadLineData(slot func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSaveFile_override_virtual_ReadLineData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSaveFile_ReadLineData
func miqt_exec_callback_QSaveFile_ReadLineData(self QSaveFile, cb intptr_t, data *char, maxlen longlong) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, maxlen int64) int64, data string, maxlen int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (int64)(maxlen)

	virtualReturn := gofunc((&QSaveFile{h: self}).callVirtualBase_ReadLineData, slotval1, slotval2)

	return (longlong)(virtualReturn)

}
