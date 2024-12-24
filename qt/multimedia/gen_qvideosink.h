#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QVIDEOSINK_H
#define MIQT_QT_MULTIMEDIA_GEN_QVIDEOSINK_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSize QSize;
typedef struct QVideoFrame QVideoFrame;
typedef struct QVideoSink QVideoSink;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QVideoSink* QVideoSink_new();
extern __declspec(dllexport) 
QVideoSink* QVideoSink_new2(QObject* parent);
extern __declspec(dllexport) 
void QVideoSink_virtbase(QVideoSink* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QVideoSink_MetaObject(const QVideoSink* self);
extern __declspec(dllexport) 
void* QVideoSink_Metacast(QVideoSink* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QVideoSink_Tr(const char* s);
extern __declspec(dllexport) 
QSize* QVideoSink_VideoSize(const QVideoSink* self);
extern __declspec(dllexport) 
struct miqt_string QVideoSink_SubtitleText(const QVideoSink* self);
extern __declspec(dllexport) 
void QVideoSink_SetSubtitleText(QVideoSink* self, struct miqt_string subtitle);
extern __declspec(dllexport) 
void QVideoSink_SetVideoFrame(QVideoSink* self, QVideoFrame* frame);
extern __declspec(dllexport) 
QVideoFrame* QVideoSink_VideoFrame(const QVideoSink* self);
extern __declspec(dllexport) 
void QVideoSink_VideoFrameChanged(const QVideoSink* self, QVideoFrame* frame);
void QVideoSink_connect_VideoFrameChanged(QVideoSink* self, intptr_t slot);
extern __declspec(dllexport) 
void QVideoSink_SubtitleTextChanged(const QVideoSink* self, struct miqt_string subtitleText);
void QVideoSink_connect_SubtitleTextChanged(QVideoSink* self, intptr_t slot);
extern __declspec(dllexport) 
void QVideoSink_VideoSizeChanged(QVideoSink* self);
void QVideoSink_connect_VideoSizeChanged(QVideoSink* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QVideoSink_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QVideoSink_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QVideoSink_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QVideoSink_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QVideoSink_override_virtual_Metacast(void* self, intptr_t slot);
void* QVideoSink_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QVideoSink_Delete(QVideoSink* self, bool isSubclass);

}
