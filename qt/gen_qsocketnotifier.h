#pragma once
#ifndef MIQT_QT_GEN_QSOCKETNOTIFIER_H
#define MIQT_QT_GEN_QSOCKETNOTIFIER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChildEvent;
class QEvent;
class QMetaMethod;
class QMetaObject;
class QObject;
class QSocketDescriptor;
class QSocketNotifier;
class QTimerEvent;
class _GUID;
class type_info;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSocketDescriptor QSocketDescriptor;
typedef struct QSocketNotifier QSocketNotifier;
typedef struct QTimerEvent QTimerEvent;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSocketNotifier* QSocketNotifier_new(Type param1);
extern __declspec(dllexport) QSocketNotifier* QSocketNotifier_new2(intptr_t socket, Type param2);
extern __declspec(dllexport) QSocketNotifier* QSocketNotifier_new3(Type param1, QObject* parent);
extern __declspec(dllexport) QSocketNotifier* QSocketNotifier_new4(intptr_t socket, Type param2, QObject* parent);
extern __declspec(dllexport) void QSocketNotifier_virtbase(QSocketNotifier* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QSocketNotifier_MetaObject(const QSocketNotifier* self);
extern __declspec(dllexport) void* QSocketNotifier_Metacast(QSocketNotifier* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QSocketNotifier_Tr(const char* s);
extern __declspec(dllexport) void QSocketNotifier_SetSocket(QSocketNotifier* self, intptr_t socket);
extern __declspec(dllexport) intptr_t QSocketNotifier_Socket(const QSocketNotifier* self);
extern __declspec(dllexport) Type QSocketNotifier_Type(const QSocketNotifier* self);
extern __declspec(dllexport) bool QSocketNotifier_IsValid(const QSocketNotifier* self);
extern __declspec(dllexport) bool QSocketNotifier_IsEnabled(const QSocketNotifier* self);
extern __declspec(dllexport) void QSocketNotifier_SetEnabled(QSocketNotifier* self, bool enabled);
extern __declspec(dllexport) bool QSocketNotifier_Event(QSocketNotifier* self, QEvent* param1);
extern __declspec(dllexport) struct miqt_string QSocketNotifier_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QSocketNotifier_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QSocketNotifier_override_virtual_Event(void* self, intptr_t slot);
bool QSocketNotifier_virtualbase_Event(void* self, QEvent* param1);
extern __declspec(dllexport) void QSocketNotifier_override_virtual_EventFilter(void* self, intptr_t slot);
bool QSocketNotifier_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QSocketNotifier_override_virtual_TimerEvent(void* self, intptr_t slot);
void QSocketNotifier_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QSocketNotifier_override_virtual_ChildEvent(void* self, intptr_t slot);
void QSocketNotifier_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QSocketNotifier_override_virtual_CustomEvent(void* self, intptr_t slot);
void QSocketNotifier_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QSocketNotifier_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QSocketNotifier_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QSocketNotifier_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QSocketNotifier_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QSocketNotifier_Delete(QSocketNotifier* self, bool isSubclass);

extern __declspec(dllexport) QSocketDescriptor* QSocketDescriptor_new();
extern __declspec(dllexport) QSocketDescriptor* QSocketDescriptor_new2(intptr_t desc);
extern __declspec(dllexport) QSocketDescriptor* QSocketDescriptor_new3(QSocketDescriptor* param1);
extern __declspec(dllexport) QSocketDescriptor* QSocketDescriptor_new4(DescriptorType descriptor);
extern __declspec(dllexport) void* QSocketDescriptor_WinHandle(const QSocketDescriptor* self);
extern __declspec(dllexport) bool QSocketDescriptor_IsValid(const QSocketDescriptor* self);
extern __declspec(dllexport) void QSocketDescriptor_Delete(QSocketDescriptor* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
