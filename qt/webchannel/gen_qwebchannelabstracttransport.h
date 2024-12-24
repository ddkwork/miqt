#pragma once
#ifndef MIQT_QT_WEBCHANNEL_GEN_QWEBCHANNELABSTRACTTRANSPORT_H
#define MIQT_QT_WEBCHANNEL_GEN_QWEBCHANNELABSTRACTTRANSPORT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QJsonObject QJsonObject;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QWebChannelAbstractTransport QWebChannelAbstractTransport;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QWebChannelAbstractTransport* QWebChannelAbstractTransport_new();
extern __declspec(dllexport) 
QWebChannelAbstractTransport* QWebChannelAbstractTransport_new2(QObject* parent);
extern __declspec(dllexport) 
void QWebChannelAbstractTransport_virtbase(QWebChannelAbstractTransport* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QWebChannelAbstractTransport_MetaObject(const QWebChannelAbstractTransport* self);
extern __declspec(dllexport) 
void* QWebChannelAbstractTransport_Metacast(QWebChannelAbstractTransport* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QWebChannelAbstractTransport_Tr(const char* s);
extern __declspec(dllexport) 
void QWebChannelAbstractTransport_SendMessage(QWebChannelAbstractTransport* self, QJsonObject* message);
extern __declspec(dllexport) 
void QWebChannelAbstractTransport_MessageReceived(QWebChannelAbstractTransport* self, QJsonObject* message, QWebChannelAbstractTransport* transport);
void QWebChannelAbstractTransport_connect_MessageReceived(QWebChannelAbstractTransport* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QWebChannelAbstractTransport_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QWebChannelAbstractTransport_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QWebChannelAbstractTransport_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QWebChannelAbstractTransport_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QWebChannelAbstractTransport_override_virtual_Metacast(void* self, intptr_t slot);
void* QWebChannelAbstractTransport_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QWebChannelAbstractTransport_Delete(QWebChannelAbstractTransport* self, bool isSubclass);

}
