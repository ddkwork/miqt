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
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QUrl QUrl;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAmbientSound* QAmbientSound_new(QAudioEngine* engine);
extern __declspec(dllexport) 
void QAmbientSound_virtbase(QAmbientSound* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QAmbientSound_MetaObject(const QAmbientSound* self);
extern __declspec(dllexport) 
void* QAmbientSound_Metacast(QAmbientSound* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAmbientSound_Tr(const char* s);
extern __declspec(dllexport) 
void QAmbientSound_SetSource(QAmbientSound* self, QUrl* url);
extern __declspec(dllexport) 
QUrl* QAmbientSound_Source(const QAmbientSound* self);
extern __declspec(dllexport) 
int QAmbientSound_Loops(const QAmbientSound* self);
extern __declspec(dllexport) 
void QAmbientSound_SetLoops(QAmbientSound* self, int loops);
extern __declspec(dllexport) 
bool QAmbientSound_AutoPlay(const QAmbientSound* self);
extern __declspec(dllexport) 
void QAmbientSound_SetAutoPlay(QAmbientSound* self, bool autoPlay);
extern __declspec(dllexport) 
void QAmbientSound_SetVolume(QAmbientSound* self, float volume);
extern __declspec(dllexport) 
float QAmbientSound_Volume(const QAmbientSound* self);
extern __declspec(dllexport) 
QAudioEngine* QAmbientSound_Engine(const QAmbientSound* self);
extern __declspec(dllexport) 
void QAmbientSound_SourceChanged(QAmbientSound* self);
void QAmbientSound_connect_SourceChanged(QAmbientSound* self, intptr_t slot);
extern __declspec(dllexport) 
void QAmbientSound_LoopsChanged(QAmbientSound* self);
void QAmbientSound_connect_LoopsChanged(QAmbientSound* self, intptr_t slot);
extern __declspec(dllexport) 
void QAmbientSound_AutoPlayChanged(QAmbientSound* self);
void QAmbientSound_connect_AutoPlayChanged(QAmbientSound* self, intptr_t slot);
extern __declspec(dllexport) 
void QAmbientSound_VolumeChanged(QAmbientSound* self);
void QAmbientSound_connect_VolumeChanged(QAmbientSound* self, intptr_t slot);
extern __declspec(dllexport) 
void QAmbientSound_Play(QAmbientSound* self);
extern __declspec(dllexport) 
void QAmbientSound_Pause(QAmbientSound* self);
extern __declspec(dllexport) 
void QAmbientSound_Stop(QAmbientSound* self);
extern __declspec(dllexport) 
struct miqt_string QAmbientSound_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAmbientSound_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAmbientSound_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAmbientSound_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAmbientSound_override_virtual_Metacast(void* self, intptr_t slot);
void* QAmbientSound_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAmbientSound_Delete(QAmbientSound* self, bool isSubclass);

}
