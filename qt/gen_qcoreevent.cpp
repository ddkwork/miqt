// +build ignore

#include <QBasicTimer>
#include <QByteArray>
#include <QChildEvent>
#include <QDynamicPropertyChangeEvent>
#include <QEvent>
#include <QObject>
#include <QTimerEvent>
#include <qcoreevent.h>
#include "gen_qcoreevent.h"

QEvent* QEvent_new(Type typeVal) {
	return new QEvent(typeVal);
}

Type QEvent_Type(const QEvent* self) {
	return self->type();
}

bool QEvent_Spontaneous(const QEvent* self) {
	return self->spontaneous();
}

void QEvent_SetAccepted(QEvent* self, bool accepted) {
	self->setAccepted(accepted);
}

bool QEvent_IsAccepted(const QEvent* self) {
	return self->isAccepted();
}

void QEvent_Accept(QEvent* self) {
	self->accept();
}

void QEvent_Ignore(QEvent* self) {
	self->ignore();
}

bool QEvent_IsInputEvent(const QEvent* self) {
	return self->isInputEvent();
}

bool QEvent_IsPointerEvent(const QEvent* self) {
	return self->isPointerEvent();
}

bool QEvent_IsSinglePointEvent(const QEvent* self) {
	return self->isSinglePointEvent();
}

int QEvent_RegisterEventType() {
	return QEvent::registerEventType();
}

QEvent* QEvent_Clone(const QEvent* self) {
	return self->clone();
}

int QEvent_RegisterEventType1(int hint) {
	return QEvent::registerEventType(static_cast<int>(hint));
}

void QEvent_Delete(QEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QEvent*>( self );
	} else {
		delete self;
	}
}

QTimerEvent* QTimerEvent_new(int timerId) {
	return new QTimerEvent(static_cast<int>(timerId));
}

QTimerEvent* QTimerEvent_new2(int timerId) {
	return new QTimerEvent(static_cast<Qt::TimerId>(timerId));
}

void QTimerEvent_virtbase(QTimerEvent* src, QEvent** outptr_QEvent) {
	*outptr_QEvent = static_cast<QEvent*>(src);
}

QTimerEvent* QTimerEvent_Clone(const QTimerEvent* self) {
	return self->clone();
}

int QTimerEvent_TimerId(const QTimerEvent* self) {
	return self->timerId();
}

int QTimerEvent_Id(const QTimerEvent* self) {
	Qt::TimerId _ret = self->id();
	return static_cast<int>(_ret);
}

bool QTimerEvent_Matches(const QTimerEvent* self, QBasicTimer* timer) {
	return self->matches(*timer);
}

void QTimerEvent_Delete(QTimerEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QTimerEvent*>( self );
	} else {
		delete self;
	}
}

QChildEvent* QChildEvent_new(Type typeVal, QObject* child) {
	return new QChildEvent(typeVal, child);
}

void QChildEvent_virtbase(QChildEvent* src, QEvent** outptr_QEvent) {
	*outptr_QEvent = static_cast<QEvent*>(src);
}

QChildEvent* QChildEvent_Clone(const QChildEvent* self) {
	return self->clone();
}

QObject* QChildEvent_Child(const QChildEvent* self) {
	return self->child();
}

bool QChildEvent_Added(const QChildEvent* self) {
	return self->added();
}

bool QChildEvent_Polished(const QChildEvent* self) {
	return self->polished();
}

bool QChildEvent_Removed(const QChildEvent* self) {
	return self->removed();
}

void QChildEvent_Delete(QChildEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QChildEvent*>( self );
	} else {
		delete self;
	}
}

QDynamicPropertyChangeEvent* QDynamicPropertyChangeEvent_new(struct miqt_string name) {
	QByteArray name_QByteArray(name.data, name.len);
	return new QDynamicPropertyChangeEvent(name_QByteArray);
}

void QDynamicPropertyChangeEvent_virtbase(QDynamicPropertyChangeEvent* src, QEvent** outptr_QEvent) {
	*outptr_QEvent = static_cast<QEvent*>(src);
}

QDynamicPropertyChangeEvent* QDynamicPropertyChangeEvent_Clone(const QDynamicPropertyChangeEvent* self) {
	return self->clone();
}

struct miqt_string QDynamicPropertyChangeEvent_PropertyName(const QDynamicPropertyChangeEvent* self) {
	QByteArray _qb = self->propertyName();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

void QDynamicPropertyChangeEvent_Delete(QDynamicPropertyChangeEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QDynamicPropertyChangeEvent*>( self );
	} else {
		delete self;
	}
}

