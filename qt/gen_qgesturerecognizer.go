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
