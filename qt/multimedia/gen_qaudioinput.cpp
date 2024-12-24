// +build ignore

#include <QAudioDevice>
#include <QAudioInput>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qaudioinput.h>
#include "gen_qaudioinput.h"

class MiqtVirtualQAudioInput : public virtual QAudioInput {
public:

	MiqtVirtualQAudioInput(): QAudioInput() {};
	MiqtVirtualQAudioInput(const QAudioDevice& deviceInfo): QAudioInput(deviceInfo) {};
	MiqtVirtualQAudioInput(QObject* parent): QAudioInput(parent) {};
	MiqtVirtualQAudioInput(const QAudioDevice& deviceInfo, QObject* parent): QAudioInput(deviceInfo, parent) {};

	virtual ~MiqtVirtualQAudioInput() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QAudioInput::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QAudioInput_MetaObject(const_cast<MiqtVirtualQAudioInput*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QAudioInput::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QAudioInput::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QAudioInput_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QAudioInput::qt_metacast(param1);

	}

};

QAudioInput* QAudioInput_new() {
	return new MiqtVirtualQAudioInput();
}

QAudioInput* QAudioInput_new2(QAudioDevice* deviceInfo) {
	return new MiqtVirtualQAudioInput(*deviceInfo);
}

QAudioInput* QAudioInput_new3(QObject* parent) {
	return new MiqtVirtualQAudioInput(parent);
}

QAudioInput* QAudioInput_new4(QAudioDevice* deviceInfo, QObject* parent) {
	return new MiqtVirtualQAudioInput(*deviceInfo, parent);
}

void QAudioInput_virtbase(QAudioInput* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QAudioInput_MetaObject(const QAudioInput* self) {
	return (QMetaObject*) self->metaObject();
}

void* QAudioInput_Metacast(QAudioInput* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QAudioInput_Tr(const char* s) {
	QString _ret = QAudioInput::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QAudioDevice* QAudioInput_Device(const QAudioInput* self) {
	return new QAudioDevice(self->device());
}

float QAudioInput_Volume(const QAudioInput* self) {
	return self->volume();
}

bool QAudioInput_IsMuted(const QAudioInput* self) {
	return self->isMuted();
}

void QAudioInput_SetDevice(QAudioInput* self, QAudioDevice* device) {
	self->setDevice(*device);
}

void QAudioInput_SetVolume(QAudioInput* self, float volume) {
	self->setVolume(static_cast<float>(volume));
}

void QAudioInput_SetMuted(QAudioInput* self, bool muted) {
	self->setMuted(muted);
}

void QAudioInput_DeviceChanged(QAudioInput* self) {
	self->deviceChanged();
}

void QAudioInput_connect_DeviceChanged(QAudioInput* self, intptr_t slot) {
	MiqtVirtualQAudioInput::connect(self, static_cast<void (QAudioInput::*)()>(&QAudioInput::deviceChanged), self, [=]() {
		miqt_exec_callback_QAudioInput_DeviceChanged(slot);
	});
}

void QAudioInput_VolumeChanged(QAudioInput* self, float volume) {
	self->volumeChanged(static_cast<float>(volume));
}

void QAudioInput_connect_VolumeChanged(QAudioInput* self, intptr_t slot) {
	MiqtVirtualQAudioInput::connect(self, static_cast<void (QAudioInput::*)(float)>(&QAudioInput::volumeChanged), self, [=](float volume) {
		float sigval1 = volume;
		miqt_exec_callback_QAudioInput_VolumeChanged(slot, sigval1);
	});
}

void QAudioInput_MutedChanged(QAudioInput* self, bool muted) {
	self->mutedChanged(muted);
}

void QAudioInput_connect_MutedChanged(QAudioInput* self, intptr_t slot) {
	MiqtVirtualQAudioInput::connect(self, static_cast<void (QAudioInput::*)(bool)>(&QAudioInput::mutedChanged), self, [=](bool muted) {
		bool sigval1 = muted;
		miqt_exec_callback_QAudioInput_MutedChanged(slot, sigval1);
	});
}

struct miqt_string QAudioInput_Tr2(const char* s, const char* c) {
	QString _ret = QAudioInput::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAudioInput_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAudioInput::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAudioInput_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAudioInput*>( (QAudioInput*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QAudioInput_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQAudioInput*)(self) )->virtualbase_MetaObject();
}

void QAudioInput_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAudioInput*>( (QAudioInput*)(self) )->handle__Metacast = slot;
}

void* QAudioInput_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQAudioInput*)(self) )->virtualbase_Metacast(param1);
}

void QAudioInput_Delete(QAudioInput* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQAudioInput*>( self );
	} else {
		delete self;
	}
}

