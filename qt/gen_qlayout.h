#pragma once
#ifndef MIQT_QT_GEN_QLAYOUT_H
#define MIQT_QT_GEN_QLAYOUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QLayout QLayout;
typedef struct QLayoutItem QLayoutItem;
typedef struct QMargins QMargins;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QSpacerItem QSpacerItem;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QLayout* QLayout_new(QWidget* parent);
extern __declspec(dllexport) QLayout* QLayout_new2();
extern __declspec(dllexport) void QLayout_virtbase(QLayout* src, QObject** outptr_QObject, QLayoutItem** outptr_QLayoutItem);
extern __declspec(dllexport) QMetaObject* QLayout_MetaObject(const QLayout* self);
extern __declspec(dllexport) void* QLayout_Metacast(QLayout* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QLayout_Tr(const char* s);
extern __declspec(dllexport) int QLayout_Spacing(const QLayout* self);
extern __declspec(dllexport) void QLayout_SetSpacing(QLayout* self, int spacing);
extern __declspec(dllexport) void QLayout_SetContentsMargins(QLayout* self, int left, int top, int right, int bottom);
extern __declspec(dllexport) void QLayout_SetContentsMarginsWithMargins(QLayout* self, QMargins* margins);
extern __declspec(dllexport) void QLayout_UnsetContentsMargins(QLayout* self);
extern __declspec(dllexport) void QLayout_GetContentsMargins(const QLayout* self, int* left, int* top, int* right, int* bottom);
extern __declspec(dllexport) QMargins* QLayout_ContentsMargins(const QLayout* self);
extern __declspec(dllexport) QRect* QLayout_ContentsRect(const QLayout* self);
extern __declspec(dllexport) bool QLayout_SetAlignment(QLayout* self, QWidget* w, int alignment);
extern __declspec(dllexport) bool QLayout_SetAlignment2(QLayout* self, QLayout* l, int alignment);
extern __declspec(dllexport) void QLayout_SetSizeConstraint(QLayout* self, SizeConstraint sizeConstraint);
extern __declspec(dllexport) SizeConstraint QLayout_SizeConstraint(const QLayout* self);
extern __declspec(dllexport) void QLayout_SetMenuBar(QLayout* self, QWidget* w);
extern __declspec(dllexport) QWidget* QLayout_MenuBar(const QLayout* self);
extern __declspec(dllexport) QWidget* QLayout_ParentWidget(const QLayout* self);
extern __declspec(dllexport) void QLayout_Invalidate(QLayout* self);
extern __declspec(dllexport) QRect* QLayout_Geometry(const QLayout* self);
extern __declspec(dllexport) bool QLayout_Activate(QLayout* self);
extern __declspec(dllexport) void QLayout_Update(QLayout* self);
extern __declspec(dllexport) void QLayout_AddWidget(QLayout* self, QWidget* w);
extern __declspec(dllexport) void QLayout_AddItem(QLayout* self, QLayoutItem* param1);
extern __declspec(dllexport) void QLayout_RemoveWidget(QLayout* self, QWidget* w);
extern __declspec(dllexport) void QLayout_RemoveItem(QLayout* self, QLayoutItem* param1);
extern __declspec(dllexport) int QLayout_ExpandingDirections(const QLayout* self);
extern __declspec(dllexport) QSize* QLayout_MinimumSize(const QLayout* self);
extern __declspec(dllexport) QSize* QLayout_MaximumSize(const QLayout* self);
extern __declspec(dllexport) void QLayout_SetGeometry(QLayout* self, QRect* geometry);
extern __declspec(dllexport) QLayoutItem* QLayout_ItemAt(const QLayout* self, int index);
extern __declspec(dllexport) QLayoutItem* QLayout_TakeAt(QLayout* self, int index);
extern __declspec(dllexport) int QLayout_IndexOf(const QLayout* self, QWidget* param1);
extern __declspec(dllexport) int QLayout_IndexOfWithQLayoutItem(const QLayout* self, QLayoutItem* param1);
extern __declspec(dllexport) int QLayout_Count(const QLayout* self);
extern __declspec(dllexport) bool QLayout_IsEmpty(const QLayout* self);
extern __declspec(dllexport) int QLayout_ControlTypes(const QLayout* self);
extern __declspec(dllexport) QLayoutItem* QLayout_ReplaceWidget(QLayout* self, QWidget* from, QWidget* to, int options);
extern __declspec(dllexport) int QLayout_TotalMinimumHeightForWidth(const QLayout* self, int w);
extern __declspec(dllexport) int QLayout_TotalHeightForWidth(const QLayout* self, int w);
extern __declspec(dllexport) QSize* QLayout_TotalMinimumSize(const QLayout* self);
extern __declspec(dllexport) QSize* QLayout_TotalMaximumSize(const QLayout* self);
extern __declspec(dllexport) QSize* QLayout_TotalSizeHint(const QLayout* self);
extern __declspec(dllexport) QLayout* QLayout_Layout(QLayout* self);
extern __declspec(dllexport) void QLayout_SetEnabled(QLayout* self, bool enabled);
extern __declspec(dllexport) bool QLayout_IsEnabled(const QLayout* self);
extern __declspec(dllexport) QSize* QLayout_ClosestAcceptableSize(QWidget* w, QSize* s);
extern __declspec(dllexport) void QLayout_ChildEvent(QLayout* self, QChildEvent* e);
extern __declspec(dllexport) struct miqt_string QLayout_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QLayout_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QLayout_override_virtual_Spacing(void* self, intptr_t slot);
int QLayout_virtualbase_Spacing(const void* self);
extern __declspec(dllexport) void QLayout_override_virtual_SetSpacing(void* self, intptr_t slot);
void QLayout_virtualbase_SetSpacing(void* self, int spacing);
extern __declspec(dllexport) void QLayout_override_virtual_Invalidate(void* self, intptr_t slot);
void QLayout_virtualbase_Invalidate(void* self);
extern __declspec(dllexport) void QLayout_override_virtual_Geometry(void* self, intptr_t slot);
QRect* QLayout_virtualbase_Geometry(const void* self);
extern __declspec(dllexport) void QLayout_override_virtual_AddItem(void* self, intptr_t slot);
void QLayout_virtualbase_AddItem(void* self, QLayoutItem* param1);
extern __declspec(dllexport) void QLayout_override_virtual_ExpandingDirections(void* self, intptr_t slot);
int QLayout_virtualbase_ExpandingDirections(const void* self);
extern __declspec(dllexport) void QLayout_override_virtual_MinimumSize(void* self, intptr_t slot);
QSize* QLayout_virtualbase_MinimumSize(const void* self);
extern __declspec(dllexport) void QLayout_override_virtual_MaximumSize(void* self, intptr_t slot);
QSize* QLayout_virtualbase_MaximumSize(const void* self);
extern __declspec(dllexport) void QLayout_override_virtual_SetGeometry(void* self, intptr_t slot);
void QLayout_virtualbase_SetGeometry(void* self, QRect* geometry);
extern __declspec(dllexport) void QLayout_override_virtual_ItemAt(void* self, intptr_t slot);
QLayoutItem* QLayout_virtualbase_ItemAt(const void* self, int index);
extern __declspec(dllexport) void QLayout_override_virtual_TakeAt(void* self, intptr_t slot);
QLayoutItem* QLayout_virtualbase_TakeAt(void* self, int index);
extern __declspec(dllexport) void QLayout_override_virtual_IndexOf(void* self, intptr_t slot);
int QLayout_virtualbase_IndexOf(const void* self, QWidget* param1);
extern __declspec(dllexport) void QLayout_override_virtual_IndexOfWithQLayoutItem(void* self, intptr_t slot);
int QLayout_virtualbase_IndexOfWithQLayoutItem(const void* self, QLayoutItem* param1);
extern __declspec(dllexport) void QLayout_override_virtual_Count(void* self, intptr_t slot);
int QLayout_virtualbase_Count(const void* self);
extern __declspec(dllexport) void QLayout_override_virtual_IsEmpty(void* self, intptr_t slot);
bool QLayout_virtualbase_IsEmpty(const void* self);
extern __declspec(dllexport) void QLayout_override_virtual_ControlTypes(void* self, intptr_t slot);
int QLayout_virtualbase_ControlTypes(const void* self);
extern __declspec(dllexport) void QLayout_override_virtual_ReplaceWidget(void* self, intptr_t slot);
QLayoutItem* QLayout_virtualbase_ReplaceWidget(void* self, QWidget* from, QWidget* to, int options);
extern __declspec(dllexport) void QLayout_override_virtual_Layout(void* self, intptr_t slot);
QLayout* QLayout_virtualbase_Layout(void* self);
extern __declspec(dllexport) void QLayout_override_virtual_ChildEvent(void* self, intptr_t slot);
void QLayout_virtualbase_ChildEvent(void* self, QChildEvent* e);
extern __declspec(dllexport) void QLayout_override_virtual_Event(void* self, intptr_t slot);
bool QLayout_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QLayout_override_virtual_EventFilter(void* self, intptr_t slot);
bool QLayout_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QLayout_override_virtual_TimerEvent(void* self, intptr_t slot);
void QLayout_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QLayout_override_virtual_CustomEvent(void* self, intptr_t slot);
void QLayout_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QLayout_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QLayout_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QLayout_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QLayout_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QLayout_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QLayout_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QLayout_override_virtual_HasHeightForWidth(void* self, intptr_t slot);
bool QLayout_virtualbase_HasHeightForWidth(const void* self);
extern __declspec(dllexport) void QLayout_override_virtual_HeightForWidth(void* self, intptr_t slot);
int QLayout_virtualbase_HeightForWidth(const void* self, int param1);
extern __declspec(dllexport) void QLayout_override_virtual_MinimumHeightForWidth(void* self, intptr_t slot);
int QLayout_virtualbase_MinimumHeightForWidth(const void* self, int param1);
extern __declspec(dllexport) void QLayout_override_virtual_Widget(void* self, intptr_t slot);
QWidget* QLayout_virtualbase_Widget(const void* self);
extern __declspec(dllexport) void QLayout_override_virtual_SpacerItem(void* self, intptr_t slot);
QSpacerItem* QLayout_virtualbase_SpacerItem(void* self);
extern __declspec(dllexport) void QLayout_Delete(QLayout* self, bool isSubclass);

} 
