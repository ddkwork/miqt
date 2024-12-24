#pragma once
#ifndef MIQT_QT_SPATIALAUDIO_GEN_QAUDIOENGINE_H
#define MIQT_QT_SPATIALAUDIO_GEN_QAUDIOENGINE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioDevice QAudioDevice;
typedef struct QAudioEngine QAudioEngine;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAudioEngine* QAudioEngine_new();
extern __declspec(dllexport) 
QAudioEngine* QAudioEngine_new2(QObject* parent);
extern __declspec(dllexport) 
QAudioEngine* QAudioEngine_new3(int sampleRate);
extern __declspec(dllexport) 
QAudioEngine* QAudioEngine_new4(int sampleRate, QObject* parent);
extern __declspec(dllexport) 
void QAudioEngine_virtbase(QAudioEngine* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QAudioEngine_MetaObject(const QAudioEngine* self);
extern __declspec(dllexport) 
void* QAudioEngine_Metacast(QAudioEngine* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAudioEngine_Tr(const char* s);
extern __declspec(dllexport) 
void QAudioEngine_SetOutputMode(QAudioEngine* self, OutputMode mode);
extern __declspec(dllexport) 
OutputMode QAudioEngine_OutputMode(const QAudioEngine* self);
extern __declspec(dllexport) 
int QAudioEngine_SampleRate(const QAudioEngine* self);
extern __declspec(dllexport) 
void QAudioEngine_SetOutputDevice(QAudioEngine* self, QAudioDevice* device);
extern __declspec(dllexport) 
QAudioDevice* QAudioEngine_OutputDevice(const QAudioEngine* self);
extern __declspec(dllexport) 
void QAudioEngine_SetMasterVolume(QAudioEngine* self, float volume);
extern __declspec(dllexport) 
float QAudioEngine_MasterVolume(const QAudioEngine* self);
extern __declspec(dllexport) 
void QAudioEngine_SetPaused(QAudioEngine* self, bool paused);
extern __declspec(dllexport) 
bool QAudioEngine_Paused(const QAudioEngine* self);
extern __declspec(dllexport) 
void QAudioEngine_SetRoomEffectsEnabled(QAudioEngine* self, bool enabled);
extern __declspec(dllexport) 
bool QAudioEngine_RoomEffectsEnabled(const QAudioEngine* self);
extern __declspec(dllexport) 
void QAudioEngine_SetDistanceScale(QAudioEngine* self, float scale);
extern __declspec(dllexport) 
float QAudioEngine_DistanceScale(const QAudioEngine* self);
extern __declspec(dllexport) 
void QAudioEngine_OutputModeChanged(QAudioEngine* self);
void QAudioEngine_connect_OutputModeChanged(QAudioEngine* self, intptr_t slot);
extern __declspec(dllexport) 
void QAudioEngine_OutputDeviceChanged(QAudioEngine* self);
void QAudioEngine_connect_OutputDeviceChanged(QAudioEngine* self, intptr_t slot);
extern __declspec(dllexport) 
void QAudioEngine_MasterVolumeChanged(QAudioEngine* self);
void QAudioEngine_connect_MasterVolumeChanged(QAudioEngine* self, intptr_t slot);
extern __declspec(dllexport) 
void QAudioEngine_PausedChanged(QAudioEngine* self);
void QAudioEngine_connect_PausedChanged(QAudioEngine* self, intptr_t slot);
extern __declspec(dllexport) 
void QAudioEngine_DistanceScaleChanged(QAudioEngine* self);
void QAudioEngine_connect_DistanceScaleChanged(QAudioEngine* self, intptr_t slot);
extern __declspec(dllexport) 
void QAudioEngine_Start(QAudioEngine* self);
extern __declspec(dllexport) 
void QAudioEngine_Stop(QAudioEngine* self);
extern __declspec(dllexport) 
void QAudioEngine_Pause(QAudioEngine* self);
extern __declspec(dllexport) 
void QAudioEngine_Resume(QAudioEngine* self);
extern __declspec(dllexport) 
struct miqt_string QAudioEngine_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAudioEngine_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAudioEngine_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAudioEngine_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAudioEngine_override_virtual_Metacast(void* self, intptr_t slot);
void* QAudioEngine_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAudioEngine_Delete(QAudioEngine* self, bool isSubclass);

}
