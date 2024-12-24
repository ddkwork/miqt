#pragma once
#ifndef MIQT_QT_GEN_QDIRITERATOR_H
#define MIQT_QT_GEN_QDIRITERATOR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QDir;
class QDirIterator;
class QFileInfo;
class _GUID;
class type_info;
#else
typedef struct QDir QDir;
typedef struct QDirIterator QDirIterator;
typedef struct QFileInfo QFileInfo;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QDirIterator* QDirIterator_new(QDir* dir);
extern __declspec(dllexport) QDirIterator* QDirIterator_new2(struct miqt_string path);
extern __declspec(dllexport) QDirIterator* QDirIterator_new3(struct miqt_string path, int filter);
extern __declspec(dllexport) QDirIterator* QDirIterator_new4(struct miqt_string path, struct miqt_array /* of struct miqt_string */  nameFilters);
extern __declspec(dllexport) QDirIterator* QDirIterator_new5(QDir* dir, IteratorFlags flags);
extern __declspec(dllexport) QDirIterator* QDirIterator_new6(struct miqt_string path, IteratorFlags flags);
extern __declspec(dllexport) QDirIterator* QDirIterator_new7(struct miqt_string path, int filter, IteratorFlags flags);
extern __declspec(dllexport) QDirIterator* QDirIterator_new8(struct miqt_string path, struct miqt_array /* of struct miqt_string */  nameFilters, int filters);
extern __declspec(dllexport) QDirIterator* QDirIterator_new9(struct miqt_string path, struct miqt_array /* of struct miqt_string */  nameFilters, int filters, IteratorFlags flags);
extern __declspec(dllexport) struct miqt_string QDirIterator_Next(QDirIterator* self);
extern __declspec(dllexport) QFileInfo* QDirIterator_NextFileInfo(QDirIterator* self);
extern __declspec(dllexport) bool QDirIterator_HasNext(const QDirIterator* self);
extern __declspec(dllexport) struct miqt_string QDirIterator_FileName(const QDirIterator* self);
extern __declspec(dllexport) struct miqt_string QDirIterator_FilePath(const QDirIterator* self);
extern __declspec(dllexport) QFileInfo* QDirIterator_FileInfo(const QDirIterator* self);
extern __declspec(dllexport) struct miqt_string QDirIterator_Path(const QDirIterator* self);
extern __declspec(dllexport) void QDirIterator_Delete(QDirIterator* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
