#pragma once
#ifndef MIQT_QT_GEN_QABSTRACTFILEICONPROVIDER_H
#define MIQT_QT_GEN_QABSTRACTFILEICONPROVIDER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractFileIconProvider;
class QFileInfo;
class QIcon;
class _GUID;
class type_info;
#else
typedef struct QAbstractFileIconProvider QAbstractFileIconProvider;
typedef struct QFileInfo QFileInfo;
typedef struct QIcon QIcon;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QAbstractFileIconProvider* QAbstractFileIconProvider_new();
extern __declspec(dllexport) QIcon* QAbstractFileIconProvider_Icon(const QAbstractFileIconProvider* self, IconType param1);
extern __declspec(dllexport) QIcon* QAbstractFileIconProvider_IconWithQFileInfo(const QAbstractFileIconProvider* self, QFileInfo* param1);
extern __declspec(dllexport) struct miqt_string QAbstractFileIconProvider_Type(const QAbstractFileIconProvider* self, QFileInfo* param1);
extern __declspec(dllexport) void QAbstractFileIconProvider_SetOptions(QAbstractFileIconProvider* self, Options options);
extern __declspec(dllexport) Options QAbstractFileIconProvider_Options(const QAbstractFileIconProvider* self);
extern __declspec(dllexport) void QAbstractFileIconProvider_override_virtual_Icon(void* self, intptr_t slot);
QIcon* QAbstractFileIconProvider_virtualbase_Icon(const void* self, IconType param1);
extern __declspec(dllexport) void QAbstractFileIconProvider_override_virtual_IconWithQFileInfo(void* self, intptr_t slot);
QIcon* QAbstractFileIconProvider_virtualbase_IconWithQFileInfo(const void* self, QFileInfo* param1);
extern __declspec(dllexport) void QAbstractFileIconProvider_override_virtual_Type(void* self, intptr_t slot);
struct miqt_string QAbstractFileIconProvider_virtualbase_Type(const void* self, QFileInfo* param1);
extern __declspec(dllexport) void QAbstractFileIconProvider_override_virtual_SetOptions(void* self, intptr_t slot);
void QAbstractFileIconProvider_virtualbase_SetOptions(void* self, Options options);
extern __declspec(dllexport) void QAbstractFileIconProvider_override_virtual_Options(void* self, intptr_t slot);
Options QAbstractFileIconProvider_virtualbase_Options(const void* self);
extern __declspec(dllexport) void QAbstractFileIconProvider_Delete(QAbstractFileIconProvider* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
