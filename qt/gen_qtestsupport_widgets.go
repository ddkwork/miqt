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
	g := newQTest__QTouchEventWidgetSequence(QTest__QTouchEventWidgetSequence_new(param1))
	g.isSubclass = true
	return g
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
