#pragma once
#ifndef MIQT_QT_SVG_GEN_QSVGWIDGET_H
#define MIQT_QT_SVG_GEN_QSVGWIDGET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
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
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QSvgRenderer QSvgRenderer;
typedef struct QSvgWidget QSvgWidget;
typedef struct QTabletEvent QTabletEvent;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSvgWidget* QSvgWidget_new(QWidget* parent);
extern __declspec(dllexport) QSvgWidget* QSvgWidget_new2();
extern __declspec(dllexport) QSvgWidget* QSvgWidget_new3(struct miqt_string file);
extern __declspec(dllexport) QSvgWidget* QSvgWidget_new4(struct miqt_string file, QWidget* parent);
extern __declspec(dllexport) void QSvgWidget_virtbase(QSvgWidget* src, QWidget** outptr_QWidget);
extern __declspec(dllexport) QMetaObject* QSvgWidget_MetaObject(const QSvgWidget* self);
extern __declspec(dllexport) void* QSvgWidget_Metacast(QSvgWidget* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QSvgWidget_Tr(const char* s);
extern __declspec(dllexport) QSvgRenderer* QSvgWidget_Renderer(const QSvgWidget* self);
extern __declspec(dllexport) QSize* QSvgWidget_SizeHint(const QSvgWidget* self);
extern __declspec(dllexport) int QSvgWidget_Options(const QSvgWidget* self);
extern __declspec(dllexport) void QSvgWidget_SetOptions(QSvgWidget* self, int options);
extern __declspec(dllexport) void QSvgWidget_Load(QSvgWidget* self, struct miqt_string file);
extern __declspec(dllexport) void QSvgWidget_LoadWithContents(QSvgWidget* self, struct miqt_string contents);
extern __declspec(dllexport) void QSvgWidget_PaintEvent(QSvgWidget* self, QPaintEvent* event);
extern __declspec(dllexport) struct miqt_string QSvgWidget_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QSvgWidget_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QSvgWidget_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QSvgWidget_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QSvgWidget_override_virtual_PaintEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_PaintEvent(void* self, QPaintEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_DevType(void* self, intptr_t slot);
int QSvgWidget_virtualbase_DevType(const void* self);
extern __declspec(dllexport) void QSvgWidget_override_virtual_SetVisible(void* self, intptr_t slot);
void QSvgWidget_virtualbase_SetVisible(void* self, bool visible);
extern __declspec(dllexport) void QSvgWidget_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QSvgWidget_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QSvgWidget_override_virtual_HeightForWidth(void* self, intptr_t slot);
int QSvgWidget_virtualbase_HeightForWidth(const void* self, int param1);
extern __declspec(dllexport) void QSvgWidget_override_virtual_HasHeightForWidth(void* self, intptr_t slot);
bool QSvgWidget_virtualbase_HasHeightForWidth(const void* self);
extern __declspec(dllexport) void QSvgWidget_override_virtual_PaintEngine(void* self, intptr_t slot);
QPaintEngine* QSvgWidget_virtualbase_PaintEngine(const void* self);
extern __declspec(dllexport) void QSvgWidget_override_virtual_Event(void* self, intptr_t slot);
bool QSvgWidget_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_MousePressEvent(void* self, QMouseEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_MouseDoubleClickEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_MouseDoubleClickEvent(void* self, QMouseEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_MouseMoveEvent(void* self, QMouseEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_WheelEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_WheelEvent(void* self, QWheelEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_KeyPressEvent(void* self, QKeyEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_KeyReleaseEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_FocusInEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_FocusInEvent(void* self, QFocusEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_FocusOutEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_FocusOutEvent(void* self, QFocusEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_EnterEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_EnterEvent(void* self, QEnterEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_LeaveEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_LeaveEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_MoveEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_MoveEvent(void* self, QMoveEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_ResizeEvent(void* self, QResizeEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_CloseEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_CloseEvent(void* self, QCloseEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_TabletEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_TabletEvent(void* self, QTabletEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_ActionEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_ActionEvent(void* self, QActionEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_DragEnterEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_DragEnterEvent(void* self, QDragEnterEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_DragMoveEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_DragLeaveEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_DropEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_DropEvent(void* self, QDropEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_ShowEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_ShowEvent(void* self, QShowEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_HideEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_HideEvent(void* self, QHideEvent* event);
extern __declspec(dllexport) void QSvgWidget_override_virtual_NativeEvent(void* self, intptr_t slot);
bool QSvgWidget_virtualbase_NativeEvent(void* self, struct miqt_string eventType, void* message, intptr_t* result);
extern __declspec(dllexport) void QSvgWidget_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_ChangeEvent(void* self, QEvent* param1);
extern __declspec(dllexport) void QSvgWidget_override_virtual_Metric(void* self, intptr_t slot);
int QSvgWidget_virtualbase_Metric(const void* self, PaintDeviceMetric param1);
extern __declspec(dllexport) void QSvgWidget_override_virtual_InitPainter(void* self, intptr_t slot);
void QSvgWidget_virtualbase_InitPainter(const void* self, QPainter* painter);
extern __declspec(dllexport) void QSvgWidget_override_virtual_Redirected(void* self, intptr_t slot);
QPaintDevice* QSvgWidget_virtualbase_Redirected(const void* self, QPoint* offset);
extern __declspec(dllexport) void QSvgWidget_override_virtual_SharedPainter(void* self, intptr_t slot);
QPainter* QSvgWidget_virtualbase_SharedPainter(const void* self);
extern __declspec(dllexport) void QSvgWidget_override_virtual_InputMethodEvent(void* self, intptr_t slot);
void QSvgWidget_virtualbase_InputMethodEvent(void* self, QInputMethodEvent* param1);
extern __declspec(dllexport) void QSvgWidget_override_virtual_InputMethodQuery(void* self, intptr_t slot);
QVariant* QSvgWidget_virtualbase_InputMethodQuery(const void* self, int param1);
extern __declspec(dllexport) void QSvgWidget_override_virtual_FocusNextPrevChild(void* self, intptr_t slot);
bool QSvgWidget_virtualbase_FocusNextPrevChild(void* self, bool next);
extern __declspec(dllexport) void QSvgWidget_Delete(QSvgWidget* self, bool isSubclass);

} 