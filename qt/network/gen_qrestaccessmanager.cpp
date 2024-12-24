// +build ignore

#include <QByteArray>
#include <QHttpMultiPart>
#include <QIODevice>
#include <QJsonDocument>
#include <QMap>
#include <QMetaObject>
#include <QNetworkAccessManager>
#include <QNetworkReply>
#include <QNetworkRequest>
#include <QObject>
#include <QRestAccessManager>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVariant>
#include <qrestaccessmanager.h>
#include "gen_qrestaccessmanager.h"
class MiqtVirtualQRestAccessManager : public virtual QRestAccessManager {
public:
MiqtVirtualQRestAccessManager(QNetworkAccessManager* manager): QRestAccessManager(manager) {};
MiqtVirtualQRestAccessManager(QNetworkAccessManager* manager, QObject* parent): QRestAccessManager(manager, parent) {};

virtual ~MiqtVirtualQRestAccessManager() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QRestAccessManager::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QRestAccessManager_MetaObject(const_cast<MiqtVirtualQRestAccessManager*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QRestAccessManager::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QRestAccessManager::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QRestAccessManager_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QRestAccessManager::qt_metacast(param1);

	}
};
QRestAccessManager* QRestAccessManager_new(QNetworkAccessManager* manager) {
return new MiqtVirtualQRestAccessManager(manager);
}
QRestAccessManager* QRestAccessManager_new2(QNetworkAccessManager* manager, QObject* parent) {
return new MiqtVirtualQRestAccessManager(manager, parent);
}
void QRestAccessManager_virtbase(QRestAccessManager* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QRestAccessManager_MetaObject(const QRestAccessManager* self) {
	return (QMetaObject*) self->metaObject();
}
void* QRestAccessManager_Metacast(QRestAccessManager* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QRestAccessManager_Tr(const char* s) {
	QString _ret = QRestAccessManager::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QNetworkAccessManager* QRestAccessManager_NetworkAccessManager(const QRestAccessManager* self) {
	return self->networkAccessManager();
}
QNetworkReply* QRestAccessManager_DeleteResource(QRestAccessManager* self, QNetworkRequest* request) {
	return self->deleteResource(*request);
}
QNetworkReply* QRestAccessManager_Head(QRestAccessManager* self, QNetworkRequest* request) {
	return self->head(*request);
}
QNetworkReply* QRestAccessManager_Get(QRestAccessManager* self, QNetworkRequest* request) {
	return self->get(*request);
}
QNetworkReply* QRestAccessManager_Get2(QRestAccessManager* self, QNetworkRequest* request, struct miqt_string data) {
	QByteArray data_QByteArray(data.data, data.len);
	return self->get(*request, data_QByteArray);
}
QNetworkReply* QRestAccessManager_Get3(QRestAccessManager* self, QNetworkRequest* request, QJsonDocument* data) {
	return self->get(*request, *data);
}
QNetworkReply* QRestAccessManager_Get4(QRestAccessManager* self, QNetworkRequest* request, QIODevice* data) {
	return self->get(*request, data);
}
QNetworkReply* QRestAccessManager_Post(QRestAccessManager* self, QNetworkRequest* request, QJsonDocument* data) {
	return self->post(*request, *data);
}
QNetworkReply* QRestAccessManager_Post2(QRestAccessManager* self, QNetworkRequest* request, struct miqt_map /* of struct miqt_string to QVariant* */  data) {
	QVariantMap data_QMap;
	struct miqt_string* data_karr = static_cast<struct miqt_string*>(data.keys);
	QVariant** data_varr = static_cast<QVariant**>(data.values);
	for(size_t i = 0; i < data.len; ++i) {
		QString data_karr_i_QString = QString::fromUtf8(data_karr[i].data, data_karr[i].len);
		data_QMap[data_karr_i_QString] = *(data_varr[i]);
	}
	return self->post(*request, data_QMap);
}
QNetworkReply* QRestAccessManager_Post3(QRestAccessManager* self, QNetworkRequest* request, struct miqt_string data) {
	QByteArray data_QByteArray(data.data, data.len);
	return self->post(*request, data_QByteArray);
}
QNetworkReply* QRestAccessManager_Post4(QRestAccessManager* self, QNetworkRequest* request, QHttpMultiPart* data) {
	return self->post(*request, data);
}
QNetworkReply* QRestAccessManager_Post5(QRestAccessManager* self, QNetworkRequest* request, QIODevice* data) {
	return self->post(*request, data);
}
QNetworkReply* QRestAccessManager_Put(QRestAccessManager* self, QNetworkRequest* request, QJsonDocument* data) {
	return self->put(*request, *data);
}
QNetworkReply* QRestAccessManager_Put2(QRestAccessManager* self, QNetworkRequest* request, struct miqt_map /* of struct miqt_string to QVariant* */  data) {
	QVariantMap data_QMap;
	struct miqt_string* data_karr = static_cast<struct miqt_string*>(data.keys);
	QVariant** data_varr = static_cast<QVariant**>(data.values);
	for(size_t i = 0; i < data.len; ++i) {
		QString data_karr_i_QString = QString::fromUtf8(data_karr[i].data, data_karr[i].len);
		data_QMap[data_karr_i_QString] = *(data_varr[i]);
	}
	return self->put(*request, data_QMap);
}
QNetworkReply* QRestAccessManager_Put3(QRestAccessManager* self, QNetworkRequest* request, struct miqt_string data) {
	QByteArray data_QByteArray(data.data, data.len);
	return self->put(*request, data_QByteArray);
}
QNetworkReply* QRestAccessManager_Put4(QRestAccessManager* self, QNetworkRequest* request, QHttpMultiPart* data) {
	return self->put(*request, data);
}
QNetworkReply* QRestAccessManager_Put5(QRestAccessManager* self, QNetworkRequest* request, QIODevice* data) {
	return self->put(*request, data);
}
QNetworkReply* QRestAccessManager_Patch(QRestAccessManager* self, QNetworkRequest* request, QJsonDocument* data) {
	return self->patch(*request, *data);
}
QNetworkReply* QRestAccessManager_Patch2(QRestAccessManager* self, QNetworkRequest* request, struct miqt_map /* of struct miqt_string to QVariant* */  data) {
	QVariantMap data_QMap;
	struct miqt_string* data_karr = static_cast<struct miqt_string*>(data.keys);
	QVariant** data_varr = static_cast<QVariant**>(data.values);
	for(size_t i = 0; i < data.len; ++i) {
		QString data_karr_i_QString = QString::fromUtf8(data_karr[i].data, data_karr[i].len);
		data_QMap[data_karr_i_QString] = *(data_varr[i]);
	}
	return self->patch(*request, data_QMap);
}
QNetworkReply* QRestAccessManager_Patch3(QRestAccessManager* self, QNetworkRequest* request, struct miqt_string data) {
	QByteArray data_QByteArray(data.data, data.len);
	return self->patch(*request, data_QByteArray);
}
QNetworkReply* QRestAccessManager_Patch4(QRestAccessManager* self, QNetworkRequest* request, QIODevice* data) {
	return self->patch(*request, data);
}
QNetworkReply* QRestAccessManager_SendCustomRequest(QRestAccessManager* self, QNetworkRequest* request, struct miqt_string method, struct miqt_string data) {
	QByteArray method_QByteArray(method.data, method.len);
	QByteArray data_QByteArray(data.data, data.len);
	return self->sendCustomRequest(*request, method_QByteArray, data_QByteArray);
}
QNetworkReply* QRestAccessManager_SendCustomRequest2(QRestAccessManager* self, QNetworkRequest* request, struct miqt_string method, QIODevice* data) {
	QByteArray method_QByteArray(method.data, method.len);
	return self->sendCustomRequest(*request, method_QByteArray, data);
}
QNetworkReply* QRestAccessManager_SendCustomRequest3(QRestAccessManager* self, QNetworkRequest* request, struct miqt_string method, QHttpMultiPart* data) {
	QByteArray method_QByteArray(method.data, method.len);
	return self->sendCustomRequest(*request, method_QByteArray, data);
}
struct miqt_string QRestAccessManager_Tr2(const char* s, const char* c) {
	QString _ret = QRestAccessManager::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QRestAccessManager_Tr3(const char* s, const char* c, int n) {
	QString _ret = QRestAccessManager::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QRestAccessManager_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRestAccessManager*>( (QRestAccessManager*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QRestAccessManager_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQRestAccessManager*)(self) )->virtualbase_MetaObject();
}
void QRestAccessManager_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRestAccessManager*>( (QRestAccessManager*)(self) )->handle__Metacast = slot;
}
void* QRestAccessManager_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQRestAccessManager*)(self) )->virtualbase_Metacast(param1);
}
void QRestAccessManager_Delete(QRestAccessManager* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQRestAccessManager*>( self );
	} else {
		delete self;
	}
}
