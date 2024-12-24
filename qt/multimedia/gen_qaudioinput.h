#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QAUDIOINPUT_H
#define MIQT_QT_MULTIMEDIA_GEN_QAUDIOINPUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioDevice QAudioDevice;
typedef struct QAudioInput QAudioInput;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAudioInput* QAudioInput_new();
extern __declspec(dllexport) 
QAudioInput* QAudioInput_new2(QAudioDevice* deviceInfo);
extern __declspec(dllexport) 
QAudioInput* QAudioInput_new3(QObject* parent);
extern __declspec(dllexport) 
QAudioInput* QAudioInput_new4(QAudioDevice* deviceInfo, QObject* parent);
extern __declspec(dllexport) 
void QAudioInput_virtbase(QAudioInput* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QAudioInput_MetaObject(const QAudioInput* self);
extern __declspec(dllexport) 
void* QAudioInput_Metacast(QAudioInput* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAudioInput_Tr(const char* s);
extern __declspec(dllexport) 
QAudioDevice* QAudioInput_Device(const QAudioInput* self);
extern __declspec(dllexport) 
float QAudioInput_Volume(const QAudioInput* self);
extern __declspec(dllexport) 
bool QAudioInput_IsMuted(const QAudioInput* self);
extern __declspec(dllexport) 
void QAudioInput_SetDevice(QAudioInput* self, QAudioDevice* device);
extern __declspec(dllexport) 
void QAudioInput_SetVolume(QAudioInput* self, float volume);
extern __declspec(dllexport) 
void QAudioInput_SetMuted(QAudioInput* self, bool muted);
extern __declspec(dllexport) 
void QAudioInput_DeviceChanged(QAudioInput* self);
void QAudioInput_connect_DeviceChanged(QAudioInput* self, intptr_t slot);
extern __declspec(dllexport) 
void QAudioInput_VolumeChanged(QAudioInput* self, float volume);
void QAudioInput_connect_VolumeChanged(QAudioInput* self, intptr_t slot);
extern __declspec(dllexport) 
void QAudioInput_MutedChanged(QAudioInput* self, bool muted);
void QAudioInput_connect_MutedChanged(QAudioInput* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QAudioInput_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAudioInput_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAudioInput_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAudioInput_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAudioInput_override_virtual_Metacast(void* self, intptr_t slot);
void* QAudioInput_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAudioInput_Delete(QAudioInput* self, bool isSubclass);

}
