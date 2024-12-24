#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QWAVEDECODER_H
#define MIQT_QT_MULTIMEDIA_GEN_QWAVEDECODER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioFormat QAudioFormat;
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QWaveDecoder QWaveDecoder;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QWaveDecoder* QWaveDecoder_new(QIODevice* device);
extern __declspec(dllexport) 
QWaveDecoder* QWaveDecoder_new2(QIODevice* device, QAudioFormat* format);
extern __declspec(dllexport) 
QWaveDecoder* QWaveDecoder_new3(QIODevice* device, QObject* parent);
extern __declspec(dllexport) 
QWaveDecoder* QWaveDecoder_new4(QIODevice* device, QAudioFormat* format, QObject* parent);
extern __declspec(dllexport) 
void QWaveDecoder_virtbase(QWaveDecoder* src
, QIODevice** outptr_QIODevice
);
extern __declspec(dllexport) 
QMetaObject* QWaveDecoder_MetaObject(const QWaveDecoder* self);
extern __declspec(dllexport) 
void* QWaveDecoder_Metacast(QWaveDecoder* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QWaveDecoder_Tr(const char* s);
extern __declspec(dllexport) 
QAudioFormat* QWaveDecoder_AudioFormat(const QWaveDecoder* self);
extern __declspec(dllexport) 
QIODevice* QWaveDecoder_GetDevice(QWaveDecoder* self);
extern __declspec(dllexport) 
int QWaveDecoder_Duration(const QWaveDecoder* self);
extern __declspec(dllexport) 
long long QWaveDecoder_HeaderLength();
extern __declspec(dllexport) 
bool QWaveDecoder_Open(QWaveDecoder* self, int mode);
extern __declspec(dllexport) 
void QWaveDecoder_Close(QWaveDecoder* self);
extern __declspec(dllexport) 
bool QWaveDecoder_Seek(QWaveDecoder* self, long long pos);
extern __declspec(dllexport) 
long long QWaveDecoder_Pos(const QWaveDecoder* self);
extern __declspec(dllexport) 
long long QWaveDecoder_Size(const QWaveDecoder* self);
extern __declspec(dllexport) 
bool QWaveDecoder_IsSequential(const QWaveDecoder* self);
extern __declspec(dllexport) 
long long QWaveDecoder_BytesAvailable(const QWaveDecoder* self);
extern __declspec(dllexport) 
void QWaveDecoder_FormatKnown(QWaveDecoder* self);
void QWaveDecoder_connect_FormatKnown(QWaveDecoder* self, intptr_t slot);
extern __declspec(dllexport) 
void QWaveDecoder_ParsingError(QWaveDecoder* self);
void QWaveDecoder_connect_ParsingError(QWaveDecoder* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QWaveDecoder_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QWaveDecoder_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QWaveDecoder_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QWaveDecoder_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QWaveDecoder_override_virtual_Metacast(void* self, intptr_t slot);
void* QWaveDecoder_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QWaveDecoder_Delete(QWaveDecoder* self, bool isSubclass);

}
