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
	g := newQSaveFile(QSaveFile_new(name_ms))
	g.isSubclass = true
	return g
}

// NewQSaveFile2 constructs a new QSaveFile object.
func NewQSaveFile2() *QSaveFile {
	g := newQSaveFile(QSaveFile_new2())
	g.isSubclass = true
	return g
}

// NewQSaveFile3 constructs a new QSaveFile object.
func NewQSaveFile3(name string, parent *QObject) *QSaveFile {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	g := newQSaveFile(QSaveFile_new3(name_ms, parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQSaveFile4 constructs a new QSaveFile object.
func NewQSaveFile4(parent *QObject) *QSaveFile {
	g := newQSaveFile(QSaveFile_new4(parent.cPointer()))
	g.isSubclass = true
	return g
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

func (this *QSaveFile) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QSaveFile_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QSaveFile) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSaveFile_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSaveFile_MetaObject
func miqt_exec_callback_QSaveFile_MetaObject(self QSaveFile, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSaveFile{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QSaveFile) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSaveFile_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSaveFile) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSaveFile_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSaveFile_Metacast
func miqt_exec_callback_QSaveFile_Metacast(self QSaveFile, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QSaveFile{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
