#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QGRAPHICSVIDEOITEM_H
#define MIQT_QT_MULTIMEDIA_GEN_QGRAPHICSVIDEOITEM_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QGraphicsItem QGraphicsItem;
typedef struct QGraphicsObject QGraphicsObject;
typedef struct QGraphicsVideoItem QGraphicsVideoItem;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPainter QPainter;
typedef struct QPointF QPointF;
typedef struct QRectF QRectF;
typedef struct QSizeF QSizeF;
typedef struct QStyleOptionGraphicsItem QStyleOptionGraphicsItem;
typedef struct QTimerEvent QTimerEvent;
typedef struct QVariant QVariant;
typedef struct QVideoSink QVideoSink;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QGraphicsVideoItem* QGraphicsVideoItem_new();
extern __declspec(dllexport) 
QGraphicsVideoItem* QGraphicsVideoItem_new2(QGraphicsItem* parent);
extern __declspec(dllexport) 
void QGraphicsVideoItem_virtbase(QGraphicsVideoItem* src
, QGraphicsObject** outptr_QGraphicsObject
);
extern __declspec(dllexport) 
QMetaObject* QGraphicsVideoItem_MetaObject(const QGraphicsVideoItem* self);
extern __declspec(dllexport) 
void* QGraphicsVideoItem_Metacast(QGraphicsVideoItem* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QGraphicsVideoItem_Tr(const char* s);
extern __declspec(dllexport) 
QVideoSink* QGraphicsVideoItem_VideoSink(const QGraphicsVideoItem* self);
extern __declspec(dllexport) 
int QGraphicsVideoItem_AspectRatioMode(const QGraphicsVideoItem* self);
extern __declspec(dllexport) 
void QGraphicsVideoItem_SetAspectRatioMode(QGraphicsVideoItem* self, int mode);
extern __declspec(dllexport) 
QPointF* QGraphicsVideoItem_Offset(const QGraphicsVideoItem* self);
extern __declspec(dllexport) 
void QGraphicsVideoItem_SetOffset(QGraphicsVideoItem* self, QPointF* offset);
extern __declspec(dllexport) 
QSizeF* QGraphicsVideoItem_Size(const QGraphicsVideoItem* self);
extern __declspec(dllexport) 
void QGraphicsVideoItem_SetSize(QGraphicsVideoItem* self, QSizeF* size);
extern __declspec(dllexport) 
QSizeF* QGraphicsVideoItem_NativeSize(const QGraphicsVideoItem* self);
extern __declspec(dllexport) 
QRectF* QGraphicsVideoItem_BoundingRect(const QGraphicsVideoItem* self);
extern __declspec(dllexport) 
void QGraphicsVideoItem_Paint(QGraphicsVideoItem* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget);
extern __declspec(dllexport) 
int QGraphicsVideoItem_Type(const QGraphicsVideoItem* self);
extern __declspec(dllexport) 
void QGraphicsVideoItem_NativeSizeChanged(QGraphicsVideoItem* self, QSizeF* size);
void QGraphicsVideoItem_connect_NativeSizeChanged(QGraphicsVideoItem* self, intptr_t slot);
extern __declspec(dllexport) 
void QGraphicsVideoItem_TimerEvent(QGraphicsVideoItem* self, QTimerEvent* event);
extern __declspec(dllexport) 
QVariant* QGraphicsVideoItem_ItemChange(QGraphicsVideoItem* self, GraphicsItemChange change, QVariant* value);
extern __declspec(dllexport) 
struct miqt_string QGraphicsVideoItem_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QGraphicsVideoItem_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QGraphicsVideoItem_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QGraphicsVideoItem_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QGraphicsVideoItem_override_virtual_Metacast(void* self, intptr_t slot);
void* QGraphicsVideoItem_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QGraphicsVideoItem_Delete(QGraphicsVideoItem* self, bool isSubclass);

}
