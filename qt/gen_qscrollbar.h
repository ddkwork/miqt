#pragma once
#ifndef MIQT_QT_GEN_QSCROLLBAR_H
#define MIQT_QT_GEN_QSCROLLBAR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractSlider QAbstractSlider;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QEvent QEvent;
typedef struct QHideEvent QHideEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QScrollBar QScrollBar;
typedef struct QSize QSize;
typedef struct QStyleOptionSlider QStyleOptionSlider;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QScrollBar* QScrollBar_new(QWidget* parent);
extern __declspec(dllexport) 
QScrollBar* QScrollBar_new2();
extern __declspec(dllexport) 
QScrollBar* QScrollBar_new3(int param1);
extern __declspec(dllexport) 
QScrollBar* QScrollBar_new4(int param1, QWidget* parent);
extern __declspec(dllexport) 
void QScrollBar_virtbase(QScrollBar* src
, QAbstractSlider** outptr_QAbstractSlider
);
extern __declspec(dllexport) 
QMetaObject* QScrollBar_MetaObject(const QScrollBar* self);
extern __declspec(dllexport) 
void* QScrollBar_Metacast(QScrollBar* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QScrollBar_Tr(const char* s);
extern __declspec(dllexport) 
QSize* QScrollBar_SizeHint(const QScrollBar* self);
extern __declspec(dllexport) 
bool QScrollBar_Event(QScrollBar* self, QEvent* event);
extern __declspec(dllexport) 
void QScrollBar_WheelEvent(QScrollBar* self, QWheelEvent* param1);
extern __declspec(dllexport) 
void QScrollBar_PaintEvent(QScrollBar* self, QPaintEvent* param1);
extern __declspec(dllexport) 
void QScrollBar_MousePressEvent(QScrollBar* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QScrollBar_MouseReleaseEvent(QScrollBar* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QScrollBar_MouseMoveEvent(QScrollBar* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QScrollBar_HideEvent(QScrollBar* self, QHideEvent* param1);
extern __declspec(dllexport) 
void QScrollBar_SliderChange(QScrollBar* self, SliderChange change);
extern __declspec(dllexport) 
void QScrollBar_ContextMenuEvent(QScrollBar* self, QContextMenuEvent* param1);
extern __declspec(dllexport) 
void QScrollBar_InitStyleOption(const QScrollBar* self, QStyleOptionSlider* option);
extern __declspec(dllexport) 
struct miqt_string QScrollBar_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QScrollBar_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QScrollBar_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QScrollBar_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QScrollBar_override_virtual_Metacast(void* self, intptr_t slot);
void* QScrollBar_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QScrollBar_Delete(QScrollBar* self, bool isSubclass);

}
