#pragma once
#ifndef MIQT_QT_GEN_QLIBRARY_H
#define MIQT_QT_GEN_QLIBRARY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QLibrary QLibrary;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QLibrary* QLibrary_new();
extern __declspec(dllexport) QLibrary* QLibrary_new2(struct miqt_string fileName);
extern __declspec(dllexport) QLibrary* QLibrary_new3(struct miqt_string fileName, int verNum);
extern __declspec(dllexport) QLibrary* QLibrary_new4(struct miqt_string fileName, struct miqt_string version);
extern __declspec(dllexport) QLibrary* QLibrary_new5(QObject* parent);
extern __declspec(dllexport) QLibrary* QLibrary_new6(struct miqt_string fileName, QObject* parent);
extern __declspec(dllexport) QLibrary* QLibrary_new7(struct miqt_string fileName, int verNum, QObject* parent);
extern __declspec(dllexport) QLibrary* QLibrary_new8(struct miqt_string fileName, struct miqt_string version, QObject* parent);
extern __declspec(dllexport) void QLibrary_virtbase(QLibrary* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QLibrary_MetaObject(const QLibrary* self);
extern __declspec(dllexport) void* QLibrary_Metacast(QLibrary* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QLibrary_Tr(const char* s);
extern __declspec(dllexport) bool QLibrary_Load(QLibrary* self);
extern __declspec(dllexport) bool QLibrary_Unload(QLibrary* self);
extern __declspec(dllexport) bool QLibrary_IsLoaded(const QLibrary* self);
extern __declspec(dllexport) bool QLibrary_IsLibrary(struct miqt_string fileName);
extern __declspec(dllexport) void QLibrary_SetFileName(QLibrary* self, struct miqt_string fileName);
extern __declspec(dllexport) struct miqt_string QLibrary_FileName(const QLibrary* self);
extern __declspec(dllexport) void QLibrary_SetFileNameAndVersion(QLibrary* self, struct miqt_string fileName, int verNum);
extern __declspec(dllexport) void QLibrary_SetFileNameAndVersion2(QLibrary* self, struct miqt_string fileName, struct miqt_string version);
extern __declspec(dllexport) struct miqt_string QLibrary_ErrorString(const QLibrary* self);
extern __declspec(dllexport) void QLibrary_SetLoadHints(QLibrary* self, LoadHints hints);
extern __declspec(dllexport) LoadHints QLibrary_LoadHints(const QLibrary* self);
extern __declspec(dllexport) struct miqt_string QLibrary_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QLibrary_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QLibrary_override_virtual_Event(void* self, intptr_t slot);
bool QLibrary_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QLibrary_override_virtual_EventFilter(void* self, intptr_t slot);
bool QLibrary_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QLibrary_override_virtual_TimerEvent(void* self, intptr_t slot);
void QLibrary_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QLibrary_override_virtual_ChildEvent(void* self, intptr_t slot);
void QLibrary_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QLibrary_override_virtual_CustomEvent(void* self, intptr_t slot);
void QLibrary_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QLibrary_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QLibrary_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QLibrary_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QLibrary_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QLibrary_Delete(QLibrary* self, bool isSubclass);

} 
