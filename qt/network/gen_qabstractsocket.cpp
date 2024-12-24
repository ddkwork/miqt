// +build ignore

#include <QAbstractSocket>
#include <QAuthenticator>
#include <QHostAddress>
#include <QIODevice>
#include <QIODeviceBase>
#include <QMetaObject>
#include <QNetworkProxy>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVariant>
#include <qabstractsocket.h>
#include "gen_qabstractsocket.h"

class MiqtVirtualQAbstractSocket : public virtual QAbstractSocket {
public:

	MiqtVirtualQAbstractSocket(SocketType socketType, QObject* parent): QAbstractSocket(socketType, parent) {};

	virtual ~MiqtVirtualQAbstractSocket() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QAbstractSocket::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QAbstractSocket_MetaObject(const_cast<MiqtVirtualQAbstractSocket*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QAbstractSocket::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QAbstractSocket::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QAbstractSocket_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QAbstractSocket::qt_metacast(param1);

	}

};

QAbstractSocket* QAbstractSocket_new(SocketType socketType, QObject* parent) {
	return new MiqtVirtualQAbstractSocket(socketType, parent);
}

void QAbstractSocket_virtbase(QAbstractSocket* src, QIODevice** outptr_QIODevice) {
	*outptr_QIODevice = static_cast<QIODevice*>(src);
}

QMetaObject* QAbstractSocket_MetaObject(const QAbstractSocket* self) {
	return (QMetaObject*) self->metaObject();
}

void* QAbstractSocket_Metacast(QAbstractSocket* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QAbstractSocket_Tr(const char* s) {
	QString _ret = QAbstractSocket::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAbstractSocket_Resume(QAbstractSocket* self) {
	self->resume();
}

PauseModes QAbstractSocket_PauseMode(const QAbstractSocket* self) {
	return self->pauseMode();
}

void QAbstractSocket_SetPauseMode(QAbstractSocket* self, PauseModes pauseMode) {
	self->setPauseMode(pauseMode);
}

bool QAbstractSocket_Bind(QAbstractSocket* self, QHostAddress* address, uint16_t port, BindMode mode) {
	return self->bind(*address, static_cast<quint16>(port), mode);
}

bool QAbstractSocket_Bind2(QAbstractSocket* self) {
	return self->bind();
}

void QAbstractSocket_ConnectToHost(QAbstractSocket* self, struct miqt_string hostName, uint16_t port, OpenMode mode, NetworkLayerProtocol protocol) {
	QString hostName_QString = QString::fromUtf8(hostName.data, hostName.len);
	self->connectToHost(hostName_QString, static_cast<quint16>(port), mode, protocol);
}

void QAbstractSocket_ConnectToHost2(QAbstractSocket* self, QHostAddress* address, uint16_t port) {
	self->connectToHost(*address, static_cast<quint16>(port));
}

void QAbstractSocket_DisconnectFromHost(QAbstractSocket* self) {
	self->disconnectFromHost();
}

bool QAbstractSocket_IsValid(const QAbstractSocket* self) {
	return self->isValid();
}

long long QAbstractSocket_BytesAvailable(const QAbstractSocket* self) {
	qint64 _ret = self->bytesAvailable();
	return static_cast<long long>(_ret);
}

long long QAbstractSocket_BytesToWrite(const QAbstractSocket* self) {
	qint64 _ret = self->bytesToWrite();
	return static_cast<long long>(_ret);
}

uint16_t QAbstractSocket_LocalPort(const QAbstractSocket* self) {
	quint16 _ret = self->localPort();
	return static_cast<uint16_t>(_ret);
}

QHostAddress* QAbstractSocket_LocalAddress(const QAbstractSocket* self) {
	return new QHostAddress(self->localAddress());
}

uint16_t QAbstractSocket_PeerPort(const QAbstractSocket* self) {
	quint16 _ret = self->peerPort();
	return static_cast<uint16_t>(_ret);
}

QHostAddress* QAbstractSocket_PeerAddress(const QAbstractSocket* self) {
	return new QHostAddress(self->peerAddress());
}

struct miqt_string QAbstractSocket_PeerName(const QAbstractSocket* self) {
	QString _ret = self->peerName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

long long QAbstractSocket_ReadBufferSize(const QAbstractSocket* self) {
	qint64 _ret = self->readBufferSize();
	return static_cast<long long>(_ret);
}

void QAbstractSocket_SetReadBufferSize(QAbstractSocket* self, long long size) {
	self->setReadBufferSize(static_cast<qint64>(size));
}

void QAbstractSocket_Abort(QAbstractSocket* self) {
	self->abort();
}

intptr_t QAbstractSocket_SocketDescriptor(const QAbstractSocket* self) {
	qintptr _ret = self->socketDescriptor();
	return (intptr_t)(_ret);
}

bool QAbstractSocket_SetSocketDescriptor(QAbstractSocket* self, intptr_t socketDescriptor, SocketState state, OpenMode openMode) {
	return self->setSocketDescriptor((qintptr)(socketDescriptor), state, openMode);
}

void QAbstractSocket_SetSocketOption(QAbstractSocket* self, int option, QVariant* value) {
	self->setSocketOption(static_cast<QAbstractSocket::SocketOption>(option), *value);
}

QVariant* QAbstractSocket_SocketOption(QAbstractSocket* self, int option) {
	return new QVariant(self->socketOption(static_cast<QAbstractSocket::SocketOption>(option)));
}

SocketType QAbstractSocket_SocketType(const QAbstractSocket* self) {
	return self->socketType();
}

SocketState QAbstractSocket_State(const QAbstractSocket* self) {
	return self->state();
}

SocketError QAbstractSocket_Error(const QAbstractSocket* self) {
	return self->error();
}

void QAbstractSocket_Close(QAbstractSocket* self) {
	self->close();
}

bool QAbstractSocket_IsSequential(const QAbstractSocket* self) {
	return self->isSequential();
}

bool QAbstractSocket_Flush(QAbstractSocket* self) {
	return self->flush();
}

bool QAbstractSocket_WaitForConnected(QAbstractSocket* self, int msecs) {
	return self->waitForConnected(static_cast<int>(msecs));
}

bool QAbstractSocket_WaitForReadyRead(QAbstractSocket* self, int msecs) {
	return self->waitForReadyRead(static_cast<int>(msecs));
}

bool QAbstractSocket_WaitForBytesWritten(QAbstractSocket* self, int msecs) {
	return self->waitForBytesWritten(static_cast<int>(msecs));
}

bool QAbstractSocket_WaitForDisconnected(QAbstractSocket* self, int msecs) {
	return self->waitForDisconnected(static_cast<int>(msecs));
}

void QAbstractSocket_SetProxy(QAbstractSocket* self, QNetworkProxy* networkProxy) {
	self->setProxy(*networkProxy);
}

QNetworkProxy* QAbstractSocket_Proxy(const QAbstractSocket* self) {
	return new QNetworkProxy(self->proxy());
}

struct miqt_string QAbstractSocket_ProtocolTag(const QAbstractSocket* self) {
	QString _ret = self->protocolTag();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAbstractSocket_SetProtocolTag(QAbstractSocket* self, struct miqt_string tag) {
	QString tag_QString = QString::fromUtf8(tag.data, tag.len);
	self->setProtocolTag(tag_QString);
}

void QAbstractSocket_HostFound(QAbstractSocket* self) {
	self->hostFound();
}

void QAbstractSocket_connect_HostFound(QAbstractSocket* self, intptr_t slot) {
	MiqtVirtualQAbstractSocket::connect(self, static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::hostFound), self, [=]() {
		miqt_exec_callback_QAbstractSocket_HostFound(slot);
	});
}

void QAbstractSocket_Connected(QAbstractSocket* self) {
	self->connected();
}

void QAbstractSocket_connect_Connected(QAbstractSocket* self, intptr_t slot) {
	MiqtVirtualQAbstractSocket::connect(self, static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::connected), self, [=]() {
		miqt_exec_callback_QAbstractSocket_Connected(slot);
	});
}

void QAbstractSocket_Disconnected(QAbstractSocket* self) {
	self->disconnected();
}

void QAbstractSocket_connect_Disconnected(QAbstractSocket* self, intptr_t slot) {
	MiqtVirtualQAbstractSocket::connect(self, static_cast<void (QAbstractSocket::*)()>(&QAbstractSocket::disconnected), self, [=]() {
		miqt_exec_callback_QAbstractSocket_Disconnected(slot);
	});
}

void QAbstractSocket_StateChanged(QAbstractSocket* self, int param1) {
	self->stateChanged(static_cast<QAbstractSocket::SocketState>(param1));
}

void QAbstractSocket_connect_StateChanged(QAbstractSocket* self, intptr_t slot) {
	MiqtVirtualQAbstractSocket::connect(self, static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketState)>(&QAbstractSocket::stateChanged), self, [=](QAbstractSocket::SocketState param1) {
		QAbstractSocket::SocketState param1_ret = param1;
		int sigval1 = static_cast<int>(param1_ret);
		miqt_exec_callback_QAbstractSocket_StateChanged(slot, sigval1);
	});
}

void QAbstractSocket_ErrorOccurred(QAbstractSocket* self, int param1) {
	self->errorOccurred(static_cast<QAbstractSocket::SocketError>(param1));
}

void QAbstractSocket_connect_ErrorOccurred(QAbstractSocket* self, intptr_t slot) {
	MiqtVirtualQAbstractSocket::connect(self, static_cast<void (QAbstractSocket::*)(QAbstractSocket::SocketError)>(&QAbstractSocket::errorOccurred), self, [=](QAbstractSocket::SocketError param1) {
		QAbstractSocket::SocketError param1_ret = param1;
		int sigval1 = static_cast<int>(param1_ret);
		miqt_exec_callback_QAbstractSocket_ErrorOccurred(slot, sigval1);
	});
}

void QAbstractSocket_ProxyAuthenticationRequired(QAbstractSocket* self, QNetworkProxy* proxy, QAuthenticator* authenticator) {
	self->proxyAuthenticationRequired(*proxy, authenticator);
}

void QAbstractSocket_connect_ProxyAuthenticationRequired(QAbstractSocket* self, intptr_t slot) {
	MiqtVirtualQAbstractSocket::connect(self, static_cast<void (QAbstractSocket::*)(const QNetworkProxy&, QAuthenticator*)>(&QAbstractSocket::proxyAuthenticationRequired), self, [=](const QNetworkProxy& proxy, QAuthenticator* authenticator) {
		const QNetworkProxy& proxy_ret = proxy;
		// Cast returned reference into pointer
		QNetworkProxy* sigval1 = const_cast<QNetworkProxy*>(&proxy_ret);
		QAuthenticator* sigval2 = authenticator;
		miqt_exec_callback_QAbstractSocket_ProxyAuthenticationRequired(slot, sigval1, sigval2);
	});
}

struct miqt_string QAbstractSocket_Tr2(const char* s, const char* c) {
	QString _ret = QAbstractSocket::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAbstractSocket_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAbstractSocket::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QAbstractSocket_Bind1(QAbstractSocket* self, uint16_t port) {
	return self->bind(static_cast<quint16>(port));
}

bool QAbstractSocket_Bind22(QAbstractSocket* self, uint16_t port, BindMode mode) {
	return self->bind(static_cast<quint16>(port), mode);
}

void QAbstractSocket_ConnectToHost3(QAbstractSocket* self, QHostAddress* address, uint16_t port, OpenMode mode) {
	self->connectToHost(*address, static_cast<quint16>(port), mode);
}

void QAbstractSocket_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAbstractSocket*>( (QAbstractSocket*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QAbstractSocket_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQAbstractSocket*)(self) )->virtualbase_MetaObject();
}

void QAbstractSocket_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAbstractSocket*>( (QAbstractSocket*)(self) )->handle__Metacast = slot;
}

void* QAbstractSocket_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQAbstractSocket*)(self) )->virtualbase_Metacast(param1);
}

void QAbstractSocket_Delete(QAbstractSocket* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQAbstractSocket*>( self );
	} else {
		delete self;
	}
}

