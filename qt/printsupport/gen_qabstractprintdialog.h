#pragma once
#ifndef MIQT_QT_PRINTSUPPORT_GEN_QABSTRACTPRINTDIALOG_H
#define MIQT_QT_PRINTSUPPORT_GEN_QABSTRACTPRINTDIALOG_H

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
typedef struct QPrinter QPrinter;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAbstractPrintDialog* QAbstractPrintDialog_new(QPrinter* printer);
extern __declspec(dllexport) 
QAbstractPrintDialog* QAbstractPrintDialog_new2(QPrinter* printer, QWidget* parent);
extern __declspec(dllexport) 
void QAbstractPrintDialog_virtbase(QAbstractPrintDialog* src
, QDialog** outptr_QDialog
);
extern __declspec(dllexport) 
QMetaObject* QAbstractPrintDialog_MetaObject(const QAbstractPrintDialog* self);
extern __declspec(dllexport) 
void* QAbstractPrintDialog_Metacast(QAbstractPrintDialog* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAbstractPrintDialog_Tr(const char* s);
extern __declspec(dllexport) 
void QAbstractPrintDialog_SetOptionTabs(QAbstractPrintDialog* self, struct miqt_array /* of QWidget* */  tabs);
extern __declspec(dllexport) 
void QAbstractPrintDialog_SetPrintRange(QAbstractPrintDialog* self, PrintRange rangeVal);
extern __declspec(dllexport) 
PrintRange QAbstractPrintDialog_PrintRange(const QAbstractPrintDialog* self);
extern __declspec(dllexport) 
void QAbstractPrintDialog_SetMinMax(QAbstractPrintDialog* self, int min, int max);
extern __declspec(dllexport) 
int QAbstractPrintDialog_MinPage(const QAbstractPrintDialog* self);
extern __declspec(dllexport) 
int QAbstractPrintDialog_MaxPage(const QAbstractPrintDialog* self);
extern __declspec(dllexport) 
void QAbstractPrintDialog_SetFromTo(QAbstractPrintDialog* self, int fromPage, int toPage);
extern __declspec(dllexport) 
int QAbstractPrintDialog_FromPage(const QAbstractPrintDialog* self);
extern __declspec(dllexport) 
int QAbstractPrintDialog_ToPage(const QAbstractPrintDialog* self);
extern __declspec(dllexport) 
QPrinter* QAbstractPrintDialog_Printer(const QAbstractPrintDialog* self);
extern __declspec(dllexport) 
struct miqt_string QAbstractPrintDialog_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAbstractPrintDialog_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAbstractPrintDialog_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAbstractPrintDialog_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAbstractPrintDialog_override_virtual_Metacast(void* self, intptr_t slot);
void* QAbstractPrintDialog_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAbstractPrintDialog_Delete(QAbstractPrintDialog* self, bool isSubclass);

}
