#pragma once
#ifndef MIQT_QT_GEN_QPUSHBUTTON_H
#define MIQT_QT_GEN_QPUSHBUTTON_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractButton QAbstractButton;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QIcon QIcon;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMenu QMenu;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPoint QPoint;
typedef struct QPushButton QPushButton;
typedef struct QSize QSize;
typedef struct QStyleOptionButton QStyleOptionButton;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QPushButton* QPushButton_new(QWidget* parent);
extern __declspec(dllexport) QPushButton* QPushButton_new2();
extern __declspec(dllexport) QPushButton* QPushButton_new3(struct miqt_string text);
extern __declspec(dllexport) QPushButton* QPushButton_new4(QIcon* icon, struct miqt_string text);
extern __declspec(dllexport) QPushButton* QPushButton_new5(struct miqt_string text, QWidget* parent);
extern __declspec(dllexport) QPushButton* QPushButton_new6(QIcon* icon, struct miqt_string text, QWidget* parent);
extern __declspec(dllexport) void QPushButton_virtbase(QPushButton* src, QAbstractButton** outptr_QAbstractButton);
extern __declspec(dllexport) QMetaObject* QPushButton_MetaObject(const QPushButton* self);
extern __declspec(dllexport) void* QPushButton_Metacast(QPushButton* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QPushButton_Tr(const char* s);
extern __declspec(dllexport) QSize* QPushButton_SizeHint(const QPushButton* self);
extern __declspec(dllexport) QSize* QPushButton_MinimumSizeHint(const QPushButton* self);
extern __declspec(dllexport) bool QPushButton_AutoDefault(const QPushButton* self);
extern __declspec(dllexport) void QPushButton_SetAutoDefault(QPushButton* self, bool autoDefault);
extern __declspec(dllexport) bool QPushButton_IsDefault(const QPushButton* self);
extern __declspec(dllexport) void QPushButton_SetDefault(QPushButton* self, bool defaultVal);
extern __declspec(dllexport) void QPushButton_SetMenu(QPushButton* self, QMenu* menu);
extern __declspec(dllexport) QMenu* QPushButton_Menu(const QPushButton* self);
extern __declspec(dllexport) void QPushButton_SetFlat(QPushButton* self, bool flat);
extern __declspec(dllexport) bool QPushButton_IsFlat(const QPushButton* self);
extern __declspec(dllexport) void QPushButton_ShowMenu(QPushButton* self);
extern __declspec(dllexport) bool QPushButton_Event(QPushButton* self, QEvent* e);
extern __declspec(dllexport) void QPushButton_PaintEvent(QPushButton* self, QPaintEvent* param1);
extern __declspec(dllexport) void QPushButton_KeyPressEvent(QPushButton* self, QKeyEvent* param1);
extern __declspec(dllexport) void QPushButton_FocusInEvent(QPushButton* self, QFocusEvent* param1);
extern __declspec(dllexport) void QPushButton_FocusOutEvent(QPushButton* self, QFocusEvent* param1);
extern __declspec(dllexport) void QPushButton_MouseMoveEvent(QPushButton* self, QMouseEvent* param1);
extern __declspec(dllexport) void QPushButton_InitStyleOption(const QPushButton* self, QStyleOptionButton* option);
extern __declspec(dllexport) bool QPushButton_HitButton(const QPushButton* self, QPoint* pos);
extern __declspec(dllexport) struct miqt_string QPushButton_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QPushButton_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QPushButton_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QPushButton_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QPushButton_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QPushButton_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QPushButton_override_virtual_Event(void* self, intptr_t slot);
bool QPushButton_virtualbase_Event(void* self, QEvent* e);
extern __declspec(dllexport) void QPushButton_override_virtual_PaintEvent(void* self, intptr_t slot);
void QPushButton_virtualbase_PaintEvent(void* self, QPaintEvent* param1);
extern __declspec(dllexport) void QPushButton_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QPushButton_virtualbase_KeyPressEvent(void* self, QKeyEvent* param1);
extern __declspec(dllexport) void QPushButton_override_virtual_FocusInEvent(void* self, intptr_t slot);
void QPushButton_virtualbase_FocusInEvent(void* self, QFocusEvent* param1);
extern __declspec(dllexport) void QPushButton_override_virtual_FocusOutEvent(void* self, intptr_t slot);
void QPushButton_virtualbase_FocusOutEvent(void* self, QFocusEvent* param1);
extern __declspec(dllexport) void QPushButton_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QPushButton_virtualbase_MouseMoveEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QPushButton_override_virtual_InitStyleOption(void* self, intptr_t slot);
void QPushButton_virtualbase_InitStyleOption(const void* self, QStyleOptionButton* option);
extern __declspec(dllexport) void QPushButton_override_virtual_HitButton(void* self, intptr_t slot);
bool QPushButton_virtualbase_HitButton(const void* self, QPoint* pos);
extern __declspec(dllexport) void QPushButton_override_virtual_CheckStateSet(void* self, intptr_t slot);
void QPushButton_virtualbase_CheckStateSet(void* self);
extern __declspec(dllexport) void QPushButton_override_virtual_NextCheckState(void* self, intptr_t slot);
void QPushButton_virtualbase_NextCheckState(void* self);
extern __declspec(dllexport) void QPushButton_override_virtual_KeyReleaseEvent(void* self, intptr_t slot);
void QPushButton_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* e);
extern __declspec(dllexport) void QPushButton_override_virtual_MousePressEvent(void* self, intptr_t slot);
void QPushButton_virtualbase_MousePressEvent(void* self, QMouseEvent* e);
extern __declspec(dllexport) void QPushButton_override_virtual_MouseReleaseEvent(void* self, intptr_t slot);
void QPushButton_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* e);
extern __declspec(dllexport) void QPushButton_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QPushButton_virtualbase_ChangeEvent(void* self, QEvent* e);
extern __declspec(dllexport) void QPushButton_override_virtual_TimerEvent(void* self, intptr_t slot);
void QPushButton_virtualbase_TimerEvent(void* self, QTimerEvent* e);
extern __declspec(dllexport) void QPushButton_Delete(QPushButton* self, bool isSubclass);

} 
