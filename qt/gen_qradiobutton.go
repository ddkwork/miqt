package qt

import (
	"unsafe"
)

type QRadioButton struct {
	h          uintptr
	isSubclass bool
}

// NewQRadioButton constructs a new QRadioButton object.
func NewQRadioButton(parent *QWidget) *QRadioButton {
	g := newQRadioButton(QRadioButton_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQRadioButton2 constructs a new QRadioButton object.
func NewQRadioButton2() *QRadioButton {
	g := newQRadioButton(QRadioButton_new2())
	g.isSubclass = true
	return g
}

// NewQRadioButton3 constructs a new QRadioButton object.
func NewQRadioButton3(text string) *QRadioButton {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQRadioButton(QRadioButton_new3(text_ms))
	g.isSubclass = true
	return g
}

// NewQRadioButton4 constructs a new QRadioButton object.
func NewQRadioButton4(text string, parent *QWidget) *QRadioButton {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQRadioButton(QRadioButton_new4(text_ms, parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QRadioButton) MetaObject() *QMetaObject {
	return newQMetaObject(QRadioButton_MetaObject(this.h))
}

func (this *QRadioButton) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QRadioButton_Metacast(this.h, param1_Cstring))
}

func QRadioButton_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QRadioButton_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRadioButton) SizeHint() *QSize {
	_goptr := newQSize(QRadioButton_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRadioButton) MinimumSizeHint() *QSize {
	_goptr := newQSize(QRadioButton_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QRadioButton_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QRadioButton_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QRadioButton_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QRadioButton_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRadioButton) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QRadioButton_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QRadioButton) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRadioButton_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRadioButton_MetaObject
func miqt_exec_callback_QRadioButton_MetaObject(self QRadioButton, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRadioButton{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QRadioButton) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QRadioButton_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QRadioButton) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRadioButton_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRadioButton_Metacast
func miqt_exec_callback_QRadioButton_Metacast(self QRadioButton, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QRadioButton{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
