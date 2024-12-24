#pragma once
#ifndef MIQT_QT_GEN_QGROUPBOX_H
#define MIQT_QT_GEN_QGROUPBOX_H

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
class QChildEvent;
class QCloseEvent;
class QContextMenuEvent;
class QDragEnterEvent;
class QDragLeaveEvent;
class QDragMoveEvent;
class QDropEvent;
class QEnterEvent;
class QEvent;
class QFocusEvent;
class QGroupBox;
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
class QStyleOptionGroupBox;
class QTabletEvent;
class QVariant;
class QWheelEvent;
class QWidget;
class _GUID;
class type_info;
#else
typedef struct QActionEvent QActionEvent;
typedef struct QChildEvent QChildEvent;
typedef struct QCloseEvent QCloseEvent;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDragEnterEvent QDragEnterEvent;
typedef struct QDragLeaveEvent QDragLeaveEvent;
typedef struct QDragMoveEvent QDragMoveEvent;
typedef struct QDropEvent QDropEvent;
typedef struct QEnterEvent QEnterEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QGroupBox QGroupBox;
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
typedef struct QStyleOptionGroupBox QStyleOptionGroupBox;
typedef struct QTabletEvent QTabletEvent;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QGroupBox* QGroupBox_new(QWidget* parent);
extern __declspec(dllexport) QGroupBox* QGroupBox_new2();
extern __declspec(dllexport) QGroupBox* QGroupBox_new3(struct miqt_string title);
extern __declspec(dllexport) QGroupBox* QGroupBox_new4(struct miqt_string title, QWidget* parent);
extern __declspec(dllexport) void QGroupBox_virtbase(QGroupBox* src, QWidget** outptr_QWidget);
extern __declspec(dllexport) QMetaObject* QGroupBox_MetaObject(const QGroupBox* self);
extern __declspec(dllexport) void* QGroupBox_Metacast(QGroupBox* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QGroupBox_Tr(const char* s);
extern __declspec(dllexport) struct miqt_string QGroupBox_Title(const QGroupBox* self);
extern __declspec(dllexport) void QGroupBox_SetTitle(QGroupBox* self, struct miqt_string title);
extern __declspec(dllexport) int QGroupBox_Alignment(const QGroupBox* self);
extern __declspec(dllexport) void QGroupBox_SetAlignment(QGroupBox* self, int alignment);
extern __declspec(dllexport) QSize* QGroupBox_MinimumSizeHint(const QGroupBox* self);
extern __declspec(dllexport) bool QGroupBox_IsFlat(const QGroupBox* self);
extern __declspec(dllexport) void QGroupBox_SetFlat(QGroupBox* self, bool flat);
extern __declspec(dllexport) bool QGroupBox_IsCheckable(const QGroupBox* self);
extern __declspec(dllexport) void QGroupBox_SetCheckable(QGroupBox* self, bool checkable);
extern __declspec(dllexport) bool QGroupBox_IsChecked(const QGroupBox* self);
extern __declspec(dllexport) void QGroupBox_SetChecked(QGroupBox* self, bool checked);
extern __declspec(dllexport) void QGroupBox_Clicked(QGroupBox* self);
void QGroupBox_connect_Clicked(QGroupBox* self, intptr_t slot);
extern __declspec(dllexport) void QGroupBox_Toggled(QGroupBox* self, bool param1);
void QGroupBox_connect_Toggled(QGroupBox* self, intptr_t slot);
extern __declspec(dllexport) bool QGroupBox_Event(QGroupBox* self, QEvent* event);
extern __declspec(dllexport) void QGroupBox_ChildEvent(QGroupBox* self, QChildEvent* event);
extern __declspec(dllexport) void QGroupBox_ResizeEvent(QGroupBox* self, QResizeEvent* event);
extern __declspec(dllexport) void QGroupBox_PaintEvent(QGroupBox* self, QPaintEvent* event);
extern __declspec(dllexport) void QGroupBox_FocusInEvent(QGroupBox* self, QFocusEvent* event);
extern __declspec(dllexport) void QGroupBox_ChangeEvent(QGroupBox* self, QEvent* event);
extern __declspec(dllexport) void QGroupBox_MousePressEvent(QGroupBox* self, QMouseEvent* event);
extern __declspec(dllexport) void QGroupBox_MouseMoveEvent(QGroupBox* self, QMouseEvent* event);
extern __declspec(dllexport) void QGroupBox_MouseReleaseEvent(QGroupBox* self, QMouseEvent* event);
extern __declspec(dllexport) void QGroupBox_InitStyleOption(const QGroupBox* self, QStyleOptionGroupBox* option);
extern __declspec(dllexport) struct miqt_string QGroupBox_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QGroupBox_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QGroupBox_Clicked1(QGroupBox* self, bool checked);
void QGroupBox_connect_Clicked1(QGroupBox* self, intptr_t slot);
extern __declspec(dllexport) void QGroupBox_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QGroupBox_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QGroupBox_override_virtual_Event(void* self, intptr_t slot);
bool QGroupBox_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_ChildEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_ResizeEvent(void* self, QResizeEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_PaintEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_PaintEvent(void* self, QPaintEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_FocusInEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_FocusInEvent(void* self, QFocusEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_ChangeEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_MousePressEvent(void* self, QMouseEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_MouseMoveEvent(void* self, QMouseEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_InitStyleOption(void* self, intptr_t slot);
void QGroupBox_virtualbase_InitStyleOption(const void* self, QStyleOptionGroupBox* option);
extern __declspec(dllexport) void QGroupBox_override_virtual_DevType(void* self, intptr_t slot);
int QGroupBox_virtualbase_DevType(const void* self);
extern __declspec(dllexport) void QGroupBox_override_virtual_SetVisible(void* self, intptr_t slot);
void QGroupBox_virtualbase_SetVisible(void* self, bool visible);
extern __declspec(dllexport) void QGroupBox_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QGroupBox_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QGroupBox_override_virtual_HeightForWidth(void* self, intptr_t slot);
int QGroupBox_virtualbase_HeightForWidth(const void* self, int param1);
extern __declspec(dllexport) void QGroupBox_override_virtual_HasHeightForWidth(void* self, intptr_t slot);
bool QGroupBox_virtualbase_HasHeightForWidth(const void* self);
extern __declspec(dllexport) void QGroupBox_override_virtual_PaintEngine(void* self, intptr_t slot);
QPaintEngine* QGroupBox_virtualbase_PaintEngine(const void* self);
extern __declspec(dllexport) void QGroupBox_override_virtual_MouseDoubleClickEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_MouseDoubleClickEvent(void* self, QMouseEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_WheelEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_WheelEvent(void* self, QWheelEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_KeyPressEvent(void* self, QKeyEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_KeyReleaseEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_FocusOutEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_FocusOutEvent(void* self, QFocusEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_EnterEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_EnterEvent(void* self, QEnterEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_LeaveEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_LeaveEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_MoveEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_MoveEvent(void* self, QMoveEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_CloseEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_CloseEvent(void* self, QCloseEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_TabletEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_TabletEvent(void* self, QTabletEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_ActionEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_ActionEvent(void* self, QActionEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_DragEnterEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_DragEnterEvent(void* self, QDragEnterEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_DragMoveEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_DragLeaveEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_DropEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_DropEvent(void* self, QDropEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_ShowEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_ShowEvent(void* self, QShowEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_HideEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_HideEvent(void* self, QHideEvent* event);
extern __declspec(dllexport) void QGroupBox_override_virtual_NativeEvent(void* self, intptr_t slot);
bool QGroupBox_virtualbase_NativeEvent(void* self, struct miqt_string eventType, void* message, intptr_t* result);
extern __declspec(dllexport) void QGroupBox_override_virtual_Metric(void* self, intptr_t slot);
int QGroupBox_virtualbase_Metric(const void* self, PaintDeviceMetric param1);
extern __declspec(dllexport) void QGroupBox_override_virtual_InitPainter(void* self, intptr_t slot);
void QGroupBox_virtualbase_InitPainter(const void* self, QPainter* painter);
extern __declspec(dllexport) void QGroupBox_override_virtual_Redirected(void* self, intptr_t slot);
QPaintDevice* QGroupBox_virtualbase_Redirected(const void* self, QPoint* offset);
extern __declspec(dllexport) void QGroupBox_override_virtual_SharedPainter(void* self, intptr_t slot);
QPainter* QGroupBox_virtualbase_SharedPainter(const void* self);
extern __declspec(dllexport) void QGroupBox_override_virtual_InputMethodEvent(void* self, intptr_t slot);
void QGroupBox_virtualbase_InputMethodEvent(void* self, QInputMethodEvent* param1);
extern __declspec(dllexport) void QGroupBox_override_virtual_InputMethodQuery(void* self, intptr_t slot);
QVariant* QGroupBox_virtualbase_InputMethodQuery(const void* self, int param1);
extern __declspec(dllexport) void QGroupBox_override_virtual_FocusNextPrevChild(void* self, intptr_t slot);
bool QGroupBox_virtualbase_FocusNextPrevChild(void* self, bool next);
extern __declspec(dllexport) void QGroupBox_Delete(QGroupBox* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
