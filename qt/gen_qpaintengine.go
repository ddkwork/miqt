package qt

import (
	"unsafe"
)

type QTextItem__RenderFlag int

const (
	QTextItem__RightToLeft QTextItem__RenderFlag = 1
	QTextItem__Overline    QTextItem__RenderFlag = 16
	QTextItem__Underline   QTextItem__RenderFlag = 32
	QTextItem__StrikeOut   QTextItem__RenderFlag = 64
	QTextItem__Dummy       QTextItem__RenderFlag = 4294967295
)

type QPaintEngine__PaintEngineFeature int

const (
	QPaintEngine__PrimitiveTransform          QPaintEngine__PaintEngineFeature = 1
	QPaintEngine__PatternTransform            QPaintEngine__PaintEngineFeature = 2
	QPaintEngine__PixmapTransform             QPaintEngine__PaintEngineFeature = 4
	QPaintEngine__PatternBrush                QPaintEngine__PaintEngineFeature = 8
	QPaintEngine__LinearGradientFill          QPaintEngine__PaintEngineFeature = 16
	QPaintEngine__RadialGradientFill          QPaintEngine__PaintEngineFeature = 32
	QPaintEngine__ConicalGradientFill         QPaintEngine__PaintEngineFeature = 64
	QPaintEngine__AlphaBlend                  QPaintEngine__PaintEngineFeature = 128
	QPaintEngine__PorterDuff                  QPaintEngine__PaintEngineFeature = 256
	QPaintEngine__PainterPaths                QPaintEngine__PaintEngineFeature = 512
	QPaintEngine__Antialiasing                QPaintEngine__PaintEngineFeature = 1024
	QPaintEngine__BrushStroke                 QPaintEngine__PaintEngineFeature = 2048
	QPaintEngine__ConstantOpacity             QPaintEngine__PaintEngineFeature = 4096
	QPaintEngine__MaskedBrush                 QPaintEngine__PaintEngineFeature = 8192
	QPaintEngine__PerspectiveTransform        QPaintEngine__PaintEngineFeature = 16384
	QPaintEngine__BlendModes                  QPaintEngine__PaintEngineFeature = 32768
	QPaintEngine__ObjectBoundingModeGradients QPaintEngine__PaintEngineFeature = 65536
	QPaintEngine__RasterOpModes               QPaintEngine__PaintEngineFeature = 131072
	QPaintEngine__PaintOutsidePaintEvent      QPaintEngine__PaintEngineFeature = 536870912
	QPaintEngine__AllFeatures                 QPaintEngine__PaintEngineFeature = 4294967295
)

type QPaintEngine__DirtyFlag int

const (
	QPaintEngine__DirtyPen             QPaintEngine__DirtyFlag = 1
	QPaintEngine__DirtyBrush           QPaintEngine__DirtyFlag = 2
	QPaintEngine__DirtyBrushOrigin     QPaintEngine__DirtyFlag = 4
	QPaintEngine__DirtyFont            QPaintEngine__DirtyFlag = 8
	QPaintEngine__DirtyBackground      QPaintEngine__DirtyFlag = 16
	QPaintEngine__DirtyBackgroundMode  QPaintEngine__DirtyFlag = 32
	QPaintEngine__DirtyTransform       QPaintEngine__DirtyFlag = 64
	QPaintEngine__DirtyClipRegion      QPaintEngine__DirtyFlag = 128
	QPaintEngine__DirtyClipPath        QPaintEngine__DirtyFlag = 256
	QPaintEngine__DirtyHints           QPaintEngine__DirtyFlag = 512
	QPaintEngine__DirtyCompositionMode QPaintEngine__DirtyFlag = 1024
	QPaintEngine__DirtyClipEnabled     QPaintEngine__DirtyFlag = 2048
	QPaintEngine__DirtyOpacity         QPaintEngine__DirtyFlag = 4096
	QPaintEngine__AllDirty             QPaintEngine__DirtyFlag = 65535
)

type QPaintEngine__PolygonDrawMode int

const (
	QPaintEngine__OddEvenMode  QPaintEngine__PolygonDrawMode = 0
	QPaintEngine__WindingMode  QPaintEngine__PolygonDrawMode = 1
	QPaintEngine__ConvexMode   QPaintEngine__PolygonDrawMode = 2
	QPaintEngine__PolylineMode QPaintEngine__PolygonDrawMode = 3
)

type QPaintEngine__Type int

const (
	QPaintEngine__X11           QPaintEngine__Type = 0
	QPaintEngine__Windows       QPaintEngine__Type = 1
	QPaintEngine__QuickDraw     QPaintEngine__Type = 2
	QPaintEngine__CoreGraphics  QPaintEngine__Type = 3
	QPaintEngine__MacPrinter    QPaintEngine__Type = 4
	QPaintEngine__QWindowSystem QPaintEngine__Type = 5
	QPaintEngine__OpenGL        QPaintEngine__Type = 6
	QPaintEngine__Picture       QPaintEngine__Type = 7
	QPaintEngine__SVG           QPaintEngine__Type = 8
	QPaintEngine__Raster        QPaintEngine__Type = 9
	QPaintEngine__Direct3D      QPaintEngine__Type = 10
	QPaintEngine__Pdf           QPaintEngine__Type = 11
	QPaintEngine__OpenVG        QPaintEngine__Type = 12
	QPaintEngine__OpenGL2       QPaintEngine__Type = 13
	QPaintEngine__PaintBuffer   QPaintEngine__Type = 14
	QPaintEngine__Blitter       QPaintEngine__Type = 15
	QPaintEngine__Direct2D      QPaintEngine__Type = 16
	QPaintEngine__User          QPaintEngine__Type = 50
	QPaintEngine__MaxUser       QPaintEngine__Type = 100
)

type QTextItem struct {
	h          uintptr
	isSubclass bool
}

// NewQTextItem constructs a new QTextItem object.
func NewQTextItem() *QTextItem {
	ret := newQTextItem(QTextItem_new())
	ret.isSubclass = true
	return ret
}

// NewQTextItem2 constructs a new QTextItem object.
func NewQTextItem2(param1 *QTextItem) *QTextItem {
	ret := newQTextItem(QTextItem_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTextItem) Descent() float64 {
	return (float64)(QTextItem_Descent(this.h))
}

func (this *QTextItem) Ascent() float64 {
	return (float64)(QTextItem_Ascent(this.h))
}

func (this *QTextItem) Width() float64 {
	return (float64)(QTextItem_Width(this.h))
}

func (this *QTextItem) RenderFlags() RenderFlags {
	xxxxxxxxx
}

func (this *QTextItem) Text() string {
	var _ms struct_miqt_string = QTextItem_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextItem) Font() *QFont {
	_goptr := newQFont(QTextItem_Font(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QPaintEngine struct {
	h          uintptr
	isSubclass bool
}

// NewQPaintEngine constructs a new QPaintEngine object.
func NewQPaintEngine() *QPaintEngine {
	ret := newQPaintEngine(QPaintEngine_new())
	ret.isSubclass = true
	return ret
}

// NewQPaintEngine2 constructs a new QPaintEngine object.
func NewQPaintEngine2(features PaintEngineFeatures) *QPaintEngine {
	ret := newQPaintEngine(QPaintEngine_new2(features))
	ret.isSubclass = true
	return ret
}

func (this *QPaintEngine) IsActive() bool {
	return (bool)(QPaintEngine_IsActive(this.h))
}

func (this *QPaintEngine) SetActive(newState bool) {
	QPaintEngine_SetActive(this.h, (bool)(newState))
}

func (this *QPaintEngine) Begin(pdev *QPaintDevice) bool {
	return (bool)(QPaintEngine_Begin(this.h, pdev.cPointer()))
}

func (this *QPaintEngine) End() bool {
	return (bool)(QPaintEngine_End(this.h))
}

func (this *QPaintEngine) UpdateState(state *QPaintEngineState) {
	QPaintEngine_UpdateState(this.h, state.cPointer())
}

func (this *QPaintEngine) DrawRects(rects *QRect, rectCount int) {
	QPaintEngine_DrawRects(this.h, rects.cPointer(), (int)(rectCount))
}

func (this *QPaintEngine) DrawRects2(rects *QRectF, rectCount int) {
	QPaintEngine_DrawRects2(this.h, rects.cPointer(), (int)(rectCount))
}

func (this *QPaintEngine) DrawLines(lines *QLine, lineCount int) {
	QPaintEngine_DrawLines(this.h, lines.cPointer(), (int)(lineCount))
}

func (this *QPaintEngine) DrawLines2(lines *QLineF, lineCount int) {
	QPaintEngine_DrawLines2(this.h, lines.cPointer(), (int)(lineCount))
}

func (this *QPaintEngine) DrawEllipse(r *QRectF) {
	QPaintEngine_DrawEllipse(this.h, r.cPointer())
}

func (this *QPaintEngine) DrawEllipseWithQRect(r *QRect) {
	QPaintEngine_DrawEllipseWithQRect(this.h, r.cPointer())
}

func (this *QPaintEngine) DrawPath(path *QPainterPath) {
	QPaintEngine_DrawPath(this.h, path.cPointer())
}

func (this *QPaintEngine) DrawPoints(points *QPointF, pointCount int) {
	QPaintEngine_DrawPoints(this.h, points.cPointer(), (int)(pointCount))
}

func (this *QPaintEngine) DrawPoints2(points *QPoint, pointCount int) {
	QPaintEngine_DrawPoints2(this.h, points.cPointer(), (int)(pointCount))
}

func (this *QPaintEngine) DrawPolygon(points *QPointF, pointCount int, mode PolygonDrawMode) {
	QPaintEngine_DrawPolygon(this.h, points.cPointer(), (int)(pointCount), mode)
}

func (this *QPaintEngine) DrawPolygon2(points *QPoint, pointCount int, mode PolygonDrawMode) {
	QPaintEngine_DrawPolygon2(this.h, points.cPointer(), (int)(pointCount), mode)
}

func (this *QPaintEngine) DrawPixmap(r *QRectF, pm *QPixmap, sr *QRectF) {
	QPaintEngine_DrawPixmap(this.h, r.cPointer(), pm.cPointer(), sr.cPointer())
}

func (this *QPaintEngine) DrawTextItem(p *QPointF, textItem *QTextItem) {
	QPaintEngine_DrawTextItem(this.h, p.cPointer(), textItem.cPointer())
}

func (this *QPaintEngine) DrawTiledPixmap(r *QRectF, pixmap *QPixmap, s *QPointF) {
	QPaintEngine_DrawTiledPixmap(this.h, r.cPointer(), pixmap.cPointer(), s.cPointer())
}

func (this *QPaintEngine) DrawImage(r *QRectF, pm *QImage, sr *QRectF, flags ImageConversionFlag) {
	QPaintEngine_DrawImage(this.h, r.cPointer(), pm.cPointer(), sr.cPointer(), (int)(flags))
}

func (this *QPaintEngine) SetPaintDevice(device *QPaintDevice) {
	QPaintEngine_SetPaintDevice(this.h, device.cPointer())
}

func (this *QPaintEngine) PaintDevice() *QPaintDevice {
	return newQPaintDevice(QPaintEngine_PaintDevice(this.h))
}

func (this *QPaintEngine) SetSystemClip(baseClip *QRegion) {
	QPaintEngine_SetSystemClip(this.h, baseClip.cPointer())
}

func (this *QPaintEngine) SystemClip() *QRegion {
	_goptr := newQRegion(QPaintEngine_SystemClip(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPaintEngine) SetSystemRect(rect *QRect) {
	QPaintEngine_SetSystemRect(this.h, rect.cPointer())
}

func (this *QPaintEngine) SystemRect() *QRect {
	_goptr := newQRect(QPaintEngine_SystemRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPaintEngine) CoordinateOffset() *QPoint {
	_goptr := newQPoint(QPaintEngine_CoordinateOffset(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPaintEngine) Type() Type {
	xxxxxxxxx
}

func (this *QPaintEngine) FixNegRect(x *int, y *int, w *int, h *int) {
	QPaintEngine_FixNegRect(this.h, (*int)(unsafe.Pointer(x)), (*int)(unsafe.Pointer(y)), (*int)(unsafe.Pointer(w)), (*int)(unsafe.Pointer(h)))
}

func (this *QPaintEngine) TestDirty(df DirtyFlags) bool {
	return (bool)(QPaintEngine_TestDirty(this.h, df))
}

func (this *QPaintEngine) SetDirty(df DirtyFlags) {
	QPaintEngine_SetDirty(this.h, df)
}

func (this *QPaintEngine) ClearDirty(df DirtyFlags) {
	QPaintEngine_ClearDirty(this.h, df)
}

func (this *QPaintEngine) HasFeature(feature PaintEngineFeatures) bool {
	return (bool)(QPaintEngine_HasFeature(this.h, feature))
}

func (this *QPaintEngine) Painter() *QPainter {
	return newQPainter(QPaintEngine_Painter(this.h))
}

func (this *QPaintEngine) SyncState() {
	QPaintEngine_SyncState(this.h)
}

func (this *QPaintEngine) IsExtended() bool {
	return (bool)(QPaintEngine_IsExtended(this.h))
}

func (this *QPaintEngine) CreatePixmap(size QSize) *QPixmap {
	_goptr := newQPixmap(QPaintEngine_CreatePixmap(this.h, size.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPaintEngine) CreatePixmapFromImage(image QImage, flags ImageConversionFlag) *QPixmap {
	_goptr := newQPixmap(QPaintEngine_CreatePixmapFromImage(this.h, image.cPointer(), (int)(flags)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QPaintEngineState struct {
	h          uintptr
	isSubclass bool
}

func (this *QPaintEngineState) State() DirtyFlag {
	return (DirtyFlag)(QPaintEngineState_State(this.h))
}

func (this *QPaintEngineState) Pen() *QPen {
	_goptr := newQPen(QPaintEngineState_Pen(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPaintEngineState) Brush() *QBrush {
	_goptr := newQBrush(QPaintEngineState_Brush(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPaintEngineState) BrushOrigin() *QPointF {
	_goptr := newQPointF(QPaintEngineState_BrushOrigin(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPaintEngineState) BackgroundBrush() *QBrush {
	_goptr := newQBrush(QPaintEngineState_BackgroundBrush(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPaintEngineState) BackgroundMode() BGMode {
	return (BGMode)(QPaintEngineState_BackgroundMode(this.h))
}

func (this *QPaintEngineState) Font() *QFont {
	_goptr := newQFont(QPaintEngineState_Font(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPaintEngineState) Transform() *QTransform {
	_goptr := newQTransform(QPaintEngineState_Transform(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPaintEngineState) ClipOperation() ClipOperation {
	return (ClipOperation)(QPaintEngineState_ClipOperation(this.h))
}

func (this *QPaintEngineState) ClipRegion() *QRegion {
	_goptr := newQRegion(QPaintEngineState_ClipRegion(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPaintEngineState) ClipPath() *QPainterPath {
	_goptr := newQPainterPath(QPaintEngineState_ClipPath(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPaintEngineState) IsClipEnabled() bool {
	return (bool)(QPaintEngineState_IsClipEnabled(this.h))
}

func (this *QPaintEngineState) RenderHints() RenderHint {
	return (RenderHint)(QPaintEngineState_RenderHints(this.h))
}

func (this *QPaintEngineState) CompositionMode() QPainter__CompositionMode {
	return (QPainter__CompositionMode)(QPaintEngineState_CompositionMode(this.h))
}

func (this *QPaintEngineState) Opacity() float64 {
	return (float64)(QPaintEngineState_Opacity(this.h))
}

func (this *QPaintEngineState) Painter() *QPainter {
	return newQPainter(QPaintEngineState_Painter(this.h))
}

func (this *QPaintEngineState) BrushNeedsResolving() bool {
	return (bool)(QPaintEngineState_BrushNeedsResolving(this.h))
}

func (this *QPaintEngineState) PenNeedsResolving() bool {
	return (bool)(QPaintEngineState_PenNeedsResolving(this.h))
}
