#pragma once
#ifndef MIQT_QT_GEN_QSCROLLAREA_H
#define MIQT_QT_GEN_QSCROLLAREA_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractScrollArea QAbstractScrollArea;
typedef struct QEvent QEvent;
typedef struct QFrame QFrame;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QResizeEvent QResizeEvent;
typedef struct QScrollArea QScrollArea;
typedef struct QSize QSize;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QScrollArea* QScrollArea_new(QWidget* parent);
extern __declspec(dllexport) 
QScrollArea* QScrollArea_new2();
extern __declspec(dllexport) 
void QScrollArea_virtbase(QScrollArea* src
, QAbstractScrollArea** outptr_QAbstractScrollArea
);
extern __declspec(dllexport) 
QMetaObject* QScrollArea_MetaObject(const QScrollArea* self);
extern __declspec(dllexport) 
void* QScrollArea_Metacast(QScrollArea* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QScrollArea_Tr(const char* s);
extern __declspec(dllexport) 
QWidget* QScrollArea_Widget(const QScrollArea* self);
extern __declspec(dllexport) 
void QScrollArea_SetWidget(QScrollArea* self, QWidget* widget);
extern __declspec(dllexport) 
QWidget* QScrollArea_TakeWidget(QScrollArea* self);
extern __declspec(dllexport) 
bool QScrollArea_WidgetResizable(const QScrollArea* self);
extern __declspec(dllexport) 
void QScrollArea_SetWidgetResizable(QScrollArea* self, bool resizable);
extern __declspec(dllexport) 
QSize* QScrollArea_SizeHint(const QScrollArea* self);
extern __declspec(dllexport) 
bool QScrollArea_FocusNextPrevChild(QScrollArea* self, bool next);
extern __declspec(dllexport) 
int QScrollArea_Alignment(const QScrollArea* self);
extern __declspec(dllexport) 
void QScrollArea_SetAlignment(QScrollArea* self, int alignment);
extern __declspec(dllexport) 
void QScrollArea_EnsureVisible(QScrollArea* self, int x, int y);
extern __declspec(dllexport) 
void QScrollArea_EnsureWidgetVisible(QScrollArea* self, QWidget* childWidget);
extern __declspec(dllexport) 
bool QScrollArea_Event(QScrollArea* self, QEvent* param1);
extern __declspec(dllexport) 
bool QScrollArea_EventFilter(QScrollArea* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) 
void QScrollArea_ResizeEvent(QScrollArea* self, QResizeEvent* param1);
extern __declspec(dllexport) 
void QScrollArea_ScrollContentsBy(QScrollArea* self, int dx, int dy);
extern __declspec(dllexport) 
QSize* QScrollArea_ViewportSizeHint(const QScrollArea* self);
extern __declspec(dllexport) 
struct miqt_string QScrollArea_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QScrollArea_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QScrollArea_EnsureVisible3(QScrollArea* self, int x, int y, int xmargin);
extern __declspec(dllexport) 
void QScrollArea_EnsureVisible4(QScrollArea* self, int x, int y, int xmargin, int ymargin);
extern __declspec(dllexport) 
void QScrollArea_EnsureWidgetVisible2(QScrollArea* self, QWidget* childWidget, int xmargin);
extern __declspec(dllexport) 
void QScrollArea_EnsureWidgetVisible3(QScrollArea* self, QWidget* childWidget, int xmargin, int ymargin);
extern __declspec(dllexport) 
void QScrollArea_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QScrollArea_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QScrollArea_override_virtual_Metacast(void* self, intptr_t slot);
void* QScrollArea_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QScrollArea_Delete(QScrollArea* self, bool isSubclass);

}
