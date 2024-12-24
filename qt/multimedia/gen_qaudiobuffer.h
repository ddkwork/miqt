#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QAUDIOBUFFER_H
#define MIQT_QT_MULTIMEDIA_GEN_QAUDIOBUFFER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioBuffer QAudioBuffer;
typedef struct QAudioFormat QAudioFormat;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAudioBuffer* QAudioBuffer_new();
extern __declspec(dllexport) 
QAudioBuffer* QAudioBuffer_new2(QAudioBuffer* other);
extern __declspec(dllexport) 
QAudioBuffer* QAudioBuffer_new3(struct miqt_string data, QAudioFormat* format);
extern __declspec(dllexport) 
QAudioBuffer* QAudioBuffer_new4(int numFrames, QAudioFormat* format);
extern __declspec(dllexport) 
QAudioBuffer* QAudioBuffer_new5(struct miqt_string data, QAudioFormat* format, long long startTime);
extern __declspec(dllexport) 
QAudioBuffer* QAudioBuffer_new6(int numFrames, QAudioFormat* format, long long startTime);
extern __declspec(dllexport) 
void QAudioBuffer_OperatorAssign(QAudioBuffer* self, QAudioBuffer* other);
extern __declspec(dllexport) 
void QAudioBuffer_Swap(QAudioBuffer* self, QAudioBuffer* other);
extern __declspec(dllexport) 
bool QAudioBuffer_IsValid(const QAudioBuffer* self);
extern __declspec(dllexport) 
void QAudioBuffer_Detach(QAudioBuffer* self);
extern __declspec(dllexport) 
QAudioFormat* QAudioBuffer_Format(const QAudioBuffer* self);
extern __declspec(dllexport) 
ptrdiff_t QAudioBuffer_FrameCount(const QAudioBuffer* self);
extern __declspec(dllexport) 
ptrdiff_t QAudioBuffer_SampleCount(const QAudioBuffer* self);
extern __declspec(dllexport) 
ptrdiff_t QAudioBuffer_ByteCount(const QAudioBuffer* self);
extern __declspec(dllexport) 
long long QAudioBuffer_Duration(const QAudioBuffer* self);
extern __declspec(dllexport) 
long long QAudioBuffer_StartTime(const QAudioBuffer* self);
extern __declspec(dllexport) 
void QAudioBuffer_Delete(QAudioBuffer* self, bool isSubclass);

}
