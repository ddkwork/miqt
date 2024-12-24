// +build ignore

#include <QEasingCurve>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimeLine>
#include <QTimerEvent>
#include <qtimeline.h>
#include "gen_qtimeline.h"
class MiqtVirtualQTimeLine : public virtual QTimeLine {
public:
MiqtVirtualQTimeLine(): QTimeLine() {};
MiqtVirtualQTimeLine(int duration): QTimeLine(duration) {};
MiqtVirtualQTimeLine(int duration, QObject* parent): QTimeLine(duration, parent) {};

virtual ~MiqtVirtualQTimeLine() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QTimeLine::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QTimeLine_MetaObject(const_cast<MiqtVirtualQTimeLine*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QTimeLine::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QTimeLine::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QTimeLine_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QTimeLine::qt_metacast(param1);

	}
};
QTimeLine* QTimeLine_new() {
return new MiqtVirtualQTimeLine();
}
QTimeLine* QTimeLine_new2(int duration) {
return new MiqtVirtualQTimeLine(static_cast<int>(duration));
}
QTimeLine* QTimeLine_new3(int duration, QObject* parent) {
return new MiqtVirtualQTimeLine(static_cast<int>(duration), parent);
}
void QTimeLine_virtbase(QTimeLine* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QTimeLine_MetaObject(const QTimeLine* self) {
	return (QMetaObject*) self->metaObject();
}
void* QTimeLine_Metacast(QTimeLine* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QTimeLine_Tr(const char* s) {
	QString _ret = QTimeLine::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
State QTimeLine_State(const QTimeLine* self) {
	return self->state();
}
int QTimeLine_LoopCount(const QTimeLine* self) {
	return self->loopCount();
}
void QTimeLine_SetLoopCount(QTimeLine* self, int count) {
	self->setLoopCount(static_cast<int>(count));
}
Direction QTimeLine_Direction(const QTimeLine* self) {
	return self->direction();
}
void QTimeLine_SetDirection(QTimeLine* self, Direction direction) {
	self->setDirection(direction);
}
int QTimeLine_Duration(const QTimeLine* self) {
	return self->duration();
}
void QTimeLine_SetDuration(QTimeLine* self, int duration) {
	self->setDuration(static_cast<int>(duration));
}
int QTimeLine_StartFrame(const QTimeLine* self) {
	return self->startFrame();
}
void QTimeLine_SetStartFrame(QTimeLine* self, int frame) {
	self->setStartFrame(static_cast<int>(frame));
}
int QTimeLine_EndFrame(const QTimeLine* self) {
	return self->endFrame();
}
void QTimeLine_SetEndFrame(QTimeLine* self, int frame) {
	self->setEndFrame(static_cast<int>(frame));
}
void QTimeLine_SetFrameRange(QTimeLine* self, int startFrame, int endFrame) {
	self->setFrameRange(static_cast<int>(startFrame), static_cast<int>(endFrame));
}
int QTimeLine_UpdateInterval(const QTimeLine* self) {
	return self->updateInterval();
}
void QTimeLine_SetUpdateInterval(QTimeLine* self, int interval) {
	self->setUpdateInterval(static_cast<int>(interval));
}
QEasingCurve* QTimeLine_EasingCurve(const QTimeLine* self) {
	return new QEasingCurve(self->easingCurve());
}
void QTimeLine_SetEasingCurve(QTimeLine* self, QEasingCurve* curve) {
	self->setEasingCurve(*curve);
}
int QTimeLine_CurrentTime(const QTimeLine* self) {
	return self->currentTime();
}
int QTimeLine_CurrentFrame(const QTimeLine* self) {
	return self->currentFrame();
}
double QTimeLine_CurrentValue(const QTimeLine* self) {
	qreal _ret = self->currentValue();
	return static_cast<double>(_ret);
}
int QTimeLine_FrameForTime(const QTimeLine* self, int msec) {
	return self->frameForTime(static_cast<int>(msec));
}
double QTimeLine_ValueForTime(const QTimeLine* self, int msec) {
	qreal _ret = self->valueForTime(static_cast<int>(msec));
	return static_cast<double>(_ret);
}
void QTimeLine_Start(QTimeLine* self) {
	self->start();
}
void QTimeLine_Resume(QTimeLine* self) {
	self->resume();
}
void QTimeLine_Stop(QTimeLine* self) {
	self->stop();
}
void QTimeLine_SetPaused(QTimeLine* self, bool paused) {
	self->setPaused(paused);
}
void QTimeLine_SetCurrentTime(QTimeLine* self, int msec) {
	self->setCurrentTime(static_cast<int>(msec));
}
void QTimeLine_ToggleDirection(QTimeLine* self) {
	self->toggleDirection();
}
struct miqt_string QTimeLine_Tr2(const char* s, const char* c) {
	QString _ret = QTimeLine::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QTimeLine_Tr3(const char* s, const char* c, int n) {
	QString _ret = QTimeLine::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QTimeLine_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTimeLine*>( (QTimeLine*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QTimeLine_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQTimeLine*)(self) )->virtualbase_MetaObject();
}
void QTimeLine_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTimeLine*>( (QTimeLine*)(self) )->handle__Metacast = slot;
}
void* QTimeLine_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQTimeLine*)(self) )->virtualbase_Metacast(param1);
}
void QTimeLine_Delete(QTimeLine* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQTimeLine*>( self );
	} else {
		delete self;
	}
}
