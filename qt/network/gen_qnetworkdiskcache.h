#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QNETWORKDISKCACHE_H
#define MIQT_QT_NETWORK_GEN_QNETWORKDISKCACHE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractNetworkCache QAbstractNetworkCache;
typedef struct QIODevice QIODevice;
typedef struct QMetaObject QMetaObject;
typedef struct QNetworkCacheMetaData QNetworkCacheMetaData;
typedef struct QNetworkDiskCache QNetworkDiskCache;
typedef struct QObject QObject;
typedef struct QUrl QUrl;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QNetworkDiskCache* QNetworkDiskCache_new();
extern __declspec(dllexport) 
QNetworkDiskCache* QNetworkDiskCache_new2(QObject* parent);
extern __declspec(dllexport) 
void QNetworkDiskCache_virtbase(QNetworkDiskCache* src
, QAbstractNetworkCache** outptr_QAbstractNetworkCache
);
extern __declspec(dllexport) 
QMetaObject* QNetworkDiskCache_MetaObject(const QNetworkDiskCache* self);
extern __declspec(dllexport) 
void* QNetworkDiskCache_Metacast(QNetworkDiskCache* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QNetworkDiskCache_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_string QNetworkDiskCache_CacheDirectory(const QNetworkDiskCache* self);
extern __declspec(dllexport) 
void QNetworkDiskCache_SetCacheDirectory(QNetworkDiskCache* self, struct miqt_string cacheDir);
extern __declspec(dllexport) 
long long QNetworkDiskCache_MaximumCacheSize(const QNetworkDiskCache* self);
extern __declspec(dllexport) 
void QNetworkDiskCache_SetMaximumCacheSize(QNetworkDiskCache* self, long long size);
extern __declspec(dllexport) 
long long QNetworkDiskCache_CacheSize(const QNetworkDiskCache* self);
extern __declspec(dllexport) 
QNetworkCacheMetaData* QNetworkDiskCache_MetaData(QNetworkDiskCache* self, QUrl* url);
extern __declspec(dllexport) 
void QNetworkDiskCache_UpdateMetaData(QNetworkDiskCache* self, QNetworkCacheMetaData* metaData);
extern __declspec(dllexport) 
QIODevice* QNetworkDiskCache_Data(QNetworkDiskCache* self, QUrl* url);
extern __declspec(dllexport) 
bool QNetworkDiskCache_Remove(QNetworkDiskCache* self, QUrl* url);
extern __declspec(dllexport) 
QIODevice* QNetworkDiskCache_Prepare(QNetworkDiskCache* self, QNetworkCacheMetaData* metaData);
extern __declspec(dllexport) 
void QNetworkDiskCache_Insert(QNetworkDiskCache* self, QIODevice* device);
extern __declspec(dllexport) 
QNetworkCacheMetaData* QNetworkDiskCache_FileMetaData(const QNetworkDiskCache* self, struct miqt_string fileName);
extern __declspec(dllexport) 
void QNetworkDiskCache_Clear(QNetworkDiskCache* self);
extern __declspec(dllexport) 
long long QNetworkDiskCache_Expire(QNetworkDiskCache* self);
extern __declspec(dllexport) 
struct miqt_string QNetworkDiskCache_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QNetworkDiskCache_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QNetworkDiskCache_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QNetworkDiskCache_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QNetworkDiskCache_override_virtual_Metacast(void* self, intptr_t slot);
void* QNetworkDiskCache_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QNetworkDiskCache_Delete(QNetworkDiskCache* self, bool isSubclass);

}
