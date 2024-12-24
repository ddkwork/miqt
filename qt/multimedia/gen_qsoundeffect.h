#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QSOUNDEFFECT_H
#define MIQT_QT_MULTIMEDIA_GEN_QSOUNDEFFECT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioDevice QAudioDevice;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSoundEffect QSoundEffect;
typedef struct QUrl QUrl;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QSoundEffect* QSoundEffect_new();
extern __declspec(dllexport) 
QSoundEffect* QSoundEffect_new2(QAudioDevice* audioDevice);
extern __declspec(dllexport) 
QSoundEffect* QSoundEffect_new3(QObject* parent);
extern __declspec(dllexport) 
QSoundEffect* QSoundEffect_new4(QAudioDevice* audioDevice, QObject* parent);
extern __declspec(dllexport) 
void QSoundEffect_virtbase(QSoundEffect* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QSoundEffect_MetaObject(const QSoundEffect* self);
extern __declspec(dllexport) 
void* QSoundEffect_Metacast(QSoundEffect* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QSoundEffect_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QSoundEffect_SupportedMimeTypes();
extern __declspec(dllexport) 
QUrl* QSoundEffect_Source(const QSoundEffect* self);
extern __declspec(dllexport) 
void QSoundEffect_SetSource(QSoundEffect* self, QUrl* url);
extern __declspec(dllexport) 
int QSoundEffect_LoopCount(const QSoundEffect* self);
extern __declspec(dllexport) 
int QSoundEffect_LoopsRemaining(const QSoundEffect* self);
extern __declspec(dllexport) 
void QSoundEffect_SetLoopCount(QSoundEffect* self, int loopCount);
extern __declspec(dllexport) 
QAudioDevice* QSoundEffect_AudioDevice(QSoundEffect* self);
extern __declspec(dllexport) 
void QSoundEffect_SetAudioDevice(QSoundEffect* self, QAudioDevice* device);
extern __declspec(dllexport) 
float QSoundEffect_Volume(const QSoundEffect* self);
extern __declspec(dllexport) 
void QSoundEffect_SetVolume(QSoundEffect* self, float volume);
extern __declspec(dllexport) 
bool QSoundEffect_IsMuted(const QSoundEffect* self);
extern __declspec(dllexport) 
void QSoundEffect_SetMuted(QSoundEffect* self, bool muted);
extern __declspec(dllexport) 
bool QSoundEffect_IsLoaded(const QSoundEffect* self);
extern __declspec(dllexport) 
bool QSoundEffect_IsPlaying(const QSoundEffect* self);
extern __declspec(dllexport) 
Status QSoundEffect_Status(const QSoundEffect* self);
extern __declspec(dllexport) 
void QSoundEffect_SourceChanged(QSoundEffect* self);
void QSoundEffect_connect_SourceChanged(QSoundEffect* self, intptr_t slot);
extern __declspec(dllexport) 
void QSoundEffect_LoopCountChanged(QSoundEffect* self);
void QSoundEffect_connect_LoopCountChanged(QSoundEffect* self, intptr_t slot);
extern __declspec(dllexport) 
void QSoundEffect_LoopsRemainingChanged(QSoundEffect* self);
void QSoundEffect_connect_LoopsRemainingChanged(QSoundEffect* self, intptr_t slot);
extern __declspec(dllexport) 
void QSoundEffect_VolumeChanged(QSoundEffect* self);
void QSoundEffect_connect_VolumeChanged(QSoundEffect* self, intptr_t slot);
extern __declspec(dllexport) 
void QSoundEffect_MutedChanged(QSoundEffect* self);
void QSoundEffect_connect_MutedChanged(QSoundEffect* self, intptr_t slot);
extern __declspec(dllexport) 
void QSoundEffect_LoadedChanged(QSoundEffect* self);
void QSoundEffect_connect_LoadedChanged(QSoundEffect* self, intptr_t slot);
extern __declspec(dllexport) 
void QSoundEffect_PlayingChanged(QSoundEffect* self);
void QSoundEffect_connect_PlayingChanged(QSoundEffect* self, intptr_t slot);
extern __declspec(dllexport) 
void QSoundEffect_StatusChanged(QSoundEffect* self);
void QSoundEffect_connect_StatusChanged(QSoundEffect* self, intptr_t slot);
extern __declspec(dllexport) 
void QSoundEffect_AudioDeviceChanged(QSoundEffect* self);
void QSoundEffect_connect_AudioDeviceChanged(QSoundEffect* self, intptr_t slot);
extern __declspec(dllexport) 
void QSoundEffect_Play(QSoundEffect* self);
extern __declspec(dllexport) 
void QSoundEffect_Stop(QSoundEffect* self);
extern __declspec(dllexport) 
struct miqt_string QSoundEffect_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QSoundEffect_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QSoundEffect_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QSoundEffect_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QSoundEffect_override_virtual_Metacast(void* self, intptr_t slot);
void* QSoundEffect_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QSoundEffect_Delete(QSoundEffect* self, bool isSubclass);

}
