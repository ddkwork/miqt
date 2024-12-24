// +build ignore

#include <QAbstractSocket>
#include <QByteArray>
#include <QHostAddress>
#include <QIODevice>
#include <QIODeviceBase>
#include <QMetaObject>
#include <QNetworkDatagram>
#include <QNetworkInterface>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QUdpSocket>
#include <qudpsocket.h>
#include "gen_qudpsocket.h"
class MiqtVirtualQUdpSocket : public virtual QUdpSocket {
public:
MiqtVirtualQUdpSocket(): QUdpSocket() {};
MiqtVirtualQUdpSocket(QObject* parent): QUdpSocket(parent) {};

virtual ~MiqtVirtualQUdpSocket() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QUdpSocket::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QUdpSocket_MetaObject(const_cast<MiqtVirtualQUdpSocket*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QUdpSocket::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QUdpSocket::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QUdpSocket_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QUdpSocket::qt_metacast(param1);

	}
};
QUdpSocket* QUdpSocket_new() {
return new MiqtVirtualQUdpSocket();
}
QUdpSocket* QUdpSocket_new2(QObject* parent) {
return new MiqtVirtualQUdpSocket(parent);
}
void QUdpSocket_virtbase(QUdpSocket* src
, QAbstractSocket** outptr_QAbstractSocket
) {
*outptr_QAbstractSocket = static_cast<QAbstractSocket*>(src);
}
QMetaObject* QUdpSocket_MetaObject(const QUdpSocket* self) {
	return (QMetaObject*) self->metaObject();
}
void* QUdpSocket_Metacast(QUdpSocket* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QUdpSocket_Tr(const char* s) {
	QString _ret = QUdpSocket::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QUdpSocket_Bind(QUdpSocket* self, int addr) {
	return self->bind(static_cast<QHostAddress::SpecialAddress>(addr));
}
bool QUdpSocket_JoinMulticastGroup(QUdpSocket* self, QHostAddress* groupAddress) {
	return self->joinMulticastGroup(*groupAddress);
}
bool QUdpSocket_JoinMulticastGroup2(QUdpSocket* self, QHostAddress* groupAddress, QNetworkInterface* iface) {
	return self->joinMulticastGroup(*groupAddress, *iface);
}
bool QUdpSocket_LeaveMulticastGroup(QUdpSocket* self, QHostAddress* groupAddress) {
	return self->leaveMulticastGroup(*groupAddress);
}
bool QUdpSocket_LeaveMulticastGroup2(QUdpSocket* self, QHostAddress* groupAddress, QNetworkInterface* iface) {
	return self->leaveMulticastGroup(*groupAddress, *iface);
}
QNetworkInterface* QUdpSocket_MulticastInterface(const QUdpSocket* self) {
	return new QNetworkInterface(self->multicastInterface());
}
void QUdpSocket_SetMulticastInterface(QUdpSocket* self, QNetworkInterface* iface) {
	self->setMulticastInterface(*iface);
}
bool QUdpSocket_HasPendingDatagrams(const QUdpSocket* self) {
	return self->hasPendingDatagrams();
}
long long QUdpSocket_PendingDatagramSize(const QUdpSocket* self) {
	qint64 _ret = self->pendingDatagramSize();
	return static_cast<long long>(_ret);
}
QNetworkDatagram* QUdpSocket_ReceiveDatagram(QUdpSocket* self) {
	return new QNetworkDatagram(self->receiveDatagram());
}
long long QUdpSocket_ReadDatagram(QUdpSocket* self, char* data, long long maxlen) {
	qint64 _ret = self->readDatagram(data, static_cast<qint64>(maxlen));
	return static_cast<long long>(_ret);
}
long long QUdpSocket_WriteDatagram(QUdpSocket* self, QNetworkDatagram* datagram) {
	qint64 _ret = self->writeDatagram(*datagram);
	return static_cast<long long>(_ret);
}
long long QUdpSocket_WriteDatagram2(QUdpSocket* self, const char* data, long long lenVal, QHostAddress* host, uint16_t port) {
	qint64 _ret = self->writeDatagram(data, static_cast<qint64>(lenVal), *host, static_cast<quint16>(port));
	return static_cast<long long>(_ret);
}
long long QUdpSocket_WriteDatagram3(QUdpSocket* self, struct miqt_string datagram, QHostAddress* host, uint16_t port) {
	QByteArray datagram_QByteArray(datagram.data, datagram.len);
	qint64 _ret = self->writeDatagram(datagram_QByteArray, *host, static_cast<quint16>(port));
	return static_cast<long long>(_ret);
}
struct miqt_string QUdpSocket_Tr2(const char* s, const char* c) {
	QString _ret = QUdpSocket::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QUdpSocket_Tr3(const char* s, const char* c, int n) {
	QString _ret = QUdpSocket::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QUdpSocket_Bind2(QUdpSocket* self, int addr, uint16_t port) {
	return self->bind(static_cast<QHostAddress::SpecialAddress>(addr), static_cast<quint16>(port));
}
bool QUdpSocket_Bind3(QUdpSocket* self, int addr, uint16_t port, BindMode mode) {
	return self->bind(static_cast<QHostAddress::SpecialAddress>(addr), static_cast<quint16>(port), mode);
}
QNetworkDatagram* QUdpSocket_ReceiveDatagram1(QUdpSocket* self, long long maxSize) {
	return new QNetworkDatagram(self->receiveDatagram(static_cast<qint64>(maxSize)));
}
long long QUdpSocket_ReadDatagram3(QUdpSocket* self, char* data, long long maxlen, QHostAddress* host) {
	qint64 _ret = self->readDatagram(data, static_cast<qint64>(maxlen), host);
	return static_cast<long long>(_ret);
}
long long QUdpSocket_ReadDatagram4(QUdpSocket* self, char* data, long long maxlen, QHostAddress* host, uint16_t* port) {
	qint64 _ret = self->readDatagram(data, static_cast<qint64>(maxlen), host, static_cast<quint16*>(port));
	return static_cast<long long>(_ret);
}
void QUdpSocket_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQUdpSocket*>( (QUdpSocket*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QUdpSocket_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQUdpSocket*)(self) )->virtualbase_MetaObject();
}
void QUdpSocket_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQUdpSocket*>( (QUdpSocket*)(self) )->handle__Metacast = slot;
}
void* QUdpSocket_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQUdpSocket*)(self) )->virtualbase_Metacast(param1);
}
void QUdpSocket_Delete(QUdpSocket* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQUdpSocket*>( self );
	} else {
		delete self;
	}
}
