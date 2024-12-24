#pragma once
#ifndef MIQT_QT_GEN_QPROGRESSBAR_H
#define MIQT_QT_GEN_QPROGRESSBAR_H

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
typedef struct QProgressBar QProgressBar;
typedef struct QSize QSize;
typedef struct QStyleOptionProgressBar QStyleOptionProgressBar;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QProgressBar* QProgressBar_new(QWidget* parent);
extern __declspec(dllexport) 
QProgressBar* QProgressBar_new2();
extern __declspec(dllexport) 
void QProgressBar_virtbase(QProgressBar* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QProgressBar_MetaObject(const QProgressBar* self);
extern __declspec(dllexport) 
void* QProgressBar_Metacast(QProgressBar* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QProgressBar_Tr(const char* s);
extern __declspec(dllexport) 
int QProgressBar_Minimum(const QProgressBar* self);
extern __declspec(dllexport) 
int QProgressBar_Maximum(const QProgressBar* self);
extern __declspec(dllexport) 
int QProgressBar_Value(const QProgressBar* self);
extern __declspec(dllexport) 
struct miqt_string QProgressBar_Text(const QProgressBar* self);
extern __declspec(dllexport) 
void QProgressBar_SetTextVisible(QProgressBar* self, bool visible);
extern __declspec(dllexport) 
bool QProgressBar_IsTextVisible(const QProgressBar* self);
extern __declspec(dllexport) 
int QProgressBar_Alignment(const QProgressBar* self);
extern __declspec(dllexport) 
void QProgressBar_SetAlignment(QProgressBar* self, int alignment);
extern __declspec(dllexport) 
QSize* QProgressBar_SizeHint(const QProgressBar* self);
extern __declspec(dllexport) 
QSize* QProgressBar_MinimumSizeHint(const QProgressBar* self);
extern __declspec(dllexport) 
int QProgressBar_Orientation(const QProgressBar* self);
extern __declspec(dllexport) 
void QProgressBar_SetInvertedAppearance(QProgressBar* self, bool invert);
extern __declspec(dllexport) 
bool QProgressBar_InvertedAppearance(const QProgressBar* self);
extern __declspec(dllexport) 
void QProgressBar_SetTextDirection(QProgressBar* self, int textDirection);
extern __declspec(dllexport) 
int QProgressBar_TextDirection(const QProgressBar* self);
extern __declspec(dllexport) 
void QProgressBar_SetFormat(QProgressBar* self, struct miqt_string format);
extern __declspec(dllexport) 
void QProgressBar_ResetFormat(QProgressBar* self);
extern __declspec(dllexport) 
struct miqt_string QProgressBar_Format(const QProgressBar* self);
extern __declspec(dllexport) 
void QProgressBar_Reset(QProgressBar* self);
extern __declspec(dllexport) 
void QProgressBar_SetRange(QProgressBar* self, int minimum, int maximum);
extern __declspec(dllexport) 
void QProgressBar_SetMinimum(QProgressBar* self, int minimum);
extern __declspec(dllexport) 
void QProgressBar_SetMaximum(QProgressBar* self, int maximum);
extern __declspec(dllexport) 
void QProgressBar_SetValue(QProgressBar* self, int value);
extern __declspec(dllexport) 
void QProgressBar_SetOrientation(QProgressBar* self, int orientation);
extern __declspec(dllexport) 
void QProgressBar_ValueChanged(QProgressBar* self, int value);
void QProgressBar_connect_ValueChanged(QProgressBar* self, intptr_t slot);
extern __declspec(dllexport) 
bool QProgressBar_Event(QProgressBar* self, QEvent* e);
extern __declspec(dllexport) 
void QProgressBar_PaintEvent(QProgressBar* self, QPaintEvent* param1);
extern __declspec(dllexport) 
void QProgressBar_InitStyleOption(const QProgressBar* self, QStyleOptionProgressBar* option);
extern __declspec(dllexport) 
struct miqt_string QProgressBar_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QProgressBar_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QProgressBar_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QProgressBar_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QProgressBar_override_virtual_Metacast(void* self, intptr_t slot);
void* QProgressBar_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QProgressBar_Delete(QProgressBar* self, bool isSubclass);

}
