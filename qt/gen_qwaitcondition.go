package qt

import (
	"unsafe"
)

type QWaitCondition struct {
	h          uintptr
	isSubclass bool
}

// NewQWaitCondition constructs a new QWaitCondition object.
func NewQWaitCondition() *QWaitCondition {
	g := newQWaitCondition(QWaitCondition_new())
	g.isSubclass = true
	return g
}

func (this *QWaitCondition) Wait(lockedMutex *QMutex) bool {
	return (bool)(QWaitCondition_Wait(this.h, lockedMutex.cPointer()))
}

func (this *QWaitCondition) Wait2(lockedMutex *QMutex, time uint32) bool {
	return (bool)(QWaitCondition_Wait2(this.h, lockedMutex.cPointer(), (ulong)(time)))
}

func (this *QWaitCondition) WaitWithLockedReadWriteLock(lockedReadWriteLock *QReadWriteLock) bool {
	return (bool)(QWaitCondition_WaitWithLockedReadWriteLock(this.h, lockedReadWriteLock.cPointer()))
}

func (this *QWaitCondition) Wait3(lockedReadWriteLock *QReadWriteLock, time uint32) bool {
	return (bool)(QWaitCondition_Wait3(this.h, lockedReadWriteLock.cPointer(), (ulong)(time)))
}

func (this *QWaitCondition) WakeOne() {
	QWaitCondition_WakeOne(this.h)
}

func (this *QWaitCondition) WakeAll() {
	QWaitCondition_WakeAll(this.h)
}

func (this *QWaitCondition) NotifyOne() {
	QWaitCondition_NotifyOne(this.h)
}

func (this *QWaitCondition) NotifyAll() {
	QWaitCondition_NotifyAll(this.h)
}

func (this *QWaitCondition) Wait22(lockedMutex *QMutex, deadline QDeadlineTimer) bool {
	return (bool)(QWaitCondition_Wait22(this.h, lockedMutex.cPointer(), deadline.cPointer()))
}

func (this *QWaitCondition) Wait23(lockedReadWriteLock *QReadWriteLock, deadline QDeadlineTimer) bool {
	return (bool)(QWaitCondition_Wait23(this.h, lockedReadWriteLock.cPointer(), deadline.cPointer()))
}
