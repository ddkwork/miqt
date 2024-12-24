#pragma once
#ifndef MIQT_QT_GEN_QREGULAREXPRESSION_H
#define MIQT_QT_GEN_QREGULAREXPRESSION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAnyStringView QAnyStringView;
typedef struct QRegularExpression QRegularExpression;
typedef struct QRegularExpressionMatch QRegularExpressionMatch;
typedef struct QRegularExpressionMatchIterator QRegularExpressionMatchIterator;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QRegularExpression* QRegularExpression_new();
extern __declspec(dllexport) QRegularExpression* QRegularExpression_new2(struct miqt_string pattern);
extern __declspec(dllexport) QRegularExpression* QRegularExpression_new3(QRegularExpression* re);
extern __declspec(dllexport) QRegularExpression* QRegularExpression_new4(struct miqt_string pattern, PatternOptions options);
extern __declspec(dllexport) PatternOptions QRegularExpression_PatternOptions(const QRegularExpression* self);
extern __declspec(dllexport) void QRegularExpression_SetPatternOptions(QRegularExpression* self, PatternOptions options);
extern __declspec(dllexport) void QRegularExpression_OperatorAssign(QRegularExpression* self, QRegularExpression* re);
extern __declspec(dllexport) void QRegularExpression_Swap(QRegularExpression* self, QRegularExpression* other);
extern __declspec(dllexport) struct miqt_string QRegularExpression_Pattern(const QRegularExpression* self);
extern __declspec(dllexport) void QRegularExpression_SetPattern(QRegularExpression* self, struct miqt_string pattern);
extern __declspec(dllexport) bool QRegularExpression_IsValid(const QRegularExpression* self);
extern __declspec(dllexport) ptrdiff_t QRegularExpression_PatternErrorOffset(const QRegularExpression* self);
extern __declspec(dllexport) struct miqt_string QRegularExpression_ErrorString(const QRegularExpression* self);
extern __declspec(dllexport) int QRegularExpression_CaptureCount(const QRegularExpression* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QRegularExpression_NamedCaptureGroups(const QRegularExpression* self);
extern __declspec(dllexport) QRegularExpressionMatch* QRegularExpression_Match(const QRegularExpression* self, struct miqt_string subject);
extern __declspec(dllexport) QRegularExpressionMatchIterator* QRegularExpression_GlobalMatch(const QRegularExpression* self, struct miqt_string subject);
extern __declspec(dllexport) void QRegularExpression_Optimize(const QRegularExpression* self);
extern __declspec(dllexport) struct miqt_string QRegularExpression_Escape(struct miqt_string str);
extern __declspec(dllexport) struct miqt_string QRegularExpression_WildcardToRegularExpression(struct miqt_string str);
extern __declspec(dllexport) struct miqt_string QRegularExpression_AnchoredPattern(struct miqt_string expression);
extern __declspec(dllexport) QRegularExpressionMatch* QRegularExpression_Match2(const QRegularExpression* self, struct miqt_string subject, ptrdiff_t offset);
extern __declspec(dllexport) QRegularExpressionMatch* QRegularExpression_Match3(const QRegularExpression* self, struct miqt_string subject, ptrdiff_t offset, MatchType matchType);
extern __declspec(dllexport) QRegularExpressionMatch* QRegularExpression_Match4(const QRegularExpression* self, struct miqt_string subject, ptrdiff_t offset, MatchType matchType, MatchOptions matchOptions);
extern __declspec(dllexport) QRegularExpressionMatchIterator* QRegularExpression_GlobalMatch2(const QRegularExpression* self, struct miqt_string subject, ptrdiff_t offset);
extern __declspec(dllexport) QRegularExpressionMatchIterator* QRegularExpression_GlobalMatch3(const QRegularExpression* self, struct miqt_string subject, ptrdiff_t offset, MatchType matchType);
extern __declspec(dllexport) QRegularExpressionMatchIterator* QRegularExpression_GlobalMatch4(const QRegularExpression* self, struct miqt_string subject, ptrdiff_t offset, MatchType matchType, MatchOptions matchOptions);
extern __declspec(dllexport) struct miqt_string QRegularExpression_WildcardToRegularExpression2(struct miqt_string str, WildcardConversionOptions options);
extern __declspec(dllexport) void QRegularExpression_Delete(QRegularExpression* self, bool isSubclass);

extern __declspec(dllexport) QRegularExpressionMatch* QRegularExpressionMatch_new();
extern __declspec(dllexport) QRegularExpressionMatch* QRegularExpressionMatch_new2(QRegularExpressionMatch* match);
extern __declspec(dllexport) void QRegularExpressionMatch_OperatorAssign(QRegularExpressionMatch* self, QRegularExpressionMatch* match);
extern __declspec(dllexport) void QRegularExpressionMatch_Swap(QRegularExpressionMatch* self, QRegularExpressionMatch* other);
extern __declspec(dllexport) QRegularExpression* QRegularExpressionMatch_RegularExpression(const QRegularExpressionMatch* self);
extern __declspec(dllexport) int QRegularExpressionMatch_MatchType(const QRegularExpressionMatch* self);
extern __declspec(dllexport) int QRegularExpressionMatch_MatchOptions(const QRegularExpressionMatch* self);
extern __declspec(dllexport) bool QRegularExpressionMatch_HasMatch(const QRegularExpressionMatch* self);
extern __declspec(dllexport) bool QRegularExpressionMatch_HasPartialMatch(const QRegularExpressionMatch* self);
extern __declspec(dllexport) bool QRegularExpressionMatch_IsValid(const QRegularExpressionMatch* self);
extern __declspec(dllexport) int QRegularExpressionMatch_LastCapturedIndex(const QRegularExpressionMatch* self);
extern __declspec(dllexport) bool QRegularExpressionMatch_HasCaptured(const QRegularExpressionMatch* self, QAnyStringView* name);
extern __declspec(dllexport) bool QRegularExpressionMatch_HasCapturedWithNth(const QRegularExpressionMatch* self, int nth);
extern __declspec(dllexport) struct miqt_string QRegularExpressionMatch_Captured(const QRegularExpressionMatch* self);
extern __declspec(dllexport) struct miqt_string QRegularExpressionMatch_CapturedWithName(const QRegularExpressionMatch* self, QAnyStringView* name);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QRegularExpressionMatch_CapturedTexts(const QRegularExpressionMatch* self);
extern __declspec(dllexport) ptrdiff_t QRegularExpressionMatch_CapturedStart(const QRegularExpressionMatch* self);
extern __declspec(dllexport) ptrdiff_t QRegularExpressionMatch_CapturedLength(const QRegularExpressionMatch* self);
extern __declspec(dllexport) ptrdiff_t QRegularExpressionMatch_CapturedEnd(const QRegularExpressionMatch* self);
extern __declspec(dllexport) ptrdiff_t QRegularExpressionMatch_CapturedStartWithName(const QRegularExpressionMatch* self, QAnyStringView* name);
extern __declspec(dllexport) ptrdiff_t QRegularExpressionMatch_CapturedLengthWithName(const QRegularExpressionMatch* self, QAnyStringView* name);
extern __declspec(dllexport) ptrdiff_t QRegularExpressionMatch_CapturedEndWithName(const QRegularExpressionMatch* self, QAnyStringView* name);
extern __declspec(dllexport) struct miqt_string QRegularExpressionMatch_Captured1(const QRegularExpressionMatch* self, int nth);
extern __declspec(dllexport) ptrdiff_t QRegularExpressionMatch_CapturedStart1(const QRegularExpressionMatch* self, int nth);
extern __declspec(dllexport) ptrdiff_t QRegularExpressionMatch_CapturedLength1(const QRegularExpressionMatch* self, int nth);
extern __declspec(dllexport) ptrdiff_t QRegularExpressionMatch_CapturedEnd1(const QRegularExpressionMatch* self, int nth);
extern __declspec(dllexport) void QRegularExpressionMatch_Delete(QRegularExpressionMatch* self, bool isSubclass);

extern __declspec(dllexport) QRegularExpressionMatchIterator* QRegularExpressionMatchIterator_new();
extern __declspec(dllexport) QRegularExpressionMatchIterator* QRegularExpressionMatchIterator_new2(QRegularExpressionMatchIterator* iterator);
extern __declspec(dllexport) void QRegularExpressionMatchIterator_OperatorAssign(QRegularExpressionMatchIterator* self, QRegularExpressionMatchIterator* iterator);
extern __declspec(dllexport) void QRegularExpressionMatchIterator_Swap(QRegularExpressionMatchIterator* self, QRegularExpressionMatchIterator* other);
extern __declspec(dllexport) bool QRegularExpressionMatchIterator_IsValid(const QRegularExpressionMatchIterator* self);
extern __declspec(dllexport) bool QRegularExpressionMatchIterator_HasNext(const QRegularExpressionMatchIterator* self);
extern __declspec(dllexport) QRegularExpressionMatch* QRegularExpressionMatchIterator_Next(QRegularExpressionMatchIterator* self);
extern __declspec(dllexport) QRegularExpressionMatch* QRegularExpressionMatchIterator_PeekNext(const QRegularExpressionMatchIterator* self);
extern __declspec(dllexport) QRegularExpression* QRegularExpressionMatchIterator_RegularExpression(const QRegularExpressionMatchIterator* self);
extern __declspec(dllexport) int QRegularExpressionMatchIterator_MatchType(const QRegularExpressionMatchIterator* self);
extern __declspec(dllexport) int QRegularExpressionMatchIterator_MatchOptions(const QRegularExpressionMatchIterator* self);
extern __declspec(dllexport) void QRegularExpressionMatchIterator_Delete(QRegularExpressionMatchIterator* self, bool isSubclass);

} 
