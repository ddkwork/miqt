#pragma once
#ifndef MIQT_QT_GEN_QLINEEDIT_H
#define MIQT_QT_GEN_QLINEEDIT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAction QAction;
typedef struct QCompleter QCompleter;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDragEnterEvent QDragEnterEvent;
typedef struct QDragLeaveEvent QDragLeaveEvent;
typedef struct QDragMoveEvent QDragMoveEvent;
typedef struct QDropEvent QDropEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QIcon QIcon;
typedef struct QInputMethodEvent QInputMethodEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QLineEdit QLineEdit;
typedef struct QMargins QMargins;
typedef struct QMenu QMenu;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPoint QPoint;
typedef struct QSize QSize;
typedef struct QStyleOptionFrame QStyleOptionFrame;
typedef struct QTimerEvent QTimerEvent;
typedef struct QValidator QValidator;
typedef struct QVariant QVariant;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QLineEdit* QLineEdit_new(QWidget* parent);
extern __declspec(dllexport) 
QLineEdit* QLineEdit_new2();
extern __declspec(dllexport) 
QLineEdit* QLineEdit_new3(struct miqt_string param1);
extern __declspec(dllexport) 
QLineEdit* QLineEdit_new4(struct miqt_string param1, QWidget* parent);
extern __declspec(dllexport) 
void QLineEdit_virtbase(QLineEdit* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QLineEdit_MetaObject(const QLineEdit* self);
extern __declspec(dllexport) 
void* QLineEdit_Metacast(QLineEdit* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QLineEdit_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_string QLineEdit_Text(const QLineEdit* self);
extern __declspec(dllexport) 
struct miqt_string QLineEdit_DisplayText(const QLineEdit* self);
extern __declspec(dllexport) 
struct miqt_string QLineEdit_PlaceholderText(const QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_SetPlaceholderText(QLineEdit* self, struct miqt_string placeholderText);
extern __declspec(dllexport) 
int QLineEdit_MaxLength(const QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_SetMaxLength(QLineEdit* self, int maxLength);
extern __declspec(dllexport) 
void QLineEdit_SetFrame(QLineEdit* self, bool frame);
extern __declspec(dllexport) 
bool QLineEdit_HasFrame(const QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_SetClearButtonEnabled(QLineEdit* self, bool enable);
extern __declspec(dllexport) 
bool QLineEdit_IsClearButtonEnabled(const QLineEdit* self);
extern __declspec(dllexport) 
EchoMode QLineEdit_EchoMode(const QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_SetEchoMode(QLineEdit* self, EchoMode echoMode);
extern __declspec(dllexport) 
bool QLineEdit_IsReadOnly(const QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_SetReadOnly(QLineEdit* self, bool readOnly);
extern __declspec(dllexport) 
void QLineEdit_SetValidator(QLineEdit* self, QValidator* validator);
extern __declspec(dllexport) 
QValidator* QLineEdit_Validator(const QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_SetCompleter(QLineEdit* self, QCompleter* completer);
extern __declspec(dllexport) 
QCompleter* QLineEdit_Completer(const QLineEdit* self);
extern __declspec(dllexport) 
QSize* QLineEdit_SizeHint(const QLineEdit* self);
extern __declspec(dllexport) 
QSize* QLineEdit_MinimumSizeHint(const QLineEdit* self);
extern __declspec(dllexport) 
int QLineEdit_CursorPosition(const QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_SetCursorPosition(QLineEdit* self, int cursorPosition);
extern __declspec(dllexport) 
int QLineEdit_CursorPositionAt(QLineEdit* self, QPoint* pos);
extern __declspec(dllexport) 
void QLineEdit_SetAlignment(QLineEdit* self, int flag);
extern __declspec(dllexport) 
int QLineEdit_Alignment(const QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_CursorForward(QLineEdit* self, bool mark);
extern __declspec(dllexport) 
void QLineEdit_CursorBackward(QLineEdit* self, bool mark);
extern __declspec(dllexport) 
void QLineEdit_CursorWordForward(QLineEdit* self, bool mark);
extern __declspec(dllexport) 
void QLineEdit_CursorWordBackward(QLineEdit* self, bool mark);
extern __declspec(dllexport) 
void QLineEdit_Backspace(QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_Del(QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_Home(QLineEdit* self, bool mark);
extern __declspec(dllexport) 
void QLineEdit_End(QLineEdit* self, bool mark);
extern __declspec(dllexport) 
bool QLineEdit_IsModified(const QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_SetModified(QLineEdit* self, bool modified);
extern __declspec(dllexport) 
void QLineEdit_SetSelection(QLineEdit* self, int param1, int param2);
extern __declspec(dllexport) 
bool QLineEdit_HasSelectedText(const QLineEdit* self);
extern __declspec(dllexport) 
struct miqt_string QLineEdit_SelectedText(const QLineEdit* self);
extern __declspec(dllexport) 
int QLineEdit_SelectionStart(const QLineEdit* self);
extern __declspec(dllexport) 
int QLineEdit_SelectionEnd(const QLineEdit* self);
extern __declspec(dllexport) 
int QLineEdit_SelectionLength(const QLineEdit* self);
extern __declspec(dllexport) 
bool QLineEdit_IsUndoAvailable(const QLineEdit* self);
extern __declspec(dllexport) 
bool QLineEdit_IsRedoAvailable(const QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_SetDragEnabled(QLineEdit* self, bool b);
extern __declspec(dllexport) 
bool QLineEdit_DragEnabled(const QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_SetCursorMoveStyle(QLineEdit* self, int style);
extern __declspec(dllexport) 
int QLineEdit_CursorMoveStyle(const QLineEdit* self);
extern __declspec(dllexport) 
struct miqt_string QLineEdit_InputMask(const QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_SetInputMask(QLineEdit* self, struct miqt_string inputMask);
extern __declspec(dllexport) 
bool QLineEdit_HasAcceptableInput(const QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_SetTextMargins(QLineEdit* self, int left, int top, int right, int bottom);
extern __declspec(dllexport) 
void QLineEdit_SetTextMarginsWithMargins(QLineEdit* self, QMargins* margins);
extern __declspec(dllexport) 
QMargins* QLineEdit_TextMargins(const QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_AddAction(QLineEdit* self, QAction* action, ActionPosition position);
extern __declspec(dllexport) 
QAction* QLineEdit_AddAction2(QLineEdit* self, QIcon* icon, ActionPosition position);
extern __declspec(dllexport) 
void QLineEdit_SetText(QLineEdit* self, struct miqt_string text);
extern __declspec(dllexport) 
void QLineEdit_Clear(QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_SelectAll(QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_Undo(QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_Redo(QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_Cut(QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_Copy(const QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_Paste(QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_Deselect(QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_Insert(QLineEdit* self, struct miqt_string param1);
extern __declspec(dllexport) 
QMenu* QLineEdit_CreateStandardContextMenu(QLineEdit* self);
extern __declspec(dllexport) 
void QLineEdit_TextChanged(QLineEdit* self, struct miqt_string param1);
void QLineEdit_connect_TextChanged(QLineEdit* self, intptr_t slot);
extern __declspec(dllexport) 
void QLineEdit_TextEdited(QLineEdit* self, struct miqt_string param1);
void QLineEdit_connect_TextEdited(QLineEdit* self, intptr_t slot);
extern __declspec(dllexport) 
void QLineEdit_CursorPositionChanged(QLineEdit* self, int param1, int param2);
void QLineEdit_connect_CursorPositionChanged(QLineEdit* self, intptr_t slot);
extern __declspec(dllexport) 
void QLineEdit_ReturnPressed(QLineEdit* self);
void QLineEdit_connect_ReturnPressed(QLineEdit* self, intptr_t slot);
extern __declspec(dllexport) 
void QLineEdit_EditingFinished(QLineEdit* self);
void QLineEdit_connect_EditingFinished(QLineEdit* self, intptr_t slot);
extern __declspec(dllexport) 
void QLineEdit_SelectionChanged(QLineEdit* self);
void QLineEdit_connect_SelectionChanged(QLineEdit* self, intptr_t slot);
extern __declspec(dllexport) 
void QLineEdit_InputRejected(QLineEdit* self);
void QLineEdit_connect_InputRejected(QLineEdit* self, intptr_t slot);
extern __declspec(dllexport) 
void QLineEdit_MousePressEvent(QLineEdit* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QLineEdit_MouseMoveEvent(QLineEdit* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QLineEdit_MouseReleaseEvent(QLineEdit* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QLineEdit_MouseDoubleClickEvent(QLineEdit* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QLineEdit_KeyPressEvent(QLineEdit* self, QKeyEvent* param1);
extern __declspec(dllexport) 
void QLineEdit_KeyReleaseEvent(QLineEdit* self, QKeyEvent* param1);
extern __declspec(dllexport) 
void QLineEdit_FocusInEvent(QLineEdit* self, QFocusEvent* param1);
extern __declspec(dllexport) 
void QLineEdit_FocusOutEvent(QLineEdit* self, QFocusEvent* param1);
extern __declspec(dllexport) 
void QLineEdit_PaintEvent(QLineEdit* self, QPaintEvent* param1);
extern __declspec(dllexport) 
void QLineEdit_DragEnterEvent(QLineEdit* self, QDragEnterEvent* param1);
extern __declspec(dllexport) 
void QLineEdit_DragMoveEvent(QLineEdit* self, QDragMoveEvent* e);
extern __declspec(dllexport) 
void QLineEdit_DragLeaveEvent(QLineEdit* self, QDragLeaveEvent* e);
extern __declspec(dllexport) 
void QLineEdit_DropEvent(QLineEdit* self, QDropEvent* param1);
extern __declspec(dllexport) 
void QLineEdit_ChangeEvent(QLineEdit* self, QEvent* param1);
extern __declspec(dllexport) 
void QLineEdit_ContextMenuEvent(QLineEdit* self, QContextMenuEvent* param1);
extern __declspec(dllexport) 
void QLineEdit_InputMethodEvent(QLineEdit* self, QInputMethodEvent* param1);
extern __declspec(dllexport) 
void QLineEdit_InitStyleOption(const QLineEdit* self, QStyleOptionFrame* option);
extern __declspec(dllexport) 
QVariant* QLineEdit_InputMethodQuery(const QLineEdit* self, int param1);
extern __declspec(dllexport) 
QVariant* QLineEdit_InputMethodQuery2(const QLineEdit* self, int property, QVariant* argument);
extern __declspec(dllexport) 
void QLineEdit_TimerEvent(QLineEdit* self, QTimerEvent* param1);
extern __declspec(dllexport) 
bool QLineEdit_Event(QLineEdit* self, QEvent* param1);
extern __declspec(dllexport) 
struct miqt_string QLineEdit_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QLineEdit_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QLineEdit_CursorForward2(QLineEdit* self, bool mark, int steps);
extern __declspec(dllexport) 
void QLineEdit_CursorBackward2(QLineEdit* self, bool mark, int steps);
extern __declspec(dllexport) 
void QLineEdit_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QLineEdit_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QLineEdit_override_virtual_Metacast(void* self, intptr_t slot);
void* QLineEdit_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QLineEdit_Delete(QLineEdit* self, bool isSubclass);

}
