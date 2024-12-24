#pragma once
#ifndef MIQT_QT_GEN_QDIALOG_H
#define MIQT_QT_GEN_QDIALOG_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QCloseEvent QCloseEvent;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDialog QDialog;
typedef struct QEvent QEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QDialog* QDialog_new(QWidget* parent);
extern __declspec(dllexport) 
QDialog* QDialog_new2();
extern __declspec(dllexport) 
QDialog* QDialog_new3(QWidget* parent, int f);
extern __declspec(dllexport) 
void QDialog_virtbase(QDialog* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QDialog_MetaObject(const QDialog* self);
extern __declspec(dllexport) 
void* QDialog_Metacast(QDialog* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QDialog_Tr(const char* s);
extern __declspec(dllexport) 
int QDialog_Result(const QDialog* self);
extern __declspec(dllexport) 
void QDialog_SetVisible(QDialog* self, bool visible);
extern __declspec(dllexport) 
QSize* QDialog_SizeHint(const QDialog* self);
extern __declspec(dllexport) 
QSize* QDialog_MinimumSizeHint(const QDialog* self);
extern __declspec(dllexport) 
void QDialog_SetSizeGripEnabled(QDialog* self, bool sizeGripEnabled);
extern __declspec(dllexport) 
bool QDialog_IsSizeGripEnabled(const QDialog* self);
extern __declspec(dllexport) 
void QDialog_SetModal(QDialog* self, bool modal);
extern __declspec(dllexport) 
void QDialog_SetResult(QDialog* self, int r);
extern __declspec(dllexport) 
void QDialog_Finished(QDialog* self, int result);
void QDialog_connect_Finished(QDialog* self, intptr_t slot);
extern __declspec(dllexport) 
void QDialog_Accepted(QDialog* self);
void QDialog_connect_Accepted(QDialog* self, intptr_t slot);
extern __declspec(dllexport) 
void QDialog_Rejected(QDialog* self);
void QDialog_connect_Rejected(QDialog* self, intptr_t slot);
extern __declspec(dllexport) 
void QDialog_Open(QDialog* self);
extern __declspec(dllexport) 
int QDialog_Exec(QDialog* self);
extern __declspec(dllexport) 
void QDialog_Done(QDialog* self, int param1);
extern __declspec(dllexport) 
void QDialog_Accept(QDialog* self);
extern __declspec(dllexport) 
void QDialog_Reject(QDialog* self);
extern __declspec(dllexport) 
void QDialog_KeyPressEvent(QDialog* self, QKeyEvent* param1);
extern __declspec(dllexport) 
void QDialog_CloseEvent(QDialog* self, QCloseEvent* param1);
extern __declspec(dllexport) 
void QDialog_ShowEvent(QDialog* self, QShowEvent* param1);
extern __declspec(dllexport) 
void QDialog_ResizeEvent(QDialog* self, QResizeEvent* param1);
extern __declspec(dllexport) 
void QDialog_ContextMenuEvent(QDialog* self, QContextMenuEvent* param1);
extern __declspec(dllexport) 
bool QDialog_EventFilter(QDialog* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) 
struct miqt_string QDialog_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QDialog_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QDialog_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QDialog_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QDialog_override_virtual_Metacast(void* self, intptr_t slot);
void* QDialog_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QDialog_Delete(QDialog* self, bool isSubclass);

}
