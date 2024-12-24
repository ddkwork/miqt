#pragma once
#ifndef MIQT_QT_GEN_QHEADERVIEW_H
#define MIQT_QT_GEN_QHEADERVIEW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QAbstractItemView QAbstractItemView;
typedef struct QAbstractScrollArea QAbstractScrollArea;
typedef struct QEvent QEvent;
typedef struct QFrame QFrame;
typedef struct QHeaderView QHeaderView;
typedef struct QItemSelection QItemSelection;
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
typedef struct QStyleOptionHeader QStyleOptionHeader;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QHeaderView* QHeaderView_new(int orientation);
extern __declspec(dllexport) 
QHeaderView* QHeaderView_new2(int orientation, QWidget* parent);
extern __declspec(dllexport) 
void QHeaderView_virtbase(QHeaderView* src
, QAbstractItemView** outptr_QAbstractItemView
);
extern __declspec(dllexport) 
QMetaObject* QHeaderView_MetaObject(const QHeaderView* self);
extern __declspec(dllexport) 
void* QHeaderView_Metacast(QHeaderView* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QHeaderView_Tr(const char* s);
extern __declspec(dllexport) 
void QHeaderView_SetModel(QHeaderView* self, QAbstractItemModel* model);
extern __declspec(dllexport) 
int QHeaderView_Orientation(const QHeaderView* self);
extern __declspec(dllexport) 
int QHeaderView_Offset(const QHeaderView* self);
extern __declspec(dllexport) 
int QHeaderView_Length(const QHeaderView* self);
extern __declspec(dllexport) 
QSize* QHeaderView_SizeHint(const QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_SetVisible(QHeaderView* self, bool v);
extern __declspec(dllexport) 
int QHeaderView_SectionSizeHint(const QHeaderView* self, int logicalIndex);
extern __declspec(dllexport) 
int QHeaderView_VisualIndexAt(const QHeaderView* self, int position);
extern __declspec(dllexport) 
int QHeaderView_LogicalIndexAt(const QHeaderView* self, int position);
extern __declspec(dllexport) 
int QHeaderView_LogicalIndexAt2(const QHeaderView* self, int x, int y);
extern __declspec(dllexport) 
int QHeaderView_LogicalIndexAtWithPos(const QHeaderView* self, QPoint* pos);
extern __declspec(dllexport) 
int QHeaderView_SectionSize(const QHeaderView* self, int logicalIndex);
extern __declspec(dllexport) 
int QHeaderView_SectionPosition(const QHeaderView* self, int logicalIndex);
extern __declspec(dllexport) 
int QHeaderView_SectionViewportPosition(const QHeaderView* self, int logicalIndex);
extern __declspec(dllexport) 
void QHeaderView_MoveSection(QHeaderView* self, int from, int to);
extern __declspec(dllexport) 
void QHeaderView_SwapSections(QHeaderView* self, int first, int second);
extern __declspec(dllexport) 
void QHeaderView_ResizeSection(QHeaderView* self, int logicalIndex, int size);
extern __declspec(dllexport) 
void QHeaderView_ResizeSections(QHeaderView* self, int mode);
extern __declspec(dllexport) 
bool QHeaderView_IsSectionHidden(const QHeaderView* self, int logicalIndex);
extern __declspec(dllexport) 
void QHeaderView_SetSectionHidden(QHeaderView* self, int logicalIndex, bool hide);
extern __declspec(dllexport) 
int QHeaderView_HiddenSectionCount(const QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_HideSection(QHeaderView* self, int logicalIndex);
extern __declspec(dllexport) 
void QHeaderView_ShowSection(QHeaderView* self, int logicalIndex);
extern __declspec(dllexport) 
int QHeaderView_Count(const QHeaderView* self);
extern __declspec(dllexport) 
int QHeaderView_VisualIndex(const QHeaderView* self, int logicalIndex);
extern __declspec(dllexport) 
int QHeaderView_LogicalIndex(const QHeaderView* self, int visualIndex);
extern __declspec(dllexport) 
void QHeaderView_SetSectionsMovable(QHeaderView* self, bool movable);
extern __declspec(dllexport) 
bool QHeaderView_SectionsMovable(const QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_SetFirstSectionMovable(QHeaderView* self, bool movable);
extern __declspec(dllexport) 
bool QHeaderView_IsFirstSectionMovable(const QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_SetSectionsClickable(QHeaderView* self, bool clickable);
extern __declspec(dllexport) 
bool QHeaderView_SectionsClickable(const QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_SetHighlightSections(QHeaderView* self, bool highlight);
extern __declspec(dllexport) 
bool QHeaderView_HighlightSections(const QHeaderView* self);
extern __declspec(dllexport) 
ResizeMode QHeaderView_SectionResizeMode(const QHeaderView* self, int logicalIndex);
extern __declspec(dllexport) 
void QHeaderView_SetSectionResizeMode(QHeaderView* self, ResizeMode mode);
extern __declspec(dllexport) 
void QHeaderView_SetSectionResizeMode2(QHeaderView* self, int logicalIndex, ResizeMode mode);
extern __declspec(dllexport) 
void QHeaderView_SetResizeContentsPrecision(QHeaderView* self, int precision);
extern __declspec(dllexport) 
int QHeaderView_ResizeContentsPrecision(const QHeaderView* self);
extern __declspec(dllexport) 
int QHeaderView_StretchSectionCount(const QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_SetSortIndicatorShown(QHeaderView* self, bool show);
extern __declspec(dllexport) 
bool QHeaderView_IsSortIndicatorShown(const QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_SetSortIndicator(QHeaderView* self, int logicalIndex, int order);
extern __declspec(dllexport) 
int QHeaderView_SortIndicatorSection(const QHeaderView* self);
extern __declspec(dllexport) 
int QHeaderView_SortIndicatorOrder(const QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_SetSortIndicatorClearable(QHeaderView* self, bool clearable);
extern __declspec(dllexport) 
bool QHeaderView_IsSortIndicatorClearable(const QHeaderView* self);
extern __declspec(dllexport) 
bool QHeaderView_StretchLastSection(const QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_SetStretchLastSection(QHeaderView* self, bool stretch);
extern __declspec(dllexport) 
bool QHeaderView_CascadingSectionResizes(const QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_SetCascadingSectionResizes(QHeaderView* self, bool enable);
extern __declspec(dllexport) 
int QHeaderView_DefaultSectionSize(const QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_SetDefaultSectionSize(QHeaderView* self, int size);
extern __declspec(dllexport) 
void QHeaderView_ResetDefaultSectionSize(QHeaderView* self);
extern __declspec(dllexport) 
int QHeaderView_MinimumSectionSize(const QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_SetMinimumSectionSize(QHeaderView* self, int size);
extern __declspec(dllexport) 
int QHeaderView_MaximumSectionSize(const QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_SetMaximumSectionSize(QHeaderView* self, int size);
extern __declspec(dllexport) 
int QHeaderView_DefaultAlignment(const QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_SetDefaultAlignment(QHeaderView* self, int alignment);
extern __declspec(dllexport) 
void QHeaderView_DoItemsLayout(QHeaderView* self);
extern __declspec(dllexport) 
bool QHeaderView_SectionsMoved(const QHeaderView* self);
extern __declspec(dllexport) 
bool QHeaderView_SectionsHidden(const QHeaderView* self);
extern __declspec(dllexport) 
struct miqt_string QHeaderView_SaveState(const QHeaderView* self);
extern __declspec(dllexport) 
bool QHeaderView_RestoreState(QHeaderView* self, struct miqt_string state);
extern __declspec(dllexport) 
void QHeaderView_Reset(QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_SetOffset(QHeaderView* self, int offset);
extern __declspec(dllexport) 
void QHeaderView_SetOffsetToSectionPosition(QHeaderView* self, int visualIndex);
extern __declspec(dllexport) 
void QHeaderView_SetOffsetToLastSection(QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_HeaderDataChanged(QHeaderView* self, int orientation, int logicalFirst, int logicalLast);
extern __declspec(dllexport) 
void QHeaderView_SectionMoved(QHeaderView* self, int logicalIndex, int oldVisualIndex, int newVisualIndex);
void QHeaderView_connect_SectionMoved(QHeaderView* self, intptr_t slot);
extern __declspec(dllexport) 
void QHeaderView_SectionResized(QHeaderView* self, int logicalIndex, int oldSize, int newSize);
void QHeaderView_connect_SectionResized(QHeaderView* self, intptr_t slot);
extern __declspec(dllexport) 
void QHeaderView_SectionPressed(QHeaderView* self, int logicalIndex);
void QHeaderView_connect_SectionPressed(QHeaderView* self, intptr_t slot);
extern __declspec(dllexport) 
void QHeaderView_SectionClicked(QHeaderView* self, int logicalIndex);
void QHeaderView_connect_SectionClicked(QHeaderView* self, intptr_t slot);
extern __declspec(dllexport) 
void QHeaderView_SectionEntered(QHeaderView* self, int logicalIndex);
void QHeaderView_connect_SectionEntered(QHeaderView* self, intptr_t slot);
extern __declspec(dllexport) 
void QHeaderView_SectionDoubleClicked(QHeaderView* self, int logicalIndex);
void QHeaderView_connect_SectionDoubleClicked(QHeaderView* self, intptr_t slot);
extern __declspec(dllexport) 
void QHeaderView_SectionCountChanged(QHeaderView* self, int oldCount, int newCount);
void QHeaderView_connect_SectionCountChanged(QHeaderView* self, intptr_t slot);
extern __declspec(dllexport) 
void QHeaderView_SectionHandleDoubleClicked(QHeaderView* self, int logicalIndex);
void QHeaderView_connect_SectionHandleDoubleClicked(QHeaderView* self, intptr_t slot);
extern __declspec(dllexport) 
void QHeaderView_GeometriesChanged(QHeaderView* self);
void QHeaderView_connect_GeometriesChanged(QHeaderView* self, intptr_t slot);
extern __declspec(dllexport) 
void QHeaderView_SortIndicatorChanged(QHeaderView* self, int logicalIndex, int order);
void QHeaderView_connect_SortIndicatorChanged(QHeaderView* self, intptr_t slot);
extern __declspec(dllexport) 
void QHeaderView_SortIndicatorClearableChanged(QHeaderView* self, bool clearable);
void QHeaderView_connect_SortIndicatorClearableChanged(QHeaderView* self, intptr_t slot);
extern __declspec(dllexport) 
void QHeaderView_CurrentChanged(QHeaderView* self, QModelIndex* current, QModelIndex* old);
extern __declspec(dllexport) 
bool QHeaderView_Event(QHeaderView* self, QEvent* e);
extern __declspec(dllexport) 
void QHeaderView_PaintEvent(QHeaderView* self, QPaintEvent* e);
extern __declspec(dllexport) 
void QHeaderView_MousePressEvent(QHeaderView* self, QMouseEvent* e);
extern __declspec(dllexport) 
void QHeaderView_MouseMoveEvent(QHeaderView* self, QMouseEvent* e);
extern __declspec(dllexport) 
void QHeaderView_MouseReleaseEvent(QHeaderView* self, QMouseEvent* e);
extern __declspec(dllexport) 
void QHeaderView_MouseDoubleClickEvent(QHeaderView* self, QMouseEvent* e);
extern __declspec(dllexport) 
bool QHeaderView_ViewportEvent(QHeaderView* self, QEvent* e);
extern __declspec(dllexport) 
void QHeaderView_PaintSection(const QHeaderView* self, QPainter* painter, QRect* rect, int logicalIndex);
extern __declspec(dllexport) 
QSize* QHeaderView_SectionSizeFromContents(const QHeaderView* self, int logicalIndex);
extern __declspec(dllexport) 
int QHeaderView_HorizontalOffset(const QHeaderView* self);
extern __declspec(dllexport) 
int QHeaderView_VerticalOffset(const QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_UpdateGeometries(QHeaderView* self);
extern __declspec(dllexport) 
void QHeaderView_ScrollContentsBy(QHeaderView* self, int dx, int dy);
extern __declspec(dllexport) 
void QHeaderView_DataChanged(QHeaderView* self, QModelIndex* topLeft, QModelIndex* bottomRight, struct miqt_array /* of int */  roles);
extern __declspec(dllexport) 
void QHeaderView_RowsInserted(QHeaderView* self, QModelIndex* parent, int start, int end);
extern __declspec(dllexport) 
QRect* QHeaderView_VisualRect(const QHeaderView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QHeaderView_ScrollTo(QHeaderView* self, QModelIndex* index, ScrollHint hint);
extern __declspec(dllexport) 
QModelIndex* QHeaderView_IndexAt(const QHeaderView* self, QPoint* p);
extern __declspec(dllexport) 
bool QHeaderView_IsIndexHidden(const QHeaderView* self, QModelIndex* index);
extern __declspec(dllexport) 
QModelIndex* QHeaderView_MoveCursor(QHeaderView* self, CursorAction param1, int param2);
extern __declspec(dllexport) 
void QHeaderView_SetSelection(QHeaderView* self, QRect* rect, int flags);
extern __declspec(dllexport) 
QRegion* QHeaderView_VisualRegionForSelection(const QHeaderView* self, QItemSelection* selection);
extern __declspec(dllexport) 
void QHeaderView_InitStyleOptionForIndex(const QHeaderView* self, QStyleOptionHeader* option, int logicalIndex);
extern __declspec(dllexport) 
void QHeaderView_InitStyleOption(const QHeaderView* self, QStyleOptionHeader* option);
extern __declspec(dllexport) 
struct miqt_string QHeaderView_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QHeaderView_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QHeaderView_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QHeaderView_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QHeaderView_override_virtual_Metacast(void* self, intptr_t slot);
void* QHeaderView_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QHeaderView_Delete(QHeaderView* self, bool isSubclass);

}
