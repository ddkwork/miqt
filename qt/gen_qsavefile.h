#pragma once
#ifndef MIQT_QT_GEN_QSAVEFILE_H
#define MIQT_QT_GEN_QSAVEFILE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QFileDevice;
class QIODevice;
class QIODeviceBase;
class QMetaObject;
class QObject;
class QSaveFile;
class _GUID;
class type_info;
#else
typedef struct QFileDevice QFileDevice;
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSaveFile QSaveFile;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSaveFile* QSaveFile_new(struct miqt_string name);
extern __declspec(dllexport) QSaveFile* QSaveFile_new2();
extern __declspec(dllexport) QSaveFile* QSaveFile_new3(struct miqt_string name, QObject* parent);
extern __declspec(dllexport) QSaveFile* QSaveFile_new4(QObject* parent);
extern __declspec(dllexport) void QSaveFile_virtbase(QSaveFile* src, QFileDevice** outptr_QFileDevice);
extern __declspec(dllexport) QMetaObject* QSaveFile_MetaObject(const QSaveFile* self);
extern __declspec(dllexport) void* QSaveFile_Metacast(QSaveFile* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QSaveFile_Tr(const char* s);
extern __declspec(dllexport) struct miqt_string QSaveFile_FileName(const QSaveFile* self);
extern __declspec(dllexport) void QSaveFile_SetFileName(QSaveFile* self, struct miqt_string name);
extern __declspec(dllexport) bool QSaveFile_Open(QSaveFile* self, OpenMode flags);
extern __declspec(dllexport) bool QSaveFile_Commit(QSaveFile* self);
extern __declspec(dllexport) void QSaveFile_CancelWriting(QSaveFile* self);
extern __declspec(dllexport) void QSaveFile_SetDirectWriteFallback(QSaveFile* self, bool enabled);
extern __declspec(dllexport) bool QSaveFile_DirectWriteFallback(const QSaveFile* self);
extern __declspec(dllexport) long long QSaveFile_WriteData(QSaveFile* self, const char* data, long long lenVal);
extern __declspec(dllexport) struct miqt_string QSaveFile_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QSaveFile_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QSaveFile_override_virtual_FileName(void* self, intptr_t slot);
struct miqt_string QSaveFile_virtualbase_FileName(const void* self);
extern __declspec(dllexport) void QSaveFile_override_virtual_Open(void* self, intptr_t slot);
bool QSaveFile_virtualbase_Open(void* self, OpenMode flags);
extern __declspec(dllexport) void QSaveFile_override_virtual_WriteData(void* self, intptr_t slot);
long long QSaveFile_virtualbase_WriteData(void* self, const char* data, long long lenVal);
extern __declspec(dllexport) void QSaveFile_override_virtual_IsSequential(void* self, intptr_t slot);
bool QSaveFile_virtualbase_IsSequential(const void* self);
extern __declspec(dllexport) void QSaveFile_override_virtual_Pos(void* self, intptr_t slot);
long long QSaveFile_virtualbase_Pos(const void* self);
extern __declspec(dllexport) void QSaveFile_override_virtual_Seek(void* self, intptr_t slot);
bool QSaveFile_virtualbase_Seek(void* self, long long offset);
extern __declspec(dllexport) void QSaveFile_override_virtual_AtEnd(void* self, intptr_t slot);
bool QSaveFile_virtualbase_AtEnd(const void* self);
extern __declspec(dllexport) void QSaveFile_override_virtual_Size(void* self, intptr_t slot);
long long QSaveFile_virtualbase_Size(const void* self);
extern __declspec(dllexport) void QSaveFile_override_virtual_Resize(void* self, intptr_t slot);
bool QSaveFile_virtualbase_Resize(void* self, long long sz);
extern __declspec(dllexport) void QSaveFile_override_virtual_Permissions(void* self, intptr_t slot);
Permissions QSaveFile_virtualbase_Permissions(const void* self);
extern __declspec(dllexport) void QSaveFile_override_virtual_SetPermissions(void* self, intptr_t slot);
bool QSaveFile_virtualbase_SetPermissions(void* self, Permissions permissionSpec);
extern __declspec(dllexport) void QSaveFile_override_virtual_ReadData(void* self, intptr_t slot);
long long QSaveFile_virtualbase_ReadData(void* self, char* data, long long maxlen);
extern __declspec(dllexport) void QSaveFile_override_virtual_ReadLineData(void* self, intptr_t slot);
long long QSaveFile_virtualbase_ReadLineData(void* self, char* data, long long maxlen);
extern __declspec(dllexport) void QSaveFile_Delete(QSaveFile* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
