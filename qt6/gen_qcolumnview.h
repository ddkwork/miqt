#pragma once
#ifndef MIQT_QT6_GEN_QCOLUMNVIEW_H
#define MIQT_QT6_GEN_QCOLUMNVIEW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractItemDelegate;
class QAbstractItemModel;
class QAbstractItemView;
class QAbstractScrollArea;
class QActionEvent;
class QChildEvent;
class QCloseEvent;
class QColumnView;
class QContextMenuEvent;
class QDragEnterEvent;
class QDragLeaveEvent;
class QDragMoveEvent;
class QDropEvent;
class QEnterEvent;
class QEvent;
class QFocusEvent;
class QFrame;
class QHideEvent;
class QInputMethodEvent;
class QItemSelection;
class QItemSelectionModel;
class QKeyEvent;
class QMargins;
class QMetaMethod;
class QMetaObject;
class QModelIndex;
class QMouseEvent;
class QMoveEvent;
class QObject;
class QPaintDevice;
class QPaintEngine;
class QPaintEvent;
class QPainter;
class QPoint;
class QRect;
class QRegion;
class QResizeEvent;
class QShowEvent;
class QSize;
class QStyleOptionFrame;
class QStyleOptionViewItem;
class QTabletEvent;
class QTimerEvent;
class QVariant;
class QWheelEvent;
class QWidget;
#else
typedef struct QAbstractItemDelegate QAbstractItemDelegate;
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QAbstractItemView QAbstractItemView;
typedef struct QAbstractScrollArea QAbstractScrollArea;
typedef struct QActionEvent QActionEvent;
typedef struct QChildEvent QChildEvent;
typedef struct QCloseEvent QCloseEvent;
typedef struct QColumnView QColumnView;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDragEnterEvent QDragEnterEvent;
typedef struct QDragLeaveEvent QDragLeaveEvent;
typedef struct QDragMoveEvent QDragMoveEvent;
typedef struct QDropEvent QDropEvent;
typedef struct QEnterEvent QEnterEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QFrame QFrame;
typedef struct QHideEvent QHideEvent;
typedef struct QInputMethodEvent QInputMethodEvent;
typedef struct QItemSelection QItemSelection;
typedef struct QItemSelectionModel QItemSelectionModel;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMargins QMargins;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QModelIndex QModelIndex;
typedef struct QMouseEvent QMouseEvent;
typedef struct QMoveEvent QMoveEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPainter QPainter;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QRegion QRegion;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QStyleOptionFrame QStyleOptionFrame;
typedef struct QStyleOptionViewItem QStyleOptionViewItem;
typedef struct QTabletEvent QTabletEvent;
typedef struct QTimerEvent QTimerEvent;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
#endif

QColumnView* QColumnView_new(QWidget* parent);
QColumnView* QColumnView_new2();
void QColumnView_virtbase(QColumnView* src, QAbstractItemView** outptr_QAbstractItemView);
QMetaObject* QColumnView_metaObject(const QColumnView* self);
void* QColumnView_metacast(QColumnView* self, const char* param1);
struct miqt_string QColumnView_tr(const char* s);
void QColumnView_updatePreviewWidget(QColumnView* self, QModelIndex* index);
void QColumnView_connect_updatePreviewWidget(QColumnView* self, intptr_t slot);
QModelIndex* QColumnView_indexAt(const QColumnView* self, QPoint* point);
void QColumnView_scrollTo(QColumnView* self, QModelIndex* index, int hint);
QSize* QColumnView_sizeHint(const QColumnView* self);
QRect* QColumnView_visualRect(const QColumnView* self, QModelIndex* index);
void QColumnView_setModel(QColumnView* self, QAbstractItemModel* model);
void QColumnView_setSelectionModel(QColumnView* self, QItemSelectionModel* selectionModel);
void QColumnView_setRootIndex(QColumnView* self, QModelIndex* index);
void QColumnView_selectAll(QColumnView* self);
void QColumnView_setResizeGripsVisible(QColumnView* self, bool visible);
bool QColumnView_resizeGripsVisible(const QColumnView* self);
QWidget* QColumnView_previewWidget(const QColumnView* self);
void QColumnView_setPreviewWidget(QColumnView* self, QWidget* widget);
void QColumnView_setColumnWidths(QColumnView* self, struct miqt_array /* of int */  list);
struct miqt_array /* of int */  QColumnView_columnWidths(const QColumnView* self);
bool QColumnView_isIndexHidden(const QColumnView* self, QModelIndex* index);
QModelIndex* QColumnView_moveCursor(QColumnView* self, int cursorAction, int modifiers);
void QColumnView_resizeEvent(QColumnView* self, QResizeEvent* event);
void QColumnView_setSelection(QColumnView* self, QRect* rect, int command);
QRegion* QColumnView_visualRegionForSelection(const QColumnView* self, QItemSelection* selection);
int QColumnView_horizontalOffset(const QColumnView* self);
int QColumnView_verticalOffset(const QColumnView* self);
void QColumnView_rowsInserted(QColumnView* self, QModelIndex* parent, int start, int end);
void QColumnView_currentChanged(QColumnView* self, QModelIndex* current, QModelIndex* previous);
void QColumnView_scrollContentsBy(QColumnView* self, int dx, int dy);
QAbstractItemView* QColumnView_createColumn(QColumnView* self, QModelIndex* rootIndex);
struct miqt_string QColumnView_tr2(const char* s, const char* c);
struct miqt_string QColumnView_tr3(const char* s, const char* c, int n);
bool QColumnView_override_virtual_indexAt(void* self, intptr_t slot);
QModelIndex* QColumnView_virtualbase_indexAt(const void* self, QPoint* point);
bool QColumnView_override_virtual_scrollTo(void* self, intptr_t slot);
void QColumnView_virtualbase_scrollTo(void* self, QModelIndex* index, int hint);
bool QColumnView_override_virtual_sizeHint(void* self, intptr_t slot);
QSize* QColumnView_virtualbase_sizeHint(const void* self);
bool QColumnView_override_virtual_visualRect(void* self, intptr_t slot);
QRect* QColumnView_virtualbase_visualRect(const void* self, QModelIndex* index);
bool QColumnView_override_virtual_setModel(void* self, intptr_t slot);
void QColumnView_virtualbase_setModel(void* self, QAbstractItemModel* model);
bool QColumnView_override_virtual_setSelectionModel(void* self, intptr_t slot);
void QColumnView_virtualbase_setSelectionModel(void* self, QItemSelectionModel* selectionModel);
bool QColumnView_override_virtual_setRootIndex(void* self, intptr_t slot);
void QColumnView_virtualbase_setRootIndex(void* self, QModelIndex* index);
bool QColumnView_override_virtual_selectAll(void* self, intptr_t slot);
void QColumnView_virtualbase_selectAll(void* self);
bool QColumnView_override_virtual_isIndexHidden(void* self, intptr_t slot);
bool QColumnView_virtualbase_isIndexHidden(const void* self, QModelIndex* index);
bool QColumnView_override_virtual_moveCursor(void* self, intptr_t slot);
QModelIndex* QColumnView_virtualbase_moveCursor(void* self, int cursorAction, int modifiers);
bool QColumnView_override_virtual_resizeEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_resizeEvent(void* self, QResizeEvent* event);
bool QColumnView_override_virtual_setSelection(void* self, intptr_t slot);
void QColumnView_virtualbase_setSelection(void* self, QRect* rect, int command);
bool QColumnView_override_virtual_visualRegionForSelection(void* self, intptr_t slot);
QRegion* QColumnView_virtualbase_visualRegionForSelection(const void* self, QItemSelection* selection);
bool QColumnView_override_virtual_horizontalOffset(void* self, intptr_t slot);
int QColumnView_virtualbase_horizontalOffset(const void* self);
bool QColumnView_override_virtual_verticalOffset(void* self, intptr_t slot);
int QColumnView_virtualbase_verticalOffset(const void* self);
bool QColumnView_override_virtual_rowsInserted(void* self, intptr_t slot);
void QColumnView_virtualbase_rowsInserted(void* self, QModelIndex* parent, int start, int end);
bool QColumnView_override_virtual_currentChanged(void* self, intptr_t slot);
void QColumnView_virtualbase_currentChanged(void* self, QModelIndex* current, QModelIndex* previous);
bool QColumnView_override_virtual_scrollContentsBy(void* self, intptr_t slot);
void QColumnView_virtualbase_scrollContentsBy(void* self, int dx, int dy);
bool QColumnView_override_virtual_createColumn(void* self, intptr_t slot);
QAbstractItemView* QColumnView_virtualbase_createColumn(void* self, QModelIndex* rootIndex);
bool QColumnView_override_virtual_keyboardSearch(void* self, intptr_t slot);
void QColumnView_virtualbase_keyboardSearch(void* self, struct miqt_string search);
bool QColumnView_override_virtual_sizeHintForRow(void* self, intptr_t slot);
int QColumnView_virtualbase_sizeHintForRow(const void* self, int row);
bool QColumnView_override_virtual_sizeHintForColumn(void* self, intptr_t slot);
int QColumnView_virtualbase_sizeHintForColumn(const void* self, int column);
bool QColumnView_override_virtual_itemDelegateForIndex(void* self, intptr_t slot);
QAbstractItemDelegate* QColumnView_virtualbase_itemDelegateForIndex(const void* self, QModelIndex* index);
bool QColumnView_override_virtual_inputMethodQuery(void* self, intptr_t slot);
QVariant* QColumnView_virtualbase_inputMethodQuery(const void* self, int query);
bool QColumnView_override_virtual_reset(void* self, intptr_t slot);
void QColumnView_virtualbase_reset(void* self);
bool QColumnView_override_virtual_doItemsLayout(void* self, intptr_t slot);
void QColumnView_virtualbase_doItemsLayout(void* self);
bool QColumnView_override_virtual_dataChanged(void* self, intptr_t slot);
void QColumnView_virtualbase_dataChanged(void* self, QModelIndex* topLeft, QModelIndex* bottomRight, struct miqt_array /* of int */  roles);
bool QColumnView_override_virtual_rowsAboutToBeRemoved(void* self, intptr_t slot);
void QColumnView_virtualbase_rowsAboutToBeRemoved(void* self, QModelIndex* parent, int start, int end);
bool QColumnView_override_virtual_selectionChanged(void* self, intptr_t slot);
void QColumnView_virtualbase_selectionChanged(void* self, QItemSelection* selected, QItemSelection* deselected);
bool QColumnView_override_virtual_updateEditorData(void* self, intptr_t slot);
void QColumnView_virtualbase_updateEditorData(void* self);
bool QColumnView_override_virtual_updateEditorGeometries(void* self, intptr_t slot);
void QColumnView_virtualbase_updateEditorGeometries(void* self);
bool QColumnView_override_virtual_updateGeometries(void* self, intptr_t slot);
void QColumnView_virtualbase_updateGeometries(void* self);
bool QColumnView_override_virtual_verticalScrollbarAction(void* self, intptr_t slot);
void QColumnView_virtualbase_verticalScrollbarAction(void* self, int action);
bool QColumnView_override_virtual_horizontalScrollbarAction(void* self, intptr_t slot);
void QColumnView_virtualbase_horizontalScrollbarAction(void* self, int action);
bool QColumnView_override_virtual_verticalScrollbarValueChanged(void* self, intptr_t slot);
void QColumnView_virtualbase_verticalScrollbarValueChanged(void* self, int value);
bool QColumnView_override_virtual_horizontalScrollbarValueChanged(void* self, intptr_t slot);
void QColumnView_virtualbase_horizontalScrollbarValueChanged(void* self, int value);
bool QColumnView_override_virtual_closeEditor(void* self, intptr_t slot);
void QColumnView_virtualbase_closeEditor(void* self, QWidget* editor, int hint);
bool QColumnView_override_virtual_commitData(void* self, intptr_t slot);
void QColumnView_virtualbase_commitData(void* self, QWidget* editor);
bool QColumnView_override_virtual_editorDestroyed(void* self, intptr_t slot);
void QColumnView_virtualbase_editorDestroyed(void* self, QObject* editor);
bool QColumnView_override_virtual_selectedIndexes(void* self, intptr_t slot);
struct miqt_array /* of QModelIndex* */  QColumnView_virtualbase_selectedIndexes(const void* self);
bool QColumnView_override_virtual_edit2(void* self, intptr_t slot);
bool QColumnView_virtualbase_edit2(void* self, QModelIndex* index, int trigger, QEvent* event);
bool QColumnView_override_virtual_selectionCommand(void* self, intptr_t slot);
int QColumnView_virtualbase_selectionCommand(const void* self, QModelIndex* index, QEvent* event);
bool QColumnView_override_virtual_startDrag(void* self, intptr_t slot);
void QColumnView_virtualbase_startDrag(void* self, int supportedActions);
bool QColumnView_override_virtual_initViewItemOption(void* self, intptr_t slot);
void QColumnView_virtualbase_initViewItemOption(const void* self, QStyleOptionViewItem* option);
bool QColumnView_override_virtual_focusNextPrevChild(void* self, intptr_t slot);
bool QColumnView_virtualbase_focusNextPrevChild(void* self, bool next);
bool QColumnView_override_virtual_event(void* self, intptr_t slot);
bool QColumnView_virtualbase_event(void* self, QEvent* event);
bool QColumnView_override_virtual_viewportEvent(void* self, intptr_t slot);
bool QColumnView_virtualbase_viewportEvent(void* self, QEvent* event);
bool QColumnView_override_virtual_mousePressEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_mousePressEvent(void* self, QMouseEvent* event);
bool QColumnView_override_virtual_mouseMoveEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_mouseMoveEvent(void* self, QMouseEvent* event);
bool QColumnView_override_virtual_mouseReleaseEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_mouseReleaseEvent(void* self, QMouseEvent* event);
bool QColumnView_override_virtual_mouseDoubleClickEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_mouseDoubleClickEvent(void* self, QMouseEvent* event);
bool QColumnView_override_virtual_dragEnterEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_dragEnterEvent(void* self, QDragEnterEvent* event);
bool QColumnView_override_virtual_dragMoveEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_dragMoveEvent(void* self, QDragMoveEvent* event);
bool QColumnView_override_virtual_dragLeaveEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_dragLeaveEvent(void* self, QDragLeaveEvent* event);
bool QColumnView_override_virtual_dropEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_dropEvent(void* self, QDropEvent* event);
bool QColumnView_override_virtual_focusInEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_focusInEvent(void* self, QFocusEvent* event);
bool QColumnView_override_virtual_focusOutEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_focusOutEvent(void* self, QFocusEvent* event);
bool QColumnView_override_virtual_keyPressEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_keyPressEvent(void* self, QKeyEvent* event);
bool QColumnView_override_virtual_timerEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_timerEvent(void* self, QTimerEvent* event);
bool QColumnView_override_virtual_inputMethodEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_inputMethodEvent(void* self, QInputMethodEvent* event);
bool QColumnView_override_virtual_eventFilter(void* self, intptr_t slot);
bool QColumnView_virtualbase_eventFilter(void* self, QObject* object, QEvent* event);
bool QColumnView_override_virtual_viewportSizeHint(void* self, intptr_t slot);
QSize* QColumnView_virtualbase_viewportSizeHint(const void* self);
bool QColumnView_override_virtual_minimumSizeHint(void* self, intptr_t slot);
QSize* QColumnView_virtualbase_minimumSizeHint(const void* self);
bool QColumnView_override_virtual_setupViewport(void* self, intptr_t slot);
void QColumnView_virtualbase_setupViewport(void* self, QWidget* viewport);
bool QColumnView_override_virtual_paintEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_paintEvent(void* self, QPaintEvent* param1);
bool QColumnView_override_virtual_wheelEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_wheelEvent(void* self, QWheelEvent* param1);
bool QColumnView_override_virtual_contextMenuEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_contextMenuEvent(void* self, QContextMenuEvent* param1);
bool QColumnView_override_virtual_changeEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_changeEvent(void* self, QEvent* param1);
bool QColumnView_override_virtual_initStyleOption(void* self, intptr_t slot);
void QColumnView_virtualbase_initStyleOption(const void* self, QStyleOptionFrame* option);
bool QColumnView_override_virtual_devType(void* self, intptr_t slot);
int QColumnView_virtualbase_devType(const void* self);
bool QColumnView_override_virtual_setVisible(void* self, intptr_t slot);
void QColumnView_virtualbase_setVisible(void* self, bool visible);
bool QColumnView_override_virtual_heightForWidth(void* self, intptr_t slot);
int QColumnView_virtualbase_heightForWidth(const void* self, int param1);
bool QColumnView_override_virtual_hasHeightForWidth(void* self, intptr_t slot);
bool QColumnView_virtualbase_hasHeightForWidth(const void* self);
bool QColumnView_override_virtual_paintEngine(void* self, intptr_t slot);
QPaintEngine* QColumnView_virtualbase_paintEngine(const void* self);
bool QColumnView_override_virtual_keyReleaseEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_keyReleaseEvent(void* self, QKeyEvent* event);
bool QColumnView_override_virtual_enterEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_enterEvent(void* self, QEnterEvent* event);
bool QColumnView_override_virtual_leaveEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_leaveEvent(void* self, QEvent* event);
bool QColumnView_override_virtual_moveEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_moveEvent(void* self, QMoveEvent* event);
bool QColumnView_override_virtual_closeEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_closeEvent(void* self, QCloseEvent* event);
bool QColumnView_override_virtual_tabletEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_tabletEvent(void* self, QTabletEvent* event);
bool QColumnView_override_virtual_actionEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_actionEvent(void* self, QActionEvent* event);
bool QColumnView_override_virtual_showEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_showEvent(void* self, QShowEvent* event);
bool QColumnView_override_virtual_hideEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_hideEvent(void* self, QHideEvent* event);
bool QColumnView_override_virtual_nativeEvent(void* self, intptr_t slot);
bool QColumnView_virtualbase_nativeEvent(void* self, struct miqt_string eventType, void* message, intptr_t* result);
bool QColumnView_override_virtual_metric(void* self, intptr_t slot);
int QColumnView_virtualbase_metric(const void* self, int param1);
bool QColumnView_override_virtual_initPainter(void* self, intptr_t slot);
void QColumnView_virtualbase_initPainter(const void* self, QPainter* painter);
bool QColumnView_override_virtual_redirected(void* self, intptr_t slot);
QPaintDevice* QColumnView_virtualbase_redirected(const void* self, QPoint* offset);
bool QColumnView_override_virtual_sharedPainter(void* self, intptr_t slot);
QPainter* QColumnView_virtualbase_sharedPainter(const void* self);
bool QColumnView_override_virtual_childEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_childEvent(void* self, QChildEvent* event);
bool QColumnView_override_virtual_customEvent(void* self, intptr_t slot);
void QColumnView_virtualbase_customEvent(void* self, QEvent* event);
bool QColumnView_override_virtual_connectNotify(void* self, intptr_t slot);
void QColumnView_virtualbase_connectNotify(void* self, QMetaMethod* signal);
bool QColumnView_override_virtual_disconnectNotify(void* self, intptr_t slot);
void QColumnView_virtualbase_disconnectNotify(void* self, QMetaMethod* signal);
void QColumnView_protectedbase_initializeColumn(bool* _dynamic_cast_ok, const void* self, QAbstractItemView* column);
int QColumnView_protectedbase_state(bool* _dynamic_cast_ok, const void* self);
void QColumnView_protectedbase_setState(bool* _dynamic_cast_ok, void* self, int state);
void QColumnView_protectedbase_scheduleDelayedItemsLayout(bool* _dynamic_cast_ok, void* self);
void QColumnView_protectedbase_executeDelayedItemsLayout(bool* _dynamic_cast_ok, void* self);
void QColumnView_protectedbase_setDirtyRegion(bool* _dynamic_cast_ok, void* self, QRegion* region);
void QColumnView_protectedbase_scrollDirtyRegion(bool* _dynamic_cast_ok, void* self, int dx, int dy);
QPoint* QColumnView_protectedbase_dirtyRegionOffset(bool* _dynamic_cast_ok, const void* self);
void QColumnView_protectedbase_startAutoScroll(bool* _dynamic_cast_ok, void* self);
void QColumnView_protectedbase_stopAutoScroll(bool* _dynamic_cast_ok, void* self);
void QColumnView_protectedbase_doAutoScroll(bool* _dynamic_cast_ok, void* self);
int QColumnView_protectedbase_dropIndicatorPosition(bool* _dynamic_cast_ok, const void* self);
void QColumnView_protectedbase_setViewportMargins(bool* _dynamic_cast_ok, void* self, int left, int top, int right, int bottom);
QMargins* QColumnView_protectedbase_viewportMargins(bool* _dynamic_cast_ok, const void* self);
void QColumnView_protectedbase_drawFrame(bool* _dynamic_cast_ok, void* self, QPainter* param1);
void QColumnView_protectedbase_updateMicroFocus(bool* _dynamic_cast_ok, void* self);
void QColumnView_protectedbase_create(bool* _dynamic_cast_ok, void* self);
void QColumnView_protectedbase_destroy(bool* _dynamic_cast_ok, void* self);
bool QColumnView_protectedbase_focusNextChild(bool* _dynamic_cast_ok, void* self);
bool QColumnView_protectedbase_focusPreviousChild(bool* _dynamic_cast_ok, void* self);
QObject* QColumnView_protectedbase_sender(bool* _dynamic_cast_ok, const void* self);
int QColumnView_protectedbase_senderSignalIndex(bool* _dynamic_cast_ok, const void* self);
int QColumnView_protectedbase_receivers(bool* _dynamic_cast_ok, const void* self, const char* signal);
bool QColumnView_protectedbase_isSignalConnected(bool* _dynamic_cast_ok, const void* self, QMetaMethod* signal);
void QColumnView_delete(QColumnView* self);

#ifdef __cplusplus
} /* extern C */
#endif

#endif
