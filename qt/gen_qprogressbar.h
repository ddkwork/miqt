#pragma once
#ifndef MIQT_QT_GEN_QPROGRESSBAR_H
#define MIQT_QT_GEN_QPROGRESSBAR_H

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
class QProgressBar;
class QResizeEvent;
class QShowEvent;
class QSize;
class QStyleOptionProgressBar;
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
typedef struct QProgressBar QProgressBar;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QStyleOptionProgressBar QStyleOptionProgressBar;
typedef struct QTabletEvent QTabletEvent;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QProgressBar* QProgressBar_new(QWidget* parent);
extern __declspec(dllexport) QProgressBar* QProgressBar_new2();
extern __declspec(dllexport) void QProgressBar_virtbase(QProgressBar* src, QWidget** outptr_QWidget);
extern __declspec(dllexport) QMetaObject* QProgressBar_MetaObject(const QProgressBar* self);
extern __declspec(dllexport) void* QProgressBar_Metacast(QProgressBar* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QProgressBar_Tr(const char* s);
extern __declspec(dllexport) int QProgressBar_Minimum(const QProgressBar* self);
extern __declspec(dllexport) int QProgressBar_Maximum(const QProgressBar* self);
extern __declspec(dllexport) int QProgressBar_Value(const QProgressBar* self);
extern __declspec(dllexport) struct miqt_string QProgressBar_Text(const QProgressBar* self);
extern __declspec(dllexport) void QProgressBar_SetTextVisible(QProgressBar* self, bool visible);
extern __declspec(dllexport) bool QProgressBar_IsTextVisible(const QProgressBar* self);
extern __declspec(dllexport) int QProgressBar_Alignment(const QProgressBar* self);
extern __declspec(dllexport) void QProgressBar_SetAlignment(QProgressBar* self, int alignment);
extern __declspec(dllexport) QSize* QProgressBar_SizeHint(const QProgressBar* self);
extern __declspec(dllexport) QSize* QProgressBar_MinimumSizeHint(const QProgressBar* self);
extern __declspec(dllexport) int QProgressBar_Orientation(const QProgressBar* self);
extern __declspec(dllexport) void QProgressBar_SetInvertedAppearance(QProgressBar* self, bool invert);
extern __declspec(dllexport) bool QProgressBar_InvertedAppearance(const QProgressBar* self);
extern __declspec(dllexport) void QProgressBar_SetTextDirection(QProgressBar* self, int textDirection);
extern __declspec(dllexport) int QProgressBar_TextDirection(const QProgressBar* self);
extern __declspec(dllexport) void QProgressBar_SetFormat(QProgressBar* self, struct miqt_string format);
extern __declspec(dllexport) void QProgressBar_ResetFormat(QProgressBar* self);
extern __declspec(dllexport) struct miqt_string QProgressBar_Format(const QProgressBar* self);
extern __declspec(dllexport) void QProgressBar_Reset(QProgressBar* self);
extern __declspec(dllexport) void QProgressBar_SetRange(QProgressBar* self, int minimum, int maximum);
extern __declspec(dllexport) void QProgressBar_SetMinimum(QProgressBar* self, int minimum);
extern __declspec(dllexport) void QProgressBar_SetMaximum(QProgressBar* self, int maximum);
extern __declspec(dllexport) void QProgressBar_SetValue(QProgressBar* self, int value);
extern __declspec(dllexport) void QProgressBar_SetOrientation(QProgressBar* self, int orientation);
extern __declspec(dllexport) void QProgressBar_ValueChanged(QProgressBar* self, int value);
void QProgressBar_connect_ValueChanged(QProgressBar* self, intptr_t slot);
extern __declspec(dllexport) bool QProgressBar_Event(QProgressBar* self, QEvent* e);
extern __declspec(dllexport) void QProgressBar_PaintEvent(QProgressBar* self, QPaintEvent* param1);
extern __declspec(dllexport) void QProgressBar_InitStyleOption(const QProgressBar* self, QStyleOptionProgressBar* option);
extern __declspec(dllexport) struct miqt_string QProgressBar_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QProgressBar_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QProgressBar_override_virtual_Text(void* self, intptr_t slot);
struct miqt_string QProgressBar_virtualbase_Text(const void* self);
extern __declspec(dllexport) void QProgressBar_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QProgressBar_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QProgressBar_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QProgressBar_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QProgressBar_override_virtual_Event(void* self, intptr_t slot);
bool QProgressBar_virtualbase_Event(void* self, QEvent* e);
extern __declspec(dllexport) void QProgressBar_override_virtual_PaintEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_PaintEvent(void* self, QPaintEvent* param1);
extern __declspec(dllexport) void QProgressBar_override_virtual_InitStyleOption(void* self, intptr_t slot);
void QProgressBar_virtualbase_InitStyleOption(const void* self, QStyleOptionProgressBar* option);
extern __declspec(dllexport) void QProgressBar_override_virtual_DevType(void* self, intptr_t slot);
int QProgressBar_virtualbase_DevType(const void* self);
extern __declspec(dllexport) void QProgressBar_override_virtual_SetVisible(void* self, intptr_t slot);
void QProgressBar_virtualbase_SetVisible(void* self, bool visible);
extern __declspec(dllexport) void QProgressBar_override_virtual_HeightForWidth(void* self, intptr_t slot);
int QProgressBar_virtualbase_HeightForWidth(const void* self, int param1);
extern __declspec(dllexport) void QProgressBar_override_virtual_HasHeightForWidth(void* self, intptr_t slot);
bool QProgressBar_virtualbase_HasHeightForWidth(const void* self);
extern __declspec(dllexport) void QProgressBar_override_virtual_PaintEngine(void* self, intptr_t slot);
QPaintEngine* QProgressBar_virtualbase_PaintEngine(const void* self);
extern __declspec(dllexport) void QProgressBar_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_MousePressEvent(void* self, QMouseEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_MouseDoubleClickEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_MouseDoubleClickEvent(void* self, QMouseEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_MouseMoveEvent(void* self, QMouseEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_WheelEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_WheelEvent(void* self, QWheelEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_KeyPressEvent(void* self, QKeyEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_KeyReleaseEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_FocusInEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_FocusInEvent(void* self, QFocusEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_FocusOutEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_FocusOutEvent(void* self, QFocusEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_EnterEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_EnterEvent(void* self, QEnterEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_LeaveEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_LeaveEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_MoveEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_MoveEvent(void* self, QMoveEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_ResizeEvent(void* self, QResizeEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_CloseEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_CloseEvent(void* self, QCloseEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_TabletEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_TabletEvent(void* self, QTabletEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_ActionEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_ActionEvent(void* self, QActionEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_DragEnterEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_DragEnterEvent(void* self, QDragEnterEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_DragMoveEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_DragLeaveEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_DropEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_DropEvent(void* self, QDropEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_ShowEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_ShowEvent(void* self, QShowEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_HideEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_HideEvent(void* self, QHideEvent* event);
extern __declspec(dllexport) void QProgressBar_override_virtual_NativeEvent(void* self, intptr_t slot);
bool QProgressBar_virtualbase_NativeEvent(void* self, struct miqt_string eventType, void* message, intptr_t* result);
extern __declspec(dllexport) void QProgressBar_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_ChangeEvent(void* self, QEvent* param1);
extern __declspec(dllexport) void QProgressBar_override_virtual_Metric(void* self, intptr_t slot);
int QProgressBar_virtualbase_Metric(const void* self, PaintDeviceMetric param1);
extern __declspec(dllexport) void QProgressBar_override_virtual_InitPainter(void* self, intptr_t slot);
void QProgressBar_virtualbase_InitPainter(const void* self, QPainter* painter);
extern __declspec(dllexport) void QProgressBar_override_virtual_Redirected(void* self, intptr_t slot);
QPaintDevice* QProgressBar_virtualbase_Redirected(const void* self, QPoint* offset);
extern __declspec(dllexport) void QProgressBar_override_virtual_SharedPainter(void* self, intptr_t slot);
QPainter* QProgressBar_virtualbase_SharedPainter(const void* self);
extern __declspec(dllexport) void QProgressBar_override_virtual_InputMethodEvent(void* self, intptr_t slot);
void QProgressBar_virtualbase_InputMethodEvent(void* self, QInputMethodEvent* param1);
extern __declspec(dllexport) void QProgressBar_override_virtual_InputMethodQuery(void* self, intptr_t slot);
QVariant* QProgressBar_virtualbase_InputMethodQuery(const void* self, int param1);
extern __declspec(dllexport) void QProgressBar_override_virtual_FocusNextPrevChild(void* self, intptr_t slot);
bool QProgressBar_virtualbase_FocusNextPrevChild(void* self, bool next);
extern __declspec(dllexport) void QProgressBar_Delete(QProgressBar* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
