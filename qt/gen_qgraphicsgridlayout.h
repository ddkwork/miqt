#pragma once
#ifndef MIQT_QT_GEN_QGRAPHICSGRIDLAYOUT_H
#define MIQT_QT_GEN_QGRAPHICSGRIDLAYOUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QGraphicsGridLayout QGraphicsGridLayout;
typedef struct QGraphicsLayout QGraphicsLayout;
typedef struct QGraphicsLayoutItem QGraphicsLayoutItem;
typedef struct QRectF QRectF;
typedef struct QSizeF QSizeF;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsGridLayout* QGraphicsGridLayout_new();
extern __declspec(dllexport) QGraphicsGridLayout* QGraphicsGridLayout_new2(QGraphicsLayoutItem* parent);
extern __declspec(dllexport) void QGraphicsGridLayout_virtbase(QGraphicsGridLayout* src, QGraphicsLayout** outptr_QGraphicsLayout);
extern __declspec(dllexport) void QGraphicsGridLayout_AddItem(QGraphicsGridLayout* self, QGraphicsLayoutItem* item, int row, int column, int rowSpan, int columnSpan);
extern __declspec(dllexport) void QGraphicsGridLayout_AddItem2(QGraphicsGridLayout* self, QGraphicsLayoutItem* item, int row, int column);
extern __declspec(dllexport) void QGraphicsGridLayout_SetHorizontalSpacing(QGraphicsGridLayout* self, double spacing);
extern __declspec(dllexport) double QGraphicsGridLayout_HorizontalSpacing(const QGraphicsGridLayout* self);
extern __declspec(dllexport) void QGraphicsGridLayout_SetVerticalSpacing(QGraphicsGridLayout* self, double spacing);
extern __declspec(dllexport) double QGraphicsGridLayout_VerticalSpacing(const QGraphicsGridLayout* self);
extern __declspec(dllexport) void QGraphicsGridLayout_SetSpacing(QGraphicsGridLayout* self, double spacing);
extern __declspec(dllexport) void QGraphicsGridLayout_SetRowSpacing(QGraphicsGridLayout* self, int row, double spacing);
extern __declspec(dllexport) double QGraphicsGridLayout_RowSpacing(const QGraphicsGridLayout* self, int row);
extern __declspec(dllexport) void QGraphicsGridLayout_SetColumnSpacing(QGraphicsGridLayout* self, int column, double spacing);
extern __declspec(dllexport) double QGraphicsGridLayout_ColumnSpacing(const QGraphicsGridLayout* self, int column);
extern __declspec(dllexport) void QGraphicsGridLayout_SetRowStretchFactor(QGraphicsGridLayout* self, int row, int stretch);
extern __declspec(dllexport) int QGraphicsGridLayout_RowStretchFactor(const QGraphicsGridLayout* self, int row);
extern __declspec(dllexport) void QGraphicsGridLayout_SetColumnStretchFactor(QGraphicsGridLayout* self, int column, int stretch);
extern __declspec(dllexport) int QGraphicsGridLayout_ColumnStretchFactor(const QGraphicsGridLayout* self, int column);
extern __declspec(dllexport) void QGraphicsGridLayout_SetRowMinimumHeight(QGraphicsGridLayout* self, int row, double height);
extern __declspec(dllexport) double QGraphicsGridLayout_RowMinimumHeight(const QGraphicsGridLayout* self, int row);
extern __declspec(dllexport) void QGraphicsGridLayout_SetRowPreferredHeight(QGraphicsGridLayout* self, int row, double height);
extern __declspec(dllexport) double QGraphicsGridLayout_RowPreferredHeight(const QGraphicsGridLayout* self, int row);
extern __declspec(dllexport) void QGraphicsGridLayout_SetRowMaximumHeight(QGraphicsGridLayout* self, int row, double height);
extern __declspec(dllexport) double QGraphicsGridLayout_RowMaximumHeight(const QGraphicsGridLayout* self, int row);
extern __declspec(dllexport) void QGraphicsGridLayout_SetRowFixedHeight(QGraphicsGridLayout* self, int row, double height);
extern __declspec(dllexport) void QGraphicsGridLayout_SetColumnMinimumWidth(QGraphicsGridLayout* self, int column, double width);
extern __declspec(dllexport) double QGraphicsGridLayout_ColumnMinimumWidth(const QGraphicsGridLayout* self, int column);
extern __declspec(dllexport) void QGraphicsGridLayout_SetColumnPreferredWidth(QGraphicsGridLayout* self, int column, double width);
extern __declspec(dllexport) double QGraphicsGridLayout_ColumnPreferredWidth(const QGraphicsGridLayout* self, int column);
extern __declspec(dllexport) void QGraphicsGridLayout_SetColumnMaximumWidth(QGraphicsGridLayout* self, int column, double width);
extern __declspec(dllexport) double QGraphicsGridLayout_ColumnMaximumWidth(const QGraphicsGridLayout* self, int column);
extern __declspec(dllexport) void QGraphicsGridLayout_SetColumnFixedWidth(QGraphicsGridLayout* self, int column, double width);
extern __declspec(dllexport) void QGraphicsGridLayout_SetRowAlignment(QGraphicsGridLayout* self, int row, int alignment);
extern __declspec(dllexport) int QGraphicsGridLayout_RowAlignment(const QGraphicsGridLayout* self, int row);
extern __declspec(dllexport) void QGraphicsGridLayout_SetColumnAlignment(QGraphicsGridLayout* self, int column, int alignment);
extern __declspec(dllexport) int QGraphicsGridLayout_ColumnAlignment(const QGraphicsGridLayout* self, int column);
extern __declspec(dllexport) void QGraphicsGridLayout_SetAlignment(QGraphicsGridLayout* self, QGraphicsLayoutItem* item, int alignment);
extern __declspec(dllexport) int QGraphicsGridLayout_Alignment(const QGraphicsGridLayout* self, QGraphicsLayoutItem* item);
extern __declspec(dllexport) int QGraphicsGridLayout_RowCount(const QGraphicsGridLayout* self);
extern __declspec(dllexport) int QGraphicsGridLayout_ColumnCount(const QGraphicsGridLayout* self);
extern __declspec(dllexport) QGraphicsLayoutItem* QGraphicsGridLayout_ItemAt(const QGraphicsGridLayout* self, int row, int column);
extern __declspec(dllexport) int QGraphicsGridLayout_Count(const QGraphicsGridLayout* self);
extern __declspec(dllexport) QGraphicsLayoutItem* QGraphicsGridLayout_ItemAtWithIndex(const QGraphicsGridLayout* self, int index);
extern __declspec(dllexport) void QGraphicsGridLayout_RemoveAt(QGraphicsGridLayout* self, int index);
extern __declspec(dllexport) void QGraphicsGridLayout_RemoveItem(QGraphicsGridLayout* self, QGraphicsLayoutItem* item);
extern __declspec(dllexport) void QGraphicsGridLayout_Invalidate(QGraphicsGridLayout* self);
extern __declspec(dllexport) void QGraphicsGridLayout_SetGeometry(QGraphicsGridLayout* self, QRectF* rect);
extern __declspec(dllexport) QSizeF* QGraphicsGridLayout_SizeHint(const QGraphicsGridLayout* self, int which, QSizeF* constraint);
extern __declspec(dllexport) void QGraphicsGridLayout_AddItem6(QGraphicsGridLayout* self, QGraphicsLayoutItem* item, int row, int column, int rowSpan, int columnSpan, int alignment);
extern __declspec(dllexport) void QGraphicsGridLayout_AddItem4(QGraphicsGridLayout* self, QGraphicsLayoutItem* item, int row, int column, int alignment);
extern __declspec(dllexport) void QGraphicsGridLayout_override_virtual_Count(void* self, intptr_t slot);
int QGraphicsGridLayout_virtualbase_Count(const void* self);
extern __declspec(dllexport) void QGraphicsGridLayout_override_virtual_ItemAtWithIndex(void* self, intptr_t slot);
QGraphicsLayoutItem* QGraphicsGridLayout_virtualbase_ItemAtWithIndex(const void* self, int index);
extern __declspec(dllexport) void QGraphicsGridLayout_override_virtual_RemoveAt(void* self, intptr_t slot);
void QGraphicsGridLayout_virtualbase_RemoveAt(void* self, int index);
extern __declspec(dllexport) void QGraphicsGridLayout_override_virtual_Invalidate(void* self, intptr_t slot);
void QGraphicsGridLayout_virtualbase_Invalidate(void* self);
extern __declspec(dllexport) void QGraphicsGridLayout_override_virtual_SetGeometry(void* self, intptr_t slot);
void QGraphicsGridLayout_virtualbase_SetGeometry(void* self, QRectF* rect);
extern __declspec(dllexport) void QGraphicsGridLayout_override_virtual_SizeHint(void* self, intptr_t slot);
QSizeF* QGraphicsGridLayout_virtualbase_SizeHint(const void* self, int which, QSizeF* constraint);
extern __declspec(dllexport) void QGraphicsGridLayout_override_virtual_GetContentsMargins(void* self, intptr_t slot);
void QGraphicsGridLayout_virtualbase_GetContentsMargins(const void* self, double* left, double* top, double* right, double* bottom);
extern __declspec(dllexport) void QGraphicsGridLayout_override_virtual_UpdateGeometry(void* self, intptr_t slot);
void QGraphicsGridLayout_virtualbase_UpdateGeometry(void* self);
extern __declspec(dllexport) void QGraphicsGridLayout_override_virtual_WidgetEvent(void* self, intptr_t slot);
void QGraphicsGridLayout_virtualbase_WidgetEvent(void* self, QEvent* e);
extern __declspec(dllexport) void QGraphicsGridLayout_Delete(QGraphicsGridLayout* self, bool isSubclass);

} 
