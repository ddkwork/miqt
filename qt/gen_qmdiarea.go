package qt

import (
	"unsafe"
)

type QMdiArea__AreaOption int

const (
	QMdiArea__DontMaximizeSubWindowOnActivation QMdiArea__AreaOption = 1
)

type QMdiArea__WindowOrder int

const (
	QMdiArea__CreationOrder          QMdiArea__WindowOrder = 0
	QMdiArea__StackingOrder          QMdiArea__WindowOrder = 1
	QMdiArea__ActivationHistoryOrder QMdiArea__WindowOrder = 2
)

type QMdiArea__ViewMode int

const (
	QMdiArea__SubWindowView QMdiArea__ViewMode = 0
	QMdiArea__TabbedView    QMdiArea__ViewMode = 1
)

type QMdiArea struct {
	h          uintptr
	isSubclass bool
}

// NewQMdiArea constructs a new QMdiArea object.
func NewQMdiArea(parent *QWidget) *QMdiArea {
	g := newQMdiArea(QMdiArea_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQMdiArea2 constructs a new QMdiArea object.
func NewQMdiArea2() *QMdiArea {
	g := newQMdiArea(QMdiArea_new2())
	g.isSubclass = true
	return g
}

func (this *QMdiArea) MetaObject() *QMetaObject {
	return newQMetaObject(QMdiArea_MetaObject(this.h))
}

func (this *QMdiArea) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QMdiArea_Metacast(this.h, param1_Cstring))
}

func QMdiArea_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QMdiArea_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMdiArea) SizeHint() *QSize {
	_goptr := newQSize(QMdiArea_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMdiArea) MinimumSizeHint() *QSize {
	_goptr := newQSize(QMdiArea_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMdiArea) CurrentSubWindow() *QMdiSubWindow {
	return newQMdiSubWindow(QMdiArea_CurrentSubWindow(this.h))
}

func (this *QMdiArea) ActiveSubWindow() *QMdiSubWindow {
	return newQMdiSubWindow(QMdiArea_ActiveSubWindow(this.h))
}

func (this *QMdiArea) SubWindowList() []*QMdiSubWindow {
	var _ma struct_miqt_array = QMdiArea_SubWindowList(this.h)
	_ret := make([]*QMdiSubWindow, int(_ma.len))
	_outCast := (*[0xffff]*QMdiSubWindow)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQMdiSubWindow(_outCast[i])
	}
	return _ret
}

func (this *QMdiArea) AddSubWindow(widget *QWidget) *QMdiSubWindow {
	return newQMdiSubWindow(QMdiArea_AddSubWindow(this.h, widget.cPointer()))
}

func (this *QMdiArea) RemoveSubWindow(widget *QWidget) {
	QMdiArea_RemoveSubWindow(this.h, widget.cPointer())
}

func (this *QMdiArea) Background() *QBrush {
	_goptr := newQBrush(QMdiArea_Background(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMdiArea) SetBackground(background *QBrush) {
	QMdiArea_SetBackground(this.h, background.cPointer())
}

func (this *QMdiArea) ActivationOrder() WindowOrder {
	xxxxxxxxx
}

func (this *QMdiArea) SetActivationOrder(order WindowOrder) {
	QMdiArea_SetActivationOrder(this.h, order)
}

func (this *QMdiArea) SetOption(option AreaOption) {
	QMdiArea_SetOption(this.h, option)
}

func (this *QMdiArea) TestOption(opton AreaOption) bool {
	return (bool)(QMdiArea_TestOption(this.h, opton))
}

func (this *QMdiArea) SetViewMode(mode ViewMode) {
	QMdiArea_SetViewMode(this.h, mode)
}

func (this *QMdiArea) ViewMode() ViewMode {
	xxxxxxxxx
}

func (this *QMdiArea) DocumentMode() bool {
	return (bool)(QMdiArea_DocumentMode(this.h))
}

func (this *QMdiArea) SetDocumentMode(enabled bool) {
	QMdiArea_SetDocumentMode(this.h, (bool)(enabled))
}

func (this *QMdiArea) SetTabsClosable(closable bool) {
	QMdiArea_SetTabsClosable(this.h, (bool)(closable))
}

func (this *QMdiArea) TabsClosable() bool {
	return (bool)(QMdiArea_TabsClosable(this.h))
}

func (this *QMdiArea) SetTabsMovable(movable bool) {
	QMdiArea_SetTabsMovable(this.h, (bool)(movable))
}

func (this *QMdiArea) TabsMovable() bool {
	return (bool)(QMdiArea_TabsMovable(this.h))
}

func (this *QMdiArea) SetTabShape(shape QTabWidget__TabShape) {
	QMdiArea_SetTabShape(this.h, (int)(shape))
}

func (this *QMdiArea) TabShape() QTabWidget__TabShape {
	return (QTabWidget__TabShape)(QMdiArea_TabShape(this.h))
}

func (this *QMdiArea) SetTabPosition(position QTabWidget__TabPosition) {
	QMdiArea_SetTabPosition(this.h, (int)(position))
}

func (this *QMdiArea) TabPosition() QTabWidget__TabPosition {
	return (QTabWidget__TabPosition)(QMdiArea_TabPosition(this.h))
}

func (this *QMdiArea) SubWindowActivated(param1 *QMdiSubWindow) {
	QMdiArea_SubWindowActivated(this.h, param1.cPointer())
}

func (this *QMdiArea) OnSubWindowActivated(slot func(param1 *QMdiSubWindow)) {
	QMdiArea_connect_SubWindowActivated(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_SubWindowActivated
func miqt_exec_callback_QMdiArea_SubWindowActivated(cb intptr_t, param1 *QMdiSubWindow) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QMdiSubWindow))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMdiSubWindow(param1)

	gofunc(slotval1)
}

func (this *QMdiArea) SetActiveSubWindow(window *QMdiSubWindow) {
	QMdiArea_SetActiveSubWindow(this.h, window.cPointer())
}

func (this *QMdiArea) TileSubWindows() {
	QMdiArea_TileSubWindows(this.h)
}

func (this *QMdiArea) CascadeSubWindows() {
	QMdiArea_CascadeSubWindows(this.h)
}

func (this *QMdiArea) CloseActiveSubWindow() {
	QMdiArea_CloseActiveSubWindow(this.h)
}

func (this *QMdiArea) CloseAllSubWindows() {
	QMdiArea_CloseAllSubWindows(this.h)
}

func (this *QMdiArea) ActivateNextSubWindow() {
	QMdiArea_ActivateNextSubWindow(this.h)
}

func (this *QMdiArea) ActivatePreviousSubWindow() {
	QMdiArea_ActivatePreviousSubWindow(this.h)
}

func QMdiArea_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMdiArea_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMdiArea_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMdiArea_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMdiArea) SubWindowList1(order WindowOrder) []*QMdiSubWindow {
	var _ma struct_miqt_array = QMdiArea_SubWindowList1(this.h, order)
	_ret := make([]*QMdiSubWindow, int(_ma.len))
	_outCast := (*[0xffff]*QMdiSubWindow)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQMdiSubWindow(_outCast[i])
	}
	return _ret
}

func (this *QMdiArea) AddSubWindow2(widget *QWidget, flags WindowType) *QMdiSubWindow {
	return newQMdiSubWindow(QMdiArea_AddSubWindow2(this.h, widget.cPointer(), (int)(flags)))
}

func (this *QMdiArea) SetOption2(option AreaOption, on bool) {
	QMdiArea_SetOption2(this.h, option, (bool)(on))
}

func (this *QMdiArea) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QMdiArea_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QMdiArea) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_MetaObject
func miqt_exec_callback_QMdiArea_MetaObject(self QMdiArea, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMdiArea{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QMdiArea) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QMdiArea_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QMdiArea) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMdiArea_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMdiArea_Metacast
func miqt_exec_callback_QMdiArea_Metacast(self QMdiArea, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QMdiArea{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
