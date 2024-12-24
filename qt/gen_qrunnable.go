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
func (this *QRunnable) OnRun(slot func()) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRunnable_override_virtual_Run(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRunnable_Run
func miqt_exec_callback_QRunnable_Run(self QRunnable, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()

}

type QGenericRunnable struct {
	h          uintptr
	isSubclass bool
}

func (this *QGenericRunnable) Run() {
	QGenericRunnable_Run(this.h)
}
