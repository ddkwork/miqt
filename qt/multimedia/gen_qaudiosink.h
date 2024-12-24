#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QAUDIOSINK_H
#define MIQT_QT_MULTIMEDIA_GEN_QAUDIOSINK_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioDevice QAudioDevice;
typedef struct QAudioFormat QAudioFormat;
typedef struct QAudioSink QAudioSink;
typedef struct QIODevice QIODevice;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAudioSink* QAudioSink_new();
extern __declspec(dllexport) 
QAudioSink* QAudioSink_new2(QAudioDevice* audioDeviceInfo);
extern __declspec(dllexport) 
QAudioSink* QAudioSink_new3(QAudioFormat* format);
extern __declspec(dllexport) 
QAudioSink* QAudioSink_new4(QAudioFormat* format, QObject* parent);
extern __declspec(dllexport) 
QAudioSink* QAudioSink_new5(QAudioDevice* audioDeviceInfo, QAudioFormat* format);
extern __declspec(dllexport) 
QAudioSink* QAudioSink_new6(QAudioDevice* audioDeviceInfo, QAudioFormat* format, QObject* parent);
extern __declspec(dllexport) 
void QAudioSink_virtbase(QAudioSink* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QAudioSink_MetaObject(const QAudioSink* self);
extern __declspec(dllexport) 
void* QAudioSink_Metacast(QAudioSink* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAudioSink_Tr(const char* s);
extern __declspec(dllexport) 
bool QAudioSink_IsNull(const QAudioSink* self);
extern __declspec(dllexport) 
QAudioFormat* QAudioSink_Format(const QAudioSink* self);
extern __declspec(dllexport) 
void QAudioSink_Start(QAudioSink* self, QIODevice* device);
extern __declspec(dllexport) 
QIODevice* QAudioSink_Start2(QAudioSink* self);
extern __declspec(dllexport) 
void QAudioSink_Stop(QAudioSink* self);
extern __declspec(dllexport) 
void QAudioSink_Reset(QAudioSink* self);
extern __declspec(dllexport) 
void QAudioSink_Suspend(QAudioSink* self);
extern __declspec(dllexport) 
void QAudioSink_Resume(QAudioSink* self);
extern __declspec(dllexport) 
void QAudioSink_SetBufferSize(QAudioSink* self, ptrdiff_t bytes);
extern __declspec(dllexport) 
ptrdiff_t QAudioSink_BufferSize(const QAudioSink* self);
extern __declspec(dllexport) 
ptrdiff_t QAudioSink_BytesFree(const QAudioSink* self);
extern __declspec(dllexport) 
long long QAudioSink_ProcessedUSecs(const QAudioSink* self);
extern __declspec(dllexport) 
long long QAudioSink_ElapsedUSecs(const QAudioSink* self);
extern __declspec(dllexport) 
QtAudio::Error QAudioSink_Error(const QAudioSink* self);
extern __declspec(dllexport) 
QtAudio::State QAudioSink_State(const QAudioSink* self);
extern __declspec(dllexport) 
void QAudioSink_SetVolume(QAudioSink* self, double volume);
extern __declspec(dllexport) 
double QAudioSink_Volume(const QAudioSink* self);
extern __declspec(dllexport) 
void QAudioSink_StateChanged(QAudioSink* self, int state);
void QAudioSink_connect_StateChanged(QAudioSink* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QAudioSink_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAudioSink_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAudioSink_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAudioSink_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAudioSink_override_virtual_Metacast(void* self, intptr_t slot);
void* QAudioSink_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAudioSink_Delete(QAudioSink* self, bool isSubclass);

}
