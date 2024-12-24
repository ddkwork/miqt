#pragma once
#ifndef MIQT_QT_GEN_QSTRINGVIEW_H
#define MIQT_QT_GEN_QSTRINGVIEW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChar QChar;
typedef struct QRegularExpression QRegularExpression;
typedef struct QRegularExpressionMatch QRegularExpressionMatch;
typedef struct QStringView QStringView;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QStringView* QStringView_new();
extern __declspec(dllexport) struct miqt_string QStringView_ToString(const QStringView* self);
extern __declspec(dllexport) ptrdiff_t QStringView_Size(const QStringView* self);
extern __declspec(dllexport) const_pointer QStringView_Data(const QStringView* self);
extern __declspec(dllexport) const_pointer QStringView_ConstData(const QStringView* self);
extern __declspec(dllexport) const storage_type* QStringView_Utf16(const QStringView* self);
extern __declspec(dllexport) QChar* QStringView_OperatorSubscript(const QStringView* self, ptrdiff_t n);
extern __declspec(dllexport) struct miqt_string QStringView_ToLatin1(const QStringView* self);
extern __declspec(dllexport) struct miqt_string QStringView_ToUtf8(const QStringView* self);
extern __declspec(dllexport) struct miqt_string QStringView_ToLocal8Bit(const QStringView* self);
extern __declspec(dllexport) struct miqt_array /* of unsigned int */  QStringView_ToUcs4(const QStringView* self);
extern __declspec(dllexport) QChar* QStringView_At(const QStringView* self, ptrdiff_t n);
extern __declspec(dllexport) void QStringView_Truncate(QStringView* self, ptrdiff_t n);
extern __declspec(dllexport) void QStringView_Chop(QStringView* self, ptrdiff_t n);
extern __declspec(dllexport) int QStringView_CompareWithQChar(const QStringView* self, QChar* c);
extern __declspec(dllexport) int QStringView_Compare3(const QStringView* self, QChar* c, int cs);
extern __declspec(dllexport) bool QStringView_StartsWithWithQChar(const QStringView* self, QChar* c);
extern __declspec(dllexport) bool QStringView_StartsWith2(const QStringView* self, QChar* c, int cs);
extern __declspec(dllexport) bool QStringView_EndsWithWithQChar(const QStringView* self, QChar* c);
extern __declspec(dllexport) bool QStringView_EndsWith2(const QStringView* self, QChar* c, int cs);
extern __declspec(dllexport) ptrdiff_t QStringView_IndexOf(const QStringView* self, QChar* c);
extern __declspec(dllexport) bool QStringView_Contains(const QStringView* self, QChar* c);
extern __declspec(dllexport) ptrdiff_t QStringView_Count(const QStringView* self, QChar* c);
extern __declspec(dllexport) ptrdiff_t QStringView_LastIndexOf(const QStringView* self, QChar* c);
extern __declspec(dllexport) ptrdiff_t QStringView_LastIndexOf2(const QStringView* self, QChar* c, ptrdiff_t from);
extern __declspec(dllexport) ptrdiff_t QStringView_IndexOfWithRe(const QStringView* self, QRegularExpression* re);
extern __declspec(dllexport) ptrdiff_t QStringView_LastIndexOf5(const QStringView* self, QRegularExpression* re, ptrdiff_t from);
extern __declspec(dllexport) bool QStringView_ContainsWithRe(const QStringView* self, QRegularExpression* re);
extern __declspec(dllexport) ptrdiff_t QStringView_CountWithRe(const QStringView* self, QRegularExpression* re);
extern __declspec(dllexport) bool QStringView_IsRightToLeft(const QStringView* self);
extern __declspec(dllexport) bool QStringView_IsValidUtf16(const QStringView* self);
extern __declspec(dllexport) bool QStringView_IsUpper(const QStringView* self);
extern __declspec(dllexport) bool QStringView_IsLower(const QStringView* self);
extern __declspec(dllexport) int16_t QStringView_ToShort(const QStringView* self);
extern __declspec(dllexport) uint16_t QStringView_ToUShort(const QStringView* self);
extern __declspec(dllexport) int QStringView_ToInt(const QStringView* self);
extern __declspec(dllexport) unsigned int QStringView_ToUInt(const QStringView* self);
extern __declspec(dllexport) long QStringView_ToLong(const QStringView* self);
extern __declspec(dllexport) unsigned long QStringView_ToULong(const QStringView* self);
extern __declspec(dllexport) long long QStringView_ToLongLong(const QStringView* self);
extern __declspec(dllexport) unsigned long long QStringView_ToULongLong(const QStringView* self);
extern __declspec(dllexport) float QStringView_ToFloat(const QStringView* self);
extern __declspec(dllexport) double QStringView_ToDouble(const QStringView* self);
extern __declspec(dllexport) const_iterator QStringView_Begin(const QStringView* self);
extern __declspec(dllexport) const_iterator QStringView_End(const QStringView* self);
extern __declspec(dllexport) const_iterator QStringView_Cbegin(const QStringView* self);
extern __declspec(dllexport) const_iterator QStringView_Cend(const QStringView* self);
extern __declspec(dllexport) const_reverse_iterator QStringView_Rbegin(const QStringView* self);
extern __declspec(dllexport) const_reverse_iterator QStringView_Rend(const QStringView* self);
extern __declspec(dllexport) const_reverse_iterator QStringView_Crbegin(const QStringView* self);
extern __declspec(dllexport) const_reverse_iterator QStringView_Crend(const QStringView* self);
extern __declspec(dllexport) bool QStringView_Empty(const QStringView* self);
extern __declspec(dllexport) QChar* QStringView_Front(const QStringView* self);
extern __declspec(dllexport) QChar* QStringView_Back(const QStringView* self);
extern __declspec(dllexport) ptrdiff_t QStringView_MaxSize(const QStringView* self);
extern __declspec(dllexport) const_iterator QStringView_ConstBegin(const QStringView* self);
extern __declspec(dllexport) const_iterator QStringView_ConstEnd(const QStringView* self);
extern __declspec(dllexport) bool QStringView_IsNull(const QStringView* self);
extern __declspec(dllexport) bool QStringView_IsEmpty(const QStringView* self);
extern __declspec(dllexport) ptrdiff_t QStringView_Length(const QStringView* self);
extern __declspec(dllexport) QChar* QStringView_First2(const QStringView* self);
extern __declspec(dllexport) QChar* QStringView_Last2(const QStringView* self);
extern __declspec(dllexport) ptrdiff_t QStringView_MaxSize2();
extern __declspec(dllexport) ptrdiff_t QStringView_IndexOf2(const QStringView* self, QChar* c, ptrdiff_t from);
extern __declspec(dllexport) ptrdiff_t QStringView_IndexOf3(const QStringView* self, QChar* c, ptrdiff_t from, int cs);
extern __declspec(dllexport) bool QStringView_Contains2(const QStringView* self, QChar* c, int cs);
extern __declspec(dllexport) ptrdiff_t QStringView_Count2(const QStringView* self, QChar* c, int cs);
extern __declspec(dllexport) ptrdiff_t QStringView_LastIndexOf22(const QStringView* self, QChar* c, int cs);
extern __declspec(dllexport) ptrdiff_t QStringView_LastIndexOf32(const QStringView* self, QChar* c, ptrdiff_t from, int cs);
extern __declspec(dllexport) ptrdiff_t QStringView_IndexOf24(const QStringView* self, QRegularExpression* re, ptrdiff_t from);
extern __declspec(dllexport) ptrdiff_t QStringView_IndexOf34(const QStringView* self, QRegularExpression* re, ptrdiff_t from, QRegularExpressionMatch* rmatch);
extern __declspec(dllexport) ptrdiff_t QStringView_LastIndexOf35(const QStringView* self, QRegularExpression* re, ptrdiff_t from, QRegularExpressionMatch* rmatch);
extern __declspec(dllexport) bool QStringView_Contains24(const QStringView* self, QRegularExpression* re, QRegularExpressionMatch* rmatch);
extern __declspec(dllexport) int16_t QStringView_ToShort1(const QStringView* self, bool* ok);
extern __declspec(dllexport) int16_t QStringView_ToShort2(const QStringView* self, bool* ok, int base);
extern __declspec(dllexport) uint16_t QStringView_ToUShort1(const QStringView* self, bool* ok);
extern __declspec(dllexport) uint16_t QStringView_ToUShort2(const QStringView* self, bool* ok, int base);
extern __declspec(dllexport) int QStringView_ToInt1(const QStringView* self, bool* ok);
extern __declspec(dllexport) int QStringView_ToInt2(const QStringView* self, bool* ok, int base);
extern __declspec(dllexport) unsigned int QStringView_ToUInt1(const QStringView* self, bool* ok);
extern __declspec(dllexport) unsigned int QStringView_ToUInt2(const QStringView* self, bool* ok, int base);
extern __declspec(dllexport) long QStringView_ToLong1(const QStringView* self, bool* ok);
extern __declspec(dllexport) long QStringView_ToLong2(const QStringView* self, bool* ok, int base);
extern __declspec(dllexport) unsigned long QStringView_ToULong1(const QStringView* self, bool* ok);
extern __declspec(dllexport) unsigned long QStringView_ToULong2(const QStringView* self, bool* ok, int base);
extern __declspec(dllexport) long long QStringView_ToLongLong1(const QStringView* self, bool* ok);
extern __declspec(dllexport) long long QStringView_ToLongLong2(const QStringView* self, bool* ok, int base);
extern __declspec(dllexport) unsigned long long QStringView_ToULongLong1(const QStringView* self, bool* ok);
extern __declspec(dllexport) unsigned long long QStringView_ToULongLong2(const QStringView* self, bool* ok, int base);
extern __declspec(dllexport) float QStringView_ToFloat1(const QStringView* self, bool* ok);
extern __declspec(dllexport) double QStringView_ToDouble1(const QStringView* self, bool* ok);
extern __declspec(dllexport) void QStringView_Delete(QStringView* self, bool isSubclass);

} 
