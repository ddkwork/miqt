#pragma once
#ifndef MIQT_QT_GEN_QCHECKBOX_H
#define MIQT_QT_GEN_QCHECKBOX_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractButton QAbstractButton;
typedef struct QCheckBox QCheckBox;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPoint QPoint;
typedef struct QSize QSize;
typedef struct QStyleOptionButton QStyleOptionButton;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QCheckBox* QCheckBox_new(QWidget* parent);
extern __declspec(dllexport) QCheckBox* QCheckBox_new2();
extern __declspec(dllexport) QCheckBox* QCheckBox_new3(struct miqt_string text);
extern __declspec(dllexport) QCheckBox* QCheckBox_new4(struct miqt_string text, QWidget* parent);
extern __declspec(dllexport) void QCheckBox_virtbase(QCheckBox* src, QAbstractButton** outptr_QAbstractButton);
extern __declspec(dllexport) QMetaObject* QCheckBox_MetaObject(const QCheckBox* self);
extern __declspec(dllexport) void* QCheckBox_Metacast(QCheckBox* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QCheckBox_Tr(const char* s);
extern __declspec(dllexport) QSize* QCheckBox_SizeHint(const QCheckBox* self);
extern __declspec(dllexport) QSize* QCheckBox_MinimumSizeHint(const QCheckBox* self);
extern __declspec(dllexport) void QCheckBox_SetTristate(QCheckBox* self);
extern __declspec(dllexport) bool QCheckBox_IsTristate(const QCheckBox* self);
extern __declspec(dllexport) int QCheckBox_CheckState(const QCheckBox* self);
extern __declspec(dllexport) void QCheckBox_SetCheckState(QCheckBox* self, int state);
extern __declspec(dllexport) void QCheckBox_StateChanged(QCheckBox* self, int param1);
void QCheckBox_connect_StateChanged(QCheckBox* self, intptr_t slot);
extern __declspec(dllexport) void QCheckBox_CheckStateChanged(QCheckBox* self, int param1);
void QCheckBox_connect_CheckStateChanged(QCheckBox* self, intptr_t slot);
extern __declspec(dllexport) bool QCheckBox_Event(QCheckBox* self, QEvent* e);
extern __declspec(dllexport) bool QCheckBox_HitButton(const QCheckBox* self, QPoint* pos);
extern __declspec(dllexport) void QCheckBox_CheckStateSet(QCheckBox* self);
extern __declspec(dllexport) void QCheckBox_NextCheckState(QCheckBox* self);
extern __declspec(dllexport) void QCheckBox_PaintEvent(QCheckBox* self, QPaintEvent* param1);
extern __declspec(dllexport) void QCheckBox_MouseMoveEvent(QCheckBox* self, QMouseEvent* param1);
extern __declspec(dllexport) void QCheckBox_InitStyleOption(const QCheckBox* self, QStyleOptionButton* option);
extern __declspec(dllexport) struct miqt_string QCheckBox_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QCheckBox_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QCheckBox_SetTristate1(QCheckBox* self, bool y);
extern __declspec(dllexport) void QCheckBox_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QCheckBox_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QCheckBox_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QCheckBox_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QCheckBox_override_virtual_Event(void* self, intptr_t slot);
bool QCheckBox_virtualbase_Event(void* self, QEvent* e);
extern __declspec(dllexport) void QCheckBox_override_virtual_HitButton(void* self, intptr_t slot);
bool QCheckBox_virtualbase_HitButton(const void* self, QPoint* pos);
extern __declspec(dllexport) void QCheckBox_override_virtual_CheckStateSet(void* self, intptr_t slot);
void QCheckBox_virtualbase_CheckStateSet(void* self);
extern __declspec(dllexport) void QCheckBox_override_virtual_NextCheckState(void* self, intptr_t slot);
void QCheckBox_virtualbase_NextCheckState(void* self);
extern __declspec(dllexport) void QCheckBox_override_virtual_PaintEvent(void* self, intptr_t slot);
void QCheckBox_virtualbase_PaintEvent(void* self, QPaintEvent* param1);
extern __declspec(dllexport) void QCheckBox_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QCheckBox_virtualbase_MouseMoveEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QCheckBox_override_virtual_InitStyleOption(void* self, intptr_t slot);
void QCheckBox_virtualbase_InitStyleOption(const void* self, QStyleOptionButton* option);
extern __declspec(dllexport) void QCheckBox_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QCheckBox_virtualbase_KeyPressEvent(void* self, QKeyEvent* e);
extern __declspec(dllexport) void QCheckBox_override_virtual_KeyReleaseEvent(void* self, intptr_t slot);
void QCheckBox_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* e);
extern __declspec(dllexport) void QCheckBox_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QCheckBox_virtualbase_MousePressEvent(void* self, QMouseEvent* e);
extern __declspec(dllexport) void QCheckBox_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QCheckBox_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* e);
extern __declspec(dllexport) void QCheckBox_override_virtual_FocusInEvent(void* self, intptr_t slot);
void QCheckBox_virtualbase_FocusInEvent(void* self, QFocusEvent* e);
extern __declspec(dllexport) void QCheckBox_override_virtual_FocusOutEvent(void* self, intptr_t slot);
void QCheckBox_virtualbase_FocusOutEvent(void* self, QFocusEvent* e);
extern __declspec(dllexport) void QCheckBox_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QCheckBox_virtualbase_ChangeEvent(void* self, QEvent* e);
extern __declspec(dllexport) void QCheckBox_override_virtual_TimerEvent(void* self, intptr_t slot);
void QCheckBox_virtualbase_TimerEvent(void* self, QTimerEvent* e);
extern __declspec(dllexport) void QCheckBox_Delete(QCheckBox* self, bool isSubclass);

} 
