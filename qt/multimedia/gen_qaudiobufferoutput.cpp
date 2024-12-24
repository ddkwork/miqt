// +build ignore

#include <QAudioBuffer>
#include <QAudioBufferOutput>
#include <QAudioFormat>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qaudiobufferoutput.h>
#include "gen_qaudiobufferoutput.h"
class MiqtVirtualQAudioBufferOutput : public virtual QAudioBufferOutput {
public:
MiqtVirtualQAudioBufferOutput(): QAudioBufferOutput() {};
MiqtVirtualQAudioBufferOutput(const QAudioFormat& format): QAudioBufferOutput(format) {};
MiqtVirtualQAudioBufferOutput(QObject* parent): QAudioBufferOutput(parent) {};
MiqtVirtualQAudioBufferOutput(const QAudioFormat& format, QObject* parent): QAudioBufferOutput(format, parent) {};

virtual ~MiqtVirtualQAudioBufferOutput() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QAudioBufferOutput::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QAudioBufferOutput_MetaObject(const_cast<MiqtVirtualQAudioBufferOutput*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QAudioBufferOutput::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QAudioBufferOutput::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QAudioBufferOutput_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QAudioBufferOutput::qt_metacast(param1);

	}
};
QAudioBufferOutput* QAudioBufferOutput_new() {
return new MiqtVirtualQAudioBufferOutput();
}
QAudioBufferOutput* QAudioBufferOutput_new2(QAudioFormat* format) {
return new MiqtVirtualQAudioBufferOutput(*format);
}
QAudioBufferOutput* QAudioBufferOutput_new3(QObject* parent) {
return new MiqtVirtualQAudioBufferOutput(parent);
}
QAudioBufferOutput* QAudioBufferOutput_new4(QAudioFormat* format, QObject* parent) {
return new MiqtVirtualQAudioBufferOutput(*format, parent);
}
void QAudioBufferOutput_virtbase(QAudioBufferOutput* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QAudioBufferOutput_MetaObject(const QAudioBufferOutput* self) {
	return (QMetaObject*) self->metaObject();
}
void* QAudioBufferOutput_Metacast(QAudioBufferOutput* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QAudioBufferOutput_Tr(const char* s) {
	QString _ret = QAudioBufferOutput::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QAudioFormat* QAudioBufferOutput_Format(const QAudioBufferOutput* self) {
	return new QAudioFormat(self->format());
}
void QAudioBufferOutput_AudioBufferReceived(QAudioBufferOutput* self, QAudioBuffer* buffer) {
	self->audioBufferReceived(*buffer);
}
void QAudioBufferOutput_connect_AudioBufferReceived(QAudioBufferOutput* self, intptr_t slot) {
	MiqtVirtualQAudioBufferOutput::connect(self, static_cast<void (QAudioBufferOutput::*)(const QAudioBuffer&)>(&QAudioBufferOutput::audioBufferReceived), self, [=](const QAudioBuffer& buffer) {
		const QAudioBuffer& buffer_ret = buffer;
		// Cast returned reference into pointer
		QAudioBuffer* sigval1 = const_cast<QAudioBuffer*>(&buffer_ret);
		miqt_exec_callback_QAudioBufferOutput_AudioBufferReceived(slot, sigval1);
	});
}
struct miqt_string QAudioBufferOutput_Tr2(const char* s, const char* c) {
	QString _ret = QAudioBufferOutput::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QAudioBufferOutput_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAudioBufferOutput::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QAudioBufferOutput_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAudioBufferOutput*>( (QAudioBufferOutput*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QAudioBufferOutput_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQAudioBufferOutput*)(self) )->virtualbase_MetaObject();
}
void QAudioBufferOutput_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAudioBufferOutput*>( (QAudioBufferOutput*)(self) )->handle__Metacast = slot;
}
void* QAudioBufferOutput_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQAudioBufferOutput*)(self) )->virtualbase_Metacast(param1);
}
void QAudioBufferOutput_Delete(QAudioBufferOutput* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQAudioBufferOutput*>( self );
	} else {
		delete self;
	}
}
