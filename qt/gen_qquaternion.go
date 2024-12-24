package qt

import (
	"unsafe"
)

type QQuaternion struct {
	h          uintptr
	isSubclass bool
}

// NewQQuaternion constructs a new QQuaternion object.
func NewQQuaternion() *QQuaternion {

	ret := newQQuaternion(QQuaternion_new())
	ret.isSubclass = true
	return ret
}

// NewQQuaternion2 constructs a new QQuaternion object.
func NewQQuaternion2(param1 Initialization) *QQuaternion {

	ret := newQQuaternion(QQuaternion_new2((int)(param1)))
	ret.isSubclass = true
	return ret
}

// NewQQuaternion3 constructs a new QQuaternion object.
func NewQQuaternion3(scalar float32, xpos float32, ypos float32, zpos float32) *QQuaternion {

	ret := newQQuaternion(QQuaternion_new3((float)(scalar), (float)(xpos), (float)(ypos), (float)(zpos)))
	ret.isSubclass = true
	return ret
}

// NewQQuaternion4 constructs a new QQuaternion object.
func NewQQuaternion4(scalar float32, vector *QVector3D) *QQuaternion {

	ret := newQQuaternion(QQuaternion_new4((float)(scalar), vector.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQQuaternion5 constructs a new QQuaternion object.
func NewQQuaternion5(vector *QVector4D) *QQuaternion {

	ret := newQQuaternion(QQuaternion_new5(vector.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQQuaternion6 constructs a new QQuaternion object.
func NewQQuaternion6(param1 *QQuaternion) *QQuaternion {

	ret := newQQuaternion(QQuaternion_new6(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QQuaternion) IsNull() bool {
	return (bool)(QQuaternion_IsNull(this.h))
}

func (this *QQuaternion) IsIdentity() bool {
	return (bool)(QQuaternion_IsIdentity(this.h))
}

func (this *QQuaternion) Vector() *QVector3D {
	_goptr := newQVector3D(QQuaternion_Vector(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QQuaternion) SetVector(vector *QVector3D) {
	QQuaternion_SetVector(this.h, vector.cPointer())
}

func (this *QQuaternion) SetVector2(x float32, y float32, z float32) {
	QQuaternion_SetVector2(this.h, (float)(x), (float)(y), (float)(z))
}

func (this *QQuaternion) X() float32 {
	return (float32)(QQuaternion_X(this.h))
}

func (this *QQuaternion) Y() float32 {
	return (float32)(QQuaternion_Y(this.h))
}

func (this *QQuaternion) Z() float32 {
	return (float32)(QQuaternion_Z(this.h))
}

func (this *QQuaternion) Scalar() float32 {
	return (float32)(QQuaternion_Scalar(this.h))
}

func (this *QQuaternion) SetX(x float32) {
	QQuaternion_SetX(this.h, (float)(x))
}

func (this *QQuaternion) SetY(y float32) {
	QQuaternion_SetY(this.h, (float)(y))
}

func (this *QQuaternion) SetZ(z float32) {
	QQuaternion_SetZ(this.h, (float)(z))
}

func (this *QQuaternion) SetScalar(scalar float32) {
	QQuaternion_SetScalar(this.h, (float)(scalar))
}

func QQuaternion_DotProduct(q1 *QQuaternion, q2 *QQuaternion) float32 {
	return (float32)(QQuaternion_DotProduct(q1.cPointer(), q2.cPointer()))
}

func (this *QQuaternion) Length() float32 {
	return (float32)(QQuaternion_Length(this.h))
}

func (this *QQuaternion) LengthSquared() float32 {
	return (float32)(QQuaternion_LengthSquared(this.h))
}

func (this *QQuaternion) Normalized() *QQuaternion {
	_goptr := newQQuaternion(QQuaternion_Normalized(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QQuaternion) Normalize() {
	QQuaternion_Normalize(this.h)
}

func (this *QQuaternion) Inverted() *QQuaternion {
	_goptr := newQQuaternion(QQuaternion_Inverted(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QQuaternion) Conjugated() *QQuaternion {
	_goptr := newQQuaternion(QQuaternion_Conjugated(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QQuaternion) RotatedVector(vector *QVector3D) *QVector3D {
	_goptr := newQVector3D(QQuaternion_RotatedVector(this.h, vector.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QQuaternion) OperatorPlusAssign(quaternion *QQuaternion) *QQuaternion {
	return newQQuaternion(QQuaternion_OperatorPlusAssign(this.h, quaternion.cPointer()))
}

func (this *QQuaternion) OperatorMinusAssign(quaternion *QQuaternion) *QQuaternion {
	return newQQuaternion(QQuaternion_OperatorMinusAssign(this.h, quaternion.cPointer()))
}

func (this *QQuaternion) OperatorMultiplyAssign(factor float32) *QQuaternion {
	return newQQuaternion(QQuaternion_OperatorMultiplyAssign(this.h, (float)(factor)))
}

func (this *QQuaternion) OperatorMultiplyAssignWithQuaternion(quaternion *QQuaternion) *QQuaternion {
	return newQQuaternion(QQuaternion_OperatorMultiplyAssignWithQuaternion(this.h, quaternion.cPointer()))
}

func (this *QQuaternion) OperatorDivideAssign(divisor float32) *QQuaternion {
	return newQQuaternion(QQuaternion_OperatorDivideAssign(this.h, (float)(divisor)))
}

func (this *QQuaternion) ToVector4D() *QVector4D {
	_goptr := newQVector4D(QQuaternion_ToVector4D(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QQuaternion) GetAxisAndAngle(axis *QVector3D, angle *float32) {
	QQuaternion_GetAxisAndAngle(this.h, axis.cPointer(), (*float)(unsafe.Pointer(angle)))
}

func QQuaternion_FromAxisAndAngle(axis *QVector3D, angle float32) *QQuaternion {
	_goptr := newQQuaternion(QQuaternion_FromAxisAndAngle(axis.cPointer(), (float)(angle)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QQuaternion) GetAxisAndAngle2(x *float32, y *float32, z *float32, angle *float32) {
	QQuaternion_GetAxisAndAngle2(this.h, (*float)(unsafe.Pointer(x)), (*float)(unsafe.Pointer(y)), (*float)(unsafe.Pointer(z)), (*float)(unsafe.Pointer(angle)))
}

func QQuaternion_FromAxisAndAngle2(x float32, y float32, z float32, angle float32) *QQuaternion {
	_goptr := newQQuaternion(QQuaternion_FromAxisAndAngle2((float)(x), (float)(y), (float)(z), (float)(angle)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QQuaternion) ToEulerAngles() *QVector3D {
	_goptr := newQVector3D(QQuaternion_ToEulerAngles(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QQuaternion_FromEulerAngles(eulerAngles *QVector3D) *QQuaternion {
	_goptr := newQQuaternion(QQuaternion_FromEulerAngles(eulerAngles.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QQuaternion) GetEulerAngles(pitch *float32, yaw *float32, roll *float32) {
	QQuaternion_GetEulerAngles(this.h, (*float)(unsafe.Pointer(pitch)), (*float)(unsafe.Pointer(yaw)), (*float)(unsafe.Pointer(roll)))
}

func QQuaternion_FromEulerAngles2(pitch float32, yaw float32, roll float32) *QQuaternion {
	_goptr := newQQuaternion(QQuaternion_FromEulerAngles2((float)(pitch), (float)(yaw), (float)(roll)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QQuaternion) GetAxes(xAxis *QVector3D, yAxis *QVector3D, zAxis *QVector3D) {
	QQuaternion_GetAxes(this.h, xAxis.cPointer(), yAxis.cPointer(), zAxis.cPointer())
}

func QQuaternion_FromAxes(xAxis *QVector3D, yAxis *QVector3D, zAxis *QVector3D) *QQuaternion {
	_goptr := newQQuaternion(QQuaternion_FromAxes(xAxis.cPointer(), yAxis.cPointer(), zAxis.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QQuaternion_FromDirection(direction *QVector3D, up *QVector3D) *QQuaternion {
	_goptr := newQQuaternion(QQuaternion_FromDirection(direction.cPointer(), up.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QQuaternion_RotationTo(from *QVector3D, to *QVector3D) *QQuaternion {
	_goptr := newQQuaternion(QQuaternion_RotationTo(from.cPointer(), to.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QQuaternion_Slerp(q1 *QQuaternion, q2 *QQuaternion, t float32) *QQuaternion {
	_goptr := newQQuaternion(QQuaternion_Slerp(q1.cPointer(), q2.cPointer(), (float)(t)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QQuaternion_Nlerp(q1 *QQuaternion, q2 *QQuaternion, t float32) *QQuaternion {
	_goptr := newQQuaternion(QQuaternion_Nlerp(q1.cPointer(), q2.cPointer(), (float)(t)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
