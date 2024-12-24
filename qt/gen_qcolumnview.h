#pragma once
#ifndef MIQT_QT_GEN_QCOLUMNVIEW_H
#define MIQT_QT_GEN_QCOLUMNVIEW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QAbstractItemView QAbstractItemView;
typedef struct QAbstractScrollArea QAbstractScrollArea;
typedef struct QColumnView QColumnView;
typedef struct QFrame QFrame;
typedef struct QItemSelection QItemSelection;
typedef struct QItemSelectionModel QItemSelectionModel;
typedef struct QMetaObject QMetaObject;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QRegion QRegion;
typedef struct QResizeEvent QResizeEvent;
typedef struct QSize QSize;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QColumnView* QColumnView_new(QWidget* parent);
extern __declspec(dllexport) 
QColumnView* QColumnView_new2();
extern __declspec(dllexport) 
void QColumnView_virtbase(QColumnView* src
, QAbstractItemView** outptr_QAbstractItemView
);
extern __declspec(dllexport) 
QMetaObject* QColumnView_MetaObject(const QColumnView* self);
extern __declspec(dllexport) 
void* QColumnView_Metacast(QColumnView* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QColumnView_Tr(const char* s);
extern __declspec(dllexport) 
void QColumnView_UpdatePreviewWidget(QColumnView* self, QModelIndex* index);
void QColumnView_connect_UpdatePreviewWidget(QColumnView* self, intptr_t slot);
extern __declspec(dllexport) 
QModelIndex* QColumnView_IndexAt(const QColumnView* self, QPoint* point);
extern __declspec(dllexport) 
void QColumnView_ScrollTo(QColumnView* self, QModelIndex* index, ScrollHint hint);
extern __declspec(dllexport) 
QSize* QColumnView_SizeHint(const QColumnView* self);
extern __declspec(dllexport) 
QRect* QColumnView_VisualRect(const QColumnView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QColumnView_SetModel(QColumnView* self, QAbstractItemModel* model);
extern __declspec(dllexport) 
void QColumnView_SetSelectionModel(QColumnView* self, QItemSelectionModel* selectionModel);
extern __declspec(dllexport) 
void QColumnView_SetRootIndex(QColumnView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QColumnView_SelectAll(QColumnView* self);
extern __declspec(dllexport) 
void QColumnView_SetResizeGripsVisible(QColumnView* self, bool visible);
extern __declspec(dllexport) 
bool QColumnView_ResizeGripsVisible(const QColumnView* self);
extern __declspec(dllexport) 
QWidget* QColumnView_PreviewWidget(const QColumnView* self);
extern __declspec(dllexport) 
void QColumnView_SetPreviewWidget(QColumnView* self, QWidget* widget);
extern __declspec(dllexport) 
void QColumnView_SetColumnWidths(QColumnView* self, struct miqt_array /* of int */  list);
extern __declspec(dllexport) 
struct miqt_array /* of int */  QColumnView_ColumnWidths(const QColumnView* self);
extern __declspec(dllexport) 
bool QColumnView_IsIndexHidden(const QColumnView* self, QModelIndex* index);
extern __declspec(dllexport) 
QModelIndex* QColumnView_MoveCursor(QColumnView* self, CursorAction cursorAction, int modifiers);
extern __declspec(dllexport) 
void QColumnView_ResizeEvent(QColumnView* self, QResizeEvent* event);
extern __declspec(dllexport) 
void QColumnView_SetSelection(QColumnView* self, QRect* rect, int command);
extern __declspec(dllexport) 
QRegion* QColumnView_VisualRegionForSelection(const QColumnView* self, QItemSelection* selection);
extern __declspec(dllexport) 
int QColumnView_HorizontalOffset(const QColumnView* self);
extern __declspec(dllexport) 
int QColumnView_VerticalOffset(const QColumnView* self);
extern __declspec(dllexport) 
void QColumnView_RowsInserted(QColumnView* self, QModelIndex* parent, int start, int end);
extern __declspec(dllexport) 
void QColumnView_CurrentChanged(QColumnView* self, QModelIndex* current, QModelIndex* previous);
extern __declspec(dllexport) 
void QColumnView_ScrollContentsBy(QColumnView* self, int dx, int dy);
extern __declspec(dllexport) 
QAbstractItemView* QColumnView_CreateColumn(QColumnView* self, QModelIndex* rootIndex);
extern __declspec(dllexport) 
struct miqt_string QColumnView_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QColumnView_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QColumnView_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QColumnView_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QColumnView_override_virtual_Metacast(void* self, intptr_t slot);
void* QColumnView_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QColumnView_Delete(QColumnView* self, bool isSubclass);

}
