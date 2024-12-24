#pragma once
#ifndef MIQT_QT_GEN_QMOVIE_H
#define MIQT_QT_GEN_QMOVIE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QColor QColor;
typedef struct QEvent QEvent;
typedef struct QIODevice QIODevice;
typedef struct QImage QImage;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QMovie QMovie;
typedef struct QObject QObject;
typedef struct QPixmap QPixmap;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QTimerEvent QTimerEvent;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QMovie* QMovie_new();
extern __declspec(dllexport) QMovie* QMovie_new2(QIODevice* device);
extern __declspec(dllexport) QMovie* QMovie_new3(struct miqt_string fileName);
extern __declspec(dllexport) QMovie* QMovie_new4(QObject* parent);
extern __declspec(dllexport) QMovie* QMovie_new5(QIODevice* device, struct miqt_string format);
extern __declspec(dllexport) QMovie* QMovie_new6(QIODevice* device, struct miqt_string format, QObject* parent);
extern __declspec(dllexport) QMovie* QMovie_new7(struct miqt_string fileName, struct miqt_string format);
extern __declspec(dllexport) QMovie* QMovie_new8(struct miqt_string fileName, struct miqt_string format, QObject* parent);
extern __declspec(dllexport) void QMovie_virtbase(QMovie* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QMovie_MetaObject(const QMovie* self);
extern __declspec(dllexport) void* QMovie_Metacast(QMovie* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QMovie_Tr(const char* s);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QMovie_SupportedFormats();
extern __declspec(dllexport) void QMovie_SetDevice(QMovie* self, QIODevice* device);
extern __declspec(dllexport) QIODevice* QMovie_Device(const QMovie* self);
extern __declspec(dllexport) void QMovie_SetFileName(QMovie* self, struct miqt_string fileName);
extern __declspec(dllexport) struct miqt_string QMovie_FileName(const QMovie* self);
extern __declspec(dllexport) void QMovie_SetFormat(QMovie* self, struct miqt_string format);
extern __declspec(dllexport) struct miqt_string QMovie_Format(const QMovie* self);
extern __declspec(dllexport) void QMovie_SetBackgroundColor(QMovie* self, QColor* color);
extern __declspec(dllexport) QColor* QMovie_BackgroundColor(const QMovie* self);
extern __declspec(dllexport) MovieState QMovie_State(const QMovie* self);
extern __declspec(dllexport) QRect* QMovie_FrameRect(const QMovie* self);
extern __declspec(dllexport) QImage* QMovie_CurrentImage(const QMovie* self);
extern __declspec(dllexport) QPixmap* QMovie_CurrentPixmap(const QMovie* self);
extern __declspec(dllexport) bool QMovie_IsValid(const QMovie* self);
extern __declspec(dllexport) int QMovie_LastError(const QMovie* self);
extern __declspec(dllexport) struct miqt_string QMovie_LastErrorString(const QMovie* self);
extern __declspec(dllexport) bool QMovie_JumpToFrame(QMovie* self, int frameNumber);
extern __declspec(dllexport) int QMovie_LoopCount(const QMovie* self);
extern __declspec(dllexport) int QMovie_FrameCount(const QMovie* self);
extern __declspec(dllexport) int QMovie_NextFrameDelay(const QMovie* self);
extern __declspec(dllexport) int QMovie_CurrentFrameNumber(const QMovie* self);
extern __declspec(dllexport) int QMovie_Speed(const QMovie* self);
extern __declspec(dllexport) QSize* QMovie_ScaledSize(QMovie* self);
extern __declspec(dllexport) void QMovie_SetScaledSize(QMovie* self, QSize* size);
extern __declspec(dllexport) CacheMode QMovie_CacheMode(const QMovie* self);
extern __declspec(dllexport) void QMovie_SetCacheMode(QMovie* self, CacheMode mode);
extern __declspec(dllexport) void QMovie_Started(QMovie* self);
void QMovie_connect_Started(QMovie* self, intptr_t slot);
extern __declspec(dllexport) void QMovie_Resized(QMovie* self, QSize* size);
void QMovie_connect_Resized(QMovie* self, intptr_t slot);
extern __declspec(dllexport) void QMovie_Updated(QMovie* self, QRect* rect);
void QMovie_connect_Updated(QMovie* self, intptr_t slot);
extern __declspec(dllexport) void QMovie_StateChanged(QMovie* self, int state);
void QMovie_connect_StateChanged(QMovie* self, intptr_t slot);
extern __declspec(dllexport) void QMovie_Error(QMovie* self, int error);
void QMovie_connect_Error(QMovie* self, intptr_t slot);
extern __declspec(dllexport) void QMovie_Finished(QMovie* self);
void QMovie_connect_Finished(QMovie* self, intptr_t slot);
extern __declspec(dllexport) void QMovie_FrameChanged(QMovie* self, int frameNumber);
void QMovie_connect_FrameChanged(QMovie* self, intptr_t slot);
extern __declspec(dllexport) void QMovie_Start(QMovie* self);
extern __declspec(dllexport) bool QMovie_JumpToNextFrame(QMovie* self);
extern __declspec(dllexport) void QMovie_SetPaused(QMovie* self, bool paused);
extern __declspec(dllexport) void QMovie_Stop(QMovie* self);
extern __declspec(dllexport) void QMovie_SetSpeed(QMovie* self, int percentSpeed);
extern __declspec(dllexport) struct miqt_string QMovie_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QMovie_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QMovie_override_virtual_Event(void* self, intptr_t slot);
bool QMovie_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QMovie_override_virtual_EventFilter(void* self, intptr_t slot);
bool QMovie_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QMovie_override_virtual_TimerEvent(void* self, intptr_t slot);
void QMovie_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QMovie_override_virtual_ChildEvent(void* self, intptr_t slot);
void QMovie_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QMovie_override_virtual_CustomEvent(void* self, intptr_t slot);
void QMovie_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QMovie_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QMovie_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QMovie_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QMovie_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QMovie_Delete(QMovie* self, bool isSubclass);

} 
