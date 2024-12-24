#pragma once
#ifndef MIQT_QT_GEN_QTEMPORARYDIR_H
#define MIQT_QT_GEN_QTEMPORARYDIR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QTemporaryDir;
class _GUID;
class type_info;
#else
typedef struct QTemporaryDir QTemporaryDir;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTemporaryDir* QTemporaryDir_new();
extern __declspec(dllexport) QTemporaryDir* QTemporaryDir_new2(struct miqt_string templateName);
extern __declspec(dllexport) void QTemporaryDir_Swap(QTemporaryDir* self, QTemporaryDir* other);
extern __declspec(dllexport) bool QTemporaryDir_IsValid(const QTemporaryDir* self);
extern __declspec(dllexport) struct miqt_string QTemporaryDir_ErrorString(const QTemporaryDir* self);
extern __declspec(dllexport) bool QTemporaryDir_AutoRemove(const QTemporaryDir* self);
extern __declspec(dllexport) void QTemporaryDir_SetAutoRemove(QTemporaryDir* self, bool b);
extern __declspec(dllexport) bool QTemporaryDir_Remove(QTemporaryDir* self);
extern __declspec(dllexport) struct miqt_string QTemporaryDir_Path(const QTemporaryDir* self);
extern __declspec(dllexport) struct miqt_string QTemporaryDir_FilePath(const QTemporaryDir* self, struct miqt_string fileName);
extern __declspec(dllexport) void QTemporaryDir_Delete(QTemporaryDir* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
