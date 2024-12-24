#pragma once
#ifndef MIQT_QT_GEN_QGRIDLAYOUT_H
#define MIQT_QT_GEN_QGRIDLAYOUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QGridLayout QGridLayout;
typedef struct QLayout QLayout;
typedef struct QLayoutItem QLayoutItem;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QGridLayout* QGridLayout_new(QWidget* parent);
extern __declspec(dllexport) QGridLayout* QGridLayout_new2();
extern __declspec(dllexport) void QGridLayout_virtbase(QGridLayout* src, QLayout** outptr_QLayout);
extern __declspec(dllexport) QMetaObject* QGridLayout_MetaObject(const QGridLayout* self);
extern __declspec(dllexport) void* QGridLayout_Metacast(QGridLayout* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QGridLayout_Tr(const char* s);
extern __declspec(dllexport) QSize* QGridLayout_SizeHint(const QGridLayout* self);
extern __declspec(dllexport) QSize* QGridLayout_MinimumSize(const QGridLayout* self);
extern __declspec(dllexport) QSize* QGridLayout_MaximumSize(const QGridLayout* self);
extern __declspec(dllexport) void QGridLayout_SetHorizontalSpacing(QGridLayout* self, int spacing);
extern __declspec(dllexport) int QGridLayout_HorizontalSpacing(const QGridLayout* self);
extern __declspec(dllexport) void QGridLayout_SetVerticalSpacing(QGridLayout* self, int spacing);
extern __declspec(dllexport) int QGridLayout_VerticalSpacing(const QGridLayout* self);
extern __declspec(dllexport) void QGridLayout_SetSpacing(QGridLayout* self, int spacing);
extern __declspec(dllexport) int QGridLayout_Spacing(const QGridLayout* self);
extern __declspec(dllexport) void QGridLayout_SetRowStretch(QGridLayout* self, int row, int stretch);
extern __declspec(dllexport) void QGridLayout_SetColumnStretch(QGridLayout* self, int column, int stretch);
extern __declspec(dllexport) int QGridLayout_RowStretch(const QGridLayout* self, int row);
extern __declspec(dllexport) int QGridLayout_ColumnStretch(const QGridLayout* self, int column);
extern __declspec(dllexport) void QGridLayout_SetRowMinimumHeight(QGridLayout* self, int row, int minSize);
extern __declspec(dllexport) void QGridLayout_SetColumnMinimumWidth(QGridLayout* self, int column, int minSize);
extern __declspec(dllexport) int QGridLayout_RowMinimumHeight(const QGridLayout* self, int row);
extern __declspec(dllexport) int QGridLayout_ColumnMinimumWidth(const QGridLayout* self, int column);
extern __declspec(dllexport) int QGridLayout_ColumnCount(const QGridLayout* self);
extern __declspec(dllexport) int QGridLayout_RowCount(const QGridLayout* self);
extern __declspec(dllexport) QRect* QGridLayout_CellRect(const QGridLayout* self, int row, int column);
extern __declspec(dllexport) bool QGridLayout_HasHeightForWidth(const QGridLayout* self);
extern __declspec(dllexport) int QGridLayout_HeightForWidth(const QGridLayout* self, int param1);
extern __declspec(dllexport) int QGridLayout_MinimumHeightForWidth(const QGridLayout* self, int param1);
extern __declspec(dllexport) int QGridLayout_ExpandingDirections(const QGridLayout* self);
extern __declspec(dllexport) void QGridLayout_Invalidate(QGridLayout* self);
extern __declspec(dllexport) void QGridLayout_AddWidget(QGridLayout* self, QWidget* w);
extern __declspec(dllexport) void QGridLayout_AddWidget2(QGridLayout* self, QWidget* param1, int row, int column);
extern __declspec(dllexport) void QGridLayout_AddWidget3(QGridLayout* self, QWidget* param1, int row, int column, int rowSpan, int columnSpan);
extern __declspec(dllexport) void QGridLayout_AddLayout(QGridLayout* self, QLayout* param1, int row, int column);
extern __declspec(dllexport) void QGridLayout_AddLayout2(QGridLayout* self, QLayout* param1, int row, int column, int rowSpan, int columnSpan);
extern __declspec(dllexport) void QGridLayout_SetOriginCorner(QGridLayout* self, int originCorner);
extern __declspec(dllexport) int QGridLayout_OriginCorner(const QGridLayout* self);
extern __declspec(dllexport) QLayoutItem* QGridLayout_ItemAt(const QGridLayout* self, int index);
extern __declspec(dllexport) QLayoutItem* QGridLayout_ItemAtPosition(const QGridLayout* self, int row, int column);
extern __declspec(dllexport) QLayoutItem* QGridLayout_TakeAt(QGridLayout* self, int index);
extern __declspec(dllexport) int QGridLayout_Count(const QGridLayout* self);
extern __declspec(dllexport) void QGridLayout_SetGeometry(QGridLayout* self, QRect* geometry);
extern __declspec(dllexport) void QGridLayout_AddItem(QGridLayout* self, QLayoutItem* item, int row, int column);
extern __declspec(dllexport) void QGridLayout_SetDefaultPositioning(QGridLayout* self, int n, int orient);
extern __declspec(dllexport) void QGridLayout_GetItemPosition(const QGridLayout* self, int idx, int* row, int* column, int* rowSpan, int* columnSpan);
extern __declspec(dllexport) void QGridLayout_AddItemWithQLayoutItem(QGridLayout* self, QLayoutItem* param1);
extern __declspec(dllexport) struct miqt_string QGridLayout_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QGridLayout_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QGridLayout_AddWidget4(QGridLayout* self, QWidget* param1, int row, int column, int param4);
extern __declspec(dllexport) void QGridLayout_AddWidget6(QGridLayout* self, QWidget* param1, int row, int column, int rowSpan, int columnSpan, int param6);
extern __declspec(dllexport) void QGridLayout_AddLayout4(QGridLayout* self, QLayout* param1, int row, int column, int param4);
extern __declspec(dllexport) void QGridLayout_AddLayout6(QGridLayout* self, QLayout* param1, int row, int column, int rowSpan, int columnSpan, int param6);
extern __declspec(dllexport) void QGridLayout_AddItem4(QGridLayout* self, QLayoutItem* item, int row, int column, int rowSpan);
extern __declspec(dllexport) void QGridLayout_AddItem5(QGridLayout* self, QLayoutItem* item, int row, int column, int rowSpan, int columnSpan);
extern __declspec(dllexport) void QGridLayout_AddItem6(QGridLayout* self, QLayoutItem* item, int row, int column, int rowSpan, int columnSpan, int param6);
extern __declspec(dllexport) void QGridLayout_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QGridLayout_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QGridLayout_override_virtual_MinimumSize(void* self, intptr_t slot);
QSize* QGridLayout_virtualbase_MinimumSize(const void* self);
extern __declspec(dllexport) void QGridLayout_override_virtual_MaximumSize(void* self, intptr_t slot);
QSize* QGridLayout_virtualbase_MaximumSize(const void* self);
extern __declspec(dllexport) void QGridLayout_override_virtual_SetSpacing(void* self, intptr_t slot);
void QGridLayout_virtualbase_SetSpacing(void* self, int spacing);
extern __declspec(dllexport) void QGridLayout_override_virtual_Spacing(void* self, intptr_t slot);
int QGridLayout_virtualbase_Spacing(const void* self);
extern __declspec(dllexport) void QGridLayout_override_virtual_HasHeightForWidth(void* self, intptr_t slot);
bool QGridLayout_virtualbase_HasHeightForWidth(const void* self);
extern __declspec(dllexport) void QGridLayout_override_virtual_HeightForWidth(void* self, intptr_t slot);
int QGridLayout_virtualbase_HeightForWidth(const void* self, int param1);
extern __declspec(dllexport) void QGridLayout_override_virtual_MinimumHeightForWidth(void* self, intptr_t slot);
int QGridLayout_virtualbase_MinimumHeightForWidth(const void* self, int param1);
extern __declspec(dllexport) void QGridLayout_override_virtual_ExpandingDirections(void* self, intptr_t slot);
int QGridLayout_virtualbase_ExpandingDirections(const void* self);
extern __declspec(dllexport) void QGridLayout_override_virtual_Invalidate(void* self, intptr_t slot);
void QGridLayout_virtualbase_Invalidate(void* self);
extern __declspec(dllexport) void QGridLayout_override_virtual_ItemAt(void* self, intptr_t slot);
QLayoutItem* QGridLayout_virtualbase_ItemAt(const void* self, int index);
extern __declspec(dllexport) void QGridLayout_override_virtual_TakeAt(void* self, intptr_t slot);
QLayoutItem* QGridLayout_virtualbase_TakeAt(void* self, int index);
extern __declspec(dllexport) void QGridLayout_override_virtual_Count(void* self, intptr_t slot);
int QGridLayout_virtualbase_Count(const void* self);
extern __declspec(dllexport) void QGridLayout_override_virtual_SetGeometry(void* self, intptr_t slot);
void QGridLayout_virtualbase_SetGeometry(void* self, QRect* geometry);
extern __declspec(dllexport) void QGridLayout_override_virtual_AddItemWithQLayoutItem(void* self, intptr_t slot);
void QGridLayout_virtualbase_AddItemWithQLayoutItem(void* self, QLayoutItem* param1);
extern __declspec(dllexport) void QGridLayout_override_virtual_Geometry(void* self, intptr_t slot);
QRect* QGridLayout_virtualbase_Geometry(const void* self);
extern __declspec(dllexport) void QGridLayout_override_virtual_IndexOf(void* self, intptr_t slot);
int QGridLayout_virtualbase_IndexOf(const void* self, QWidget* param1);
extern __declspec(dllexport) void QGridLayout_override_virtual_IsEmpty(void* self, intptr_t slot);
bool QGridLayout_virtualbase_IsEmpty(const void* self);
extern __declspec(dllexport) void QGridLayout_override_virtual_ControlTypes(void* self, intptr_t slot);
int QGridLayout_virtualbase_ControlTypes(const void* self);
extern __declspec(dllexport) void QGridLayout_override_virtual_ReplaceWidget(void* self, intptr_t slot);
QLayoutItem* QGridLayout_virtualbase_ReplaceWidget(void* self, QWidget* from, QWidget* to, int options);
extern __declspec(dllexport) void QGridLayout_override_virtual_Layout(void* self, intptr_t slot);
QLayout* QGridLayout_virtualbase_Layout(void* self);
extern __declspec(dllexport) void QGridLayout_override_virtual_ChildEvent(void* self, intptr_t slot);
void QGridLayout_virtualbase_ChildEvent(void* self, QChildEvent* e);
extern __declspec(dllexport) void QGridLayout_Delete(QGridLayout* self, bool isSubclass);

} 
