#pragma once
#ifndef MIQT_QT_GEN_QFONTCOMBOBOX_H
#define MIQT_QT_GEN_QFONTCOMBOBOX_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QComboBox QComboBox;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QFont QFont;
typedef struct QFontComboBox QFontComboBox;
typedef struct QHideEvent QHideEvent;
typedef struct QInputMethodEvent QInputMethodEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QStyleOptionComboBox QStyleOptionComboBox;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QFontComboBox* QFontComboBox_new(QWidget* parent);
extern __declspec(dllexport) QFontComboBox* QFontComboBox_new2();
extern __declspec(dllexport) void QFontComboBox_virtbase(QFontComboBox* src, QComboBox** outptr_QComboBox);
extern __declspec(dllexport) QMetaObject* QFontComboBox_MetaObject(const QFontComboBox* self);
extern __declspec(dllexport) void* QFontComboBox_Metacast(QFontComboBox* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QFontComboBox_Tr(const char* s);
extern __declspec(dllexport) void QFontComboBox_SetWritingSystem(QFontComboBox* self, int writingSystem);
extern __declspec(dllexport) int QFontComboBox_WritingSystem(const QFontComboBox* self);
extern __declspec(dllexport) void QFontComboBox_SetFontFilters(QFontComboBox* self, FontFilters filters);
extern __declspec(dllexport) FontFilters QFontComboBox_FontFilters(const QFontComboBox* self);
extern __declspec(dllexport) QFont* QFontComboBox_CurrentFont(const QFontComboBox* self);
extern __declspec(dllexport) QSize* QFontComboBox_SizeHint(const QFontComboBox* self);
extern __declspec(dllexport) void QFontComboBox_SetSampleTextForSystem(QFontComboBox* self, int writingSystem, struct miqt_string sampleText);
extern __declspec(dllexport) struct miqt_string QFontComboBox_SampleTextForSystem(const QFontComboBox* self, int writingSystem);
extern __declspec(dllexport) void QFontComboBox_SetSampleTextForFont(QFontComboBox* self, struct miqt_string fontFamily, struct miqt_string sampleText);
extern __declspec(dllexport) struct miqt_string QFontComboBox_SampleTextForFont(const QFontComboBox* self, struct miqt_string fontFamily);
extern __declspec(dllexport) void QFontComboBox_SetDisplayFont(QFontComboBox* self, struct miqt_string fontFamily, QFont* font);
extern __declspec(dllexport) void QFontComboBox_SetCurrentFont(QFontComboBox* self, QFont* f);
extern __declspec(dllexport) void QFontComboBox_CurrentFontChanged(QFontComboBox* self, QFont* f);
void QFontComboBox_connect_CurrentFontChanged(QFontComboBox* self, intptr_t slot);
extern __declspec(dllexport) bool QFontComboBox_Event(QFontComboBox* self, QEvent* e);
extern __declspec(dllexport) struct miqt_string QFontComboBox_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QFontComboBox_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QFontComboBox_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QFontComboBox_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QFontComboBox_override_virtual_Event(void* self, intptr_t slot);
bool QFontComboBox_virtualbase_Event(void* self, QEvent* e);
extern __declspec(dllexport) void QFontComboBox_override_virtual_SetModel(void* self, intptr_t slot);
void QFontComboBox_virtualbase_SetModel(void* self, QAbstractItemModel* model);
extern __declspec(dllexport) void QFontComboBox_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QFontComboBox_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QFontComboBox_override_virtual_ShowPopup(void* self, intptr_t slot);
void QFontComboBox_virtualbase_ShowPopup(void* self);
extern __declspec(dllexport) void QFontComboBox_override_virtual_HidePopup(void* self, intptr_t slot);
void QFontComboBox_virtualbase_HidePopup(void* self);
extern __declspec(dllexport) void QFontComboBox_override_virtual_InputMethodQuery(void* self, intptr_t slot);
QVariant* QFontComboBox_virtualbase_InputMethodQuery(const void* self, int param1);
extern __declspec(dllexport) void QFontComboBox_override_virtual_FocusInEvent(void* self, intptr_t slot);
void QFontComboBox_virtualbase_FocusInEvent(void* self, QFocusEvent* e);
extern __declspec(dllexport) void QFontComboBox_override_virtual_FocusOutEvent(void* self, intptr_t slot);
void QFontComboBox_virtualbase_FocusOutEvent(void* self, QFocusEvent* e);
extern __declspec(dllexport) void QFontComboBox_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QFontComboBox_virtualbase_ChangeEvent(void* self, QEvent* e);
extern __declspec(dllexport) void QFontComboBox_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QFontComboBox_virtualbase_ResizeEvent(void* self, QResizeEvent* e);
extern __declspec(dllexport) void QFontComboBox_override_virtual_PaintEvent(void* self, intptr_t slot);
void QFontComboBox_virtualbase_PaintEvent(void* self, QPaintEvent* e);
extern __declspec(dllexport) void QFontComboBox_override_virtual_ShowEvent(void* self, intptr_t slot);
void QFontComboBox_virtualbase_ShowEvent(void* self, QShowEvent* e);
extern __declspec(dllexport) void QFontComboBox_override_virtual_HideEvent(void* self, intptr_t slot);
void QFontComboBox_virtualbase_HideEvent(void* self, QHideEvent* e);
extern __declspec(dllexport) void QFontComboBox_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QFontComboBox_virtualbase_MousePressEvent(void* self, QMouseEvent* e);
extern __declspec(dllexport) void QFontComboBox_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QFontComboBox_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* e);
extern __declspec(dllexport) void QFontComboBox_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QFontComboBox_virtualbase_KeyPressEvent(void* self, QKeyEvent* e);
extern __declspec(dllexport) void QFontComboBox_override_virtual_KeyReleaseEvent(void* self, intptr_t slot);
void QFontComboBox_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* e);
extern __declspec(dllexport) void QFontComboBox_override_virtual_WheelEvent(void* self, intptr_t slot);
void QFontComboBox_virtualbase_WheelEvent(void* self, QWheelEvent* e);
extern __declspec(dllexport) void QFontComboBox_override_virtual_ContextMenuEvent(void* self, intptr_t slot);
void QFontComboBox_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* e);
extern __declspec(dllexport) void QFontComboBox_override_virtual_InputMethodEvent(void* self, intptr_t slot);
void QFontComboBox_virtualbase_InputMethodEvent(void* self, QInputMethodEvent* param1);
extern __declspec(dllexport) void QFontComboBox_override_virtual_InitStyleOption(void* self, intptr_t slot);
void QFontComboBox_virtualbase_InitStyleOption(const void* self, QStyleOptionComboBox* option);
extern __declspec(dllexport) void QFontComboBox_Delete(QFontComboBox* self, bool isSubclass);

} 
