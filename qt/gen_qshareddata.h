#pragma once
#ifndef MIQT_QT_GEN_QSHAREDDATA_H
#define MIQT_QT_GEN_QSHAREDDATA_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAdoptSharedDataTag;
class QSharedData;
class _GUID;
class type_info;
#else
typedef struct QAdoptSharedDataTag QAdoptSharedDataTag;
typedef struct QSharedData QSharedData;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSharedData* QSharedData_new();
extern __declspec(dllexport) QSharedData* QSharedData_new2(QSharedData* param1);
extern __declspec(dllexport) void QSharedData_Delete(QSharedData* self, bool isSubclass);

extern __declspec(dllexport) QAdoptSharedDataTag* QAdoptSharedDataTag_new();
extern __declspec(dllexport) void QAdoptSharedDataTag_Delete(QAdoptSharedDataTag* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
