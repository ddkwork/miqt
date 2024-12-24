package qt

import (
	"unsafe"
)

type QBasicMutex struct {
	h          uintptr
	isSubclass bool
}

// NewQBasicMutex constructs a new QBasicMutex object.
func NewQBasicMutex() *QBasicMutex {
	ret := newQBasicMutex(QBasicMutex_new())
	ret.isSubclass = true
	return ret
}

func (this *QBasicMutex) Lock() {
	QBasicMutex_Lock(this.h)
}

func (this *QBasicMutex) Unlock() {
	QBasicMutex_Unlock(this.h)
}

func (this *QBasicMutex) TryLock() bool {
	return (bool)(QBasicMutex_TryLock(this.h))
}

func (this *QBasicMutex) TryLock2() bool {
	return (bool)(QBasicMutex_TryLock2(this.h))
}

type QMutex struct {
	h          uintptr
	isSubclass bool
}

// NewQMutex constructs a new QMutex object.
func NewQMutex() *QMutex {
	ret := newQMutex(QMutex_new())
	ret.isSubclass = true
	return ret
}

func (this *QMutex) TryLock() bool {
	return (bool)(QMutex_TryLock(this.h))
}

func (this *QMutex) TryLockWithTimeout(timeout int) bool {
	return (bool)(QMutex_TryLockWithTimeout(this.h, (int)(timeout)))
}

func (this *QMutex) TryLock2(timeout QDeadlineTimer) bool {
	return (bool)(QMutex_TryLock2(this.h, timeout.cPointer()))
}

type QRecursiveMutex struct {
	h          uintptr
	isSubclass bool
}

// NewQRecursiveMutex constructs a new QRecursiveMutex object.
func NewQRecursiveMutex() *QRecursiveMutex {
	ret := newQRecursiveMutex(QRecursiveMutex_new())
	ret.isSubclass = true
	return ret
}

func (this *QRecursiveMutex) Lock() {
	QRecursiveMutex_Lock(this.h)
}

func (this *QRecursiveMutex) TryLock(timeout int) bool {
	return (bool)(QRecursiveMutex_TryLock(this.h, (int)(timeout)))
}

func (this *QRecursiveMutex) TryLock2() bool {
	return (bool)(QRecursiveMutex_TryLock2(this.h))
}

func (this *QRecursiveMutex) Unlock() {
	QRecursiveMutex_Unlock(this.h)
}

func (this *QRecursiveMutex) TryLock3() bool {
	return (bool)(QRecursiveMutex_TryLock3(this.h))
}

func (this *QRecursiveMutex) TryLock1(timer QDeadlineTimer) bool {
	return (bool)(QRecursiveMutex_TryLock1(this.h, timer.cPointer()))
}
