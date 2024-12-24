package qt

import (
	"unsafe"
)

type QCheckBox struct {
	h          uintptr
	isSubclass bool
}

// NewQCheckBox constructs a new QCheckBox object.
func NewQCheckBox(parent *QWidget) *QCheckBox {
	ret := newQCheckBox(QCheckBox_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQCheckBox2 constructs a new QCheckBox object.
func NewQCheckBox2() *QCheckBox {
	ret := newQCheckBox(QCheckBox_new2())
	ret.isSubclass = true
	return ret
}

// NewQCheckBox3 constructs a new QCheckBox object.
func NewQCheckBox3(text string) *QCheckBox {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQCheckBox(QCheckBox_new3(text_ms))
	ret.isSubclass = true
	return ret
}

// NewQCheckBox4 constructs a new QCheckBox object.
func NewQCheckBox4(text string, parent *QWidget) *QCheckBox {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQCheckBox(QCheckBox_new4(text_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QCheckBox) MetaObject() *QMetaObject {
	return newQMetaObject(QCheckBox_MetaObject(this.h))
}

func (this *QCheckBox) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QCheckBox_Metacast(this.h, param1_Cstring))
}

func QCheckBox_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QCheckBox_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCheckBox) SizeHint() *QSize {
	_goptr := newQSize(QCheckBox_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCheckBox) MinimumSizeHint() *QSize {
	_goptr := newQSize(QCheckBox_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCheckBox) SetTristate() {
	QCheckBox_SetTristate(this.h)
}

func (this *QCheckBox) IsTristate() bool {
	return (bool)(QCheckBox_IsTristate(this.h))
}

func (this *QCheckBox) CheckState() CheckState {
	return (CheckState)(QCheckBox_CheckState(this.h))
}

func (this *QCheckBox) SetCheckState(state CheckState) {
	QCheckBox_SetCheckState(this.h, (int)(state))
}

func (this *QCheckBox) StateChanged(param1 int) {
	QCheckBox_StateChanged(this.h, (int)(param1))
}

func (this *QCheckBox) OnStateChanged(slot func(param1 int)) {
	QCheckBox_connect_StateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_StateChanged
func miqt_exec_callback_QCheckBox_StateChanged(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc(slotval1)
}

func (this *QCheckBox) CheckStateChanged(param1 CheckState) {
	QCheckBox_CheckStateChanged(this.h, (int)(param1))
}

func (this *QCheckBox) OnCheckStateChanged(slot func(param1 CheckState)) {
	QCheckBox_connect_CheckStateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_CheckStateChanged
func miqt_exec_callback_QCheckBox_CheckStateChanged(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 CheckState))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (CheckState)(param1)

	gofunc(slotval1)
}

func QCheckBox_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QCheckBox_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCheckBox_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QCheckBox_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCheckBox) SetTristate1(y bool) {
	QCheckBox_SetTristate1(this.h, (bool)(y))
}

func (this *QCheckBox) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QCheckBox_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QCheckBox) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_MetaObject
func miqt_exec_callback_QCheckBox_MetaObject(self QCheckBox, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCheckBox{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QCheckBox) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QCheckBox_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QCheckBox) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCheckBox_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCheckBox_Metacast
func miqt_exec_callback_QCheckBox_Metacast(self QCheckBox, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QCheckBox{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
