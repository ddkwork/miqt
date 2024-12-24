// +build ignore

#include <QAudioDevice>
#include <QAudioFormat>
#include <QAudioSink>
#include <QIODevice>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qaudiosink.h>
#include "gen_qaudiosink.h"

class MiqtVirtualQAudioSink : public virtual QAudioSink {
public:

	MiqtVirtualQAudioSink(): QAudioSink() {};
	MiqtVirtualQAudioSink(const QAudioDevice& audioDeviceInfo): QAudioSink(audioDeviceInfo) {};
	MiqtVirtualQAudioSink(const QAudioFormat& format): QAudioSink(format) {};
	MiqtVirtualQAudioSink(const QAudioFormat& format, QObject* parent): QAudioSink(format, parent) {};
	MiqtVirtualQAudioSink(const QAudioDevice& audioDeviceInfo, const QAudioFormat& format): QAudioSink(audioDeviceInfo, format) {};
	MiqtVirtualQAudioSink(const QAudioDevice& audioDeviceInfo, const QAudioFormat& format, QObject* parent): QAudioSink(audioDeviceInfo, format, parent) {};

	virtual ~MiqtVirtualQAudioSink() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QAudioSink::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QAudioSink_MetaObject(const_cast<MiqtVirtualQAudioSink*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QAudioSink::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QAudioSink::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QAudioSink_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QAudioSink::qt_metacast(param1);

	}

};

QAudioSink* QAudioSink_new() {
	return new MiqtVirtualQAudioSink();
}

QAudioSink* QAudioSink_new2(QAudioDevice* audioDeviceInfo) {
	return new MiqtVirtualQAudioSink(*audioDeviceInfo);
}

QAudioSink* QAudioSink_new3(QAudioFormat* format) {
	return new MiqtVirtualQAudioSink(*format);
}

QAudioSink* QAudioSink_new4(QAudioFormat* format, QObject* parent) {
	return new MiqtVirtualQAudioSink(*format, parent);
}

QAudioSink* QAudioSink_new5(QAudioDevice* audioDeviceInfo, QAudioFormat* format) {
	return new MiqtVirtualQAudioSink(*audioDeviceInfo, *format);
}

QAudioSink* QAudioSink_new6(QAudioDevice* audioDeviceInfo, QAudioFormat* format, QObject* parent) {
	return new MiqtVirtualQAudioSink(*audioDeviceInfo, *format, parent);
}

void QAudioSink_virtbase(QAudioSink* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QAudioSink_MetaObject(const QAudioSink* self) {
	return (QMetaObject*) self->metaObject();
}

void* QAudioSink_Metacast(QAudioSink* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QAudioSink_Tr(const char* s) {
	QString _ret = QAudioSink::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QAudioSink_IsNull(const QAudioSink* self) {
	return self->isNull();
}

QAudioFormat* QAudioSink_Format(const QAudioSink* self) {
	return new QAudioFormat(self->format());
}

void QAudioSink_Start(QAudioSink* self, QIODevice* device) {
	self->start(device);
}

QIODevice* QAudioSink_Start2(QAudioSink* self) {
	return self->start();
}

void QAudioSink_Stop(QAudioSink* self) {
	self->stop();
}

void QAudioSink_Reset(QAudioSink* self) {
	self->reset();
}

void QAudioSink_Suspend(QAudioSink* self) {
	self->suspend();
}

void QAudioSink_Resume(QAudioSink* self) {
	self->resume();
}

void QAudioSink_SetBufferSize(QAudioSink* self, ptrdiff_t bytes) {
	self->setBufferSize((qsizetype)(bytes));
}

ptrdiff_t QAudioSink_BufferSize(const QAudioSink* self) {
	qsizetype _ret = self->bufferSize();
	return static_cast<ptrdiff_t>(_ret);
}

ptrdiff_t QAudioSink_BytesFree(const QAudioSink* self) {
	qsizetype _ret = self->bytesFree();
	return static_cast<ptrdiff_t>(_ret);
}

long long QAudioSink_ProcessedUSecs(const QAudioSink* self) {
	qint64 _ret = self->processedUSecs();
	return static_cast<long long>(_ret);
}

long long QAudioSink_ElapsedUSecs(const QAudioSink* self) {
	qint64 _ret = self->elapsedUSecs();
	return static_cast<long long>(_ret);
}

QtAudio::Error QAudioSink_Error(const QAudioSink* self) {
	return self->error();
}

QtAudio::State QAudioSink_State(const QAudioSink* self) {
	return self->state();
}

void QAudioSink_SetVolume(QAudioSink* self, double volume) {
	self->setVolume(static_cast<qreal>(volume));
}

double QAudioSink_Volume(const QAudioSink* self) {
	qreal _ret = self->volume();
	return static_cast<double>(_ret);
}

void QAudioSink_StateChanged(QAudioSink* self, int state) {
	self->stateChanged(static_cast<QAudio::State>(state));
}

void QAudioSink_connect_StateChanged(QAudioSink* self, intptr_t slot) {
	MiqtVirtualQAudioSink::connect(self, static_cast<void (QAudioSink::*)(QAudio::State)>(&QAudioSink::stateChanged), self, [=](QAudio::State state) {
		QAudio::State state_ret = state;
		int sigval1 = static_cast<int>(state_ret);
		miqt_exec_callback_QAudioSink_StateChanged(slot, sigval1);
	});
}

struct miqt_string QAudioSink_Tr2(const char* s, const char* c) {
	QString _ret = QAudioSink::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAudioSink_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAudioSink::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAudioSink_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAudioSink*>( (QAudioSink*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QAudioSink_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQAudioSink*)(self) )->virtualbase_MetaObject();
}

void QAudioSink_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAudioSink*>( (QAudioSink*)(self) )->handle__Metacast = slot;
}

void* QAudioSink_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQAudioSink*)(self) )->virtualbase_Metacast(param1);
}

void QAudioSink_Delete(QAudioSink* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQAudioSink*>( self );
	} else {
		delete self;
	}
}

