#pragma once
#ifndef MIQT_QT_GEN_QTIMER_H
#define MIQT_QT_GEN_QTIMER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimer QTimer;
typedef struct QTimerEvent QTimerEvent;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QTimer* QTimer_new();
extern __declspec(dllexport) 
QTimer* QTimer_new2(QObject* parent);
extern __declspec(dllexport) 
void QTimer_virtbase(QTimer* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QTimer_MetaObject(const QTimer* self);
extern __declspec(dllexport) 
void* QTimer_Metacast(QTimer* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QTimer_Tr(const char* s);
extern __declspec(dllexport) 
bool QTimer_IsActive(const QTimer* self);
extern __declspec(dllexport) 
int QTimer_TimerId(const QTimer* self);
extern __declspec(dllexport) 
int QTimer_Id(const QTimer* self);
extern __declspec(dllexport) 
void QTimer_SetInterval(QTimer* self, int msec);
extern __declspec(dllexport) 
int QTimer_Interval(const QTimer* self);
extern __declspec(dllexport) 
int QTimer_RemainingTime(const QTimer* self);
extern __declspec(dllexport) 
void QTimer_SetTimerType(QTimer* self, int atype);
extern __declspec(dllexport) 
int QTimer_TimerType(const QTimer* self);
extern __declspec(dllexport) 
void QTimer_SetSingleShot(QTimer* self, bool singleShot);
extern __declspec(dllexport) 
bool QTimer_IsSingleShot(const QTimer* self);
extern __declspec(dllexport) 
void QTimer_Start(QTimer* self, int msec);
extern __declspec(dllexport) 
void QTimer_Start2(QTimer* self);
extern __declspec(dllexport) 
void QTimer_Stop(QTimer* self);
extern __declspec(dllexport) 
void QTimer_TimerEvent(QTimer* self, QTimerEvent* param1);
extern __declspec(dllexport) 
struct miqt_string QTimer_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QTimer_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QTimer_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QTimer_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QTimer_override_virtual_Metacast(void* self, intptr_t slot);
void* QTimer_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QTimer_Delete(QTimer* self, bool isSubclass);

}
