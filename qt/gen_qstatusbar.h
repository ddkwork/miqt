#pragma once
#ifndef MIQT_QT_GEN_QSTATUSBAR_H
#define MIQT_QT_GEN_QSTATUSBAR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QStatusBar QStatusBar;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QStatusBar* QStatusBar_new(QWidget* parent);
extern __declspec(dllexport) 
QStatusBar* QStatusBar_new2();
extern __declspec(dllexport) 
void QStatusBar_virtbase(QStatusBar* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QStatusBar_MetaObject(const QStatusBar* self);
extern __declspec(dllexport) 
void* QStatusBar_Metacast(QStatusBar* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QStatusBar_Tr(const char* s);
extern __declspec(dllexport) 
void QStatusBar_AddWidget(QStatusBar* self, QWidget* widget);
extern __declspec(dllexport) 
int QStatusBar_InsertWidget(QStatusBar* self, int index, QWidget* widget);
extern __declspec(dllexport) 
void QStatusBar_AddPermanentWidget(QStatusBar* self, QWidget* widget);
extern __declspec(dllexport) 
int QStatusBar_InsertPermanentWidget(QStatusBar* self, int index, QWidget* widget);
extern __declspec(dllexport) 
void QStatusBar_RemoveWidget(QStatusBar* self, QWidget* widget);
extern __declspec(dllexport) 
void QStatusBar_SetSizeGripEnabled(QStatusBar* self, bool sizeGripEnabled);
extern __declspec(dllexport) 
bool QStatusBar_IsSizeGripEnabled(const QStatusBar* self);
extern __declspec(dllexport) 
struct miqt_string QStatusBar_CurrentMessage(const QStatusBar* self);
extern __declspec(dllexport) 
void QStatusBar_ShowMessage(QStatusBar* self, struct miqt_string text);
extern __declspec(dllexport) 
void QStatusBar_ClearMessage(QStatusBar* self);
extern __declspec(dllexport) 
void QStatusBar_MessageChanged(QStatusBar* self, struct miqt_string text);
void QStatusBar_connect_MessageChanged(QStatusBar* self, intptr_t slot);
extern __declspec(dllexport) 
void QStatusBar_ShowEvent(QStatusBar* self, QShowEvent* param1);
extern __declspec(dllexport) 
void QStatusBar_PaintEvent(QStatusBar* self, QPaintEvent* param1);
extern __declspec(dllexport) 
void QStatusBar_ResizeEvent(QStatusBar* self, QResizeEvent* param1);
extern __declspec(dllexport) 
bool QStatusBar_Event(QStatusBar* self, QEvent* param1);
extern __declspec(dllexport) 
struct miqt_string QStatusBar_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QStatusBar_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QStatusBar_AddWidget2(QStatusBar* self, QWidget* widget, int stretch);
extern __declspec(dllexport) 
int QStatusBar_InsertWidget3(QStatusBar* self, int index, QWidget* widget, int stretch);
extern __declspec(dllexport) 
void QStatusBar_AddPermanentWidget2(QStatusBar* self, QWidget* widget, int stretch);
extern __declspec(dllexport) 
int QStatusBar_InsertPermanentWidget3(QStatusBar* self, int index, QWidget* widget, int stretch);
extern __declspec(dllexport) 
void QStatusBar_ShowMessage2(QStatusBar* self, struct miqt_string text, int timeout);
extern __declspec(dllexport) 
void QStatusBar_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QStatusBar_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QStatusBar_override_virtual_Metacast(void* self, intptr_t slot);
void* QStatusBar_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QStatusBar_Delete(QStatusBar* self, bool isSubclass);

}
