package qt

import (
	"unsafe"
)

type QBasicTimer struct {
	h          uintptr
	isSubclass bool
}

// NewQBasicTimer constructs a new QBasicTimer object.
func NewQBasicTimer() *QBasicTimer {
	ret := newQBasicTimer(QBasicTimer_new())
	ret.isSubclass = true
	return ret
}

func (this *QBasicTimer) Swap(other *QBasicTimer) {
	QBasicTimer_Swap(this.h, other.cPointer())
}

func (this *QBasicTimer) IsActive() bool {
	return (bool)(QBasicTimer_IsActive(this.h))
}

func (this *QBasicTimer) TimerId() int {
	return (int)(QBasicTimer_TimerId(this.h))
}

func (this *QBasicTimer) Id() TimerId {
	return (TimerId)(QBasicTimer_Id(this.h))
}

func (this *QBasicTimer) Start(msec int, obj *QObject) {
	QBasicTimer_Start(this.h, (int)(msec), obj.cPointer())
}

func (this *QBasicTimer) Start2(msec int, timerType TimerType, obj *QObject) {
	QBasicTimer_Start2(this.h, (int)(msec), (int)(timerType), obj.cPointer())
}

func (this *QBasicTimer) Start3(duration Duration, obj *QObject) {
	QBasicTimer_Start3(this.h, duration, obj.cPointer())
}

func (this *QBasicTimer) Start4(duration Duration, timerType TimerType, obj *QObject) {
	QBasicTimer_Start4(this.h, duration, (int)(timerType), obj.cPointer())
}

func (this *QBasicTimer) Stop() {
	QBasicTimer_Stop(this.h)
}
