#pragma once
#ifndef MIQT_QT_GEN_QOBJECT_H
#define MIQT_QT_GEN_QOBJECT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAnyStringView;
class QBindingStorage;
class QChildEvent;
class QEvent;
class QMetaMethod;
class QMetaObject;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QMetaObject__Connection)
typedef QMetaObject::Connection QMetaObject__Connection;
#else
class QMetaObject__Connection;
#endif
class QObject;
class QObjectData;
class QSignalBlocker;
class QThread;
class QTimerEvent;
class QVariant;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_Disambiguated_t)
typedef Qt::Disambiguated_t Disambiguated_t;
#else
class Disambiguated_t;
#endif
class _GUID;
class type_info;
#else
typedef struct QAnyStringView QAnyStringView;
typedef struct QBindingStorage QBindingStorage;
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QMetaObject__Connection QMetaObject__Connection;
typedef struct QObject QObject;
typedef struct QObjectData QObjectData;
typedef struct QSignalBlocker QSignalBlocker;
typedef struct QThread QThread;
typedef struct QTimerEvent QTimerEvent;
typedef struct QVariant QVariant;
typedef struct Disambiguated_t Disambiguated_t;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QMetaObject* QObjectData_DynamicMetaObject(const QObjectData* self);
extern __declspec(dllexport) void QObjectData_Delete(QObjectData* self, bool isSubclass);

extern __declspec(dllexport) QObject* QObject_new();
extern __declspec(dllexport) QObject* QObject_new2(QObject* parent);
extern __declspec(dllexport) QMetaObject* QObject_MetaObject(const QObject* self);
extern __declspec(dllexport) void* QObject_Metacast(QObject* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QObject_Tr(const char* s);
extern __declspec(dllexport) bool QObject_Event(QObject* self, QEvent* event);
extern __declspec(dllexport) bool QObject_EventFilter(QObject* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) struct miqt_string QObject_ObjectName(const QObject* self);
extern __declspec(dllexport) void QObject_SetObjectName(QObject* self, QAnyStringView* name);
extern __declspec(dllexport) bool QObject_IsWidgetType(const QObject* self);
extern __declspec(dllexport) bool QObject_IsWindowType(const QObject* self);
extern __declspec(dllexport) bool QObject_IsQuickItemType(const QObject* self);
extern __declspec(dllexport) bool QObject_SignalsBlocked(const QObject* self);
extern __declspec(dllexport) bool QObject_BlockSignals(QObject* self, bool b);
extern __declspec(dllexport) QThread* QObject_Thread(const QObject* self);
extern __declspec(dllexport) bool QObject_MoveToThread(QObject* self, QThread* thread);
extern __declspec(dllexport) int QObject_StartTimer(QObject* self, int interval);
extern __declspec(dllexport) void QObject_KillTimer(QObject* self, int id);
extern __declspec(dllexport) void QObject_KillTimerWithId(QObject* self, int id);
extern __declspec(dllexport) struct miqt_array /* of QObject* */  QObject_Children(const QObject* self);
extern __declspec(dllexport) void QObject_SetParent(QObject* self, QObject* parent);
extern __declspec(dllexport) void QObject_InstallEventFilter(QObject* self, QObject* filterObj);
extern __declspec(dllexport) void QObject_RemoveEventFilter(QObject* self, QObject* obj);
extern __declspec(dllexport) QMetaObject__Connection* QObject_Connect(QObject* sender, QMetaMethod* signal, QObject* receiver, QMetaMethod* method);
extern __declspec(dllexport) QMetaObject__Connection* QObject_Connect2(const QObject* self, QObject* sender, const char* signal, const char* member);
extern __declspec(dllexport) bool QObject_Disconnect(QObject* sender, QMetaMethod* signal, QObject* receiver, QMetaMethod* member);
extern __declspec(dllexport) bool QObject_DisconnectWithQMetaObjectConnection(QMetaObject__Connection* param1);
extern __declspec(dllexport) void QObject_DumpObjectTree(const QObject* self);
extern __declspec(dllexport) void QObject_DumpObjectInfo(const QObject* self);
extern __declspec(dllexport) bool QObject_SetProperty(QObject* self, const char* name, QVariant* value);
extern __declspec(dllexport) QVariant* QObject_Property(const QObject* self, const char* name);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QObject_DynamicPropertyNames(const QObject* self);
extern __declspec(dllexport) QBindingStorage* QObject_BindingStorage(QObject* self);
extern __declspec(dllexport) QBindingStorage* QObject_BindingStorage2(const QObject* self);
extern __declspec(dllexport) void QObject_Destroyed(QObject* self);
void QObject_connect_Destroyed(QObject* self, intptr_t slot);
extern __declspec(dllexport) QObject* QObject_Parent(const QObject* self);
extern __declspec(dllexport) bool QObject_Inherits(const QObject* self, const char* classname);
extern __declspec(dllexport) void QObject_DeleteLater(QObject* self);
extern __declspec(dllexport) void QObject_TimerEvent(QObject* self, QTimerEvent* event);
extern __declspec(dllexport) void QObject_ChildEvent(QObject* self, QChildEvent* event);
extern __declspec(dllexport) void QObject_CustomEvent(QObject* self, QEvent* event);
extern __declspec(dllexport) void QObject_ConnectNotify(QObject* self, QMetaMethod* signal);
extern __declspec(dllexport) void QObject_DisconnectNotify(QObject* self, QMetaMethod* signal);
extern __declspec(dllexport) struct miqt_string QObject_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QObject_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) bool QObject_MoveToThread2(QObject* self, QThread* thread, Disambiguated_t* param2);
extern __declspec(dllexport) int QObject_StartTimer2(QObject* self, int interval, int timerType);
extern __declspec(dllexport) QMetaObject__Connection* QObject_Connect5(QObject* sender, QMetaMethod* signal, QObject* receiver, QMetaMethod* method, int typeVal);
extern __declspec(dllexport) QMetaObject__Connection* QObject_Connect4(const QObject* self, QObject* sender, const char* signal, const char* member, int typeVal);
extern __declspec(dllexport) void QObject_Destroyed1(QObject* self, QObject* param1);
void QObject_connect_Destroyed1(QObject* self, intptr_t slot);
extern __declspec(dllexport) void QObject_override_virtual_Event(void* self, intptr_t slot);
bool QObject_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QObject_override_virtual_EventFilter(void* self, intptr_t slot);
bool QObject_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QObject_override_virtual_TimerEvent(void* self, intptr_t slot);
void QObject_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QObject_override_virtual_ChildEvent(void* self, intptr_t slot);
void QObject_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QObject_override_virtual_CustomEvent(void* self, intptr_t slot);
void QObject_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QObject_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QObject_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QObject_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QObject_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QObject_Delete(QObject* self, bool isSubclass);

extern __declspec(dllexport) QSignalBlocker* QSignalBlocker_new(QObject* o);
extern __declspec(dllexport) QSignalBlocker* QSignalBlocker_new2(QObject* o);
extern __declspec(dllexport) void QSignalBlocker_Reblock(QSignalBlocker* self);
extern __declspec(dllexport) void QSignalBlocker_Unblock(QSignalBlocker* self);
extern __declspec(dllexport) void QSignalBlocker_Dismiss(QSignalBlocker* self);
extern __declspec(dllexport) void QSignalBlocker_Delete(QSignalBlocker* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
