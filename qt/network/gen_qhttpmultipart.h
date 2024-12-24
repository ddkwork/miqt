#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QHTTPMULTIPART_H
#define MIQT_QT_NETWORK_GEN_QHTTPMULTIPART_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QHttpMultiPart QHttpMultiPart;
typedef struct QHttpPart QHttpPart;
typedef struct QIODevice QIODevice;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;
typedef struct QVariant QVariant;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QHttpPart* QHttpPart_new();
extern __declspec(dllexport) QHttpPart* QHttpPart_new2(QHttpPart* other);
extern __declspec(dllexport) void QHttpPart_OperatorAssign(QHttpPart* self, QHttpPart* other);
extern __declspec(dllexport) void QHttpPart_Swap(QHttpPart* self, QHttpPart* other);
extern __declspec(dllexport) bool QHttpPart_OperatorEqual(const QHttpPart* self, QHttpPart* other);
extern __declspec(dllexport) bool QHttpPart_OperatorNotEqual(const QHttpPart* self, QHttpPart* other);
extern __declspec(dllexport) void QHttpPart_SetHeader(QHttpPart* self, int header, QVariant* value);
extern __declspec(dllexport) void QHttpPart_SetRawHeader(QHttpPart* self, struct miqt_string headerName, struct miqt_string headerValue);
extern __declspec(dllexport) void QHttpPart_SetBody(QHttpPart* self, struct miqt_string body);
extern __declspec(dllexport) void QHttpPart_SetBodyDevice(QHttpPart* self, QIODevice* device);
extern __declspec(dllexport) void QHttpPart_Delete(QHttpPart* self, bool isSubclass);

extern __declspec(dllexport) QHttpMultiPart* QHttpMultiPart_new();
extern __declspec(dllexport) QHttpMultiPart* QHttpMultiPart_new2(ContentType contentType);
extern __declspec(dllexport) QHttpMultiPart* QHttpMultiPart_new3(QObject* parent);
extern __declspec(dllexport) QHttpMultiPart* QHttpMultiPart_new4(ContentType contentType, QObject* parent);
extern __declspec(dllexport) void QHttpMultiPart_virtbase(QHttpMultiPart* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QHttpMultiPart_MetaObject(const QHttpMultiPart* self);
extern __declspec(dllexport) void* QHttpMultiPart_Metacast(QHttpMultiPart* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QHttpMultiPart_Tr(const char* s);
extern __declspec(dllexport) void QHttpMultiPart_Append(QHttpMultiPart* self, QHttpPart* httpPart);
extern __declspec(dllexport) void QHttpMultiPart_SetContentType(QHttpMultiPart* self, ContentType contentType);
extern __declspec(dllexport) struct miqt_string QHttpMultiPart_Boundary(const QHttpMultiPart* self);
extern __declspec(dllexport) void QHttpMultiPart_SetBoundary(QHttpMultiPart* self, struct miqt_string boundary);
extern __declspec(dllexport) struct miqt_string QHttpMultiPart_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QHttpMultiPart_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QHttpMultiPart_override_virtual_Event(void* self, intptr_t slot);
bool QHttpMultiPart_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QHttpMultiPart_override_virtual_EventFilter(void* self, intptr_t slot);
bool QHttpMultiPart_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QHttpMultiPart_override_virtual_TimerEvent(void* self, intptr_t slot);
void QHttpMultiPart_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QHttpMultiPart_override_virtual_ChildEvent(void* self, intptr_t slot);
void QHttpMultiPart_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QHttpMultiPart_override_virtual_CustomEvent(void* self, intptr_t slot);
void QHttpMultiPart_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QHttpMultiPart_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QHttpMultiPart_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QHttpMultiPart_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QHttpMultiPart_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QHttpMultiPart_Delete(QHttpMultiPart* self, bool isSubclass);

} 
