// +build ignore

#include <QAudioDevice>
#include <QList>
#include <QMetaObject>
#include <QObject>
#include <QSoundEffect>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QUrl>
#include <qsoundeffect.h>
#include "gen_qsoundeffect.h"
class MiqtVirtualQSoundEffect : public virtual QSoundEffect {
public:
MiqtVirtualQSoundEffect(): QSoundEffect() {};
MiqtVirtualQSoundEffect(const QAudioDevice& audioDevice): QSoundEffect(audioDevice) {};
MiqtVirtualQSoundEffect(QObject* parent): QSoundEffect(parent) {};
MiqtVirtualQSoundEffect(const QAudioDevice& audioDevice, QObject* parent): QSoundEffect(audioDevice, parent) {};

virtual ~MiqtVirtualQSoundEffect() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QSoundEffect::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QSoundEffect_MetaObject(const_cast<MiqtVirtualQSoundEffect*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QSoundEffect::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QSoundEffect::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QSoundEffect_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QSoundEffect::qt_metacast(param1);

	}
};
QSoundEffect* QSoundEffect_new() {
return new MiqtVirtualQSoundEffect();
}
QSoundEffect* QSoundEffect_new2(QAudioDevice* audioDevice) {
return new MiqtVirtualQSoundEffect(*audioDevice);
}
QSoundEffect* QSoundEffect_new3(QObject* parent) {
return new MiqtVirtualQSoundEffect(parent);
}
QSoundEffect* QSoundEffect_new4(QAudioDevice* audioDevice, QObject* parent) {
return new MiqtVirtualQSoundEffect(*audioDevice, parent);
}
void QSoundEffect_virtbase(QSoundEffect* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QSoundEffect_MetaObject(const QSoundEffect* self) {
	return (QMetaObject*) self->metaObject();
}
void* QSoundEffect_Metacast(QSoundEffect* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QSoundEffect_Tr(const char* s) {
	QString _ret = QSoundEffect::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_array /* of struct miqt_string */  QSoundEffect_SupportedMimeTypes() {
	QStringList _ret = QSoundEffect::supportedMimeTypes();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
QUrl* QSoundEffect_Source(const QSoundEffect* self) {
	return new QUrl(self->source());
}
void QSoundEffect_SetSource(QSoundEffect* self, QUrl* url) {
	self->setSource(*url);
}
int QSoundEffect_LoopCount(const QSoundEffect* self) {
	return self->loopCount();
}
int QSoundEffect_LoopsRemaining(const QSoundEffect* self) {
	return self->loopsRemaining();
}
void QSoundEffect_SetLoopCount(QSoundEffect* self, int loopCount) {
	self->setLoopCount(static_cast<int>(loopCount));
}
QAudioDevice* QSoundEffect_AudioDevice(QSoundEffect* self) {
	return new QAudioDevice(self->audioDevice());
}
void QSoundEffect_SetAudioDevice(QSoundEffect* self, QAudioDevice* device) {
	self->setAudioDevice(*device);
}
float QSoundEffect_Volume(const QSoundEffect* self) {
	return self->volume();
}
void QSoundEffect_SetVolume(QSoundEffect* self, float volume) {
	self->setVolume(static_cast<float>(volume));
}
bool QSoundEffect_IsMuted(const QSoundEffect* self) {
	return self->isMuted();
}
void QSoundEffect_SetMuted(QSoundEffect* self, bool muted) {
	self->setMuted(muted);
}
bool QSoundEffect_IsLoaded(const QSoundEffect* self) {
	return self->isLoaded();
}
bool QSoundEffect_IsPlaying(const QSoundEffect* self) {
	return self->isPlaying();
}
Status QSoundEffect_Status(const QSoundEffect* self) {
	return self->status();
}
void QSoundEffect_SourceChanged(QSoundEffect* self) {
	self->sourceChanged();
}
void QSoundEffect_connect_SourceChanged(QSoundEffect* self, intptr_t slot) {
	MiqtVirtualQSoundEffect::connect(self, static_cast<void (QSoundEffect::*)()>(&QSoundEffect::sourceChanged), self, [=]() {
		miqt_exec_callback_QSoundEffect_SourceChanged(slot);
	});
}
void QSoundEffect_LoopCountChanged(QSoundEffect* self) {
	self->loopCountChanged();
}
void QSoundEffect_connect_LoopCountChanged(QSoundEffect* self, intptr_t slot) {
	MiqtVirtualQSoundEffect::connect(self, static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopCountChanged), self, [=]() {
		miqt_exec_callback_QSoundEffect_LoopCountChanged(slot);
	});
}
void QSoundEffect_LoopsRemainingChanged(QSoundEffect* self) {
	self->loopsRemainingChanged();
}
void QSoundEffect_connect_LoopsRemainingChanged(QSoundEffect* self, intptr_t slot) {
	MiqtVirtualQSoundEffect::connect(self, static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loopsRemainingChanged), self, [=]() {
		miqt_exec_callback_QSoundEffect_LoopsRemainingChanged(slot);
	});
}
void QSoundEffect_VolumeChanged(QSoundEffect* self) {
	self->volumeChanged();
}
void QSoundEffect_connect_VolumeChanged(QSoundEffect* self, intptr_t slot) {
	MiqtVirtualQSoundEffect::connect(self, static_cast<void (QSoundEffect::*)()>(&QSoundEffect::volumeChanged), self, [=]() {
		miqt_exec_callback_QSoundEffect_VolumeChanged(slot);
	});
}
void QSoundEffect_MutedChanged(QSoundEffect* self) {
	self->mutedChanged();
}
void QSoundEffect_connect_MutedChanged(QSoundEffect* self, intptr_t slot) {
	MiqtVirtualQSoundEffect::connect(self, static_cast<void (QSoundEffect::*)()>(&QSoundEffect::mutedChanged), self, [=]() {
		miqt_exec_callback_QSoundEffect_MutedChanged(slot);
	});
}
void QSoundEffect_LoadedChanged(QSoundEffect* self) {
	self->loadedChanged();
}
void QSoundEffect_connect_LoadedChanged(QSoundEffect* self, intptr_t slot) {
	MiqtVirtualQSoundEffect::connect(self, static_cast<void (QSoundEffect::*)()>(&QSoundEffect::loadedChanged), self, [=]() {
		miqt_exec_callback_QSoundEffect_LoadedChanged(slot);
	});
}
void QSoundEffect_PlayingChanged(QSoundEffect* self) {
	self->playingChanged();
}
void QSoundEffect_connect_PlayingChanged(QSoundEffect* self, intptr_t slot) {
	MiqtVirtualQSoundEffect::connect(self, static_cast<void (QSoundEffect::*)()>(&QSoundEffect::playingChanged), self, [=]() {
		miqt_exec_callback_QSoundEffect_PlayingChanged(slot);
	});
}
void QSoundEffect_StatusChanged(QSoundEffect* self) {
	self->statusChanged();
}
void QSoundEffect_connect_StatusChanged(QSoundEffect* self, intptr_t slot) {
	MiqtVirtualQSoundEffect::connect(self, static_cast<void (QSoundEffect::*)()>(&QSoundEffect::statusChanged), self, [=]() {
		miqt_exec_callback_QSoundEffect_StatusChanged(slot);
	});
}
void QSoundEffect_AudioDeviceChanged(QSoundEffect* self) {
	self->audioDeviceChanged();
}
void QSoundEffect_connect_AudioDeviceChanged(QSoundEffect* self, intptr_t slot) {
	MiqtVirtualQSoundEffect::connect(self, static_cast<void (QSoundEffect::*)()>(&QSoundEffect::audioDeviceChanged), self, [=]() {
		miqt_exec_callback_QSoundEffect_AudioDeviceChanged(slot);
	});
}
void QSoundEffect_Play(QSoundEffect* self) {
	self->play();
}
void QSoundEffect_Stop(QSoundEffect* self) {
	self->stop();
}
struct miqt_string QSoundEffect_Tr2(const char* s, const char* c) {
	QString _ret = QSoundEffect::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QSoundEffect_Tr3(const char* s, const char* c, int n) {
	QString _ret = QSoundEffect::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSoundEffect_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSoundEffect*>( (QSoundEffect*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QSoundEffect_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQSoundEffect*)(self) )->virtualbase_MetaObject();
}
void QSoundEffect_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSoundEffect*>( (QSoundEffect*)(self) )->handle__Metacast = slot;
}
void* QSoundEffect_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQSoundEffect*)(self) )->virtualbase_Metacast(param1);
}
void QSoundEffect_Delete(QSoundEffect* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQSoundEffect*>( self );
	} else {
		delete self;
	}
}
