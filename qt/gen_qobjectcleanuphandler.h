#pragma once
#ifndef MIQT_QT_GEN_QOBJECTCLEANUPHANDLER_H
#define MIQT_QT_GEN_QOBJECTCLEANUPHANDLER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QObjectCleanupHandler QObjectCleanupHandler;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QObjectCleanupHandler* QObjectCleanupHandler_new();
extern __declspec(dllexport) 
void QObjectCleanupHandler_virtbase(QObjectCleanupHandler* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QObjectCleanupHandler_MetaObject(const QObjectCleanupHandler* self);
extern __declspec(dllexport) 
void* QObjectCleanupHandler_Metacast(QObjectCleanupHandler* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QObjectCleanupHandler_Tr(const char* s);
extern __declspec(dllexport) 
QObject* QObjectCleanupHandler_Add(QObjectCleanupHandler* self, QObject* object);
extern __declspec(dllexport) 
void QObjectCleanupHandler_Remove(QObjectCleanupHandler* self, QObject* object);
extern __declspec(dllexport) 
bool QObjectCleanupHandler_IsEmpty(const QObjectCleanupHandler* self);
extern __declspec(dllexport) 
void QObjectCleanupHandler_Clear(QObjectCleanupHandler* self);
extern __declspec(dllexport) 
struct miqt_string QObjectCleanupHandler_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QObjectCleanupHandler_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QObjectCleanupHandler_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QObjectCleanupHandler_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QObjectCleanupHandler_override_virtual_Metacast(void* self, intptr_t slot);
void* QObjectCleanupHandler_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QObjectCleanupHandler_Delete(QObjectCleanupHandler* self, bool isSubclass);

}
