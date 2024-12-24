#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QAUDIOBUFFERINPUT_H
#define MIQT_QT_MULTIMEDIA_GEN_QAUDIOBUFFERINPUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioBuffer QAudioBuffer;
typedef struct QAudioBufferInput QAudioBufferInput;
typedef struct QAudioFormat QAudioFormat;
typedef struct QMediaCaptureSession QMediaCaptureSession;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAudioBufferInput* QAudioBufferInput_new();
extern __declspec(dllexport) 
QAudioBufferInput* QAudioBufferInput_new2(QAudioFormat* format);
extern __declspec(dllexport) 
QAudioBufferInput* QAudioBufferInput_new3(QObject* parent);
extern __declspec(dllexport) 
QAudioBufferInput* QAudioBufferInput_new4(QAudioFormat* format, QObject* parent);
extern __declspec(dllexport) 
void QAudioBufferInput_virtbase(QAudioBufferInput* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QAudioBufferInput_MetaObject(const QAudioBufferInput* self);
extern __declspec(dllexport) 
void* QAudioBufferInput_Metacast(QAudioBufferInput* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAudioBufferInput_Tr(const char* s);
extern __declspec(dllexport) 
bool QAudioBufferInput_SendAudioBuffer(QAudioBufferInput* self, QAudioBuffer* audioBuffer);
extern __declspec(dllexport) 
QAudioFormat* QAudioBufferInput_Format(const QAudioBufferInput* self);
extern __declspec(dllexport) 
QMediaCaptureSession* QAudioBufferInput_CaptureSession(const QAudioBufferInput* self);
extern __declspec(dllexport) 
void QAudioBufferInput_ReadyToSendAudioBuffer(QAudioBufferInput* self);
void QAudioBufferInput_connect_ReadyToSendAudioBuffer(QAudioBufferInput* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QAudioBufferInput_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAudioBufferInput_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAudioBufferInput_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAudioBufferInput_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAudioBufferInput_override_virtual_Metacast(void* self, intptr_t slot);
void* QAudioBufferInput_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAudioBufferInput_Delete(QAudioBufferInput* self, bool isSubclass);

}
