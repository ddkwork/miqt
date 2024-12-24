#pragma once
#ifndef MIQT_QT_GEN_QGRAPHICSANCHORLAYOUT_H
#define MIQT_QT_GEN_QGRAPHICSANCHORLAYOUT_H

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
class QGraphicsAnchor;
class QGraphicsAnchorLayout;
class QGraphicsLayout;
class QGraphicsLayoutItem;
class QMetaObject;
class QObject;
class QRectF;
class QSizeF;
class _GUID;
class type_info;
#else
typedef struct QEvent QEvent;
typedef struct QGraphicsAnchor QGraphicsAnchor;
typedef struct QGraphicsAnchorLayout QGraphicsAnchorLayout;
typedef struct QGraphicsLayout QGraphicsLayout;
typedef struct QGraphicsLayoutItem QGraphicsLayoutItem;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QRectF QRectF;
typedef struct QSizeF QSizeF;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QGraphicsAnchor_virtbase(QGraphicsAnchor* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QGraphicsAnchor_MetaObject(const QGraphicsAnchor* self);
extern __declspec(dllexport) void* QGraphicsAnchor_Metacast(QGraphicsAnchor* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QGraphicsAnchor_Tr(const char* s);
extern __declspec(dllexport) void QGraphicsAnchor_SetSpacing(QGraphicsAnchor* self, double spacing);
extern __declspec(dllexport) void QGraphicsAnchor_UnsetSpacing(QGraphicsAnchor* self);
extern __declspec(dllexport) double QGraphicsAnchor_Spacing(const QGraphicsAnchor* self);
extern __declspec(dllexport) void QGraphicsAnchor_SetSizePolicy(QGraphicsAnchor* self, int policy);
extern __declspec(dllexport) int QGraphicsAnchor_SizePolicy(const QGraphicsAnchor* self);
extern __declspec(dllexport) struct miqt_string QGraphicsAnchor_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QGraphicsAnchor_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QGraphicsAnchor_Delete(QGraphicsAnchor* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsAnchorLayout* QGraphicsAnchorLayout_new();
extern __declspec(dllexport) QGraphicsAnchorLayout* QGraphicsAnchorLayout_new2(QGraphicsLayoutItem* parent);
extern __declspec(dllexport) void QGraphicsAnchorLayout_virtbase(QGraphicsAnchorLayout* src, QGraphicsLayout** outptr_QGraphicsLayout);
extern __declspec(dllexport) QGraphicsAnchor* QGraphicsAnchorLayout_AddAnchor(QGraphicsAnchorLayout* self, QGraphicsLayoutItem* firstItem, int firstEdge, QGraphicsLayoutItem* secondItem, int secondEdge);
extern __declspec(dllexport) QGraphicsAnchor* QGraphicsAnchorLayout_Anchor(QGraphicsAnchorLayout* self, QGraphicsLayoutItem* firstItem, int firstEdge, QGraphicsLayoutItem* secondItem, int secondEdge);
extern __declspec(dllexport) void QGraphicsAnchorLayout_AddCornerAnchors(QGraphicsAnchorLayout* self, QGraphicsLayoutItem* firstItem, int firstCorner, QGraphicsLayoutItem* secondItem, int secondCorner);
extern __declspec(dllexport) void QGraphicsAnchorLayout_AddAnchors(QGraphicsAnchorLayout* self, QGraphicsLayoutItem* firstItem, QGraphicsLayoutItem* secondItem);
extern __declspec(dllexport) void QGraphicsAnchorLayout_SetHorizontalSpacing(QGraphicsAnchorLayout* self, double spacing);
extern __declspec(dllexport) void QGraphicsAnchorLayout_SetVerticalSpacing(QGraphicsAnchorLayout* self, double spacing);
extern __declspec(dllexport) void QGraphicsAnchorLayout_SetSpacing(QGraphicsAnchorLayout* self, double spacing);
extern __declspec(dllexport) double QGraphicsAnchorLayout_HorizontalSpacing(const QGraphicsAnchorLayout* self);
extern __declspec(dllexport) double QGraphicsAnchorLayout_VerticalSpacing(const QGraphicsAnchorLayout* self);
extern __declspec(dllexport) void QGraphicsAnchorLayout_RemoveAt(QGraphicsAnchorLayout* self, int index);
extern __declspec(dllexport) void QGraphicsAnchorLayout_SetGeometry(QGraphicsAnchorLayout* self, QRectF* rect);
extern __declspec(dllexport) int QGraphicsAnchorLayout_Count(const QGraphicsAnchorLayout* self);
extern __declspec(dllexport) QGraphicsLayoutItem* QGraphicsAnchorLayout_ItemAt(const QGraphicsAnchorLayout* self, int index);
extern __declspec(dllexport) void QGraphicsAnchorLayout_Invalidate(QGraphicsAnchorLayout* self);
extern __declspec(dllexport) QSizeF* QGraphicsAnchorLayout_SizeHint(const QGraphicsAnchorLayout* self, int which, QSizeF* constraint);
extern __declspec(dllexport) void QGraphicsAnchorLayout_AddAnchors3(QGraphicsAnchorLayout* self, QGraphicsLayoutItem* firstItem, QGraphicsLayoutItem* secondItem, int orientations);
extern __declspec(dllexport) void QGraphicsAnchorLayout_override_virtual_RemoveAt(void* self, intptr_t slot);
void QGraphicsAnchorLayout_virtualbase_RemoveAt(void* self, int index);
extern __declspec(dllexport) void QGraphicsAnchorLayout_override_virtual_SetGeometry(void* self, intptr_t slot);
void QGraphicsAnchorLayout_virtualbase_SetGeometry(void* self, QRectF* rect);
extern __declspec(dllexport) void QGraphicsAnchorLayout_override_virtual_Count(void* self, intptr_t slot);
int QGraphicsAnchorLayout_virtualbase_Count(const void* self);
extern __declspec(dllexport) void QGraphicsAnchorLayout_override_virtual_ItemAt(void* self, intptr_t slot);
QGraphicsLayoutItem* QGraphicsAnchorLayout_virtualbase_ItemAt(const void* self, int index);
extern __declspec(dllexport) void QGraphicsAnchorLayout_override_virtual_Invalidate(void* self, intptr_t slot);
void QGraphicsAnchorLayout_virtualbase_Invalidate(void* self);
extern __declspec(dllexport) void QGraphicsAnchorLayout_override_virtual_SizeHint(void* self, intptr_t slot);
QSizeF* QGraphicsAnchorLayout_virtualbase_SizeHint(const void* self, int which, QSizeF* constraint);
extern __declspec(dllexport) void QGraphicsAnchorLayout_override_virtual_GetContentsMargins(void* self, intptr_t slot);
void QGraphicsAnchorLayout_virtualbase_GetContentsMargins(const void* self, double* left, double* top, double* right, double* bottom);
extern __declspec(dllexport) void QGraphicsAnchorLayout_override_virtual_UpdateGeometry(void* self, intptr_t slot);
void QGraphicsAnchorLayout_virtualbase_UpdateGeometry(void* self);
extern __declspec(dllexport) void QGraphicsAnchorLayout_override_virtual_WidgetEvent(void* self, intptr_t slot);
void QGraphicsAnchorLayout_virtualbase_WidgetEvent(void* self, QEvent* e);
extern __declspec(dllexport) void QGraphicsAnchorLayout_Delete(QGraphicsAnchorLayout* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
