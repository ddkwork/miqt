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

func (this *QWindow) callVirtualBase_SurfaceType() SurfaceType {

	xxxxxxxxx
}
func (this *QWindow) OnSurfaceType(slot func(super func() SurfaceType) SurfaceType) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_SurfaceType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_SurfaceType
func miqt_exec_callback_QWindow_SurfaceType(self QWindow, cb intptr_t) SurfaceType {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() SurfaceType) SurfaceType)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWindow{h: self}).callVirtualBase_SurfaceType)

	return virtualReturn

}

func (this *QWindow) callVirtualBase_Format() *QSurfaceFormat {

	_goptr := newQSurfaceFormat(QWindow_virtualbase_Format(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWindow) OnFormat(slot func(super func() *QSurfaceFormat) *QSurfaceFormat) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_Format(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_Format
func miqt_exec_callback_QWindow_Format(self QWindow, cb intptr_t) *QSurfaceFormat {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSurfaceFormat) *QSurfaceFormat)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWindow{h: self}).callVirtualBase_Format)

	return virtualReturn.cPointer()

}

func (this *QWindow) callVirtualBase_Size() *QSize {

	_goptr := newQSize(QWindow_virtualbase_Size(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWindow) OnSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_Size(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_Size
func miqt_exec_callback_QWindow_Size(self QWindow, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWindow{h: self}).callVirtualBase_Size)

	return virtualReturn.cPointer()

}

func (this *QWindow) callVirtualBase_AccessibleRoot() *QAccessibleInterface {

	xxxxxxxxx
}
func (this *QWindow) OnAccessibleRoot(slot func(super func() *QAccessibleInterface) *QAccessibleInterface) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_AccessibleRoot(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_AccessibleRoot
func miqt_exec_callback_QWindow_AccessibleRoot(self QWindow, cb intptr_t) *QAccessibleInterface {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QAccessibleInterface) *QAccessibleInterface)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWindow{h: self}).callVirtualBase_AccessibleRoot)

	return virtualReturn

}

func (this *QWindow) callVirtualBase_FocusObject() *QObject {

	return newQObject(QWindow_virtualbase_FocusObject(unsafe.Pointer(this.h)))

}
func (this *QWindow) OnFocusObject(slot func(super func() *QObject) *QObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_FocusObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_FocusObject
func miqt_exec_callback_QWindow_FocusObject(self QWindow, cb intptr_t) *QObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QObject) *QObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWindow{h: self}).callVirtualBase_FocusObject)

	return virtualReturn.cPointer()

}

func (this *QWindow) callVirtualBase_ExposeEvent(param1 *QExposeEvent) {

	QWindow_virtualbase_ExposeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnExposeEvent(slot func(super func(param1 *QExposeEvent), param1 *QExposeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_ExposeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_ExposeEvent
func miqt_exec_callback_QWindow_ExposeEvent(self QWindow, cb intptr_t, param1 *QExposeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QExposeEvent), param1 *QExposeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQExposeEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_ExposeEvent, slotval1)

}

func (this *QWindow) callVirtualBase_ResizeEvent(param1 *QResizeEvent) {

	QWindow_virtualbase_ResizeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnResizeEvent(slot func(super func(param1 *QResizeEvent), param1 *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_ResizeEvent
func miqt_exec_callback_QWindow_ResizeEvent(self QWindow, cb intptr_t, param1 *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QResizeEvent), param1 *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QWindow) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QWindow_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_PaintEvent
func miqt_exec_callback_QWindow_PaintEvent(self QWindow, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QWindow) callVirtualBase_MoveEvent(param1 *QMoveEvent) {

	QWindow_virtualbase_MoveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnMoveEvent(slot func(super func(param1 *QMoveEvent), param1 *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_MoveEvent
func miqt_exec_callback_QWindow_MoveEvent(self QWindow, cb intptr_t, param1 *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMoveEvent), param1 *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QWindow) callVirtualBase_FocusInEvent(param1 *QFocusEvent) {

	QWindow_virtualbase_FocusInEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnFocusInEvent(slot func(super func(param1 *QFocusEvent), param1 *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_FocusInEvent
func miqt_exec_callback_QWindow_FocusInEvent(self QWindow, cb intptr_t, param1 *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QFocusEvent), param1 *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QWindow) callVirtualBase_FocusOutEvent(param1 *QFocusEvent) {

	QWindow_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnFocusOutEvent(slot func(super func(param1 *QFocusEvent), param1 *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_FocusOutEvent
func miqt_exec_callback_QWindow_FocusOutEvent(self QWindow, cb intptr_t, param1 *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QFocusEvent), param1 *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QWindow) callVirtualBase_ShowEvent(param1 *QShowEvent) {

	QWindow_virtualbase_ShowEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnShowEvent(slot func(super func(param1 *QShowEvent), param1 *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_ShowEvent
func miqt_exec_callback_QWindow_ShowEvent(self QWindow, cb intptr_t, param1 *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QShowEvent), param1 *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QWindow) callVirtualBase_HideEvent(param1 *QHideEvent) {

	QWindow_virtualbase_HideEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnHideEvent(slot func(super func(param1 *QHideEvent), param1 *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_HideEvent
func miqt_exec_callback_QWindow_HideEvent(self QWindow, cb intptr_t, param1 *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QHideEvent), param1 *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QWindow) callVirtualBase_CloseEvent(param1 *QCloseEvent) {

	QWindow_virtualbase_CloseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnCloseEvent(slot func(super func(param1 *QCloseEvent), param1 *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_CloseEvent
func miqt_exec_callback_QWindow_CloseEvent(self QWindow, cb intptr_t, param1 *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QCloseEvent), param1 *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QWindow) callVirtualBase_Event(param1 *QEvent) bool {

	return (bool)(QWindow_virtualbase_Event(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QWindow) OnEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_Event
func miqt_exec_callback_QWindow_Event(self QWindow, cb intptr_t, param1 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QWindow{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QWindow) callVirtualBase_KeyPressEvent(param1 *QKeyEvent) {

	QWindow_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnKeyPressEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_KeyPressEvent
func miqt_exec_callback_QWindow_KeyPressEvent(self QWindow, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QWindow) callVirtualBase_KeyReleaseEvent(param1 *QKeyEvent) {

	QWindow_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnKeyReleaseEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_KeyReleaseEvent
func miqt_exec_callback_QWindow_KeyReleaseEvent(self QWindow, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QWindow) callVirtualBase_MousePressEvent(param1 *QMouseEvent) {

	QWindow_virtualbase_MousePressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnMousePressEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_MousePressEvent
func miqt_exec_callback_QWindow_MousePressEvent(self QWindow, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QWindow) callVirtualBase_MouseReleaseEvent(param1 *QMouseEvent) {

	QWindow_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnMouseReleaseEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_MouseReleaseEvent
func miqt_exec_callback_QWindow_MouseReleaseEvent(self QWindow, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QWindow) callVirtualBase_MouseDoubleClickEvent(param1 *QMouseEvent) {

	QWindow_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnMouseDoubleClickEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_MouseDoubleClickEvent
func miqt_exec_callback_QWindow_MouseDoubleClickEvent(self QWindow, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QWindow) callVirtualBase_MouseMoveEvent(param1 *QMouseEvent) {

	QWindow_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnMouseMoveEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_MouseMoveEvent
func miqt_exec_callback_QWindow_MouseMoveEvent(self QWindow, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QWindow) callVirtualBase_WheelEvent(param1 *QWheelEvent) {

	QWindow_virtualbase_WheelEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnWheelEvent(slot func(super func(param1 *QWheelEvent), param1 *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_WheelEvent
func miqt_exec_callback_QWindow_WheelEvent(self QWindow, cb intptr_t, param1 *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QWheelEvent), param1 *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QWindow) callVirtualBase_TouchEvent(param1 *QTouchEvent) {

	QWindow_virtualbase_TouchEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnTouchEvent(slot func(super func(param1 *QTouchEvent), param1 *QTouchEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_TouchEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_TouchEvent
func miqt_exec_callback_QWindow_TouchEvent(self QWindow, cb intptr_t, param1 *QTouchEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QTouchEvent), param1 *QTouchEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTouchEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_TouchEvent, slotval1)

}

func (this *QWindow) callVirtualBase_TabletEvent(param1 *QTabletEvent) {

	QWindow_virtualbase_TabletEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWindow) OnTabletEvent(slot func(super func(param1 *QTabletEvent), param1 *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_TabletEvent
func miqt_exec_callback_QWindow_TabletEvent(self QWindow, cb intptr_t, param1 *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QTabletEvent), param1 *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(param1)

	gofunc((&QWindow{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QWindow) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QWindow_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QWindow) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_NativeEvent
func miqt_exec_callback_QWindow_NativeEvent(self QWindow, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
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

	virtualReturn := gofunc((&QWindow{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QWindow) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QWindow_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QWindow) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_EventFilter
func miqt_exec_callback_QWindow_EventFilter(self QWindow, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QWindow{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QWindow) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QWindow_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWindow) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_TimerEvent
func miqt_exec_callback_QWindow_TimerEvent(self QWindow, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QWindow{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QWindow) callVirtualBase_ChildEvent(event *QChildEvent) {

	QWindow_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWindow) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_ChildEvent
func miqt_exec_callback_QWindow_ChildEvent(self QWindow, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QWindow{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QWindow) callVirtualBase_CustomEvent(event *QEvent) {

	QWindow_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWindow) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_CustomEvent
func miqt_exec_callback_QWindow_CustomEvent(self QWindow, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QWindow{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QWindow) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QWindow_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QWindow) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_ConnectNotify
func miqt_exec_callback_QWindow_ConnectNotify(self QWindow, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QWindow{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QWindow) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QWindow_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QWindow) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindow_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindow_DisconnectNotify
func miqt_exec_callback_QWindow_DisconnectNotify(self QWindow, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QWindow{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
