#pragma once
#ifndef MIQT_QT_GEN_QGRAPHICSSCENE_H
#define MIQT_QT_GEN_QGRAPHICSSCENE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QBrush QBrush;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QFont QFont;
typedef struct QGraphicsEllipseItem QGraphicsEllipseItem;
typedef struct QGraphicsItem QGraphicsItem;
typedef struct QGraphicsItemGroup QGraphicsItemGroup;
typedef struct QGraphicsLineItem QGraphicsLineItem;
typedef struct QGraphicsPathItem QGraphicsPathItem;
typedef struct QGraphicsPixmapItem QGraphicsPixmapItem;
typedef struct QGraphicsProxyWidget QGraphicsProxyWidget;
typedef struct QGraphicsRectItem QGraphicsRectItem;
typedef struct QGraphicsScene QGraphicsScene;
typedef struct QGraphicsSceneContextMenuEvent QGraphicsSceneContextMenuEvent;
typedef struct QGraphicsSceneDragDropEvent QGraphicsSceneDragDropEvent;
typedef struct QGraphicsSceneHelpEvent QGraphicsSceneHelpEvent;
typedef struct QGraphicsSceneMouseEvent QGraphicsSceneMouseEvent;
typedef struct QGraphicsSceneWheelEvent QGraphicsSceneWheelEvent;
typedef struct QGraphicsSimpleTextItem QGraphicsSimpleTextItem;
typedef struct QGraphicsTextItem QGraphicsTextItem;
typedef struct QGraphicsView QGraphicsView;
typedef struct QGraphicsWidget QGraphicsWidget;
typedef struct QInputMethodEvent QInputMethodEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QLineF QLineF;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPainter QPainter;
typedef struct QPainterPath QPainterPath;
typedef struct QPalette QPalette;
typedef struct QPen QPen;
typedef struct QPixmap QPixmap;
typedef struct QPointF QPointF;
typedef struct QRectF QRectF;
typedef struct QStyle QStyle;
typedef struct QTransform QTransform;
typedef struct QVariant QVariant;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QGraphicsScene* QGraphicsScene_new();
extern __declspec(dllexport) 
QGraphicsScene* QGraphicsScene_new2(QRectF* sceneRect);
extern __declspec(dllexport) 
QGraphicsScene* QGraphicsScene_new3(double x, double y, double width, double height);
extern __declspec(dllexport) 
QGraphicsScene* QGraphicsScene_new4(QObject* parent);
extern __declspec(dllexport) 
QGraphicsScene* QGraphicsScene_new5(QRectF* sceneRect, QObject* parent);
extern __declspec(dllexport) 
QGraphicsScene* QGraphicsScene_new6(double x, double y, double width, double height, QObject* parent);
extern __declspec(dllexport) 
void QGraphicsScene_virtbase(QGraphicsScene* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QGraphicsScene_MetaObject(const QGraphicsScene* self);
extern __declspec(dllexport) 
void* QGraphicsScene_Metacast(QGraphicsScene* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QGraphicsScene_Tr(const char* s);
extern __declspec(dllexport) 
QRectF* QGraphicsScene_SceneRect(const QGraphicsScene* self);
extern __declspec(dllexport) 
double QGraphicsScene_Width(const QGraphicsScene* self);
extern __declspec(dllexport) 
double QGraphicsScene_Height(const QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_SetSceneRect(QGraphicsScene* self, QRectF* rect);
extern __declspec(dllexport) 
void QGraphicsScene_SetSceneRect2(QGraphicsScene* self, double x, double y, double w, double h);
extern __declspec(dllexport) 
void QGraphicsScene_Render(QGraphicsScene* self, QPainter* painter);
extern __declspec(dllexport) 
ItemIndexMethod QGraphicsScene_ItemIndexMethod(const QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_SetItemIndexMethod(QGraphicsScene* self, ItemIndexMethod method);
extern __declspec(dllexport) 
int QGraphicsScene_BspTreeDepth(const QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_SetBspTreeDepth(QGraphicsScene* self, int depth);
extern __declspec(dllexport) 
QRectF* QGraphicsScene_ItemsBoundingRect(const QGraphicsScene* self);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_Items(const QGraphicsScene* self);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_ItemsWithPos(const QGraphicsScene* self, QPointF* pos);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_ItemsWithRect(const QGraphicsScene* self, QRectF* rect);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_ItemsWithPath(const QGraphicsScene* self, QPainterPath* path);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_Items2(const QGraphicsScene* self, double x, double y, double w, double h, int mode, int order);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_CollidingItems(const QGraphicsScene* self, QGraphicsItem* item);
extern __declspec(dllexport) 
QGraphicsItem* QGraphicsScene_ItemAt(const QGraphicsScene* self, QPointF* pos, QTransform* deviceTransform);
extern __declspec(dllexport) 
QGraphicsItem* QGraphicsScene_ItemAt2(const QGraphicsScene* self, double x, double y, QTransform* deviceTransform);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_SelectedItems(const QGraphicsScene* self);
extern __declspec(dllexport) 
QPainterPath* QGraphicsScene_SelectionArea(const QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_SetSelectionArea(QGraphicsScene* self, QPainterPath* path, QTransform* deviceTransform);
extern __declspec(dllexport) 
void QGraphicsScene_SetSelectionAreaWithPath(QGraphicsScene* self, QPainterPath* path);
extern __declspec(dllexport) 
QGraphicsItemGroup* QGraphicsScene_CreateItemGroup(QGraphicsScene* self, struct miqt_array /* of QGraphicsItem* */  items);
extern __declspec(dllexport) 
void QGraphicsScene_DestroyItemGroup(QGraphicsScene* self, QGraphicsItemGroup* group);
extern __declspec(dllexport) 
void QGraphicsScene_AddItem(QGraphicsScene* self, QGraphicsItem* item);
extern __declspec(dllexport) 
QGraphicsEllipseItem* QGraphicsScene_AddEllipse(QGraphicsScene* self, QRectF* rect);
extern __declspec(dllexport) 
QGraphicsLineItem* QGraphicsScene_AddLine(QGraphicsScene* self, QLineF* line);
extern __declspec(dllexport) 
QGraphicsPathItem* QGraphicsScene_AddPath(QGraphicsScene* self, QPainterPath* path);
extern __declspec(dllexport) 
QGraphicsPixmapItem* QGraphicsScene_AddPixmap(QGraphicsScene* self, QPixmap* pixmap);
extern __declspec(dllexport) 
QGraphicsRectItem* QGraphicsScene_AddRect(QGraphicsScene* self, QRectF* rect);
extern __declspec(dllexport) 
QGraphicsTextItem* QGraphicsScene_AddText(QGraphicsScene* self, struct miqt_string text);
extern __declspec(dllexport) 
QGraphicsSimpleTextItem* QGraphicsScene_AddSimpleText(QGraphicsScene* self, struct miqt_string text);
extern __declspec(dllexport) 
QGraphicsProxyWidget* QGraphicsScene_AddWidget(QGraphicsScene* self, QWidget* widget);
extern __declspec(dllexport) 
QGraphicsEllipseItem* QGraphicsScene_AddEllipse2(QGraphicsScene* self, double x, double y, double w, double h);
extern __declspec(dllexport) 
QGraphicsLineItem* QGraphicsScene_AddLine2(QGraphicsScene* self, double x1, double y1, double x2, double y2);
extern __declspec(dllexport) 
QGraphicsRectItem* QGraphicsScene_AddRect2(QGraphicsScene* self, double x, double y, double w, double h);
extern __declspec(dllexport) 
void QGraphicsScene_RemoveItem(QGraphicsScene* self, QGraphicsItem* item);
extern __declspec(dllexport) 
QGraphicsItem* QGraphicsScene_FocusItem(const QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_SetFocusItem(QGraphicsScene* self, QGraphicsItem* item);
extern __declspec(dllexport) 
bool QGraphicsScene_HasFocus(const QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_SetFocus(QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_ClearFocus(QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_SetStickyFocus(QGraphicsScene* self, bool enabled);
extern __declspec(dllexport) 
bool QGraphicsScene_StickyFocus(const QGraphicsScene* self);
extern __declspec(dllexport) 
QGraphicsItem* QGraphicsScene_MouseGrabberItem(const QGraphicsScene* self);
extern __declspec(dllexport) 
QBrush* QGraphicsScene_BackgroundBrush(const QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_SetBackgroundBrush(QGraphicsScene* self, QBrush* brush);
extern __declspec(dllexport) 
QBrush* QGraphicsScene_ForegroundBrush(const QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_SetForegroundBrush(QGraphicsScene* self, QBrush* brush);
extern __declspec(dllexport) 
QVariant* QGraphicsScene_InputMethodQuery(const QGraphicsScene* self, int query);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsView* */  QGraphicsScene_Views(const QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_Update(QGraphicsScene* self, double x, double y, double w, double h);
extern __declspec(dllexport) 
void QGraphicsScene_Invalidate(QGraphicsScene* self, double x, double y, double w, double h);
extern __declspec(dllexport) 
QStyle* QGraphicsScene_Style(const QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_SetStyle(QGraphicsScene* self, QStyle* style);
extern __declspec(dllexport) 
QFont* QGraphicsScene_Font(const QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_SetFont(QGraphicsScene* self, QFont* font);
extern __declspec(dllexport) 
QPalette* QGraphicsScene_Palette(const QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_SetPalette(QGraphicsScene* self, QPalette* palette);
extern __declspec(dllexport) 
bool QGraphicsScene_IsActive(const QGraphicsScene* self);
extern __declspec(dllexport) 
QGraphicsItem* QGraphicsScene_ActivePanel(const QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_SetActivePanel(QGraphicsScene* self, QGraphicsItem* item);
extern __declspec(dllexport) 
QGraphicsWidget* QGraphicsScene_ActiveWindow(const QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_SetActiveWindow(QGraphicsScene* self, QGraphicsWidget* widget);
extern __declspec(dllexport) 
bool QGraphicsScene_SendEvent(QGraphicsScene* self, QGraphicsItem* item, QEvent* event);
extern __declspec(dllexport) 
double QGraphicsScene_MinimumRenderSize(const QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_SetMinimumRenderSize(QGraphicsScene* self, double minSize);
extern __declspec(dllexport) 
bool QGraphicsScene_FocusOnTouch(const QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_SetFocusOnTouch(QGraphicsScene* self, bool enabled);
extern __declspec(dllexport) 
void QGraphicsScene_Update2(QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_Invalidate2(QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_Advance(QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_ClearSelection(QGraphicsScene* self);
extern __declspec(dllexport) 
void QGraphicsScene_Clear(QGraphicsScene* self);
extern __declspec(dllexport) 
bool QGraphicsScene_Event(QGraphicsScene* self, QEvent* event);
extern __declspec(dllexport) 
bool QGraphicsScene_EventFilter(QGraphicsScene* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) 
void QGraphicsScene_ContextMenuEvent(QGraphicsScene* self, QGraphicsSceneContextMenuEvent* event);
extern __declspec(dllexport) 
void QGraphicsScene_DragEnterEvent(QGraphicsScene* self, QGraphicsSceneDragDropEvent* event);
extern __declspec(dllexport) 
void QGraphicsScene_DragMoveEvent(QGraphicsScene* self, QGraphicsSceneDragDropEvent* event);
extern __declspec(dllexport) 
void QGraphicsScene_DragLeaveEvent(QGraphicsScene* self, QGraphicsSceneDragDropEvent* event);
extern __declspec(dllexport) 
void QGraphicsScene_DropEvent(QGraphicsScene* self, QGraphicsSceneDragDropEvent* event);
extern __declspec(dllexport) 
void QGraphicsScene_FocusInEvent(QGraphicsScene* self, QFocusEvent* event);
extern __declspec(dllexport) 
void QGraphicsScene_FocusOutEvent(QGraphicsScene* self, QFocusEvent* event);
extern __declspec(dllexport) 
void QGraphicsScene_HelpEvent(QGraphicsScene* self, QGraphicsSceneHelpEvent* event);
extern __declspec(dllexport) 
void QGraphicsScene_KeyPressEvent(QGraphicsScene* self, QKeyEvent* event);
extern __declspec(dllexport) 
void QGraphicsScene_KeyReleaseEvent(QGraphicsScene* self, QKeyEvent* event);
extern __declspec(dllexport) 
void QGraphicsScene_MousePressEvent(QGraphicsScene* self, QGraphicsSceneMouseEvent* event);
extern __declspec(dllexport) 
void QGraphicsScene_MouseMoveEvent(QGraphicsScene* self, QGraphicsSceneMouseEvent* event);
extern __declspec(dllexport) 
void QGraphicsScene_MouseReleaseEvent(QGraphicsScene* self, QGraphicsSceneMouseEvent* event);
extern __declspec(dllexport) 
void QGraphicsScene_MouseDoubleClickEvent(QGraphicsScene* self, QGraphicsSceneMouseEvent* event);
extern __declspec(dllexport) 
void QGraphicsScene_WheelEvent(QGraphicsScene* self, QGraphicsSceneWheelEvent* event);
extern __declspec(dllexport) 
void QGraphicsScene_InputMethodEvent(QGraphicsScene* self, QInputMethodEvent* event);
extern __declspec(dllexport) 
void QGraphicsScene_DrawBackground(QGraphicsScene* self, QPainter* painter, QRectF* rect);
extern __declspec(dllexport) 
void QGraphicsScene_DrawForeground(QGraphicsScene* self, QPainter* painter, QRectF* rect);
extern __declspec(dllexport) 
bool QGraphicsScene_FocusNextPrevChild(QGraphicsScene* self, bool next);
extern __declspec(dllexport) 
void QGraphicsScene_Changed(QGraphicsScene* self, struct miqt_array /* of QRectF* */  region);
void QGraphicsScene_connect_Changed(QGraphicsScene* self, intptr_t slot);
extern __declspec(dllexport) 
void QGraphicsScene_SceneRectChanged(QGraphicsScene* self, QRectF* rect);
void QGraphicsScene_connect_SceneRectChanged(QGraphicsScene* self, intptr_t slot);
extern __declspec(dllexport) 
void QGraphicsScene_SelectionChanged(QGraphicsScene* self);
void QGraphicsScene_connect_SelectionChanged(QGraphicsScene* self, intptr_t slot);
extern __declspec(dllexport) 
void QGraphicsScene_FocusItemChanged(QGraphicsScene* self, QGraphicsItem* newFocus, QGraphicsItem* oldFocus, int reason);
void QGraphicsScene_connect_FocusItemChanged(QGraphicsScene* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QGraphicsScene_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QGraphicsScene_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QGraphicsScene_Render2(QGraphicsScene* self, QPainter* painter, QRectF* target);
extern __declspec(dllexport) 
void QGraphicsScene_Render3(QGraphicsScene* self, QPainter* painter, QRectF* target, QRectF* source);
extern __declspec(dllexport) 
void QGraphicsScene_Render4(QGraphicsScene* self, QPainter* painter, QRectF* target, QRectF* source, int aspectRatioMode);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_Items1(const QGraphicsScene* self, int order);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_Items22(const QGraphicsScene* self, QPointF* pos, int mode);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_Items3(const QGraphicsScene* self, QPointF* pos, int mode, int order);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_Items4(const QGraphicsScene* self, QPointF* pos, int mode, int order, QTransform* deviceTransform);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_Items23(const QGraphicsScene* self, QRectF* rect, int mode);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_Items32(const QGraphicsScene* self, QRectF* rect, int mode, int order);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_Items42(const QGraphicsScene* self, QRectF* rect, int mode, int order, QTransform* deviceTransform);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_Items25(const QGraphicsScene* self, QPainterPath* path, int mode);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_Items34(const QGraphicsScene* self, QPainterPath* path, int mode, int order);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_Items44(const QGraphicsScene* self, QPainterPath* path, int mode, int order, QTransform* deviceTransform);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_Items7(const QGraphicsScene* self, double x, double y, double w, double h, int mode, int order, QTransform* deviceTransform);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsScene_CollidingItems2(const QGraphicsScene* self, QGraphicsItem* item, int mode);
extern __declspec(dllexport) 
void QGraphicsScene_SetSelectionArea2(QGraphicsScene* self, QPainterPath* path, int selectionOperation);
extern __declspec(dllexport) 
void QGraphicsScene_SetSelectionArea3(QGraphicsScene* self, QPainterPath* path, int selectionOperation, int mode);
extern __declspec(dllexport) 
void QGraphicsScene_SetSelectionArea4(QGraphicsScene* self, QPainterPath* path, int selectionOperation, int mode, QTransform* deviceTransform);
extern __declspec(dllexport) 
QGraphicsEllipseItem* QGraphicsScene_AddEllipse22(QGraphicsScene* self, QRectF* rect, QPen* pen);
extern __declspec(dllexport) 
QGraphicsEllipseItem* QGraphicsScene_AddEllipse3(QGraphicsScene* self, QRectF* rect, QPen* pen, QBrush* brush);
extern __declspec(dllexport) 
QGraphicsLineItem* QGraphicsScene_AddLine22(QGraphicsScene* self, QLineF* line, QPen* pen);
extern __declspec(dllexport) 
QGraphicsPathItem* QGraphicsScene_AddPath2(QGraphicsScene* self, QPainterPath* path, QPen* pen);
extern __declspec(dllexport) 
QGraphicsPathItem* QGraphicsScene_AddPath3(QGraphicsScene* self, QPainterPath* path, QPen* pen, QBrush* brush);
extern __declspec(dllexport) 
QGraphicsRectItem* QGraphicsScene_AddRect22(QGraphicsScene* self, QRectF* rect, QPen* pen);
extern __declspec(dllexport) 
QGraphicsRectItem* QGraphicsScene_AddRect3(QGraphicsScene* self, QRectF* rect, QPen* pen, QBrush* brush);
extern __declspec(dllexport) 
QGraphicsTextItem* QGraphicsScene_AddText2(QGraphicsScene* self, struct miqt_string text, QFont* font);
extern __declspec(dllexport) 
QGraphicsSimpleTextItem* QGraphicsScene_AddSimpleText2(QGraphicsScene* self, struct miqt_string text, QFont* font);
extern __declspec(dllexport) 
QGraphicsProxyWidget* QGraphicsScene_AddWidget2(QGraphicsScene* self, QWidget* widget, int wFlags);
extern __declspec(dllexport) 
QGraphicsEllipseItem* QGraphicsScene_AddEllipse5(QGraphicsScene* self, double x, double y, double w, double h, QPen* pen);
extern __declspec(dllexport) 
QGraphicsEllipseItem* QGraphicsScene_AddEllipse6(QGraphicsScene* self, double x, double y, double w, double h, QPen* pen, QBrush* brush);
extern __declspec(dllexport) 
QGraphicsLineItem* QGraphicsScene_AddLine5(QGraphicsScene* self, double x1, double y1, double x2, double y2, QPen* pen);
extern __declspec(dllexport) 
QGraphicsRectItem* QGraphicsScene_AddRect5(QGraphicsScene* self, double x, double y, double w, double h, QPen* pen);
extern __declspec(dllexport) 
QGraphicsRectItem* QGraphicsScene_AddRect6(QGraphicsScene* self, double x, double y, double w, double h, QPen* pen, QBrush* brush);
extern __declspec(dllexport) 
void QGraphicsScene_SetFocusItem2(QGraphicsScene* self, QGraphicsItem* item, int focusReason);
extern __declspec(dllexport) 
void QGraphicsScene_SetFocus1(QGraphicsScene* self, int focusReason);
extern __declspec(dllexport) 
void QGraphicsScene_Invalidate5(QGraphicsScene* self, double x, double y, double w, double h, SceneLayers layers);
extern __declspec(dllexport) 
void QGraphicsScene_Update1(QGraphicsScene* self, QRectF* rect);
extern __declspec(dllexport) 
void QGraphicsScene_Invalidate1(QGraphicsScene* self, QRectF* rect);
extern __declspec(dllexport) 
void QGraphicsScene_Invalidate22(QGraphicsScene* self, QRectF* rect, SceneLayers layers);
extern __declspec(dllexport) 
void QGraphicsScene_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QGraphicsScene_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QGraphicsScene_override_virtual_Metacast(void* self, intptr_t slot);
void* QGraphicsScene_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QGraphicsScene_Delete(QGraphicsScene* self, bool isSubclass);

}
