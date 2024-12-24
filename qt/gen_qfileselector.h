#pragma once
#ifndef MIQT_QT_GEN_QFILESELECTOR_H
#define MIQT_QT_GEN_QFILESELECTOR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QFileSelector QFileSelector;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;
typedef struct QUrl QUrl;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QFileSelector* QFileSelector_new();
extern __declspec(dllexport) QFileSelector* QFileSelector_new2(QObject* parent);
extern __declspec(dllexport) void QFileSelector_virtbase(QFileSelector* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QFileSelector_MetaObject(const QFileSelector* self);
extern __declspec(dllexport) void* QFileSelector_Metacast(QFileSelector* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QFileSelector_Tr(const char* s);
extern __declspec(dllexport) struct miqt_string QFileSelector_Select(const QFileSelector* self, struct miqt_string filePath);
extern __declspec(dllexport) QUrl* QFileSelector_SelectWithFilePath(const QFileSelector* self, QUrl* filePath);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QFileSelector_ExtraSelectors(const QFileSelector* self);
extern __declspec(dllexport) void QFileSelector_SetExtraSelectors(QFileSelector* self, struct miqt_array /* of struct miqt_string */  list);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QFileSelector_AllSelectors(const QFileSelector* self);
extern __declspec(dllexport) struct miqt_string QFileSelector_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QFileSelector_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QFileSelector_override_virtual_Event(void* self, intptr_t slot);
bool QFileSelector_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QFileSelector_override_virtual_EventFilter(void* self, intptr_t slot);
bool QFileSelector_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QFileSelector_override_virtual_TimerEvent(void* self, intptr_t slot);
void QFileSelector_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QFileSelector_override_virtual_ChildEvent(void* self, intptr_t slot);
void QFileSelector_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QFileSelector_override_virtual_CustomEvent(void* self, intptr_t slot);
void QFileSelector_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QFileSelector_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QFileSelector_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QFileSelector_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QFileSelector_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QFileSelector_Delete(QFileSelector* self, bool isSubclass);

} 
