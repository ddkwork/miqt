#pragma once
#ifndef MIQT_QT_GEN_QVERSIONNUMBER_H
#define MIQT_QT_GEN_QVERSIONNUMBER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAnyStringView QAnyStringView;
typedef struct QVersionNumber QVersionNumber;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QVersionNumber* QVersionNumber_new();
extern __declspec(dllexport) QVersionNumber* QVersionNumber_new2(QSpan<const int> args);
extern __declspec(dllexport) QVersionNumber* QVersionNumber_new3(int maj);
extern __declspec(dllexport) QVersionNumber* QVersionNumber_new4(int maj, int min);
extern __declspec(dllexport) QVersionNumber* QVersionNumber_new5(int maj, int min, int mic);
extern __declspec(dllexport) bool QVersionNumber_IsNull(const QVersionNumber* self);
extern __declspec(dllexport) bool QVersionNumber_IsNormalized(const QVersionNumber* self);
extern __declspec(dllexport) int QVersionNumber_MajorVersion(const QVersionNumber* self);
extern __declspec(dllexport) int QVersionNumber_MinorVersion(const QVersionNumber* self);
extern __declspec(dllexport) int QVersionNumber_MicroVersion(const QVersionNumber* self);
extern __declspec(dllexport) QVersionNumber* QVersionNumber_Normalized(const QVersionNumber* self);
extern __declspec(dllexport) struct miqt_array /* of int */  QVersionNumber_Segments(const QVersionNumber* self);
extern __declspec(dllexport) int QVersionNumber_SegmentAt(const QVersionNumber* self, ptrdiff_t index);
extern __declspec(dllexport) ptrdiff_t QVersionNumber_SegmentCount(const QVersionNumber* self);
extern __declspec(dllexport) const_iterator QVersionNumber_Begin(const QVersionNumber* self);
extern __declspec(dllexport) const_iterator QVersionNumber_End(const QVersionNumber* self);
extern __declspec(dllexport) const_iterator QVersionNumber_Cbegin(const QVersionNumber* self);
extern __declspec(dllexport) const_iterator QVersionNumber_Cend(const QVersionNumber* self);
extern __declspec(dllexport) const_reverse_iterator QVersionNumber_Rbegin(const QVersionNumber* self);
extern __declspec(dllexport) const_reverse_iterator QVersionNumber_Rend(const QVersionNumber* self);
extern __declspec(dllexport) const_reverse_iterator QVersionNumber_Crbegin(const QVersionNumber* self);
extern __declspec(dllexport) const_reverse_iterator QVersionNumber_Crend(const QVersionNumber* self);
extern __declspec(dllexport) const_iterator QVersionNumber_ConstBegin(const QVersionNumber* self);
extern __declspec(dllexport) const_iterator QVersionNumber_ConstEnd(const QVersionNumber* self);
extern __declspec(dllexport) bool QVersionNumber_IsPrefixOf(const QVersionNumber* self, QVersionNumber* other);
extern __declspec(dllexport) int QVersionNumber_Compare(QVersionNumber* v1, QVersionNumber* v2);
extern __declspec(dllexport) QVersionNumber* QVersionNumber_CommonPrefix(QVersionNumber* v1, QVersionNumber* v2);
extern __declspec(dllexport) struct miqt_string QVersionNumber_ToString(const QVersionNumber* self);
extern __declspec(dllexport) QVersionNumber* QVersionNumber_FromString(QAnyStringView* stringVal);
extern __declspec(dllexport) QVersionNumber* QVersionNumber_FromString2(QAnyStringView* stringVal, ptrdiff_t* suffixIndex);
extern __declspec(dllexport) void QVersionNumber_Delete(QVersionNumber* self, bool isSubclass);

} 
