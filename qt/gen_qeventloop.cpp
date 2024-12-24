// +build ignore

#include <QDeadlineTimer>
#include <QEvent>
#include <QEventLoop>
#include <QEventLoopLocker>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QThread>
#include <qeventloop.h>
#include "gen_qeventloop.h"

class MiqtVirtualQEventLoop : public virtual QEventLoop {
public:

	MiqtVirtualQEventLoop(): QEventLoop() {};
	MiqtVirtualQEventLoop(QObject* parent): QEventLoop(parent) {};

	virtual ~MiqtVirtualQEventLoop() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QEventLoop::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QEventLoop_MetaObject(const_cast<MiqtVirtualQEventLoop*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QEventLoop::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QEventLoop::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QEventLoop_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QEventLoop::qt_metacast(param1);

	}

};

QEventLoop* QEventLoop_new() {
	return new MiqtVirtualQEventLoop();
}

QEventLoop* QEventLoop_new2(QObject* parent) {
	return new MiqtVirtualQEventLoop(parent);
}

void QEventLoop_virtbase(QEventLoop* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QEventLoop_MetaObject(const QEventLoop* self) {
	return (QMetaObject*) self->metaObject();
}

void* QEventLoop_Metacast(QEventLoop* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QEventLoop_Tr(const char* s) {
	QString _ret = QEventLoop::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QEventLoop_ProcessEvents(QEventLoop* self) {
	return self->processEvents();
}

void QEventLoop_ProcessEvents2(QEventLoop* self, ProcessEventsFlags flags, int maximumTime) {
	self->processEvents(flags, static_cast<int>(maximumTime));
}

void QEventLoop_ProcessEvents3(QEventLoop* self, ProcessEventsFlags flags, QDeadlineTimer* deadline) {
	self->processEvents(flags, *deadline);
}

int QEventLoop_Exec(QEventLoop* self) {
	return self->exec();
}

bool QEventLoop_IsRunning(const QEventLoop* self) {
	return self->isRunning();
}

void QEventLoop_WakeUp(QEventLoop* self) {
	self->wakeUp();
}

bool QEventLoop_Event(QEventLoop* self, QEvent* event) {
	return self->event(event);
}

void QEventLoop_Exit(QEventLoop* self) {
	self->exit();
}

void QEventLoop_Quit(QEventLoop* self) {
	self->quit();
}

struct miqt_string QEventLoop_Tr2(const char* s, const char* c) {
	QString _ret = QEventLoop::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QEventLoop_Tr3(const char* s, const char* c, int n) {
	QString _ret = QEventLoop::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QEventLoop_ProcessEvents1(QEventLoop* self, ProcessEventsFlags flags) {
	return self->processEvents(flags);
}

int QEventLoop_Exec1(QEventLoop* self, ProcessEventsFlags flags) {
	return self->exec(flags);
}

void QEventLoop_Exit1(QEventLoop* self, int returnCode) {
	self->exit(static_cast<int>(returnCode));
}

void QEventLoop_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQEventLoop*>( (QEventLoop*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QEventLoop_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQEventLoop*)(self) )->virtualbase_MetaObject();
}

void QEventLoop_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQEventLoop*>( (QEventLoop*)(self) )->handle__Metacast = slot;
}

void* QEventLoop_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQEventLoop*)(self) )->virtualbase_Metacast(param1);
}

void QEventLoop_Delete(QEventLoop* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQEventLoop*>( self );
	} else {
		delete self;
	}
}

QEventLoopLocker* QEventLoopLocker_new() {
	return new QEventLoopLocker();
}

QEventLoopLocker* QEventLoopLocker_new2(QEventLoop* loop) {
	return new QEventLoopLocker(loop);
}

QEventLoopLocker* QEventLoopLocker_new3(QThread* thread) {
	return new QEventLoopLocker(thread);
}

void QEventLoopLocker_Swap(QEventLoopLocker* self, QEventLoopLocker* other) {
	self->swap(*other);
}

void QEventLoopLocker_Delete(QEventLoopLocker* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QEventLoopLocker*>( self );
	} else {
		delete self;
	}
}

