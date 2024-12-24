package qt

import (
	"unsafe"
)

type QTest__QTouchEventSequence struct {
	h          uintptr
	isSubclass bool
}

func (this *QTest__QTouchEventSequence) Press(touchId int, pt *QPoint) *QTouchEventSequence {
	xxxxxxxxx
}

func (this *QTest__QTouchEventSequence) Move(touchId int, pt *QPoint) *QTouchEventSequence {
	xxxxxxxxx
}

func (this *QTest__QTouchEventSequence) Release(touchId int, pt *QPoint) *QTouchEventSequence {
	xxxxxxxxx
}

func (this *QTest__QTouchEventSequence) Stationary(touchId int) *QTouchEventSequence {
	xxxxxxxxx
}

func (this *QTest__QTouchEventSequence) Commit(processEvents bool) bool {
	return (bool)(QTest__QTouchEventSequence_Commit(this.h, (bool)(processEvents)))
}

func (this *QTest__QTouchEventSequence) Press3(touchId int, pt *QPoint, window *QWindow) *QTouchEventSequence {
	xxxxxxxxx
}

func (this *QTest__QTouchEventSequence) Move3(touchId int, pt *QPoint, window *QWindow) *QTouchEventSequence {
	xxxxxxxxx
}

func (this *QTest__QTouchEventSequence) Release3(touchId int, pt *QPoint, window *QWindow) *QTouchEventSequence {
	xxxxxxxxx
}
