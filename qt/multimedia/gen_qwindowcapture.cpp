// +build ignore

#include <QCapturableWindow>
#include <QList>
#include <QMediaCaptureSession>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWindowCapture>
#include <qwindowcapture.h>
#include "gen_qwindowcapture.h"

class MiqtVirtualQWindowCapture : public virtual QWindowCapture {
public:

	MiqtVirtualQWindowCapture(): QWindowCapture() {};
	MiqtVirtualQWindowCapture(QObject* parent): QWindowCapture(parent) {};

	virtual ~MiqtVirtualQWindowCapture() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QWindowCapture::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QWindowCapture_MetaObject(const_cast<MiqtVirtualQWindowCapture*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QWindowCapture::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QWindowCapture::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QWindowCapture_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QWindowCapture::qt_metacast(param1);

	}

};

QWindowCapture* QWindowCapture_new() {
	return new MiqtVirtualQWindowCapture();
}

QWindowCapture* QWindowCapture_new2(QObject* parent) {
	return new MiqtVirtualQWindowCapture(parent);
}

void QWindowCapture_virtbase(QWindowCapture* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QWindowCapture_MetaObject(const QWindowCapture* self) {
	return (QMetaObject*) self->metaObject();
}

void* QWindowCapture_Metacast(QWindowCapture* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QWindowCapture_Tr(const char* s) {
	QString _ret = QWindowCapture::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_array /* of QCapturableWindow* */  QWindowCapture_CapturableWindows() {
	QList<QCapturableWindow> _ret = QWindowCapture::capturableWindows();
	// Convert QList<> from C++ memory to manually-managed C memory
	QCapturableWindow** _arr = static_cast<QCapturableWindow**>(malloc(sizeof(QCapturableWindow*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QCapturableWindow(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

QMediaCaptureSession* QWindowCapture_CaptureSession(const QWindowCapture* self) {
	return self->captureSession();
}

void QWindowCapture_SetWindow(QWindowCapture* self, QCapturableWindow* window) {
	self->setWindow(*window);
}

QCapturableWindow* QWindowCapture_Window(const QWindowCapture* self) {
	return new QCapturableWindow(self->window());
}

bool QWindowCapture_IsActive(const QWindowCapture* self) {
	return self->isActive();
}

Error QWindowCapture_Error(const QWindowCapture* self) {
	return self->error();
}

struct miqt_string QWindowCapture_ErrorString(const QWindowCapture* self) {
	QString _ret = self->errorString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QWindowCapture_SetActive(QWindowCapture* self, bool active) {
	self->setActive(active);
}

void QWindowCapture_Start(QWindowCapture* self) {
	self->start();
}

void QWindowCapture_Stop(QWindowCapture* self) {
	self->stop();
}

void QWindowCapture_ActiveChanged(QWindowCapture* self, bool param1) {
	self->activeChanged(param1);
}

void QWindowCapture_connect_ActiveChanged(QWindowCapture* self, intptr_t slot) {
	MiqtVirtualQWindowCapture::connect(self, static_cast<void (QWindowCapture::*)(bool)>(&QWindowCapture::activeChanged), self, [=](bool param1) {
		bool sigval1 = param1;
		miqt_exec_callback_QWindowCapture_ActiveChanged(slot, sigval1);
	});
}

void QWindowCapture_WindowChanged(QWindowCapture* self, QCapturableWindow* window) {
	self->windowChanged(*window);
}

void QWindowCapture_connect_WindowChanged(QWindowCapture* self, intptr_t slot) {
	MiqtVirtualQWindowCapture::connect(self, static_cast<void (QWindowCapture::*)(QCapturableWindow)>(&QWindowCapture::windowChanged), self, [=](QCapturableWindow window) {
		QCapturableWindow* sigval1 = new QCapturableWindow(window);
		miqt_exec_callback_QWindowCapture_WindowChanged(slot, sigval1);
	});
}

void QWindowCapture_ErrorChanged(QWindowCapture* self) {
	self->errorChanged();
}

void QWindowCapture_connect_ErrorChanged(QWindowCapture* self, intptr_t slot) {
	MiqtVirtualQWindowCapture::connect(self, static_cast<void (QWindowCapture::*)()>(&QWindowCapture::errorChanged), self, [=]() {
		miqt_exec_callback_QWindowCapture_ErrorChanged(slot);
	});
}

void QWindowCapture_ErrorOccurred(QWindowCapture* self, int error, struct miqt_string errorString) {
	QString errorString_QString = QString::fromUtf8(errorString.data, errorString.len);
	self->errorOccurred(static_cast<QWindowCapture::Error>(error), errorString_QString);
}

void QWindowCapture_connect_ErrorOccurred(QWindowCapture* self, intptr_t slot) {
	MiqtVirtualQWindowCapture::connect(self, static_cast<void (QWindowCapture::*)(QWindowCapture::Error, const QString&)>(&QWindowCapture::errorOccurred), self, [=](QWindowCapture::Error error, const QString& errorString) {
		QWindowCapture::Error error_ret = error;
		int sigval1 = static_cast<int>(error_ret);
		const QString errorString_ret = errorString;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray errorString_b = errorString_ret.toUtf8();
		struct miqt_string errorString_ms;
		errorString_ms.len = errorString_b.length();
		errorString_ms.data = static_cast<char*>(malloc(errorString_ms.len));
		memcpy(errorString_ms.data, errorString_b.data(), errorString_ms.len);
		struct miqt_string sigval2 = errorString_ms;
		miqt_exec_callback_QWindowCapture_ErrorOccurred(slot, sigval1, sigval2);
	});
}

struct miqt_string QWindowCapture_Tr2(const char* s, const char* c) {
	QString _ret = QWindowCapture::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QWindowCapture_Tr3(const char* s, const char* c, int n) {
	QString _ret = QWindowCapture::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QWindowCapture_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWindowCapture*>( (QWindowCapture*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QWindowCapture_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQWindowCapture*)(self) )->virtualbase_MetaObject();
}

void QWindowCapture_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWindowCapture*>( (QWindowCapture*)(self) )->handle__Metacast = slot;
}

void* QWindowCapture_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQWindowCapture*)(self) )->virtualbase_Metacast(param1);
}

void QWindowCapture_Delete(QWindowCapture* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQWindowCapture*>( self );
	} else {
		delete self;
	}
}

