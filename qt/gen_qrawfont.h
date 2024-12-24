#pragma once
#ifndef MIQT_QT_GEN_QRAWFONT_H
#define MIQT_QT_GEN_QRAWFONT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QFont__Tag)
typedef QFont::Tag QFont__Tag;
typedef struct QChar QChar;
typedef struct QFont QFont;
typedef struct QFont__Tag QFont__Tag;
typedef struct QImage QImage;
typedef struct QPainterPath QPainterPath;
typedef struct QPointF QPointF;
typedef struct QRawFont QRawFont;
typedef struct QRectF QRectF;
typedef struct QTransform QTransform;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QRawFont* QRawFont_new();
extern __declspec(dllexport) QRawFont* QRawFont_new2(struct miqt_string fileName, double pixelSize);
extern __declspec(dllexport) QRawFont* QRawFont_new3(struct miqt_string fontData, double pixelSize);
extern __declspec(dllexport) QRawFont* QRawFont_new4(QRawFont* other);
extern __declspec(dllexport) QRawFont* QRawFont_new5(struct miqt_string fileName, double pixelSize, int hintingPreference);
extern __declspec(dllexport) QRawFont* QRawFont_new6(struct miqt_string fontData, double pixelSize, int hintingPreference);
extern __declspec(dllexport) void QRawFont_OperatorAssign(QRawFont* self, QRawFont* other);
extern __declspec(dllexport) void QRawFont_Swap(QRawFont* self, QRawFont* other);
extern __declspec(dllexport) bool QRawFont_IsValid(const QRawFont* self);
extern __declspec(dllexport) bool QRawFont_OperatorEqual(const QRawFont* self, QRawFont* other);
extern __declspec(dllexport) bool QRawFont_OperatorNotEqual(const QRawFont* self, QRawFont* other);
extern __declspec(dllexport) struct miqt_string QRawFont_FamilyName(const QRawFont* self);
extern __declspec(dllexport) struct miqt_string QRawFont_StyleName(const QRawFont* self);
extern __declspec(dllexport) int QRawFont_Style(const QRawFont* self);
extern __declspec(dllexport) int QRawFont_Weight(const QRawFont* self);
extern __declspec(dllexport) struct miqt_array /* of unsigned int */  QRawFont_GlyphIndexesForString(const QRawFont* self, struct miqt_string text);
extern __declspec(dllexport) struct miqt_array /* of QPointF* */  QRawFont_AdvancesForGlyphIndexes(const QRawFont* self, struct miqt_array /* of unsigned int */  glyphIndexes);
extern __declspec(dllexport) struct miqt_array /* of QPointF* */  QRawFont_AdvancesForGlyphIndexes2(const QRawFont* self, struct miqt_array /* of unsigned int */  glyphIndexes, LayoutFlags layoutFlags);
extern __declspec(dllexport) bool QRawFont_GlyphIndexesForChars(const QRawFont* self, QChar* chars, int numChars, unsigned int* glyphIndexes, int* numGlyphs);
extern __declspec(dllexport) bool QRawFont_AdvancesForGlyphIndexes3(const QRawFont* self, const unsigned int* glyphIndexes, QPointF* advances, int numGlyphs);
extern __declspec(dllexport) bool QRawFont_AdvancesForGlyphIndexes4(const QRawFont* self, const unsigned int* glyphIndexes, QPointF* advances, int numGlyphs, LayoutFlags layoutFlags);
extern __declspec(dllexport) QImage* QRawFont_AlphaMapForGlyph(const QRawFont* self, unsigned int glyphIndex);
extern __declspec(dllexport) QPainterPath* QRawFont_PathForGlyph(const QRawFont* self, unsigned int glyphIndex);
extern __declspec(dllexport) QRectF* QRawFont_BoundingRect(const QRawFont* self, unsigned int glyphIndex);
extern __declspec(dllexport) void QRawFont_SetPixelSize(QRawFont* self, double pixelSize);
extern __declspec(dllexport) double QRawFont_PixelSize(const QRawFont* self);
extern __declspec(dllexport) int QRawFont_HintingPreference(const QRawFont* self);
extern __declspec(dllexport) double QRawFont_Ascent(const QRawFont* self);
extern __declspec(dllexport) double QRawFont_CapHeight(const QRawFont* self);
extern __declspec(dllexport) double QRawFont_Descent(const QRawFont* self);
extern __declspec(dllexport) double QRawFont_Leading(const QRawFont* self);
extern __declspec(dllexport) double QRawFont_XHeight(const QRawFont* self);
extern __declspec(dllexport) double QRawFont_AverageCharWidth(const QRawFont* self);
extern __declspec(dllexport) double QRawFont_MaxCharWidth(const QRawFont* self);
extern __declspec(dllexport) double QRawFont_LineThickness(const QRawFont* self);
extern __declspec(dllexport) double QRawFont_UnderlinePosition(const QRawFont* self);
extern __declspec(dllexport) double QRawFont_UnitsPerEm(const QRawFont* self);
extern __declspec(dllexport) void QRawFont_LoadFromFile(QRawFont* self, struct miqt_string fileName, double pixelSize, int hintingPreference);
extern __declspec(dllexport) void QRawFont_LoadFromData(QRawFont* self, struct miqt_string fontData, double pixelSize, int hintingPreference);
extern __declspec(dllexport) bool QRawFont_SupportsCharacter(const QRawFont* self, unsigned int ucs4);
extern __declspec(dllexport) bool QRawFont_SupportsCharacterWithCharacter(const QRawFont* self, QChar* character);
extern __declspec(dllexport) struct miqt_array /* of int */  QRawFont_SupportedWritingSystems(const QRawFont* self);
extern __declspec(dllexport) struct miqt_string QRawFont_FontTable(const QRawFont* self, const char* tagName);
extern __declspec(dllexport) struct miqt_string QRawFont_FontTableWithTag(const QRawFont* self, QFont__Tag* tag);
extern __declspec(dllexport) QRawFont* QRawFont_FromFont(QFont* font);
extern __declspec(dllexport) QImage* QRawFont_AlphaMapForGlyph2(const QRawFont* self, unsigned int glyphIndex, AntialiasingType antialiasingType);
extern __declspec(dllexport) QImage* QRawFont_AlphaMapForGlyph3(const QRawFont* self, unsigned int glyphIndex, AntialiasingType antialiasingType, QTransform* transform);
extern __declspec(dllexport) QRawFont* QRawFont_FromFont2(QFont* font, int writingSystem);
extern __declspec(dllexport) void QRawFont_Delete(QRawFont* self, bool isSubclass);

} 
