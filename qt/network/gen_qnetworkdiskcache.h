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

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QNetworkDiskCache* QNetworkDiskCache_new();
extern __declspec(dllexport) QNetworkDiskCache* QNetworkDiskCache_new2(QObject* parent);
extern __declspec(dllexport) void QNetworkDiskCache_virtbase(QNetworkDiskCache* src, QAbstractNetworkCache** outptr_QAbstractNetworkCache);
extern __declspec(dllexport) QMetaObject* QNetworkDiskCache_MetaObject(const QNetworkDiskCache* self);
extern __declspec(dllexport) void* QNetworkDiskCache_Metacast(QNetworkDiskCache* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QNetworkDiskCache_Tr(const char* s);
extern __declspec(dllexport) struct miqt_string QNetworkDiskCache_CacheDirectory(const QNetworkDiskCache* self);
extern __declspec(dllexport) void QNetworkDiskCache_SetCacheDirectory(QNetworkDiskCache* self, struct miqt_string cacheDir);
extern __declspec(dllexport) long long QNetworkDiskCache_MaximumCacheSize(const QNetworkDiskCache* self);
extern __declspec(dllexport) void QNetworkDiskCache_SetMaximumCacheSize(QNetworkDiskCache* self, long long size);
extern __declspec(dllexport) long long QNetworkDiskCache_CacheSize(const QNetworkDiskCache* self);
extern __declspec(dllexport) QNetworkCacheMetaData* QNetworkDiskCache_MetaData(QNetworkDiskCache* self, QUrl* url);
extern __declspec(dllexport) void QNetworkDiskCache_UpdateMetaData(QNetworkDiskCache* self, QNetworkCacheMetaData* metaData);
extern __declspec(dllexport) QIODevice* QNetworkDiskCache_Data(QNetworkDiskCache* self, QUrl* url);
extern __declspec(dllexport) bool QNetworkDiskCache_Remove(QNetworkDiskCache* self, QUrl* url);
extern __declspec(dllexport) QIODevice* QNetworkDiskCache_Prepare(QNetworkDiskCache* self, QNetworkCacheMetaData* metaData);
extern __declspec(dllexport) void QNetworkDiskCache_Insert(QNetworkDiskCache* self, QIODevice* device);
extern __declspec(dllexport) QNetworkCacheMetaData* QNetworkDiskCache_FileMetaData(const QNetworkDiskCache* self, struct miqt_string fileName);
extern __declspec(dllexport) void QNetworkDiskCache_Clear(QNetworkDiskCache* self);
extern __declspec(dllexport) long long QNetworkDiskCache_Expire(QNetworkDiskCache* self);
extern __declspec(dllexport) struct miqt_string QNetworkDiskCache_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QNetworkDiskCache_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QNetworkDiskCache_override_virtual_CacheSize(void* self, intptr_t slot);
long long QNetworkDiskCache_virtualbase_CacheSize(const void* self);
extern __declspec(dllexport) void QNetworkDiskCache_override_virtual_MetaData(void* self, intptr_t slot);
QNetworkCacheMetaData* QNetworkDiskCache_virtualbase_MetaData(void* self, QUrl* url);
extern __declspec(dllexport) void QNetworkDiskCache_override_virtual_UpdateMetaData(void* self, intptr_t slot);
void QNetworkDiskCache_virtualbase_UpdateMetaData(void* self, QNetworkCacheMetaData* metaData);
extern __declspec(dllexport) void QNetworkDiskCache_override_virtual_Data(void* self, intptr_t slot);
QIODevice* QNetworkDiskCache_virtualbase_Data(void* self, QUrl* url);
extern __declspec(dllexport) void QNetworkDiskCache_override_virtual_Remove(void* self, intptr_t slot);
bool QNetworkDiskCache_virtualbase_Remove(void* self, QUrl* url);
extern __declspec(dllexport) void QNetworkDiskCache_override_virtual_Prepare(void* self, intptr_t slot);
QIODevice* QNetworkDiskCache_virtualbase_Prepare(void* self, QNetworkCacheMetaData* metaData);
extern __declspec(dllexport) void QNetworkDiskCache_override_virtual_Insert(void* self, intptr_t slot);
void QNetworkDiskCache_virtualbase_Insert(void* self, QIODevice* device);
extern __declspec(dllexport) void QNetworkDiskCache_override_virtual_Clear(void* self, intptr_t slot);
void QNetworkDiskCache_virtualbase_Clear(void* self);
extern __declspec(dllexport) void QNetworkDiskCache_override_virtual_Expire(void* self, intptr_t slot);
long long QNetworkDiskCache_virtualbase_Expire(void* self);
extern __declspec(dllexport) void QNetworkDiskCache_Delete(QNetworkDiskCache* self, bool isSubclass);

} 
