#pragma once
#ifndef MIQT_QT_GEN_QRADIOBUTTON_H
#define MIQT_QT_GEN_QRADIOBUTTON_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractButton QAbstractButton;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPoint QPoint;
typedef struct QRadioButton QRadioButton;
typedef struct QSize QSize;
typedef struct QStyleOptionButton QStyleOptionButton;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QRadioButton* QRadioButton_new(QWidget* parent);
extern __declspec(dllexport) QRadioButton* QRadioButton_new2();
extern __declspec(dllexport) QRadioButton* QRadioButton_new3(struct miqt_string text);
extern __declspec(dllexport) QRadioButton* QRadioButton_new4(struct miqt_string text, QWidget* parent);
extern __declspec(dllexport) void QRadioButton_virtbase(QRadioButton* src, QAbstractButton** outptr_QAbstractButton);
extern __declspec(dllexport) QMetaObject* QRadioButton_MetaObject(const QRadioButton* self);
extern __declspec(dllexport) void* QRadioButton_Metacast(QRadioButton* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QRadioButton_Tr(const char* s);
extern __declspec(dllexport) QSize* QRadioButton_SizeHint(const QRadioButton* self);
extern __declspec(dllexport) QSize* QRadioButton_MinimumSizeHint(const QRadioButton* self);
extern __declspec(dllexport) bool QRadioButton_Event(QRadioButton* self, QEvent* e);
extern __declspec(dllexport) bool QRadioButton_HitButton(const QRadioButton* self, QPoint* param1);
extern __declspec(dllexport) void QRadioButton_PaintEvent(QRadioButton* self, QPaintEvent* param1);
extern __declspec(dllexport) void QRadioButton_MouseMoveEvent(QRadioButton* self, QMouseEvent* param1);
extern __declspec(dllexport) void QRadioButton_InitStyleOption(const QRadioButton* self, QStyleOptionButton* button);
extern __declspec(dllexport) struct miqt_string QRadioButton_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QRadioButton_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QRadioButton_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QRadioButton_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QRadioButton_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QRadioButton_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QRadioButton_override_virtual_Event(void* self, intptr_t slot);
bool QRadioButton_virtualbase_Event(void* self, QEvent* e);
extern __declspec(dllexport) void QRadioButton_override_virtual_HitButton(void* self, intptr_t slot);
bool QRadioButton_virtualbase_HitButton(const void* self, QPoint* param1);
extern __declspec(dllexport) void QRadioButton_override_virtual_PaintEvent(void* self, intptr_t slot);
void QRadioButton_virtualbase_PaintEvent(void* self, QPaintEvent* param1);
extern __declspec(dllexport) void QRadioButton_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QRadioButton_virtualbase_MouseMoveEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QRadioButton_override_virtual_InitStyleOption(void* self, intptr_t slot);
void QRadioButton_virtualbase_InitStyleOption(const void* self, QStyleOptionButton* button);
extern __declspec(dllexport) void QRadioButton_override_virtual_CheckStateSet(void* self, intptr_t slot);
void QRadioButton_virtualbase_CheckStateSet(void* self);
extern __declspec(dllexport) void QRadioButton_override_virtual_NextCheckState(void* self, intptr_t slot);
void QRadioButton_virtualbase_NextCheckState(void* self);
extern __declspec(dllexport) void QRadioButton_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QRadioButton_virtualbase_KeyPressEvent(void* self, QKeyEvent* e);
extern __declspec(dllexport) void QRadioButton_override_virtual_KeyReleaseEvent(void* self, intptr_t slot);
void QRadioButton_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* e);
extern __declspec(dllexport) void QRadioButton_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QRadioButton_virtualbase_MousePressEvent(void* self, QMouseEvent* e);
extern __declspec(dllexport) void QRadioButton_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QRadioButton_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* e);
extern __declspec(dllexport) void QRadioButton_override_virtual_FocusInEvent(void* self, intptr_t slot);
void QRadioButton_virtualbase_FocusInEvent(void* self, QFocusEvent* e);
extern __declspec(dllexport) void QRadioButton_override_virtual_FocusOutEvent(void* self, intptr_t slot);
void QRadioButton_virtualbase_FocusOutEvent(void* self, QFocusEvent* e);
extern __declspec(dllexport) void QRadioButton_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QRadioButton_virtualbase_ChangeEvent(void* self, QEvent* e);
extern __declspec(dllexport) void QRadioButton_override_virtual_TimerEvent(void* self, intptr_t slot);
void QRadioButton_virtualbase_TimerEvent(void* self, QTimerEvent* e);
extern __declspec(dllexport) void QRadioButton_Delete(QRadioButton* self, bool isSubclass);

} 
