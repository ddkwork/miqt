#pragma once
#ifndef MIQT_QT_PRINTSUPPORT_GEN_QPRINTPREVIEWDIALOG_H
#define MIQT_QT_PRINTSUPPORT_GEN_QPRINTPREVIEWDIALOG_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDialog QDialog;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPrintPreviewDialog QPrintPreviewDialog;
typedef struct QPrinter QPrinter;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QPrintPreviewDialog* QPrintPreviewDialog_new(QWidget* parent);
extern __declspec(dllexport) 
QPrintPreviewDialog* QPrintPreviewDialog_new2();
extern __declspec(dllexport) 
QPrintPreviewDialog* QPrintPreviewDialog_new3(QPrinter* printer);
extern __declspec(dllexport) 
QPrintPreviewDialog* QPrintPreviewDialog_new4(QWidget* parent, int flags);
extern __declspec(dllexport) 
QPrintPreviewDialog* QPrintPreviewDialog_new5(QPrinter* printer, QWidget* parent);
extern __declspec(dllexport) 
QPrintPreviewDialog* QPrintPreviewDialog_new6(QPrinter* printer, QWidget* parent, int flags);
extern __declspec(dllexport) 
void QPrintPreviewDialog_virtbase(QPrintPreviewDialog* src
, QDialog** outptr_QDialog
);
extern __declspec(dllexport) 
QMetaObject* QPrintPreviewDialog_MetaObject(const QPrintPreviewDialog* self);
extern __declspec(dllexport) 
void* QPrintPreviewDialog_Metacast(QPrintPreviewDialog* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QPrintPreviewDialog_Tr(const char* s);
extern __declspec(dllexport) 
QPrinter* QPrintPreviewDialog_Printer(QPrintPreviewDialog* self);
extern __declspec(dllexport) 
void QPrintPreviewDialog_SetVisible(QPrintPreviewDialog* self, bool visible);
extern __declspec(dllexport) 
void QPrintPreviewDialog_Done(QPrintPreviewDialog* self, int result);
extern __declspec(dllexport) 
void QPrintPreviewDialog_PaintRequested(QPrintPreviewDialog* self, QPrinter* printer);
void QPrintPreviewDialog_connect_PaintRequested(QPrintPreviewDialog* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QPrintPreviewDialog_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QPrintPreviewDialog_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QPrintPreviewDialog_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QPrintPreviewDialog_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QPrintPreviewDialog_override_virtual_Metacast(void* self, intptr_t slot);
void* QPrintPreviewDialog_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QPrintPreviewDialog_Delete(QPrintPreviewDialog* self, bool isSubclass);

}
