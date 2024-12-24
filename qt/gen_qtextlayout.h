#pragma once
#ifndef MIQT_QT_GEN_QTEXTLAYOUT_H
#define MIQT_QT_GEN_QTEXTLAYOUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QTextLayout__FormatRange)
typedef QTextLayout::FormatRange QTextLayout__FormatRange;
typedef struct QFont QFont;
typedef struct QGlyphRun QGlyphRun;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPainter QPainter;
typedef struct QPointF QPointF;
typedef struct QRawFont QRawFont;
typedef struct QRectF QRectF;
typedef struct QTextBlock QTextBlock;
typedef struct QTextFormat QTextFormat;
typedef struct QTextInlineObject QTextInlineObject;
typedef struct QTextLayout QTextLayout;
typedef struct QTextLayout__FormatRange QTextLayout__FormatRange;
typedef struct QTextLine QTextLine;
typedef struct QTextOption QTextOption;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTextInlineObject* QTextInlineObject_new();
extern __declspec(dllexport) bool QTextInlineObject_IsValid(const QTextInlineObject* self);
extern __declspec(dllexport) QRectF* QTextInlineObject_Rect(const QTextInlineObject* self);
extern __declspec(dllexport) double QTextInlineObject_Width(const QTextInlineObject* self);
extern __declspec(dllexport) double QTextInlineObject_Ascent(const QTextInlineObject* self);
extern __declspec(dllexport) double QTextInlineObject_Descent(const QTextInlineObject* self);
extern __declspec(dllexport) double QTextInlineObject_Height(const QTextInlineObject* self);
extern __declspec(dllexport) int QTextInlineObject_TextDirection(const QTextInlineObject* self);
extern __declspec(dllexport) void QTextInlineObject_SetWidth(QTextInlineObject* self, double w);
extern __declspec(dllexport) void QTextInlineObject_SetAscent(QTextInlineObject* self, double a);
extern __declspec(dllexport) void QTextInlineObject_SetDescent(QTextInlineObject* self, double d);
extern __declspec(dllexport) int QTextInlineObject_TextPosition(const QTextInlineObject* self);
extern __declspec(dllexport) int QTextInlineObject_FormatIndex(const QTextInlineObject* self);
extern __declspec(dllexport) QTextFormat* QTextInlineObject_Format(const QTextInlineObject* self);
extern __declspec(dllexport) void QTextInlineObject_Delete(QTextInlineObject* self, bool isSubclass);

extern __declspec(dllexport) QTextLayout* QTextLayout_new();
extern __declspec(dllexport) QTextLayout* QTextLayout_new2(struct miqt_string text);
extern __declspec(dllexport) QTextLayout* QTextLayout_new3(struct miqt_string text, QFont* font);
extern __declspec(dllexport) QTextLayout* QTextLayout_new4(QTextBlock* b);
extern __declspec(dllexport) QTextLayout* QTextLayout_new5(struct miqt_string text, QFont* font, QPaintDevice* paintdevice);
extern __declspec(dllexport) void QTextLayout_SetFont(QTextLayout* self, QFont* f);
extern __declspec(dllexport) QFont* QTextLayout_Font(const QTextLayout* self);
extern __declspec(dllexport) void QTextLayout_SetRawFont(QTextLayout* self, QRawFont* rawFont);
extern __declspec(dllexport) void QTextLayout_SetText(QTextLayout* self, struct miqt_string stringVal);
extern __declspec(dllexport) struct miqt_string QTextLayout_Text(const QTextLayout* self);
extern __declspec(dllexport) void QTextLayout_SetTextOption(QTextLayout* self, QTextOption* option);
extern __declspec(dllexport) QTextOption* QTextLayout_TextOption(const QTextLayout* self);
extern __declspec(dllexport) void QTextLayout_SetPreeditArea(QTextLayout* self, int position, struct miqt_string text);
extern __declspec(dllexport) int QTextLayout_PreeditAreaPosition(const QTextLayout* self);
extern __declspec(dllexport) struct miqt_string QTextLayout_PreeditAreaText(const QTextLayout* self);
extern __declspec(dllexport) void QTextLayout_SetFormats(QTextLayout* self, struct miqt_array /* of FormatRange */  overrides);
extern __declspec(dllexport) struct miqt_array /* of FormatRange */  QTextLayout_Formats(const QTextLayout* self);
extern __declspec(dllexport) void QTextLayout_ClearFormats(QTextLayout* self);
extern __declspec(dllexport) void QTextLayout_SetCacheEnabled(QTextLayout* self, bool enable);
extern __declspec(dllexport) bool QTextLayout_CacheEnabled(const QTextLayout* self);
extern __declspec(dllexport) void QTextLayout_SetCursorMoveStyle(QTextLayout* self, int style);
extern __declspec(dllexport) int QTextLayout_CursorMoveStyle(const QTextLayout* self);
extern __declspec(dllexport) void QTextLayout_BeginLayout(QTextLayout* self);
extern __declspec(dllexport) void QTextLayout_EndLayout(QTextLayout* self);
extern __declspec(dllexport) void QTextLayout_ClearLayout(QTextLayout* self);
extern __declspec(dllexport) QTextLine* QTextLayout_CreateLine(QTextLayout* self);
extern __declspec(dllexport) int QTextLayout_LineCount(const QTextLayout* self);
extern __declspec(dllexport) QTextLine* QTextLayout_LineAt(const QTextLayout* self, int i);
extern __declspec(dllexport) QTextLine* QTextLayout_LineForTextPosition(const QTextLayout* self, int pos);
extern __declspec(dllexport) bool QTextLayout_IsValidCursorPosition(const QTextLayout* self, int pos);
extern __declspec(dllexport) int QTextLayout_NextCursorPosition(const QTextLayout* self, int oldPos);
extern __declspec(dllexport) int QTextLayout_PreviousCursorPosition(const QTextLayout* self, int oldPos);
extern __declspec(dllexport) int QTextLayout_LeftCursorPosition(const QTextLayout* self, int oldPos);
extern __declspec(dllexport) int QTextLayout_RightCursorPosition(const QTextLayout* self, int oldPos);
extern __declspec(dllexport) void QTextLayout_Draw(const QTextLayout* self, QPainter* p, QPointF* pos);
extern __declspec(dllexport) void QTextLayout_DrawCursor(const QTextLayout* self, QPainter* p, QPointF* pos, int cursorPosition);
extern __declspec(dllexport) void QTextLayout_DrawCursor2(const QTextLayout* self, QPainter* p, QPointF* pos, int cursorPosition, int width);
extern __declspec(dllexport) QPointF* QTextLayout_Position(const QTextLayout* self);
extern __declspec(dllexport) void QTextLayout_SetPosition(QTextLayout* self, QPointF* p);
extern __declspec(dllexport) QRectF* QTextLayout_BoundingRect(const QTextLayout* self);
extern __declspec(dllexport) double QTextLayout_MinimumWidth(const QTextLayout* self);
extern __declspec(dllexport) double QTextLayout_MaximumWidth(const QTextLayout* self);
extern __declspec(dllexport) struct miqt_array /* of QGlyphRun* */  QTextLayout_GlyphRuns(const QTextLayout* self, int from, int length, GlyphRunRetrievalFlags flags);
extern __declspec(dllexport) struct miqt_array /* of QGlyphRun* */  QTextLayout_GlyphRuns2(const QTextLayout* self);
extern __declspec(dllexport) void QTextLayout_SetFlags(QTextLayout* self, int flags);
extern __declspec(dllexport) int QTextLayout_NextCursorPosition2(const QTextLayout* self, int oldPos, CursorMode mode);
extern __declspec(dllexport) int QTextLayout_PreviousCursorPosition2(const QTextLayout* self, int oldPos, CursorMode mode);
extern __declspec(dllexport) void QTextLayout_Draw3(const QTextLayout* self, QPainter* p, QPointF* pos, struct miqt_array /* of FormatRange */  selections);
extern __declspec(dllexport) void QTextLayout_Draw4(const QTextLayout* self, QPainter* p, QPointF* pos, struct miqt_array /* of FormatRange */  selections, QRectF* clip);
extern __declspec(dllexport) struct miqt_array /* of QGlyphRun* */  QTextLayout_GlyphRuns1(const QTextLayout* self, int from);
extern __declspec(dllexport) struct miqt_array /* of QGlyphRun* */  QTextLayout_GlyphRuns22(const QTextLayout* self, int from, int length);
extern __declspec(dllexport) void QTextLayout_Delete(QTextLayout* self, bool isSubclass);

extern __declspec(dllexport) QTextLine* QTextLine_new();
extern __declspec(dllexport) bool QTextLine_IsValid(const QTextLine* self);
extern __declspec(dllexport) QRectF* QTextLine_Rect(const QTextLine* self);
extern __declspec(dllexport) double QTextLine_X(const QTextLine* self);
extern __declspec(dllexport) double QTextLine_Y(const QTextLine* self);
extern __declspec(dllexport) double QTextLine_Width(const QTextLine* self);
extern __declspec(dllexport) double QTextLine_Ascent(const QTextLine* self);
extern __declspec(dllexport) double QTextLine_Descent(const QTextLine* self);
extern __declspec(dllexport) double QTextLine_Height(const QTextLine* self);
extern __declspec(dllexport) double QTextLine_Leading(const QTextLine* self);
extern __declspec(dllexport) void QTextLine_SetLeadingIncluded(QTextLine* self, bool included);
extern __declspec(dllexport) bool QTextLine_LeadingIncluded(const QTextLine* self);
extern __declspec(dllexport) double QTextLine_NaturalTextWidth(const QTextLine* self);
extern __declspec(dllexport) double QTextLine_HorizontalAdvance(const QTextLine* self);
extern __declspec(dllexport) QRectF* QTextLine_NaturalTextRect(const QTextLine* self);
extern __declspec(dllexport) double QTextLine_CursorToX(const QTextLine* self, int* cursorPos);
extern __declspec(dllexport) double QTextLine_CursorToXWithCursorPos(const QTextLine* self, int cursorPos);
extern __declspec(dllexport) int QTextLine_XToCursor(const QTextLine* self, double x);
extern __declspec(dllexport) void QTextLine_SetLineWidth(QTextLine* self, double width);
extern __declspec(dllexport) void QTextLine_SetNumColumns(QTextLine* self, int columns);
extern __declspec(dllexport) void QTextLine_SetNumColumns2(QTextLine* self, int columns, double alignmentWidth);
extern __declspec(dllexport) void QTextLine_SetPosition(QTextLine* self, QPointF* pos);
extern __declspec(dllexport) QPointF* QTextLine_Position(const QTextLine* self);
extern __declspec(dllexport) int QTextLine_TextStart(const QTextLine* self);
extern __declspec(dllexport) int QTextLine_TextLength(const QTextLine* self);
extern __declspec(dllexport) int QTextLine_LineNumber(const QTextLine* self);
extern __declspec(dllexport) void QTextLine_Draw(const QTextLine* self, QPainter* painter, QPointF* position);
extern __declspec(dllexport) struct miqt_array /* of QGlyphRun* */  QTextLine_GlyphRuns(const QTextLine* self, int from, int length, int flags);
extern __declspec(dllexport) struct miqt_array /* of QGlyphRun* */  QTextLine_GlyphRuns2(const QTextLine* self);
extern __declspec(dllexport) double QTextLine_CursorToX2(const QTextLine* self, int* cursorPos, Edge edge);
extern __declspec(dllexport) double QTextLine_CursorToX22(const QTextLine* self, int cursorPos, Edge edge);
extern __declspec(dllexport) int QTextLine_XToCursor2(const QTextLine* self, double x, CursorPosition param2);
extern __declspec(dllexport) struct miqt_array /* of QGlyphRun* */  QTextLine_GlyphRuns1(const QTextLine* self, int from);
extern __declspec(dllexport) struct miqt_array /* of QGlyphRun* */  QTextLine_GlyphRuns22(const QTextLine* self, int from, int length);
extern __declspec(dllexport) void QTextLine_Delete(QTextLine* self, bool isSubclass);

extern __declspec(dllexport) QTextLayout__FormatRange* QTextLayout__FormatRange_new();
extern __declspec(dllexport) void QTextLayout__FormatRange_Delete(QTextLayout__FormatRange* self, bool isSubclass);

} 
