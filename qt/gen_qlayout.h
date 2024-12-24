#pragma once
#ifndef MIQT_QT_GEN_QLAYOUT_H
#define MIQT_QT_GEN_QLAYOUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QLayout QLayout;
typedef struct QLayoutItem QLayoutItem;
typedef struct QMargins QMargins;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QLayout* QLayout_new(QWidget* parent);
extern __declspec(dllexport) 
QLayout* QLayout_new2();
extern __declspec(dllexport) 
void QLayout_virtbase(QLayout* src
, QObject** outptr_QObject
, QLayoutItem** outptr_QLayoutItem
);
extern __declspec(dllexport) 
QMetaObject* QLayout_MetaObject(const QLayout* self);
extern __declspec(dllexport) 
void* QLayout_Metacast(QLayout* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QLayout_Tr(const char* s);
extern __declspec(dllexport) 
int QLayout_Spacing(const QLayout* self);
extern __declspec(dllexport) 
void QLayout_SetSpacing(QLayout* self, int spacing);
extern __declspec(dllexport) 
void QLayout_SetContentsMargins(QLayout* self, int left, int top, int right, int bottom);
extern __declspec(dllexport) 
void QLayout_SetContentsMarginsWithMargins(QLayout* self, QMargins* margins);
extern __declspec(dllexport) 
void QLayout_UnsetContentsMargins(QLayout* self);
extern __declspec(dllexport) 
void QLayout_GetContentsMargins(const QLayout* self, int* left, int* top, int* right, int* bottom);
extern __declspec(dllexport) 
QMargins* QLayout_ContentsMargins(const QLayout* self);
extern __declspec(dllexport) 
QRect* QLayout_ContentsRect(const QLayout* self);
extern __declspec(dllexport) 
bool QLayout_SetAlignment(QLayout* self, QWidget* w, int alignment);
extern __declspec(dllexport) 
bool QLayout_SetAlignment2(QLayout* self, QLayout* l, int alignment);
extern __declspec(dllexport) 
void QLayout_SetSizeConstraint(QLayout* self, SizeConstraint sizeConstraint);
extern __declspec(dllexport) 
SizeConstraint QLayout_SizeConstraint(const QLayout* self);
extern __declspec(dllexport) 
void QLayout_SetMenuBar(QLayout* self, QWidget* w);
extern __declspec(dllexport) 
QWidget* QLayout_MenuBar(const QLayout* self);
extern __declspec(dllexport) 
QWidget* QLayout_ParentWidget(const QLayout* self);
extern __declspec(dllexport) 
void QLayout_Invalidate(QLayout* self);
extern __declspec(dllexport) 
QRect* QLayout_Geometry(const QLayout* self);
extern __declspec(dllexport) 
bool QLayout_Activate(QLayout* self);
extern __declspec(dllexport) 
void QLayout_Update(QLayout* self);
extern __declspec(dllexport) 
void QLayout_AddWidget(QLayout* self, QWidget* w);
extern __declspec(dllexport) 
void QLayout_AddItem(QLayout* self, QLayoutItem* param1);
extern __declspec(dllexport) 
void QLayout_RemoveWidget(QLayout* self, QWidget* w);
extern __declspec(dllexport) 
void QLayout_RemoveItem(QLayout* self, QLayoutItem* param1);
extern __declspec(dllexport) 
int QLayout_ExpandingDirections(const QLayout* self);
extern __declspec(dllexport) 
QSize* QLayout_MinimumSize(const QLayout* self);
extern __declspec(dllexport) 
QSize* QLayout_MaximumSize(const QLayout* self);
extern __declspec(dllexport) 
void QLayout_SetGeometry(QLayout* self, QRect* geometry);
extern __declspec(dllexport) 
QLayoutItem* QLayout_ItemAt(const QLayout* self, int index);
extern __declspec(dllexport) 
QLayoutItem* QLayout_TakeAt(QLayout* self, int index);
extern __declspec(dllexport) 
int QLayout_IndexOf(const QLayout* self, QWidget* param1);
extern __declspec(dllexport) 
int QLayout_IndexOfWithQLayoutItem(const QLayout* self, QLayoutItem* param1);
extern __declspec(dllexport) 
int QLayout_Count(const QLayout* self);
extern __declspec(dllexport) 
bool QLayout_IsEmpty(const QLayout* self);
extern __declspec(dllexport) 
int QLayout_ControlTypes(const QLayout* self);
extern __declspec(dllexport) 
QLayoutItem* QLayout_ReplaceWidget(QLayout* self, QWidget* from, QWidget* to, int options);
extern __declspec(dllexport) 
int QLayout_TotalMinimumHeightForWidth(const QLayout* self, int w);
extern __declspec(dllexport) 
int QLayout_TotalHeightForWidth(const QLayout* self, int w);
extern __declspec(dllexport) 
QSize* QLayout_TotalMinimumSize(const QLayout* self);
extern __declspec(dllexport) 
QSize* QLayout_TotalMaximumSize(const QLayout* self);
extern __declspec(dllexport) 
QSize* QLayout_TotalSizeHint(const QLayout* self);
extern __declspec(dllexport) 
QLayout* QLayout_Layout(QLayout* self);
extern __declspec(dllexport) 
void QLayout_SetEnabled(QLayout* self, bool enabled);
extern __declspec(dllexport) 
bool QLayout_IsEnabled(const QLayout* self);
extern __declspec(dllexport) 
QSize* QLayout_ClosestAcceptableSize(QWidget* w, QSize* s);
extern __declspec(dllexport) 
void QLayout_ChildEvent(QLayout* self, QChildEvent* e);
extern __declspec(dllexport) 
struct miqt_string QLayout_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QLayout_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QLayout_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QLayout_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QLayout_override_virtual_Metacast(void* self, intptr_t slot);
void* QLayout_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QLayout_Delete(QLayout* self, bool isSubclass);

}
