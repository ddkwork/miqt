#pragma once
#ifndef MIQT_QT_GEN_QGRAPHICSVIEW_H
#define MIQT_QT_GEN_QGRAPHICSVIEW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractScrollArea QAbstractScrollArea;
typedef struct QBrush QBrush;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDragEnterEvent QDragEnterEvent;
typedef struct QDragLeaveEvent QDragLeaveEvent;
typedef struct QDragMoveEvent QDragMoveEvent;
typedef struct QDropEvent QDropEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QFrame QFrame;
typedef struct QGraphicsItem QGraphicsItem;
typedef struct QGraphicsScene QGraphicsScene;
typedef struct QGraphicsView QGraphicsView;
typedef struct QInputMethodEvent QInputMethodEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPainter QPainter;
typedef struct QPainterPath QPainterPath;
typedef struct QPoint QPoint;
typedef struct QPointF QPointF;
typedef struct QRect QRect;
typedef struct QRectF QRectF;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QTransform QTransform;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QGraphicsView* QGraphicsView_new(QWidget* parent);
extern __declspec(dllexport) 
QGraphicsView* QGraphicsView_new2();
extern __declspec(dllexport) 
QGraphicsView* QGraphicsView_new3(QGraphicsScene* scene);
extern __declspec(dllexport) 
QGraphicsView* QGraphicsView_new4(QGraphicsScene* scene, QWidget* parent);
extern __declspec(dllexport) 
void QGraphicsView_virtbase(QGraphicsView* src
, QAbstractScrollArea** outptr_QAbstractScrollArea
);
extern __declspec(dllexport) 
QMetaObject* QGraphicsView_MetaObject(const QGraphicsView* self);
extern __declspec(dllexport) 
void* QGraphicsView_Metacast(QGraphicsView* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QGraphicsView_Tr(const char* s);
extern __declspec(dllexport) 
QSize* QGraphicsView_SizeHint(const QGraphicsView* self);
extern __declspec(dllexport) 
int QGraphicsView_RenderHints(const QGraphicsView* self);
extern __declspec(dllexport) 
void QGraphicsView_SetRenderHint(QGraphicsView* self, int hint);
extern __declspec(dllexport) 
void QGraphicsView_SetRenderHints(QGraphicsView* self, int hints);
extern __declspec(dllexport) 
int QGraphicsView_Alignment(const QGraphicsView* self);
extern __declspec(dllexport) 
void QGraphicsView_SetAlignment(QGraphicsView* self, int alignment);
extern __declspec(dllexport) 
ViewportAnchor QGraphicsView_TransformationAnchor(const QGraphicsView* self);
extern __declspec(dllexport) 
void QGraphicsView_SetTransformationAnchor(QGraphicsView* self, ViewportAnchor anchor);
extern __declspec(dllexport) 
ViewportAnchor QGraphicsView_ResizeAnchor(const QGraphicsView* self);
extern __declspec(dllexport) 
void QGraphicsView_SetResizeAnchor(QGraphicsView* self, ViewportAnchor anchor);
extern __declspec(dllexport) 
ViewportUpdateMode QGraphicsView_ViewportUpdateMode(const QGraphicsView* self);
extern __declspec(dllexport) 
void QGraphicsView_SetViewportUpdateMode(QGraphicsView* self, ViewportUpdateMode mode);
extern __declspec(dllexport) 
OptimizationFlags QGraphicsView_OptimizationFlags(const QGraphicsView* self);
extern __declspec(dllexport) 
void QGraphicsView_SetOptimizationFlag(QGraphicsView* self, OptimizationFlag flag);
extern __declspec(dllexport) 
void QGraphicsView_SetOptimizationFlags(QGraphicsView* self, OptimizationFlags flags);
extern __declspec(dllexport) 
DragMode QGraphicsView_DragMode(const QGraphicsView* self);
extern __declspec(dllexport) 
void QGraphicsView_SetDragMode(QGraphicsView* self, DragMode mode);
extern __declspec(dllexport) 
int QGraphicsView_RubberBandSelectionMode(const QGraphicsView* self);
extern __declspec(dllexport) 
void QGraphicsView_SetRubberBandSelectionMode(QGraphicsView* self, int mode);
extern __declspec(dllexport) 
QRect* QGraphicsView_RubberBandRect(const QGraphicsView* self);
extern __declspec(dllexport) 
CacheMode QGraphicsView_CacheMode(const QGraphicsView* self);
extern __declspec(dllexport) 
void QGraphicsView_SetCacheMode(QGraphicsView* self, CacheMode mode);
extern __declspec(dllexport) 
void QGraphicsView_ResetCachedContent(QGraphicsView* self);
extern __declspec(dllexport) 
bool QGraphicsView_IsInteractive(const QGraphicsView* self);
extern __declspec(dllexport) 
void QGraphicsView_SetInteractive(QGraphicsView* self, bool allowed);
extern __declspec(dllexport) 
QGraphicsScene* QGraphicsView_Scene(const QGraphicsView* self);
extern __declspec(dllexport) 
void QGraphicsView_SetScene(QGraphicsView* self, QGraphicsScene* scene);
extern __declspec(dllexport) 
QRectF* QGraphicsView_SceneRect(const QGraphicsView* self);
extern __declspec(dllexport) 
void QGraphicsView_SetSceneRect(QGraphicsView* self, QRectF* rect);
extern __declspec(dllexport) 
void QGraphicsView_SetSceneRect2(QGraphicsView* self, double x, double y, double w, double h);
extern __declspec(dllexport) 
QTransform* QGraphicsView_Transform(const QGraphicsView* self);
extern __declspec(dllexport) 
QTransform* QGraphicsView_ViewportTransform(const QGraphicsView* self);
extern __declspec(dllexport) 
bool QGraphicsView_IsTransformed(const QGraphicsView* self);
extern __declspec(dllexport) 
void QGraphicsView_SetTransform(QGraphicsView* self, QTransform* matrix);
extern __declspec(dllexport) 
void QGraphicsView_ResetTransform(QGraphicsView* self);
extern __declspec(dllexport) 
void QGraphicsView_Rotate(QGraphicsView* self, double angle);
extern __declspec(dllexport) 
void QGraphicsView_Scale(QGraphicsView* self, double sx, double sy);
extern __declspec(dllexport) 
void QGraphicsView_Shear(QGraphicsView* self, double sh, double sv);
extern __declspec(dllexport) 
void QGraphicsView_Translate(QGraphicsView* self, double dx, double dy);
extern __declspec(dllexport) 
void QGraphicsView_CenterOn(QGraphicsView* self, QPointF* pos);
extern __declspec(dllexport) 
void QGraphicsView_CenterOn2(QGraphicsView* self, double x, double y);
extern __declspec(dllexport) 
void QGraphicsView_CenterOnWithItem(QGraphicsView* self, QGraphicsItem* item);
extern __declspec(dllexport) 
void QGraphicsView_EnsureVisible(QGraphicsView* self, QRectF* rect);
extern __declspec(dllexport) 
void QGraphicsView_EnsureVisible2(QGraphicsView* self, double x, double y, double w, double h);
extern __declspec(dllexport) 
void QGraphicsView_EnsureVisibleWithItem(QGraphicsView* self, QGraphicsItem* item);
extern __declspec(dllexport) 
void QGraphicsView_FitInView(QGraphicsView* self, QRectF* rect);
extern __declspec(dllexport) 
void QGraphicsView_FitInView2(QGraphicsView* self, double x, double y, double w, double h);
extern __declspec(dllexport) 
void QGraphicsView_FitInViewWithItem(QGraphicsView* self, QGraphicsItem* item);
extern __declspec(dllexport) 
void QGraphicsView_Render(QGraphicsView* self, QPainter* painter);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsView_Items(const QGraphicsView* self);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsView_ItemsWithPos(const QGraphicsView* self, QPoint* pos);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsView_Items2(const QGraphicsView* self, int x, int y);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsView_ItemsWithRect(const QGraphicsView* self, QRect* rect);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsView_Items3(const QGraphicsView* self, int x, int y, int w, int h);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsView_ItemsWithPath(const QGraphicsView* self, QPainterPath* path);
extern __declspec(dllexport) 
QGraphicsItem* QGraphicsView_ItemAt(const QGraphicsView* self, QPoint* pos);
extern __declspec(dllexport) 
QGraphicsItem* QGraphicsView_ItemAt2(const QGraphicsView* self, int x, int y);
extern __declspec(dllexport) 
QPointF* QGraphicsView_MapToScene(const QGraphicsView* self, QPoint* point);
extern __declspec(dllexport) 
QPainterPath* QGraphicsView_MapToSceneWithPath(const QGraphicsView* self, QPainterPath* path);
extern __declspec(dllexport) 
QPoint* QGraphicsView_MapFromScene(const QGraphicsView* self, QPointF* point);
extern __declspec(dllexport) 
QPainterPath* QGraphicsView_MapFromSceneWithPath(const QGraphicsView* self, QPainterPath* path);
extern __declspec(dllexport) 
QPointF* QGraphicsView_MapToScene2(const QGraphicsView* self, int x, int y);
extern __declspec(dllexport) 
QPoint* QGraphicsView_MapFromScene2(const QGraphicsView* self, double x, double y);
extern __declspec(dllexport) 
QVariant* QGraphicsView_InputMethodQuery(const QGraphicsView* self, int query);
extern __declspec(dllexport) 
QBrush* QGraphicsView_BackgroundBrush(const QGraphicsView* self);
extern __declspec(dllexport) 
void QGraphicsView_SetBackgroundBrush(QGraphicsView* self, QBrush* brush);
extern __declspec(dllexport) 
QBrush* QGraphicsView_ForegroundBrush(const QGraphicsView* self);
extern __declspec(dllexport) 
void QGraphicsView_SetForegroundBrush(QGraphicsView* self, QBrush* brush);
extern __declspec(dllexport) 
void QGraphicsView_UpdateScene(QGraphicsView* self, struct miqt_array /* of QRectF* */  rects);
extern __declspec(dllexport) 
void QGraphicsView_InvalidateScene(QGraphicsView* self);
extern __declspec(dllexport) 
void QGraphicsView_UpdateSceneRect(QGraphicsView* self, QRectF* rect);
extern __declspec(dllexport) 
void QGraphicsView_RubberBandChanged(QGraphicsView* self, QRect* viewportRect, QPointF* fromScenePoint, QPointF* toScenePoint);
void QGraphicsView_connect_RubberBandChanged(QGraphicsView* self, intptr_t slot);
extern __declspec(dllexport) 
void QGraphicsView_SetupViewport(QGraphicsView* self, QWidget* widget);
extern __declspec(dllexport) 
bool QGraphicsView_Event(QGraphicsView* self, QEvent* event);
extern __declspec(dllexport) 
bool QGraphicsView_ViewportEvent(QGraphicsView* self, QEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_ContextMenuEvent(QGraphicsView* self, QContextMenuEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_DragEnterEvent(QGraphicsView* self, QDragEnterEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_DragLeaveEvent(QGraphicsView* self, QDragLeaveEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_DragMoveEvent(QGraphicsView* self, QDragMoveEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_DropEvent(QGraphicsView* self, QDropEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_FocusInEvent(QGraphicsView* self, QFocusEvent* event);
extern __declspec(dllexport) 
bool QGraphicsView_FocusNextPrevChild(QGraphicsView* self, bool next);
extern __declspec(dllexport) 
void QGraphicsView_FocusOutEvent(QGraphicsView* self, QFocusEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_KeyPressEvent(QGraphicsView* self, QKeyEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_KeyReleaseEvent(QGraphicsView* self, QKeyEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_MouseDoubleClickEvent(QGraphicsView* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_MousePressEvent(QGraphicsView* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_MouseMoveEvent(QGraphicsView* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_MouseReleaseEvent(QGraphicsView* self, QMouseEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_WheelEvent(QGraphicsView* self, QWheelEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_PaintEvent(QGraphicsView* self, QPaintEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_ResizeEvent(QGraphicsView* self, QResizeEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_ScrollContentsBy(QGraphicsView* self, int dx, int dy);
extern __declspec(dllexport) 
void QGraphicsView_ShowEvent(QGraphicsView* self, QShowEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_InputMethodEvent(QGraphicsView* self, QInputMethodEvent* event);
extern __declspec(dllexport) 
void QGraphicsView_DrawBackground(QGraphicsView* self, QPainter* painter, QRectF* rect);
extern __declspec(dllexport) 
void QGraphicsView_DrawForeground(QGraphicsView* self, QPainter* painter, QRectF* rect);
extern __declspec(dllexport) 
struct miqt_string QGraphicsView_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QGraphicsView_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QGraphicsView_SetRenderHint2(QGraphicsView* self, int hint, bool enabled);
extern __declspec(dllexport) 
void QGraphicsView_SetOptimizationFlag2(QGraphicsView* self, OptimizationFlag flag, bool enabled);
extern __declspec(dllexport) 
void QGraphicsView_SetTransform2(QGraphicsView* self, QTransform* matrix, bool combine);
extern __declspec(dllexport) 
void QGraphicsView_EnsureVisible22(QGraphicsView* self, QRectF* rect, int xmargin);
extern __declspec(dllexport) 
void QGraphicsView_EnsureVisible3(QGraphicsView* self, QRectF* rect, int xmargin, int ymargin);
extern __declspec(dllexport) 
void QGraphicsView_EnsureVisible5(QGraphicsView* self, double x, double y, double w, double h, int xmargin);
extern __declspec(dllexport) 
void QGraphicsView_EnsureVisible6(QGraphicsView* self, double x, double y, double w, double h, int xmargin, int ymargin);
extern __declspec(dllexport) 
void QGraphicsView_EnsureVisible23(QGraphicsView* self, QGraphicsItem* item, int xmargin);
extern __declspec(dllexport) 
void QGraphicsView_EnsureVisible32(QGraphicsView* self, QGraphicsItem* item, int xmargin, int ymargin);
extern __declspec(dllexport) 
void QGraphicsView_FitInView22(QGraphicsView* self, QRectF* rect, int aspectRadioMode);
extern __declspec(dllexport) 
void QGraphicsView_FitInView5(QGraphicsView* self, double x, double y, double w, double h, int aspectRadioMode);
extern __declspec(dllexport) 
void QGraphicsView_FitInView23(QGraphicsView* self, QGraphicsItem* item, int aspectRadioMode);
extern __declspec(dllexport) 
void QGraphicsView_Render2(QGraphicsView* self, QPainter* painter, QRectF* target);
extern __declspec(dllexport) 
void QGraphicsView_Render3(QGraphicsView* self, QPainter* painter, QRectF* target, QRect* source);
extern __declspec(dllexport) 
void QGraphicsView_Render4(QGraphicsView* self, QPainter* painter, QRectF* target, QRect* source, int aspectRatioMode);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsView_Items22(const QGraphicsView* self, QRect* rect, int mode);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsView_Items5(const QGraphicsView* self, int x, int y, int w, int h, int mode);
extern __declspec(dllexport) 
struct miqt_array /* of QGraphicsItem* */  QGraphicsView_Items24(const QGraphicsView* self, QPainterPath* path, int mode);
extern __declspec(dllexport) 
void QGraphicsView_InvalidateScene1(QGraphicsView* self, QRectF* rect);
extern __declspec(dllexport) 
void QGraphicsView_InvalidateScene2(QGraphicsView* self, QRectF* rect, int layers);
extern __declspec(dllexport) 
void QGraphicsView_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QGraphicsView_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QGraphicsView_override_virtual_Metacast(void* self, intptr_t slot);
void* QGraphicsView_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QGraphicsView_Delete(QGraphicsView* self, bool isSubclass);

}
