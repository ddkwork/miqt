#pragma once
#ifndef MIQT_QT_GEN_QDIAL_H
#define MIQT_QT_GEN_QDIAL_H

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
class QDial;
class QEvent;
class QKeyEvent;
class QMetaObject;
class QMouseEvent;
class QObject;
class QPaintDevice;
class QPaintEvent;
class QResizeEvent;
class QSize;
class QStyleOptionSlider;
class QTimerEvent;
class QWheelEvent;
class QWidget;
class _GUID;
class type_info;
#else
typedef struct QAbstractSlider QAbstractSlider;
typedef struct QDial QDial;
typedef struct QEvent QEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QResizeEvent QResizeEvent;
typedef struct QSize QSize;
typedef struct QStyleOptionSlider QStyleOptionSlider;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QDial* QDial_new(QWidget* parent);
extern __declspec(dllexport) QDial* QDial_new2();
extern __declspec(dllexport) void QDial_virtbase(QDial* src, QAbstractSlider** outptr_QAbstractSlider);
extern __declspec(dllexport) QMetaObject* QDial_MetaObject(const QDial* self);
extern __declspec(dllexport) void* QDial_Metacast(QDial* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QDial_Tr(const char* s);
extern __declspec(dllexport) bool QDial_Wrapping(const QDial* self);
extern __declspec(dllexport) int QDial_NotchSize(const QDial* self);
extern __declspec(dllexport) void QDial_SetNotchTarget(QDial* self, double target);
extern __declspec(dllexport) double QDial_NotchTarget(const QDial* self);
extern __declspec(dllexport) bool QDial_NotchesVisible(const QDial* self);
extern __declspec(dllexport) QSize* QDial_SizeHint(const QDial* self);
extern __declspec(dllexport) QSize* QDial_MinimumSizeHint(const QDial* self);
extern __declspec(dllexport) void QDial_SetNotchesVisible(QDial* self, bool visible);
extern __declspec(dllexport) void QDial_SetWrapping(QDial* self, bool on);
extern __declspec(dllexport) bool QDial_Event(QDial* self, QEvent* e);
extern __declspec(dllexport) void QDial_ResizeEvent(QDial* self, QResizeEvent* re);
extern __declspec(dllexport) void QDial_PaintEvent(QDial* self, QPaintEvent* pe);
extern __declspec(dllexport) void QDial_MousePressEvent(QDial* self, QMouseEvent* me);
extern __declspec(dllexport) void QDial_MouseReleaseEvent(QDial* self, QMouseEvent* me);
extern __declspec(dllexport) void QDial_MouseMoveEvent(QDial* self, QMouseEvent* me);
extern __declspec(dllexport) void QDial_SliderChange(QDial* self, SliderChange change);
extern __declspec(dllexport) void QDial_InitStyleOption(const QDial* self, QStyleOptionSlider* option);
extern __declspec(dllexport) struct miqt_string QDial_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QDial_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QDial_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QDial_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QDial_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QDial_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QDial_override_virtual_Event(void* self, intptr_t slot);
bool QDial_virtualbase_Event(void* self, QEvent* e);
extern __declspec(dllexport) void QDial_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QDial_virtualbase_ResizeEvent(void* self, QResizeEvent* re);
extern __declspec(dllexport) void QDial_override_virtual_PaintEvent(void* self, intptr_t slot);
void QDial_virtualbase_PaintEvent(void* self, QPaintEvent* pe);
extern __declspec(dllexport) void QDial_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QDial_virtualbase_MousePressEvent(void* self, QMouseEvent* me);
extern __declspec(dllexport) void QDial_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QDial_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* me);
extern __declspec(dllexport) void QDial_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QDial_virtualbase_MouseMoveEvent(void* self, QMouseEvent* me);
extern __declspec(dllexport) void QDial_override_virtual_SliderChange(void* self, intptr_t slot);
void QDial_virtualbase_SliderChange(void* self, SliderChange change);
extern __declspec(dllexport) void QDial_override_virtual_InitStyleOption(void* self, intptr_t slot);
void QDial_virtualbase_InitStyleOption(const void* self, QStyleOptionSlider* option);
extern __declspec(dllexport) void QDial_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QDial_virtualbase_KeyPressEvent(void* self, QKeyEvent* ev);
extern __declspec(dllexport) void QDial_override_virtual_TimerEvent(void* self, intptr_t slot);
void QDial_virtualbase_TimerEvent(void* self, QTimerEvent* param1);
extern __declspec(dllexport) void QDial_override_virtual_WheelEvent(void* self, intptr_t slot);
void QDial_virtualbase_WheelEvent(void* self, QWheelEvent* e);
extern __declspec(dllexport) void QDial_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QDial_virtualbase_ChangeEvent(void* self, QEvent* e);
extern __declspec(dllexport) void QDial_Delete(QDial* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
