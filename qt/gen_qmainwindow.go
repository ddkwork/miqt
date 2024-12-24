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

func (this *QMainWindow) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QMainWindow_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QMainWindow) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_MetaObject
func miqt_exec_callback_QMainWindow_MetaObject(self QMainWindow, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMainWindow{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QMainWindow) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QMainWindow_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QMainWindow) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMainWindow_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMainWindow_Metacast
func miqt_exec_callback_QMainWindow_Metacast(self QMainWindow, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QMainWindow{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
