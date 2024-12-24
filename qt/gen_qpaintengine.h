#pragma once
#ifndef MIQT_QT_GEN_QPAINTENGINE_H
#define MIQT_QT_GEN_QPAINTENGINE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QBrush QBrush;
typedef struct QFont QFont;
typedef struct QImage QImage;
typedef struct QLine QLine;
typedef struct QLineF QLineF;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPaintEngineState QPaintEngineState;
typedef struct QPainter QPainter;
typedef struct QPainterPath QPainterPath;
typedef struct QPen QPen;
typedef struct QPixmap QPixmap;
typedef struct QPoint QPoint;
typedef struct QPointF QPointF;
typedef struct QRect QRect;
typedef struct QRectF QRectF;
typedef struct QRegion QRegion;
typedef struct QSize QSize;
typedef struct QTextItem QTextItem;
typedef struct QTransform QTransform;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QTextItem* QTextItem_new();
extern __declspec(dllexport) 
QTextItem* QTextItem_new2(QTextItem* param1);
extern __declspec(dllexport) 
double QTextItem_Descent(const QTextItem* self);
extern __declspec(dllexport) 
double QTextItem_Ascent(const QTextItem* self);
extern __declspec(dllexport) 
double QTextItem_Width(const QTextItem* self);
extern __declspec(dllexport) 
RenderFlags QTextItem_RenderFlags(const QTextItem* self);
extern __declspec(dllexport) 
struct miqt_string QTextItem_Text(const QTextItem* self);
extern __declspec(dllexport) 
QFont* QTextItem_Font(const QTextItem* self);
extern __declspec(dllexport) 
void QTextItem_Delete(QTextItem* self, bool isSubclass);

extern __declspec(dllexport) 
QPaintEngine* QPaintEngine_new();
extern __declspec(dllexport) 
QPaintEngine* QPaintEngine_new2(PaintEngineFeatures features);
extern __declspec(dllexport) 
bool QPaintEngine_IsActive(const QPaintEngine* self);
extern __declspec(dllexport) 
void QPaintEngine_SetActive(QPaintEngine* self, bool newState);
extern __declspec(dllexport) 
bool QPaintEngine_Begin(QPaintEngine* self, QPaintDevice* pdev);
extern __declspec(dllexport) 
bool QPaintEngine_End(QPaintEngine* self);
extern __declspec(dllexport) 
void QPaintEngine_UpdateState(QPaintEngine* self, QPaintEngineState* state);
extern __declspec(dllexport) 
void QPaintEngine_DrawRects(QPaintEngine* self, QRect* rects, int rectCount);
extern __declspec(dllexport) 
void QPaintEngine_DrawRects2(QPaintEngine* self, QRectF* rects, int rectCount);
extern __declspec(dllexport) 
void QPaintEngine_DrawLines(QPaintEngine* self, QLine* lines, int lineCount);
extern __declspec(dllexport) 
void QPaintEngine_DrawLines2(QPaintEngine* self, QLineF* lines, int lineCount);
extern __declspec(dllexport) 
void QPaintEngine_DrawEllipse(QPaintEngine* self, QRectF* r);
extern __declspec(dllexport) 
void QPaintEngine_DrawEllipseWithQRect(QPaintEngine* self, QRect* r);
extern __declspec(dllexport) 
void QPaintEngine_DrawPath(QPaintEngine* self, QPainterPath* path);
extern __declspec(dllexport) 
void QPaintEngine_DrawPoints(QPaintEngine* self, QPointF* points, int pointCount);
extern __declspec(dllexport) 
void QPaintEngine_DrawPoints2(QPaintEngine* self, QPoint* points, int pointCount);
extern __declspec(dllexport) 
void QPaintEngine_DrawPolygon(QPaintEngine* self, QPointF* points, int pointCount, PolygonDrawMode mode);
extern __declspec(dllexport) 
void QPaintEngine_DrawPolygon2(QPaintEngine* self, QPoint* points, int pointCount, PolygonDrawMode mode);
extern __declspec(dllexport) 
void QPaintEngine_DrawPixmap(QPaintEngine* self, QRectF* r, QPixmap* pm, QRectF* sr);
extern __declspec(dllexport) 
void QPaintEngine_DrawTextItem(QPaintEngine* self, QPointF* p, QTextItem* textItem);
extern __declspec(dllexport) 
void QPaintEngine_DrawTiledPixmap(QPaintEngine* self, QRectF* r, QPixmap* pixmap, QPointF* s);
extern __declspec(dllexport) 
void QPaintEngine_DrawImage(QPaintEngine* self, QRectF* r, QImage* pm, QRectF* sr, int flags);
extern __declspec(dllexport) 
void QPaintEngine_SetPaintDevice(QPaintEngine* self, QPaintDevice* device);
extern __declspec(dllexport) 
QPaintDevice* QPaintEngine_PaintDevice(const QPaintEngine* self);
extern __declspec(dllexport) 
void QPaintEngine_SetSystemClip(QPaintEngine* self, QRegion* baseClip);
extern __declspec(dllexport) 
QRegion* QPaintEngine_SystemClip(const QPaintEngine* self);
extern __declspec(dllexport) 
void QPaintEngine_SetSystemRect(QPaintEngine* self, QRect* rect);
extern __declspec(dllexport) 
QRect* QPaintEngine_SystemRect(const QPaintEngine* self);
extern __declspec(dllexport) 
QPoint* QPaintEngine_CoordinateOffset(const QPaintEngine* self);
extern __declspec(dllexport) 
Type QPaintEngine_Type(const QPaintEngine* self);
extern __declspec(dllexport) 
void QPaintEngine_FixNegRect(QPaintEngine* self, int* x, int* y, int* w, int* h);
extern __declspec(dllexport) 
bool QPaintEngine_TestDirty(QPaintEngine* self, DirtyFlags df);
extern __declspec(dllexport) 
void QPaintEngine_SetDirty(QPaintEngine* self, DirtyFlags df);
extern __declspec(dllexport) 
void QPaintEngine_ClearDirty(QPaintEngine* self, DirtyFlags df);
extern __declspec(dllexport) 
bool QPaintEngine_HasFeature(const QPaintEngine* self, PaintEngineFeatures feature);
extern __declspec(dllexport) 
QPainter* QPaintEngine_Painter(const QPaintEngine* self);
extern __declspec(dllexport) 
void QPaintEngine_SyncState(QPaintEngine* self);
extern __declspec(dllexport) 
bool QPaintEngine_IsExtended(const QPaintEngine* self);
extern __declspec(dllexport) 
QPixmap* QPaintEngine_CreatePixmap(QPaintEngine* self, QSize* size);
extern __declspec(dllexport) 
QPixmap* QPaintEngine_CreatePixmapFromImage(QPaintEngine* self, QImage* image, int flags);
extern __declspec(dllexport) 
void QPaintEngine_Delete(QPaintEngine* self, bool isSubclass);

extern __declspec(dllexport) 
int QPaintEngineState_State(const QPaintEngineState* self);
extern __declspec(dllexport) 
QPen* QPaintEngineState_Pen(const QPaintEngineState* self);
extern __declspec(dllexport) 
QBrush* QPaintEngineState_Brush(const QPaintEngineState* self);
extern __declspec(dllexport) 
QPointF* QPaintEngineState_BrushOrigin(const QPaintEngineState* self);
extern __declspec(dllexport) 
QBrush* QPaintEngineState_BackgroundBrush(const QPaintEngineState* self);
extern __declspec(dllexport) 
int QPaintEngineState_BackgroundMode(const QPaintEngineState* self);
extern __declspec(dllexport) 
QFont* QPaintEngineState_Font(const QPaintEngineState* self);
extern __declspec(dllexport) 
QTransform* QPaintEngineState_Transform(const QPaintEngineState* self);
extern __declspec(dllexport) 
int QPaintEngineState_ClipOperation(const QPaintEngineState* self);
extern __declspec(dllexport) 
QRegion* QPaintEngineState_ClipRegion(const QPaintEngineState* self);
extern __declspec(dllexport) 
QPainterPath* QPaintEngineState_ClipPath(const QPaintEngineState* self);
extern __declspec(dllexport) 
bool QPaintEngineState_IsClipEnabled(const QPaintEngineState* self);
extern __declspec(dllexport) 
int QPaintEngineState_RenderHints(const QPaintEngineState* self);
extern __declspec(dllexport) 
int QPaintEngineState_CompositionMode(const QPaintEngineState* self);
extern __declspec(dllexport) 
double QPaintEngineState_Opacity(const QPaintEngineState* self);
extern __declspec(dllexport) 
QPainter* QPaintEngineState_Painter(const QPaintEngineState* self);
extern __declspec(dllexport) 
bool QPaintEngineState_BrushNeedsResolving(const QPaintEngineState* self);
extern __declspec(dllexport) 
bool QPaintEngineState_PenNeedsResolving(const QPaintEngineState* self);
extern __declspec(dllexport) 
void QPaintEngineState_Delete(QPaintEngineState* self, bool isSubclass);

}
