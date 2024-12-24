// +build ignore

#include <QChildEvent>
#include <QEvent>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimerEvent>
#include <QWinEventNotifier>
#include <qwineventnotifier.h>
#include "gen_qwineventnotifier.h"

#ifndef _Bool
#define _Bool bool
#endif

void _GUID_Delete(_GUID* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<_GUID*>( self );
	} else {
		delete self;
	}
}

void type_info_Delete(type_info* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<type_info*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQWinEventNotifier : public virtual QWinEventNotifier {
public:

	MiqtVirtualQWinEventNotifier(): QWinEventNotifier() {};
	MiqtVirtualQWinEventNotifier(HANDLE hEvent): QWinEventNotifier(hEvent) {};
	MiqtVirtualQWinEventNotifier(QObject* parent): QWinEventNotifier(parent) {};
	MiqtVirtualQWinEventNotifier(HANDLE hEvent, QObject* parent): QWinEventNotifier(hEvent, parent) {};

	virtual ~MiqtVirtualQWinEventNotifier() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Event = 0;

	// Subclass to allow providing a Go implementation
	virtual bool event(QEvent* e) override {
		if (handle__Event == 0) {
			return QWinEventNotifier::event(e);
		}
		
		QEvent* sigval1 = e;

		bool callback_return_value = miqt_exec_callback_QWinEventNotifier_Event(this, handle__Event, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_Event(QEvent* e) {

		return QWinEventNotifier::event(e);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__EventFilter = 0;

	// Subclass to allow providing a Go implementation
	virtual bool eventFilter(QObject* watched, QEvent* event) override {
		if (handle__EventFilter == 0) {
			return QWinEventNotifier::eventFilter(watched, event);
		}
		
		QObject* sigval1 = watched;
		QEvent* sigval2 = event;

		bool callback_return_value = miqt_exec_callback_QWinEventNotifier_EventFilter(this, handle__EventFilter, sigval1, sigval2);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_EventFilter(QObject* watched, QEvent* event) {

		return QWinEventNotifier::eventFilter(watched, event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__TimerEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void timerEvent(QTimerEvent* event) override {
		if (handle__TimerEvent == 0) {
			QWinEventNotifier::timerEvent(event);
			return;
		}
		
		QTimerEvent* sigval1 = event;

		miqt_exec_callback_QWinEventNotifier_TimerEvent(this, handle__TimerEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_TimerEvent(QTimerEvent* event) {

		QWinEventNotifier::timerEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ChildEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void childEvent(QChildEvent* event) override {
		if (handle__ChildEvent == 0) {
			QWinEventNotifier::childEvent(event);
			return;
		}
		
		QChildEvent* sigval1 = event;

		miqt_exec_callback_QWinEventNotifier_ChildEvent(this, handle__ChildEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ChildEvent(QChildEvent* event) {

		QWinEventNotifier::childEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__CustomEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void customEvent(QEvent* event) override {
		if (handle__CustomEvent == 0) {
			QWinEventNotifier::customEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QWinEventNotifier_CustomEvent(this, handle__CustomEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_CustomEvent(QEvent* event) {

		QWinEventNotifier::customEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ConnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void connectNotify(const QMetaMethod& signal) override {
		if (handle__ConnectNotify == 0) {
			QWinEventNotifier::connectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QWinEventNotifier_ConnectNotify(this, handle__ConnectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ConnectNotify(QMetaMethod* signal) {

		QWinEventNotifier::connectNotify(*signal);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DisconnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void disconnectNotify(const QMetaMethod& signal) override {
		if (handle__DisconnectNotify == 0) {
			QWinEventNotifier::disconnectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QWinEventNotifier_DisconnectNotify(this, handle__DisconnectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DisconnectNotify(QMetaMethod* signal) {

		QWinEventNotifier::disconnectNotify(*signal);

	}

};

QWinEventNotifier* QWinEventNotifier_new() {
	return new MiqtVirtualQWinEventNotifier();
}

QWinEventNotifier* QWinEventNotifier_new2(HANDLE hEvent) {
	return new MiqtVirtualQWinEventNotifier(hEvent);
}

QWinEventNotifier* QWinEventNotifier_new3(QObject* parent) {
	return new MiqtVirtualQWinEventNotifier(parent);
}

QWinEventNotifier* QWinEventNotifier_new4(HANDLE hEvent, QObject* parent) {
	return new MiqtVirtualQWinEventNotifier(hEvent, parent);
}

void QWinEventNotifier_virtbase(QWinEventNotifier* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QWinEventNotifier_MetaObject(const QWinEventNotifier* self) {
	return (QMetaObject*) self->metaObject();
}

void* QWinEventNotifier_Metacast(QWinEventNotifier* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QWinEventNotifier_Tr(const char* s) {
	QString _ret = QWinEventNotifier::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QWinEventNotifier_SetHandle(QWinEventNotifier* self, HANDLE hEvent) {
	self->setHandle(hEvent);
}

HANDLE QWinEventNotifier_Handle(const QWinEventNotifier* self) {
	return self->handle();
}

bool QWinEventNotifier_IsEnabled(const QWinEventNotifier* self) {
	return self->isEnabled();
}

void QWinEventNotifier_SetEnabled(QWinEventNotifier* self, bool enable) {
	self->setEnabled(enable);
}

struct miqt_string QWinEventNotifier_Tr2(const char* s, const char* c) {
	QString _ret = QWinEventNotifier::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QWinEventNotifier_Tr3(const char* s, const char* c, int n) {
	QString _ret = QWinEventNotifier::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QWinEventNotifier_override_virtual_Event(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWinEventNotifier*>( (QWinEventNotifier*)(self) )->handle__Event = slot;
}

bool QWinEventNotifier_virtualbase_Event(void* self, QEvent* e) {
	return ( (MiqtVirtualQWinEventNotifier*)(self) )->virtualbase_Event(e);
}

void QWinEventNotifier_override_virtual_EventFilter(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWinEventNotifier*>( (QWinEventNotifier*)(self) )->handle__EventFilter = slot;
}

bool QWinEventNotifier_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event) {
	return ( (MiqtVirtualQWinEventNotifier*)(self) )->virtualbase_EventFilter(watched, event);
}

void QWinEventNotifier_override_virtual_TimerEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWinEventNotifier*>( (QWinEventNotifier*)(self) )->handle__TimerEvent = slot;
}

void QWinEventNotifier_virtualbase_TimerEvent(void* self, QTimerEvent* event) {
	( (MiqtVirtualQWinEventNotifier*)(self) )->virtualbase_TimerEvent(event);
}

void QWinEventNotifier_override_virtual_ChildEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWinEventNotifier*>( (QWinEventNotifier*)(self) )->handle__ChildEvent = slot;
}

void QWinEventNotifier_virtualbase_ChildEvent(void* self, QChildEvent* event) {
	( (MiqtVirtualQWinEventNotifier*)(self) )->virtualbase_ChildEvent(event);
}

void QWinEventNotifier_override_virtual_CustomEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWinEventNotifier*>( (QWinEventNotifier*)(self) )->handle__CustomEvent = slot;
}

void QWinEventNotifier_virtualbase_CustomEvent(void* self, QEvent* event) {
	( (MiqtVirtualQWinEventNotifier*)(self) )->virtualbase_CustomEvent(event);
}

void QWinEventNotifier_override_virtual_ConnectNotify(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWinEventNotifier*>( (QWinEventNotifier*)(self) )->handle__ConnectNotify = slot;
}

void QWinEventNotifier_virtualbase_ConnectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQWinEventNotifier*)(self) )->virtualbase_ConnectNotify(signal);
}

void QWinEventNotifier_override_virtual_DisconnectNotify(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWinEventNotifier*>( (QWinEventNotifier*)(self) )->handle__DisconnectNotify = slot;
}

void QWinEventNotifier_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQWinEventNotifier*)(self) )->virtualbase_DisconnectNotify(signal);
}

void QWinEventNotifier_Delete(QWinEventNotifier* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQWinEventNotifier*>( self );
	} else {
		delete self;
	}
}

