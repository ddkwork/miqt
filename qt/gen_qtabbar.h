#pragma once
#ifndef MIQT_QT_GEN_QTABBAR_H
#define MIQT_QT_GEN_QTABBAR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QColor QColor;
typedef struct QEvent QEvent;
typedef struct QHideEvent QHideEvent;
typedef struct QIcon QIcon;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QStyleOptionTab QStyleOptionTab;
typedef struct QTabBar QTabBar;
typedef struct QTimerEvent QTimerEvent;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QTabBar* QTabBar_new(QWidget* parent);
extern __declspec(dllexport) 
QTabBar* QTabBar_new2();
extern __declspec(dllexport) 
void QTabBar_virtbase(QTabBar* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QTabBar_MetaObject(const QTabBar* self);
extern __declspec(dllexport) 
void* QTabBar_Metacast(QTabBar* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QTabBar_Tr(const char* s);
extern __declspec(dllexport) 
Shape QTabBar_Shape(const QTabBar* self);
extern __declspec(dllexport) 
void QTabBar_SetShape(QTabBar* self, Shape shape);
extern __declspec(dllexport) 
int QTabBar_AddTab(QTabBar* self, struct miqt_string text);
extern __declspec(dllexport) 
int QTabBar_AddTab2(QTabBar* self, QIcon* icon, struct miqt_string text);
extern __declspec(dllexport) 
int QTabBar_InsertTab(QTabBar* self, int index, struct miqt_string text);
extern __declspec(dllexport) 
int QTabBar_InsertTab2(QTabBar* self, int index, QIcon* icon, struct miqt_string text);
extern __declspec(dllexport) 
void QTabBar_RemoveTab(QTabBar* self, int index);
extern __declspec(dllexport) 
void QTabBar_MoveTab(QTabBar* self, int from, int to);
extern __declspec(dllexport) 
bool QTabBar_IsTabEnabled(const QTabBar* self, int index);
extern __declspec(dllexport) 
void QTabBar_SetTabEnabled(QTabBar* self, int index, bool enabled);
extern __declspec(dllexport) 
bool QTabBar_IsTabVisible(const QTabBar* self, int index);
extern __declspec(dllexport) 
void QTabBar_SetTabVisible(QTabBar* self, int index, bool visible);
extern __declspec(dllexport) 
struct miqt_string QTabBar_TabText(const QTabBar* self, int index);
extern __declspec(dllexport) 
void QTabBar_SetTabText(QTabBar* self, int index, struct miqt_string text);
extern __declspec(dllexport) 
QColor* QTabBar_TabTextColor(const QTabBar* self, int index);
extern __declspec(dllexport) 
void QTabBar_SetTabTextColor(QTabBar* self, int index, QColor* color);
extern __declspec(dllexport) 
QIcon* QTabBar_TabIcon(const QTabBar* self, int index);
extern __declspec(dllexport) 
void QTabBar_SetTabIcon(QTabBar* self, int index, QIcon* icon);
extern __declspec(dllexport) 
int QTabBar_ElideMode(const QTabBar* self);
extern __declspec(dllexport) 
void QTabBar_SetElideMode(QTabBar* self, int mode);
extern __declspec(dllexport) 
void QTabBar_SetTabToolTip(QTabBar* self, int index, struct miqt_string tip);
extern __declspec(dllexport) 
struct miqt_string QTabBar_TabToolTip(const QTabBar* self, int index);
extern __declspec(dllexport) 
void QTabBar_SetTabWhatsThis(QTabBar* self, int index, struct miqt_string text);
extern __declspec(dllexport) 
struct miqt_string QTabBar_TabWhatsThis(const QTabBar* self, int index);
extern __declspec(dllexport) 
void QTabBar_SetTabData(QTabBar* self, int index, QVariant* data);
extern __declspec(dllexport) 
QVariant* QTabBar_TabData(const QTabBar* self, int index);
extern __declspec(dllexport) 
QRect* QTabBar_TabRect(const QTabBar* self, int index);
extern __declspec(dllexport) 
int QTabBar_TabAt(const QTabBar* self, QPoint* pos);
extern __declspec(dllexport) 
int QTabBar_CurrentIndex(const QTabBar* self);
extern __declspec(dllexport) 
int QTabBar_Count(const QTabBar* self);
extern __declspec(dllexport) 
QSize* QTabBar_SizeHint(const QTabBar* self);
extern __declspec(dllexport) 
QSize* QTabBar_MinimumSizeHint(const QTabBar* self);
extern __declspec(dllexport) 
void QTabBar_SetDrawBase(QTabBar* self, bool drawTheBase);
extern __declspec(dllexport) 
bool QTabBar_DrawBase(const QTabBar* self);
extern __declspec(dllexport) 
QSize* QTabBar_IconSize(const QTabBar* self);
extern __declspec(dllexport) 
void QTabBar_SetIconSize(QTabBar* self, QSize* size);
extern __declspec(dllexport) 
bool QTabBar_UsesScrollButtons(const QTabBar* self);
extern __declspec(dllexport) 
void QTabBar_SetUsesScrollButtons(QTabBar* self, bool useButtons);
extern __declspec(dllexport) 
bool QTabBar_TabsClosable(const QTabBar* self);
extern __declspec(dllexport) 
void QTabBar_SetTabsClosable(QTabBar* self, bool closable);
extern __declspec(dllexport) 
void QTabBar_SetTabButton(QTabBar* self, int index, ButtonPosition position, QWidget* widget);
extern __declspec(dllexport) 
QWidget* QTabBar_TabButton(const QTabBar* self, int index, ButtonPosition position);
extern __declspec(dllexport) 
SelectionBehavior QTabBar_SelectionBehaviorOnRemove(const QTabBar* self);
extern __declspec(dllexport) 
void QTabBar_SetSelectionBehaviorOnRemove(QTabBar* self, SelectionBehavior behavior);
extern __declspec(dllexport) 
bool QTabBar_Expanding(const QTabBar* self);
extern __declspec(dllexport) 
void QTabBar_SetExpanding(QTabBar* self, bool enabled);
extern __declspec(dllexport) 
bool QTabBar_IsMovable(const QTabBar* self);
extern __declspec(dllexport) 
void QTabBar_SetMovable(QTabBar* self, bool movable);
extern __declspec(dllexport) 
bool QTabBar_DocumentMode(const QTabBar* self);
extern __declspec(dllexport) 
void QTabBar_SetDocumentMode(QTabBar* self, bool set);
extern __declspec(dllexport) 
bool QTabBar_AutoHide(const QTabBar* self);
extern __declspec(dllexport) 
void QTabBar_SetAutoHide(QTabBar* self, bool hide);
extern __declspec(dllexport) 
bool QTabBar_ChangeCurrentOnDrag(const QTabBar* self);
extern __declspec(dllexport) 
void QTabBar_SetChangeCurrentOnDrag(QTabBar* self, bool change);
extern __declspec(dllexport) 
void QTabBar_SetCurrentIndex(QTabBar* self, int index);
extern __declspec(dllexport) 
void QTabBar_CurrentChanged(QTabBar* self, int index);
void QTabBar_connect_CurrentChanged(QTabBar* self, intptr_t slot);
extern __declspec(dllexport) 
void QTabBar_TabCloseRequested(QTabBar* self, int index);
void QTabBar_connect_TabCloseRequested(QTabBar* self, intptr_t slot);
extern __declspec(dllexport) 
void QTabBar_TabMoved(QTabBar* self, int from, int to);
void QTabBar_connect_TabMoved(QTabBar* self, intptr_t slot);
extern __declspec(dllexport) 
void QTabBar_TabBarClicked(QTabBar* self, int index);
void QTabBar_connect_TabBarClicked(QTabBar* self, intptr_t slot);
extern __declspec(dllexport) 
void QTabBar_TabBarDoubleClicked(QTabBar* self, int index);
void QTabBar_connect_TabBarDoubleClicked(QTabBar* self, intptr_t slot);
extern __declspec(dllexport) 
QSize* QTabBar_TabSizeHint(const QTabBar* self, int index);
extern __declspec(dllexport) 
QSize* QTabBar_MinimumTabSizeHint(const QTabBar* self, int index);
extern __declspec(dllexport) 
void QTabBar_TabInserted(QTabBar* self, int index);
extern __declspec(dllexport) 
void QTabBar_TabRemoved(QTabBar* self, int index);
extern __declspec(dllexport) 
void QTabBar_TabLayoutChange(QTabBar* self);
extern __declspec(dllexport) 
bool QTabBar_Event(QTabBar* self, QEvent* param1);
extern __declspec(dllexport) 
void QTabBar_ResizeEvent(QTabBar* self, QResizeEvent* param1);
extern __declspec(dllexport) 
void QTabBar_ShowEvent(QTabBar* self, QShowEvent* param1);
extern __declspec(dllexport) 
void QTabBar_HideEvent(QTabBar* self, QHideEvent* param1);
extern __declspec(dllexport) 
void QTabBar_PaintEvent(QTabBar* self, QPaintEvent* param1);
extern __declspec(dllexport) 
void QTabBar_MousePressEvent(QTabBar* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QTabBar_MouseMoveEvent(QTabBar* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QTabBar_MouseReleaseEvent(QTabBar* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QTabBar_MouseDoubleClickEvent(QTabBar* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QTabBar_WheelEvent(QTabBar* self, QWheelEvent* event);
extern __declspec(dllexport) 
void QTabBar_KeyPressEvent(QTabBar* self, QKeyEvent* param1);
extern __declspec(dllexport) 
void QTabBar_ChangeEvent(QTabBar* self, QEvent* param1);
extern __declspec(dllexport) 
void QTabBar_TimerEvent(QTabBar* self, QTimerEvent* event);
extern __declspec(dllexport) 
void QTabBar_InitStyleOption(const QTabBar* self, QStyleOptionTab* option, int tabIndex);
extern __declspec(dllexport) 
struct miqt_string QTabBar_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QTabBar_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QTabBar_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QTabBar_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QTabBar_override_virtual_Metacast(void* self, intptr_t slot);
void* QTabBar_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QTabBar_Delete(QTabBar* self, bool isSubclass);

}
