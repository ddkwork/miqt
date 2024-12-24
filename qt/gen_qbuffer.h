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

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QBuffer* QBuffer_new();
extern __declspec(dllexport) 
QBuffer* QBuffer_new2(QObject* parent);
extern __declspec(dllexport) 
void QBuffer_virtbase(QBuffer* src
, QIODevice** outptr_QIODevice
);
extern __declspec(dllexport) 
QMetaObject* QBuffer_MetaObject(const QBuffer* self);
extern __declspec(dllexport) 
void* QBuffer_Metacast(QBuffer* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QBuffer_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_string QBuffer_Buffer(QBuffer* self);
extern __declspec(dllexport) 
struct miqt_string QBuffer_Buffer2(const QBuffer* self);
extern __declspec(dllexport) 
void QBuffer_SetData(QBuffer* self, struct miqt_string data);
extern __declspec(dllexport) 
void QBuffer_SetData2(QBuffer* self, const char* data, ptrdiff_t lenVal);
extern __declspec(dllexport) 
struct miqt_string QBuffer_Data(const QBuffer* self);
extern __declspec(dllexport) 
bool QBuffer_Open(QBuffer* self, OpenMode openMode);
extern __declspec(dllexport) 
void QBuffer_Close(QBuffer* self);
extern __declspec(dllexport) 
long long QBuffer_Size(const QBuffer* self);
extern __declspec(dllexport) 
long long QBuffer_Pos(const QBuffer* self);
extern __declspec(dllexport) 
bool QBuffer_Seek(QBuffer* self, long long off);
extern __declspec(dllexport) 
bool QBuffer_AtEnd(const QBuffer* self);
extern __declspec(dllexport) 
bool QBuffer_CanReadLine(const QBuffer* self);
extern __declspec(dllexport) 
void QBuffer_ConnectNotify(QBuffer* self, QMetaMethod* param1);
extern __declspec(dllexport) 
void QBuffer_DisconnectNotify(QBuffer* self, QMetaMethod* param1);
extern __declspec(dllexport) 
long long QBuffer_ReadData(QBuffer* self, char* data, long long maxlen);
extern __declspec(dllexport) 
long long QBuffer_WriteData(QBuffer* self, const char* data, long long lenVal);
extern __declspec(dllexport) 
struct miqt_string QBuffer_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QBuffer_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QBuffer_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QBuffer_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QBuffer_override_virtual_Metacast(void* self, intptr_t slot);
void* QBuffer_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QBuffer_Delete(QBuffer* self, bool isSubclass);

}
