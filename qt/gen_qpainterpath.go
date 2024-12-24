package qt

import (
	"unsafe"
)

type QPainterPath__ElementType int

const (
	QPainterPath__MoveToElement      QPainterPath__ElementType = 0
	QPainterPath__LineToElement      QPainterPath__ElementType = 1
	QPainterPath__CurveToElement     QPainterPath__ElementType = 2
	QPainterPath__CurveToDataElement QPainterPath__ElementType = 3
)

type QPainterPath struct {
	h          uintptr
	isSubclass bool
}

// NewQPainterPath constructs a new QPainterPath object.
func NewQPainterPath() *QPainterPath {
	ret := newQPainterPath(QPainterPath_new())
	ret.isSubclass = true
	return ret
}

// NewQPainterPath2 constructs a new QPainterPath object.
func NewQPainterPath2(startPoint *QPointF) *QPainterPath {
	ret := newQPainterPath(QPainterPath_new2(startPoint.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPainterPath3 constructs a new QPainterPath object.
func NewQPainterPath3(other *QPainterPath) *QPainterPath {
	ret := newQPainterPath(QPainterPath_new3(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QPainterPath) OperatorAssign(other *QPainterPath) {
	QPainterPath_OperatorAssign(this.h, other.cPointer())
}

func (this *QPainterPath) Swap(other *QPainterPath) {
	QPainterPath_Swap(this.h, other.cPointer())
}

func (this *QPainterPath) Clear() {
	QPainterPath_Clear(this.h)
}

func (this *QPainterPath) Reserve(size int) {
	QPainterPath_Reserve(this.h, (int)(size))
}

func (this *QPainterPath) Capacity() int {
	return (int)(QPainterPath_Capacity(this.h))
}

func (this *QPainterPath) CloseSubpath() {
	QPainterPath_CloseSubpath(this.h)
}

func (this *QPainterPath) MoveTo(p *QPointF) {
	QPainterPath_MoveTo(this.h, p.cPointer())
}

func (this *QPainterPath) MoveTo2(x float64, y float64) {
	QPainterPath_MoveTo2(this.h, (double)(x), (double)(y))
}

func (this *QPainterPath) LineTo(p *QPointF) {
	QPainterPath_LineTo(this.h, p.cPointer())
}

func (this *QPainterPath) LineTo2(x float64, y float64) {
	QPainterPath_LineTo2(this.h, (double)(x), (double)(y))
}

func (this *QPainterPath) ArcMoveTo(rect *QRectF, angle float64) {
	QPainterPath_ArcMoveTo(this.h, rect.cPointer(), (double)(angle))
}

func (this *QPainterPath) ArcMoveTo2(x float64, y float64, w float64, h float64, angle float64) {
	QPainterPath_ArcMoveTo2(this.h, (double)(x), (double)(y), (double)(w), (double)(h), (double)(angle))
}

func (this *QPainterPath) ArcTo(rect *QRectF, startAngle float64, arcLength float64) {
	QPainterPath_ArcTo(this.h, rect.cPointer(), (double)(startAngle), (double)(arcLength))
}

func (this *QPainterPath) ArcTo2(x float64, y float64, w float64, h float64, startAngle float64, arcLength float64) {
	QPainterPath_ArcTo2(this.h, (double)(x), (double)(y), (double)(w), (double)(h), (double)(startAngle), (double)(arcLength))
}

func (this *QPainterPath) CubicTo(ctrlPt1 *QPointF, ctrlPt2 *QPointF, endPt *QPointF) {
	QPainterPath_CubicTo(this.h, ctrlPt1.cPointer(), ctrlPt2.cPointer(), endPt.cPointer())
}

func (this *QPainterPath) CubicTo2(ctrlPt1x float64, ctrlPt1y float64, ctrlPt2x float64, ctrlPt2y float64, endPtx float64, endPty float64) {
	QPainterPath_CubicTo2(this.h, (double)(ctrlPt1x), (double)(ctrlPt1y), (double)(ctrlPt2x), (double)(ctrlPt2y), (double)(endPtx), (double)(endPty))
}

func (this *QPainterPath) QuadTo(ctrlPt *QPointF, endPt *QPointF) {
	QPainterPath_QuadTo(this.h, ctrlPt.cPointer(), endPt.cPointer())
}

func (this *QPainterPath) QuadTo2(ctrlPtx float64, ctrlPty float64, endPtx float64, endPty float64) {
	QPainterPath_QuadTo2(this.h, (double)(ctrlPtx), (double)(ctrlPty), (double)(endPtx), (double)(endPty))
}

func (this *QPainterPath) CurrentPosition() *QPointF {
	_goptr := newQPointF(QPainterPath_CurrentPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainterPath) AddRect(rect *QRectF) {
	QPainterPath_AddRect(this.h, rect.cPointer())
}

func (this *QPainterPath) AddRect2(x float64, y float64, w float64, h float64) {
	QPainterPath_AddRect2(this.h, (double)(x), (double)(y), (double)(w), (double)(h))
}

func (this *QPainterPath) AddEllipse(rect *QRectF) {
	QPainterPath_AddEllipse(this.h, rect.cPointer())
}

func (this *QPainterPath) AddEllipse2(x float64, y float64, w float64, h float64) {
	QPainterPath_AddEllipse2(this.h, (double)(x), (double)(y), (double)(w), (double)(h))
}

func (this *QPainterPath) AddEllipse3(center *QPointF, rx float64, ry float64) {
	QPainterPath_AddEllipse3(this.h, center.cPointer(), (double)(rx), (double)(ry))
}

func (this *QPainterPath) AddText(point *QPointF, f *QFont, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QPainterPath_AddText(this.h, point.cPointer(), f.cPointer(), text_ms)
}

func (this *QPainterPath) AddText2(x float64, y float64, f *QFont, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QPainterPath_AddText2(this.h, (double)(x), (double)(y), f.cPointer(), text_ms)
}

func (this *QPainterPath) AddPath(path *QPainterPath) {
	QPainterPath_AddPath(this.h, path.cPointer())
}

func (this *QPainterPath) AddRegion(region *QRegion) {
	QPainterPath_AddRegion(this.h, region.cPointer())
}

func (this *QPainterPath) AddRoundedRect(rect *QRectF, xRadius float64, yRadius float64) {
	QPainterPath_AddRoundedRect(this.h, rect.cPointer(), (double)(xRadius), (double)(yRadius))
}

func (this *QPainterPath) AddRoundedRect2(x float64, y float64, w float64, h float64, xRadius float64, yRadius float64) {
	QPainterPath_AddRoundedRect2(this.h, (double)(x), (double)(y), (double)(w), (double)(h), (double)(xRadius), (double)(yRadius))
}

func (this *QPainterPath) ConnectPath(path *QPainterPath) {
	QPainterPath_ConnectPath(this.h, path.cPointer())
}

func (this *QPainterPath) Contains(pt *QPointF) bool {
	return (bool)(QPainterPath_Contains(this.h, pt.cPointer()))
}

func (this *QPainterPath) ContainsWithRect(rect *QRectF) bool {
	return (bool)(QPainterPath_ContainsWithRect(this.h, rect.cPointer()))
}

func (this *QPainterPath) Intersects(rect *QRectF) bool {
	return (bool)(QPainterPath_Intersects(this.h, rect.cPointer()))
}

func (this *QPainterPath) Translate(dx float64, dy float64) {
	QPainterPath_Translate(this.h, (double)(dx), (double)(dy))
}

func (this *QPainterPath) TranslateWithOffset(offset *QPointF) {
	QPainterPath_TranslateWithOffset(this.h, offset.cPointer())
}

func (this *QPainterPath) Translated(dx float64, dy float64) *QPainterPath {
	_goptr := newQPainterPath(QPainterPath_Translated(this.h, (double)(dx), (double)(dy)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainterPath) TranslatedWithOffset(offset *QPointF) *QPainterPath {
	_goptr := newQPainterPath(QPainterPath_TranslatedWithOffset(this.h, offset.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainterPath) BoundingRect() *QRectF {
	_goptr := newQRectF(QPainterPath_BoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainterPath) ControlPointRect() *QRectF {
	_goptr := newQRectF(QPainterPath_ControlPointRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainterPath) FillRule() FillRule {
	return (FillRule)(QPainterPath_FillRule(this.h))
}

func (this *QPainterPath) SetFillRule(fillRule FillRule) {
	QPainterPath_SetFillRule(this.h, (int)(fillRule))
}

func (this *QPainterPath) IsEmpty() bool {
	return (bool)(QPainterPath_IsEmpty(this.h))
}

func (this *QPainterPath) ToReversed() *QPainterPath {
	_goptr := newQPainterPath(QPainterPath_ToReversed(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainterPath) ElementCount() int {
	return (int)(QPainterPath_ElementCount(this.h))
}

func (this *QPainterPath) ElementAt(i int) *QPainterPath__Element {
	_goptr := newQPainterPath__Element(QPainterPath_ElementAt(this.h, (int)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainterPath) SetElementPositionAt(i int, x float64, y float64) {
	QPainterPath_SetElementPositionAt(this.h, (int)(i), (double)(x), (double)(y))
}

func (this *QPainterPath) Length() float64 {
	return (float64)(QPainterPath_Length(this.h))
}

func (this *QPainterPath) PercentAtLength(t float64) float64 {
	return (float64)(QPainterPath_PercentAtLength(this.h, (double)(t)))
}

func (this *QPainterPath) PointAtPercent(t float64) *QPointF {
	_goptr := newQPointF(QPainterPath_PointAtPercent(this.h, (double)(t)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainterPath) AngleAtPercent(t float64) float64 {
	return (float64)(QPainterPath_AngleAtPercent(this.h, (double)(t)))
}

func (this *QPainterPath) SlopeAtPercent(t float64) float64 {
	return (float64)(QPainterPath_SlopeAtPercent(this.h, (double)(t)))
}

func (this *QPainterPath) IntersectsWithQPainterPath(p *QPainterPath) bool {
	return (bool)(QPainterPath_IntersectsWithQPainterPath(this.h, p.cPointer()))
}

func (this *QPainterPath) ContainsWithQPainterPath(p *QPainterPath) bool {
	return (bool)(QPainterPath_ContainsWithQPainterPath(this.h, p.cPointer()))
}

func (this *QPainterPath) United(r *QPainterPath) *QPainterPath {
	_goptr := newQPainterPath(QPainterPath_United(this.h, r.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainterPath) Intersected(r *QPainterPath) *QPainterPath {
	_goptr := newQPainterPath(QPainterPath_Intersected(this.h, r.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainterPath) Subtracted(r *QPainterPath) *QPainterPath {
	_goptr := newQPainterPath(QPainterPath_Subtracted(this.h, r.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainterPath) Simplified() *QPainterPath {
	_goptr := newQPainterPath(QPainterPath_Simplified(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainterPath) OperatorEqual(other *QPainterPath) bool {
	return (bool)(QPainterPath_OperatorEqual(this.h, other.cPointer()))
}

func (this *QPainterPath) OperatorNotEqual(other *QPainterPath) bool {
	return (bool)(QPainterPath_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QPainterPath) OperatorBitwiseAnd(other *QPainterPath) *QPainterPath {
	_goptr := newQPainterPath(QPainterPath_OperatorBitwiseAnd(this.h, other.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainterPath) OperatorBitwiseOr(other *QPainterPath) *QPainterPath {
	_goptr := newQPainterPath(QPainterPath_OperatorBitwiseOr(this.h, other.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainterPath) OperatorPlus(other *QPainterPath) *QPainterPath {
	_goptr := newQPainterPath(QPainterPath_OperatorPlus(this.h, other.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainterPath) OperatorMinus(other *QPainterPath) *QPainterPath {
	_goptr := newQPainterPath(QPainterPath_OperatorMinus(this.h, other.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPainterPath) OperatorBitwiseAndAssign(other *QPainterPath) {
	QPainterPath_OperatorBitwiseAndAssign(this.h, other.cPointer())
}

func (this *QPainterPath) OperatorBitwiseOrAssign(other *QPainterPath) {
	QPainterPath_OperatorBitwiseOrAssign(this.h, other.cPointer())
}

func (this *QPainterPath) OperatorPlusAssign(other *QPainterPath) *QPainterPath {
	return newQPainterPath(QPainterPath_OperatorPlusAssign(this.h, other.cPointer()))
}

func (this *QPainterPath) OperatorMinusAssign(other *QPainterPath) *QPainterPath {
	return newQPainterPath(QPainterPath_OperatorMinusAssign(this.h, other.cPointer()))
}

func (this *QPainterPath) AddRoundedRect4(rect *QRectF, xRadius float64, yRadius float64, mode SizeMode) {
	QPainterPath_AddRoundedRect4(this.h, rect.cPointer(), (double)(xRadius), (double)(yRadius), (int)(mode))
}

func (this *QPainterPath) AddRoundedRect7(x float64, y float64, w float64, h float64, xRadius float64, yRadius float64, mode SizeMode) {
	QPainterPath_AddRoundedRect7(this.h, (double)(x), (double)(y), (double)(w), (double)(h), (double)(xRadius), (double)(yRadius), (int)(mode))
}

type QPainterPathStroker struct {
	h          uintptr
	isSubclass bool
}

// NewQPainterPathStroker constructs a new QPainterPathStroker object.
func NewQPainterPathStroker() *QPainterPathStroker {
	ret := newQPainterPathStroker(QPainterPathStroker_new())
	ret.isSubclass = true
	return ret
}

// NewQPainterPathStroker2 constructs a new QPainterPathStroker object.
func NewQPainterPathStroker2(pen *QPen) *QPainterPathStroker {
	ret := newQPainterPathStroker(QPainterPathStroker_new2(pen.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QPainterPathStroker) SetWidth(width float64) {
	QPainterPathStroker_SetWidth(this.h, (double)(width))
}

func (this *QPainterPathStroker) Width() float64 {
	return (float64)(QPainterPathStroker_Width(this.h))
}

func (this *QPainterPathStroker) SetCapStyle(style PenCapStyle) {
	QPainterPathStroker_SetCapStyle(this.h, (int)(style))
}

func (this *QPainterPathStroker) CapStyle() PenCapStyle {
	return (PenCapStyle)(QPainterPathStroker_CapStyle(this.h))
}

func (this *QPainterPathStroker) SetJoinStyle(style PenJoinStyle) {
	QPainterPathStroker_SetJoinStyle(this.h, (int)(style))
}

func (this *QPainterPathStroker) JoinStyle() PenJoinStyle {
	return (PenJoinStyle)(QPainterPathStroker_JoinStyle(this.h))
}

func (this *QPainterPathStroker) SetMiterLimit(length float64) {
	QPainterPathStroker_SetMiterLimit(this.h, (double)(length))
}

func (this *QPainterPathStroker) MiterLimit() float64 {
	return (float64)(QPainterPathStroker_MiterLimit(this.h))
}

func (this *QPainterPathStroker) SetCurveThreshold(threshold float64) {
	QPainterPathStroker_SetCurveThreshold(this.h, (double)(threshold))
}

func (this *QPainterPathStroker) CurveThreshold() float64 {
	return (float64)(QPainterPathStroker_CurveThreshold(this.h))
}

func (this *QPainterPathStroker) SetDashPattern(dashPattern PenStyle) {
	QPainterPathStroker_SetDashPattern(this.h, (int)(dashPattern))
}

func (this *QPainterPathStroker) SetDashPatternWithDashPattern(dashPattern []float64) {
	dashPattern_CArray := (*[0xffff]double)(malloc(size_t(8 * len(dashPattern))))
	defer free(unsafe.Pointer(dashPattern_CArray))
	for i := range dashPattern {
		dashPattern_CArray[i] = (double)(dashPattern[i])
	}
	dashPattern_ma := struct_miqt_array{len: size_t(len(dashPattern)), data: unsafe.Pointer(dashPattern_CArray)}
	QPainterPathStroker_SetDashPatternWithDashPattern(this.h, dashPattern_ma)
}

func (this *QPainterPathStroker) DashPattern() []float64 {
	var _ma struct_miqt_array = QPainterPathStroker_DashPattern(this.h)
	_ret := make([]float64, int(_ma.len))
	_outCast := (*[0xffff]double)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (float64)(_outCast[i])
	}
	return _ret
}

func (this *QPainterPathStroker) SetDashOffset(offset float64) {
	QPainterPathStroker_SetDashOffset(this.h, (double)(offset))
}

func (this *QPainterPathStroker) DashOffset() float64 {
	return (float64)(QPainterPathStroker_DashOffset(this.h))
}

func (this *QPainterPathStroker) CreateStroke(path *QPainterPath) *QPainterPath {
	_goptr := newQPainterPath(QPainterPathStroker_CreateStroke(this.h, path.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QPainterPath__Element struct {
	h          uintptr
	isSubclass bool
}

// NewQPainterPath__Element constructs a new QPainterPath::Element object.
func NewQPainterPath__Element() *QPainterPath__Element {
	ret := newQPainterPath__Element(QPainterPath__Element_new())
	ret.isSubclass = true
	return ret
}

// NewQPainterPath__Element2 constructs a new QPainterPath::Element object.
func NewQPainterPath__Element2(param1 *Element) *QPainterPath__Element {
	ret := newQPainterPath__Element(QPainterPath__Element_new2(param1))
	ret.isSubclass = true
	return ret
}

func (this *QPainterPath__Element) IsMoveTo() bool {
	return (bool)(QPainterPath__Element_IsMoveTo(this.h))
}

func (this *QPainterPath__Element) IsLineTo() bool {
	return (bool)(QPainterPath__Element_IsLineTo(this.h))
}

func (this *QPainterPath__Element) IsCurveTo() bool {
	return (bool)(QPainterPath__Element_IsCurveTo(this.h))
}

func (this *QPainterPath__Element) OperatorEqual(e *Element) bool {
	return (bool)(QPainterPath__Element_OperatorEqual(this.h, e))
}

func (this *QPainterPath__Element) OperatorNotEqual(e *Element) bool {
	return (bool)(QPainterPath__Element_OperatorNotEqual(this.h, e))
}
