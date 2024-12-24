package qt

import (
	"unsafe"
)

type QMargins struct {
	h          uintptr
	isSubclass bool
}

// NewQMargins constructs a new QMargins object.
func NewQMargins() *QMargins {

	ret := newQMargins(QMargins_new())
	ret.isSubclass = true
	return ret
}

// NewQMargins2 constructs a new QMargins object.
func NewQMargins2(left int, top int, right int, bottom int) *QMargins {

	ret := newQMargins(QMargins_new2((int)(left), (int)(top), (int)(right), (int)(bottom)))
	ret.isSubclass = true
	return ret
}

// NewQMargins3 constructs a new QMargins object.
func NewQMargins3(param1 *QMargins) *QMargins {

	ret := newQMargins(QMargins_new3(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QMargins) IsNull() bool {
	return (bool)(QMargins_IsNull(this.h))
}

func (this *QMargins) Left() int {
	return (int)(QMargins_Left(this.h))
}

func (this *QMargins) Top() int {
	return (int)(QMargins_Top(this.h))
}

func (this *QMargins) Right() int {
	return (int)(QMargins_Right(this.h))
}

func (this *QMargins) Bottom() int {
	return (int)(QMargins_Bottom(this.h))
}

func (this *QMargins) SetLeft(left int) {
	QMargins_SetLeft(this.h, (int)(left))
}

func (this *QMargins) SetTop(top int) {
	QMargins_SetTop(this.h, (int)(top))
}

func (this *QMargins) SetRight(right int) {
	QMargins_SetRight(this.h, (int)(right))
}

func (this *QMargins) SetBottom(bottom int) {
	QMargins_SetBottom(this.h, (int)(bottom))
}

func (this *QMargins) OperatorPlusAssign(margins *QMargins) *QMargins {
	return newQMargins(QMargins_OperatorPlusAssign(this.h, margins.cPointer()))
}

func (this *QMargins) OperatorMinusAssign(margins *QMargins) *QMargins {
	return newQMargins(QMargins_OperatorMinusAssign(this.h, margins.cPointer()))
}

func (this *QMargins) OperatorPlusAssignWithInt(param1 int) *QMargins {
	return newQMargins(QMargins_OperatorPlusAssignWithInt(this.h, (int)(param1)))
}

func (this *QMargins) OperatorMinusAssignWithInt(param1 int) *QMargins {
	return newQMargins(QMargins_OperatorMinusAssignWithInt(this.h, (int)(param1)))
}

func (this *QMargins) OperatorMultiplyAssign(param1 int) *QMargins {
	return newQMargins(QMargins_OperatorMultiplyAssign(this.h, (int)(param1)))
}

func (this *QMargins) OperatorDivideAssign(param1 int) *QMargins {
	return newQMargins(QMargins_OperatorDivideAssign(this.h, (int)(param1)))
}

func (this *QMargins) OperatorMultiplyAssignWithQreal(param1 float64) *QMargins {
	return newQMargins(QMargins_OperatorMultiplyAssignWithQreal(this.h, (double)(param1)))
}

func (this *QMargins) OperatorDivideAssignWithQreal(param1 float64) *QMargins {
	return newQMargins(QMargins_OperatorDivideAssignWithQreal(this.h, (double)(param1)))
}

func (this *QMargins) ToMarginsF() *QMarginsF {
	_goptr := newQMarginsF(QMargins_ToMarginsF(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QMarginsF struct {
	h          uintptr
	isSubclass bool
}

// NewQMarginsF constructs a new QMarginsF object.
func NewQMarginsF() *QMarginsF {

	ret := newQMarginsF(QMarginsF_new())
	ret.isSubclass = true
	return ret
}

// NewQMarginsF2 constructs a new QMarginsF object.
func NewQMarginsF2(left float64, top float64, right float64, bottom float64) *QMarginsF {

	ret := newQMarginsF(QMarginsF_new2((double)(left), (double)(top), (double)(right), (double)(bottom)))
	ret.isSubclass = true
	return ret
}

// NewQMarginsF3 constructs a new QMarginsF object.
func NewQMarginsF3(margins *QMargins) *QMarginsF {

	ret := newQMarginsF(QMarginsF_new3(margins.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQMarginsF4 constructs a new QMarginsF object.
func NewQMarginsF4(param1 *QMarginsF) *QMarginsF {

	ret := newQMarginsF(QMarginsF_new4(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QMarginsF) IsNull() bool {
	return (bool)(QMarginsF_IsNull(this.h))
}

func (this *QMarginsF) Left() float64 {
	return (float64)(QMarginsF_Left(this.h))
}

func (this *QMarginsF) Top() float64 {
	return (float64)(QMarginsF_Top(this.h))
}

func (this *QMarginsF) Right() float64 {
	return (float64)(QMarginsF_Right(this.h))
}

func (this *QMarginsF) Bottom() float64 {
	return (float64)(QMarginsF_Bottom(this.h))
}

func (this *QMarginsF) SetLeft(aleft float64) {
	QMarginsF_SetLeft(this.h, (double)(aleft))
}

func (this *QMarginsF) SetTop(atop float64) {
	QMarginsF_SetTop(this.h, (double)(atop))
}

func (this *QMarginsF) SetRight(aright float64) {
	QMarginsF_SetRight(this.h, (double)(aright))
}

func (this *QMarginsF) SetBottom(abottom float64) {
	QMarginsF_SetBottom(this.h, (double)(abottom))
}

func (this *QMarginsF) OperatorPlusAssign(margins *QMarginsF) *QMarginsF {
	return newQMarginsF(QMarginsF_OperatorPlusAssign(this.h, margins.cPointer()))
}

func (this *QMarginsF) OperatorMinusAssign(margins *QMarginsF) *QMarginsF {
	return newQMarginsF(QMarginsF_OperatorMinusAssign(this.h, margins.cPointer()))
}

func (this *QMarginsF) OperatorPlusAssignWithAddend(addend float64) *QMarginsF {
	return newQMarginsF(QMarginsF_OperatorPlusAssignWithAddend(this.h, (double)(addend)))
}

func (this *QMarginsF) OperatorMinusAssignWithSubtrahend(subtrahend float64) *QMarginsF {
	return newQMarginsF(QMarginsF_OperatorMinusAssignWithSubtrahend(this.h, (double)(subtrahend)))
}

func (this *QMarginsF) OperatorMultiplyAssign(factor float64) *QMarginsF {
	return newQMarginsF(QMarginsF_OperatorMultiplyAssign(this.h, (double)(factor)))
}

func (this *QMarginsF) OperatorDivideAssign(divisor float64) *QMarginsF {
	return newQMarginsF(QMarginsF_OperatorDivideAssign(this.h, (double)(divisor)))
}

func (this *QMarginsF) ToMargins() *QMargins {
	_goptr := newQMargins(QMarginsF_ToMargins(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
