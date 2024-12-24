#pragma once
#ifndef MIQT_QT_GEN_QANYSTRINGVIEW_H
#define MIQT_QT_GEN_QANYSTRINGVIEW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAnyStringView QAnyStringView;
typedef struct QChar QChar;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QAnyStringView* QAnyStringView_new();
extern __declspec(dllexport) QAnyStringView* QAnyStringView_new2(struct miqt_string str);
extern __declspec(dllexport) QAnyStringView* QAnyStringView_new3(struct miqt_string str);
extern __declspec(dllexport) QAnyStringView* QAnyStringView_new4(QAnyStringView* param1);
extern __declspec(dllexport) QAnyStringView* QAnyStringView_Mid(const QAnyStringView* self, ptrdiff_t pos);
extern __declspec(dllexport) QAnyStringView* QAnyStringView_Left(const QAnyStringView* self, ptrdiff_t n);
extern __declspec(dllexport) QAnyStringView* QAnyStringView_Right(const QAnyStringView* self, ptrdiff_t n);
extern __declspec(dllexport) QAnyStringView* QAnyStringView_Sliced(const QAnyStringView* self, ptrdiff_t pos);
extern __declspec(dllexport) QAnyStringView* QAnyStringView_Sliced2(const QAnyStringView* self, ptrdiff_t pos, ptrdiff_t n);
extern __declspec(dllexport) QAnyStringView* QAnyStringView_First(const QAnyStringView* self, ptrdiff_t n);
extern __declspec(dllexport) QAnyStringView* QAnyStringView_Last(const QAnyStringView* self, ptrdiff_t n);
extern __declspec(dllexport) QAnyStringView* QAnyStringView_Chopped(const QAnyStringView* self, ptrdiff_t n);
extern __declspec(dllexport) QAnyStringView* QAnyStringView_Slice(QAnyStringView* self, ptrdiff_t pos);
extern __declspec(dllexport) QAnyStringView* QAnyStringView_Slice2(QAnyStringView* self, ptrdiff_t pos, ptrdiff_t n);
extern __declspec(dllexport) void QAnyStringView_Truncate(QAnyStringView* self, ptrdiff_t n);
extern __declspec(dllexport) void QAnyStringView_Chop(QAnyStringView* self, ptrdiff_t n);
extern __declspec(dllexport) struct miqt_string QAnyStringView_ToString(const QAnyStringView* self);
extern __declspec(dllexport) ptrdiff_t QAnyStringView_Size(const QAnyStringView* self);
extern __declspec(dllexport) const void* QAnyStringView_Data(const QAnyStringView* self);
extern __declspec(dllexport) int QAnyStringView_Compare(QAnyStringView* lhs, QAnyStringView* rhs);
extern __declspec(dllexport) bool QAnyStringView_Equal(QAnyStringView* lhs, QAnyStringView* rhs);
extern __declspec(dllexport) QChar* QAnyStringView_Front(const QAnyStringView* self);
extern __declspec(dllexport) QChar* QAnyStringView_Back(const QAnyStringView* self);
extern __declspec(dllexport) bool QAnyStringView_Empty(const QAnyStringView* self);
extern __declspec(dllexport) ptrdiff_t QAnyStringView_SizeBytes(const QAnyStringView* self);
extern __declspec(dllexport) ptrdiff_t QAnyStringView_MaxSize(const QAnyStringView* self);
extern __declspec(dllexport) bool QAnyStringView_IsNull(const QAnyStringView* self);
extern __declspec(dllexport) bool QAnyStringView_IsEmpty(const QAnyStringView* self);
extern __declspec(dllexport) ptrdiff_t QAnyStringView_Length(const QAnyStringView* self);
extern __declspec(dllexport) QAnyStringView* QAnyStringView_Mid2(const QAnyStringView* self, ptrdiff_t pos, ptrdiff_t n);
extern __declspec(dllexport) int QAnyStringView_Compare3(QAnyStringView* lhs, QAnyStringView* rhs, int cs);
extern __declspec(dllexport) void QAnyStringView_Delete(QAnyStringView* self, bool isSubclass);

} 
