// +build ignore

#include <QPoint>
#define WORKAROUND_INNER_CLASS_DEFINITION_QTest__QTouchEventSequence
#include <QWindow>
#include <qtestsupport_gui.h>
#include "gen_qtestsupport_gui.h"

QTouchEventSequence* QTest__QTouchEventSequence_Press(QTest__QTouchEventSequence* self, int touchId, QPoint* pt) {
	return &self->press(static_cast<int>(touchId), *pt);
}

QTouchEventSequence* QTest__QTouchEventSequence_Move(QTest__QTouchEventSequence* self, int touchId, QPoint* pt) {
	return &self->move(static_cast<int>(touchId), *pt);
}

QTouchEventSequence* QTest__QTouchEventSequence_Release(QTest__QTouchEventSequence* self, int touchId, QPoint* pt) {
	return &self->release(static_cast<int>(touchId), *pt);
}

QTouchEventSequence* QTest__QTouchEventSequence_Stationary(QTest__QTouchEventSequence* self, int touchId) {
	return &self->stationary(static_cast<int>(touchId));
}

bool QTest__QTouchEventSequence_Commit(QTest__QTouchEventSequence* self, bool processEvents) {
	return self->commit(processEvents);
}

QTouchEventSequence* QTest__QTouchEventSequence_Press3(QTest__QTouchEventSequence* self, int touchId, QPoint* pt, QWindow* window) {
	return &self->press(static_cast<int>(touchId), *pt, window);
}

QTouchEventSequence* QTest__QTouchEventSequence_Move3(QTest__QTouchEventSequence* self, int touchId, QPoint* pt, QWindow* window) {
	return &self->move(static_cast<int>(touchId), *pt, window);
}

QTouchEventSequence* QTest__QTouchEventSequence_Release3(QTest__QTouchEventSequence* self, int touchId, QPoint* pt, QWindow* window) {
	return &self->release(static_cast<int>(touchId), *pt, window);
}

void QTest__QTouchEventSequence_Delete(QTest__QTouchEventSequence* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QTest::QTouchEventSequence*>( self );
	} else {
		delete self;
	}
}

