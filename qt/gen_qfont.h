#pragma once
#ifndef MIQT_QT_GEN_QFONT_H
#define MIQT_QT_GEN_QFONT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QFont;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QFont__Tag)
typedef QFont::Tag QFont__Tag;
#else
class QFont__Tag;
#endif
class QPaintDevice;
class _GUID;
class type_info;
#else
typedef struct QFont QFont;
typedef struct QFont__Tag QFont__Tag;
typedef struct QPaintDevice QPaintDevice;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QFont* QFont_new();
extern __declspec(dllexport) QFont* QFont_new2(struct miqt_string family);
extern __declspec(dllexport) QFont* QFont_new3(struct miqt_array /* of struct miqt_string */  families);
extern __declspec(dllexport) QFont* QFont_new4(QFont* font, QPaintDevice* pd);
extern __declspec(dllexport) QFont* QFont_new5(QFont* font);
extern __declspec(dllexport) QFont* QFont_new6(struct miqt_string family, int pointSize);
extern __declspec(dllexport) QFont* QFont_new7(struct miqt_string family, int pointSize, int weight);
extern __declspec(dllexport) QFont* QFont_new8(struct miqt_string family, int pointSize, int weight, bool italic);
extern __declspec(dllexport) QFont* QFont_new9(struct miqt_array /* of struct miqt_string */  families, int pointSize);
extern __declspec(dllexport) QFont* QFont_new10(struct miqt_array /* of struct miqt_string */  families, int pointSize, int weight);
extern __declspec(dllexport) QFont* QFont_new11(struct miqt_array /* of struct miqt_string */  families, int pointSize, int weight, bool italic);
extern __declspec(dllexport) void QFont_Swap(QFont* self, QFont* other);
extern __declspec(dllexport) struct miqt_string QFont_Family(const QFont* self);
extern __declspec(dllexport) void QFont_SetFamily(QFont* self, struct miqt_string family);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QFont_Families(const QFont* self);
extern __declspec(dllexport) void QFont_SetFamilies(QFont* self, struct miqt_array /* of struct miqt_string */  families);
extern __declspec(dllexport) struct miqt_string QFont_StyleName(const QFont* self);
extern __declspec(dllexport) void QFont_SetStyleName(QFont* self, struct miqt_string styleName);
extern __declspec(dllexport) int QFont_PointSize(const QFont* self);
extern __declspec(dllexport) void QFont_SetPointSize(QFont* self, int pointSize);
extern __declspec(dllexport) double QFont_PointSizeF(const QFont* self);
extern __declspec(dllexport) void QFont_SetPointSizeF(QFont* self, double pointSizeF);
extern __declspec(dllexport) int QFont_PixelSize(const QFont* self);
extern __declspec(dllexport) void QFont_SetPixelSize(QFont* self, int pixelSize);
extern __declspec(dllexport) Weight QFont_Weight(const QFont* self);
extern __declspec(dllexport) void QFont_SetWeight(QFont* self, Weight weight);
extern __declspec(dllexport) bool QFont_Bold(const QFont* self);
extern __declspec(dllexport) void QFont_SetBold(QFont* self, bool bold);
extern __declspec(dllexport) void QFont_SetStyle(QFont* self, Style style);
extern __declspec(dllexport) Style QFont_Style(const QFont* self);
extern __declspec(dllexport) bool QFont_Italic(const QFont* self);
extern __declspec(dllexport) void QFont_SetItalic(QFont* self, bool b);
extern __declspec(dllexport) bool QFont_Underline(const QFont* self);
extern __declspec(dllexport) void QFont_SetUnderline(QFont* self, bool underline);
extern __declspec(dllexport) bool QFont_Overline(const QFont* self);
extern __declspec(dllexport) void QFont_SetOverline(QFont* self, bool overline);
extern __declspec(dllexport) bool QFont_StrikeOut(const QFont* self);
extern __declspec(dllexport) void QFont_SetStrikeOut(QFont* self, bool strikeOut);
extern __declspec(dllexport) bool QFont_FixedPitch(const QFont* self);
extern __declspec(dllexport) void QFont_SetFixedPitch(QFont* self, bool fixedPitch);
extern __declspec(dllexport) bool QFont_Kerning(const QFont* self);
extern __declspec(dllexport) void QFont_SetKerning(QFont* self, bool kerning);
extern __declspec(dllexport) StyleHint QFont_StyleHint(const QFont* self);
extern __declspec(dllexport) StyleStrategy QFont_StyleStrategy(const QFont* self);
extern __declspec(dllexport) void QFont_SetStyleHint(QFont* self, StyleHint param1);
extern __declspec(dllexport) void QFont_SetStyleStrategy(QFont* self, StyleStrategy s);
extern __declspec(dllexport) int QFont_Stretch(const QFont* self);
extern __declspec(dllexport) void QFont_SetStretch(QFont* self, int stretch);
extern __declspec(dllexport) double QFont_LetterSpacing(const QFont* self);
extern __declspec(dllexport) SpacingType QFont_LetterSpacingType(const QFont* self);
extern __declspec(dllexport) void QFont_SetLetterSpacing(QFont* self, SpacingType typeVal, double spacing);
extern __declspec(dllexport) double QFont_WordSpacing(const QFont* self);
extern __declspec(dllexport) void QFont_SetWordSpacing(QFont* self, double spacing);
extern __declspec(dllexport) void QFont_SetCapitalization(QFont* self, Capitalization capitalization);
extern __declspec(dllexport) Capitalization QFont_Capitalization(const QFont* self);
extern __declspec(dllexport) void QFont_SetHintingPreference(QFont* self, HintingPreference hintingPreference);
extern __declspec(dllexport) HintingPreference QFont_HintingPreference(const QFont* self);
extern __declspec(dllexport) void QFont_SetFeature(QFont* self, Tag tag, unsigned int value);
extern __declspec(dllexport) void QFont_UnsetFeature(QFont* self, Tag tag);
extern __declspec(dllexport) unsigned int QFont_FeatureValue(const QFont* self, Tag tag);
extern __declspec(dllexport) bool QFont_IsFeatureSet(const QFont* self, Tag tag);
extern __declspec(dllexport) struct miqt_array /* of Tag */  QFont_FeatureTags(const QFont* self);
extern __declspec(dllexport) void QFont_ClearFeatures(QFont* self);
extern __declspec(dllexport) void QFont_SetVariableAxis(QFont* self, Tag tag, float value);
extern __declspec(dllexport) void QFont_UnsetVariableAxis(QFont* self, Tag tag);
extern __declspec(dllexport) bool QFont_IsVariableAxisSet(const QFont* self, Tag tag);
extern __declspec(dllexport) float QFont_VariableAxisValue(const QFont* self, Tag tag);
extern __declspec(dllexport) void QFont_ClearVariableAxes(QFont* self);
extern __declspec(dllexport) struct miqt_array /* of Tag */  QFont_VariableAxisTags(const QFont* self);
extern __declspec(dllexport) bool QFont_ExactMatch(const QFont* self);
extern __declspec(dllexport) void QFont_OperatorAssign(QFont* self, QFont* param1);
extern __declspec(dllexport) bool QFont_OperatorEqual(const QFont* self, QFont* param1);
extern __declspec(dllexport) bool QFont_OperatorNotEqual(const QFont* self, QFont* param1);
extern __declspec(dllexport) bool QFont_OperatorLesser(const QFont* self, QFont* param1);
extern __declspec(dllexport) bool QFont_IsCopyOf(const QFont* self, QFont* param1);
extern __declspec(dllexport) struct miqt_string QFont_Key(const QFont* self);
extern __declspec(dllexport) struct miqt_string QFont_ToString(const QFont* self);
extern __declspec(dllexport) bool QFont_FromString(QFont* self, struct miqt_string param1);
extern __declspec(dllexport) struct miqt_string QFont_Substitute(struct miqt_string param1);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QFont_Substitutes(struct miqt_string param1);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QFont_Substitutions();
extern __declspec(dllexport) void QFont_InsertSubstitution(struct miqt_string param1, struct miqt_string param2);
extern __declspec(dllexport) void QFont_InsertSubstitutions(struct miqt_string param1, struct miqt_array /* of struct miqt_string */  param2);
extern __declspec(dllexport) void QFont_RemoveSubstitutions(struct miqt_string param1);
extern __declspec(dllexport) void QFont_Initialize();
extern __declspec(dllexport) void QFont_Cleanup();
extern __declspec(dllexport) void QFont_CacheStatistics();
extern __declspec(dllexport) struct miqt_string QFont_DefaultFamily(const QFont* self);
extern __declspec(dllexport) QFont* QFont_Resolve(const QFont* self, QFont* param1);
extern __declspec(dllexport) unsigned int QFont_ResolveMask(const QFont* self);
extern __declspec(dllexport) void QFont_SetResolveMask(QFont* self, unsigned int mask);
extern __declspec(dllexport) void QFont_SetLegacyWeight(QFont* self, int legacyWeight);
extern __declspec(dllexport) int QFont_LegacyWeight(const QFont* self);
extern __declspec(dllexport) void QFont_SetStyleHint2(QFont* self, StyleHint param1, StyleStrategy param2);
extern __declspec(dllexport) void QFont_Delete(QFont* self, bool isSubclass);

extern __declspec(dllexport) QFont__Tag* QFont__Tag_new();
extern __declspec(dllexport) QFont__Tag* QFont__Tag_new2(const Tag* param1);
extern __declspec(dllexport) bool QFont__Tag_IsValid(const QFont__Tag* self);
extern __declspec(dllexport) unsigned int QFont__Tag_Value(const QFont__Tag* self);
extern __declspec(dllexport) struct miqt_string QFont__Tag_ToString(const QFont__Tag* self);
extern __declspec(dllexport) void QFont__Tag_Delete(QFont__Tag* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
