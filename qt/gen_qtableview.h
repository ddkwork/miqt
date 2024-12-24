#pragma once
#ifndef MIQT_QT_GEN_QTABLEVIEW_H
#define MIQT_QT_GEN_QTABLEVIEW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QAbstractItemView QAbstractItemView;
typedef struct QAbstractScrollArea QAbstractScrollArea;
typedef struct QDropEvent QDropEvent;
typedef struct QFrame QFrame;
typedef struct QHeaderView QHeaderView;
typedef struct QItemSelection QItemSelection;
typedef struct QItemSelectionModel QItemSelectionModel;
typedef struct QMetaObject QMetaObject;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QRegion QRegion;
typedef struct QSize QSize;
typedef struct QStyleOptionViewItem QStyleOptionViewItem;
typedef struct QTableView QTableView;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QTableView* QTableView_new(QWidget* parent);
extern __declspec(dllexport) 
QTableView* QTableView_new2();
extern __declspec(dllexport) 
void QTableView_virtbase(QTableView* src
, QAbstractItemView** outptr_QAbstractItemView
);
extern __declspec(dllexport) 
QMetaObject* QTableView_MetaObject(const QTableView* self);
extern __declspec(dllexport) 
void* QTableView_Metacast(QTableView* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QTableView_Tr(const char* s);
extern __declspec(dllexport) 
void QTableView_SetModel(QTableView* self, QAbstractItemModel* model);
extern __declspec(dllexport) 
void QTableView_SetRootIndex(QTableView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QTableView_SetSelectionModel(QTableView* self, QItemSelectionModel* selectionModel);
extern __declspec(dllexport) 
void QTableView_DoItemsLayout(QTableView* self);
extern __declspec(dllexport) 
QHeaderView* QTableView_HorizontalHeader(const QTableView* self);
extern __declspec(dllexport) 
QHeaderView* QTableView_VerticalHeader(const QTableView* self);
extern __declspec(dllexport) 
void QTableView_SetHorizontalHeader(QTableView* self, QHeaderView* header);
extern __declspec(dllexport) 
void QTableView_SetVerticalHeader(QTableView* self, QHeaderView* header);
extern __declspec(dllexport) 
int QTableView_RowViewportPosition(const QTableView* self, int row);
extern __declspec(dllexport) 
int QTableView_RowAt(const QTableView* self, int y);
extern __declspec(dllexport) 
void QTableView_SetRowHeight(QTableView* self, int row, int height);
extern __declspec(dllexport) 
int QTableView_RowHeight(const QTableView* self, int row);
extern __declspec(dllexport) 
int QTableView_ColumnViewportPosition(const QTableView* self, int column);
extern __declspec(dllexport) 
int QTableView_ColumnAt(const QTableView* self, int x);
extern __declspec(dllexport) 
void QTableView_SetColumnWidth(QTableView* self, int column, int width);
extern __declspec(dllexport) 
int QTableView_ColumnWidth(const QTableView* self, int column);
extern __declspec(dllexport) 
bool QTableView_IsRowHidden(const QTableView* self, int row);
extern __declspec(dllexport) 
void QTableView_SetRowHidden(QTableView* self, int row, bool hide);
extern __declspec(dllexport) 
bool QTableView_IsColumnHidden(const QTableView* self, int column);
extern __declspec(dllexport) 
void QTableView_SetColumnHidden(QTableView* self, int column, bool hide);
extern __declspec(dllexport) 
void QTableView_SetSortingEnabled(QTableView* self, bool enable);
extern __declspec(dllexport) 
bool QTableView_IsSortingEnabled(const QTableView* self);
extern __declspec(dllexport) 
bool QTableView_ShowGrid(const QTableView* self);
extern __declspec(dllexport) 
int QTableView_GridStyle(const QTableView* self);
extern __declspec(dllexport) 
void QTableView_SetGridStyle(QTableView* self, int style);
extern __declspec(dllexport) 
void QTableView_SetWordWrap(QTableView* self, bool on);
extern __declspec(dllexport) 
bool QTableView_WordWrap(const QTableView* self);
extern __declspec(dllexport) 
void QTableView_SetCornerButtonEnabled(QTableView* self, bool enable);
extern __declspec(dllexport) 
bool QTableView_IsCornerButtonEnabled(const QTableView* self);
extern __declspec(dllexport) 
QRect* QTableView_VisualRect(const QTableView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QTableView_ScrollTo(QTableView* self, QModelIndex* index, ScrollHint hint);
extern __declspec(dllexport) 
QModelIndex* QTableView_IndexAt(const QTableView* self, QPoint* p);
extern __declspec(dllexport) 
void QTableView_SetSpan(QTableView* self, int row, int column, int rowSpan, int columnSpan);
extern __declspec(dllexport) 
int QTableView_RowSpan(const QTableView* self, int row, int column);
extern __declspec(dllexport) 
int QTableView_ColumnSpan(const QTableView* self, int row, int column);
extern __declspec(dllexport) 
void QTableView_ClearSpans(QTableView* self);
extern __declspec(dllexport) 
void QTableView_SelectRow(QTableView* self, int row);
extern __declspec(dllexport) 
void QTableView_SelectColumn(QTableView* self, int column);
extern __declspec(dllexport) 
void QTableView_HideRow(QTableView* self, int row);
extern __declspec(dllexport) 
void QTableView_HideColumn(QTableView* self, int column);
extern __declspec(dllexport) 
void QTableView_ShowRow(QTableView* self, int row);
extern __declspec(dllexport) 
void QTableView_ShowColumn(QTableView* self, int column);
extern __declspec(dllexport) 
void QTableView_ResizeRowToContents(QTableView* self, int row);
extern __declspec(dllexport) 
void QTableView_ResizeRowsToContents(QTableView* self);
extern __declspec(dllexport) 
void QTableView_ResizeColumnToContents(QTableView* self, int column);
extern __declspec(dllexport) 
void QTableView_ResizeColumnsToContents(QTableView* self);
extern __declspec(dllexport) 
void QTableView_SortByColumn(QTableView* self, int column, int order);
extern __declspec(dllexport) 
void QTableView_SetShowGrid(QTableView* self, bool show);
extern __declspec(dllexport) 
void QTableView_ScrollContentsBy(QTableView* self, int dx, int dy);
extern __declspec(dllexport) 
void QTableView_InitViewItemOption(const QTableView* self, QStyleOptionViewItem* option);
extern __declspec(dllexport) 
void QTableView_PaintEvent(QTableView* self, QPaintEvent* e);
extern __declspec(dllexport) 
void QTableView_TimerEvent(QTableView* self, QTimerEvent* event);
extern __declspec(dllexport) 
void QTableView_DropEvent(QTableView* self, QDropEvent* event);
extern __declspec(dllexport) 
int QTableView_HorizontalOffset(const QTableView* self);
extern __declspec(dllexport) 
int QTableView_VerticalOffset(const QTableView* self);
extern __declspec(dllexport) 
QModelIndex* QTableView_MoveCursor(QTableView* self, CursorAction cursorAction, int modifiers);
extern __declspec(dllexport) 
void QTableView_SetSelection(QTableView* self, QRect* rect, int command);
extern __declspec(dllexport) 
QRegion* QTableView_VisualRegionForSelection(const QTableView* self, QItemSelection* selection);
extern __declspec(dllexport) 
struct miqt_array /* of QModelIndex* */  QTableView_SelectedIndexes(const QTableView* self);
extern __declspec(dllexport) 
void QTableView_UpdateGeometries(QTableView* self);
extern __declspec(dllexport) 
QSize* QTableView_ViewportSizeHint(const QTableView* self);
extern __declspec(dllexport) 
int QTableView_SizeHintForRow(const QTableView* self, int row);
extern __declspec(dllexport) 
int QTableView_SizeHintForColumn(const QTableView* self, int column);
extern __declspec(dllexport) 
void QTableView_VerticalScrollbarAction(QTableView* self, int action);
extern __declspec(dllexport) 
void QTableView_HorizontalScrollbarAction(QTableView* self, int action);
extern __declspec(dllexport) 
bool QTableView_IsIndexHidden(const QTableView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QTableView_SelectionChanged(QTableView* self, QItemSelection* selected, QItemSelection* deselected);
extern __declspec(dllexport) 
void QTableView_CurrentChanged(QTableView* self, QModelIndex* current, QModelIndex* previous);
extern __declspec(dllexport) 
struct miqt_string QTableView_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QTableView_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QTableView_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QTableView_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QTableView_override_virtual_Metacast(void* self, intptr_t slot);
void* QTableView_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QTableView_Delete(QTableView* self, bool isSubclass);

}
