#pragma once
#ifndef MIQT_QT_PRINTSUPPORT_GEN_QPRINTDIALOG_H
#define MIQT_QT_PRINTSUPPORT_GEN_QPRINTDIALOG_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractPrintDialog QAbstractPrintDialog;
typedef struct QDialog QDialog;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPrintDialog QPrintDialog;
typedef struct QPrinter QPrinter;
typedef struct QWidget QWidget;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QPrintDialog* QPrintDialog_new(QWidget* parent);
extern __declspec(dllexport) QPrintDialog* QPrintDialog_new2(QPrinter* printer);
extern __declspec(dllexport) QPrintDialog* QPrintDialog_new3();
extern __declspec(dllexport) QPrintDialog* QPrintDialog_new4(QPrinter* printer, QWidget* parent);
extern __declspec(dllexport) void QPrintDialog_virtbase(QPrintDialog* src, QAbstractPrintDialog** outptr_QAbstractPrintDialog);
extern __declspec(dllexport) QMetaObject* QPrintDialog_MetaObject(const QPrintDialog* self);
extern __declspec(dllexport) void* QPrintDialog_Metacast(QPrintDialog* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QPrintDialog_Tr(const char* s);
extern __declspec(dllexport) int QPrintDialog_Exec(QPrintDialog* self);
extern __declspec(dllexport) void QPrintDialog_Done(QPrintDialog* self, int result);
extern __declspec(dllexport) void QPrintDialog_SetOption(QPrintDialog* self, PrintDialogOption option);
extern __declspec(dllexport) bool QPrintDialog_TestOption(const QPrintDialog* self, PrintDialogOption option);
extern __declspec(dllexport) void QPrintDialog_SetOptions(QPrintDialog* self, PrintDialogOptions options);
extern __declspec(dllexport) PrintDialogOptions QPrintDialog_Options(const QPrintDialog* self);
extern __declspec(dllexport) void QPrintDialog_SetVisible(QPrintDialog* self, bool visible);
extern __declspec(dllexport) void QPrintDialog_Accepted(QPrintDialog* self, QPrinter* printer);
void QPrintDialog_connect_Accepted(QPrintDialog* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QPrintDialog_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QPrintDialog_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QPrintDialog_SetOption2(QPrintDialog* self, PrintDialogOption option, bool on);
extern __declspec(dllexport) void QPrintDialog_override_virtual_Exec(void* self, intptr_t slot);
int QPrintDialog_virtualbase_Exec(void* self);
extern __declspec(dllexport) void QPrintDialog_override_virtual_Done(void* self, intptr_t slot);
void QPrintDialog_virtualbase_Done(void* self, int result);
extern __declspec(dllexport) void QPrintDialog_override_virtual_SetVisible(void* self, intptr_t slot);
void QPrintDialog_virtualbase_SetVisible(void* self, bool visible);
extern __declspec(dllexport) void QPrintDialog_Delete(QPrintDialog* self, bool isSubclass);

} 
