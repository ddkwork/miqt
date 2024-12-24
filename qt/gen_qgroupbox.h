#pragma once
#ifndef MIQT_QT_GEN_QGROUPBOX_H
#define MIQT_QT_GEN_QGROUPBOX_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QGroupBox QGroupBox;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QResizeEvent QResizeEvent;
typedef struct QSize QSize;
typedef struct QStyleOptionGroupBox QStyleOptionGroupBox;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QGroupBox* QGroupBox_new(QWidget* parent);
extern __declspec(dllexport) 
QGroupBox* QGroupBox_new2();
extern __declspec(dllexport) 
QGroupBox* QGroupBox_new3(struct miqt_string title);
extern __declspec(dllexport) 
QGroupBox* QGroupBox_new4(struct miqt_string title, QWidget* parent);
extern __declspec(dllexport) 
void QGroupBox_virtbase(QGroupBox* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QGroupBox_MetaObject(const QGroupBox* self);
extern __declspec(dllexport) 
void* QGroupBox_Metacast(QGroupBox* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QGroupBox_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_string QGroupBox_Title(const QGroupBox* self);
extern __declspec(dllexport) 
void QGroupBox_SetTitle(QGroupBox* self, struct miqt_string title);
extern __declspec(dllexport) 
int QGroupBox_Alignment(const QGroupBox* self);
extern __declspec(dllexport) 
void QGroupBox_SetAlignment(QGroupBox* self, int alignment);
extern __declspec(dllexport) 
QSize* QGroupBox_MinimumSizeHint(const QGroupBox* self);
extern __declspec(dllexport) 
bool QGroupBox_IsFlat(const QGroupBox* self);
extern __declspec(dllexport) 
void QGroupBox_SetFlat(QGroupBox* self, bool flat);
extern __declspec(dllexport) 
bool QGroupBox_IsCheckable(const QGroupBox* self);
extern __declspec(dllexport) 
void QGroupBox_SetCheckable(QGroupBox* self, bool checkable);
extern __declspec(dllexport) 
bool QGroupBox_IsChecked(const QGroupBox* self);
extern __declspec(dllexport) 
void QGroupBox_SetChecked(QGroupBox* self, bool checked);
extern __declspec(dllexport) 
void QGroupBox_Clicked(QGroupBox* self);
void QGroupBox_connect_Clicked(QGroupBox* self, intptr_t slot);
extern __declspec(dllexport) 
void QGroupBox_Toggled(QGroupBox* self, bool param1);
void QGroupBox_connect_Toggled(QGroupBox* self, intptr_t slot);
extern __declspec(dllexport) 
bool QGroupBox_Event(QGroupBox* self, QEvent* event);
extern __declspec(dllexport) 
void QGroupBox_ChildEvent(QGroupBox* self, QChildEvent* event);
extern __declspec(dllexport) 
void QGroupBox_ResizeEvent(QGroupBox* self, QResizeEvent* event);
extern __declspec(dllexport) 
void QGroupBox_PaintEvent(QGroupBox* self, QPaintEvent* event);
extern __declspec(dllexport) 
void QGroupBox_FocusInEvent(QGroupBox* self, QFocusEvent* event);
extern __declspec(dllexport) 
void QGroupBox_ChangeEvent(QGroupBox* self, QEvent* event);
extern __declspec(dllexport) 
void QGroupBox_MousePressEvent(QGroupBox* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QGroupBox_MouseMoveEvent(QGroupBox* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QGroupBox_MouseReleaseEvent(QGroupBox* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QGroupBox_InitStyleOption(const QGroupBox* self, QStyleOptionGroupBox* option);
extern __declspec(dllexport) 
struct miqt_string QGroupBox_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QGroupBox_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QGroupBox_Clicked1(QGroupBox* self, bool checked);
void QGroupBox_connect_Clicked1(QGroupBox* self, intptr_t slot);
extern __declspec(dllexport) 
void QGroupBox_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QGroupBox_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QGroupBox_override_virtual_Metacast(void* self, intptr_t slot);
void* QGroupBox_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QGroupBox_Delete(QGroupBox* self, bool isSubclass);

}
