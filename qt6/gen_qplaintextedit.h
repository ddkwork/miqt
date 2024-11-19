#pragma once
#ifndef MIQT_QT6_GEN_QPLAINTEXTEDIT_H
#define MIQT_QT6_GEN_QPLAINTEXTEDIT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractScrollArea;
class QAbstractTextDocumentLayout;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QAbstractTextDocumentLayout__PaintContext)
typedef QAbstractTextDocumentLayout::PaintContext QAbstractTextDocumentLayout__PaintContext;
#else
class QAbstractTextDocumentLayout__PaintContext;
#endif
class QContextMenuEvent;
class QDragEnterEvent;
class QDragLeaveEvent;
class QDragMoveEvent;
class QDropEvent;
class QEvent;
class QFocusEvent;
class QFrame;
class QInputMethodEvent;
class QKeyEvent;
class QMenu;
class QMetaObject;
class QMimeData;
class QMouseEvent;
class QObject;
class QPagedPaintDevice;
class QPaintDevice;
class QPaintEvent;
class QPainter;
class QPlainTextDocumentLayout;
class QPlainTextEdit;
class QPoint;
class QPointF;
class QRect;
class QRectF;
class QRegularExpression;
class QResizeEvent;
class QShowEvent;
class QSize;
class QSizeF;
class QTextBlock;
class QTextCharFormat;
class QTextCursor;
class QTextDocument;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QTextEdit__ExtraSelection)
typedef QTextEdit::ExtraSelection QTextEdit__ExtraSelection;
#else
class QTextEdit__ExtraSelection;
#endif
class QTextFormat;
class QTextFrame;
class QTextInlineObject;
class QTimerEvent;
class QUrl;
class QVariant;
class QWheelEvent;
class QWidget;
#else
typedef struct QAbstractScrollArea QAbstractScrollArea;
typedef struct QAbstractTextDocumentLayout QAbstractTextDocumentLayout;
typedef struct QAbstractTextDocumentLayout__PaintContext QAbstractTextDocumentLayout__PaintContext;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDragEnterEvent QDragEnterEvent;
typedef struct QDragLeaveEvent QDragLeaveEvent;
typedef struct QDragMoveEvent QDragMoveEvent;
typedef struct QDropEvent QDropEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QFrame QFrame;
typedef struct QInputMethodEvent QInputMethodEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMenu QMenu;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPagedPaintDevice QPagedPaintDevice;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPainter QPainter;
typedef struct QPlainTextDocumentLayout QPlainTextDocumentLayout;
typedef struct QPlainTextEdit QPlainTextEdit;
typedef struct QPoint QPoint;
typedef struct QPointF QPointF;
typedef struct QRect QRect;
typedef struct QRectF QRectF;
typedef struct QRegularExpression QRegularExpression;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QSizeF QSizeF;
typedef struct QTextBlock QTextBlock;
typedef struct QTextCharFormat QTextCharFormat;
typedef struct QTextCursor QTextCursor;
typedef struct QTextDocument QTextDocument;
typedef struct QTextEdit__ExtraSelection QTextEdit__ExtraSelection;
typedef struct QTextFormat QTextFormat;
typedef struct QTextFrame QTextFrame;
typedef struct QTextInlineObject QTextInlineObject;
typedef struct QTimerEvent QTimerEvent;
typedef struct QUrl QUrl;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
#endif

void QPlainTextEdit_new(QWidget* parent, QPlainTextEdit** outptr_QPlainTextEdit, QAbstractScrollArea** outptr_QAbstractScrollArea, QFrame** outptr_QFrame, QWidget** outptr_QWidget, QObject** outptr_QObject, QPaintDevice** outptr_QPaintDevice);
void QPlainTextEdit_new2(QPlainTextEdit** outptr_QPlainTextEdit, QAbstractScrollArea** outptr_QAbstractScrollArea, QFrame** outptr_QFrame, QWidget** outptr_QWidget, QObject** outptr_QObject, QPaintDevice** outptr_QPaintDevice);
void QPlainTextEdit_new3(struct miqt_string text, QPlainTextEdit** outptr_QPlainTextEdit, QAbstractScrollArea** outptr_QAbstractScrollArea, QFrame** outptr_QFrame, QWidget** outptr_QWidget, QObject** outptr_QObject, QPaintDevice** outptr_QPaintDevice);
void QPlainTextEdit_new4(struct miqt_string text, QWidget* parent, QPlainTextEdit** outptr_QPlainTextEdit, QAbstractScrollArea** outptr_QAbstractScrollArea, QFrame** outptr_QFrame, QWidget** outptr_QWidget, QObject** outptr_QObject, QPaintDevice** outptr_QPaintDevice);
QMetaObject* QPlainTextEdit_MetaObject(const QPlainTextEdit* self);
void* QPlainTextEdit_Metacast(QPlainTextEdit* self, const char* param1);
struct miqt_string QPlainTextEdit_Tr(const char* s);
void QPlainTextEdit_SetDocument(QPlainTextEdit* self, QTextDocument* document);
QTextDocument* QPlainTextEdit_Document(const QPlainTextEdit* self);
void QPlainTextEdit_SetPlaceholderText(QPlainTextEdit* self, struct miqt_string placeholderText);
struct miqt_string QPlainTextEdit_PlaceholderText(const QPlainTextEdit* self);
void QPlainTextEdit_SetTextCursor(QPlainTextEdit* self, QTextCursor* cursor);
QTextCursor* QPlainTextEdit_TextCursor(const QPlainTextEdit* self);
bool QPlainTextEdit_IsReadOnly(const QPlainTextEdit* self);
void QPlainTextEdit_SetReadOnly(QPlainTextEdit* self, bool ro);
void QPlainTextEdit_SetTextInteractionFlags(QPlainTextEdit* self, int flags);
int QPlainTextEdit_TextInteractionFlags(const QPlainTextEdit* self);
void QPlainTextEdit_MergeCurrentCharFormat(QPlainTextEdit* self, QTextCharFormat* modifier);
void QPlainTextEdit_SetCurrentCharFormat(QPlainTextEdit* self, QTextCharFormat* format);
QTextCharFormat* QPlainTextEdit_CurrentCharFormat(const QPlainTextEdit* self);
bool QPlainTextEdit_TabChangesFocus(const QPlainTextEdit* self);
void QPlainTextEdit_SetTabChangesFocus(QPlainTextEdit* self, bool b);
void QPlainTextEdit_SetDocumentTitle(QPlainTextEdit* self, struct miqt_string title);
struct miqt_string QPlainTextEdit_DocumentTitle(const QPlainTextEdit* self);
bool QPlainTextEdit_IsUndoRedoEnabled(const QPlainTextEdit* self);
void QPlainTextEdit_SetUndoRedoEnabled(QPlainTextEdit* self, bool enable);
void QPlainTextEdit_SetMaximumBlockCount(QPlainTextEdit* self, int maximum);
int QPlainTextEdit_MaximumBlockCount(const QPlainTextEdit* self);
int QPlainTextEdit_LineWrapMode(const QPlainTextEdit* self);
void QPlainTextEdit_SetLineWrapMode(QPlainTextEdit* self, int mode);
int QPlainTextEdit_WordWrapMode(const QPlainTextEdit* self);
void QPlainTextEdit_SetWordWrapMode(QPlainTextEdit* self, int policy);
void QPlainTextEdit_SetBackgroundVisible(QPlainTextEdit* self, bool visible);
bool QPlainTextEdit_BackgroundVisible(const QPlainTextEdit* self);
void QPlainTextEdit_SetCenterOnScroll(QPlainTextEdit* self, bool enabled);
bool QPlainTextEdit_CenterOnScroll(const QPlainTextEdit* self);
bool QPlainTextEdit_Find(QPlainTextEdit* self, struct miqt_string exp);
bool QPlainTextEdit_FindWithExp(QPlainTextEdit* self, QRegularExpression* exp);
struct miqt_string QPlainTextEdit_ToPlainText(const QPlainTextEdit* self);
void QPlainTextEdit_EnsureCursorVisible(QPlainTextEdit* self);
QVariant* QPlainTextEdit_LoadResource(QPlainTextEdit* self, int typeVal, QUrl* name);
QMenu* QPlainTextEdit_CreateStandardContextMenu(QPlainTextEdit* self);
QMenu* QPlainTextEdit_CreateStandardContextMenuWithPosition(QPlainTextEdit* self, QPoint* position);
QTextCursor* QPlainTextEdit_CursorForPosition(const QPlainTextEdit* self, QPoint* pos);
QRect* QPlainTextEdit_CursorRect(const QPlainTextEdit* self, QTextCursor* cursor);
QRect* QPlainTextEdit_CursorRect2(const QPlainTextEdit* self);
struct miqt_string QPlainTextEdit_AnchorAt(const QPlainTextEdit* self, QPoint* pos);
bool QPlainTextEdit_OverwriteMode(const QPlainTextEdit* self);
void QPlainTextEdit_SetOverwriteMode(QPlainTextEdit* self, bool overwrite);
double QPlainTextEdit_TabStopDistance(const QPlainTextEdit* self);
void QPlainTextEdit_SetTabStopDistance(QPlainTextEdit* self, double distance);
int QPlainTextEdit_CursorWidth(const QPlainTextEdit* self);
void QPlainTextEdit_SetCursorWidth(QPlainTextEdit* self, int width);
void QPlainTextEdit_SetExtraSelections(QPlainTextEdit* self, struct miqt_array /* of QTextEdit__ExtraSelection* */  selections);
struct miqt_array /* of QTextEdit__ExtraSelection* */  QPlainTextEdit_ExtraSelections(const QPlainTextEdit* self);
void QPlainTextEdit_MoveCursor(QPlainTextEdit* self, int operation);
bool QPlainTextEdit_CanPaste(const QPlainTextEdit* self);
void QPlainTextEdit_Print(const QPlainTextEdit* self, QPagedPaintDevice* printer);
int QPlainTextEdit_BlockCount(const QPlainTextEdit* self);
QVariant* QPlainTextEdit_InputMethodQuery(const QPlainTextEdit* self, int property);
QVariant* QPlainTextEdit_InputMethodQuery2(const QPlainTextEdit* self, int query, QVariant* argument);
void QPlainTextEdit_SetPlainText(QPlainTextEdit* self, struct miqt_string text);
void QPlainTextEdit_Cut(QPlainTextEdit* self);
void QPlainTextEdit_Copy(QPlainTextEdit* self);
void QPlainTextEdit_Paste(QPlainTextEdit* self);
void QPlainTextEdit_Undo(QPlainTextEdit* self);
void QPlainTextEdit_Redo(QPlainTextEdit* self);
void QPlainTextEdit_Clear(QPlainTextEdit* self);
void QPlainTextEdit_SelectAll(QPlainTextEdit* self);
void QPlainTextEdit_InsertPlainText(QPlainTextEdit* self, struct miqt_string text);
void QPlainTextEdit_AppendPlainText(QPlainTextEdit* self, struct miqt_string text);
void QPlainTextEdit_AppendHtml(QPlainTextEdit* self, struct miqt_string html);
void QPlainTextEdit_CenterCursor(QPlainTextEdit* self);
void QPlainTextEdit_ZoomIn(QPlainTextEdit* self);
void QPlainTextEdit_ZoomOut(QPlainTextEdit* self);
void QPlainTextEdit_TextChanged(QPlainTextEdit* self);
void QPlainTextEdit_connect_TextChanged(QPlainTextEdit* self, intptr_t slot);
void QPlainTextEdit_UndoAvailable(QPlainTextEdit* self, bool b);
void QPlainTextEdit_connect_UndoAvailable(QPlainTextEdit* self, intptr_t slot);
void QPlainTextEdit_RedoAvailable(QPlainTextEdit* self, bool b);
void QPlainTextEdit_connect_RedoAvailable(QPlainTextEdit* self, intptr_t slot);
void QPlainTextEdit_CopyAvailable(QPlainTextEdit* self, bool b);
void QPlainTextEdit_connect_CopyAvailable(QPlainTextEdit* self, intptr_t slot);
void QPlainTextEdit_SelectionChanged(QPlainTextEdit* self);
void QPlainTextEdit_connect_SelectionChanged(QPlainTextEdit* self, intptr_t slot);
void QPlainTextEdit_CursorPositionChanged(QPlainTextEdit* self);
void QPlainTextEdit_connect_CursorPositionChanged(QPlainTextEdit* self, intptr_t slot);
void QPlainTextEdit_UpdateRequest(QPlainTextEdit* self, QRect* rect, int dy);
void QPlainTextEdit_connect_UpdateRequest(QPlainTextEdit* self, intptr_t slot);
void QPlainTextEdit_BlockCountChanged(QPlainTextEdit* self, int newBlockCount);
void QPlainTextEdit_connect_BlockCountChanged(QPlainTextEdit* self, intptr_t slot);
void QPlainTextEdit_ModificationChanged(QPlainTextEdit* self, bool param1);
void QPlainTextEdit_connect_ModificationChanged(QPlainTextEdit* self, intptr_t slot);
bool QPlainTextEdit_Event(QPlainTextEdit* self, QEvent* e);
void QPlainTextEdit_TimerEvent(QPlainTextEdit* self, QTimerEvent* e);
void QPlainTextEdit_KeyPressEvent(QPlainTextEdit* self, QKeyEvent* e);
void QPlainTextEdit_KeyReleaseEvent(QPlainTextEdit* self, QKeyEvent* e);
void QPlainTextEdit_ResizeEvent(QPlainTextEdit* self, QResizeEvent* e);
void QPlainTextEdit_PaintEvent(QPlainTextEdit* self, QPaintEvent* e);
void QPlainTextEdit_MousePressEvent(QPlainTextEdit* self, QMouseEvent* e);
void QPlainTextEdit_MouseMoveEvent(QPlainTextEdit* self, QMouseEvent* e);
void QPlainTextEdit_MouseReleaseEvent(QPlainTextEdit* self, QMouseEvent* e);
void QPlainTextEdit_MouseDoubleClickEvent(QPlainTextEdit* self, QMouseEvent* e);
bool QPlainTextEdit_FocusNextPrevChild(QPlainTextEdit* self, bool next);
void QPlainTextEdit_ContextMenuEvent(QPlainTextEdit* self, QContextMenuEvent* e);
void QPlainTextEdit_DragEnterEvent(QPlainTextEdit* self, QDragEnterEvent* e);
void QPlainTextEdit_DragLeaveEvent(QPlainTextEdit* self, QDragLeaveEvent* e);
void QPlainTextEdit_DragMoveEvent(QPlainTextEdit* self, QDragMoveEvent* e);
void QPlainTextEdit_DropEvent(QPlainTextEdit* self, QDropEvent* e);
void QPlainTextEdit_FocusInEvent(QPlainTextEdit* self, QFocusEvent* e);
void QPlainTextEdit_FocusOutEvent(QPlainTextEdit* self, QFocusEvent* e);
void QPlainTextEdit_ShowEvent(QPlainTextEdit* self, QShowEvent* param1);
void QPlainTextEdit_ChangeEvent(QPlainTextEdit* self, QEvent* e);
void QPlainTextEdit_WheelEvent(QPlainTextEdit* self, QWheelEvent* e);
QMimeData* QPlainTextEdit_CreateMimeDataFromSelection(const QPlainTextEdit* self);
bool QPlainTextEdit_CanInsertFromMimeData(const QPlainTextEdit* self, QMimeData* source);
void QPlainTextEdit_InsertFromMimeData(QPlainTextEdit* self, QMimeData* source);
void QPlainTextEdit_InputMethodEvent(QPlainTextEdit* self, QInputMethodEvent* param1);
void QPlainTextEdit_ScrollContentsBy(QPlainTextEdit* self, int dx, int dy);
void QPlainTextEdit_DoSetTextCursor(QPlainTextEdit* self, QTextCursor* cursor);
struct miqt_string QPlainTextEdit_Tr2(const char* s, const char* c);
struct miqt_string QPlainTextEdit_Tr3(const char* s, const char* c, int n);
bool QPlainTextEdit_Find2(QPlainTextEdit* self, struct miqt_string exp, int options);
bool QPlainTextEdit_Find22(QPlainTextEdit* self, QRegularExpression* exp, int options);
void QPlainTextEdit_MoveCursor2(QPlainTextEdit* self, int operation, int mode);
void QPlainTextEdit_ZoomIn1(QPlainTextEdit* self, int rangeVal);
void QPlainTextEdit_ZoomOut1(QPlainTextEdit* self, int rangeVal);
void QPlainTextEdit_override_virtual_LoadResource(void* self, intptr_t slot);
QVariant* QPlainTextEdit_virtualbase_LoadResource(void* self, int typeVal, QUrl* name);
void QPlainTextEdit_override_virtual_InputMethodQuery(void* self, intptr_t slot);
QVariant* QPlainTextEdit_virtualbase_InputMethodQuery(const void* self, int property);
void QPlainTextEdit_override_virtual_Event(void* self, intptr_t slot);
bool QPlainTextEdit_virtualbase_Event(void* self, QEvent* e);
void QPlainTextEdit_override_virtual_TimerEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_TimerEvent(void* self, QTimerEvent* e);
void QPlainTextEdit_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_KeyPressEvent(void* self, QKeyEvent* e);
void QPlainTextEdit_override_virtual_KeyReleaseEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* e);
void QPlainTextEdit_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_ResizeEvent(void* self, QResizeEvent* e);
void QPlainTextEdit_override_virtual_PaintEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_PaintEvent(void* self, QPaintEvent* e);
void QPlainTextEdit_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_MousePressEvent(void* self, QMouseEvent* e);
void QPlainTextEdit_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_MouseMoveEvent(void* self, QMouseEvent* e);
void QPlainTextEdit_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* e);
void QPlainTextEdit_override_virtual_MouseDoubleClickEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_MouseDoubleClickEvent(void* self, QMouseEvent* e);
void QPlainTextEdit_override_virtual_FocusNextPrevChild(void* self, intptr_t slot);
bool QPlainTextEdit_virtualbase_FocusNextPrevChild(void* self, bool next);
void QPlainTextEdit_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* e);
void QPlainTextEdit_override_virtual_DragEnterEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_DragEnterEvent(void* self, QDragEnterEvent* e);
void QPlainTextEdit_override_virtual_DragLeaveEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* e);
void QPlainTextEdit_override_virtual_DragMoveEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* e);
void QPlainTextEdit_override_virtual_DropEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_DropEvent(void* self, QDropEvent* e);
void QPlainTextEdit_override_virtual_FocusInEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_FocusInEvent(void* self, QFocusEvent* e);
void QPlainTextEdit_override_virtual_FocusOutEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_FocusOutEvent(void* self, QFocusEvent* e);
void QPlainTextEdit_override_virtual_ShowEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_ShowEvent(void* self, QShowEvent* param1);
void QPlainTextEdit_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_ChangeEvent(void* self, QEvent* e);
void QPlainTextEdit_override_virtual_WheelEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_WheelEvent(void* self, QWheelEvent* e);
void QPlainTextEdit_override_virtual_CreateMimeDataFromSelection(void* self, intptr_t slot);
QMimeData* QPlainTextEdit_virtualbase_CreateMimeDataFromSelection(const void* self);
void QPlainTextEdit_override_virtual_CanInsertFromMimeData(void* self, intptr_t slot);
bool QPlainTextEdit_virtualbase_CanInsertFromMimeData(const void* self, QMimeData* source);
void QPlainTextEdit_override_virtual_InsertFromMimeData(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_InsertFromMimeData(void* self, QMimeData* source);
void QPlainTextEdit_override_virtual_InputMethodEvent(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_InputMethodEvent(void* self, QInputMethodEvent* param1);
void QPlainTextEdit_override_virtual_ScrollContentsBy(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_ScrollContentsBy(void* self, int dx, int dy);
void QPlainTextEdit_override_virtual_DoSetTextCursor(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_DoSetTextCursor(void* self, QTextCursor* cursor);
void QPlainTextEdit_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QPlainTextEdit_virtualbase_MinimumSizeHint(const void* self);
void QPlainTextEdit_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QPlainTextEdit_virtualbase_SizeHint(const void* self);
void QPlainTextEdit_override_virtual_SetupViewport(void* self, intptr_t slot);
void QPlainTextEdit_virtualbase_SetupViewport(void* self, QWidget* viewport);
void QPlainTextEdit_override_virtual_EventFilter(void* self, intptr_t slot);
bool QPlainTextEdit_virtualbase_EventFilter(void* self, QObject* param1, QEvent* param2);
void QPlainTextEdit_override_virtual_ViewportEvent(void* self, intptr_t slot);
bool QPlainTextEdit_virtualbase_ViewportEvent(void* self, QEvent* param1);
void QPlainTextEdit_override_virtual_ViewportSizeHint(void* self, intptr_t slot);
QSize* QPlainTextEdit_virtualbase_ViewportSizeHint(const void* self);
void QPlainTextEdit_Delete(QPlainTextEdit* self, bool isSubclass);

void QPlainTextDocumentLayout_new(QTextDocument* document, QPlainTextDocumentLayout** outptr_QPlainTextDocumentLayout, QAbstractTextDocumentLayout** outptr_QAbstractTextDocumentLayout, QObject** outptr_QObject);
QMetaObject* QPlainTextDocumentLayout_MetaObject(const QPlainTextDocumentLayout* self);
void* QPlainTextDocumentLayout_Metacast(QPlainTextDocumentLayout* self, const char* param1);
struct miqt_string QPlainTextDocumentLayout_Tr(const char* s);
void QPlainTextDocumentLayout_Draw(QPlainTextDocumentLayout* self, QPainter* param1, QAbstractTextDocumentLayout__PaintContext* param2);
int QPlainTextDocumentLayout_HitTest(const QPlainTextDocumentLayout* self, QPointF* param1, int param2);
int QPlainTextDocumentLayout_PageCount(const QPlainTextDocumentLayout* self);
QSizeF* QPlainTextDocumentLayout_DocumentSize(const QPlainTextDocumentLayout* self);
QRectF* QPlainTextDocumentLayout_FrameBoundingRect(const QPlainTextDocumentLayout* self, QTextFrame* param1);
QRectF* QPlainTextDocumentLayout_BlockBoundingRect(const QPlainTextDocumentLayout* self, QTextBlock* block);
void QPlainTextDocumentLayout_EnsureBlockLayout(const QPlainTextDocumentLayout* self, QTextBlock* block);
void QPlainTextDocumentLayout_SetCursorWidth(QPlainTextDocumentLayout* self, int width);
int QPlainTextDocumentLayout_CursorWidth(const QPlainTextDocumentLayout* self);
void QPlainTextDocumentLayout_RequestUpdate(QPlainTextDocumentLayout* self);
void QPlainTextDocumentLayout_DocumentChanged(QPlainTextDocumentLayout* self, int from, int param2, int charsAdded);
struct miqt_string QPlainTextDocumentLayout_Tr2(const char* s, const char* c);
struct miqt_string QPlainTextDocumentLayout_Tr3(const char* s, const char* c, int n);
void QPlainTextDocumentLayout_override_virtual_Draw(void* self, intptr_t slot);
void QPlainTextDocumentLayout_virtualbase_Draw(void* self, QPainter* param1, QAbstractTextDocumentLayout__PaintContext* param2);
void QPlainTextDocumentLayout_override_virtual_HitTest(void* self, intptr_t slot);
int QPlainTextDocumentLayout_virtualbase_HitTest(const void* self, QPointF* param1, int param2);
void QPlainTextDocumentLayout_override_virtual_PageCount(void* self, intptr_t slot);
int QPlainTextDocumentLayout_virtualbase_PageCount(const void* self);
void QPlainTextDocumentLayout_override_virtual_DocumentSize(void* self, intptr_t slot);
QSizeF* QPlainTextDocumentLayout_virtualbase_DocumentSize(const void* self);
void QPlainTextDocumentLayout_override_virtual_FrameBoundingRect(void* self, intptr_t slot);
QRectF* QPlainTextDocumentLayout_virtualbase_FrameBoundingRect(const void* self, QTextFrame* param1);
void QPlainTextDocumentLayout_override_virtual_BlockBoundingRect(void* self, intptr_t slot);
QRectF* QPlainTextDocumentLayout_virtualbase_BlockBoundingRect(const void* self, QTextBlock* block);
void QPlainTextDocumentLayout_override_virtual_DocumentChanged(void* self, intptr_t slot);
void QPlainTextDocumentLayout_virtualbase_DocumentChanged(void* self, int from, int param2, int charsAdded);
void QPlainTextDocumentLayout_override_virtual_ResizeInlineObject(void* self, intptr_t slot);
void QPlainTextDocumentLayout_virtualbase_ResizeInlineObject(void* self, QTextInlineObject* item, int posInDocument, QTextFormat* format);
void QPlainTextDocumentLayout_override_virtual_PositionInlineObject(void* self, intptr_t slot);
void QPlainTextDocumentLayout_virtualbase_PositionInlineObject(void* self, QTextInlineObject* item, int posInDocument, QTextFormat* format);
void QPlainTextDocumentLayout_override_virtual_DrawInlineObject(void* self, intptr_t slot);
void QPlainTextDocumentLayout_virtualbase_DrawInlineObject(void* self, QPainter* painter, QRectF* rect, QTextInlineObject* object, int posInDocument, QTextFormat* format);
void QPlainTextDocumentLayout_Delete(QPlainTextDocumentLayout* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
