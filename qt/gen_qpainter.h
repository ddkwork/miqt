#pragma once
#ifndef MIQT_QT_GEN_QPAINTER_H
#define MIQT_QT_GEN_QPAINTER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QPainter__PixmapFragment)
typedef QPainter::PixmapFragment QPainter__PixmapFragment;
typedef struct QBrush QBrush;
typedef struct QColor QColor;
typedef struct QFont QFont;
typedef struct QFontInfo QFontInfo;
typedef struct QFontMetrics QFontMetrics;
typedef struct QGlyphRun QGlyphRun;
typedef struct QImage QImage;
typedef struct QLine QLine;
typedef struct QLineF QLineF;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPainter QPainter;
typedef struct QPainter__PixmapFragment QPainter__PixmapFragment;
typedef struct QPainterPath QPainterPath;
typedef struct QPen QPen;
typedef struct QPicture QPicture;
typedef struct QPixmap QPixmap;
typedef struct QPoint QPoint;
typedef struct QPointF QPointF;
typedef struct QRect QRect;
typedef struct QRectF QRectF;
typedef struct QRegion QRegion;
typedef struct QStaticText QStaticText;
typedef struct QTextItem QTextItem;
typedef struct QTextOption QTextOption;
typedef struct QTransform QTransform;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QPainter* QPainter_new();
extern __declspec(dllexport) 
QPainter* QPainter_new2(QPaintDevice* param1);
extern __declspec(dllexport) 
QPaintDevice* QPainter_Device(const QPainter* self);
extern __declspec(dllexport) 
bool QPainter_Begin(QPainter* self, QPaintDevice* param1);
extern __declspec(dllexport) 
bool QPainter_End(QPainter* self);
extern __declspec(dllexport) 
bool QPainter_IsActive(const QPainter* self);
extern __declspec(dllexport) 
void QPainter_SetCompositionMode(QPainter* self, CompositionMode mode);
extern __declspec(dllexport) 
CompositionMode QPainter_CompositionMode(const QPainter* self);
extern __declspec(dllexport) 
QFont* QPainter_Font(const QPainter* self);
extern __declspec(dllexport) 
void QPainter_SetFont(QPainter* self, QFont* f);
extern __declspec(dllexport) 
QFontMetrics* QPainter_FontMetrics(const QPainter* self);
extern __declspec(dllexport) 
QFontInfo* QPainter_FontInfo(const QPainter* self);
extern __declspec(dllexport) 
void QPainter_SetPen(QPainter* self, QColor* color);
extern __declspec(dllexport) 
void QPainter_SetPenWithPen(QPainter* self, QPen* pen);
extern __declspec(dllexport) 
void QPainter_SetPenWithStyle(QPainter* self, int style);
extern __declspec(dllexport) 
QPen* QPainter_Pen(const QPainter* self);
extern __declspec(dllexport) 
void QPainter_SetBrush(QPainter* self, QBrush* brush);
extern __declspec(dllexport) 
void QPainter_SetBrushWithStyle(QPainter* self, int style);
extern __declspec(dllexport) 
void QPainter_SetBrushWithColor(QPainter* self, QColor* color);
extern __declspec(dllexport) 
void QPainter_SetBrush2(QPainter* self, int color);
extern __declspec(dllexport) 
QBrush* QPainter_Brush(const QPainter* self);
extern __declspec(dllexport) 
void QPainter_SetBackgroundMode(QPainter* self, int mode);
extern __declspec(dllexport) 
int QPainter_BackgroundMode(const QPainter* self);
extern __declspec(dllexport) 
QPoint* QPainter_BrushOrigin(const QPainter* self);
extern __declspec(dllexport) 
void QPainter_SetBrushOrigin(QPainter* self, int x, int y);
extern __declspec(dllexport) 
void QPainter_SetBrushOriginWithBrushOrigin(QPainter* self, QPoint* brushOrigin);
extern __declspec(dllexport) 
void QPainter_SetBrushOrigin2(QPainter* self, QPointF* brushOrigin);
extern __declspec(dllexport) 
void QPainter_SetBackground(QPainter* self, QBrush* bg);
extern __declspec(dllexport) 
QBrush* QPainter_Background(const QPainter* self);
extern __declspec(dllexport) 
double QPainter_Opacity(const QPainter* self);
extern __declspec(dllexport) 
void QPainter_SetOpacity(QPainter* self, double opacity);
extern __declspec(dllexport) 
QRegion* QPainter_ClipRegion(const QPainter* self);
extern __declspec(dllexport) 
QPainterPath* QPainter_ClipPath(const QPainter* self);
extern __declspec(dllexport) 
void QPainter_SetClipRect(QPainter* self, QRectF* param1);
extern __declspec(dllexport) 
void QPainter_SetClipRectWithQRect(QPainter* self, QRect* param1);
extern __declspec(dllexport) 
void QPainter_SetClipRect2(QPainter* self, int x, int y, int w, int h);
extern __declspec(dllexport) 
void QPainter_SetClipRegion(QPainter* self, QRegion* param1);
extern __declspec(dllexport) 
void QPainter_SetClipPath(QPainter* self, QPainterPath* path);
extern __declspec(dllexport) 
void QPainter_SetClipping(QPainter* self, bool enable);
extern __declspec(dllexport) 
bool QPainter_HasClipping(const QPainter* self);
extern __declspec(dllexport) 
QRectF* QPainter_ClipBoundingRect(const QPainter* self);
extern __declspec(dllexport) 
void QPainter_Save(QPainter* self);
extern __declspec(dllexport) 
void QPainter_Restore(QPainter* self);
extern __declspec(dllexport) 
void QPainter_SetTransform(QPainter* self, QTransform* transform);
extern __declspec(dllexport) 
QTransform* QPainter_Transform(const QPainter* self);
extern __declspec(dllexport) 
QTransform* QPainter_DeviceTransform(const QPainter* self);
extern __declspec(dllexport) 
void QPainter_ResetTransform(QPainter* self);
extern __declspec(dllexport) 
void QPainter_SetWorldTransform(QPainter* self, QTransform* matrix);
extern __declspec(dllexport) 
QTransform* QPainter_WorldTransform(const QPainter* self);
extern __declspec(dllexport) 
QTransform* QPainter_CombinedTransform(const QPainter* self);
extern __declspec(dllexport) 
void QPainter_SetWorldMatrixEnabled(QPainter* self, bool enabled);
extern __declspec(dllexport) 
bool QPainter_WorldMatrixEnabled(const QPainter* self);
extern __declspec(dllexport) 
void QPainter_Scale(QPainter* self, double sx, double sy);
extern __declspec(dllexport) 
void QPainter_Shear(QPainter* self, double sh, double sv);
extern __declspec(dllexport) 
void QPainter_Rotate(QPainter* self, double a);
extern __declspec(dllexport) 
void QPainter_Translate(QPainter* self, QPointF* offset);
extern __declspec(dllexport) 
void QPainter_TranslateWithOffset(QPainter* self, QPoint* offset);
extern __declspec(dllexport) 
void QPainter_Translate2(QPainter* self, double dx, double dy);
extern __declspec(dllexport) 
QRect* QPainter_Window(const QPainter* self);
extern __declspec(dllexport) 
void QPainter_SetWindow(QPainter* self, QRect* window);
extern __declspec(dllexport) 
void QPainter_SetWindow2(QPainter* self, int x, int y, int w, int h);
extern __declspec(dllexport) 
QRect* QPainter_Viewport(const QPainter* self);
extern __declspec(dllexport) 
void QPainter_SetViewport(QPainter* self, QRect* viewport);
extern __declspec(dllexport) 
void QPainter_SetViewport2(QPainter* self, int x, int y, int w, int h);
extern __declspec(dllexport) 
void QPainter_SetViewTransformEnabled(QPainter* self, bool enable);
extern __declspec(dllexport) 
bool QPainter_ViewTransformEnabled(const QPainter* self);
extern __declspec(dllexport) 
void QPainter_StrokePath(QPainter* self, QPainterPath* path, QPen* pen);
extern __declspec(dllexport) 
void QPainter_FillPath(QPainter* self, QPainterPath* path, QBrush* brush);
extern __declspec(dllexport) 
void QPainter_DrawPath(QPainter* self, QPainterPath* path);
extern __declspec(dllexport) 
void QPainter_DrawPoint(QPainter* self, QPointF* pt);
extern __declspec(dllexport) 
void QPainter_DrawPointWithQPoint(QPainter* self, QPoint* p);
extern __declspec(dllexport) 
void QPainter_DrawPoint2(QPainter* self, int x, int y);
extern __declspec(dllexport) 
void QPainter_DrawPoints(QPainter* self, QPointF* points, int pointCount);
extern __declspec(dllexport) 
void QPainter_DrawPoints2(QPainter* self, QPoint* points, int pointCount);
extern __declspec(dllexport) 
void QPainter_DrawLine(QPainter* self, QLineF* line);
extern __declspec(dllexport) 
void QPainter_DrawLineWithLine(QPainter* self, QLine* line);
extern __declspec(dllexport) 
void QPainter_DrawLine2(QPainter* self, int x1, int y1, int x2, int y2);
extern __declspec(dllexport) 
void QPainter_DrawLine3(QPainter* self, QPoint* p1, QPoint* p2);
extern __declspec(dllexport) 
void QPainter_DrawLine4(QPainter* self, QPointF* p1, QPointF* p2);
extern __declspec(dllexport) 
void QPainter_DrawLines(QPainter* self, QLineF* lines, int lineCount);
extern __declspec(dllexport) 
void QPainter_DrawLinesWithLines(QPainter* self, struct miqt_array /* of QLineF* */  lines);
extern __declspec(dllexport) 
void QPainter_DrawLines2(QPainter* self, QPointF* pointPairs, int lineCount);
extern __declspec(dllexport) 
void QPainter_DrawLinesWithPointPairs(QPainter* self, struct miqt_array /* of QPointF* */  pointPairs);
extern __declspec(dllexport) 
void QPainter_DrawLines3(QPainter* self, QLine* lines, int lineCount);
extern __declspec(dllexport) 
void QPainter_DrawLines4(QPainter* self, struct miqt_array /* of QLine* */  lines);
extern __declspec(dllexport) 
void QPainter_DrawLines5(QPainter* self, QPoint* pointPairs, int lineCount);
extern __declspec(dllexport) 
void QPainter_DrawLines6(QPainter* self, struct miqt_array /* of QPoint* */  pointPairs);
extern __declspec(dllexport) 
void QPainter_DrawRect(QPainter* self, QRectF* rect);
extern __declspec(dllexport) 
void QPainter_DrawRect2(QPainter* self, int x1, int y1, int w, int h);
extern __declspec(dllexport) 
void QPainter_DrawRectWithRect(QPainter* self, QRect* rect);
extern __declspec(dllexport) 
void QPainter_DrawRects(QPainter* self, QRectF* rects, int rectCount);
extern __declspec(dllexport) 
void QPainter_DrawRectsWithRectangles(QPainter* self, struct miqt_array /* of QRectF* */  rectangles);
extern __declspec(dllexport) 
void QPainter_DrawRects2(QPainter* self, QRect* rects, int rectCount);
extern __declspec(dllexport) 
void QPainter_DrawRects3(QPainter* self, struct miqt_array /* of QRect* */  rectangles);
extern __declspec(dllexport) 
void QPainter_DrawEllipse(QPainter* self, QRectF* r);
extern __declspec(dllexport) 
void QPainter_DrawEllipseWithQRect(QPainter* self, QRect* r);
extern __declspec(dllexport) 
void QPainter_DrawEllipse2(QPainter* self, int x, int y, int w, int h);
extern __declspec(dllexport) 
void QPainter_DrawEllipse3(QPainter* self, QPointF* center, double rx, double ry);
extern __declspec(dllexport) 
void QPainter_DrawEllipse4(QPainter* self, QPoint* center, int rx, int ry);
extern __declspec(dllexport) 
void QPainter_DrawPolyline(QPainter* self, QPointF* points, int pointCount);
extern __declspec(dllexport) 
void QPainter_DrawPolyline2(QPainter* self, QPoint* points, int pointCount);
extern __declspec(dllexport) 
void QPainter_DrawPolygon(QPainter* self, QPointF* points, int pointCount);
extern __declspec(dllexport) 
void QPainter_DrawPolygon2(QPainter* self, QPoint* points, int pointCount);
extern __declspec(dllexport) 
void QPainter_DrawConvexPolygon(QPainter* self, QPointF* points, int pointCount);
extern __declspec(dllexport) 
void QPainter_DrawConvexPolygon2(QPainter* self, QPoint* points, int pointCount);
extern __declspec(dllexport) 
void QPainter_DrawArc(QPainter* self, QRectF* rect, int a, int alen);
extern __declspec(dllexport) 
void QPainter_DrawArc2(QPainter* self, QRect* param1, int a, int alen);
extern __declspec(dllexport) 
void QPainter_DrawArc3(QPainter* self, int x, int y, int w, int h, int a, int alen);
extern __declspec(dllexport) 
void QPainter_DrawPie(QPainter* self, QRectF* rect, int a, int alen);
extern __declspec(dllexport) 
void QPainter_DrawPie2(QPainter* self, int x, int y, int w, int h, int a, int alen);
extern __declspec(dllexport) 
void QPainter_DrawPie3(QPainter* self, QRect* param1, int a, int alen);
extern __declspec(dllexport) 
void QPainter_DrawChord(QPainter* self, QRectF* rect, int a, int alen);
extern __declspec(dllexport) 
void QPainter_DrawChord2(QPainter* self, int x, int y, int w, int h, int a, int alen);
extern __declspec(dllexport) 
void QPainter_DrawChord3(QPainter* self, QRect* param1, int a, int alen);
extern __declspec(dllexport) 
void QPainter_DrawRoundedRect(QPainter* self, QRectF* rect, double xRadius, double yRadius);
extern __declspec(dllexport) 
void QPainter_DrawRoundedRect2(QPainter* self, int x, int y, int w, int h, double xRadius, double yRadius);
extern __declspec(dllexport) 
void QPainter_DrawRoundedRect3(QPainter* self, QRect* rect, double xRadius, double yRadius);
extern __declspec(dllexport) 
void QPainter_DrawTiledPixmap(QPainter* self, QRectF* rect, QPixmap* pm);
extern __declspec(dllexport) 
void QPainter_DrawTiledPixmap2(QPainter* self, int x, int y, int w, int h, QPixmap* param5);
extern __declspec(dllexport) 
void QPainter_DrawTiledPixmap3(QPainter* self, QRect* param1, QPixmap* param2);
extern __declspec(dllexport) 
void QPainter_DrawPicture(QPainter* self, QPointF* p, QPicture* picture);
extern __declspec(dllexport) 
void QPainter_DrawPicture2(QPainter* self, int x, int y, QPicture* picture);
extern __declspec(dllexport) 
void QPainter_DrawPicture3(QPainter* self, QPoint* p, QPicture* picture);
extern __declspec(dllexport) 
void QPainter_DrawPixmap(QPainter* self, QRectF* targetRect, QPixmap* pixmap, QRectF* sourceRect);
extern __declspec(dllexport) 
void QPainter_DrawPixmap2(QPainter* self, QRect* targetRect, QPixmap* pixmap, QRect* sourceRect);
extern __declspec(dllexport) 
void QPainter_DrawPixmap3(QPainter* self, int x, int y, int w, int h, QPixmap* pm, int sx, int sy, int sw, int sh);
extern __declspec(dllexport) 
void QPainter_DrawPixmap4(QPainter* self, int x, int y, QPixmap* pm, int sx, int sy, int sw, int sh);
extern __declspec(dllexport) 
void QPainter_DrawPixmap5(QPainter* self, QPointF* p, QPixmap* pm, QRectF* sr);
extern __declspec(dllexport) 
void QPainter_DrawPixmap6(QPainter* self, QPoint* p, QPixmap* pm, QRect* sr);
extern __declspec(dllexport) 
void QPainter_DrawPixmap7(QPainter* self, QPointF* p, QPixmap* pm);
extern __declspec(dllexport) 
void QPainter_DrawPixmap8(QPainter* self, QPoint* p, QPixmap* pm);
extern __declspec(dllexport) 
void QPainter_DrawPixmap9(QPainter* self, int x, int y, QPixmap* pm);
extern __declspec(dllexport) 
void QPainter_DrawPixmap10(QPainter* self, QRect* r, QPixmap* pm);
extern __declspec(dllexport) 
void QPainter_DrawPixmap11(QPainter* self, int x, int y, int w, int h, QPixmap* pm);
extern __declspec(dllexport) 
void QPainter_DrawPixmapFragments(QPainter* self, const PixmapFragment* fragments, int fragmentCount, QPixmap* pixmap);
extern __declspec(dllexport) 
void QPainter_DrawImage(QPainter* self, QRectF* targetRect, QImage* image, QRectF* sourceRect);
extern __declspec(dllexport) 
void QPainter_DrawImage2(QPainter* self, QRect* targetRect, QImage* image, QRect* sourceRect);
extern __declspec(dllexport) 
void QPainter_DrawImage3(QPainter* self, QPointF* p, QImage* image, QRectF* sr);
extern __declspec(dllexport) 
void QPainter_DrawImage4(QPainter* self, QPoint* p, QImage* image, QRect* sr);
extern __declspec(dllexport) 
void QPainter_DrawImage5(QPainter* self, QRectF* r, QImage* image);
extern __declspec(dllexport) 
void QPainter_DrawImage6(QPainter* self, QRect* r, QImage* image);
extern __declspec(dllexport) 
void QPainter_DrawImage7(QPainter* self, QPointF* p, QImage* image);
extern __declspec(dllexport) 
void QPainter_DrawImage8(QPainter* self, QPoint* p, QImage* image);
extern __declspec(dllexport) 
void QPainter_DrawImage9(QPainter* self, int x, int y, QImage* image);
extern __declspec(dllexport) 
void QPainter_SetLayoutDirection(QPainter* self, int direction);
extern __declspec(dllexport) 
int QPainter_LayoutDirection(const QPainter* self);
extern __declspec(dllexport) 
void QPainter_DrawGlyphRun(QPainter* self, QPointF* position, QGlyphRun* glyphRun);
extern __declspec(dllexport) 
void QPainter_DrawStaticText(QPainter* self, QPointF* topLeftPosition, QStaticText* staticText);
extern __declspec(dllexport) 
void QPainter_DrawStaticText2(QPainter* self, QPoint* topLeftPosition, QStaticText* staticText);
extern __declspec(dllexport) 
void QPainter_DrawStaticText3(QPainter* self, int left, int top, QStaticText* staticText);
extern __declspec(dllexport) 
void QPainter_DrawText(QPainter* self, QPointF* p, struct miqt_string s);
extern __declspec(dllexport) 
void QPainter_DrawText2(QPainter* self, QPoint* p, struct miqt_string s);
extern __declspec(dllexport) 
void QPainter_DrawText3(QPainter* self, int x, int y, struct miqt_string s);
extern __declspec(dllexport) 
void QPainter_DrawText4(QPainter* self, QPointF* p, struct miqt_string str, int tf, int justificationPadding);
extern __declspec(dllexport) 
void QPainter_DrawText5(QPainter* self, QRectF* r, int flags, struct miqt_string text);
extern __declspec(dllexport) 
void QPainter_DrawText6(QPainter* self, QRect* r, int flags, struct miqt_string text);
extern __declspec(dllexport) 
void QPainter_DrawText7(QPainter* self, int x, int y, int w, int h, int flags, struct miqt_string text);
extern __declspec(dllexport) 
void QPainter_DrawText8(QPainter* self, QRectF* r, struct miqt_string text);
extern __declspec(dllexport) 
QRectF* QPainter_BoundingRect(QPainter* self, QRectF* rect, int flags, struct miqt_string text);
extern __declspec(dllexport) 
QRect* QPainter_BoundingRect2(QPainter* self, QRect* rect, int flags, struct miqt_string text);
extern __declspec(dllexport) 
QRect* QPainter_BoundingRect3(QPainter* self, int x, int y, int w, int h, int flags, struct miqt_string text);
extern __declspec(dllexport) 
QRectF* QPainter_BoundingRect4(QPainter* self, QRectF* rect, struct miqt_string text);
extern __declspec(dllexport) 
void QPainter_DrawTextItem(QPainter* self, QPointF* p, QTextItem* ti);
extern __declspec(dllexport) 
void QPainter_DrawTextItem2(QPainter* self, int x, int y, QTextItem* ti);
extern __declspec(dllexport) 
void QPainter_DrawTextItem3(QPainter* self, QPoint* p, QTextItem* ti);
extern __declspec(dllexport) 
void QPainter_FillRect(QPainter* self, QRectF* param1, QBrush* param2);
extern __declspec(dllexport) 
void QPainter_FillRect2(QPainter* self, int x, int y, int w, int h, QBrush* param5);
extern __declspec(dllexport) 
void QPainter_FillRect3(QPainter* self, QRect* param1, QBrush* param2);
extern __declspec(dllexport) 
void QPainter_FillRect4(QPainter* self, QRectF* param1, QColor* color);
extern __declspec(dllexport) 
void QPainter_FillRect5(QPainter* self, int x, int y, int w, int h, QColor* color);
extern __declspec(dllexport) 
void QPainter_FillRect6(QPainter* self, QRect* param1, QColor* color);
extern __declspec(dllexport) 
void QPainter_FillRect7(QPainter* self, int x, int y, int w, int h, int c);
extern __declspec(dllexport) 
void QPainter_FillRect8(QPainter* self, QRect* r, int c);
extern __declspec(dllexport) 
void QPainter_FillRect9(QPainter* self, QRectF* r, int c);
extern __declspec(dllexport) 
void QPainter_FillRect10(QPainter* self, int x, int y, int w, int h, int style);
extern __declspec(dllexport) 
void QPainter_FillRect11(QPainter* self, QRect* r, int style);
extern __declspec(dllexport) 
void QPainter_FillRect12(QPainter* self, QRectF* r, int style);
extern __declspec(dllexport) 
void QPainter_FillRect13(QPainter* self, int x, int y, int w, int h, int preset);
extern __declspec(dllexport) 
void QPainter_FillRect14(QPainter* self, QRect* r, int preset);
extern __declspec(dllexport) 
void QPainter_FillRect15(QPainter* self, QRectF* r, int preset);
extern __declspec(dllexport) 
void QPainter_EraseRect(QPainter* self, QRectF* param1);
extern __declspec(dllexport) 
void QPainter_EraseRect2(QPainter* self, int x, int y, int w, int h);
extern __declspec(dllexport) 
void QPainter_EraseRectWithQRect(QPainter* self, QRect* param1);
extern __declspec(dllexport) 
void QPainter_SetRenderHint(QPainter* self, RenderHint hint);
extern __declspec(dllexport) 
void QPainter_SetRenderHints(QPainter* self, RenderHints hints);
extern __declspec(dllexport) 
RenderHints QPainter_RenderHints(const QPainter* self);
extern __declspec(dllexport) 
bool QPainter_TestRenderHint(const QPainter* self, RenderHint hint);
extern __declspec(dllexport) 
QPaintEngine* QPainter_PaintEngine(const QPainter* self);
extern __declspec(dllexport) 
void QPainter_BeginNativePainting(QPainter* self);
extern __declspec(dllexport) 
void QPainter_EndNativePainting(QPainter* self);
extern __declspec(dllexport) 
void QPainter_SetClipRect22(QPainter* self, QRectF* param1, int op);
extern __declspec(dllexport) 
void QPainter_SetClipRect23(QPainter* self, QRect* param1, int op);
extern __declspec(dllexport) 
void QPainter_SetClipRect5(QPainter* self, int x, int y, int w, int h, int op);
extern __declspec(dllexport) 
void QPainter_SetClipRegion2(QPainter* self, QRegion* param1, int op);
extern __declspec(dllexport) 
void QPainter_SetClipPath2(QPainter* self, QPainterPath* path, int op);
extern __declspec(dllexport) 
void QPainter_SetTransform2(QPainter* self, QTransform* transform, bool combine);
extern __declspec(dllexport) 
void QPainter_SetWorldTransform2(QPainter* self, QTransform* matrix, bool combine);
extern __declspec(dllexport) 
void QPainter_DrawPolygon32(QPainter* self, QPointF* points, int pointCount, int fillRule);
extern __declspec(dllexport) 
void QPainter_DrawPolygon33(QPainter* self, QPoint* points, int pointCount, int fillRule);
extern __declspec(dllexport) 
void QPainter_DrawRoundedRect4(QPainter* self, QRectF* rect, double xRadius, double yRadius, int mode);
extern __declspec(dllexport) 
void QPainter_DrawRoundedRect7(QPainter* self, int x, int y, int w, int h, double xRadius, double yRadius, int mode);
extern __declspec(dllexport) 
void QPainter_DrawRoundedRect42(QPainter* self, QRect* rect, double xRadius, double yRadius, int mode);
extern __declspec(dllexport) 
void QPainter_DrawTiledPixmap32(QPainter* self, QRectF* rect, QPixmap* pm, QPointF* offset);
extern __declspec(dllexport) 
void QPainter_DrawTiledPixmap6(QPainter* self, int x, int y, int w, int h, QPixmap* param5, int sx);
extern __declspec(dllexport) 
void QPainter_DrawTiledPixmap7(QPainter* self, int x, int y, int w, int h, QPixmap* param5, int sx, int sy);
extern __declspec(dllexport) 
void QPainter_DrawTiledPixmap33(QPainter* self, QRect* param1, QPixmap* param2, QPoint* param3);
extern __declspec(dllexport) 
void QPainter_DrawPixmapFragments4(QPainter* self, const PixmapFragment* fragments, int fragmentCount, QPixmap* pixmap, PixmapFragmentHints hints);
extern __declspec(dllexport) 
void QPainter_DrawImage42(QPainter* self, QRectF* targetRect, QImage* image, QRectF* sourceRect, int flags);
extern __declspec(dllexport) 
void QPainter_DrawImage43(QPainter* self, QRect* targetRect, QImage* image, QRect* sourceRect, int flags);
extern __declspec(dllexport) 
void QPainter_DrawImage44(QPainter* self, QPointF* p, QImage* image, QRectF* sr, int flags);
extern __declspec(dllexport) 
void QPainter_DrawImage45(QPainter* self, QPoint* p, QImage* image, QRect* sr, int flags);
extern __declspec(dllexport) 
void QPainter_DrawImage46(QPainter* self, int x, int y, QImage* image, int sx);
extern __declspec(dllexport) 
void QPainter_DrawImage52(QPainter* self, int x, int y, QImage* image, int sx, int sy);
extern __declspec(dllexport) 
void QPainter_DrawImage62(QPainter* self, int x, int y, QImage* image, int sx, int sy, int sw);
extern __declspec(dllexport) 
void QPainter_DrawImage72(QPainter* self, int x, int y, QImage* image, int sx, int sy, int sw, int sh);
extern __declspec(dllexport) 
void QPainter_DrawImage82(QPainter* self, int x, int y, QImage* image, int sx, int sy, int sw, int sh, int flags);
extern __declspec(dllexport) 
void QPainter_DrawText42(QPainter* self, QRectF* r, int flags, struct miqt_string text, QRectF* br);
extern __declspec(dllexport) 
void QPainter_DrawText43(QPainter* self, QRect* r, int flags, struct miqt_string text, QRect* br);
extern __declspec(dllexport) 
void QPainter_DrawText72(QPainter* self, int x, int y, int w, int h, int flags, struct miqt_string text, QRect* br);
extern __declspec(dllexport) 
void QPainter_DrawText32(QPainter* self, QRectF* r, struct miqt_string text, QTextOption* o);
extern __declspec(dllexport) 
QRectF* QPainter_BoundingRect32(QPainter* self, QRectF* rect, struct miqt_string text, QTextOption* o);
extern __declspec(dllexport) 
void QPainter_SetRenderHint2(QPainter* self, RenderHint hint, bool on);
extern __declspec(dllexport) 
void QPainter_SetRenderHints2(QPainter* self, RenderHints hints, bool on);
extern __declspec(dllexport) 
void QPainter_Delete(QPainter* self, bool isSubclass);

extern __declspec(dllexport) 
QPainter__PixmapFragment* QPainter__PixmapFragment_new();
extern __declspec(dllexport) 
QPainter__PixmapFragment* QPainter__PixmapFragment_new2(const PixmapFragment* param1);
extern __declspec(dllexport) 
PixmapFragment QPainter__PixmapFragment_Create(QPointF* pos, QRectF* sourceRect);
extern __declspec(dllexport) 
PixmapFragment QPainter__PixmapFragment_Create3(QPointF* pos, QRectF* sourceRect, double scaleX);
extern __declspec(dllexport) 
PixmapFragment QPainter__PixmapFragment_Create4(QPointF* pos, QRectF* sourceRect, double scaleX, double scaleY);
extern __declspec(dllexport) 
PixmapFragment QPainter__PixmapFragment_Create5(QPointF* pos, QRectF* sourceRect, double scaleX, double scaleY, double rotation);
extern __declspec(dllexport) 
PixmapFragment QPainter__PixmapFragment_Create6(QPointF* pos, QRectF* sourceRect, double scaleX, double scaleY, double rotation, double opacity);
extern __declspec(dllexport) 
void QPainter__PixmapFragment_Delete(QPainter__PixmapFragment* self, bool isSubclass);

}
