#pragma once
#ifndef MIQT_QT_GEN_QSTORAGEINFO_H
#define MIQT_QT_GEN_QSTORAGEINFO_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDir QDir;
typedef struct QStorageInfo QStorageInfo;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QStorageInfo* QStorageInfo_new();
extern __declspec(dllexport) QStorageInfo* QStorageInfo_new2(struct miqt_string path);
extern __declspec(dllexport) QStorageInfo* QStorageInfo_new3(QDir* dir);
extern __declspec(dllexport) QStorageInfo* QStorageInfo_new4(QStorageInfo* other);
extern __declspec(dllexport) void QStorageInfo_OperatorAssign(QStorageInfo* self, QStorageInfo* other);
extern __declspec(dllexport) void QStorageInfo_Swap(QStorageInfo* self, QStorageInfo* other);
extern __declspec(dllexport) void QStorageInfo_SetPath(QStorageInfo* self, struct miqt_string path);
extern __declspec(dllexport) struct miqt_string QStorageInfo_RootPath(const QStorageInfo* self);
extern __declspec(dllexport) struct miqt_string QStorageInfo_Device(const QStorageInfo* self);
extern __declspec(dllexport) struct miqt_string QStorageInfo_Subvolume(const QStorageInfo* self);
extern __declspec(dllexport) struct miqt_string QStorageInfo_FileSystemType(const QStorageInfo* self);
extern __declspec(dllexport) struct miqt_string QStorageInfo_Name(const QStorageInfo* self);
extern __declspec(dllexport) struct miqt_string QStorageInfo_DisplayName(const QStorageInfo* self);
extern __declspec(dllexport) long long QStorageInfo_BytesTotal(const QStorageInfo* self);
extern __declspec(dllexport) long long QStorageInfo_BytesFree(const QStorageInfo* self);
extern __declspec(dllexport) long long QStorageInfo_BytesAvailable(const QStorageInfo* self);
extern __declspec(dllexport) int QStorageInfo_BlockSize(const QStorageInfo* self);
extern __declspec(dllexport) bool QStorageInfo_IsRoot(const QStorageInfo* self);
extern __declspec(dllexport) bool QStorageInfo_IsReadOnly(const QStorageInfo* self);
extern __declspec(dllexport) bool QStorageInfo_IsReady(const QStorageInfo* self);
extern __declspec(dllexport) bool QStorageInfo_IsValid(const QStorageInfo* self);
extern __declspec(dllexport) void QStorageInfo_Refresh(QStorageInfo* self);
extern __declspec(dllexport) struct miqt_array /* of QStorageInfo* */  QStorageInfo_MountedVolumes();
extern __declspec(dllexport) QStorageInfo* QStorageInfo_Root();
extern __declspec(dllexport) void QStorageInfo_Delete(QStorageInfo* self, bool isSubclass);

} 
