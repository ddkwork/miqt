package qt

import (
	"unsafe"
)

type QRunnable struct {
	h          uintptr
	isSubclass bool
}

// NewQRunnable constructs a new QRunnable object.
func NewQRunnable() *QRunnable {
	ret := newQRunnable(QRunnable_new())
	ret.isSubclass = true
	return ret
}

func (this *QRunnable) Run() {
	QRunnable_Run(this.h)
}

func (this *QRunnable) AutoDelete() bool {
	return (bool)(QRunnable_AutoDelete(this.h))
}

func (this *QRunnable) SetAutoDelete(autoDelete bool) {
	QRunnable_SetAutoDelete(this.h, (bool)(autoDelete))
}

type QGenericRunnable struct {
	h          uintptr
	isSubclass bool
}

func (this *QGenericRunnable) Run() {
	QGenericRunnable_Run(this.h)
}
