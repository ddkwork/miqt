#pragma once
#ifndef MIQT_QT_GEN_QABSTRACTNATIVEEVENTFILTER_H
#define MIQT_QT_GEN_QABSTRACTNATIVEEVENTFILTER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractNativeEventFilter;
class _GUID;
class type_info;
#else
typedef struct QAbstractNativeEventFilter QAbstractNativeEventFilter;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QAbstractNativeEventFilter* QAbstractNativeEventFilter_new();
extern __declspec(dllexport) bool QAbstractNativeEventFilter_NativeEventFilter(QAbstractNativeEventFilter* self, struct miqt_string eventType, void* message, intptr_t* result);
extern __declspec(dllexport) void QAbstractNativeEventFilter_override_virtual_NativeEventFilter(void* self, intptr_t slot);
bool QAbstractNativeEventFilter_virtualbase_NativeEventFilter(void* self, struct miqt_string eventType, void* message, intptr_t* result);
extern __declspec(dllexport) void QAbstractNativeEventFilter_Delete(QAbstractNativeEventFilter* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
