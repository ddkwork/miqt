package qt

import (
	"unsafe"
)

type QWindow__Visibility int

const (
	QWindow__Hidden              QWindow__Visibility = 0
	QWindow__AutomaticVisibility QWindow__Visibility = 1
	QWindow__Windowed            QWindow__Visibility = 2
	QWindow__Minimized           QWindow__Visibility = 3
	QWindow__Maximized           QWindow__Visibility = 4
	QWindow__FullScreen          QWindow__Visibility = 5
)

type QWindow__AncestorMode int

const (
	QWindow__ExcludeTransients QWindow__AncestorMode = 0
	QWindow__IncludeTransients QWindow__AncestorMode = 1
)

type QWindow struct {
	h          uintptr
	isSubclass bool
}

// NewQWindow constructs a new QWindow object.
func NewQWindow() *QWindow {
	ret := newQWindow(QWindow_new())
	ret.isSubclass = true
	return ret
}

// NewQWindow2 constructs a new QWindow object.
func NewQWindow2(parent *QWindow) *QWindow {
	ret := newQWindow(QWindow_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQWindow3 constructs a new QWindow object.
func NewQWindow3(screen *QScreen) *QWindow {
	ret := newQWindow(QWindow_new3(screen.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QWindow) MetaObject() *QMetaObject {
	return newQMetaObject(QWindow_MetaObject(this.h))
}

func (this *QWindow) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QWindow_Metacast(this.h, param1_Cstring))
}

func QWindow_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QWindow_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWindow) SetSurfaceType(surfaceType SurfaceType) {
	QWindow_SetSurfaceType(this.h, surfaceType)
}

func (this *QWindow) SurfaceType() SurfaceType {
	xxxxxxxxx
}

func (this *QWindow) IsVisible() bool {
	return (bool)(QWindow_IsVisible(this.h))
}

func (this *QWindow) Visibility() Visibility {
	xxxxxxxxx
}

func (this *QWindow) SetVisibility(v Visibility) {
	QWindow_SetVisibility(this.h, v)
}

func (this *QWindow) Create() {
	QWindow_Create(this.h)
}

func (this *QWindow) WinId() uintptr {
	return (uintptr)(QWindow_WinId(this.h))
}

func (this *QWindow) Parent() *QWindow {
	return newQWindow(QWindow_Parent(this.h))
}

func (this *QWindow) SetParent(parent *QWindow) {
	QWindow_SetParent(this.h, parent.cPointer())
}

func (this *QWindow) IsTopLevel() bool {
	return (bool)(QWindow_IsTopLevel(this.h))
}

func (this *QWindow) IsModal() bool {
	return (bool)(QWindow_IsModal(this.h))
}

func (this *QWindow) Modality() WindowModality {
	return (WindowModality)(QWindow_Modality(this.h))
}

func (this *QWindow) SetModality(modality WindowModality) {
	QWindow_SetModality(this.h, (int)(modality))
}

func (this *QWindow) SetFormat(format *QSurfaceFormat) {
	QWindow_SetFormat(this.h, format.cPointer())
}

func (this *QWindow) Format() *QSurfaceFormat {
	_goptr := newQSurfaceFormat(QWindow_Format(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) RequestedFormat() *QSurfaceFormat {
	_goptr := newQSurfaceFormat(QWindow_RequestedFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) SetFlags(flags WindowType) {
	QWindow_SetFlags(this.h, (int)(flags))
}

func (this *QWindow) Flags() WindowType {
	return (WindowType)(QWindow_Flags(this.h))
}

func (this *QWindow) SetFlag(param1 WindowType) {
	QWindow_SetFlag(this.h, (int)(param1))
}

func (this *QWindow) Type() WindowType {
	return (WindowType)(QWindow_Type(this.h))
}

func (this *QWindow) Title() string {
	var _ms struct_miqt_string = QWindow_Title(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWindow) SetOpacity(level float64) {
	QWindow_SetOpacity(this.h, (double)(level))
}

func (this *QWindow) Opacity() float64 {
	return (float64)(QWindow_Opacity(this.h))
}

func (this *QWindow) SetMask(region *QRegion) {
	QWindow_SetMask(this.h, region.cPointer())
}

func (this *QWindow) Mask() *QRegion {
	_goptr := newQRegion(QWindow_Mask(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) IsActive() bool {
	return (bool)(QWindow_IsActive(this.h))
}

func (this *QWindow) ReportContentOrientationChange(orientation ScreenOrientation) {
	QWindow_ReportContentOrientationChange(this.h, (int)(orientation))
}

func (this *QWindow) ContentOrientation() ScreenOrientation {
	return (ScreenOrientation)(QWindow_ContentOrientation(this.h))
}

func (this *QWindow) DevicePixelRatio() float64 {
	return (float64)(QWindow_DevicePixelRatio(this.h))
}

func (this *QWindow) WindowState() WindowState {
	return (WindowState)(QWindow_WindowState(this.h))
}

func (this *QWindow) WindowStates() WindowState {
	return (WindowState)(QWindow_WindowStates(this.h))
}

func (this *QWindow) SetWindowState(state WindowState) {
	QWindow_SetWindowState(this.h, (int)(state))
}

func (this *QWindow) SetWindowStates(states WindowState) {
	QWindow_SetWindowStates(this.h, (int)(states))
}

func (this *QWindow) SetTransientParent(parent *QWindow) {
	QWindow_SetTransientParent(this.h, parent.cPointer())
}

func (this *QWindow) TransientParent() *QWindow {
	return newQWindow(QWindow_TransientParent(this.h))
}

func (this *QWindow) IsAncestorOf(child *QWindow) bool {
	return (bool)(QWindow_IsAncestorOf(this.h, child.cPointer()))
}

func (this *QWindow) IsExposed() bool {
	return (bool)(QWindow_IsExposed(this.h))
}

func (this *QWindow) MinimumWidth() int {
	return (int)(QWindow_MinimumWidth(this.h))
}

func (this *QWindow) MinimumHeight() int {
	return (int)(QWindow_MinimumHeight(this.h))
}

func (this *QWindow) MaximumWidth() int {
	return (int)(QWindow_MaximumWidth(this.h))
}

func (this *QWindow) MaximumHeight() int {
	return (int)(QWindow_MaximumHeight(this.h))
}

func (this *QWindow) MinimumSize() *QSize {
	_goptr := newQSize(QWindow_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) MaximumSize() *QSize {
	_goptr := newQSize(QWindow_MaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) BaseSize() *QSize {
	_goptr := newQSize(QWindow_BaseSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) SizeIncrement() *QSize {
	_goptr := newQSize(QWindow_SizeIncrement(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) SetMinimumSize(size *QSize) {
	QWindow_SetMinimumSize(this.h, size.cPointer())
}

func (this *QWindow) SetMaximumSize(size *QSize) {
	QWindow_SetMaximumSize(this.h, size.cPointer())
}

func (this *QWindow) SetBaseSize(size *QSize) {
	QWindow_SetBaseSize(this.h, size.cPointer())
}

func (this *QWindow) SetSizeIncrement(size *QSize) {
	QWindow_SetSizeIncrement(this.h, size.cPointer())
}

func (this *QWindow) Geometry() *QRect {
	_goptr := newQRect(QWindow_Geometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) FrameMargins() *QMargins {
	_goptr := newQMargins(QWindow_FrameMargins(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) FrameGeometry() *QRect {
	_goptr := newQRect(QWindow_FrameGeometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) FramePosition() *QPoint {
	_goptr := newQPoint(QWindow_FramePosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) SetFramePosition(point *QPoint) {
	QWindow_SetFramePosition(this.h, point.cPointer())
}

func (this *QWindow) SafeAreaMargins() *QMargins {
	_goptr := newQMargins(QWindow_SafeAreaMargins(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) Width() int {
	return (int)(QWindow_Width(this.h))
}

func (this *QWindow) Height() int {
	return (int)(QWindow_Height(this.h))
}

func (this *QWindow) X() int {
	return (int)(QWindow_X(this.h))
}

func (this *QWindow) Y() int {
	return (int)(QWindow_Y(this.h))
}

func (this *QWindow) Size() *QSize {
	_goptr := newQSize(QWindow_Size(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) Position() *QPoint {
	_goptr := newQPoint(QWindow_Position(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) SetPosition(pt *QPoint) {
	QWindow_SetPosition(this.h, pt.cPointer())
}

func (this *QWindow) SetPosition2(posx int, posy int) {
	QWindow_SetPosition2(this.h, (int)(posx), (int)(posy))
}

func (this *QWindow) Resize(newSize *QSize) {
	QWindow_Resize(this.h, newSize.cPointer())
}

func (this *QWindow) Resize2(w int, h int) {
	QWindow_Resize2(this.h, (int)(w), (int)(h))
}

func (this *QWindow) SetFilePath(filePath string) {
	filePath_ms := struct_miqt_string{}
	filePath_ms.data = CString(filePath)
	filePath_ms.len = size_t(len(filePath))
	defer free(unsafe.Pointer(filePath_ms.data))
	QWindow_SetFilePath(this.h, filePath_ms)
}

func (this *QWindow) FilePath() string {
	var _ms struct_miqt_string = QWindow_FilePath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWindow) SetIcon(icon *QIcon) {
	QWindow_SetIcon(this.h, icon.cPointer())
}

func (this *QWindow) Icon() *QIcon {
	_goptr := newQIcon(QWindow_Icon(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) Destroy() {
	QWindow_Destroy(this.h)
}

func (this *QWindow) SetKeyboardGrabEnabled(grab bool) bool {
	return (bool)(QWindow_SetKeyboardGrabEnabled(this.h, (bool)(grab)))
}

func (this *QWindow) SetMouseGrabEnabled(grab bool) bool {
	return (bool)(QWindow_SetMouseGrabEnabled(this.h, (bool)(grab)))
}

func (this *QWindow) Screen() *QScreen {
	return newQScreen(QWindow_Screen(this.h))
}

func (this *QWindow) SetScreen(screen *QScreen) {
	QWindow_SetScreen(this.h, screen.cPointer())
}

func (this *QWindow) AccessibleRoot() *QAccessibleInterface {
	xxxxxxxxx
}

func (this *QWindow) FocusObject() *QObject {
	return newQObject(QWindow_FocusObject(this.h))
}

func (this *QWindow) MapToGlobal(pos *QPointF) *QPointF {
	_goptr := newQPointF(QWindow_MapToGlobal(this.h, pos.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) MapFromGlobal(pos *QPointF) *QPointF {
	_goptr := newQPointF(QWindow_MapFromGlobal(this.h, pos.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) MapToGlobalWithPos(pos *QPoint) *QPoint {
	_goptr := newQPoint(QWindow_MapToGlobalWithPos(this.h, pos.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) MapFromGlobalWithPos(pos *QPoint) *QPoint {
	_goptr := newQPoint(QWindow_MapFromGlobalWithPos(this.h, pos.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) Cursor() *QCursor {
	_goptr := newQCursor(QWindow_Cursor(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindow) SetCursor(cursor *QCursor) {
	QWindow_SetCursor(this.h, cursor.cPointer())
}

func (this *QWindow) UnsetCursor() {
	QWindow_UnsetCursor(this.h)
}

func QWindow_FromWinId(id uintptr) *QWindow {
	return newQWindow(QWindow_FromWinId((uintptr_t)(id)))
}

func (this *QWindow) RequestActivate() {
	QWindow_RequestActivate(this.h)
}

func (this *QWindow) SetVisible(visible bool) {
	QWindow_SetVisible(this.h, (bool)(visible))
}

func (this *QWindow) Show() {
	QWindow_Show(this.h)
}

func (this *QWindow) Hide() {
	QWindow_Hide(this.h)
}

func (this *QWindow) ShowMinimized() {
	QWindow_ShowMinimized(this.h)
}

func (this *QWindow) ShowMaximized() {
	QWindow_ShowMaximized(this.h)
}

func (this *QWindow) ShowFullScreen() {
	QWindow_ShowFullScreen(this.h)
}

func (this *QWindow) ShowNormal() {
	QWindow_ShowNormal(this.h)
}

func (this *QWindow) Close() bool {
	return (bool)(QWindow_Close(this.h))
}

func (this *QWindow) Raise() {
	QWindow_Raise(this.h)
}

func (this *QWindow) Lower() {
	QWindow_Lower(this.h)
}

func (this *QWindow) StartSystemResize(edges Edge) bool {
	return (bool)(QWindow_StartSystemResize(this.h, (int)(edges)))
}

func (this *QWindow) StartSystemMove() bool {
	return (bool)(QWindow_StartSystemMove(this.h))
}

func (this *QWindow) SetTitle(title string) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	QWindow_SetTitle(this.h, title_ms)
}

func (this *QWindow) SetX(arg int) {
	QWindow_SetX(this.h, (int)(arg))
}

func (this *QWindow) SetY(arg int) {
	QWindow_SetY(this.h, (int)(arg))
}

func (this *QWindow) SetWidth(arg int) {
	QWindow_SetWidth(this.h, (int)(arg))
}

func (this *QWindow) SetHeight(arg int) {
	QWindow_SetHeight(this.h, (int)(arg))
}

func (this *QWindow) SetGeometry(posx int, posy int, w int, h int) {
	QWindow_SetGeometry(this.h, (int)(posx), (int)(posy), (int)(w), (int)(h))
}

func (this *QWindow) SetGeometryWithRect(rect *QRect) {
	QWindow_SetGeometryWithRect(this.h, rect.cPointer())
}

func (this *QWindow) SetMinimumWidth(w int) {
	QWindow_SetMinimumWidth(this.h, (int)(w))
}

func (this *QWindow) SetMinimumHeight(h int) {
	QWindow_SetMinimumHeight(this.h, (int)(h))
}

func (this *QWindow) SetMaximumWidth(w int) {
	QWindow_SetMaximumWidth(this.h, (int)(w))
}

func (this *QWindow) SetMaximumHeight(h int) {
	QWindow_SetMaximumHeight(this.h, (int)(h))
}

func (this *QWindow) Alert(msec int) {
	QWindow_Alert(this.h, (int)(msec))
}

func (this *QWindow) RequestUpdate() {
	QWindow_RequestUpdate(this.h)
}

func (this *QWindow) ScreenChanged(screen *QScreen) {
	QWindow_ScreenChanged(this.h, screen.cPointer())
}

func (this *QWindow) OnScreenChanged(slot func(screen *QScreen)) {
	QWindow_connect_ScreenChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_ScreenChanged
func miqt_exec_callback_QWindow_ScreenChanged(cb intptr_t, screen *QScreen) {
	gofunc, ok := cgo.Handle(cb).Value().(func(screen *QScreen))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQScreen(screen)

	gofunc(slotval1)
}

func (this *QWindow) ModalityChanged(modality WindowModality) {
	QWindow_ModalityChanged(this.h, (int)(modality))
}

func (this *QWindow) OnModalityChanged(slot func(modality WindowModality)) {
	QWindow_connect_ModalityChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_ModalityChanged
func miqt_exec_callback_QWindow_ModalityChanged(cb intptr_t, modality int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(modality WindowModality))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (WindowModality)(modality)

	gofunc(slotval1)
}

func (this *QWindow) WindowStateChanged(windowState WindowState) {
	QWindow_WindowStateChanged(this.h, (int)(windowState))
}

func (this *QWindow) OnWindowStateChanged(slot func(windowState WindowState)) {
	QWindow_connect_WindowStateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_WindowStateChanged
func miqt_exec_callback_QWindow_WindowStateChanged(cb intptr_t, windowState int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(windowState WindowState))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (WindowState)(windowState)

	gofunc(slotval1)
}

func (this *QWindow) WindowTitleChanged(title string) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	QWindow_WindowTitleChanged(this.h, title_ms)
}

func (this *QWindow) OnWindowTitleChanged(slot func(title string)) {
	QWindow_connect_WindowTitleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_WindowTitleChanged
func miqt_exec_callback_QWindow_WindowTitleChanged(cb intptr_t, title struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(title string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var title_ms struct_miqt_string = title
	title_ret := GoStringN(title_ms.data, int(int64(title_ms.len)))
	free(unsafe.Pointer(title_ms.data))
	slotval1 := title_ret

	gofunc(slotval1)
}

func (this *QWindow) XChanged(arg int) {
	QWindow_XChanged(this.h, (int)(arg))
}

func (this *QWindow) OnXChanged(slot func(arg int)) {
	QWindow_connect_XChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_XChanged
func miqt_exec_callback_QWindow_XChanged(cb intptr_t, arg int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(arg int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(arg)

	gofunc(slotval1)
}

func (this *QWindow) YChanged(arg int) {
	QWindow_YChanged(this.h, (int)(arg))
}

func (this *QWindow) OnYChanged(slot func(arg int)) {
	QWindow_connect_YChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_YChanged
func miqt_exec_callback_QWindow_YChanged(cb intptr_t, arg int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(arg int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(arg)

	gofunc(slotval1)
}

func (this *QWindow) WidthChanged(arg int) {
	QWindow_WidthChanged(this.h, (int)(arg))
}

func (this *QWindow) OnWidthChanged(slot func(arg int)) {
	QWindow_connect_WidthChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_WidthChanged
func miqt_exec_callback_QWindow_WidthChanged(cb intptr_t, arg int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(arg int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(arg)

	gofunc(slotval1)
}

func (this *QWindow) HeightChanged(arg int) {
	QWindow_HeightChanged(this.h, (int)(arg))
}

func (this *QWindow) OnHeightChanged(slot func(arg int)) {
	QWindow_connect_HeightChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_HeightChanged
func miqt_exec_callback_QWindow_HeightChanged(cb intptr_t, arg int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(arg int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(arg)

	gofunc(slotval1)
}

func (this *QWindow) MinimumWidthChanged(arg int) {
	QWindow_MinimumWidthChanged(this.h, (int)(arg))
}

func (this *QWindow) OnMinimumWidthChanged(slot func(arg int)) {
	QWindow_connect_MinimumWidthChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_MinimumWidthChanged
func miqt_exec_callback_QWindow_MinimumWidthChanged(cb intptr_t, arg int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(arg int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(arg)

	gofunc(slotval1)
}

func (this *QWindow) MinimumHeightChanged(arg int) {
	QWindow_MinimumHeightChanged(this.h, (int)(arg))
}

func (this *QWindow) OnMinimumHeightChanged(slot func(arg int)) {
	QWindow_connect_MinimumHeightChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_MinimumHeightChanged
func miqt_exec_callback_QWindow_MinimumHeightChanged(cb intptr_t, arg int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(arg int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(arg)

	gofunc(slotval1)
}

func (this *QWindow) MaximumWidthChanged(arg int) {
	QWindow_MaximumWidthChanged(this.h, (int)(arg))
}

func (this *QWindow) OnMaximumWidthChanged(slot func(arg int)) {
	QWindow_connect_MaximumWidthChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_MaximumWidthChanged
func miqt_exec_callback_QWindow_MaximumWidthChanged(cb intptr_t, arg int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(arg int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(arg)

	gofunc(slotval1)
}

func (this *QWindow) MaximumHeightChanged(arg int) {
	QWindow_MaximumHeightChanged(this.h, (int)(arg))
}

func (this *QWindow) OnMaximumHeightChanged(slot func(arg int)) {
	QWindow_connect_MaximumHeightChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_MaximumHeightChanged
func miqt_exec_callback_QWindow_MaximumHeightChanged(cb intptr_t, arg int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(arg int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(arg)

	gofunc(slotval1)
}

func (this *QWindow) SafeAreaMarginsChanged(arg QMargins) {
	QWindow_SafeAreaMarginsChanged(this.h, arg.cPointer())
}

func (this *QWindow) OnSafeAreaMarginsChanged(slot func(arg QMargins)) {
	QWindow_connect_SafeAreaMarginsChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_SafeAreaMarginsChanged
func miqt_exec_callback_QWindow_SafeAreaMarginsChanged(cb intptr_t, arg *QMargins) {
	gofunc, ok := cgo.Handle(cb).Value().(func(arg QMargins))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	arg_goptr := newQMargins(arg)
	arg_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval1 := *arg_goptr

	gofunc(slotval1)
}

func (this *QWindow) VisibleChanged(arg bool) {
	QWindow_VisibleChanged(this.h, (bool)(arg))
}

func (this *QWindow) OnVisibleChanged(slot func(arg bool)) {
	QWindow_connect_VisibleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_VisibleChanged
func miqt_exec_callback_QWindow_VisibleChanged(cb intptr_t, arg bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(arg bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(arg)

	gofunc(slotval1)
}

func (this *QWindow) VisibilityChanged(visibility QWindow__Visibility) {
	QWindow_VisibilityChanged(this.h, (int)(visibility))
}

func (this *QWindow) OnVisibilityChanged(slot func(visibility QWindow__Visibility)) {
	QWindow_connect_VisibilityChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_VisibilityChanged
func miqt_exec_callback_QWindow_VisibilityChanged(cb intptr_t, visibility int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(visibility QWindow__Visibility))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QWindow__Visibility)(visibility)

	gofunc(slotval1)
}

func (this *QWindow) ActiveChanged() {
	QWindow_ActiveChanged(this.h)
}

func (this *QWindow) OnActiveChanged(slot func()) {
	QWindow_connect_ActiveChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_ActiveChanged
func miqt_exec_callback_QWindow_ActiveChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QWindow) ContentOrientationChanged(orientation ScreenOrientation) {
	QWindow_ContentOrientationChanged(this.h, (int)(orientation))
}

func (this *QWindow) OnContentOrientationChanged(slot func(orientation ScreenOrientation)) {
	QWindow_connect_ContentOrientationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_ContentOrientationChanged
func miqt_exec_callback_QWindow_ContentOrientationChanged(cb intptr_t, orientation int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(orientation ScreenOrientation))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (ScreenOrientation)(orientation)

	gofunc(slotval1)
}

func (this *QWindow) FocusObjectChanged(object *QObject) {
	QWindow_FocusObjectChanged(this.h, object.cPointer())
}

func (this *QWindow) OnFocusObjectChanged(slot func(object *QObject)) {
	QWindow_connect_FocusObjectChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_FocusObjectChanged
func miqt_exec_callback_QWindow_FocusObjectChanged(cb intptr_t, object *QObject) {
	gofunc, ok := cgo.Handle(cb).Value().(func(object *QObject))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(object)

	gofunc(slotval1)
}

func (this *QWindow) OpacityChanged(opacity float64) {
	QWindow_OpacityChanged(this.h, (double)(opacity))
}

func (this *QWindow) OnOpacityChanged(slot func(opacity float64)) {
	QWindow_connect_OpacityChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_OpacityChanged
func miqt_exec_callback_QWindow_OpacityChanged(cb intptr_t, opacity double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(opacity float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float64)(opacity)

	gofunc(slotval1)
}

func (this *QWindow) TransientParentChanged(transientParent *QWindow) {
	QWindow_TransientParentChanged(this.h, transientParent.cPointer())
}

func (this *QWindow) OnTransientParentChanged(slot func(transientParent *QWindow)) {
	QWindow_connect_TransientParentChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_TransientParentChanged
func miqt_exec_callback_QWindow_TransientParentChanged(cb intptr_t, transientParent *QWindow) {
	gofunc, ok := cgo.Handle(cb).Value().(func(transientParent *QWindow))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWindow(transientParent)

	gofunc(slotval1)
}

func QWindow_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWindow_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWindow_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWindow_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWindow) Parent1(mode AncestorMode) *QWindow {
	return newQWindow(QWindow_Parent1(this.h, mode))
}

func (this *QWindow) SetFlag2(param1 WindowType, on bool) {
	QWindow_SetFlag2(this.h, (int)(param1), (bool)(on))
}

func (this *QWindow) IsAncestorOf2(child *QWindow, mode AncestorMode) bool {
	return (bool)(QWindow_IsAncestorOf2(this.h, child.cPointer(), mode))
}

func (this *QWindow) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QWindow_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QWindow) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_MetaObject
func miqt_exec_callback_QWindow_MetaObject(self QWindow, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWindow{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QWindow) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QWindow_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QWindow) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_Metacast
func miqt_exec_callback_QWindow_Metacast(self QWindow, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QWindow{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
