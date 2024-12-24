#pragma once
#ifndef MIQT_QT_GEN_QTEXTBROWSER_H
#define MIQT_QT_GEN_QTEXTBROWSER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractScrollArea QAbstractScrollArea;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QFrame QFrame;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QTextBrowser QTextBrowser;
typedef struct QTextEdit QTextEdit;
typedef struct QUrl QUrl;
typedef struct QVariant QVariant;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QTextBrowser* QTextBrowser_new(QWidget* parent);
extern __declspec(dllexport) 
QTextBrowser* QTextBrowser_new2();
extern __declspec(dllexport) 
void QTextBrowser_virtbase(QTextBrowser* src
, QTextEdit** outptr_QTextEdit
);
extern __declspec(dllexport) 
QMetaObject* QTextBrowser_MetaObject(const QTextBrowser* self);
extern __declspec(dllexport) 
void* QTextBrowser_Metacast(QTextBrowser* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QTextBrowser_Tr(const char* s);
extern __declspec(dllexport) 
QUrl* QTextBrowser_Source(const QTextBrowser* self);
extern __declspec(dllexport) 
int QTextBrowser_SourceType(const QTextBrowser* self);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QTextBrowser_SearchPaths(const QTextBrowser* self);
extern __declspec(dllexport) 
void QTextBrowser_SetSearchPaths(QTextBrowser* self, struct miqt_array /* of struct miqt_string */  paths);
extern __declspec(dllexport) 
QVariant* QTextBrowser_LoadResource(QTextBrowser* self, int typeVal, QUrl* name);
extern __declspec(dllexport) 
bool QTextBrowser_IsBackwardAvailable(const QTextBrowser* self);
extern __declspec(dllexport) 
bool QTextBrowser_IsForwardAvailable(const QTextBrowser* self);
extern __declspec(dllexport) 
void QTextBrowser_ClearHistory(QTextBrowser* self);
extern __declspec(dllexport) 
struct miqt_string QTextBrowser_HistoryTitle(const QTextBrowser* self, int param1);
extern __declspec(dllexport) 
QUrl* QTextBrowser_HistoryUrl(const QTextBrowser* self, int param1);
extern __declspec(dllexport) 
int QTextBrowser_BackwardHistoryCount(const QTextBrowser* self);
extern __declspec(dllexport) 
int QTextBrowser_ForwardHistoryCount(const QTextBrowser* self);
extern __declspec(dllexport) 
bool QTextBrowser_OpenExternalLinks(const QTextBrowser* self);
extern __declspec(dllexport) 
void QTextBrowser_SetOpenExternalLinks(QTextBrowser* self, bool open);
extern __declspec(dllexport) 
bool QTextBrowser_OpenLinks(const QTextBrowser* self);
extern __declspec(dllexport) 
void QTextBrowser_SetOpenLinks(QTextBrowser* self, bool open);
extern __declspec(dllexport) 
void QTextBrowser_SetSource(QTextBrowser* self, QUrl* name);
extern __declspec(dllexport) 
void QTextBrowser_Backward(QTextBrowser* self);
extern __declspec(dllexport) 
void QTextBrowser_Forward(QTextBrowser* self);
extern __declspec(dllexport) 
void QTextBrowser_Home(QTextBrowser* self);
extern __declspec(dllexport) 
void QTextBrowser_Reload(QTextBrowser* self);
extern __declspec(dllexport) 
void QTextBrowser_BackwardAvailable(QTextBrowser* self, bool param1);
void QTextBrowser_connect_BackwardAvailable(QTextBrowser* self, intptr_t slot);
extern __declspec(dllexport) 
void QTextBrowser_ForwardAvailable(QTextBrowser* self, bool param1);
void QTextBrowser_connect_ForwardAvailable(QTextBrowser* self, intptr_t slot);
extern __declspec(dllexport) 
void QTextBrowser_HistoryChanged(QTextBrowser* self);
void QTextBrowser_connect_HistoryChanged(QTextBrowser* self, intptr_t slot);
extern __declspec(dllexport) 
void QTextBrowser_SourceChanged(QTextBrowser* self, QUrl* param1);
void QTextBrowser_connect_SourceChanged(QTextBrowser* self, intptr_t slot);
extern __declspec(dllexport) 
void QTextBrowser_Highlighted(QTextBrowser* self, QUrl* param1);
void QTextBrowser_connect_Highlighted(QTextBrowser* self, intptr_t slot);
extern __declspec(dllexport) 
void QTextBrowser_AnchorClicked(QTextBrowser* self, QUrl* param1);
void QTextBrowser_connect_AnchorClicked(QTextBrowser* self, intptr_t slot);
extern __declspec(dllexport) 
bool QTextBrowser_Event(QTextBrowser* self, QEvent* e);
extern __declspec(dllexport) 
void QTextBrowser_KeyPressEvent(QTextBrowser* self, QKeyEvent* ev);
extern __declspec(dllexport) 
void QTextBrowser_MouseMoveEvent(QTextBrowser* self, QMouseEvent* ev);
extern __declspec(dllexport) 
void QTextBrowser_MousePressEvent(QTextBrowser* self, QMouseEvent* ev);
extern __declspec(dllexport) 
void QTextBrowser_MouseReleaseEvent(QTextBrowser* self, QMouseEvent* ev);
extern __declspec(dllexport) 
void QTextBrowser_FocusOutEvent(QTextBrowser* self, QFocusEvent* ev);
extern __declspec(dllexport) 
bool QTextBrowser_FocusNextPrevChild(QTextBrowser* self, bool next);
extern __declspec(dllexport) 
void QTextBrowser_PaintEvent(QTextBrowser* self, QPaintEvent* e);
extern __declspec(dllexport) 
void QTextBrowser_DoSetSource(QTextBrowser* self, QUrl* name, int typeVal);
extern __declspec(dllexport) 
struct miqt_string QTextBrowser_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QTextBrowser_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QTextBrowser_SetSource2(QTextBrowser* self, QUrl* name, int typeVal);
extern __declspec(dllexport) 
void QTextBrowser_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QTextBrowser_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QTextBrowser_override_virtual_Metacast(void* self, intptr_t slot);
void* QTextBrowser_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QTextBrowser_Delete(QTextBrowser* self, bool isSubclass);

}
