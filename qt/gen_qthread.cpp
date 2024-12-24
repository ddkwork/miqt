// +build ignore

#include <QAbstractEventDispatcher>
#include <QDeadlineTimer>
#include <QEvent>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QThread>
#include <qthread.h>
#include "gen_qthread.h"
class MiqtVirtualQThread : public virtual QThread {
public:
MiqtVirtualQThread(): QThread() {};
MiqtVirtualQThread(QObject* parent): QThread(parent) {};

virtual ~MiqtVirtualQThread() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QThread::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QThread_MetaObject(const_cast<MiqtVirtualQThread*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QThread::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QThread::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QThread_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QThread::qt_metacast(param1);

	}
};
QThread* QThread_new() {
return new MiqtVirtualQThread();
}
QThread* QThread_new2(QObject* parent) {
return new MiqtVirtualQThread(parent);
}
void QThread_virtbase(QThread* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QThread_MetaObject(const QThread* self) {
	return (QMetaObject*) self->metaObject();
}
void* QThread_Metacast(QThread* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QThread_Tr(const char* s) {
	QString _ret = QThread::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void* QThread_CurrentThreadId() {
	Qt::HANDLE _ret = QThread::currentThreadId();
	return static_cast<void*>(_ret);
}
QThread* QThread_CurrentThread() {
	return QThread::currentThread();
}
bool QThread_IsMainThread() {
	return QThread::isMainThread();
}
int QThread_IdealThreadCount() {
	return QThread::idealThreadCount();
}
void QThread_YieldCurrentThread() {
	QThread::yieldCurrentThread();
}
void QThread_SetPriority(QThread* self, Priority priority) {
	self->setPriority(priority);
}
Priority QThread_Priority(const QThread* self) {
	return self->priority();
}
bool QThread_IsFinished(const QThread* self) {
	return self->isFinished();
}
bool QThread_IsRunning(const QThread* self) {
	return self->isRunning();
}
void QThread_RequestInterruption(QThread* self) {
	self->requestInterruption();
}
bool QThread_IsInterruptionRequested(const QThread* self) {
	return self->isInterruptionRequested();
}
void QThread_SetStackSize(QThread* self, unsigned int stackSize) {
	self->setStackSize(static_cast<uint>(stackSize));
}
unsigned int QThread_StackSize(const QThread* self) {
	uint _ret = self->stackSize();
	return static_cast<unsigned int>(_ret);
}
QAbstractEventDispatcher* QThread_EventDispatcher(const QThread* self) {
	return self->eventDispatcher();
}
void QThread_SetEventDispatcher(QThread* self, QAbstractEventDispatcher* eventDispatcher) {
	self->setEventDispatcher(eventDispatcher);
}
bool QThread_Event(QThread* self, QEvent* event) {
	return self->event(event);
}
int QThread_LoopLevel(const QThread* self) {
	return self->loopLevel();
}
bool QThread_IsCurrentThread(const QThread* self) {
	return self->isCurrentThread();
}
void QThread_SetServiceLevel(QThread* self, QualityOfService serviceLevel) {
	self->setServiceLevel(serviceLevel);
}
QualityOfService QThread_ServiceLevel(const QThread* self) {
	return self->serviceLevel();
}
void QThread_Start(QThread* self) {
	self->start();
}
void QThread_Terminate(QThread* self) {
	self->terminate();
}
void QThread_Exit(QThread* self) {
	self->exit();
}
void QThread_Quit(QThread* self) {
	self->quit();
}
bool QThread_Wait(QThread* self) {
	return self->wait();
}
bool QThread_WaitWithTime(QThread* self, unsigned long time) {
	return self->wait(static_cast<unsigned long>(time));
}
void QThread_Sleep(unsigned long param1) {
	QThread::sleep(static_cast<unsigned long>(param1));
}
void QThread_Msleep(unsigned long param1) {
	QThread::msleep(static_cast<unsigned long>(param1));
}
void QThread_Usleep(unsigned long param1) {
	QThread::usleep(static_cast<unsigned long>(param1));
}
struct miqt_string QThread_Tr2(const char* s, const char* c) {
	QString _ret = QThread::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QThread_Tr3(const char* s, const char* c, int n) {
	QString _ret = QThread::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QThread_Start1(QThread* self, Priority param1) {
	self->start(param1);
}
void QThread_Exit1(QThread* self, int retcode) {
	self->exit(static_cast<int>(retcode));
}
bool QThread_Wait1(QThread* self, QDeadlineTimer* deadline) {
	return self->wait(*deadline);
}
void QThread_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQThread*>( (QThread*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QThread_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQThread*)(self) )->virtualbase_MetaObject();
}
void QThread_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQThread*>( (QThread*)(self) )->handle__Metacast = slot;
}
void* QThread_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQThread*)(self) )->virtualbase_Metacast(param1);
}
void QThread_Delete(QThread* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQThread*>( self );
	} else {
		delete self;
	}
}
