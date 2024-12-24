package qt

import (
	"unsafe"
)

type QSemaphore struct {
	h          uintptr
	isSubclass bool
}

// NewQSemaphore constructs a new QSemaphore object.
func NewQSemaphore() *QSemaphore {
	g := newQSemaphore(QSemaphore_new())
	g.isSubclass = true
	return g
}

// NewQSemaphore2 constructs a new QSemaphore object.
func NewQSemaphore2(n int) *QSemaphore {
	g := newQSemaphore(QSemaphore_new2((int)(n)))
	g.isSubclass = true
	return g
}

func (this *QSemaphore) Acquire() {
	QSemaphore_Acquire(this.h)
}

func (this *QSemaphore) TryAcquire() bool {
	return (bool)(QSemaphore_TryAcquire(this.h))
}

func (this *QSemaphore) TryAcquire2(n int, timeout int) bool {
	return (bool)(QSemaphore_TryAcquire2(this.h, (int)(n), (int)(timeout)))
}

func (this *QSemaphore) TryAcquire3(n int, timeout QDeadlineTimer) bool {
	return (bool)(QSemaphore_TryAcquire3(this.h, (int)(n), timeout.cPointer()))
}

func (this *QSemaphore) Release() {
	QSemaphore_Release(this.h)
}

func (this *QSemaphore) Available() int {
	return (int)(QSemaphore_Available(this.h))
}

func (this *QSemaphore) TryAcquire4() bool {
	return (bool)(QSemaphore_TryAcquire4(this.h))
}

func (this *QSemaphore) Acquire1(n int) {
	QSemaphore_Acquire1(this.h, (int)(n))
}

func (this *QSemaphore) TryAcquire1(n int) bool {
	return (bool)(QSemaphore_TryAcquire1(this.h, (int)(n)))
}

func (this *QSemaphore) Release1(n int) {
	QSemaphore_Release1(this.h, (int)(n))
}

type QSemaphoreReleaser struct {
	h          uintptr
	isSubclass bool
}

// NewQSemaphoreReleaser constructs a new QSemaphoreReleaser object.
func NewQSemaphoreReleaser() *QSemaphoreReleaser {
	g := newQSemaphoreReleaser(QSemaphoreReleaser_new())
	g.isSubclass = true
	return g
}

// NewQSemaphoreReleaser2 constructs a new QSemaphoreReleaser object.
func NewQSemaphoreReleaser2(sem *QSemaphore) *QSemaphoreReleaser {
	g := newQSemaphoreReleaser(QSemaphoreReleaser_new2(sem.cPointer()))
	g.isSubclass = true
	return g
}

// NewQSemaphoreReleaser3 constructs a new QSemaphoreReleaser object.
func NewQSemaphoreReleaser3(sem *QSemaphore) *QSemaphoreReleaser {
	g := newQSemaphoreReleaser(QSemaphoreReleaser_new3(sem.cPointer()))
	g.isSubclass = true
	return g
}

// NewQSemaphoreReleaser4 constructs a new QSemaphoreReleaser object.
func NewQSemaphoreReleaser4(sem *QSemaphore, n int) *QSemaphoreReleaser {
	g := newQSemaphoreReleaser(QSemaphoreReleaser_new4(sem.cPointer(), (int)(n)))
	g.isSubclass = true
	return g
}

// NewQSemaphoreReleaser5 constructs a new QSemaphoreReleaser object.
func NewQSemaphoreReleaser5(sem *QSemaphore, n int) *QSemaphoreReleaser {
	g := newQSemaphoreReleaser(QSemaphoreReleaser_new5(sem.cPointer(), (int)(n)))
	g.isSubclass = true
	return g
}

func (this *QSemaphoreReleaser) Swap(other *QSemaphoreReleaser) {
	QSemaphoreReleaser_Swap(this.h, other.cPointer())
}

func (this *QSemaphoreReleaser) Semaphore() *QSemaphore {
	return newQSemaphore(QSemaphoreReleaser_Semaphore(this.h))
}

func (this *QSemaphoreReleaser) Cancel() *QSemaphore {
	return newQSemaphore(QSemaphoreReleaser_Cancel(this.h))
}
