package qt

import (
	"unsafe"
)

type QMdiSubWindow__SubWindowOption int

const (
	QMdiSubWindow__AllowOutsideAreaHorizontally QMdiSubWindow__SubWindowOption = 1
	QMdiSubWindow__AllowOutsideAreaVertically   QMdiSubWindow__SubWindowOption = 2
	QMdiSubWindow__RubberBandResize             QMdiSubWindow__SubWindowOption = 4
	QMdiSubWindow__RubberBandMove               QMdiSubWindow__SubWindowOption = 8
)

type QMdiSubWindow struct {
	h          uintptr
	isSubclass bool
}

// NewQMdiSubWindow constructs a new QMdiSubWindow object.
func NewQMdiSubWindow(parent *QWidget) *QMdiSubWindow {
	g := newQMdiSubWindow(QMdiSubWindow_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQMdiSubWindow2 constructs a new QMdiSubWindow object.
func NewQMdiSubWindow2() *QMdiSubWindow {
	g := newQMdiSubWindow(QMdiSubWindow_new2())
	g.isSubclass = true
	return g
}

// NewQMdiSubWindow3 constructs a new QMdiSubWindow object.
func NewQMdiSubWindow3(parent *QWidget, flags WindowType) *QMdiSubWindow {
	g := newQMdiSubWindow(QMdiSubWindow_new3(parent.cPointer(), (int)(flags)))
	g.isSubclass = true
	return g
}

func (this *QMdiSubWindow) MetaObject() *QMetaObject {
	return newQMetaObject(QMdiSubWindow_MetaObject(this.h))
}

func (this *QMdiSubWindow) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QMdiSubWindow_Metacast(this.h, param1_Cstring))
}

func QMdiSubWindow_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QMdiSubWindow_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMdiSubWindow) SizeHint() *QSize {
	_goptr := newQSize(QMdiSubWindow_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMdiSubWindow) MinimumSizeHint() *QSize {
	_goptr := newQSize(QMdiSubWindow_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMdiSubWindow) SetWidget(widget *QWidget) {
	QMdiSubWindow_SetWidget(this.h, widget.cPointer())
}

func (this *QMdiSubWindow) Widget() *QWidget {
	return newQWidget(QMdiSubWindow_Widget(this.h))
}

func (this *QMdiSubWindow) MaximizedButtonsWidget() *QWidget {
	return newQWidget(QMdiSubWindow_MaximizedButtonsWidget(this.h))
}

func (this *QMdiSubWindow) MaximizedSystemMenuIconWidget() *QWidget {
	return newQWidget(QMdiSubWindow_MaximizedSystemMenuIconWidget(this.h))
}

func (this *QMdiSubWindow) IsShaded() bool {
	return (bool)(QMdiSubWindow_IsShaded(this.h))
}

func (this *QMdiSubWindow) SetOption(option SubWindowOption) {
	QMdiSubWindow_SetOption(this.h, option)
}

func (this *QMdiSubWindow) TestOption(param1 SubWindowOption) bool {
	return (bool)(QMdiSubWindow_TestOption(this.h, param1))
}

func (this *QMdiSubWindow) SetKeyboardSingleStep(step int) {
	QMdiSubWindow_SetKeyboardSingleStep(this.h, (int)(step))
}

func (this *QMdiSubWindow) KeyboardSingleStep() int {
	return (int)(QMdiSubWindow_KeyboardSingleStep(this.h))
}

func (this *QMdiSubWindow) SetKeyboardPageStep(step int) {
	QMdiSubWindow_SetKeyboardPageStep(this.h, (int)(step))
}

func (this *QMdiSubWindow) KeyboardPageStep() int {
	return (int)(QMdiSubWindow_KeyboardPageStep(this.h))
}

func (this *QMdiSubWindow) SetSystemMenu(systemMenu *QMenu) {
	QMdiSubWindow_SetSystemMenu(this.h, systemMenu.cPointer())
}

func (this *QMdiSubWindow) SystemMenu() *QMenu {
	return newQMenu(QMdiSubWindow_SystemMenu(this.h))
}

func (this *QMdiSubWindow) MdiArea() *QMdiArea {
	return newQMdiArea(QMdiSubWindow_MdiArea(this.h))
}

func (this *QMdiSubWindow) WindowStateChanged(oldState WindowState, newState WindowState) {
	QMdiSubWindow_WindowStateChanged(this.h, (int)(oldState), (int)(newState))
}

func (this *QMdiSubWindow) OnWindowStateChanged(slot func(oldState WindowState, newState WindowState)) {
	QMdiSubWindow_connect_WindowStateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_WindowStateChanged
func miqt_exec_callback_QMdiSubWindow_WindowStateChanged(cb intptr_t, oldState int, newState int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(oldState WindowState, newState WindowState))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (WindowState)(oldState)

	slotval2 := (WindowState)(newState)

	gofunc(slotval1, slotval2)
}

func (this *QMdiSubWindow) AboutToActivate() {
	QMdiSubWindow_AboutToActivate(this.h)
}

func (this *QMdiSubWindow) OnAboutToActivate(slot func()) {
	QMdiSubWindow_connect_AboutToActivate(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_AboutToActivate
func miqt_exec_callback_QMdiSubWindow_AboutToActivate(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMdiSubWindow) ShowSystemMenu() {
	QMdiSubWindow_ShowSystemMenu(this.h)
}

func (this *QMdiSubWindow) ShowShaded() {
	QMdiSubWindow_ShowShaded(this.h)
}

func QMdiSubWindow_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMdiSubWindow_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMdiSubWindow_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMdiSubWindow_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMdiSubWindow) SetOption2(option SubWindowOption, on bool) {
	QMdiSubWindow_SetOption2(this.h, option, (bool)(on))
}

func (this *QMdiSubWindow) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QMdiSubWindow_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QMdiSubWindow) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_MetaObject
func miqt_exec_callback_QMdiSubWindow_MetaObject(self QMdiSubWindow, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMdiSubWindow{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QMdiSubWindow) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QMdiSubWindow_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QMdiSubWindow) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiSubWindow_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiSubWindow_Metacast
func miqt_exec_callback_QMdiSubWindow_Metacast(self QMdiSubWindow, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QMdiSubWindow{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
