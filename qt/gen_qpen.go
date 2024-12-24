package qt

import (
	"unsafe"
)

type QPen struct {
	h          uintptr
	isSubclass bool
}

// NewQPen constructs a new QPen object.
func NewQPen() *QPen {
	g := newQPen(QPen_new())
	g.isSubclass = true
	return g
}

// NewQPen2 constructs a new QPen object.
func NewQPen2(param1 PenStyle) *QPen {
	g := newQPen(QPen_new2((int)(param1)))
	g.isSubclass = true
	return g
}

// NewQPen3 constructs a new QPen object.
func NewQPen3(color *QColor) *QPen {
	g := newQPen(QPen_new3(color.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPen4 constructs a new QPen object.
func NewQPen4(brush *QBrush, width float64) *QPen {
	g := newQPen(QPen_new4(brush.cPointer(), (double)(width)))
	g.isSubclass = true
	return g
}

// NewQPen5 constructs a new QPen object.
func NewQPen5(pen *QPen) *QPen {
	g := newQPen(QPen_new5(pen.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPen6 constructs a new QPen object.
func NewQPen6(brush *QBrush, width float64, s PenStyle) *QPen {
	g := newQPen(QPen_new6(brush.cPointer(), (double)(width), (int)(s)))
	g.isSubclass = true
	return g
}

// NewQPen7 constructs a new QPen object.
func NewQPen7(brush *QBrush, width float64, s PenStyle, c PenCapStyle) *QPen {
	g := newQPen(QPen_new7(brush.cPointer(), (double)(width), (int)(s), (int)(c)))
	g.isSubclass = true
	return g
}

// NewQPen8 constructs a new QPen object.
func NewQPen8(brush *QBrush, width float64, s PenStyle, c PenCapStyle, j PenJoinStyle) *QPen {
	g := newQPen(QPen_new8(brush.cPointer(), (double)(width), (int)(s), (int)(c), (int)(j)))
	g.isSubclass = true
	return g
}

func (this *QPen) OperatorAssign(pen *QPen) {
	QPen_OperatorAssign(this.h, pen.cPointer())
}

func (this *QPen) Swap(other *QPen) {
	QPen_Swap(this.h, other.cPointer())
}

func (this *QPen) OperatorAssignWithColor(color QColor) {
	QPen_OperatorAssignWithColor(this.h, color.cPointer())
}

func (this *QPen) OperatorAssignWithStyle(style PenStyle) {
	QPen_OperatorAssignWithStyle(this.h, (int)(style))
}

func (this *QPen) Style() PenStyle {
	return (PenStyle)(QPen_Style(this.h))
}

func (this *QPen) SetStyle(style PenStyle) {
	QPen_SetStyle(this.h, (int)(style))
}

func (this *QPen) DashPattern() []float64 {
	var _ma struct_miqt_array = QPen_DashPattern(this.h)
	_ret := make([]float64, int(_ma.len))
	_outCast := (*[0xffff]double)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (float64)(_outCast[i])
	}
	return _ret
}

func (this *QPen) SetDashPattern(pattern []float64) {
	pattern_CArray := (*[0xffff]double)(malloc(size_t(8 * len(pattern))))
	defer free(unsafe.Pointer(pattern_CArray))
	for i := range pattern {
		pattern_CArray[i] = (double)(pattern[i])
	}
	pattern_ma := struct_miqt_array{len: size_t(len(pattern)), data: unsafe.Pointer(pattern_CArray)}
	QPen_SetDashPattern(this.h, pattern_ma)
}

func (this *QPen) DashOffset() float64 {
	return (float64)(QPen_DashOffset(this.h))
}

func (this *QPen) SetDashOffset(doffset float64) {
	QPen_SetDashOffset(this.h, (double)(doffset))
}

func (this *QPen) MiterLimit() float64 {
	return (float64)(QPen_MiterLimit(this.h))
}

func (this *QPen) SetMiterLimit(limit float64) {
	QPen_SetMiterLimit(this.h, (double)(limit))
}

func (this *QPen) WidthF() float64 {
	return (float64)(QPen_WidthF(this.h))
}

func (this *QPen) SetWidthF(width float64) {
	QPen_SetWidthF(this.h, (double)(width))
}

func (this *QPen) Width() int {
	return (int)(QPen_Width(this.h))
}

func (this *QPen) SetWidth(width int) {
	QPen_SetWidth(this.h, (int)(width))
}

func (this *QPen) Color() *QColor {
	_goptr := newQColor(QPen_Color(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPen) SetColor(color *QColor) {
	QPen_SetColor(this.h, color.cPointer())
}

func (this *QPen) Brush() *QBrush {
	_goptr := newQBrush(QPen_Brush(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPen) SetBrush(brush *QBrush) {
	QPen_SetBrush(this.h, brush.cPointer())
}

func (this *QPen) IsSolid() bool {
	return (bool)(QPen_IsSolid(this.h))
}

func (this *QPen) CapStyle() PenCapStyle {
	return (PenCapStyle)(QPen_CapStyle(this.h))
}

func (this *QPen) SetCapStyle(pcs PenCapStyle) {
	QPen_SetCapStyle(this.h, (int)(pcs))
}

func (this *QPen) JoinStyle() PenJoinStyle {
	return (PenJoinStyle)(QPen_JoinStyle(this.h))
}

func (this *QPen) SetJoinStyle(pcs PenJoinStyle) {
	QPen_SetJoinStyle(this.h, (int)(pcs))
}

func (this *QPen) IsCosmetic() bool {
	return (bool)(QPen_IsCosmetic(this.h))
}

func (this *QPen) SetCosmetic(cosmetic bool) {
	QPen_SetCosmetic(this.h, (bool)(cosmetic))
}

func (this *QPen) OperatorEqual(p *QPen) bool {
	return (bool)(QPen_OperatorEqual(this.h, p.cPointer()))
}

func (this *QPen) OperatorNotEqual(p *QPen) bool {
	return (bool)(QPen_OperatorNotEqual(this.h, p.cPointer()))
}

func (this *QPen) IsDetached() bool {
	return (bool)(QPen_IsDetached(this.h))
}

func (this *QPen) DataPtr() *DataPtr {
	xxxxxxxxx
}
