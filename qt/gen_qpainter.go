package qt

import (
	"unsafe"
)

type QPainter__RenderHint int

const (
	QPainter__Antialiasing                QPainter__RenderHint = 1
	QPainter__TextAntialiasing            QPainter__RenderHint = 2
	QPainter__SmoothPixmapTransform       QPainter__RenderHint = 4
	QPainter__VerticalSubpixelPositioning QPainter__RenderHint = 8
	QPainter__LosslessImageRendering      QPainter__RenderHint = 64
	QPainter__NonCosmeticBrushPatterns    QPainter__RenderHint = 128
)

type QPainter__PixmapFragmentHint int

const (
	QPainter__OpaqueHint QPainter__PixmapFragmentHint = 1
)

type QPainter__CompositionMode int

const (
	QPainter__CompositionMode_SourceOver          QPainter__CompositionMode = 0
	QPainter__CompositionMode_DestinationOver     QPainter__CompositionMode = 1
	QPainter__CompositionMode_Clear               QPainter__CompositionMode = 2
	QPainter__CompositionMode_Source              QPainter__CompositionMode = 3
	QPainter__CompositionMode_Destination         QPainter__CompositionMode = 4
	QPainter__CompositionMode_SourceIn            QPainter__CompositionMode = 5
	QPainter__CompositionMode_DestinationIn       QPainter__CompositionMode = 6
	QPainter__CompositionMode_SourceOut           QPainter__CompositionMode = 7
	QPainter__CompositionMode_DestinationOut      QPainter__CompositionMode = 8
	QPainter__CompositionMode_SourceAtop          QPainter__CompositionMode = 9
	QPainter__CompositionMode_DestinationAtop     QPainter__CompositionMode = 10
	QPainter__CompositionMode_Xor                 QPainter__CompositionMode = 11
	QPainter__CompositionMode_Plus                QPainter__CompositionMode = 12
	QPainter__CompositionMode_Multiply            QPainter__CompositionMode = 13
	QPainter__CompositionMode_Screen              QPainter__CompositionMode = 14
	QPainter__CompositionMode_Overlay             QPainter__CompositionMode = 15
	QPainter__CompositionMode_Darken              QPainter__CompositionMode = 16
	QPainter__CompositionMode_Lighten             QPainter__CompositionMode = 17
	QPainter__CompositionMode_ColorDodge          QPainter__CompositionMode = 18
	QPainter__CompositionMode_ColorBurn           QPainter__CompositionMode = 19
	QPainter__CompositionMode_HardLight           QPainter__CompositionMode = 20
	QPainter__CompositionMode_SoftLight           QPainter__CompositionMode = 21
	QPainter__CompositionMode_Difference          QPainter__CompositionMode = 22
	QPainter__CompositionMode_Exclusion           QPainter__CompositionMode = 23
	QPainter__RasterOp_SourceOrDestination        QPainter__CompositionMode = 24
	QPainter__RasterOp_SourceAndDestination       QPainter__CompositionMode = 25
	QPainter__RasterOp_SourceXorDestination       QPainter__CompositionMode = 26
	QPainter__RasterOp_NotSourceAndNotDestination QPainter__CompositionMode = 27
	QPainter__RasterOp_NotSourceOrNotDestination  QPainter__CompositionMode = 28
	QPainter__RasterOp_NotSourceXorDestination    QPainter__CompositionMode = 29
	QPainter__RasterOp_NotSource                  QPainter__CompositionMode = 30
	QPainter__RasterOp_NotSourceAndDestination    QPainter__CompositionMode = 31
	QPainter__RasterOp_SourceAndNotDestination    QPainter__CompositionMode = 32
	QPainter__RasterOp_NotSourceOrDestination     QPainter__CompositionMode = 33
	QPainter__RasterOp_SourceOrNotDestination     QPainter__CompositionMode = 34
	QPainter__RasterOp_ClearDestination           QPainter__CompositionMode = 35
	QPainter__RasterOp_SetDestination             QPainter__CompositionMode = 36
	QPainter__RasterOp_NotDestination             QPainter__CompositionMode = 37
	QPainter__NCompositionModes                   QPainter__CompositionMode = 38
)

type QPainter struct {
	h          uintptr
	isSubclass bool
}

// NewQPainter constructs a new QPainter object.
func NewQPainter() *QPainter {
	g := newQPainter(QPainter_new())
	g.isSubclass = true
	return g
}

// NewQPainter2 constructs a new QPainter object.
func NewQPainter2(param1 *QPaintDevice) *QPainter {
	g := newQPainter(QPainter_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QPainter) Device() *QPaintDevice {
	return newQPaintDevice(QPainter_Device(this.h))
}

func (this *QPainter) Begin(param1 *QPaintDevice) bool {
	return (bool)(QPainter_Begin(this.h, param1.cPointer()))
}

func (this *QPainter) End() bool {
	return (bool)(QPainter_End(this.h))
}

func (this *QPainter) IsActive() bool {
	return (bool)(QPainter_IsActive(this.h))
}

func (this *QPainter) SetCompositionMode(mode CompositionMode) {
	QPainter_SetCompositionMode(this.h, mode)
}

func (this *QPainter) CompositionMode() CompositionMode {
	xxxxxxxxx
}

func (this *QPainter) Font() *QFont {
	return newQFont(QPainter_Font(this.h))
}

func (this *QPainter) SetFont(f *QFont) {
	QPainter_SetFont(this.h, f.cPointer())
}

func (this *QPainter) FontMetrics() *QFontMetrics {
	_goptr := newQFontMetrics(QPainter_FontMetrics(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainter) FontInfo() *QFontInfo {
	_goptr := newQFontInfo(QPainter_FontInfo(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainter) SetPen(color *QColor) {
	QPainter_SetPen(this.h, color.cPointer())
}

func (this *QPainter) SetPenWithPen(pen *QPen) {
	QPainter_SetPenWithPen(this.h, pen.cPointer())
}

func (this *QPainter) SetPenWithStyle(style PenStyle) {
	QPainter_SetPenWithStyle(this.h, (int)(style))
}

func (this *QPainter) Pen() *QPen {
	return newQPen(QPainter_Pen(this.h))
}

func (this *QPainter) SetBrush(brush *QBrush) {
	QPainter_SetBrush(this.h, brush.cPointer())
}

func (this *QPainter) SetBrushWithStyle(style BrushStyle) {
	QPainter_SetBrushWithStyle(this.h, (int)(style))
}

func (this *QPainter) SetBrushWithColor(color QColor) {
	QPainter_SetBrushWithColor(this.h, color.cPointer())
}

func (this *QPainter) SetBrush2(color GlobalColor) {
	QPainter_SetBrush2(this.h, (int)(color))
}

func (this *QPainter) Brush() *QBrush {
	return newQBrush(QPainter_Brush(this.h))
}

func (this *QPainter) SetBackgroundMode(mode BGMode) {
	QPainter_SetBackgroundMode(this.h, (int)(mode))
}

func (this *QPainter) BackgroundMode() BGMode {
	return (BGMode)(QPainter_BackgroundMode(this.h))
}

func (this *QPainter) BrushOrigin() *QPoint {
	_goptr := newQPoint(QPainter_BrushOrigin(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainter) SetBrushOrigin(x int, y int) {
	QPainter_SetBrushOrigin(this.h, (int)(x), (int)(y))
}

func (this *QPainter) SetBrushOriginWithBrushOrigin(brushOrigin *QPoint) {
	QPainter_SetBrushOriginWithBrushOrigin(this.h, brushOrigin.cPointer())
}

func (this *QPainter) SetBrushOrigin2(brushOrigin *QPointF) {
	QPainter_SetBrushOrigin2(this.h, brushOrigin.cPointer())
}

func (this *QPainter) SetBackground(bg *QBrush) {
	QPainter_SetBackground(this.h, bg.cPointer())
}

func (this *QPainter) Background() *QBrush {
	return newQBrush(QPainter_Background(this.h))
}

func (this *QPainter) Opacity() float64 {
	return (float64)(QPainter_Opacity(this.h))
}

func (this *QPainter) SetOpacity(opacity float64) {
	QPainter_SetOpacity(this.h, (double)(opacity))
}

func (this *QPainter) ClipRegion() *QRegion {
	_goptr := newQRegion(QPainter_ClipRegion(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainter) ClipPath() *QPainterPath {
	_goptr := newQPainterPath(QPainter_ClipPath(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainter) SetClipRect(param1 *QRectF) {
	QPainter_SetClipRect(this.h, param1.cPointer())
}

func (this *QPainter) SetClipRectWithQRect(param1 *QRect) {
	QPainter_SetClipRectWithQRect(this.h, param1.cPointer())
}

func (this *QPainter) SetClipRect2(x int, y int, w int, h int) {
	QPainter_SetClipRect2(this.h, (int)(x), (int)(y), (int)(w), (int)(h))
}

func (this *QPainter) SetClipRegion(param1 *QRegion) {
	QPainter_SetClipRegion(this.h, param1.cPointer())
}

func (this *QPainter) SetClipPath(path *QPainterPath) {
	QPainter_SetClipPath(this.h, path.cPointer())
}

func (this *QPainter) SetClipping(enable bool) {
	QPainter_SetClipping(this.h, (bool)(enable))
}

func (this *QPainter) HasClipping() bool {
	return (bool)(QPainter_HasClipping(this.h))
}

func (this *QPainter) ClipBoundingRect() *QRectF {
	_goptr := newQRectF(QPainter_ClipBoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainter) Save() {
	QPainter_Save(this.h)
}

func (this *QPainter) Restore() {
	QPainter_Restore(this.h)
}

func (this *QPainter) SetTransform(transform *QTransform) {
	QPainter_SetTransform(this.h, transform.cPointer())
}

func (this *QPainter) Transform() *QTransform {
	return newQTransform(QPainter_Transform(this.h))
}

func (this *QPainter) DeviceTransform() *QTransform {
	return newQTransform(QPainter_DeviceTransform(this.h))
}

func (this *QPainter) ResetTransform() {
	QPainter_ResetTransform(this.h)
}

func (this *QPainter) SetWorldTransform(matrix *QTransform) {
	QPainter_SetWorldTransform(this.h, matrix.cPointer())
}

func (this *QPainter) WorldTransform() *QTransform {
	return newQTransform(QPainter_WorldTransform(this.h))
}

func (this *QPainter) CombinedTransform() *QTransform {
	_goptr := newQTransform(QPainter_CombinedTransform(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainter) SetWorldMatrixEnabled(enabled bool) {
	QPainter_SetWorldMatrixEnabled(this.h, (bool)(enabled))
}

func (this *QPainter) WorldMatrixEnabled() bool {
	return (bool)(QPainter_WorldMatrixEnabled(this.h))
}

func (this *QPainter) Scale(sx float64, sy float64) {
	QPainter_Scale(this.h, (double)(sx), (double)(sy))
}

func (this *QPainter) Shear(sh float64, sv float64) {
	QPainter_Shear(this.h, (double)(sh), (double)(sv))
}

func (this *QPainter) Rotate(a float64) {
	QPainter_Rotate(this.h, (double)(a))
}

func (this *QPainter) Translate(offset *QPointF) {
	QPainter_Translate(this.h, offset.cPointer())
}

func (this *QPainter) TranslateWithOffset(offset *QPoint) {
	QPainter_TranslateWithOffset(this.h, offset.cPointer())
}

func (this *QPainter) Translate2(dx float64, dy float64) {
	QPainter_Translate2(this.h, (double)(dx), (double)(dy))
}

func (this *QPainter) Window() *QRect {
	_goptr := newQRect(QPainter_Window(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainter) SetWindow(window *QRect) {
	QPainter_SetWindow(this.h, window.cPointer())
}

func (this *QPainter) SetWindow2(x int, y int, w int, h int) {
	QPainter_SetWindow2(this.h, (int)(x), (int)(y), (int)(w), (int)(h))
}

func (this *QPainter) Viewport() *QRect {
	_goptr := newQRect(QPainter_Viewport(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainter) SetViewport(viewport *QRect) {
	QPainter_SetViewport(this.h, viewport.cPointer())
}

func (this *QPainter) SetViewport2(x int, y int, w int, h int) {
	QPainter_SetViewport2(this.h, (int)(x), (int)(y), (int)(w), (int)(h))
}

func (this *QPainter) SetViewTransformEnabled(enable bool) {
	QPainter_SetViewTransformEnabled(this.h, (bool)(enable))
}

func (this *QPainter) ViewTransformEnabled() bool {
	return (bool)(QPainter_ViewTransformEnabled(this.h))
}

func (this *QPainter) StrokePath(path *QPainterPath, pen *QPen) {
	QPainter_StrokePath(this.h, path.cPointer(), pen.cPointer())
}

func (this *QPainter) FillPath(path *QPainterPath, brush *QBrush) {
	QPainter_FillPath(this.h, path.cPointer(), brush.cPointer())
}

func (this *QPainter) DrawPath(path *QPainterPath) {
	QPainter_DrawPath(this.h, path.cPointer())
}

func (this *QPainter) DrawPoint(pt *QPointF) {
	QPainter_DrawPoint(this.h, pt.cPointer())
}

func (this *QPainter) DrawPointWithQPoint(p *QPoint) {
	QPainter_DrawPointWithQPoint(this.h, p.cPointer())
}

func (this *QPainter) DrawPoint2(x int, y int) {
	QPainter_DrawPoint2(this.h, (int)(x), (int)(y))
}

func (this *QPainter) DrawPoints(points *QPointF, pointCount int) {
	QPainter_DrawPoints(this.h, points.cPointer(), (int)(pointCount))
}

func (this *QPainter) DrawPoints2(points *QPoint, pointCount int) {
	QPainter_DrawPoints2(this.h, points.cPointer(), (int)(pointCount))
}

func (this *QPainter) DrawLine(line *QLineF) {
	QPainter_DrawLine(this.h, line.cPointer())
}

func (this *QPainter) DrawLineWithLine(line *QLine) {
	QPainter_DrawLineWithLine(this.h, line.cPointer())
}

func (this *QPainter) DrawLine2(x1 int, y1 int, x2 int, y2 int) {
	QPainter_DrawLine2(this.h, (int)(x1), (int)(y1), (int)(x2), (int)(y2))
}

func (this *QPainter) DrawLine3(p1 *QPoint, p2 *QPoint) {
	QPainter_DrawLine3(this.h, p1.cPointer(), p2.cPointer())
}

func (this *QPainter) DrawLine4(p1 *QPointF, p2 *QPointF) {
	QPainter_DrawLine4(this.h, p1.cPointer(), p2.cPointer())
}

func (this *QPainter) DrawLines(lines *QLineF, lineCount int) {
	QPainter_DrawLines(this.h, lines.cPointer(), (int)(lineCount))
}

func (this *QPainter) DrawLinesWithLines(lines []QLineF) {
	lines_CArray := (*[0xffff]*QLineF)(malloc(size_t(8 * len(lines))))
	defer free(unsafe.Pointer(lines_CArray))
	for i := range lines {
		lines_CArray[i] = lines[i].cPointer()
	}
	lines_ma := struct_miqt_array{len: size_t(len(lines)), data: unsafe.Pointer(lines_CArray)}
	QPainter_DrawLinesWithLines(this.h, lines_ma)
}

func (this *QPainter) DrawLines2(pointPairs *QPointF, lineCount int) {
	QPainter_DrawLines2(this.h, pointPairs.cPointer(), (int)(lineCount))
}

func (this *QPainter) DrawLinesWithPointPairs(pointPairs []QPointF) {
	pointPairs_CArray := (*[0xffff]*QPointF)(malloc(size_t(8 * len(pointPairs))))
	defer free(unsafe.Pointer(pointPairs_CArray))
	for i := range pointPairs {
		pointPairs_CArray[i] = pointPairs[i].cPointer()
	}
	pointPairs_ma := struct_miqt_array{len: size_t(len(pointPairs)), data: unsafe.Pointer(pointPairs_CArray)}
	QPainter_DrawLinesWithPointPairs(this.h, pointPairs_ma)
}

func (this *QPainter) DrawLines3(lines *QLine, lineCount int) {
	QPainter_DrawLines3(this.h, lines.cPointer(), (int)(lineCount))
}

func (this *QPainter) DrawLines4(lines []QLine) {
	lines_CArray := (*[0xffff]*QLine)(malloc(size_t(8 * len(lines))))
	defer free(unsafe.Pointer(lines_CArray))
	for i := range lines {
		lines_CArray[i] = lines[i].cPointer()
	}
	lines_ma := struct_miqt_array{len: size_t(len(lines)), data: unsafe.Pointer(lines_CArray)}
	QPainter_DrawLines4(this.h, lines_ma)
}

func (this *QPainter) DrawLines5(pointPairs *QPoint, lineCount int) {
	QPainter_DrawLines5(this.h, pointPairs.cPointer(), (int)(lineCount))
}

func (this *QPainter) DrawLines6(pointPairs []QPoint) {
	pointPairs_CArray := (*[0xffff]*QPoint)(malloc(size_t(8 * len(pointPairs))))
	defer free(unsafe.Pointer(pointPairs_CArray))
	for i := range pointPairs {
		pointPairs_CArray[i] = pointPairs[i].cPointer()
	}
	pointPairs_ma := struct_miqt_array{len: size_t(len(pointPairs)), data: unsafe.Pointer(pointPairs_CArray)}
	QPainter_DrawLines6(this.h, pointPairs_ma)
}

func (this *QPainter) DrawRect(rect *QRectF) {
	QPainter_DrawRect(this.h, rect.cPointer())
}

func (this *QPainter) DrawRect2(x1 int, y1 int, w int, h int) {
	QPainter_DrawRect2(this.h, (int)(x1), (int)(y1), (int)(w), (int)(h))
}

func (this *QPainter) DrawRectWithRect(rect *QRect) {
	QPainter_DrawRectWithRect(this.h, rect.cPointer())
}

func (this *QPainter) DrawRects(rects *QRectF, rectCount int) {
	QPainter_DrawRects(this.h, rects.cPointer(), (int)(rectCount))
}

func (this *QPainter) DrawRectsWithRectangles(rectangles []QRectF) {
	rectangles_CArray := (*[0xffff]*QRectF)(malloc(size_t(8 * len(rectangles))))
	defer free(unsafe.Pointer(rectangles_CArray))
	for i := range rectangles {
		rectangles_CArray[i] = rectangles[i].cPointer()
	}
	rectangles_ma := struct_miqt_array{len: size_t(len(rectangles)), data: unsafe.Pointer(rectangles_CArray)}
	QPainter_DrawRectsWithRectangles(this.h, rectangles_ma)
}

func (this *QPainter) DrawRects2(rects *QRect, rectCount int) {
	QPainter_DrawRects2(this.h, rects.cPointer(), (int)(rectCount))
}

func (this *QPainter) DrawRects3(rectangles []QRect) {
	rectangles_CArray := (*[0xffff]*QRect)(malloc(size_t(8 * len(rectangles))))
	defer free(unsafe.Pointer(rectangles_CArray))
	for i := range rectangles {
		rectangles_CArray[i] = rectangles[i].cPointer()
	}
	rectangles_ma := struct_miqt_array{len: size_t(len(rectangles)), data: unsafe.Pointer(rectangles_CArray)}
	QPainter_DrawRects3(this.h, rectangles_ma)
}

func (this *QPainter) DrawEllipse(r *QRectF) {
	QPainter_DrawEllipse(this.h, r.cPointer())
}

func (this *QPainter) DrawEllipseWithQRect(r *QRect) {
	QPainter_DrawEllipseWithQRect(this.h, r.cPointer())
}

func (this *QPainter) DrawEllipse2(x int, y int, w int, h int) {
	QPainter_DrawEllipse2(this.h, (int)(x), (int)(y), (int)(w), (int)(h))
}

func (this *QPainter) DrawEllipse3(center *QPointF, rx float64, ry float64) {
	QPainter_DrawEllipse3(this.h, center.cPointer(), (double)(rx), (double)(ry))
}

func (this *QPainter) DrawEllipse4(center *QPoint, rx int, ry int) {
	QPainter_DrawEllipse4(this.h, center.cPointer(), (int)(rx), (int)(ry))
}

func (this *QPainter) DrawPolyline(points *QPointF, pointCount int) {
	QPainter_DrawPolyline(this.h, points.cPointer(), (int)(pointCount))
}

func (this *QPainter) DrawPolyline2(points *QPoint, pointCount int) {
	QPainter_DrawPolyline2(this.h, points.cPointer(), (int)(pointCount))
}

func (this *QPainter) DrawPolygon(points *QPointF, pointCount int) {
	QPainter_DrawPolygon(this.h, points.cPointer(), (int)(pointCount))
}

func (this *QPainter) DrawPolygon2(points *QPoint, pointCount int) {
	QPainter_DrawPolygon2(this.h, points.cPointer(), (int)(pointCount))
}

func (this *QPainter) DrawConvexPolygon(points *QPointF, pointCount int) {
	QPainter_DrawConvexPolygon(this.h, points.cPointer(), (int)(pointCount))
}

func (this *QPainter) DrawConvexPolygon2(points *QPoint, pointCount int) {
	QPainter_DrawConvexPolygon2(this.h, points.cPointer(), (int)(pointCount))
}

func (this *QPainter) DrawArc(rect *QRectF, a int, alen int) {
	QPainter_DrawArc(this.h, rect.cPointer(), (int)(a), (int)(alen))
}

func (this *QPainter) DrawArc2(param1 *QRect, a int, alen int) {
	QPainter_DrawArc2(this.h, param1.cPointer(), (int)(a), (int)(alen))
}

func (this *QPainter) DrawArc3(x int, y int, w int, h int, a int, alen int) {
	QPainter_DrawArc3(this.h, (int)(x), (int)(y), (int)(w), (int)(h), (int)(a), (int)(alen))
}

func (this *QPainter) DrawPie(rect *QRectF, a int, alen int) {
	QPainter_DrawPie(this.h, rect.cPointer(), (int)(a), (int)(alen))
}

func (this *QPainter) DrawPie2(x int, y int, w int, h int, a int, alen int) {
	QPainter_DrawPie2(this.h, (int)(x), (int)(y), (int)(w), (int)(h), (int)(a), (int)(alen))
}

func (this *QPainter) DrawPie3(param1 *QRect, a int, alen int) {
	QPainter_DrawPie3(this.h, param1.cPointer(), (int)(a), (int)(alen))
}

func (this *QPainter) DrawChord(rect *QRectF, a int, alen int) {
	QPainter_DrawChord(this.h, rect.cPointer(), (int)(a), (int)(alen))
}

func (this *QPainter) DrawChord2(x int, y int, w int, h int, a int, alen int) {
	QPainter_DrawChord2(this.h, (int)(x), (int)(y), (int)(w), (int)(h), (int)(a), (int)(alen))
}

func (this *QPainter) DrawChord3(param1 *QRect, a int, alen int) {
	QPainter_DrawChord3(this.h, param1.cPointer(), (int)(a), (int)(alen))
}

func (this *QPainter) DrawRoundedRect(rect *QRectF, xRadius float64, yRadius float64) {
	QPainter_DrawRoundedRect(this.h, rect.cPointer(), (double)(xRadius), (double)(yRadius))
}

func (this *QPainter) DrawRoundedRect2(x int, y int, w int, h int, xRadius float64, yRadius float64) {
	QPainter_DrawRoundedRect2(this.h, (int)(x), (int)(y), (int)(w), (int)(h), (double)(xRadius), (double)(yRadius))
}

func (this *QPainter) DrawRoundedRect3(rect *QRect, xRadius float64, yRadius float64) {
	QPainter_DrawRoundedRect3(this.h, rect.cPointer(), (double)(xRadius), (double)(yRadius))
}

func (this *QPainter) DrawTiledPixmap(rect *QRectF, pm *QPixmap) {
	QPainter_DrawTiledPixmap(this.h, rect.cPointer(), pm.cPointer())
}

func (this *QPainter) DrawTiledPixmap2(x int, y int, w int, h int, param5 *QPixmap) {
	QPainter_DrawTiledPixmap2(this.h, (int)(x), (int)(y), (int)(w), (int)(h), param5.cPointer())
}

func (this *QPainter) DrawTiledPixmap3(param1 *QRect, param2 *QPixmap) {
	QPainter_DrawTiledPixmap3(this.h, param1.cPointer(), param2.cPointer())
}

func (this *QPainter) DrawPicture(p *QPointF, picture *QPicture) {
	QPainter_DrawPicture(this.h, p.cPointer(), picture.cPointer())
}

func (this *QPainter) DrawPicture2(x int, y int, picture *QPicture) {
	QPainter_DrawPicture2(this.h, (int)(x), (int)(y), picture.cPointer())
}

func (this *QPainter) DrawPicture3(p *QPoint, picture *QPicture) {
	QPainter_DrawPicture3(this.h, p.cPointer(), picture.cPointer())
}

func (this *QPainter) DrawPixmap(targetRect *QRectF, pixmap *QPixmap, sourceRect *QRectF) {
	QPainter_DrawPixmap(this.h, targetRect.cPointer(), pixmap.cPointer(), sourceRect.cPointer())
}

func (this *QPainter) DrawPixmap2(targetRect *QRect, pixmap *QPixmap, sourceRect *QRect) {
	QPainter_DrawPixmap2(this.h, targetRect.cPointer(), pixmap.cPointer(), sourceRect.cPointer())
}

func (this *QPainter) DrawPixmap3(x int, y int, w int, h int, pm *QPixmap, sx int, sy int, sw int, sh int) {
	QPainter_DrawPixmap3(this.h, (int)(x), (int)(y), (int)(w), (int)(h), pm.cPointer(), (int)(sx), (int)(sy), (int)(sw), (int)(sh))
}

func (this *QPainter) DrawPixmap4(x int, y int, pm *QPixmap, sx int, sy int, sw int, sh int) {
	QPainter_DrawPixmap4(this.h, (int)(x), (int)(y), pm.cPointer(), (int)(sx), (int)(sy), (int)(sw), (int)(sh))
}

func (this *QPainter) DrawPixmap5(p *QPointF, pm *QPixmap, sr *QRectF) {
	QPainter_DrawPixmap5(this.h, p.cPointer(), pm.cPointer(), sr.cPointer())
}

func (this *QPainter) DrawPixmap6(p *QPoint, pm *QPixmap, sr *QRect) {
	QPainter_DrawPixmap6(this.h, p.cPointer(), pm.cPointer(), sr.cPointer())
}

func (this *QPainter) DrawPixmap7(p *QPointF, pm *QPixmap) {
	QPainter_DrawPixmap7(this.h, p.cPointer(), pm.cPointer())
}

func (this *QPainter) DrawPixmap8(p *QPoint, pm *QPixmap) {
	QPainter_DrawPixmap8(this.h, p.cPointer(), pm.cPointer())
}

func (this *QPainter) DrawPixmap9(x int, y int, pm *QPixmap) {
	QPainter_DrawPixmap9(this.h, (int)(x), (int)(y), pm.cPointer())
}

func (this *QPainter) DrawPixmap10(r *QRect, pm *QPixmap) {
	QPainter_DrawPixmap10(this.h, r.cPointer(), pm.cPointer())
}

func (this *QPainter) DrawPixmap11(x int, y int, w int, h int, pm *QPixmap) {
	QPainter_DrawPixmap11(this.h, (int)(x), (int)(y), (int)(w), (int)(h), pm.cPointer())
}

func (this *QPainter) DrawPixmapFragments(fragments *PixmapFragment, fragmentCount int, pixmap *QPixmap) {
	QPainter_DrawPixmapFragments(this.h, fragments, (int)(fragmentCount), pixmap.cPointer())
}

func (this *QPainter) DrawImage(targetRect *QRectF, image *QImage, sourceRect *QRectF) {
	QPainter_DrawImage(this.h, targetRect.cPointer(), image.cPointer(), sourceRect.cPointer())
}

func (this *QPainter) DrawImage2(targetRect *QRect, image *QImage, sourceRect *QRect) {
	QPainter_DrawImage2(this.h, targetRect.cPointer(), image.cPointer(), sourceRect.cPointer())
}

func (this *QPainter) DrawImage3(p *QPointF, image *QImage, sr *QRectF) {
	QPainter_DrawImage3(this.h, p.cPointer(), image.cPointer(), sr.cPointer())
}

func (this *QPainter) DrawImage4(p *QPoint, image *QImage, sr *QRect) {
	QPainter_DrawImage4(this.h, p.cPointer(), image.cPointer(), sr.cPointer())
}

func (this *QPainter) DrawImage5(r *QRectF, image *QImage) {
	QPainter_DrawImage5(this.h, r.cPointer(), image.cPointer())
}

func (this *QPainter) DrawImage6(r *QRect, image *QImage) {
	QPainter_DrawImage6(this.h, r.cPointer(), image.cPointer())
}

func (this *QPainter) DrawImage7(p *QPointF, image *QImage) {
	QPainter_DrawImage7(this.h, p.cPointer(), image.cPointer())
}

func (this *QPainter) DrawImage8(p *QPoint, image *QImage) {
	QPainter_DrawImage8(this.h, p.cPointer(), image.cPointer())
}

func (this *QPainter) DrawImage9(x int, y int, image *QImage) {
	QPainter_DrawImage9(this.h, (int)(x), (int)(y), image.cPointer())
}

func (this *QPainter) SetLayoutDirection(direction LayoutDirection) {
	QPainter_SetLayoutDirection(this.h, (int)(direction))
}

func (this *QPainter) LayoutDirection() LayoutDirection {
	return (LayoutDirection)(QPainter_LayoutDirection(this.h))
}

func (this *QPainter) DrawGlyphRun(position *QPointF, glyphRun *QGlyphRun) {
	QPainter_DrawGlyphRun(this.h, position.cPointer(), glyphRun.cPointer())
}

func (this *QPainter) DrawStaticText(topLeftPosition *QPointF, staticText *QStaticText) {
	QPainter_DrawStaticText(this.h, topLeftPosition.cPointer(), staticText.cPointer())
}

func (this *QPainter) DrawStaticText2(topLeftPosition *QPoint, staticText *QStaticText) {
	QPainter_DrawStaticText2(this.h, topLeftPosition.cPointer(), staticText.cPointer())
}

func (this *QPainter) DrawStaticText3(left int, top int, staticText *QStaticText) {
	QPainter_DrawStaticText3(this.h, (int)(left), (int)(top), staticText.cPointer())
}

func (this *QPainter) DrawText(p *QPointF, s string) {
	s_ms := struct_miqt_string{}
	s_ms.data = CString(s)
	s_ms.len = size_t(len(s))
	defer free(unsafe.Pointer(s_ms.data))
	QPainter_DrawText(this.h, p.cPointer(), s_ms)
}

func (this *QPainter) DrawText2(p *QPoint, s string) {
	s_ms := struct_miqt_string{}
	s_ms.data = CString(s)
	s_ms.len = size_t(len(s))
	defer free(unsafe.Pointer(s_ms.data))
	QPainter_DrawText2(this.h, p.cPointer(), s_ms)
}

func (this *QPainter) DrawText3(x int, y int, s string) {
	s_ms := struct_miqt_string{}
	s_ms.data = CString(s)
	s_ms.len = size_t(len(s))
	defer free(unsafe.Pointer(s_ms.data))
	QPainter_DrawText3(this.h, (int)(x), (int)(y), s_ms)
}

func (this *QPainter) DrawText4(p *QPointF, str string, tf int, justificationPadding int) {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))
	QPainter_DrawText4(this.h, p.cPointer(), str_ms, (int)(tf), (int)(justificationPadding))
}

func (this *QPainter) DrawText5(r *QRectF, flags int, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QPainter_DrawText5(this.h, r.cPointer(), (int)(flags), text_ms)
}

func (this *QPainter) DrawText6(r *QRect, flags int, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QPainter_DrawText6(this.h, r.cPointer(), (int)(flags), text_ms)
}

func (this *QPainter) DrawText7(x int, y int, w int, h int, flags int, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QPainter_DrawText7(this.h, (int)(x), (int)(y), (int)(w), (int)(h), (int)(flags), text_ms)
}

func (this *QPainter) DrawText8(r *QRectF, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QPainter_DrawText8(this.h, r.cPointer(), text_ms)
}

func (this *QPainter) BoundingRect(rect *QRectF, flags int, text string) *QRectF {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRectF(QPainter_BoundingRect(this.h, rect.cPointer(), (int)(flags), text_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainter) BoundingRect2(rect *QRect, flags int, text string) *QRect {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRect(QPainter_BoundingRect2(this.h, rect.cPointer(), (int)(flags), text_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainter) BoundingRect3(x int, y int, w int, h int, flags int, text string) *QRect {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRect(QPainter_BoundingRect3(this.h, (int)(x), (int)(y), (int)(w), (int)(h), (int)(flags), text_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainter) BoundingRect4(rect *QRectF, text string) *QRectF {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRectF(QPainter_BoundingRect4(this.h, rect.cPointer(), text_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainter) DrawTextItem(p *QPointF, ti *QTextItem) {
	QPainter_DrawTextItem(this.h, p.cPointer(), ti.cPointer())
}

func (this *QPainter) DrawTextItem2(x int, y int, ti *QTextItem) {
	QPainter_DrawTextItem2(this.h, (int)(x), (int)(y), ti.cPointer())
}

func (this *QPainter) DrawTextItem3(p *QPoint, ti *QTextItem) {
	QPainter_DrawTextItem3(this.h, p.cPointer(), ti.cPointer())
}

func (this *QPainter) FillRect(param1 *QRectF, param2 *QBrush) {
	QPainter_FillRect(this.h, param1.cPointer(), param2.cPointer())
}

func (this *QPainter) FillRect2(x int, y int, w int, h int, param5 *QBrush) {
	QPainter_FillRect2(this.h, (int)(x), (int)(y), (int)(w), (int)(h), param5.cPointer())
}

func (this *QPainter) FillRect3(param1 *QRect, param2 *QBrush) {
	QPainter_FillRect3(this.h, param1.cPointer(), param2.cPointer())
}

func (this *QPainter) FillRect4(param1 *QRectF, color *QColor) {
	QPainter_FillRect4(this.h, param1.cPointer(), color.cPointer())
}

func (this *QPainter) FillRect5(x int, y int, w int, h int, color *QColor) {
	QPainter_FillRect5(this.h, (int)(x), (int)(y), (int)(w), (int)(h), color.cPointer())
}

func (this *QPainter) FillRect6(param1 *QRect, color *QColor) {
	QPainter_FillRect6(this.h, param1.cPointer(), color.cPointer())
}

func (this *QPainter) FillRect7(x int, y int, w int, h int, c GlobalColor) {
	QPainter_FillRect7(this.h, (int)(x), (int)(y), (int)(w), (int)(h), (int)(c))
}

func (this *QPainter) FillRect8(r *QRect, c GlobalColor) {
	QPainter_FillRect8(this.h, r.cPointer(), (int)(c))
}

func (this *QPainter) FillRect9(r *QRectF, c GlobalColor) {
	QPainter_FillRect9(this.h, r.cPointer(), (int)(c))
}

func (this *QPainter) FillRect10(x int, y int, w int, h int, style BrushStyle) {
	QPainter_FillRect10(this.h, (int)(x), (int)(y), (int)(w), (int)(h), (int)(style))
}

func (this *QPainter) FillRect11(r *QRect, style BrushStyle) {
	QPainter_FillRect11(this.h, r.cPointer(), (int)(style))
}

func (this *QPainter) FillRect12(r *QRectF, style BrushStyle) {
	QPainter_FillRect12(this.h, r.cPointer(), (int)(style))
}

func (this *QPainter) FillRect13(x int, y int, w int, h int, preset QGradient__Preset) {
	QPainter_FillRect13(this.h, (int)(x), (int)(y), (int)(w), (int)(h), (int)(preset))
}

func (this *QPainter) FillRect14(r *QRect, preset QGradient__Preset) {
	QPainter_FillRect14(this.h, r.cPointer(), (int)(preset))
}

func (this *QPainter) FillRect15(r *QRectF, preset QGradient__Preset) {
	QPainter_FillRect15(this.h, r.cPointer(), (int)(preset))
}

func (this *QPainter) EraseRect(param1 *QRectF) {
	QPainter_EraseRect(this.h, param1.cPointer())
}

func (this *QPainter) EraseRect2(x int, y int, w int, h int) {
	QPainter_EraseRect2(this.h, (int)(x), (int)(y), (int)(w), (int)(h))
}

func (this *QPainter) EraseRectWithQRect(param1 *QRect) {
	QPainter_EraseRectWithQRect(this.h, param1.cPointer())
}

func (this *QPainter) SetRenderHint(hint RenderHint) {
	QPainter_SetRenderHint(this.h, hint)
}

func (this *QPainter) SetRenderHints(hints RenderHints) {
	QPainter_SetRenderHints(this.h, hints)
}

func (this *QPainter) RenderHints() RenderHints {
	xxxxxxxxx
}

func (this *QPainter) TestRenderHint(hint RenderHint) bool {
	return (bool)(QPainter_TestRenderHint(this.h, hint))
}

func (this *QPainter) PaintEngine() *QPaintEngine {
	return newQPaintEngine(QPainter_PaintEngine(this.h))
}

func (this *QPainter) BeginNativePainting() {
	QPainter_BeginNativePainting(this.h)
}

func (this *QPainter) EndNativePainting() {
	QPainter_EndNativePainting(this.h)
}

func (this *QPainter) SetClipRect22(param1 *QRectF, op ClipOperation) {
	QPainter_SetClipRect22(this.h, param1.cPointer(), (int)(op))
}

func (this *QPainter) SetClipRect23(param1 *QRect, op ClipOperation) {
	QPainter_SetClipRect23(this.h, param1.cPointer(), (int)(op))
}

func (this *QPainter) SetClipRect5(x int, y int, w int, h int, op ClipOperation) {
	QPainter_SetClipRect5(this.h, (int)(x), (int)(y), (int)(w), (int)(h), (int)(op))
}

func (this *QPainter) SetClipRegion2(param1 *QRegion, op ClipOperation) {
	QPainter_SetClipRegion2(this.h, param1.cPointer(), (int)(op))
}

func (this *QPainter) SetClipPath2(path *QPainterPath, op ClipOperation) {
	QPainter_SetClipPath2(this.h, path.cPointer(), (int)(op))
}

func (this *QPainter) SetTransform2(transform *QTransform, combine bool) {
	QPainter_SetTransform2(this.h, transform.cPointer(), (bool)(combine))
}

func (this *QPainter) SetWorldTransform2(matrix *QTransform, combine bool) {
	QPainter_SetWorldTransform2(this.h, matrix.cPointer(), (bool)(combine))
}

func (this *QPainter) DrawPolygon32(points *QPointF, pointCount int, fillRule FillRule) {
	QPainter_DrawPolygon32(this.h, points.cPointer(), (int)(pointCount), (int)(fillRule))
}

func (this *QPainter) DrawPolygon33(points *QPoint, pointCount int, fillRule FillRule) {
	QPainter_DrawPolygon33(this.h, points.cPointer(), (int)(pointCount), (int)(fillRule))
}

func (this *QPainter) DrawRoundedRect4(rect *QRectF, xRadius float64, yRadius float64, mode SizeMode) {
	QPainter_DrawRoundedRect4(this.h, rect.cPointer(), (double)(xRadius), (double)(yRadius), (int)(mode))
}

func (this *QPainter) DrawRoundedRect7(x int, y int, w int, h int, xRadius float64, yRadius float64, mode SizeMode) {
	QPainter_DrawRoundedRect7(this.h, (int)(x), (int)(y), (int)(w), (int)(h), (double)(xRadius), (double)(yRadius), (int)(mode))
}

func (this *QPainter) DrawRoundedRect42(rect *QRect, xRadius float64, yRadius float64, mode SizeMode) {
	QPainter_DrawRoundedRect42(this.h, rect.cPointer(), (double)(xRadius), (double)(yRadius), (int)(mode))
}

func (this *QPainter) DrawTiledPixmap32(rect *QRectF, pm *QPixmap, offset *QPointF) {
	QPainter_DrawTiledPixmap32(this.h, rect.cPointer(), pm.cPointer(), offset.cPointer())
}

func (this *QPainter) DrawTiledPixmap6(x int, y int, w int, h int, param5 *QPixmap, sx int) {
	QPainter_DrawTiledPixmap6(this.h, (int)(x), (int)(y), (int)(w), (int)(h), param5.cPointer(), (int)(sx))
}

func (this *QPainter) DrawTiledPixmap7(x int, y int, w int, h int, param5 *QPixmap, sx int, sy int) {
	QPainter_DrawTiledPixmap7(this.h, (int)(x), (int)(y), (int)(w), (int)(h), param5.cPointer(), (int)(sx), (int)(sy))
}

func (this *QPainter) DrawTiledPixmap33(param1 *QRect, param2 *QPixmap, param3 *QPoint) {
	QPainter_DrawTiledPixmap33(this.h, param1.cPointer(), param2.cPointer(), param3.cPointer())
}

func (this *QPainter) DrawPixmapFragments4(fragments *PixmapFragment, fragmentCount int, pixmap *QPixmap, hints PixmapFragmentHints) {
	QPainter_DrawPixmapFragments4(this.h, fragments, (int)(fragmentCount), pixmap.cPointer(), hints)
}

func (this *QPainter) DrawImage42(targetRect *QRectF, image *QImage, sourceRect *QRectF, flags ImageConversionFlag) {
	QPainter_DrawImage42(this.h, targetRect.cPointer(), image.cPointer(), sourceRect.cPointer(), (int)(flags))
}

func (this *QPainter) DrawImage43(targetRect *QRect, image *QImage, sourceRect *QRect, flags ImageConversionFlag) {
	QPainter_DrawImage43(this.h, targetRect.cPointer(), image.cPointer(), sourceRect.cPointer(), (int)(flags))
}

func (this *QPainter) DrawImage44(p *QPointF, image *QImage, sr *QRectF, flags ImageConversionFlag) {
	QPainter_DrawImage44(this.h, p.cPointer(), image.cPointer(), sr.cPointer(), (int)(flags))
}

func (this *QPainter) DrawImage45(p *QPoint, image *QImage, sr *QRect, flags ImageConversionFlag) {
	QPainter_DrawImage45(this.h, p.cPointer(), image.cPointer(), sr.cPointer(), (int)(flags))
}

func (this *QPainter) DrawImage46(x int, y int, image *QImage, sx int) {
	QPainter_DrawImage46(this.h, (int)(x), (int)(y), image.cPointer(), (int)(sx))
}

func (this *QPainter) DrawImage52(x int, y int, image *QImage, sx int, sy int) {
	QPainter_DrawImage52(this.h, (int)(x), (int)(y), image.cPointer(), (int)(sx), (int)(sy))
}

func (this *QPainter) DrawImage62(x int, y int, image *QImage, sx int, sy int, sw int) {
	QPainter_DrawImage62(this.h, (int)(x), (int)(y), image.cPointer(), (int)(sx), (int)(sy), (int)(sw))
}

func (this *QPainter) DrawImage72(x int, y int, image *QImage, sx int, sy int, sw int, sh int) {
	QPainter_DrawImage72(this.h, (int)(x), (int)(y), image.cPointer(), (int)(sx), (int)(sy), (int)(sw), (int)(sh))
}

func (this *QPainter) DrawImage82(x int, y int, image *QImage, sx int, sy int, sw int, sh int, flags ImageConversionFlag) {
	QPainter_DrawImage82(this.h, (int)(x), (int)(y), image.cPointer(), (int)(sx), (int)(sy), (int)(sw), (int)(sh), (int)(flags))
}

func (this *QPainter) DrawText42(r *QRectF, flags int, text string, br *QRectF) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QPainter_DrawText42(this.h, r.cPointer(), (int)(flags), text_ms, br.cPointer())
}

func (this *QPainter) DrawText43(r *QRect, flags int, text string, br *QRect) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QPainter_DrawText43(this.h, r.cPointer(), (int)(flags), text_ms, br.cPointer())
}

func (this *QPainter) DrawText72(x int, y int, w int, h int, flags int, text string, br *QRect) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QPainter_DrawText72(this.h, (int)(x), (int)(y), (int)(w), (int)(h), (int)(flags), text_ms, br.cPointer())
}

func (this *QPainter) DrawText32(r *QRectF, text string, o *QTextOption) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QPainter_DrawText32(this.h, r.cPointer(), text_ms, o.cPointer())
}

func (this *QPainter) BoundingRect32(rect *QRectF, text string, o *QTextOption) *QRectF {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	_goptr := newQRectF(QPainter_BoundingRect32(this.h, rect.cPointer(), text_ms, o.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainter) SetRenderHint2(hint RenderHint, on bool) {
	QPainter_SetRenderHint2(this.h, hint, (bool)(on))
}

func (this *QPainter) SetRenderHints2(hints RenderHints, on bool) {
	QPainter_SetRenderHints2(this.h, hints, (bool)(on))
}

type QPainter__PixmapFragment struct {
	h          uintptr
	isSubclass bool
}

// NewQPainter__PixmapFragment constructs a new QPainter::PixmapFragment object.
func NewQPainter__PixmapFragment() *QPainter__PixmapFragment {
	g := newQPainter__PixmapFragment(QPainter__PixmapFragment_new())
	g.isSubclass = true
	return g
}

// NewQPainter__PixmapFragment2 constructs a new QPainter::PixmapFragment object.
func NewQPainter__PixmapFragment2(param1 *PixmapFragment) *QPainter__PixmapFragment {
	g := newQPainter__PixmapFragment(QPainter__PixmapFragment_new2(param1))
	g.isSubclass = true
	return g
}

func QPainter__PixmapFragment_Create(pos *QPointF, sourceRect *QRectF) PixmapFragment {
	xxxxxxxxx
}

func QPainter__PixmapFragment_Create3(pos *QPointF, sourceRect *QRectF, scaleX float64) PixmapFragment {
	xxxxxxxxx
}

func QPainter__PixmapFragment_Create4(pos *QPointF, sourceRect *QRectF, scaleX float64, scaleY float64) PixmapFragment {
	xxxxxxxxx
}

func QPainter__PixmapFragment_Create5(pos *QPointF, sourceRect *QRectF, scaleX float64, scaleY float64, rotation float64) PixmapFragment {
	xxxxxxxxx
}

func QPainter__PixmapFragment_Create6(pos *QPointF, sourceRect *QRectF, scaleX float64, scaleY float64, rotation float64, opacity float64) PixmapFragment {
	xxxxxxxxx
}
