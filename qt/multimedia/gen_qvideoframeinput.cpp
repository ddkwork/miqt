// +build ignore

#include <QMediaCaptureSession>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVideoFrame>
#include <QVideoFrameFormat>
#include <QVideoFrameInput>
#include <qvideoframeinput.h>
#include "gen_qvideoframeinput.h"
class MiqtVirtualQVideoFrameInput : public virtual QVideoFrameInput {
public:
MiqtVirtualQVideoFrameInput(): QVideoFrameInput() {};
MiqtVirtualQVideoFrameInput(const QVideoFrameFormat& format): QVideoFrameInput(format) {};
MiqtVirtualQVideoFrameInput(QObject* parent): QVideoFrameInput(parent) {};
MiqtVirtualQVideoFrameInput(const QVideoFrameFormat& format, QObject* parent): QVideoFrameInput(format, parent) {};

virtual ~MiqtVirtualQVideoFrameInput() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QVideoFrameInput::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QVideoFrameInput_MetaObject(const_cast<MiqtVirtualQVideoFrameInput*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QVideoFrameInput::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QVideoFrameInput::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QVideoFrameInput_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QVideoFrameInput::qt_metacast(param1);

	}
};
QVideoFrameInput* QVideoFrameInput_new() {
return new MiqtVirtualQVideoFrameInput();
}
QVideoFrameInput* QVideoFrameInput_new2(QVideoFrameFormat* format) {
return new MiqtVirtualQVideoFrameInput(*format);
}
QVideoFrameInput* QVideoFrameInput_new3(QObject* parent) {
return new MiqtVirtualQVideoFrameInput(parent);
}
QVideoFrameInput* QVideoFrameInput_new4(QVideoFrameFormat* format, QObject* parent) {
return new MiqtVirtualQVideoFrameInput(*format, parent);
}
void QVideoFrameInput_virtbase(QVideoFrameInput* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QVideoFrameInput_MetaObject(const QVideoFrameInput* self) {
	return (QMetaObject*) self->metaObject();
}
void* QVideoFrameInput_Metacast(QVideoFrameInput* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QVideoFrameInput_Tr(const char* s) {
	QString _ret = QVideoFrameInput::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QVideoFrameInput_SendVideoFrame(QVideoFrameInput* self, QVideoFrame* frame) {
	return self->sendVideoFrame(*frame);
}
QVideoFrameFormat* QVideoFrameInput_Format(const QVideoFrameInput* self) {
	return new QVideoFrameFormat(self->format());
}
QMediaCaptureSession* QVideoFrameInput_CaptureSession(const QVideoFrameInput* self) {
	return self->captureSession();
}
void QVideoFrameInput_ReadyToSendVideoFrame(QVideoFrameInput* self) {
	self->readyToSendVideoFrame();
}
void QVideoFrameInput_connect_ReadyToSendVideoFrame(QVideoFrameInput* self, intptr_t slot) {
	MiqtVirtualQVideoFrameInput::connect(self, static_cast<void (QVideoFrameInput::*)()>(&QVideoFrameInput::readyToSendVideoFrame), self, [=]() {
		miqt_exec_callback_QVideoFrameInput_ReadyToSendVideoFrame(slot);
	});
}
struct miqt_string QVideoFrameInput_Tr2(const char* s, const char* c) {
	QString _ret = QVideoFrameInput::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QVideoFrameInput_Tr3(const char* s, const char* c, int n) {
	QString _ret = QVideoFrameInput::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QVideoFrameInput_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQVideoFrameInput*>( (QVideoFrameInput*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QVideoFrameInput_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQVideoFrameInput*)(self) )->virtualbase_MetaObject();
}
void QVideoFrameInput_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQVideoFrameInput*>( (QVideoFrameInput*)(self) )->handle__Metacast = slot;
}
void* QVideoFrameInput_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQVideoFrameInput*)(self) )->virtualbase_Metacast(param1);
}
void QVideoFrameInput_Delete(QVideoFrameInput* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQVideoFrameInput*>( self );
	} else {
		delete self;
	}
}
