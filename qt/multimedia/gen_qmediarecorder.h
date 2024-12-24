#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QMEDIARECORDER_H
#define MIQT_QT_MULTIMEDIA_GEN_QMEDIARECORDER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QIODevice QIODevice;
typedef struct QMediaCaptureSession QMediaCaptureSession;
typedef struct QMediaFormat QMediaFormat;
typedef struct QMediaMetaData QMediaMetaData;
typedef struct QMediaRecorder QMediaRecorder;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSize QSize;
typedef struct QUrl QUrl;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QMediaRecorder* QMediaRecorder_new();
extern __declspec(dllexport) 
QMediaRecorder* QMediaRecorder_new2(QObject* parent);
extern __declspec(dllexport) 
void QMediaRecorder_virtbase(QMediaRecorder* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QMediaRecorder_MetaObject(const QMediaRecorder* self);
extern __declspec(dllexport) 
void* QMediaRecorder_Metacast(QMediaRecorder* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QMediaRecorder_Tr(const char* s);
extern __declspec(dllexport) 
bool QMediaRecorder_IsAvailable(const QMediaRecorder* self);
extern __declspec(dllexport) 
QUrl* QMediaRecorder_OutputLocation(const QMediaRecorder* self);
extern __declspec(dllexport) 
void QMediaRecorder_SetOutputLocation(QMediaRecorder* self, QUrl* location);
extern __declspec(dllexport) 
void QMediaRecorder_SetOutputDevice(QMediaRecorder* self, QIODevice* device);
extern __declspec(dllexport) 
QIODevice* QMediaRecorder_OutputDevice(const QMediaRecorder* self);
extern __declspec(dllexport) 
QUrl* QMediaRecorder_ActualLocation(const QMediaRecorder* self);
extern __declspec(dllexport) 
RecorderState QMediaRecorder_RecorderState(const QMediaRecorder* self);
extern __declspec(dllexport) 
Error QMediaRecorder_Error(const QMediaRecorder* self);
extern __declspec(dllexport) 
struct miqt_string QMediaRecorder_ErrorString(const QMediaRecorder* self);
extern __declspec(dllexport) 
long long QMediaRecorder_Duration(const QMediaRecorder* self);
extern __declspec(dllexport) 
QMediaFormat* QMediaRecorder_MediaFormat(const QMediaRecorder* self);
extern __declspec(dllexport) 
void QMediaRecorder_SetMediaFormat(QMediaRecorder* self, QMediaFormat* format);
extern __declspec(dllexport) 
EncodingMode QMediaRecorder_EncodingMode(const QMediaRecorder* self);
extern __declspec(dllexport) 
void QMediaRecorder_SetEncodingMode(QMediaRecorder* self, EncodingMode encodingMode);
extern __declspec(dllexport) 
Quality QMediaRecorder_Quality(const QMediaRecorder* self);
extern __declspec(dllexport) 
void QMediaRecorder_SetQuality(QMediaRecorder* self, Quality quality);
extern __declspec(dllexport) 
QSize* QMediaRecorder_VideoResolution(const QMediaRecorder* self);
extern __declspec(dllexport) 
void QMediaRecorder_SetVideoResolution(QMediaRecorder* self, QSize* videoResolution);
extern __declspec(dllexport) 
void QMediaRecorder_SetVideoResolution2(QMediaRecorder* self, int width, int height);
extern __declspec(dllexport) 
double QMediaRecorder_VideoFrameRate(const QMediaRecorder* self);
extern __declspec(dllexport) 
void QMediaRecorder_SetVideoFrameRate(QMediaRecorder* self, double frameRate);
extern __declspec(dllexport) 
int QMediaRecorder_VideoBitRate(const QMediaRecorder* self);
extern __declspec(dllexport) 
void QMediaRecorder_SetVideoBitRate(QMediaRecorder* self, int bitRate);
extern __declspec(dllexport) 
int QMediaRecorder_AudioBitRate(const QMediaRecorder* self);
extern __declspec(dllexport) 
void QMediaRecorder_SetAudioBitRate(QMediaRecorder* self, int bitRate);
extern __declspec(dllexport) 
int QMediaRecorder_AudioChannelCount(const QMediaRecorder* self);
extern __declspec(dllexport) 
void QMediaRecorder_SetAudioChannelCount(QMediaRecorder* self, int channels);
extern __declspec(dllexport) 
int QMediaRecorder_AudioSampleRate(const QMediaRecorder* self);
extern __declspec(dllexport) 
void QMediaRecorder_SetAudioSampleRate(QMediaRecorder* self, int sampleRate);
extern __declspec(dllexport) 
QMediaMetaData* QMediaRecorder_MetaData(const QMediaRecorder* self);
extern __declspec(dllexport) 
void QMediaRecorder_SetMetaData(QMediaRecorder* self, QMediaMetaData* metaData);
extern __declspec(dllexport) 
void QMediaRecorder_AddMetaData(QMediaRecorder* self, QMediaMetaData* metaData);
extern __declspec(dllexport) 
bool QMediaRecorder_AutoStop(const QMediaRecorder* self);
extern __declspec(dllexport) 
void QMediaRecorder_SetAutoStop(QMediaRecorder* self, bool autoStop);
extern __declspec(dllexport) 
QMediaCaptureSession* QMediaRecorder_CaptureSession(const QMediaRecorder* self);
extern __declspec(dllexport) 
void QMediaRecorder_Record(QMediaRecorder* self);
extern __declspec(dllexport) 
void QMediaRecorder_Pause(QMediaRecorder* self);
extern __declspec(dllexport) 
void QMediaRecorder_Stop(QMediaRecorder* self);
extern __declspec(dllexport) 
void QMediaRecorder_RecorderStateChanged(QMediaRecorder* self, RecorderState state);
void QMediaRecorder_connect_RecorderStateChanged(QMediaRecorder* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaRecorder_DurationChanged(QMediaRecorder* self, long long duration);
void QMediaRecorder_connect_DurationChanged(QMediaRecorder* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaRecorder_ActualLocationChanged(QMediaRecorder* self, QUrl* location);
void QMediaRecorder_connect_ActualLocationChanged(QMediaRecorder* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaRecorder_EncoderSettingsChanged(QMediaRecorder* self);
void QMediaRecorder_connect_EncoderSettingsChanged(QMediaRecorder* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaRecorder_ErrorOccurred(QMediaRecorder* self, Error error, struct miqt_string errorString);
void QMediaRecorder_connect_ErrorOccurred(QMediaRecorder* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaRecorder_ErrorChanged(QMediaRecorder* self);
void QMediaRecorder_connect_ErrorChanged(QMediaRecorder* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaRecorder_MetaDataChanged(QMediaRecorder* self);
void QMediaRecorder_connect_MetaDataChanged(QMediaRecorder* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaRecorder_MediaFormatChanged(QMediaRecorder* self);
void QMediaRecorder_connect_MediaFormatChanged(QMediaRecorder* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaRecorder_EncodingModeChanged(QMediaRecorder* self);
void QMediaRecorder_connect_EncodingModeChanged(QMediaRecorder* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaRecorder_QualityChanged(QMediaRecorder* self);
void QMediaRecorder_connect_QualityChanged(QMediaRecorder* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaRecorder_VideoResolutionChanged(QMediaRecorder* self);
void QMediaRecorder_connect_VideoResolutionChanged(QMediaRecorder* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaRecorder_VideoFrameRateChanged(QMediaRecorder* self);
void QMediaRecorder_connect_VideoFrameRateChanged(QMediaRecorder* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaRecorder_VideoBitRateChanged(QMediaRecorder* self);
void QMediaRecorder_connect_VideoBitRateChanged(QMediaRecorder* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaRecorder_AudioBitRateChanged(QMediaRecorder* self);
void QMediaRecorder_connect_AudioBitRateChanged(QMediaRecorder* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaRecorder_AudioChannelCountChanged(QMediaRecorder* self);
void QMediaRecorder_connect_AudioChannelCountChanged(QMediaRecorder* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaRecorder_AudioSampleRateChanged(QMediaRecorder* self);
void QMediaRecorder_connect_AudioSampleRateChanged(QMediaRecorder* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaRecorder_AutoStopChanged(QMediaRecorder* self);
void QMediaRecorder_connect_AutoStopChanged(QMediaRecorder* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QMediaRecorder_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QMediaRecorder_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QMediaRecorder_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QMediaRecorder_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QMediaRecorder_override_virtual_Metacast(void* self, intptr_t slot);
void* QMediaRecorder_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QMediaRecorder_Delete(QMediaRecorder* self, bool isSubclass);

}
