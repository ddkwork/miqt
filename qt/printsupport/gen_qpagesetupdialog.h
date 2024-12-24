#pragma once
#ifndef MIQT_QT_PRINTSUPPORT_GEN_QPAGESETUPDIALOG_H
#define MIQT_QT_PRINTSUPPORT_GEN_QPAGESETUPDIALOG_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QCloseEvent;
class QContextMenuEvent;
class QDialog;
class QEvent;
class QKeyEvent;
class QMetaObject;
class QObject;
class QPageSetupDialog;
class QPaintDevice;
class QPrinter;
class QResizeEvent;
class QShowEvent;
class QSize;
class QWidget;
class _GUID;
class type_info;
#else
typedef struct QCloseEvent QCloseEvent;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDialog QDialog;
typedef struct QEvent QEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPageSetupDialog QPageSetupDialog;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPrinter QPrinter;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QPageSetupDialog* QPageSetupDialog_new(QWidget* parent);
extern __declspec(dllexport) QPageSetupDialog* QPageSetupDialog_new2(QPrinter* printer);
extern __declspec(dllexport) QPageSetupDialog* QPageSetupDialog_new3();
extern __declspec(dllexport) QPageSetupDialog* QPageSetupDialog_new4(QPrinter* printer, QWidget* parent);
extern __declspec(dllexport) void QPageSetupDialog_virtbase(QPageSetupDialog* src, QDialog** outptr_QDialog);
extern __declspec(dllexport) QMetaObject* QPageSetupDialog_MetaObject(const QPageSetupDialog* self);
extern __declspec(dllexport) void* QPageSetupDialog_Metacast(QPageSetupDialog* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QPageSetupDialog_Tr(const char* s);
extern __declspec(dllexport) void QPageSetupDialog_SetVisible(QPageSetupDialog* self, bool visible);
extern __declspec(dllexport) int QPageSetupDialog_Exec(QPageSetupDialog* self);
extern __declspec(dllexport) void QPageSetupDialog_Done(QPageSetupDialog* self, int result);
extern __declspec(dllexport) QPrinter* QPageSetupDialog_Printer(QPageSetupDialog* self);
extern __declspec(dllexport) struct miqt_string QPageSetupDialog_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QPageSetupDialog_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QPageSetupDialog_override_virtual_SetVisible(void* self, intptr_t slot);
void QPageSetupDialog_virtualbase_SetVisible(void* self, bool visible);
extern __declspec(dllexport) void QPageSetupDialog_override_virtual_Exec(void* self, intptr_t slot);
int QPageSetupDialog_virtualbase_Exec(void* self);
extern __declspec(dllexport) void QPageSetupDialog_override_virtual_Done(void* self, intptr_t slot);
void QPageSetupDialog_virtualbase_Done(void* self, int result);
extern __declspec(dllexport) void QPageSetupDialog_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QPageSetupDialog_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QPageSetupDialog_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QPageSetupDialog_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QPageSetupDialog_override_virtual_Open(void* self, intptr_t slot);
void QPageSetupDialog_virtualbase_Open(void* self);
extern __declspec(dllexport) void QPageSetupDialog_override_virtual_Accept(void* self, intptr_t slot);
void QPageSetupDialog_virtualbase_Accept(void* self);
extern __declspec(dllexport) void QPageSetupDialog_override_virtual_Reject(void* self, intptr_t slot);
void QPageSetupDialog_virtualbase_Reject(void* self);
extern __declspec(dllexport) void QPageSetupDialog_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QPageSetupDialog_virtualbase_KeyPressEvent(void* self, QKeyEvent* param1);
extern __declspec(dllexport) void QPageSetupDialog_override_virtual_CloseEvent(void* self, intptr_t slot);
void QPageSetupDialog_virtualbase_CloseEvent(void* self, QCloseEvent* param1);
extern __declspec(dllexport) void QPageSetupDialog_override_virtual_ShowEvent(void* self, intptr_t slot);
void QPageSetupDialog_virtualbase_ShowEvent(void* self, QShowEvent* param1);
extern __declspec(dllexport) void QPageSetupDialog_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QPageSetupDialog_virtualbase_ResizeEvent(void* self, QResizeEvent* param1);
extern __declspec(dllexport) void QPageSetupDialog_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QPageSetupDialog_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* param1);
extern __declspec(dllexport) void QPageSetupDialog_override_virtual_EventFilter(void* self, intptr_t slot);
bool QPageSetupDialog_virtualbase_EventFilter(void* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) void QPageSetupDialog_Delete(QPageSetupDialog* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
