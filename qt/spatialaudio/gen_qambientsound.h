#pragma once
#ifndef MIQT_QT_SPATIALAUDIO_GEN_QAMBIENTSOUND_H
#define MIQT_QT_SPATIALAUDIO_GEN_QAMBIENTSOUND_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAmbientSound QAmbientSound;
typedef struct QAudioEngine QAudioEngine;
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;
typedef struct QUrl QUrl;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QAmbientSound* QAmbientSound_new(QAudioEngine* engine);
extern __declspec(dllexport) void QAmbientSound_virtbase(QAmbientSound* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QAmbientSound_MetaObject(const QAmbientSound* self);
extern __declspec(dllexport) void* QAmbientSound_Metacast(QAmbientSound* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QAmbientSound_Tr(const char* s);
extern __declspec(dllexport) void QAmbientSound_SetSource(QAmbientSound* self, QUrl* url);
extern __declspec(dllexport) QUrl* QAmbientSound_Source(const QAmbientSound* self);
extern __declspec(dllexport) int QAmbientSound_Loops(const QAmbientSound* self);
extern __declspec(dllexport) void QAmbientSound_SetLoops(QAmbientSound* self, int loops);
extern __declspec(dllexport) bool QAmbientSound_AutoPlay(const QAmbientSound* self);
extern __declspec(dllexport) void QAmbientSound_SetAutoPlay(QAmbientSound* self, bool autoPlay);
extern __declspec(dllexport) void QAmbientSound_SetVolume(QAmbientSound* self, float volume);
extern __declspec(dllexport) float QAmbientSound_Volume(const QAmbientSound* self);
extern __declspec(dllexport) QAudioEngine* QAmbientSound_Engine(const QAmbientSound* self);
extern __declspec(dllexport) void QAmbientSound_SourceChanged(QAmbientSound* self);
void QAmbientSound_connect_SourceChanged(QAmbientSound* self, intptr_t slot);
extern __declspec(dllexport) void QAmbientSound_LoopsChanged(QAmbientSound* self);
void QAmbientSound_connect_LoopsChanged(QAmbientSound* self, intptr_t slot);
extern __declspec(dllexport) void QAmbientSound_AutoPlayChanged(QAmbientSound* self);
void QAmbientSound_connect_AutoPlayChanged(QAmbientSound* self, intptr_t slot);
extern __declspec(dllexport) void QAmbientSound_VolumeChanged(QAmbientSound* self);
void QAmbientSound_connect_VolumeChanged(QAmbientSound* self, intptr_t slot);
extern __declspec(dllexport) void QAmbientSound_Play(QAmbientSound* self);
extern __declspec(dllexport) void QAmbientSound_Pause(QAmbientSound* self);
extern __declspec(dllexport) void QAmbientSound_Stop(QAmbientSound* self);
extern __declspec(dllexport) struct miqt_string QAmbientSound_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QAmbientSound_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QAmbientSound_override_virtual_Event(void* self, intptr_t slot);
bool QAmbientSound_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QAmbientSound_override_virtual_EventFilter(void* self, intptr_t slot);
bool QAmbientSound_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QAmbientSound_override_virtual_TimerEvent(void* self, intptr_t slot);
void QAmbientSound_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QAmbientSound_override_virtual_ChildEvent(void* self, intptr_t slot);
void QAmbientSound_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QAmbientSound_override_virtual_CustomEvent(void* self, intptr_t slot);
void QAmbientSound_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QAmbientSound_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QAmbientSound_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QAmbientSound_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QAmbientSound_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QAmbientSound_Delete(QAmbientSound* self, bool isSubclass);

} 
