package qt

import (
	"unsafe"
)

type QTransform__TransformationType int

const (
	QTransform__TxNone      QTransform__TransformationType = 0
	QTransform__TxTranslate QTransform__TransformationType = 1
	QTransform__TxScale     QTransform__TransformationType = 2
	QTransform__TxRotate    QTransform__TransformationType = 4
	QTransform__TxShear     QTransform__TransformationType = 8
	QTransform__TxProject   QTransform__TransformationType = 16
)

type QTransform struct {
	h          uintptr
	isSubclass bool
}

// NewQTransform constructs a new QTransform object.
func NewQTransform(param1 Initialization) *QTransform {

	ret := newQTransform(QTransform_new((int)(param1)))
	ret.isSubclass = true
	return ret
}

// NewQTransform2 constructs a new QTransform object.
func NewQTransform2() *QTransform {

	ret := newQTransform(QTransform_new2())
	ret.isSubclass = true
	return ret
}

// NewQTransform3 constructs a new QTransform object.
func NewQTransform3(h11 float64, h12 float64, h13 float64, h21 float64, h22 float64, h23 float64, h31 float64, h32 float64, h33 float64) *QTransform {

	ret := newQTransform(QTransform_new3((double)(h11), (double)(h12), (double)(h13), (double)(h21), (double)(h22), (double)(h23), (double)(h31), (double)(h32), (double)(h33)))
	ret.isSubclass = true
	return ret
}

// NewQTransform4 constructs a new QTransform object.
func NewQTransform4(h11 float64, h12 float64, h21 float64, h22 float64, dx float64, dy float64) *QTransform {

	ret := newQTransform(QTransform_new4((double)(h11), (double)(h12), (double)(h21), (double)(h22), (double)(dx), (double)(dy)))
	ret.isSubclass = true
	return ret
}

// NewQTransform5 constructs a new QTransform object.
func NewQTransform5(other *QTransform) *QTransform {

	ret := newQTransform(QTransform_new5(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTransform) OperatorAssign(param1 *QTransform) {
	QTransform_OperatorAssign(this.h, param1.cPointer())
}

func (this *QTransform) IsAffine() bool {
	return (bool)(QTransform_IsAffine(this.h))
}

func (this *QTransform) IsIdentity() bool {
	return (bool)(QTransform_IsIdentity(this.h))
}

func (this *QTransform) IsInvertible() bool {
	return (bool)(QTransform_IsInvertible(this.h))
}

func (this *QTransform) IsScaling() bool {
	return (bool)(QTransform_IsScaling(this.h))
}

func (this *QTransform) IsRotating() bool {
	return (bool)(QTransform_IsRotating(this.h))
}

func (this *QTransform) IsTranslating() bool {
	return (bool)(QTransform_IsTranslating(this.h))
}

func (this *QTransform) Type() TransformationType {
	xxxxxxxxx
}

func (this *QTransform) Determinant() float64 {
	return (float64)(QTransform_Determinant(this.h))
}

func (this *QTransform) M11() float64 {
	return (float64)(QTransform_M11(this.h))
}

func (this *QTransform) M12() float64 {
	return (float64)(QTransform_M12(this.h))
}

func (this *QTransform) M13() float64 {
	return (float64)(QTransform_M13(this.h))
}

func (this *QTransform) M21() float64 {
	return (float64)(QTransform_M21(this.h))
}

func (this *QTransform) M22() float64 {
	return (float64)(QTransform_M22(this.h))
}

func (this *QTransform) M23() float64 {
	return (float64)(QTransform_M23(this.h))
}

func (this *QTransform) M31() float64 {
	return (float64)(QTransform_M31(this.h))
}

func (this *QTransform) M32() float64 {
	return (float64)(QTransform_M32(this.h))
}

func (this *QTransform) M33() float64 {
	return (float64)(QTransform_M33(this.h))
}

func (this *QTransform) Dx() float64 {
	return (float64)(QTransform_Dx(this.h))
}

func (this *QTransform) Dy() float64 {
	return (float64)(QTransform_Dy(this.h))
}

func (this *QTransform) SetMatrix(m11 float64, m12 float64, m13 float64, m21 float64, m22 float64, m23 float64, m31 float64, m32 float64, m33 float64) {
	QTransform_SetMatrix(this.h, (double)(m11), (double)(m12), (double)(m13), (double)(m21), (double)(m22), (double)(m23), (double)(m31), (double)(m32), (double)(m33))
}

func (this *QTransform) Inverted() *QTransform {
	_goptr := newQTransform(QTransform_Inverted(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransform) Adjoint() *QTransform {
	_goptr := newQTransform(QTransform_Adjoint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransform) Transposed() *QTransform {
	_goptr := newQTransform(QTransform_Transposed(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransform) Translate(dx float64, dy float64) *QTransform {
	return newQTransform(QTransform_Translate(this.h, (double)(dx), (double)(dy)))
}

func (this *QTransform) Scale(sx float64, sy float64) *QTransform {
	return newQTransform(QTransform_Scale(this.h, (double)(sx), (double)(sy)))
}

func (this *QTransform) Shear(sh float64, sv float64) *QTransform {
	return newQTransform(QTransform_Shear(this.h, (double)(sh), (double)(sv)))
}

func (this *QTransform) Rotate(a float64, axis Axis, distanceToPlane float64) *QTransform {
	return newQTransform(QTransform_Rotate(this.h, (double)(a), (int)(axis), (double)(distanceToPlane)))
}

func (this *QTransform) RotateWithQreal(a float64) *QTransform {
	return newQTransform(QTransform_RotateWithQreal(this.h, (double)(a)))
}

func (this *QTransform) RotateRadians(a float64, axis Axis, distanceToPlane float64) *QTransform {
	return newQTransform(QTransform_RotateRadians(this.h, (double)(a), (int)(axis), (double)(distanceToPlane)))
}

func (this *QTransform) RotateRadiansWithQreal(a float64) *QTransform {
	return newQTransform(QTransform_RotateRadiansWithQreal(this.h, (double)(a)))
}

func (this *QTransform) OperatorEqual(param1 *QTransform) bool {
	return (bool)(QTransform_OperatorEqual(this.h, param1.cPointer()))
}

func (this *QTransform) OperatorNotEqual(param1 *QTransform) bool {
	return (bool)(QTransform_OperatorNotEqual(this.h, param1.cPointer()))
}

func (this *QTransform) OperatorMultiplyAssign(param1 *QTransform) *QTransform {
	return newQTransform(QTransform_OperatorMultiplyAssign(this.h, param1.cPointer()))
}

func (this *QTransform) OperatorMultiply(o *QTransform) *QTransform {
	_goptr := newQTransform(QTransform_OperatorMultiply(this.h, o.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransform) Reset() {
	QTransform_Reset(this.h)
}

func (this *QTransform) Map(p *QPoint) *QPoint {
	_goptr := newQPoint(QTransform_Map(this.h, p.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransform) MapWithQPointF(p *QPointF) *QPointF {
	_goptr := newQPointF(QTransform_MapWithQPointF(this.h, p.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransform) MapWithQLine(l *QLine) *QLine {
	_goptr := newQLine(QTransform_MapWithQLine(this.h, l.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransform) MapWithQLineF(l *QLineF) *QLineF {
	_goptr := newQLineF(QTransform_MapWithQLineF(this.h, l.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransform) MapWithQRegion(r *QRegion) *QRegion {
	_goptr := newQRegion(QTransform_MapWithQRegion(this.h, r.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransform) MapWithQPainterPath(p *QPainterPath) *QPainterPath {
	_goptr := newQPainterPath(QTransform_MapWithQPainterPath(this.h, p.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransform) MapRect(param1 *QRect) *QRect {
	_goptr := newQRect(QTransform_MapRect(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransform) MapRectWithQRectF(param1 *QRectF) *QRectF {
	_goptr := newQRectF(QTransform_MapRectWithQRectF(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransform) Map2(x int, y int, tx *int, ty *int) {
	QTransform_Map2(this.h, (int)(x), (int)(y), (*int)(unsafe.Pointer(tx)), (*int)(unsafe.Pointer(ty)))
}

func (this *QTransform) Map3(x float64, y float64, tx *float64, ty *float64) {
	QTransform_Map3(this.h, (double)(x), (double)(y), (*double)(unsafe.Pointer(tx)), (*double)(unsafe.Pointer(ty)))
}

func (this *QTransform) OperatorMultiplyAssignWithDiv(div float64) *QTransform {
	return newQTransform(QTransform_OperatorMultiplyAssignWithDiv(this.h, (double)(div)))
}

func (this *QTransform) OperatorDivideAssign(div float64) *QTransform {
	return newQTransform(QTransform_OperatorDivideAssign(this.h, (double)(div)))
}

func (this *QTransform) OperatorPlusAssign(div float64) *QTransform {
	return newQTransform(QTransform_OperatorPlusAssign(this.h, (double)(div)))
}

func (this *QTransform) OperatorMinusAssign(div float64) *QTransform {
	return newQTransform(QTransform_OperatorMinusAssign(this.h, (double)(div)))
}

func QTransform_FromTranslate(dx float64, dy float64) *QTransform {
	_goptr := newQTransform(QTransform_FromTranslate((double)(dx), (double)(dy)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QTransform_FromScale(dx float64, dy float64) *QTransform {
	_goptr := newQTransform(QTransform_FromScale((double)(dx), (double)(dy)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransform) AsAffineMatrix() Affine {
	xxxxxxxxx
}

func (this *QTransform) Inverted1(invertible *bool) *QTransform {
	_goptr := newQTransform(QTransform_Inverted1(this.h, (*bool)(unsafe.Pointer(invertible))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTransform) Rotate2(a float64, axis Axis) *QTransform {
	return newQTransform(QTransform_Rotate2(this.h, (double)(a), (int)(axis)))
}

func (this *QTransform) RotateRadians2(a float64, axis Axis) *QTransform {
	return newQTransform(QTransform_RotateRadians2(this.h, (double)(a), (int)(axis)))
}
