#pragma once
#ifndef MIQT_QT_GEN_QLINEEDIT_H
#define MIQT_QT_GEN_QLINEEDIT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAction;
class QActionEvent;
class QChildEvent;
class QCloseEvent;
class QCompleter;
class QContextMenuEvent;
class QDragEnterEvent;
class QDragLeaveEvent;
class QDragMoveEvent;
class QDropEvent;
class QEvent;
class QFocusEvent;
class QHideEvent;
class QIcon;
class QInputMethodEvent;
class QKeyEvent;
class QLineEdit;
class QMargins;
class QMenu;
class QMetaMethod;
class QMetaObject;
class QMouseEvent;
class QMoveEvent;
class QObject;
class QPaintDevice;
class QPaintEngine;
class QPaintEvent;
class QPainter;
class QPoint;
class QRect;
class QResizeEvent;
class QShowEvent;
class QSize;
class QStyleOptionFrame;
class QTabletEvent;
class QTimerEvent;
class QValidator;
class QVariant;
class QWheelEvent;
class QWidget;
#else
typedef struct QAction QAction;
typedef struct QActionEvent QActionEvent;
typedef struct QChildEvent QChildEvent;
typedef struct QCloseEvent QCloseEvent;
typedef struct QCompleter QCompleter;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDragEnterEvent QDragEnterEvent;
typedef struct QDragLeaveEvent QDragLeaveEvent;
typedef struct QDragMoveEvent QDragMoveEvent;
typedef struct QDropEvent QDropEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QHideEvent QHideEvent;
typedef struct QIcon QIcon;
typedef struct QInputMethodEvent QInputMethodEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QLineEdit QLineEdit;
typedef struct QMargins QMargins;
typedef struct QMenu QMenu;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QMoveEvent QMoveEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPainter QPainter;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QStyleOptionFrame QStyleOptionFrame;
typedef struct QTabletEvent QTabletEvent;
typedef struct QTimerEvent QTimerEvent;
typedef struct QValidator QValidator;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
#endif

QLineEdit* QLineEdit_new(QWidget* parent);
QLineEdit* QLineEdit_new2();
QLineEdit* QLineEdit_new3(struct miqt_string param1);
QLineEdit* QLineEdit_new4(struct miqt_string param1, QWidget* parent);
void QLineEdit_virtbase(QLineEdit* src, QWidget** outptr_QWidget);
QMetaObject* QLineEdit_metaObject(const QLineEdit* self);
void* QLineEdit_metacast(QLineEdit* self, const char* param1);
struct miqt_string QLineEdit_tr(const char* s);
struct miqt_string QLineEdit_trUtf8(const char* s);
struct miqt_string QLineEdit_text(const QLineEdit* self);
struct miqt_string QLineEdit_displayText(const QLineEdit* self);
struct miqt_string QLineEdit_placeholderText(const QLineEdit* self);
void QLineEdit_setPlaceholderText(QLineEdit* self, struct miqt_string placeholderText);
int QLineEdit_maxLength(const QLineEdit* self);
void QLineEdit_setMaxLength(QLineEdit* self, int maxLength);
void QLineEdit_setFrame(QLineEdit* self, bool frame);
bool QLineEdit_hasFrame(const QLineEdit* self);
void QLineEdit_setClearButtonEnabled(QLineEdit* self, bool enable);
bool QLineEdit_isClearButtonEnabled(const QLineEdit* self);
int QLineEdit_echoMode(const QLineEdit* self);
void QLineEdit_setEchoMode(QLineEdit* self, int echoMode);
bool QLineEdit_isReadOnly(const QLineEdit* self);
void QLineEdit_setReadOnly(QLineEdit* self, bool readOnly);
void QLineEdit_setValidator(QLineEdit* self, QValidator* validator);
QValidator* QLineEdit_validator(const QLineEdit* self);
void QLineEdit_setCompleter(QLineEdit* self, QCompleter* completer);
QCompleter* QLineEdit_completer(const QLineEdit* self);
QSize* QLineEdit_sizeHint(const QLineEdit* self);
QSize* QLineEdit_minimumSizeHint(const QLineEdit* self);
int QLineEdit_cursorPosition(const QLineEdit* self);
void QLineEdit_setCursorPosition(QLineEdit* self, int cursorPosition);
int QLineEdit_cursorPositionAt(QLineEdit* self, QPoint* pos);
void QLineEdit_setAlignment(QLineEdit* self, int flag);
int QLineEdit_alignment(const QLineEdit* self);
void QLineEdit_cursorForward(QLineEdit* self, bool mark);
void QLineEdit_cursorBackward(QLineEdit* self, bool mark);
void QLineEdit_cursorWordForward(QLineEdit* self, bool mark);
void QLineEdit_cursorWordBackward(QLineEdit* self, bool mark);
void QLineEdit_backspace(QLineEdit* self);
void QLineEdit_del(QLineEdit* self);
void QLineEdit_home(QLineEdit* self, bool mark);
void QLineEdit_end(QLineEdit* self, bool mark);
bool QLineEdit_isModified(const QLineEdit* self);
void QLineEdit_setModified(QLineEdit* self, bool modified);
void QLineEdit_setSelection(QLineEdit* self, int param1, int param2);
bool QLineEdit_hasSelectedText(const QLineEdit* self);
struct miqt_string QLineEdit_selectedText(const QLineEdit* self);
int QLineEdit_selectionStart(const QLineEdit* self);
int QLineEdit_selectionEnd(const QLineEdit* self);
int QLineEdit_selectionLength(const QLineEdit* self);
bool QLineEdit_isUndoAvailable(const QLineEdit* self);
bool QLineEdit_isRedoAvailable(const QLineEdit* self);
void QLineEdit_setDragEnabled(QLineEdit* self, bool b);
bool QLineEdit_dragEnabled(const QLineEdit* self);
void QLineEdit_setCursorMoveStyle(QLineEdit* self, int style);
int QLineEdit_cursorMoveStyle(const QLineEdit* self);
struct miqt_string QLineEdit_inputMask(const QLineEdit* self);
void QLineEdit_setInputMask(QLineEdit* self, struct miqt_string inputMask);
bool QLineEdit_hasAcceptableInput(const QLineEdit* self);
void QLineEdit_setTextMargins(QLineEdit* self, int left, int top, int right, int bottom);
void QLineEdit_setTextMarginsWithMargins(QLineEdit* self, QMargins* margins);
void QLineEdit_getTextMargins(const QLineEdit* self, int* left, int* top, int* right, int* bottom);
QMargins* QLineEdit_textMargins(const QLineEdit* self);
void QLineEdit_addAction(QLineEdit* self, QAction* action, int position);
QAction* QLineEdit_addAction2(QLineEdit* self, QIcon* icon, int position);
void QLineEdit_setText(QLineEdit* self, struct miqt_string text);
void QLineEdit_clear(QLineEdit* self);
void QLineEdit_selectAll(QLineEdit* self);
void QLineEdit_undo(QLineEdit* self);
void QLineEdit_redo(QLineEdit* self);
void QLineEdit_cut(QLineEdit* self);
void QLineEdit_copy(const QLineEdit* self);
void QLineEdit_paste(QLineEdit* self);
void QLineEdit_deselect(QLineEdit* self);
void QLineEdit_insert(QLineEdit* self, struct miqt_string param1);
QMenu* QLineEdit_createStandardContextMenu(QLineEdit* self);
void QLineEdit_textChanged(QLineEdit* self, struct miqt_string param1);
void QLineEdit_connect_textChanged(QLineEdit* self, intptr_t slot);
void QLineEdit_textEdited(QLineEdit* self, struct miqt_string param1);
void QLineEdit_connect_textEdited(QLineEdit* self, intptr_t slot);
void QLineEdit_cursorPositionChanged(QLineEdit* self, int param1, int param2);
void QLineEdit_connect_cursorPositionChanged(QLineEdit* self, intptr_t slot);
void QLineEdit_returnPressed(QLineEdit* self);
void QLineEdit_connect_returnPressed(QLineEdit* self, intptr_t slot);
void QLineEdit_editingFinished(QLineEdit* self);
void QLineEdit_connect_editingFinished(QLineEdit* self, intptr_t slot);
void QLineEdit_selectionChanged(QLineEdit* self);
void QLineEdit_connect_selectionChanged(QLineEdit* self, intptr_t slot);
void QLineEdit_inputRejected(QLineEdit* self);
void QLineEdit_connect_inputRejected(QLineEdit* self, intptr_t slot);
void QLineEdit_mousePressEvent(QLineEdit* self, QMouseEvent* param1);
void QLineEdit_mouseMoveEvent(QLineEdit* self, QMouseEvent* param1);
void QLineEdit_mouseReleaseEvent(QLineEdit* self, QMouseEvent* param1);
void QLineEdit_mouseDoubleClickEvent(QLineEdit* self, QMouseEvent* param1);
void QLineEdit_keyPressEvent(QLineEdit* self, QKeyEvent* param1);
void QLineEdit_focusInEvent(QLineEdit* self, QFocusEvent* param1);
void QLineEdit_focusOutEvent(QLineEdit* self, QFocusEvent* param1);
void QLineEdit_paintEvent(QLineEdit* self, QPaintEvent* param1);
void QLineEdit_dragEnterEvent(QLineEdit* self, QDragEnterEvent* param1);
void QLineEdit_dragMoveEvent(QLineEdit* self, QDragMoveEvent* e);
void QLineEdit_dragLeaveEvent(QLineEdit* self, QDragLeaveEvent* e);
void QLineEdit_dropEvent(QLineEdit* self, QDropEvent* param1);
void QLineEdit_changeEvent(QLineEdit* self, QEvent* param1);
void QLineEdit_contextMenuEvent(QLineEdit* self, QContextMenuEvent* param1);
void QLineEdit_inputMethodEvent(QLineEdit* self, QInputMethodEvent* param1);
QVariant* QLineEdit_inputMethodQuery(const QLineEdit* self, int param1);
QVariant* QLineEdit_inputMethodQuery2(const QLineEdit* self, int property, QVariant* argument);
bool QLineEdit_event(QLineEdit* self, QEvent* param1);
struct miqt_string QLineEdit_tr2(const char* s, const char* c);
struct miqt_string QLineEdit_tr3(const char* s, const char* c, int n);
struct miqt_string QLineEdit_trUtf82(const char* s, const char* c);
struct miqt_string QLineEdit_trUtf83(const char* s, const char* c, int n);
void QLineEdit_cursorForward2(QLineEdit* self, bool mark, int steps);
void QLineEdit_cursorBackward2(QLineEdit* self, bool mark, int steps);
bool QLineEdit_override_virtual_sizeHint(void* self, intptr_t slot);
QSize* QLineEdit_virtualbase_sizeHint(const void* self);
bool QLineEdit_override_virtual_minimumSizeHint(void* self, intptr_t slot);
QSize* QLineEdit_virtualbase_minimumSizeHint(const void* self);
bool QLineEdit_override_virtual_mousePressEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_mousePressEvent(void* self, QMouseEvent* param1);
bool QLineEdit_override_virtual_mouseMoveEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_mouseMoveEvent(void* self, QMouseEvent* param1);
bool QLineEdit_override_virtual_mouseReleaseEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_mouseReleaseEvent(void* self, QMouseEvent* param1);
bool QLineEdit_override_virtual_mouseDoubleClickEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_mouseDoubleClickEvent(void* self, QMouseEvent* param1);
bool QLineEdit_override_virtual_keyPressEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_keyPressEvent(void* self, QKeyEvent* param1);
bool QLineEdit_override_virtual_focusInEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_focusInEvent(void* self, QFocusEvent* param1);
bool QLineEdit_override_virtual_focusOutEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_focusOutEvent(void* self, QFocusEvent* param1);
bool QLineEdit_override_virtual_paintEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_paintEvent(void* self, QPaintEvent* param1);
bool QLineEdit_override_virtual_dragEnterEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_dragEnterEvent(void* self, QDragEnterEvent* param1);
bool QLineEdit_override_virtual_dragMoveEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_dragMoveEvent(void* self, QDragMoveEvent* e);
bool QLineEdit_override_virtual_dragLeaveEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_dragLeaveEvent(void* self, QDragLeaveEvent* e);
bool QLineEdit_override_virtual_dropEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_dropEvent(void* self, QDropEvent* param1);
bool QLineEdit_override_virtual_changeEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_changeEvent(void* self, QEvent* param1);
bool QLineEdit_override_virtual_contextMenuEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_contextMenuEvent(void* self, QContextMenuEvent* param1);
bool QLineEdit_override_virtual_inputMethodEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_inputMethodEvent(void* self, QInputMethodEvent* param1);
bool QLineEdit_override_virtual_inputMethodQuery(void* self, intptr_t slot);
QVariant* QLineEdit_virtualbase_inputMethodQuery(const void* self, int param1);
bool QLineEdit_override_virtual_event(void* self, intptr_t slot);
bool QLineEdit_virtualbase_event(void* self, QEvent* param1);
bool QLineEdit_override_virtual_devType(void* self, intptr_t slot);
int QLineEdit_virtualbase_devType(const void* self);
bool QLineEdit_override_virtual_setVisible(void* self, intptr_t slot);
void QLineEdit_virtualbase_setVisible(void* self, bool visible);
bool QLineEdit_override_virtual_heightForWidth(void* self, intptr_t slot);
int QLineEdit_virtualbase_heightForWidth(const void* self, int param1);
bool QLineEdit_override_virtual_hasHeightForWidth(void* self, intptr_t slot);
bool QLineEdit_virtualbase_hasHeightForWidth(const void* self);
bool QLineEdit_override_virtual_paintEngine(void* self, intptr_t slot);
QPaintEngine* QLineEdit_virtualbase_paintEngine(const void* self);
bool QLineEdit_override_virtual_wheelEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_wheelEvent(void* self, QWheelEvent* event);
bool QLineEdit_override_virtual_keyReleaseEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_keyReleaseEvent(void* self, QKeyEvent* event);
bool QLineEdit_override_virtual_enterEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_enterEvent(void* self, QEvent* event);
bool QLineEdit_override_virtual_leaveEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_leaveEvent(void* self, QEvent* event);
bool QLineEdit_override_virtual_moveEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_moveEvent(void* self, QMoveEvent* event);
bool QLineEdit_override_virtual_resizeEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_resizeEvent(void* self, QResizeEvent* event);
bool QLineEdit_override_virtual_closeEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_closeEvent(void* self, QCloseEvent* event);
bool QLineEdit_override_virtual_tabletEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_tabletEvent(void* self, QTabletEvent* event);
bool QLineEdit_override_virtual_actionEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_actionEvent(void* self, QActionEvent* event);
bool QLineEdit_override_virtual_showEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_showEvent(void* self, QShowEvent* event);
bool QLineEdit_override_virtual_hideEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_hideEvent(void* self, QHideEvent* event);
bool QLineEdit_override_virtual_nativeEvent(void* self, intptr_t slot);
bool QLineEdit_virtualbase_nativeEvent(void* self, struct miqt_string eventType, void* message, long* result);
bool QLineEdit_override_virtual_metric(void* self, intptr_t slot);
int QLineEdit_virtualbase_metric(const void* self, int param1);
bool QLineEdit_override_virtual_initPainter(void* self, intptr_t slot);
void QLineEdit_virtualbase_initPainter(const void* self, QPainter* painter);
bool QLineEdit_override_virtual_redirected(void* self, intptr_t slot);
QPaintDevice* QLineEdit_virtualbase_redirected(const void* self, QPoint* offset);
bool QLineEdit_override_virtual_sharedPainter(void* self, intptr_t slot);
QPainter* QLineEdit_virtualbase_sharedPainter(const void* self);
bool QLineEdit_override_virtual_focusNextPrevChild(void* self, intptr_t slot);
bool QLineEdit_virtualbase_focusNextPrevChild(void* self, bool next);
bool QLineEdit_override_virtual_eventFilter(void* self, intptr_t slot);
bool QLineEdit_virtualbase_eventFilter(void* self, QObject* watched, QEvent* event);
bool QLineEdit_override_virtual_timerEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_timerEvent(void* self, QTimerEvent* event);
bool QLineEdit_override_virtual_childEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_childEvent(void* self, QChildEvent* event);
bool QLineEdit_override_virtual_customEvent(void* self, intptr_t slot);
void QLineEdit_virtualbase_customEvent(void* self, QEvent* event);
bool QLineEdit_override_virtual_connectNotify(void* self, intptr_t slot);
void QLineEdit_virtualbase_connectNotify(void* self, QMetaMethod* signal);
bool QLineEdit_override_virtual_disconnectNotify(void* self, intptr_t slot);
void QLineEdit_virtualbase_disconnectNotify(void* self, QMetaMethod* signal);
void QLineEdit_protectedbase_initStyleOption(bool* _dynamic_cast_ok, const void* self, QStyleOptionFrame* option);
QRect* QLineEdit_protectedbase_cursorRect(bool* _dynamic_cast_ok, const void* self);
void QLineEdit_protectedbase_updateMicroFocus(bool* _dynamic_cast_ok, void* self);
void QLineEdit_protectedbase_create(bool* _dynamic_cast_ok, void* self);
void QLineEdit_protectedbase_destroy(bool* _dynamic_cast_ok, void* self);
bool QLineEdit_protectedbase_focusNextChild(bool* _dynamic_cast_ok, void* self);
bool QLineEdit_protectedbase_focusPreviousChild(bool* _dynamic_cast_ok, void* self);
QObject* QLineEdit_protectedbase_sender(bool* _dynamic_cast_ok, const void* self);
int QLineEdit_protectedbase_senderSignalIndex(bool* _dynamic_cast_ok, const void* self);
int QLineEdit_protectedbase_receivers(bool* _dynamic_cast_ok, const void* self, const char* signal);
bool QLineEdit_protectedbase_isSignalConnected(bool* _dynamic_cast_ok, const void* self, QMetaMethod* signal);
void QLineEdit_delete(QLineEdit* self);

#ifdef __cplusplus
} /* extern C */
#endif

#endif
