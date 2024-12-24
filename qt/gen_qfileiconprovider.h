#pragma once
#ifndef MIQT_QT_GEN_QFILEICONPROVIDER_H
#define MIQT_QT_GEN_QFILEICONPROVIDER_H

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
class QFileIconProvider;
class QFileInfo;
class QIcon;
class _GUID;
class type_info;
#else
typedef struct QAbstractFileIconProvider QAbstractFileIconProvider;
typedef struct QFileIconProvider QFileIconProvider;
typedef struct QFileInfo QFileInfo;
typedef struct QIcon QIcon;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QFileIconProvider* QFileIconProvider_new();
extern __declspec(dllexport) void QFileIconProvider_virtbase(QFileIconProvider* src, QAbstractFileIconProvider** outptr_QAbstractFileIconProvider);
extern __declspec(dllexport) QIcon* QFileIconProvider_Icon(const QFileIconProvider* self, IconType typeVal);
extern __declspec(dllexport) QIcon* QFileIconProvider_IconWithInfo(const QFileIconProvider* self, QFileInfo* info);
extern __declspec(dllexport) void QFileIconProvider_override_virtual_Icon(void* self, intptr_t slot);
QIcon* QFileIconProvider_virtualbase_Icon(const void* self, IconType typeVal);
extern __declspec(dllexport) void QFileIconProvider_override_virtual_IconWithInfo(void* self, intptr_t slot);
QIcon* QFileIconProvider_virtualbase_IconWithInfo(const void* self, QFileInfo* info);
extern __declspec(dllexport) void QFileIconProvider_override_virtual_Type(void* self, intptr_t slot);
struct miqt_string QFileIconProvider_virtualbase_Type(const void* self, QFileInfo* param1);
extern __declspec(dllexport) void QFileIconProvider_override_virtual_SetOptions(void* self, intptr_t slot);
void QFileIconProvider_virtualbase_SetOptions(void* self, Options options);
extern __declspec(dllexport) void QFileIconProvider_override_virtual_Options(void* self, intptr_t slot);
Options QFileIconProvider_virtualbase_Options(const void* self);
extern __declspec(dllexport) void QFileIconProvider_Delete(QFileIconProvider* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
