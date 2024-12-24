#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QVIDEOWIDGET_H
#define MIQT_QT_MULTIMEDIA_GEN_QVIDEOWIDGET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QHideEvent QHideEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMoveEvent QMoveEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QVideoSink QVideoSink;
typedef struct QVideoWidget QVideoWidget;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QVideoWidget* QVideoWidget_new(QWidget* parent);
extern __declspec(dllexport) 
QVideoWidget* QVideoWidget_new2();
extern __declspec(dllexport) 
void QVideoWidget_virtbase(QVideoWidget* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QVideoWidget_MetaObject(const QVideoWidget* self);
extern __declspec(dllexport) 
void* QVideoWidget_Metacast(QVideoWidget* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QVideoWidget_Tr(const char* s);
extern __declspec(dllexport) 
QVideoSink* QVideoWidget_VideoSink(const QVideoWidget* self);
extern __declspec(dllexport) 
int QVideoWidget_AspectRatioMode(const QVideoWidget* self);
extern __declspec(dllexport) 
QSize* QVideoWidget_SizeHint(const QVideoWidget* self);
extern __declspec(dllexport) 
void QVideoWidget_SetFullScreen(QVideoWidget* self, bool fullScreen);
extern __declspec(dllexport) 
void QVideoWidget_SetAspectRatioMode(QVideoWidget* self, int mode);
extern __declspec(dllexport) 
void QVideoWidget_FullScreenChanged(QVideoWidget* self, bool fullScreen);
void QVideoWidget_connect_FullScreenChanged(QVideoWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QVideoWidget_AspectRatioModeChanged(QVideoWidget* self, int mode);
void QVideoWidget_connect_AspectRatioModeChanged(QVideoWidget* self, intptr_t slot);
extern __declspec(dllexport) 
bool QVideoWidget_Event(QVideoWidget* self, QEvent* event);
extern __declspec(dllexport) 
void QVideoWidget_ShowEvent(QVideoWidget* self, QShowEvent* event);
extern __declspec(dllexport) 
void QVideoWidget_HideEvent(QVideoWidget* self, QHideEvent* event);
extern __declspec(dllexport) 
void QVideoWidget_ResizeEvent(QVideoWidget* self, QResizeEvent* event);
extern __declspec(dllexport) 
void QVideoWidget_MoveEvent(QVideoWidget* self, QMoveEvent* event);
extern __declspec(dllexport) 
struct miqt_string QVideoWidget_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QVideoWidget_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QVideoWidget_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QVideoWidget_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QVideoWidget_override_virtual_Metacast(void* self, intptr_t slot);
void* QVideoWidget_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QVideoWidget_Delete(QVideoWidget* self, bool isSubclass);

}
