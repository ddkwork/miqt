#pragma once
#ifndef MIQT_QT_GEN_QSIZEGRIP_H
#define MIQT_QT_GEN_QSIZEGRIP_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QHideEvent QHideEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QMoveEvent QMoveEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QSizeGrip QSizeGrip;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QSizeGrip* QSizeGrip_new(QWidget* parent);
extern __declspec(dllexport) 
void QSizeGrip_virtbase(QSizeGrip* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QSizeGrip_MetaObject(const QSizeGrip* self);
extern __declspec(dllexport) 
void* QSizeGrip_Metacast(QSizeGrip* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QSizeGrip_Tr(const char* s);
extern __declspec(dllexport) 
QSize* QSizeGrip_SizeHint(const QSizeGrip* self);
extern __declspec(dllexport) 
void QSizeGrip_SetVisible(QSizeGrip* self, bool visible);
extern __declspec(dllexport) 
void QSizeGrip_PaintEvent(QSizeGrip* self, QPaintEvent* param1);
extern __declspec(dllexport) 
void QSizeGrip_MousePressEvent(QSizeGrip* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QSizeGrip_MouseMoveEvent(QSizeGrip* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QSizeGrip_MouseReleaseEvent(QSizeGrip* self, QMouseEvent* mouseEvent);
extern __declspec(dllexport) 
void QSizeGrip_MoveEvent(QSizeGrip* self, QMoveEvent* moveEvent);
extern __declspec(dllexport) 
void QSizeGrip_ShowEvent(QSizeGrip* self, QShowEvent* showEvent);
extern __declspec(dllexport) 
void QSizeGrip_HideEvent(QSizeGrip* self, QHideEvent* hideEvent);
extern __declspec(dllexport) 
bool QSizeGrip_EventFilter(QSizeGrip* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) 
bool QSizeGrip_Event(QSizeGrip* self, QEvent* param1);
extern __declspec(dllexport) 
struct miqt_string QSizeGrip_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QSizeGrip_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QSizeGrip_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QSizeGrip_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QSizeGrip_override_virtual_Metacast(void* self, intptr_t slot);
void* QSizeGrip_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QSizeGrip_Delete(QSizeGrip* self, bool isSubclass);

}
