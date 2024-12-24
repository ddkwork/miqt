#pragma once
#ifndef MIQT_QT_GEN_QFONTDATABASE_H
#define MIQT_QT_GEN_QFONTDATABASE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QFont QFont;
typedef struct QFontDatabase QFontDatabase;
typedef struct QFontInfo QFontInfo;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QFontDatabase* QFontDatabase_new();
extern __declspec(dllexport) struct miqt_array /* of int */  QFontDatabase_StandardSizes();
extern __declspec(dllexport) struct miqt_array /* of WritingSystem */  QFontDatabase_WritingSystems();
extern __declspec(dllexport) struct miqt_array /* of WritingSystem */  QFontDatabase_WritingSystemsWithFamily(struct miqt_string family);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QFontDatabase_Families();
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QFontDatabase_Styles(struct miqt_string family);
extern __declspec(dllexport) struct miqt_array /* of int */  QFontDatabase_PointSizes(struct miqt_string family);
extern __declspec(dllexport) struct miqt_array /* of int */  QFontDatabase_SmoothSizes(struct miqt_string family, struct miqt_string style);
extern __declspec(dllexport) struct miqt_string QFontDatabase_StyleString(QFont* font);
extern __declspec(dllexport) struct miqt_string QFontDatabase_StyleStringWithFontInfo(QFontInfo* fontInfo);
extern __declspec(dllexport) QFont* QFontDatabase_Font(struct miqt_string family, struct miqt_string style, int pointSize);
extern __declspec(dllexport) bool QFontDatabase_IsBitmapScalable(struct miqt_string family);
extern __declspec(dllexport) bool QFontDatabase_IsSmoothlyScalable(struct miqt_string family);
extern __declspec(dllexport) bool QFontDatabase_IsScalable(struct miqt_string family);
extern __declspec(dllexport) bool QFontDatabase_IsFixedPitch(struct miqt_string family);
extern __declspec(dllexport) bool QFontDatabase_Italic(struct miqt_string family, struct miqt_string style);
extern __declspec(dllexport) bool QFontDatabase_Bold(struct miqt_string family, struct miqt_string style);
extern __declspec(dllexport) int QFontDatabase_Weight(struct miqt_string family, struct miqt_string style);
extern __declspec(dllexport) bool QFontDatabase_HasFamily(struct miqt_string family);
extern __declspec(dllexport) bool QFontDatabase_IsPrivateFamily(struct miqt_string family);
extern __declspec(dllexport) struct miqt_string QFontDatabase_WritingSystemName(WritingSystem writingSystem);
extern __declspec(dllexport) struct miqt_string QFontDatabase_WritingSystemSample(WritingSystem writingSystem);
extern __declspec(dllexport) int QFontDatabase_AddApplicationFont(struct miqt_string fileName);
extern __declspec(dllexport) int QFontDatabase_AddApplicationFontFromData(struct miqt_string fontData);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QFontDatabase_ApplicationFontFamilies(int id);
extern __declspec(dllexport) bool QFontDatabase_RemoveApplicationFont(int id);
extern __declspec(dllexport) bool QFontDatabase_RemoveAllApplicationFonts();
extern __declspec(dllexport) void QFontDatabase_AddApplicationFallbackFontFamily(int script, struct miqt_string familyName);
extern __declspec(dllexport) bool QFontDatabase_RemoveApplicationFallbackFontFamily(int script, struct miqt_string familyName);
extern __declspec(dllexport) void QFontDatabase_SetApplicationFallbackFontFamilies(int param1, struct miqt_array /* of struct miqt_string */  familyNames);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QFontDatabase_ApplicationFallbackFontFamilies(int script);
extern __declspec(dllexport) void QFontDatabase_AddApplicationEmojiFontFamily(struct miqt_string familyName);
extern __declspec(dllexport) bool QFontDatabase_RemoveApplicationEmojiFontFamily(struct miqt_string familyName);
extern __declspec(dllexport) void QFontDatabase_SetApplicationEmojiFontFamilies(struct miqt_array /* of struct miqt_string */  familyNames);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QFontDatabase_ApplicationEmojiFontFamilies();
extern __declspec(dllexport) QFont* QFontDatabase_SystemFont(SystemFont typeVal);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QFontDatabase_Families1(WritingSystem writingSystem);
extern __declspec(dllexport) struct miqt_array /* of int */  QFontDatabase_PointSizes2(struct miqt_string family, struct miqt_string style);
extern __declspec(dllexport) bool QFontDatabase_IsBitmapScalable2(struct miqt_string family, struct miqt_string style);
extern __declspec(dllexport) bool QFontDatabase_IsSmoothlyScalable2(struct miqt_string family, struct miqt_string style);
extern __declspec(dllexport) bool QFontDatabase_IsScalable2(struct miqt_string family, struct miqt_string style);
extern __declspec(dllexport) bool QFontDatabase_IsFixedPitch2(struct miqt_string family, struct miqt_string style);
extern __declspec(dllexport) void QFontDatabase_Delete(QFontDatabase* self, bool isSubclass);

} 
