#pragma once
#ifndef MIQT_QT_WEBCHANNEL_GEN_QWEBCHANNEL_H
#define MIQT_QT_WEBCHANNEL_GEN_QWEBCHANNEL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QWebChannel QWebChannel;
typedef struct QWebChannelAbstractTransport QWebChannelAbstractTransport;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QWebChannel* QWebChannel_new();
extern __declspec(dllexport) 
QWebChannel* QWebChannel_new2(QObject* parent);
extern __declspec(dllexport) 
void QWebChannel_virtbase(QWebChannel* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QWebChannel_MetaObject(const QWebChannel* self);
extern __declspec(dllexport) 
void* QWebChannel_Metacast(QWebChannel* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QWebChannel_Tr(const char* s);
extern __declspec(dllexport) 
void QWebChannel_RegisterObjects(QWebChannel* self, struct miqt_map /* of struct miqt_string to QObject* */  objects);
extern __declspec(dllexport) 
struct miqt_map /* of struct miqt_string to QObject* */  QWebChannel_RegisteredObjects(const QWebChannel* self);
extern __declspec(dllexport) 
void QWebChannel_RegisterObject(QWebChannel* self, struct miqt_string id, QObject* object);
extern __declspec(dllexport) 
void QWebChannel_DeregisterObject(QWebChannel* self, QObject* object);
extern __declspec(dllexport) 
bool QWebChannel_BlockUpdates(const QWebChannel* self);
extern __declspec(dllexport) 
void QWebChannel_SetBlockUpdates(QWebChannel* self, bool block);
extern __declspec(dllexport) 
int QWebChannel_PropertyUpdateInterval(const QWebChannel* self);
extern __declspec(dllexport) 
void QWebChannel_SetPropertyUpdateInterval(QWebChannel* self, int ms);
extern __declspec(dllexport) 
void QWebChannel_BlockUpdatesChanged(QWebChannel* self, bool block);
void QWebChannel_connect_BlockUpdatesChanged(QWebChannel* self, intptr_t slot);
extern __declspec(dllexport) 
void QWebChannel_ConnectTo(QWebChannel* self, QWebChannelAbstractTransport* transport);
extern __declspec(dllexport) 
void QWebChannel_DisconnectFrom(QWebChannel* self, QWebChannelAbstractTransport* transport);
extern __declspec(dllexport) 
struct miqt_string QWebChannel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QWebChannel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QWebChannel_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QWebChannel_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QWebChannel_override_virtual_Metacast(void* self, intptr_t slot);
void* QWebChannel_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QWebChannel_Delete(QWebChannel* self, bool isSubclass);

}
