#pragma once
#ifndef MIQT_QT_GEN_QMIMETYPE_H
#define MIQT_QT_GEN_QMIMETYPE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QMimeType;
class _GUID;
class type_info;
#else
typedef struct QMimeType QMimeType;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QMimeType* QMimeType_new();
extern __declspec(dllexport) QMimeType* QMimeType_new2(QMimeType* other);
extern __declspec(dllexport) void QMimeType_OperatorAssign(QMimeType* self, QMimeType* other);
extern __declspec(dllexport) void QMimeType_Swap(QMimeType* self, QMimeType* other);
extern __declspec(dllexport) bool QMimeType_IsValid(const QMimeType* self);
extern __declspec(dllexport) bool QMimeType_IsDefault(const QMimeType* self);
extern __declspec(dllexport) struct miqt_string QMimeType_Name(const QMimeType* self);
extern __declspec(dllexport) struct miqt_string QMimeType_Comment(const QMimeType* self);
extern __declspec(dllexport) struct miqt_string QMimeType_GenericIconName(const QMimeType* self);
extern __declspec(dllexport) struct miqt_string QMimeType_IconName(const QMimeType* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QMimeType_GlobPatterns(const QMimeType* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QMimeType_ParentMimeTypes(const QMimeType* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QMimeType_AllAncestors(const QMimeType* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QMimeType_Aliases(const QMimeType* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QMimeType_Suffixes(const QMimeType* self);
extern __declspec(dllexport) struct miqt_string QMimeType_PreferredSuffix(const QMimeType* self);
extern __declspec(dllexport) bool QMimeType_Inherits(const QMimeType* self, struct miqt_string mimeTypeName);
extern __declspec(dllexport) struct miqt_string QMimeType_FilterString(const QMimeType* self);
extern __declspec(dllexport) void QMimeType_Delete(QMimeType* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
