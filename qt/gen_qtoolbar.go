package qt

import (
	"unsafe"
)

type QToolBar struct {
	h          uintptr
	isSubclass bool
}

// NewQToolBar constructs a new QToolBar object.
func NewQToolBar(parent *QWidget) *QToolBar {
	ret := newQToolBar(QToolBar_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQToolBar2 constructs a new QToolBar object.
func NewQToolBar2(title string) *QToolBar {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))

	ret := newQToolBar(QToolBar_new2(title_ms))
	ret.isSubclass = true
	return ret
}

// NewQToolBar3 constructs a new QToolBar object.
func NewQToolBar3() *QToolBar {
	ret := newQToolBar(QToolBar_new3())
	ret.isSubclass = true
	return ret
}

// NewQToolBar4 constructs a new QToolBar object.
func NewQToolBar4(title string, parent *QWidget) *QToolBar {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))

	ret := newQToolBar(QToolBar_new4(title_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QToolBar) MetaObject() *QMetaObject {
	return newQMetaObject(QToolBar_MetaObject(this.h))
}

func (this *QToolBar) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QToolBar_Metacast(this.h, param1_Cstring))
}

func QToolBar_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QToolBar_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QToolBar) SetMovable(movable bool) {
	QToolBar_SetMovable(this.h, (bool)(movable))
}

func (this *QToolBar) IsMovable() bool {
	return (bool)(QToolBar_IsMovable(this.h))
}

func (this *QToolBar) SetAllowedAreas(areas ToolBarArea) {
	QToolBar_SetAllowedAreas(this.h, (int)(areas))
}

func (this *QToolBar) AllowedAreas() ToolBarArea {
	return (ToolBarArea)(QToolBar_AllowedAreas(this.h))
}

func (this *QToolBar) IsAreaAllowed(area ToolBarArea) bool {
	return (bool)(QToolBar_IsAreaAllowed(this.h, (int)(area)))
}

func (this *QToolBar) SetOrientation(orientation Orientation) {
	QToolBar_SetOrientation(this.h, (int)(orientation))
}

func (this *QToolBar) Orientation() Orientation {
	return (Orientation)(QToolBar_Orientation(this.h))
}

func (this *QToolBar) Clear() {
	QToolBar_Clear(this.h)
}

func (this *QToolBar) AddSeparator() *QAction {
	return newQAction(QToolBar_AddSeparator(this.h))
}

func (this *QToolBar) InsertSeparator(before *QAction) *QAction {
	return newQAction(QToolBar_InsertSeparator(this.h, before.cPointer()))
}

func (this *QToolBar) AddWidget(widget *QWidget) *QAction {
	return newQAction(QToolBar_AddWidget(this.h, widget.cPointer()))
}

func (this *QToolBar) InsertWidget(before *QAction, widget *QWidget) *QAction {
	return newQAction(QToolBar_InsertWidget(this.h, before.cPointer(), widget.cPointer()))
}

func (this *QToolBar) ActionGeometry(action *QAction) *QRect {
	_goptr := newQRect(QToolBar_ActionGeometry(this.h, action.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QToolBar) ActionAt(p *QPoint) *QAction {
	return newQAction(QToolBar_ActionAt(this.h, p.cPointer()))
}

func (this *QToolBar) ActionAt2(x int, y int) *QAction {
	return newQAction(QToolBar_ActionAt2(this.h, (int)(x), (int)(y)))
}

func (this *QToolBar) ToggleViewAction() *QAction {
	return newQAction(QToolBar_ToggleViewAction(this.h))
}

func (this *QToolBar) IconSize() *QSize {
	_goptr := newQSize(QToolBar_IconSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QToolBar) ToolButtonStyle() ToolButtonStyle {
	return (ToolButtonStyle)(QToolBar_ToolButtonStyle(this.h))
}

func (this *QToolBar) WidgetForAction(action *QAction) *QWidget {
	return newQWidget(QToolBar_WidgetForAction(this.h, action.cPointer()))
}

func (this *QToolBar) IsFloatable() bool {
	return (bool)(QToolBar_IsFloatable(this.h))
}

func (this *QToolBar) SetFloatable(floatable bool) {
	QToolBar_SetFloatable(this.h, (bool)(floatable))
}

func (this *QToolBar) IsFloating() bool {
	return (bool)(QToolBar_IsFloating(this.h))
}

func (this *QToolBar) SetIconSize(iconSize *QSize) {
	QToolBar_SetIconSize(this.h, iconSize.cPointer())
}

func (this *QToolBar) SetToolButtonStyle(toolButtonStyle ToolButtonStyle) {
	QToolBar_SetToolButtonStyle(this.h, (int)(toolButtonStyle))
}

func (this *QToolBar) ActionTriggered(action *QAction) {
	QToolBar_ActionTriggered(this.h, action.cPointer())
}

func (this *QToolBar) OnActionTriggered(slot func(action *QAction)) {
	QToolBar_connect_ActionTriggered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_ActionTriggered
func miqt_exec_callback_QToolBar_ActionTriggered(cb intptr_t, action *QAction) {
	gofunc, ok := cgo.Handle(cb).Value().(func(action *QAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAction(action)

	gofunc(slotval1)
}

func (this *QToolBar) MovableChanged(movable bool) {
	QToolBar_MovableChanged(this.h, (bool)(movable))
}

func (this *QToolBar) OnMovableChanged(slot func(movable bool)) {
	QToolBar_connect_MovableChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_MovableChanged
func miqt_exec_callback_QToolBar_MovableChanged(cb intptr_t, movable bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(movable bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(movable)

	gofunc(slotval1)
}

func (this *QToolBar) AllowedAreasChanged(allowedAreas ToolBarArea) {
	QToolBar_AllowedAreasChanged(this.h, (int)(allowedAreas))
}

func (this *QToolBar) OnAllowedAreasChanged(slot func(allowedAreas ToolBarArea)) {
	QToolBar_connect_AllowedAreasChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_AllowedAreasChanged
func miqt_exec_callback_QToolBar_AllowedAreasChanged(cb intptr_t, allowedAreas int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(allowedAreas ToolBarArea))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (ToolBarArea)(allowedAreas)

	gofunc(slotval1)
}

func (this *QToolBar) OrientationChanged(orientation Orientation) {
	QToolBar_OrientationChanged(this.h, (int)(orientation))
}

func (this *QToolBar) OnOrientationChanged(slot func(orientation Orientation)) {
	QToolBar_connect_OrientationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_OrientationChanged
func miqt_exec_callback_QToolBar_OrientationChanged(cb intptr_t, orientation int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(orientation Orientation))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (Orientation)(orientation)

	gofunc(slotval1)
}

func (this *QToolBar) IconSizeChanged(iconSize *QSize) {
	QToolBar_IconSizeChanged(this.h, iconSize.cPointer())
}

func (this *QToolBar) OnIconSizeChanged(slot func(iconSize *QSize)) {
	QToolBar_connect_IconSizeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_IconSizeChanged
func miqt_exec_callback_QToolBar_IconSizeChanged(cb intptr_t, iconSize *QSize) {
	gofunc, ok := cgo.Handle(cb).Value().(func(iconSize *QSize))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSize(iconSize)

	gofunc(slotval1)
}

func (this *QToolBar) ToolButtonStyleChanged(toolButtonStyle ToolButtonStyle) {
	QToolBar_ToolButtonStyleChanged(this.h, (int)(toolButtonStyle))
}

func (this *QToolBar) OnToolButtonStyleChanged(slot func(toolButtonStyle ToolButtonStyle)) {
	QToolBar_connect_ToolButtonStyleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_ToolButtonStyleChanged
func miqt_exec_callback_QToolBar_ToolButtonStyleChanged(cb intptr_t, toolButtonStyle int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(toolButtonStyle ToolButtonStyle))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (ToolButtonStyle)(toolButtonStyle)

	gofunc(slotval1)
}

func (this *QToolBar) TopLevelChanged(topLevel bool) {
	QToolBar_TopLevelChanged(this.h, (bool)(topLevel))
}

func (this *QToolBar) OnTopLevelChanged(slot func(topLevel bool)) {
	QToolBar_connect_TopLevelChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_TopLevelChanged
func miqt_exec_callback_QToolBar_TopLevelChanged(cb intptr_t, topLevel bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(topLevel bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(topLevel)

	gofunc(slotval1)
}

func (this *QToolBar) VisibilityChanged(visible bool) {
	QToolBar_VisibilityChanged(this.h, (bool)(visible))
}

func (this *QToolBar) OnVisibilityChanged(slot func(visible bool)) {
	QToolBar_connect_VisibilityChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_VisibilityChanged
func miqt_exec_callback_QToolBar_VisibilityChanged(cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc(slotval1)
}

func QToolBar_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QToolBar_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QToolBar_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QToolBar_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QToolBar) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QToolBar_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QToolBar) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_MetaObject
func miqt_exec_callback_QToolBar_MetaObject(self QToolBar, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QToolBar{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QToolBar) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QToolBar_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QToolBar) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_Metacast
func miqt_exec_callback_QToolBar_Metacast(self QToolBar, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QToolBar{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
