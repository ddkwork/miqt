#pragma once
#ifndef MIQT_QT_GEN_QSCROLLAREA_H
#define MIQT_QT_GEN_QSCROLLAREA_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractScrollArea;
class QContextMenuEvent;
class QDragEnterEvent;
class QDragLeaveEvent;
class QDragMoveEvent;
class QDropEvent;
class QEvent;
class QFrame;
class QKeyEvent;
class QMetaObject;
class QMouseEvent;
class QObject;
class QPaintDevice;
class QPaintEvent;
class QResizeEvent;
class QScrollArea;
class QSize;
class QWheelEvent;
class QWidget;
class _GUID;
class type_info;
#else
typedef struct QAbstractScrollArea QAbstractScrollArea;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDragEnterEvent QDragEnterEvent;
typedef struct QDragLeaveEvent QDragLeaveEvent;
typedef struct QDragMoveEvent QDragMoveEvent;
typedef struct QDropEvent QDropEvent;
typedef struct QEvent QEvent;
typedef struct QFrame QFrame;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QResizeEvent QResizeEvent;
typedef struct QScrollArea QScrollArea;
typedef struct QSize QSize;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QScrollArea* QScrollArea_new(QWidget* parent);
extern __declspec(dllexport) QScrollArea* QScrollArea_new2();
extern __declspec(dllexport) void QScrollArea_virtbase(QScrollArea* src, QAbstractScrollArea** outptr_QAbstractScrollArea);
extern __declspec(dllexport) QMetaObject* QScrollArea_MetaObject(const QScrollArea* self);
extern __declspec(dllexport) void* QScrollArea_Metacast(QScrollArea* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QScrollArea_Tr(const char* s);
extern __declspec(dllexport) QWidget* QScrollArea_Widget(const QScrollArea* self);
extern __declspec(dllexport) void QScrollArea_SetWidget(QScrollArea* self, QWidget* widget);
extern __declspec(dllexport) QWidget* QScrollArea_TakeWidget(QScrollArea* self);
extern __declspec(dllexport) bool QScrollArea_WidgetResizable(const QScrollArea* self);
extern __declspec(dllexport) void QScrollArea_SetWidgetResizable(QScrollArea* self, bool resizable);
extern __declspec(dllexport) QSize* QScrollArea_SizeHint(const QScrollArea* self);
extern __declspec(dllexport) bool QScrollArea_FocusNextPrevChild(QScrollArea* self, bool next);
extern __declspec(dllexport) int QScrollArea_Alignment(const QScrollArea* self);
extern __declspec(dllexport) void QScrollArea_SetAlignment(QScrollArea* self, int alignment);
extern __declspec(dllexport) void QScrollArea_EnsureVisible(QScrollArea* self, int x, int y);
extern __declspec(dllexport) void QScrollArea_EnsureWidgetVisible(QScrollArea* self, QWidget* childWidget);
extern __declspec(dllexport) bool QScrollArea_Event(QScrollArea* self, QEvent* param1);
extern __declspec(dllexport) bool QScrollArea_EventFilter(QScrollArea* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) void QScrollArea_ResizeEvent(QScrollArea* self, QResizeEvent* param1);
extern __declspec(dllexport) void QScrollArea_ScrollContentsBy(QScrollArea* self, int dx, int dy);
extern __declspec(dllexport) QSize* QScrollArea_ViewportSizeHint(const QScrollArea* self);
extern __declspec(dllexport) struct miqt_string QScrollArea_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QScrollArea_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QScrollArea_EnsureVisible3(QScrollArea* self, int x, int y, int xmargin);
extern __declspec(dllexport) void QScrollArea_EnsureVisible4(QScrollArea* self, int x, int y, int xmargin, int ymargin);
extern __declspec(dllexport) void QScrollArea_EnsureWidgetVisible2(QScrollArea* self, QWidget* childWidget, int xmargin);
extern __declspec(dllexport) void QScrollArea_EnsureWidgetVisible3(QScrollArea* self, QWidget* childWidget, int xmargin, int ymargin);
extern __declspec(dllexport) void QScrollArea_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QScrollArea_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QScrollArea_override_virtual_FocusNextPrevChild(void* self, intptr_t slot);
bool QScrollArea_virtualbase_FocusNextPrevChild(void* self, bool next);
extern __declspec(dllexport) void QScrollArea_override_virtual_Event(void* self, intptr_t slot);
bool QScrollArea_virtualbase_Event(void* self, QEvent* param1);
extern __declspec(dllexport) void QScrollArea_override_virtual_EventFilter(void* self, intptr_t slot);
bool QScrollArea_virtualbase_EventFilter(void* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) void QScrollArea_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QScrollArea_virtualbase_ResizeEvent(void* self, QResizeEvent* param1);
extern __declspec(dllexport) void QScrollArea_override_virtual_ScrollContentsBy(void* self, intptr_t slot);
void QScrollArea_virtualbase_ScrollContentsBy(void* self, int dx, int dy);
extern __declspec(dllexport) void QScrollArea_override_virtual_ViewportSizeHint(void* self, intptr_t slot);
QSize* QScrollArea_virtualbase_ViewportSizeHint(const void* self);
extern __declspec(dllexport) void QScrollArea_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QScrollArea_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QScrollArea_override_virtual_SetupViewport(void* self, intptr_t slot);
void QScrollArea_virtualbase_SetupViewport(void* self, QWidget* viewport);
extern __declspec(dllexport) void QScrollArea_override_virtual_ViewportEvent(void* self, intptr_t slot);
bool QScrollArea_virtualbase_ViewportEvent(void* self, QEvent* param1);
extern __declspec(dllexport) void QScrollArea_override_virtual_PaintEvent(void* self, intptr_t slot);
void QScrollArea_virtualbase_PaintEvent(void* self, QPaintEvent* param1);
extern __declspec(dllexport) void QScrollArea_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QScrollArea_virtualbase_MousePressEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QScrollArea_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QScrollArea_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QScrollArea_override_virtual_MouseDoubleClickEvent(void* self, intptr_t slot);
void QScrollArea_virtualbase_MouseDoubleClickEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QScrollArea_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QScrollArea_virtualbase_MouseMoveEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QScrollArea_override_virtual_WheelEvent(void* self, intptr_t slot);
void QScrollArea_virtualbase_WheelEvent(void* self, QWheelEvent* param1);
extern __declspec(dllexport) void QScrollArea_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QScrollArea_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* param1);
extern __declspec(dllexport) void QScrollArea_override_virtual_DragEnterEvent(void* self, intptr_t slot);
void QScrollArea_virtualbase_DragEnterEvent(void* self, QDragEnterEvent* param1);
extern __declspec(dllexport) void QScrollArea_override_virtual_DragMoveEvent(void* self, intptr_t slot);
void QScrollArea_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* param1);
extern __declspec(dllexport) void QScrollArea_override_virtual_DragLeaveEvent(void* self, intptr_t slot);
void QScrollArea_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* param1);
extern __declspec(dllexport) void QScrollArea_override_virtual_DropEvent(void* self, intptr_t slot);
void QScrollArea_virtualbase_DropEvent(void* self, QDropEvent* param1);
extern __declspec(dllexport) void QScrollArea_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QScrollArea_virtualbase_KeyPressEvent(void* self, QKeyEvent* param1);
extern __declspec(dllexport) void QScrollArea_Delete(QScrollArea* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
