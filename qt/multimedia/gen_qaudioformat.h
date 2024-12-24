#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QAUDIOFORMAT_H
#define MIQT_QT_MULTIMEDIA_GEN_QAUDIOFORMAT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioFormat QAudioFormat;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAudioFormat* QAudioFormat_new();
extern __declspec(dllexport) 
QAudioFormat* QAudioFormat_new2(QAudioFormat* param1);
extern __declspec(dllexport) 
bool QAudioFormat_IsValid(const QAudioFormat* self);
extern __declspec(dllexport) 
void QAudioFormat_SetSampleRate(QAudioFormat* self, int sampleRate);
extern __declspec(dllexport) 
int QAudioFormat_SampleRate(const QAudioFormat* self);
extern __declspec(dllexport) 
void QAudioFormat_SetChannelConfig(QAudioFormat* self, ChannelConfig config);
extern __declspec(dllexport) 
ChannelConfig QAudioFormat_ChannelConfig(const QAudioFormat* self);
extern __declspec(dllexport) 
void QAudioFormat_SetChannelCount(QAudioFormat* self, int channelCount);
extern __declspec(dllexport) 
int QAudioFormat_ChannelCount(const QAudioFormat* self);
extern __declspec(dllexport) 
int QAudioFormat_ChannelOffset(const QAudioFormat* self, AudioChannelPosition channel);
extern __declspec(dllexport) 
void QAudioFormat_SetSampleFormat(QAudioFormat* self, SampleFormat f);
extern __declspec(dllexport) 
SampleFormat QAudioFormat_SampleFormat(const QAudioFormat* self);
extern __declspec(dllexport) 
int QAudioFormat_BytesForDuration(const QAudioFormat* self, long long microseconds);
extern __declspec(dllexport) 
long long QAudioFormat_DurationForBytes(const QAudioFormat* self, int byteCount);
extern __declspec(dllexport) 
int QAudioFormat_BytesForFrames(const QAudioFormat* self, int frameCount);
extern __declspec(dllexport) 
int QAudioFormat_FramesForBytes(const QAudioFormat* self, int byteCount);
extern __declspec(dllexport) 
int QAudioFormat_FramesForDuration(const QAudioFormat* self, long long microseconds);
extern __declspec(dllexport) 
long long QAudioFormat_DurationForFrames(const QAudioFormat* self, int frameCount);
extern __declspec(dllexport) 
int QAudioFormat_BytesPerFrame(const QAudioFormat* self);
extern __declspec(dllexport) 
int QAudioFormat_BytesPerSample(const QAudioFormat* self);
extern __declspec(dllexport) 
float QAudioFormat_NormalizedSampleValue(const QAudioFormat* self, const void* sample);
extern __declspec(dllexport) 
ChannelConfig QAudioFormat_DefaultChannelConfigForChannelCount(int channelCount);
extern __declspec(dllexport) 
void QAudioFormat_Delete(QAudioFormat* self, bool isSubclass);

}
