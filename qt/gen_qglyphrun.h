#ifndef GEN_QGLYPHRUN_H
#define GEN_QGLYPHRUN_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QGlyphRun;
class QPointF;
class QRawFont;
class QRectF;
#else
typedef struct QGlyphRun QGlyphRun;
typedef struct QPointF QPointF;
typedef struct QRawFont QRawFont;
typedef struct QRectF QRectF;
#endif

QGlyphRun* QGlyphRun_new();
QGlyphRun* QGlyphRun_new2(QGlyphRun* other);
void QGlyphRun_OperatorAssign(QGlyphRun* self, QGlyphRun* other);
void QGlyphRun_Swap(QGlyphRun* self, QGlyphRun* other);
QRawFont* QGlyphRun_RawFont(const QGlyphRun* self);
void QGlyphRun_SetRawFont(QGlyphRun* self, QRawFont* rawFont);
void QGlyphRun_SetRawData(QGlyphRun* self, const unsigned int* glyphIndexArray, QPointF* glyphPositionArray, int size);
void QGlyphRun_GlyphIndexes(const QGlyphRun* self, unsigned int** _out, size_t* _out_len);
void QGlyphRun_SetGlyphIndexes(QGlyphRun* self, unsigned int* glyphIndexes, size_t glyphIndexes_len);
void QGlyphRun_Positions(const QGlyphRun* self, QPointF*** _out, size_t* _out_len);
void QGlyphRun_SetPositions(QGlyphRun* self, QPointF** positions, size_t positions_len);
void QGlyphRun_Clear(QGlyphRun* self);
bool QGlyphRun_OperatorEqual(const QGlyphRun* self, QGlyphRun* other);
bool QGlyphRun_OperatorNotEqual(const QGlyphRun* self, QGlyphRun* other);
void QGlyphRun_SetOverline(QGlyphRun* self, bool overline);
bool QGlyphRun_Overline(const QGlyphRun* self);
void QGlyphRun_SetUnderline(QGlyphRun* self, bool underline);
bool QGlyphRun_Underline(const QGlyphRun* self);
void QGlyphRun_SetStrikeOut(QGlyphRun* self, bool strikeOut);
bool QGlyphRun_StrikeOut(const QGlyphRun* self);
void QGlyphRun_SetRightToLeft(QGlyphRun* self, bool on);
bool QGlyphRun_IsRightToLeft(const QGlyphRun* self);
void QGlyphRun_SetFlag(QGlyphRun* self, uintptr_t flag);
void QGlyphRun_SetFlags(QGlyphRun* self, int flags);
int QGlyphRun_Flags(const QGlyphRun* self);
void QGlyphRun_SetBoundingRect(QGlyphRun* self, QRectF* boundingRect);
QRectF* QGlyphRun_BoundingRect(const QGlyphRun* self);
bool QGlyphRun_IsEmpty(const QGlyphRun* self);
void QGlyphRun_SetFlag2(QGlyphRun* self, uintptr_t flag, bool enabled);
void QGlyphRun_Delete(QGlyphRun* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
