#pragma once
#ifndef MIQT_QT_GEN_QCOLORDIALOG_H
#define MIQT_QT_GEN_QCOLORDIALOG_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QColor QColor;
typedef struct QColorDialog QColorDialog;
typedef struct QDialog QDialog;
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
QColorDialog* QColorDialog_new(QWidget* parent);
extern __declspec(dllexport) 
QColorDialog* QColorDialog_new2();
extern __declspec(dllexport) 
QColorDialog* QColorDialog_new3(QColor* initial);
extern __declspec(dllexport) 
QColorDialog* QColorDialog_new4(QColor* initial, QWidget* parent);
extern __declspec(dllexport) 
void QColorDialog_virtbase(QColorDialog* src
, QDialog** outptr_QDialog
);
extern __declspec(dllexport) 
QMetaObject* QColorDialog_MetaObject(const QColorDialog* self);
extern __declspec(dllexport) 
void* QColorDialog_Metacast(QColorDialog* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QColorDialog_Tr(const char* s);
extern __declspec(dllexport) 
void QColorDialog_SetCurrentColor(QColorDialog* self, QColor* color);
extern __declspec(dllexport) 
QColor* QColorDialog_CurrentColor(const QColorDialog* self);
extern __declspec(dllexport) 
QColor* QColorDialog_SelectedColor(const QColorDialog* self);
extern __declspec(dllexport) 
void QColorDialog_SetOption(QColorDialog* self, ColorDialogOption option);
extern __declspec(dllexport) 
bool QColorDialog_TestOption(const QColorDialog* self, ColorDialogOption option);
extern __declspec(dllexport) 
void QColorDialog_SetOptions(QColorDialog* self, ColorDialogOptions options);
extern __declspec(dllexport) 
ColorDialogOptions QColorDialog_Options(const QColorDialog* self);
extern __declspec(dllexport) 
void QColorDialog_SetVisible(QColorDialog* self, bool visible);
extern __declspec(dllexport) 
QColor* QColorDialog_GetColor();
extern __declspec(dllexport) 
int QColorDialog_CustomCount();
extern __declspec(dllexport) 
QColor* QColorDialog_CustomColor(int index);
extern __declspec(dllexport) 
void QColorDialog_SetCustomColor(int index, QColor* color);
extern __declspec(dllexport) 
QColor* QColorDialog_StandardColor(int index);
extern __declspec(dllexport) 
void QColorDialog_SetStandardColor(int index, QColor* color);
extern __declspec(dllexport) 
void QColorDialog_CurrentColorChanged(QColorDialog* self, QColor* color);
void QColorDialog_connect_CurrentColorChanged(QColorDialog* self, intptr_t slot);
extern __declspec(dllexport) 
void QColorDialog_ColorSelected(QColorDialog* self, QColor* color);
void QColorDialog_connect_ColorSelected(QColorDialog* self, intptr_t slot);
extern __declspec(dllexport) 
void QColorDialog_ChangeEvent(QColorDialog* self, QEvent* event);
extern __declspec(dllexport) 
void QColorDialog_Done(QColorDialog* self, int result);
extern __declspec(dllexport) 
struct miqt_string QColorDialog_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QColorDialog_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QColorDialog_SetOption2(QColorDialog* self, ColorDialogOption option, bool on);
extern __declspec(dllexport) 
QColor* QColorDialog_GetColor1(QColor* initial);
extern __declspec(dllexport) 
QColor* QColorDialog_GetColor2(QColor* initial, QWidget* parent);
extern __declspec(dllexport) 
QColor* QColorDialog_GetColor3(QColor* initial, QWidget* parent, struct miqt_string title);
extern __declspec(dllexport) 
QColor* QColorDialog_GetColor4(QColor* initial, QWidget* parent, struct miqt_string title, ColorDialogOptions options);
extern __declspec(dllexport) 
void QColorDialog_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QColorDialog_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QColorDialog_override_virtual_Metacast(void* self, intptr_t slot);
void* QColorDialog_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QColorDialog_Delete(QColorDialog* self, bool isSubclass);

}
