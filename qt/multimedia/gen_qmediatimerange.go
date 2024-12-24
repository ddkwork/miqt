package multimedia

import (
	"unsafe"
)

type QMediaTimeRange struct {
	h          uintptr
	isSubclass bool
}

// NewQMediaTimeRange constructs a new QMediaTimeRange object.
func NewQMediaTimeRange() *QMediaTimeRange {
	ret := newQMediaTimeRange(QMediaTimeRange_new())
	ret.isSubclass = true
	return ret
}

// NewQMediaTimeRange2 constructs a new QMediaTimeRange object.
func NewQMediaTimeRange2(start int64, end int64) *QMediaTimeRange {
	ret := newQMediaTimeRange(QMediaTimeRange_new2((longlong)(start), (longlong)(end)))
	ret.isSubclass = true
	return ret
}

// NewQMediaTimeRange3 constructs a new QMediaTimeRange object.
func NewQMediaTimeRange3(param1 *Interval) *QMediaTimeRange {
	ret := newQMediaTimeRange(QMediaTimeRange_new3(param1))
	ret.isSubclass = true
	return ret
}

// NewQMediaTimeRange4 constructs a new QMediaTimeRange object.
func NewQMediaTimeRange4(rangeVal *QMediaTimeRange) *QMediaTimeRange {
	ret := newQMediaTimeRange(QMediaTimeRange_new4(rangeVal.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QMediaTimeRange) OperatorAssign(param1 *QMediaTimeRange) {
	QMediaTimeRange_OperatorAssign(this.h, param1.cPointer())
}

func (this *QMediaTimeRange) Swap(other *QMediaTimeRange) {
	QMediaTimeRange_Swap(this.h, other.cPointer())
}

func (this *QMediaTimeRange) Detach() {
	QMediaTimeRange_Detach(this.h)
}

func (this *QMediaTimeRange) OperatorAssignWithInterval(param1 *Interval) {
	QMediaTimeRange_OperatorAssignWithInterval(this.h, param1)
}

func (this *QMediaTimeRange) EarliestTime() int64 {
	return (int64)(QMediaTimeRange_EarliestTime(this.h))
}

func (this *QMediaTimeRange) LatestTime() int64 {
	return (int64)(QMediaTimeRange_LatestTime(this.h))
}

func (this *QMediaTimeRange) Intervals() []QMediaTimeRange__Interval {
	var _ma struct_miqt_array = QMediaTimeRange_Intervals(this.h)
	_ret := make([]QMediaTimeRange__Interval, int(_ma.len))
	_outCast := (*[0xffff]*QMediaTimeRange__Interval)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQMediaTimeRange__Interval(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QMediaTimeRange) IsEmpty() bool {
	return (bool)(QMediaTimeRange_IsEmpty(this.h))
}

func (this *QMediaTimeRange) IsContinuous() bool {
	return (bool)(QMediaTimeRange_IsContinuous(this.h))
}

func (this *QMediaTimeRange) Contains(time int64) bool {
	return (bool)(QMediaTimeRange_Contains(this.h, (longlong)(time)))
}

func (this *QMediaTimeRange) AddInterval(start int64, end int64) {
	QMediaTimeRange_AddInterval(this.h, (longlong)(start), (longlong)(end))
}

func (this *QMediaTimeRange) AddIntervalWithInterval(interval *Interval) {
	QMediaTimeRange_AddIntervalWithInterval(this.h, interval)
}

func (this *QMediaTimeRange) AddTimeRange(param1 *QMediaTimeRange) {
	QMediaTimeRange_AddTimeRange(this.h, param1.cPointer())
}

func (this *QMediaTimeRange) RemoveInterval(start int64, end int64) {
	QMediaTimeRange_RemoveInterval(this.h, (longlong)(start), (longlong)(end))
}

func (this *QMediaTimeRange) RemoveIntervalWithInterval(interval *Interval) {
	QMediaTimeRange_RemoveIntervalWithInterval(this.h, interval)
}

func (this *QMediaTimeRange) RemoveTimeRange(param1 *QMediaTimeRange) {
	QMediaTimeRange_RemoveTimeRange(this.h, param1.cPointer())
}

func (this *QMediaTimeRange) OperatorPlusAssign(param1 *QMediaTimeRange) *QMediaTimeRange {
	return newQMediaTimeRange(QMediaTimeRange_OperatorPlusAssign(this.h, param1.cPointer()))
}

func (this *QMediaTimeRange) OperatorPlusAssignWithInterval(param1 *Interval) *QMediaTimeRange {
	return newQMediaTimeRange(QMediaTimeRange_OperatorPlusAssignWithInterval(this.h, param1))
}

func (this *QMediaTimeRange) OperatorMinusAssign(param1 *QMediaTimeRange) *QMediaTimeRange {
	return newQMediaTimeRange(QMediaTimeRange_OperatorMinusAssign(this.h, param1.cPointer()))
}

func (this *QMediaTimeRange) OperatorMinusAssignWithInterval(param1 *Interval) *QMediaTimeRange {
	return newQMediaTimeRange(QMediaTimeRange_OperatorMinusAssignWithInterval(this.h, param1))
}

func (this *QMediaTimeRange) Clear() {
	QMediaTimeRange_Clear(this.h)
}

type QMediaTimeRange__Interval struct {
	h          uintptr
	isSubclass bool
}

// NewQMediaTimeRange__Interval constructs a new QMediaTimeRange::Interval object.
func NewQMediaTimeRange__Interval() *QMediaTimeRange__Interval {
	ret := newQMediaTimeRange__Interval(QMediaTimeRange__Interval_new())
	ret.isSubclass = true
	return ret
}

// NewQMediaTimeRange__Interval2 constructs a new QMediaTimeRange::Interval object.
func NewQMediaTimeRange__Interval2(start int64, end int64) *QMediaTimeRange__Interval {
	ret := newQMediaTimeRange__Interval(QMediaTimeRange__Interval_new2((longlong)(start), (longlong)(end)))
	ret.isSubclass = true
	return ret
}

// NewQMediaTimeRange__Interval3 constructs a new QMediaTimeRange::Interval object.
func NewQMediaTimeRange__Interval3(param1 *Interval) *QMediaTimeRange__Interval {
	ret := newQMediaTimeRange__Interval(QMediaTimeRange__Interval_new3(param1))
	ret.isSubclass = true
	return ret
}

func (this *QMediaTimeRange__Interval) Start() int64 {
	return (int64)(QMediaTimeRange__Interval_Start(this.h))
}

func (this *QMediaTimeRange__Interval) End() int64 {
	return (int64)(QMediaTimeRange__Interval_End(this.h))
}

func (this *QMediaTimeRange__Interval) Contains(time int64) bool {
	return (bool)(QMediaTimeRange__Interval_Contains(this.h, (longlong)(time)))
}

func (this *QMediaTimeRange__Interval) IsNormal() bool {
	return (bool)(QMediaTimeRange__Interval_IsNormal(this.h))
}

func (this *QMediaTimeRange__Interval) Normalized() Interval {
	xxxxxxxxx
}

func (this *QMediaTimeRange__Interval) Translated(offset int64) Interval {
	xxxxxxxxx
}
