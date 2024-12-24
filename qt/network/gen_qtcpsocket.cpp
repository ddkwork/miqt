// +build ignore

#include <QAbstractSocket>
#include <QIODevice>
#include <QIODeviceBase>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTcpSocket>
#include <qtcpsocket.h>
#include "gen_qtcpsocket.h"

class MiqtVirtualQTcpSocket : public virtual QTcpSocket {
public:

	MiqtVirtualQTcpSocket(): QTcpSocket() {};
	MiqtVirtualQTcpSocket(QObject* parent): QTcpSocket(parent) {};

	virtual ~MiqtVirtualQTcpSocket() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QTcpSocket::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QTcpSocket_MetaObject(const_cast<MiqtVirtualQTcpSocket*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QTcpSocket::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QTcpSocket::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QTcpSocket_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QTcpSocket::qt_metacast(param1);

	}

};

QTcpSocket* QTcpSocket_new() {
	return new MiqtVirtualQTcpSocket();
}

QTcpSocket* QTcpSocket_new2(QObject* parent) {
	return new MiqtVirtualQTcpSocket(parent);
}

void QTcpSocket_virtbase(QTcpSocket* src, QAbstractSocket** outptr_QAbstractSocket) {
	*outptr_QAbstractSocket = static_cast<QAbstractSocket*>(src);
}

QMetaObject* QTcpSocket_MetaObject(const QTcpSocket* self) {
	return (QMetaObject*) self->metaObject();
}

void* QTcpSocket_Metacast(QTcpSocket* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QTcpSocket_Tr(const char* s) {
	QString _ret = QTcpSocket::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QTcpSocket_Bind(QTcpSocket* self, int addr) {
	return self->bind(static_cast<QHostAddress::SpecialAddress>(addr));
}

struct miqt_string QTcpSocket_Tr2(const char* s, const char* c) {
	QString _ret = QTcpSocket::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QTcpSocket_Tr3(const char* s, const char* c, int n) {
	QString _ret = QTcpSocket::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QTcpSocket_Bind2(QTcpSocket* self, int addr, uint16_t port) {
	return self->bind(static_cast<QHostAddress::SpecialAddress>(addr), static_cast<quint16>(port));
}

bool QTcpSocket_Bind3(QTcpSocket* self, int addr, uint16_t port, BindMode mode) {
	return self->bind(static_cast<QHostAddress::SpecialAddress>(addr), static_cast<quint16>(port), mode);
}

void QTcpSocket_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTcpSocket*>( (QTcpSocket*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QTcpSocket_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQTcpSocket*)(self) )->virtualbase_MetaObject();
}

void QTcpSocket_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTcpSocket*>( (QTcpSocket*)(self) )->handle__Metacast = slot;
}

void* QTcpSocket_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQTcpSocket*)(self) )->virtualbase_Metacast(param1);
}

void QTcpSocket_Delete(QTcpSocket* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQTcpSocket*>( self );
	} else {
		delete self;
	}
}

