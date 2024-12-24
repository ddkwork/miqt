#pragma once
#ifndef MIQT_QT_GEN_QABSTRACTSCROLLAREA_H
#define MIQT_QT_GEN_QABSTRACTSCROLLAREA_H

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
class QScrollBar;
class QSize;
class QStyleOptionFrame;
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
typedef struct QScrollBar QScrollBar;
typedef struct QSize QSize;
typedef struct QStyleOptionFrame QStyleOptionFrame;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QAbstractScrollArea* QAbstractScrollArea_new(QWidget* parent);
extern __declspec(dllexport) QAbstractScrollArea* QAbstractScrollArea_new2();
extern __declspec(dllexport) void QAbstractScrollArea_virtbase(QAbstractScrollArea* src, QFrame** outptr_QFrame);
extern __declspec(dllexport) QMetaObject* QAbstractScrollArea_MetaObject(const QAbstractScrollArea* self);
extern __declspec(dllexport) void* QAbstractScrollArea_Metacast(QAbstractScrollArea* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QAbstractScrollArea_Tr(const char* s);
extern __declspec(dllexport) int QAbstractScrollArea_VerticalScrollBarPolicy(const QAbstractScrollArea* self);
extern __declspec(dllexport) void QAbstractScrollArea_SetVerticalScrollBarPolicy(QAbstractScrollArea* self, int verticalScrollBarPolicy);
extern __declspec(dllexport) QScrollBar* QAbstractScrollArea_VerticalScrollBar(const QAbstractScrollArea* self);
extern __declspec(dllexport) void QAbstractScrollArea_SetVerticalScrollBar(QAbstractScrollArea* self, QScrollBar* scrollbar);
extern __declspec(dllexport) int QAbstractScrollArea_HorizontalScrollBarPolicy(const QAbstractScrollArea* self);
extern __declspec(dllexport) void QAbstractScrollArea_SetHorizontalScrollBarPolicy(QAbstractScrollArea* self, int horizontalScrollBarPolicy);
extern __declspec(dllexport) QScrollBar* QAbstractScrollArea_HorizontalScrollBar(const QAbstractScrollArea* self);
extern __declspec(dllexport) void QAbstractScrollArea_SetHorizontalScrollBar(QAbstractScrollArea* self, QScrollBar* scrollbar);
extern __declspec(dllexport) QWidget* QAbstractScrollArea_CornerWidget(const QAbstractScrollArea* self);
extern __declspec(dllexport) void QAbstractScrollArea_SetCornerWidget(QAbstractScrollArea* self, QWidget* widget);
extern __declspec(dllexport) void QAbstractScrollArea_AddScrollBarWidget(QAbstractScrollArea* self, QWidget* widget, int alignment);
extern __declspec(dllexport) struct miqt_array /* of QWidget* */  QAbstractScrollArea_ScrollBarWidgets(QAbstractScrollArea* self, int alignment);
extern __declspec(dllexport) QWidget* QAbstractScrollArea_Viewport(const QAbstractScrollArea* self);
extern __declspec(dllexport) void QAbstractScrollArea_SetViewport(QAbstractScrollArea* self, QWidget* widget);
extern __declspec(dllexport) QSize* QAbstractScrollArea_MaximumViewportSize(const QAbstractScrollArea* self);
extern __declspec(dllexport) QSize* QAbstractScrollArea_MinimumSizeHint(const QAbstractScrollArea* self);
extern __declspec(dllexport) QSize* QAbstractScrollArea_SizeHint(const QAbstractScrollArea* self);
extern __declspec(dllexport) void QAbstractScrollArea_SetupViewport(QAbstractScrollArea* self, QWidget* viewport);
extern __declspec(dllexport) SizeAdjustPolicy QAbstractScrollArea_SizeAdjustPolicy(const QAbstractScrollArea* self);
extern __declspec(dllexport) void QAbstractScrollArea_SetSizeAdjustPolicy(QAbstractScrollArea* self, SizeAdjustPolicy policy);
extern __declspec(dllexport) bool QAbstractScrollArea_EventFilter(QAbstractScrollArea* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) bool QAbstractScrollArea_Event(QAbstractScrollArea* self, QEvent* param1);
extern __declspec(dllexport) bool QAbstractScrollArea_ViewportEvent(QAbstractScrollArea* self, QEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_ResizeEvent(QAbstractScrollArea* self, QResizeEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_PaintEvent(QAbstractScrollArea* self, QPaintEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_MousePressEvent(QAbstractScrollArea* self, QMouseEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_MouseReleaseEvent(QAbstractScrollArea* self, QMouseEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_MouseDoubleClickEvent(QAbstractScrollArea* self, QMouseEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_MouseMoveEvent(QAbstractScrollArea* self, QMouseEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_WheelEvent(QAbstractScrollArea* self, QWheelEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_ContextMenuEvent(QAbstractScrollArea* self, QContextMenuEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_DragEnterEvent(QAbstractScrollArea* self, QDragEnterEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_DragMoveEvent(QAbstractScrollArea* self, QDragMoveEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_DragLeaveEvent(QAbstractScrollArea* self, QDragLeaveEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_DropEvent(QAbstractScrollArea* self, QDropEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_KeyPressEvent(QAbstractScrollArea* self, QKeyEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_ScrollContentsBy(QAbstractScrollArea* self, int dx, int dy);
extern __declspec(dllexport) QSize* QAbstractScrollArea_ViewportSizeHint(const QAbstractScrollArea* self);
extern __declspec(dllexport) struct miqt_string QAbstractScrollArea_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QAbstractScrollArea_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QAbstractScrollArea_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QAbstractScrollArea_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_SetupViewport(void* self, intptr_t slot);
void QAbstractScrollArea_virtualbase_SetupViewport(void* self, QWidget* viewport);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_EventFilter(void* self, intptr_t slot);
bool QAbstractScrollArea_virtualbase_EventFilter(void* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_Event(void* self, intptr_t slot);
bool QAbstractScrollArea_virtualbase_Event(void* self, QEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_ViewportEvent(void* self, intptr_t slot);
bool QAbstractScrollArea_virtualbase_ViewportEvent(void* self, QEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QAbstractScrollArea_virtualbase_ResizeEvent(void* self, QResizeEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_PaintEvent(void* self, intptr_t slot);
void QAbstractScrollArea_virtualbase_PaintEvent(void* self, QPaintEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QAbstractScrollArea_virtualbase_MousePressEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QAbstractScrollArea_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_MouseDoubleClickEvent(void* self, intptr_t slot);
void QAbstractScrollArea_virtualbase_MouseDoubleClickEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QAbstractScrollArea_virtualbase_MouseMoveEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_WheelEvent(void* self, intptr_t slot);
void QAbstractScrollArea_virtualbase_WheelEvent(void* self, QWheelEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QAbstractScrollArea_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_DragEnterEvent(void* self, intptr_t slot);
void QAbstractScrollArea_virtualbase_DragEnterEvent(void* self, QDragEnterEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_DragMoveEvent(void* self, intptr_t slot);
void QAbstractScrollArea_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_DragLeaveEvent(void* self, intptr_t slot);
void QAbstractScrollArea_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_DropEvent(void* self, intptr_t slot);
void QAbstractScrollArea_virtualbase_DropEvent(void* self, QDropEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QAbstractScrollArea_virtualbase_KeyPressEvent(void* self, QKeyEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_ScrollContentsBy(void* self, intptr_t slot);
void QAbstractScrollArea_virtualbase_ScrollContentsBy(void* self, int dx, int dy);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_ViewportSizeHint(void* self, intptr_t slot);
QSize* QAbstractScrollArea_virtualbase_ViewportSizeHint(const void* self);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QAbstractScrollArea_virtualbase_ChangeEvent(void* self, QEvent* param1);
extern __declspec(dllexport) void QAbstractScrollArea_override_virtual_InitStyleOption(void* self, intptr_t slot);
void QAbstractScrollArea_virtualbase_InitStyleOption(const void* self, QStyleOptionFrame* option);
extern __declspec(dllexport) void QAbstractScrollArea_Delete(QAbstractScrollArea* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
