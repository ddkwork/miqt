// +build ignore

#include <QAudioDevice>
#include <QCameraDevice>
#include <QList>
#include <QMediaDevices>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qmediadevices.h>
#include "gen_qmediadevices.h"

class MiqtVirtualQMediaDevices : public virtual QMediaDevices {
public:

	MiqtVirtualQMediaDevices(): QMediaDevices() {};
	MiqtVirtualQMediaDevices(QObject* parent): QMediaDevices(parent) {};

	virtual ~MiqtVirtualQMediaDevices() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QMediaDevices::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QMediaDevices_MetaObject(const_cast<MiqtVirtualQMediaDevices*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QMediaDevices::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QMediaDevices::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QMediaDevices_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QMediaDevices::qt_metacast(param1);

	}

};

QMediaDevices* QMediaDevices_new() {
	return new MiqtVirtualQMediaDevices();
}

QMediaDevices* QMediaDevices_new2(QObject* parent) {
	return new MiqtVirtualQMediaDevices(parent);
}

void QMediaDevices_virtbase(QMediaDevices* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QMediaDevices_MetaObject(const QMediaDevices* self) {
	return (QMetaObject*) self->metaObject();
}

void* QMediaDevices_Metacast(QMediaDevices* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QMediaDevices_Tr(const char* s) {
	QString _ret = QMediaDevices::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_array /* of QAudioDevice* */  QMediaDevices_AudioInputs() {
	QList<QAudioDevice> _ret = QMediaDevices::audioInputs();
	// Convert QList<> from C++ memory to manually-managed C memory
	QAudioDevice** _arr = static_cast<QAudioDevice**>(malloc(sizeof(QAudioDevice*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QAudioDevice(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of QAudioDevice* */  QMediaDevices_AudioOutputs() {
	QList<QAudioDevice> _ret = QMediaDevices::audioOutputs();
	// Convert QList<> from C++ memory to manually-managed C memory
	QAudioDevice** _arr = static_cast<QAudioDevice**>(malloc(sizeof(QAudioDevice*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QAudioDevice(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of QCameraDevice* */  QMediaDevices_VideoInputs() {
	QList<QCameraDevice> _ret = QMediaDevices::videoInputs();
	// Convert QList<> from C++ memory to manually-managed C memory
	QCameraDevice** _arr = static_cast<QCameraDevice**>(malloc(sizeof(QCameraDevice*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QCameraDevice(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

QAudioDevice* QMediaDevices_DefaultAudioInput() {
	return new QAudioDevice(QMediaDevices::defaultAudioInput());
}

QAudioDevice* QMediaDevices_DefaultAudioOutput() {
	return new QAudioDevice(QMediaDevices::defaultAudioOutput());
}

QCameraDevice* QMediaDevices_DefaultVideoInput() {
	return new QCameraDevice(QMediaDevices::defaultVideoInput());
}

void QMediaDevices_AudioInputsChanged(QMediaDevices* self) {
	self->audioInputsChanged();
}

void QMediaDevices_connect_AudioInputsChanged(QMediaDevices* self, intptr_t slot) {
	MiqtVirtualQMediaDevices::connect(self, static_cast<void (QMediaDevices::*)()>(&QMediaDevices::audioInputsChanged), self, [=]() {
		miqt_exec_callback_QMediaDevices_AudioInputsChanged(slot);
	});
}

void QMediaDevices_AudioOutputsChanged(QMediaDevices* self) {
	self->audioOutputsChanged();
}

void QMediaDevices_connect_AudioOutputsChanged(QMediaDevices* self, intptr_t slot) {
	MiqtVirtualQMediaDevices::connect(self, static_cast<void (QMediaDevices::*)()>(&QMediaDevices::audioOutputsChanged), self, [=]() {
		miqt_exec_callback_QMediaDevices_AudioOutputsChanged(slot);
	});
}

void QMediaDevices_VideoInputsChanged(QMediaDevices* self) {
	self->videoInputsChanged();
}

void QMediaDevices_connect_VideoInputsChanged(QMediaDevices* self, intptr_t slot) {
	MiqtVirtualQMediaDevices::connect(self, static_cast<void (QMediaDevices::*)()>(&QMediaDevices::videoInputsChanged), self, [=]() {
		miqt_exec_callback_QMediaDevices_VideoInputsChanged(slot);
	});
}

struct miqt_string QMediaDevices_Tr2(const char* s, const char* c) {
	QString _ret = QMediaDevices::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QMediaDevices_Tr3(const char* s, const char* c, int n) {
	QString _ret = QMediaDevices::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QMediaDevices_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQMediaDevices*>( (QMediaDevices*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QMediaDevices_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQMediaDevices*)(self) )->virtualbase_MetaObject();
}

void QMediaDevices_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQMediaDevices*>( (QMediaDevices*)(self) )->handle__Metacast = slot;
}

void* QMediaDevices_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQMediaDevices*)(self) )->virtualbase_Metacast(param1);
}

void QMediaDevices_Delete(QMediaDevices* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQMediaDevices*>( self );
	} else {
		delete self;
	}
}

