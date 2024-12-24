package qt

import (
	"unsafe"
)

type QReadWriteLock__RecursionMode int

const (
	QReadWriteLock__NonRecursive QReadWriteLock__RecursionMode = 0
	QReadWriteLock__Recursive    QReadWriteLock__RecursionMode = 1
)

type QReadWriteLock struct {
	h          uintptr
	isSubclass bool
}

// NewQReadWriteLock constructs a new QReadWriteLock object.
func NewQReadWriteLock() *QReadWriteLock {

	ret := newQReadWriteLock(QReadWriteLock_new())
	ret.isSubclass = true
	return ret
}

// NewQReadWriteLock2 constructs a new QReadWriteLock object.
func NewQReadWriteLock2(recursionMode RecursionMode) *QReadWriteLock {

	ret := newQReadWriteLock(QReadWriteLock_new2(recursionMode))
	ret.isSubclass = true
	return ret
}

func (this *QReadWriteLock) LockForRead() {
	QReadWriteLock_LockForRead(this.h)
}

func (this *QReadWriteLock) TryLockForRead(timeout int) bool {
	return (bool)(QReadWriteLock_TryLockForRead(this.h, (int)(timeout)))
}

func (this *QReadWriteLock) TryLockForRead2() bool {
	return (bool)(QReadWriteLock_TryLockForRead2(this.h))
}

func (this *QReadWriteLock) LockForWrite() {
	QReadWriteLock_LockForWrite(this.h)
}

func (this *QReadWriteLock) TryLockForWrite(timeout int) bool {
	return (bool)(QReadWriteLock_TryLockForWrite(this.h, (int)(timeout)))
}

func (this *QReadWriteLock) TryLockForWrite2() bool {
	return (bool)(QReadWriteLock_TryLockForWrite2(this.h))
}

func (this *QReadWriteLock) Unlock() {
	QReadWriteLock_Unlock(this.h)
}

func (this *QReadWriteLock) TryLockForRead1(timeout QDeadlineTimer) bool {
	return (bool)(QReadWriteLock_TryLockForRead1(this.h, timeout.cPointer()))
}

func (this *QReadWriteLock) TryLockForWrite1(timeout QDeadlineTimer) bool {
	return (bool)(QReadWriteLock_TryLockForWrite1(this.h, timeout.cPointer()))
}

type QReadLocker struct {
	h          uintptr
	isSubclass bool
}

// NewQReadLocker constructs a new QReadLocker object.
func NewQReadLocker(readWriteLock *QReadWriteLock) *QReadLocker {

	ret := newQReadLocker(QReadLocker_new(readWriteLock.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QReadLocker) Unlock() {
	QReadLocker_Unlock(this.h)
}

func (this *QReadLocker) Relock() {
	QReadLocker_Relock(this.h)
}

func (this *QReadLocker) ReadWriteLock() *QReadWriteLock {
	return newQReadWriteLock(QReadLocker_ReadWriteLock(this.h))
}

type QWriteLocker struct {
	h          uintptr
	isSubclass bool
}

// NewQWriteLocker constructs a new QWriteLocker object.
func NewQWriteLocker(readWriteLock *QReadWriteLock) *QWriteLocker {

	ret := newQWriteLocker(QWriteLocker_new(readWriteLock.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QWriteLocker) Unlock() {
	QWriteLocker_Unlock(this.h)
}

func (this *QWriteLocker) Relock() {
	QWriteLocker_Relock(this.h)
}

func (this *QWriteLocker) ReadWriteLock() *QReadWriteLock {
	return newQReadWriteLock(QWriteLocker_ReadWriteLock(this.h))
}
