// +build ignore

#include <QChronoTimer>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimerEvent>
#include <qchronotimer.h>
#include "gen_qchronotimer.h"
class MiqtVirtualQChronoTimer : public virtual QChronoTimer {
public:
MiqtVirtualQChronoTimer(): QChronoTimer() {};
MiqtVirtualQChronoTimer(QObject* parent): QChronoTimer(parent) {};

virtual ~MiqtVirtualQChronoTimer() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QChronoTimer::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QChronoTimer_MetaObject(const_cast<MiqtVirtualQChronoTimer*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QChronoTimer::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QChronoTimer::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QChronoTimer_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QChronoTimer::qt_metacast(param1);

	}
};
QChronoTimer* QChronoTimer_new() {
return new MiqtVirtualQChronoTimer();
}
QChronoTimer* QChronoTimer_new2(QObject* parent) {
return new MiqtVirtualQChronoTimer(parent);
}
void QChronoTimer_virtbase(QChronoTimer* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QChronoTimer_MetaObject(const QChronoTimer* self) {
	return (QMetaObject*) self->metaObject();
}
void* QChronoTimer_Metacast(QChronoTimer* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QChronoTimer_Tr(const char* s) {
	QString _ret = QChronoTimer::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QChronoTimer_IsActive(const QChronoTimer* self) {
	return self->isActive();
}
int QChronoTimer_Id(const QChronoTimer* self) {
	Qt::TimerId _ret = self->id();
	return static_cast<int>(_ret);
}
void QChronoTimer_SetTimerType(QChronoTimer* self, int atype) {
	self->setTimerType(static_cast<Qt::TimerType>(atype));
}
int QChronoTimer_TimerType(const QChronoTimer* self) {
	Qt::TimerType _ret = self->timerType();
	return static_cast<int>(_ret);
}
void QChronoTimer_SetSingleShot(QChronoTimer* self, bool singleShot) {
	self->setSingleShot(singleShot);
}
bool QChronoTimer_IsSingleShot(const QChronoTimer* self) {
	return self->isSingleShot();
}
void QChronoTimer_Start(QChronoTimer* self) {
	self->start();
}
void QChronoTimer_Stop(QChronoTimer* self) {
	self->stop();
}
struct miqt_string QChronoTimer_Tr2(const char* s, const char* c) {
	QString _ret = QChronoTimer::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QChronoTimer_Tr3(const char* s, const char* c, int n) {
	QString _ret = QChronoTimer::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QChronoTimer_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQChronoTimer*>( (QChronoTimer*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QChronoTimer_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQChronoTimer*)(self) )->virtualbase_MetaObject();
}
void QChronoTimer_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQChronoTimer*>( (QChronoTimer*)(self) )->handle__Metacast = slot;
}
void* QChronoTimer_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQChronoTimer*)(self) )->virtualbase_Metacast(param1);
}
void QChronoTimer_Delete(QChronoTimer* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQChronoTimer*>( self );
	} else {
		delete self;
	}
}
