// +build ignore

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QHttpMultiPart>
#include <QIODevice>
#include <QJsonDocument>
#include <QMap>
#include <QMetaMethod>
#include <QMetaObject>
#include <QNetworkAccessManager>
#include <QNetworkReply>
#include <QNetworkRequest>
#include <QObject>
#include <QRestAccessManager>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimerEvent>
#include <QVariant>
#include <qrestaccessmanager.h>
#include "gen_qrestaccessmanager.h"

#ifndef _Bool
#define _Bool bool
#endif

void _GUID_Delete(_GUID* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<_GUID*>( self );
	} else {
		delete self;
	}
}

void type_info_Delete(type_info* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<type_info*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQRestAccessManager : public virtual QRestAccessManager {
public:

	MiqtVirtualQRestAccessManager(QNetworkAccessManager* manager): QRestAccessManager(manager) {};
	MiqtVirtualQRestAccessManager(QNetworkAccessManager* manager, QObject* parent): QRestAccessManager(manager, parent) {};

	virtual ~MiqtVirtualQRestAccessManager() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Event = 0;

	// Subclass to allow providing a Go implementation
	virtual bool event(QEvent* event) override {
		if (handle__Event == 0) {
			return QRestAccessManager::event(event);
		}
		
		QEvent* sigval1 = event;

		bool callback_return_value = miqt_exec_callback_QRestAccessManager_Event(this, handle__Event, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_Event(QEvent* event) {

		return QRestAccessManager::event(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__EventFilter = 0;

	// Subclass to allow providing a Go implementation
	virtual bool eventFilter(QObject* watched, QEvent* event) override {
		if (handle__EventFilter == 0) {
			return QRestAccessManager::eventFilter(watched, event);
		}
		
		QObject* sigval1 = watched;
		QEvent* sigval2 = event;

		bool callback_return_value = miqt_exec_callback_QRestAccessManager_EventFilter(this, handle__EventFilter, sigval1, sigval2);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_EventFilter(QObject* watched, QEvent* event) {

		return QRestAccessManager::eventFilter(watched, event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__TimerEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void timerEvent(QTimerEvent* event) override {
		if (handle__TimerEvent == 0) {
			QRestAccessManager::timerEvent(event);
			return;
		}
		
		QTimerEvent* sigval1 = event;

		miqt_exec_callback_QRestAccessManager_TimerEvent(this, handle__TimerEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_TimerEvent(QTimerEvent* event) {

		QRestAccessManager::timerEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ChildEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void childEvent(QChildEvent* event) override {
		if (handle__ChildEvent == 0) {
			QRestAccessManager::childEvent(event);
			return;
		}
		
		QChildEvent* sigval1 = event;

		miqt_exec_callback_QRestAccessManager_ChildEvent(this, handle__ChildEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ChildEvent(QChildEvent* event) {

		QRestAccessManager::childEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__CustomEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void customEvent(QEvent* event) override {
		if (handle__CustomEvent == 0) {
			QRestAccessManager::customEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QRestAccessManager_CustomEvent(this, handle__CustomEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_CustomEvent(QEvent* event) {

		QRestAccessManager::customEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ConnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void connectNotify(const QMetaMethod& signal) override {
		if (handle__ConnectNotify == 0) {
			QRestAccessManager::connectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QRestAccessManager_ConnectNotify(this, handle__ConnectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ConnectNotify(QMetaMethod* signal) {

		QRestAccessManager::connectNotify(*signal);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DisconnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void disconnectNotify(const QMetaMethod& signal) override {
		if (handle__DisconnectNotify == 0) {
			QRestAccessManager::disconnectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QRestAccessManager_DisconnectNotify(this, handle__DisconnectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DisconnectNotify(QMetaMethod* signal) {

		QRestAccessManager::disconnectNotify(*signal);

	}

};

QRestAccessManager* QRestAccessManager_new(QNetworkAccessManager* manager) {
	return new MiqtVirtualQRestAccessManager(manager);
}

QRestAccessManager* QRestAccessManager_new2(QNetworkAccessManager* manager, QObject* parent) {
	return new MiqtVirtualQRestAccessManager(manager, parent);
}

void QRestAccessManager_virtbase(QRestAccessManager* src, QObject** outptr_QObject) {
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

void QRestAccessManager_override_virtual_Event(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRestAccessManager*>( (QRestAccessManager*)(self) )->handle__Event = slot;
}

bool QRestAccessManager_virtualbase_Event(void* self, QEvent* event) {
	return ( (MiqtVirtualQRestAccessManager*)(self) )->virtualbase_Event(event);
}

void QRestAccessManager_override_virtual_EventFilter(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRestAccessManager*>( (QRestAccessManager*)(self) )->handle__EventFilter = slot;
}

bool QRestAccessManager_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event) {
	return ( (MiqtVirtualQRestAccessManager*)(self) )->virtualbase_EventFilter(watched, event);
}

void QRestAccessManager_override_virtual_TimerEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRestAccessManager*>( (QRestAccessManager*)(self) )->handle__TimerEvent = slot;
}

void QRestAccessManager_virtualbase_TimerEvent(void* self, QTimerEvent* event) {
	( (MiqtVirtualQRestAccessManager*)(self) )->virtualbase_TimerEvent(event);
}

void QRestAccessManager_override_virtual_ChildEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRestAccessManager*>( (QRestAccessManager*)(self) )->handle__ChildEvent = slot;
}

void QRestAccessManager_virtualbase_ChildEvent(void* self, QChildEvent* event) {
	( (MiqtVirtualQRestAccessManager*)(self) )->virtualbase_ChildEvent(event);
}

void QRestAccessManager_override_virtual_CustomEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRestAccessManager*>( (QRestAccessManager*)(self) )->handle__CustomEvent = slot;
}

void QRestAccessManager_virtualbase_CustomEvent(void* self, QEvent* event) {
	( (MiqtVirtualQRestAccessManager*)(self) )->virtualbase_CustomEvent(event);
}

void QRestAccessManager_override_virtual_ConnectNotify(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRestAccessManager*>( (QRestAccessManager*)(self) )->handle__ConnectNotify = slot;
}

void QRestAccessManager_virtualbase_ConnectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQRestAccessManager*)(self) )->virtualbase_ConnectNotify(signal);
}

void QRestAccessManager_override_virtual_DisconnectNotify(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRestAccessManager*>( (QRestAccessManager*)(self) )->handle__DisconnectNotify = slot;
}

void QRestAccessManager_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQRestAccessManager*)(self) )->virtualbase_DisconnectNotify(signal);
}

void QRestAccessManager_Delete(QRestAccessManager* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQRestAccessManager*>( self );
	} else {
		delete self;
	}
}

