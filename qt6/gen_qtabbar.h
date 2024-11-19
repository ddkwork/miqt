#pragma once
#ifndef MIQT_QT6_GEN_QTABBAR_H
#define MIQT_QT6_GEN_QTABBAR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QActionEvent;
class QByteArray;
class QCloseEvent;
class QColor;
class QContextMenuEvent;
class QDragEnterEvent;
class QDragLeaveEvent;
class QDragMoveEvent;
class QDropEvent;
class QEnterEvent;
class QEvent;
class QFocusEvent;
class QHideEvent;
class QIcon;
class QInputMethodEvent;
class QKeyEvent;
class QMetaObject;
class QMouseEvent;
class QMoveEvent;
class QObject;
class QPaintDevice;
class QPaintEngine;
class QPaintEvent;
class QPainter;
class QPoint;
class QRect;
class QResizeEvent;
class QShowEvent;
class QSize;
class QStyleOptionTab;
class QTabBar;
class QTabletEvent;
class QTimerEvent;
class QVariant;
class QWheelEvent;
class QWidget;
#else
typedef struct QActionEvent QActionEvent;
typedef struct QByteArray QByteArray;
typedef struct QCloseEvent QCloseEvent;
typedef struct QColor QColor;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDragEnterEvent QDragEnterEvent;
typedef struct QDragLeaveEvent QDragLeaveEvent;
typedef struct QDragMoveEvent QDragMoveEvent;
typedef struct QDropEvent QDropEvent;
typedef struct QEnterEvent QEnterEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QHideEvent QHideEvent;
typedef struct QIcon QIcon;
typedef struct QInputMethodEvent QInputMethodEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QMoveEvent QMoveEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPainter QPainter;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QStyleOptionTab QStyleOptionTab;
typedef struct QTabBar QTabBar;
typedef struct QTabletEvent QTabletEvent;
typedef struct QTimerEvent QTimerEvent;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
#endif

void QTabBar_new(QWidget* parent, QTabBar** outptr_QTabBar, QWidget** outptr_QWidget, QObject** outptr_QObject, QPaintDevice** outptr_QPaintDevice);
void QTabBar_new2(QTabBar** outptr_QTabBar, QWidget** outptr_QWidget, QObject** outptr_QObject, QPaintDevice** outptr_QPaintDevice);
QMetaObject* QTabBar_MetaObject(const QTabBar* self);
void* QTabBar_Metacast(QTabBar* self, const char* param1);
struct miqt_string QTabBar_Tr(const char* s);
int QTabBar_Shape(const QTabBar* self);
void QTabBar_SetShape(QTabBar* self, int shape);
int QTabBar_AddTab(QTabBar* self, struct miqt_string text);
int QTabBar_AddTab2(QTabBar* self, QIcon* icon, struct miqt_string text);
int QTabBar_InsertTab(QTabBar* self, int index, struct miqt_string text);
int QTabBar_InsertTab2(QTabBar* self, int index, QIcon* icon, struct miqt_string text);
void QTabBar_RemoveTab(QTabBar* self, int index);
void QTabBar_MoveTab(QTabBar* self, int from, int to);
bool QTabBar_IsTabEnabled(const QTabBar* self, int index);
void QTabBar_SetTabEnabled(QTabBar* self, int index, bool enabled);
bool QTabBar_IsTabVisible(const QTabBar* self, int index);
void QTabBar_SetTabVisible(QTabBar* self, int index, bool visible);
struct miqt_string QTabBar_TabText(const QTabBar* self, int index);
void QTabBar_SetTabText(QTabBar* self, int index, struct miqt_string text);
QColor* QTabBar_TabTextColor(const QTabBar* self, int index);
void QTabBar_SetTabTextColor(QTabBar* self, int index, QColor* color);
QIcon* QTabBar_TabIcon(const QTabBar* self, int index);
void QTabBar_SetTabIcon(QTabBar* self, int index, QIcon* icon);
int QTabBar_ElideMode(const QTabBar* self);
void QTabBar_SetElideMode(QTabBar* self, int mode);
void QTabBar_SetTabToolTip(QTabBar* self, int index, struct miqt_string tip);
struct miqt_string QTabBar_TabToolTip(const QTabBar* self, int index);
void QTabBar_SetTabWhatsThis(QTabBar* self, int index, struct miqt_string text);
struct miqt_string QTabBar_TabWhatsThis(const QTabBar* self, int index);
void QTabBar_SetTabData(QTabBar* self, int index, QVariant* data);
QVariant* QTabBar_TabData(const QTabBar* self, int index);
QRect* QTabBar_TabRect(const QTabBar* self, int index);
int QTabBar_TabAt(const QTabBar* self, QPoint* pos);
int QTabBar_CurrentIndex(const QTabBar* self);
int QTabBar_Count(const QTabBar* self);
QSize* QTabBar_SizeHint(const QTabBar* self);
QSize* QTabBar_MinimumSizeHint(const QTabBar* self);
void QTabBar_SetDrawBase(QTabBar* self, bool drawTheBase);
bool QTabBar_DrawBase(const QTabBar* self);
QSize* QTabBar_IconSize(const QTabBar* self);
void QTabBar_SetIconSize(QTabBar* self, QSize* size);
bool QTabBar_UsesScrollButtons(const QTabBar* self);
void QTabBar_SetUsesScrollButtons(QTabBar* self, bool useButtons);
bool QTabBar_TabsClosable(const QTabBar* self);
void QTabBar_SetTabsClosable(QTabBar* self, bool closable);
void QTabBar_SetTabButton(QTabBar* self, int index, int position, QWidget* widget);
QWidget* QTabBar_TabButton(const QTabBar* self, int index, int position);
int QTabBar_SelectionBehaviorOnRemove(const QTabBar* self);
void QTabBar_SetSelectionBehaviorOnRemove(QTabBar* self, int behavior);
bool QTabBar_Expanding(const QTabBar* self);
void QTabBar_SetExpanding(QTabBar* self, bool enabled);
bool QTabBar_IsMovable(const QTabBar* self);
void QTabBar_SetMovable(QTabBar* self, bool movable);
bool QTabBar_DocumentMode(const QTabBar* self);
void QTabBar_SetDocumentMode(QTabBar* self, bool set);
bool QTabBar_AutoHide(const QTabBar* self);
void QTabBar_SetAutoHide(QTabBar* self, bool hide);
bool QTabBar_ChangeCurrentOnDrag(const QTabBar* self);
void QTabBar_SetChangeCurrentOnDrag(QTabBar* self, bool change);
struct miqt_string QTabBar_AccessibleTabName(const QTabBar* self, int index);
void QTabBar_SetAccessibleTabName(QTabBar* self, int index, struct miqt_string name);
void QTabBar_SetCurrentIndex(QTabBar* self, int index);
void QTabBar_CurrentChanged(QTabBar* self, int index);
void QTabBar_connect_CurrentChanged(QTabBar* self, intptr_t slot);
void QTabBar_TabCloseRequested(QTabBar* self, int index);
void QTabBar_connect_TabCloseRequested(QTabBar* self, intptr_t slot);
void QTabBar_TabMoved(QTabBar* self, int from, int to);
void QTabBar_connect_TabMoved(QTabBar* self, intptr_t slot);
void QTabBar_TabBarClicked(QTabBar* self, int index);
void QTabBar_connect_TabBarClicked(QTabBar* self, intptr_t slot);
void QTabBar_TabBarDoubleClicked(QTabBar* self, int index);
void QTabBar_connect_TabBarDoubleClicked(QTabBar* self, intptr_t slot);
QSize* QTabBar_TabSizeHint(const QTabBar* self, int index);
QSize* QTabBar_MinimumTabSizeHint(const QTabBar* self, int index);
void QTabBar_TabInserted(QTabBar* self, int index);
void QTabBar_TabRemoved(QTabBar* self, int index);
void QTabBar_TabLayoutChange(QTabBar* self);
bool QTabBar_Event(QTabBar* self, QEvent* param1);
void QTabBar_ResizeEvent(QTabBar* self, QResizeEvent* param1);
void QTabBar_ShowEvent(QTabBar* self, QShowEvent* param1);
void QTabBar_HideEvent(QTabBar* self, QHideEvent* param1);
void QTabBar_PaintEvent(QTabBar* self, QPaintEvent* param1);
void QTabBar_MousePressEvent(QTabBar* self, QMouseEvent* param1);
void QTabBar_MouseMoveEvent(QTabBar* self, QMouseEvent* param1);
void QTabBar_MouseReleaseEvent(QTabBar* self, QMouseEvent* param1);
void QTabBar_MouseDoubleClickEvent(QTabBar* self, QMouseEvent* param1);
void QTabBar_WheelEvent(QTabBar* self, QWheelEvent* event);
void QTabBar_KeyPressEvent(QTabBar* self, QKeyEvent* param1);
void QTabBar_ChangeEvent(QTabBar* self, QEvent* param1);
void QTabBar_TimerEvent(QTabBar* self, QTimerEvent* event);
void QTabBar_InitStyleOption(const QTabBar* self, QStyleOptionTab* option, int tabIndex);
struct miqt_string QTabBar_Tr2(const char* s, const char* c);
struct miqt_string QTabBar_Tr3(const char* s, const char* c, int n);
void QTabBar_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QTabBar_virtualbase_SizeHint(const void* self);
void QTabBar_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QTabBar_virtualbase_MinimumSizeHint(const void* self);
void QTabBar_override_virtual_TabSizeHint(void* self, intptr_t slot);
QSize* QTabBar_virtualbase_TabSizeHint(const void* self, int index);
void QTabBar_override_virtual_MinimumTabSizeHint(void* self, intptr_t slot);
QSize* QTabBar_virtualbase_MinimumTabSizeHint(const void* self, int index);
void QTabBar_override_virtual_TabInserted(void* self, intptr_t slot);
void QTabBar_virtualbase_TabInserted(void* self, int index);
void QTabBar_override_virtual_TabRemoved(void* self, intptr_t slot);
void QTabBar_virtualbase_TabRemoved(void* self, int index);
void QTabBar_override_virtual_TabLayoutChange(void* self, intptr_t slot);
void QTabBar_virtualbase_TabLayoutChange(void* self);
void QTabBar_override_virtual_Event(void* self, intptr_t slot);
bool QTabBar_virtualbase_Event(void* self, QEvent* param1);
void QTabBar_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_ResizeEvent(void* self, QResizeEvent* param1);
void QTabBar_override_virtual_ShowEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_ShowEvent(void* self, QShowEvent* param1);
void QTabBar_override_virtual_HideEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_HideEvent(void* self, QHideEvent* param1);
void QTabBar_override_virtual_PaintEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_PaintEvent(void* self, QPaintEvent* param1);
void QTabBar_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_MousePressEvent(void* self, QMouseEvent* param1);
void QTabBar_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_MouseMoveEvent(void* self, QMouseEvent* param1);
void QTabBar_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* param1);
void QTabBar_override_virtual_MouseDoubleClickEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_MouseDoubleClickEvent(void* self, QMouseEvent* param1);
void QTabBar_override_virtual_WheelEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_WheelEvent(void* self, QWheelEvent* event);
void QTabBar_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_KeyPressEvent(void* self, QKeyEvent* param1);
void QTabBar_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_ChangeEvent(void* self, QEvent* param1);
void QTabBar_override_virtual_TimerEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_TimerEvent(void* self, QTimerEvent* event);
void QTabBar_override_virtual_InitStyleOption(void* self, intptr_t slot);
void QTabBar_virtualbase_InitStyleOption(const void* self, QStyleOptionTab* option, int tabIndex);
void QTabBar_override_virtual_DevType(void* self, intptr_t slot);
int QTabBar_virtualbase_DevType(const void* self);
void QTabBar_override_virtual_SetVisible(void* self, intptr_t slot);
void QTabBar_virtualbase_SetVisible(void* self, bool visible);
void QTabBar_override_virtual_HeightForWidth(void* self, intptr_t slot);
int QTabBar_virtualbase_HeightForWidth(const void* self, int param1);
void QTabBar_override_virtual_HasHeightForWidth(void* self, intptr_t slot);
bool QTabBar_virtualbase_HasHeightForWidth(const void* self);
void QTabBar_override_virtual_PaintEngine(void* self, intptr_t slot);
QPaintEngine* QTabBar_virtualbase_PaintEngine(const void* self);
void QTabBar_override_virtual_KeyReleaseEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* event);
void QTabBar_override_virtual_FocusInEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_FocusInEvent(void* self, QFocusEvent* event);
void QTabBar_override_virtual_FocusOutEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_FocusOutEvent(void* self, QFocusEvent* event);
void QTabBar_override_virtual_EnterEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_EnterEvent(void* self, QEnterEvent* event);
void QTabBar_override_virtual_LeaveEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_LeaveEvent(void* self, QEvent* event);
void QTabBar_override_virtual_MoveEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_MoveEvent(void* self, QMoveEvent* event);
void QTabBar_override_virtual_CloseEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_CloseEvent(void* self, QCloseEvent* event);
void QTabBar_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* event);
void QTabBar_override_virtual_TabletEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_TabletEvent(void* self, QTabletEvent* event);
void QTabBar_override_virtual_ActionEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_ActionEvent(void* self, QActionEvent* event);
void QTabBar_override_virtual_DragEnterEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_DragEnterEvent(void* self, QDragEnterEvent* event);
void QTabBar_override_virtual_DragMoveEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* event);
void QTabBar_override_virtual_DragLeaveEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* event);
void QTabBar_override_virtual_DropEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_DropEvent(void* self, QDropEvent* event);
void QTabBar_override_virtual_NativeEvent(void* self, intptr_t slot);
bool QTabBar_virtualbase_NativeEvent(void* self, struct miqt_string eventType, void* message, intptr_t* result);
void QTabBar_override_virtual_Metric(void* self, intptr_t slot);
int QTabBar_virtualbase_Metric(const void* self, int param1);
void QTabBar_override_virtual_InitPainter(void* self, intptr_t slot);
void QTabBar_virtualbase_InitPainter(const void* self, QPainter* painter);
void QTabBar_override_virtual_Redirected(void* self, intptr_t slot);
QPaintDevice* QTabBar_virtualbase_Redirected(const void* self, QPoint* offset);
void QTabBar_override_virtual_SharedPainter(void* self, intptr_t slot);
QPainter* QTabBar_virtualbase_SharedPainter(const void* self);
void QTabBar_override_virtual_InputMethodEvent(void* self, intptr_t slot);
void QTabBar_virtualbase_InputMethodEvent(void* self, QInputMethodEvent* param1);
void QTabBar_override_virtual_InputMethodQuery(void* self, intptr_t slot);
QVariant* QTabBar_virtualbase_InputMethodQuery(const void* self, int param1);
void QTabBar_override_virtual_FocusNextPrevChild(void* self, intptr_t slot);
bool QTabBar_virtualbase_FocusNextPrevChild(void* self, bool next);
void QTabBar_Delete(QTabBar* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
