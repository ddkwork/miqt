// +build ignore

#include <QJsonObject>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWebChannelAbstractTransport>
#include <qwebchannelabstracttransport.h>
#include "gen_qwebchannelabstracttransport.h"
class MiqtVirtualQWebChannelAbstractTransport : public virtual QWebChannelAbstractTransport {
public:
MiqtVirtualQWebChannelAbstractTransport(): QWebChannelAbstractTransport() {};
MiqtVirtualQWebChannelAbstractTransport(QObject* parent): QWebChannelAbstractTransport(parent) {};

virtual ~MiqtVirtualQWebChannelAbstractTransport() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QWebChannelAbstractTransport::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QWebChannelAbstractTransport_MetaObject(const_cast<MiqtVirtualQWebChannelAbstractTransport*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QWebChannelAbstractTransport::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QWebChannelAbstractTransport::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QWebChannelAbstractTransport_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QWebChannelAbstractTransport::qt_metacast(param1);

	}
};
QWebChannelAbstractTransport* QWebChannelAbstractTransport_new() {
return new MiqtVirtualQWebChannelAbstractTransport();
}
QWebChannelAbstractTransport* QWebChannelAbstractTransport_new2(QObject* parent) {
return new MiqtVirtualQWebChannelAbstractTransport(parent);
}
void QWebChannelAbstractTransport_virtbase(QWebChannelAbstractTransport* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QWebChannelAbstractTransport_MetaObject(const QWebChannelAbstractTransport* self) {
	return (QMetaObject*) self->metaObject();
}
void* QWebChannelAbstractTransport_Metacast(QWebChannelAbstractTransport* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QWebChannelAbstractTransport_Tr(const char* s) {
	QString _ret = QWebChannelAbstractTransport::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QWebChannelAbstractTransport_SendMessage(QWebChannelAbstractTransport* self, QJsonObject* message) {
	self->sendMessage(*message);
}
void QWebChannelAbstractTransport_MessageReceived(QWebChannelAbstractTransport* self, QJsonObject* message, QWebChannelAbstractTransport* transport) {
	self->messageReceived(*message, transport);
}
void QWebChannelAbstractTransport_connect_MessageReceived(QWebChannelAbstractTransport* self, intptr_t slot) {
	MiqtVirtualQWebChannelAbstractTransport::connect(self, static_cast<void (QWebChannelAbstractTransport::*)(const QJsonObject&, QWebChannelAbstractTransport*)>(&QWebChannelAbstractTransport::messageReceived), self, [=](const QJsonObject& message, QWebChannelAbstractTransport* transport) {
		const QJsonObject& message_ret = message;
		// Cast returned reference into pointer
		QJsonObject* sigval1 = const_cast<QJsonObject*>(&message_ret);
		QWebChannelAbstractTransport* sigval2 = transport;
		miqt_exec_callback_QWebChannelAbstractTransport_MessageReceived(slot, sigval1, sigval2);
	});
}
struct miqt_string QWebChannelAbstractTransport_Tr2(const char* s, const char* c) {
	QString _ret = QWebChannelAbstractTransport::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QWebChannelAbstractTransport_Tr3(const char* s, const char* c, int n) {
	QString _ret = QWebChannelAbstractTransport::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QWebChannelAbstractTransport_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWebChannelAbstractTransport*>( (QWebChannelAbstractTransport*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QWebChannelAbstractTransport_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQWebChannelAbstractTransport*)(self) )->virtualbase_MetaObject();
}
void QWebChannelAbstractTransport_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWebChannelAbstractTransport*>( (QWebChannelAbstractTransport*)(self) )->handle__Metacast = slot;
}
void* QWebChannelAbstractTransport_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQWebChannelAbstractTransport*)(self) )->virtualbase_Metacast(param1);
}
void QWebChannelAbstractTransport_Delete(QWebChannelAbstractTransport* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQWebChannelAbstractTransport*>( self );
	} else {
		delete self;
	}
}
