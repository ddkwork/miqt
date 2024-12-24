package qt

import (
	"unsafe"
)

type QScrollArea struct {
	h          uintptr
	isSubclass bool
}

// NewQScrollArea constructs a new QScrollArea object.
func NewQScrollArea(parent *QWidget) *QScrollArea {
	ret := newQScrollArea(QScrollArea_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQScrollArea2 constructs a new QScrollArea object.
func NewQScrollArea2() *QScrollArea {
	ret := newQScrollArea(QScrollArea_new2())
	ret.isSubclass = true
	return ret
}

func (this *QScrollArea) MetaObject() *QMetaObject {
	return newQMetaObject(QScrollArea_MetaObject(this.h))
}

func (this *QScrollArea) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QScrollArea_Metacast(this.h, param1_Cstring))
}

func QScrollArea_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QScrollArea_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QScrollArea) Widget() *QWidget {
	return newQWidget(QScrollArea_Widget(this.h))
}

func (this *QScrollArea) SetWidget(widget *QWidget) {
	QScrollArea_SetWidget(this.h, widget.cPointer())
}

func (this *QScrollArea) TakeWidget() *QWidget {
	return newQWidget(QScrollArea_TakeWidget(this.h))
}

func (this *QScrollArea) WidgetResizable() bool {
	return (bool)(QScrollArea_WidgetResizable(this.h))
}

func (this *QScrollArea) SetWidgetResizable(resizable bool) {
	QScrollArea_SetWidgetResizable(this.h, (bool)(resizable))
}

func (this *QScrollArea) SizeHint() *QSize {
	_goptr := newQSize(QScrollArea_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QScrollArea) FocusNextPrevChild(next bool) bool {
	return (bool)(QScrollArea_FocusNextPrevChild(this.h, (bool)(next)))
}

func (this *QScrollArea) Alignment() AlignmentFlag {
	return (AlignmentFlag)(QScrollArea_Alignment(this.h))
}

func (this *QScrollArea) SetAlignment(alignment AlignmentFlag) {
	QScrollArea_SetAlignment(this.h, (int)(alignment))
}

func (this *QScrollArea) EnsureVisible(x int, y int) {
	QScrollArea_EnsureVisible(this.h, (int)(x), (int)(y))
}

func (this *QScrollArea) EnsureWidgetVisible(childWidget *QWidget) {
	QScrollArea_EnsureWidgetVisible(this.h, childWidget.cPointer())
}

func QScrollArea_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QScrollArea_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QScrollArea_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QScrollArea_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QScrollArea) EnsureVisible3(x int, y int, xmargin int) {
	QScrollArea_EnsureVisible3(this.h, (int)(x), (int)(y), (int)(xmargin))
}

func (this *QScrollArea) EnsureVisible4(x int, y int, xmargin int, ymargin int) {
	QScrollArea_EnsureVisible4(this.h, (int)(x), (int)(y), (int)(xmargin), (int)(ymargin))
}

func (this *QScrollArea) EnsureWidgetVisible2(childWidget *QWidget, xmargin int) {
	QScrollArea_EnsureWidgetVisible2(this.h, childWidget.cPointer(), (int)(xmargin))
}

func (this *QScrollArea) EnsureWidgetVisible3(childWidget *QWidget, xmargin int, ymargin int) {
	QScrollArea_EnsureWidgetVisible3(this.h, childWidget.cPointer(), (int)(xmargin), (int)(ymargin))
}

func (this *QScrollArea) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QScrollArea_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QScrollArea) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_MetaObject
func miqt_exec_callback_QScrollArea_MetaObject(self QScrollArea, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QScrollArea{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QScrollArea) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QScrollArea_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QScrollArea) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScrollArea_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScrollArea_Metacast
func miqt_exec_callback_QScrollArea_Metacast(self QScrollArea, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QScrollArea{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
