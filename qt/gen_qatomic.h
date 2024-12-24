#pragma once
#ifndef MIQT_QT_GEN_QATOMIC_H
#define MIQT_QT_GEN_QATOMIC_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAtomicInt;
class _GUID;
class type_info;
#else
typedef struct QAtomicInt QAtomicInt;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QAtomicInt* QAtomicInt_new();
extern __declspec(dllexport) QAtomicInt* QAtomicInt_new2(QAtomicInt* param1);
extern __declspec(dllexport) QAtomicInt* QAtomicInt_new3(int value);
extern __declspec(dllexport) void QAtomicInt_OperatorAssign(QAtomicInt* self, QAtomicInt* param1);
extern __declspec(dllexport) void QAtomicInt_Delete(QAtomicInt* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
