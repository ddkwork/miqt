package qt

import (
	"unsafe"
)

type QGroupBox struct {
	h          uintptr
	isSubclass bool
}

// NewQGroupBox constructs a new QGroupBox object.
func NewQGroupBox(parent *QWidget) *QGroupBox {
	g := newQGroupBox(QGroupBox_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGroupBox2 constructs a new QGroupBox object.
func NewQGroupBox2() *QGroupBox {
	g := newQGroupBox(QGroupBox_new2())
	g.isSubclass = true
	return g
}

// NewQGroupBox3 constructs a new QGroupBox object.
func NewQGroupBox3(title string) *QGroupBox {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	g := newQGroupBox(QGroupBox_new3(title_ms))
	g.isSubclass = true
	return g
}

// NewQGroupBox4 constructs a new QGroupBox object.
func NewQGroupBox4(title string, parent *QWidget) *QGroupBox {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	g := newQGroupBox(QGroupBox_new4(title_ms, parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QGroupBox) MetaObject() *QMetaObject {
	return newQMetaObject(QGroupBox_MetaObject(this.h))
}

func (this *QGroupBox) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGroupBox_Metacast(this.h, param1_Cstring))
}

func QGroupBox_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGroupBox_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGroupBox) Title() string {
	var _ms struct_miqt_string = QGroupBox_Title(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGroupBox) SetTitle(title string) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	QGroupBox_SetTitle(this.h, title_ms)
}

func (this *QGroupBox) Alignment() AlignmentFlag {
	return (AlignmentFlag)(QGroupBox_Alignment(this.h))
}

func (this *QGroupBox) SetAlignment(alignment int) {
	QGroupBox_SetAlignment(this.h, (int)(alignment))
}

func (this *QGroupBox) MinimumSizeHint() *QSize {
	_goptr := newQSize(QGroupBox_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGroupBox) IsFlat() bool {
	return (bool)(QGroupBox_IsFlat(this.h))
}

func (this *QGroupBox) SetFlat(flat bool) {
	QGroupBox_SetFlat(this.h, (bool)(flat))
}

func (this *QGroupBox) IsCheckable() bool {
	return (bool)(QGroupBox_IsCheckable(this.h))
}

func (this *QGroupBox) SetCheckable(checkable bool) {
	QGroupBox_SetCheckable(this.h, (bool)(checkable))
}

func (this *QGroupBox) IsChecked() bool {
	return (bool)(QGroupBox_IsChecked(this.h))
}

func (this *QGroupBox) SetChecked(checked bool) {
	QGroupBox_SetChecked(this.h, (bool)(checked))
}

func (this *QGroupBox) Clicked() {
	QGroupBox_Clicked(this.h)
}

func (this *QGroupBox) OnClicked(slot func()) {
	QGroupBox_connect_Clicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_Clicked
func miqt_exec_callback_QGroupBox_Clicked(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGroupBox) Toggled(param1 bool) {
	QGroupBox_Toggled(this.h, (bool)(param1))
}

func (this *QGroupBox) OnToggled(slot func(param1 bool)) {
	QGroupBox_connect_Toggled(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_Toggled
func miqt_exec_callback_QGroupBox_Toggled(cb intptr_t, param1 bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func QGroupBox_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGroupBox_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGroupBox_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGroupBox_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGroupBox) Clicked1(checked bool) {
	QGroupBox_Clicked1(this.h, (bool)(checked))
}

func (this *QGroupBox) OnClicked1(slot func(checked bool)) {
	QGroupBox_connect_Clicked1(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_Clicked1
func miqt_exec_callback_QGroupBox_Clicked1(cb intptr_t, checked bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(checked bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(checked)

	gofunc(slotval1)
}

func (this *QGroupBox) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QGroupBox_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QGroupBox) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_MetaObject
func miqt_exec_callback_QGroupBox_MetaObject(self QGroupBox, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGroupBox{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QGroupBox) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QGroupBox_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QGroupBox) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGroupBox_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGroupBox_Metacast
func miqt_exec_callback_QGroupBox_Metacast(self QGroupBox, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QGroupBox{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
