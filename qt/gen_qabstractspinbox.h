#pragma once
#ifndef MIQT_QT_GEN_QABSTRACTSPINBOX_H
#define MIQT_QT_GEN_QABSTRACTSPINBOX_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractSpinBox QAbstractSpinBox;
typedef struct QCloseEvent QCloseEvent;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QHideEvent QHideEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QStyleOptionSpinBox QStyleOptionSpinBox;
typedef struct QTimerEvent QTimerEvent;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAbstractSpinBox* QAbstractSpinBox_new(QWidget* parent);
extern __declspec(dllexport) 
QAbstractSpinBox* QAbstractSpinBox_new2();
extern __declspec(dllexport) 
void QAbstractSpinBox_virtbase(QAbstractSpinBox* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QAbstractSpinBox_MetaObject(const QAbstractSpinBox* self);
extern __declspec(dllexport) 
void* QAbstractSpinBox_Metacast(QAbstractSpinBox* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAbstractSpinBox_Tr(const char* s);
extern __declspec(dllexport) 
ButtonSymbols QAbstractSpinBox_ButtonSymbols(const QAbstractSpinBox* self);
extern __declspec(dllexport) 
void QAbstractSpinBox_SetButtonSymbols(QAbstractSpinBox* self, ButtonSymbols bs);
extern __declspec(dllexport) 
void QAbstractSpinBox_SetCorrectionMode(QAbstractSpinBox* self, CorrectionMode cm);
extern __declspec(dllexport) 
CorrectionMode QAbstractSpinBox_CorrectionMode(const QAbstractSpinBox* self);
extern __declspec(dllexport) 
bool QAbstractSpinBox_HasAcceptableInput(const QAbstractSpinBox* self);
extern __declspec(dllexport) 
struct miqt_string QAbstractSpinBox_Text(const QAbstractSpinBox* self);
extern __declspec(dllexport) 
struct miqt_string QAbstractSpinBox_SpecialValueText(const QAbstractSpinBox* self);
extern __declspec(dllexport) 
void QAbstractSpinBox_SetSpecialValueText(QAbstractSpinBox* self, struct miqt_string txt);
extern __declspec(dllexport) 
bool QAbstractSpinBox_Wrapping(const QAbstractSpinBox* self);
extern __declspec(dllexport) 
void QAbstractSpinBox_SetWrapping(QAbstractSpinBox* self, bool w);
extern __declspec(dllexport) 
void QAbstractSpinBox_SetReadOnly(QAbstractSpinBox* self, bool r);
extern __declspec(dllexport) 
bool QAbstractSpinBox_IsReadOnly(const QAbstractSpinBox* self);
extern __declspec(dllexport) 
void QAbstractSpinBox_SetKeyboardTracking(QAbstractSpinBox* self, bool kt);
extern __declspec(dllexport) 
bool QAbstractSpinBox_KeyboardTracking(const QAbstractSpinBox* self);
extern __declspec(dllexport) 
void QAbstractSpinBox_SetAlignment(QAbstractSpinBox* self, int flag);
extern __declspec(dllexport) 
int QAbstractSpinBox_Alignment(const QAbstractSpinBox* self);
extern __declspec(dllexport) 
void QAbstractSpinBox_SetFrame(QAbstractSpinBox* self, bool frame);
extern __declspec(dllexport) 
bool QAbstractSpinBox_HasFrame(const QAbstractSpinBox* self);
extern __declspec(dllexport) 
void QAbstractSpinBox_SetAccelerated(QAbstractSpinBox* self, bool on);
extern __declspec(dllexport) 
bool QAbstractSpinBox_IsAccelerated(const QAbstractSpinBox* self);
extern __declspec(dllexport) 
void QAbstractSpinBox_SetGroupSeparatorShown(QAbstractSpinBox* self, bool shown);
extern __declspec(dllexport) 
bool QAbstractSpinBox_IsGroupSeparatorShown(const QAbstractSpinBox* self);
extern __declspec(dllexport) 
QSize* QAbstractSpinBox_SizeHint(const QAbstractSpinBox* self);
extern __declspec(dllexport) 
QSize* QAbstractSpinBox_MinimumSizeHint(const QAbstractSpinBox* self);
extern __declspec(dllexport) 
void QAbstractSpinBox_InterpretText(QAbstractSpinBox* self);
extern __declspec(dllexport) 
bool QAbstractSpinBox_Event(QAbstractSpinBox* self, QEvent* event);
extern __declspec(dllexport) 
QVariant* QAbstractSpinBox_InputMethodQuery(const QAbstractSpinBox* self, int param1);
extern __declspec(dllexport) 
int QAbstractSpinBox_Validate(const QAbstractSpinBox* self, struct miqt_string input, int* pos);
extern __declspec(dllexport) 
void QAbstractSpinBox_Fixup(const QAbstractSpinBox* self, struct miqt_string input);
extern __declspec(dllexport) 
void QAbstractSpinBox_StepBy(QAbstractSpinBox* self, int steps);
extern __declspec(dllexport) 
void QAbstractSpinBox_StepUp(QAbstractSpinBox* self);
extern __declspec(dllexport) 
void QAbstractSpinBox_StepDown(QAbstractSpinBox* self);
extern __declspec(dllexport) 
void QAbstractSpinBox_SelectAll(QAbstractSpinBox* self);
extern __declspec(dllexport) 
void QAbstractSpinBox_Clear(QAbstractSpinBox* self);
extern __declspec(dllexport) 
void QAbstractSpinBox_ResizeEvent(QAbstractSpinBox* self, QResizeEvent* event);
extern __declspec(dllexport) 
void QAbstractSpinBox_KeyPressEvent(QAbstractSpinBox* self, QKeyEvent* event);
extern __declspec(dllexport) 
void QAbstractSpinBox_KeyReleaseEvent(QAbstractSpinBox* self, QKeyEvent* event);
extern __declspec(dllexport) 
void QAbstractSpinBox_WheelEvent(QAbstractSpinBox* self, QWheelEvent* event);
extern __declspec(dllexport) 
void QAbstractSpinBox_FocusInEvent(QAbstractSpinBox* self, QFocusEvent* event);
extern __declspec(dllexport) 
void QAbstractSpinBox_FocusOutEvent(QAbstractSpinBox* self, QFocusEvent* event);
extern __declspec(dllexport) 
void QAbstractSpinBox_ContextMenuEvent(QAbstractSpinBox* self, QContextMenuEvent* event);
extern __declspec(dllexport) 
void QAbstractSpinBox_ChangeEvent(QAbstractSpinBox* self, QEvent* event);
extern __declspec(dllexport) 
void QAbstractSpinBox_CloseEvent(QAbstractSpinBox* self, QCloseEvent* event);
extern __declspec(dllexport) 
void QAbstractSpinBox_HideEvent(QAbstractSpinBox* self, QHideEvent* event);
extern __declspec(dllexport) 
void QAbstractSpinBox_MousePressEvent(QAbstractSpinBox* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QAbstractSpinBox_MouseReleaseEvent(QAbstractSpinBox* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QAbstractSpinBox_MouseMoveEvent(QAbstractSpinBox* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QAbstractSpinBox_TimerEvent(QAbstractSpinBox* self, QTimerEvent* event);
extern __declspec(dllexport) 
void QAbstractSpinBox_PaintEvent(QAbstractSpinBox* self, QPaintEvent* event);
extern __declspec(dllexport) 
void QAbstractSpinBox_ShowEvent(QAbstractSpinBox* self, QShowEvent* event);
extern __declspec(dllexport) 
void QAbstractSpinBox_InitStyleOption(const QAbstractSpinBox* self, QStyleOptionSpinBox* option);
extern __declspec(dllexport) 
StepEnabled QAbstractSpinBox_StepEnabled(const QAbstractSpinBox* self);
extern __declspec(dllexport) 
void QAbstractSpinBox_EditingFinished(QAbstractSpinBox* self);
void QAbstractSpinBox_connect_EditingFinished(QAbstractSpinBox* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QAbstractSpinBox_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAbstractSpinBox_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAbstractSpinBox_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAbstractSpinBox_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAbstractSpinBox_override_virtual_Metacast(void* self, intptr_t slot);
void* QAbstractSpinBox_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAbstractSpinBox_Delete(QAbstractSpinBox* self, bool isSubclass);

}
