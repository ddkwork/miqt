#pragma once
#ifndef MIQT_QT_GEN_QPROGRESSDIALOG_H
#define MIQT_QT_GEN_QPROGRESSDIALOG_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QCloseEvent QCloseEvent;
typedef struct QDialog QDialog;
typedef struct QEvent QEvent;
typedef struct QLabel QLabel;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QProgressBar QProgressBar;
typedef struct QProgressDialog QProgressDialog;
typedef struct QPushButton QPushButton;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QProgressDialog* QProgressDialog_new(QWidget* parent);
extern __declspec(dllexport) 
QProgressDialog* QProgressDialog_new2();
extern __declspec(dllexport) 
QProgressDialog* QProgressDialog_new3(struct miqt_string labelText, struct miqt_string cancelButtonText, int minimum, int maximum);
extern __declspec(dllexport) 
QProgressDialog* QProgressDialog_new4(QWidget* parent, int flags);
extern __declspec(dllexport) 
QProgressDialog* QProgressDialog_new5(struct miqt_string labelText, struct miqt_string cancelButtonText, int minimum, int maximum, QWidget* parent);
extern __declspec(dllexport) 
QProgressDialog* QProgressDialog_new6(struct miqt_string labelText, struct miqt_string cancelButtonText, int minimum, int maximum, QWidget* parent, int flags);
extern __declspec(dllexport) 
void QProgressDialog_virtbase(QProgressDialog* src
, QDialog** outptr_QDialog
);
extern __declspec(dllexport) 
QMetaObject* QProgressDialog_MetaObject(const QProgressDialog* self);
extern __declspec(dllexport) 
void* QProgressDialog_Metacast(QProgressDialog* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QProgressDialog_Tr(const char* s);
extern __declspec(dllexport) 
void QProgressDialog_SetLabel(QProgressDialog* self, QLabel* label);
extern __declspec(dllexport) 
void QProgressDialog_SetCancelButton(QProgressDialog* self, QPushButton* button);
extern __declspec(dllexport) 
void QProgressDialog_SetBar(QProgressDialog* self, QProgressBar* bar);
extern __declspec(dllexport) 
bool QProgressDialog_WasCanceled(const QProgressDialog* self);
extern __declspec(dllexport) 
int QProgressDialog_Minimum(const QProgressDialog* self);
extern __declspec(dllexport) 
int QProgressDialog_Maximum(const QProgressDialog* self);
extern __declspec(dllexport) 
int QProgressDialog_Value(const QProgressDialog* self);
extern __declspec(dllexport) 
QSize* QProgressDialog_SizeHint(const QProgressDialog* self);
extern __declspec(dllexport) 
struct miqt_string QProgressDialog_LabelText(const QProgressDialog* self);
extern __declspec(dllexport) 
int QProgressDialog_MinimumDuration(const QProgressDialog* self);
extern __declspec(dllexport) 
void QProgressDialog_SetAutoReset(QProgressDialog* self, bool reset);
extern __declspec(dllexport) 
bool QProgressDialog_AutoReset(const QProgressDialog* self);
extern __declspec(dllexport) 
void QProgressDialog_SetAutoClose(QProgressDialog* self, bool close);
extern __declspec(dllexport) 
bool QProgressDialog_AutoClose(const QProgressDialog* self);
extern __declspec(dllexport) 
void QProgressDialog_Cancel(QProgressDialog* self);
extern __declspec(dllexport) 
void QProgressDialog_Reset(QProgressDialog* self);
extern __declspec(dllexport) 
void QProgressDialog_SetMaximum(QProgressDialog* self, int maximum);
extern __declspec(dllexport) 
void QProgressDialog_SetMinimum(QProgressDialog* self, int minimum);
extern __declspec(dllexport) 
void QProgressDialog_SetRange(QProgressDialog* self, int minimum, int maximum);
extern __declspec(dllexport) 
void QProgressDialog_SetValue(QProgressDialog* self, int progress);
extern __declspec(dllexport) 
void QProgressDialog_SetLabelText(QProgressDialog* self, struct miqt_string text);
extern __declspec(dllexport) 
void QProgressDialog_SetCancelButtonText(QProgressDialog* self, struct miqt_string text);
extern __declspec(dllexport) 
void QProgressDialog_SetMinimumDuration(QProgressDialog* self, int ms);
extern __declspec(dllexport) 
void QProgressDialog_Canceled(QProgressDialog* self);
void QProgressDialog_connect_Canceled(QProgressDialog* self, intptr_t slot);
extern __declspec(dllexport) 
void QProgressDialog_ResizeEvent(QProgressDialog* self, QResizeEvent* event);
extern __declspec(dllexport) 
void QProgressDialog_CloseEvent(QProgressDialog* self, QCloseEvent* event);
extern __declspec(dllexport) 
void QProgressDialog_ChangeEvent(QProgressDialog* self, QEvent* event);
extern __declspec(dllexport) 
void QProgressDialog_ShowEvent(QProgressDialog* self, QShowEvent* event);
extern __declspec(dllexport) 
struct miqt_string QProgressDialog_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QProgressDialog_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QProgressDialog_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QProgressDialog_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QProgressDialog_override_virtual_Metacast(void* self, intptr_t slot);
void* QProgressDialog_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QProgressDialog_Delete(QProgressDialog* self, bool isSubclass);

}
