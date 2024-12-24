#pragma once
#ifndef MIQT_QT_PRINTSUPPORT_GEN_QPAGESETUPDIALOG_H
#define MIQT_QT_PRINTSUPPORT_GEN_QPAGESETUPDIALOG_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDialog QDialog;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPageSetupDialog QPageSetupDialog;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPrinter QPrinter;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QPageSetupDialog* QPageSetupDialog_new(QWidget* parent);
extern __declspec(dllexport) 
QPageSetupDialog* QPageSetupDialog_new2(QPrinter* printer);
extern __declspec(dllexport) 
QPageSetupDialog* QPageSetupDialog_new3();
extern __declspec(dllexport) 
QPageSetupDialog* QPageSetupDialog_new4(QPrinter* printer, QWidget* parent);
extern __declspec(dllexport) 
void QPageSetupDialog_virtbase(QPageSetupDialog* src
, QDialog** outptr_QDialog
);
extern __declspec(dllexport) 
QMetaObject* QPageSetupDialog_MetaObject(const QPageSetupDialog* self);
extern __declspec(dllexport) 
void* QPageSetupDialog_Metacast(QPageSetupDialog* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QPageSetupDialog_Tr(const char* s);
extern __declspec(dllexport) 
void QPageSetupDialog_SetVisible(QPageSetupDialog* self, bool visible);
extern __declspec(dllexport) 
int QPageSetupDialog_Exec(QPageSetupDialog* self);
extern __declspec(dllexport) 
void QPageSetupDialog_Done(QPageSetupDialog* self, int result);
extern __declspec(dllexport) 
QPrinter* QPageSetupDialog_Printer(QPageSetupDialog* self);
extern __declspec(dllexport) 
struct miqt_string QPageSetupDialog_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QPageSetupDialog_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QPageSetupDialog_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QPageSetupDialog_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QPageSetupDialog_override_virtual_Metacast(void* self, intptr_t slot);
void* QPageSetupDialog_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QPageSetupDialog_Delete(QPageSetupDialog* self, bool isSubclass);

}
