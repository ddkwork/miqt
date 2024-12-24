package qt

import (
	"unsafe"
)

type QMenuBar struct {
	h          uintptr
	isSubclass bool
}

// NewQMenuBar constructs a new QMenuBar object.
func NewQMenuBar(parent *QWidget) *QMenuBar {
	g := newQMenuBar(QMenuBar_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQMenuBar2 constructs a new QMenuBar object.
func NewQMenuBar2() *QMenuBar {
	g := newQMenuBar(QMenuBar_new2())
	g.isSubclass = true
	return g
}

func (this *QMenuBar) MetaObject() *QMetaObject {
	return newQMetaObject(QMenuBar_MetaObject(this.h))
}

func (this *QMenuBar) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QMenuBar_Metacast(this.h, param1_Cstring))
}

func QMenuBar_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QMenuBar_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMenuBar) AddMenu(menu *QMenu) *QAction {
	return newQAction(QMenuBar_AddMenu(this.h, menu.cPointer()))
}

func (this *QMenuBar) AddMenuWithTitle(title string) *QMenu {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	return newQMenu(QMenuBar_AddMenuWithTitle(this.h, title_ms))
}

func (this *QMenuBar) AddMenu2(icon *QIcon, title string) *QMenu {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	return newQMenu(QMenuBar_AddMenu2(this.h, icon.cPointer(), title_ms))
}

func (this *QMenuBar) AddSeparator() *QAction {
	return newQAction(QMenuBar_AddSeparator(this.h))
}

func (this *QMenuBar) InsertSeparator(before *QAction) *QAction {
	return newQAction(QMenuBar_InsertSeparator(this.h, before.cPointer()))
}

func (this *QMenuBar) InsertMenu(before *QAction, menu *QMenu) *QAction {
	return newQAction(QMenuBar_InsertMenu(this.h, before.cPointer(), menu.cPointer()))
}

func (this *QMenuBar) Clear() {
	QMenuBar_Clear(this.h)
}

func (this *QMenuBar) ActiveAction() *QAction {
	return newQAction(QMenuBar_ActiveAction(this.h))
}

func (this *QMenuBar) SetActiveAction(action *QAction) {
	QMenuBar_SetActiveAction(this.h, action.cPointer())
}

func (this *QMenuBar) SetDefaultUp(defaultUp bool) {
	QMenuBar_SetDefaultUp(this.h, (bool)(defaultUp))
}

func (this *QMenuBar) IsDefaultUp() bool {
	return (bool)(QMenuBar_IsDefaultUp(this.h))
}

func (this *QMenuBar) SizeHint() *QSize {
	_goptr := newQSize(QMenuBar_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMenuBar) MinimumSizeHint() *QSize {
	_goptr := newQSize(QMenuBar_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMenuBar) HeightForWidth(param1 int) int {
	return (int)(QMenuBar_HeightForWidth(this.h, (int)(param1)))
}

func (this *QMenuBar) ActionGeometry(param1 *QAction) *QRect {
	_goptr := newQRect(QMenuBar_ActionGeometry(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMenuBar) ActionAt(param1 *QPoint) *QAction {
	return newQAction(QMenuBar_ActionAt(this.h, param1.cPointer()))
}

func (this *QMenuBar) SetCornerWidget(w *QWidget) {
	QMenuBar_SetCornerWidget(this.h, w.cPointer())
}

func (this *QMenuBar) CornerWidget() *QWidget {
	return newQWidget(QMenuBar_CornerWidget(this.h))
}

func (this *QMenuBar) IsNativeMenuBar() bool {
	return (bool)(QMenuBar_IsNativeMenuBar(this.h))
}

func (this *QMenuBar) SetNativeMenuBar(nativeMenuBar bool) {
	QMenuBar_SetNativeMenuBar(this.h, (bool)(nativeMenuBar))
}

func (this *QMenuBar) SetVisible(visible bool) {
	QMenuBar_SetVisible(this.h, (bool)(visible))
}

func (this *QMenuBar) Triggered(action *QAction) {
	QMenuBar_Triggered(this.h, action.cPointer())
}

func (this *QMenuBar) OnTriggered(slot func(action *QAction)) {
	QMenuBar_connect_Triggered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMenuBar_Triggered
func miqt_exec_callback_QMenuBar_Triggered(cb intptr_t, action *QAction) {
	gofunc, ok := cgo.Handle(cb).Value().(func(action *QAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAction(action)

	gofunc(slotval1)
}

func (this *QMenuBar) Hovered(action *QAction) {
	QMenuBar_Hovered(this.h, action.cPointer())
}

func (this *QMenuBar) OnHovered(slot func(action *QAction)) {
	QMenuBar_connect_Hovered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMenuBar_Hovered
func miqt_exec_callback_QMenuBar_Hovered(cb intptr_t, action *QAction) {
	gofunc, ok := cgo.Handle(cb).Value().(func(action *QAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAction(action)

	gofunc(slotval1)
}

func QMenuBar_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMenuBar_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMenuBar_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMenuBar_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMenuBar) SetCornerWidget2(w *QWidget, corner Corner) {
	QMenuBar_SetCornerWidget2(this.h, w.cPointer(), (int)(corner))
}

func (this *QMenuBar) CornerWidget1(corner Corner) *QWidget {
	return newQWidget(QMenuBar_CornerWidget1(this.h, (int)(corner)))
}

func (this *QMenuBar) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QMenuBar_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QMenuBar) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMenuBar_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMenuBar_MetaObject
func miqt_exec_callback_QMenuBar_MetaObject(self QMenuBar, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMenuBar{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QMenuBar) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QMenuBar_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QMenuBar) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMenuBar_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMenuBar_Metacast
func miqt_exec_callback_QMenuBar_Metacast(self QMenuBar, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QMenuBar{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
