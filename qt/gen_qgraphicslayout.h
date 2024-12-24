#pragma once
#ifndef MIQT_QT_GEN_QGRAPHICSLAYOUT_H
#define MIQT_QT_GEN_QGRAPHICSLAYOUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QEvent;
class QGraphicsLayout;
class QGraphicsLayoutItem;
class QRectF;
class QSizeF;
class _GUID;
class type_info;
#else
typedef struct QEvent QEvent;
typedef struct QGraphicsLayout QGraphicsLayout;
typedef struct QGraphicsLayoutItem QGraphicsLayoutItem;
typedef struct QRectF QRectF;
typedef struct QSizeF QSizeF;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsLayout* QGraphicsLayout_new();
extern __declspec(dllexport) QGraphicsLayout* QGraphicsLayout_new2(QGraphicsLayoutItem* parent);
extern __declspec(dllexport) void QGraphicsLayout_virtbase(QGraphicsLayout* src, QGraphicsLayoutItem** outptr_QGraphicsLayoutItem);
extern __declspec(dllexport) void QGraphicsLayout_SetContentsMargins(QGraphicsLayout* self, double left, double top, double right, double bottom);
extern __declspec(dllexport) void QGraphicsLayout_GetContentsMargins(const QGraphicsLayout* self, double* left, double* top, double* right, double* bottom);
extern __declspec(dllexport) void QGraphicsLayout_Activate(QGraphicsLayout* self);
extern __declspec(dllexport) bool QGraphicsLayout_IsActivated(const QGraphicsLayout* self);
extern __declspec(dllexport) void QGraphicsLayout_Invalidate(QGraphicsLayout* self);
extern __declspec(dllexport) void QGraphicsLayout_UpdateGeometry(QGraphicsLayout* self);
extern __declspec(dllexport) void QGraphicsLayout_WidgetEvent(QGraphicsLayout* self, QEvent* e);
extern __declspec(dllexport) int QGraphicsLayout_Count(const QGraphicsLayout* self);
extern __declspec(dllexport) QGraphicsLayoutItem* QGraphicsLayout_ItemAt(const QGraphicsLayout* self, int i);
extern __declspec(dllexport) void QGraphicsLayout_RemoveAt(QGraphicsLayout* self, int index);
extern __declspec(dllexport) void QGraphicsLayout_SetInstantInvalidatePropagation(bool enable);
extern __declspec(dllexport) bool QGraphicsLayout_InstantInvalidatePropagation();
extern __declspec(dllexport) void QGraphicsLayout_override_virtual_GetContentsMargins(void* self, intptr_t slot);
void QGraphicsLayout_virtualbase_GetContentsMargins(const void* self, double* left, double* top, double* right, double* bottom);
extern __declspec(dllexport) void QGraphicsLayout_override_virtual_Invalidate(void* self, intptr_t slot);
void QGraphicsLayout_virtualbase_Invalidate(void* self);
extern __declspec(dllexport) void QGraphicsLayout_override_virtual_UpdateGeometry(void* self, intptr_t slot);
void QGraphicsLayout_virtualbase_UpdateGeometry(void* self);
extern __declspec(dllexport) void QGraphicsLayout_override_virtual_WidgetEvent(void* self, intptr_t slot);
void QGraphicsLayout_virtualbase_WidgetEvent(void* self, QEvent* e);
extern __declspec(dllexport) void QGraphicsLayout_override_virtual_Count(void* self, intptr_t slot);
int QGraphicsLayout_virtualbase_Count(const void* self);
extern __declspec(dllexport) void QGraphicsLayout_override_virtual_ItemAt(void* self, intptr_t slot);
QGraphicsLayoutItem* QGraphicsLayout_virtualbase_ItemAt(const void* self, int i);
extern __declspec(dllexport) void QGraphicsLayout_override_virtual_RemoveAt(void* self, intptr_t slot);
void QGraphicsLayout_virtualbase_RemoveAt(void* self, int index);
extern __declspec(dllexport) void QGraphicsLayout_override_virtual_SetGeometry(void* self, intptr_t slot);
void QGraphicsLayout_virtualbase_SetGeometry(void* self, QRectF* rect);
extern __declspec(dllexport) void QGraphicsLayout_override_virtual_IsEmpty(void* self, intptr_t slot);
bool QGraphicsLayout_virtualbase_IsEmpty(const void* self);
extern __declspec(dllexport) void QGraphicsLayout_override_virtual_SizeHint(void* self, intptr_t slot);
QSizeF* QGraphicsLayout_virtualbase_SizeHint(const void* self, int which, QSizeF* constraint);
extern __declspec(dllexport) void QGraphicsLayout_Delete(QGraphicsLayout* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
