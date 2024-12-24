#pragma once
#ifndef MIQT_QT_GEN_QPIXMAPCACHE_H
#define MIQT_QT_GEN_QPIXMAPCACHE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QPixmapCache__Key)
typedef QPixmapCache::Key QPixmapCache__Key;
typedef struct QPixmap QPixmap;
typedef struct QPixmapCache QPixmapCache;
typedef struct QPixmapCache__Key QPixmapCache__Key;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) int QPixmapCache_CacheLimit();
extern __declspec(dllexport) void QPixmapCache_SetCacheLimit(int cacheLimit);
extern __declspec(dllexport) bool QPixmapCache_Find(struct miqt_string key, QPixmap* pixmap);
extern __declspec(dllexport) bool QPixmapCache_Find2(const Key* key, QPixmap* pixmap);
extern __declspec(dllexport) bool QPixmapCache_Insert(struct miqt_string key, QPixmap* pixmap);
extern __declspec(dllexport) Key QPixmapCache_InsertWithPixmap(QPixmap* pixmap);
extern __declspec(dllexport) bool QPixmapCache_Replace(const Key* key, QPixmap* pixmap);
extern __declspec(dllexport) void QPixmapCache_Remove(struct miqt_string key);
extern __declspec(dllexport) void QPixmapCache_RemoveWithKey(const Key* key);
extern __declspec(dllexport) void QPixmapCache_Clear();
extern __declspec(dllexport) void QPixmapCache_Delete(QPixmapCache* self, bool isSubclass);

extern __declspec(dllexport) QPixmapCache__Key* QPixmapCache__Key_new();
extern __declspec(dllexport) QPixmapCache__Key* QPixmapCache__Key_new2(const Key* other);
extern __declspec(dllexport) bool QPixmapCache__Key_OperatorEqual(const QPixmapCache__Key* self, const Key* key);
extern __declspec(dllexport) bool QPixmapCache__Key_OperatorNotEqual(const QPixmapCache__Key* self, const Key* key);
extern __declspec(dllexport) void QPixmapCache__Key_OperatorAssign(QPixmapCache__Key* self, const Key* other);
extern __declspec(dllexport) void QPixmapCache__Key_Swap(QPixmapCache__Key* self, Key* other);
extern __declspec(dllexport) bool QPixmapCache__Key_IsValid(const QPixmapCache__Key* self);
extern __declspec(dllexport) void QPixmapCache__Key_Delete(QPixmapCache__Key* self, bool isSubclass);

} 
