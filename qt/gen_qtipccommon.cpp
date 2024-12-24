// +build ignore

#include <QNativeIpcKey>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qtipccommon.h>
#include "gen_qtipccommon.h"

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

QNativeIpcKey* QNativeIpcKey_new() {
	return new QNativeIpcKey();
}

QNativeIpcKey* QNativeIpcKey_new2(Type typeVal) {
	return new QNativeIpcKey(typeVal);
}

QNativeIpcKey* QNativeIpcKey_new3(struct miqt_string k) {
	QString k_QString = QString::fromUtf8(k.data, k.len);
	return new QNativeIpcKey(k_QString);
}

QNativeIpcKey* QNativeIpcKey_new4(QNativeIpcKey* other) {
	return new QNativeIpcKey(*other);
}

QNativeIpcKey* QNativeIpcKey_new5(struct miqt_string k, Type typeVal) {
	QString k_QString = QString::fromUtf8(k.data, k.len);
	return new QNativeIpcKey(k_QString, typeVal);
}

Type QNativeIpcKey_LegacyDefaultTypeForOs() {
	return QNativeIpcKey::legacyDefaultTypeForOs();
}

void QNativeIpcKey_OperatorAssign(QNativeIpcKey* self, QNativeIpcKey* other) {
	self->operator=(*other);
}

void QNativeIpcKey_Swap(QNativeIpcKey* self, QNativeIpcKey* other) {
	self->swap(*other);
}

bool QNativeIpcKey_IsEmpty(const QNativeIpcKey* self) {
	return self->isEmpty();
}

bool QNativeIpcKey_IsValid(const QNativeIpcKey* self) {
	return self->isValid();
}

Type QNativeIpcKey_Type(const QNativeIpcKey* self) {
	return self->type();
}

void QNativeIpcKey_SetType(QNativeIpcKey* self, Type typeVal) {
	self->setType(typeVal);
}

struct miqt_string QNativeIpcKey_NativeKey(const QNativeIpcKey* self) {
	QString _ret = self->nativeKey();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QNativeIpcKey_SetNativeKey(QNativeIpcKey* self, struct miqt_string newKey) {
	QString newKey_QString = QString::fromUtf8(newKey.data, newKey.len);
	self->setNativeKey(newKey_QString);
}

struct miqt_string QNativeIpcKey_ToString(const QNativeIpcKey* self) {
	QString _ret = self->toString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QNativeIpcKey* QNativeIpcKey_FromString(struct miqt_string stringVal) {
	QString stringVal_QString = QString::fromUtf8(stringVal.data, stringVal.len);
	return new QNativeIpcKey(QNativeIpcKey::fromString(stringVal_QString));
}

void QNativeIpcKey_Delete(QNativeIpcKey* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QNativeIpcKey*>( self );
	} else {
		delete self;
	}
}

