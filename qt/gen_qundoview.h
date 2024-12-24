#pragma once
#ifndef MIQT_QT_GEN_QUNDOVIEW_H
#define MIQT_QT_GEN_QUNDOVIEW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemView QAbstractItemView;
typedef struct QAbstractScrollArea QAbstractScrollArea;
typedef struct QDragLeaveEvent QDragLeaveEvent;
typedef struct QDragMoveEvent QDragMoveEvent;
typedef struct QDropEvent QDropEvent;
typedef struct QEvent QEvent;
typedef struct QFrame QFrame;
typedef struct QIcon QIcon;
typedef struct QItemSelection QItemSelection;
typedef struct QListView QListView;
typedef struct QMetaObject QMetaObject;
typedef struct QModelIndex QModelIndex;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QRegion QRegion;
typedef struct QResizeEvent QResizeEvent;
typedef struct QSize QSize;
typedef struct QStyleOptionViewItem QStyleOptionViewItem;
typedef struct QTimerEvent QTimerEvent;
typedef struct QUndoGroup QUndoGroup;
typedef struct QUndoStack QUndoStack;
typedef struct QUndoView QUndoView;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QUndoView* QUndoView_new(QWidget* parent);
extern __declspec(dllexport) QUndoView* QUndoView_new2();
extern __declspec(dllexport) QUndoView* QUndoView_new3(QUndoStack* stack);
extern __declspec(dllexport) QUndoView* QUndoView_new4(QUndoGroup* group);
extern __declspec(dllexport) QUndoView* QUndoView_new5(QUndoStack* stack, QWidget* parent);
extern __declspec(dllexport) QUndoView* QUndoView_new6(QUndoGroup* group, QWidget* parent);
extern __declspec(dllexport) void QUndoView_virtbase(QUndoView* src, QListView** outptr_QListView);
extern __declspec(dllexport) QMetaObject* QUndoView_MetaObject(const QUndoView* self);
extern __declspec(dllexport) void* QUndoView_Metacast(QUndoView* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QUndoView_Tr(const char* s);
extern __declspec(dllexport) QUndoStack* QUndoView_Stack(const QUndoView* self);
extern __declspec(dllexport) QUndoGroup* QUndoView_Group(const QUndoView* self);
extern __declspec(dllexport) void QUndoView_SetEmptyLabel(QUndoView* self, struct miqt_string label);
extern __declspec(dllexport) struct miqt_string QUndoView_EmptyLabel(const QUndoView* self);
extern __declspec(dllexport) void QUndoView_SetCleanIcon(QUndoView* self, QIcon* icon);
extern __declspec(dllexport) QIcon* QUndoView_CleanIcon(const QUndoView* self);
extern __declspec(dllexport) void QUndoView_SetStack(QUndoView* self, QUndoStack* stack);
extern __declspec(dllexport) void QUndoView_SetGroup(QUndoView* self, QUndoGroup* group);
extern __declspec(dllexport) struct miqt_string QUndoView_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QUndoView_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QUndoView_override_virtual_VisualRect(void* self, intptr_t slot);
QRect* QUndoView_virtualbase_VisualRect(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QUndoView_override_virtual_ScrollTo(void* self, intptr_t slot);
void QUndoView_virtualbase_ScrollTo(void* self, QModelIndex* index, ScrollHint hint);
extern __declspec(dllexport) void QUndoView_override_virtual_IndexAt(void* self, intptr_t slot);
QModelIndex* QUndoView_virtualbase_IndexAt(const void* self, QPoint* p);
extern __declspec(dllexport) void QUndoView_override_virtual_DoItemsLayout(void* self, intptr_t slot);
void QUndoView_virtualbase_DoItemsLayout(void* self);
extern __declspec(dllexport) void QUndoView_override_virtual_Reset(void* self, intptr_t slot);
void QUndoView_virtualbase_Reset(void* self);
extern __declspec(dllexport) void QUndoView_override_virtual_SetRootIndex(void* self, intptr_t slot);
void QUndoView_virtualbase_SetRootIndex(void* self, QModelIndex* index);
extern __declspec(dllexport) void QUndoView_override_virtual_Event(void* self, intptr_t slot);
bool QUndoView_virtualbase_Event(void* self, QEvent* e);
extern __declspec(dllexport) void QUndoView_override_virtual_ScrollContentsBy(void* self, intptr_t slot);
void QUndoView_virtualbase_ScrollContentsBy(void* self, int dx, int dy);
extern __declspec(dllexport) void QUndoView_override_virtual_DataChanged(void* self, intptr_t slot);
void QUndoView_virtualbase_DataChanged(void* self, QModelIndex* topLeft, QModelIndex* bottomRight, struct miqt_array /* of int */  roles);
extern __declspec(dllexport) void QUndoView_override_virtual_RowsInserted(void* self, intptr_t slot);
void QUndoView_virtualbase_RowsInserted(void* self, QModelIndex* parent, int start, int end);
extern __declspec(dllexport) void QUndoView_override_virtual_RowsAboutToBeRemoved(void* self, intptr_t slot);
void QUndoView_virtualbase_RowsAboutToBeRemoved(void* self, QModelIndex* parent, int start, int end);
extern __declspec(dllexport) void QUndoView_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QUndoView_virtualbase_MouseMoveEvent(void* self, QMouseEvent* e);
extern __declspec(dllexport) void QUndoView_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QUndoView_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* e);
extern __declspec(dllexport) void QUndoView_override_virtual_WheelEvent(void* self, intptr_t slot);
void QUndoView_virtualbase_WheelEvent(void* self, QWheelEvent* e);
extern __declspec(dllexport) void QUndoView_override_virtual_TimerEvent(void* self, intptr_t slot);
void QUndoView_virtualbase_TimerEvent(void* self, QTimerEvent* e);
extern __declspec(dllexport) void QUndoView_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QUndoView_virtualbase_ResizeEvent(void* self, QResizeEvent* e);
extern __declspec(dllexport) void QUndoView_override_virtual_DragMoveEvent(void* self, intptr_t slot);
void QUndoView_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* e);
extern __declspec(dllexport) void QUndoView_override_virtual_DragLeaveEvent(void* self, intptr_t slot);
void QUndoView_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* e);
extern __declspec(dllexport) void QUndoView_override_virtual_DropEvent(void* self, intptr_t slot);
void QUndoView_virtualbase_DropEvent(void* self, QDropEvent* e);
extern __declspec(dllexport) void QUndoView_override_virtual_StartDrag(void* self, intptr_t slot);
void QUndoView_virtualbase_StartDrag(void* self, int supportedActions);
extern __declspec(dllexport) void QUndoView_override_virtual_InitViewItemOption(void* self, intptr_t slot);
void QUndoView_virtualbase_InitViewItemOption(const void* self, QStyleOptionViewItem* option);
extern __declspec(dllexport) void QUndoView_override_virtual_PaintEvent(void* self, intptr_t slot);
void QUndoView_virtualbase_PaintEvent(void* self, QPaintEvent* e);
extern __declspec(dllexport) void QUndoView_override_virtual_HorizontalOffset(void* self, intptr_t slot);
int QUndoView_virtualbase_HorizontalOffset(const void* self);
extern __declspec(dllexport) void QUndoView_override_virtual_VerticalOffset(void* self, intptr_t slot);
int QUndoView_virtualbase_VerticalOffset(const void* self);
extern __declspec(dllexport) void QUndoView_override_virtual_MoveCursor(void* self, intptr_t slot);
QModelIndex* QUndoView_virtualbase_MoveCursor(void* self, CursorAction cursorAction, int modifiers);
extern __declspec(dllexport) void QUndoView_override_virtual_SetSelection(void* self, intptr_t slot);
void QUndoView_virtualbase_SetSelection(void* self, QRect* rect, int command);
extern __declspec(dllexport) void QUndoView_override_virtual_VisualRegionForSelection(void* self, intptr_t slot);
QRegion* QUndoView_virtualbase_VisualRegionForSelection(const void* self, QItemSelection* selection);
extern __declspec(dllexport) void QUndoView_override_virtual_SelectedIndexes(void* self, intptr_t slot);
struct miqt_array /* of QModelIndex* */  QUndoView_virtualbase_SelectedIndexes(const void* self);
extern __declspec(dllexport) void QUndoView_override_virtual_UpdateGeometries(void* self, intptr_t slot);
void QUndoView_virtualbase_UpdateGeometries(void* self);
extern __declspec(dllexport) void QUndoView_override_virtual_IsIndexHidden(void* self, intptr_t slot);
bool QUndoView_virtualbase_IsIndexHidden(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QUndoView_override_virtual_SelectionChanged(void* self, intptr_t slot);
void QUndoView_virtualbase_SelectionChanged(void* self, QItemSelection* selected, QItemSelection* deselected);
extern __declspec(dllexport) void QUndoView_override_virtual_CurrentChanged(void* self, intptr_t slot);
void QUndoView_virtualbase_CurrentChanged(void* self, QModelIndex* current, QModelIndex* previous);
extern __declspec(dllexport) void QUndoView_override_virtual_ViewportSizeHint(void* self, intptr_t slot);
QSize* QUndoView_virtualbase_ViewportSizeHint(const void* self);
extern __declspec(dllexport) void QUndoView_Delete(QUndoView* self, bool isSubclass);

} 
