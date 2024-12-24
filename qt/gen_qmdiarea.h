#pragma once
#ifndef MIQT_QT_GEN_QMDIAREA_H
#define MIQT_QT_GEN_QMDIAREA_H

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
class QBrush;
class QChildEvent;
class QContextMenuEvent;
class QDragEnterEvent;
class QDragLeaveEvent;
class QDragMoveEvent;
class QDropEvent;
class QEvent;
class QFrame;
class QKeyEvent;
class QMdiArea;
class QMdiSubWindow;
class QMetaObject;
class QMouseEvent;
class QObject;
class QPaintDevice;
class QPaintEvent;
class QResizeEvent;
class QShowEvent;
class QSize;
class QTimerEvent;
class QWheelEvent;
class QWidget;
class _GUID;
class type_info;
#else
typedef struct QAbstractScrollArea QAbstractScrollArea;
typedef struct QBrush QBrush;
typedef struct QChildEvent QChildEvent;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDragEnterEvent QDragEnterEvent;
typedef struct QDragLeaveEvent QDragLeaveEvent;
typedef struct QDragMoveEvent QDragMoveEvent;
typedef struct QDropEvent QDropEvent;
typedef struct QEvent QEvent;
typedef struct QFrame QFrame;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMdiArea QMdiArea;
typedef struct QMdiSubWindow QMdiSubWindow;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QMdiArea* QMdiArea_new(QWidget* parent);
extern __declspec(dllexport) QMdiArea* QMdiArea_new2();
extern __declspec(dllexport) void QMdiArea_virtbase(QMdiArea* src, QAbstractScrollArea** outptr_QAbstractScrollArea);
extern __declspec(dllexport) QMetaObject* QMdiArea_MetaObject(const QMdiArea* self);
extern __declspec(dllexport) void* QMdiArea_Metacast(QMdiArea* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QMdiArea_Tr(const char* s);
extern __declspec(dllexport) QSize* QMdiArea_SizeHint(const QMdiArea* self);
extern __declspec(dllexport) QSize* QMdiArea_MinimumSizeHint(const QMdiArea* self);
extern __declspec(dllexport) QMdiSubWindow* QMdiArea_CurrentSubWindow(const QMdiArea* self);
extern __declspec(dllexport) QMdiSubWindow* QMdiArea_ActiveSubWindow(const QMdiArea* self);
extern __declspec(dllexport) struct miqt_array /* of QMdiSubWindow* */  QMdiArea_SubWindowList(const QMdiArea* self);
extern __declspec(dllexport) QMdiSubWindow* QMdiArea_AddSubWindow(QMdiArea* self, QWidget* widget);
extern __declspec(dllexport) void QMdiArea_RemoveSubWindow(QMdiArea* self, QWidget* widget);
extern __declspec(dllexport) QBrush* QMdiArea_Background(const QMdiArea* self);
extern __declspec(dllexport) void QMdiArea_SetBackground(QMdiArea* self, QBrush* background);
extern __declspec(dllexport) WindowOrder QMdiArea_ActivationOrder(const QMdiArea* self);
extern __declspec(dllexport) void QMdiArea_SetActivationOrder(QMdiArea* self, WindowOrder order);
extern __declspec(dllexport) void QMdiArea_SetOption(QMdiArea* self, AreaOption option);
extern __declspec(dllexport) bool QMdiArea_TestOption(const QMdiArea* self, AreaOption opton);
extern __declspec(dllexport) void QMdiArea_SetViewMode(QMdiArea* self, ViewMode mode);
extern __declspec(dllexport) ViewMode QMdiArea_ViewMode(const QMdiArea* self);
extern __declspec(dllexport) bool QMdiArea_DocumentMode(const QMdiArea* self);
extern __declspec(dllexport) void QMdiArea_SetDocumentMode(QMdiArea* self, bool enabled);
extern __declspec(dllexport) void QMdiArea_SetTabsClosable(QMdiArea* self, bool closable);
extern __declspec(dllexport) bool QMdiArea_TabsClosable(const QMdiArea* self);
extern __declspec(dllexport) void QMdiArea_SetTabsMovable(QMdiArea* self, bool movable);
extern __declspec(dllexport) bool QMdiArea_TabsMovable(const QMdiArea* self);
extern __declspec(dllexport) void QMdiArea_SetTabShape(QMdiArea* self, int shape);
extern __declspec(dllexport) int QMdiArea_TabShape(const QMdiArea* self);
extern __declspec(dllexport) void QMdiArea_SetTabPosition(QMdiArea* self, int position);
extern __declspec(dllexport) int QMdiArea_TabPosition(const QMdiArea* self);
extern __declspec(dllexport) void QMdiArea_SubWindowActivated(QMdiArea* self, QMdiSubWindow* param1);
void QMdiArea_connect_SubWindowActivated(QMdiArea* self, intptr_t slot);
extern __declspec(dllexport) void QMdiArea_SetActiveSubWindow(QMdiArea* self, QMdiSubWindow* window);
extern __declspec(dllexport) void QMdiArea_TileSubWindows(QMdiArea* self);
extern __declspec(dllexport) void QMdiArea_CascadeSubWindows(QMdiArea* self);
extern __declspec(dllexport) void QMdiArea_CloseActiveSubWindow(QMdiArea* self);
extern __declspec(dllexport) void QMdiArea_CloseAllSubWindows(QMdiArea* self);
extern __declspec(dllexport) void QMdiArea_ActivateNextSubWindow(QMdiArea* self);
extern __declspec(dllexport) void QMdiArea_ActivatePreviousSubWindow(QMdiArea* self);
extern __declspec(dllexport) void QMdiArea_SetupViewport(QMdiArea* self, QWidget* viewport);
extern __declspec(dllexport) bool QMdiArea_Event(QMdiArea* self, QEvent* event);
extern __declspec(dllexport) bool QMdiArea_EventFilter(QMdiArea* self, QObject* object, QEvent* event);
extern __declspec(dllexport) void QMdiArea_PaintEvent(QMdiArea* self, QPaintEvent* paintEvent);
extern __declspec(dllexport) void QMdiArea_ChildEvent(QMdiArea* self, QChildEvent* childEvent);
extern __declspec(dllexport) void QMdiArea_ResizeEvent(QMdiArea* self, QResizeEvent* resizeEvent);
extern __declspec(dllexport) void QMdiArea_TimerEvent(QMdiArea* self, QTimerEvent* timerEvent);
extern __declspec(dllexport) void QMdiArea_ShowEvent(QMdiArea* self, QShowEvent* showEvent);
extern __declspec(dllexport) bool QMdiArea_ViewportEvent(QMdiArea* self, QEvent* event);
extern __declspec(dllexport) void QMdiArea_ScrollContentsBy(QMdiArea* self, int dx, int dy);
extern __declspec(dllexport) struct miqt_string QMdiArea_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QMdiArea_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) struct miqt_array /* of QMdiSubWindow* */  QMdiArea_SubWindowList1(const QMdiArea* self, WindowOrder order);
extern __declspec(dllexport) QMdiSubWindow* QMdiArea_AddSubWindow2(QMdiArea* self, QWidget* widget, int flags);
extern __declspec(dllexport) void QMdiArea_SetOption2(QMdiArea* self, AreaOption option, bool on);
extern __declspec(dllexport) void QMdiArea_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QMdiArea_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QMdiArea_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QMdiArea_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QMdiArea_override_virtual_SetupViewport(void* self, intptr_t slot);
void QMdiArea_virtualbase_SetupViewport(void* self, QWidget* viewport);
extern __declspec(dllexport) void QMdiArea_override_virtual_Event(void* self, intptr_t slot);
bool QMdiArea_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QMdiArea_override_virtual_EventFilter(void* self, intptr_t slot);
bool QMdiArea_virtualbase_EventFilter(void* self, QObject* object, QEvent* event);
extern __declspec(dllexport) void QMdiArea_override_virtual_PaintEvent(void* self, intptr_t slot);
void QMdiArea_virtualbase_PaintEvent(void* self, QPaintEvent* paintEvent);
extern __declspec(dllexport) void QMdiArea_override_virtual_ChildEvent(void* self, intptr_t slot);
void QMdiArea_virtualbase_ChildEvent(void* self, QChildEvent* childEvent);
extern __declspec(dllexport) void QMdiArea_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QMdiArea_virtualbase_ResizeEvent(void* self, QResizeEvent* resizeEvent);
extern __declspec(dllexport) void QMdiArea_override_virtual_TimerEvent(void* self, intptr_t slot);
void QMdiArea_virtualbase_TimerEvent(void* self, QTimerEvent* timerEvent);
extern __declspec(dllexport) void QMdiArea_override_virtual_ShowEvent(void* self, intptr_t slot);
void QMdiArea_virtualbase_ShowEvent(void* self, QShowEvent* showEvent);
extern __declspec(dllexport) void QMdiArea_override_virtual_ViewportEvent(void* self, intptr_t slot);
bool QMdiArea_virtualbase_ViewportEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QMdiArea_override_virtual_ScrollContentsBy(void* self, intptr_t slot);
void QMdiArea_virtualbase_ScrollContentsBy(void* self, int dx, int dy);
extern __declspec(dllexport) void QMdiArea_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QMdiArea_virtualbase_MousePressEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QMdiArea_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QMdiArea_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QMdiArea_override_virtual_MouseDoubleClickEvent(void* self, intptr_t slot);
void QMdiArea_virtualbase_MouseDoubleClickEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QMdiArea_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QMdiArea_virtualbase_MouseMoveEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QMdiArea_override_virtual_WheelEvent(void* self, intptr_t slot);
void QMdiArea_virtualbase_WheelEvent(void* self, QWheelEvent* param1);
extern __declspec(dllexport) void QMdiArea_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QMdiArea_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* param1);
extern __declspec(dllexport) void QMdiArea_override_virtual_DragEnterEvent(void* self, intptr_t slot);
void QMdiArea_virtualbase_DragEnterEvent(void* self, QDragEnterEvent* param1);
extern __declspec(dllexport) void QMdiArea_override_virtual_DragMoveEvent(void* self, intptr_t slot);
void QMdiArea_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* param1);
extern __declspec(dllexport) void QMdiArea_override_virtual_DragLeaveEvent(void* self, intptr_t slot);
void QMdiArea_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* param1);
extern __declspec(dllexport) void QMdiArea_override_virtual_DropEvent(void* self, intptr_t slot);
void QMdiArea_virtualbase_DropEvent(void* self, QDropEvent* param1);
extern __declspec(dllexport) void QMdiArea_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QMdiArea_virtualbase_KeyPressEvent(void* self, QKeyEvent* param1);
extern __declspec(dllexport) void QMdiArea_override_virtual_ViewportSizeHint(void* self, intptr_t slot);
QSize* QMdiArea_virtualbase_ViewportSizeHint(const void* self);
extern __declspec(dllexport) void QMdiArea_Delete(QMdiArea* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
