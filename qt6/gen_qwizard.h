#pragma once
#ifndef MIQT_QT6_GEN_QWIZARD_H
#define MIQT_QT6_GEN_QWIZARD_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractButton;
class QActionEvent;
class QChildEvent;
class QCloseEvent;
class QContextMenuEvent;
class QDialog;
class QDragEnterEvent;
class QDragLeaveEvent;
class QDragMoveEvent;
class QDropEvent;
class QEnterEvent;
class QEvent;
class QFocusEvent;
class QHideEvent;
class QInputMethodEvent;
class QKeyEvent;
class QMetaMethod;
class QMetaObject;
class QMouseEvent;
class QMoveEvent;
class QObject;
class QPaintDevice;
class QPaintEngine;
class QPaintEvent;
class QPainter;
class QPixmap;
class QPoint;
class QResizeEvent;
class QShowEvent;
class QSize;
class QTabletEvent;
class QTimerEvent;
class QVariant;
class QWheelEvent;
class QWidget;
class QWizard;
class QWizardPage;
#else
typedef struct QAbstractButton QAbstractButton;
typedef struct QActionEvent QActionEvent;
typedef struct QChildEvent QChildEvent;
typedef struct QCloseEvent QCloseEvent;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDialog QDialog;
typedef struct QDragEnterEvent QDragEnterEvent;
typedef struct QDragLeaveEvent QDragLeaveEvent;
typedef struct QDragMoveEvent QDragMoveEvent;
typedef struct QDropEvent QDropEvent;
typedef struct QEnterEvent QEnterEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QHideEvent QHideEvent;
typedef struct QInputMethodEvent QInputMethodEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QMoveEvent QMoveEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPainter QPainter;
typedef struct QPixmap QPixmap;
typedef struct QPoint QPoint;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QTabletEvent QTabletEvent;
typedef struct QTimerEvent QTimerEvent;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
typedef struct QWizard QWizard;
typedef struct QWizardPage QWizardPage;
#endif

QWizard* QWizard_new(QWidget* parent);
QWizard* QWizard_new2();
QWizard* QWizard_new3(QWidget* parent, int flags);
void QWizard_virtbase(QWizard* src, QDialog** outptr_QDialog);
QMetaObject* QWizard_MetaObject(const QWizard* self);
void* QWizard_Metacast(QWizard* self, const char* param1);
struct miqt_string QWizard_Tr(const char* s);
int QWizard_AddPage(QWizard* self, QWizardPage* page);
void QWizard_SetPage(QWizard* self, int id, QWizardPage* page);
void QWizard_RemovePage(QWizard* self, int id);
QWizardPage* QWizard_Page(const QWizard* self, int id);
bool QWizard_HasVisitedPage(const QWizard* self, int id);
struct miqt_array /* of int */  QWizard_VisitedIds(const QWizard* self);
struct miqt_array /* of int */  QWizard_PageIds(const QWizard* self);
void QWizard_SetStartId(QWizard* self, int id);
int QWizard_StartId(const QWizard* self);
QWizardPage* QWizard_CurrentPage(const QWizard* self);
int QWizard_CurrentId(const QWizard* self);
bool QWizard_ValidateCurrentPage(QWizard* self);
int QWizard_NextId(const QWizard* self);
void QWizard_SetField(QWizard* self, struct miqt_string name, QVariant* value);
QVariant* QWizard_Field(const QWizard* self, struct miqt_string name);
void QWizard_SetWizardStyle(QWizard* self, int style);
int QWizard_WizardStyle(const QWizard* self);
void QWizard_SetOption(QWizard* self, int option);
bool QWizard_TestOption(const QWizard* self, int option);
void QWizard_SetOptions(QWizard* self, int options);
int QWizard_Options(const QWizard* self);
void QWizard_SetButtonText(QWizard* self, int which, struct miqt_string text);
struct miqt_string QWizard_ButtonText(const QWizard* self, int which);
void QWizard_SetButtonLayout(QWizard* self, struct miqt_array /* of int */  layout);
void QWizard_SetButton(QWizard* self, int which, QAbstractButton* button);
QAbstractButton* QWizard_Button(const QWizard* self, int which);
void QWizard_SetTitleFormat(QWizard* self, int format);
int QWizard_TitleFormat(const QWizard* self);
void QWizard_SetSubTitleFormat(QWizard* self, int format);
int QWizard_SubTitleFormat(const QWizard* self);
void QWizard_SetPixmap(QWizard* self, int which, QPixmap* pixmap);
QPixmap* QWizard_Pixmap(const QWizard* self, int which);
void QWizard_SetSideWidget(QWizard* self, QWidget* widget);
QWidget* QWizard_SideWidget(const QWizard* self);
void QWizard_SetDefaultProperty(QWizard* self, const char* className, const char* property, const char* changedSignal);
void QWizard_SetVisible(QWizard* self, bool visible);
QSize* QWizard_SizeHint(const QWizard* self);
void QWizard_CurrentIdChanged(QWizard* self, int id);
void QWizard_connect_CurrentIdChanged(QWizard* self, intptr_t slot);
void QWizard_HelpRequested(QWizard* self);
void QWizard_connect_HelpRequested(QWizard* self, intptr_t slot);
void QWizard_CustomButtonClicked(QWizard* self, int which);
void QWizard_connect_CustomButtonClicked(QWizard* self, intptr_t slot);
void QWizard_PageAdded(QWizard* self, int id);
void QWizard_connect_PageAdded(QWizard* self, intptr_t slot);
void QWizard_PageRemoved(QWizard* self, int id);
void QWizard_connect_PageRemoved(QWizard* self, intptr_t slot);
void QWizard_Back(QWizard* self);
void QWizard_Next(QWizard* self);
void QWizard_SetCurrentId(QWizard* self, int id);
void QWizard_Restart(QWizard* self);
bool QWizard_Event(QWizard* self, QEvent* event);
void QWizard_ResizeEvent(QWizard* self, QResizeEvent* event);
void QWizard_PaintEvent(QWizard* self, QPaintEvent* event);
void QWizard_Done(QWizard* self, int result);
void QWizard_InitializePage(QWizard* self, int id);
void QWizard_CleanupPage(QWizard* self, int id);
struct miqt_string QWizard_Tr2(const char* s, const char* c);
struct miqt_string QWizard_Tr3(const char* s, const char* c, int n);
void QWizard_SetOption2(QWizard* self, int option, bool on);
bool QWizard_override_virtual_ValidateCurrentPage(void* self, intptr_t slot);
bool QWizard_virtualbase_ValidateCurrentPage(void* self);
bool QWizard_override_virtual_NextId(void* self, intptr_t slot);
int QWizard_virtualbase_NextId(const void* self);
bool QWizard_override_virtual_SetVisible(void* self, intptr_t slot);
void QWizard_virtualbase_SetVisible(void* self, bool visible);
bool QWizard_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QWizard_virtualbase_SizeHint(const void* self);
bool QWizard_override_virtual_Event(void* self, intptr_t slot);
bool QWizard_virtualbase_Event(void* self, QEvent* event);
bool QWizard_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QWizard_virtualbase_ResizeEvent(void* self, QResizeEvent* event);
bool QWizard_override_virtual_PaintEvent(void* self, intptr_t slot);
void QWizard_virtualbase_PaintEvent(void* self, QPaintEvent* event);
bool QWizard_override_virtual_Done(void* self, intptr_t slot);
void QWizard_virtualbase_Done(void* self, int result);
bool QWizard_override_virtual_InitializePage(void* self, intptr_t slot);
void QWizard_virtualbase_InitializePage(void* self, int id);
bool QWizard_override_virtual_CleanupPage(void* self, intptr_t slot);
void QWizard_virtualbase_CleanupPage(void* self, int id);
bool QWizard_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QWizard_virtualbase_MinimumSizeHint(const void* self);
bool QWizard_override_virtual_Open(void* self, intptr_t slot);
void QWizard_virtualbase_Open(void* self);
bool QWizard_override_virtual_Exec(void* self, intptr_t slot);
int QWizard_virtualbase_Exec(void* self);
bool QWizard_override_virtual_Accept(void* self, intptr_t slot);
void QWizard_virtualbase_Accept(void* self);
bool QWizard_override_virtual_Reject(void* self, intptr_t slot);
void QWizard_virtualbase_Reject(void* self);
bool QWizard_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QWizard_virtualbase_KeyPressEvent(void* self, QKeyEvent* param1);
bool QWizard_override_virtual_CloseEvent(void* self, intptr_t slot);
void QWizard_virtualbase_CloseEvent(void* self, QCloseEvent* param1);
bool QWizard_override_virtual_ShowEvent(void* self, intptr_t slot);
void QWizard_virtualbase_ShowEvent(void* self, QShowEvent* param1);
bool QWizard_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QWizard_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* param1);
bool QWizard_override_virtual_EventFilter(void* self, intptr_t slot);
bool QWizard_virtualbase_EventFilter(void* self, QObject* param1, QEvent* param2);
bool QWizard_override_virtual_DevType(void* self, intptr_t slot);
int QWizard_virtualbase_DevType(const void* self);
bool QWizard_override_virtual_HeightForWidth(void* self, intptr_t slot);
int QWizard_virtualbase_HeightForWidth(const void* self, int param1);
bool QWizard_override_virtual_HasHeightForWidth(void* self, intptr_t slot);
bool QWizard_virtualbase_HasHeightForWidth(const void* self);
bool QWizard_override_virtual_PaintEngine(void* self, intptr_t slot);
QPaintEngine* QWizard_virtualbase_PaintEngine(const void* self);
bool QWizard_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QWizard_virtualbase_MousePressEvent(void* self, QMouseEvent* event);
bool QWizard_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QWizard_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* event);
bool QWizard_override_virtual_MouseDoubleClickEvent(void* self, intptr_t slot);
void QWizard_virtualbase_MouseDoubleClickEvent(void* self, QMouseEvent* event);
bool QWizard_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QWizard_virtualbase_MouseMoveEvent(void* self, QMouseEvent* event);
bool QWizard_override_virtual_WheelEvent(void* self, intptr_t slot);
void QWizard_virtualbase_WheelEvent(void* self, QWheelEvent* event);
bool QWizard_override_virtual_KeyReleaseEvent(void* self, intptr_t slot);
void QWizard_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* event);
bool QWizard_override_virtual_FocusInEvent(void* self, intptr_t slot);
void QWizard_virtualbase_FocusInEvent(void* self, QFocusEvent* event);
bool QWizard_override_virtual_FocusOutEvent(void* self, intptr_t slot);
void QWizard_virtualbase_FocusOutEvent(void* self, QFocusEvent* event);
bool QWizard_override_virtual_EnterEvent(void* self, intptr_t slot);
void QWizard_virtualbase_EnterEvent(void* self, QEnterEvent* event);
bool QWizard_override_virtual_LeaveEvent(void* self, intptr_t slot);
void QWizard_virtualbase_LeaveEvent(void* self, QEvent* event);
bool QWizard_override_virtual_MoveEvent(void* self, intptr_t slot);
void QWizard_virtualbase_MoveEvent(void* self, QMoveEvent* event);
bool QWizard_override_virtual_TabletEvent(void* self, intptr_t slot);
void QWizard_virtualbase_TabletEvent(void* self, QTabletEvent* event);
bool QWizard_override_virtual_ActionEvent(void* self, intptr_t slot);
void QWizard_virtualbase_ActionEvent(void* self, QActionEvent* event);
bool QWizard_override_virtual_DragEnterEvent(void* self, intptr_t slot);
void QWizard_virtualbase_DragEnterEvent(void* self, QDragEnterEvent* event);
bool QWizard_override_virtual_DragMoveEvent(void* self, intptr_t slot);
void QWizard_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* event);
bool QWizard_override_virtual_DragLeaveEvent(void* self, intptr_t slot);
void QWizard_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* event);
bool QWizard_override_virtual_DropEvent(void* self, intptr_t slot);
void QWizard_virtualbase_DropEvent(void* self, QDropEvent* event);
bool QWizard_override_virtual_HideEvent(void* self, intptr_t slot);
void QWizard_virtualbase_HideEvent(void* self, QHideEvent* event);
bool QWizard_override_virtual_NativeEvent(void* self, intptr_t slot);
bool QWizard_virtualbase_NativeEvent(void* self, struct miqt_string eventType, void* message, intptr_t* result);
bool QWizard_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QWizard_virtualbase_ChangeEvent(void* self, QEvent* param1);
bool QWizard_override_virtual_Metric(void* self, intptr_t slot);
int QWizard_virtualbase_Metric(const void* self, int param1);
bool QWizard_override_virtual_InitPainter(void* self, intptr_t slot);
void QWizard_virtualbase_InitPainter(const void* self, QPainter* painter);
bool QWizard_override_virtual_Redirected(void* self, intptr_t slot);
QPaintDevice* QWizard_virtualbase_Redirected(const void* self, QPoint* offset);
bool QWizard_override_virtual_SharedPainter(void* self, intptr_t slot);
QPainter* QWizard_virtualbase_SharedPainter(const void* self);
bool QWizard_override_virtual_InputMethodEvent(void* self, intptr_t slot);
void QWizard_virtualbase_InputMethodEvent(void* self, QInputMethodEvent* param1);
bool QWizard_override_virtual_InputMethodQuery(void* self, intptr_t slot);
QVariant* QWizard_virtualbase_InputMethodQuery(const void* self, int param1);
bool QWizard_override_virtual_FocusNextPrevChild(void* self, intptr_t slot);
bool QWizard_virtualbase_FocusNextPrevChild(void* self, bool next);
bool QWizard_override_virtual_TimerEvent(void* self, intptr_t slot);
void QWizard_virtualbase_TimerEvent(void* self, QTimerEvent* event);
bool QWizard_override_virtual_ChildEvent(void* self, intptr_t slot);
void QWizard_virtualbase_ChildEvent(void* self, QChildEvent* event);
bool QWizard_override_virtual_CustomEvent(void* self, intptr_t slot);
void QWizard_virtualbase_CustomEvent(void* self, QEvent* event);
bool QWizard_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QWizard_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
bool QWizard_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QWizard_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
void QWizard_Delete(QWizard* self);

QWizardPage* QWizardPage_new(QWidget* parent);
QWizardPage* QWizardPage_new2();
void QWizardPage_virtbase(QWizardPage* src, QWidget** outptr_QWidget);
QMetaObject* QWizardPage_MetaObject(const QWizardPage* self);
void* QWizardPage_Metacast(QWizardPage* self, const char* param1);
struct miqt_string QWizardPage_Tr(const char* s);
void QWizardPage_SetTitle(QWizardPage* self, struct miqt_string title);
struct miqt_string QWizardPage_Title(const QWizardPage* self);
void QWizardPage_SetSubTitle(QWizardPage* self, struct miqt_string subTitle);
struct miqt_string QWizardPage_SubTitle(const QWizardPage* self);
void QWizardPage_SetPixmap(QWizardPage* self, int which, QPixmap* pixmap);
QPixmap* QWizardPage_Pixmap(const QWizardPage* self, int which);
void QWizardPage_SetFinalPage(QWizardPage* self, bool finalPage);
bool QWizardPage_IsFinalPage(const QWizardPage* self);
void QWizardPage_SetCommitPage(QWizardPage* self, bool commitPage);
bool QWizardPage_IsCommitPage(const QWizardPage* self);
void QWizardPage_SetButtonText(QWizardPage* self, int which, struct miqt_string text);
struct miqt_string QWizardPage_ButtonText(const QWizardPage* self, int which);
void QWizardPage_InitializePage(QWizardPage* self);
void QWizardPage_CleanupPage(QWizardPage* self);
bool QWizardPage_ValidatePage(QWizardPage* self);
bool QWizardPage_IsComplete(const QWizardPage* self);
int QWizardPage_NextId(const QWizardPage* self);
void QWizardPage_CompleteChanged(QWizardPage* self);
void QWizardPage_connect_CompleteChanged(QWizardPage* self, intptr_t slot);
struct miqt_string QWizardPage_Tr2(const char* s, const char* c);
struct miqt_string QWizardPage_Tr3(const char* s, const char* c, int n);
bool QWizardPage_override_virtual_InitializePage(void* self, intptr_t slot);
void QWizardPage_virtualbase_InitializePage(void* self);
bool QWizardPage_override_virtual_CleanupPage(void* self, intptr_t slot);
void QWizardPage_virtualbase_CleanupPage(void* self);
bool QWizardPage_override_virtual_ValidatePage(void* self, intptr_t slot);
bool QWizardPage_virtualbase_ValidatePage(void* self);
bool QWizardPage_override_virtual_IsComplete(void* self, intptr_t slot);
bool QWizardPage_virtualbase_IsComplete(const void* self);
bool QWizardPage_override_virtual_NextId(void* self, intptr_t slot);
int QWizardPage_virtualbase_NextId(const void* self);
bool QWizardPage_override_virtual_DevType(void* self, intptr_t slot);
int QWizardPage_virtualbase_DevType(const void* self);
bool QWizardPage_override_virtual_SetVisible(void* self, intptr_t slot);
void QWizardPage_virtualbase_SetVisible(void* self, bool visible);
bool QWizardPage_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QWizardPage_virtualbase_SizeHint(const void* self);
bool QWizardPage_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QWizardPage_virtualbase_MinimumSizeHint(const void* self);
bool QWizardPage_override_virtual_HeightForWidth(void* self, intptr_t slot);
int QWizardPage_virtualbase_HeightForWidth(const void* self, int param1);
bool QWizardPage_override_virtual_HasHeightForWidth(void* self, intptr_t slot);
bool QWizardPage_virtualbase_HasHeightForWidth(const void* self);
bool QWizardPage_override_virtual_PaintEngine(void* self, intptr_t slot);
QPaintEngine* QWizardPage_virtualbase_PaintEngine(const void* self);
bool QWizardPage_override_virtual_Event(void* self, intptr_t slot);
bool QWizardPage_virtualbase_Event(void* self, QEvent* event);
bool QWizardPage_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_MousePressEvent(void* self, QMouseEvent* event);
bool QWizardPage_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* event);
bool QWizardPage_override_virtual_MouseDoubleClickEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_MouseDoubleClickEvent(void* self, QMouseEvent* event);
bool QWizardPage_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_MouseMoveEvent(void* self, QMouseEvent* event);
bool QWizardPage_override_virtual_WheelEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_WheelEvent(void* self, QWheelEvent* event);
bool QWizardPage_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_KeyPressEvent(void* self, QKeyEvent* event);
bool QWizardPage_override_virtual_KeyReleaseEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* event);
bool QWizardPage_override_virtual_FocusInEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_FocusInEvent(void* self, QFocusEvent* event);
bool QWizardPage_override_virtual_FocusOutEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_FocusOutEvent(void* self, QFocusEvent* event);
bool QWizardPage_override_virtual_EnterEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_EnterEvent(void* self, QEnterEvent* event);
bool QWizardPage_override_virtual_LeaveEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_LeaveEvent(void* self, QEvent* event);
bool QWizardPage_override_virtual_PaintEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_PaintEvent(void* self, QPaintEvent* event);
bool QWizardPage_override_virtual_MoveEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_MoveEvent(void* self, QMoveEvent* event);
bool QWizardPage_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_ResizeEvent(void* self, QResizeEvent* event);
bool QWizardPage_override_virtual_CloseEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_CloseEvent(void* self, QCloseEvent* event);
bool QWizardPage_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* event);
bool QWizardPage_override_virtual_TabletEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_TabletEvent(void* self, QTabletEvent* event);
bool QWizardPage_override_virtual_ActionEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_ActionEvent(void* self, QActionEvent* event);
bool QWizardPage_override_virtual_DragEnterEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_DragEnterEvent(void* self, QDragEnterEvent* event);
bool QWizardPage_override_virtual_DragMoveEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* event);
bool QWizardPage_override_virtual_DragLeaveEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* event);
bool QWizardPage_override_virtual_DropEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_DropEvent(void* self, QDropEvent* event);
bool QWizardPage_override_virtual_ShowEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_ShowEvent(void* self, QShowEvent* event);
bool QWizardPage_override_virtual_HideEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_HideEvent(void* self, QHideEvent* event);
bool QWizardPage_override_virtual_NativeEvent(void* self, intptr_t slot);
bool QWizardPage_virtualbase_NativeEvent(void* self, struct miqt_string eventType, void* message, intptr_t* result);
bool QWizardPage_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_ChangeEvent(void* self, QEvent* param1);
bool QWizardPage_override_virtual_Metric(void* self, intptr_t slot);
int QWizardPage_virtualbase_Metric(const void* self, int param1);
bool QWizardPage_override_virtual_InitPainter(void* self, intptr_t slot);
void QWizardPage_virtualbase_InitPainter(const void* self, QPainter* painter);
bool QWizardPage_override_virtual_Redirected(void* self, intptr_t slot);
QPaintDevice* QWizardPage_virtualbase_Redirected(const void* self, QPoint* offset);
bool QWizardPage_override_virtual_SharedPainter(void* self, intptr_t slot);
QPainter* QWizardPage_virtualbase_SharedPainter(const void* self);
bool QWizardPage_override_virtual_InputMethodEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_InputMethodEvent(void* self, QInputMethodEvent* param1);
bool QWizardPage_override_virtual_InputMethodQuery(void* self, intptr_t slot);
QVariant* QWizardPage_virtualbase_InputMethodQuery(const void* self, int param1);
bool QWizardPage_override_virtual_FocusNextPrevChild(void* self, intptr_t slot);
bool QWizardPage_virtualbase_FocusNextPrevChild(void* self, bool next);
bool QWizardPage_override_virtual_EventFilter(void* self, intptr_t slot);
bool QWizardPage_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
bool QWizardPage_override_virtual_TimerEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_TimerEvent(void* self, QTimerEvent* event);
bool QWizardPage_override_virtual_ChildEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_ChildEvent(void* self, QChildEvent* event);
bool QWizardPage_override_virtual_CustomEvent(void* self, intptr_t slot);
void QWizardPage_virtualbase_CustomEvent(void* self, QEvent* event);
bool QWizardPage_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QWizardPage_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
bool QWizardPage_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QWizardPage_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
void QWizardPage_Delete(QWizardPage* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
