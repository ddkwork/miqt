// +build ignore

#include <QHostAddress>
#include <QIPv6Address>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qhostaddress.h>
#include "gen_qhostaddress.h"

#ifndef _Bool
#define _Bool bool
#endif

QIPv6Address* QIPv6Address_new() {
	return new QIPv6Address();
}

QIPv6Address* QIPv6Address_new2(QIPv6Address* param1) {
	return new QIPv6Address(*param1);
}

unsigned char QIPv6Address_OperatorSubscript(const QIPv6Address* self, int index) {
	quint8 _ret = self->operator[](static_cast<int>(index));
	return static_cast<unsigned char>(_ret);
}

void QIPv6Address_Delete(QIPv6Address* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QIPv6Address*>( self );
	} else {
		delete self;
	}
}

QHostAddress* QHostAddress_new() {
	return new QHostAddress();
}

QHostAddress* QHostAddress_new2(unsigned int ip4Addr) {
	return new QHostAddress(static_cast<quint32>(ip4Addr));
}

QHostAddress* QHostAddress_new3(const unsigned char* ip6Addr) {
	return new QHostAddress(static_cast<const quint8*>(ip6Addr));
}

QHostAddress* QHostAddress_new4(QIPv6Address* ip6Addr) {
	return new QHostAddress(*ip6Addr);
}

QHostAddress* QHostAddress_new5(struct miqt_string address) {
	QString address_QString = QString::fromUtf8(address.data, address.len);
	return new QHostAddress(address_QString);
}

QHostAddress* QHostAddress_new6(QHostAddress* copyVal) {
	return new QHostAddress(*copyVal);
}

QHostAddress* QHostAddress_new7(SpecialAddress address) {
	return new QHostAddress(address);
}

void QHostAddress_OperatorAssign(QHostAddress* self, QHostAddress* other) {
	self->operator=(*other);
}

void QHostAddress_OperatorAssignWithAddress(QHostAddress* self, SpecialAddress address) {
	self->operator=(address);
}

void QHostAddress_Swap(QHostAddress* self, QHostAddress* other) {
	self->swap(*other);
}

void QHostAddress_SetAddress(QHostAddress* self, unsigned int ip4Addr) {
	self->setAddress(static_cast<quint32>(ip4Addr));
}

void QHostAddress_SetAddressWithIp6Addr(QHostAddress* self, const unsigned char* ip6Addr) {
	self->setAddress(static_cast<const quint8*>(ip6Addr));
}

void QHostAddress_SetAddress2(QHostAddress* self, QIPv6Address* ip6Addr) {
	self->setAddress(*ip6Addr);
}

bool QHostAddress_SetAddress3(QHostAddress* self, struct miqt_string address) {
	QString address_QString = QString::fromUtf8(address.data, address.len);
	return self->setAddress(address_QString);
}

void QHostAddress_SetAddress4(QHostAddress* self, SpecialAddress address) {
	self->setAddress(address);
}

NetworkLayerProtocol QHostAddress_Protocol(const QHostAddress* self) {
	return self->protocol();
}

unsigned int QHostAddress_ToIPv4Address(const QHostAddress* self) {
	quint32 _ret = self->toIPv4Address();
	return static_cast<unsigned int>(_ret);
}

QIPv6Address* QHostAddress_ToIPv6Address(const QHostAddress* self) {
	return new QIPv6Address(self->toIPv6Address());
}

struct miqt_string QHostAddress_ToString(const QHostAddress* self) {
	QString _ret = self->toString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QHostAddress_ScopeId(const QHostAddress* self) {
	QString _ret = self->scopeId();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QHostAddress_SetScopeId(QHostAddress* self, struct miqt_string id) {
	QString id_QString = QString::fromUtf8(id.data, id.len);
	self->setScopeId(id_QString);
}

bool QHostAddress_IsEqual(const QHostAddress* self, QHostAddress* address) {
	return self->isEqual(*address);
}

bool QHostAddress_OperatorEqual(const QHostAddress* self, QHostAddress* address) {
	return (*self == *address);
}

bool QHostAddress_OperatorEqualWithAddress(const QHostAddress* self, SpecialAddress address) {
	return (*self == address);
}

bool QHostAddress_OperatorNotEqual(const QHostAddress* self, QHostAddress* address) {
	return (*self != *address);
}

bool QHostAddress_OperatorNotEqualWithAddress(const QHostAddress* self, SpecialAddress address) {
	return (*self != address);
}

bool QHostAddress_IsNull(const QHostAddress* self) {
	return self->isNull();
}

void QHostAddress_Clear(QHostAddress* self) {
	self->clear();
}

bool QHostAddress_IsInSubnet(const QHostAddress* self, QHostAddress* subnet, int netmask) {
	return self->isInSubnet(*subnet, static_cast<int>(netmask));
}

bool QHostAddress_IsLoopback(const QHostAddress* self) {
	return self->isLoopback();
}

bool QHostAddress_IsGlobal(const QHostAddress* self) {
	return self->isGlobal();
}

bool QHostAddress_IsLinkLocal(const QHostAddress* self) {
	return self->isLinkLocal();
}

bool QHostAddress_IsSiteLocal(const QHostAddress* self) {
	return self->isSiteLocal();
}

bool QHostAddress_IsUniqueLocalUnicast(const QHostAddress* self) {
	return self->isUniqueLocalUnicast();
}

bool QHostAddress_IsMulticast(const QHostAddress* self) {
	return self->isMulticast();
}

bool QHostAddress_IsBroadcast(const QHostAddress* self) {
	return self->isBroadcast();
}

bool QHostAddress_IsPrivateUse(const QHostAddress* self) {
	return self->isPrivateUse();
}

unsigned int QHostAddress_ToIPv4Address1(const QHostAddress* self, bool* ok) {
	quint32 _ret = self->toIPv4Address(ok);
	return static_cast<unsigned int>(_ret);
}

bool QHostAddress_IsEqual2(const QHostAddress* self, QHostAddress* address, ConversionMode mode) {
	return self->isEqual(*address, mode);
}

void QHostAddress_Delete(QHostAddress* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QHostAddress*>( self );
	} else {
		delete self;
	}
}
