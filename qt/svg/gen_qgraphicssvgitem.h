#pragma once
#ifndef MIQT_QT_SVG_GEN_QGRAPHICSSVGITEM_H
#define MIQT_QT_SVG_GEN_QGRAPHICSSVGITEM_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QGraphicsItem QGraphicsItem;
typedef struct QGraphicsObject QGraphicsObject;
typedef struct QGraphicsSvgItem QGraphicsSvgItem;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPainter QPainter;
typedef struct QRectF QRectF;
typedef struct QSize QSize;
typedef struct QStyleOptionGraphicsItem QStyleOptionGraphicsItem;
typedef struct QSvgRenderer QSvgRenderer;
typedef struct QWidget QWidget;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsSvgItem* QGraphicsSvgItem_new();
extern __declspec(dllexport) QGraphicsSvgItem* QGraphicsSvgItem_new2(struct miqt_string fileName);
extern __declspec(dllexport) QGraphicsSvgItem* QGraphicsSvgItem_new3(QGraphicsItem* parentItem);
extern __declspec(dllexport) QGraphicsSvgItem* QGraphicsSvgItem_new4(struct miqt_string fileName, QGraphicsItem* parentItem);
extern __declspec(dllexport) void QGraphicsSvgItem_virtbase(QGraphicsSvgItem* src, QGraphicsObject** outptr_QGraphicsObject);
extern __declspec(dllexport) QMetaObject* QGraphicsSvgItem_MetaObject(const QGraphicsSvgItem* self);
extern __declspec(dllexport) void* QGraphicsSvgItem_Metacast(QGraphicsSvgItem* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QGraphicsSvgItem_Tr(const char* s);
extern __declspec(dllexport) void QGraphicsSvgItem_SetSharedRenderer(QGraphicsSvgItem* self, QSvgRenderer* renderer);
extern __declspec(dllexport) QSvgRenderer* QGraphicsSvgItem_Renderer(const QGraphicsSvgItem* self);
extern __declspec(dllexport) void QGraphicsSvgItem_SetElementId(QGraphicsSvgItem* self, struct miqt_string id);
extern __declspec(dllexport) struct miqt_string QGraphicsSvgItem_ElementId(const QGraphicsSvgItem* self);
extern __declspec(dllexport) void QGraphicsSvgItem_SetCachingEnabled(QGraphicsSvgItem* self, bool cachingEnabled);
extern __declspec(dllexport) bool QGraphicsSvgItem_IsCachingEnabled(const QGraphicsSvgItem* self);
extern __declspec(dllexport) void QGraphicsSvgItem_SetMaximumCacheSize(QGraphicsSvgItem* self, QSize* size);
extern __declspec(dllexport) QSize* QGraphicsSvgItem_MaximumCacheSize(const QGraphicsSvgItem* self);
extern __declspec(dllexport) QRectF* QGraphicsSvgItem_BoundingRect(const QGraphicsSvgItem* self);
extern __declspec(dllexport) void QGraphicsSvgItem_Paint(QGraphicsSvgItem* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget);
extern __declspec(dllexport) int QGraphicsSvgItem_Type(const QGraphicsSvgItem* self);
extern __declspec(dllexport) struct miqt_string QGraphicsSvgItem_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QGraphicsSvgItem_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QGraphicsSvgItem_override_virtual_BoundingRect(void* self, intptr_t slot);
QRectF* QGraphicsSvgItem_virtualbase_BoundingRect(const void* self);
extern __declspec(dllexport) void QGraphicsSvgItem_override_virtual_Paint(void* self, intptr_t slot);
void QGraphicsSvgItem_virtualbase_Paint(void* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget);
extern __declspec(dllexport) void QGraphicsSvgItem_override_virtual_Type(void* self, intptr_t slot);
int QGraphicsSvgItem_virtualbase_Type(const void* self);
extern __declspec(dllexport) void QGraphicsSvgItem_override_virtual_Event(void* self, intptr_t slot);
bool QGraphicsSvgItem_virtualbase_Event(void* self, QEvent* ev);
extern __declspec(dllexport) void QGraphicsSvgItem_Delete(QGraphicsSvgItem* self, bool isSubclass);

} 
