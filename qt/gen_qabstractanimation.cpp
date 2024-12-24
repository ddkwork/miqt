// +build ignore

#include <QAbstractAnimation>
#include <QAnimationDriver>
#include <QAnimationGroup>
#include <QEvent>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qabstractanimation.h>
#include "gen_qabstractanimation.h"
class MiqtVirtualQAbstractAnimation : public virtual QAbstractAnimation {
public:
MiqtVirtualQAbstractAnimation(): QAbstractAnimation() {};
MiqtVirtualQAbstractAnimation(QObject* parent): QAbstractAnimation(parent) {};

virtual ~MiqtVirtualQAbstractAnimation() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QAbstractAnimation::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QAbstractAnimation_MetaObject(const_cast<MiqtVirtualQAbstractAnimation*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QAbstractAnimation::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QAbstractAnimation::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QAbstractAnimation_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QAbstractAnimation::qt_metacast(param1);

	}
};
QAbstractAnimation* QAbstractAnimation_new() {
return new MiqtVirtualQAbstractAnimation();
}
QAbstractAnimation* QAbstractAnimation_new2(QObject* parent) {
return new MiqtVirtualQAbstractAnimation(parent);
}
void QAbstractAnimation_virtbase(QAbstractAnimation* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QAbstractAnimation_MetaObject(const QAbstractAnimation* self) {
	return (QMetaObject*) self->metaObject();
}
void* QAbstractAnimation_Metacast(QAbstractAnimation* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QAbstractAnimation_Tr(const char* s) {
	QString _ret = QAbstractAnimation::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
State QAbstractAnimation_State(const QAbstractAnimation* self) {
	return self->state();
}
QAnimationGroup* QAbstractAnimation_Group(const QAbstractAnimation* self) {
	return self->group();
}
Direction QAbstractAnimation_Direction(const QAbstractAnimation* self) {
	return self->direction();
}
void QAbstractAnimation_SetDirection(QAbstractAnimation* self, Direction direction) {
	self->setDirection(direction);
}
int QAbstractAnimation_CurrentTime(const QAbstractAnimation* self) {
	return self->currentTime();
}
int QAbstractAnimation_CurrentLoopTime(const QAbstractAnimation* self) {
	return self->currentLoopTime();
}
int QAbstractAnimation_LoopCount(const QAbstractAnimation* self) {
	return self->loopCount();
}
void QAbstractAnimation_SetLoopCount(QAbstractAnimation* self, int loopCount) {
	self->setLoopCount(static_cast<int>(loopCount));
}
int QAbstractAnimation_CurrentLoop(const QAbstractAnimation* self) {
	return self->currentLoop();
}
int QAbstractAnimation_Duration(const QAbstractAnimation* self) {
	return self->duration();
}
int QAbstractAnimation_TotalDuration(const QAbstractAnimation* self) {
	return self->totalDuration();
}
void QAbstractAnimation_Finished(QAbstractAnimation* self) {
	self->finished();
}
void QAbstractAnimation_connect_Finished(QAbstractAnimation* self, intptr_t slot) {
	MiqtVirtualQAbstractAnimation::connect(self, static_cast<void (QAbstractAnimation::*)()>(&QAbstractAnimation::finished), self, [=]() {
		miqt_exec_callback_QAbstractAnimation_Finished(slot);
	});
}
void QAbstractAnimation_StateChanged(QAbstractAnimation* self, int newState, int oldState) {
	self->stateChanged(static_cast<QAbstractAnimation::State>(newState), static_cast<QAbstractAnimation::State>(oldState));
}
void QAbstractAnimation_connect_StateChanged(QAbstractAnimation* self, intptr_t slot) {
	MiqtVirtualQAbstractAnimation::connect(self, static_cast<void (QAbstractAnimation::*)(QAbstractAnimation::State, QAbstractAnimation::State)>(&QAbstractAnimation::stateChanged), self, [=](QAbstractAnimation::State newState, QAbstractAnimation::State oldState) {
		QAbstractAnimation::State newState_ret = newState;
		int sigval1 = static_cast<int>(newState_ret);
		QAbstractAnimation::State oldState_ret = oldState;
		int sigval2 = static_cast<int>(oldState_ret);
		miqt_exec_callback_QAbstractAnimation_StateChanged(slot, sigval1, sigval2);
	});
}
void QAbstractAnimation_CurrentLoopChanged(QAbstractAnimation* self, int currentLoop) {
	self->currentLoopChanged(static_cast<int>(currentLoop));
}
void QAbstractAnimation_connect_CurrentLoopChanged(QAbstractAnimation* self, intptr_t slot) {
	MiqtVirtualQAbstractAnimation::connect(self, static_cast<void (QAbstractAnimation::*)(int)>(&QAbstractAnimation::currentLoopChanged), self, [=](int currentLoop) {
		int sigval1 = currentLoop;
		miqt_exec_callback_QAbstractAnimation_CurrentLoopChanged(slot, sigval1);
	});
}
void QAbstractAnimation_DirectionChanged(QAbstractAnimation* self, int param1) {
	self->directionChanged(static_cast<QAbstractAnimation::Direction>(param1));
}
void QAbstractAnimation_connect_DirectionChanged(QAbstractAnimation* self, intptr_t slot) {
	MiqtVirtualQAbstractAnimation::connect(self, static_cast<void (QAbstractAnimation::*)(QAbstractAnimation::Direction)>(&QAbstractAnimation::directionChanged), self, [=](QAbstractAnimation::Direction param1) {
		QAbstractAnimation::Direction param1_ret = param1;
		int sigval1 = static_cast<int>(param1_ret);
		miqt_exec_callback_QAbstractAnimation_DirectionChanged(slot, sigval1);
	});
}
void QAbstractAnimation_Start(QAbstractAnimation* self) {
	self->start();
}
void QAbstractAnimation_Pause(QAbstractAnimation* self) {
	self->pause();
}
void QAbstractAnimation_Resume(QAbstractAnimation* self) {
	self->resume();
}
void QAbstractAnimation_SetPaused(QAbstractAnimation* self, bool paused) {
	self->setPaused(paused);
}
void QAbstractAnimation_Stop(QAbstractAnimation* self) {
	self->stop();
}
void QAbstractAnimation_SetCurrentTime(QAbstractAnimation* self, int msecs) {
	self->setCurrentTime(static_cast<int>(msecs));
}
struct miqt_string QAbstractAnimation_Tr2(const char* s, const char* c) {
	QString _ret = QAbstractAnimation::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QAbstractAnimation_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAbstractAnimation::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QAbstractAnimation_Start1(QAbstractAnimation* self, int policy) {
	self->start(static_cast<QAbstractAnimation::DeletionPolicy>(policy));
}
void QAbstractAnimation_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAbstractAnimation*>( (QAbstractAnimation*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QAbstractAnimation_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQAbstractAnimation*)(self) )->virtualbase_MetaObject();
}
void QAbstractAnimation_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAbstractAnimation*>( (QAbstractAnimation*)(self) )->handle__Metacast = slot;
}
void* QAbstractAnimation_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQAbstractAnimation*)(self) )->virtualbase_Metacast(param1);
}
void QAbstractAnimation_Delete(QAbstractAnimation* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQAbstractAnimation*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQAnimationDriver : public virtual QAnimationDriver {
public:
MiqtVirtualQAnimationDriver(): QAnimationDriver() {};
MiqtVirtualQAnimationDriver(QObject* parent): QAnimationDriver(parent) {};

virtual ~MiqtVirtualQAnimationDriver() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QAnimationDriver::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QAnimationDriver_MetaObject(const_cast<MiqtVirtualQAnimationDriver*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QAnimationDriver::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QAnimationDriver::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QAnimationDriver_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QAnimationDriver::qt_metacast(param1);

	}
};
QAnimationDriver* QAnimationDriver_new() {
return new MiqtVirtualQAnimationDriver();
}
QAnimationDriver* QAnimationDriver_new2(QObject* parent) {
return new MiqtVirtualQAnimationDriver(parent);
}
void QAnimationDriver_virtbase(QAnimationDriver* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QAnimationDriver_MetaObject(const QAnimationDriver* self) {
	return (QMetaObject*) self->metaObject();
}
void* QAnimationDriver_Metacast(QAnimationDriver* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QAnimationDriver_Tr(const char* s) {
	QString _ret = QAnimationDriver::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QAnimationDriver_Advance(QAnimationDriver* self) {
	self->advance();
}
void QAnimationDriver_Install(QAnimationDriver* self) {
	self->install();
}
void QAnimationDriver_Uninstall(QAnimationDriver* self) {
	self->uninstall();
}
bool QAnimationDriver_IsRunning(const QAnimationDriver* self) {
	return self->isRunning();
}
long long QAnimationDriver_Elapsed(const QAnimationDriver* self) {
	qint64 _ret = self->elapsed();
	return static_cast<long long>(_ret);
}
void QAnimationDriver_Started(QAnimationDriver* self) {
	self->started();
}
void QAnimationDriver_connect_Started(QAnimationDriver* self, intptr_t slot) {
	MiqtVirtualQAnimationDriver::connect(self, static_cast<void (QAnimationDriver::*)()>(&QAnimationDriver::started), self, [=]() {
		miqt_exec_callback_QAnimationDriver_Started(slot);
	});
}
void QAnimationDriver_Stopped(QAnimationDriver* self) {
	self->stopped();
}
void QAnimationDriver_connect_Stopped(QAnimationDriver* self, intptr_t slot) {
	MiqtVirtualQAnimationDriver::connect(self, static_cast<void (QAnimationDriver::*)()>(&QAnimationDriver::stopped), self, [=]() {
		miqt_exec_callback_QAnimationDriver_Stopped(slot);
	});
}
struct miqt_string QAnimationDriver_Tr2(const char* s, const char* c) {
	QString _ret = QAnimationDriver::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QAnimationDriver_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAnimationDriver::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QAnimationDriver_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAnimationDriver*>( (QAnimationDriver*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QAnimationDriver_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQAnimationDriver*)(self) )->virtualbase_MetaObject();
}
void QAnimationDriver_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAnimationDriver*>( (QAnimationDriver*)(self) )->handle__Metacast = slot;
}
void* QAnimationDriver_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQAnimationDriver*)(self) )->virtualbase_Metacast(param1);
}
void QAnimationDriver_Delete(QAnimationDriver* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQAnimationDriver*>( self );
	} else {
		delete self;
	}
}
