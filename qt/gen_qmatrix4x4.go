package qt

import (
	"unsafe"
)

type QMatrix4x4__Flag int

const (
	QMatrix4x4__Identity    QMatrix4x4__Flag = 0
	QMatrix4x4__Translation QMatrix4x4__Flag = 1
	QMatrix4x4__Scale       QMatrix4x4__Flag = 2
	QMatrix4x4__Rotation2D  QMatrix4x4__Flag = 4
	QMatrix4x4__Rotation    QMatrix4x4__Flag = 8
	QMatrix4x4__Perspective QMatrix4x4__Flag = 16
	QMatrix4x4__General     QMatrix4x4__Flag = 31
)

type QMatrix4x4 struct {
	h          uintptr
	isSubclass bool
}

// NewQMatrix4x4 constructs a new QMatrix4x4 object.
func NewQMatrix4x4() *QMatrix4x4 {

	ret := newQMatrix4x4(QMatrix4x4_new())
	ret.isSubclass = true
	return ret
}

// NewQMatrix4x42 constructs a new QMatrix4x4 object.
func NewQMatrix4x42(param1 Initialization) *QMatrix4x4 {

	ret := newQMatrix4x4(QMatrix4x4_new2((int)(param1)))
	ret.isSubclass = true
	return ret
}

// NewQMatrix4x43 constructs a new QMatrix4x4 object.
func NewQMatrix4x43(values *float32) *QMatrix4x4 {

	ret := newQMatrix4x4(QMatrix4x4_new3((*float)(unsafe.Pointer(values))))
	ret.isSubclass = true
	return ret
}

// NewQMatrix4x44 constructs a new QMatrix4x4 object.
func NewQMatrix4x44(m11 float32, m12 float32, m13 float32, m14 float32, m21 float32, m22 float32, m23 float32, m24 float32, m31 float32, m32 float32, m33 float32, m34 float32, m41 float32, m42 float32, m43 float32, m44 float32) *QMatrix4x4 {

	ret := newQMatrix4x4(QMatrix4x4_new4((float)(m11), (float)(m12), (float)(m13), (float)(m14), (float)(m21), (float)(m22), (float)(m23), (float)(m24), (float)(m31), (float)(m32), (float)(m33), (float)(m34), (float)(m41), (float)(m42), (float)(m43), (float)(m44)))
	ret.isSubclass = true
	return ret
}

// NewQMatrix4x45 constructs a new QMatrix4x4 object.
func NewQMatrix4x45(values *float32, cols int, rows int) *QMatrix4x4 {

	ret := newQMatrix4x4(QMatrix4x4_new5((*float)(unsafe.Pointer(values)), (int)(cols), (int)(rows)))
	ret.isSubclass = true
	return ret
}

// NewQMatrix4x46 constructs a new QMatrix4x4 object.
func NewQMatrix4x46(transform *QTransform) *QMatrix4x4 {

	ret := newQMatrix4x4(QMatrix4x4_new6(transform.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQMatrix4x47 constructs a new QMatrix4x4 object.
func NewQMatrix4x47(param1 *QMatrix4x4) *QMatrix4x4 {

	ret := newQMatrix4x4(QMatrix4x4_new7(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QMatrix4x4) Column(index int) *QVector4D {
	_goptr := newQVector4D(QMatrix4x4_Column(this.h, (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMatrix4x4) SetColumn(index int, value *QVector4D) {
	QMatrix4x4_SetColumn(this.h, (int)(index), value.cPointer())
}

func (this *QMatrix4x4) Row(index int) *QVector4D {
	_goptr := newQVector4D(QMatrix4x4_Row(this.h, (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMatrix4x4) SetRow(index int, value *QVector4D) {
	QMatrix4x4_SetRow(this.h, (int)(index), value.cPointer())
}

func (this *QMatrix4x4) IsAffine() bool {
	return (bool)(QMatrix4x4_IsAffine(this.h))
}

func (this *QMatrix4x4) IsIdentity() bool {
	return (bool)(QMatrix4x4_IsIdentity(this.h))
}

func (this *QMatrix4x4) SetToIdentity() {
	QMatrix4x4_SetToIdentity(this.h)
}

func (this *QMatrix4x4) Fill(value float32) {
	QMatrix4x4_Fill(this.h, (float)(value))
}

func (this *QMatrix4x4) Determinant() float64 {
	return (float64)(QMatrix4x4_Determinant(this.h))
}

func (this *QMatrix4x4) Inverted() *QMatrix4x4 {
	_goptr := newQMatrix4x4(QMatrix4x4_Inverted(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMatrix4x4) Transposed() *QMatrix4x4 {
	_goptr := newQMatrix4x4(QMatrix4x4_Transposed(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMatrix4x4) OperatorPlusAssign(other *QMatrix4x4) *QMatrix4x4 {
	return newQMatrix4x4(QMatrix4x4_OperatorPlusAssign(this.h, other.cPointer()))
}

func (this *QMatrix4x4) OperatorMinusAssign(other *QMatrix4x4) *QMatrix4x4 {
	return newQMatrix4x4(QMatrix4x4_OperatorMinusAssign(this.h, other.cPointer()))
}

func (this *QMatrix4x4) OperatorMultiplyAssign(other *QMatrix4x4) *QMatrix4x4 {
	return newQMatrix4x4(QMatrix4x4_OperatorMultiplyAssign(this.h, other.cPointer()))
}

func (this *QMatrix4x4) OperatorMultiplyAssignWithFactor(factor float32) *QMatrix4x4 {
	return newQMatrix4x4(QMatrix4x4_OperatorMultiplyAssignWithFactor(this.h, (float)(factor)))
}

func (this *QMatrix4x4) OperatorDivideAssign(divisor float32) *QMatrix4x4 {
	return newQMatrix4x4(QMatrix4x4_OperatorDivideAssign(this.h, (float)(divisor)))
}

func (this *QMatrix4x4) OperatorEqual(other *QMatrix4x4) bool {
	return (bool)(QMatrix4x4_OperatorEqual(this.h, other.cPointer()))
}

func (this *QMatrix4x4) OperatorNotEqual(other *QMatrix4x4) bool {
	return (bool)(QMatrix4x4_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QMatrix4x4) Scale(vector *QVector3D) {
	QMatrix4x4_Scale(this.h, vector.cPointer())
}

func (this *QMatrix4x4) Translate(vector *QVector3D) {
	QMatrix4x4_Translate(this.h, vector.cPointer())
}

func (this *QMatrix4x4) Rotate(angle float32, vector *QVector3D) {
	QMatrix4x4_Rotate(this.h, (float)(angle), vector.cPointer())
}

func (this *QMatrix4x4) Scale2(x float32, y float32) {
	QMatrix4x4_Scale2(this.h, (float)(x), (float)(y))
}

func (this *QMatrix4x4) Scale3(x float32, y float32, z float32) {
	QMatrix4x4_Scale3(this.h, (float)(x), (float)(y), (float)(z))
}

func (this *QMatrix4x4) ScaleWithFactor(factor float32) {
	QMatrix4x4_ScaleWithFactor(this.h, (float)(factor))
}

func (this *QMatrix4x4) Translate2(x float32, y float32) {
	QMatrix4x4_Translate2(this.h, (float)(x), (float)(y))
}

func (this *QMatrix4x4) Translate3(x float32, y float32, z float32) {
	QMatrix4x4_Translate3(this.h, (float)(x), (float)(y), (float)(z))
}

func (this *QMatrix4x4) Rotate2(angle float32, x float32, y float32) {
	QMatrix4x4_Rotate2(this.h, (float)(angle), (float)(x), (float)(y))
}

func (this *QMatrix4x4) RotateWithQuaternion(quaternion *QQuaternion) {
	QMatrix4x4_RotateWithQuaternion(this.h, quaternion.cPointer())
}

func (this *QMatrix4x4) Ortho(rect *QRect) {
	QMatrix4x4_Ortho(this.h, rect.cPointer())
}

func (this *QMatrix4x4) OrthoWithRect(rect *QRectF) {
	QMatrix4x4_OrthoWithRect(this.h, rect.cPointer())
}

func (this *QMatrix4x4) Ortho2(left float32, right float32, bottom float32, top float32, nearPlane float32, farPlane float32) {
	QMatrix4x4_Ortho2(this.h, (float)(left), (float)(right), (float)(bottom), (float)(top), (float)(nearPlane), (float)(farPlane))
}

func (this *QMatrix4x4) Frustum(left float32, right float32, bottom float32, top float32, nearPlane float32, farPlane float32) {
	QMatrix4x4_Frustum(this.h, (float)(left), (float)(right), (float)(bottom), (float)(top), (float)(nearPlane), (float)(farPlane))
}

func (this *QMatrix4x4) Perspective(verticalAngle float32, aspectRatio float32, nearPlane float32, farPlane float32) {
	QMatrix4x4_Perspective(this.h, (float)(verticalAngle), (float)(aspectRatio), (float)(nearPlane), (float)(farPlane))
}

func (this *QMatrix4x4) LookAt(eye *QVector3D, center *QVector3D, up *QVector3D) {
	QMatrix4x4_LookAt(this.h, eye.cPointer(), center.cPointer(), up.cPointer())
}

func (this *QMatrix4x4) Viewport(rect *QRectF) {
	QMatrix4x4_Viewport(this.h, rect.cPointer())
}

func (this *QMatrix4x4) Viewport2(left float32, bottom float32, width float32, height float32) {
	QMatrix4x4_Viewport2(this.h, (float)(left), (float)(bottom), (float)(width), (float)(height))
}

func (this *QMatrix4x4) FlipCoordinates() {
	QMatrix4x4_FlipCoordinates(this.h)
}

func (this *QMatrix4x4) CopyDataTo(values *float32) {
	QMatrix4x4_CopyDataTo(this.h, (*float)(unsafe.Pointer(values)))
}

func (this *QMatrix4x4) ToTransform() *QTransform {
	_goptr := newQTransform(QMatrix4x4_ToTransform(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMatrix4x4) ToTransformWithDistanceToPlane(distanceToPlane float32) *QTransform {
	_goptr := newQTransform(QMatrix4x4_ToTransformWithDistanceToPlane(this.h, (float)(distanceToPlane)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMatrix4x4) Map(point *QPoint) *QPoint {
	_goptr := newQPoint(QMatrix4x4_Map(this.h, point.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMatrix4x4) MapWithPoint(point *QPointF) *QPointF {
	_goptr := newQPointF(QMatrix4x4_MapWithPoint(this.h, point.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMatrix4x4) Map2(point *QVector3D) *QVector3D {
	_goptr := newQVector3D(QMatrix4x4_Map2(this.h, point.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMatrix4x4) MapVector(vector *QVector3D) *QVector3D {
	_goptr := newQVector3D(QMatrix4x4_MapVector(this.h, vector.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMatrix4x4) Map3(point *QVector4D) *QVector4D {
	_goptr := newQVector4D(QMatrix4x4_Map3(this.h, point.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMatrix4x4) MapRect(rect *QRect) *QRect {
	_goptr := newQRect(QMatrix4x4_MapRect(this.h, rect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMatrix4x4) MapRectWithRect(rect *QRectF) *QRectF {
	_goptr := newQRectF(QMatrix4x4_MapRectWithRect(this.h, rect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMatrix4x4) Data() *float32 {
	return (*float32)(unsafe.Pointer(QMatrix4x4_Data(this.h)))
}

func (this *QMatrix4x4) Data2() *float32 {
	return (*float32)(unsafe.Pointer(QMatrix4x4_Data2(this.h)))
}

func (this *QMatrix4x4) ConstData() *float32 {
	return (*float32)(unsafe.Pointer(QMatrix4x4_ConstData(this.h)))
}

func (this *QMatrix4x4) Optimize() {
	QMatrix4x4_Optimize(this.h)
}

func (this *QMatrix4x4) ProjectedRotate(angle float32, x float32, y float32, z float32, distanceToPlane float32) {
	QMatrix4x4_ProjectedRotate(this.h, (float)(angle), (float)(x), (float)(y), (float)(z), (float)(distanceToPlane))
}

func (this *QMatrix4x4) ProjectedRotate2(angle float32, x float32, y float32, z float32) {
	QMatrix4x4_ProjectedRotate2(this.h, (float)(angle), (float)(x), (float)(y), (float)(z))
}

func (this *QMatrix4x4) Flags() Flags {
	xxxxxxxxx
}

func (this *QMatrix4x4) Inverted1(invertible *bool) *QMatrix4x4 {
	_goptr := newQMatrix4x4(QMatrix4x4_Inverted1(this.h, (*bool)(unsafe.Pointer(invertible))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMatrix4x4) Rotate4(angle float32, x float32, y float32, z float32) {
	QMatrix4x4_Rotate4(this.h, (float)(angle), (float)(x), (float)(y), (float)(z))
}

func (this *QMatrix4x4) Viewport5(left float32, bottom float32, width float32, height float32, nearPlane float32) {
	QMatrix4x4_Viewport5(this.h, (float)(left), (float)(bottom), (float)(width), (float)(height), (float)(nearPlane))
}

func (this *QMatrix4x4) Viewport6(left float32, bottom float32, width float32, height float32, nearPlane float32, farPlane float32) {
	QMatrix4x4_Viewport6(this.h, (float)(left), (float)(bottom), (float)(width), (float)(height), (float)(nearPlane), (float)(farPlane))
}
