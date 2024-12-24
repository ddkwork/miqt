#pragma once
#ifndef MIQT_QT_GEN_QSCOPEDPOINTER_H
#define MIQT_QT_GEN_QSCOPEDPOINTER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QScopedPointerPodDeleter;
class _GUID;
class type_info;
#else
typedef struct QScopedPointerPodDeleter QScopedPointerPodDeleter;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QScopedPointerPodDeleter_Cleanup(void* pointer);
extern __declspec(dllexport) void QScopedPointerPodDeleter_OperatorCall(const QScopedPointerPodDeleter* self, void* pointer);
extern __declspec(dllexport) void QScopedPointerPodDeleter_Delete(QScopedPointerPodDeleter* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
