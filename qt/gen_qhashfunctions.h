#pragma once
#ifndef MIQT_QT_GEN_QHASHFUNCTIONS_H
#define MIQT_QT_GEN_QHASHFUNCTIONS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QHashSeed;
class _GUID;
class type_info;
#else
typedef struct QHashSeed QHashSeed;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QHashSeed* QHashSeed_new();
extern __declspec(dllexport) QHashSeed* QHashSeed_new2(unsigned long long d);
extern __declspec(dllexport) QHashSeed* QHashSeed_GlobalSeed();
extern __declspec(dllexport) void QHashSeed_SetDeterministicGlobalSeed();
extern __declspec(dllexport) void QHashSeed_ResetRandomGlobalSeed();
extern __declspec(dllexport) void QHashSeed_Delete(QHashSeed* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
