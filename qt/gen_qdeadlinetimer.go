package qt

import (
	"unsafe"
)

type QDeadlineTimer__ForeverConstant int

const (
	QDeadlineTimer__Forever QDeadlineTimer__ForeverConstant = 0
)

type QDeadlineTimer struct {
	h          uintptr
	isSubclass bool
}

// NewQDeadlineTimer constructs a new QDeadlineTimer object.
func NewQDeadlineTimer() *QDeadlineTimer {

	ret := newQDeadlineTimer(QDeadlineTimer_new())
	ret.isSubclass = true
	return ret
}

// NewQDeadlineTimer2 constructs a new QDeadlineTimer object.
func NewQDeadlineTimer2(type_ TimerType) *QDeadlineTimer {

	ret := newQDeadlineTimer(QDeadlineTimer_new2((int)(type_)))
	ret.isSubclass = true
	return ret
}

// NewQDeadlineTimer3 constructs a new QDeadlineTimer object.
func NewQDeadlineTimer3(param1 ForeverConstant) *QDeadlineTimer {

	ret := newQDeadlineTimer(QDeadlineTimer_new3(param1))
	ret.isSubclass = true
	return ret
}

// NewQDeadlineTimer4 constructs a new QDeadlineTimer object.
func NewQDeadlineTimer4(msecs int64) *QDeadlineTimer {

	ret := newQDeadlineTimer(QDeadlineTimer_new4((longlong)(msecs)))
	ret.isSubclass = true
	return ret
}

// NewQDeadlineTimer5 constructs a new QDeadlineTimer object.
func NewQDeadlineTimer5(param1 *QDeadlineTimer) *QDeadlineTimer {

	ret := newQDeadlineTimer(QDeadlineTimer_new5(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDeadlineTimer6 constructs a new QDeadlineTimer object.
func NewQDeadlineTimer6(param1 ForeverConstant, type_ TimerType) *QDeadlineTimer {

	ret := newQDeadlineTimer(QDeadlineTimer_new6(param1, (int)(type_)))
	ret.isSubclass = true
	return ret
}

// NewQDeadlineTimer7 constructs a new QDeadlineTimer object.
func NewQDeadlineTimer7(msecs int64, typeVal TimerType) *QDeadlineTimer {

	ret := newQDeadlineTimer(QDeadlineTimer_new7((longlong)(msecs), (int)(typeVal)))
	ret.isSubclass = true
	return ret
}

func (this *QDeadlineTimer) Swap(other *QDeadlineTimer) {
	QDeadlineTimer_Swap(this.h, other.cPointer())
}

func (this *QDeadlineTimer) IsForever() bool {
	return (bool)(QDeadlineTimer_IsForever(this.h))
}

func (this *QDeadlineTimer) HasExpired() bool {
	return (bool)(QDeadlineTimer_HasExpired(this.h))
}

func (this *QDeadlineTimer) TimerType() TimerType {
	return (TimerType)(QDeadlineTimer_TimerType(this.h))
}

func (this *QDeadlineTimer) SetTimerType(typeVal TimerType) {
	QDeadlineTimer_SetTimerType(this.h, (int)(typeVal))
}

func (this *QDeadlineTimer) RemainingTime() int64 {
	return (int64)(QDeadlineTimer_RemainingTime(this.h))
}

func (this *QDeadlineTimer) RemainingTimeNSecs() int64 {
	return (int64)(QDeadlineTimer_RemainingTimeNSecs(this.h))
}

func (this *QDeadlineTimer) SetRemainingTime(msecs int64) {
	QDeadlineTimer_SetRemainingTime(this.h, (longlong)(msecs))
}

func (this *QDeadlineTimer) SetPreciseRemainingTime(secs int64) {
	QDeadlineTimer_SetPreciseRemainingTime(this.h, (longlong)(secs))
}

func (this *QDeadlineTimer) Deadline() int64 {
	return (int64)(QDeadlineTimer_Deadline(this.h))
}

func (this *QDeadlineTimer) DeadlineNSecs() int64 {
	return (int64)(QDeadlineTimer_DeadlineNSecs(this.h))
}

func (this *QDeadlineTimer) SetDeadline(msecs int64) {
	QDeadlineTimer_SetDeadline(this.h, (longlong)(msecs))
}

func (this *QDeadlineTimer) SetPreciseDeadline(secs int64) {
	QDeadlineTimer_SetPreciseDeadline(this.h, (longlong)(secs))
}

func QDeadlineTimer_AddNSecs(dt QDeadlineTimer, nsecs int64) *QDeadlineTimer {
	_goptr := newQDeadlineTimer(QDeadlineTimer_AddNSecs(dt.cPointer(), (longlong)(nsecs)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDeadlineTimer_Current() *QDeadlineTimer {
	_goptr := newQDeadlineTimer(QDeadlineTimer_Current())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDeadlineTimer) OperatorPlusAssign(msecs int64) *QDeadlineTimer {
	return newQDeadlineTimer(QDeadlineTimer_OperatorPlusAssign(this.h, (longlong)(msecs)))
}

func (this *QDeadlineTimer) OperatorMinusAssign(msecs int64) *QDeadlineTimer {
	return newQDeadlineTimer(QDeadlineTimer_OperatorMinusAssign(this.h, (longlong)(msecs)))
}

func (this *QDeadlineTimer) SetRemainingTime2(msecs int64, typeVal TimerType) {
	QDeadlineTimer_SetRemainingTime2(this.h, (longlong)(msecs), (int)(typeVal))
}

func (this *QDeadlineTimer) SetPreciseRemainingTime2(secs int64, nsecs int64) {
	QDeadlineTimer_SetPreciseRemainingTime2(this.h, (longlong)(secs), (longlong)(nsecs))
}

func (this *QDeadlineTimer) SetPreciseRemainingTime3(secs int64, nsecs int64, typeVal TimerType) {
	QDeadlineTimer_SetPreciseRemainingTime3(this.h, (longlong)(secs), (longlong)(nsecs), (int)(typeVal))
}

func (this *QDeadlineTimer) SetDeadline2(msecs int64, timerType TimerType) {
	QDeadlineTimer_SetDeadline2(this.h, (longlong)(msecs), (int)(timerType))
}

func (this *QDeadlineTimer) SetPreciseDeadline2(secs int64, nsecs int64) {
	QDeadlineTimer_SetPreciseDeadline2(this.h, (longlong)(secs), (longlong)(nsecs))
}

func (this *QDeadlineTimer) SetPreciseDeadline3(secs int64, nsecs int64, typeVal TimerType) {
	QDeadlineTimer_SetPreciseDeadline3(this.h, (longlong)(secs), (longlong)(nsecs), (int)(typeVal))
}

func QDeadlineTimer_Current1(timerType TimerType) *QDeadlineTimer {
	_goptr := newQDeadlineTimer(QDeadlineTimer_Current1((int)(timerType)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
