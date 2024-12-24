// +build ignore

#include <QBasicTimer>
#include <QObject>
#include <qbasictimer.h>
#include "gen_qbasictimer.h"

#ifndef _Bool
#define _Bool bool
#endif

QBasicTimer* QBasicTimer_new() {
	return new QBasicTimer();
}

void QBasicTimer_Swap(QBasicTimer* self, QBasicTimer* other) {
	self->swap(*other);
}

bool QBasicTimer_IsActive(const QBasicTimer* self) {
	return self->isActive();
}

int QBasicTimer_TimerId(const QBasicTimer* self) {
	return self->timerId();
}

int QBasicTimer_Id(const QBasicTimer* self) {
	Qt::TimerId _ret = self->id();
	return static_cast<int>(_ret);
}

void QBasicTimer_Start(QBasicTimer* self, int msec, QObject* obj) {
	self->start(static_cast<int>(msec), obj);
}

void QBasicTimer_Start2(QBasicTimer* self, int msec, int timerType, QObject* obj) {
	self->start(static_cast<int>(msec), static_cast<Qt::TimerType>(timerType), obj);
}

void QBasicTimer_Start3(QBasicTimer* self, Duration duration, QObject* obj) {
	self->start(duration, obj);
}

void QBasicTimer_Start4(QBasicTimer* self, Duration duration, int timerType, QObject* obj) {
	self->start(duration, static_cast<Qt::TimerType>(timerType), obj);
}

void QBasicTimer_Stop(QBasicTimer* self) {
	self->stop();
}

void QBasicTimer_Delete(QBasicTimer* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QBasicTimer*>( self );
	} else {
		delete self;
	}
}

