#pragma once
#ifndef MIQT_QT_GEN_QGRAPHICSPROXYWIDGET_H
#define MIQT_QT_GEN_QGRAPHICSPROXYWIDGET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QGraphicsItem QGraphicsItem;
typedef struct QGraphicsLayoutItem QGraphicsLayoutItem;
typedef struct QGraphicsObject QGraphicsObject;
typedef struct QGraphicsProxyWidget QGraphicsProxyWidget;
typedef struct QGraphicsSceneContextMenuEvent QGraphicsSceneContextMenuEvent;
typedef struct QGraphicsSceneDragDropEvent QGraphicsSceneDragDropEvent;
typedef struct QGraphicsSceneHoverEvent QGraphicsSceneHoverEvent;
typedef struct QGraphicsSceneMouseEvent QGraphicsSceneMouseEvent;
typedef struct QGraphicsSceneResizeEvent QGraphicsSceneResizeEvent;
typedef struct QGraphicsSceneWheelEvent QGraphicsSceneWheelEvent;
typedef struct QGraphicsWidget QGraphicsWidget;
typedef struct QHideEvent QHideEvent;
typedef struct QInputMethodEvent QInputMethodEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPainter QPainter;
typedef struct QRectF QRectF;
typedef struct QShowEvent QShowEvent;
typedef struct QSizeF QSizeF;
typedef struct QStyleOptionGraphicsItem QStyleOptionGraphicsItem;
typedef struct QVariant QVariant;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QGraphicsProxyWidget* QGraphicsProxyWidget_new();
extern __declspec(dllexport) 
QGraphicsProxyWidget* QGraphicsProxyWidget_new2(QGraphicsItem* parent);
extern __declspec(dllexport) 
QGraphicsProxyWidget* QGraphicsProxyWidget_new3(QGraphicsItem* parent, int wFlags);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_virtbase(QGraphicsProxyWidget* src
, QGraphicsWidget** outptr_QGraphicsWidget
);
extern __declspec(dllexport) 
QMetaObject* QGraphicsProxyWidget_MetaObject(const QGraphicsProxyWidget* self);
extern __declspec(dllexport) 
void* QGraphicsProxyWidget_Metacast(QGraphicsProxyWidget* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QGraphicsProxyWidget_Tr(const char* s);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_SetWidget(QGraphicsProxyWidget* self, QWidget* widget);
extern __declspec(dllexport) 
QWidget* QGraphicsProxyWidget_Widget(const QGraphicsProxyWidget* self);
extern __declspec(dllexport) 
QRectF* QGraphicsProxyWidget_SubWidgetRect(const QGraphicsProxyWidget* self, QWidget* widget);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_SetGeometry(QGraphicsProxyWidget* self, QRectF* rect);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_Paint(QGraphicsProxyWidget* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget);
extern __declspec(dllexport) 
int QGraphicsProxyWidget_Type(const QGraphicsProxyWidget* self);
extern __declspec(dllexport) 
QGraphicsProxyWidget* QGraphicsProxyWidget_CreateProxyForChildWidget(QGraphicsProxyWidget* self, QWidget* child);
extern __declspec(dllexport) 
QVariant* QGraphicsProxyWidget_ItemChange(QGraphicsProxyWidget* self, GraphicsItemChange change, QVariant* value);
extern __declspec(dllexport) 
bool QGraphicsProxyWidget_Event(QGraphicsProxyWidget* self, QEvent* event);
extern __declspec(dllexport) 
bool QGraphicsProxyWidget_EventFilter(QGraphicsProxyWidget* self, QObject* object, QEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_ShowEvent(QGraphicsProxyWidget* self, QShowEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_HideEvent(QGraphicsProxyWidget* self, QHideEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_ContextMenuEvent(QGraphicsProxyWidget* self, QGraphicsSceneContextMenuEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_DragEnterEvent(QGraphicsProxyWidget* self, QGraphicsSceneDragDropEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_DragLeaveEvent(QGraphicsProxyWidget* self, QGraphicsSceneDragDropEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_DragMoveEvent(QGraphicsProxyWidget* self, QGraphicsSceneDragDropEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_DropEvent(QGraphicsProxyWidget* self, QGraphicsSceneDragDropEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_HoverEnterEvent(QGraphicsProxyWidget* self, QGraphicsSceneHoverEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_HoverLeaveEvent(QGraphicsProxyWidget* self, QGraphicsSceneHoverEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_HoverMoveEvent(QGraphicsProxyWidget* self, QGraphicsSceneHoverEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_GrabMouseEvent(QGraphicsProxyWidget* self, QEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_UngrabMouseEvent(QGraphicsProxyWidget* self, QEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_MouseMoveEvent(QGraphicsProxyWidget* self, QGraphicsSceneMouseEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_MousePressEvent(QGraphicsProxyWidget* self, QGraphicsSceneMouseEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_MouseReleaseEvent(QGraphicsProxyWidget* self, QGraphicsSceneMouseEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_MouseDoubleClickEvent(QGraphicsProxyWidget* self, QGraphicsSceneMouseEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_WheelEvent(QGraphicsProxyWidget* self, QGraphicsSceneWheelEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_KeyPressEvent(QGraphicsProxyWidget* self, QKeyEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_KeyReleaseEvent(QGraphicsProxyWidget* self, QKeyEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_FocusInEvent(QGraphicsProxyWidget* self, QFocusEvent* event);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_FocusOutEvent(QGraphicsProxyWidget* self, QFocusEvent* event);
extern __declspec(dllexport) 
bool QGraphicsProxyWidget_FocusNextPrevChild(QGraphicsProxyWidget* self, bool next);
extern __declspec(dllexport) 
QVariant* QGraphicsProxyWidget_InputMethodQuery(const QGraphicsProxyWidget* self, int query);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_InputMethodEvent(QGraphicsProxyWidget* self, QInputMethodEvent* event);
extern __declspec(dllexport) 
QSizeF* QGraphicsProxyWidget_SizeHint(const QGraphicsProxyWidget* self, int which, QSizeF* constraint);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_ResizeEvent(QGraphicsProxyWidget* self, QGraphicsSceneResizeEvent* event);
extern __declspec(dllexport) 
struct miqt_string QGraphicsProxyWidget_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QGraphicsProxyWidget_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QGraphicsProxyWidget_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_override_virtual_Metacast(void* self, intptr_t slot);
void* QGraphicsProxyWidget_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QGraphicsProxyWidget_Delete(QGraphicsProxyWidget* self, bool isSubclass);

}
