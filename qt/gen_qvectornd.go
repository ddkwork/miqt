package qt

import (
	"unsafe"
)

type QVector2D struct {
	h          uintptr
	isSubclass bool
}

// NewQVector2D constructs a new QVector2D object.
func NewQVector2D() *QVector2D {

	ret := newQVector2D(QVector2D_new())
	ret.isSubclass = true
	return ret
}

// NewQVector2D2 constructs a new QVector2D object.
func NewQVector2D2(param1 Initialization) *QVector2D {

	ret := newQVector2D(QVector2D_new2((int)(param1)))
	ret.isSubclass = true
	return ret
}

// NewQVector2D3 constructs a new QVector2D object.
func NewQVector2D3(xpos float32, ypos float32) *QVector2D {

	ret := newQVector2D(QVector2D_new3((float)(xpos), (float)(ypos)))
	ret.isSubclass = true
	return ret
}

// NewQVector2D4 constructs a new QVector2D object.
func NewQVector2D4(point QPoint) *QVector2D {

	ret := newQVector2D(QVector2D_new4(point.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVector2D5 constructs a new QVector2D object.
func NewQVector2D5(point QPointF) *QVector2D {

	ret := newQVector2D(QVector2D_new5(point.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVector2D6 constructs a new QVector2D object.
func NewQVector2D6(vector QVector3D) *QVector2D {

	ret := newQVector2D(QVector2D_new6(vector.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVector2D7 constructs a new QVector2D object.
func NewQVector2D7(vector QVector4D) *QVector2D {

	ret := newQVector2D(QVector2D_new7(vector.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVector2D8 constructs a new QVector2D object.
func NewQVector2D8(param1 *QVector2D) *QVector2D {

	ret := newQVector2D(QVector2D_new8(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QVector2D) IsNull() bool {
	return (bool)(QVector2D_IsNull(this.h))
}

func (this *QVector2D) X() float32 {
	return (float32)(QVector2D_X(this.h))
}

func (this *QVector2D) Y() float32 {
	return (float32)(QVector2D_Y(this.h))
}

func (this *QVector2D) SetX(x float32) {
	QVector2D_SetX(this.h, (float)(x))
}

func (this *QVector2D) SetY(y float32) {
	QVector2D_SetY(this.h, (float)(y))
}

func (this *QVector2D) OperatorSubscript(i int) float32 {
	return (float32)(QVector2D_OperatorSubscript(this.h, (int)(i)))
}

func (this *QVector2D) Length() float32 {
	return (float32)(QVector2D_Length(this.h))
}

func (this *QVector2D) LengthSquared() float32 {
	return (float32)(QVector2D_LengthSquared(this.h))
}

func (this *QVector2D) Normalized() *QVector2D {
	_goptr := newQVector2D(QVector2D_Normalized(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVector2D) Normalize() {
	QVector2D_Normalize(this.h)
}

func (this *QVector2D) DistanceToPoint(point QVector2D) float32 {
	return (float32)(QVector2D_DistanceToPoint(this.h, point.cPointer()))
}

func (this *QVector2D) DistanceToLine(point QVector2D, direction QVector2D) float32 {
	return (float32)(QVector2D_DistanceToLine(this.h, point.cPointer(), direction.cPointer()))
}

func (this *QVector2D) OperatorPlusAssign(vector QVector2D) *QVector2D {
	return newQVector2D(QVector2D_OperatorPlusAssign(this.h, vector.cPointer()))
}

func (this *QVector2D) OperatorMinusAssign(vector QVector2D) *QVector2D {
	return newQVector2D(QVector2D_OperatorMinusAssign(this.h, vector.cPointer()))
}

func (this *QVector2D) OperatorMultiplyAssign(factor float32) *QVector2D {
	return newQVector2D(QVector2D_OperatorMultiplyAssign(this.h, (float)(factor)))
}

func (this *QVector2D) OperatorMultiplyAssignWithVector(vector QVector2D) *QVector2D {
	return newQVector2D(QVector2D_OperatorMultiplyAssignWithVector(this.h, vector.cPointer()))
}

func (this *QVector2D) OperatorDivideAssign(divisor float32) *QVector2D {
	return newQVector2D(QVector2D_OperatorDivideAssign(this.h, (float)(divisor)))
}

func (this *QVector2D) OperatorDivideAssignWithVector(vector QVector2D) *QVector2D {
	return newQVector2D(QVector2D_OperatorDivideAssignWithVector(this.h, vector.cPointer()))
}

func QVector2D_DotProduct(v1 QVector2D, v2 QVector2D) float32 {
	return (float32)(QVector2D_DotProduct(v1.cPointer(), v2.cPointer()))
}

func (this *QVector2D) ToVector3D() *QVector3D {
	_goptr := newQVector3D(QVector2D_ToVector3D(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVector2D) ToVector4D() *QVector4D {
	_goptr := newQVector4D(QVector2D_ToVector4D(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVector2D) ToPoint() *QPoint {
	_goptr := newQPoint(QVector2D_ToPoint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVector2D) ToPointF() *QPointF {
	_goptr := newQPointF(QVector2D_ToPointF(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QVector3D struct {
	h          uintptr
	isSubclass bool
}

// NewQVector3D constructs a new QVector3D object.
func NewQVector3D() *QVector3D {

	ret := newQVector3D(QVector3D_new())
	ret.isSubclass = true
	return ret
}

// NewQVector3D2 constructs a new QVector3D object.
func NewQVector3D2(param1 Initialization) *QVector3D {

	ret := newQVector3D(QVector3D_new2((int)(param1)))
	ret.isSubclass = true
	return ret
}

// NewQVector3D3 constructs a new QVector3D object.
func NewQVector3D3(xpos float32, ypos float32, zpos float32) *QVector3D {

	ret := newQVector3D(QVector3D_new3((float)(xpos), (float)(ypos), (float)(zpos)))
	ret.isSubclass = true
	return ret
}

// NewQVector3D4 constructs a new QVector3D object.
func NewQVector3D4(point QPoint) *QVector3D {

	ret := newQVector3D(QVector3D_new4(point.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVector3D5 constructs a new QVector3D object.
func NewQVector3D5(point QPointF) *QVector3D {

	ret := newQVector3D(QVector3D_new5(point.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVector3D6 constructs a new QVector3D object.
func NewQVector3D6(vector QVector2D) *QVector3D {

	ret := newQVector3D(QVector3D_new6(vector.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVector3D7 constructs a new QVector3D object.
func NewQVector3D7(vector QVector2D, zpos float32) *QVector3D {

	ret := newQVector3D(QVector3D_new7(vector.cPointer(), (float)(zpos)))
	ret.isSubclass = true
	return ret
}

// NewQVector3D8 constructs a new QVector3D object.
func NewQVector3D8(vector QVector4D) *QVector3D {

	ret := newQVector3D(QVector3D_new8(vector.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVector3D9 constructs a new QVector3D object.
func NewQVector3D9(param1 *QVector3D) *QVector3D {

	ret := newQVector3D(QVector3D_new9(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QVector3D) IsNull() bool {
	return (bool)(QVector3D_IsNull(this.h))
}

func (this *QVector3D) X() float32 {
	return (float32)(QVector3D_X(this.h))
}

func (this *QVector3D) Y() float32 {
	return (float32)(QVector3D_Y(this.h))
}

func (this *QVector3D) Z() float32 {
	return (float32)(QVector3D_Z(this.h))
}

func (this *QVector3D) SetX(x float32) {
	QVector3D_SetX(this.h, (float)(x))
}

func (this *QVector3D) SetY(y float32) {
	QVector3D_SetY(this.h, (float)(y))
}

func (this *QVector3D) SetZ(z float32) {
	QVector3D_SetZ(this.h, (float)(z))
}

func (this *QVector3D) OperatorSubscript(i int) float32 {
	return (float32)(QVector3D_OperatorSubscript(this.h, (int)(i)))
}

func (this *QVector3D) Length() float32 {
	return (float32)(QVector3D_Length(this.h))
}

func (this *QVector3D) LengthSquared() float32 {
	return (float32)(QVector3D_LengthSquared(this.h))
}

func (this *QVector3D) Normalized() *QVector3D {
	_goptr := newQVector3D(QVector3D_Normalized(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVector3D) Normalize() {
	QVector3D_Normalize(this.h)
}

func (this *QVector3D) OperatorPlusAssign(vector QVector3D) *QVector3D {
	return newQVector3D(QVector3D_OperatorPlusAssign(this.h, vector.cPointer()))
}

func (this *QVector3D) OperatorMinusAssign(vector QVector3D) *QVector3D {
	return newQVector3D(QVector3D_OperatorMinusAssign(this.h, vector.cPointer()))
}

func (this *QVector3D) OperatorMultiplyAssign(factor float32) *QVector3D {
	return newQVector3D(QVector3D_OperatorMultiplyAssign(this.h, (float)(factor)))
}

func (this *QVector3D) OperatorMultiplyAssignWithVector(vector QVector3D) *QVector3D {
	return newQVector3D(QVector3D_OperatorMultiplyAssignWithVector(this.h, vector.cPointer()))
}

func (this *QVector3D) OperatorDivideAssign(divisor float32) *QVector3D {
	return newQVector3D(QVector3D_OperatorDivideAssign(this.h, (float)(divisor)))
}

func (this *QVector3D) OperatorDivideAssignWithVector(vector QVector3D) *QVector3D {
	return newQVector3D(QVector3D_OperatorDivideAssignWithVector(this.h, vector.cPointer()))
}

func QVector3D_DotProduct(v1 QVector3D, v2 QVector3D) float32 {
	return (float32)(QVector3D_DotProduct(v1.cPointer(), v2.cPointer()))
}

func QVector3D_CrossProduct(v1 QVector3D, v2 QVector3D) *QVector3D {
	_goptr := newQVector3D(QVector3D_CrossProduct(v1.cPointer(), v2.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QVector3D_Normal(v1 QVector3D, v2 QVector3D) *QVector3D {
	_goptr := newQVector3D(QVector3D_Normal(v1.cPointer(), v2.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QVector3D_Normal2(v1 QVector3D, v2 QVector3D, v3 QVector3D) *QVector3D {
	_goptr := newQVector3D(QVector3D_Normal2(v1.cPointer(), v2.cPointer(), v3.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVector3D) Project(modelView *QMatrix4x4, projection *QMatrix4x4, viewport *QRect) *QVector3D {
	_goptr := newQVector3D(QVector3D_Project(this.h, modelView.cPointer(), projection.cPointer(), viewport.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVector3D) Unproject(modelView *QMatrix4x4, projection *QMatrix4x4, viewport *QRect) *QVector3D {
	_goptr := newQVector3D(QVector3D_Unproject(this.h, modelView.cPointer(), projection.cPointer(), viewport.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVector3D) DistanceToPoint(point QVector3D) float32 {
	return (float32)(QVector3D_DistanceToPoint(this.h, point.cPointer()))
}

func (this *QVector3D) DistanceToPlane(plane QVector3D, normal QVector3D) float32 {
	return (float32)(QVector3D_DistanceToPlane(this.h, plane.cPointer(), normal.cPointer()))
}

func (this *QVector3D) DistanceToPlane2(plane1 QVector3D, plane2 QVector3D, plane3 QVector3D) float32 {
	return (float32)(QVector3D_DistanceToPlane2(this.h, plane1.cPointer(), plane2.cPointer(), plane3.cPointer()))
}

func (this *QVector3D) DistanceToLine(point QVector3D, direction QVector3D) float32 {
	return (float32)(QVector3D_DistanceToLine(this.h, point.cPointer(), direction.cPointer()))
}

func (this *QVector3D) ToVector2D() *QVector2D {
	_goptr := newQVector2D(QVector3D_ToVector2D(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVector3D) ToVector4D() *QVector4D {
	_goptr := newQVector4D(QVector3D_ToVector4D(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVector3D) ToPoint() *QPoint {
	_goptr := newQPoint(QVector3D_ToPoint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVector3D) ToPointF() *QPointF {
	_goptr := newQPointF(QVector3D_ToPointF(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QVector4D struct {
	h          uintptr
	isSubclass bool
}

// NewQVector4D constructs a new QVector4D object.
func NewQVector4D() *QVector4D {

	ret := newQVector4D(QVector4D_new())
	ret.isSubclass = true
	return ret
}

// NewQVector4D2 constructs a new QVector4D object.
func NewQVector4D2(param1 Initialization) *QVector4D {

	ret := newQVector4D(QVector4D_new2((int)(param1)))
	ret.isSubclass = true
	return ret
}

// NewQVector4D3 constructs a new QVector4D object.
func NewQVector4D3(xpos float32, ypos float32, zpos float32, wpos float32) *QVector4D {

	ret := newQVector4D(QVector4D_new3((float)(xpos), (float)(ypos), (float)(zpos), (float)(wpos)))
	ret.isSubclass = true
	return ret
}

// NewQVector4D4 constructs a new QVector4D object.
func NewQVector4D4(point QPoint) *QVector4D {

	ret := newQVector4D(QVector4D_new4(point.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVector4D5 constructs a new QVector4D object.
func NewQVector4D5(point QPointF) *QVector4D {

	ret := newQVector4D(QVector4D_new5(point.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVector4D6 constructs a new QVector4D object.
func NewQVector4D6(vector QVector2D) *QVector4D {

	ret := newQVector4D(QVector4D_new6(vector.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVector4D7 constructs a new QVector4D object.
func NewQVector4D7(vector QVector2D, zpos float32, wpos float32) *QVector4D {

	ret := newQVector4D(QVector4D_new7(vector.cPointer(), (float)(zpos), (float)(wpos)))
	ret.isSubclass = true
	return ret
}

// NewQVector4D8 constructs a new QVector4D object.
func NewQVector4D8(vector QVector3D) *QVector4D {

	ret := newQVector4D(QVector4D_new8(vector.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVector4D9 constructs a new QVector4D object.
func NewQVector4D9(vector QVector3D, wpos float32) *QVector4D {

	ret := newQVector4D(QVector4D_new9(vector.cPointer(), (float)(wpos)))
	ret.isSubclass = true
	return ret
}

// NewQVector4D10 constructs a new QVector4D object.
func NewQVector4D10(param1 *QVector4D) *QVector4D {

	ret := newQVector4D(QVector4D_new10(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QVector4D) IsNull() bool {
	return (bool)(QVector4D_IsNull(this.h))
}

func (this *QVector4D) X() float32 {
	return (float32)(QVector4D_X(this.h))
}

func (this *QVector4D) Y() float32 {
	return (float32)(QVector4D_Y(this.h))
}

func (this *QVector4D) Z() float32 {
	return (float32)(QVector4D_Z(this.h))
}

func (this *QVector4D) W() float32 {
	return (float32)(QVector4D_W(this.h))
}

func (this *QVector4D) SetX(x float32) {
	QVector4D_SetX(this.h, (float)(x))
}

func (this *QVector4D) SetY(y float32) {
	QVector4D_SetY(this.h, (float)(y))
}

func (this *QVector4D) SetZ(z float32) {
	QVector4D_SetZ(this.h, (float)(z))
}

func (this *QVector4D) SetW(w float32) {
	QVector4D_SetW(this.h, (float)(w))
}

func (this *QVector4D) OperatorSubscript(i int) float32 {
	return (float32)(QVector4D_OperatorSubscript(this.h, (int)(i)))
}

func (this *QVector4D) Length() float32 {
	return (float32)(QVector4D_Length(this.h))
}

func (this *QVector4D) LengthSquared() float32 {
	return (float32)(QVector4D_LengthSquared(this.h))
}

func (this *QVector4D) Normalized() *QVector4D {
	_goptr := newQVector4D(QVector4D_Normalized(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVector4D) Normalize() {
	QVector4D_Normalize(this.h)
}

func (this *QVector4D) OperatorPlusAssign(vector QVector4D) *QVector4D {
	return newQVector4D(QVector4D_OperatorPlusAssign(this.h, vector.cPointer()))
}

func (this *QVector4D) OperatorMinusAssign(vector QVector4D) *QVector4D {
	return newQVector4D(QVector4D_OperatorMinusAssign(this.h, vector.cPointer()))
}

func (this *QVector4D) OperatorMultiplyAssign(factor float32) *QVector4D {
	return newQVector4D(QVector4D_OperatorMultiplyAssign(this.h, (float)(factor)))
}

func (this *QVector4D) OperatorMultiplyAssignWithVector(vector QVector4D) *QVector4D {
	return newQVector4D(QVector4D_OperatorMultiplyAssignWithVector(this.h, vector.cPointer()))
}

func (this *QVector4D) OperatorDivideAssign(divisor float32) *QVector4D {
	return newQVector4D(QVector4D_OperatorDivideAssign(this.h, (float)(divisor)))
}

func (this *QVector4D) OperatorDivideAssignWithVector(vector QVector4D) *QVector4D {
	return newQVector4D(QVector4D_OperatorDivideAssignWithVector(this.h, vector.cPointer()))
}

func QVector4D_DotProduct(v1 QVector4D, v2 QVector4D) float32 {
	return (float32)(QVector4D_DotProduct(v1.cPointer(), v2.cPointer()))
}

func (this *QVector4D) ToVector2D() *QVector2D {
	_goptr := newQVector2D(QVector4D_ToVector2D(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVector4D) ToVector2DAffine() *QVector2D {
	_goptr := newQVector2D(QVector4D_ToVector2DAffine(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVector4D) ToVector3D() *QVector3D {
	_goptr := newQVector3D(QVector4D_ToVector3D(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVector4D) ToVector3DAffine() *QVector3D {
	_goptr := newQVector3D(QVector4D_ToVector3DAffine(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVector4D) ToPoint() *QPoint {
	_goptr := newQPoint(QVector4D_ToPoint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVector4D) ToPointF() *QPointF {
	_goptr := newQPointF(QVector4D_ToPointF(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}