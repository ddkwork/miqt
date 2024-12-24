#pragma once
#ifndef MIQT_QT_GEN_QERRORMESSAGE_H
#define MIQT_QT_GEN_QERRORMESSAGE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDialog QDialog;
typedef struct QErrorMessage QErrorMessage;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QErrorMessage* QErrorMessage_new(QWidget* parent);
extern __declspec(dllexport) 
QErrorMessage* QErrorMessage_new2();
extern __declspec(dllexport) 
void QErrorMessage_virtbase(QErrorMessage* src
, QDialog** outptr_QDialog
);
extern __declspec(dllexport) 
QMetaObject* QErrorMessage_MetaObject(const QErrorMessage* self);
extern __declspec(dllexport) 
void* QErrorMessage_Metacast(QErrorMessage* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QErrorMessage_Tr(const char* s);
extern __declspec(dllexport) 
QErrorMessage* QErrorMessage_QtHandler();
extern __declspec(dllexport) 
void QErrorMessage_ShowMessage(QErrorMessage* self, struct miqt_string message);
extern __declspec(dllexport) 
void QErrorMessage_ShowMessage2(QErrorMessage* self, struct miqt_string message, struct miqt_string typeVal);
extern __declspec(dllexport) 
void QErrorMessage_Done(QErrorMessage* self, int param1);
extern __declspec(dllexport) 
void QErrorMessage_ChangeEvent(QErrorMessage* self, QEvent* e);
extern __declspec(dllexport) 
struct miqt_string QErrorMessage_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QErrorMessage_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QErrorMessage_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QErrorMessage_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QErrorMessage_override_virtual_Metacast(void* self, intptr_t slot);
void* QErrorMessage_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QErrorMessage_Delete(QErrorMessage* self, bool isSubclass);

}
