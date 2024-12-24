#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QAUDIODECODER_H
#define MIQT_QT_MULTIMEDIA_GEN_QAUDIODECODER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioBuffer QAudioBuffer;
typedef struct QAudioDecoder QAudioDecoder;
typedef struct QAudioFormat QAudioFormat;
typedef struct QIODevice QIODevice;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QUrl QUrl;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAudioDecoder* QAudioDecoder_new();
extern __declspec(dllexport) 
QAudioDecoder* QAudioDecoder_new2(QObject* parent);
extern __declspec(dllexport) 
void QAudioDecoder_virtbase(QAudioDecoder* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QAudioDecoder_MetaObject(const QAudioDecoder* self);
extern __declspec(dllexport) 
void* QAudioDecoder_Metacast(QAudioDecoder* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAudioDecoder_Tr(const char* s);
extern __declspec(dllexport) 
bool QAudioDecoder_IsSupported(const QAudioDecoder* self);
extern __declspec(dllexport) 
bool QAudioDecoder_IsDecoding(const QAudioDecoder* self);
extern __declspec(dllexport) 
QUrl* QAudioDecoder_Source(const QAudioDecoder* self);
extern __declspec(dllexport) 
void QAudioDecoder_SetSource(QAudioDecoder* self, QUrl* fileName);
extern __declspec(dllexport) 
QIODevice* QAudioDecoder_SourceDevice(const QAudioDecoder* self);
extern __declspec(dllexport) 
void QAudioDecoder_SetSourceDevice(QAudioDecoder* self, QIODevice* device);
extern __declspec(dllexport) 
QAudioFormat* QAudioDecoder_AudioFormat(const QAudioDecoder* self);
extern __declspec(dllexport) 
void QAudioDecoder_SetAudioFormat(QAudioDecoder* self, QAudioFormat* format);
extern __declspec(dllexport) 
Error QAudioDecoder_Error(const QAudioDecoder* self);
extern __declspec(dllexport) 
struct miqt_string QAudioDecoder_ErrorString(const QAudioDecoder* self);
extern __declspec(dllexport) 
QAudioBuffer* QAudioDecoder_Read(const QAudioDecoder* self);
extern __declspec(dllexport) 
bool QAudioDecoder_BufferAvailable(const QAudioDecoder* self);
extern __declspec(dllexport) 
long long QAudioDecoder_Position(const QAudioDecoder* self);
extern __declspec(dllexport) 
long long QAudioDecoder_Duration(const QAudioDecoder* self);
extern __declspec(dllexport) 
void QAudioDecoder_Start(QAudioDecoder* self);
extern __declspec(dllexport) 
void QAudioDecoder_Stop(QAudioDecoder* self);
extern __declspec(dllexport) 
void QAudioDecoder_BufferAvailableChanged(QAudioDecoder* self, bool param1);
void QAudioDecoder_connect_BufferAvailableChanged(QAudioDecoder* self, intptr_t slot);
extern __declspec(dllexport) 
void QAudioDecoder_BufferReady(QAudioDecoder* self);
void QAudioDecoder_connect_BufferReady(QAudioDecoder* self, intptr_t slot);
extern __declspec(dllexport) 
void QAudioDecoder_Finished(QAudioDecoder* self);
void QAudioDecoder_connect_Finished(QAudioDecoder* self, intptr_t slot);
extern __declspec(dllexport) 
void QAudioDecoder_IsDecodingChanged(QAudioDecoder* self, bool param1);
void QAudioDecoder_connect_IsDecodingChanged(QAudioDecoder* self, intptr_t slot);
extern __declspec(dllexport) 
void QAudioDecoder_FormatChanged(QAudioDecoder* self, QAudioFormat* format);
void QAudioDecoder_connect_FormatChanged(QAudioDecoder* self, intptr_t slot);
extern __declspec(dllexport) 
void QAudioDecoder_ErrorWithError(QAudioDecoder* self, int error);
void QAudioDecoder_connect_ErrorWithError(QAudioDecoder* self, intptr_t slot);
extern __declspec(dllexport) 
void QAudioDecoder_SourceChanged(QAudioDecoder* self);
void QAudioDecoder_connect_SourceChanged(QAudioDecoder* self, intptr_t slot);
extern __declspec(dllexport) 
void QAudioDecoder_PositionChanged(QAudioDecoder* self, long long position);
void QAudioDecoder_connect_PositionChanged(QAudioDecoder* self, intptr_t slot);
extern __declspec(dllexport) 
void QAudioDecoder_DurationChanged(QAudioDecoder* self, long long duration);
void QAudioDecoder_connect_DurationChanged(QAudioDecoder* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QAudioDecoder_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAudioDecoder_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAudioDecoder_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAudioDecoder_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAudioDecoder_override_virtual_Metacast(void* self, intptr_t slot);
void* QAudioDecoder_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAudioDecoder_Delete(QAudioDecoder* self, bool isSubclass);

}
