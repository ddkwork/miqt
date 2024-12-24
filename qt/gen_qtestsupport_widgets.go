package qt

import (
	"unsafe"
)

type QTest__QTouchEventWidgetSequence struct {
	h          uintptr
	isSubclass bool
}

// NewQTest__QTouchEventWidgetSequence constructs a new QTest::QTouchEventWidgetSequence object.
func NewQTest__QTouchEventWidgetSequence(param1 *QTouchEventWidgetSequence) *QTest__QTouchEventWidgetSequence {

	ret := newQTest__QTouchEventWidgetSequence(QTest__QTouchEventWidgetSequence_new(param1))
	ret.isSubclass = true
	return ret
}

func (this *QTest__QTouchEventWidgetSequence) Press(touchId int, pt *QPoint) *QTouchEventWidgetSequence {
	xxxxxxxxx
}

func (this *QTest__QTouchEventWidgetSequence) Move(touchId int, pt *QPoint) *QTouchEventWidgetSequence {
	xxxxxxxxx
}

func (this *QTest__QTouchEventWidgetSequence) Release(touchId int, pt *QPoint) *QTouchEventWidgetSequence {
	xxxxxxxxx
}

func (this *QTest__QTouchEventWidgetSequence) Stationary(touchId int) *QTouchEventWidgetSequence {
	xxxxxxxxx
}

func (this *QTest__QTouchEventWidgetSequence) Commit(processEvents bool) bool {
	return (bool)(QTest__QTouchEventWidgetSequence_Commit(this.h, (bool)(processEvents)))
}

func (this *QTest__QTouchEventWidgetSequence) Press3(touchId int, pt *QPoint, widget *QWidget) *QTouchEventWidgetSequence {
	xxxxxxxxx
}

func (this *QTest__QTouchEventWidgetSequence) Move3(touchId int, pt *QPoint, widget *QWidget) *QTouchEventWidgetSequence {
	xxxxxxxxx
}

func (this *QTest__QTouchEventWidgetSequence) Release3(touchId int, pt *QPoint, widget *QWidget) *QTouchEventWidgetSequence {
	xxxxxxxxx
}

func (this *QTest__QTouchEventWidgetSequence) callVirtualBase_Stationary(touchId int) *QTouchEventWidgetSequence {

	xxxxxxxxx
}
func (this *QTest__QTouchEventWidgetSequence) OnStationary(slot func(super func(touchId int) *QTouchEventWidgetSequence, touchId int) *QTouchEventWidgetSequence) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTest__QTouchEventWidgetSequence_override_virtual_Stationary(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTest__QTouchEventWidgetSequence_Stationary
func miqt_exec_callback_QTest__QTouchEventWidgetSequence_Stationary(self QTest__QTouchEventWidgetSequence, cb intptr_t, touchId int) *QTouchEventWidgetSequence {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(touchId int) *QTouchEventWidgetSequence, touchId int) *QTouchEventWidgetSequence)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(touchId)

	virtualReturn := gofunc((&QTest__QTouchEventWidgetSequence{h: self}).callVirtualBase_Stationary, slotval1)

	return virtualReturn

}

func (this *QTest__QTouchEventWidgetSequence) callVirtualBase_Commit(processEvents bool) bool {

	return (bool)(QTest__QTouchEventWidgetSequence_virtualbase_Commit(unsafe.Pointer(this.h), (bool)(processEvents)))

}
func (this *QTest__QTouchEventWidgetSequence) OnCommit(slot func(super func(processEvents bool) bool, processEvents bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTest__QTouchEventWidgetSequence_override_virtual_Commit(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTest__QTouchEventWidgetSequence_Commit
func miqt_exec_callback_QTest__QTouchEventWidgetSequence_Commit(self QTest__QTouchEventWidgetSequence, cb intptr_t, processEvents bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(processEvents bool) bool, processEvents bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(processEvents)

	virtualReturn := gofunc((&QTest__QTouchEventWidgetSequence{h: self}).callVirtualBase_Commit, slotval1)

	return (bool)(virtualReturn)

}
