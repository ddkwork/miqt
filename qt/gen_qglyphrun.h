#pragma once
#ifndef MIQT_QT_GEN_QGLYPHRUN_H
#define MIQT_QT_GEN_QGLYPHRUN_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QGlyphRun;
class QPointF;
class QRawFont;
class QRectF;
class _GUID;
class type_info;
#else
typedef struct QGlyphRun QGlyphRun;
typedef struct QPointF QPointF;
typedef struct QRawFont QRawFont;
typedef struct QRectF QRectF;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QGlyphRun* QGlyphRun_new();
extern __declspec(dllexport) QGlyphRun* QGlyphRun_new2(QGlyphRun* other);
extern __declspec(dllexport) void QGlyphRun_OperatorAssign(QGlyphRun* self, QGlyphRun* other);
extern __declspec(dllexport) void QGlyphRun_Swap(QGlyphRun* self, QGlyphRun* other);
extern __declspec(dllexport) QRawFont* QGlyphRun_RawFont(const QGlyphRun* self);
extern __declspec(dllexport) void QGlyphRun_SetRawFont(QGlyphRun* self, QRawFont* rawFont);
extern __declspec(dllexport) void QGlyphRun_SetRawData(QGlyphRun* self, const unsigned int* glyphIndexArray, QPointF* glyphPositionArray, int size);
extern __declspec(dllexport) struct miqt_array /* of unsigned int */  QGlyphRun_GlyphIndexes(const QGlyphRun* self);
extern __declspec(dllexport) void QGlyphRun_SetGlyphIndexes(QGlyphRun* self, struct miqt_array /* of unsigned int */  glyphIndexes);
extern __declspec(dllexport) struct miqt_array /* of QPointF* */  QGlyphRun_Positions(const QGlyphRun* self);
extern __declspec(dllexport) void QGlyphRun_SetPositions(QGlyphRun* self, struct miqt_array /* of QPointF* */  positions);
extern __declspec(dllexport) void QGlyphRun_Clear(QGlyphRun* self);
extern __declspec(dllexport) bool QGlyphRun_OperatorEqual(const QGlyphRun* self, QGlyphRun* other);
extern __declspec(dllexport) bool QGlyphRun_OperatorNotEqual(const QGlyphRun* self, QGlyphRun* other);
extern __declspec(dllexport) void QGlyphRun_SetOverline(QGlyphRun* self, bool overline);
extern __declspec(dllexport) bool QGlyphRun_Overline(const QGlyphRun* self);
extern __declspec(dllexport) void QGlyphRun_SetUnderline(QGlyphRun* self, bool underline);
extern __declspec(dllexport) bool QGlyphRun_Underline(const QGlyphRun* self);
extern __declspec(dllexport) void QGlyphRun_SetStrikeOut(QGlyphRun* self, bool strikeOut);
extern __declspec(dllexport) bool QGlyphRun_StrikeOut(const QGlyphRun* self);
extern __declspec(dllexport) void QGlyphRun_SetRightToLeft(QGlyphRun* self, bool on);
extern __declspec(dllexport) bool QGlyphRun_IsRightToLeft(const QGlyphRun* self);
extern __declspec(dllexport) void QGlyphRun_SetFlag(QGlyphRun* self, GlyphRunFlag flag);
extern __declspec(dllexport) void QGlyphRun_SetFlags(QGlyphRun* self, GlyphRunFlags flags);
extern __declspec(dllexport) GlyphRunFlags QGlyphRun_Flags(const QGlyphRun* self);
extern __declspec(dllexport) void QGlyphRun_SetBoundingRect(QGlyphRun* self, QRectF* boundingRect);
extern __declspec(dllexport) QRectF* QGlyphRun_BoundingRect(const QGlyphRun* self);
extern __declspec(dllexport) struct miqt_array /* of ptrdiff_t */  QGlyphRun_StringIndexes(const QGlyphRun* self);
extern __declspec(dllexport) void QGlyphRun_SetStringIndexes(QGlyphRun* self, struct miqt_array /* of ptrdiff_t */  stringIndexes);
extern __declspec(dllexport) void QGlyphRun_SetSourceString(QGlyphRun* self, struct miqt_string sourceString);
extern __declspec(dllexport) struct miqt_string QGlyphRun_SourceString(const QGlyphRun* self);
extern __declspec(dllexport) bool QGlyphRun_IsEmpty(const QGlyphRun* self);
extern __declspec(dllexport) void QGlyphRun_SetFlag2(QGlyphRun* self, GlyphRunFlag flag, bool enabled);
extern __declspec(dllexport) void QGlyphRun_Delete(QGlyphRun* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
