#pragma once
#ifndef MIQT_QT_GEN_QFONTCOMBOBOX_H
#define MIQT_QT_GEN_QFONTCOMBOBOX_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QComboBox QComboBox;
typedef struct QEvent QEvent;
typedef struct QFont QFont;
typedef struct QFontComboBox QFontComboBox;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QSize QSize;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QFontComboBox* QFontComboBox_new(QWidget* parent);
extern __declspec(dllexport) 
QFontComboBox* QFontComboBox_new2();
extern __declspec(dllexport) 
void QFontComboBox_virtbase(QFontComboBox* src
, QComboBox** outptr_QComboBox
);
extern __declspec(dllexport) 
QMetaObject* QFontComboBox_MetaObject(const QFontComboBox* self);
extern __declspec(dllexport) 
void* QFontComboBox_Metacast(QFontComboBox* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QFontComboBox_Tr(const char* s);
extern __declspec(dllexport) 
void QFontComboBox_SetWritingSystem(QFontComboBox* self, int writingSystem);
extern __declspec(dllexport) 
int QFontComboBox_WritingSystem(const QFontComboBox* self);
extern __declspec(dllexport) 
void QFontComboBox_SetFontFilters(QFontComboBox* self, FontFilters filters);
extern __declspec(dllexport) 
FontFilters QFontComboBox_FontFilters(const QFontComboBox* self);
extern __declspec(dllexport) 
QFont* QFontComboBox_CurrentFont(const QFontComboBox* self);
extern __declspec(dllexport) 
QSize* QFontComboBox_SizeHint(const QFontComboBox* self);
extern __declspec(dllexport) 
void QFontComboBox_SetSampleTextForSystem(QFontComboBox* self, int writingSystem, struct miqt_string sampleText);
extern __declspec(dllexport) 
struct miqt_string QFontComboBox_SampleTextForSystem(const QFontComboBox* self, int writingSystem);
extern __declspec(dllexport) 
void QFontComboBox_SetSampleTextForFont(QFontComboBox* self, struct miqt_string fontFamily, struct miqt_string sampleText);
extern __declspec(dllexport) 
struct miqt_string QFontComboBox_SampleTextForFont(const QFontComboBox* self, struct miqt_string fontFamily);
extern __declspec(dllexport) 
void QFontComboBox_SetDisplayFont(QFontComboBox* self, struct miqt_string fontFamily, QFont* font);
extern __declspec(dllexport) 
void QFontComboBox_SetCurrentFont(QFontComboBox* self, QFont* f);
extern __declspec(dllexport) 
void QFontComboBox_CurrentFontChanged(QFontComboBox* self, QFont* f);
void QFontComboBox_connect_CurrentFontChanged(QFontComboBox* self, intptr_t slot);
extern __declspec(dllexport) 
bool QFontComboBox_Event(QFontComboBox* self, QEvent* e);
extern __declspec(dllexport) 
struct miqt_string QFontComboBox_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QFontComboBox_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QFontComboBox_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QFontComboBox_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QFontComboBox_override_virtual_Metacast(void* self, intptr_t slot);
void* QFontComboBox_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QFontComboBox_Delete(QFontComboBox* self, bool isSubclass);

}
