#pragma once
#ifndef MIQT_QT_GEN_QARRAYDATA_H
#define MIQT_QT_GEN_QARRAYDATA_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QArrayData;
class _GUID;
class type_info;
#else
typedef struct QArrayData QArrayData;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) ptrdiff_t QArrayData_AllocatedCapacity(QArrayData* self);
extern __declspec(dllexport) ptrdiff_t QArrayData_ConstAllocatedCapacity(const QArrayData* self);
extern __declspec(dllexport) bool QArrayData_Ref(QArrayData* self);
extern __declspec(dllexport) bool QArrayData_Deref(QArrayData* self);
extern __declspec(dllexport) bool QArrayData_IsShared(const QArrayData* self);
extern __declspec(dllexport) bool QArrayData_NeedsDetach(QArrayData* self);
extern __declspec(dllexport) ptrdiff_t QArrayData_DetachCapacity(const QArrayData* self, ptrdiff_t newSize);
extern __declspec(dllexport) void QArrayData_Deallocate(QArrayData* data, ptrdiff_t objectSize, ptrdiff_t alignment);
extern __declspec(dllexport) void QArrayData_Delete(QArrayData* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
