#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QVIDEOFRAMEINPUT_H
#define MIQT_QT_MULTIMEDIA_GEN_QVIDEOFRAMEINPUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMediaCaptureSession QMediaCaptureSession;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QVideoFrame QVideoFrame;
typedef struct QVideoFrameFormat QVideoFrameFormat;
typedef struct QVideoFrameInput QVideoFrameInput;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QVideoFrameInput* QVideoFrameInput_new();
extern __declspec(dllexport) 
QVideoFrameInput* QVideoFrameInput_new2(QVideoFrameFormat* format);
extern __declspec(dllexport) 
QVideoFrameInput* QVideoFrameInput_new3(QObject* parent);
extern __declspec(dllexport) 
QVideoFrameInput* QVideoFrameInput_new4(QVideoFrameFormat* format, QObject* parent);
extern __declspec(dllexport) 
void QVideoFrameInput_virtbase(QVideoFrameInput* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QVideoFrameInput_MetaObject(const QVideoFrameInput* self);
extern __declspec(dllexport) 
void* QVideoFrameInput_Metacast(QVideoFrameInput* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QVideoFrameInput_Tr(const char* s);
extern __declspec(dllexport) 
bool QVideoFrameInput_SendVideoFrame(QVideoFrameInput* self, QVideoFrame* frame);
extern __declspec(dllexport) 
QVideoFrameFormat* QVideoFrameInput_Format(const QVideoFrameInput* self);
extern __declspec(dllexport) 
QMediaCaptureSession* QVideoFrameInput_CaptureSession(const QVideoFrameInput* self);
extern __declspec(dllexport) 
void QVideoFrameInput_ReadyToSendVideoFrame(QVideoFrameInput* self);
void QVideoFrameInput_connect_ReadyToSendVideoFrame(QVideoFrameInput* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QVideoFrameInput_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QVideoFrameInput_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QVideoFrameInput_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QVideoFrameInput_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QVideoFrameInput_override_virtual_Metacast(void* self, intptr_t slot);
void* QVideoFrameInput_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QVideoFrameInput_Delete(QVideoFrameInput* self, bool isSubclass);

}
