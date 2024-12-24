#pragma once
#ifndef MIQT_QT_GEN_QFOCUSFRAME_H
#define MIQT_QT_GEN_QFOCUSFRAME_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QFocusFrame QFocusFrame;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QStyleOption QStyleOption;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QFocusFrame* QFocusFrame_new(QWidget* parent);
extern __declspec(dllexport) 
QFocusFrame* QFocusFrame_new2();
extern __declspec(dllexport) 
void QFocusFrame_virtbase(QFocusFrame* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QFocusFrame_MetaObject(const QFocusFrame* self);
extern __declspec(dllexport) 
void* QFocusFrame_Metacast(QFocusFrame* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QFocusFrame_Tr(const char* s);
extern __declspec(dllexport) 
void QFocusFrame_SetWidget(QFocusFrame* self, QWidget* widget);
extern __declspec(dllexport) 
QWidget* QFocusFrame_Widget(const QFocusFrame* self);
extern __declspec(dllexport) 
bool QFocusFrame_Event(QFocusFrame* self, QEvent* e);
extern __declspec(dllexport) 
bool QFocusFrame_EventFilter(QFocusFrame* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) 
void QFocusFrame_PaintEvent(QFocusFrame* self, QPaintEvent* param1);
extern __declspec(dllexport) 
void QFocusFrame_InitStyleOption(const QFocusFrame* self, QStyleOption* option);
extern __declspec(dllexport) 
struct miqt_string QFocusFrame_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QFocusFrame_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QFocusFrame_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QFocusFrame_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QFocusFrame_override_virtual_Metacast(void* self, intptr_t slot);
void* QFocusFrame_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QFocusFrame_Delete(QFocusFrame* self, bool isSubclass);

}
