// +build ignore

#include <QIODevice>
#include <QIODeviceBase>
#include <QLocalSocket>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qlocalsocket.h>
#include "gen_qlocalsocket.h"

class MiqtVirtualQLocalSocket : public virtual QLocalSocket {
public:

	MiqtVirtualQLocalSocket(): QLocalSocket() {};
	MiqtVirtualQLocalSocket(QObject* parent): QLocalSocket(parent) {};

	virtual ~MiqtVirtualQLocalSocket() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QLocalSocket::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QLocalSocket_MetaObject(const_cast<MiqtVirtualQLocalSocket*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QLocalSocket::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QLocalSocket::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QLocalSocket_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QLocalSocket::qt_metacast(param1);

	}

};

QLocalSocket* QLocalSocket_new() {
	return new MiqtVirtualQLocalSocket();
}

QLocalSocket* QLocalSocket_new2(QObject* parent) {
	return new MiqtVirtualQLocalSocket(parent);
}

void QLocalSocket_virtbase(QLocalSocket* src, QIODevice** outptr_QIODevice) {
	*outptr_QIODevice = static_cast<QIODevice*>(src);
}

QMetaObject* QLocalSocket_MetaObject(const QLocalSocket* self) {
	return (QMetaObject*) self->metaObject();
}

void* QLocalSocket_Metacast(QLocalSocket* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QLocalSocket_Tr(const char* s) {
	QString _ret = QLocalSocket::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QLocalSocket_ConnectToServer(QLocalSocket* self) {
	self->connectToServer();
}

void QLocalSocket_ConnectToServerWithName(QLocalSocket* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	self->connectToServer(name_QString);
}

void QLocalSocket_DisconnectFromServer(QLocalSocket* self) {
	self->disconnectFromServer();
}

void QLocalSocket_SetServerName(QLocalSocket* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	self->setServerName(name_QString);
}

struct miqt_string QLocalSocket_ServerName(const QLocalSocket* self) {
	QString _ret = self->serverName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocalSocket_FullServerName(const QLocalSocket* self) {
	QString _ret = self->fullServerName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QLocalSocket_Abort(QLocalSocket* self) {
	self->abort();
}

bool QLocalSocket_IsSequential(const QLocalSocket* self) {
	return self->isSequential();
}

long long QLocalSocket_BytesAvailable(const QLocalSocket* self) {
	qint64 _ret = self->bytesAvailable();
	return static_cast<long long>(_ret);
}

long long QLocalSocket_BytesToWrite(const QLocalSocket* self) {
	qint64 _ret = self->bytesToWrite();
	return static_cast<long long>(_ret);
}

bool QLocalSocket_CanReadLine(const QLocalSocket* self) {
	return self->canReadLine();
}

bool QLocalSocket_Open(QLocalSocket* self, OpenMode openMode) {
	return self->open(openMode);
}

void QLocalSocket_Close(QLocalSocket* self) {
	self->close();
}

LocalSocketError QLocalSocket_Error(const QLocalSocket* self) {
	return self->error();
}

bool QLocalSocket_Flush(QLocalSocket* self) {
	return self->flush();
}

bool QLocalSocket_IsValid(const QLocalSocket* self) {
	return self->isValid();
}

long long QLocalSocket_ReadBufferSize(const QLocalSocket* self) {
	qint64 _ret = self->readBufferSize();
	return static_cast<long long>(_ret);
}

void QLocalSocket_SetReadBufferSize(QLocalSocket* self, long long size) {
	self->setReadBufferSize(static_cast<qint64>(size));
}

bool QLocalSocket_SetSocketDescriptor(QLocalSocket* self, intptr_t socketDescriptor) {
	return self->setSocketDescriptor((qintptr)(socketDescriptor));
}

intptr_t QLocalSocket_SocketDescriptor(const QLocalSocket* self) {
	qintptr _ret = self->socketDescriptor();
	return (intptr_t)(_ret);
}

void QLocalSocket_SetSocketOptions(QLocalSocket* self, SocketOptions option) {
	self->setSocketOptions(option);
}

SocketOptions QLocalSocket_SocketOptions(const QLocalSocket* self) {
	return self->socketOptions();
}

LocalSocketState QLocalSocket_State(const QLocalSocket* self) {
	return self->state();
}

bool QLocalSocket_WaitForBytesWritten(QLocalSocket* self, int msecs) {
	return self->waitForBytesWritten(static_cast<int>(msecs));
}

bool QLocalSocket_WaitForConnected(QLocalSocket* self) {
	return self->waitForConnected();
}

bool QLocalSocket_WaitForDisconnected(QLocalSocket* self) {
	return self->waitForDisconnected();
}

bool QLocalSocket_WaitForReadyRead(QLocalSocket* self, int msecs) {
	return self->waitForReadyRead(static_cast<int>(msecs));
}

void QLocalSocket_Connected(QLocalSocket* self) {
	self->connected();
}

void QLocalSocket_connect_Connected(QLocalSocket* self, intptr_t slot) {
	MiqtVirtualQLocalSocket::connect(self, static_cast<void (QLocalSocket::*)()>(&QLocalSocket::connected), self, [=]() {
		miqt_exec_callback_QLocalSocket_Connected(slot);
	});
}

void QLocalSocket_Disconnected(QLocalSocket* self) {
	self->disconnected();
}

void QLocalSocket_connect_Disconnected(QLocalSocket* self, intptr_t slot) {
	MiqtVirtualQLocalSocket::connect(self, static_cast<void (QLocalSocket::*)()>(&QLocalSocket::disconnected), self, [=]() {
		miqt_exec_callback_QLocalSocket_Disconnected(slot);
	});
}

void QLocalSocket_ErrorOccurred(QLocalSocket* self, int socketError) {
	self->errorOccurred(static_cast<QLocalSocket::LocalSocketError>(socketError));
}

void QLocalSocket_connect_ErrorOccurred(QLocalSocket* self, intptr_t slot) {
	MiqtVirtualQLocalSocket::connect(self, static_cast<void (QLocalSocket::*)(QLocalSocket::LocalSocketError)>(&QLocalSocket::errorOccurred), self, [=](QLocalSocket::LocalSocketError socketError) {
		QLocalSocket::LocalSocketError socketError_ret = socketError;
		int sigval1 = static_cast<int>(socketError_ret);
		miqt_exec_callback_QLocalSocket_ErrorOccurred(slot, sigval1);
	});
}

void QLocalSocket_StateChanged(QLocalSocket* self, int socketState) {
	self->stateChanged(static_cast<QLocalSocket::LocalSocketState>(socketState));
}

void QLocalSocket_connect_StateChanged(QLocalSocket* self, intptr_t slot) {
	MiqtVirtualQLocalSocket::connect(self, static_cast<void (QLocalSocket::*)(QLocalSocket::LocalSocketState)>(&QLocalSocket::stateChanged), self, [=](QLocalSocket::LocalSocketState socketState) {
		QLocalSocket::LocalSocketState socketState_ret = socketState;
		int sigval1 = static_cast<int>(socketState_ret);
		miqt_exec_callback_QLocalSocket_StateChanged(slot, sigval1);
	});
}

struct miqt_string QLocalSocket_Tr2(const char* s, const char* c) {
	QString _ret = QLocalSocket::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLocalSocket_Tr3(const char* s, const char* c, int n) {
	QString _ret = QLocalSocket::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QLocalSocket_ConnectToServer1(QLocalSocket* self, OpenMode openMode) {
	self->connectToServer(openMode);
}

void QLocalSocket_ConnectToServer2(QLocalSocket* self, struct miqt_string name, OpenMode openMode) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	self->connectToServer(name_QString, openMode);
}

bool QLocalSocket_SetSocketDescriptor2(QLocalSocket* self, intptr_t socketDescriptor, LocalSocketState socketState) {
	return self->setSocketDescriptor((qintptr)(socketDescriptor), socketState);
}

bool QLocalSocket_SetSocketDescriptor3(QLocalSocket* self, intptr_t socketDescriptor, LocalSocketState socketState, OpenMode openMode) {
	return self->setSocketDescriptor((qintptr)(socketDescriptor), socketState, openMode);
}

bool QLocalSocket_WaitForConnected1(QLocalSocket* self, int msecs) {
	return self->waitForConnected(static_cast<int>(msecs));
}

bool QLocalSocket_WaitForDisconnected1(QLocalSocket* self, int msecs) {
	return self->waitForDisconnected(static_cast<int>(msecs));
}

void QLocalSocket_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLocalSocket*>( (QLocalSocket*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QLocalSocket_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQLocalSocket*)(self) )->virtualbase_MetaObject();
}

void QLocalSocket_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLocalSocket*>( (QLocalSocket*)(self) )->handle__Metacast = slot;
}

void* QLocalSocket_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQLocalSocket*)(self) )->virtualbase_Metacast(param1);
}

void QLocalSocket_Delete(QLocalSocket* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQLocalSocket*>( self );
	} else {
		delete self;
	}
}

