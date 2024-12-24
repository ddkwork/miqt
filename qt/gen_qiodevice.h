#pragma once
#ifndef MIQT_QT_GEN_QIODEVICE_H
#define MIQT_QT_GEN_QIODEVICE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QByteArrayView QByteArrayView;
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QIODevice* QIODevice_new();
extern __declspec(dllexport) QIODevice* QIODevice_new2(QObject* parent);
extern __declspec(dllexport) void QIODevice_virtbase(QIODevice* src, QObject** outptr_QObject, QIODeviceBase** outptr_QIODeviceBase);
extern __declspec(dllexport) QMetaObject* QIODevice_MetaObject(const QIODevice* self);
extern __declspec(dllexport) void* QIODevice_Metacast(QIODevice* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QIODevice_Tr(const char* s);
extern __declspec(dllexport) int QIODevice_OpenMode(const QIODevice* self);
extern __declspec(dllexport) void QIODevice_SetTextModeEnabled(QIODevice* self, bool enabled);
extern __declspec(dllexport) bool QIODevice_IsTextModeEnabled(const QIODevice* self);
extern __declspec(dllexport) bool QIODevice_IsOpen(const QIODevice* self);
extern __declspec(dllexport) bool QIODevice_IsReadable(const QIODevice* self);
extern __declspec(dllexport) bool QIODevice_IsWritable(const QIODevice* self);
extern __declspec(dllexport) bool QIODevice_IsSequential(const QIODevice* self);
extern __declspec(dllexport) int QIODevice_ReadChannelCount(const QIODevice* self);
extern __declspec(dllexport) int QIODevice_WriteChannelCount(const QIODevice* self);
extern __declspec(dllexport) int QIODevice_CurrentReadChannel(const QIODevice* self);
extern __declspec(dllexport) void QIODevice_SetCurrentReadChannel(QIODevice* self, int channel);
extern __declspec(dllexport) int QIODevice_CurrentWriteChannel(const QIODevice* self);
extern __declspec(dllexport) void QIODevice_SetCurrentWriteChannel(QIODevice* self, int channel);
extern __declspec(dllexport) bool QIODevice_Open(QIODevice* self, int mode);
extern __declspec(dllexport) void QIODevice_Close(QIODevice* self);
extern __declspec(dllexport) long long QIODevice_Pos(const QIODevice* self);
extern __declspec(dllexport) long long QIODevice_Size(const QIODevice* self);
extern __declspec(dllexport) bool QIODevice_Seek(QIODevice* self, long long pos);
extern __declspec(dllexport) bool QIODevice_AtEnd(const QIODevice* self);
extern __declspec(dllexport) bool QIODevice_Reset(QIODevice* self);
extern __declspec(dllexport) long long QIODevice_BytesAvailable(const QIODevice* self);
extern __declspec(dllexport) long long QIODevice_BytesToWrite(const QIODevice* self);
extern __declspec(dllexport) long long QIODevice_Read(QIODevice* self, char* data, long long maxlen);
extern __declspec(dllexport) struct miqt_string QIODevice_ReadWithMaxlen(QIODevice* self, long long maxlen);
extern __declspec(dllexport) struct miqt_string QIODevice_ReadAll(QIODevice* self);
extern __declspec(dllexport) long long QIODevice_ReadLine(QIODevice* self, char* data, long long maxlen);
extern __declspec(dllexport) struct miqt_string QIODevice_ReadLine2(QIODevice* self);
extern __declspec(dllexport) QByteArrayView* QIODevice_ReadLineIntoWithBuffer(QIODevice* self, QSpan<char> buffer);
extern __declspec(dllexport) QByteArrayView* QIODevice_ReadLineInto2(QIODevice* self, QSpan<uchar> buffer);
extern __declspec(dllexport) QByteArrayView* QIODevice_ReadLineInto3(QIODevice* self, QSpan<std::byte> buffer);
extern __declspec(dllexport) bool QIODevice_CanReadLine(const QIODevice* self);
extern __declspec(dllexport) void QIODevice_StartTransaction(QIODevice* self);
extern __declspec(dllexport) void QIODevice_CommitTransaction(QIODevice* self);
extern __declspec(dllexport) void QIODevice_RollbackTransaction(QIODevice* self);
extern __declspec(dllexport) bool QIODevice_IsTransactionStarted(const QIODevice* self);
extern __declspec(dllexport) long long QIODevice_Write(QIODevice* self, const char* data, long long lenVal);
extern __declspec(dllexport) long long QIODevice_WriteWithData(QIODevice* self, const char* data);
extern __declspec(dllexport) long long QIODevice_Write2(QIODevice* self, struct miqt_string data);
extern __declspec(dllexport) long long QIODevice_Peek(QIODevice* self, char* data, long long maxlen);
extern __declspec(dllexport) struct miqt_string QIODevice_PeekWithMaxlen(QIODevice* self, long long maxlen);
extern __declspec(dllexport) long long QIODevice_Skip(QIODevice* self, long long maxSize);
extern __declspec(dllexport) bool QIODevice_WaitForReadyRead(QIODevice* self, int msecs);
extern __declspec(dllexport) bool QIODevice_WaitForBytesWritten(QIODevice* self, int msecs);
extern __declspec(dllexport) void QIODevice_UngetChar(QIODevice* self, char c);
extern __declspec(dllexport) bool QIODevice_PutChar(QIODevice* self, char c);
extern __declspec(dllexport) bool QIODevice_GetChar(QIODevice* self, char* c);
extern __declspec(dllexport) struct miqt_string QIODevice_ErrorString(const QIODevice* self);
extern __declspec(dllexport) void QIODevice_ReadyRead(QIODevice* self);
void QIODevice_connect_ReadyRead(QIODevice* self, intptr_t slot);
extern __declspec(dllexport) void QIODevice_ChannelReadyRead(QIODevice* self, int channel);
void QIODevice_connect_ChannelReadyRead(QIODevice* self, intptr_t slot);
extern __declspec(dllexport) void QIODevice_BytesWritten(QIODevice* self, long long bytes);
void QIODevice_connect_BytesWritten(QIODevice* self, intptr_t slot);
extern __declspec(dllexport) void QIODevice_ChannelBytesWritten(QIODevice* self, int channel, long long bytes);
void QIODevice_connect_ChannelBytesWritten(QIODevice* self, intptr_t slot);
extern __declspec(dllexport) void QIODevice_AboutToClose(QIODevice* self);
void QIODevice_connect_AboutToClose(QIODevice* self, intptr_t slot);
extern __declspec(dllexport) void QIODevice_ReadChannelFinished(QIODevice* self);
void QIODevice_connect_ReadChannelFinished(QIODevice* self, intptr_t slot);
extern __declspec(dllexport) long long QIODevice_ReadData(QIODevice* self, char* data, long long maxlen);
extern __declspec(dllexport) long long QIODevice_ReadLineData(QIODevice* self, char* data, long long maxlen);
extern __declspec(dllexport) long long QIODevice_SkipData(QIODevice* self, long long maxSize);
extern __declspec(dllexport) long long QIODevice_WriteData(QIODevice* self, const char* data, long long lenVal);
extern __declspec(dllexport) struct miqt_string QIODevice_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QIODevice_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) struct miqt_string QIODevice_ReadLine1(QIODevice* self, long long maxlen);
extern __declspec(dllexport) void QIODevice_override_virtual_IsSequential(void* self, intptr_t slot);
bool QIODevice_virtualbase_IsSequential(const void* self);
extern __declspec(dllexport) void QIODevice_override_virtual_Open(void* self, intptr_t slot);
bool QIODevice_virtualbase_Open(void* self, int mode);
extern __declspec(dllexport) void QIODevice_override_virtual_Close(void* self, intptr_t slot);
void QIODevice_virtualbase_Close(void* self);
extern __declspec(dllexport) void QIODevice_override_virtual_Pos(void* self, intptr_t slot);
long long QIODevice_virtualbase_Pos(const void* self);
extern __declspec(dllexport) void QIODevice_override_virtual_Size(void* self, intptr_t slot);
long long QIODevice_virtualbase_Size(const void* self);
extern __declspec(dllexport) void QIODevice_override_virtual_Seek(void* self, intptr_t slot);
bool QIODevice_virtualbase_Seek(void* self, long long pos);
extern __declspec(dllexport) void QIODevice_override_virtual_AtEnd(void* self, intptr_t slot);
bool QIODevice_virtualbase_AtEnd(const void* self);
extern __declspec(dllexport) void QIODevice_override_virtual_Reset(void* self, intptr_t slot);
bool QIODevice_virtualbase_Reset(void* self);
extern __declspec(dllexport) void QIODevice_override_virtual_BytesAvailable(void* self, intptr_t slot);
long long QIODevice_virtualbase_BytesAvailable(const void* self);
extern __declspec(dllexport) void QIODevice_override_virtual_BytesToWrite(void* self, intptr_t slot);
long long QIODevice_virtualbase_BytesToWrite(const void* self);
extern __declspec(dllexport) void QIODevice_override_virtual_CanReadLine(void* self, intptr_t slot);
bool QIODevice_virtualbase_CanReadLine(const void* self);
extern __declspec(dllexport) void QIODevice_override_virtual_WaitForReadyRead(void* self, intptr_t slot);
bool QIODevice_virtualbase_WaitForReadyRead(void* self, int msecs);
extern __declspec(dllexport) void QIODevice_override_virtual_WaitForBytesWritten(void* self, intptr_t slot);
bool QIODevice_virtualbase_WaitForBytesWritten(void* self, int msecs);
extern __declspec(dllexport) void QIODevice_override_virtual_ReadData(void* self, intptr_t slot);
long long QIODevice_virtualbase_ReadData(void* self, char* data, long long maxlen);
extern __declspec(dllexport) void QIODevice_override_virtual_ReadLineData(void* self, intptr_t slot);
long long QIODevice_virtualbase_ReadLineData(void* self, char* data, long long maxlen);
extern __declspec(dllexport) void QIODevice_override_virtual_SkipData(void* self, intptr_t slot);
long long QIODevice_virtualbase_SkipData(void* self, long long maxSize);
extern __declspec(dllexport) void QIODevice_override_virtual_WriteData(void* self, intptr_t slot);
long long QIODevice_virtualbase_WriteData(void* self, const char* data, long long lenVal);
extern __declspec(dllexport) void QIODevice_override_virtual_Event(void* self, intptr_t slot);
bool QIODevice_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QIODevice_override_virtual_EventFilter(void* self, intptr_t slot);
bool QIODevice_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QIODevice_override_virtual_TimerEvent(void* self, intptr_t slot);
void QIODevice_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QIODevice_override_virtual_ChildEvent(void* self, intptr_t slot);
void QIODevice_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QIODevice_override_virtual_CustomEvent(void* self, intptr_t slot);
void QIODevice_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QIODevice_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QIODevice_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QIODevice_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QIODevice_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QIODevice_Delete(QIODevice* self, bool isSubclass);

} 
