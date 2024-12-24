package qt

import (
	"unsafe"
)

type QGestureRecognizer__ResultFlag int

const (
	QGestureRecognizer__Ignore           QGestureRecognizer__ResultFlag = 1
	QGestureRecognizer__MayBeGesture     QGestureRecognizer__ResultFlag = 2
	QGestureRecognizer__TriggerGesture   QGestureRecognizer__ResultFlag = 4
	QGestureRecognizer__FinishGesture    QGestureRecognizer__ResultFlag = 8
	QGestureRecognizer__CancelGesture    QGestureRecognizer__ResultFlag = 16
	QGestureRecognizer__ResultState_Mask QGestureRecognizer__ResultFlag = 255
	QGestureRecognizer__ConsumeEventHint QGestureRecognizer__ResultFlag = 256
	QGestureRecognizer__ResultHint_Mask  QGestureRecognizer__ResultFlag = 65280
)

type QGestureRecognizer struct {
	h          uintptr
	isSubclass bool
}

// NewQGestureRecognizer constructs a new QGestureRecognizer object.
func NewQGestureRecognizer() *QGestureRecognizer {

	ret := newQGestureRecognizer(QGestureRecognizer_new())
	ret.isSubclass = true
	return ret
}

func (this *QGestureRecognizer) Create(target *QObject) *QGesture {
	return newQGesture(QGestureRecognizer_Create(this.h, target.cPointer()))
}

func (this *QGestureRecognizer) Recognize(state *QGesture, watched *QObject, event *QEvent) Result {
	xxxxxxxxx
}

func (this *QGestureRecognizer) Reset(state *QGesture) {
	QGestureRecognizer_Reset(this.h, state.cPointer())
}

func QGestureRecognizer_RegisterRecognizer(recognizer *QGestureRecognizer) GestureType {
	return (GestureType)(QGestureRecognizer_RegisterRecognizer(recognizer.cPointer()))
}

func QGestureRecognizer_UnregisterRecognizer(typeVal GestureType) {
	QGestureRecognizer_UnregisterRecognizer((int)(typeVal))
}

func (this *QGestureRecognizer) OperatorAssign(param1 *QGestureRecognizer) {
	QGestureRecognizer_OperatorAssign(this.h, param1.cPointer())
}

func (this *QGestureRecognizer) callVirtualBase_Create(target *QObject) *QGesture {

	return newQGesture(QGestureRecognizer_virtualbase_Create(unsafe.Pointer(this.h), target.cPointer()))

}
func (this *QGestureRecognizer) OnCreate(slot func(super func(target *QObject) *QGesture, target *QObject) *QGesture) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGestureRecognizer_override_virtual_Create(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGestureRecognizer_Create
func miqt_exec_callback_QGestureRecognizer_Create(self QGestureRecognizer, cb intptr_t, target *QObject) *QGesture {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(target *QObject) *QGesture, target *QObject) *QGesture)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(target)

	virtualReturn := gofunc((&QGestureRecognizer{h: self}).callVirtualBase_Create, slotval1)

	return virtualReturn.cPointer()

}
func (this *QGestureRecognizer) OnRecognize(slot func(state *QGesture, watched *QObject, event *QEvent) Result) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGestureRecognizer_override_virtual_Recognize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGestureRecognizer_Recognize
func miqt_exec_callback_QGestureRecognizer_Recognize(self QGestureRecognizer, cb intptr_t, state *QGesture, watched *QObject, event *QEvent) Result {
	gofunc, ok := cgo.Handle(cb).Value().(func(state *QGesture, watched *QObject, event *QEvent) Result)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGesture(state)

	slotval2 := newQObject(watched)

	slotval3 := newQEvent(event)

	virtualReturn := gofunc(slotval1, slotval2, slotval3)

	return virtualReturn

}

func (this *QGestureRecognizer) callVirtualBase_Reset(state *QGesture) {

	QGestureRecognizer_virtualbase_Reset(unsafe.Pointer(this.h), state.cPointer())

}
func (this *QGestureRecognizer) OnReset(slot func(super func(state *QGesture), state *QGesture)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGestureRecognizer_override_virtual_Reset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGestureRecognizer_Reset
func miqt_exec_callback_QGestureRecognizer_Reset(self QGestureRecognizer, cb intptr_t, state *QGesture) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(state *QGesture), state *QGesture))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGesture(state)

	gofunc((&QGestureRecognizer{h: self}).callVirtualBase_Reset, slotval1)

}
