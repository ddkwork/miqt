#pragma once
#ifndef MIQT_QT_GEN_QMENUBAR_H
#define MIQT_QT_GEN_QMENUBAR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAction QAction;
typedef struct QActionEvent QActionEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QIcon QIcon;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMenu QMenu;
typedef struct QMenuBar QMenuBar;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QResizeEvent QResizeEvent;
typedef struct QSize QSize;
typedef struct QStyleOptionMenuItem QStyleOptionMenuItem;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QMenuBar* QMenuBar_new(QWidget* parent);
extern __declspec(dllexport) 
QMenuBar* QMenuBar_new2();
extern __declspec(dllexport) 
void QMenuBar_virtbase(QMenuBar* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QMenuBar_MetaObject(const QMenuBar* self);
extern __declspec(dllexport) 
void* QMenuBar_Metacast(QMenuBar* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QMenuBar_Tr(const char* s);
extern __declspec(dllexport) 
QAction* QMenuBar_AddMenu(QMenuBar* self, QMenu* menu);
extern __declspec(dllexport) 
QMenu* QMenuBar_AddMenuWithTitle(QMenuBar* self, struct miqt_string title);
extern __declspec(dllexport) 
QMenu* QMenuBar_AddMenu2(QMenuBar* self, QIcon* icon, struct miqt_string title);
extern __declspec(dllexport) 
QAction* QMenuBar_AddSeparator(QMenuBar* self);
extern __declspec(dllexport) 
QAction* QMenuBar_InsertSeparator(QMenuBar* self, QAction* before);
extern __declspec(dllexport) 
QAction* QMenuBar_InsertMenu(QMenuBar* self, QAction* before, QMenu* menu);
extern __declspec(dllexport) 
void QMenuBar_Clear(QMenuBar* self);
extern __declspec(dllexport) 
QAction* QMenuBar_ActiveAction(const QMenuBar* self);
extern __declspec(dllexport) 
void QMenuBar_SetActiveAction(QMenuBar* self, QAction* action);
extern __declspec(dllexport) 
void QMenuBar_SetDefaultUp(QMenuBar* self, bool defaultUp);
extern __declspec(dllexport) 
bool QMenuBar_IsDefaultUp(const QMenuBar* self);
extern __declspec(dllexport) 
QSize* QMenuBar_SizeHint(const QMenuBar* self);
extern __declspec(dllexport) 
QSize* QMenuBar_MinimumSizeHint(const QMenuBar* self);
extern __declspec(dllexport) 
int QMenuBar_HeightForWidth(const QMenuBar* self, int param1);
extern __declspec(dllexport) 
QRect* QMenuBar_ActionGeometry(const QMenuBar* self, QAction* param1);
extern __declspec(dllexport) 
QAction* QMenuBar_ActionAt(const QMenuBar* self, QPoint* param1);
extern __declspec(dllexport) 
void QMenuBar_SetCornerWidget(QMenuBar* self, QWidget* w);
extern __declspec(dllexport) 
QWidget* QMenuBar_CornerWidget(const QMenuBar* self);
extern __declspec(dllexport) 
bool QMenuBar_IsNativeMenuBar(const QMenuBar* self);
extern __declspec(dllexport) 
void QMenuBar_SetNativeMenuBar(QMenuBar* self, bool nativeMenuBar);
extern __declspec(dllexport) 
void QMenuBar_SetVisible(QMenuBar* self, bool visible);
extern __declspec(dllexport) 
void QMenuBar_Triggered(QMenuBar* self, QAction* action);
void QMenuBar_connect_Triggered(QMenuBar* self, intptr_t slot);
extern __declspec(dllexport) 
void QMenuBar_Hovered(QMenuBar* self, QAction* action);
void QMenuBar_connect_Hovered(QMenuBar* self, intptr_t slot);
extern __declspec(dllexport) 
void QMenuBar_ChangeEvent(QMenuBar* self, QEvent* param1);
extern __declspec(dllexport) 
void QMenuBar_KeyPressEvent(QMenuBar* self, QKeyEvent* param1);
extern __declspec(dllexport) 
void QMenuBar_MouseReleaseEvent(QMenuBar* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QMenuBar_MousePressEvent(QMenuBar* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QMenuBar_MouseMoveEvent(QMenuBar* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QMenuBar_LeaveEvent(QMenuBar* self, QEvent* param1);
extern __declspec(dllexport) 
void QMenuBar_PaintEvent(QMenuBar* self, QPaintEvent* param1);
extern __declspec(dllexport) 
void QMenuBar_ResizeEvent(QMenuBar* self, QResizeEvent* param1);
extern __declspec(dllexport) 
void QMenuBar_ActionEvent(QMenuBar* self, QActionEvent* param1);
extern __declspec(dllexport) 
void QMenuBar_FocusOutEvent(QMenuBar* self, QFocusEvent* param1);
extern __declspec(dllexport) 
void QMenuBar_FocusInEvent(QMenuBar* self, QFocusEvent* param1);
extern __declspec(dllexport) 
void QMenuBar_TimerEvent(QMenuBar* self, QTimerEvent* param1);
extern __declspec(dllexport) 
bool QMenuBar_EventFilter(QMenuBar* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) 
bool QMenuBar_Event(QMenuBar* self, QEvent* param1);
extern __declspec(dllexport) 
void QMenuBar_InitStyleOption(const QMenuBar* self, QStyleOptionMenuItem* option, QAction* action);
extern __declspec(dllexport) 
struct miqt_string QMenuBar_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QMenuBar_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QMenuBar_SetCornerWidget2(QMenuBar* self, QWidget* w, int corner);
extern __declspec(dllexport) 
QWidget* QMenuBar_CornerWidget1(const QMenuBar* self, int corner);
extern __declspec(dllexport) 
void QMenuBar_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QMenuBar_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QMenuBar_override_virtual_Metacast(void* self, intptr_t slot);
void* QMenuBar_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QMenuBar_Delete(QMenuBar* self, bool isSubclass);

}
