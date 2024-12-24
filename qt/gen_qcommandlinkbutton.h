#pragma once
#ifndef MIQT_QT_GEN_QCOMMANDLINKBUTTON_H
#define MIQT_QT_GEN_QCOMMANDLINKBUTTON_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractButton QAbstractButton;
typedef struct QCommandLinkButton QCommandLinkButton;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPoint QPoint;
typedef struct QPushButton QPushButton;
typedef struct QSize QSize;
typedef struct QStyleOptionButton QStyleOptionButton;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QCommandLinkButton* QCommandLinkButton_new(QWidget* parent);
extern __declspec(dllexport) QCommandLinkButton* QCommandLinkButton_new2();
extern __declspec(dllexport) QCommandLinkButton* QCommandLinkButton_new3(struct miqt_string text);
extern __declspec(dllexport) QCommandLinkButton* QCommandLinkButton_new4(struct miqt_string text, struct miqt_string description);
extern __declspec(dllexport) QCommandLinkButton* QCommandLinkButton_new5(struct miqt_string text, QWidget* parent);
extern __declspec(dllexport) QCommandLinkButton* QCommandLinkButton_new6(struct miqt_string text, struct miqt_string description, QWidget* parent);
extern __declspec(dllexport) void QCommandLinkButton_virtbase(QCommandLinkButton* src, QPushButton** outptr_QPushButton);
extern __declspec(dllexport) QMetaObject* QCommandLinkButton_MetaObject(const QCommandLinkButton* self);
extern __declspec(dllexport) void* QCommandLinkButton_Metacast(QCommandLinkButton* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QCommandLinkButton_Tr(const char* s);
extern __declspec(dllexport) struct miqt_string QCommandLinkButton_Description(const QCommandLinkButton* self);
extern __declspec(dllexport) void QCommandLinkButton_SetDescription(QCommandLinkButton* self, struct miqt_string description);
extern __declspec(dllexport) QSize* QCommandLinkButton_SizeHint(const QCommandLinkButton* self);
extern __declspec(dllexport) int QCommandLinkButton_HeightForWidth(const QCommandLinkButton* self, int param1);
extern __declspec(dllexport) QSize* QCommandLinkButton_MinimumSizeHint(const QCommandLinkButton* self);
extern __declspec(dllexport) void QCommandLinkButton_InitStyleOption(const QCommandLinkButton* self, QStyleOptionButton* option);
extern __declspec(dllexport) bool QCommandLinkButton_Event(QCommandLinkButton* self, QEvent* e);
extern __declspec(dllexport) void QCommandLinkButton_PaintEvent(QCommandLinkButton* self, QPaintEvent* param1);
extern __declspec(dllexport) struct miqt_string QCommandLinkButton_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QCommandLinkButton_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QCommandLinkButton_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QCommandLinkButton_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QCommandLinkButton_override_virtual_HeightForWidth(void* self, intptr_t slot);
int QCommandLinkButton_virtualbase_HeightForWidth(const void* self, int param1);
extern __declspec(dllexport) void QCommandLinkButton_override_virtual_MinimumSizeHint(void* self, intptr_t slot);
QSize* QCommandLinkButton_virtualbase_MinimumSizeHint(const void* self);
extern __declspec(dllexport) void QCommandLinkButton_override_virtual_InitStyleOption(void* self, intptr_t slot);
void QCommandLinkButton_virtualbase_InitStyleOption(const void* self, QStyleOptionButton* option);
extern __declspec(dllexport) void QCommandLinkButton_override_virtual_Event(void* self, intptr_t slot);
bool QCommandLinkButton_virtualbase_Event(void* self, QEvent* e);
extern __declspec(dllexport) void QCommandLinkButton_override_virtual_PaintEvent(void* self, intptr_t slot);
void QCommandLinkButton_virtualbase_PaintEvent(void* self, QPaintEvent* param1);
extern __declspec(dllexport) void QCommandLinkButton_override_virtual_KeyPressEvent(void* self, intptr_t slot);
void QCommandLinkButton_virtualbase_KeyPressEvent(void* self, QKeyEvent* param1);
extern __declspec(dllexport) void QCommandLinkButton_override_virtual_FocusInEvent(void* self, intptr_t slot);
void QCommandLinkButton_virtualbase_FocusInEvent(void* self, QFocusEvent* param1);
extern __declspec(dllexport) void QCommandLinkButton_override_virtual_FocusOutEvent(void* self, intptr_t slot);
void QCommandLinkButton_virtualbase_FocusOutEvent(void* self, QFocusEvent* param1);
extern __declspec(dllexport) void QCommandLinkButton_override_virtual_MouseMoveEvent(void* self, intptr_t slot);
void QCommandLinkButton_virtualbase_MouseMoveEvent(void* self, QMouseEvent* param1);
extern __declspec(dllexport) void QCommandLinkButton_override_virtual_HitButton(void* self, intptr_t slot);
bool QCommandLinkButton_virtualbase_HitButton(const void* self, QPoint* pos);
extern __declspec(dllexport) void QCommandLinkButton_Delete(QCommandLinkButton* self, bool isSubclass);

} 
