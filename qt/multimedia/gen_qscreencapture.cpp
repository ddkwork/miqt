// +build ignore

#include <QMediaCaptureSession>
#include <QMetaObject>
#include <QObject>
#include <QScreen>
#include <QScreenCapture>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qscreencapture.h>
#include "gen_qscreencapture.h"

class MiqtVirtualQScreenCapture : public virtual QScreenCapture {
public:

	MiqtVirtualQScreenCapture(): QScreenCapture() {};
	MiqtVirtualQScreenCapture(QObject* parent): QScreenCapture(parent) {};

	virtual ~MiqtVirtualQScreenCapture() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QScreenCapture::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QScreenCapture_MetaObject(const_cast<MiqtVirtualQScreenCapture*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QScreenCapture::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QScreenCapture::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QScreenCapture_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QScreenCapture::qt_metacast(param1);

	}

};

QScreenCapture* QScreenCapture_new() {
	return new MiqtVirtualQScreenCapture();
}

QScreenCapture* QScreenCapture_new2(QObject* parent) {
	return new MiqtVirtualQScreenCapture(parent);
}

void QScreenCapture_virtbase(QScreenCapture* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QScreenCapture_MetaObject(const QScreenCapture* self) {
	return (QMetaObject*) self->metaObject();
}

void* QScreenCapture_Metacast(QScreenCapture* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QScreenCapture_Tr(const char* s) {
	QString _ret = QScreenCapture::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QMediaCaptureSession* QScreenCapture_CaptureSession(const QScreenCapture* self) {
	return self->captureSession();
}

void QScreenCapture_SetScreen(QScreenCapture* self, QScreen* screen) {
	self->setScreen(screen);
}

QScreen* QScreenCapture_Screen(const QScreenCapture* self) {
	return self->screen();
}

bool QScreenCapture_IsActive(const QScreenCapture* self) {
	return self->isActive();
}

Error QScreenCapture_Error(const QScreenCapture* self) {
	return self->error();
}

struct miqt_string QScreenCapture_ErrorString(const QScreenCapture* self) {
	QString _ret = self->errorString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QScreenCapture_SetActive(QScreenCapture* self, bool active) {
	self->setActive(active);
}

void QScreenCapture_Start(QScreenCapture* self) {
	self->start();
}

void QScreenCapture_Stop(QScreenCapture* self) {
	self->stop();
}

void QScreenCapture_ActiveChanged(QScreenCapture* self, bool param1) {
	self->activeChanged(param1);
}

void QScreenCapture_connect_ActiveChanged(QScreenCapture* self, intptr_t slot) {
	MiqtVirtualQScreenCapture::connect(self, static_cast<void (QScreenCapture::*)(bool)>(&QScreenCapture::activeChanged), self, [=](bool param1) {
		bool sigval1 = param1;
		miqt_exec_callback_QScreenCapture_ActiveChanged(slot, sigval1);
	});
}

void QScreenCapture_ErrorChanged(QScreenCapture* self) {
	self->errorChanged();
}

void QScreenCapture_connect_ErrorChanged(QScreenCapture* self, intptr_t slot) {
	MiqtVirtualQScreenCapture::connect(self, static_cast<void (QScreenCapture::*)()>(&QScreenCapture::errorChanged), self, [=]() {
		miqt_exec_callback_QScreenCapture_ErrorChanged(slot);
	});
}

void QScreenCapture_ScreenChanged(QScreenCapture* self, QScreen* param1) {
	self->screenChanged(param1);
}

void QScreenCapture_connect_ScreenChanged(QScreenCapture* self, intptr_t slot) {
	MiqtVirtualQScreenCapture::connect(self, static_cast<void (QScreenCapture::*)(QScreen*)>(&QScreenCapture::screenChanged), self, [=](QScreen* param1) {
		QScreen* sigval1 = param1;
		miqt_exec_callback_QScreenCapture_ScreenChanged(slot, sigval1);
	});
}

void QScreenCapture_ErrorOccurred(QScreenCapture* self, int error, struct miqt_string errorString) {
	QString errorString_QString = QString::fromUtf8(errorString.data, errorString.len);
	self->errorOccurred(static_cast<QScreenCapture::Error>(error), errorString_QString);
}

void QScreenCapture_connect_ErrorOccurred(QScreenCapture* self, intptr_t slot) {
	MiqtVirtualQScreenCapture::connect(self, static_cast<void (QScreenCapture::*)(QScreenCapture::Error, const QString&)>(&QScreenCapture::errorOccurred), self, [=](QScreenCapture::Error error, const QString& errorString) {
		QScreenCapture::Error error_ret = error;
		int sigval1 = static_cast<int>(error_ret);
		const QString errorString_ret = errorString;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray errorString_b = errorString_ret.toUtf8();
		struct miqt_string errorString_ms;
		errorString_ms.len = errorString_b.length();
		errorString_ms.data = static_cast<char*>(malloc(errorString_ms.len));
		memcpy(errorString_ms.data, errorString_b.data(), errorString_ms.len);
		struct miqt_string sigval2 = errorString_ms;
		miqt_exec_callback_QScreenCapture_ErrorOccurred(slot, sigval1, sigval2);
	});
}

struct miqt_string QScreenCapture_Tr2(const char* s, const char* c) {
	QString _ret = QScreenCapture::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QScreenCapture_Tr3(const char* s, const char* c, int n) {
	QString _ret = QScreenCapture::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QScreenCapture_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQScreenCapture*>( (QScreenCapture*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QScreenCapture_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQScreenCapture*)(self) )->virtualbase_MetaObject();
}

void QScreenCapture_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQScreenCapture*>( (QScreenCapture*)(self) )->handle__Metacast = slot;
}

void* QScreenCapture_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQScreenCapture*)(self) )->virtualbase_Metacast(param1);
}

void QScreenCapture_Delete(QScreenCapture* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQScreenCapture*>( self );
	} else {
		delete self;
	}
}

