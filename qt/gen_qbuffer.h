#pragma once
#ifndef MIQT_QT_GEN_QBUFFER_H
#define MIQT_QT_GEN_QBUFFER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QBuffer QBuffer;
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QBuffer* QBuffer_new();
extern __declspec(dllexport) QBuffer* QBuffer_new2(QObject* parent);
extern __declspec(dllexport) void QBuffer_virtbase(QBuffer* src, QIODevice** outptr_QIODevice);
extern __declspec(dllexport) QMetaObject* QBuffer_MetaObject(const QBuffer* self);
extern __declspec(dllexport) void* QBuffer_Metacast(QBuffer* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QBuffer_Tr(const char* s);
extern __declspec(dllexport) struct miqt_string QBuffer_Buffer(QBuffer* self);
extern __declspec(dllexport) struct miqt_string QBuffer_Buffer2(const QBuffer* self);
extern __declspec(dllexport) void QBuffer_SetData(QBuffer* self, struct miqt_string data);
extern __declspec(dllexport) void QBuffer_SetData2(QBuffer* self, const char* data, ptrdiff_t lenVal);
extern __declspec(dllexport) struct miqt_string QBuffer_Data(const QBuffer* self);
extern __declspec(dllexport) bool QBuffer_Open(QBuffer* self, OpenMode openMode);
extern __declspec(dllexport) void QBuffer_Close(QBuffer* self);
extern __declspec(dllexport) long long QBuffer_Size(const QBuffer* self);
extern __declspec(dllexport) long long QBuffer_Pos(const QBuffer* self);
extern __declspec(dllexport) bool QBuffer_Seek(QBuffer* self, long long off);
extern __declspec(dllexport) bool QBuffer_AtEnd(const QBuffer* self);
extern __declspec(dllexport) bool QBuffer_CanReadLine(const QBuffer* self);
extern __declspec(dllexport) void QBuffer_ConnectNotify(QBuffer* self, QMetaMethod* param1);
extern __declspec(dllexport) void QBuffer_DisconnectNotify(QBuffer* self, QMetaMethod* param1);
extern __declspec(dllexport) long long QBuffer_ReadData(QBuffer* self, char* data, long long maxlen);
extern __declspec(dllexport) long long QBuffer_WriteData(QBuffer* self, const char* data, long long lenVal);
extern __declspec(dllexport) struct miqt_string QBuffer_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QBuffer_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QBuffer_override_virtual_Open(void* self, intptr_t slot);
bool QBuffer_virtualbase_Open(void* self, OpenMode openMode);
extern __declspec(dllexport) void QBuffer_override_virtual_Close(void* self, intptr_t slot);
void QBuffer_virtualbase_Close(void* self);
extern __declspec(dllexport) void QBuffer_override_virtual_Size(void* self, intptr_t slot);
long long QBuffer_virtualbase_Size(const void* self);
extern __declspec(dllexport) void QBuffer_override_virtual_Pos(void* self, intptr_t slot);
long long QBuffer_virtualbase_Pos(const void* self);
extern __declspec(dllexport) void QBuffer_override_virtual_Seek(void* self, intptr_t slot);
bool QBuffer_virtualbase_Seek(void* self, long long off);
extern __declspec(dllexport) void QBuffer_override_virtual_AtEnd(void* self, intptr_t slot);
bool QBuffer_virtualbase_AtEnd(const void* self);
extern __declspec(dllexport) void QBuffer_override_virtual_CanReadLine(void* self, intptr_t slot);
bool QBuffer_virtualbase_CanReadLine(const void* self);
extern __declspec(dllexport) void QBuffer_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QBuffer_virtualbase_ConnectNotify(void* self, QMetaMethod* param1);
extern __declspec(dllexport) void QBuffer_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QBuffer_virtualbase_DisconnectNotify(void* self, QMetaMethod* param1);
extern __declspec(dllexport) void QBuffer_override_virtual_ReadData(void* self, intptr_t slot);
long long QBuffer_virtualbase_ReadData(void* self, char* data, long long maxlen);
extern __declspec(dllexport) void QBuffer_override_virtual_WriteData(void* self, intptr_t slot);
long long QBuffer_virtualbase_WriteData(void* self, const char* data, long long lenVal);
extern __declspec(dllexport) void QBuffer_override_virtual_IsSequential(void* self, intptr_t slot);
bool QBuffer_virtualbase_IsSequential(const void* self);
extern __declspec(dllexport) void QBuffer_override_virtual_Reset(void* self, intptr_t slot);
bool QBuffer_virtualbase_Reset(void* self);
extern __declspec(dllexport) void QBuffer_override_virtual_BytesAvailable(void* self, intptr_t slot);
long long QBuffer_virtualbase_BytesAvailable(const void* self);
extern __declspec(dllexport) void QBuffer_override_virtual_BytesToWrite(void* self, intptr_t slot);
long long QBuffer_virtualbase_BytesToWrite(const void* self);
extern __declspec(dllexport) void QBuffer_override_virtual_WaitForReadyRead(void* self, intptr_t slot);
bool QBuffer_virtualbase_WaitForReadyRead(void* self, int msecs);
extern __declspec(dllexport) void QBuffer_override_virtual_WaitForBytesWritten(void* self, intptr_t slot);
bool QBuffer_virtualbase_WaitForBytesWritten(void* self, int msecs);
extern __declspec(dllexport) void QBuffer_override_virtual_ReadLineData(void* self, intptr_t slot);
long long QBuffer_virtualbase_ReadLineData(void* self, char* data, long long maxlen);
extern __declspec(dllexport) void QBuffer_override_virtual_SkipData(void* self, intptr_t slot);
long long QBuffer_virtualbase_SkipData(void* self, long long maxSize);
extern __declspec(dllexport) void QBuffer_Delete(QBuffer* self, bool isSubclass);

} 
