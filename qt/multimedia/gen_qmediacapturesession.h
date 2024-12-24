#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QMEDIACAPTURESESSION_H
#define MIQT_QT_MULTIMEDIA_GEN_QMEDIACAPTURESESSION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioBufferInput QAudioBufferInput;
typedef struct QAudioInput QAudioInput;
typedef struct QAudioOutput QAudioOutput;
typedef struct QCamera QCamera;
typedef struct QImageCapture QImageCapture;
typedef struct QMediaCaptureSession QMediaCaptureSession;
typedef struct QMediaRecorder QMediaRecorder;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QScreenCapture QScreenCapture;
typedef struct QVideoFrameInput QVideoFrameInput;
typedef struct QVideoSink QVideoSink;
typedef struct QWindowCapture QWindowCapture;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QMediaCaptureSession* QMediaCaptureSession_new();
extern __declspec(dllexport) 
QMediaCaptureSession* QMediaCaptureSession_new2(QObject* parent);
extern __declspec(dllexport) 
void QMediaCaptureSession_virtbase(QMediaCaptureSession* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QMediaCaptureSession_MetaObject(const QMediaCaptureSession* self);
extern __declspec(dllexport) 
void* QMediaCaptureSession_Metacast(QMediaCaptureSession* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QMediaCaptureSession_Tr(const char* s);
extern __declspec(dllexport) 
QAudioInput* QMediaCaptureSession_AudioInput(const QMediaCaptureSession* self);
extern __declspec(dllexport) 
void QMediaCaptureSession_SetAudioInput(QMediaCaptureSession* self, QAudioInput* input);
extern __declspec(dllexport) 
QAudioBufferInput* QMediaCaptureSession_AudioBufferInput(const QMediaCaptureSession* self);
extern __declspec(dllexport) 
void QMediaCaptureSession_SetAudioBufferInput(QMediaCaptureSession* self, QAudioBufferInput* input);
extern __declspec(dllexport) 
QCamera* QMediaCaptureSession_Camera(const QMediaCaptureSession* self);
extern __declspec(dllexport) 
void QMediaCaptureSession_SetCamera(QMediaCaptureSession* self, QCamera* camera);
extern __declspec(dllexport) 
QImageCapture* QMediaCaptureSession_ImageCapture(QMediaCaptureSession* self);
extern __declspec(dllexport) 
void QMediaCaptureSession_SetImageCapture(QMediaCaptureSession* self, QImageCapture* imageCapture);
extern __declspec(dllexport) 
QScreenCapture* QMediaCaptureSession_ScreenCapture(QMediaCaptureSession* self);
extern __declspec(dllexport) 
void QMediaCaptureSession_SetScreenCapture(QMediaCaptureSession* self, QScreenCapture* screenCapture);
extern __declspec(dllexport) 
QWindowCapture* QMediaCaptureSession_WindowCapture(QMediaCaptureSession* self);
extern __declspec(dllexport) 
void QMediaCaptureSession_SetWindowCapture(QMediaCaptureSession* self, QWindowCapture* windowCapture);
extern __declspec(dllexport) 
QVideoFrameInput* QMediaCaptureSession_VideoFrameInput(const QMediaCaptureSession* self);
extern __declspec(dllexport) 
void QMediaCaptureSession_SetVideoFrameInput(QMediaCaptureSession* self, QVideoFrameInput* input);
extern __declspec(dllexport) 
QMediaRecorder* QMediaCaptureSession_Recorder(QMediaCaptureSession* self);
extern __declspec(dllexport) 
void QMediaCaptureSession_SetRecorder(QMediaCaptureSession* self, QMediaRecorder* recorder);
extern __declspec(dllexport) 
void QMediaCaptureSession_SetVideoOutput(QMediaCaptureSession* self, QObject* output);
extern __declspec(dllexport) 
QObject* QMediaCaptureSession_VideoOutput(const QMediaCaptureSession* self);
extern __declspec(dllexport) 
void QMediaCaptureSession_SetVideoSink(QMediaCaptureSession* self, QVideoSink* sink);
extern __declspec(dllexport) 
QVideoSink* QMediaCaptureSession_VideoSink(const QMediaCaptureSession* self);
extern __declspec(dllexport) 
void QMediaCaptureSession_SetAudioOutput(QMediaCaptureSession* self, QAudioOutput* output);
extern __declspec(dllexport) 
QAudioOutput* QMediaCaptureSession_AudioOutput(const QMediaCaptureSession* self);
extern __declspec(dllexport) 
void QMediaCaptureSession_AudioInputChanged(QMediaCaptureSession* self);
void QMediaCaptureSession_connect_AudioInputChanged(QMediaCaptureSession* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaCaptureSession_AudioBufferInputChanged(QMediaCaptureSession* self);
void QMediaCaptureSession_connect_AudioBufferInputChanged(QMediaCaptureSession* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaCaptureSession_CameraChanged(QMediaCaptureSession* self);
void QMediaCaptureSession_connect_CameraChanged(QMediaCaptureSession* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaCaptureSession_ScreenCaptureChanged(QMediaCaptureSession* self);
void QMediaCaptureSession_connect_ScreenCaptureChanged(QMediaCaptureSession* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaCaptureSession_WindowCaptureChanged(QMediaCaptureSession* self);
void QMediaCaptureSession_connect_WindowCaptureChanged(QMediaCaptureSession* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaCaptureSession_VideoFrameInputChanged(QMediaCaptureSession* self);
void QMediaCaptureSession_connect_VideoFrameInputChanged(QMediaCaptureSession* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaCaptureSession_ImageCaptureChanged(QMediaCaptureSession* self);
void QMediaCaptureSession_connect_ImageCaptureChanged(QMediaCaptureSession* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaCaptureSession_RecorderChanged(QMediaCaptureSession* self);
void QMediaCaptureSession_connect_RecorderChanged(QMediaCaptureSession* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaCaptureSession_VideoOutputChanged(QMediaCaptureSession* self);
void QMediaCaptureSession_connect_VideoOutputChanged(QMediaCaptureSession* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaCaptureSession_AudioOutputChanged(QMediaCaptureSession* self);
void QMediaCaptureSession_connect_AudioOutputChanged(QMediaCaptureSession* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QMediaCaptureSession_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QMediaCaptureSession_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QMediaCaptureSession_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QMediaCaptureSession_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QMediaCaptureSession_override_virtual_Metacast(void* self, intptr_t slot);
void* QMediaCaptureSession_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QMediaCaptureSession_Delete(QMediaCaptureSession* self, bool isSubclass);

}
