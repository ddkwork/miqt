#pragma once
#ifndef MIQT_QT_GEN_QABSTRACTANIMATION_H
#define MIQT_QT_GEN_QABSTRACTANIMATION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractAnimation QAbstractAnimation;
typedef struct QAnimationDriver QAnimationDriver;
typedef struct QAnimationGroup QAnimationGroup;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAbstractAnimation* QAbstractAnimation_new();
extern __declspec(dllexport) 
QAbstractAnimation* QAbstractAnimation_new2(QObject* parent);
extern __declspec(dllexport) 
void QAbstractAnimation_virtbase(QAbstractAnimation* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QAbstractAnimation_MetaObject(const QAbstractAnimation* self);
extern __declspec(dllexport) 
void* QAbstractAnimation_Metacast(QAbstractAnimation* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAbstractAnimation_Tr(const char* s);
extern __declspec(dllexport) 
State QAbstractAnimation_State(const QAbstractAnimation* self);
extern __declspec(dllexport) 
QAnimationGroup* QAbstractAnimation_Group(const QAbstractAnimation* self);
extern __declspec(dllexport) 
Direction QAbstractAnimation_Direction(const QAbstractAnimation* self);
extern __declspec(dllexport) 
void QAbstractAnimation_SetDirection(QAbstractAnimation* self, Direction direction);
extern __declspec(dllexport) 
int QAbstractAnimation_CurrentTime(const QAbstractAnimation* self);
extern __declspec(dllexport) 
int QAbstractAnimation_CurrentLoopTime(const QAbstractAnimation* self);
extern __declspec(dllexport) 
int QAbstractAnimation_LoopCount(const QAbstractAnimation* self);
extern __declspec(dllexport) 
void QAbstractAnimation_SetLoopCount(QAbstractAnimation* self, int loopCount);
extern __declspec(dllexport) 
int QAbstractAnimation_CurrentLoop(const QAbstractAnimation* self);
extern __declspec(dllexport) 
int QAbstractAnimation_Duration(const QAbstractAnimation* self);
extern __declspec(dllexport) 
int QAbstractAnimation_TotalDuration(const QAbstractAnimation* self);
extern __declspec(dllexport) 
void QAbstractAnimation_Finished(QAbstractAnimation* self);
void QAbstractAnimation_connect_Finished(QAbstractAnimation* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractAnimation_StateChanged(QAbstractAnimation* self, int newState, int oldState);
void QAbstractAnimation_connect_StateChanged(QAbstractAnimation* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractAnimation_CurrentLoopChanged(QAbstractAnimation* self, int currentLoop);
void QAbstractAnimation_connect_CurrentLoopChanged(QAbstractAnimation* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractAnimation_DirectionChanged(QAbstractAnimation* self, int param1);
void QAbstractAnimation_connect_DirectionChanged(QAbstractAnimation* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractAnimation_Start(QAbstractAnimation* self);
extern __declspec(dllexport) 
void QAbstractAnimation_Pause(QAbstractAnimation* self);
extern __declspec(dllexport) 
void QAbstractAnimation_Resume(QAbstractAnimation* self);
extern __declspec(dllexport) 
void QAbstractAnimation_SetPaused(QAbstractAnimation* self, bool paused);
extern __declspec(dllexport) 
void QAbstractAnimation_Stop(QAbstractAnimation* self);
extern __declspec(dllexport) 
void QAbstractAnimation_SetCurrentTime(QAbstractAnimation* self, int msecs);
extern __declspec(dllexport) 
bool QAbstractAnimation_Event(QAbstractAnimation* self, QEvent* event);
extern __declspec(dllexport) 
void QAbstractAnimation_UpdateCurrentTime(QAbstractAnimation* self, int currentTime);
extern __declspec(dllexport) 
void QAbstractAnimation_UpdateState(QAbstractAnimation* self, int newState, int oldState);
extern __declspec(dllexport) 
void QAbstractAnimation_UpdateDirection(QAbstractAnimation* self, int direction);
extern __declspec(dllexport) 
struct miqt_string QAbstractAnimation_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAbstractAnimation_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAbstractAnimation_Start1(QAbstractAnimation* self, int policy);
extern __declspec(dllexport) 
void QAbstractAnimation_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAbstractAnimation_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAbstractAnimation_override_virtual_Metacast(void* self, intptr_t slot);
void* QAbstractAnimation_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAbstractAnimation_Delete(QAbstractAnimation* self, bool isSubclass);

extern __declspec(dllexport) 
QAnimationDriver* QAnimationDriver_new();
extern __declspec(dllexport) 
QAnimationDriver* QAnimationDriver_new2(QObject* parent);
extern __declspec(dllexport) 
void QAnimationDriver_virtbase(QAnimationDriver* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QAnimationDriver_MetaObject(const QAnimationDriver* self);
extern __declspec(dllexport) 
void* QAnimationDriver_Metacast(QAnimationDriver* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAnimationDriver_Tr(const char* s);
extern __declspec(dllexport) 
void QAnimationDriver_Advance(QAnimationDriver* self);
extern __declspec(dllexport) 
void QAnimationDriver_Install(QAnimationDriver* self);
extern __declspec(dllexport) 
void QAnimationDriver_Uninstall(QAnimationDriver* self);
extern __declspec(dllexport) 
bool QAnimationDriver_IsRunning(const QAnimationDriver* self);
extern __declspec(dllexport) 
long long QAnimationDriver_Elapsed(const QAnimationDriver* self);
extern __declspec(dllexport) 
void QAnimationDriver_Started(QAnimationDriver* self);
void QAnimationDriver_connect_Started(QAnimationDriver* self, intptr_t slot);
extern __declspec(dllexport) 
void QAnimationDriver_Stopped(QAnimationDriver* self);
void QAnimationDriver_connect_Stopped(QAnimationDriver* self, intptr_t slot);
extern __declspec(dllexport) 
void QAnimationDriver_Start(QAnimationDriver* self);
extern __declspec(dllexport) 
void QAnimationDriver_Stop(QAnimationDriver* self);
extern __declspec(dllexport) 
struct miqt_string QAnimationDriver_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAnimationDriver_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAnimationDriver_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAnimationDriver_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAnimationDriver_override_virtual_Metacast(void* self, intptr_t slot);
void* QAnimationDriver_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAnimationDriver_Delete(QAnimationDriver* self, bool isSubclass);

}
