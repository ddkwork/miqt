package qt

import (
	"unsafe"
)

type QProgressBar__Direction int

const (
	QProgressBar__TopToBottom QProgressBar__Direction = 0
	QProgressBar__BottomToTop QProgressBar__Direction = 1
)

type QProgressBar struct {
	h          uintptr
	isSubclass bool
}

// NewQProgressBar constructs a new QProgressBar object.
func NewQProgressBar(parent *QWidget) *QProgressBar {
	g := newQProgressBar(QProgressBar_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQProgressBar2 constructs a new QProgressBar object.
func NewQProgressBar2() *QProgressBar {
	g := newQProgressBar(QProgressBar_new2())
	g.isSubclass = true
	return g
}

func (this *QProgressBar) MetaObject() *QMetaObject {
	return newQMetaObject(QProgressBar_MetaObject(this.h))
}

func (this *QProgressBar) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QProgressBar_Metacast(this.h, param1_Cstring))
}

func QProgressBar_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QProgressBar_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProgressBar) Minimum() int {
	return (int)(QProgressBar_Minimum(this.h))
}

func (this *QProgressBar) Maximum() int {
	return (int)(QProgressBar_Maximum(this.h))
}

func (this *QProgressBar) Value() int {
	return (int)(QProgressBar_Value(this.h))
}

func (this *QProgressBar) Text() string {
	var _ms struct_miqt_string = QProgressBar_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProgressBar) SetTextVisible(visible bool) {
	QProgressBar_SetTextVisible(this.h, (bool)(visible))
}

func (this *QProgressBar) IsTextVisible() bool {
	return (bool)(QProgressBar_IsTextVisible(this.h))
}

func (this *QProgressBar) Alignment() AlignmentFlag {
	return (AlignmentFlag)(QProgressBar_Alignment(this.h))
}

func (this *QProgressBar) SetAlignment(alignment AlignmentFlag) {
	QProgressBar_SetAlignment(this.h, (int)(alignment))
}

func (this *QProgressBar) SizeHint() *QSize {
	_goptr := newQSize(QProgressBar_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProgressBar) MinimumSizeHint() *QSize {
	_goptr := newQSize(QProgressBar_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProgressBar) Orientation() Orientation {
	return (Orientation)(QProgressBar_Orientation(this.h))
}

func (this *QProgressBar) SetInvertedAppearance(invert bool) {
	QProgressBar_SetInvertedAppearance(this.h, (bool)(invert))
}

func (this *QProgressBar) InvertedAppearance() bool {
	return (bool)(QProgressBar_InvertedAppearance(this.h))
}

func (this *QProgressBar) SetTextDirection(textDirection QProgressBar__Direction) {
	QProgressBar_SetTextDirection(this.h, (int)(textDirection))
}

func (this *QProgressBar) TextDirection() QProgressBar__Direction {
	return (QProgressBar__Direction)(QProgressBar_TextDirection(this.h))
}

func (this *QProgressBar) SetFormat(format string) {
	format_ms := struct_miqt_string{}
	format_ms.data = CString(format)
	format_ms.len = size_t(len(format))
	defer free(unsafe.Pointer(format_ms.data))
	QProgressBar_SetFormat(this.h, format_ms)
}

func (this *QProgressBar) ResetFormat() {
	QProgressBar_ResetFormat(this.h)
}

func (this *QProgressBar) Format() string {
	var _ms struct_miqt_string = QProgressBar_Format(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProgressBar) Reset() {
	QProgressBar_Reset(this.h)
}

func (this *QProgressBar) SetRange(minimum int, maximum int) {
	QProgressBar_SetRange(this.h, (int)(minimum), (int)(maximum))
}

func (this *QProgressBar) SetMinimum(minimum int) {
	QProgressBar_SetMinimum(this.h, (int)(minimum))
}

func (this *QProgressBar) SetMaximum(maximum int) {
	QProgressBar_SetMaximum(this.h, (int)(maximum))
}

func (this *QProgressBar) SetValue(value int) {
	QProgressBar_SetValue(this.h, (int)(value))
}

func (this *QProgressBar) SetOrientation(orientation Orientation) {
	QProgressBar_SetOrientation(this.h, (int)(orientation))
}

func (this *QProgressBar) ValueChanged(value int) {
	QProgressBar_ValueChanged(this.h, (int)(value))
}

func (this *QProgressBar) OnValueChanged(slot func(value int)) {
	QProgressBar_connect_ValueChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_ValueChanged
func miqt_exec_callback_QProgressBar_ValueChanged(cb intptr_t, value int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(value int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(value)

	gofunc(slotval1)
}

func QProgressBar_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QProgressBar_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QProgressBar_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QProgressBar_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProgressBar) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QProgressBar_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QProgressBar) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_MetaObject
func miqt_exec_callback_QProgressBar_MetaObject(self QProgressBar, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QProgressBar{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QProgressBar) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QProgressBar_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QProgressBar) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressBar_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressBar_Metacast
func miqt_exec_callback_QProgressBar_Metacast(self QProgressBar, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QProgressBar{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
