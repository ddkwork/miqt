// +build ignore

#include <QMap>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWebChannel>
#include <QWebChannelAbstractTransport>
#include <qwebchannel.h>
#include "gen_qwebchannel.h"
class MiqtVirtualQWebChannel : public virtual QWebChannel {
public:
MiqtVirtualQWebChannel(): QWebChannel() {};
MiqtVirtualQWebChannel(QObject* parent): QWebChannel(parent) {};

virtual ~MiqtVirtualQWebChannel() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QWebChannel::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QWebChannel_MetaObject(const_cast<MiqtVirtualQWebChannel*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QWebChannel::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QWebChannel::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QWebChannel_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QWebChannel::qt_metacast(param1);

	}
};
QWebChannel* QWebChannel_new() {
return new MiqtVirtualQWebChannel();
}
QWebChannel* QWebChannel_new2(QObject* parent) {
return new MiqtVirtualQWebChannel(parent);
}
void QWebChannel_virtbase(QWebChannel* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QWebChannel_MetaObject(const QWebChannel* self) {
	return (QMetaObject*) self->metaObject();
}
void* QWebChannel_Metacast(QWebChannel* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QWebChannel_Tr(const char* s) {
	QString _ret = QWebChannel::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QWebChannel_RegisterObjects(QWebChannel* self, struct miqt_map /* of struct miqt_string to QObject* */  objects) {
	QHash<QString, QObject *> objects_QMap;
	objects_QMap.reserve(objects.len);
	struct miqt_string* objects_karr = static_cast<struct miqt_string*>(objects.keys);
	QObject** objects_varr = static_cast<QObject**>(objects.values);
	for(size_t i = 0; i < objects.len; ++i) {
		QString objects_karr_i_QString = QString::fromUtf8(objects_karr[i].data, objects_karr[i].len);
		objects_QMap[objects_karr_i_QString] = objects_varr[i];
	}
	self->registerObjects(objects_QMap);
}
struct miqt_map /* of struct miqt_string to QObject* */  QWebChannel_RegisteredObjects(const QWebChannel* self) {
	QHash<QString, QObject *> _ret = self->registeredObjects();
	// Convert QMap<> from C++ memory to manually-managed C memory
	struct miqt_string* _karr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.size()));
	QObject** _varr = static_cast<QObject**>(malloc(sizeof(QObject*) * _ret.size()));
	int _ctr = 0;
	for (auto _itr = _ret.keyValueBegin(); _itr != _ret.keyValueEnd(); ++_itr) {
		QString _hashkey_ret = _itr->first;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _hashkey_b = _hashkey_ret.toUtf8();
		struct miqt_string _hashkey_ms;
		_hashkey_ms.len = _hashkey_b.length();
		_hashkey_ms.data = static_cast<char*>(malloc(_hashkey_ms.len));
		memcpy(_hashkey_ms.data, _hashkey_b.data(), _hashkey_ms.len);
		_karr[_ctr] = _hashkey_ms;
		_varr[_ctr] = _itr->second;
		_ctr++;
	}
	struct miqt_map _out;
	_out.len = _ret.size();
	_out.keys = static_cast<void*>(_karr);
	_out.values = static_cast<void*>(_varr);
	return _out;
}
void QWebChannel_RegisterObject(QWebChannel* self, struct miqt_string id, QObject* object) {
	QString id_QString = QString::fromUtf8(id.data, id.len);
	self->registerObject(id_QString, object);
}
void QWebChannel_DeregisterObject(QWebChannel* self, QObject* object) {
	self->deregisterObject(object);
}
bool QWebChannel_BlockUpdates(const QWebChannel* self) {
	return self->blockUpdates();
}
void QWebChannel_SetBlockUpdates(QWebChannel* self, bool block) {
	self->setBlockUpdates(block);
}
int QWebChannel_PropertyUpdateInterval(const QWebChannel* self) {
	return self->propertyUpdateInterval();
}
void QWebChannel_SetPropertyUpdateInterval(QWebChannel* self, int ms) {
	self->setPropertyUpdateInterval(static_cast<int>(ms));
}
void QWebChannel_BlockUpdatesChanged(QWebChannel* self, bool block) {
	self->blockUpdatesChanged(block);
}
void QWebChannel_connect_BlockUpdatesChanged(QWebChannel* self, intptr_t slot) {
	MiqtVirtualQWebChannel::connect(self, static_cast<void (QWebChannel::*)(bool)>(&QWebChannel::blockUpdatesChanged), self, [=](bool block) {
		bool sigval1 = block;
		miqt_exec_callback_QWebChannel_BlockUpdatesChanged(slot, sigval1);
	});
}
void QWebChannel_ConnectTo(QWebChannel* self, QWebChannelAbstractTransport* transport) {
	self->connectTo(transport);
}
void QWebChannel_DisconnectFrom(QWebChannel* self, QWebChannelAbstractTransport* transport) {
	self->disconnectFrom(transport);
}
struct miqt_string QWebChannel_Tr2(const char* s, const char* c) {
	QString _ret = QWebChannel::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QWebChannel_Tr3(const char* s, const char* c, int n) {
	QString _ret = QWebChannel::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QWebChannel_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWebChannel*>( (QWebChannel*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QWebChannel_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQWebChannel*)(self) )->virtualbase_MetaObject();
}
void QWebChannel_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWebChannel*>( (QWebChannel*)(self) )->handle__Metacast = slot;
}
void* QWebChannel_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQWebChannel*)(self) )->virtualbase_Metacast(param1);
}
void QWebChannel_Delete(QWebChannel* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQWebChannel*>( self );
	} else {
		delete self;
	}
}
