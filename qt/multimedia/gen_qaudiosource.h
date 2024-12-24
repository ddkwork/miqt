#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QAUDIOSOURCE_H
#define MIQT_QT_MULTIMEDIA_GEN_QAUDIOSOURCE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioDevice QAudioDevice;
typedef struct QAudioFormat QAudioFormat;
typedef struct QAudioSource QAudioSource;
typedef struct QIODevice QIODevice;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAudioSource* QAudioSource_new();
extern __declspec(dllexport) 
QAudioSource* QAudioSource_new2(QAudioDevice* audioDeviceInfo);
extern __declspec(dllexport) 
QAudioSource* QAudioSource_new3(QAudioFormat* format);
extern __declspec(dllexport) 
QAudioSource* QAudioSource_new4(QAudioFormat* format, QObject* parent);
extern __declspec(dllexport) 
QAudioSource* QAudioSource_new5(QAudioDevice* audioDeviceInfo, QAudioFormat* format);
extern __declspec(dllexport) 
QAudioSource* QAudioSource_new6(QAudioDevice* audioDeviceInfo, QAudioFormat* format, QObject* parent);
extern __declspec(dllexport) 
void QAudioSource_virtbase(QAudioSource* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QAudioSource_MetaObject(const QAudioSource* self);
extern __declspec(dllexport) 
void* QAudioSource_Metacast(QAudioSource* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAudioSource_Tr(const char* s);
extern __declspec(dllexport) 
bool QAudioSource_IsNull(const QAudioSource* self);
extern __declspec(dllexport) 
QAudioFormat* QAudioSource_Format(const QAudioSource* self);
extern __declspec(dllexport) 
void QAudioSource_Start(QAudioSource* self, QIODevice* device);
extern __declspec(dllexport) 
QIODevice* QAudioSource_Start2(QAudioSource* self);
extern __declspec(dllexport) 
void QAudioSource_Stop(QAudioSource* self);
extern __declspec(dllexport) 
void QAudioSource_Reset(QAudioSource* self);
extern __declspec(dllexport) 
void QAudioSource_Suspend(QAudioSource* self);
extern __declspec(dllexport) 
void QAudioSource_Resume(QAudioSource* self);
extern __declspec(dllexport) 
void QAudioSource_SetBufferSize(QAudioSource* self, ptrdiff_t bytes);
extern __declspec(dllexport) 
ptrdiff_t QAudioSource_BufferSize(const QAudioSource* self);
extern __declspec(dllexport) 
ptrdiff_t QAudioSource_BytesAvailable(const QAudioSource* self);
extern __declspec(dllexport) 
void QAudioSource_SetVolume(QAudioSource* self, double volume);
extern __declspec(dllexport) 
double QAudioSource_Volume(const QAudioSource* self);
extern __declspec(dllexport) 
long long QAudioSource_ProcessedUSecs(const QAudioSource* self);
extern __declspec(dllexport) 
long long QAudioSource_ElapsedUSecs(const QAudioSource* self);
extern __declspec(dllexport) 
QtAudio::Error QAudioSource_Error(const QAudioSource* self);
extern __declspec(dllexport) 
QtAudio::State QAudioSource_State(const QAudioSource* self);
extern __declspec(dllexport) 
void QAudioSource_StateChanged(QAudioSource* self, int state);
void QAudioSource_connect_StateChanged(QAudioSource* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QAudioSource_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAudioSource_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAudioSource_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAudioSource_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAudioSource_override_virtual_Metacast(void* self, intptr_t slot);
void* QAudioSource_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAudioSource_Delete(QAudioSource* self, bool isSubclass);

}
