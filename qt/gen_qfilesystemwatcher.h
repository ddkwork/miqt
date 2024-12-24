#pragma once
#ifndef MIQT_QT_GEN_QFILESYSTEMWATCHER_H
#define MIQT_QT_GEN_QFILESYSTEMWATCHER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QFileSystemWatcher QFileSystemWatcher;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QFileSystemWatcher* QFileSystemWatcher_new();
extern __declspec(dllexport) QFileSystemWatcher* QFileSystemWatcher_new2(struct miqt_array /* of struct miqt_string */  paths);
extern __declspec(dllexport) QFileSystemWatcher* QFileSystemWatcher_new3(QObject* parent);
extern __declspec(dllexport) QFileSystemWatcher* QFileSystemWatcher_new4(struct miqt_array /* of struct miqt_string */  paths, QObject* parent);
extern __declspec(dllexport) void QFileSystemWatcher_virtbase(QFileSystemWatcher* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QFileSystemWatcher_MetaObject(const QFileSystemWatcher* self);
extern __declspec(dllexport) void* QFileSystemWatcher_Metacast(QFileSystemWatcher* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QFileSystemWatcher_Tr(const char* s);
extern __declspec(dllexport) bool QFileSystemWatcher_AddPath(QFileSystemWatcher* self, struct miqt_string file);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QFileSystemWatcher_AddPaths(QFileSystemWatcher* self, struct miqt_array /* of struct miqt_string */  files);
extern __declspec(dllexport) bool QFileSystemWatcher_RemovePath(QFileSystemWatcher* self, struct miqt_string file);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QFileSystemWatcher_RemovePaths(QFileSystemWatcher* self, struct miqt_array /* of struct miqt_string */  files);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QFileSystemWatcher_Files(const QFileSystemWatcher* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QFileSystemWatcher_Directories(const QFileSystemWatcher* self);
extern __declspec(dllexport) struct miqt_string QFileSystemWatcher_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QFileSystemWatcher_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QFileSystemWatcher_override_virtual_Event(void* self, intptr_t slot);
bool QFileSystemWatcher_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QFileSystemWatcher_override_virtual_EventFilter(void* self, intptr_t slot);
bool QFileSystemWatcher_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QFileSystemWatcher_override_virtual_TimerEvent(void* self, intptr_t slot);
void QFileSystemWatcher_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QFileSystemWatcher_override_virtual_ChildEvent(void* self, intptr_t slot);
void QFileSystemWatcher_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QFileSystemWatcher_override_virtual_CustomEvent(void* self, intptr_t slot);
void QFileSystemWatcher_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QFileSystemWatcher_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QFileSystemWatcher_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QFileSystemWatcher_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QFileSystemWatcher_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QFileSystemWatcher_Delete(QFileSystemWatcher* self, bool isSubclass);

} 
