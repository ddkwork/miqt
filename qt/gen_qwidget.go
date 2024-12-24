package qt

import (
	"unsafe"
)

type QWidget__RenderFlag int

const (
	QWidget__DrawWindowBackground QWidget__RenderFlag = 1
	QWidget__DrawChildren         QWidget__RenderFlag = 2
	QWidget__IgnoreMask           QWidget__RenderFlag = 4
)

type QWidgetData struct {
	h          uintptr
	isSubclass bool
}

// NewQWidgetData constructs a new QWidgetData object.
func NewQWidgetData(param1 *QWidgetData) *QWidgetData {
	g := newQWidgetData(QWidgetData_new(param1.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QWidgetData) OperatorAssign(param1 *QWidgetData) {
	QWidgetData_OperatorAssign(this.h, param1.cPointer())
}

type QWidget struct {
	h          uintptr
	isSubclass bool
}

// NewQWidget constructs a new QWidget object.
func NewQWidget(parent *QWidget) *QWidget {
	g := newQWidget(QWidget_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQWidget2 constructs a new QWidget object.
func NewQWidget2() *QWidget {
	g := newQWidget(QWidget_new2())
	g.isSubclass = true
	return g
}

// NewQWidget3 constructs a new QWidget object.
func NewQWidget3(parent *QWidget, f WindowType) *QWidget {
	g := newQWidget(QWidget_new3(parent.cPointer(), (int)(f)))
	g.isSubclass = true
	return g
}

func (this *QWidget) MetaObject() *QMetaObject {
	return newQMetaObject(QWidget_MetaObject(this.h))
}

func (this *QWidget) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QWidget_Metacast(this.h, param1_Cstring))
}

func QWidget_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QWidget_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWidget) DevType() int {
	return (int)(QWidget_DevType(this.h))
}

func (this *QWidget) WinId() uintptr {
	return (uintptr)(QWidget_WinId(this.h))
}

func (this *QWidget) CreateWinId() {
	QWidget_CreateWinId(this.h)
}

func (this *QWidget) InternalWinId() uintptr {
	return (uintptr)(QWidget_InternalWinId(this.h))
}

func (this *QWidget) EffectiveWinId() uintptr {
	return (uintptr)(QWidget_EffectiveWinId(this.h))
}

func (this *QWidget) Style() *QStyle {
	return newQStyle(QWidget_Style(this.h))
}

func (this *QWidget) SetStyle(style *QStyle) {
	QWidget_SetStyle(this.h, style.cPointer())
}

func (this *QWidget) IsTopLevel() bool {
	return (bool)(QWidget_IsTopLevel(this.h))
}

func (this *QWidget) IsWindow() bool {
	return (bool)(QWidget_IsWindow(this.h))
}

func (this *QWidget) IsModal() bool {
	return (bool)(QWidget_IsModal(this.h))
}

func (this *QWidget) WindowModality() WindowModality {
	return (WindowModality)(QWidget_WindowModality(this.h))
}

func (this *QWidget) SetWindowModality(windowModality WindowModality) {
	QWidget_SetWindowModality(this.h, (int)(windowModality))
}

func (this *QWidget) IsEnabled() bool {
	return (bool)(QWidget_IsEnabled(this.h))
}

func (this *QWidget) IsEnabledTo(param1 *QWidget) bool {
	return (bool)(QWidget_IsEnabledTo(this.h, param1.cPointer()))
}

func (this *QWidget) SetEnabled(enabled bool) {
	QWidget_SetEnabled(this.h, (bool)(enabled))
}

func (this *QWidget) SetDisabled(disabled bool) {
	QWidget_SetDisabled(this.h, (bool)(disabled))
}

func (this *QWidget) SetWindowModified(windowModified bool) {
	QWidget_SetWindowModified(this.h, (bool)(windowModified))
}

func (this *QWidget) FrameGeometry() *QRect {
	_goptr := newQRect(QWidget_FrameGeometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) Geometry() *QRect {
	return newQRect(QWidget_Geometry(this.h))
}

func (this *QWidget) NormalGeometry() *QRect {
	_goptr := newQRect(QWidget_NormalGeometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) X() int {
	return (int)(QWidget_X(this.h))
}

func (this *QWidget) Y() int {
	return (int)(QWidget_Y(this.h))
}

func (this *QWidget) Pos() *QPoint {
	_goptr := newQPoint(QWidget_Pos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) FrameSize() *QSize {
	_goptr := newQSize(QWidget_FrameSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) Size() *QSize {
	_goptr := newQSize(QWidget_Size(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) Width() int {
	return (int)(QWidget_Width(this.h))
}

func (this *QWidget) Height() int {
	return (int)(QWidget_Height(this.h))
}

func (this *QWidget) Rect() *QRect {
	_goptr := newQRect(QWidget_Rect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) ChildrenRect() *QRect {
	_goptr := newQRect(QWidget_ChildrenRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) ChildrenRegion() *QRegion {
	_goptr := newQRegion(QWidget_ChildrenRegion(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) MinimumSize() *QSize {
	_goptr := newQSize(QWidget_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) MaximumSize() *QSize {
	_goptr := newQSize(QWidget_MaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) MinimumWidth() int {
	return (int)(QWidget_MinimumWidth(this.h))
}

func (this *QWidget) MinimumHeight() int {
	return (int)(QWidget_MinimumHeight(this.h))
}

func (this *QWidget) MaximumWidth() int {
	return (int)(QWidget_MaximumWidth(this.h))
}

func (this *QWidget) MaximumHeight() int {
	return (int)(QWidget_MaximumHeight(this.h))
}

func (this *QWidget) SetMinimumSize(minimumSize *QSize) {
	QWidget_SetMinimumSize(this.h, minimumSize.cPointer())
}

func (this *QWidget) SetMinimumSize2(minw int, minh int) {
	QWidget_SetMinimumSize2(this.h, (int)(minw), (int)(minh))
}

func (this *QWidget) SetMaximumSize(maximumSize *QSize) {
	QWidget_SetMaximumSize(this.h, maximumSize.cPointer())
}

func (this *QWidget) SetMaximumSize2(maxw int, maxh int) {
	QWidget_SetMaximumSize2(this.h, (int)(maxw), (int)(maxh))
}

func (this *QWidget) SetMinimumWidth(minw int) {
	QWidget_SetMinimumWidth(this.h, (int)(minw))
}

func (this *QWidget) SetMinimumHeight(minh int) {
	QWidget_SetMinimumHeight(this.h, (int)(minh))
}

func (this *QWidget) SetMaximumWidth(maxw int) {
	QWidget_SetMaximumWidth(this.h, (int)(maxw))
}

func (this *QWidget) SetMaximumHeight(maxh int) {
	QWidget_SetMaximumHeight(this.h, (int)(maxh))
}

func (this *QWidget) SizeIncrement() *QSize {
	_goptr := newQSize(QWidget_SizeIncrement(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) SetSizeIncrement(sizeIncrement *QSize) {
	QWidget_SetSizeIncrement(this.h, sizeIncrement.cPointer())
}

func (this *QWidget) SetSizeIncrement2(w int, h int) {
	QWidget_SetSizeIncrement2(this.h, (int)(w), (int)(h))
}

func (this *QWidget) BaseSize() *QSize {
	_goptr := newQSize(QWidget_BaseSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) SetBaseSize(baseSize *QSize) {
	QWidget_SetBaseSize(this.h, baseSize.cPointer())
}

func (this *QWidget) SetBaseSize2(basew int, baseh int) {
	QWidget_SetBaseSize2(this.h, (int)(basew), (int)(baseh))
}

func (this *QWidget) SetFixedSize(fixedSize *QSize) {
	QWidget_SetFixedSize(this.h, fixedSize.cPointer())
}

func (this *QWidget) SetFixedSize2(w int, h int) {
	QWidget_SetFixedSize2(this.h, (int)(w), (int)(h))
}

func (this *QWidget) SetFixedWidth(w int) {
	QWidget_SetFixedWidth(this.h, (int)(w))
}

func (this *QWidget) SetFixedHeight(h int) {
	QWidget_SetFixedHeight(this.h, (int)(h))
}

func (this *QWidget) MapToGlobal(param1 *QPointF) *QPointF {
	_goptr := newQPointF(QWidget_MapToGlobal(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) MapToGlobalWithQPoint(param1 *QPoint) *QPoint {
	_goptr := newQPoint(QWidget_MapToGlobalWithQPoint(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) MapFromGlobal(param1 *QPointF) *QPointF {
	_goptr := newQPointF(QWidget_MapFromGlobal(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) MapFromGlobalWithQPoint(param1 *QPoint) *QPoint {
	_goptr := newQPoint(QWidget_MapFromGlobalWithQPoint(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) MapToParent(param1 *QPointF) *QPointF {
	_goptr := newQPointF(QWidget_MapToParent(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) MapToParentWithQPoint(param1 *QPoint) *QPoint {
	_goptr := newQPoint(QWidget_MapToParentWithQPoint(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) MapFromParent(param1 *QPointF) *QPointF {
	_goptr := newQPointF(QWidget_MapFromParent(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) MapFromParentWithQPoint(param1 *QPoint) *QPoint {
	_goptr := newQPoint(QWidget_MapFromParentWithQPoint(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) MapTo(param1 *QWidget, param2 *QPointF) *QPointF {
	_goptr := newQPointF(QWidget_MapTo(this.h, param1.cPointer(), param2.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) MapTo2(param1 *QWidget, param2 *QPoint) *QPoint {
	_goptr := newQPoint(QWidget_MapTo2(this.h, param1.cPointer(), param2.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) MapFrom(param1 *QWidget, param2 *QPointF) *QPointF {
	_goptr := newQPointF(QWidget_MapFrom(this.h, param1.cPointer(), param2.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) MapFrom2(param1 *QWidget, param2 *QPoint) *QPoint {
	_goptr := newQPoint(QWidget_MapFrom2(this.h, param1.cPointer(), param2.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) Window() *QWidget {
	return newQWidget(QWidget_Window(this.h))
}

func (this *QWidget) NativeParentWidget() *QWidget {
	return newQWidget(QWidget_NativeParentWidget(this.h))
}

func (this *QWidget) TopLevelWidget() *QWidget {
	return newQWidget(QWidget_TopLevelWidget(this.h))
}

func (this *QWidget) Palette() *QPalette {
	return newQPalette(QWidget_Palette(this.h))
}

func (this *QWidget) SetPalette(palette *QPalette) {
	QWidget_SetPalette(this.h, palette.cPointer())
}

func (this *QWidget) SetBackgroundRole(backgroundRole QPalette__ColorRole) {
	QWidget_SetBackgroundRole(this.h, (int)(backgroundRole))
}

func (this *QWidget) BackgroundRole() QPalette__ColorRole {
	return (QPalette__ColorRole)(QWidget_BackgroundRole(this.h))
}

func (this *QWidget) SetForegroundRole(foregroundRole QPalette__ColorRole) {
	QWidget_SetForegroundRole(this.h, (int)(foregroundRole))
}

func (this *QWidget) ForegroundRole() QPalette__ColorRole {
	return (QPalette__ColorRole)(QWidget_ForegroundRole(this.h))
}

func (this *QWidget) Font() *QFont {
	return newQFont(QWidget_Font(this.h))
}

func (this *QWidget) SetFont(font *QFont) {
	QWidget_SetFont(this.h, font.cPointer())
}

func (this *QWidget) FontMetrics() *QFontMetrics {
	_goptr := newQFontMetrics(QWidget_FontMetrics(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) FontInfo() *QFontInfo {
	_goptr := newQFontInfo(QWidget_FontInfo(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) Cursor() *QCursor {
	_goptr := newQCursor(QWidget_Cursor(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) SetCursor(cursor *QCursor) {
	QWidget_SetCursor(this.h, cursor.cPointer())
}

func (this *QWidget) UnsetCursor() {
	QWidget_UnsetCursor(this.h)
}

func (this *QWidget) SetMouseTracking(enable bool) {
	QWidget_SetMouseTracking(this.h, (bool)(enable))
}

func (this *QWidget) HasMouseTracking() bool {
	return (bool)(QWidget_HasMouseTracking(this.h))
}

func (this *QWidget) UnderMouse() bool {
	return (bool)(QWidget_UnderMouse(this.h))
}

func (this *QWidget) SetTabletTracking(enable bool) {
	QWidget_SetTabletTracking(this.h, (bool)(enable))
}

func (this *QWidget) HasTabletTracking() bool {
	return (bool)(QWidget_HasTabletTracking(this.h))
}

func (this *QWidget) SetMask(mask *QBitmap) {
	QWidget_SetMask(this.h, mask.cPointer())
}

func (this *QWidget) SetMaskWithMask(mask *QRegion) {
	QWidget_SetMaskWithMask(this.h, mask.cPointer())
}

func (this *QWidget) Mask() *QRegion {
	_goptr := newQRegion(QWidget_Mask(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) ClearMask() {
	QWidget_ClearMask(this.h)
}

func (this *QWidget) Render(target *QPaintDevice) {
	QWidget_Render(this.h, target.cPointer())
}

func (this *QWidget) RenderWithPainter(painter *QPainter) {
	QWidget_RenderWithPainter(this.h, painter.cPointer())
}

func (this *QWidget) Grab() *QPixmap {
	_goptr := newQPixmap(QWidget_Grab(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) GraphicsEffect() *QGraphicsEffect {
	return newQGraphicsEffect(QWidget_GraphicsEffect(this.h))
}

func (this *QWidget) SetGraphicsEffect(effect *QGraphicsEffect) {
	QWidget_SetGraphicsEffect(this.h, effect.cPointer())
}

func (this *QWidget) GrabGesture(typeVal GestureType) {
	QWidget_GrabGesture(this.h, (int)(typeVal))
}

func (this *QWidget) UngrabGesture(typeVal GestureType) {
	QWidget_UngrabGesture(this.h, (int)(typeVal))
}

func (this *QWidget) SetWindowTitle(windowTitle string) {
	windowTitle_ms := struct_miqt_string{}
	windowTitle_ms.data = CString(windowTitle)
	windowTitle_ms.len = size_t(len(windowTitle))
	defer free(unsafe.Pointer(windowTitle_ms.data))
	QWidget_SetWindowTitle(this.h, windowTitle_ms)
}

func (this *QWidget) SetStyleSheet(styleSheet string) {
	styleSheet_ms := struct_miqt_string{}
	styleSheet_ms.data = CString(styleSheet)
	styleSheet_ms.len = size_t(len(styleSheet))
	defer free(unsafe.Pointer(styleSheet_ms.data))
	QWidget_SetStyleSheet(this.h, styleSheet_ms)
}

func (this *QWidget) StyleSheet() string {
	var _ms struct_miqt_string = QWidget_StyleSheet(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWidget) WindowTitle() string {
	var _ms struct_miqt_string = QWidget_WindowTitle(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWidget) SetWindowIcon(icon *QIcon) {
	QWidget_SetWindowIcon(this.h, icon.cPointer())
}

func (this *QWidget) WindowIcon() *QIcon {
	_goptr := newQIcon(QWidget_WindowIcon(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) SetWindowIconText(windowIconText string) {
	windowIconText_ms := struct_miqt_string{}
	windowIconText_ms.data = CString(windowIconText)
	windowIconText_ms.len = size_t(len(windowIconText))
	defer free(unsafe.Pointer(windowIconText_ms.data))
	QWidget_SetWindowIconText(this.h, windowIconText_ms)
}

func (this *QWidget) WindowIconText() string {
	var _ms struct_miqt_string = QWidget_WindowIconText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWidget) SetWindowRole(windowRole string) {
	windowRole_ms := struct_miqt_string{}
	windowRole_ms.data = CString(windowRole)
	windowRole_ms.len = size_t(len(windowRole))
	defer free(unsafe.Pointer(windowRole_ms.data))
	QWidget_SetWindowRole(this.h, windowRole_ms)
}

func (this *QWidget) WindowRole() string {
	var _ms struct_miqt_string = QWidget_WindowRole(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWidget) SetWindowFilePath(filePath string) {
	filePath_ms := struct_miqt_string{}
	filePath_ms.data = CString(filePath)
	filePath_ms.len = size_t(len(filePath))
	defer free(unsafe.Pointer(filePath_ms.data))
	QWidget_SetWindowFilePath(this.h, filePath_ms)
}

func (this *QWidget) WindowFilePath() string {
	var _ms struct_miqt_string = QWidget_WindowFilePath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWidget) SetWindowOpacity(level float64) {
	QWidget_SetWindowOpacity(this.h, (double)(level))
}

func (this *QWidget) WindowOpacity() float64 {
	return (float64)(QWidget_WindowOpacity(this.h))
}

func (this *QWidget) IsWindowModified() bool {
	return (bool)(QWidget_IsWindowModified(this.h))
}

func (this *QWidget) SetToolTip(toolTip string) {
	toolTip_ms := struct_miqt_string{}
	toolTip_ms.data = CString(toolTip)
	toolTip_ms.len = size_t(len(toolTip))
	defer free(unsafe.Pointer(toolTip_ms.data))
	QWidget_SetToolTip(this.h, toolTip_ms)
}

func (this *QWidget) ToolTip() string {
	var _ms struct_miqt_string = QWidget_ToolTip(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWidget) SetToolTipDuration(msec int) {
	QWidget_SetToolTipDuration(this.h, (int)(msec))
}

func (this *QWidget) ToolTipDuration() int {
	return (int)(QWidget_ToolTipDuration(this.h))
}

func (this *QWidget) SetStatusTip(statusTip string) {
	statusTip_ms := struct_miqt_string{}
	statusTip_ms.data = CString(statusTip)
	statusTip_ms.len = size_t(len(statusTip))
	defer free(unsafe.Pointer(statusTip_ms.data))
	QWidget_SetStatusTip(this.h, statusTip_ms)
}

func (this *QWidget) StatusTip() string {
	var _ms struct_miqt_string = QWidget_StatusTip(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWidget) SetWhatsThis(whatsThis string) {
	whatsThis_ms := struct_miqt_string{}
	whatsThis_ms.data = CString(whatsThis)
	whatsThis_ms.len = size_t(len(whatsThis))
	defer free(unsafe.Pointer(whatsThis_ms.data))
	QWidget_SetWhatsThis(this.h, whatsThis_ms)
}

func (this *QWidget) WhatsThis() string {
	var _ms struct_miqt_string = QWidget_WhatsThis(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWidget) SetLayoutDirection(direction LayoutDirection) {
	QWidget_SetLayoutDirection(this.h, (int)(direction))
}

func (this *QWidget) LayoutDirection() LayoutDirection {
	return (LayoutDirection)(QWidget_LayoutDirection(this.h))
}

func (this *QWidget) UnsetLayoutDirection() {
	QWidget_UnsetLayoutDirection(this.h)
}

func (this *QWidget) SetLocale(locale *QLocale) {
	QWidget_SetLocale(this.h, locale.cPointer())
}

func (this *QWidget) Locale() *QLocale {
	_goptr := newQLocale(QWidget_Locale(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) UnsetLocale() {
	QWidget_UnsetLocale(this.h)
}

func (this *QWidget) IsRightToLeft() bool {
	return (bool)(QWidget_IsRightToLeft(this.h))
}

func (this *QWidget) IsLeftToRight() bool {
	return (bool)(QWidget_IsLeftToRight(this.h))
}

func (this *QWidget) SetFocus() {
	QWidget_SetFocus(this.h)
}

func (this *QWidget) IsActiveWindow() bool {
	return (bool)(QWidget_IsActiveWindow(this.h))
}

func (this *QWidget) ActivateWindow() {
	QWidget_ActivateWindow(this.h)
}

func (this *QWidget) ClearFocus() {
	QWidget_ClearFocus(this.h)
}

func (this *QWidget) SetFocusWithReason(reason FocusReason) {
	QWidget_SetFocusWithReason(this.h, (int)(reason))
}

func (this *QWidget) FocusPolicy() FocusPolicy {
	return (FocusPolicy)(QWidget_FocusPolicy(this.h))
}

func (this *QWidget) SetFocusPolicy(policy FocusPolicy) {
	QWidget_SetFocusPolicy(this.h, (int)(policy))
}

func (this *QWidget) HasFocus() bool {
	return (bool)(QWidget_HasFocus(this.h))
}

func QWidget_SetTabOrder(param1 *QWidget, param2 *QWidget) {
	QWidget_SetTabOrder(param1.cPointer(), param2.cPointer())
}

func (this *QWidget) SetFocusProxy(focusProxy *QWidget) {
	QWidget_SetFocusProxy(this.h, focusProxy.cPointer())
}

func (this *QWidget) FocusProxy() *QWidget {
	return newQWidget(QWidget_FocusProxy(this.h))
}

func (this *QWidget) ContextMenuPolicy() ContextMenuPolicy {
	return (ContextMenuPolicy)(QWidget_ContextMenuPolicy(this.h))
}

func (this *QWidget) SetContextMenuPolicy(policy ContextMenuPolicy) {
	QWidget_SetContextMenuPolicy(this.h, (int)(policy))
}

func (this *QWidget) GrabMouse() {
	QWidget_GrabMouse(this.h)
}

func (this *QWidget) GrabMouseWithQCursor(param1 *QCursor) {
	QWidget_GrabMouseWithQCursor(this.h, param1.cPointer())
}

func (this *QWidget) ReleaseMouse() {
	QWidget_ReleaseMouse(this.h)
}

func (this *QWidget) GrabKeyboard() {
	QWidget_GrabKeyboard(this.h)
}

func (this *QWidget) ReleaseKeyboard() {
	QWidget_ReleaseKeyboard(this.h)
}

func (this *QWidget) GrabShortcut(key *QKeySequence) int {
	return (int)(QWidget_GrabShortcut(this.h, key.cPointer()))
}

func (this *QWidget) ReleaseShortcut(id int) {
	QWidget_ReleaseShortcut(this.h, (int)(id))
}

func (this *QWidget) SetShortcutEnabled(id int) {
	QWidget_SetShortcutEnabled(this.h, (int)(id))
}

func (this *QWidget) SetShortcutAutoRepeat(id int) {
	QWidget_SetShortcutAutoRepeat(this.h, (int)(id))
}

func QWidget_MouseGrabber() *QWidget {
	return newQWidget(QWidget_MouseGrabber())
}

func QWidget_KeyboardGrabber() *QWidget {
	return newQWidget(QWidget_KeyboardGrabber())
}

func (this *QWidget) UpdatesEnabled() bool {
	return (bool)(QWidget_UpdatesEnabled(this.h))
}

func (this *QWidget) SetUpdatesEnabled(enable bool) {
	QWidget_SetUpdatesEnabled(this.h, (bool)(enable))
}

func (this *QWidget) GraphicsProxyWidget() *QGraphicsProxyWidget {
	return newQGraphicsProxyWidget(QWidget_GraphicsProxyWidget(this.h))
}

func (this *QWidget) Update() {
	QWidget_Update(this.h)
}

func (this *QWidget) Repaint() {
	QWidget_Repaint(this.h)
}

func (this *QWidget) Update2(x int, y int, w int, h int) {
	QWidget_Update2(this.h, (int)(x), (int)(y), (int)(w), (int)(h))
}

func (this *QWidget) UpdateWithQRect(param1 *QRect) {
	QWidget_UpdateWithQRect(this.h, param1.cPointer())
}

func (this *QWidget) UpdateWithQRegion(param1 *QRegion) {
	QWidget_UpdateWithQRegion(this.h, param1.cPointer())
}

func (this *QWidget) Repaint2(x int, y int, w int, h int) {
	QWidget_Repaint2(this.h, (int)(x), (int)(y), (int)(w), (int)(h))
}

func (this *QWidget) RepaintWithQRect(param1 *QRect) {
	QWidget_RepaintWithQRect(this.h, param1.cPointer())
}

func (this *QWidget) RepaintWithQRegion(param1 *QRegion) {
	QWidget_RepaintWithQRegion(this.h, param1.cPointer())
}

func (this *QWidget) SetVisible(visible bool) {
	QWidget_SetVisible(this.h, (bool)(visible))
}

func (this *QWidget) SetHidden(hidden bool) {
	QWidget_SetHidden(this.h, (bool)(hidden))
}

func (this *QWidget) Show() {
	QWidget_Show(this.h)
}

func (this *QWidget) Hide() {
	QWidget_Hide(this.h)
}

func (this *QWidget) ShowMinimized() {
	QWidget_ShowMinimized(this.h)
}

func (this *QWidget) ShowMaximized() {
	QWidget_ShowMaximized(this.h)
}

func (this *QWidget) ShowFullScreen() {
	QWidget_ShowFullScreen(this.h)
}

func (this *QWidget) ShowNormal() {
	QWidget_ShowNormal(this.h)
}

func (this *QWidget) Close() bool {
	return (bool)(QWidget_Close(this.h))
}

func (this *QWidget) Raise() {
	QWidget_Raise(this.h)
}

func (this *QWidget) Lower() {
	QWidget_Lower(this.h)
}

func (this *QWidget) StackUnder(param1 *QWidget) {
	QWidget_StackUnder(this.h, param1.cPointer())
}

func (this *QWidget) Move(x int, y int) {
	QWidget_Move(this.h, (int)(x), (int)(y))
}

func (this *QWidget) MoveWithQPoint(param1 *QPoint) {
	QWidget_MoveWithQPoint(this.h, param1.cPointer())
}

func (this *QWidget) Resize(w int, h int) {
	QWidget_Resize(this.h, (int)(w), (int)(h))
}

func (this *QWidget) ResizeWithQSize(param1 *QSize) {
	QWidget_ResizeWithQSize(this.h, param1.cPointer())
}

func (this *QWidget) SetGeometry(x int, y int, w int, h int) {
	QWidget_SetGeometry(this.h, (int)(x), (int)(y), (int)(w), (int)(h))
}

func (this *QWidget) SetGeometryWithGeometry(geometry *QRect) {
	QWidget_SetGeometryWithGeometry(this.h, geometry.cPointer())
}

func (this *QWidget) SaveGeometry() []byte {
	var _bytearray struct_miqt_string = QWidget_SaveGeometry(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QWidget) RestoreGeometry(geometry []byte) bool {
	geometry_alias := struct_miqt_string{}
	geometry_alias.data = (char)(unsafe.Pointer(&geometry[0]))
	geometry_alias.len = size_t(len(geometry))
	return (bool)(QWidget_RestoreGeometry(this.h, geometry_alias))
}

func (this *QWidget) AdjustSize() {
	QWidget_AdjustSize(this.h)
}

func (this *QWidget) IsVisible() bool {
	return (bool)(QWidget_IsVisible(this.h))
}

func (this *QWidget) IsVisibleTo(param1 *QWidget) bool {
	return (bool)(QWidget_IsVisibleTo(this.h, param1.cPointer()))
}

func (this *QWidget) IsHidden() bool {
	return (bool)(QWidget_IsHidden(this.h))
}

func (this *QWidget) IsMinimized() bool {
	return (bool)(QWidget_IsMinimized(this.h))
}

func (this *QWidget) IsMaximized() bool {
	return (bool)(QWidget_IsMaximized(this.h))
}

func (this *QWidget) IsFullScreen() bool {
	return (bool)(QWidget_IsFullScreen(this.h))
}

func (this *QWidget) WindowState() WindowState {
	return (WindowState)(QWidget_WindowState(this.h))
}

func (this *QWidget) SetWindowState(state WindowState) {
	QWidget_SetWindowState(this.h, (int)(state))
}

func (this *QWidget) OverrideWindowState(state WindowState) {
	QWidget_OverrideWindowState(this.h, (int)(state))
}

func (this *QWidget) SizeHint() *QSize {
	_goptr := newQSize(QWidget_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) MinimumSizeHint() *QSize {
	_goptr := newQSize(QWidget_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) SizePolicy() *QSizePolicy {
	_goptr := newQSizePolicy(QWidget_SizePolicy(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) SetSizePolicy(sizePolicy QSizePolicy) {
	QWidget_SetSizePolicy(this.h, sizePolicy.cPointer())
}

func (this *QWidget) SetSizePolicy2(horizontal QSizePolicy__Policy, vertical QSizePolicy__Policy) {
	QWidget_SetSizePolicy2(this.h, (int)(horizontal), (int)(vertical))
}

func (this *QWidget) HeightForWidth(param1 int) int {
	return (int)(QWidget_HeightForWidth(this.h, (int)(param1)))
}

func (this *QWidget) HasHeightForWidth() bool {
	return (bool)(QWidget_HasHeightForWidth(this.h))
}

func (this *QWidget) VisibleRegion() *QRegion {
	_goptr := newQRegion(QWidget_VisibleRegion(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) SetContentsMargins(left int, top int, right int, bottom int) {
	QWidget_SetContentsMargins(this.h, (int)(left), (int)(top), (int)(right), (int)(bottom))
}

func (this *QWidget) SetContentsMarginsWithMargins(margins *QMargins) {
	QWidget_SetContentsMarginsWithMargins(this.h, margins.cPointer())
}

func (this *QWidget) ContentsMargins() *QMargins {
	_goptr := newQMargins(QWidget_ContentsMargins(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) ContentsRect() *QRect {
	_goptr := newQRect(QWidget_ContentsRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) Layout() *QLayout {
	return newQLayout(QWidget_Layout(this.h))
}

func (this *QWidget) SetLayout(layout *QLayout) {
	QWidget_SetLayout(this.h, layout.cPointer())
}

func (this *QWidget) UpdateGeometry() {
	QWidget_UpdateGeometry(this.h)
}

func (this *QWidget) SetParent(parent *QWidget) {
	QWidget_SetParent(this.h, parent.cPointer())
}

func (this *QWidget) SetParent2(parent *QWidget, f WindowType) {
	QWidget_SetParent2(this.h, parent.cPointer(), (int)(f))
}

func (this *QWidget) Scroll(dx int, dy int) {
	QWidget_Scroll(this.h, (int)(dx), (int)(dy))
}

func (this *QWidget) Scroll2(dx int, dy int, param3 *QRect) {
	QWidget_Scroll2(this.h, (int)(dx), (int)(dy), param3.cPointer())
}

func (this *QWidget) FocusWidget() *QWidget {
	return newQWidget(QWidget_FocusWidget(this.h))
}

func (this *QWidget) NextInFocusChain() *QWidget {
	return newQWidget(QWidget_NextInFocusChain(this.h))
}

func (this *QWidget) PreviousInFocusChain() *QWidget {
	return newQWidget(QWidget_PreviousInFocusChain(this.h))
}

func (this *QWidget) AcceptDrops() bool {
	return (bool)(QWidget_AcceptDrops(this.h))
}

func (this *QWidget) SetAcceptDrops(on bool) {
	QWidget_SetAcceptDrops(this.h, (bool)(on))
}

func (this *QWidget) AddAction(action *QAction) {
	QWidget_AddAction(this.h, action.cPointer())
}

func (this *QWidget) AddActions(actions []*QAction) {
	actions_CArray := (*[0xffff]*QAction)(malloc(size_t(8 * len(actions))))
	defer free(unsafe.Pointer(actions_CArray))
	for i := range actions {
		actions_CArray[i] = actions[i].cPointer()
	}
	actions_ma := struct_miqt_array{len: size_t(len(actions)), data: unsafe.Pointer(actions_CArray)}
	QWidget_AddActions(this.h, actions_ma)
}

func (this *QWidget) InsertActions(before *QAction, actions []*QAction) {
	actions_CArray := (*[0xffff]*QAction)(malloc(size_t(8 * len(actions))))
	defer free(unsafe.Pointer(actions_CArray))
	for i := range actions {
		actions_CArray[i] = actions[i].cPointer()
	}
	actions_ma := struct_miqt_array{len: size_t(len(actions)), data: unsafe.Pointer(actions_CArray)}
	QWidget_InsertActions(this.h, before.cPointer(), actions_ma)
}

func (this *QWidget) InsertAction(before *QAction, action *QAction) {
	QWidget_InsertAction(this.h, before.cPointer(), action.cPointer())
}

func (this *QWidget) RemoveAction(action *QAction) {
	QWidget_RemoveAction(this.h, action.cPointer())
}

func (this *QWidget) Actions() []*QAction {
	var _ma struct_miqt_array = QWidget_Actions(this.h)
	_ret := make([]*QAction, int(_ma.len))
	_outCast := (*[0xffff]*QAction)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQAction(_outCast[i])
	}
	return _ret
}

func (this *QWidget) AddActionWithText(text string) *QAction {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQAction(QWidget_AddActionWithText(this.h, text_ms))
}

func (this *QWidget) AddAction2(icon *QIcon, text string) *QAction {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQAction(QWidget_AddAction2(this.h, icon.cPointer(), text_ms))
}

func (this *QWidget) AddAction3(text string, shortcut *QKeySequence) *QAction {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQAction(QWidget_AddAction3(this.h, text_ms, shortcut.cPointer()))
}

func (this *QWidget) AddAction4(icon *QIcon, text string, shortcut *QKeySequence) *QAction {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQAction(QWidget_AddAction4(this.h, icon.cPointer(), text_ms, shortcut.cPointer()))
}

func (this *QWidget) ParentWidget() *QWidget {
	return newQWidget(QWidget_ParentWidget(this.h))
}

func (this *QWidget) SetWindowFlags(typeVal WindowType) {
	QWidget_SetWindowFlags(this.h, (int)(typeVal))
}

func (this *QWidget) WindowFlags() WindowType {
	return (WindowType)(QWidget_WindowFlags(this.h))
}

func (this *QWidget) SetWindowFlag(param1 WindowType) {
	QWidget_SetWindowFlag(this.h, (int)(param1))
}

func (this *QWidget) OverrideWindowFlags(typeVal WindowType) {
	QWidget_OverrideWindowFlags(this.h, (int)(typeVal))
}

func (this *QWidget) WindowType() WindowType {
	return (WindowType)(QWidget_WindowType(this.h))
}

func QWidget_Find(param1 uintptr) *QWidget {
	return newQWidget(QWidget_Find((uintptr_t)(param1)))
}

func (this *QWidget) ChildAt(x int, y int) *QWidget {
	return newQWidget(QWidget_ChildAt(this.h, (int)(x), (int)(y)))
}

func (this *QWidget) ChildAtWithQPoint(p *QPoint) *QWidget {
	return newQWidget(QWidget_ChildAtWithQPoint(this.h, p.cPointer()))
}

func (this *QWidget) ChildAtWithQPointF(p *QPointF) *QWidget {
	return newQWidget(QWidget_ChildAtWithQPointF(this.h, p.cPointer()))
}

func (this *QWidget) SetAttribute(param1 WidgetAttribute) {
	QWidget_SetAttribute(this.h, (int)(param1))
}

func (this *QWidget) TestAttribute(param1 WidgetAttribute) bool {
	return (bool)(QWidget_TestAttribute(this.h, (int)(param1)))
}

func (this *QWidget) PaintEngine() *QPaintEngine {
	return newQPaintEngine(QWidget_PaintEngine(this.h))
}

func (this *QWidget) EnsurePolished() {
	QWidget_EnsurePolished(this.h)
}

func (this *QWidget) IsAncestorOf(child *QWidget) bool {
	return (bool)(QWidget_IsAncestorOf(this.h, child.cPointer()))
}

func (this *QWidget) AutoFillBackground() bool {
	return (bool)(QWidget_AutoFillBackground(this.h))
}

func (this *QWidget) SetAutoFillBackground(enabled bool) {
	QWidget_SetAutoFillBackground(this.h, (bool)(enabled))
}

func (this *QWidget) BackingStore() *QBackingStore {
	return newQBackingStore(QWidget_BackingStore(this.h))
}

func (this *QWidget) WindowHandle() *QWindow {
	return newQWindow(QWidget_WindowHandle(this.h))
}

func (this *QWidget) Screen() *QScreen {
	return newQScreen(QWidget_Screen(this.h))
}

func (this *QWidget) SetScreen(screen *QScreen) {
	QWidget_SetScreen(this.h, screen.cPointer())
}

func QWidget_CreateWindowContainer(window *QWindow) *QWidget {
	return newQWidget(QWidget_CreateWindowContainer(window.cPointer()))
}

func (this *QWidget) WindowTitleChanged(title string) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	QWidget_WindowTitleChanged(this.h, title_ms)
}

func (this *QWidget) OnWindowTitleChanged(slot func(title string)) {
	QWidget_connect_WindowTitleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidget_WindowTitleChanged
func miqt_exec_callback_QWidget_WindowTitleChanged(cb intptr_t, title struct_miqt_string) {
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

func (this *QWidget) WindowIconChanged(icon *QIcon) {
	QWidget_WindowIconChanged(this.h, icon.cPointer())
}

func (this *QWidget) OnWindowIconChanged(slot func(icon *QIcon)) {
	QWidget_connect_WindowIconChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidget_WindowIconChanged
func miqt_exec_callback_QWidget_WindowIconChanged(cb intptr_t, icon *QIcon) {
	gofunc, ok := cgo.Handle(cb).Value().(func(icon *QIcon))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQIcon(icon)

	gofunc(slotval1)
}

func (this *QWidget) WindowIconTextChanged(iconText string) {
	iconText_ms := struct_miqt_string{}
	iconText_ms.data = CString(iconText)
	iconText_ms.len = size_t(len(iconText))
	defer free(unsafe.Pointer(iconText_ms.data))
	QWidget_WindowIconTextChanged(this.h, iconText_ms)
}

func (this *QWidget) OnWindowIconTextChanged(slot func(iconText string)) {
	QWidget_connect_WindowIconTextChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidget_WindowIconTextChanged
func miqt_exec_callback_QWidget_WindowIconTextChanged(cb intptr_t, iconText struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(iconText string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var iconText_ms struct_miqt_string = iconText
	iconText_ret := GoStringN(iconText_ms.data, int(int64(iconText_ms.len)))
	free(unsafe.Pointer(iconText_ms.data))
	slotval1 := iconText_ret

	gofunc(slotval1)
}

func (this *QWidget) CustomContextMenuRequested(pos *QPoint) {
	QWidget_CustomContextMenuRequested(this.h, pos.cPointer())
}

func (this *QWidget) OnCustomContextMenuRequested(slot func(pos *QPoint)) {
	QWidget_connect_CustomContextMenuRequested(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidget_CustomContextMenuRequested
func miqt_exec_callback_QWidget_CustomContextMenuRequested(cb intptr_t, pos *QPoint) {
	gofunc, ok := cgo.Handle(cb).Value().(func(pos *QPoint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(pos)

	gofunc(slotval1)
}

func (this *QWidget) InputMethodQuery(param1 InputMethodQuery) *QVariant {
	_goptr := newQVariant(QWidget_InputMethodQuery(this.h, (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) InputMethodHints() InputMethodHint {
	return (InputMethodHint)(QWidget_InputMethodHints(this.h))
}

func (this *QWidget) SetInputMethodHints(hints InputMethodHint) {
	QWidget_SetInputMethodHints(this.h, (int)(hints))
}

func QWidget_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWidget_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWidget_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWidget_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWidget) Render2(target *QPaintDevice, targetOffset *QPoint) {
	QWidget_Render2(this.h, target.cPointer(), targetOffset.cPointer())
}

func (this *QWidget) Render3(target *QPaintDevice, targetOffset *QPoint, sourceRegion *QRegion) {
	QWidget_Render3(this.h, target.cPointer(), targetOffset.cPointer(), sourceRegion.cPointer())
}

func (this *QWidget) Render4(target *QPaintDevice, targetOffset *QPoint, sourceRegion *QRegion, renderFlags RenderFlags) {
	QWidget_Render4(this.h, target.cPointer(), targetOffset.cPointer(), sourceRegion.cPointer(), renderFlags)
}

func (this *QWidget) Render22(painter *QPainter, targetOffset *QPoint) {
	QWidget_Render22(this.h, painter.cPointer(), targetOffset.cPointer())
}

func (this *QWidget) Render32(painter *QPainter, targetOffset *QPoint, sourceRegion *QRegion) {
	QWidget_Render32(this.h, painter.cPointer(), targetOffset.cPointer(), sourceRegion.cPointer())
}

func (this *QWidget) Render42(painter *QPainter, targetOffset *QPoint, sourceRegion *QRegion, renderFlags RenderFlags) {
	QWidget_Render42(this.h, painter.cPointer(), targetOffset.cPointer(), sourceRegion.cPointer(), renderFlags)
}

func (this *QWidget) Grab1(rectangle *QRect) *QPixmap {
	_goptr := newQPixmap(QWidget_Grab1(this.h, rectangle.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidget) GrabGesture2(typeVal GestureType, flags GestureFlag) {
	QWidget_GrabGesture2(this.h, (int)(typeVal), (int)(flags))
}

func (this *QWidget) GrabShortcut2(key *QKeySequence, context ShortcutContext) int {
	return (int)(QWidget_GrabShortcut2(this.h, key.cPointer(), (int)(context)))
}

func (this *QWidget) SetShortcutEnabled2(id int, enable bool) {
	QWidget_SetShortcutEnabled2(this.h, (int)(id), (bool)(enable))
}

func (this *QWidget) SetShortcutAutoRepeat2(id int, enable bool) {
	QWidget_SetShortcutAutoRepeat2(this.h, (int)(id), (bool)(enable))
}

func (this *QWidget) SetWindowFlag2(param1 WindowType, on bool) {
	QWidget_SetWindowFlag2(this.h, (int)(param1), (bool)(on))
}

func (this *QWidget) SetAttribute2(param1 WidgetAttribute, on bool) {
	QWidget_SetAttribute2(this.h, (int)(param1), (bool)(on))
}

func QWidget_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	return newQWidget(QWidget_CreateWindowContainer2(window.cPointer(), parent.cPointer()))
}

func QWidget_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	return newQWidget(QWidget_CreateWindowContainer3(window.cPointer(), parent.cPointer(), (int)(flags)))
}

func (this *QWidget) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QWidget_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QWidget) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidget_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidget_MetaObject
func miqt_exec_callback_QWidget_MetaObject(self QWidget, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidget{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QWidget) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QWidget_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QWidget) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidget_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidget_Metacast
func miqt_exec_callback_QWidget_Metacast(self QWidget, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QWidget{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
