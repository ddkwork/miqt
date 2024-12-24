#pragma once
#ifndef MIQT_QT_GEN_QABSTRACTITEMVIEW_H
#define MIQT_QT_GEN_QABSTRACTITEMVIEW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemDelegate QAbstractItemDelegate;
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QAbstractItemView QAbstractItemView;
typedef struct QAbstractScrollArea QAbstractScrollArea;
typedef struct QDragEnterEvent QDragEnterEvent;
typedef struct QDragLeaveEvent QDragLeaveEvent;
typedef struct QDragMoveEvent QDragMoveEvent;
typedef struct QDropEvent QDropEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QFrame QFrame;
typedef struct QInputMethodEvent QInputMethodEvent;
typedef struct QItemSelection QItemSelection;
typedef struct QItemSelectionModel QItemSelectionModel;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QModelIndex QModelIndex;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QRegion QRegion;
typedef struct QResizeEvent QResizeEvent;
typedef struct QSize QSize;
typedef struct QStyleOptionViewItem QStyleOptionViewItem;
typedef struct QTimerEvent QTimerEvent;
typedef struct QVariant QVariant;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAbstractItemView* QAbstractItemView_new(QWidget* parent);
extern __declspec(dllexport) 
QAbstractItemView* QAbstractItemView_new2();
extern __declspec(dllexport) 
void QAbstractItemView_virtbase(QAbstractItemView* src
, QAbstractScrollArea** outptr_QAbstractScrollArea
);
extern __declspec(dllexport) 
QMetaObject* QAbstractItemView_MetaObject(const QAbstractItemView* self);
extern __declspec(dllexport) 
void* QAbstractItemView_Metacast(QAbstractItemView* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAbstractItemView_Tr(const char* s);
extern __declspec(dllexport) 
void QAbstractItemView_SetModel(QAbstractItemView* self, QAbstractItemModel* model);
extern __declspec(dllexport) 
QAbstractItemModel* QAbstractItemView_Model(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetSelectionModel(QAbstractItemView* self, QItemSelectionModel* selectionModel);
extern __declspec(dllexport) 
QItemSelectionModel* QAbstractItemView_SelectionModel(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetItemDelegate(QAbstractItemView* self, QAbstractItemDelegate* delegate);
extern __declspec(dllexport) 
QAbstractItemDelegate* QAbstractItemView_ItemDelegate(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetSelectionMode(QAbstractItemView* self, int mode);
extern __declspec(dllexport) 
int QAbstractItemView_SelectionMode(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetSelectionBehavior(QAbstractItemView* self, int behavior);
extern __declspec(dllexport) 
int QAbstractItemView_SelectionBehavior(const QAbstractItemView* self);
extern __declspec(dllexport) 
QModelIndex* QAbstractItemView_CurrentIndex(const QAbstractItemView* self);
extern __declspec(dllexport) 
QModelIndex* QAbstractItemView_RootIndex(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetEditTriggers(QAbstractItemView* self, EditTriggers triggers);
extern __declspec(dllexport) 
EditTriggers QAbstractItemView_EditTriggers(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetVerticalScrollMode(QAbstractItemView* self, ScrollMode mode);
extern __declspec(dllexport) 
ScrollMode QAbstractItemView_VerticalScrollMode(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_ResetVerticalScrollMode(QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetHorizontalScrollMode(QAbstractItemView* self, ScrollMode mode);
extern __declspec(dllexport) 
ScrollMode QAbstractItemView_HorizontalScrollMode(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_ResetHorizontalScrollMode(QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetAutoScroll(QAbstractItemView* self, bool enable);
extern __declspec(dllexport) 
bool QAbstractItemView_HasAutoScroll(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetAutoScrollMargin(QAbstractItemView* self, int margin);
extern __declspec(dllexport) 
int QAbstractItemView_AutoScrollMargin(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetTabKeyNavigation(QAbstractItemView* self, bool enable);
extern __declspec(dllexport) 
bool QAbstractItemView_TabKeyNavigation(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetDropIndicatorShown(QAbstractItemView* self, bool enable);
extern __declspec(dllexport) 
bool QAbstractItemView_ShowDropIndicator(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetDragEnabled(QAbstractItemView* self, bool enable);
extern __declspec(dllexport) 
bool QAbstractItemView_DragEnabled(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetDragDropOverwriteMode(QAbstractItemView* self, bool overwrite);
extern __declspec(dllexport) 
bool QAbstractItemView_DragDropOverwriteMode(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetDragDropMode(QAbstractItemView* self, DragDropMode behavior);
extern __declspec(dllexport) 
DragDropMode QAbstractItemView_DragDropMode(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetDefaultDropAction(QAbstractItemView* self, int dropAction);
extern __declspec(dllexport) 
int QAbstractItemView_DefaultDropAction(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetAlternatingRowColors(QAbstractItemView* self, bool enable);
extern __declspec(dllexport) 
bool QAbstractItemView_AlternatingRowColors(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetIconSize(QAbstractItemView* self, QSize* size);
extern __declspec(dllexport) 
QSize* QAbstractItemView_IconSize(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetTextElideMode(QAbstractItemView* self, int mode);
extern __declspec(dllexport) 
int QAbstractItemView_TextElideMode(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_KeyboardSearch(QAbstractItemView* self, struct miqt_string search);
extern __declspec(dllexport) 
QRect* QAbstractItemView_VisualRect(const QAbstractItemView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QAbstractItemView_ScrollTo(QAbstractItemView* self, QModelIndex* index, ScrollHint hint);
extern __declspec(dllexport) 
QModelIndex* QAbstractItemView_IndexAt(const QAbstractItemView* self, QPoint* point);
extern __declspec(dllexport) 
QSize* QAbstractItemView_SizeHintForIndex(const QAbstractItemView* self, QModelIndex* index);
extern __declspec(dllexport) 
int QAbstractItemView_SizeHintForRow(const QAbstractItemView* self, int row);
extern __declspec(dllexport) 
int QAbstractItemView_SizeHintForColumn(const QAbstractItemView* self, int column);
extern __declspec(dllexport) 
uint32_t QAbstractItemView_UpdateThreshold(const QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetUpdateThreshold(QAbstractItemView* self, uint32_t threshold);
extern __declspec(dllexport) 
void QAbstractItemView_OpenPersistentEditor(QAbstractItemView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QAbstractItemView_ClosePersistentEditor(QAbstractItemView* self, QModelIndex* index);
extern __declspec(dllexport) 
bool QAbstractItemView_IsPersistentEditorOpen(const QAbstractItemView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QAbstractItemView_SetIndexWidget(QAbstractItemView* self, QModelIndex* index, QWidget* widget);
extern __declspec(dllexport) 
QWidget* QAbstractItemView_IndexWidget(const QAbstractItemView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QAbstractItemView_SetItemDelegateForRow(QAbstractItemView* self, int row, QAbstractItemDelegate* delegate);
extern __declspec(dllexport) 
QAbstractItemDelegate* QAbstractItemView_ItemDelegateForRow(const QAbstractItemView* self, int row);
extern __declspec(dllexport) 
void QAbstractItemView_SetItemDelegateForColumn(QAbstractItemView* self, int column, QAbstractItemDelegate* delegate);
extern __declspec(dllexport) 
QAbstractItemDelegate* QAbstractItemView_ItemDelegateForColumn(const QAbstractItemView* self, int column);
extern __declspec(dllexport) 
QAbstractItemDelegate* QAbstractItemView_ItemDelegateWithIndex(const QAbstractItemView* self, QModelIndex* index);
extern __declspec(dllexport) 
QAbstractItemDelegate* QAbstractItemView_ItemDelegateForIndex(const QAbstractItemView* self, QModelIndex* index);
extern __declspec(dllexport) 
QVariant* QAbstractItemView_InputMethodQuery(const QAbstractItemView* self, int query);
extern __declspec(dllexport) 
void QAbstractItemView_Reset(QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetRootIndex(QAbstractItemView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QAbstractItemView_DoItemsLayout(QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SelectAll(QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_Edit(QAbstractItemView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QAbstractItemView_ClearSelection(QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_SetCurrentIndex(QAbstractItemView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QAbstractItemView_ScrollToTop(QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_ScrollToBottom(QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_Update(QAbstractItemView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QAbstractItemView_DataChanged(QAbstractItemView* self, QModelIndex* topLeft, QModelIndex* bottomRight, struct miqt_array /* of int */  roles);
extern __declspec(dllexport) 
void QAbstractItemView_RowsInserted(QAbstractItemView* self, QModelIndex* parent, int start, int end);
extern __declspec(dllexport) 
void QAbstractItemView_RowsAboutToBeRemoved(QAbstractItemView* self, QModelIndex* parent, int start, int end);
extern __declspec(dllexport) 
void QAbstractItemView_SelectionChanged(QAbstractItemView* self, QItemSelection* selected, QItemSelection* deselected);
extern __declspec(dllexport) 
void QAbstractItemView_CurrentChanged(QAbstractItemView* self, QModelIndex* current, QModelIndex* previous);
extern __declspec(dllexport) 
void QAbstractItemView_UpdateEditorData(QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_UpdateEditorGeometries(QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_UpdateGeometries(QAbstractItemView* self);
extern __declspec(dllexport) 
void QAbstractItemView_VerticalScrollbarAction(QAbstractItemView* self, int action);
extern __declspec(dllexport) 
void QAbstractItemView_HorizontalScrollbarAction(QAbstractItemView* self, int action);
extern __declspec(dllexport) 
void QAbstractItemView_VerticalScrollbarValueChanged(QAbstractItemView* self, int value);
extern __declspec(dllexport) 
void QAbstractItemView_HorizontalScrollbarValueChanged(QAbstractItemView* self, int value);
extern __declspec(dllexport) 
void QAbstractItemView_CloseEditor(QAbstractItemView* self, QWidget* editor, int hint);
extern __declspec(dllexport) 
void QAbstractItemView_CommitData(QAbstractItemView* self, QWidget* editor);
extern __declspec(dllexport) 
void QAbstractItemView_EditorDestroyed(QAbstractItemView* self, QObject* editor);
extern __declspec(dllexport) 
void QAbstractItemView_Pressed(QAbstractItemView* self, QModelIndex* index);
void QAbstractItemView_connect_Pressed(QAbstractItemView* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractItemView_Clicked(QAbstractItemView* self, QModelIndex* index);
void QAbstractItemView_connect_Clicked(QAbstractItemView* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractItemView_DoubleClicked(QAbstractItemView* self, QModelIndex* index);
void QAbstractItemView_connect_DoubleClicked(QAbstractItemView* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractItemView_Activated(QAbstractItemView* self, QModelIndex* index);
void QAbstractItemView_connect_Activated(QAbstractItemView* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractItemView_Entered(QAbstractItemView* self, QModelIndex* index);
void QAbstractItemView_connect_Entered(QAbstractItemView* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractItemView_ViewportEntered(QAbstractItemView* self);
void QAbstractItemView_connect_ViewportEntered(QAbstractItemView* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractItemView_IconSizeChanged(QAbstractItemView* self, QSize* size);
void QAbstractItemView_connect_IconSizeChanged(QAbstractItemView* self, intptr_t slot);
extern __declspec(dllexport) 
QModelIndex* QAbstractItemView_MoveCursor(QAbstractItemView* self, CursorAction cursorAction, int modifiers);
extern __declspec(dllexport) 
int QAbstractItemView_HorizontalOffset(const QAbstractItemView* self);
extern __declspec(dllexport) 
int QAbstractItemView_VerticalOffset(const QAbstractItemView* self);
extern __declspec(dllexport) 
bool QAbstractItemView_IsIndexHidden(const QAbstractItemView* self, QModelIndex* index);
extern __declspec(dllexport) 
void QAbstractItemView_SetSelection(QAbstractItemView* self, QRect* rect, int command);
extern __declspec(dllexport) 
QRegion* QAbstractItemView_VisualRegionForSelection(const QAbstractItemView* self, QItemSelection* selection);
extern __declspec(dllexport) 
struct miqt_array /* of QModelIndex* */  QAbstractItemView_SelectedIndexes(const QAbstractItemView* self);
extern __declspec(dllexport) 
bool QAbstractItemView_Edit2(QAbstractItemView* self, QModelIndex* index, EditTrigger trigger, QEvent* event);
extern __declspec(dllexport) 
int QAbstractItemView_SelectionCommand(const QAbstractItemView* self, QModelIndex* index, QEvent* event);
extern __declspec(dllexport) 
void QAbstractItemView_StartDrag(QAbstractItemView* self, int supportedActions);
extern __declspec(dllexport) 
void QAbstractItemView_InitViewItemOption(const QAbstractItemView* self, QStyleOptionViewItem* option);
extern __declspec(dllexport) 
bool QAbstractItemView_FocusNextPrevChild(QAbstractItemView* self, bool next);
extern __declspec(dllexport) 
bool QAbstractItemView_Event(QAbstractItemView* self, QEvent* event);
extern __declspec(dllexport) 
bool QAbstractItemView_ViewportEvent(QAbstractItemView* self, QEvent* event);
extern __declspec(dllexport) 
void QAbstractItemView_MousePressEvent(QAbstractItemView* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QAbstractItemView_MouseMoveEvent(QAbstractItemView* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QAbstractItemView_MouseReleaseEvent(QAbstractItemView* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QAbstractItemView_MouseDoubleClickEvent(QAbstractItemView* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QAbstractItemView_DragEnterEvent(QAbstractItemView* self, QDragEnterEvent* event);
extern __declspec(dllexport) 
void QAbstractItemView_DragMoveEvent(QAbstractItemView* self, QDragMoveEvent* event);
extern __declspec(dllexport) 
void QAbstractItemView_DragLeaveEvent(QAbstractItemView* self, QDragLeaveEvent* event);
extern __declspec(dllexport) 
void QAbstractItemView_DropEvent(QAbstractItemView* self, QDropEvent* event);
extern __declspec(dllexport) 
void QAbstractItemView_FocusInEvent(QAbstractItemView* self, QFocusEvent* event);
extern __declspec(dllexport) 
void QAbstractItemView_FocusOutEvent(QAbstractItemView* self, QFocusEvent* event);
extern __declspec(dllexport) 
void QAbstractItemView_KeyPressEvent(QAbstractItemView* self, QKeyEvent* event);
extern __declspec(dllexport) 
void QAbstractItemView_ResizeEvent(QAbstractItemView* self, QResizeEvent* event);
extern __declspec(dllexport) 
void QAbstractItemView_TimerEvent(QAbstractItemView* self, QTimerEvent* event);
extern __declspec(dllexport) 
void QAbstractItemView_InputMethodEvent(QAbstractItemView* self, QInputMethodEvent* event);
extern __declspec(dllexport) 
bool QAbstractItemView_EventFilter(QAbstractItemView* self, QObject* object, QEvent* event);
extern __declspec(dllexport) 
QSize* QAbstractItemView_ViewportSizeHint(const QAbstractItemView* self);
extern __declspec(dllexport) 
struct miqt_string QAbstractItemView_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAbstractItemView_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAbstractItemView_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAbstractItemView_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAbstractItemView_override_virtual_Metacast(void* self, intptr_t slot);
void* QAbstractItemView_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAbstractItemView_Delete(QAbstractItemView* self, bool isSubclass);

}
