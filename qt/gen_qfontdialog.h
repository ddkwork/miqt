#pragma once
#ifndef MIQT_QT_GEN_QFONTDIALOG_H
#define MIQT_QT_GEN_QFONTDIALOG_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDialog QDialog;
typedef struct QEvent QEvent;
typedef struct QFont QFont;
typedef struct QFontDialog QFontDialog;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QFontDialog* QFontDialog_new(QWidget* parent);
extern __declspec(dllexport) 
QFontDialog* QFontDialog_new2();
extern __declspec(dllexport) 
QFontDialog* QFontDialog_new3(QFont* initial);
extern __declspec(dllexport) 
QFontDialog* QFontDialog_new4(QFont* initial, QWidget* parent);
extern __declspec(dllexport) 
void QFontDialog_virtbase(QFontDialog* src
, QDialog** outptr_QDialog
);
extern __declspec(dllexport) 
QMetaObject* QFontDialog_MetaObject(const QFontDialog* self);
extern __declspec(dllexport) 
void* QFontDialog_Metacast(QFontDialog* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QFontDialog_Tr(const char* s);
extern __declspec(dllexport) 
void QFontDialog_SetCurrentFont(QFontDialog* self, QFont* font);
extern __declspec(dllexport) 
QFont* QFontDialog_CurrentFont(const QFontDialog* self);
extern __declspec(dllexport) 
QFont* QFontDialog_SelectedFont(const QFontDialog* self);
extern __declspec(dllexport) 
void QFontDialog_SetOption(QFontDialog* self, FontDialogOption option);
extern __declspec(dllexport) 
bool QFontDialog_TestOption(const QFontDialog* self, FontDialogOption option);
extern __declspec(dllexport) 
void QFontDialog_SetOptions(QFontDialog* self, FontDialogOptions options);
extern __declspec(dllexport) 
FontDialogOptions QFontDialog_Options(const QFontDialog* self);
extern __declspec(dllexport) 
void QFontDialog_SetVisible(QFontDialog* self, bool visible);
extern __declspec(dllexport) 
QFont* QFontDialog_GetFont(bool* ok);
extern __declspec(dllexport) 
QFont* QFontDialog_GetFont2(bool* ok, QFont* initial);
extern __declspec(dllexport) 
void QFontDialog_CurrentFontChanged(QFontDialog* self, QFont* font);
void QFontDialog_connect_CurrentFontChanged(QFontDialog* self, intptr_t slot);
extern __declspec(dllexport) 
void QFontDialog_FontSelected(QFontDialog* self, QFont* font);
void QFontDialog_connect_FontSelected(QFontDialog* self, intptr_t slot);
extern __declspec(dllexport) 
void QFontDialog_ChangeEvent(QFontDialog* self, QEvent* event);
extern __declspec(dllexport) 
void QFontDialog_Done(QFontDialog* self, int result);
extern __declspec(dllexport) 
bool QFontDialog_EventFilter(QFontDialog* self, QObject* object, QEvent* event);
extern __declspec(dllexport) 
struct miqt_string QFontDialog_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QFontDialog_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QFontDialog_SetOption2(QFontDialog* self, FontDialogOption option, bool on);
extern __declspec(dllexport) 
QFont* QFontDialog_GetFont22(bool* ok, QWidget* parent);
extern __declspec(dllexport) 
QFont* QFontDialog_GetFont3(bool* ok, QFont* initial, QWidget* parent);
extern __declspec(dllexport) 
QFont* QFontDialog_GetFont4(bool* ok, QFont* initial, QWidget* parent, struct miqt_string title);
extern __declspec(dllexport) 
QFont* QFontDialog_GetFont5(bool* ok, QFont* initial, QWidget* parent, struct miqt_string title, FontDialogOptions options);
extern __declspec(dllexport) 
void QFontDialog_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QFontDialog_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QFontDialog_override_virtual_Metacast(void* self, intptr_t slot);
void* QFontDialog_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QFontDialog_Delete(QFontDialog* self, bool isSubclass);

}
