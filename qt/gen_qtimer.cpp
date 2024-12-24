// +build ignore

#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimer>
#include <QTimerEvent>
#include <qtimer.h>
#include "gen_qtimer.h"
class MiqtVirtualQTimer : public virtual QTimer {
public:
MiqtVirtualQTimer(): QTimer() {};
MiqtVirtualQTimer(QObject* parent): QTimer(parent) {};

virtual ~MiqtVirtualQTimer() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QTimer::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QTimer_MetaObject(const_cast<MiqtVirtualQTimer*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QTimer::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QTimer::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QTimer_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QTimer::qt_metacast(param1);

	}
};
QTimer* QTimer_new() {
return new MiqtVirtualQTimer();
}
QTimer* QTimer_new2(QObject* parent) {
return new MiqtVirtualQTimer(parent);
}
void QTimer_virtbase(QTimer* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QTimer_MetaObject(const QTimer* self) {
	return (QMetaObject*) self->metaObject();
}
void* QTimer_Metacast(QTimer* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QTimer_Tr(const char* s) {
	QString _ret = QTimer::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QTimer_IsActive(const QTimer* self) {
	return self->isActive();
}
int QTimer_TimerId(const QTimer* self) {
	return self->timerId();
}
int QTimer_Id(const QTimer* self) {
	Qt::TimerId _ret = self->id();
	return static_cast<int>(_ret);
}
void QTimer_SetInterval(QTimer* self, int msec) {
	self->setInterval(static_cast<int>(msec));
}
int QTimer_Interval(const QTimer* self) {
	return self->interval();
}
int QTimer_RemainingTime(const QTimer* self) {
	return self->remainingTime();
}
void QTimer_SetTimerType(QTimer* self, int atype) {
	self->setTimerType(static_cast<Qt::TimerType>(atype));
}
int QTimer_TimerType(const QTimer* self) {
	Qt::TimerType _ret = self->timerType();
	return static_cast<int>(_ret);
}
void QTimer_SetSingleShot(QTimer* self, bool singleShot) {
	self->setSingleShot(singleShot);
}
bool QTimer_IsSingleShot(const QTimer* self) {
	return self->isSingleShot();
}
void QTimer_Start(QTimer* self, int msec) {
	self->start(static_cast<int>(msec));
}
void QTimer_Start2(QTimer* self) {
	self->start();
}
void QTimer_Stop(QTimer* self) {
	self->stop();
}
struct miqt_string QTimer_Tr2(const char* s, const char* c) {
	QString _ret = QTimer::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QTimer_Tr3(const char* s, const char* c, int n) {
	QString _ret = QTimer::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QTimer_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTimer*>( (QTimer*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QTimer_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQTimer*)(self) )->virtualbase_MetaObject();
}
void QTimer_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTimer*>( (QTimer*)(self) )->handle__Metacast = slot;
}
void* QTimer_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQTimer*)(self) )->virtualbase_Metacast(param1);
}
void QTimer_Delete(QTimer* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQTimer*>( self );
	} else {
		delete self;
	}
}
