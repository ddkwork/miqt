#pragma once
#ifndef MIQT_QT_GEN_QERRORMESSAGE_H
#define MIQT_QT_GEN_QERRORMESSAGE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QCloseEvent QCloseEvent;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDialog QDialog;
typedef struct QErrorMessage QErrorMessage;
typedef struct QEvent QEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QWidget QWidget;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QErrorMessage* QErrorMessage_new(QWidget* parent);
extern __declspec(dllexport) QErrorMessage* QErrorMessage_new2();
extern __declspec(dllexport) void QErrorMessage_virtbase(QErrorMessage* src, QDialog** outptr_QDialog);
extern __declspec(dllexport) QMetaObject* QErrorMessage_MetaObject(const QErrorMessage* self);
extern __declspec(dllexport) void* QErrorMessage_Metacast(QErrorMessage* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QErrorMessage_Tr(const char* s);
extern __declspec(dllexport) QErrorMessage* QErrorMessage_QtHandler();
extern __declspec(dllexport) void QErrorMessage_ShowMessage(QErrorMessage* self, struct miqt_string message);
extern __declspec(dllexport) void QErrorMessage_ShowMessage2(QErrorMessage* self, struct miqt_string message, struct miqt_string typeVal);
extern __declspec(dllexport) void QErrorMessage_Done(QErrorMessage* self, int param1);
extern __declspec(dllexport) void QErrorMessage_ChangeEvent(QErrorMessage* self, QEvent* e);
extern __declspec(dllexport) struct miqt_string QErrorMessage_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QErrorMessage_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QErrorMessage_override_virtual_Done(void* self, intptr_t slot);
void QErrorMessage_virtualbase_Done(void* self, int param1);
extern __declspec(dllexport) void QErrorMessage_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QErrorMessage_virtualbase_ChangeEvent(void* self, QEvent* e);
extern __declspec(dllexport) void QErrorMessage_override_virtual_SetVisible(void* self, intptr_t slot);
void QErrorMessage_virtualbase_SetVisible(void* self, bool visible);
extern __declspec(dllexport) void QErrorMessage_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QErrorMessage_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QErrorMessage_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QErrorMessage_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QErrorMessage_override_virtual_Open(void* self, intptr_t slot);
void QErrorMessage_virtualbase_Open(void* self);
extern __declspec(dllexport) void QErrorMessage_override_virtual_Exec(void* self, intptr_t slot);
int QErrorMessage_virtualbase_Exec(void* self);
extern __declspec(dllexport) void QErrorMessage_override_virtual_Accept(void* self, intptr_t slot);
void QErrorMessage_virtualbase_Accept(void* self);
extern __declspec(dllexport) void QErrorMessage_override_virtual_Reject(void* self, intptr_t slot);
void QErrorMessage_virtualbase_Reject(void* self);
extern __declspec(dllexport) void QErrorMessage_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QErrorMessage_virtualbase_KeyPressEvent(void* self, QKeyEvent* param1);
extern __declspec(dllexport) void QErrorMessage_override_virtual_CloseEvent(void* self, intptr_t slot);
void QErrorMessage_virtualbase_CloseEvent(void* self, QCloseEvent* param1);
extern __declspec(dllexport) void QErrorMessage_override_virtual_ShowEvent(void* self, intptr_t slot);
void QErrorMessage_virtualbase_ShowEvent(void* self, QShowEvent* param1);
extern __declspec(dllexport) void QErrorMessage_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QErrorMessage_virtualbase_ResizeEvent(void* self, QResizeEvent* param1);
extern __declspec(dllexport) void QErrorMessage_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QErrorMessage_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* param1);
extern __declspec(dllexport) void QErrorMessage_override_virtual_EventFilter(void* self, intptr_t slot);
bool QErrorMessage_virtualbase_EventFilter(void* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) void QErrorMessage_Delete(QErrorMessage* self, bool isSubclass);

} 
