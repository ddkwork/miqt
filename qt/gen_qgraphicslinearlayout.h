#pragma once
#ifndef MIQT_QT_GEN_QGRAPHICSLINEARLAYOUT_H
#define MIQT_QT_GEN_QGRAPHICSLINEARLAYOUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QGraphicsLayout QGraphicsLayout;
typedef struct QGraphicsLayoutItem QGraphicsLayoutItem;
typedef struct QGraphicsLinearLayout QGraphicsLinearLayout;
typedef struct QRectF QRectF;
typedef struct QSizeF QSizeF;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsLinearLayout* QGraphicsLinearLayout_new();
extern __declspec(dllexport) QGraphicsLinearLayout* QGraphicsLinearLayout_new2(int orientation);
extern __declspec(dllexport) QGraphicsLinearLayout* QGraphicsLinearLayout_new3(QGraphicsLayoutItem* parent);
extern __declspec(dllexport) QGraphicsLinearLayout* QGraphicsLinearLayout_new4(int orientation, QGraphicsLayoutItem* parent);
extern __declspec(dllexport) void QGraphicsLinearLayout_virtbase(QGraphicsLinearLayout* src, QGraphicsLayout** outptr_QGraphicsLayout);
extern __declspec(dllexport) void QGraphicsLinearLayout_SetOrientation(QGraphicsLinearLayout* self, int orientation);
extern __declspec(dllexport) int QGraphicsLinearLayout_Orientation(const QGraphicsLinearLayout* self);
extern __declspec(dllexport) void QGraphicsLinearLayout_AddItem(QGraphicsLinearLayout* self, QGraphicsLayoutItem* item);
extern __declspec(dllexport) void QGraphicsLinearLayout_AddStretch(QGraphicsLinearLayout* self);
extern __declspec(dllexport) void QGraphicsLinearLayout_InsertItem(QGraphicsLinearLayout* self, int index, QGraphicsLayoutItem* item);
extern __declspec(dllexport) void QGraphicsLinearLayout_InsertStretch(QGraphicsLinearLayout* self, int index);
extern __declspec(dllexport) void QGraphicsLinearLayout_RemoveItem(QGraphicsLinearLayout* self, QGraphicsLayoutItem* item);
extern __declspec(dllexport) void QGraphicsLinearLayout_RemoveAt(QGraphicsLinearLayout* self, int index);
extern __declspec(dllexport) void QGraphicsLinearLayout_SetSpacing(QGraphicsLinearLayout* self, double spacing);
extern __declspec(dllexport) double QGraphicsLinearLayout_Spacing(const QGraphicsLinearLayout* self);
extern __declspec(dllexport) void QGraphicsLinearLayout_SetItemSpacing(QGraphicsLinearLayout* self, int index, double spacing);
extern __declspec(dllexport) double QGraphicsLinearLayout_ItemSpacing(const QGraphicsLinearLayout* self, int index);
extern __declspec(dllexport) void QGraphicsLinearLayout_SetStretchFactor(QGraphicsLinearLayout* self, QGraphicsLayoutItem* item, int stretch);
extern __declspec(dllexport) int QGraphicsLinearLayout_StretchFactor(const QGraphicsLinearLayout* self, QGraphicsLayoutItem* item);
extern __declspec(dllexport) void QGraphicsLinearLayout_SetAlignment(QGraphicsLinearLayout* self, QGraphicsLayoutItem* item, int alignment);
extern __declspec(dllexport) int QGraphicsLinearLayout_Alignment(const QGraphicsLinearLayout* self, QGraphicsLayoutItem* item);
extern __declspec(dllexport) void QGraphicsLinearLayout_SetGeometry(QGraphicsLinearLayout* self, QRectF* rect);
extern __declspec(dllexport) int QGraphicsLinearLayout_Count(const QGraphicsLinearLayout* self);
extern __declspec(dllexport) QGraphicsLayoutItem* QGraphicsLinearLayout_ItemAt(const QGraphicsLinearLayout* self, int index);
extern __declspec(dllexport) void QGraphicsLinearLayout_Invalidate(QGraphicsLinearLayout* self);
extern __declspec(dllexport) QSizeF* QGraphicsLinearLayout_SizeHint(const QGraphicsLinearLayout* self, int which, QSizeF* constraint);
extern __declspec(dllexport) void QGraphicsLinearLayout_Dump(const QGraphicsLinearLayout* self);
extern __declspec(dllexport) void QGraphicsLinearLayout_AddStretch1(QGraphicsLinearLayout* self, int stretch);
extern __declspec(dllexport) void QGraphicsLinearLayout_InsertStretch2(QGraphicsLinearLayout* self, int index, int stretch);
extern __declspec(dllexport) void QGraphicsLinearLayout_Dump1(const QGraphicsLinearLayout* self, int indent);
extern __declspec(dllexport) void QGraphicsLinearLayout_override_virtual_RemoveAt(void* self, intptr_t slot);
void QGraphicsLinearLayout_virtualbase_RemoveAt(void* self, int index);
extern __declspec(dllexport) void QGraphicsLinearLayout_override_virtual_SetGeometry(void* self, intptr_t slot);
void QGraphicsLinearLayout_virtualbase_SetGeometry(void* self, QRectF* rect);
extern __declspec(dllexport) void QGraphicsLinearLayout_override_virtual_Count(void* self, intptr_t slot);
int QGraphicsLinearLayout_virtualbase_Count(const void* self);
extern __declspec(dllexport) void QGraphicsLinearLayout_override_virtual_ItemAt(void* self, intptr_t slot);
QGraphicsLayoutItem* QGraphicsLinearLayout_virtualbase_ItemAt(const void* self, int index);
extern __declspec(dllexport) void QGraphicsLinearLayout_override_virtual_Invalidate(void* self, intptr_t slot);
void QGraphicsLinearLayout_virtualbase_Invalidate(void* self);
extern __declspec(dllexport) void QGraphicsLinearLayout_override_virtual_SizeHint(void* self, intptr_t slot);
QSizeF* QGraphicsLinearLayout_virtualbase_SizeHint(const void* self, int which, QSizeF* constraint);
extern __declspec(dllexport) void QGraphicsLinearLayout_override_virtual_GetContentsMargins(void* self, intptr_t slot);
void QGraphicsLinearLayout_virtualbase_GetContentsMargins(const void* self, double* left, double* top, double* right, double* bottom);
extern __declspec(dllexport) void QGraphicsLinearLayout_override_virtual_UpdateGeometry(void* self, intptr_t slot);
void QGraphicsLinearLayout_virtualbase_UpdateGeometry(void* self);
extern __declspec(dllexport) void QGraphicsLinearLayout_override_virtual_WidgetEvent(void* self, intptr_t slot);
void QGraphicsLinearLayout_virtualbase_WidgetEvent(void* self, QEvent* e);
extern __declspec(dllexport) void QGraphicsLinearLayout_Delete(QGraphicsLinearLayout* self, bool isSubclass);

} 
