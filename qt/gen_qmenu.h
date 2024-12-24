#pragma once
#ifndef MIQT_QT_GEN_QMENU_H
#define MIQT_QT_GEN_QMENU_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAction QAction;
typedef struct QActionEvent QActionEvent;
typedef struct QEnterEvent QEnterEvent;
typedef struct QEvent QEvent;
typedef struct QHideEvent QHideEvent;
typedef struct QIcon QIcon;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMenu QMenu;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QStyleOptionMenuItem QStyleOptionMenuItem;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QMenu* QMenu_new(QWidget* parent);
extern __declspec(dllexport) 
QMenu* QMenu_new2();
extern __declspec(dllexport) 
QMenu* QMenu_new3(struct miqt_string title);
extern __declspec(dllexport) 
QMenu* QMenu_new4(struct miqt_string title, QWidget* parent);
extern __declspec(dllexport) 
void QMenu_virtbase(QMenu* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QMenu_MetaObject(const QMenu* self);
extern __declspec(dllexport) 
void* QMenu_Metacast(QMenu* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QMenu_Tr(const char* s);
extern __declspec(dllexport) 
QAction* QMenu_AddMenu(QMenu* self, QMenu* menu);
extern __declspec(dllexport) 
QMenu* QMenu_AddMenuWithTitle(QMenu* self, struct miqt_string title);
extern __declspec(dllexport) 
QMenu* QMenu_AddMenu2(QMenu* self, QIcon* icon, struct miqt_string title);
extern __declspec(dllexport) 
QAction* QMenu_AddSeparator(QMenu* self);
extern __declspec(dllexport) 
QAction* QMenu_AddSection(QMenu* self, struct miqt_string text);
extern __declspec(dllexport) 
QAction* QMenu_AddSection2(QMenu* self, QIcon* icon, struct miqt_string text);
extern __declspec(dllexport) 
QAction* QMenu_InsertMenu(QMenu* self, QAction* before, QMenu* menu);
extern __declspec(dllexport) 
QAction* QMenu_InsertSeparator(QMenu* self, QAction* before);
extern __declspec(dllexport) 
QAction* QMenu_InsertSection(QMenu* self, QAction* before, struct miqt_string text);
extern __declspec(dllexport) 
QAction* QMenu_InsertSection2(QMenu* self, QAction* before, QIcon* icon, struct miqt_string text);
extern __declspec(dllexport) 
bool QMenu_IsEmpty(const QMenu* self);
extern __declspec(dllexport) 
void QMenu_Clear(QMenu* self);
extern __declspec(dllexport) 
void QMenu_SetTearOffEnabled(QMenu* self, bool tearOffEnabled);
extern __declspec(dllexport) 
bool QMenu_IsTearOffEnabled(const QMenu* self);
extern __declspec(dllexport) 
bool QMenu_IsTearOffMenuVisible(const QMenu* self);
extern __declspec(dllexport) 
void QMenu_ShowTearOffMenu(QMenu* self);
extern __declspec(dllexport) 
void QMenu_ShowTearOffMenuWithPos(QMenu* self, QPoint* pos);
extern __declspec(dllexport) 
void QMenu_HideTearOffMenu(QMenu* self);
extern __declspec(dllexport) 
void QMenu_SetDefaultAction(QMenu* self, QAction* defaultAction);
extern __declspec(dllexport) 
QAction* QMenu_DefaultAction(const QMenu* self);
extern __declspec(dllexport) 
void QMenu_SetActiveAction(QMenu* self, QAction* act);
extern __declspec(dllexport) 
QAction* QMenu_ActiveAction(const QMenu* self);
extern __declspec(dllexport) 
void QMenu_Popup(QMenu* self, QPoint* pos);
extern __declspec(dllexport) 
QAction* QMenu_Exec(QMenu* self);
extern __declspec(dllexport) 
QAction* QMenu_ExecWithPos(QMenu* self, QPoint* pos);
extern __declspec(dllexport) 
QAction* QMenu_Exec2(struct miqt_array /* of QAction* */  actions, QPoint* pos);
extern __declspec(dllexport) 
QSize* QMenu_SizeHint(const QMenu* self);
extern __declspec(dllexport) 
QRect* QMenu_ActionGeometry(const QMenu* self, QAction* param1);
extern __declspec(dllexport) 
QAction* QMenu_ActionAt(const QMenu* self, QPoint* param1);
extern __declspec(dllexport) 
QAction* QMenu_MenuAction(const QMenu* self);
extern __declspec(dllexport) 
QMenu* QMenu_MenuInAction(QAction* action);
extern __declspec(dllexport) 
struct miqt_string QMenu_Title(const QMenu* self);
extern __declspec(dllexport) 
void QMenu_SetTitle(QMenu* self, struct miqt_string title);
extern __declspec(dllexport) 
QIcon* QMenu_Icon(const QMenu* self);
extern __declspec(dllexport) 
void QMenu_SetIcon(QMenu* self, QIcon* icon);
extern __declspec(dllexport) 
void QMenu_SetNoReplayFor(QMenu* self, QWidget* widget);
extern __declspec(dllexport) 
bool QMenu_SeparatorsCollapsible(const QMenu* self);
extern __declspec(dllexport) 
void QMenu_SetSeparatorsCollapsible(QMenu* self, bool collapse);
extern __declspec(dllexport) 
bool QMenu_ToolTipsVisible(const QMenu* self);
extern __declspec(dllexport) 
void QMenu_SetToolTipsVisible(QMenu* self, bool visible);
extern __declspec(dllexport) 
void QMenu_AboutToShow(QMenu* self);
void QMenu_connect_AboutToShow(QMenu* self, intptr_t slot);
extern __declspec(dllexport) 
void QMenu_AboutToHide(QMenu* self);
void QMenu_connect_AboutToHide(QMenu* self, intptr_t slot);
extern __declspec(dllexport) 
void QMenu_Triggered(QMenu* self, QAction* action);
void QMenu_connect_Triggered(QMenu* self, intptr_t slot);
extern __declspec(dllexport) 
void QMenu_Hovered(QMenu* self, QAction* action);
void QMenu_connect_Hovered(QMenu* self, intptr_t slot);
extern __declspec(dllexport) 
void QMenu_ChangeEvent(QMenu* self, QEvent* param1);
extern __declspec(dllexport) 
void QMenu_KeyPressEvent(QMenu* self, QKeyEvent* param1);
extern __declspec(dllexport) 
void QMenu_MouseReleaseEvent(QMenu* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QMenu_MousePressEvent(QMenu* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QMenu_MouseMoveEvent(QMenu* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QMenu_WheelEvent(QMenu* self, QWheelEvent* param1);
extern __declspec(dllexport) 
void QMenu_EnterEvent(QMenu* self, QEnterEvent* param1);
extern __declspec(dllexport) 
void QMenu_LeaveEvent(QMenu* self, QEvent* param1);
extern __declspec(dllexport) 
void QMenu_HideEvent(QMenu* self, QHideEvent* param1);
extern __declspec(dllexport) 
void QMenu_PaintEvent(QMenu* self, QPaintEvent* param1);
extern __declspec(dllexport) 
void QMenu_ActionEvent(QMenu* self, QActionEvent* param1);
extern __declspec(dllexport) 
void QMenu_TimerEvent(QMenu* self, QTimerEvent* param1);
extern __declspec(dllexport) 
bool QMenu_Event(QMenu* self, QEvent* param1);
extern __declspec(dllexport) 
bool QMenu_FocusNextPrevChild(QMenu* self, bool next);
extern __declspec(dllexport) 
void QMenu_InitStyleOption(const QMenu* self, QStyleOptionMenuItem* option, QAction* action);
extern __declspec(dllexport) 
struct miqt_string QMenu_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QMenu_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QMenu_Popup2(QMenu* self, QPoint* pos, QAction* at);
extern __declspec(dllexport) 
QAction* QMenu_Exec22(QMenu* self, QPoint* pos, QAction* at);
extern __declspec(dllexport) 
QAction* QMenu_Exec3(struct miqt_array /* of QAction* */  actions, QPoint* pos, QAction* at);
extern __declspec(dllexport) 
QAction* QMenu_Exec4(struct miqt_array /* of QAction* */  actions, QPoint* pos, QAction* at, QWidget* parent);
extern __declspec(dllexport) 
void QMenu_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QMenu_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QMenu_override_virtual_Metacast(void* self, intptr_t slot);
void* QMenu_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QMenu_Delete(QMenu* self, bool isSubclass);

}
