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

func (this *QToolBar) callVirtualBase_ActionEvent(event *QActionEvent) {

	QToolBar_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_ActionEvent
func miqt_exec_callback_QToolBar_ActionEvent(self QToolBar, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_ChangeEvent(event *QEvent) {

	QToolBar_virtualbase_ChangeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnChangeEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_ChangeEvent
func miqt_exec_callback_QToolBar_ChangeEvent(self QToolBar, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QToolBar_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_PaintEvent
func miqt_exec_callback_QToolBar_PaintEvent(self QToolBar, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QToolBar_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QToolBar) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_Event
func miqt_exec_callback_QToolBar_Event(self QToolBar, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QToolBar{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QToolBar) callVirtualBase_InitStyleOption(option *QStyleOptionToolBar) {

	QToolBar_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QToolBar) OnInitStyleOption(slot func(super func(option *QStyleOptionToolBar), option *QStyleOptionToolBar)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_InitStyleOption
func miqt_exec_callback_QToolBar_InitStyleOption(self QToolBar, cb intptr_t, option *QStyleOptionToolBar) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionToolBar), option *QStyleOptionToolBar))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionToolBar(option)

	gofunc((&QToolBar{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QToolBar) callVirtualBase_DevType() int {

	return (int)(QToolBar_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QToolBar) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_DevType
func miqt_exec_callback_QToolBar_DevType(self QToolBar, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QToolBar{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QToolBar) callVirtualBase_SetVisible(visible bool) {

	QToolBar_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QToolBar) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_SetVisible
func miqt_exec_callback_QToolBar_SetVisible(self QToolBar, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QToolBar{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QToolBar) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QToolBar_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QToolBar) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_SizeHint
func miqt_exec_callback_QToolBar_SizeHint(self QToolBar, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QToolBar{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QToolBar) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QToolBar_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QToolBar) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_MinimumSizeHint
func miqt_exec_callback_QToolBar_MinimumSizeHint(self QToolBar, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QToolBar{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QToolBar) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QToolBar_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QToolBar) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_HeightForWidth
func miqt_exec_callback_QToolBar_HeightForWidth(self QToolBar, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QToolBar{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QToolBar) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QToolBar_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QToolBar) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_HasHeightForWidth
func miqt_exec_callback_QToolBar_HasHeightForWidth(self QToolBar, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QToolBar{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QToolBar) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QToolBar_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QToolBar) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_PaintEngine
func miqt_exec_callback_QToolBar_PaintEngine(self QToolBar, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QToolBar{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QToolBar) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QToolBar_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_MousePressEvent
func miqt_exec_callback_QToolBar_MousePressEvent(self QToolBar, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QToolBar_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_MouseReleaseEvent
func miqt_exec_callback_QToolBar_MouseReleaseEvent(self QToolBar, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QToolBar_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_MouseDoubleClickEvent
func miqt_exec_callback_QToolBar_MouseDoubleClickEvent(self QToolBar, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QToolBar_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_MouseMoveEvent
func miqt_exec_callback_QToolBar_MouseMoveEvent(self QToolBar, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QToolBar_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_WheelEvent
func miqt_exec_callback_QToolBar_WheelEvent(self QToolBar, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QToolBar_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_KeyPressEvent
func miqt_exec_callback_QToolBar_KeyPressEvent(self QToolBar, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QToolBar_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_KeyReleaseEvent
func miqt_exec_callback_QToolBar_KeyReleaseEvent(self QToolBar, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QToolBar_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_FocusInEvent
func miqt_exec_callback_QToolBar_FocusInEvent(self QToolBar, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QToolBar_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_FocusOutEvent
func miqt_exec_callback_QToolBar_FocusOutEvent(self QToolBar, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QToolBar_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_EnterEvent
func miqt_exec_callback_QToolBar_EnterEvent(self QToolBar, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_LeaveEvent(event *QEvent) {

	QToolBar_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_LeaveEvent
func miqt_exec_callback_QToolBar_LeaveEvent(self QToolBar, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QToolBar_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_MoveEvent
func miqt_exec_callback_QToolBar_MoveEvent(self QToolBar, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QToolBar_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_ResizeEvent
func miqt_exec_callback_QToolBar_ResizeEvent(self QToolBar, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QToolBar_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_CloseEvent
func miqt_exec_callback_QToolBar_CloseEvent(self QToolBar, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QToolBar_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_ContextMenuEvent
func miqt_exec_callback_QToolBar_ContextMenuEvent(self QToolBar, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QToolBar_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_TabletEvent
func miqt_exec_callback_QToolBar_TabletEvent(self QToolBar, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QToolBar_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_DragEnterEvent
func miqt_exec_callback_QToolBar_DragEnterEvent(self QToolBar, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QToolBar_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_DragMoveEvent
func miqt_exec_callback_QToolBar_DragMoveEvent(self QToolBar, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QToolBar_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_DragLeaveEvent
func miqt_exec_callback_QToolBar_DragLeaveEvent(self QToolBar, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_DropEvent(event *QDropEvent) {

	QToolBar_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_DropEvent
func miqt_exec_callback_QToolBar_DropEvent(self QToolBar, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_ShowEvent(event *QShowEvent) {

	QToolBar_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_ShowEvent
func miqt_exec_callback_QToolBar_ShowEvent(self QToolBar, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_HideEvent(event *QHideEvent) {

	QToolBar_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QToolBar) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_HideEvent
func miqt_exec_callback_QToolBar_HideEvent(self QToolBar, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QToolBar{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QToolBar_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QToolBar) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_NativeEvent
func miqt_exec_callback_QToolBar_NativeEvent(self QToolBar, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray struct_miqt_string = eventType
	eventType_ret := GoBytes(unsafe.Pointer(eventType_bytearray.data), int(int64(eventType_bytearray.len)))
	free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*uintptr)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QToolBar{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QToolBar) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QToolBar_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QToolBar) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_Metric
func miqt_exec_callback_QToolBar_Metric(self QToolBar, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QToolBar{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QToolBar) callVirtualBase_InitPainter(painter *QPainter) {

	QToolBar_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QToolBar) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_InitPainter
func miqt_exec_callback_QToolBar_InitPainter(self QToolBar, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QToolBar{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QToolBar) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QToolBar_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QToolBar) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_Redirected
func miqt_exec_callback_QToolBar_Redirected(self QToolBar, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QToolBar{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QToolBar) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QToolBar_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QToolBar) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_SharedPainter
func miqt_exec_callback_QToolBar_SharedPainter(self QToolBar, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QToolBar{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QToolBar) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QToolBar_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QToolBar) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_InputMethodEvent
func miqt_exec_callback_QToolBar_InputMethodEvent(self QToolBar, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QToolBar{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QToolBar) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QToolBar_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QToolBar) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_InputMethodQuery
func miqt_exec_callback_QToolBar_InputMethodQuery(self QToolBar, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QToolBar{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QToolBar) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QToolBar_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QToolBar) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QToolBar_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QToolBar_FocusNextPrevChild
func miqt_exec_callback_QToolBar_FocusNextPrevChild(self QToolBar, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QToolBar{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
