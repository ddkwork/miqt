#pragma once
#ifndef MIQT_QT_GEN_QFILE_H
#define MIQT_QT_GEN_QFILE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QFile QFile;
typedef struct QFileDevice QFileDevice;
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;
typedef struct QMetaObject QMetaObject;
typedef struct QNtfsPermissionCheckGuard QNtfsPermissionCheckGuard;
typedef struct QObject QObject;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QNtfsPermissionCheckGuard* QNtfsPermissionCheckGuard_new();
extern __declspec(dllexport) void QNtfsPermissionCheckGuard_Delete(QNtfsPermissionCheckGuard* self, bool isSubclass);

extern __declspec(dllexport) QFile* QFile_new();
extern __declspec(dllexport) QFile* QFile_new2(struct miqt_string name);
extern __declspec(dllexport) QFile* QFile_new3(QObject* parent);
extern __declspec(dllexport) QFile* QFile_new4(struct miqt_string name, QObject* parent);
extern __declspec(dllexport) void QFile_virtbase(QFile* src, QFileDevice** outptr_QFileDevice);
extern __declspec(dllexport) QMetaObject* QFile_MetaObject(const QFile* self);
extern __declspec(dllexport) void* QFile_Metacast(QFile* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QFile_Tr(const char* s);
extern __declspec(dllexport) struct miqt_string QFile_FileName(const QFile* self);
extern __declspec(dllexport) void QFile_SetFileName(QFile* self, struct miqt_string name);
extern __declspec(dllexport) struct miqt_string QFile_EncodeName(struct miqt_string fileName);
extern __declspec(dllexport) struct miqt_string QFile_DecodeName(struct miqt_string localFileName);
extern __declspec(dllexport) struct miqt_string QFile_DecodeNameWithLocalFileName(const char* localFileName);
extern __declspec(dllexport) bool QFile_Exists(const QFile* self);
extern __declspec(dllexport) bool QFile_ExistsWithFileName(struct miqt_string fileName);
extern __declspec(dllexport) struct miqt_string QFile_SymLinkTarget(const QFile* self);
extern __declspec(dllexport) struct miqt_string QFile_SymLinkTargetWithFileName(struct miqt_string fileName);
extern __declspec(dllexport) bool QFile_Remove(QFile* self);
extern __declspec(dllexport) bool QFile_RemoveWithFileName(struct miqt_string fileName);
extern __declspec(dllexport) bool QFile_SupportsMoveToTrash();
extern __declspec(dllexport) bool QFile_MoveToTrash(QFile* self);
extern __declspec(dllexport) bool QFile_MoveToTrashWithFileName(struct miqt_string fileName);
extern __declspec(dllexport) bool QFile_Rename(QFile* self, struct miqt_string newName);
extern __declspec(dllexport) bool QFile_Rename2(struct miqt_string oldName, struct miqt_string newName);
extern __declspec(dllexport) bool QFile_Link(QFile* self, struct miqt_string newName);
extern __declspec(dllexport) bool QFile_Link2(struct miqt_string fileName, struct miqt_string newName);
extern __declspec(dllexport) bool QFile_Copy(QFile* self, struct miqt_string newName);
extern __declspec(dllexport) bool QFile_Copy2(struct miqt_string fileName, struct miqt_string newName);
extern __declspec(dllexport) bool QFile_Open(QFile* self, OpenMode flags);
extern __declspec(dllexport) bool QFile_Open2(QFile* self, OpenMode flags, Permissions permissions);
extern __declspec(dllexport) bool QFile_Open4(QFile* self, int fd, OpenMode ioFlags);
extern __declspec(dllexport) long long QFile_Size(const QFile* self);
extern __declspec(dllexport) bool QFile_Resize(QFile* self, long long sz);
extern __declspec(dllexport) bool QFile_Resize2(struct miqt_string filename, long long sz);
extern __declspec(dllexport) Permissions QFile_Permissions(const QFile* self);
extern __declspec(dllexport) Permissions QFile_PermissionsWithFilename(struct miqt_string filename);
extern __declspec(dllexport) bool QFile_SetPermissions(QFile* self, Permissions permissionSpec);
extern __declspec(dllexport) bool QFile_SetPermissions2(struct miqt_string filename, Permissions permissionSpec);
extern __declspec(dllexport) struct miqt_string QFile_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QFile_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) bool QFile_Open33(QFile* self, int fd, OpenMode ioFlags, FileHandleFlags handleFlags);
extern __declspec(dllexport) void QFile_override_virtual_FileName(void* self, intptr_t slot);
struct miqt_string QFile_virtualbase_FileName(const void* self);
extern __declspec(dllexport) void QFile_override_virtual_Open(void* self, intptr_t slot);
bool QFile_virtualbase_Open(void* self, OpenMode flags);
extern __declspec(dllexport) void QFile_override_virtual_Size(void* self, intptr_t slot);
long long QFile_virtualbase_Size(const void* self);
extern __declspec(dllexport) void QFile_override_virtual_Resize(void* self, intptr_t slot);
bool QFile_virtualbase_Resize(void* self, long long sz);
extern __declspec(dllexport) void QFile_override_virtual_Permissions(void* self, intptr_t slot);
Permissions QFile_virtualbase_Permissions(const void* self);
extern __declspec(dllexport) void QFile_override_virtual_SetPermissions(void* self, intptr_t slot);
bool QFile_virtualbase_SetPermissions(void* self, Permissions permissionSpec);
extern __declspec(dllexport) void QFile_override_virtual_Close(void* self, intptr_t slot);
void QFile_virtualbase_Close(void* self);
extern __declspec(dllexport) void QFile_override_virtual_IsSequential(void* self, intptr_t slot);
bool QFile_virtualbase_IsSequential(const void* self);
extern __declspec(dllexport) void QFile_override_virtual_Pos(void* self, intptr_t slot);
long long QFile_virtualbase_Pos(const void* self);
extern __declspec(dllexport) void QFile_override_virtual_Seek(void* self, intptr_t slot);
bool QFile_virtualbase_Seek(void* self, long long offset);
extern __declspec(dllexport) void QFile_override_virtual_AtEnd(void* self, intptr_t slot);
bool QFile_virtualbase_AtEnd(const void* self);
extern __declspec(dllexport) void QFile_override_virtual_ReadData(void* self, intptr_t slot);
long long QFile_virtualbase_ReadData(void* self, char* data, long long maxlen);
extern __declspec(dllexport) void QFile_override_virtual_WriteData(void* self, intptr_t slot);
long long QFile_virtualbase_WriteData(void* self, const char* data, long long lenVal);
extern __declspec(dllexport) void QFile_override_virtual_ReadLineData(void* self, intptr_t slot);
long long QFile_virtualbase_ReadLineData(void* self, char* data, long long maxlen);
extern __declspec(dllexport) void QFile_Delete(QFile* self, bool isSubclass);

} 
