#pragma once
#ifndef MIQT_QT_GEN_QFILEINFO_H
#define MIQT_QT_GEN_QFILEINFO_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDateTime QDateTime;
typedef struct QDir QDir;
typedef struct QFileDevice QFileDevice;
typedef struct QFileInfo QFileInfo;
typedef struct QTimeZone QTimeZone;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QFileInfo* QFileInfo_new();
extern __declspec(dllexport) QFileInfo* QFileInfo_new2(struct miqt_string file);
extern __declspec(dllexport) QFileInfo* QFileInfo_new3(QFileDevice* file);
extern __declspec(dllexport) QFileInfo* QFileInfo_new4(QDir* dir, struct miqt_string file);
extern __declspec(dllexport) QFileInfo* QFileInfo_new5(QFileInfo* fileinfo);
extern __declspec(dllexport) void QFileInfo_OperatorAssign(QFileInfo* self, QFileInfo* fileinfo);
extern __declspec(dllexport) void QFileInfo_Swap(QFileInfo* self, QFileInfo* other);
extern __declspec(dllexport) void QFileInfo_SetFile(QFileInfo* self, struct miqt_string file);
extern __declspec(dllexport) void QFileInfo_SetFileWithFile(QFileInfo* self, QFileDevice* file);
extern __declspec(dllexport) void QFileInfo_SetFile2(QFileInfo* self, QDir* dir, struct miqt_string file);
extern __declspec(dllexport) bool QFileInfo_Exists(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_ExistsWithFile(struct miqt_string file);
extern __declspec(dllexport) void QFileInfo_Refresh(QFileInfo* self);
extern __declspec(dllexport) struct miqt_string QFileInfo_FilePath(const QFileInfo* self);
extern __declspec(dllexport) struct miqt_string QFileInfo_AbsoluteFilePath(const QFileInfo* self);
extern __declspec(dllexport) struct miqt_string QFileInfo_CanonicalFilePath(const QFileInfo* self);
extern __declspec(dllexport) struct miqt_string QFileInfo_FileName(const QFileInfo* self);
extern __declspec(dllexport) struct miqt_string QFileInfo_BaseName(const QFileInfo* self);
extern __declspec(dllexport) struct miqt_string QFileInfo_CompleteBaseName(const QFileInfo* self);
extern __declspec(dllexport) struct miqt_string QFileInfo_Suffix(const QFileInfo* self);
extern __declspec(dllexport) struct miqt_string QFileInfo_BundleName(const QFileInfo* self);
extern __declspec(dllexport) struct miqt_string QFileInfo_CompleteSuffix(const QFileInfo* self);
extern __declspec(dllexport) struct miqt_string QFileInfo_Path(const QFileInfo* self);
extern __declspec(dllexport) struct miqt_string QFileInfo_AbsolutePath(const QFileInfo* self);
extern __declspec(dllexport) struct miqt_string QFileInfo_CanonicalPath(const QFileInfo* self);
extern __declspec(dllexport) QDir* QFileInfo_Dir(const QFileInfo* self);
extern __declspec(dllexport) QDir* QFileInfo_AbsoluteDir(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_IsReadable(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_IsWritable(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_IsExecutable(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_IsHidden(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_IsNativePath(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_IsRelative(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_IsAbsolute(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_MakeAbsolute(QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_IsFile(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_IsDir(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_IsSymLink(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_IsSymbolicLink(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_IsShortcut(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_IsAlias(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_IsJunction(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_IsRoot(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_IsBundle(const QFileInfo* self);
extern __declspec(dllexport) struct miqt_string QFileInfo_SymLinkTarget(const QFileInfo* self);
extern __declspec(dllexport) struct miqt_string QFileInfo_ReadSymLink(const QFileInfo* self);
extern __declspec(dllexport) struct miqt_string QFileInfo_JunctionTarget(const QFileInfo* self);
extern __declspec(dllexport) struct miqt_string QFileInfo_Owner(const QFileInfo* self);
extern __declspec(dllexport) unsigned int QFileInfo_OwnerId(const QFileInfo* self);
extern __declspec(dllexport) struct miqt_string QFileInfo_Group(const QFileInfo* self);
extern __declspec(dllexport) unsigned int QFileInfo_GroupId(const QFileInfo* self);
extern __declspec(dllexport) bool QFileInfo_Permission(const QFileInfo* self, int permissions);
extern __declspec(dllexport) int QFileInfo_Permissions(const QFileInfo* self);
extern __declspec(dllexport) long long QFileInfo_Size(const QFileInfo* self);
extern __declspec(dllexport) QDateTime* QFileInfo_BirthTime(const QFileInfo* self);
extern __declspec(dllexport) QDateTime* QFileInfo_MetadataChangeTime(const QFileInfo* self);
extern __declspec(dllexport) QDateTime* QFileInfo_LastModified(const QFileInfo* self);
extern __declspec(dllexport) QDateTime* QFileInfo_LastRead(const QFileInfo* self);
extern __declspec(dllexport) QDateTime* QFileInfo_FileTime(const QFileInfo* self, int time);
extern __declspec(dllexport) QDateTime* QFileInfo_BirthTimeWithTz(const QFileInfo* self, QTimeZone* tz);
extern __declspec(dllexport) QDateTime* QFileInfo_MetadataChangeTimeWithTz(const QFileInfo* self, QTimeZone* tz);
extern __declspec(dllexport) QDateTime* QFileInfo_LastModifiedWithTz(const QFileInfo* self, QTimeZone* tz);
extern __declspec(dllexport) QDateTime* QFileInfo_LastReadWithTz(const QFileInfo* self, QTimeZone* tz);
extern __declspec(dllexport) QDateTime* QFileInfo_FileTime2(const QFileInfo* self, int time, QTimeZone* tz);
extern __declspec(dllexport) bool QFileInfo_Caching(const QFileInfo* self);
extern __declspec(dllexport) void QFileInfo_SetCaching(QFileInfo* self, bool on);
extern __declspec(dllexport) void QFileInfo_Stat(QFileInfo* self);
extern __declspec(dllexport) void QFileInfo_Delete(QFileInfo* self, bool isSubclass);

} 
