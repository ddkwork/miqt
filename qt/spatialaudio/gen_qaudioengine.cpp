// +build ignore

#include <QAudioDevice>
#include <QAudioEngine>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qaudioengine.h>
#include "gen_qaudioengine.h"
class MiqtVirtualQAudioEngine : public virtual QAudioEngine {
public:
MiqtVirtualQAudioEngine(): QAudioEngine() {};
MiqtVirtualQAudioEngine(QObject* parent): QAudioEngine(parent) {};
MiqtVirtualQAudioEngine(int sampleRate): QAudioEngine(sampleRate) {};
MiqtVirtualQAudioEngine(int sampleRate, QObject* parent): QAudioEngine(sampleRate, parent) {};

virtual ~MiqtVirtualQAudioEngine() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QAudioEngine::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QAudioEngine_MetaObject(const_cast<MiqtVirtualQAudioEngine*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QAudioEngine::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QAudioEngine::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QAudioEngine_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QAudioEngine::qt_metacast(param1);

	}
};
QAudioEngine* QAudioEngine_new() {
return new MiqtVirtualQAudioEngine();
}
QAudioEngine* QAudioEngine_new2(QObject* parent) {
return new MiqtVirtualQAudioEngine(parent);
}
QAudioEngine* QAudioEngine_new3(int sampleRate) {
return new MiqtVirtualQAudioEngine(static_cast<int>(sampleRate));
}
QAudioEngine* QAudioEngine_new4(int sampleRate, QObject* parent) {
return new MiqtVirtualQAudioEngine(static_cast<int>(sampleRate), parent);
}
void QAudioEngine_virtbase(QAudioEngine* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QAudioEngine_MetaObject(const QAudioEngine* self) {
	return (QMetaObject*) self->metaObject();
}
void* QAudioEngine_Metacast(QAudioEngine* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QAudioEngine_Tr(const char* s) {
	QString _ret = QAudioEngine::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QAudioEngine_SetOutputMode(QAudioEngine* self, OutputMode mode) {
	self->setOutputMode(mode);
}
OutputMode QAudioEngine_OutputMode(const QAudioEngine* self) {
	return self->outputMode();
}
int QAudioEngine_SampleRate(const QAudioEngine* self) {
	return self->sampleRate();
}
void QAudioEngine_SetOutputDevice(QAudioEngine* self, QAudioDevice* device) {
	self->setOutputDevice(*device);
}
QAudioDevice* QAudioEngine_OutputDevice(const QAudioEngine* self) {
	return new QAudioDevice(self->outputDevice());
}
void QAudioEngine_SetMasterVolume(QAudioEngine* self, float volume) {
	self->setMasterVolume(static_cast<float>(volume));
}
float QAudioEngine_MasterVolume(const QAudioEngine* self) {
	return self->masterVolume();
}
void QAudioEngine_SetPaused(QAudioEngine* self, bool paused) {
	self->setPaused(paused);
}
bool QAudioEngine_Paused(const QAudioEngine* self) {
	return self->paused();
}
void QAudioEngine_SetRoomEffectsEnabled(QAudioEngine* self, bool enabled) {
	self->setRoomEffectsEnabled(enabled);
}
bool QAudioEngine_RoomEffectsEnabled(const QAudioEngine* self) {
	return self->roomEffectsEnabled();
}
void QAudioEngine_SetDistanceScale(QAudioEngine* self, float scale) {
	self->setDistanceScale(static_cast<float>(scale));
}
float QAudioEngine_DistanceScale(const QAudioEngine* self) {
	return self->distanceScale();
}
void QAudioEngine_OutputModeChanged(QAudioEngine* self) {
	self->outputModeChanged();
}
void QAudioEngine_connect_OutputModeChanged(QAudioEngine* self, intptr_t slot) {
	MiqtVirtualQAudioEngine::connect(self, static_cast<void (QAudioEngine::*)()>(&QAudioEngine::outputModeChanged), self, [=]() {
		miqt_exec_callback_QAudioEngine_OutputModeChanged(slot);
	});
}
void QAudioEngine_OutputDeviceChanged(QAudioEngine* self) {
	self->outputDeviceChanged();
}
void QAudioEngine_connect_OutputDeviceChanged(QAudioEngine* self, intptr_t slot) {
	MiqtVirtualQAudioEngine::connect(self, static_cast<void (QAudioEngine::*)()>(&QAudioEngine::outputDeviceChanged), self, [=]() {
		miqt_exec_callback_QAudioEngine_OutputDeviceChanged(slot);
	});
}
void QAudioEngine_MasterVolumeChanged(QAudioEngine* self) {
	self->masterVolumeChanged();
}
void QAudioEngine_connect_MasterVolumeChanged(QAudioEngine* self, intptr_t slot) {
	MiqtVirtualQAudioEngine::connect(self, static_cast<void (QAudioEngine::*)()>(&QAudioEngine::masterVolumeChanged), self, [=]() {
		miqt_exec_callback_QAudioEngine_MasterVolumeChanged(slot);
	});
}
void QAudioEngine_PausedChanged(QAudioEngine* self) {
	self->pausedChanged();
}
void QAudioEngine_connect_PausedChanged(QAudioEngine* self, intptr_t slot) {
	MiqtVirtualQAudioEngine::connect(self, static_cast<void (QAudioEngine::*)()>(&QAudioEngine::pausedChanged), self, [=]() {
		miqt_exec_callback_QAudioEngine_PausedChanged(slot);
	});
}
void QAudioEngine_DistanceScaleChanged(QAudioEngine* self) {
	self->distanceScaleChanged();
}
void QAudioEngine_connect_DistanceScaleChanged(QAudioEngine* self, intptr_t slot) {
	MiqtVirtualQAudioEngine::connect(self, static_cast<void (QAudioEngine::*)()>(&QAudioEngine::distanceScaleChanged), self, [=]() {
		miqt_exec_callback_QAudioEngine_DistanceScaleChanged(slot);
	});
}
void QAudioEngine_Start(QAudioEngine* self) {
	self->start();
}
void QAudioEngine_Stop(QAudioEngine* self) {
	self->stop();
}
void QAudioEngine_Pause(QAudioEngine* self) {
	self->pause();
}
void QAudioEngine_Resume(QAudioEngine* self) {
	self->resume();
}
struct miqt_string QAudioEngine_Tr2(const char* s, const char* c) {
	QString _ret = QAudioEngine::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QAudioEngine_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAudioEngine::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QAudioEngine_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAudioEngine*>( (QAudioEngine*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QAudioEngine_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQAudioEngine*)(self) )->virtualbase_MetaObject();
}
void QAudioEngine_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAudioEngine*>( (QAudioEngine*)(self) )->handle__Metacast = slot;
}
void* QAudioEngine_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQAudioEngine*)(self) )->virtualbase_Metacast(param1);
}
void QAudioEngine_Delete(QAudioEngine* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQAudioEngine*>( self );
	} else {
		delete self;
	}
}
