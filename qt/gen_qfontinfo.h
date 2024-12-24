#pragma once
#ifndef MIQT_QT_GEN_QFONTINFO_H
#define MIQT_QT_GEN_QFONTINFO_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QFont QFont;
typedef struct QFontInfo QFontInfo;
typedef struct QFontVariableAxis QFontVariableAxis;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QFontInfo* QFontInfo_new(QFont* param1);
extern __declspec(dllexport) QFontInfo* QFontInfo_new2(QFontInfo* param1);
extern __declspec(dllexport) void QFontInfo_OperatorAssign(QFontInfo* self, QFontInfo* param1);
extern __declspec(dllexport) void QFontInfo_Swap(QFontInfo* self, QFontInfo* other);
extern __declspec(dllexport) struct miqt_string QFontInfo_Family(const QFontInfo* self);
extern __declspec(dllexport) struct miqt_string QFontInfo_StyleName(const QFontInfo* self);
extern __declspec(dllexport) int QFontInfo_PixelSize(const QFontInfo* self);
extern __declspec(dllexport) int QFontInfo_PointSize(const QFontInfo* self);
extern __declspec(dllexport) double QFontInfo_PointSizeF(const QFontInfo* self);
extern __declspec(dllexport) bool QFontInfo_Italic(const QFontInfo* self);
extern __declspec(dllexport) int QFontInfo_Style(const QFontInfo* self);
extern __declspec(dllexport) int QFontInfo_Weight(const QFontInfo* self);
extern __declspec(dllexport) bool QFontInfo_Bold(const QFontInfo* self);
extern __declspec(dllexport) bool QFontInfo_Underline(const QFontInfo* self);
extern __declspec(dllexport) bool QFontInfo_Overline(const QFontInfo* self);
extern __declspec(dllexport) bool QFontInfo_StrikeOut(const QFontInfo* self);
extern __declspec(dllexport) bool QFontInfo_FixedPitch(const QFontInfo* self);
extern __declspec(dllexport) int QFontInfo_StyleHint(const QFontInfo* self);
extern __declspec(dllexport) struct miqt_array /* of QFontVariableAxis* */  QFontInfo_VariableAxes(const QFontInfo* self);
extern __declspec(dllexport) int QFontInfo_LegacyWeight(const QFontInfo* self);
extern __declspec(dllexport) bool QFontInfo_ExactMatch(const QFontInfo* self);
extern __declspec(dllexport) void QFontInfo_Delete(QFontInfo* self, bool isSubclass);

} 
