#pragma once
#ifndef MIQT_QT_GEN_QTIMELINE_H
#define MIQT_QT_GEN_QTIMELINE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEasingCurve QEasingCurve;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimeLine QTimeLine;
typedef struct QTimerEvent QTimerEvent;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QTimeLine* QTimeLine_new();
extern __declspec(dllexport) 
QTimeLine* QTimeLine_new2(int duration);
extern __declspec(dllexport) 
QTimeLine* QTimeLine_new3(int duration, QObject* parent);
extern __declspec(dllexport) 
void QTimeLine_virtbase(QTimeLine* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QTimeLine_MetaObject(const QTimeLine* self);
extern __declspec(dllexport) 
void* QTimeLine_Metacast(QTimeLine* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QTimeLine_Tr(const char* s);
extern __declspec(dllexport) 
State QTimeLine_State(const QTimeLine* self);
extern __declspec(dllexport) 
int QTimeLine_LoopCount(const QTimeLine* self);
extern __declspec(dllexport) 
void QTimeLine_SetLoopCount(QTimeLine* self, int count);
extern __declspec(dllexport) 
Direction QTimeLine_Direction(const QTimeLine* self);
extern __declspec(dllexport) 
void QTimeLine_SetDirection(QTimeLine* self, Direction direction);
extern __declspec(dllexport) 
int QTimeLine_Duration(const QTimeLine* self);
extern __declspec(dllexport) 
void QTimeLine_SetDuration(QTimeLine* self, int duration);
extern __declspec(dllexport) 
int QTimeLine_StartFrame(const QTimeLine* self);
extern __declspec(dllexport) 
void QTimeLine_SetStartFrame(QTimeLine* self, int frame);
extern __declspec(dllexport) 
int QTimeLine_EndFrame(const QTimeLine* self);
extern __declspec(dllexport) 
void QTimeLine_SetEndFrame(QTimeLine* self, int frame);
extern __declspec(dllexport) 
void QTimeLine_SetFrameRange(QTimeLine* self, int startFrame, int endFrame);
extern __declspec(dllexport) 
int QTimeLine_UpdateInterval(const QTimeLine* self);
extern __declspec(dllexport) 
void QTimeLine_SetUpdateInterval(QTimeLine* self, int interval);
extern __declspec(dllexport) 
QEasingCurve* QTimeLine_EasingCurve(const QTimeLine* self);
extern __declspec(dllexport) 
void QTimeLine_SetEasingCurve(QTimeLine* self, QEasingCurve* curve);
extern __declspec(dllexport) 
int QTimeLine_CurrentTime(const QTimeLine* self);
extern __declspec(dllexport) 
int QTimeLine_CurrentFrame(const QTimeLine* self);
extern __declspec(dllexport) 
double QTimeLine_CurrentValue(const QTimeLine* self);
extern __declspec(dllexport) 
int QTimeLine_FrameForTime(const QTimeLine* self, int msec);
extern __declspec(dllexport) 
double QTimeLine_ValueForTime(const QTimeLine* self, int msec);
extern __declspec(dllexport) 
void QTimeLine_Start(QTimeLine* self);
extern __declspec(dllexport) 
void QTimeLine_Resume(QTimeLine* self);
extern __declspec(dllexport) 
void QTimeLine_Stop(QTimeLine* self);
extern __declspec(dllexport) 
void QTimeLine_SetPaused(QTimeLine* self, bool paused);
extern __declspec(dllexport) 
void QTimeLine_SetCurrentTime(QTimeLine* self, int msec);
extern __declspec(dllexport) 
void QTimeLine_ToggleDirection(QTimeLine* self);
extern __declspec(dllexport) 
void QTimeLine_TimerEvent(QTimeLine* self, QTimerEvent* event);
extern __declspec(dllexport) 
struct miqt_string QTimeLine_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QTimeLine_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QTimeLine_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QTimeLine_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QTimeLine_override_virtual_Metacast(void* self, intptr_t slot);
void* QTimeLine_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QTimeLine_Delete(QTimeLine* self, bool isSubclass);

}
