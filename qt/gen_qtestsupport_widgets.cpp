// +build ignore

#include <QPoint>
#define WORKAROUND_INNER_CLASS_DEFINITION_QTest__QTouchEventWidgetSequence
#include <QWidget>
#include <qtestsupport_widgets.h>
#include "gen_qtestsupport_widgets.h"
QTest__QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_new(const QTouchEventWidgetSequence* param1) {
return new QTest::QTouchEventWidgetSequence(*param1);
}
QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Press(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt) {
	return &self->press(static_cast<int>(touchId), *pt);
}
QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Move(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt) {
	return &self->move(static_cast<int>(touchId), *pt);
}
QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Release(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt) {
	return &self->release(static_cast<int>(touchId), *pt);
}
QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Stationary(QTest__QTouchEventWidgetSequence* self, int touchId) {
	return &self->stationary(static_cast<int>(touchId));
}
bool QTest__QTouchEventWidgetSequence_Commit(QTest__QTouchEventWidgetSequence* self, bool processEvents) {
	return self->commit(processEvents);
}
QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Press3(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt, QWidget* widget) {
	return &self->press(static_cast<int>(touchId), *pt, widget);
}
QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Move3(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt, QWidget* widget) {
	return &self->move(static_cast<int>(touchId), *pt, widget);
}
QTouchEventWidgetSequence* QTest__QTouchEventWidgetSequence_Release3(QTest__QTouchEventWidgetSequence* self, int touchId, QPoint* pt, QWidget* widget) {
	return &self->release(static_cast<int>(touchId), *pt, widget);
}
void QTest__QTouchEventWidgetSequence_Delete(QTest__QTouchEventWidgetSequence* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QTest::QTouchEventWidgetSequence*>( self );
	} else {
		delete self;
	}
}
