#pragma once
#ifndef MIQT_QT_GEN_QVARLENGTHARRAY_H
#define MIQT_QT_GEN_QVARLENGTHARRAY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QVLABaseBase;
class _GUID;
class type_info;
#else
typedef struct QVLABaseBase QVLABaseBase;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) size_type QVLABaseBase_Capacity(const QVLABaseBase* self);
extern __declspec(dllexport) size_type QVLABaseBase_Size(const QVLABaseBase* self);
extern __declspec(dllexport) bool QVLABaseBase_Empty(const QVLABaseBase* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
