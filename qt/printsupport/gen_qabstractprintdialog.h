#pragma once
#ifndef MIQT_QT_PRINTSUPPORT_GEN_QABSTRACTPRINTDIALOG_H
#define MIQT_QT_PRINTSUPPORT_GEN_QABSTRACTPRINTDIALOG_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractPrintDialog QAbstractPrintDialog;
typedef struct QCloseEvent QCloseEvent;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDialog QDialog;
typedef struct QEvent QEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPrinter QPrinter;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QAbstractPrintDialog* QAbstractPrintDialog_new(QPrinter* printer);
extern __declspec(dllexport) QAbstractPrintDialog* QAbstractPrintDialog_new2(QPrinter* printer, QWidget* parent);
extern __declspec(dllexport) void QAbstractPrintDialog_virtbase(QAbstractPrintDialog* src, QDialog** outptr_QDialog);
extern __declspec(dllexport) QMetaObject* QAbstractPrintDialog_MetaObject(const QAbstractPrintDialog* self);
extern __declspec(dllexport) void* QAbstractPrintDialog_Metacast(QAbstractPrintDialog* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QAbstractPrintDialog_Tr(const char* s);
extern __declspec(dllexport) void QAbstractPrintDialog_SetOptionTabs(QAbstractPrintDialog* self, struct miqt_array /* of QWidget* */  tabs);
extern __declspec(dllexport) void QAbstractPrintDialog_SetPrintRange(QAbstractPrintDialog* self, PrintRange rangeVal);
extern __declspec(dllexport) PrintRange QAbstractPrintDialog_PrintRange(const QAbstractPrintDialog* self);
extern __declspec(dllexport) void QAbstractPrintDialog_SetMinMax(QAbstractPrintDialog* self, int min, int max);
extern __declspec(dllexport) int QAbstractPrintDialog_MinPage(const QAbstractPrintDialog* self);
extern __declspec(dllexport) int QAbstractPrintDialog_MaxPage(const QAbstractPrintDialog* self);
extern __declspec(dllexport) void QAbstractPrintDialog_SetFromTo(QAbstractPrintDialog* self, int fromPage, int toPage);
extern __declspec(dllexport) int QAbstractPrintDialog_FromPage(const QAbstractPrintDialog* self);
extern __declspec(dllexport) int QAbstractPrintDialog_ToPage(const QAbstractPrintDialog* self);
extern __declspec(dllexport) QPrinter* QAbstractPrintDialog_Printer(const QAbstractPrintDialog* self);
extern __declspec(dllexport) struct miqt_string QAbstractPrintDialog_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QAbstractPrintDialog_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QAbstractPrintDialog_override_virtual_SetVisible(void* self, intptr_t slot);
void QAbstractPrintDialog_virtualbase_SetVisible(void* self, bool visible);
extern __declspec(dllexport) void QAbstractPrintDialog_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QAbstractPrintDialog_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QAbstractPrintDialog_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QAbstractPrintDialog_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QAbstractPrintDialog_override_virtual_Open(void* self, intptr_t slot);
void QAbstractPrintDialog_virtualbase_Open(void* self);
extern __declspec(dllexport) void QAbstractPrintDialog_override_virtual_Exec(void* self, intptr_t slot);
int QAbstractPrintDialog_virtualbase_Exec(void* self);
extern __declspec(dllexport) void QAbstractPrintDialog_override_virtual_Done(void* self, intptr_t slot);
void QAbstractPrintDialog_virtualbase_Done(void* self, int param1);
extern __declspec(dllexport) void QAbstractPrintDialog_override_virtual_Accept(void* self, intptr_t slot);
void QAbstractPrintDialog_virtualbase_Accept(void* self);
extern __declspec(dllexport) void QAbstractPrintDialog_override_virtual_Reject(void* self, intptr_t slot);
void QAbstractPrintDialog_virtualbase_Reject(void* self);
extern __declspec(dllexport) void QAbstractPrintDialog_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QAbstractPrintDialog_virtualbase_KeyPressEvent(void* self, QKeyEvent* param1);
extern __declspec(dllexport) void QAbstractPrintDialog_override_virtual_CloseEvent(void* self, intptr_t slot);
void QAbstractPrintDialog_virtualbase_CloseEvent(void* self, QCloseEvent* param1);
extern __declspec(dllexport) void QAbstractPrintDialog_override_virtual_ShowEvent(void* self, intptr_t slot);
void QAbstractPrintDialog_virtualbase_ShowEvent(void* self, QShowEvent* param1);
extern __declspec(dllexport) void QAbstractPrintDialog_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QAbstractPrintDialog_virtualbase_ResizeEvent(void* self, QResizeEvent* param1);
extern __declspec(dllexport) void QAbstractPrintDialog_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QAbstractPrintDialog_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* param1);
extern __declspec(dllexport) void QAbstractPrintDialog_override_virtual_EventFilter(void* self, intptr_t slot);
bool QAbstractPrintDialog_virtualbase_EventFilter(void* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) void QAbstractPrintDialog_Delete(QAbstractPrintDialog* self, bool isSubclass);

} 
