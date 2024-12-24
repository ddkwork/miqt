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
func (this *QPaintEngine) OnBegin(slot func(pdev *QPaintDevice) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_Begin(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_Begin
func miqt_exec_callback_QPaintEngine_Begin(self QPaintEngine, cb intptr_t, pdev *QPaintDevice) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(pdev *QPaintDevice) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintDevice(pdev)

	virtualReturn := gofunc(slotval1)

	return (bool)(virtualReturn)

}
func (this *QPaintEngine) OnEnd(slot func() bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_End(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_End
func miqt_exec_callback_QPaintEngine_End(self QPaintEngine, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func() bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return (bool)(virtualReturn)

}
func (this *QPaintEngine) OnUpdateState(slot func(state *QPaintEngineState)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_UpdateState(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_UpdateState
func miqt_exec_callback_QPaintEngine_UpdateState(self QPaintEngine, cb intptr_t, state *QPaintEngineState) {
	gofunc, ok := cgo.Handle(cb).Value().(func(state *QPaintEngineState))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEngineState(state)

	gofunc(slotval1)

}

func (this *QPaintEngine) callVirtualBase_DrawRects(rects *QRect, rectCount int) {

	QPaintEngine_virtualbase_DrawRects(unsafe.Pointer(this.h), rects.cPointer(), (int)(rectCount))

}
func (this *QPaintEngine) OnDrawRects(slot func(super func(rects *QRect, rectCount int), rects *QRect, rectCount int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_DrawRects(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_DrawRects
func miqt_exec_callback_QPaintEngine_DrawRects(self QPaintEngine, cb intptr_t, rects *QRect, rectCount int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rects *QRect, rectCount int), rects *QRect, rectCount int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(rects)

	slotval2 := (int)(rectCount)

	gofunc((&QPaintEngine{h: self}).callVirtualBase_DrawRects, slotval1, slotval2)

}

func (this *QPaintEngine) callVirtualBase_DrawRects2(rects *QRectF, rectCount int) {

	QPaintEngine_virtualbase_DrawRects2(unsafe.Pointer(this.h), rects.cPointer(), (int)(rectCount))

}
func (this *QPaintEngine) OnDrawRects2(slot func(super func(rects *QRectF, rectCount int), rects *QRectF, rectCount int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_DrawRects2(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_DrawRects2
func miqt_exec_callback_QPaintEngine_DrawRects2(self QPaintEngine, cb intptr_t, rects *QRectF, rectCount int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rects *QRectF, rectCount int), rects *QRectF, rectCount int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(rects)

	slotval2 := (int)(rectCount)

	gofunc((&QPaintEngine{h: self}).callVirtualBase_DrawRects2, slotval1, slotval2)

}

func (this *QPaintEngine) callVirtualBase_DrawLines(lines *QLine, lineCount int) {

	QPaintEngine_virtualbase_DrawLines(unsafe.Pointer(this.h), lines.cPointer(), (int)(lineCount))

}
func (this *QPaintEngine) OnDrawLines(slot func(super func(lines *QLine, lineCount int), lines *QLine, lineCount int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_DrawLines(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_DrawLines
func miqt_exec_callback_QPaintEngine_DrawLines(self QPaintEngine, cb intptr_t, lines *QLine, lineCount int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(lines *QLine, lineCount int), lines *QLine, lineCount int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQLine(lines)

	slotval2 := (int)(lineCount)

	gofunc((&QPaintEngine{h: self}).callVirtualBase_DrawLines, slotval1, slotval2)

}

func (this *QPaintEngine) callVirtualBase_DrawLines2(lines *QLineF, lineCount int) {

	QPaintEngine_virtualbase_DrawLines2(unsafe.Pointer(this.h), lines.cPointer(), (int)(lineCount))

}
func (this *QPaintEngine) OnDrawLines2(slot func(super func(lines *QLineF, lineCount int), lines *QLineF, lineCount int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_DrawLines2(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_DrawLines2
func miqt_exec_callback_QPaintEngine_DrawLines2(self QPaintEngine, cb intptr_t, lines *QLineF, lineCount int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(lines *QLineF, lineCount int), lines *QLineF, lineCount int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQLineF(lines)

	slotval2 := (int)(lineCount)

	gofunc((&QPaintEngine{h: self}).callVirtualBase_DrawLines2, slotval1, slotval2)

}

func (this *QPaintEngine) callVirtualBase_DrawEllipse(r *QRectF) {

	QPaintEngine_virtualbase_DrawEllipse(unsafe.Pointer(this.h), r.cPointer())

}
func (this *QPaintEngine) OnDrawEllipse(slot func(super func(r *QRectF), r *QRectF)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_DrawEllipse(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_DrawEllipse
func miqt_exec_callback_QPaintEngine_DrawEllipse(self QPaintEngine, cb intptr_t, r *QRectF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(r *QRectF), r *QRectF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(r)

	gofunc((&QPaintEngine{h: self}).callVirtualBase_DrawEllipse, slotval1)

}

func (this *QPaintEngine) callVirtualBase_DrawEllipseWithQRect(r *QRect) {

	QPaintEngine_virtualbase_DrawEllipseWithQRect(unsafe.Pointer(this.h), r.cPointer())

}
func (this *QPaintEngine) OnDrawEllipseWithQRect(slot func(super func(r *QRect), r *QRect)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_DrawEllipseWithQRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_DrawEllipseWithQRect
func miqt_exec_callback_QPaintEngine_DrawEllipseWithQRect(self QPaintEngine, cb intptr_t, r *QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(r *QRect), r *QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(r)

	gofunc((&QPaintEngine{h: self}).callVirtualBase_DrawEllipseWithQRect, slotval1)

}

func (this *QPaintEngine) callVirtualBase_DrawPath(path *QPainterPath) {

	QPaintEngine_virtualbase_DrawPath(unsafe.Pointer(this.h), path.cPointer())

}
func (this *QPaintEngine) OnDrawPath(slot func(super func(path *QPainterPath), path *QPainterPath)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_DrawPath(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_DrawPath
func miqt_exec_callback_QPaintEngine_DrawPath(self QPaintEngine, cb intptr_t, path *QPainterPath) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(path *QPainterPath), path *QPainterPath))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainterPath(path)

	gofunc((&QPaintEngine{h: self}).callVirtualBase_DrawPath, slotval1)

}

func (this *QPaintEngine) callVirtualBase_DrawPoints(points *QPointF, pointCount int) {

	QPaintEngine_virtualbase_DrawPoints(unsafe.Pointer(this.h), points.cPointer(), (int)(pointCount))

}
func (this *QPaintEngine) OnDrawPoints(slot func(super func(points *QPointF, pointCount int), points *QPointF, pointCount int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_DrawPoints(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_DrawPoints
func miqt_exec_callback_QPaintEngine_DrawPoints(self QPaintEngine, cb intptr_t, points *QPointF, pointCount int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(points *QPointF, pointCount int), points *QPointF, pointCount int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPointF(points)

	slotval2 := (int)(pointCount)

	gofunc((&QPaintEngine{h: self}).callVirtualBase_DrawPoints, slotval1, slotval2)

}

func (this *QPaintEngine) callVirtualBase_DrawPoints2(points *QPoint, pointCount int) {

	QPaintEngine_virtualbase_DrawPoints2(unsafe.Pointer(this.h), points.cPointer(), (int)(pointCount))

}
func (this *QPaintEngine) OnDrawPoints2(slot func(super func(points *QPoint, pointCount int), points *QPoint, pointCount int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_DrawPoints2(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_DrawPoints2
func miqt_exec_callback_QPaintEngine_DrawPoints2(self QPaintEngine, cb intptr_t, points *QPoint, pointCount int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(points *QPoint, pointCount int), points *QPoint, pointCount int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(points)

	slotval2 := (int)(pointCount)

	gofunc((&QPaintEngine{h: self}).callVirtualBase_DrawPoints2, slotval1, slotval2)

}

func (this *QPaintEngine) callVirtualBase_DrawPolygon(points *QPointF, pointCount int, mode PolygonDrawMode) {

	QPaintEngine_virtualbase_DrawPolygon(unsafe.Pointer(this.h), points.cPointer(), (int)(pointCount), mode)

}
func (this *QPaintEngine) OnDrawPolygon(slot func(super func(points *QPointF, pointCount int, mode PolygonDrawMode), points *QPointF, pointCount int, mode PolygonDrawMode)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_DrawPolygon(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_DrawPolygon
func miqt_exec_callback_QPaintEngine_DrawPolygon(self QPaintEngine, cb intptr_t, points *QPointF, pointCount int, mode PolygonDrawMode) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(points *QPointF, pointCount int, mode PolygonDrawMode), points *QPointF, pointCount int, mode PolygonDrawMode))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPointF(points)

	slotval2 := (int)(pointCount)

	xxxxxxxxx

	gofunc((&QPaintEngine{h: self}).callVirtualBase_DrawPolygon, slotval1, slotval2, slotval3)

}

func (this *QPaintEngine) callVirtualBase_DrawPolygon2(points *QPoint, pointCount int, mode PolygonDrawMode) {

	QPaintEngine_virtualbase_DrawPolygon2(unsafe.Pointer(this.h), points.cPointer(), (int)(pointCount), mode)

}
func (this *QPaintEngine) OnDrawPolygon2(slot func(super func(points *QPoint, pointCount int, mode PolygonDrawMode), points *QPoint, pointCount int, mode PolygonDrawMode)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_DrawPolygon2(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_DrawPolygon2
func miqt_exec_callback_QPaintEngine_DrawPolygon2(self QPaintEngine, cb intptr_t, points *QPoint, pointCount int, mode PolygonDrawMode) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(points *QPoint, pointCount int, mode PolygonDrawMode), points *QPoint, pointCount int, mode PolygonDrawMode))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(points)

	slotval2 := (int)(pointCount)

	xxxxxxxxx

	gofunc((&QPaintEngine{h: self}).callVirtualBase_DrawPolygon2, slotval1, slotval2, slotval3)

}
func (this *QPaintEngine) OnDrawPixmap(slot func(r *QRectF, pm *QPixmap, sr *QRectF)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_DrawPixmap(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_DrawPixmap
func miqt_exec_callback_QPaintEngine_DrawPixmap(self QPaintEngine, cb intptr_t, r *QRectF, pm *QPixmap, sr *QRectF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(r *QRectF, pm *QPixmap, sr *QRectF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(r)

	slotval2 := newQPixmap(pm)

	slotval3 := newQRectF(sr)

	gofunc(slotval1, slotval2, slotval3)

}

func (this *QPaintEngine) callVirtualBase_DrawTextItem(p *QPointF, textItem *QTextItem) {

	QPaintEngine_virtualbase_DrawTextItem(unsafe.Pointer(this.h), p.cPointer(), textItem.cPointer())

}
func (this *QPaintEngine) OnDrawTextItem(slot func(super func(p *QPointF, textItem *QTextItem), p *QPointF, textItem *QTextItem)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_DrawTextItem(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_DrawTextItem
func miqt_exec_callback_QPaintEngine_DrawTextItem(self QPaintEngine, cb intptr_t, p *QPointF, textItem *QTextItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(p *QPointF, textItem *QTextItem), p *QPointF, textItem *QTextItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPointF(p)

	slotval2 := newQTextItem(textItem)

	gofunc((&QPaintEngine{h: self}).callVirtualBase_DrawTextItem, slotval1, slotval2)

}

func (this *QPaintEngine) callVirtualBase_DrawTiledPixmap(r *QRectF, pixmap *QPixmap, s *QPointF) {

	QPaintEngine_virtualbase_DrawTiledPixmap(unsafe.Pointer(this.h), r.cPointer(), pixmap.cPointer(), s.cPointer())

}
func (this *QPaintEngine) OnDrawTiledPixmap(slot func(super func(r *QRectF, pixmap *QPixmap, s *QPointF), r *QRectF, pixmap *QPixmap, s *QPointF)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_DrawTiledPixmap(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_DrawTiledPixmap
func miqt_exec_callback_QPaintEngine_DrawTiledPixmap(self QPaintEngine, cb intptr_t, r *QRectF, pixmap *QPixmap, s *QPointF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(r *QRectF, pixmap *QPixmap, s *QPointF), r *QRectF, pixmap *QPixmap, s *QPointF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(r)

	slotval2 := newQPixmap(pixmap)

	slotval3 := newQPointF(s)

	gofunc((&QPaintEngine{h: self}).callVirtualBase_DrawTiledPixmap, slotval1, slotval2, slotval3)

}

func (this *QPaintEngine) callVirtualBase_DrawImage(r *QRectF, pm *QImage, sr *QRectF, flags ImageConversionFlag) {

	QPaintEngine_virtualbase_DrawImage(unsafe.Pointer(this.h), r.cPointer(), pm.cPointer(), sr.cPointer(), (int)(flags))

}
func (this *QPaintEngine) OnDrawImage(slot func(super func(r *QRectF, pm *QImage, sr *QRectF, flags ImageConversionFlag), r *QRectF, pm *QImage, sr *QRectF, flags ImageConversionFlag)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_DrawImage(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_DrawImage
func miqt_exec_callback_QPaintEngine_DrawImage(self QPaintEngine, cb intptr_t, r *QRectF, pm *QImage, sr *QRectF, flags int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(r *QRectF, pm *QImage, sr *QRectF, flags ImageConversionFlag), r *QRectF, pm *QImage, sr *QRectF, flags ImageConversionFlag))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(r)

	slotval2 := newQImage(pm)

	slotval3 := newQRectF(sr)

	slotval4 := (ImageConversionFlag)(flags)

	gofunc((&QPaintEngine{h: self}).callVirtualBase_DrawImage, slotval1, slotval2, slotval3, slotval4)

}

func (this *QPaintEngine) callVirtualBase_CoordinateOffset() *QPoint {

	_goptr := newQPoint(QPaintEngine_virtualbase_CoordinateOffset(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QPaintEngine) OnCoordinateOffset(slot func(super func() *QPoint) *QPoint) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_CoordinateOffset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_CoordinateOffset
func miqt_exec_callback_QPaintEngine_CoordinateOffset(self QPaintEngine, cb intptr_t) *QPoint {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPoint) *QPoint)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPaintEngine{h: self}).callVirtualBase_CoordinateOffset)

	return virtualReturn.cPointer()

}
func (this *QPaintEngine) OnType(slot func() Type) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_Type(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_Type
func miqt_exec_callback_QPaintEngine_Type(self QPaintEngine, cb intptr_t) Type {
	gofunc, ok := cgo.Handle(cb).Value().(func() Type)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return virtualReturn

}

func (this *QPaintEngine) callVirtualBase_CreatePixmap(size QSize) *QPixmap {

	_goptr := newQPixmap(QPaintEngine_virtualbase_CreatePixmap(unsafe.Pointer(this.h), size.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QPaintEngine) OnCreatePixmap(slot func(super func(size QSize) *QPixmap, size QSize) *QPixmap) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_CreatePixmap(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_CreatePixmap
func miqt_exec_callback_QPaintEngine_CreatePixmap(self QPaintEngine, cb intptr_t, size *QSize) *QPixmap {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(size QSize) *QPixmap, size QSize) *QPixmap)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	size_goptr := newQSize(size)
	size_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval1 := *size_goptr

	virtualReturn := gofunc((&QPaintEngine{h: self}).callVirtualBase_CreatePixmap, slotval1)

	return virtualReturn.cPointer()

}

func (this *QPaintEngine) callVirtualBase_CreatePixmapFromImage(image QImage, flags ImageConversionFlag) *QPixmap {

	_goptr := newQPixmap(QPaintEngine_virtualbase_CreatePixmapFromImage(unsafe.Pointer(this.h), image.cPointer(), (int)(flags)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QPaintEngine) OnCreatePixmapFromImage(slot func(super func(image QImage, flags ImageConversionFlag) *QPixmap, image QImage, flags ImageConversionFlag) *QPixmap) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPaintEngine_override_virtual_CreatePixmapFromImage(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPaintEngine_CreatePixmapFromImage
func miqt_exec_callback_QPaintEngine_CreatePixmapFromImage(self QPaintEngine, cb intptr_t, image *QImage, flags int) *QPixmap {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(image QImage, flags ImageConversionFlag) *QPixmap, image QImage, flags ImageConversionFlag) *QPixmap)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	image_goptr := newQImage(image)
	image_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval1 := *image_goptr

	slotval2 := (ImageConversionFlag)(flags)

	virtualReturn := gofunc((&QPaintEngine{h: self}).callVirtualBase_CreatePixmapFromImage, slotval1, slotval2)

	return virtualReturn.cPointer()

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
