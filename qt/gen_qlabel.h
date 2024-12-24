#pragma once
#ifndef MIQT_QT_GEN_QLABEL_H
#define MIQT_QT_GEN_QLABEL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QFrame QFrame;
typedef struct QKeyEvent QKeyEvent;
typedef struct QLabel QLabel;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QMovie QMovie;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPicture QPicture;
typedef struct QPixmap QPixmap;
typedef struct QSize QSize;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QLabel* QLabel_new(QWidget* parent);
extern __declspec(dllexport) 
QLabel* QLabel_new2();
extern __declspec(dllexport) 
QLabel* QLabel_new3(struct miqt_string text);
extern __declspec(dllexport) 
QLabel* QLabel_new4(QWidget* parent, int f);
extern __declspec(dllexport) 
QLabel* QLabel_new5(struct miqt_string text, QWidget* parent);
extern __declspec(dllexport) 
QLabel* QLabel_new6(struct miqt_string text, QWidget* parent, int f);
extern __declspec(dllexport) 
void QLabel_virtbase(QLabel* src
, QFrame** outptr_QFrame
);
extern __declspec(dllexport) 
QMetaObject* QLabel_MetaObject(const QLabel* self);
extern __declspec(dllexport) 
void* QLabel_Metacast(QLabel* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QLabel_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_string QLabel_Text(const QLabel* self);
extern __declspec(dllexport) 
QPixmap* QLabel_Pixmap(const QLabel* self, int param1);
extern __declspec(dllexport) 
QPixmap* QLabel_Pixmap2(const QLabel* self);
extern __declspec(dllexport) 
QPicture* QLabel_Picture(const QLabel* self, int param1);
extern __declspec(dllexport) 
QPicture* QLabel_Picture2(const QLabel* self);
extern __declspec(dllexport) 
QMovie* QLabel_Movie(const QLabel* self);
extern __declspec(dllexport) 
int QLabel_TextFormat(const QLabel* self);
extern __declspec(dllexport) 
void QLabel_SetTextFormat(QLabel* self, int textFormat);
extern __declspec(dllexport) 
int QLabel_Alignment(const QLabel* self);
extern __declspec(dllexport) 
void QLabel_SetAlignment(QLabel* self, int alignment);
extern __declspec(dllexport) 
void QLabel_SetWordWrap(QLabel* self, bool on);
extern __declspec(dllexport) 
bool QLabel_WordWrap(const QLabel* self);
extern __declspec(dllexport) 
int QLabel_Indent(const QLabel* self);
extern __declspec(dllexport) 
void QLabel_SetIndent(QLabel* self, int indent);
extern __declspec(dllexport) 
int QLabel_Margin(const QLabel* self);
extern __declspec(dllexport) 
void QLabel_SetMargin(QLabel* self, int margin);
extern __declspec(dllexport) 
bool QLabel_HasScaledContents(const QLabel* self);
extern __declspec(dllexport) 
void QLabel_SetScaledContents(QLabel* self, bool scaledContents);
extern __declspec(dllexport) 
QSize* QLabel_SizeHint(const QLabel* self);
extern __declspec(dllexport) 
QSize* QLabel_MinimumSizeHint(const QLabel* self);
extern __declspec(dllexport) 
void QLabel_SetBuddy(QLabel* self, QWidget* buddy);
extern __declspec(dllexport) 
QWidget* QLabel_Buddy(const QLabel* self);
extern __declspec(dllexport) 
int QLabel_HeightForWidth(const QLabel* self, int param1);
extern __declspec(dllexport) 
bool QLabel_OpenExternalLinks(const QLabel* self);
extern __declspec(dllexport) 
void QLabel_SetOpenExternalLinks(QLabel* self, bool open);
extern __declspec(dllexport) 
void QLabel_SetTextInteractionFlags(QLabel* self, int flags);
extern __declspec(dllexport) 
int QLabel_TextInteractionFlags(const QLabel* self);
extern __declspec(dllexport) 
void QLabel_SetSelection(QLabel* self, int param1, int param2);
extern __declspec(dllexport) 
bool QLabel_HasSelectedText(const QLabel* self);
extern __declspec(dllexport) 
struct miqt_string QLabel_SelectedText(const QLabel* self);
extern __declspec(dllexport) 
int QLabel_SelectionStart(const QLabel* self);
extern __declspec(dllexport) 
void QLabel_SetText(QLabel* self, struct miqt_string text);
extern __declspec(dllexport) 
void QLabel_SetPixmap(QLabel* self, QPixmap* pixmap);
extern __declspec(dllexport) 
void QLabel_SetPicture(QLabel* self, QPicture* picture);
extern __declspec(dllexport) 
void QLabel_SetMovie(QLabel* self, QMovie* movie);
extern __declspec(dllexport) 
void QLabel_SetNum(QLabel* self, int num);
extern __declspec(dllexport) 
void QLabel_SetNumWithNum(QLabel* self, double num);
extern __declspec(dllexport) 
void QLabel_Clear(QLabel* self);
extern __declspec(dllexport) 
void QLabel_LinkActivated(QLabel* self, struct miqt_string link);
void QLabel_connect_LinkActivated(QLabel* self, intptr_t slot);
extern __declspec(dllexport) 
void QLabel_LinkHovered(QLabel* self, struct miqt_string link);
void QLabel_connect_LinkHovered(QLabel* self, intptr_t slot);
extern __declspec(dllexport) 
bool QLabel_Event(QLabel* self, QEvent* e);
extern __declspec(dllexport) 
void QLabel_KeyPressEvent(QLabel* self, QKeyEvent* ev);
extern __declspec(dllexport) 
void QLabel_PaintEvent(QLabel* self, QPaintEvent* param1);
extern __declspec(dllexport) 
void QLabel_ChangeEvent(QLabel* self, QEvent* param1);
extern __declspec(dllexport) 
void QLabel_MousePressEvent(QLabel* self, QMouseEvent* ev);
extern __declspec(dllexport) 
void QLabel_MouseMoveEvent(QLabel* self, QMouseEvent* ev);
extern __declspec(dllexport) 
void QLabel_MouseReleaseEvent(QLabel* self, QMouseEvent* ev);
extern __declspec(dllexport) 
void QLabel_ContextMenuEvent(QLabel* self, QContextMenuEvent* ev);
extern __declspec(dllexport) 
void QLabel_FocusInEvent(QLabel* self, QFocusEvent* ev);
extern __declspec(dllexport) 
void QLabel_FocusOutEvent(QLabel* self, QFocusEvent* ev);
extern __declspec(dllexport) 
bool QLabel_FocusNextPrevChild(QLabel* self, bool next);
extern __declspec(dllexport) 
struct miqt_string QLabel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QLabel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QLabel_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QLabel_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QLabel_override_virtual_Metacast(void* self, intptr_t slot);
void* QLabel_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QLabel_Delete(QLabel* self, bool isSubclass);

}
