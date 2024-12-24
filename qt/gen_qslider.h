#pragma once
#ifndef MIQT_QT_GEN_QSLIDER_H
#define MIQT_QT_GEN_QSLIDER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractSlider;
class QEvent;
class QKeyEvent;
class QMetaObject;
class QMouseEvent;
class QObject;
class QPaintDevice;
class QPaintEvent;
class QSize;
class QSlider;
class QStyleOptionSlider;
class QTimerEvent;
class QWheelEvent;
class QWidget;
class _GUID;
class type_info;
#else
typedef struct QAbstractSlider QAbstractSlider;
typedef struct QEvent QEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QSize QSize;
typedef struct QSlider QSlider;
typedef struct QStyleOptionSlider QStyleOptionSlider;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSlider* QSlider_new(QWidget* parent);
extern __declspec(dllexport) QSlider* QSlider_new2();
extern __declspec(dllexport) QSlider* QSlider_new3(int orientation);
extern __declspec(dllexport) QSlider* QSlider_new4(int orientation, QWidget* parent);
extern __declspec(dllexport) void QSlider_virtbase(QSlider* src, QAbstractSlider** outptr_QAbstractSlider);
extern __declspec(dllexport) QMetaObject* QSlider_MetaObject(const QSlider* self);
extern __declspec(dllexport) void* QSlider_Metacast(QSlider* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QSlider_Tr(const char* s);
extern __declspec(dllexport) QSize* QSlider_SizeHint(const QSlider* self);
extern __declspec(dllexport) QSize* QSlider_MinimumSizeHint(const QSlider* self);
extern __declspec(dllexport) void QSlider_SetTickPosition(QSlider* self, TickPosition position);
extern __declspec(dllexport) TickPosition QSlider_TickPosition(const QSlider* self);
extern __declspec(dllexport) void QSlider_SetTickInterval(QSlider* self, int ti);
extern __declspec(dllexport) int QSlider_TickInterval(const QSlider* self);
extern __declspec(dllexport) bool QSlider_Event(QSlider* self, QEvent* event);
extern __declspec(dllexport) void QSlider_PaintEvent(QSlider* self, QPaintEvent* ev);
extern __declspec(dllexport) void QSlider_MousePressEvent(QSlider* self, QMouseEvent* ev);
extern __declspec(dllexport) void QSlider_MouseReleaseEvent(QSlider* self, QMouseEvent* ev);
extern __declspec(dllexport) void QSlider_MouseMoveEvent(QSlider* self, QMouseEvent* ev);
extern __declspec(dllexport) void QSlider_InitStyleOption(const QSlider* self, QStyleOptionSlider* option);
extern __declspec(dllexport) struct miqt_string QSlider_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QSlider_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QSlider_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QSlider_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QSlider_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QSlider_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QSlider_override_virtual_Event(void* self, intptr_t slot);
bool QSlider_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QSlider_override_virtual_PaintEvent(void* self, intptr_t slot);
void QSlider_virtualbase_PaintEvent(void* self, QPaintEvent* ev);
extern __declspec(dllexport) void QSlider_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QSlider_virtualbase_MousePressEvent(void* self, QMouseEvent* ev);
extern __declspec(dllexport) void QSlider_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QSlider_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* ev);
extern __declspec(dllexport) void QSlider_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QSlider_virtualbase_MouseMoveEvent(void* self, QMouseEvent* ev);
extern __declspec(dllexport) void QSlider_override_virtual_InitStyleOption(void* self, intptr_t slot);
void QSlider_virtualbase_InitStyleOption(const void* self, QStyleOptionSlider* option);
extern __declspec(dllexport) void QSlider_override_virtual_SliderChange(void* self, intptr_t slot);
void QSlider_virtualbase_SliderChange(void* self, SliderChange change);
extern __declspec(dllexport) void QSlider_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QSlider_virtualbase_KeyPressEvent(void* self, QKeyEvent* ev);
extern __declspec(dllexport) void QSlider_override_virtual_TimerEvent(void* self, intptr_t slot);
void QSlider_virtualbase_TimerEvent(void* self, QTimerEvent* param1);
extern __declspec(dllexport) void QSlider_override_virtual_WheelEvent(void* self, intptr_t slot);
void QSlider_virtualbase_WheelEvent(void* self, QWheelEvent* e);
extern __declspec(dllexport) void QSlider_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QSlider_virtualbase_ChangeEvent(void* self, QEvent* e);
extern __declspec(dllexport) void QSlider_Delete(QSlider* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
