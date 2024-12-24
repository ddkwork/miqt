#pragma once
#ifndef MIQT_QT_GEN_QPAINTENGINE_H
#define MIQT_QT_GEN_QPAINTENGINE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QBrush;
class QFont;
class QImage;
class QLine;
class QLineF;
class QPaintDevice;
class QPaintEngine;
class QPaintEngineState;
class QPainter;
class QPainterPath;
class QPen;
class QPixmap;
class QPoint;
class QPointF;
class QRect;
class QRectF;
class QRegion;
class QSize;
class QTextItem;
class QTransform;
class _GUID;
class type_info;
#else
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
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTextItem* QTextItem_new();
extern __declspec(dllexport) QTextItem* QTextItem_new2(QTextItem* param1);
extern __declspec(dllexport) double QTextItem_Descent(const QTextItem* self);
extern __declspec(dllexport) double QTextItem_Ascent(const QTextItem* self);
extern __declspec(dllexport) double QTextItem_Width(const QTextItem* self);
extern __declspec(dllexport) RenderFlags QTextItem_RenderFlags(const QTextItem* self);
extern __declspec(dllexport) struct miqt_string QTextItem_Text(const QTextItem* self);
extern __declspec(dllexport) QFont* QTextItem_Font(const QTextItem* self);
extern __declspec(dllexport) void QTextItem_Delete(QTextItem* self, bool isSubclass);

extern __declspec(dllexport) QPaintEngine* QPaintEngine_new();
extern __declspec(dllexport) QPaintEngine* QPaintEngine_new2(PaintEngineFeatures features);
extern __declspec(dllexport) bool QPaintEngine_IsActive(const QPaintEngine* self);
extern __declspec(dllexport) void QPaintEngine_SetActive(QPaintEngine* self, bool newState);
extern __declspec(dllexport) bool QPaintEngine_Begin(QPaintEngine* self, QPaintDevice* pdev);
extern __declspec(dllexport) bool QPaintEngine_End(QPaintEngine* self);
extern __declspec(dllexport) void QPaintEngine_UpdateState(QPaintEngine* self, QPaintEngineState* state);
extern __declspec(dllexport) void QPaintEngine_DrawRects(QPaintEngine* self, QRect* rects, int rectCount);
extern __declspec(dllexport) void QPaintEngine_DrawRects2(QPaintEngine* self, QRectF* rects, int rectCount);
extern __declspec(dllexport) void QPaintEngine_DrawLines(QPaintEngine* self, QLine* lines, int lineCount);
extern __declspec(dllexport) void QPaintEngine_DrawLines2(QPaintEngine* self, QLineF* lines, int lineCount);
extern __declspec(dllexport) void QPaintEngine_DrawEllipse(QPaintEngine* self, QRectF* r);
extern __declspec(dllexport) void QPaintEngine_DrawEllipseWithQRect(QPaintEngine* self, QRect* r);
extern __declspec(dllexport) void QPaintEngine_DrawPath(QPaintEngine* self, QPainterPath* path);
extern __declspec(dllexport) void QPaintEngine_DrawPoints(QPaintEngine* self, QPointF* points, int pointCount);
extern __declspec(dllexport) void QPaintEngine_DrawPoints2(QPaintEngine* self, QPoint* points, int pointCount);
extern __declspec(dllexport) void QPaintEngine_DrawPolygon(QPaintEngine* self, QPointF* points, int pointCount, PolygonDrawMode mode);
extern __declspec(dllexport) void QPaintEngine_DrawPolygon2(QPaintEngine* self, QPoint* points, int pointCount, PolygonDrawMode mode);
extern __declspec(dllexport) void QPaintEngine_DrawPixmap(QPaintEngine* self, QRectF* r, QPixmap* pm, QRectF* sr);
extern __declspec(dllexport) void QPaintEngine_DrawTextItem(QPaintEngine* self, QPointF* p, QTextItem* textItem);
extern __declspec(dllexport) void QPaintEngine_DrawTiledPixmap(QPaintEngine* self, QRectF* r, QPixmap* pixmap, QPointF* s);
extern __declspec(dllexport) void QPaintEngine_DrawImage(QPaintEngine* self, QRectF* r, QImage* pm, QRectF* sr, int flags);
extern __declspec(dllexport) void QPaintEngine_SetPaintDevice(QPaintEngine* self, QPaintDevice* device);
extern __declspec(dllexport) QPaintDevice* QPaintEngine_PaintDevice(const QPaintEngine* self);
extern __declspec(dllexport) void QPaintEngine_SetSystemClip(QPaintEngine* self, QRegion* baseClip);
extern __declspec(dllexport) QRegion* QPaintEngine_SystemClip(const QPaintEngine* self);
extern __declspec(dllexport) void QPaintEngine_SetSystemRect(QPaintEngine* self, QRect* rect);
extern __declspec(dllexport) QRect* QPaintEngine_SystemRect(const QPaintEngine* self);
extern __declspec(dllexport) QPoint* QPaintEngine_CoordinateOffset(const QPaintEngine* self);
extern __declspec(dllexport) Type QPaintEngine_Type(const QPaintEngine* self);
extern __declspec(dllexport) void QPaintEngine_FixNegRect(QPaintEngine* self, int* x, int* y, int* w, int* h);
extern __declspec(dllexport) bool QPaintEngine_TestDirty(QPaintEngine* self, DirtyFlags df);
extern __declspec(dllexport) void QPaintEngine_SetDirty(QPaintEngine* self, DirtyFlags df);
extern __declspec(dllexport) void QPaintEngine_ClearDirty(QPaintEngine* self, DirtyFlags df);
extern __declspec(dllexport) bool QPaintEngine_HasFeature(const QPaintEngine* self, PaintEngineFeatures feature);
extern __declspec(dllexport) QPainter* QPaintEngine_Painter(const QPaintEngine* self);
extern __declspec(dllexport) void QPaintEngine_SyncState(QPaintEngine* self);
extern __declspec(dllexport) bool QPaintEngine_IsExtended(const QPaintEngine* self);
extern __declspec(dllexport) QPixmap* QPaintEngine_CreatePixmap(QPaintEngine* self, QSize* size);
extern __declspec(dllexport) QPixmap* QPaintEngine_CreatePixmapFromImage(QPaintEngine* self, QImage* image, int flags);
extern __declspec(dllexport) void QPaintEngine_override_virtual_Begin(void* self, intptr_t slot);
bool QPaintEngine_virtualbase_Begin(void* self, QPaintDevice* pdev);
extern __declspec(dllexport) void QPaintEngine_override_virtual_End(void* self, intptr_t slot);
bool QPaintEngine_virtualbase_End(void* self);
extern __declspec(dllexport) void QPaintEngine_override_virtual_UpdateState(void* self, intptr_t slot);
void QPaintEngine_virtualbase_UpdateState(void* self, QPaintEngineState* state);
extern __declspec(dllexport) void QPaintEngine_override_virtual_DrawRects(void* self, intptr_t slot);
void QPaintEngine_virtualbase_DrawRects(void* self, QRect* rects, int rectCount);
extern __declspec(dllexport) void QPaintEngine_override_virtual_DrawRects2(void* self, intptr_t slot);
void QPaintEngine_virtualbase_DrawRects2(void* self, QRectF* rects, int rectCount);
extern __declspec(dllexport) void QPaintEngine_override_virtual_DrawLines(void* self, intptr_t slot);
void QPaintEngine_virtualbase_DrawLines(void* self, QLine* lines, int lineCount);
extern __declspec(dllexport) void QPaintEngine_override_virtual_DrawLines2(void* self, intptr_t slot);
void QPaintEngine_virtualbase_DrawLines2(void* self, QLineF* lines, int lineCount);
extern __declspec(dllexport) void QPaintEngine_override_virtual_DrawEllipse(void* self, intptr_t slot);
void QPaintEngine_virtualbase_DrawEllipse(void* self, QRectF* r);
extern __declspec(dllexport) void QPaintEngine_override_virtual_DrawEllipseWithQRect(void* self, intptr_t slot);
void QPaintEngine_virtualbase_DrawEllipseWithQRect(void* self, QRect* r);
extern __declspec(dllexport) void QPaintEngine_override_virtual_DrawPath(void* self, intptr_t slot);
void QPaintEngine_virtualbase_DrawPath(void* self, QPainterPath* path);
extern __declspec(dllexport) void QPaintEngine_override_virtual_DrawPoints(void* self, intptr_t slot);
void QPaintEngine_virtualbase_DrawPoints(void* self, QPointF* points, int pointCount);
extern __declspec(dllexport) void QPaintEngine_override_virtual_DrawPoints2(void* self, intptr_t slot);
void QPaintEngine_virtualbase_DrawPoints2(void* self, QPoint* points, int pointCount);
extern __declspec(dllexport) void QPaintEngine_override_virtual_DrawPolygon(void* self, intptr_t slot);
void QPaintEngine_virtualbase_DrawPolygon(void* self, QPointF* points, int pointCount, PolygonDrawMode mode);
extern __declspec(dllexport) void QPaintEngine_override_virtual_DrawPolygon2(void* self, intptr_t slot);
void QPaintEngine_virtualbase_DrawPolygon2(void* self, QPoint* points, int pointCount, PolygonDrawMode mode);
extern __declspec(dllexport) void QPaintEngine_override_virtual_DrawPixmap(void* self, intptr_t slot);
void QPaintEngine_virtualbase_DrawPixmap(void* self, QRectF* r, QPixmap* pm, QRectF* sr);
extern __declspec(dllexport) void QPaintEngine_override_virtual_DrawTextItem(void* self, intptr_t slot);
void QPaintEngine_virtualbase_DrawTextItem(void* self, QPointF* p, QTextItem* textItem);
extern __declspec(dllexport) void QPaintEngine_override_virtual_DrawTiledPixmap(void* self, intptr_t slot);
void QPaintEngine_virtualbase_DrawTiledPixmap(void* self, QRectF* r, QPixmap* pixmap, QPointF* s);
extern __declspec(dllexport) void QPaintEngine_override_virtual_DrawImage(void* self, intptr_t slot);
void QPaintEngine_virtualbase_DrawImage(void* self, QRectF* r, QImage* pm, QRectF* sr, int flags);
extern __declspec(dllexport) void QPaintEngine_override_virtual_CoordinateOffset(void* self, intptr_t slot);
QPoint* QPaintEngine_virtualbase_CoordinateOffset(const void* self);
extern __declspec(dllexport) void QPaintEngine_override_virtual_Type(void* self, intptr_t slot);
Type QPaintEngine_virtualbase_Type(const void* self);
extern __declspec(dllexport) void QPaintEngine_override_virtual_CreatePixmap(void* self, intptr_t slot);
QPixmap* QPaintEngine_virtualbase_CreatePixmap(void* self, QSize* size);
extern __declspec(dllexport) void QPaintEngine_override_virtual_CreatePixmapFromImage(void* self, intptr_t slot);
QPixmap* QPaintEngine_virtualbase_CreatePixmapFromImage(void* self, QImage* image, int flags);
extern __declspec(dllexport) void QPaintEngine_Delete(QPaintEngine* self, bool isSubclass);

extern __declspec(dllexport) int QPaintEngineState_State(const QPaintEngineState* self);
extern __declspec(dllexport) QPen* QPaintEngineState_Pen(const QPaintEngineState* self);
extern __declspec(dllexport) QBrush* QPaintEngineState_Brush(const QPaintEngineState* self);
extern __declspec(dllexport) QPointF* QPaintEngineState_BrushOrigin(const QPaintEngineState* self);
extern __declspec(dllexport) QBrush* QPaintEngineState_BackgroundBrush(const QPaintEngineState* self);
extern __declspec(dllexport) int QPaintEngineState_BackgroundMode(const QPaintEngineState* self);
extern __declspec(dllexport) QFont* QPaintEngineState_Font(const QPaintEngineState* self);
extern __declspec(dllexport) QTransform* QPaintEngineState_Transform(const QPaintEngineState* self);
extern __declspec(dllexport) int QPaintEngineState_ClipOperation(const QPaintEngineState* self);
extern __declspec(dllexport) QRegion* QPaintEngineState_ClipRegion(const QPaintEngineState* self);
extern __declspec(dllexport) QPainterPath* QPaintEngineState_ClipPath(const QPaintEngineState* self);
extern __declspec(dllexport) bool QPaintEngineState_IsClipEnabled(const QPaintEngineState* self);
extern __declspec(dllexport) int QPaintEngineState_RenderHints(const QPaintEngineState* self);
extern __declspec(dllexport) int QPaintEngineState_CompositionMode(const QPaintEngineState* self);
extern __declspec(dllexport) double QPaintEngineState_Opacity(const QPaintEngineState* self);
extern __declspec(dllexport) QPainter* QPaintEngineState_Painter(const QPaintEngineState* self);
extern __declspec(dllexport) bool QPaintEngineState_BrushNeedsResolving(const QPaintEngineState* self);
extern __declspec(dllexport) bool QPaintEngineState_PenNeedsResolving(const QPaintEngineState* self);
extern __declspec(dllexport) void QPaintEngineState_Delete(QPaintEngineState* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
