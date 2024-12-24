package qt

import (
	"unsafe"
)

type QPoint struct {
	h          uintptr
	isSubclass bool
}

// NewQPoint constructs a new QPoint object.
func NewQPoint() *QPoint {
	ret := newQPoint(QPoint_new())
	ret.isSubclass = true
	return ret
}

// NewQPoint2 constructs a new QPoint object.
func NewQPoint2(xpos int, ypos int) *QPoint {
	ret := newQPoint(QPoint_new2((int)(xpos), (int)(ypos)))
	ret.isSubclass = true
	return ret
}

// NewQPoint3 constructs a new QPoint object.
func NewQPoint3(param1 *QPoint) *QPoint {
	ret := newQPoint(QPoint_new3(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QPoint) IsNull() bool {
	return (bool)(QPoint_IsNull(this.h))
}

func (this *QPoint) X() int {
	return (int)(QPoint_X(this.h))
}

func (this *QPoint) Y() int {
	return (int)(QPoint_Y(this.h))
}

func (this *QPoint) SetX(x int) {
	QPoint_SetX(this.h, (int)(x))
}

func (this *QPoint) SetY(y int) {
	QPoint_SetY(this.h, (int)(y))
}

func (this *QPoint) ManhattanLength() int {
	return (int)(QPoint_ManhattanLength(this.h))
}

func (this *QPoint) Transposed() *QPoint {
	_goptr := newQPoint(QPoint_Transposed(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPoint) OperatorPlusAssign(p *QPoint) *QPoint {
	return newQPoint(QPoint_OperatorPlusAssign(this.h, p.cPointer()))
}

func (this *QPoint) OperatorMinusAssign(p *QPoint) *QPoint {
	return newQPoint(QPoint_OperatorMinusAssign(this.h, p.cPointer()))
}

func (this *QPoint) OperatorMultiplyAssign(factor float32) *QPoint {
	return newQPoint(QPoint_OperatorMultiplyAssign(this.h, (float)(factor)))
}

func (this *QPoint) OperatorMultiplyAssignWithFactor(factor float64) *QPoint {
	return newQPoint(QPoint_OperatorMultiplyAssignWithFactor(this.h, (double)(factor)))
}

func (this *QPoint) OperatorMultiplyAssign2(factor int) *QPoint {
	return newQPoint(QPoint_OperatorMultiplyAssign2(this.h, (int)(factor)))
}

func (this *QPoint) OperatorDivideAssign(divisor float64) *QPoint {
	return newQPoint(QPoint_OperatorDivideAssign(this.h, (double)(divisor)))
}

func QPoint_DotProduct(p1 *QPoint, p2 *QPoint) int {
	return (int)(QPoint_DotProduct(p1.cPointer(), p2.cPointer()))
}

func (this *QPoint) ToPointF() *QPointF {
	_goptr := newQPointF(QPoint_ToPointF(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QPointF struct {
	h          uintptr
	isSubclass bool
}

// NewQPointF constructs a new QPointF object.
func NewQPointF() *QPointF {
	ret := newQPointF(QPointF_new())
	ret.isSubclass = true
	return ret
}

// NewQPointF2 constructs a new QPointF object.
func NewQPointF2(p *QPoint) *QPointF {
	ret := newQPointF(QPointF_new2(p.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPointF3 constructs a new QPointF object.
func NewQPointF3(xpos float64, ypos float64) *QPointF {
	ret := newQPointF(QPointF_new3((double)(xpos), (double)(ypos)))
	ret.isSubclass = true
	return ret
}

// NewQPointF4 constructs a new QPointF object.
func NewQPointF4(param1 *QPointF) *QPointF {
	ret := newQPointF(QPointF_new4(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QPointF) ManhattanLength() float64 {
	return (float64)(QPointF_ManhattanLength(this.h))
}

func (this *QPointF) IsNull() bool {
	return (bool)(QPointF_IsNull(this.h))
}

func (this *QPointF) X() float64 {
	return (float64)(QPointF_X(this.h))
}

func (this *QPointF) Y() float64 {
	return (float64)(QPointF_Y(this.h))
}

func (this *QPointF) SetX(x float64) {
	QPointF_SetX(this.h, (double)(x))
}

func (this *QPointF) SetY(y float64) {
	QPointF_SetY(this.h, (double)(y))
}

func (this *QPointF) Transposed() *QPointF {
	_goptr := newQPointF(QPointF_Transposed(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPointF) OperatorPlusAssign(p *QPointF) *QPointF {
	return newQPointF(QPointF_OperatorPlusAssign(this.h, p.cPointer()))
}

func (this *QPointF) OperatorMinusAssign(p *QPointF) *QPointF {
	return newQPointF(QPointF_OperatorMinusAssign(this.h, p.cPointer()))
}

func (this *QPointF) OperatorMultiplyAssign(c float64) *QPointF {
	return newQPointF(QPointF_OperatorMultiplyAssign(this.h, (double)(c)))
}

func (this *QPointF) OperatorDivideAssign(c float64) *QPointF {
	return newQPointF(QPointF_OperatorDivideAssign(this.h, (double)(c)))
}

func QPointF_DotProduct(p1 *QPointF, p2 *QPointF) float64 {
	return (float64)(QPointF_DotProduct(p1.cPointer(), p2.cPointer()))
}

func (this *QPointF) ToPoint() *QPoint {
	_goptr := newQPoint(QPointF_ToPoint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
