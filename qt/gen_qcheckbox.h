#pragma once
#ifndef MIQT_QT_GEN_QCHECKBOX_H
#define MIQT_QT_GEN_QCHECKBOX_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractButton QAbstractButton;
typedef struct QCheckBox QCheckBox;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPoint QPoint;
typedef struct QSize QSize;
typedef struct QStyleOptionButton QStyleOptionButton;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QCheckBox* QCheckBox_new(QWidget* parent);
extern __declspec(dllexport) 
QCheckBox* QCheckBox_new2();
extern __declspec(dllexport) 
QCheckBox* QCheckBox_new3(struct miqt_string text);
extern __declspec(dllexport) 
QCheckBox* QCheckBox_new4(struct miqt_string text, QWidget* parent);
extern __declspec(dllexport) 
void QCheckBox_virtbase(QCheckBox* src
, QAbstractButton** outptr_QAbstractButton
);
extern __declspec(dllexport) 
QMetaObject* QCheckBox_MetaObject(const QCheckBox* self);
extern __declspec(dllexport) 
void* QCheckBox_Metacast(QCheckBox* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QCheckBox_Tr(const char* s);
extern __declspec(dllexport) 
QSize* QCheckBox_SizeHint(const QCheckBox* self);
extern __declspec(dllexport) 
QSize* QCheckBox_MinimumSizeHint(const QCheckBox* self);
extern __declspec(dllexport) 
void QCheckBox_SetTristate(QCheckBox* self);
extern __declspec(dllexport) 
bool QCheckBox_IsTristate(const QCheckBox* self);
extern __declspec(dllexport) 
int QCheckBox_CheckState(const QCheckBox* self);
extern __declspec(dllexport) 
void QCheckBox_SetCheckState(QCheckBox* self, int state);
extern __declspec(dllexport) 
void QCheckBox_StateChanged(QCheckBox* self, int param1);
void QCheckBox_connect_StateChanged(QCheckBox* self, intptr_t slot);
extern __declspec(dllexport) 
void QCheckBox_CheckStateChanged(QCheckBox* self, int param1);
void QCheckBox_connect_CheckStateChanged(QCheckBox* self, intptr_t slot);
extern __declspec(dllexport) 
bool QCheckBox_Event(QCheckBox* self, QEvent* e);
extern __declspec(dllexport) 
bool QCheckBox_HitButton(const QCheckBox* self, QPoint* pos);
extern __declspec(dllexport) 
void QCheckBox_CheckStateSet(QCheckBox* self);
extern __declspec(dllexport) 
void QCheckBox_NextCheckState(QCheckBox* self);
extern __declspec(dllexport) 
void QCheckBox_PaintEvent(QCheckBox* self, QPaintEvent* param1);
extern __declspec(dllexport) 
void QCheckBox_MouseMoveEvent(QCheckBox* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QCheckBox_InitStyleOption(const QCheckBox* self, QStyleOptionButton* option);
extern __declspec(dllexport) 
struct miqt_string QCheckBox_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QCheckBox_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QCheckBox_SetTristate1(QCheckBox* self, bool y);
extern __declspec(dllexport) 
void QCheckBox_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QCheckBox_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QCheckBox_override_virtual_Metacast(void* self, intptr_t slot);
void* QCheckBox_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QCheckBox_Delete(QCheckBox* self, bool isSubclass);

}
