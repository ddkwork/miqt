#pragma once
#ifndef MIQT_QT_GEN_QTOOLBOX_H
#define MIQT_QT_GEN_QTOOLBOX_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QFrame QFrame;
typedef struct QIcon QIcon;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QStyleOptionFrame QStyleOptionFrame;
typedef struct QToolBox QToolBox;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QToolBox* QToolBox_new(QWidget* parent);
extern __declspec(dllexport) QToolBox* QToolBox_new2();
extern __declspec(dllexport) QToolBox* QToolBox_new3(QWidget* parent, int f);
extern __declspec(dllexport) void QToolBox_virtbase(QToolBox* src, QFrame** outptr_QFrame);
extern __declspec(dllexport) QMetaObject* QToolBox_MetaObject(const QToolBox* self);
extern __declspec(dllexport) void* QToolBox_Metacast(QToolBox* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QToolBox_Tr(const char* s);
extern __declspec(dllexport) int QToolBox_AddItem(QToolBox* self, QWidget* widget, struct miqt_string text);
extern __declspec(dllexport) int QToolBox_AddItem2(QToolBox* self, QWidget* widget, QIcon* icon, struct miqt_string text);
extern __declspec(dllexport) int QToolBox_InsertItem(QToolBox* self, int index, QWidget* widget, struct miqt_string text);
extern __declspec(dllexport) int QToolBox_InsertItem2(QToolBox* self, int index, QWidget* widget, QIcon* icon, struct miqt_string text);
extern __declspec(dllexport) void QToolBox_RemoveItem(QToolBox* self, int index);
extern __declspec(dllexport) void QToolBox_SetItemEnabled(QToolBox* self, int index, bool enabled);
extern __declspec(dllexport) bool QToolBox_IsItemEnabled(const QToolBox* self, int index);
extern __declspec(dllexport) void QToolBox_SetItemText(QToolBox* self, int index, struct miqt_string text);
extern __declspec(dllexport) struct miqt_string QToolBox_ItemText(const QToolBox* self, int index);
extern __declspec(dllexport) void QToolBox_SetItemIcon(QToolBox* self, int index, QIcon* icon);
extern __declspec(dllexport) QIcon* QToolBox_ItemIcon(const QToolBox* self, int index);
extern __declspec(dllexport) void QToolBox_SetItemToolTip(QToolBox* self, int index, struct miqt_string toolTip);
extern __declspec(dllexport) struct miqt_string QToolBox_ItemToolTip(const QToolBox* self, int index);
extern __declspec(dllexport) int QToolBox_CurrentIndex(const QToolBox* self);
extern __declspec(dllexport) QWidget* QToolBox_CurrentWidget(const QToolBox* self);
extern __declspec(dllexport) QWidget* QToolBox_Widget(const QToolBox* self, int index);
extern __declspec(dllexport) int QToolBox_IndexOf(const QToolBox* self, QWidget* widget);
extern __declspec(dllexport) int QToolBox_Count(const QToolBox* self);
extern __declspec(dllexport) void QToolBox_SetCurrentIndex(QToolBox* self, int index);
extern __declspec(dllexport) void QToolBox_SetCurrentWidget(QToolBox* self, QWidget* widget);
extern __declspec(dllexport) void QToolBox_CurrentChanged(QToolBox* self, int index);
void QToolBox_connect_CurrentChanged(QToolBox* self, intptr_t slot);
extern __declspec(dllexport) bool QToolBox_Event(QToolBox* self, QEvent* e);
extern __declspec(dllexport) void QToolBox_ItemInserted(QToolBox* self, int index);
extern __declspec(dllexport) void QToolBox_ItemRemoved(QToolBox* self, int index);
extern __declspec(dllexport) void QToolBox_ShowEvent(QToolBox* self, QShowEvent* e);
extern __declspec(dllexport) void QToolBox_ChangeEvent(QToolBox* self, QEvent* param1);
extern __declspec(dllexport) struct miqt_string QToolBox_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QToolBox_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QToolBox_override_virtual_Event(void* self, intptr_t slot);
bool QToolBox_virtualbase_Event(void* self, QEvent* e);
extern __declspec(dllexport) void QToolBox_override_virtual_ItemInserted(void* self, intptr_t slot);
void QToolBox_virtualbase_ItemInserted(void* self, int index);
extern __declspec(dllexport) void QToolBox_override_virtual_ItemRemoved(void* self, intptr_t slot);
void QToolBox_virtualbase_ItemRemoved(void* self, int index);
extern __declspec(dllexport) void QToolBox_override_virtual_ShowEvent(void* self, intptr_t slot);
void QToolBox_virtualbase_ShowEvent(void* self, QShowEvent* e);
extern __declspec(dllexport) void QToolBox_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QToolBox_virtualbase_ChangeEvent(void* self, QEvent* param1);
extern __declspec(dllexport) void QToolBox_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QToolBox_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QToolBox_override_virtual_PaintEvent(void* self, intptr_t slot);
void QToolBox_virtualbase_PaintEvent(void* self, QPaintEvent* param1);
extern __declspec(dllexport) void QToolBox_override_virtual_InitStyleOption(void* self, intptr_t slot);
void QToolBox_virtualbase_InitStyleOption(const void* self, QStyleOptionFrame* option);
extern __declspec(dllexport) void QToolBox_Delete(QToolBox* self, bool isSubclass);

} 
