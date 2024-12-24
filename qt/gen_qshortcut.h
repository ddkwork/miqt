#pragma once
#ifndef MIQT_QT_GEN_QSHORTCUT_H
#define MIQT_QT_GEN_QSHORTCUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QKeySequence QKeySequence;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QShortcut QShortcut;
typedef struct QTimerEvent QTimerEvent;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QShortcut* QShortcut_new(QObject* parent);
extern __declspec(dllexport) QShortcut* QShortcut_new2(QKeySequence* key, QObject* parent);
extern __declspec(dllexport) QShortcut* QShortcut_new3(int key, QObject* parent);
extern __declspec(dllexport) QShortcut* QShortcut_new4(QKeySequence* key, QObject* parent, const char* member);
extern __declspec(dllexport) QShortcut* QShortcut_new5(QKeySequence* key, QObject* parent, const char* member, const char* ambiguousMember);
extern __declspec(dllexport) QShortcut* QShortcut_new6(QKeySequence* key, QObject* parent, const char* member, const char* ambiguousMember, int context);
extern __declspec(dllexport) QShortcut* QShortcut_new7(int key, QObject* parent, const char* member);
extern __declspec(dllexport) QShortcut* QShortcut_new8(int key, QObject* parent, const char* member, const char* ambiguousMember);
extern __declspec(dllexport) QShortcut* QShortcut_new9(int key, QObject* parent, const char* member, const char* ambiguousMember, int context);
extern __declspec(dllexport) void QShortcut_virtbase(QShortcut* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QShortcut_MetaObject(const QShortcut* self);
extern __declspec(dllexport) void* QShortcut_Metacast(QShortcut* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QShortcut_Tr(const char* s);
extern __declspec(dllexport) void QShortcut_SetKey(QShortcut* self, QKeySequence* key);
extern __declspec(dllexport) QKeySequence* QShortcut_Key(const QShortcut* self);
extern __declspec(dllexport) void QShortcut_SetKeys(QShortcut* self, int key);
extern __declspec(dllexport) void QShortcut_SetKeysWithKeys(QShortcut* self, struct miqt_array /* of QKeySequence* */  keys);
extern __declspec(dllexport) struct miqt_array /* of QKeySequence* */  QShortcut_Keys(const QShortcut* self);
extern __declspec(dllexport) void QShortcut_SetEnabled(QShortcut* self, bool enable);
extern __declspec(dllexport) bool QShortcut_IsEnabled(const QShortcut* self);
extern __declspec(dllexport) void QShortcut_SetContext(QShortcut* self, int context);
extern __declspec(dllexport) int QShortcut_Context(const QShortcut* self);
extern __declspec(dllexport) void QShortcut_SetAutoRepeat(QShortcut* self, bool on);
extern __declspec(dllexport) bool QShortcut_AutoRepeat(const QShortcut* self);
extern __declspec(dllexport) int QShortcut_Id(const QShortcut* self);
extern __declspec(dllexport) void QShortcut_SetWhatsThis(QShortcut* self, struct miqt_string text);
extern __declspec(dllexport) struct miqt_string QShortcut_WhatsThis(const QShortcut* self);
extern __declspec(dllexport) void QShortcut_Activated(QShortcut* self);
void QShortcut_connect_Activated(QShortcut* self, intptr_t slot);
extern __declspec(dllexport) void QShortcut_ActivatedAmbiguously(QShortcut* self);
void QShortcut_connect_ActivatedAmbiguously(QShortcut* self, intptr_t slot);
extern __declspec(dllexport) bool QShortcut_Event(QShortcut* self, QEvent* e);
extern __declspec(dllexport) struct miqt_string QShortcut_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QShortcut_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QShortcut_override_virtual_Event(void* self, intptr_t slot);
bool QShortcut_virtualbase_Event(void* self, QEvent* e);
extern __declspec(dllexport) void QShortcut_override_virtual_EventFilter(void* self, intptr_t slot);
bool QShortcut_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QShortcut_override_virtual_TimerEvent(void* self, intptr_t slot);
void QShortcut_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QShortcut_override_virtual_ChildEvent(void* self, intptr_t slot);
void QShortcut_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QShortcut_override_virtual_CustomEvent(void* self, intptr_t slot);
void QShortcut_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QShortcut_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QShortcut_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QShortcut_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QShortcut_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QShortcut_Delete(QShortcut* self, bool isSubclass);

} 
