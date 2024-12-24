#pragma once
#ifndef MIQT_QT_GEN_QEVENTLOOP_H
#define MIQT_QT_GEN_QEVENTLOOP_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDeadlineTimer QDeadlineTimer;
typedef struct QEvent QEvent;
typedef struct QEventLoop QEventLoop;
typedef struct QEventLoopLocker QEventLoopLocker;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QThread QThread;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QEventLoop* QEventLoop_new();
extern __declspec(dllexport) 
QEventLoop* QEventLoop_new2(QObject* parent);
extern __declspec(dllexport) 
void QEventLoop_virtbase(QEventLoop* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QEventLoop_MetaObject(const QEventLoop* self);
extern __declspec(dllexport) 
void* QEventLoop_Metacast(QEventLoop* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QEventLoop_Tr(const char* s);
extern __declspec(dllexport) 
bool QEventLoop_ProcessEvents(QEventLoop* self);
extern __declspec(dllexport) 
void QEventLoop_ProcessEvents2(QEventLoop* self, ProcessEventsFlags flags, int maximumTime);
extern __declspec(dllexport) 
void QEventLoop_ProcessEvents3(QEventLoop* self, ProcessEventsFlags flags, QDeadlineTimer* deadline);
extern __declspec(dllexport) 
int QEventLoop_Exec(QEventLoop* self);
extern __declspec(dllexport) 
bool QEventLoop_IsRunning(const QEventLoop* self);
extern __declspec(dllexport) 
void QEventLoop_WakeUp(QEventLoop* self);
extern __declspec(dllexport) 
bool QEventLoop_Event(QEventLoop* self, QEvent* event);
extern __declspec(dllexport) 
void QEventLoop_Exit(QEventLoop* self);
extern __declspec(dllexport) 
void QEventLoop_Quit(QEventLoop* self);
extern __declspec(dllexport) 
struct miqt_string QEventLoop_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QEventLoop_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
bool QEventLoop_ProcessEvents1(QEventLoop* self, ProcessEventsFlags flags);
extern __declspec(dllexport) 
int QEventLoop_Exec1(QEventLoop* self, ProcessEventsFlags flags);
extern __declspec(dllexport) 
void QEventLoop_Exit1(QEventLoop* self, int returnCode);
extern __declspec(dllexport) 
void QEventLoop_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QEventLoop_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QEventLoop_override_virtual_Metacast(void* self, intptr_t slot);
void* QEventLoop_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QEventLoop_Delete(QEventLoop* self, bool isSubclass);

extern __declspec(dllexport) 
QEventLoopLocker* QEventLoopLocker_new();
extern __declspec(dllexport) 
QEventLoopLocker* QEventLoopLocker_new2(QEventLoop* loop);
extern __declspec(dllexport) 
QEventLoopLocker* QEventLoopLocker_new3(QThread* thread);
extern __declspec(dllexport) 
void QEventLoopLocker_Swap(QEventLoopLocker* self, QEventLoopLocker* other);
extern __declspec(dllexport) 
void QEventLoopLocker_Delete(QEventLoopLocker* self, bool isSubclass);

}
