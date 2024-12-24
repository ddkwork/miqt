#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QAUDIOOUTPUT_H
#define MIQT_QT_MULTIMEDIA_GEN_QAUDIOOUTPUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioDevice QAudioDevice;
typedef struct QAudioOutput QAudioOutput;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAudioOutput* QAudioOutput_new();
extern __declspec(dllexport) 
QAudioOutput* QAudioOutput_new2(QAudioDevice* device);
extern __declspec(dllexport) 
QAudioOutput* QAudioOutput_new3(QObject* parent);
extern __declspec(dllexport) 
QAudioOutput* QAudioOutput_new4(QAudioDevice* device, QObject* parent);
extern __declspec(dllexport) 
void QAudioOutput_virtbase(QAudioOutput* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QAudioOutput_MetaObject(const QAudioOutput* self);
extern __declspec(dllexport) 
void* QAudioOutput_Metacast(QAudioOutput* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAudioOutput_Tr(const char* s);
extern __declspec(dllexport) 
QAudioDevice* QAudioOutput_Device(const QAudioOutput* self);
extern __declspec(dllexport) 
float QAudioOutput_Volume(const QAudioOutput* self);
extern __declspec(dllexport) 
bool QAudioOutput_IsMuted(const QAudioOutput* self);
extern __declspec(dllexport) 
void QAudioOutput_SetDevice(QAudioOutput* self, QAudioDevice* device);
extern __declspec(dllexport) 
void QAudioOutput_SetVolume(QAudioOutput* self, float volume);
extern __declspec(dllexport) 
void QAudioOutput_SetMuted(QAudioOutput* self, bool muted);
extern __declspec(dllexport) 
void QAudioOutput_DeviceChanged(QAudioOutput* self);
void QAudioOutput_connect_DeviceChanged(QAudioOutput* self, intptr_t slot);
extern __declspec(dllexport) 
void QAudioOutput_VolumeChanged(QAudioOutput* self, float volume);
void QAudioOutput_connect_VolumeChanged(QAudioOutput* self, intptr_t slot);
extern __declspec(dllexport) 
void QAudioOutput_MutedChanged(QAudioOutput* self, bool muted);
void QAudioOutput_connect_MutedChanged(QAudioOutput* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QAudioOutput_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAudioOutput_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAudioOutput_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAudioOutput_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAudioOutput_override_virtual_Metacast(void* self, intptr_t slot);
void* QAudioOutput_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAudioOutput_Delete(QAudioOutput* self, bool isSubclass);

}
