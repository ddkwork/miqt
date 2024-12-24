#pragma once
#ifndef MIQT_QT_GEN_QPROPERTYPRIVATE_H
#define MIQT_QT_GEN_QPROPERTYPRIVATE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QPropertyProxyBindingData;
class QUntypedPropertyData;
class _GUID;
class type_info;
#else
typedef struct QPropertyProxyBindingData QPropertyProxyBindingData;
typedef struct QUntypedPropertyData QUntypedPropertyData;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QUntypedPropertyData_Delete(QUntypedPropertyData* self, bool isSubclass);

extern __declspec(dllexport) void QPropertyProxyBindingData_Delete(QPropertyProxyBindingData* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
