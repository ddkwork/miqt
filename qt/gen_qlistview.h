#pragma once
#ifndef MIQT_QT_GEN_QLISTVIEW_H
#define MIQT_QT_GEN_QLISTVIEW_H

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
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QListView* QListView_new(QWidget* parent);
extern __declspec(dllexport) 
QListView* QListView_new2();
extern __declspec(dllexport) 
void QListView_virtbase(QListView* src
, QAbstractItemView** outptr_QAbstractItemView
);
extern __declspec(dllexport) 
QMetaObject* QListView_MetaObject(const QListView* self);
extern __declspec(dllexport) 
void* QListView_Metacast(QListView* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QListView_Tr(const char* s);
extern __declspec(dllexport) 
void QListView_SetMovement(QListView* self, Movement movement);
extern __declspec(dllexport) 
Movement QListView_Movement(const QListView* self);
extern __declspec(dllexport) 
void QListView_SetFlow(QListView* self, Flow flow);
extern __declspec(dllexport) 
Flow QListView_Flow(const QListView* self);
extern __declspec(dllexport) 
void QListView_SetWrapping(QListView* self, bool enable);
extern __declspec(dllexport) 
bool QListView_IsWrapping(const QListView* self);
extern __declspec(dllexport) 
void QListView_SetResizeMode(QListView* self, ResizeMode mode);
extern __declspec(dllexport) 
ResizeMode QListView_ResizeMode(const QListView* self);
extern __declspec(dllexport) 
void QListView_SetLayoutMode(QListView* self, LayoutMode mode);
extern __declspec(dllexport) 
LayoutMode QListView_LayoutMode(const QListView* self);
extern __declspec(dllexport) 
void QListView_SetSpacing(QListView* self, int space);
extern __declspec(dllexport) 
int QListView_Spacing(const QListView* self);
extern __declspec(dllexport) 
void QListView_SetBatchSize(QListView* self, int batchSize);
extern __declspec(dllexport) 
int QListView_BatchSize(const QListView* self);
extern __declspec(dllexport) 
void QListView_SetGridSize(QListView* self, QSize* size);
extern __declspec(dllexport) 
QSize* QListView_GridSize(const QListView* self);
extern __declspec(dllexport) 
void QListView_SetViewMode(QListView* self, ViewMode mode);
extern __declspec(dllexport) 
ViewMode QListView_ViewMode(const QListView* self);
extern __declspec(dllexport) 
void QListView_ClearPropertyFlags(QListView* self);
extern __declspec(dllexport) 
bool QListView_IsRowHidden(const QListView* self, int row);
extern __declspec(dllexport) 
void QListView_SetRowHidden(QListView* self, int row, bool hide);
extern __declspec(dllexport) 
void QListView_SetModelColumn(QListView* self, int column);
extern __declspec(dllexport) 
int QListView_ModelColumn(const QListView* self);
extern __declspec(dllexport) 
void QListView_SetUniformItemSizes(QListView* self, bool enable);
extern __declspec(dllexport) 
bool QListView_UniformItemSizes(const QListView* self);
extern __declspec(dllexport) 
void QListView_SetWordWrap(QListView* self, bool on);
extern __declspec(dllexport) 
bool QListView_WordWrap(const QListView* self);
extern __declspec(dllexport) 
void QListView_SetSelectionRectVisible(QListView* self, bool show);
extern __declspec(dllexport) 
bool QListView_IsSelectionRectVisible(const QListView* self);
extern __declspec(dllexport) 
void QListView_SetItemAlignment(QListView* self, int alignment);
extern __declspec(dllexport) 
int QListView_ItemAlignment(const QListView* self);
extern __declspec(dllexport) 
QRect* QListView_VisualRect(const QListView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QListView_ScrollTo(QListView* self, QModelIndex* index, ScrollHint hint);
extern __declspec(dllexport) 
QModelIndex* QListView_IndexAt(const QListView* self, QPoint* p);
extern __declspec(dllexport) 
void QListView_DoItemsLayout(QListView* self);
extern __declspec(dllexport) 
void QListView_Reset(QListView* self);
extern __declspec(dllexport) 
void QListView_SetRootIndex(QListView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QListView_IndexesMoved(QListView* self, struct miqt_array /* of QModelIndex* */  indexes);
void QListView_connect_IndexesMoved(QListView* self, intptr_t slot);
extern __declspec(dllexport) 
bool QListView_Event(QListView* self, QEvent* e);
extern __declspec(dllexport) 
void QListView_ScrollContentsBy(QListView* self, int dx, int dy);
extern __declspec(dllexport) 
void QListView_DataChanged(QListView* self, QModelIndex* topLeft, QModelIndex* bottomRight, struct miqt_array /* of int */  roles);
extern __declspec(dllexport) 
void QListView_RowsInserted(QListView* self, QModelIndex* parent, int start, int end);
extern __declspec(dllexport) 
void QListView_RowsAboutToBeRemoved(QListView* self, QModelIndex* parent, int start, int end);
extern __declspec(dllexport) 
void QListView_MouseMoveEvent(QListView* self, QMouseEvent* e);
extern __declspec(dllexport) 
void QListView_MouseReleaseEvent(QListView* self, QMouseEvent* e);
extern __declspec(dllexport) 
void QListView_WheelEvent(QListView* self, QWheelEvent* e);
extern __declspec(dllexport) 
void QListView_TimerEvent(QListView* self, QTimerEvent* e);
extern __declspec(dllexport) 
void QListView_ResizeEvent(QListView* self, QResizeEvent* e);
extern __declspec(dllexport) 
void QListView_DragMoveEvent(QListView* self, QDragMoveEvent* e);
extern __declspec(dllexport) 
void QListView_DragLeaveEvent(QListView* self, QDragLeaveEvent* e);
extern __declspec(dllexport) 
void QListView_DropEvent(QListView* self, QDropEvent* e);
extern __declspec(dllexport) 
void QListView_StartDrag(QListView* self, int supportedActions);
extern __declspec(dllexport) 
void QListView_InitViewItemOption(const QListView* self, QStyleOptionViewItem* option);
extern __declspec(dllexport) 
void QListView_PaintEvent(QListView* self, QPaintEvent* e);
extern __declspec(dllexport) 
int QListView_HorizontalOffset(const QListView* self);
extern __declspec(dllexport) 
int QListView_VerticalOffset(const QListView* self);
extern __declspec(dllexport) 
QModelIndex* QListView_MoveCursor(QListView* self, CursorAction cursorAction, int modifiers);
extern __declspec(dllexport) 
void QListView_SetSelection(QListView* self, QRect* rect, int command);
extern __declspec(dllexport) 
QRegion* QListView_VisualRegionForSelection(const QListView* self, QItemSelection* selection);
extern __declspec(dllexport) 
struct miqt_array /* of QModelIndex* */  QListView_SelectedIndexes(const QListView* self);
extern __declspec(dllexport) 
void QListView_UpdateGeometries(QListView* self);
extern __declspec(dllexport) 
bool QListView_IsIndexHidden(const QListView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QListView_SelectionChanged(QListView* self, QItemSelection* selected, QItemSelection* deselected);
extern __declspec(dllexport) 
void QListView_CurrentChanged(QListView* self, QModelIndex* current, QModelIndex* previous);
extern __declspec(dllexport) 
QSize* QListView_ViewportSizeHint(const QListView* self);
extern __declspec(dllexport) 
struct miqt_string QListView_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QListView_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QListView_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QListView_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QListView_override_virtual_Metacast(void* self, intptr_t slot);
void* QListView_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QListView_Delete(QListView* self, bool isSubclass);

}
