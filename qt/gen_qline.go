package qt

import (
	"unsafe"
)

type QLineF__IntersectionType int

const (
	QLineF__NoIntersection        QLineF__IntersectionType = 0
	QLineF__BoundedIntersection   QLineF__IntersectionType = 1
	QLineF__UnboundedIntersection QLineF__IntersectionType = 2
)

type QLine struct {
	h          uintptr
	isSubclass bool
}

// NewQLine constructs a new QLine object.
func NewQLine() *QLine {
	ret := newQLine(QLine_new())
	ret.isSubclass = true
	return ret
}

// NewQLine2 constructs a new QLine object.
func NewQLine2(pt1 *QPoint, pt2 *QPoint) *QLine {
	ret := newQLine(QLine_new2(pt1.cPointer(), pt2.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQLine3 constructs a new QLine object.
func NewQLine3(x1 int, y1 int, x2 int, y2 int) *QLine {
	ret := newQLine(QLine_new3((int)(x1), (int)(y1), (int)(x2), (int)(y2)))
	ret.isSubclass = true
	return ret
}

// NewQLine4 constructs a new QLine object.
func NewQLine4(param1 *QLine) *QLine {
	ret := newQLine(QLine_new4(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QLine) IsNull() bool {
	return (bool)(QLine_IsNull(this.h))
}

func (this *QLine) P1() *QPoint {
	_goptr := newQPoint(QLine_P1(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLine) P2() *QPoint {
	_goptr := newQPoint(QLine_P2(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLine) X1() int {
	return (int)(QLine_X1(this.h))
}

func (this *QLine) Y1() int {
	return (int)(QLine_Y1(this.h))
}

func (this *QLine) X2() int {
	return (int)(QLine_X2(this.h))
}

func (this *QLine) Y2() int {
	return (int)(QLine_Y2(this.h))
}

func (this *QLine) Dx() int {
	return (int)(QLine_Dx(this.h))
}

func (this *QLine) Dy() int {
	return (int)(QLine_Dy(this.h))
}

func (this *QLine) Translate(p *QPoint) {
	QLine_Translate(this.h, p.cPointer())
}

func (this *QLine) Translate2(dx int, dy int) {
	QLine_Translate2(this.h, (int)(dx), (int)(dy))
}

func (this *QLine) Translated(p *QPoint) *QLine {
	_goptr := newQLine(QLine_Translated(this.h, p.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLine) Translated2(dx int, dy int) *QLine {
	_goptr := newQLine(QLine_Translated2(this.h, (int)(dx), (int)(dy)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLine) Center() *QPoint {
	_goptr := newQPoint(QLine_Center(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLine) SetP1(p1 *QPoint) {
	QLine_SetP1(this.h, p1.cPointer())
}

func (this *QLine) SetP2(p2 *QPoint) {
	QLine_SetP2(this.h, p2.cPointer())
}

func (this *QLine) SetPoints(p1 *QPoint, p2 *QPoint) {
	QLine_SetPoints(this.h, p1.cPointer(), p2.cPointer())
}

func (this *QLine) SetLine(x1 int, y1 int, x2 int, y2 int) {
	QLine_SetLine(this.h, (int)(x1), (int)(y1), (int)(x2), (int)(y2))
}

func (this *QLine) ToLineF() *QLineF {
	_goptr := newQLineF(QLine_ToLineF(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QLineF struct {
	h          uintptr
	isSubclass bool
}

// NewQLineF constructs a new QLineF object.
func NewQLineF() *QLineF {
	ret := newQLineF(QLineF_new())
	ret.isSubclass = true
	return ret
}

// NewQLineF2 constructs a new QLineF object.
func NewQLineF2(pt1 *QPointF, pt2 *QPointF) *QLineF {
	ret := newQLineF(QLineF_new2(pt1.cPointer(), pt2.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQLineF3 constructs a new QLineF object.
func NewQLineF3(x1 float64, y1 float64, x2 float64, y2 float64) *QLineF {
	ret := newQLineF(QLineF_new3((double)(x1), (double)(y1), (double)(x2), (double)(y2)))
	ret.isSubclass = true
	return ret
}

// NewQLineF4 constructs a new QLineF object.
func NewQLineF4(line *QLine) *QLineF {
	ret := newQLineF(QLineF_new4(line.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQLineF5 constructs a new QLineF object.
func NewQLineF5(param1 *QLineF) *QLineF {
	ret := newQLineF(QLineF_new5(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func QLineF_FromPolar(length float64, angle float64) *QLineF {
	_goptr := newQLineF(QLineF_FromPolar((double)(length), (double)(angle)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLineF) IsNull() bool {
	return (bool)(QLineF_IsNull(this.h))
}

func (this *QLineF) P1() *QPointF {
	_goptr := newQPointF(QLineF_P1(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLineF) P2() *QPointF {
	_goptr := newQPointF(QLineF_P2(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLineF) X1() float64 {
	return (float64)(QLineF_X1(this.h))
}

func (this *QLineF) Y1() float64 {
	return (float64)(QLineF_Y1(this.h))
}

func (this *QLineF) X2() float64 {
	return (float64)(QLineF_X2(this.h))
}

func (this *QLineF) Y2() float64 {
	return (float64)(QLineF_Y2(this.h))
}

func (this *QLineF) Dx() float64 {
	return (float64)(QLineF_Dx(this.h))
}

func (this *QLineF) Dy() float64 {
	return (float64)(QLineF_Dy(this.h))
}

func (this *QLineF) Length() float64 {
	return (float64)(QLineF_Length(this.h))
}

func (this *QLineF) SetLength(lenVal float64) {
	QLineF_SetLength(this.h, (double)(lenVal))
}

func (this *QLineF) Angle() float64 {
	return (float64)(QLineF_Angle(this.h))
}

func (this *QLineF) SetAngle(angle float64) {
	QLineF_SetAngle(this.h, (double)(angle))
}

func (this *QLineF) AngleTo(l *QLineF) float64 {
	return (float64)(QLineF_AngleTo(this.h, l.cPointer()))
}

func (this *QLineF) UnitVector() *QLineF {
	_goptr := newQLineF(QLineF_UnitVector(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLineF) NormalVector() *QLineF {
	_goptr := newQLineF(QLineF_NormalVector(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLineF) Intersects(l *QLineF) IntersectionType {
	xxxxxxxxx
}

func (this *QLineF) PointAt(t float64) *QPointF {
	_goptr := newQPointF(QLineF_PointAt(this.h, (double)(t)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLineF) Translate(p *QPointF) {
	QLineF_Translate(this.h, p.cPointer())
}

func (this *QLineF) Translate2(dx float64, dy float64) {
	QLineF_Translate2(this.h, (double)(dx), (double)(dy))
}

func (this *QLineF) Translated(p *QPointF) *QLineF {
	_goptr := newQLineF(QLineF_Translated(this.h, p.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLineF) Translated2(dx float64, dy float64) *QLineF {
	_goptr := newQLineF(QLineF_Translated2(this.h, (double)(dx), (double)(dy)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLineF) Center() *QPointF {
	_goptr := newQPointF(QLineF_Center(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLineF) SetP1(p1 *QPointF) {
	QLineF_SetP1(this.h, p1.cPointer())
}

func (this *QLineF) SetP2(p2 *QPointF) {
	QLineF_SetP2(this.h, p2.cPointer())
}

func (this *QLineF) SetPoints(p1 *QPointF, p2 *QPointF) {
	QLineF_SetPoints(this.h, p1.cPointer(), p2.cPointer())
}

func (this *QLineF) SetLine(x1 float64, y1 float64, x2 float64, y2 float64) {
	QLineF_SetLine(this.h, (double)(x1), (double)(y1), (double)(x2), (double)(y2))
}

func (this *QLineF) ToLine() *QLine {
	_goptr := newQLine(QLineF_ToLine(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLineF) Intersects2(l *QLineF, intersectionPoint *QPointF) IntersectionType {
	xxxxxxxxx
}
