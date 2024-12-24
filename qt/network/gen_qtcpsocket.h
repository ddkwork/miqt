#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QTCPSOCKET_H
#define MIQT_QT_NETWORK_GEN_QTCPSOCKET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractSocket QAbstractSocket;
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTcpSocket QTcpSocket;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QTcpSocket* QTcpSocket_new();
extern __declspec(dllexport) 
QTcpSocket* QTcpSocket_new2(QObject* parent);
extern __declspec(dllexport) 
void QTcpSocket_virtbase(QTcpSocket* src
, QAbstractSocket** outptr_QAbstractSocket
);
extern __declspec(dllexport) 
QMetaObject* QTcpSocket_MetaObject(const QTcpSocket* self);
extern __declspec(dllexport) 
void* QTcpSocket_Metacast(QTcpSocket* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QTcpSocket_Tr(const char* s);
extern __declspec(dllexport) 
bool QTcpSocket_Bind(QTcpSocket* self, int addr);
extern __declspec(dllexport) 
struct miqt_string QTcpSocket_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QTcpSocket_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
bool QTcpSocket_Bind2(QTcpSocket* self, int addr, uint16_t port);
extern __declspec(dllexport) 
bool QTcpSocket_Bind3(QTcpSocket* self, int addr, uint16_t port, BindMode mode);
extern __declspec(dllexport) 
void QTcpSocket_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QTcpSocket_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QTcpSocket_override_virtual_Metacast(void* self, intptr_t slot);
void* QTcpSocket_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QTcpSocket_Delete(QTcpSocket* self, bool isSubclass);

}
