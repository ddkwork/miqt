#pragma once
#ifndef MIQT_QT_GEN_QMDISUBWINDOW_H
#define MIQT_QT_GEN_QMDISUBWINDOW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QCloseEvent QCloseEvent;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QHideEvent QHideEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMdiArea QMdiArea;
typedef struct QMdiSubWindow QMdiSubWindow;
typedef struct QMenu QMenu;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QMoveEvent QMoveEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QMdiSubWindow* QMdiSubWindow_new(QWidget* parent);
extern __declspec(dllexport) 
QMdiSubWindow* QMdiSubWindow_new2();
extern __declspec(dllexport) 
QMdiSubWindow* QMdiSubWindow_new3(QWidget* parent, int flags);
extern __declspec(dllexport) 
void QMdiSubWindow_virtbase(QMdiSubWindow* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QMdiSubWindow_MetaObject(const QMdiSubWindow* self);
extern __declspec(dllexport) 
void* QMdiSubWindow_Metacast(QMdiSubWindow* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QMdiSubWindow_Tr(const char* s);
extern __declspec(dllexport) 
QSize* QMdiSubWindow_SizeHint(const QMdiSubWindow* self);
extern __declspec(dllexport) 
QSize* QMdiSubWindow_MinimumSizeHint(const QMdiSubWindow* self);
extern __declspec(dllexport) 
void QMdiSubWindow_SetWidget(QMdiSubWindow* self, QWidget* widget);
extern __declspec(dllexport) 
QWidget* QMdiSubWindow_Widget(const QMdiSubWindow* self);
extern __declspec(dllexport) 
QWidget* QMdiSubWindow_MaximizedButtonsWidget(const QMdiSubWindow* self);
extern __declspec(dllexport) 
QWidget* QMdiSubWindow_MaximizedSystemMenuIconWidget(const QMdiSubWindow* self);
extern __declspec(dllexport) 
bool QMdiSubWindow_IsShaded(const QMdiSubWindow* self);
extern __declspec(dllexport) 
void QMdiSubWindow_SetOption(QMdiSubWindow* self, SubWindowOption option);
extern __declspec(dllexport) 
bool QMdiSubWindow_TestOption(const QMdiSubWindow* self, SubWindowOption param1);
extern __declspec(dllexport) 
void QMdiSubWindow_SetKeyboardSingleStep(QMdiSubWindow* self, int step);
extern __declspec(dllexport) 
int QMdiSubWindow_KeyboardSingleStep(const QMdiSubWindow* self);
extern __declspec(dllexport) 
void QMdiSubWindow_SetKeyboardPageStep(QMdiSubWindow* self, int step);
extern __declspec(dllexport) 
int QMdiSubWindow_KeyboardPageStep(const QMdiSubWindow* self);
extern __declspec(dllexport) 
void QMdiSubWindow_SetSystemMenu(QMdiSubWindow* self, QMenu* systemMenu);
extern __declspec(dllexport) 
QMenu* QMdiSubWindow_SystemMenu(const QMdiSubWindow* self);
extern __declspec(dllexport) 
QMdiArea* QMdiSubWindow_MdiArea(const QMdiSubWindow* self);
extern __declspec(dllexport) 
void QMdiSubWindow_WindowStateChanged(QMdiSubWindow* self, int oldState, int newState);
void QMdiSubWindow_connect_WindowStateChanged(QMdiSubWindow* self, intptr_t slot);
extern __declspec(dllexport) 
void QMdiSubWindow_AboutToActivate(QMdiSubWindow* self);
void QMdiSubWindow_connect_AboutToActivate(QMdiSubWindow* self, intptr_t slot);
extern __declspec(dllexport) 
void QMdiSubWindow_ShowSystemMenu(QMdiSubWindow* self);
extern __declspec(dllexport) 
void QMdiSubWindow_ShowShaded(QMdiSubWindow* self);
extern __declspec(dllexport) 
bool QMdiSubWindow_EventFilter(QMdiSubWindow* self, QObject* object, QEvent* event);
extern __declspec(dllexport) 
bool QMdiSubWindow_Event(QMdiSubWindow* self, QEvent* event);
extern __declspec(dllexport) 
void QMdiSubWindow_ShowEvent(QMdiSubWindow* self, QShowEvent* showEvent);
extern __declspec(dllexport) 
void QMdiSubWindow_HideEvent(QMdiSubWindow* self, QHideEvent* hideEvent);
extern __declspec(dllexport) 
void QMdiSubWindow_ChangeEvent(QMdiSubWindow* self, QEvent* changeEvent);
extern __declspec(dllexport) 
void QMdiSubWindow_CloseEvent(QMdiSubWindow* self, QCloseEvent* closeEvent);
extern __declspec(dllexport) 
void QMdiSubWindow_LeaveEvent(QMdiSubWindow* self, QEvent* leaveEvent);
extern __declspec(dllexport) 
void QMdiSubWindow_ResizeEvent(QMdiSubWindow* self, QResizeEvent* resizeEvent);
extern __declspec(dllexport) 
void QMdiSubWindow_TimerEvent(QMdiSubWindow* self, QTimerEvent* timerEvent);
extern __declspec(dllexport) 
void QMdiSubWindow_MoveEvent(QMdiSubWindow* self, QMoveEvent* moveEvent);
extern __declspec(dllexport) 
void QMdiSubWindow_PaintEvent(QMdiSubWindow* self, QPaintEvent* paintEvent);
extern __declspec(dllexport) 
void QMdiSubWindow_MousePressEvent(QMdiSubWindow* self, QMouseEvent* mouseEvent);
extern __declspec(dllexport) 
void QMdiSubWindow_MouseDoubleClickEvent(QMdiSubWindow* self, QMouseEvent* mouseEvent);
extern __declspec(dllexport) 
void QMdiSubWindow_MouseReleaseEvent(QMdiSubWindow* self, QMouseEvent* mouseEvent);
extern __declspec(dllexport) 
void QMdiSubWindow_MouseMoveEvent(QMdiSubWindow* self, QMouseEvent* mouseEvent);
extern __declspec(dllexport) 
void QMdiSubWindow_KeyPressEvent(QMdiSubWindow* self, QKeyEvent* keyEvent);
extern __declspec(dllexport) 
void QMdiSubWindow_ContextMenuEvent(QMdiSubWindow* self, QContextMenuEvent* contextMenuEvent);
extern __declspec(dllexport) 
void QMdiSubWindow_FocusInEvent(QMdiSubWindow* self, QFocusEvent* focusInEvent);
extern __declspec(dllexport) 
void QMdiSubWindow_FocusOutEvent(QMdiSubWindow* self, QFocusEvent* focusOutEvent);
extern __declspec(dllexport) 
void QMdiSubWindow_ChildEvent(QMdiSubWindow* self, QChildEvent* childEvent);
extern __declspec(dllexport) 
struct miqt_string QMdiSubWindow_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QMdiSubWindow_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QMdiSubWindow_SetOption2(QMdiSubWindow* self, SubWindowOption option, bool on);
extern __declspec(dllexport) 
void QMdiSubWindow_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QMdiSubWindow_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QMdiSubWindow_override_virtual_Metacast(void* self, intptr_t slot);
void* QMdiSubWindow_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QMdiSubWindow_Delete(QMdiSubWindow* self, bool isSubclass);

}
