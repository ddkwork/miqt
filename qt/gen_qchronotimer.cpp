// +build ignore

#include <QChildEvent>
#include <QChronoTimer>
#include <QEvent>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimerEvent>
#include <qchronotimer.h>
#include "gen_qchronotimer.h"

#ifndef _Bool
#define _Bool bool
#endif

class MiqtVirtualQChronoTimer : public virtual QChronoTimer {
public:

	MiqtVirtualQChronoTimer(): QChronoTimer() {};
	MiqtVirtualQChronoTimer(QObject* parent): QChronoTimer(parent) {};

	virtual ~MiqtVirtualQChronoTimer() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__TimerEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void timerEvent(QTimerEvent* param1) override {
		if (handle__TimerEvent == 0) {
			QChronoTimer::timerEvent(param1);
			return;
		}
		
		QTimerEvent* sigval1 = param1;

		miqt_exec_callback_QChronoTimer_TimerEvent(this, handle__TimerEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_TimerEvent(QTimerEvent* param1) {

		QChronoTimer::timerEvent(param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Event = 0;

	// Subclass to allow providing a Go implementation
	virtual bool event(QEvent* event) override {
		if (handle__Event == 0) {
			return QChronoTimer::event(event);
		}
		
		QEvent* sigval1 = event;

		bool callback_return_value = miqt_exec_callback_QChronoTimer_Event(this, handle__Event, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_Event(QEvent* event) {

		return QChronoTimer::event(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__EventFilter = 0;

	// Subclass to allow providing a Go implementation
	virtual bool eventFilter(QObject* watched, QEvent* event) override {
		if (handle__EventFilter == 0) {
			return QChronoTimer::eventFilter(watched, event);
		}
		
		QObject* sigval1 = watched;
		QEvent* sigval2 = event;

		bool callback_return_value = miqt_exec_callback_QChronoTimer_EventFilter(this, handle__EventFilter, sigval1, sigval2);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_EventFilter(QObject* watched, QEvent* event) {

		return QChronoTimer::eventFilter(watched, event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ChildEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void childEvent(QChildEvent* event) override {
		if (handle__ChildEvent == 0) {
			QChronoTimer::childEvent(event);
			return;
		}
		
		QChildEvent* sigval1 = event;

		miqt_exec_callback_QChronoTimer_ChildEvent(this, handle__ChildEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ChildEvent(QChildEvent* event) {

		QChronoTimer::childEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__CustomEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void customEvent(QEvent* event) override {
		if (handle__CustomEvent == 0) {
			QChronoTimer::customEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QChronoTimer_CustomEvent(this, handle__CustomEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_CustomEvent(QEvent* event) {

		QChronoTimer::customEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ConnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void connectNotify(const QMetaMethod& signal) override {
		if (handle__ConnectNotify == 0) {
			QChronoTimer::connectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QChronoTimer_ConnectNotify(this, handle__ConnectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ConnectNotify(QMetaMethod* signal) {

		QChronoTimer::connectNotify(*signal);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DisconnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void disconnectNotify(const QMetaMethod& signal) override {
		if (handle__DisconnectNotify == 0) {
			QChronoTimer::disconnectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QChronoTimer_DisconnectNotify(this, handle__DisconnectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DisconnectNotify(QMetaMethod* signal) {

		QChronoTimer::disconnectNotify(*signal);

	}

};

QChronoTimer* QChronoTimer_new() {
	return new MiqtVirtualQChronoTimer();
}

QChronoTimer* QChronoTimer_new2(QObject* parent) {
	return new MiqtVirtualQChronoTimer(parent);
}

void QChronoTimer_virtbase(QChronoTimer* src, QObject** outptr_QObject) {
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

void QChronoTimer_override_virtual_TimerEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQChronoTimer*>( (QChronoTimer*)(self) )->handle__TimerEvent = slot;
}

void QChronoTimer_virtualbase_TimerEvent(void* self, QTimerEvent* param1) {
	( (MiqtVirtualQChronoTimer*)(self) )->virtualbase_TimerEvent(param1);
}

void QChronoTimer_override_virtual_Event(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQChronoTimer*>( (QChronoTimer*)(self) )->handle__Event = slot;
}

bool QChronoTimer_virtualbase_Event(void* self, QEvent* event) {
	return ( (MiqtVirtualQChronoTimer*)(self) )->virtualbase_Event(event);
}

void QChronoTimer_override_virtual_EventFilter(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQChronoTimer*>( (QChronoTimer*)(self) )->handle__EventFilter = slot;
}

bool QChronoTimer_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event) {
	return ( (MiqtVirtualQChronoTimer*)(self) )->virtualbase_EventFilter(watched, event);
}

void QChronoTimer_override_virtual_ChildEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQChronoTimer*>( (QChronoTimer*)(self) )->handle__ChildEvent = slot;
}

void QChronoTimer_virtualbase_ChildEvent(void* self, QChildEvent* event) {
	( (MiqtVirtualQChronoTimer*)(self) )->virtualbase_ChildEvent(event);
}

void QChronoTimer_override_virtual_CustomEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQChronoTimer*>( (QChronoTimer*)(self) )->handle__CustomEvent = slot;
}

void QChronoTimer_virtualbase_CustomEvent(void* self, QEvent* event) {
	( (MiqtVirtualQChronoTimer*)(self) )->virtualbase_CustomEvent(event);
}

void QChronoTimer_override_virtual_ConnectNotify(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQChronoTimer*>( (QChronoTimer*)(self) )->handle__ConnectNotify = slot;
}

void QChronoTimer_virtualbase_ConnectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQChronoTimer*)(self) )->virtualbase_ConnectNotify(signal);
}

void QChronoTimer_override_virtual_DisconnectNotify(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQChronoTimer*>( (QChronoTimer*)(self) )->handle__DisconnectNotify = slot;
}

void QChronoTimer_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQChronoTimer*)(self) )->virtualbase_DisconnectNotify(signal);
}

void QChronoTimer_Delete(QChronoTimer* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQChronoTimer*>( self );
	} else {
		delete self;
	}
}

