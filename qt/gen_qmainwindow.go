package qt

import (
	"unsafe"
)

type QMainWindow__DockOption int

const (
	QMainWindow__AnimatedDocks    QMainWindow__DockOption = 1
	QMainWindow__AllowNestedDocks QMainWindow__DockOption = 2
	QMainWindow__AllowTabbedDocks QMainWindow__DockOption = 4
	QMainWindow__ForceTabbedDocks QMainWindow__DockOption = 8
	QMainWindow__VerticalTabs     QMainWindow__DockOption = 16
	QMainWindow__GroupedDragging  QMainWindow__DockOption = 32
)

type QMainWindow struct {
	h          uintptr
	isSubclass bool
}

// NewQMainWindow constructs a new QMainWindow object.
func NewQMainWindow(parent *QWidget) *QMainWindow {

	ret := newQMainWindow(QMainWindow_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQMainWindow2 constructs a new QMainWindow object.
func NewQMainWindow2() *QMainWindow {

	ret := newQMainWindow(QMainWindow_new2())
	ret.isSubclass = true
	return ret
}

// NewQMainWindow3 constructs a new QMainWindow object.
func NewQMainWindow3(parent *QWidget, flags WindowType) *QMainWindow {

	ret := newQMainWindow(QMainWindow_new3(parent.cPointer(), (int)(flags)))
	ret.isSubclass = true
	return ret
}

func (this *QMainWindow) MetaObject() *QMetaObject {
	return newQMetaObject(QMainWindow_MetaObject(this.h))
}

func (this *QMainWindow) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QMainWindow_Metacast(this.h, param1_Cstring))
}

func QMainWindow_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QMainWindow_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMainWindow) IconSize() *QSize {
	_goptr := newQSize(QMainWindow_IconSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMainWindow) SetIconSize(iconSize *QSize) {
	QMainWindow_SetIconSize(this.h, iconSize.cPointer())
}

func (this *QMainWindow) ToolButtonStyle() ToolButtonStyle {
	return (ToolButtonStyle)(QMainWindow_ToolButtonStyle(this.h))
}

func (this *QMainWindow) SetToolButtonStyle(toolButtonStyle ToolButtonStyle) {
	QMainWindow_SetToolButtonStyle(this.h, (int)(toolButtonStyle))
}

func (this *QMainWindow) IsAnimated() bool {
	return (bool)(QMainWindow_IsAnimated(this.h))
}

func (this *QMainWindow) IsDockNestingEnabled() bool {
	return (bool)(QMainWindow_IsDockNestingEnabled(this.h))
}

func (this *QMainWindow) DocumentMode() bool {
	return (bool)(QMainWindow_DocumentMode(this.h))
}

func (this *QMainWindow) SetDocumentMode(enabled bool) {
	QMainWindow_SetDocumentMode(this.h, (bool)(enabled))
}

func (this *QMainWindow) TabShape() QTabWidget__TabShape {
	return (QTabWidget__TabShape)(QMainWindow_TabShape(this.h))
}

func (this *QMainWindow) SetTabShape(tabShape QTabWidget__TabShape) {
	QMainWindow_SetTabShape(this.h, (int)(tabShape))
}

func (this *QMainWindow) TabPosition(area DockWidgetArea) QTabWidget__TabPosition {
	return (QTabWidget__TabPosition)(QMainWindow_TabPosition(this.h, (int)(area)))
}

func (this *QMainWindow) SetTabPosition(areas DockWidgetArea, tabPosition QTabWidget__TabPosition) {
	QMainWindow_SetTabPosition(this.h, (int)(areas), (int)(tabPosition))
}

func (this *QMainWindow) SetDockOptions(options DockOptions) {
	QMainWindow_SetDockOptions(this.h, options)
}

func (this *QMainWindow) DockOptions() DockOptions {
	xxxxxxxxx
}

func (this *QMainWindow) IsSeparator(pos *QPoint) bool {
	return (bool)(QMainWindow_IsSeparator(this.h, pos.cPointer()))
}

func (this *QMainWindow) MenuBar() *QMenuBar {
	return newQMenuBar(QMainWindow_MenuBar(this.h))
}

func (this *QMainWindow) SetMenuBar(menubar *QMenuBar) {
	QMainWindow_SetMenuBar(this.h, menubar.cPointer())
}

func (this *QMainWindow) MenuWidget() *QWidget {
	return newQWidget(QMainWindow_MenuWidget(this.h))
}

func (this *QMainWindow) SetMenuWidget(menubar *QWidget) {
	QMainWindow_SetMenuWidget(this.h, menubar.cPointer())
}

func (this *QMainWindow) StatusBar() *QStatusBar {
	return newQStatusBar(QMainWindow_StatusBar(this.h))
}

func (this *QMainWindow) SetStatusBar(statusbar *QStatusBar) {
	QMainWindow_SetStatusBar(this.h, statusbar.cPointer())
}

func (this *QMainWindow) CentralWidget() *QWidget {
	return newQWidget(QMainWindow_CentralWidget(this.h))
}

func (this *QMainWindow) SetCentralWidget(widget *QWidget) {
	QMainWindow_SetCentralWidget(this.h, widget.cPointer())
}

func (this *QMainWindow) TakeCentralWidget() *QWidget {
	return newQWidget(QMainWindow_TakeCentralWidget(this.h))
}

func (this *QMainWindow) SetCorner(corner Corner, area DockWidgetArea) {
	QMainWindow_SetCorner(this.h, (int)(corner), (int)(area))
}

func (this *QMainWindow) Corner(corner Corner) DockWidgetArea {
	return (DockWidgetArea)(QMainWindow_Corner(this.h, (int)(corner)))
}

func (this *QMainWindow) AddToolBarBreak() {
	QMainWindow_AddToolBarBreak(this.h)
}

func (this *QMainWindow) InsertToolBarBreak(before *QToolBar) {
	QMainWindow_InsertToolBarBreak(this.h, before.cPointer())
}

func (this *QMainWindow) AddToolBar(area ToolBarArea, toolbar *QToolBar) {
	QMainWindow_AddToolBar(this.h, (int)(area), toolbar.cPointer())
}

func (this *QMainWindow) AddToolBarWithToolbar(toolbar *QToolBar) {
	QMainWindow_AddToolBarWithToolbar(this.h, toolbar.cPointer())
}

func (this *QMainWindow) AddToolBarWithTitle(title string) *QToolBar {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	return newQToolBar(QMainWindow_AddToolBarWithTitle(this.h, title_ms))
}

func (this *QMainWindow) InsertToolBar(before *QToolBar, toolbar *QToolBar) {
	QMainWindow_InsertToolBar(this.h, before.cPointer(), toolbar.cPointer())
}

func (this *QMainWindow) RemoveToolBar(toolbar *QToolBar) {
	QMainWindow_RemoveToolBar(this.h, toolbar.cPointer())
}

func (this *QMainWindow) RemoveToolBarBreak(before *QToolBar) {
	QMainWindow_RemoveToolBarBreak(this.h, before.cPointer())
}

func (this *QMainWindow) UnifiedTitleAndToolBarOnMac() bool {
	return (bool)(QMainWindow_UnifiedTitleAndToolBarOnMac(this.h))
}

func (this *QMainWindow) ToolBarArea(toolbar *QToolBar) ToolBarArea {
	return (ToolBarArea)(QMainWindow_ToolBarArea(this.h, toolbar.cPointer()))
}

func (this *QMainWindow) ToolBarBreak(toolbar *QToolBar) bool {
	return (bool)(QMainWindow_ToolBarBreak(this.h, toolbar.cPointer()))
}

func (this *QMainWindow) AddDockWidget(area DockWidgetArea, dockwidget *QDockWidget) {
	QMainWindow_AddDockWidget(this.h, (int)(area), dockwidget.cPointer())
}

func (this *QMainWindow) AddDockWidget2(area DockWidgetArea, dockwidget *QDockWidget, orientation Orientation) {
	QMainWindow_AddDockWidget2(this.h, (int)(area), dockwidget.cPointer(), (int)(orientation))
}

func (this *QMainWindow) SplitDockWidget(after *QDockWidget, dockwidget *QDockWidget, orientation Orientation) {
	QMainWindow_SplitDockWidget(this.h, after.cPointer(), dockwidget.cPointer(), (int)(orientation))
}

func (this *QMainWindow) TabifyDockWidget(first *QDockWidget, second *QDockWidget) {
	QMainWindow_TabifyDockWidget(this.h, first.cPointer(), second.cPointer())
}

func (this *QMainWindow) TabifiedDockWidgets(dockwidget *QDockWidget) []*QDockWidget {
	var _ma struct_miqt_array = QMainWindow_TabifiedDockWidgets(this.h, dockwidget.cPointer())
	_ret := make([]*QDockWidget, int(_ma.len))
	_outCast := (*[0xffff]*QDockWidget)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQDockWidget(_outCast[i])
	}
	return _ret
}

func (this *QMainWindow) RemoveDockWidget(dockwidget *QDockWidget) {
	QMainWindow_RemoveDockWidget(this.h, dockwidget.cPointer())
}

func (this *QMainWindow) RestoreDockWidget(dockwidget *QDockWidget) bool {
	return (bool)(QMainWindow_RestoreDockWidget(this.h, dockwidget.cPointer()))
}

func (this *QMainWindow) DockWidgetArea(dockwidget *QDockWidget) DockWidgetArea {
	return (DockWidgetArea)(QMainWindow_DockWidgetArea(this.h, dockwidget.cPointer()))
}

func (this *QMainWindow) ResizeDocks(docks []*QDockWidget, sizes []int, orientation Orientation) {
	docks_CArray := (*[0xffff]*QDockWidget)(malloc(size_t(8 * len(docks))))
	defer free(unsafe.Pointer(docks_CArray))
	for i := range docks {
		docks_CArray[i] = docks[i].cPointer()
	}
	docks_ma := struct_miqt_array{len: size_t(len(docks)), data: unsafe.Pointer(docks_CArray)}
	sizes_CArray := (*[0xffff]int)(malloc(size_t(8 * len(sizes))))
	defer free(unsafe.Pointer(sizes_CArray))
	for i := range sizes {
		sizes_CArray[i] = (int)(sizes[i])
	}
	sizes_ma := struct_miqt_array{len: size_t(len(sizes)), data: unsafe.Pointer(sizes_CArray)}
	QMainWindow_ResizeDocks(this.h, docks_ma, sizes_ma, (int)(orientation))
}

func (this *QMainWindow) SaveState() []byte {
	var _bytearray struct_miqt_string = QMainWindow_SaveState(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QMainWindow) RestoreState(state []byte) bool {
	state_alias := struct_miqt_string{}
	state_alias.data = (char)(unsafe.Pointer(&state[0]))
	state_alias.len = size_t(len(state))
	return (bool)(QMainWindow_RestoreState(this.h, state_alias))
}

func (this *QMainWindow) CreatePopupMenu() *QMenu {
	return newQMenu(QMainWindow_CreatePopupMenu(this.h))
}

func (this *QMainWindow) SetAnimated(enabled bool) {
	QMainWindow_SetAnimated(this.h, (bool)(enabled))
}

func (this *QMainWindow) SetDockNestingEnabled(enabled bool) {
	QMainWindow_SetDockNestingEnabled(this.h, (bool)(enabled))
}

func (this *QMainWindow) SetUnifiedTitleAndToolBarOnMac(set bool) {
	QMainWindow_SetUnifiedTitleAndToolBarOnMac(this.h, (bool)(set))
}

func (this *QMainWindow) IconSizeChanged(iconSize *QSize) {
	QMainWindow_IconSizeChanged(this.h, iconSize.cPointer())
}
func (this *QMainWindow) OnIconSizeChanged(slot func(iconSize *QSize)) {
	QMainWindow_connect_IconSizeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_IconSizeChanged
func miqt_exec_callback_QMainWindow_IconSizeChanged(cb intptr_t, iconSize *QSize) {
	gofunc, ok := cgo.Handle(cb).Value().(func(iconSize *QSize))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSize(iconSize)

	gofunc(slotval1)
}

func (this *QMainWindow) ToolButtonStyleChanged(toolButtonStyle ToolButtonStyle) {
	QMainWindow_ToolButtonStyleChanged(this.h, (int)(toolButtonStyle))
}
func (this *QMainWindow) OnToolButtonStyleChanged(slot func(toolButtonStyle ToolButtonStyle)) {
	QMainWindow_connect_ToolButtonStyleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_ToolButtonStyleChanged
func miqt_exec_callback_QMainWindow_ToolButtonStyleChanged(cb intptr_t, toolButtonStyle int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(toolButtonStyle ToolButtonStyle))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (ToolButtonStyle)(toolButtonStyle)

	gofunc(slotval1)
}

func (this *QMainWindow) TabifiedDockWidgetActivated(dockWidget *QDockWidget) {
	QMainWindow_TabifiedDockWidgetActivated(this.h, dockWidget.cPointer())
}
func (this *QMainWindow) OnTabifiedDockWidgetActivated(slot func(dockWidget *QDockWidget)) {
	QMainWindow_connect_TabifiedDockWidgetActivated(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_TabifiedDockWidgetActivated
func miqt_exec_callback_QMainWindow_TabifiedDockWidgetActivated(cb intptr_t, dockWidget *QDockWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(dockWidget *QDockWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDockWidget(dockWidget)

	gofunc(slotval1)
}

func QMainWindow_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMainWindow_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMainWindow_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMainWindow_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMainWindow) AddToolBarBreak1(area ToolBarArea) {
	QMainWindow_AddToolBarBreak1(this.h, (int)(area))
}

func (this *QMainWindow) SaveState1(version int) []byte {
	var _bytearray struct_miqt_string = QMainWindow_SaveState1(this.h, (int)(version))
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QMainWindow) RestoreState2(state []byte, version int) bool {
	state_alias := struct_miqt_string{}
	state_alias.data = (char)(unsafe.Pointer(&state[0]))
	state_alias.len = size_t(len(state))
	return (bool)(QMainWindow_RestoreState2(this.h, state_alias, (int)(version)))
}

func (this *QMainWindow) callVirtualBase_CreatePopupMenu() *QMenu {

	return newQMenu(QMainWindow_virtualbase_CreatePopupMenu(unsafe.Pointer(this.h)))

}
func (this *QMainWindow) OnCreatePopupMenu(slot func(super func() *QMenu) *QMenu) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_CreatePopupMenu(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_CreatePopupMenu
func miqt_exec_callback_QMainWindow_CreatePopupMenu(self QMainWindow, cb intptr_t) *QMenu {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMenu) *QMenu)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMainWindow{h: self}).callVirtualBase_CreatePopupMenu)

	return virtualReturn.cPointer()

}

func (this *QMainWindow) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QMainWindow_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_ContextMenuEvent
func miqt_exec_callback_QMainWindow_ContextMenuEvent(self QMainWindow, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QMainWindow_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QMainWindow) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_Event
func miqt_exec_callback_QMainWindow_Event(self QMainWindow, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QMainWindow{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QMainWindow) callVirtualBase_DevType() int {

	return (int)(QMainWindow_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QMainWindow) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_DevType
func miqt_exec_callback_QMainWindow_DevType(self QMainWindow, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMainWindow{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QMainWindow) callVirtualBase_SetVisible(visible bool) {

	QMainWindow_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QMainWindow) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_SetVisible
func miqt_exec_callback_QMainWindow_SetVisible(self QMainWindow, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QMainWindow{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QMainWindow) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QMainWindow_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QMainWindow) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_SizeHint
func miqt_exec_callback_QMainWindow_SizeHint(self QMainWindow, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMainWindow{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QMainWindow) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QMainWindow_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QMainWindow) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_MinimumSizeHint
func miqt_exec_callback_QMainWindow_MinimumSizeHint(self QMainWindow, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMainWindow{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QMainWindow) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QMainWindow_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QMainWindow) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_HeightForWidth
func miqt_exec_callback_QMainWindow_HeightForWidth(self QMainWindow, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QMainWindow{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QMainWindow) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QMainWindow_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QMainWindow) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_HasHeightForWidth
func miqt_exec_callback_QMainWindow_HasHeightForWidth(self QMainWindow, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMainWindow{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QMainWindow) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QMainWindow_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QMainWindow) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_PaintEngine
func miqt_exec_callback_QMainWindow_PaintEngine(self QMainWindow, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMainWindow{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QMainWindow) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QMainWindow_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_MousePressEvent
func miqt_exec_callback_QMainWindow_MousePressEvent(self QMainWindow, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QMainWindow_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_MouseReleaseEvent
func miqt_exec_callback_QMainWindow_MouseReleaseEvent(self QMainWindow, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QMainWindow_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_MouseDoubleClickEvent
func miqt_exec_callback_QMainWindow_MouseDoubleClickEvent(self QMainWindow, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QMainWindow_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_MouseMoveEvent
func miqt_exec_callback_QMainWindow_MouseMoveEvent(self QMainWindow, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QMainWindow_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_WheelEvent
func miqt_exec_callback_QMainWindow_WheelEvent(self QMainWindow, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QMainWindow_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_KeyPressEvent
func miqt_exec_callback_QMainWindow_KeyPressEvent(self QMainWindow, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QMainWindow_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_KeyReleaseEvent
func miqt_exec_callback_QMainWindow_KeyReleaseEvent(self QMainWindow, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QMainWindow_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_FocusInEvent
func miqt_exec_callback_QMainWindow_FocusInEvent(self QMainWindow, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QMainWindow_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_FocusOutEvent
func miqt_exec_callback_QMainWindow_FocusOutEvent(self QMainWindow, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QMainWindow_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_EnterEvent
func miqt_exec_callback_QMainWindow_EnterEvent(self QMainWindow, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_LeaveEvent(event *QEvent) {

	QMainWindow_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_LeaveEvent
func miqt_exec_callback_QMainWindow_LeaveEvent(self QMainWindow, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QMainWindow_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_PaintEvent
func miqt_exec_callback_QMainWindow_PaintEvent(self QMainWindow, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QMainWindow_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_MoveEvent
func miqt_exec_callback_QMainWindow_MoveEvent(self QMainWindow, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QMainWindow_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_ResizeEvent
func miqt_exec_callback_QMainWindow_ResizeEvent(self QMainWindow, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QMainWindow_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_CloseEvent
func miqt_exec_callback_QMainWindow_CloseEvent(self QMainWindow, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QMainWindow_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_TabletEvent
func miqt_exec_callback_QMainWindow_TabletEvent(self QMainWindow, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_ActionEvent(event *QActionEvent) {

	QMainWindow_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_ActionEvent
func miqt_exec_callback_QMainWindow_ActionEvent(self QMainWindow, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QMainWindow_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_DragEnterEvent
func miqt_exec_callback_QMainWindow_DragEnterEvent(self QMainWindow, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QMainWindow_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_DragMoveEvent
func miqt_exec_callback_QMainWindow_DragMoveEvent(self QMainWindow, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QMainWindow_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_DragLeaveEvent
func miqt_exec_callback_QMainWindow_DragLeaveEvent(self QMainWindow, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_DropEvent(event *QDropEvent) {

	QMainWindow_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_DropEvent
func miqt_exec_callback_QMainWindow_DropEvent(self QMainWindow, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_ShowEvent(event *QShowEvent) {

	QMainWindow_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_ShowEvent
func miqt_exec_callback_QMainWindow_ShowEvent(self QMainWindow, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_HideEvent(event *QHideEvent) {

	QMainWindow_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMainWindow) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_HideEvent
func miqt_exec_callback_QMainWindow_HideEvent(self QMainWindow, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QMainWindow{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QMainWindow_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QMainWindow) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_NativeEvent
func miqt_exec_callback_QMainWindow_NativeEvent(self QMainWindow, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
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

	virtualReturn := gofunc((&QMainWindow{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QMainWindow) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QMainWindow_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QMainWindow) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_ChangeEvent
func miqt_exec_callback_QMainWindow_ChangeEvent(self QMainWindow, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QMainWindow{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QMainWindow_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QMainWindow) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_Metric
func miqt_exec_callback_QMainWindow_Metric(self QMainWindow, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QMainWindow{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QMainWindow) callVirtualBase_InitPainter(painter *QPainter) {

	QMainWindow_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QMainWindow) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_InitPainter
func miqt_exec_callback_QMainWindow_InitPainter(self QMainWindow, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QMainWindow{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QMainWindow) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QMainWindow_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QMainWindow) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_Redirected
func miqt_exec_callback_QMainWindow_Redirected(self QMainWindow, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QMainWindow{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QMainWindow) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QMainWindow_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QMainWindow) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_SharedPainter
func miqt_exec_callback_QMainWindow_SharedPainter(self QMainWindow, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMainWindow{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QMainWindow) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QMainWindow_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QMainWindow) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_InputMethodEvent
func miqt_exec_callback_QMainWindow_InputMethodEvent(self QMainWindow, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QMainWindow{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QMainWindow) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QMainWindow_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QMainWindow) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_InputMethodQuery
func miqt_exec_callback_QMainWindow_InputMethodQuery(self QMainWindow, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QMainWindow{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QMainWindow) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QMainWindow_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QMainWindow) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_FocusNextPrevChild
func miqt_exec_callback_QMainWindow_FocusNextPrevChild(self QMainWindow, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QMainWindow{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
