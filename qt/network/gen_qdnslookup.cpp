// +build ignore

#include <QByteArray>
#include <QChildEvent>
#include <QDnsDomainNameRecord>
#include <QDnsHostAddressRecord>
#include <QDnsLookup>
#include <QDnsMailExchangeRecord>
#include <QDnsServiceRecord>
#include <QDnsTextRecord>
#include <QDnsTlsAssociationRecord>
#include <QEvent>
#include <QHostAddress>
#include <QList>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QSslConfiguration>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimerEvent>
#include <qdnslookup.h>
#include "gen_qdnslookup.h"

#ifndef _Bool
#define _Bool bool
#endif

QDnsDomainNameRecord* QDnsDomainNameRecord_new() {
	return new QDnsDomainNameRecord();
}

QDnsDomainNameRecord* QDnsDomainNameRecord_new2(QDnsDomainNameRecord* other) {
	return new QDnsDomainNameRecord(*other);
}

void QDnsDomainNameRecord_OperatorAssign(QDnsDomainNameRecord* self, QDnsDomainNameRecord* other) {
	self->operator=(*other);
}

void QDnsDomainNameRecord_Swap(QDnsDomainNameRecord* self, QDnsDomainNameRecord* other) {
	self->swap(*other);
}

struct miqt_string QDnsDomainNameRecord_Name(const QDnsDomainNameRecord* self) {
	QString _ret = self->name();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

unsigned int QDnsDomainNameRecord_TimeToLive(const QDnsDomainNameRecord* self) {
	quint32 _ret = self->timeToLive();
	return static_cast<unsigned int>(_ret);
}

struct miqt_string QDnsDomainNameRecord_Value(const QDnsDomainNameRecord* self) {
	QString _ret = self->value();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QDnsDomainNameRecord_Delete(QDnsDomainNameRecord* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QDnsDomainNameRecord*>( self );
	} else {
		delete self;
	}
}

QDnsHostAddressRecord* QDnsHostAddressRecord_new() {
	return new QDnsHostAddressRecord();
}

QDnsHostAddressRecord* QDnsHostAddressRecord_new2(QDnsHostAddressRecord* other) {
	return new QDnsHostAddressRecord(*other);
}

void QDnsHostAddressRecord_OperatorAssign(QDnsHostAddressRecord* self, QDnsHostAddressRecord* other) {
	self->operator=(*other);
}

void QDnsHostAddressRecord_Swap(QDnsHostAddressRecord* self, QDnsHostAddressRecord* other) {
	self->swap(*other);
}

struct miqt_string QDnsHostAddressRecord_Name(const QDnsHostAddressRecord* self) {
	QString _ret = self->name();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

unsigned int QDnsHostAddressRecord_TimeToLive(const QDnsHostAddressRecord* self) {
	quint32 _ret = self->timeToLive();
	return static_cast<unsigned int>(_ret);
}

QHostAddress* QDnsHostAddressRecord_Value(const QDnsHostAddressRecord* self) {
	return new QHostAddress(self->value());
}

void QDnsHostAddressRecord_Delete(QDnsHostAddressRecord* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QDnsHostAddressRecord*>( self );
	} else {
		delete self;
	}
}

QDnsMailExchangeRecord* QDnsMailExchangeRecord_new() {
	return new QDnsMailExchangeRecord();
}

QDnsMailExchangeRecord* QDnsMailExchangeRecord_new2(QDnsMailExchangeRecord* other) {
	return new QDnsMailExchangeRecord(*other);
}

void QDnsMailExchangeRecord_OperatorAssign(QDnsMailExchangeRecord* self, QDnsMailExchangeRecord* other) {
	self->operator=(*other);
}

void QDnsMailExchangeRecord_Swap(QDnsMailExchangeRecord* self, QDnsMailExchangeRecord* other) {
	self->swap(*other);
}

struct miqt_string QDnsMailExchangeRecord_Exchange(const QDnsMailExchangeRecord* self) {
	QString _ret = self->exchange();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDnsMailExchangeRecord_Name(const QDnsMailExchangeRecord* self) {
	QString _ret = self->name();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

uint16_t QDnsMailExchangeRecord_Preference(const QDnsMailExchangeRecord* self) {
	quint16 _ret = self->preference();
	return static_cast<uint16_t>(_ret);
}

unsigned int QDnsMailExchangeRecord_TimeToLive(const QDnsMailExchangeRecord* self) {
	quint32 _ret = self->timeToLive();
	return static_cast<unsigned int>(_ret);
}

void QDnsMailExchangeRecord_Delete(QDnsMailExchangeRecord* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QDnsMailExchangeRecord*>( self );
	} else {
		delete self;
	}
}

QDnsServiceRecord* QDnsServiceRecord_new() {
	return new QDnsServiceRecord();
}

QDnsServiceRecord* QDnsServiceRecord_new2(QDnsServiceRecord* other) {
	return new QDnsServiceRecord(*other);
}

void QDnsServiceRecord_OperatorAssign(QDnsServiceRecord* self, QDnsServiceRecord* other) {
	self->operator=(*other);
}

void QDnsServiceRecord_Swap(QDnsServiceRecord* self, QDnsServiceRecord* other) {
	self->swap(*other);
}

struct miqt_string QDnsServiceRecord_Name(const QDnsServiceRecord* self) {
	QString _ret = self->name();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

uint16_t QDnsServiceRecord_Port(const QDnsServiceRecord* self) {
	quint16 _ret = self->port();
	return static_cast<uint16_t>(_ret);
}

uint16_t QDnsServiceRecord_Priority(const QDnsServiceRecord* self) {
	quint16 _ret = self->priority();
	return static_cast<uint16_t>(_ret);
}

struct miqt_string QDnsServiceRecord_Target(const QDnsServiceRecord* self) {
	QString _ret = self->target();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

unsigned int QDnsServiceRecord_TimeToLive(const QDnsServiceRecord* self) {
	quint32 _ret = self->timeToLive();
	return static_cast<unsigned int>(_ret);
}

uint16_t QDnsServiceRecord_Weight(const QDnsServiceRecord* self) {
	quint16 _ret = self->weight();
	return static_cast<uint16_t>(_ret);
}

void QDnsServiceRecord_Delete(QDnsServiceRecord* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QDnsServiceRecord*>( self );
	} else {
		delete self;
	}
}

QDnsTextRecord* QDnsTextRecord_new() {
	return new QDnsTextRecord();
}

QDnsTextRecord* QDnsTextRecord_new2(QDnsTextRecord* other) {
	return new QDnsTextRecord(*other);
}

void QDnsTextRecord_OperatorAssign(QDnsTextRecord* self, QDnsTextRecord* other) {
	self->operator=(*other);
}

void QDnsTextRecord_Swap(QDnsTextRecord* self, QDnsTextRecord* other) {
	self->swap(*other);
}

struct miqt_string QDnsTextRecord_Name(const QDnsTextRecord* self) {
	QString _ret = self->name();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

unsigned int QDnsTextRecord_TimeToLive(const QDnsTextRecord* self) {
	quint32 _ret = self->timeToLive();
	return static_cast<unsigned int>(_ret);
}

struct miqt_array /* of struct miqt_string */  QDnsTextRecord_Values(const QDnsTextRecord* self) {
	QList<QByteArray> _ret = self->values();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QByteArray _lv_qb = _ret[i];
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_qb.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_qb.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

void QDnsTextRecord_Delete(QDnsTextRecord* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QDnsTextRecord*>( self );
	} else {
		delete self;
	}
}

QDnsTlsAssociationRecord* QDnsTlsAssociationRecord_new() {
	return new QDnsTlsAssociationRecord();
}

QDnsTlsAssociationRecord* QDnsTlsAssociationRecord_new2(QDnsTlsAssociationRecord* other) {
	return new QDnsTlsAssociationRecord(*other);
}

void QDnsTlsAssociationRecord_OperatorAssign(QDnsTlsAssociationRecord* self, QDnsTlsAssociationRecord* other) {
	self->operator=(*other);
}

void QDnsTlsAssociationRecord_Swap(QDnsTlsAssociationRecord* self, QDnsTlsAssociationRecord* other) {
	self->swap(*other);
}

struct miqt_string QDnsTlsAssociationRecord_Name(const QDnsTlsAssociationRecord* self) {
	QString _ret = self->name();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

unsigned int QDnsTlsAssociationRecord_TimeToLive(const QDnsTlsAssociationRecord* self) {
	quint32 _ret = self->timeToLive();
	return static_cast<unsigned int>(_ret);
}

CertificateUsage QDnsTlsAssociationRecord_Usage(const QDnsTlsAssociationRecord* self) {
	return self->usage();
}

Selector QDnsTlsAssociationRecord_Selector(const QDnsTlsAssociationRecord* self) {
	return self->selector();
}

MatchingType QDnsTlsAssociationRecord_MatchType(const QDnsTlsAssociationRecord* self) {
	return self->matchType();
}

struct miqt_string QDnsTlsAssociationRecord_Value(const QDnsTlsAssociationRecord* self) {
	QByteArray _qb = self->value();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

void QDnsTlsAssociationRecord_Delete(QDnsTlsAssociationRecord* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QDnsTlsAssociationRecord*>( self );
	} else {
		delete self;
	}
}

class MiqtVirtualQDnsLookup : public virtual QDnsLookup {
public:

	MiqtVirtualQDnsLookup(): QDnsLookup() {};
	MiqtVirtualQDnsLookup(Type typeVal, const QString& name): QDnsLookup(typeVal, name) {};
	MiqtVirtualQDnsLookup(Type typeVal, const QString& name, const QHostAddress& nameserver): QDnsLookup(typeVal, name, nameserver) {};
	MiqtVirtualQDnsLookup(Type typeVal, const QString& name, const QHostAddress& nameserver, quint16 port): QDnsLookup(typeVal, name, nameserver, port) {};
	MiqtVirtualQDnsLookup(Type typeVal, const QString& name, Protocol protocol, const QHostAddress& nameserver): QDnsLookup(typeVal, name, protocol, nameserver) {};
	MiqtVirtualQDnsLookup(QObject* parent): QDnsLookup(parent) {};
	MiqtVirtualQDnsLookup(Type typeVal, const QString& name, QObject* parent): QDnsLookup(typeVal, name, parent) {};
	MiqtVirtualQDnsLookup(Type typeVal, const QString& name, const QHostAddress& nameserver, QObject* parent): QDnsLookup(typeVal, name, nameserver, parent) {};
	MiqtVirtualQDnsLookup(Type typeVal, const QString& name, const QHostAddress& nameserver, quint16 port, QObject* parent): QDnsLookup(typeVal, name, nameserver, port, parent) {};
	MiqtVirtualQDnsLookup(Type typeVal, const QString& name, Protocol protocol, const QHostAddress& nameserver, quint16 port): QDnsLookup(typeVal, name, protocol, nameserver, port) {};
	MiqtVirtualQDnsLookup(Type typeVal, const QString& name, Protocol protocol, const QHostAddress& nameserver, quint16 port, QObject* parent): QDnsLookup(typeVal, name, protocol, nameserver, port, parent) {};

	virtual ~MiqtVirtualQDnsLookup() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Event = 0;

	// Subclass to allow providing a Go implementation
	virtual bool event(QEvent* event) override {
		if (handle__Event == 0) {
			return QDnsLookup::event(event);
		}
		
		QEvent* sigval1 = event;

		bool callback_return_value = miqt_exec_callback_QDnsLookup_Event(this, handle__Event, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_Event(QEvent* event) {

		return QDnsLookup::event(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__EventFilter = 0;

	// Subclass to allow providing a Go implementation
	virtual bool eventFilter(QObject* watched, QEvent* event) override {
		if (handle__EventFilter == 0) {
			return QDnsLookup::eventFilter(watched, event);
		}
		
		QObject* sigval1 = watched;
		QEvent* sigval2 = event;

		bool callback_return_value = miqt_exec_callback_QDnsLookup_EventFilter(this, handle__EventFilter, sigval1, sigval2);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_EventFilter(QObject* watched, QEvent* event) {

		return QDnsLookup::eventFilter(watched, event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__TimerEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void timerEvent(QTimerEvent* event) override {
		if (handle__TimerEvent == 0) {
			QDnsLookup::timerEvent(event);
			return;
		}
		
		QTimerEvent* sigval1 = event;

		miqt_exec_callback_QDnsLookup_TimerEvent(this, handle__TimerEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_TimerEvent(QTimerEvent* event) {

		QDnsLookup::timerEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ChildEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void childEvent(QChildEvent* event) override {
		if (handle__ChildEvent == 0) {
			QDnsLookup::childEvent(event);
			return;
		}
		
		QChildEvent* sigval1 = event;

		miqt_exec_callback_QDnsLookup_ChildEvent(this, handle__ChildEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ChildEvent(QChildEvent* event) {

		QDnsLookup::childEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__CustomEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void customEvent(QEvent* event) override {
		if (handle__CustomEvent == 0) {
			QDnsLookup::customEvent(event);
			return;
		}
		
		QEvent* sigval1 = event;

		miqt_exec_callback_QDnsLookup_CustomEvent(this, handle__CustomEvent, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_CustomEvent(QEvent* event) {

		QDnsLookup::customEvent(event);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ConnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void connectNotify(const QMetaMethod& signal) override {
		if (handle__ConnectNotify == 0) {
			QDnsLookup::connectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QDnsLookup_ConnectNotify(this, handle__ConnectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_ConnectNotify(QMetaMethod* signal) {

		QDnsLookup::connectNotify(*signal);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DisconnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void disconnectNotify(const QMetaMethod& signal) override {
		if (handle__DisconnectNotify == 0) {
			QDnsLookup::disconnectNotify(signal);
			return;
		}
		
		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);

		miqt_exec_callback_QDnsLookup_DisconnectNotify(this, handle__DisconnectNotify, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DisconnectNotify(QMetaMethod* signal) {

		QDnsLookup::disconnectNotify(*signal);

	}

};

QDnsLookup* QDnsLookup_new() {
	return new MiqtVirtualQDnsLookup();
}

QDnsLookup* QDnsLookup_new2(Type typeVal, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new MiqtVirtualQDnsLookup(typeVal, name_QString);
}

QDnsLookup* QDnsLookup_new3(Type typeVal, struct miqt_string name, QHostAddress* nameserver) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new MiqtVirtualQDnsLookup(typeVal, name_QString, *nameserver);
}

QDnsLookup* QDnsLookup_new4(Type typeVal, struct miqt_string name, QHostAddress* nameserver, uint16_t port) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new MiqtVirtualQDnsLookup(typeVal, name_QString, *nameserver, static_cast<quint16>(port));
}

QDnsLookup* QDnsLookup_new5(Type typeVal, struct miqt_string name, Protocol protocol, QHostAddress* nameserver) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new MiqtVirtualQDnsLookup(typeVal, name_QString, protocol, *nameserver);
}

QDnsLookup* QDnsLookup_new6(QObject* parent) {
	return new MiqtVirtualQDnsLookup(parent);
}

QDnsLookup* QDnsLookup_new7(Type typeVal, struct miqt_string name, QObject* parent) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new MiqtVirtualQDnsLookup(typeVal, name_QString, parent);
}

QDnsLookup* QDnsLookup_new8(Type typeVal, struct miqt_string name, QHostAddress* nameserver, QObject* parent) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new MiqtVirtualQDnsLookup(typeVal, name_QString, *nameserver, parent);
}

QDnsLookup* QDnsLookup_new9(Type typeVal, struct miqt_string name, QHostAddress* nameserver, uint16_t port, QObject* parent) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new MiqtVirtualQDnsLookup(typeVal, name_QString, *nameserver, static_cast<quint16>(port), parent);
}

QDnsLookup* QDnsLookup_new10(Type typeVal, struct miqt_string name, Protocol protocol, QHostAddress* nameserver, uint16_t port) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new MiqtVirtualQDnsLookup(typeVal, name_QString, protocol, *nameserver, static_cast<quint16>(port));
}

QDnsLookup* QDnsLookup_new11(Type typeVal, struct miqt_string name, Protocol protocol, QHostAddress* nameserver, uint16_t port, QObject* parent) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new MiqtVirtualQDnsLookup(typeVal, name_QString, protocol, *nameserver, static_cast<quint16>(port), parent);
}

void QDnsLookup_virtbase(QDnsLookup* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QDnsLookup_MetaObject(const QDnsLookup* self) {
	return (QMetaObject*) self->metaObject();
}

void* QDnsLookup_Metacast(QDnsLookup* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QDnsLookup_Tr(const char* s) {
	QString _ret = QDnsLookup::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QDnsLookup_IsAuthenticData(const QDnsLookup* self) {
	return self->isAuthenticData();
}

Error QDnsLookup_Error(const QDnsLookup* self) {
	return self->error();
}

struct miqt_string QDnsLookup_ErrorString(const QDnsLookup* self) {
	QString _ret = self->errorString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QDnsLookup_IsFinished(const QDnsLookup* self) {
	return self->isFinished();
}

struct miqt_string QDnsLookup_Name(const QDnsLookup* self) {
	QString _ret = self->name();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QDnsLookup_SetName(QDnsLookup* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	self->setName(name_QString);
}

Type QDnsLookup_Type(const QDnsLookup* self) {
	return self->type();
}

void QDnsLookup_SetType(QDnsLookup* self, int typeVal) {
	self->setType(static_cast<QDnsLookup::Type>(typeVal));
}

QHostAddress* QDnsLookup_Nameserver(const QDnsLookup* self) {
	return new QHostAddress(self->nameserver());
}

void QDnsLookup_SetNameserver(QDnsLookup* self, QHostAddress* nameserver) {
	self->setNameserver(*nameserver);
}

uint16_t QDnsLookup_NameserverPort(const QDnsLookup* self) {
	quint16 _ret = self->nameserverPort();
	return static_cast<uint16_t>(_ret);
}

void QDnsLookup_SetNameserverPort(QDnsLookup* self, uint16_t port) {
	self->setNameserverPort(static_cast<quint16>(port));
}

Protocol QDnsLookup_NameserverProtocol(const QDnsLookup* self) {
	return self->nameserverProtocol();
}

void QDnsLookup_SetNameserverProtocol(QDnsLookup* self, Protocol protocol) {
	self->setNameserverProtocol(protocol);
}

void QDnsLookup_SetNameserver2(QDnsLookup* self, Protocol protocol, QHostAddress* nameserver) {
	self->setNameserver(protocol, *nameserver);
}

void QDnsLookup_SetNameserver3(QDnsLookup* self, QHostAddress* nameserver, uint16_t port) {
	self->setNameserver(*nameserver, static_cast<quint16>(port));
}

struct miqt_array /* of QDnsDomainNameRecord* */  QDnsLookup_CanonicalNameRecords(const QDnsLookup* self) {
	QList<QDnsDomainNameRecord> _ret = self->canonicalNameRecords();
	// Convert QList<> from C++ memory to manually-managed C memory
	QDnsDomainNameRecord** _arr = static_cast<QDnsDomainNameRecord**>(malloc(sizeof(QDnsDomainNameRecord*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QDnsDomainNameRecord(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of QDnsHostAddressRecord* */  QDnsLookup_HostAddressRecords(const QDnsLookup* self) {
	QList<QDnsHostAddressRecord> _ret = self->hostAddressRecords();
	// Convert QList<> from C++ memory to manually-managed C memory
	QDnsHostAddressRecord** _arr = static_cast<QDnsHostAddressRecord**>(malloc(sizeof(QDnsHostAddressRecord*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QDnsHostAddressRecord(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of QDnsMailExchangeRecord* */  QDnsLookup_MailExchangeRecords(const QDnsLookup* self) {
	QList<QDnsMailExchangeRecord> _ret = self->mailExchangeRecords();
	// Convert QList<> from C++ memory to manually-managed C memory
	QDnsMailExchangeRecord** _arr = static_cast<QDnsMailExchangeRecord**>(malloc(sizeof(QDnsMailExchangeRecord*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QDnsMailExchangeRecord(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of QDnsDomainNameRecord* */  QDnsLookup_NameServerRecords(const QDnsLookup* self) {
	QList<QDnsDomainNameRecord> _ret = self->nameServerRecords();
	// Convert QList<> from C++ memory to manually-managed C memory
	QDnsDomainNameRecord** _arr = static_cast<QDnsDomainNameRecord**>(malloc(sizeof(QDnsDomainNameRecord*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QDnsDomainNameRecord(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of QDnsDomainNameRecord* */  QDnsLookup_PointerRecords(const QDnsLookup* self) {
	QList<QDnsDomainNameRecord> _ret = self->pointerRecords();
	// Convert QList<> from C++ memory to manually-managed C memory
	QDnsDomainNameRecord** _arr = static_cast<QDnsDomainNameRecord**>(malloc(sizeof(QDnsDomainNameRecord*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QDnsDomainNameRecord(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of QDnsServiceRecord* */  QDnsLookup_ServiceRecords(const QDnsLookup* self) {
	QList<QDnsServiceRecord> _ret = self->serviceRecords();
	// Convert QList<> from C++ memory to manually-managed C memory
	QDnsServiceRecord** _arr = static_cast<QDnsServiceRecord**>(malloc(sizeof(QDnsServiceRecord*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QDnsServiceRecord(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of QDnsTextRecord* */  QDnsLookup_TextRecords(const QDnsLookup* self) {
	QList<QDnsTextRecord> _ret = self->textRecords();
	// Convert QList<> from C++ memory to manually-managed C memory
	QDnsTextRecord** _arr = static_cast<QDnsTextRecord**>(malloc(sizeof(QDnsTextRecord*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QDnsTextRecord(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of QDnsTlsAssociationRecord* */  QDnsLookup_TlsAssociationRecords(const QDnsLookup* self) {
	QList<QDnsTlsAssociationRecord> _ret = self->tlsAssociationRecords();
	// Convert QList<> from C++ memory to manually-managed C memory
	QDnsTlsAssociationRecord** _arr = static_cast<QDnsTlsAssociationRecord**>(malloc(sizeof(QDnsTlsAssociationRecord*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QDnsTlsAssociationRecord(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

void QDnsLookup_SetSslConfiguration(QDnsLookup* self, QSslConfiguration* sslConfiguration) {
	self->setSslConfiguration(*sslConfiguration);
}

QSslConfiguration* QDnsLookup_SslConfiguration(const QDnsLookup* self) {
	return new QSslConfiguration(self->sslConfiguration());
}

bool QDnsLookup_IsProtocolSupported(Protocol protocol) {
	return QDnsLookup::isProtocolSupported(protocol);
}

uint16_t QDnsLookup_DefaultPortForProtocol(Protocol protocol) {
	quint16 _ret = QDnsLookup::defaultPortForProtocol(protocol);
	return static_cast<uint16_t>(_ret);
}

void QDnsLookup_Abort(QDnsLookup* self) {
	self->abort();
}

void QDnsLookup_Lookup(QDnsLookup* self) {
	self->lookup();
}

void QDnsLookup_Finished(QDnsLookup* self) {
	self->finished();
}

void QDnsLookup_connect_Finished(QDnsLookup* self, intptr_t slot) {
	MiqtVirtualQDnsLookup::connect(self, static_cast<void (QDnsLookup::*)()>(&QDnsLookup::finished), self, [=]() {
		miqt_exec_callback_QDnsLookup_Finished(slot);
	});
}

void QDnsLookup_NameChanged(QDnsLookup* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	self->nameChanged(name_QString);
}

void QDnsLookup_connect_NameChanged(QDnsLookup* self, intptr_t slot) {
	MiqtVirtualQDnsLookup::connect(self, static_cast<void (QDnsLookup::*)(const QString&)>(&QDnsLookup::nameChanged), self, [=](const QString& name) {
		const QString name_ret = name;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray name_b = name_ret.toUtf8();
		struct miqt_string name_ms;
		name_ms.len = name_b.length();
		name_ms.data = static_cast<char*>(malloc(name_ms.len));
		memcpy(name_ms.data, name_b.data(), name_ms.len);
		struct miqt_string sigval1 = name_ms;
		miqt_exec_callback_QDnsLookup_NameChanged(slot, sigval1);
	});
}

void QDnsLookup_TypeChanged(QDnsLookup* self, int typeVal) {
	self->typeChanged(static_cast<QDnsLookup::Type>(typeVal));
}

void QDnsLookup_connect_TypeChanged(QDnsLookup* self, intptr_t slot) {
	MiqtVirtualQDnsLookup::connect(self, static_cast<void (QDnsLookup::*)(QDnsLookup::Type)>(&QDnsLookup::typeChanged), self, [=](QDnsLookup::Type typeVal) {
		QDnsLookup::Type typeVal_ret = typeVal;
		int sigval1 = static_cast<int>(typeVal_ret);
		miqt_exec_callback_QDnsLookup_TypeChanged(slot, sigval1);
	});
}

void QDnsLookup_NameserverChanged(QDnsLookup* self, QHostAddress* nameserver) {
	self->nameserverChanged(*nameserver);
}

void QDnsLookup_connect_NameserverChanged(QDnsLookup* self, intptr_t slot) {
	MiqtVirtualQDnsLookup::connect(self, static_cast<void (QDnsLookup::*)(const QHostAddress&)>(&QDnsLookup::nameserverChanged), self, [=](const QHostAddress& nameserver) {
		const QHostAddress& nameserver_ret = nameserver;
		// Cast returned reference into pointer
		QHostAddress* sigval1 = const_cast<QHostAddress*>(&nameserver_ret);
		miqt_exec_callback_QDnsLookup_NameserverChanged(slot, sigval1);
	});
}

void QDnsLookup_NameserverPortChanged(QDnsLookup* self, uint16_t port) {
	self->nameserverPortChanged(static_cast<quint16>(port));
}

void QDnsLookup_connect_NameserverPortChanged(QDnsLookup* self, intptr_t slot) {
	MiqtVirtualQDnsLookup::connect(self, static_cast<void (QDnsLookup::*)(quint16)>(&QDnsLookup::nameserverPortChanged), self, [=](quint16 port) {
		quint16 port_ret = port;
		uint16_t sigval1 = static_cast<uint16_t>(port_ret);
		miqt_exec_callback_QDnsLookup_NameserverPortChanged(slot, sigval1);
	});
}

void QDnsLookup_NameserverProtocolChanged(QDnsLookup* self, uint8_t protocol) {
	self->nameserverProtocolChanged(static_cast<QDnsLookup::Protocol>(protocol));
}

void QDnsLookup_connect_NameserverProtocolChanged(QDnsLookup* self, intptr_t slot) {
	MiqtVirtualQDnsLookup::connect(self, static_cast<void (QDnsLookup::*)(QDnsLookup::Protocol)>(&QDnsLookup::nameserverProtocolChanged), self, [=](QDnsLookup::Protocol protocol) {
		QDnsLookup::Protocol protocol_ret = protocol;
		uint8_t sigval1 = static_cast<uint8_t>(protocol_ret);
		miqt_exec_callback_QDnsLookup_NameserverProtocolChanged(slot, sigval1);
	});
}

struct miqt_string QDnsLookup_Tr2(const char* s, const char* c) {
	QString _ret = QDnsLookup::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDnsLookup_Tr3(const char* s, const char* c, int n) {
	QString _ret = QDnsLookup::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QDnsLookup_SetNameserver32(QDnsLookup* self, Protocol protocol, QHostAddress* nameserver, uint16_t port) {
	self->setNameserver(protocol, *nameserver, static_cast<quint16>(port));
}

void QDnsLookup_override_virtual_Event(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDnsLookup*>( (QDnsLookup*)(self) )->handle__Event = slot;
}

bool QDnsLookup_virtualbase_Event(void* self, QEvent* event) {
	return ( (MiqtVirtualQDnsLookup*)(self) )->virtualbase_Event(event);
}

void QDnsLookup_override_virtual_EventFilter(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDnsLookup*>( (QDnsLookup*)(self) )->handle__EventFilter = slot;
}

bool QDnsLookup_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event) {
	return ( (MiqtVirtualQDnsLookup*)(self) )->virtualbase_EventFilter(watched, event);
}

void QDnsLookup_override_virtual_TimerEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDnsLookup*>( (QDnsLookup*)(self) )->handle__TimerEvent = slot;
}

void QDnsLookup_virtualbase_TimerEvent(void* self, QTimerEvent* event) {
	( (MiqtVirtualQDnsLookup*)(self) )->virtualbase_TimerEvent(event);
}

void QDnsLookup_override_virtual_ChildEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDnsLookup*>( (QDnsLookup*)(self) )->handle__ChildEvent = slot;
}

void QDnsLookup_virtualbase_ChildEvent(void* self, QChildEvent* event) {
	( (MiqtVirtualQDnsLookup*)(self) )->virtualbase_ChildEvent(event);
}

void QDnsLookup_override_virtual_CustomEvent(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDnsLookup*>( (QDnsLookup*)(self) )->handle__CustomEvent = slot;
}

void QDnsLookup_virtualbase_CustomEvent(void* self, QEvent* event) {
	( (MiqtVirtualQDnsLookup*)(self) )->virtualbase_CustomEvent(event);
}

void QDnsLookup_override_virtual_ConnectNotify(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDnsLookup*>( (QDnsLookup*)(self) )->handle__ConnectNotify = slot;
}

void QDnsLookup_virtualbase_ConnectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQDnsLookup*)(self) )->virtualbase_ConnectNotify(signal);
}

void QDnsLookup_override_virtual_DisconnectNotify(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDnsLookup*>( (QDnsLookup*)(self) )->handle__DisconnectNotify = slot;
}

void QDnsLookup_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal) {
	( (MiqtVirtualQDnsLookup*)(self) )->virtualbase_DisconnectNotify(signal);
}

void QDnsLookup_Delete(QDnsLookup* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQDnsLookup*>( self );
	} else {
		delete self;
	}
}
