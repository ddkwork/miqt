package multimedia

import (
	"unsafe"
)

type QCapturableWindow struct {
	h          uintptr
	isSubclass bool
}

// NewQCapturableWindow constructs a new QCapturableWindow object.
func NewQCapturableWindow() *QCapturableWindow {
	g := newQCapturableWindow(QCapturableWindow_new())
	g.isSubclass = true
	return g
}

// NewQCapturableWindow2 constructs a new QCapturableWindow object.
func NewQCapturableWindow2(other *QCapturableWindow) *QCapturableWindow {
	g := newQCapturableWindow(QCapturableWindow_new2(other.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QCapturableWindow) OperatorAssign(other *QCapturableWindow) {
	QCapturableWindow_OperatorAssign(this.h, other.cPointer())
}

func (this *QCapturableWindow) Swap(other *QCapturableWindow) {
	QCapturableWindow_Swap(this.h, other.cPointer())
}

func (this *QCapturableWindow) IsValid() bool {
	return (bool)(QCapturableWindow_IsValid(this.h))
}

func (this *QCapturableWindow) Description() string {
	var _ms struct_miqt_string = QCapturableWindow_Description(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}