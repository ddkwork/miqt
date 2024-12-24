// +build ignore

#include <QAudioBuffer>
#include <QAudioBufferInput>
#include <QAudioFormat>
#include <QMediaCaptureSession>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qaudiobufferinput.h>
#include "gen_qaudiobufferinput.h"
class MiqtVirtualQAudioBufferInput : public virtual QAudioBufferInput {
public:
MiqtVirtualQAudioBufferInput(): QAudioBufferInput() {};
MiqtVirtualQAudioBufferInput(const QAudioFormat& format): QAudioBufferInput(format) {};
MiqtVirtualQAudioBufferInput(QObject* parent): QAudioBufferInput(parent) {};
MiqtVirtualQAudioBufferInput(const QAudioFormat& format, QObject* parent): QAudioBufferInput(format, parent) {};

virtual ~MiqtVirtualQAudioBufferInput() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QAudioBufferInput::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QAudioBufferInput_MetaObject(const_cast<MiqtVirtualQAudioBufferInput*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QAudioBufferInput::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QAudioBufferInput::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QAudioBufferInput_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QAudioBufferInput::qt_metacast(param1);

	}
};
QAudioBufferInput* QAudioBufferInput_new() {
return new MiqtVirtualQAudioBufferInput();
}
QAudioBufferInput* QAudioBufferInput_new2(QAudioFormat* format) {
return new MiqtVirtualQAudioBufferInput(*format);
}
QAudioBufferInput* QAudioBufferInput_new3(QObject* parent) {
return new MiqtVirtualQAudioBufferInput(parent);
}
QAudioBufferInput* QAudioBufferInput_new4(QAudioFormat* format, QObject* parent) {
return new MiqtVirtualQAudioBufferInput(*format, parent);
}
void QAudioBufferInput_virtbase(QAudioBufferInput* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QAudioBufferInput_MetaObject(const QAudioBufferInput* self) {
	return (QMetaObject*) self->metaObject();
}
void* QAudioBufferInput_Metacast(QAudioBufferInput* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QAudioBufferInput_Tr(const char* s) {
	QString _ret = QAudioBufferInput::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QAudioBufferInput_SendAudioBuffer(QAudioBufferInput* self, QAudioBuffer* audioBuffer) {
	return self->sendAudioBuffer(*audioBuffer);
}
QAudioFormat* QAudioBufferInput_Format(const QAudioBufferInput* self) {
	return new QAudioFormat(self->format());
}
QMediaCaptureSession* QAudioBufferInput_CaptureSession(const QAudioBufferInput* self) {
	return self->captureSession();
}
void QAudioBufferInput_ReadyToSendAudioBuffer(QAudioBufferInput* self) {
	self->readyToSendAudioBuffer();
}
void QAudioBufferInput_connect_ReadyToSendAudioBuffer(QAudioBufferInput* self, intptr_t slot) {
	MiqtVirtualQAudioBufferInput::connect(self, static_cast<void (QAudioBufferInput::*)()>(&QAudioBufferInput::readyToSendAudioBuffer), self, [=]() {
		miqt_exec_callback_QAudioBufferInput_ReadyToSendAudioBuffer(slot);
	});
}
struct miqt_string QAudioBufferInput_Tr2(const char* s, const char* c) {
	QString _ret = QAudioBufferInput::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QAudioBufferInput_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAudioBufferInput::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QAudioBufferInput_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAudioBufferInput*>( (QAudioBufferInput*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QAudioBufferInput_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQAudioBufferInput*)(self) )->virtualbase_MetaObject();
}
void QAudioBufferInput_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAudioBufferInput*>( (QAudioBufferInput*)(self) )->handle__Metacast = slot;
}
void* QAudioBufferInput_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQAudioBufferInput*)(self) )->virtualbase_Metacast(param1);
}
void QAudioBufferInput_Delete(QAudioBufferInput* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQAudioBufferInput*>( self );
	} else {
		delete self;
	}
}
