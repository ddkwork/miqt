#pragma once
#ifndef MIQT_QT_GEN_QSLIDER_H
#define MIQT_QT_GEN_QSLIDER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractSlider QAbstractSlider;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QSize QSize;
typedef struct QSlider QSlider;
typedef struct QStyleOptionSlider QStyleOptionSlider;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QSlider* QSlider_new(QWidget* parent);
extern __declspec(dllexport) 
QSlider* QSlider_new2();
extern __declspec(dllexport) 
QSlider* QSlider_new3(int orientation);
extern __declspec(dllexport) 
QSlider* QSlider_new4(int orientation, QWidget* parent);
extern __declspec(dllexport) 
void QSlider_virtbase(QSlider* src
, QAbstractSlider** outptr_QAbstractSlider
);
extern __declspec(dllexport) 
QMetaObject* QSlider_MetaObject(const QSlider* self);
extern __declspec(dllexport) 
void* QSlider_Metacast(QSlider* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QSlider_Tr(const char* s);
extern __declspec(dllexport) 
QSize* QSlider_SizeHint(const QSlider* self);
extern __declspec(dllexport) 
QSize* QSlider_MinimumSizeHint(const QSlider* self);
extern __declspec(dllexport) 
void QSlider_SetTickPosition(QSlider* self, TickPosition position);
extern __declspec(dllexport) 
TickPosition QSlider_TickPosition(const QSlider* self);
extern __declspec(dllexport) 
void QSlider_SetTickInterval(QSlider* self, int ti);
extern __declspec(dllexport) 
int QSlider_TickInterval(const QSlider* self);
extern __declspec(dllexport) 
bool QSlider_Event(QSlider* self, QEvent* event);
extern __declspec(dllexport) 
void QSlider_PaintEvent(QSlider* self, QPaintEvent* ev);
extern __declspec(dllexport) 
void QSlider_MousePressEvent(QSlider* self, QMouseEvent* ev);
extern __declspec(dllexport) 
void QSlider_MouseReleaseEvent(QSlider* self, QMouseEvent* ev);
extern __declspec(dllexport) 
void QSlider_MouseMoveEvent(QSlider* self, QMouseEvent* ev);
extern __declspec(dllexport) 
void QSlider_InitStyleOption(const QSlider* self, QStyleOptionSlider* option);
extern __declspec(dllexport) 
struct miqt_string QSlider_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QSlider_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QSlider_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QSlider_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QSlider_override_virtual_Metacast(void* self, intptr_t slot);
void* QSlider_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QSlider_Delete(QSlider* self, bool isSubclass);

}
