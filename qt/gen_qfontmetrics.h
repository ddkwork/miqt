#pragma once
#ifndef MIQT_QT_GEN_QFONTMETRICS_H
#define MIQT_QT_GEN_QFONTMETRICS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChar;
class QFont;
class QFontMetrics;
class QFontMetricsF;
class QPaintDevice;
class QRect;
class QRectF;
class QSize;
class QSizeF;
class QTextOption;
class _GUID;
class type_info;
#else
typedef struct QChar QChar;
typedef struct QFont QFont;
typedef struct QFontMetrics QFontMetrics;
typedef struct QFontMetricsF QFontMetricsF;
typedef struct QPaintDevice QPaintDevice;
typedef struct QRect QRect;
typedef struct QRectF QRectF;
typedef struct QSize QSize;
typedef struct QSizeF QSizeF;
typedef struct QTextOption QTextOption;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QFontMetrics* QFontMetrics_new(QFont* param1);
extern __declspec(dllexport) QFontMetrics* QFontMetrics_new2(QFont* font, QPaintDevice* pd);
extern __declspec(dllexport) QFontMetrics* QFontMetrics_new3(QFontMetrics* param1);
extern __declspec(dllexport) void QFontMetrics_OperatorAssign(QFontMetrics* self, QFontMetrics* param1);
extern __declspec(dllexport) void QFontMetrics_Swap(QFontMetrics* self, QFontMetrics* other);
extern __declspec(dllexport) int QFontMetrics_Ascent(const QFontMetrics* self);
extern __declspec(dllexport) int QFontMetrics_CapHeight(const QFontMetrics* self);
extern __declspec(dllexport) int QFontMetrics_Descent(const QFontMetrics* self);
extern __declspec(dllexport) int QFontMetrics_Height(const QFontMetrics* self);
extern __declspec(dllexport) int QFontMetrics_Leading(const QFontMetrics* self);
extern __declspec(dllexport) int QFontMetrics_LineSpacing(const QFontMetrics* self);
extern __declspec(dllexport) int QFontMetrics_MinLeftBearing(const QFontMetrics* self);
extern __declspec(dllexport) int QFontMetrics_MinRightBearing(const QFontMetrics* self);
extern __declspec(dllexport) int QFontMetrics_MaxWidth(const QFontMetrics* self);
extern __declspec(dllexport) int QFontMetrics_XHeight(const QFontMetrics* self);
extern __declspec(dllexport) int QFontMetrics_AverageCharWidth(const QFontMetrics* self);
extern __declspec(dllexport) bool QFontMetrics_InFont(const QFontMetrics* self, QChar* param1);
extern __declspec(dllexport) bool QFontMetrics_InFontUcs4(const QFontMetrics* self, unsigned int ucs4);
extern __declspec(dllexport) int QFontMetrics_LeftBearing(const QFontMetrics* self, QChar* param1);
extern __declspec(dllexport) int QFontMetrics_RightBearing(const QFontMetrics* self, QChar* param1);
extern __declspec(dllexport) int QFontMetrics_HorizontalAdvance(const QFontMetrics* self, struct miqt_string param1);
extern __declspec(dllexport) int QFontMetrics_HorizontalAdvance2(const QFontMetrics* self, struct miqt_string param1, QTextOption* textOption);
extern __declspec(dllexport) int QFontMetrics_HorizontalAdvanceWithQChar(const QFontMetrics* self, QChar* param1);
extern __declspec(dllexport) QRect* QFontMetrics_BoundingRect(const QFontMetrics* self, QChar* param1);
extern __declspec(dllexport) QRect* QFontMetrics_BoundingRectWithText(const QFontMetrics* self, struct miqt_string text);
extern __declspec(dllexport) QRect* QFontMetrics_BoundingRect2(const QFontMetrics* self, struct miqt_string text, QTextOption* textOption);
extern __declspec(dllexport) QRect* QFontMetrics_BoundingRect3(const QFontMetrics* self, QRect* r, int flags, struct miqt_string text);
extern __declspec(dllexport) QRect* QFontMetrics_BoundingRect4(const QFontMetrics* self, int x, int y, int w, int h, int flags, struct miqt_string text);
extern __declspec(dllexport) QSize* QFontMetrics_Size(const QFontMetrics* self, int flags, struct miqt_string str);
extern __declspec(dllexport) QRect* QFontMetrics_TightBoundingRect(const QFontMetrics* self, struct miqt_string text);
extern __declspec(dllexport) QRect* QFontMetrics_TightBoundingRect2(const QFontMetrics* self, struct miqt_string text, QTextOption* textOption);
extern __declspec(dllexport) struct miqt_string QFontMetrics_ElidedText(const QFontMetrics* self, struct miqt_string text, int mode, int width);
extern __declspec(dllexport) int QFontMetrics_UnderlinePos(const QFontMetrics* self);
extern __declspec(dllexport) int QFontMetrics_OverlinePos(const QFontMetrics* self);
extern __declspec(dllexport) int QFontMetrics_StrikeOutPos(const QFontMetrics* self);
extern __declspec(dllexport) int QFontMetrics_LineWidth(const QFontMetrics* self);
extern __declspec(dllexport) double QFontMetrics_FontDpi(const QFontMetrics* self);
extern __declspec(dllexport) bool QFontMetrics_OperatorEqual(const QFontMetrics* self, QFontMetrics* other);
extern __declspec(dllexport) bool QFontMetrics_OperatorNotEqual(const QFontMetrics* self, QFontMetrics* other);
extern __declspec(dllexport) int QFontMetrics_HorizontalAdvance22(const QFontMetrics* self, struct miqt_string param1, int lenVal);
extern __declspec(dllexport) QRect* QFontMetrics_BoundingRect42(const QFontMetrics* self, QRect* r, int flags, struct miqt_string text, int tabstops);
extern __declspec(dllexport) QRect* QFontMetrics_BoundingRect5(const QFontMetrics* self, QRect* r, int flags, struct miqt_string text, int tabstops, int* tabarray);
extern __declspec(dllexport) QRect* QFontMetrics_BoundingRect7(const QFontMetrics* self, int x, int y, int w, int h, int flags, struct miqt_string text, int tabstops);
extern __declspec(dllexport) QRect* QFontMetrics_BoundingRect8(const QFontMetrics* self, int x, int y, int w, int h, int flags, struct miqt_string text, int tabstops, int* tabarray);
extern __declspec(dllexport) QSize* QFontMetrics_Size3(const QFontMetrics* self, int flags, struct miqt_string str, int tabstops);
extern __declspec(dllexport) QSize* QFontMetrics_Size4(const QFontMetrics* self, int flags, struct miqt_string str, int tabstops, int* tabarray);
extern __declspec(dllexport) struct miqt_string QFontMetrics_ElidedText4(const QFontMetrics* self, struct miqt_string text, int mode, int width, int flags);
extern __declspec(dllexport) void QFontMetrics_Delete(QFontMetrics* self, bool isSubclass);

extern __declspec(dllexport) QFontMetricsF* QFontMetricsF_new(QFont* font);
extern __declspec(dllexport) QFontMetricsF* QFontMetricsF_new2(QFont* font, QPaintDevice* pd);
extern __declspec(dllexport) QFontMetricsF* QFontMetricsF_new3(QFontMetrics* param1);
extern __declspec(dllexport) QFontMetricsF* QFontMetricsF_new4(QFontMetricsF* param1);
extern __declspec(dllexport) void QFontMetricsF_OperatorAssign(QFontMetricsF* self, QFontMetricsF* param1);
extern __declspec(dllexport) void QFontMetricsF_OperatorAssignWithQFontMetrics(QFontMetricsF* self, QFontMetrics* param1);
extern __declspec(dllexport) void QFontMetricsF_Swap(QFontMetricsF* self, QFontMetricsF* other);
extern __declspec(dllexport) double QFontMetricsF_Ascent(const QFontMetricsF* self);
extern __declspec(dllexport) double QFontMetricsF_CapHeight(const QFontMetricsF* self);
extern __declspec(dllexport) double QFontMetricsF_Descent(const QFontMetricsF* self);
extern __declspec(dllexport) double QFontMetricsF_Height(const QFontMetricsF* self);
extern __declspec(dllexport) double QFontMetricsF_Leading(const QFontMetricsF* self);
extern __declspec(dllexport) double QFontMetricsF_LineSpacing(const QFontMetricsF* self);
extern __declspec(dllexport) double QFontMetricsF_MinLeftBearing(const QFontMetricsF* self);
extern __declspec(dllexport) double QFontMetricsF_MinRightBearing(const QFontMetricsF* self);
extern __declspec(dllexport) double QFontMetricsF_MaxWidth(const QFontMetricsF* self);
extern __declspec(dllexport) double QFontMetricsF_XHeight(const QFontMetricsF* self);
extern __declspec(dllexport) double QFontMetricsF_AverageCharWidth(const QFontMetricsF* self);
extern __declspec(dllexport) bool QFontMetricsF_InFont(const QFontMetricsF* self, QChar* param1);
extern __declspec(dllexport) bool QFontMetricsF_InFontUcs4(const QFontMetricsF* self, unsigned int ucs4);
extern __declspec(dllexport) double QFontMetricsF_LeftBearing(const QFontMetricsF* self, QChar* param1);
extern __declspec(dllexport) double QFontMetricsF_RightBearing(const QFontMetricsF* self, QChar* param1);
extern __declspec(dllexport) double QFontMetricsF_HorizontalAdvance(const QFontMetricsF* self, struct miqt_string stringVal);
extern __declspec(dllexport) double QFontMetricsF_HorizontalAdvanceWithQChar(const QFontMetricsF* self, QChar* param1);
extern __declspec(dllexport) double QFontMetricsF_HorizontalAdvance2(const QFontMetricsF* self, struct miqt_string stringVal, QTextOption* textOption);
extern __declspec(dllexport) QRectF* QFontMetricsF_BoundingRect(const QFontMetricsF* self, struct miqt_string stringVal);
extern __declspec(dllexport) QRectF* QFontMetricsF_BoundingRect2(const QFontMetricsF* self, struct miqt_string text, QTextOption* textOption);
extern __declspec(dllexport) QRectF* QFontMetricsF_BoundingRectWithQChar(const QFontMetricsF* self, QChar* param1);
extern __declspec(dllexport) QRectF* QFontMetricsF_BoundingRect3(const QFontMetricsF* self, QRectF* r, int flags, struct miqt_string stringVal);
extern __declspec(dllexport) QSizeF* QFontMetricsF_Size(const QFontMetricsF* self, int flags, struct miqt_string str);
extern __declspec(dllexport) QRectF* QFontMetricsF_TightBoundingRect(const QFontMetricsF* self, struct miqt_string text);
extern __declspec(dllexport) QRectF* QFontMetricsF_TightBoundingRect2(const QFontMetricsF* self, struct miqt_string text, QTextOption* textOption);
extern __declspec(dllexport) struct miqt_string QFontMetricsF_ElidedText(const QFontMetricsF* self, struct miqt_string text, int mode, double width);
extern __declspec(dllexport) double QFontMetricsF_UnderlinePos(const QFontMetricsF* self);
extern __declspec(dllexport) double QFontMetricsF_OverlinePos(const QFontMetricsF* self);
extern __declspec(dllexport) double QFontMetricsF_StrikeOutPos(const QFontMetricsF* self);
extern __declspec(dllexport) double QFontMetricsF_LineWidth(const QFontMetricsF* self);
extern __declspec(dllexport) double QFontMetricsF_FontDpi(const QFontMetricsF* self);
extern __declspec(dllexport) bool QFontMetricsF_OperatorEqual(const QFontMetricsF* self, QFontMetricsF* other);
extern __declspec(dllexport) bool QFontMetricsF_OperatorNotEqual(const QFontMetricsF* self, QFontMetricsF* other);
extern __declspec(dllexport) double QFontMetricsF_HorizontalAdvance22(const QFontMetricsF* self, struct miqt_string stringVal, int length);
extern __declspec(dllexport) QRectF* QFontMetricsF_BoundingRect4(const QFontMetricsF* self, QRectF* r, int flags, struct miqt_string stringVal, int tabstops);
extern __declspec(dllexport) QRectF* QFontMetricsF_BoundingRect5(const QFontMetricsF* self, QRectF* r, int flags, struct miqt_string stringVal, int tabstops, int* tabarray);
extern __declspec(dllexport) QSizeF* QFontMetricsF_Size3(const QFontMetricsF* self, int flags, struct miqt_string str, int tabstops);
extern __declspec(dllexport) QSizeF* QFontMetricsF_Size4(const QFontMetricsF* self, int flags, struct miqt_string str, int tabstops, int* tabarray);
extern __declspec(dllexport) struct miqt_string QFontMetricsF_ElidedText4(const QFontMetricsF* self, struct miqt_string text, int mode, double width, int flags);
extern __declspec(dllexport) void QFontMetricsF_Delete(QFontMetricsF* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
