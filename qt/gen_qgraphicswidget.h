#pragma once
#ifndef MIQT_QT_GEN_QGRAPHICSWIDGET_H
#define MIQT_QT_GEN_QGRAPHICSWIDGET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAction QAction;
typedef struct QCloseEvent QCloseEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QFont QFont;
typedef struct QGraphicsItem QGraphicsItem;
typedef struct QGraphicsLayout QGraphicsLayout;
typedef struct QGraphicsLayoutItem QGraphicsLayoutItem;
typedef struct QGraphicsObject QGraphicsObject;
typedef struct QGraphicsSceneHoverEvent QGraphicsSceneHoverEvent;
typedef struct QGraphicsSceneMoveEvent QGraphicsSceneMoveEvent;
typedef struct QGraphicsSceneResizeEvent QGraphicsSceneResizeEvent;
typedef struct QGraphicsWidget QGraphicsWidget;
typedef struct QHideEvent QHideEvent;
typedef struct QKeySequence QKeySequence;
typedef struct QMarginsF QMarginsF;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPainter QPainter;
typedef struct QPainterPath QPainterPath;
typedef struct QPalette QPalette;
typedef struct QPointF QPointF;
typedef struct QRectF QRectF;
typedef struct QShowEvent QShowEvent;
typedef struct QSizeF QSizeF;
typedef struct QStyle QStyle;
typedef struct QStyleOption QStyleOption;
typedef struct QStyleOptionGraphicsItem QStyleOptionGraphicsItem;
typedef struct QVariant QVariant;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QGraphicsWidget* QGraphicsWidget_new();
extern __declspec(dllexport) 
QGraphicsWidget* QGraphicsWidget_new2(QGraphicsItem* parent);
extern __declspec(dllexport) 
QGraphicsWidget* QGraphicsWidget_new3(QGraphicsItem* parent, int wFlags);
extern __declspec(dllexport) 
void QGraphicsWidget_virtbase(QGraphicsWidget* src
, QGraphicsObject** outptr_QGraphicsObject
, QGraphicsLayoutItem** outptr_QGraphicsLayoutItem
);
extern __declspec(dllexport) 
QMetaObject* QGraphicsWidget_MetaObject(const QGraphicsWidget* self);
extern __declspec(dllexport) 
void* QGraphicsWidget_Metacast(QGraphicsWidget* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QGraphicsWidget_Tr(const char* s);
extern __declspec(dllexport) 
QGraphicsLayout* QGraphicsWidget_Layout(const QGraphicsWidget* self);
extern __declspec(dllexport) 
void QGraphicsWidget_SetLayout(QGraphicsWidget* self, QGraphicsLayout* layout);
extern __declspec(dllexport) 
void QGraphicsWidget_AdjustSize(QGraphicsWidget* self);
extern __declspec(dllexport) 
int QGraphicsWidget_LayoutDirection(const QGraphicsWidget* self);
extern __declspec(dllexport) 
void QGraphicsWidget_SetLayoutDirection(QGraphicsWidget* self, int direction);
extern __declspec(dllexport) 
void QGraphicsWidget_UnsetLayoutDirection(QGraphicsWidget* self);
extern __declspec(dllexport) 
QStyle* QGraphicsWidget_Style(const QGraphicsWidget* self);
extern __declspec(dllexport) 
void QGraphicsWidget_SetStyle(QGraphicsWidget* self, QStyle* style);
extern __declspec(dllexport) 
QFont* QGraphicsWidget_Font(const QGraphicsWidget* self);
extern __declspec(dllexport) 
void QGraphicsWidget_SetFont(QGraphicsWidget* self, QFont* font);
extern __declspec(dllexport) 
QPalette* QGraphicsWidget_Palette(const QGraphicsWidget* self);
extern __declspec(dllexport) 
void QGraphicsWidget_SetPalette(QGraphicsWidget* self, QPalette* palette);
extern __declspec(dllexport) 
bool QGraphicsWidget_AutoFillBackground(const QGraphicsWidget* self);
extern __declspec(dllexport) 
void QGraphicsWidget_SetAutoFillBackground(QGraphicsWidget* self, bool enabled);
extern __declspec(dllexport) 
void QGraphicsWidget_Resize(QGraphicsWidget* self, QSizeF* size);
extern __declspec(dllexport) 
void QGraphicsWidget_Resize2(QGraphicsWidget* self, double w, double h);
extern __declspec(dllexport) 
QSizeF* QGraphicsWidget_Size(const QGraphicsWidget* self);
extern __declspec(dllexport) 
void QGraphicsWidget_SetGeometry(QGraphicsWidget* self, QRectF* rect);
extern __declspec(dllexport) 
void QGraphicsWidget_SetGeometry2(QGraphicsWidget* self, double x, double y, double w, double h);
extern __declspec(dllexport) 
QRectF* QGraphicsWidget_Rect(const QGraphicsWidget* self);
extern __declspec(dllexport) 
void QGraphicsWidget_SetContentsMargins(QGraphicsWidget* self, double left, double top, double right, double bottom);
extern __declspec(dllexport) 
void QGraphicsWidget_SetContentsMarginsWithMargins(QGraphicsWidget* self, QMarginsF* margins);
extern __declspec(dllexport) 
void QGraphicsWidget_GetContentsMargins(const QGraphicsWidget* self, double* left, double* top, double* right, double* bottom);
extern __declspec(dllexport) 
void QGraphicsWidget_SetWindowFrameMargins(QGraphicsWidget* self, double left, double top, double right, double bottom);
extern __declspec(dllexport) 
void QGraphicsWidget_SetWindowFrameMarginsWithMargins(QGraphicsWidget* self, QMarginsF* margins);
extern __declspec(dllexport) 
void QGraphicsWidget_GetWindowFrameMargins(const QGraphicsWidget* self, double* left, double* top, double* right, double* bottom);
extern __declspec(dllexport) 
void QGraphicsWidget_UnsetWindowFrameMargins(QGraphicsWidget* self);
extern __declspec(dllexport) 
QRectF* QGraphicsWidget_WindowFrameGeometry(const QGraphicsWidget* self);
extern __declspec(dllexport) 
QRectF* QGraphicsWidget_WindowFrameRect(const QGraphicsWidget* self);
extern __declspec(dllexport) 
int QGraphicsWidget_WindowFlags(const QGraphicsWidget* self);
extern __declspec(dllexport) 
int QGraphicsWidget_WindowType(const QGraphicsWidget* self);
extern __declspec(dllexport) 
void QGraphicsWidget_SetWindowFlags(QGraphicsWidget* self, int wFlags);
extern __declspec(dllexport) 
bool QGraphicsWidget_IsActiveWindow(const QGraphicsWidget* self);
extern __declspec(dllexport) 
void QGraphicsWidget_SetWindowTitle(QGraphicsWidget* self, struct miqt_string title);
extern __declspec(dllexport) 
struct miqt_string QGraphicsWidget_WindowTitle(const QGraphicsWidget* self);
extern __declspec(dllexport) 
int QGraphicsWidget_FocusPolicy(const QGraphicsWidget* self);
extern __declspec(dllexport) 
void QGraphicsWidget_SetFocusPolicy(QGraphicsWidget* self, int policy);
extern __declspec(dllexport) 
void QGraphicsWidget_SetTabOrder(QGraphicsWidget* first, QGraphicsWidget* second);
extern __declspec(dllexport) 
QGraphicsWidget* QGraphicsWidget_FocusWidget(const QGraphicsWidget* self);
extern __declspec(dllexport) 
int QGraphicsWidget_GrabShortcut(QGraphicsWidget* self, QKeySequence* sequence);
extern __declspec(dllexport) 
void QGraphicsWidget_ReleaseShortcut(QGraphicsWidget* self, int id);
extern __declspec(dllexport) 
void QGraphicsWidget_SetShortcutEnabled(QGraphicsWidget* self, int id);
extern __declspec(dllexport) 
void QGraphicsWidget_SetShortcutAutoRepeat(QGraphicsWidget* self, int id);
extern __declspec(dllexport) 
void QGraphicsWidget_AddAction(QGraphicsWidget* self, QAction* action);
extern __declspec(dllexport) 
void QGraphicsWidget_AddActions(QGraphicsWidget* self, struct miqt_array /* of QAction* */  actions);
extern __declspec(dllexport) 
void QGraphicsWidget_InsertActions(QGraphicsWidget* self, QAction* before, struct miqt_array /* of QAction* */  actions);
extern __declspec(dllexport) 
void QGraphicsWidget_InsertAction(QGraphicsWidget* self, QAction* before, QAction* action);
extern __declspec(dllexport) 
void QGraphicsWidget_RemoveAction(QGraphicsWidget* self, QAction* action);
extern __declspec(dllexport) 
struct miqt_array /* of QAction* */  QGraphicsWidget_Actions(const QGraphicsWidget* self);
extern __declspec(dllexport) 
void QGraphicsWidget_SetAttribute(QGraphicsWidget* self, int attribute);
extern __declspec(dllexport) 
bool QGraphicsWidget_TestAttribute(const QGraphicsWidget* self, int attribute);
extern __declspec(dllexport) 
int QGraphicsWidget_Type(const QGraphicsWidget* self);
extern __declspec(dllexport) 
void QGraphicsWidget_Paint(QGraphicsWidget* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget);
extern __declspec(dllexport) 
void QGraphicsWidget_PaintWindowFrame(QGraphicsWidget* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget);
extern __declspec(dllexport) 
QRectF* QGraphicsWidget_BoundingRect(const QGraphicsWidget* self);
extern __declspec(dllexport) 
QPainterPath* QGraphicsWidget_Shape(const QGraphicsWidget* self);
extern __declspec(dllexport) 
void QGraphicsWidget_GeometryChanged(QGraphicsWidget* self);
void QGraphicsWidget_connect_GeometryChanged(QGraphicsWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QGraphicsWidget_LayoutChanged(QGraphicsWidget* self);
void QGraphicsWidget_connect_LayoutChanged(QGraphicsWidget* self, intptr_t slot);
extern __declspec(dllexport) 
bool QGraphicsWidget_Close(QGraphicsWidget* self);
extern __declspec(dllexport) 
void QGraphicsWidget_InitStyleOption(const QGraphicsWidget* self, QStyleOption* option);
extern __declspec(dllexport) 
QSizeF* QGraphicsWidget_SizeHint(const QGraphicsWidget* self, int which, QSizeF* constraint);
extern __declspec(dllexport) 
void QGraphicsWidget_UpdateGeometry(QGraphicsWidget* self);
extern __declspec(dllexport) 
QVariant* QGraphicsWidget_ItemChange(QGraphicsWidget* self, GraphicsItemChange change, QVariant* value);
extern __declspec(dllexport) 
QVariant* QGraphicsWidget_PropertyChange(QGraphicsWidget* self, struct miqt_string propertyName, QVariant* value);
extern __declspec(dllexport) 
bool QGraphicsWidget_SceneEvent(QGraphicsWidget* self, QEvent* event);
extern __declspec(dllexport) 
bool QGraphicsWidget_WindowFrameEvent(QGraphicsWidget* self, QEvent* e);
extern __declspec(dllexport) 
int QGraphicsWidget_WindowFrameSectionAt(const QGraphicsWidget* self, QPointF* pos);
extern __declspec(dllexport) 
bool QGraphicsWidget_Event(QGraphicsWidget* self, QEvent* event);
extern __declspec(dllexport) 
void QGraphicsWidget_ChangeEvent(QGraphicsWidget* self, QEvent* event);
extern __declspec(dllexport) 
void QGraphicsWidget_CloseEvent(QGraphicsWidget* self, QCloseEvent* event);
extern __declspec(dllexport) 
void QGraphicsWidget_FocusInEvent(QGraphicsWidget* self, QFocusEvent* event);
extern __declspec(dllexport) 
bool QGraphicsWidget_FocusNextPrevChild(QGraphicsWidget* self, bool next);
extern __declspec(dllexport) 
void QGraphicsWidget_FocusOutEvent(QGraphicsWidget* self, QFocusEvent* event);
extern __declspec(dllexport) 
void QGraphicsWidget_HideEvent(QGraphicsWidget* self, QHideEvent* event);
extern __declspec(dllexport) 
void QGraphicsWidget_MoveEvent(QGraphicsWidget* self, QGraphicsSceneMoveEvent* event);
extern __declspec(dllexport) 
void QGraphicsWidget_PolishEvent(QGraphicsWidget* self);
extern __declspec(dllexport) 
void QGraphicsWidget_ResizeEvent(QGraphicsWidget* self, QGraphicsSceneResizeEvent* event);
extern __declspec(dllexport) 
void QGraphicsWidget_ShowEvent(QGraphicsWidget* self, QShowEvent* event);
extern __declspec(dllexport) 
void QGraphicsWidget_HoverMoveEvent(QGraphicsWidget* self, QGraphicsSceneHoverEvent* event);
extern __declspec(dllexport) 
void QGraphicsWidget_HoverLeaveEvent(QGraphicsWidget* self, QGraphicsSceneHoverEvent* event);
extern __declspec(dllexport) 
void QGraphicsWidget_GrabMouseEvent(QGraphicsWidget* self, QEvent* event);
extern __declspec(dllexport) 
void QGraphicsWidget_UngrabMouseEvent(QGraphicsWidget* self, QEvent* event);
extern __declspec(dllexport) 
void QGraphicsWidget_GrabKeyboardEvent(QGraphicsWidget* self, QEvent* event);
extern __declspec(dllexport) 
void QGraphicsWidget_UngrabKeyboardEvent(QGraphicsWidget* self, QEvent* event);
extern __declspec(dllexport) 
struct miqt_string QGraphicsWidget_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QGraphicsWidget_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
int QGraphicsWidget_GrabShortcut2(QGraphicsWidget* self, QKeySequence* sequence, int context);
extern __declspec(dllexport) 
void QGraphicsWidget_SetShortcutEnabled2(QGraphicsWidget* self, int id, bool enabled);
extern __declspec(dllexport) 
void QGraphicsWidget_SetShortcutAutoRepeat2(QGraphicsWidget* self, int id, bool enabled);
extern __declspec(dllexport) 
void QGraphicsWidget_SetAttribute2(QGraphicsWidget* self, int attribute, bool on);
extern __declspec(dllexport) 
void QGraphicsWidget_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QGraphicsWidget_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QGraphicsWidget_override_virtual_Metacast(void* self, intptr_t slot);
void* QGraphicsWidget_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QGraphicsWidget_Delete(QGraphicsWidget* self, bool isSubclass);

}
