#pragma once
#ifndef MIQT_QT_GEN_QCHRONOTIMER_H
#define MIQT_QT_GEN_QCHRONOTIMER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChronoTimer QChronoTimer;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QChronoTimer* QChronoTimer_new();
extern __declspec(dllexport) 
QChronoTimer* QChronoTimer_new2(QObject* parent);
extern __declspec(dllexport) 
void QChronoTimer_virtbase(QChronoTimer* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QChronoTimer_MetaObject(const QChronoTimer* self);
extern __declspec(dllexport) 
void* QChronoTimer_Metacast(QChronoTimer* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QChronoTimer_Tr(const char* s);
extern __declspec(dllexport) 
bool QChronoTimer_IsActive(const QChronoTimer* self);
extern __declspec(dllexport) 
int QChronoTimer_Id(const QChronoTimer* self);
extern __declspec(dllexport) 
void QChronoTimer_SetTimerType(QChronoTimer* self, int atype);
extern __declspec(dllexport) 
int QChronoTimer_TimerType(const QChronoTimer* self);
extern __declspec(dllexport) 
void QChronoTimer_SetSingleShot(QChronoTimer* self, bool singleShot);
extern __declspec(dllexport) 
bool QChronoTimer_IsSingleShot(const QChronoTimer* self);
extern __declspec(dllexport) 
void QChronoTimer_Start(QChronoTimer* self);
extern __declspec(dllexport) 
void QChronoTimer_Stop(QChronoTimer* self);
extern __declspec(dllexport) 
void QChronoTimer_TimerEvent(QChronoTimer* self, QTimerEvent* param1);
extern __declspec(dllexport) 
struct miqt_string QChronoTimer_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QChronoTimer_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QChronoTimer_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QChronoTimer_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QChronoTimer_override_virtual_Metacast(void* self, intptr_t slot);
void* QChronoTimer_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QChronoTimer_Delete(QChronoTimer* self, bool isSubclass);

}
