#pragma once
#ifndef MIQT_QT_PRINTSUPPORT_GEN_QPRINTPREVIEWDIALOG_H
#define MIQT_QT_PRINTSUPPORT_GEN_QPRINTPREVIEWDIALOG_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QCloseEvent QCloseEvent;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDialog QDialog;
typedef struct QEvent QEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPrintPreviewDialog QPrintPreviewDialog;
typedef struct QPrinter QPrinter;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QWidget QWidget;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QPrintPreviewDialog* QPrintPreviewDialog_new(QWidget* parent);
extern __declspec(dllexport) QPrintPreviewDialog* QPrintPreviewDialog_new2();
extern __declspec(dllexport) QPrintPreviewDialog* QPrintPreviewDialog_new3(QPrinter* printer);
extern __declspec(dllexport) QPrintPreviewDialog* QPrintPreviewDialog_new4(QWidget* parent, int flags);
extern __declspec(dllexport) QPrintPreviewDialog* QPrintPreviewDialog_new5(QPrinter* printer, QWidget* parent);
extern __declspec(dllexport) QPrintPreviewDialog* QPrintPreviewDialog_new6(QPrinter* printer, QWidget* parent, int flags);
extern __declspec(dllexport) void QPrintPreviewDialog_virtbase(QPrintPreviewDialog* src, QDialog** outptr_QDialog);
extern __declspec(dllexport) QMetaObject* QPrintPreviewDialog_MetaObject(const QPrintPreviewDialog* self);
extern __declspec(dllexport) void* QPrintPreviewDialog_Metacast(QPrintPreviewDialog* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QPrintPreviewDialog_Tr(const char* s);
extern __declspec(dllexport) QPrinter* QPrintPreviewDialog_Printer(QPrintPreviewDialog* self);
extern __declspec(dllexport) void QPrintPreviewDialog_SetVisible(QPrintPreviewDialog* self, bool visible);
extern __declspec(dllexport) void QPrintPreviewDialog_Done(QPrintPreviewDialog* self, int result);
extern __declspec(dllexport) void QPrintPreviewDialog_PaintRequested(QPrintPreviewDialog* self, QPrinter* printer);
void QPrintPreviewDialog_connect_PaintRequested(QPrintPreviewDialog* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QPrintPreviewDialog_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QPrintPreviewDialog_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QPrintPreviewDialog_override_virtual_SetVisible(void* self, intptr_t slot);
void QPrintPreviewDialog_virtualbase_SetVisible(void* self, bool visible);
extern __declspec(dllexport) void QPrintPreviewDialog_override_virtual_Done(void* self, intptr_t slot);
void QPrintPreviewDialog_virtualbase_Done(void* self, int result);
extern __declspec(dllexport) void QPrintPreviewDialog_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QPrintPreviewDialog_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QPrintPreviewDialog_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QPrintPreviewDialog_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QPrintPreviewDialog_override_virtual_Open(void* self, intptr_t slot);
void QPrintPreviewDialog_virtualbase_Open(void* self);
extern __declspec(dllexport) void QPrintPreviewDialog_override_virtual_Exec(void* self, intptr_t slot);
int QPrintPreviewDialog_virtualbase_Exec(void* self);
extern __declspec(dllexport) void QPrintPreviewDialog_override_virtual_Accept(void* self, intptr_t slot);
void QPrintPreviewDialog_virtualbase_Accept(void* self);
extern __declspec(dllexport) void QPrintPreviewDialog_override_virtual_Reject(void* self, intptr_t slot);
void QPrintPreviewDialog_virtualbase_Reject(void* self);
extern __declspec(dllexport) void QPrintPreviewDialog_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QPrintPreviewDialog_virtualbase_KeyPressEvent(void* self, QKeyEvent* param1);
extern __declspec(dllexport) void QPrintPreviewDialog_override_virtual_CloseEvent(void* self, intptr_t slot);
void QPrintPreviewDialog_virtualbase_CloseEvent(void* self, QCloseEvent* param1);
extern __declspec(dllexport) void QPrintPreviewDialog_override_virtual_ShowEvent(void* self, intptr_t slot);
void QPrintPreviewDialog_virtualbase_ShowEvent(void* self, QShowEvent* param1);
extern __declspec(dllexport) void QPrintPreviewDialog_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QPrintPreviewDialog_virtualbase_ResizeEvent(void* self, QResizeEvent* param1);
extern __declspec(dllexport) void QPrintPreviewDialog_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QPrintPreviewDialog_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* param1);
extern __declspec(dllexport) void QPrintPreviewDialog_override_virtual_EventFilter(void* self, intptr_t slot);
bool QPrintPreviewDialog_virtualbase_EventFilter(void* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) void QPrintPreviewDialog_Delete(QPrintPreviewDialog* self, bool isSubclass);

} 
