// +build ignore

#include <QAudioDevice>
#include <QAudioFormat>
#include <QAudioSource>
#include <QIODevice>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qaudiosource.h>
#include "gen_qaudiosource.h"

class MiqtVirtualQAudioSource : public virtual QAudioSource {
public:

	MiqtVirtualQAudioSource(): QAudioSource() {};
	MiqtVirtualQAudioSource(const QAudioDevice& audioDeviceInfo): QAudioSource(audioDeviceInfo) {};
	MiqtVirtualQAudioSource(const QAudioFormat& format): QAudioSource(format) {};
	MiqtVirtualQAudioSource(const QAudioFormat& format, QObject* parent): QAudioSource(format, parent) {};
	MiqtVirtualQAudioSource(const QAudioDevice& audioDeviceInfo, const QAudioFormat& format): QAudioSource(audioDeviceInfo, format) {};
	MiqtVirtualQAudioSource(const QAudioDevice& audioDeviceInfo, const QAudioFormat& format, QObject* parent): QAudioSource(audioDeviceInfo, format, parent) {};

	virtual ~MiqtVirtualQAudioSource() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QAudioSource::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QAudioSource_MetaObject(const_cast<MiqtVirtualQAudioSource*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QAudioSource::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QAudioSource::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QAudioSource_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QAudioSource::qt_metacast(param1);

	}

};

QAudioSource* QAudioSource_new() {
	return new MiqtVirtualQAudioSource();
}

QAudioSource* QAudioSource_new2(QAudioDevice* audioDeviceInfo) {
	return new MiqtVirtualQAudioSource(*audioDeviceInfo);
}

QAudioSource* QAudioSource_new3(QAudioFormat* format) {
	return new MiqtVirtualQAudioSource(*format);
}

QAudioSource* QAudioSource_new4(QAudioFormat* format, QObject* parent) {
	return new MiqtVirtualQAudioSource(*format, parent);
}

QAudioSource* QAudioSource_new5(QAudioDevice* audioDeviceInfo, QAudioFormat* format) {
	return new MiqtVirtualQAudioSource(*audioDeviceInfo, *format);
}

QAudioSource* QAudioSource_new6(QAudioDevice* audioDeviceInfo, QAudioFormat* format, QObject* parent) {
	return new MiqtVirtualQAudioSource(*audioDeviceInfo, *format, parent);
}

void QAudioSource_virtbase(QAudioSource* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QAudioSource_MetaObject(const QAudioSource* self) {
	return (QMetaObject*) self->metaObject();
}

void* QAudioSource_Metacast(QAudioSource* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QAudioSource_Tr(const char* s) {
	QString _ret = QAudioSource::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QAudioSource_IsNull(const QAudioSource* self) {
	return self->isNull();
}

QAudioFormat* QAudioSource_Format(const QAudioSource* self) {
	return new QAudioFormat(self->format());
}

void QAudioSource_Start(QAudioSource* self, QIODevice* device) {
	self->start(device);
}

QIODevice* QAudioSource_Start2(QAudioSource* self) {
	return self->start();
}

void QAudioSource_Stop(QAudioSource* self) {
	self->stop();
}

void QAudioSource_Reset(QAudioSource* self) {
	self->reset();
}

void QAudioSource_Suspend(QAudioSource* self) {
	self->suspend();
}

void QAudioSource_Resume(QAudioSource* self) {
	self->resume();
}

void QAudioSource_SetBufferSize(QAudioSource* self, ptrdiff_t bytes) {
	self->setBufferSize((qsizetype)(bytes));
}

ptrdiff_t QAudioSource_BufferSize(const QAudioSource* self) {
	qsizetype _ret = self->bufferSize();
	return static_cast<ptrdiff_t>(_ret);
}

ptrdiff_t QAudioSource_BytesAvailable(const QAudioSource* self) {
	qsizetype _ret = self->bytesAvailable();
	return static_cast<ptrdiff_t>(_ret);
}

void QAudioSource_SetVolume(QAudioSource* self, double volume) {
	self->setVolume(static_cast<qreal>(volume));
}

double QAudioSource_Volume(const QAudioSource* self) {
	qreal _ret = self->volume();
	return static_cast<double>(_ret);
}

long long QAudioSource_ProcessedUSecs(const QAudioSource* self) {
	qint64 _ret = self->processedUSecs();
	return static_cast<long long>(_ret);
}

long long QAudioSource_ElapsedUSecs(const QAudioSource* self) {
	qint64 _ret = self->elapsedUSecs();
	return static_cast<long long>(_ret);
}

QtAudio::Error QAudioSource_Error(const QAudioSource* self) {
	return self->error();
}

QtAudio::State QAudioSource_State(const QAudioSource* self) {
	return self->state();
}

void QAudioSource_StateChanged(QAudioSource* self, int state) {
	self->stateChanged(static_cast<QAudio::State>(state));
}

void QAudioSource_connect_StateChanged(QAudioSource* self, intptr_t slot) {
	MiqtVirtualQAudioSource::connect(self, static_cast<void (QAudioSource::*)(QAudio::State)>(&QAudioSource::stateChanged), self, [=](QAudio::State state) {
		QAudio::State state_ret = state;
		int sigval1 = static_cast<int>(state_ret);
		miqt_exec_callback_QAudioSource_StateChanged(slot, sigval1);
	});
}

struct miqt_string QAudioSource_Tr2(const char* s, const char* c) {
	QString _ret = QAudioSource::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAudioSource_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAudioSource::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAudioSource_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAudioSource*>( (QAudioSource*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QAudioSource_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQAudioSource*)(self) )->virtualbase_MetaObject();
}

void QAudioSource_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAudioSource*>( (QAudioSource*)(self) )->handle__Metacast = slot;
}

void* QAudioSource_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQAudioSource*)(self) )->virtualbase_Metacast(param1);
}

void QAudioSource_Delete(QAudioSource* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQAudioSource*>( self );
	} else {
		delete self;
	}
}

