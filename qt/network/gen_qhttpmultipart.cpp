// +build ignore

#include <QByteArray>
#include <QHttpMultiPart>
#include <QHttpPart>
#include <QIODevice>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVariant>
#include <qhttpmultipart.h>
#include "gen_qhttpmultipart.h"

QHttpPart* QHttpPart_new() {
	return new QHttpPart();
}

QHttpPart* QHttpPart_new2(QHttpPart* other) {
	return new QHttpPart(*other);
}

void QHttpPart_OperatorAssign(QHttpPart* self, QHttpPart* other) {
	self->operator=(*other);
}

void QHttpPart_Swap(QHttpPart* self, QHttpPart* other) {
	self->swap(*other);
}

bool QHttpPart_OperatorEqual(const QHttpPart* self, QHttpPart* other) {
	return (*self == *other);
}

bool QHttpPart_OperatorNotEqual(const QHttpPart* self, QHttpPart* other) {
	return (*self != *other);
}

void QHttpPart_SetHeader(QHttpPart* self, int header, QVariant* value) {
	self->setHeader(static_cast<QNetworkRequest::KnownHeaders>(header), *value);
}

void QHttpPart_SetRawHeader(QHttpPart* self, struct miqt_string headerName, struct miqt_string headerValue) {
	QByteArray headerName_QByteArray(headerName.data, headerName.len);
	QByteArray headerValue_QByteArray(headerValue.data, headerValue.len);
	self->setRawHeader(headerName_QByteArray, headerValue_QByteArray);
}

void QHttpPart_SetBody(QHttpPart* self, struct miqt_string body) {
	QByteArray body_QByteArray(body.data, body.len);
	self->setBody(body_QByteArray);
}

void QHttpPart_SetBodyDevice(QHttpPart* self, QIODevice* device) {
	self->setBodyDevice(device);
}

void QHttpPart_Delete(QHttpPart* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QHttpPart*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQHttpMultiPart : public virtual QHttpMultiPart {
public:

	MiqtVirtualQHttpMultiPart(): QHttpMultiPart() {};
	MiqtVirtualQHttpMultiPart(ContentType contentType): QHttpMultiPart(contentType) {};
	MiqtVirtualQHttpMultiPart(QObject* parent): QHttpMultiPart(parent) {};
	MiqtVirtualQHttpMultiPart(ContentType contentType, QObject* parent): QHttpMultiPart(contentType, parent) {};

	virtual ~MiqtVirtualQHttpMultiPart() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QHttpMultiPart::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QHttpMultiPart_MetaObject(const_cast<MiqtVirtualQHttpMultiPart*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QHttpMultiPart::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QHttpMultiPart::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QHttpMultiPart_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QHttpMultiPart::qt_metacast(param1);

	}

};

QHttpMultiPart* QHttpMultiPart_new() {
	return new MiqtVirtualQHttpMultiPart();
}

QHttpMultiPart* QHttpMultiPart_new2(ContentType contentType) {
	return new MiqtVirtualQHttpMultiPart(contentType);
}

QHttpMultiPart* QHttpMultiPart_new3(QObject* parent) {
	return new MiqtVirtualQHttpMultiPart(parent);
}

QHttpMultiPart* QHttpMultiPart_new4(ContentType contentType, QObject* parent) {
	return new MiqtVirtualQHttpMultiPart(contentType, parent);
}

void QHttpMultiPart_virtbase(QHttpMultiPart* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QHttpMultiPart_MetaObject(const QHttpMultiPart* self) {
	return (QMetaObject*) self->metaObject();
}

void* QHttpMultiPart_Metacast(QHttpMultiPart* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QHttpMultiPart_Tr(const char* s) {
	QString _ret = QHttpMultiPart::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QHttpMultiPart_Append(QHttpMultiPart* self, QHttpPart* httpPart) {
	self->append(*httpPart);
}

void QHttpMultiPart_SetContentType(QHttpMultiPart* self, ContentType contentType) {
	self->setContentType(contentType);
}

struct miqt_string QHttpMultiPart_Boundary(const QHttpMultiPart* self) {
	QByteArray _qb = self->boundary();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

void QHttpMultiPart_SetBoundary(QHttpMultiPart* self, struct miqt_string boundary) {
	QByteArray boundary_QByteArray(boundary.data, boundary.len);
	self->setBoundary(boundary_QByteArray);
}

struct miqt_string QHttpMultiPart_Tr2(const char* s, const char* c) {
	QString _ret = QHttpMultiPart::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QHttpMultiPart_Tr3(const char* s, const char* c, int n) {
	QString _ret = QHttpMultiPart::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QHttpMultiPart_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQHttpMultiPart*>( (QHttpMultiPart*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QHttpMultiPart_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQHttpMultiPart*)(self) )->virtualbase_MetaObject();
}

void QHttpMultiPart_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQHttpMultiPart*>( (QHttpMultiPart*)(self) )->handle__Metacast = slot;
}

void* QHttpMultiPart_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQHttpMultiPart*)(self) )->virtualbase_Metacast(param1);
}

void QHttpMultiPart_Delete(QHttpMultiPart* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQHttpMultiPart*>( self );
	} else {
		delete self;
	}
}

