#pragma once
#ifndef MIQT_QT_GEN_QTOOLBUTTON_H
#define MIQT_QT_GEN_QTOOLBUTTON_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractButton QAbstractButton;
typedef struct QAction QAction;
typedef struct QActionEvent QActionEvent;
typedef struct QEnterEvent QEnterEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMenu QMenu;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPoint QPoint;
typedef struct QSize QSize;
typedef struct QStyleOptionToolButton QStyleOptionToolButton;
typedef struct QTimerEvent QTimerEvent;
typedef struct QToolButton QToolButton;
typedef struct QWidget QWidget;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QToolButton* QToolButton_new(QWidget* parent);
extern __declspec(dllexport) QToolButton* QToolButton_new2();
extern __declspec(dllexport) void QToolButton_virtbase(QToolButton* src, QAbstractButton** outptr_QAbstractButton);
extern __declspec(dllexport) QMetaObject* QToolButton_MetaObject(const QToolButton* self);
extern __declspec(dllexport) void* QToolButton_Metacast(QToolButton* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QToolButton_Tr(const char* s);
extern __declspec(dllexport) QSize* QToolButton_SizeHint(const QToolButton* self);
extern __declspec(dllexport) QSize* QToolButton_MinimumSizeHint(const QToolButton* self);
extern __declspec(dllexport) int QToolButton_ToolButtonStyle(const QToolButton* self);
extern __declspec(dllexport) int QToolButton_ArrowType(const QToolButton* self);
extern __declspec(dllexport) void QToolButton_SetArrowType(QToolButton* self, int typeVal);
extern __declspec(dllexport) void QToolButton_SetMenu(QToolButton* self, QMenu* menu);
extern __declspec(dllexport) QMenu* QToolButton_Menu(const QToolButton* self);
extern __declspec(dllexport) void QToolButton_SetPopupMode(QToolButton* self, ToolButtonPopupMode mode);
extern __declspec(dllexport) ToolButtonPopupMode QToolButton_PopupMode(const QToolButton* self);
extern __declspec(dllexport) QAction* QToolButton_DefaultAction(const QToolButton* self);
extern __declspec(dllexport) void QToolButton_SetAutoRaise(QToolButton* self, bool enable);
extern __declspec(dllexport) bool QToolButton_AutoRaise(const QToolButton* self);
extern __declspec(dllexport) void QToolButton_ShowMenu(QToolButton* self);
extern __declspec(dllexport) void QToolButton_SetToolButtonStyle(QToolButton* self, int style);
extern __declspec(dllexport) void QToolButton_SetDefaultAction(QToolButton* self, QAction* defaultAction);
extern __declspec(dllexport) void QToolButton_Triggered(QToolButton* self, QAction* param1);
void QToolButton_connect_Triggered(QToolButton* self, intptr_t slot);
extern __declspec(dllexport) bool QToolButton_Event(QToolButton* self, QEvent* e);
extern __declspec(dllexport) void QToolButton_MousePressEvent(QToolButton* self, QMouseEvent* param1);
extern __declspec(dllexport) void QToolButton_MouseReleaseEvent(QToolButton* self, QMouseEvent* param1);
extern __declspec(dllexport) void QToolButton_PaintEvent(QToolButton* self, QPaintEvent* param1);
extern __declspec(dllexport) void QToolButton_ActionEvent(QToolButton* self, QActionEvent* param1);
extern __declspec(dllexport) void QToolButton_EnterEvent(QToolButton* self, QEnterEvent* param1);
extern __declspec(dllexport) void QToolButton_LeaveEvent(QToolButton* self, QEvent* param1);
extern __declspec(dllexport) void QToolButton_TimerEvent(QToolButton* self, QTimerEvent* param1);
extern __declspec(dllexport) void QToolButton_ChangeEvent(QToolButton* self, QEvent* param1);
extern __declspec(dllexport) bool QToolButton_HitButton(const QToolButton* self, QPoint* pos);
extern __declspec(dllexport) void QToolButton_CheckStateSet(QToolButton* self);
extern __declspec(dllexport) void QToolButton_NextCheckState(QToolButton* self);
extern __declspec(dllexport) void QToolButton_InitStyleOption(const QToolButton* self, QStyleOptionToolButton* option);
extern __declspec(dllexport) struct miqt_string QToolButton_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QToolButton_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QToolButton_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QToolButton_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QToolButton_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QToolButton_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QToolButton_override_virtual_Event(void* self, intptr_t slot);
bool QToolButton_virtualbase_Event(void* self, QEvent* e);
extern __declspec(dllexport) void QToolButton_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QToolButton_virtualbase_MousePressEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QToolButton_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QToolButton_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QToolButton_override_virtual_PaintEvent(void* self, intptr_t slot);
void QToolButton_virtualbase_PaintEvent(void* self, QPaintEvent* param1);
extern __declspec(dllexport) void QToolButton_override_virtual_ActionEvent(void* self, intptr_t slot);
void QToolButton_virtualbase_ActionEvent(void* self, QActionEvent* param1);
extern __declspec(dllexport) void QToolButton_override_virtual_EnterEvent(void* self, intptr_t slot);
void QToolButton_virtualbase_EnterEvent(void* self, QEnterEvent* param1);
extern __declspec(dllexport) void QToolButton_override_virtual_LeaveEvent(void* self, intptr_t slot);
void QToolButton_virtualbase_LeaveEvent(void* self, QEvent* param1);
extern __declspec(dllexport) void QToolButton_override_virtual_TimerEvent(void* self, intptr_t slot);
void QToolButton_virtualbase_TimerEvent(void* self, QTimerEvent* param1);
extern __declspec(dllexport) void QToolButton_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QToolButton_virtualbase_ChangeEvent(void* self, QEvent* param1);
extern __declspec(dllexport) void QToolButton_override_virtual_HitButton(void* self, intptr_t slot);
bool QToolButton_virtualbase_HitButton(const void* self, QPoint* pos);
extern __declspec(dllexport) void QToolButton_override_virtual_CheckStateSet(void* self, intptr_t slot);
void QToolButton_virtualbase_CheckStateSet(void* self);
extern __declspec(dllexport) void QToolButton_override_virtual_NextCheckState(void* self, intptr_t slot);
void QToolButton_virtualbase_NextCheckState(void* self);
extern __declspec(dllexport) void QToolButton_override_virtual_InitStyleOption(void* self, intptr_t slot);
void QToolButton_virtualbase_InitStyleOption(const void* self, QStyleOptionToolButton* option);
extern __declspec(dllexport) void QToolButton_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QToolButton_virtualbase_KeyPressEvent(void* self, QKeyEvent* e);
extern __declspec(dllexport) void QToolButton_override_virtual_KeyReleaseEvent(void* self, intptr_t slot);
void QToolButton_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* e);
extern __declspec(dllexport) void QToolButton_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QToolButton_virtualbase_MouseMoveEvent(void* self, QMouseEvent* e);
extern __declspec(dllexport) void QToolButton_override_virtual_FocusInEvent(void* self, intptr_t slot);
void QToolButton_virtualbase_FocusInEvent(void* self, QFocusEvent* e);
extern __declspec(dllexport) void QToolButton_override_virtual_FocusOutEvent(void* self, intptr_t slot);
void QToolButton_virtualbase_FocusOutEvent(void* self, QFocusEvent* e);
extern __declspec(dllexport) void QToolButton_Delete(QToolButton* self, bool isSubclass);

} 
