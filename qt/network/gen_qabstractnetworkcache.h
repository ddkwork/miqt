#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QABSTRACTNETWORKCACHE_H
#define MIQT_QT_NETWORK_GEN_QABSTRACTNETWORKCACHE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractNetworkCache QAbstractNetworkCache;
typedef struct QDateTime QDateTime;
typedef struct QHttpHeaders QHttpHeaders;
typedef struct QIODevice QIODevice;
typedef struct QMetaObject QMetaObject;
typedef struct QNetworkCacheMetaData QNetworkCacheMetaData;
typedef struct QObject QObject;
typedef struct QUrl QUrl;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QNetworkCacheMetaData* QNetworkCacheMetaData_new();
extern __declspec(dllexport) 
QNetworkCacheMetaData* QNetworkCacheMetaData_new2(QNetworkCacheMetaData* other);
extern __declspec(dllexport) 
void QNetworkCacheMetaData_OperatorAssign(QNetworkCacheMetaData* self, QNetworkCacheMetaData* other);
extern __declspec(dllexport) 
void QNetworkCacheMetaData_Swap(QNetworkCacheMetaData* self, QNetworkCacheMetaData* other);
extern __declspec(dllexport) 
bool QNetworkCacheMetaData_OperatorEqual(const QNetworkCacheMetaData* self, QNetworkCacheMetaData* other);
extern __declspec(dllexport) 
bool QNetworkCacheMetaData_OperatorNotEqual(const QNetworkCacheMetaData* self, QNetworkCacheMetaData* other);
extern __declspec(dllexport) 
bool QNetworkCacheMetaData_IsValid(const QNetworkCacheMetaData* self);
extern __declspec(dllexport) 
QUrl* QNetworkCacheMetaData_Url(const QNetworkCacheMetaData* self);
extern __declspec(dllexport) 
void QNetworkCacheMetaData_SetUrl(QNetworkCacheMetaData* self, QUrl* url);
extern __declspec(dllexport) 
RawHeaderList QNetworkCacheMetaData_RawHeaders(const QNetworkCacheMetaData* self);
extern __declspec(dllexport) 
void QNetworkCacheMetaData_SetRawHeaders(QNetworkCacheMetaData* self, const RawHeaderList* headers);
extern __declspec(dllexport) 
QHttpHeaders* QNetworkCacheMetaData_Headers(const QNetworkCacheMetaData* self);
extern __declspec(dllexport) 
void QNetworkCacheMetaData_SetHeaders(QNetworkCacheMetaData* self, QHttpHeaders* headers);
extern __declspec(dllexport) 
QDateTime* QNetworkCacheMetaData_LastModified(const QNetworkCacheMetaData* self);
extern __declspec(dllexport) 
void QNetworkCacheMetaData_SetLastModified(QNetworkCacheMetaData* self, QDateTime* dateTime);
extern __declspec(dllexport) 
QDateTime* QNetworkCacheMetaData_ExpirationDate(const QNetworkCacheMetaData* self);
extern __declspec(dllexport) 
void QNetworkCacheMetaData_SetExpirationDate(QNetworkCacheMetaData* self, QDateTime* dateTime);
extern __declspec(dllexport) 
bool QNetworkCacheMetaData_SaveToDisk(const QNetworkCacheMetaData* self);
extern __declspec(dllexport) 
void QNetworkCacheMetaData_SetSaveToDisk(QNetworkCacheMetaData* self, bool allow);
extern __declspec(dllexport) 
AttributesMap QNetworkCacheMetaData_Attributes(const QNetworkCacheMetaData* self);
extern __declspec(dllexport) 
void QNetworkCacheMetaData_SetAttributes(QNetworkCacheMetaData* self, const AttributesMap* attributes);
extern __declspec(dllexport) 
void QNetworkCacheMetaData_Delete(QNetworkCacheMetaData* self, bool isSubclass);

extern __declspec(dllexport) 
void QAbstractNetworkCache_virtbase(QAbstractNetworkCache* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QAbstractNetworkCache_MetaObject(const QAbstractNetworkCache* self);
extern __declspec(dllexport) 
void* QAbstractNetworkCache_Metacast(QAbstractNetworkCache* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAbstractNetworkCache_Tr(const char* s);
extern __declspec(dllexport) 
QNetworkCacheMetaData* QAbstractNetworkCache_MetaData(QAbstractNetworkCache* self, QUrl* url);
extern __declspec(dllexport) 
void QAbstractNetworkCache_UpdateMetaData(QAbstractNetworkCache* self, QNetworkCacheMetaData* metaData);
extern __declspec(dllexport) 
QIODevice* QAbstractNetworkCache_Data(QAbstractNetworkCache* self, QUrl* url);
extern __declspec(dllexport) 
bool QAbstractNetworkCache_Remove(QAbstractNetworkCache* self, QUrl* url);
extern __declspec(dllexport) 
long long QAbstractNetworkCache_CacheSize(const QAbstractNetworkCache* self);
extern __declspec(dllexport) 
QIODevice* QAbstractNetworkCache_Prepare(QAbstractNetworkCache* self, QNetworkCacheMetaData* metaData);
extern __declspec(dllexport) 
void QAbstractNetworkCache_Insert(QAbstractNetworkCache* self, QIODevice* device);
extern __declspec(dllexport) 
void QAbstractNetworkCache_Clear(QAbstractNetworkCache* self);
extern __declspec(dllexport) 
struct miqt_string QAbstractNetworkCache_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAbstractNetworkCache_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAbstractNetworkCache_Delete(QAbstractNetworkCache* self, bool isSubclass);

}
