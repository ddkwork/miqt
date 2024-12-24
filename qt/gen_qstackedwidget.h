#pragma once
#ifndef MIQT_QT_GEN_QSTACKEDWIDGET_H
#define MIQT_QT_GEN_QSTACKEDWIDGET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QFrame QFrame;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QStackedWidget QStackedWidget;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QStackedWidget* QStackedWidget_new(QWidget* parent);
extern __declspec(dllexport) 
QStackedWidget* QStackedWidget_new2();
extern __declspec(dllexport) 
void QStackedWidget_virtbase(QStackedWidget* src
, QFrame** outptr_QFrame
);
extern __declspec(dllexport) 
QMetaObject* QStackedWidget_MetaObject(const QStackedWidget* self);
extern __declspec(dllexport) 
void* QStackedWidget_Metacast(QStackedWidget* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QStackedWidget_Tr(const char* s);
extern __declspec(dllexport) 
int QStackedWidget_AddWidget(QStackedWidget* self, QWidget* w);
extern __declspec(dllexport) 
int QStackedWidget_InsertWidget(QStackedWidget* self, int index, QWidget* w);
extern __declspec(dllexport) 
void QStackedWidget_RemoveWidget(QStackedWidget* self, QWidget* w);
extern __declspec(dllexport) 
QWidget* QStackedWidget_CurrentWidget(const QStackedWidget* self);
extern __declspec(dllexport) 
int QStackedWidget_CurrentIndex(const QStackedWidget* self);
extern __declspec(dllexport) 
int QStackedWidget_IndexOf(const QStackedWidget* self, QWidget* param1);
extern __declspec(dllexport) 
QWidget* QStackedWidget_Widget(const QStackedWidget* self, int param1);
extern __declspec(dllexport) 
int QStackedWidget_Count(const QStackedWidget* self);
extern __declspec(dllexport) 
void QStackedWidget_SetCurrentIndex(QStackedWidget* self, int index);
extern __declspec(dllexport) 
void QStackedWidget_SetCurrentWidget(QStackedWidget* self, QWidget* w);
extern __declspec(dllexport) 
void QStackedWidget_CurrentChanged(QStackedWidget* self, int param1);
void QStackedWidget_connect_CurrentChanged(QStackedWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QStackedWidget_WidgetRemoved(QStackedWidget* self, int index);
void QStackedWidget_connect_WidgetRemoved(QStackedWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QStackedWidget_WidgetAdded(QStackedWidget* self, int index);
void QStackedWidget_connect_WidgetAdded(QStackedWidget* self, intptr_t slot);
extern __declspec(dllexport) 
bool QStackedWidget_Event(QStackedWidget* self, QEvent* e);
extern __declspec(dllexport) 
struct miqt_string QStackedWidget_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QStackedWidget_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QStackedWidget_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QStackedWidget_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QStackedWidget_override_virtual_Metacast(void* self, intptr_t slot);
void* QStackedWidget_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QStackedWidget_Delete(QStackedWidget* self, bool isSubclass);

}
