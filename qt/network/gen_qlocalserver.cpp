// +build ignore

#include <QLocalServer>
#include <QLocalSocket>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qlocalserver.h>
#include "gen_qlocalserver.h"
class MiqtVirtualQLocalServer : public virtual QLocalServer {
public:
MiqtVirtualQLocalServer(): QLocalServer() {};
MiqtVirtualQLocalServer(QObject* parent): QLocalServer(parent) {};

virtual ~MiqtVirtualQLocalServer() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QLocalServer::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QLocalServer_MetaObject(const_cast<MiqtVirtualQLocalServer*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QLocalServer::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QLocalServer::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QLocalServer_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QLocalServer::qt_metacast(param1);

	}
};
QLocalServer* QLocalServer_new() {
return new MiqtVirtualQLocalServer();
}
QLocalServer* QLocalServer_new2(QObject* parent) {
return new MiqtVirtualQLocalServer(parent);
}
void QLocalServer_virtbase(QLocalServer* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QLocalServer_MetaObject(const QLocalServer* self) {
	return (QMetaObject*) self->metaObject();
}
void* QLocalServer_Metacast(QLocalServer* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QLocalServer_Tr(const char* s) {
	QString _ret = QLocalServer::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QLocalServer_NewConnection(QLocalServer* self) {
	self->newConnection();
}
void QLocalServer_connect_NewConnection(QLocalServer* self, intptr_t slot) {
	MiqtVirtualQLocalServer::connect(self, static_cast<void (QLocalServer::*)()>(&QLocalServer::newConnection), self, [=]() {
		miqt_exec_callback_QLocalServer_NewConnection(slot);
	});
}
void QLocalServer_Close(QLocalServer* self) {
	self->close();
}
struct miqt_string QLocalServer_ErrorString(const QLocalServer* self) {
	QString _ret = self->errorString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QLocalServer_HasPendingConnections(const QLocalServer* self) {
	return self->hasPendingConnections();
}
bool QLocalServer_IsListening(const QLocalServer* self) {
	return self->isListening();
}
bool QLocalServer_Listen(QLocalServer* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return self->listen(name_QString);
}
bool QLocalServer_ListenWithSocketDescriptor(QLocalServer* self, intptr_t socketDescriptor) {
	return self->listen((qintptr)(socketDescriptor));
}
int QLocalServer_MaxPendingConnections(const QLocalServer* self) {
	return self->maxPendingConnections();
}
QLocalSocket* QLocalServer_NextPendingConnection(QLocalServer* self) {
	return self->nextPendingConnection();
}
struct miqt_string QLocalServer_ServerName(const QLocalServer* self) {
	QString _ret = self->serverName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QLocalServer_FullServerName(const QLocalServer* self) {
	QString _ret = self->fullServerName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QLocalServer_RemoveServer(struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return QLocalServer::removeServer(name_QString);
}
int QLocalServer_ServerError(const QLocalServer* self) {
	QAbstractSocket::SocketError _ret = self->serverError();
	return static_cast<int>(_ret);
}
void QLocalServer_SetMaxPendingConnections(QLocalServer* self, int numConnections) {
	self->setMaxPendingConnections(static_cast<int>(numConnections));
}
bool QLocalServer_WaitForNewConnection(QLocalServer* self) {
	return self->waitForNewConnection();
}
void QLocalServer_SetListenBacklogSize(QLocalServer* self, int size) {
	self->setListenBacklogSize(static_cast<int>(size));
}
int QLocalServer_ListenBacklogSize(const QLocalServer* self) {
	return self->listenBacklogSize();
}
void QLocalServer_SetSocketOptions(QLocalServer* self, SocketOptions options) {
	self->setSocketOptions(options);
}
SocketOptions QLocalServer_SocketOptions(const QLocalServer* self) {
	return self->socketOptions();
}
intptr_t QLocalServer_SocketDescriptor(const QLocalServer* self) {
	qintptr _ret = self->socketDescriptor();
	return (intptr_t)(_ret);
}
struct miqt_string QLocalServer_Tr2(const char* s, const char* c) {
	QString _ret = QLocalServer::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QLocalServer_Tr3(const char* s, const char* c, int n) {
	QString _ret = QLocalServer::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QLocalServer_WaitForNewConnection1(QLocalServer* self, int msec) {
	return self->waitForNewConnection(static_cast<int>(msec));
}
bool QLocalServer_WaitForNewConnection2(QLocalServer* self, int msec, bool* timedOut) {
	return self->waitForNewConnection(static_cast<int>(msec), timedOut);
}
void QLocalServer_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLocalServer*>( (QLocalServer*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QLocalServer_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQLocalServer*)(self) )->virtualbase_MetaObject();
}
void QLocalServer_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLocalServer*>( (QLocalServer*)(self) )->handle__Metacast = slot;
}
void* QLocalServer_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQLocalServer*)(self) )->virtualbase_Metacast(param1);
}
void QLocalServer_Delete(QLocalServer* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQLocalServer*>( self );
	} else {
		delete self;
	}
}
