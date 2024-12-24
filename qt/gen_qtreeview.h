#pragma once
#ifndef MIQT_QT_GEN_QTREEVIEW_H
#define MIQT_QT_GEN_QTREEVIEW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QAbstractItemView QAbstractItemView;
typedef struct QAbstractScrollArea QAbstractScrollArea;
typedef struct QDragMoveEvent QDragMoveEvent;
typedef struct QEvent QEvent;
typedef struct QFrame QFrame;
typedef struct QHeaderView QHeaderView;
typedef struct QItemSelection QItemSelection;
typedef struct QItemSelectionModel QItemSelectionModel;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QModelIndex QModelIndex;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPainter QPainter;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QRegion QRegion;
typedef struct QSize QSize;
typedef struct QStyleOptionViewItem QStyleOptionViewItem;
typedef struct QTimerEvent QTimerEvent;
typedef struct QTreeView QTreeView;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QTreeView* QTreeView_new(QWidget* parent);
extern __declspec(dllexport) 
QTreeView* QTreeView_new2();
extern __declspec(dllexport) 
void QTreeView_virtbase(QTreeView* src
, QAbstractItemView** outptr_QAbstractItemView
);
extern __declspec(dllexport) 
QMetaObject* QTreeView_MetaObject(const QTreeView* self);
extern __declspec(dllexport) 
void* QTreeView_Metacast(QTreeView* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QTreeView_Tr(const char* s);
extern __declspec(dllexport) 
void QTreeView_SetModel(QTreeView* self, QAbstractItemModel* model);
extern __declspec(dllexport) 
void QTreeView_SetRootIndex(QTreeView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QTreeView_SetSelectionModel(QTreeView* self, QItemSelectionModel* selectionModel);
extern __declspec(dllexport) 
QHeaderView* QTreeView_Header(const QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_SetHeader(QTreeView* self, QHeaderView* header);
extern __declspec(dllexport) 
int QTreeView_AutoExpandDelay(const QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_SetAutoExpandDelay(QTreeView* self, int delay);
extern __declspec(dllexport) 
int QTreeView_Indentation(const QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_SetIndentation(QTreeView* self, int i);
extern __declspec(dllexport) 
void QTreeView_ResetIndentation(QTreeView* self);
extern __declspec(dllexport) 
bool QTreeView_RootIsDecorated(const QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_SetRootIsDecorated(QTreeView* self, bool show);
extern __declspec(dllexport) 
bool QTreeView_UniformRowHeights(const QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_SetUniformRowHeights(QTreeView* self, bool uniform);
extern __declspec(dllexport) 
bool QTreeView_ItemsExpandable(const QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_SetItemsExpandable(QTreeView* self, bool enable);
extern __declspec(dllexport) 
bool QTreeView_ExpandsOnDoubleClick(const QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_SetExpandsOnDoubleClick(QTreeView* self, bool enable);
extern __declspec(dllexport) 
int QTreeView_ColumnViewportPosition(const QTreeView* self, int column);
extern __declspec(dllexport) 
int QTreeView_ColumnWidth(const QTreeView* self, int column);
extern __declspec(dllexport) 
void QTreeView_SetColumnWidth(QTreeView* self, int column, int width);
extern __declspec(dllexport) 
int QTreeView_ColumnAt(const QTreeView* self, int x);
extern __declspec(dllexport) 
bool QTreeView_IsColumnHidden(const QTreeView* self, int column);
extern __declspec(dllexport) 
void QTreeView_SetColumnHidden(QTreeView* self, int column, bool hide);
extern __declspec(dllexport) 
bool QTreeView_IsHeaderHidden(const QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_SetHeaderHidden(QTreeView* self, bool hide);
extern __declspec(dllexport) 
bool QTreeView_IsRowHidden(const QTreeView* self, int row, QModelIndex* parent);
extern __declspec(dllexport) 
void QTreeView_SetRowHidden(QTreeView* self, int row, QModelIndex* parent, bool hide);
extern __declspec(dllexport) 
bool QTreeView_IsFirstColumnSpanned(const QTreeView* self, int row, QModelIndex* parent);
extern __declspec(dllexport) 
void QTreeView_SetFirstColumnSpanned(QTreeView* self, int row, QModelIndex* parent, bool span);
extern __declspec(dllexport) 
bool QTreeView_IsExpanded(const QTreeView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QTreeView_SetExpanded(QTreeView* self, QModelIndex* index, bool expand);
extern __declspec(dllexport) 
void QTreeView_SetSortingEnabled(QTreeView* self, bool enable);
extern __declspec(dllexport) 
bool QTreeView_IsSortingEnabled(const QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_SetAnimated(QTreeView* self, bool enable);
extern __declspec(dllexport) 
bool QTreeView_IsAnimated(const QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_SetAllColumnsShowFocus(QTreeView* self, bool enable);
extern __declspec(dllexport) 
bool QTreeView_AllColumnsShowFocus(const QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_SetWordWrap(QTreeView* self, bool on);
extern __declspec(dllexport) 
bool QTreeView_WordWrap(const QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_SetTreePosition(QTreeView* self, int logicalIndex);
extern __declspec(dllexport) 
int QTreeView_TreePosition(const QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_KeyboardSearch(QTreeView* self, struct miqt_string search);
extern __declspec(dllexport) 
QRect* QTreeView_VisualRect(const QTreeView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QTreeView_ScrollTo(QTreeView* self, QModelIndex* index, ScrollHint hint);
extern __declspec(dllexport) 
QModelIndex* QTreeView_IndexAt(const QTreeView* self, QPoint* p);
extern __declspec(dllexport) 
QModelIndex* QTreeView_IndexAbove(const QTreeView* self, QModelIndex* index);
extern __declspec(dllexport) 
QModelIndex* QTreeView_IndexBelow(const QTreeView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QTreeView_DoItemsLayout(QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_Reset(QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_DataChanged(QTreeView* self, QModelIndex* topLeft, QModelIndex* bottomRight, struct miqt_array /* of int */  roles);
extern __declspec(dllexport) 
void QTreeView_SelectAll(QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_Expanded(QTreeView* self, QModelIndex* index);
void QTreeView_connect_Expanded(QTreeView* self, intptr_t slot);
extern __declspec(dllexport) 
void QTreeView_Collapsed(QTreeView* self, QModelIndex* index);
void QTreeView_connect_Collapsed(QTreeView* self, intptr_t slot);
extern __declspec(dllexport) 
void QTreeView_HideColumn(QTreeView* self, int column);
extern __declspec(dllexport) 
void QTreeView_ShowColumn(QTreeView* self, int column);
extern __declspec(dllexport) 
void QTreeView_Expand(QTreeView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QTreeView_Collapse(QTreeView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QTreeView_ResizeColumnToContents(QTreeView* self, int column);
extern __declspec(dllexport) 
void QTreeView_SortByColumn(QTreeView* self, int column, int order);
extern __declspec(dllexport) 
void QTreeView_ExpandAll(QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_ExpandRecursively(QTreeView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QTreeView_CollapseAll(QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_ExpandToDepth(QTreeView* self, int depth);
extern __declspec(dllexport) 
void QTreeView_VerticalScrollbarValueChanged(QTreeView* self, int value);
extern __declspec(dllexport) 
void QTreeView_ScrollContentsBy(QTreeView* self, int dx, int dy);
extern __declspec(dllexport) 
void QTreeView_RowsInserted(QTreeView* self, QModelIndex* parent, int start, int end);
extern __declspec(dllexport) 
void QTreeView_RowsAboutToBeRemoved(QTreeView* self, QModelIndex* parent, int start, int end);
extern __declspec(dllexport) 
QModelIndex* QTreeView_MoveCursor(QTreeView* self, CursorAction cursorAction, int modifiers);
extern __declspec(dllexport) 
int QTreeView_HorizontalOffset(const QTreeView* self);
extern __declspec(dllexport) 
int QTreeView_VerticalOffset(const QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_SetSelection(QTreeView* self, QRect* rect, int command);
extern __declspec(dllexport) 
QRegion* QTreeView_VisualRegionForSelection(const QTreeView* self, QItemSelection* selection);
extern __declspec(dllexport) 
struct miqt_array /* of QModelIndex* */  QTreeView_SelectedIndexes(const QTreeView* self);
extern __declspec(dllexport) 
void QTreeView_ChangeEvent(QTreeView* self, QEvent* event);
extern __declspec(dllexport) 
void QTreeView_TimerEvent(QTreeView* self, QTimerEvent* event);
extern __declspec(dllexport) 
void QTreeView_PaintEvent(QTreeView* self, QPaintEvent* event);
extern __declspec(dllexport) 
void QTreeView_DrawRow(const QTreeView* self, QPainter* painter, QStyleOptionViewItem* options, QModelIndex* index);
extern __declspec(dllexport) 
void QTreeView_DrawBranches(const QTreeView* self, QPainter* painter, QRect* rect, QModelIndex* index);
extern __declspec(dllexport) 
void QTreeView_MousePressEvent(QTreeView* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QTreeView_MouseReleaseEvent(QTreeView* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QTreeView_MouseDoubleClickEvent(QTreeView* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QTreeView_MouseMoveEvent(QTreeView* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QTreeView_KeyPressEvent(QTreeView* self, QKeyEvent* event);
extern __declspec(dllexport) 
void QTreeView_DragMoveEvent(QTreeView* self, QDragMoveEvent* event);
extern __declspec(dllexport) 
bool QTreeView_ViewportEvent(QTreeView* self, QEvent* event);
extern __declspec(dllexport) 
void QTreeView_UpdateGeometries(QTreeView* self);
extern __declspec(dllexport) 
QSize* QTreeView_ViewportSizeHint(const QTreeView* self);
extern __declspec(dllexport) 
int QTreeView_SizeHintForColumn(const QTreeView* self, int column);
extern __declspec(dllexport) 
void QTreeView_HorizontalScrollbarAction(QTreeView* self, int action);
extern __declspec(dllexport) 
bool QTreeView_IsIndexHidden(const QTreeView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QTreeView_SelectionChanged(QTreeView* self, QItemSelection* selected, QItemSelection* deselected);
extern __declspec(dllexport) 
void QTreeView_CurrentChanged(QTreeView* self, QModelIndex* current, QModelIndex* previous);
extern __declspec(dllexport) 
struct miqt_string QTreeView_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QTreeView_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QTreeView_ExpandRecursively2(QTreeView* self, QModelIndex* index, int depth);
extern __declspec(dllexport) 
void QTreeView_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QTreeView_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QTreeView_override_virtual_Metacast(void* self, intptr_t slot);
void* QTreeView_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QTreeView_Delete(QTreeView* self, bool isSubclass);

}
