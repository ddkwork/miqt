#pragma once
#ifndef MIQT_QT_GEN_QFILEDEVICE_H
#define MIQT_QT_GEN_QFILEDEVICE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDateTime QDateTime;
typedef struct QFileDevice QFileDevice;
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QFileDevice_virtbase(QFileDevice* src, QIODevice** outptr_QIODevice);
extern __declspec(dllexport) QMetaObject* QFileDevice_MetaObject(const QFileDevice* self);
extern __declspec(dllexport) void* QFileDevice_Metacast(QFileDevice* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QFileDevice_Tr(const char* s);
extern __declspec(dllexport) FileError QFileDevice_Error(const QFileDevice* self);
extern __declspec(dllexport) void QFileDevice_UnsetError(QFileDevice* self);
extern __declspec(dllexport) void QFileDevice_Close(QFileDevice* self);
extern __declspec(dllexport) bool QFileDevice_IsSequential(const QFileDevice* self);
extern __declspec(dllexport) int QFileDevice_Handle(const QFileDevice* self);
extern __declspec(dllexport) struct miqt_string QFileDevice_FileName(const QFileDevice* self);
extern __declspec(dllexport) long long QFileDevice_Pos(const QFileDevice* self);
extern __declspec(dllexport) bool QFileDevice_Seek(QFileDevice* self, long long offset);
extern __declspec(dllexport) bool QFileDevice_AtEnd(const QFileDevice* self);
extern __declspec(dllexport) bool QFileDevice_Flush(QFileDevice* self);
extern __declspec(dllexport) long long QFileDevice_Size(const QFileDevice* self);
extern __declspec(dllexport) bool QFileDevice_Resize(QFileDevice* self, long long sz);
extern __declspec(dllexport) Permissions QFileDevice_Permissions(const QFileDevice* self);
extern __declspec(dllexport) bool QFileDevice_SetPermissions(QFileDevice* self, Permissions permissionSpec);
extern __declspec(dllexport) unsigned char* QFileDevice_Map(QFileDevice* self, long long offset, long long size);
extern __declspec(dllexport) bool QFileDevice_Unmap(QFileDevice* self, unsigned char* address);
extern __declspec(dllexport) QDateTime* QFileDevice_FileTime(const QFileDevice* self, int time);
extern __declspec(dllexport) bool QFileDevice_SetFileTime(QFileDevice* self, QDateTime* newDate, int fileTime);
extern __declspec(dllexport) long long QFileDevice_ReadData(QFileDevice* self, char* data, long long maxlen);
extern __declspec(dllexport) long long QFileDevice_WriteData(QFileDevice* self, const char* data, long long lenVal);
extern __declspec(dllexport) long long QFileDevice_ReadLineData(QFileDevice* self, char* data, long long maxlen);
extern __declspec(dllexport) struct miqt_string QFileDevice_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QFileDevice_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) unsigned char* QFileDevice_Map3(QFileDevice* self, long long offset, long long size, MemoryMapFlags flags);
extern __declspec(dllexport) void QFileDevice_Delete(QFileDevice* self, bool isSubclass);

} 
