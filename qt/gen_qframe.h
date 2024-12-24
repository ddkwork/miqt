#pragma once
#ifndef MIQT_QT_GEN_QFRAME_H
#define MIQT_QT_GEN_QFRAME_H

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
typedef struct QPaintEvent QPaintEvent;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QStyleOptionFrame QStyleOptionFrame;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QFrame* QFrame_new(QWidget* parent);
extern __declspec(dllexport) 
QFrame* QFrame_new2();
extern __declspec(dllexport) 
QFrame* QFrame_new3(QWidget* parent, int f);
extern __declspec(dllexport) 
void QFrame_virtbase(QFrame* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QFrame_MetaObject(const QFrame* self);
extern __declspec(dllexport) 
void* QFrame_Metacast(QFrame* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QFrame_Tr(const char* s);
extern __declspec(dllexport) 
int QFrame_FrameStyle(const QFrame* self);
extern __declspec(dllexport) 
void QFrame_SetFrameStyle(QFrame* self, int frameStyle);
extern __declspec(dllexport) 
int QFrame_FrameWidth(const QFrame* self);
extern __declspec(dllexport) 
QSize* QFrame_SizeHint(const QFrame* self);
extern __declspec(dllexport) 
Shape QFrame_FrameShape(const QFrame* self);
extern __declspec(dllexport) 
void QFrame_SetFrameShape(QFrame* self, Shape frameShape);
extern __declspec(dllexport) 
Shadow QFrame_FrameShadow(const QFrame* self);
extern __declspec(dllexport) 
void QFrame_SetFrameShadow(QFrame* self, Shadow frameShadow);
extern __declspec(dllexport) 
int QFrame_LineWidth(const QFrame* self);
extern __declspec(dllexport) 
void QFrame_SetLineWidth(QFrame* self, int lineWidth);
extern __declspec(dllexport) 
int QFrame_MidLineWidth(const QFrame* self);
extern __declspec(dllexport) 
void QFrame_SetMidLineWidth(QFrame* self, int midLineWidth);
extern __declspec(dllexport) 
QRect* QFrame_FrameRect(const QFrame* self);
extern __declspec(dllexport) 
void QFrame_SetFrameRect(QFrame* self, QRect* frameRect);
extern __declspec(dllexport) 
bool QFrame_Event(QFrame* self, QEvent* e);
extern __declspec(dllexport) 
void QFrame_PaintEvent(QFrame* self, QPaintEvent* param1);
extern __declspec(dllexport) 
void QFrame_ChangeEvent(QFrame* self, QEvent* param1);
extern __declspec(dllexport) 
void QFrame_InitStyleOption(const QFrame* self, QStyleOptionFrame* option);
extern __declspec(dllexport) 
struct miqt_string QFrame_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QFrame_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QFrame_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QFrame_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QFrame_override_virtual_Metacast(void* self, intptr_t slot);
void* QFrame_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QFrame_Delete(QFrame* self, bool isSubclass);

}
