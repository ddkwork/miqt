#pragma once
#ifndef MIQT_QT_GEN_QDIALOG_H
#define MIQT_QT_GEN_QDIALOG_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QActionEvent;
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
class QMetaObject;
class QMouseEvent;
class QMoveEvent;
class QObject;
class QPaintDevice;
class QPaintEngine;
class QPaintEvent;
class QPainter;
class QPoint;
class QResizeEvent;
class QShowEvent;
class QSize;
class QTabletEvent;
class QVariant;
class QWheelEvent;
class QWidget;
class _GUID;
class type_info;
#else
typedef struct QActionEvent QActionEvent;
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
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QMoveEvent QMoveEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPainter QPainter;
typedef struct QPoint QPoint;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QTabletEvent QTabletEvent;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QDialog* QDialog_new(QWidget* parent);
extern __declspec(dllexport) QDialog* QDialog_new2();
extern __declspec(dllexport) QDialog* QDialog_new3(QWidget* parent, int f);
extern __declspec(dllexport) void QDialog_virtbase(QDialog* src, QWidget** outptr_QWidget);
extern __declspec(dllexport) QMetaObject* QDialog_MetaObject(const QDialog* self);
extern __declspec(dllexport) void* QDialog_Metacast(QDialog* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QDialog_Tr(const char* s);
extern __declspec(dllexport) int QDialog_Result(const QDialog* self);
extern __declspec(dllexport) void QDialog_SetVisible(QDialog* self, bool visible);
extern __declspec(dllexport) QSize* QDialog_SizeHint(const QDialog* self);
extern __declspec(dllexport) QSize* QDialog_MinimumSizeHint(const QDialog* self);
extern __declspec(dllexport) void QDialog_SetSizeGripEnabled(QDialog* self, bool sizeGripEnabled);
extern __declspec(dllexport) bool QDialog_IsSizeGripEnabled(const QDialog* self);
extern __declspec(dllexport) void QDialog_SetModal(QDialog* self, bool modal);
extern __declspec(dllexport) void QDialog_SetResult(QDialog* self, int r);
extern __declspec(dllexport) void QDialog_Finished(QDialog* self, int result);
void QDialog_connect_Finished(QDialog* self, intptr_t slot);
extern __declspec(dllexport) void QDialog_Accepted(QDialog* self);
void QDialog_connect_Accepted(QDialog* self, intptr_t slot);
extern __declspec(dllexport) void QDialog_Rejected(QDialog* self);
void QDialog_connect_Rejected(QDialog* self, intptr_t slot);
extern __declspec(dllexport) void QDialog_Open(QDialog* self);
extern __declspec(dllexport) int QDialog_Exec(QDialog* self);
extern __declspec(dllexport) void QDialog_Done(QDialog* self, int param1);
extern __declspec(dllexport) void QDialog_Accept(QDialog* self);
extern __declspec(dllexport) void QDialog_Reject(QDialog* self);
extern __declspec(dllexport) void QDialog_KeyPressEvent(QDialog* self, QKeyEvent* param1);
extern __declspec(dllexport) void QDialog_CloseEvent(QDialog* self, QCloseEvent* param1);
extern __declspec(dllexport) void QDialog_ShowEvent(QDialog* self, QShowEvent* param1);
extern __declspec(dllexport) void QDialog_ResizeEvent(QDialog* self, QResizeEvent* param1);
extern __declspec(dllexport) void QDialog_ContextMenuEvent(QDialog* self, QContextMenuEvent* param1);
extern __declspec(dllexport) bool QDialog_EventFilter(QDialog* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) struct miqt_string QDialog_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QDialog_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QDialog_override_virtual_SetVisible(void* self, intptr_t slot);
void QDialog_virtualbase_SetVisible(void* self, bool visible);
extern __declspec(dllexport) void QDialog_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QDialog_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QDialog_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QDialog_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QDialog_override_virtual_Open(void* self, intptr_t slot);
void QDialog_virtualbase_Open(void* self);
extern __declspec(dllexport) void QDialog_override_virtual_Exec(void* self, intptr_t slot);
int QDialog_virtualbase_Exec(void* self);
extern __declspec(dllexport) void QDialog_override_virtual_Done(void* self, intptr_t slot);
void QDialog_virtualbase_Done(void* self, int param1);
extern __declspec(dllexport) void QDialog_override_virtual_Accept(void* self, intptr_t slot);
void QDialog_virtualbase_Accept(void* self);
extern __declspec(dllexport) void QDialog_override_virtual_Reject(void* self, intptr_t slot);
void QDialog_virtualbase_Reject(void* self);
extern __declspec(dllexport) void QDialog_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QDialog_virtualbase_KeyPressEvent(void* self, QKeyEvent* param1);
extern __declspec(dllexport) void QDialog_override_virtual_CloseEvent(void* self, intptr_t slot);
void QDialog_virtualbase_CloseEvent(void* self, QCloseEvent* param1);
extern __declspec(dllexport) void QDialog_override_virtual_ShowEvent(void* self, intptr_t slot);
void QDialog_virtualbase_ShowEvent(void* self, QShowEvent* param1);
extern __declspec(dllexport) void QDialog_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QDialog_virtualbase_ResizeEvent(void* self, QResizeEvent* param1);
extern __declspec(dllexport) void QDialog_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QDialog_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* param1);
extern __declspec(dllexport) void QDialog_override_virtual_EventFilter(void* self, intptr_t slot);
bool QDialog_virtualbase_EventFilter(void* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) void QDialog_override_virtual_DevType(void* self, intptr_t slot);
int QDialog_virtualbase_DevType(const void* self);
extern __declspec(dllexport) void QDialog_override_virtual_HeightForWidth(void* self, intptr_t slot);
int QDialog_virtualbase_HeightForWidth(const void* self, int param1);
extern __declspec(dllexport) void QDialog_override_virtual_HasHeightForWidth(void* self, intptr_t slot);
bool QDialog_virtualbase_HasHeightForWidth(const void* self);
extern __declspec(dllexport) void QDialog_override_virtual_PaintEngine(void* self, intptr_t slot);
QPaintEngine* QDialog_virtualbase_PaintEngine(const void* self);
extern __declspec(dllexport) void QDialog_override_virtual_Event(void* self, intptr_t slot);
bool QDialog_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QDialog_virtualbase_MousePressEvent(void* self, QMouseEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QDialog_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_MouseDoubleClickEvent(void* self, intptr_t slot);
void QDialog_virtualbase_MouseDoubleClickEvent(void* self, QMouseEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QDialog_virtualbase_MouseMoveEvent(void* self, QMouseEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_WheelEvent(void* self, intptr_t slot);
void QDialog_virtualbase_WheelEvent(void* self, QWheelEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_KeyReleaseEvent(void* self, intptr_t slot);
void QDialog_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_FocusInEvent(void* self, intptr_t slot);
void QDialog_virtualbase_FocusInEvent(void* self, QFocusEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_FocusOutEvent(void* self, intptr_t slot);
void QDialog_virtualbase_FocusOutEvent(void* self, QFocusEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_EnterEvent(void* self, intptr_t slot);
void QDialog_virtualbase_EnterEvent(void* self, QEnterEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_LeaveEvent(void* self, intptr_t slot);
void QDialog_virtualbase_LeaveEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_PaintEvent(void* self, intptr_t slot);
void QDialog_virtualbase_PaintEvent(void* self, QPaintEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_MoveEvent(void* self, intptr_t slot);
void QDialog_virtualbase_MoveEvent(void* self, QMoveEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_TabletEvent(void* self, intptr_t slot);
void QDialog_virtualbase_TabletEvent(void* self, QTabletEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_ActionEvent(void* self, intptr_t slot);
void QDialog_virtualbase_ActionEvent(void* self, QActionEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_DragEnterEvent(void* self, intptr_t slot);
void QDialog_virtualbase_DragEnterEvent(void* self, QDragEnterEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_DragMoveEvent(void* self, intptr_t slot);
void QDialog_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_DragLeaveEvent(void* self, intptr_t slot);
void QDialog_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_DropEvent(void* self, intptr_t slot);
void QDialog_virtualbase_DropEvent(void* self, QDropEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_HideEvent(void* self, intptr_t slot);
void QDialog_virtualbase_HideEvent(void* self, QHideEvent* event);
extern __declspec(dllexport) void QDialog_override_virtual_NativeEvent(void* self, intptr_t slot);
bool QDialog_virtualbase_NativeEvent(void* self, struct miqt_string eventType, void* message, intptr_t* result);
extern __declspec(dllexport) void QDialog_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QDialog_virtualbase_ChangeEvent(void* self, QEvent* param1);
extern __declspec(dllexport) void QDialog_override_virtual_Metric(void* self, intptr_t slot);
int QDialog_virtualbase_Metric(const void* self, PaintDeviceMetric param1);
extern __declspec(dllexport) void QDialog_override_virtual_InitPainter(void* self, intptr_t slot);
void QDialog_virtualbase_InitPainter(const void* self, QPainter* painter);
extern __declspec(dllexport) void QDialog_override_virtual_Redirected(void* self, intptr_t slot);
QPaintDevice* QDialog_virtualbase_Redirected(const void* self, QPoint* offset);
extern __declspec(dllexport) void QDialog_override_virtual_SharedPainter(void* self, intptr_t slot);
QPainter* QDialog_virtualbase_SharedPainter(const void* self);
extern __declspec(dllexport) void QDialog_override_virtual_InputMethodEvent(void* self, intptr_t slot);
void QDialog_virtualbase_InputMethodEvent(void* self, QInputMethodEvent* param1);
extern __declspec(dllexport) void QDialog_override_virtual_InputMethodQuery(void* self, intptr_t slot);
QVariant* QDialog_virtualbase_InputMethodQuery(const void* self, int param1);
extern __declspec(dllexport) void QDialog_override_virtual_FocusNextPrevChild(void* self, intptr_t slot);
bool QDialog_virtualbase_FocusNextPrevChild(void* self, bool next);
extern __declspec(dllexport) void QDialog_Delete(QDialog* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
