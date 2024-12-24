#pragma once
#ifndef MIQT_QT_GEN_QGRAPHICSLAYOUTITEM_H
#define MIQT_QT_GEN_QGRAPHICSLAYOUTITEM_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QGraphicsItem QGraphicsItem;
typedef struct QGraphicsLayoutItem QGraphicsLayoutItem;
typedef struct QRectF QRectF;
typedef struct QSizeF QSizeF;
typedef struct QSizePolicy QSizePolicy;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsLayoutItem* QGraphicsLayoutItem_new();
extern __declspec(dllexport) QGraphicsLayoutItem* QGraphicsLayoutItem_new2(QGraphicsLayoutItem* parent);
extern __declspec(dllexport) QGraphicsLayoutItem* QGraphicsLayoutItem_new3(QGraphicsLayoutItem* parent, bool isLayout);
extern __declspec(dllexport) void QGraphicsLayoutItem_SetSizePolicy(QGraphicsLayoutItem* self, QSizePolicy* policy);
extern __declspec(dllexport) void QGraphicsLayoutItem_SetSizePolicy2(QGraphicsLayoutItem* self, int hPolicy, int vPolicy);
extern __declspec(dllexport) QSizePolicy* QGraphicsLayoutItem_SizePolicy(const QGraphicsLayoutItem* self);
extern __declspec(dllexport) void QGraphicsLayoutItem_SetMinimumSize(QGraphicsLayoutItem* self, QSizeF* size);
extern __declspec(dllexport) void QGraphicsLayoutItem_SetMinimumSize2(QGraphicsLayoutItem* self, double w, double h);
extern __declspec(dllexport) QSizeF* QGraphicsLayoutItem_MinimumSize(const QGraphicsLayoutItem* self);
extern __declspec(dllexport) void QGraphicsLayoutItem_SetMinimumWidth(QGraphicsLayoutItem* self, double width);
extern __declspec(dllexport) double QGraphicsLayoutItem_MinimumWidth(const QGraphicsLayoutItem* self);
extern __declspec(dllexport) void QGraphicsLayoutItem_SetMinimumHeight(QGraphicsLayoutItem* self, double height);
extern __declspec(dllexport) double QGraphicsLayoutItem_MinimumHeight(const QGraphicsLayoutItem* self);
extern __declspec(dllexport) void QGraphicsLayoutItem_SetPreferredSize(QGraphicsLayoutItem* self, QSizeF* size);
extern __declspec(dllexport) void QGraphicsLayoutItem_SetPreferredSize2(QGraphicsLayoutItem* self, double w, double h);
extern __declspec(dllexport) QSizeF* QGraphicsLayoutItem_PreferredSize(const QGraphicsLayoutItem* self);
extern __declspec(dllexport) void QGraphicsLayoutItem_SetPreferredWidth(QGraphicsLayoutItem* self, double width);
extern __declspec(dllexport) double QGraphicsLayoutItem_PreferredWidth(const QGraphicsLayoutItem* self);
extern __declspec(dllexport) void QGraphicsLayoutItem_SetPreferredHeight(QGraphicsLayoutItem* self, double height);
extern __declspec(dllexport) double QGraphicsLayoutItem_PreferredHeight(const QGraphicsLayoutItem* self);
extern __declspec(dllexport) void QGraphicsLayoutItem_SetMaximumSize(QGraphicsLayoutItem* self, QSizeF* size);
extern __declspec(dllexport) void QGraphicsLayoutItem_SetMaximumSize2(QGraphicsLayoutItem* self, double w, double h);
extern __declspec(dllexport) QSizeF* QGraphicsLayoutItem_MaximumSize(const QGraphicsLayoutItem* self);
extern __declspec(dllexport) void QGraphicsLayoutItem_SetMaximumWidth(QGraphicsLayoutItem* self, double width);
extern __declspec(dllexport) double QGraphicsLayoutItem_MaximumWidth(const QGraphicsLayoutItem* self);
extern __declspec(dllexport) void QGraphicsLayoutItem_SetMaximumHeight(QGraphicsLayoutItem* self, double height);
extern __declspec(dllexport) double QGraphicsLayoutItem_MaximumHeight(const QGraphicsLayoutItem* self);
extern __declspec(dllexport) void QGraphicsLayoutItem_SetGeometry(QGraphicsLayoutItem* self, QRectF* rect);
extern __declspec(dllexport) QRectF* QGraphicsLayoutItem_Geometry(const QGraphicsLayoutItem* self);
extern __declspec(dllexport) void QGraphicsLayoutItem_GetContentsMargins(const QGraphicsLayoutItem* self, double* left, double* top, double* right, double* bottom);
extern __declspec(dllexport) QRectF* QGraphicsLayoutItem_ContentsRect(const QGraphicsLayoutItem* self);
extern __declspec(dllexport) QSizeF* QGraphicsLayoutItem_EffectiveSizeHint(const QGraphicsLayoutItem* self, int which);
extern __declspec(dllexport) void QGraphicsLayoutItem_UpdateGeometry(QGraphicsLayoutItem* self);
extern __declspec(dllexport) bool QGraphicsLayoutItem_IsEmpty(const QGraphicsLayoutItem* self);
extern __declspec(dllexport) QGraphicsLayoutItem* QGraphicsLayoutItem_ParentLayoutItem(const QGraphicsLayoutItem* self);
extern __declspec(dllexport) void QGraphicsLayoutItem_SetParentLayoutItem(QGraphicsLayoutItem* self, QGraphicsLayoutItem* parent);
extern __declspec(dllexport) bool QGraphicsLayoutItem_IsLayout(const QGraphicsLayoutItem* self);
extern __declspec(dllexport) QGraphicsItem* QGraphicsLayoutItem_GraphicsItem(const QGraphicsLayoutItem* self);
extern __declspec(dllexport) bool QGraphicsLayoutItem_OwnedByLayout(const QGraphicsLayoutItem* self);
extern __declspec(dllexport) QSizeF* QGraphicsLayoutItem_SizeHint(const QGraphicsLayoutItem* self, int which, QSizeF* constraint);
extern __declspec(dllexport) void QGraphicsLayoutItem_SetSizePolicy3(QGraphicsLayoutItem* self, int hPolicy, int vPolicy, int controlType);
extern __declspec(dllexport) QSizeF* QGraphicsLayoutItem_EffectiveSizeHint2(const QGraphicsLayoutItem* self, int which, QSizeF* constraint);
extern __declspec(dllexport) void QGraphicsLayoutItem_override_virtual_SetGeometry(void* self, intptr_t slot);
void QGraphicsLayoutItem_virtualbase_SetGeometry(void* self, QRectF* rect);
extern __declspec(dllexport) void QGraphicsLayoutItem_override_virtual_GetContentsMargins(void* self, intptr_t slot);
void QGraphicsLayoutItem_virtualbase_GetContentsMargins(const void* self, double* left, double* top, double* right, double* bottom);
extern __declspec(dllexport) void QGraphicsLayoutItem_override_virtual_UpdateGeometry(void* self, intptr_t slot);
void QGraphicsLayoutItem_virtualbase_UpdateGeometry(void* self);
extern __declspec(dllexport) void QGraphicsLayoutItem_override_virtual_IsEmpty(void* self, intptr_t slot);
bool QGraphicsLayoutItem_virtualbase_IsEmpty(const void* self);
extern __declspec(dllexport) void QGraphicsLayoutItem_override_virtual_SizeHint(void* self, intptr_t slot);
QSizeF* QGraphicsLayoutItem_virtualbase_SizeHint(const void* self, int which, QSizeF* constraint);
extern __declspec(dllexport) void QGraphicsLayoutItem_Delete(QGraphicsLayoutItem* self, bool isSubclass);

} 
