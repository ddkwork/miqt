#pragma once
#ifndef MIQT_QT_GEN_QFORMLAYOUT_H
#define MIQT_QT_GEN_QFORMLAYOUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QFormLayout__TakeRowResult)
typedef QFormLayout::TakeRowResult QFormLayout__TakeRowResult;
typedef struct QChildEvent QChildEvent;
typedef struct QFormLayout QFormLayout;
typedef struct QFormLayout__TakeRowResult QFormLayout__TakeRowResult;
typedef struct QLayout QLayout;
typedef struct QLayoutItem QLayoutItem;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QWidget QWidget;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QFormLayout* QFormLayout_new(QWidget* parent);
extern __declspec(dllexport) QFormLayout* QFormLayout_new2();
extern __declspec(dllexport) void QFormLayout_virtbase(QFormLayout* src, QLayout** outptr_QLayout);
extern __declspec(dllexport) QMetaObject* QFormLayout_MetaObject(const QFormLayout* self);
extern __declspec(dllexport) void* QFormLayout_Metacast(QFormLayout* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QFormLayout_Tr(const char* s);
extern __declspec(dllexport) void QFormLayout_SetFieldGrowthPolicy(QFormLayout* self, FieldGrowthPolicy policy);
extern __declspec(dllexport) FieldGrowthPolicy QFormLayout_FieldGrowthPolicy(const QFormLayout* self);
extern __declspec(dllexport) void QFormLayout_SetRowWrapPolicy(QFormLayout* self, RowWrapPolicy policy);
extern __declspec(dllexport) RowWrapPolicy QFormLayout_RowWrapPolicy(const QFormLayout* self);
extern __declspec(dllexport) void QFormLayout_SetLabelAlignment(QFormLayout* self, int alignment);
extern __declspec(dllexport) int QFormLayout_LabelAlignment(const QFormLayout* self);
extern __declspec(dllexport) void QFormLayout_SetFormAlignment(QFormLayout* self, int alignment);
extern __declspec(dllexport) int QFormLayout_FormAlignment(const QFormLayout* self);
extern __declspec(dllexport) void QFormLayout_SetHorizontalSpacing(QFormLayout* self, int spacing);
extern __declspec(dllexport) int QFormLayout_HorizontalSpacing(const QFormLayout* self);
extern __declspec(dllexport) void QFormLayout_SetVerticalSpacing(QFormLayout* self, int spacing);
extern __declspec(dllexport) int QFormLayout_VerticalSpacing(const QFormLayout* self);
extern __declspec(dllexport) int QFormLayout_Spacing(const QFormLayout* self);
extern __declspec(dllexport) void QFormLayout_SetSpacing(QFormLayout* self, int spacing);
extern __declspec(dllexport) void QFormLayout_AddRow(QFormLayout* self, QWidget* label, QWidget* field);
extern __declspec(dllexport) void QFormLayout_AddRow2(QFormLayout* self, QWidget* label, QLayout* field);
extern __declspec(dllexport) void QFormLayout_AddRow3(QFormLayout* self, struct miqt_string labelText, QWidget* field);
extern __declspec(dllexport) void QFormLayout_AddRow4(QFormLayout* self, struct miqt_string labelText, QLayout* field);
extern __declspec(dllexport) void QFormLayout_AddRowWithWidget(QFormLayout* self, QWidget* widget);
extern __declspec(dllexport) void QFormLayout_AddRowWithLayout(QFormLayout* self, QLayout* layout);
extern __declspec(dllexport) void QFormLayout_InsertRow(QFormLayout* self, int row, QWidget* label, QWidget* field);
extern __declspec(dllexport) void QFormLayout_InsertRow2(QFormLayout* self, int row, QWidget* label, QLayout* field);
extern __declspec(dllexport) void QFormLayout_InsertRow3(QFormLayout* self, int row, struct miqt_string labelText, QWidget* field);
extern __declspec(dllexport) void QFormLayout_InsertRow4(QFormLayout* self, int row, struct miqt_string labelText, QLayout* field);
extern __declspec(dllexport) void QFormLayout_InsertRow5(QFormLayout* self, int row, QWidget* widget);
extern __declspec(dllexport) void QFormLayout_InsertRow6(QFormLayout* self, int row, QLayout* layout);
extern __declspec(dllexport) void QFormLayout_RemoveRow(QFormLayout* self, int row);
extern __declspec(dllexport) void QFormLayout_RemoveRowWithWidget(QFormLayout* self, QWidget* widget);
extern __declspec(dllexport) void QFormLayout_RemoveRowWithLayout(QFormLayout* self, QLayout* layout);
extern __declspec(dllexport) TakeRowResult QFormLayout_TakeRow(QFormLayout* self, int row);
extern __declspec(dllexport) TakeRowResult QFormLayout_TakeRowWithWidget(QFormLayout* self, QWidget* widget);
extern __declspec(dllexport) TakeRowResult QFormLayout_TakeRowWithLayout(QFormLayout* self, QLayout* layout);
extern __declspec(dllexport) void QFormLayout_SetItem(QFormLayout* self, int row, ItemRole role, QLayoutItem* item);
extern __declspec(dllexport) void QFormLayout_SetWidget(QFormLayout* self, int row, ItemRole role, QWidget* widget);
extern __declspec(dllexport) void QFormLayout_SetLayout(QFormLayout* self, int row, ItemRole role, QLayout* layout);
extern __declspec(dllexport) void QFormLayout_SetRowVisible(QFormLayout* self, int row, bool on);
extern __declspec(dllexport) void QFormLayout_SetRowVisible2(QFormLayout* self, QWidget* widget, bool on);
extern __declspec(dllexport) void QFormLayout_SetRowVisible3(QFormLayout* self, QLayout* layout, bool on);
extern __declspec(dllexport) bool QFormLayout_IsRowVisible(const QFormLayout* self, int row);
extern __declspec(dllexport) bool QFormLayout_IsRowVisibleWithWidget(const QFormLayout* self, QWidget* widget);
extern __declspec(dllexport) bool QFormLayout_IsRowVisibleWithLayout(const QFormLayout* self, QLayout* layout);
extern __declspec(dllexport) QLayoutItem* QFormLayout_ItemAt(const QFormLayout* self, int row, ItemRole role);
extern __declspec(dllexport) void QFormLayout_GetItemPosition(const QFormLayout* self, int index, int* rowPtr, ItemRole* rolePtr);
extern __declspec(dllexport) void QFormLayout_GetWidgetPosition(const QFormLayout* self, QWidget* widget, int* rowPtr, ItemRole* rolePtr);
extern __declspec(dllexport) void QFormLayout_GetLayoutPosition(const QFormLayout* self, QLayout* layout, int* rowPtr, ItemRole* rolePtr);
extern __declspec(dllexport) QWidget* QFormLayout_LabelForField(const QFormLayout* self, QWidget* field);
extern __declspec(dllexport) QWidget* QFormLayout_LabelForFieldWithField(const QFormLayout* self, QLayout* field);
extern __declspec(dllexport) void QFormLayout_AddItem(QFormLayout* self, QLayoutItem* item);
extern __declspec(dllexport) QLayoutItem* QFormLayout_ItemAtWithIndex(const QFormLayout* self, int index);
extern __declspec(dllexport) QLayoutItem* QFormLayout_TakeAt(QFormLayout* self, int index);
extern __declspec(dllexport) void QFormLayout_SetGeometry(QFormLayout* self, QRect* rect);
extern __declspec(dllexport) QSize* QFormLayout_MinimumSize(const QFormLayout* self);
extern __declspec(dllexport) QSize* QFormLayout_SizeHint(const QFormLayout* self);
extern __declspec(dllexport) void QFormLayout_Invalidate(QFormLayout* self);
extern __declspec(dllexport) bool QFormLayout_HasHeightForWidth(const QFormLayout* self);
extern __declspec(dllexport) int QFormLayout_HeightForWidth(const QFormLayout* self, int width);
extern __declspec(dllexport) int QFormLayout_ExpandingDirections(const QFormLayout* self);
extern __declspec(dllexport) int QFormLayout_Count(const QFormLayout* self);
extern __declspec(dllexport) int QFormLayout_RowCount(const QFormLayout* self);
extern __declspec(dllexport) struct miqt_string QFormLayout_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QFormLayout_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QFormLayout_override_virtual_Spacing(void* self, intptr_t slot);
int QFormLayout_virtualbase_Spacing(const void* self);
extern __declspec(dllexport) void QFormLayout_override_virtual_SetSpacing(void* self, intptr_t slot);
void QFormLayout_virtualbase_SetSpacing(void* self, int spacing);
extern __declspec(dllexport) void QFormLayout_override_virtual_AddItem(void* self, intptr_t slot);
void QFormLayout_virtualbase_AddItem(void* self, QLayoutItem* item);
extern __declspec(dllexport) void QFormLayout_override_virtual_ItemAtWithIndex(void* self, intptr_t slot);
QLayoutItem* QFormLayout_virtualbase_ItemAtWithIndex(const void* self, int index);
extern __declspec(dllexport) void QFormLayout_override_virtual_TakeAt(void* self, intptr_t slot);
QLayoutItem* QFormLayout_virtualbase_TakeAt(void* self, int index);
extern __declspec(dllexport) void QFormLayout_override_virtual_SetGeometry(void* self, intptr_t slot);
void QFormLayout_virtualbase_SetGeometry(void* self, QRect* rect);
extern __declspec(dllexport) void QFormLayout_override_virtual_MinimumSize(void* self, intptr_t slot);
QSize* QFormLayout_virtualbase_MinimumSize(const void* self);
extern __declspec(dllexport) void QFormLayout_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QFormLayout_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QFormLayout_override_virtual_Invalidate(void* self, intptr_t slot);
void QFormLayout_virtualbase_Invalidate(void* self);
extern __declspec(dllexport) void QFormLayout_override_virtual_HasHeightForWidth(void* self, intptr_t slot);
bool QFormLayout_virtualbase_HasHeightForWidth(const void* self);
extern __declspec(dllexport) void QFormLayout_override_virtual_HeightForWidth(void* self, intptr_t slot);
int QFormLayout_virtualbase_HeightForWidth(const void* self, int width);
extern __declspec(dllexport) void QFormLayout_override_virtual_ExpandingDirections(void* self, intptr_t slot);
int QFormLayout_virtualbase_ExpandingDirections(const void* self);
extern __declspec(dllexport) void QFormLayout_override_virtual_Count(void* self, intptr_t slot);
int QFormLayout_virtualbase_Count(const void* self);
extern __declspec(dllexport) void QFormLayout_override_virtual_Geometry(void* self, intptr_t slot);
QRect* QFormLayout_virtualbase_Geometry(const void* self);
extern __declspec(dllexport) void QFormLayout_override_virtual_MaximumSize(void* self, intptr_t slot);
QSize* QFormLayout_virtualbase_MaximumSize(const void* self);
extern __declspec(dllexport) void QFormLayout_override_virtual_IndexOf(void* self, intptr_t slot);
int QFormLayout_virtualbase_IndexOf(const void* self, QWidget* param1);
extern __declspec(dllexport) void QFormLayout_override_virtual_IsEmpty(void* self, intptr_t slot);
bool QFormLayout_virtualbase_IsEmpty(const void* self);
extern __declspec(dllexport) void QFormLayout_override_virtual_ControlTypes(void* self, intptr_t slot);
int QFormLayout_virtualbase_ControlTypes(const void* self);
extern __declspec(dllexport) void QFormLayout_override_virtual_ReplaceWidget(void* self, intptr_t slot);
QLayoutItem* QFormLayout_virtualbase_ReplaceWidget(void* self, QWidget* from, QWidget* to, int options);
extern __declspec(dllexport) void QFormLayout_override_virtual_Layout(void* self, intptr_t slot);
QLayout* QFormLayout_virtualbase_Layout(void* self);
extern __declspec(dllexport) void QFormLayout_override_virtual_ChildEvent(void* self, intptr_t slot);
void QFormLayout_virtualbase_ChildEvent(void* self, QChildEvent* e);
extern __declspec(dllexport) void QFormLayout_Delete(QFormLayout* self, bool isSubclass);

extern __declspec(dllexport) QFormLayout__TakeRowResult* QFormLayout__TakeRowResult_new();
extern __declspec(dllexport) QFormLayout__TakeRowResult* QFormLayout__TakeRowResult_new2(const TakeRowResult* param1);
extern __declspec(dllexport) void QFormLayout__TakeRowResult_Delete(QFormLayout__TakeRowResult* self, bool isSubclass);

} 
