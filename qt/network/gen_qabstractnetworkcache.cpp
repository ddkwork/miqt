// +build ignore

#include <QAbstractNetworkCache>
#include <QDateTime>
#include <QHttpHeaders>
#include <QIODevice>
#include <QMetaObject>
#include <QNetworkCacheMetaData>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QUrl>
#include <qabstractnetworkcache.h>
#include "gen_qabstractnetworkcache.h"

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

QNetworkCacheMetaData* QNetworkCacheMetaData_new() {
	return new QNetworkCacheMetaData();
}

QNetworkCacheMetaData* QNetworkCacheMetaData_new2(QNetworkCacheMetaData* other) {
	return new QNetworkCacheMetaData(*other);
}

void QNetworkCacheMetaData_OperatorAssign(QNetworkCacheMetaData* self, QNetworkCacheMetaData* other) {
	self->operator=(*other);
}

void QNetworkCacheMetaData_Swap(QNetworkCacheMetaData* self, QNetworkCacheMetaData* other) {
	self->swap(*other);
}

bool QNetworkCacheMetaData_OperatorEqual(const QNetworkCacheMetaData* self, QNetworkCacheMetaData* other) {
	return (*self == *other);
}

bool QNetworkCacheMetaData_OperatorNotEqual(const QNetworkCacheMetaData* self, QNetworkCacheMetaData* other) {
	return (*self != *other);
}

bool QNetworkCacheMetaData_IsValid(const QNetworkCacheMetaData* self) {
	return self->isValid();
}

QUrl* QNetworkCacheMetaData_Url(const QNetworkCacheMetaData* self) {
	return new QUrl(self->url());
}

void QNetworkCacheMetaData_SetUrl(QNetworkCacheMetaData* self, QUrl* url) {
	self->setUrl(*url);
}

RawHeaderList QNetworkCacheMetaData_RawHeaders(const QNetworkCacheMetaData* self) {
	return self->rawHeaders();
}

void QNetworkCacheMetaData_SetRawHeaders(QNetworkCacheMetaData* self, const RawHeaderList* headers) {
	self->setRawHeaders(*headers);
}

QHttpHeaders* QNetworkCacheMetaData_Headers(const QNetworkCacheMetaData* self) {
	return new QHttpHeaders(self->headers());
}

void QNetworkCacheMetaData_SetHeaders(QNetworkCacheMetaData* self, QHttpHeaders* headers) {
	self->setHeaders(*headers);
}

QDateTime* QNetworkCacheMetaData_LastModified(const QNetworkCacheMetaData* self) {
	return new QDateTime(self->lastModified());
}

void QNetworkCacheMetaData_SetLastModified(QNetworkCacheMetaData* self, QDateTime* dateTime) {
	self->setLastModified(*dateTime);
}

QDateTime* QNetworkCacheMetaData_ExpirationDate(const QNetworkCacheMetaData* self) {
	return new QDateTime(self->expirationDate());
}

void QNetworkCacheMetaData_SetExpirationDate(QNetworkCacheMetaData* self, QDateTime* dateTime) {
	self->setExpirationDate(*dateTime);
}

bool QNetworkCacheMetaData_SaveToDisk(const QNetworkCacheMetaData* self) {
	return self->saveToDisk();
}

void QNetworkCacheMetaData_SetSaveToDisk(QNetworkCacheMetaData* self, bool allow) {
	self->setSaveToDisk(allow);
}

AttributesMap QNetworkCacheMetaData_Attributes(const QNetworkCacheMetaData* self) {
	return self->attributes();
}

void QNetworkCacheMetaData_SetAttributes(QNetworkCacheMetaData* self, const AttributesMap* attributes) {
	self->setAttributes(*attributes);
}

void QNetworkCacheMetaData_Delete(QNetworkCacheMetaData* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QNetworkCacheMetaData*>( self );
	} else {
		delete self;
	}
}

void QAbstractNetworkCache_virtbase(QAbstractNetworkCache* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QAbstractNetworkCache_MetaObject(const QAbstractNetworkCache* self) {
	return (QMetaObject*) self->metaObject();
}

void* QAbstractNetworkCache_Metacast(QAbstractNetworkCache* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QAbstractNetworkCache_Tr(const char* s) {
	QString _ret = QAbstractNetworkCache::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QNetworkCacheMetaData* QAbstractNetworkCache_MetaData(QAbstractNetworkCache* self, QUrl* url) {
	return new QNetworkCacheMetaData(self->metaData(*url));
}

void QAbstractNetworkCache_UpdateMetaData(QAbstractNetworkCache* self, QNetworkCacheMetaData* metaData) {
	self->updateMetaData(*metaData);
}

QIODevice* QAbstractNetworkCache_Data(QAbstractNetworkCache* self, QUrl* url) {
	return self->data(*url);
}

bool QAbstractNetworkCache_Remove(QAbstractNetworkCache* self, QUrl* url) {
	return self->remove(*url);
}

long long QAbstractNetworkCache_CacheSize(const QAbstractNetworkCache* self) {
	qint64 _ret = self->cacheSize();
	return static_cast<long long>(_ret);
}

QIODevice* QAbstractNetworkCache_Prepare(QAbstractNetworkCache* self, QNetworkCacheMetaData* metaData) {
	return self->prepare(*metaData);
}

void QAbstractNetworkCache_Insert(QAbstractNetworkCache* self, QIODevice* device) {
	self->insert(device);
}

void QAbstractNetworkCache_Clear(QAbstractNetworkCache* self) {
	self->clear();
}

struct miqt_string QAbstractNetworkCache_Tr2(const char* s, const char* c) {
	QString _ret = QAbstractNetworkCache::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAbstractNetworkCache_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAbstractNetworkCache::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAbstractNetworkCache_Delete(QAbstractNetworkCache* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QAbstractNetworkCache*>( self );
	} else {
		delete self;
	}
}

