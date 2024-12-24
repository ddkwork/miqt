// +build ignore

#include <QAnyStringView>
#include <QByteArray>
#include <QByteArrayView>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QUuid>
#define WORKAROUND_INNER_CLASS_DEFINITION_QUuid__Id128Bytes
#define WORKAROUND_INNER_CLASS_DEFINITION_Disambiguated_t
#include <quuid.h>
#include "gen_quuid.h"

QUuid* QUuid_new() {
	return new QUuid();
}

QUuid* QUuid_new2(unsigned int l, uint16_t w1, uint16_t w2, unsigned char b1, unsigned char b2, unsigned char b3, unsigned char b4, unsigned char b5, unsigned char b6, unsigned char b7, unsigned char b8) {
	return new QUuid(static_cast<uint>(l), static_cast<ushort>(w1), static_cast<ushort>(w2), static_cast<uchar>(b1), static_cast<uchar>(b2), static_cast<uchar>(b3), static_cast<uchar>(b4), static_cast<uchar>(b5), static_cast<uchar>(b6), static_cast<uchar>(b7), static_cast<uchar>(b8));
}

QUuid* QUuid_new3(Id128Bytes id128) {
	return new QUuid(id128);
}

QUuid* QUuid_new4(QAnyStringView* stringVal) {
	return new QUuid(*stringVal);
}

QUuid* QUuid_new5(const struct _GUID* guid) {
	return new QUuid(*guid);
}

QUuid* QUuid_new6(QUuid* param1) {
	return new QUuid(*param1);
}

QUuid* QUuid_new7(Id128Bytes id128, int order) {
	return new QUuid(id128, static_cast<QSysInfo::Endian>(order));
}

QUuid* QUuid_FromString(QAnyStringView* stringVal) {
	return new QUuid(QUuid::fromString(*stringVal));
}

struct miqt_string QUuid_ToString(const QUuid* self) {
	QString _ret = self->toString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QUuid_ToByteArray(const QUuid* self) {
	QByteArray _qb = self->toByteArray();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

Id128Bytes QUuid_ToBytes(const QUuid* self) {
	return self->toBytes();
}

struct miqt_string QUuid_ToRfc4122(const QUuid* self) {
	QByteArray _qb = self->toRfc4122();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

QUuid* QUuid_FromBytes(const void* bytes) {
	return new QUuid(QUuid::fromBytes(bytes));
}

QUuid* QUuid_FromRfc4122(QByteArrayView* param1) {
	return new QUuid(QUuid::fromRfc4122(*param1));
}

bool QUuid_IsNull(const QUuid* self) {
	return self->isNull();
}

void QUuid_OperatorAssign(QUuid* self, const struct _GUID* guid) {
	self->operator=(*guid);
}

QUuid* QUuid_CreateUuid() {
	return new QUuid(QUuid::createUuid());
}

QUuid* QUuid_CreateUuidV5(QUuid* ns, QByteArrayView* baseData) {
	return new QUuid(QUuid::createUuidV5(*ns, *baseData));
}

QUuid* QUuid_CreateUuidV3(QUuid* ns, QByteArrayView* baseData) {
	return new QUuid(QUuid::createUuidV3(*ns, *baseData));
}

QUuid* QUuid_CreateUuidV7() {
	return new QUuid(QUuid::createUuidV7());
}

Variant QUuid_Variant(const QUuid* self) {
	return self->variant();
}

Version QUuid_Version(const QUuid* self) {
	return self->version();
}

void QUuid_OperatorAssignWithQUuid(QUuid* self, QUuid* param1) {
	self->operator=(*param1);
}

struct miqt_string QUuid_ToString1(const QUuid* self, StringFormat mode) {
	QString _ret = self->toString(mode);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QUuid_ToByteArray1(const QUuid* self, StringFormat mode) {
	QByteArray _qb = self->toByteArray(mode);
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

Id128Bytes QUuid_ToBytes1(const QUuid* self, int order) {
	return self->toBytes(static_cast<QSysInfo::Endian>(order));
}

QUuid* QUuid_FromBytes2(const void* bytes, int order) {
	return new QUuid(QUuid::fromBytes(bytes, static_cast<QSysInfo::Endian>(order)));
}

bool QUuid_IsNull1(const QUuid* self, Disambiguated_t* param1) {
	return self->isNull(*param1);
}

Variant QUuid_Variant1(const QUuid* self, Disambiguated_t* param1) {
	return self->variant(*param1);
}

Version QUuid_Version1(const QUuid* self, Disambiguated_t* param1) {
	return self->version(*param1);
}

void QUuid_Delete(QUuid* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QUuid*>( self );
	} else {
		delete self;
	}
}

QUuid__Id128Bytes* QUuid__Id128Bytes_new() {
	return new QUuid::Id128Bytes();
}

QUuid__Id128Bytes* QUuid__Id128Bytes_new2(const Id128Bytes* param1) {
	return new QUuid::Id128Bytes(*param1);
}

void QUuid__Id128Bytes_Delete(QUuid__Id128Bytes* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QUuid::Id128Bytes*>( self );
	} else {
		delete self;
	}
}

