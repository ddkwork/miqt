package qt

import (
	"unsafe"
)

type QElapsedTimer__ClockType int

const (
	QElapsedTimer__SystemTime         QElapsedTimer__ClockType = 0
	QElapsedTimer__MonotonicClock     QElapsedTimer__ClockType = 1
	QElapsedTimer__MachAbsoluteTime   QElapsedTimer__ClockType = 2
	QElapsedTimer__PerformanceCounter QElapsedTimer__ClockType = 3
)

type QElapsedTimer struct {
	h          uintptr
	isSubclass bool
}

// NewQElapsedTimer constructs a new QElapsedTimer object.
func NewQElapsedTimer() *QElapsedTimer {
	ret := newQElapsedTimer(QElapsedTimer_new())
	ret.isSubclass = true
	return ret
}

func QElapsedTimer_ClockType() ClockType {
	xxxxxxxxx
}

func QElapsedTimer_IsMonotonic() bool {
	return (bool)(QElapsedTimer_IsMonotonic())
}

func (this *QElapsedTimer) Start() {
	QElapsedTimer_Start(this.h)
}

func (this *QElapsedTimer) Restart() int64 {
	return (int64)(QElapsedTimer_Restart(this.h))
}

func (this *QElapsedTimer) Invalidate() {
	QElapsedTimer_Invalidate(this.h)
}

func (this *QElapsedTimer) IsValid() bool {
	return (bool)(QElapsedTimer_IsValid(this.h))
}

func (this *QElapsedTimer) DurationElapsed() Duration {
	xxxxxxxxx
}

func (this *QElapsedTimer) NsecsElapsed() int64 {
	return (int64)(QElapsedTimer_NsecsElapsed(this.h))
}

func (this *QElapsedTimer) Elapsed() int64 {
	return (int64)(QElapsedTimer_Elapsed(this.h))
}

func (this *QElapsedTimer) HasExpired(timeout int64) bool {
	return (bool)(QElapsedTimer_HasExpired(this.h, (longlong)(timeout)))
}

func (this *QElapsedTimer) MsecsSinceReference() int64 {
	return (int64)(QElapsedTimer_MsecsSinceReference(this.h))
}

func (this *QElapsedTimer) DurationTo(other *QElapsedTimer) Duration {
	xxxxxxxxx
}

func (this *QElapsedTimer) MsecsTo(other *QElapsedTimer) int64 {
	return (int64)(QElapsedTimer_MsecsTo(this.h, other.cPointer()))
}

func (this *QElapsedTimer) SecsTo(other *QElapsedTimer) int64 {
	return (int64)(QElapsedTimer_SecsTo(this.h, other.cPointer()))
}
