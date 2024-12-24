#pragma once
#ifndef MIQT_QT_GEN_QRHIWIDGET_H
#define MIQT_QT_GEN_QRHIWIDGET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QImage QImage;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QResizeEvent QResizeEvent;
typedef struct QRhiWidget QRhiWidget;
typedef struct QSize QSize;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QRhiWidget* QRhiWidget_new(QWidget* parent);
extern __declspec(dllexport) 
QRhiWidget* QRhiWidget_new2();
extern __declspec(dllexport) 
QRhiWidget* QRhiWidget_new3(QWidget* parent, int f);
extern __declspec(dllexport) 
void QRhiWidget_virtbase(QRhiWidget* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QRhiWidget_MetaObject(const QRhiWidget* self);
extern __declspec(dllexport) 
void* QRhiWidget_Metacast(QRhiWidget* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QRhiWidget_Tr(const char* s);
extern __declspec(dllexport) 
Api QRhiWidget_Api(const QRhiWidget* self);
extern __declspec(dllexport) 
void QRhiWidget_SetApi(QRhiWidget* self, Api api);
extern __declspec(dllexport) 
bool QRhiWidget_IsDebugLayerEnabled(const QRhiWidget* self);
extern __declspec(dllexport) 
void QRhiWidget_SetDebugLayerEnabled(QRhiWidget* self, bool enable);
extern __declspec(dllexport) 
int QRhiWidget_SampleCount(const QRhiWidget* self);
extern __declspec(dllexport) 
void QRhiWidget_SetSampleCount(QRhiWidget* self, int samples);
extern __declspec(dllexport) 
TextureFormat QRhiWidget_ColorBufferFormat(const QRhiWidget* self);
extern __declspec(dllexport) 
void QRhiWidget_SetColorBufferFormat(QRhiWidget* self, TextureFormat format);
extern __declspec(dllexport) 
QSize* QRhiWidget_FixedColorBufferSize(const QRhiWidget* self);
extern __declspec(dllexport) 
void QRhiWidget_SetFixedColorBufferSize(QRhiWidget* self, QSize* pixelSize);
extern __declspec(dllexport) 
void QRhiWidget_SetFixedColorBufferSize2(QRhiWidget* self, int w, int h);
extern __declspec(dllexport) 
bool QRhiWidget_IsMirrorVerticallyEnabled(const QRhiWidget* self);
extern __declspec(dllexport) 
void QRhiWidget_SetMirrorVertically(QRhiWidget* self, bool enabled);
extern __declspec(dllexport) 
QImage* QRhiWidget_GrabFramebuffer(const QRhiWidget* self);
extern __declspec(dllexport) 
void QRhiWidget_Initialize(QRhiWidget* self, QRhiCommandBuffer* cb);
extern __declspec(dllexport) 
void QRhiWidget_Render(QRhiWidget* self, QRhiCommandBuffer* cb);
extern __declspec(dllexport) 
void QRhiWidget_ReleaseResources(QRhiWidget* self);
extern __declspec(dllexport) 
void QRhiWidget_ResizeEvent(QRhiWidget* self, QResizeEvent* e);
extern __declspec(dllexport) 
void QRhiWidget_PaintEvent(QRhiWidget* self, QPaintEvent* e);
extern __declspec(dllexport) 
bool QRhiWidget_Event(QRhiWidget* self, QEvent* e);
extern __declspec(dllexport) 
void QRhiWidget_FrameSubmitted(QRhiWidget* self);
void QRhiWidget_connect_FrameSubmitted(QRhiWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QRhiWidget_RenderFailed(QRhiWidget* self);
void QRhiWidget_connect_RenderFailed(QRhiWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QRhiWidget_SampleCountChanged(QRhiWidget* self, int samples);
void QRhiWidget_connect_SampleCountChanged(QRhiWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QRhiWidget_ColorBufferFormatChanged(QRhiWidget* self, TextureFormat format);
void QRhiWidget_connect_ColorBufferFormatChanged(QRhiWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QRhiWidget_FixedColorBufferSizeChanged(QRhiWidget* self, QSize* pixelSize);
void QRhiWidget_connect_FixedColorBufferSizeChanged(QRhiWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QRhiWidget_MirrorVerticallyChanged(QRhiWidget* self, bool enabled);
void QRhiWidget_connect_MirrorVerticallyChanged(QRhiWidget* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QRhiWidget_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QRhiWidget_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QRhiWidget_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QRhiWidget_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QRhiWidget_override_virtual_Metacast(void* self, intptr_t slot);
void* QRhiWidget_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QRhiWidget_Delete(QRhiWidget* self, bool isSubclass);

}
