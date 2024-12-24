// +build ignore

#include <QEvent>
#include <QGesture>
#include <QGestureRecognizer>
#include <QObject>
#include <qgesturerecognizer.h>
#include "gen_qgesturerecognizer.h"

QGestureRecognizer* QGestureRecognizer_new() {
	return new QGestureRecognizer();
}

QGesture* QGestureRecognizer_Create(QGestureRecognizer* self, QObject* target) {
	return self->create(target);
}

Result QGestureRecognizer_Recognize(QGestureRecognizer* self, QGesture* state, QObject* watched, QEvent* event) {
	return self->recognize(state, watched, event);
}

void QGestureRecognizer_Reset(QGestureRecognizer* self, QGesture* state) {
	self->reset(state);
}

int QGestureRecognizer_RegisterRecognizer(QGestureRecognizer* recognizer) {
	Qt::GestureType _ret = QGestureRecognizer::registerRecognizer(recognizer);
	return static_cast<int>(_ret);
}

void QGestureRecognizer_UnregisterRecognizer(int typeVal) {
	QGestureRecognizer::unregisterRecognizer(static_cast<Qt::GestureType>(typeVal));
}

void QGestureRecognizer_OperatorAssign(QGestureRecognizer* self, QGestureRecognizer* param1) {
	self->operator=(*param1);
}

void QGestureRecognizer_Delete(QGestureRecognizer* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QGestureRecognizer*>( self );
	} else {
		delete self;
	}
}

