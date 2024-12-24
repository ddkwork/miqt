#pragma once
#ifndef MIQT_QT_GEN_QWINEVENTNOTIFIER_H
#define MIQT_QT_GEN_QWINEVENTNOTIFIER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QWinEventNotifier QWinEventNotifier;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QWinEventNotifier* QWinEventNotifier_new();
extern __declspec(dllexport) 
QWinEventNotifier* QWinEventNotifier_new2(HANDLE hEvent);
extern __declspec(dllexport) 
QWinEventNotifier* QWinEventNotifier_new3(QObject* parent);
extern __declspec(dllexport) 
QWinEventNotifier* QWinEventNotifier_new4(HANDLE hEvent, QObject* parent);
extern __declspec(dllexport) 
void QWinEventNotifier_virtbase(QWinEventNotifier* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QWinEventNotifier_MetaObject(const QWinEventNotifier* self);
extern __declspec(dllexport) 
void* QWinEventNotifier_Metacast(QWinEventNotifier* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QWinEventNotifier_Tr(const char* s);
extern __declspec(dllexport) 
void QWinEventNotifier_SetHandle(QWinEventNotifier* self, HANDLE hEvent);
extern __declspec(dllexport) 
HANDLE QWinEventNotifier_Handle(const QWinEventNotifier* self);
extern __declspec(dllexport) 
bool QWinEventNotifier_IsEnabled(const QWinEventNotifier* self);
extern __declspec(dllexport) 
void QWinEventNotifier_SetEnabled(QWinEventNotifier* self, bool enable);
extern __declspec(dllexport) 
bool QWinEventNotifier_Event(QWinEventNotifier* self, QEvent* e);
extern __declspec(dllexport) 
struct miqt_string QWinEventNotifier_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QWinEventNotifier_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QWinEventNotifier_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QWinEventNotifier_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QWinEventNotifier_override_virtual_Metacast(void* self, intptr_t slot);
void* QWinEventNotifier_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QWinEventNotifier_Delete(QWinEventNotifier* self, bool isSubclass);

}
