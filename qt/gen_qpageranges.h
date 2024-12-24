#pragma once
#ifndef MIQT_QT_GEN_QPAGERANGES_H
#define MIQT_QT_GEN_QPAGERANGES_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QPageRanges__Range)
typedef QPageRanges::Range QPageRanges__Range;
typedef struct QPageRanges QPageRanges;
typedef struct QPageRanges__Range QPageRanges__Range;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QPageRanges* QPageRanges_new();
extern __declspec(dllexport) QPageRanges* QPageRanges_new2(QPageRanges* other);
extern __declspec(dllexport) void QPageRanges_OperatorAssign(QPageRanges* self, QPageRanges* other);
extern __declspec(dllexport) void QPageRanges_Swap(QPageRanges* self, QPageRanges* other);
extern __declspec(dllexport) void QPageRanges_AddPage(QPageRanges* self, int pageNumber);
extern __declspec(dllexport) void QPageRanges_AddRange(QPageRanges* self, int from, int to);
extern __declspec(dllexport) struct miqt_array /* of Range */  QPageRanges_ToRangeList(const QPageRanges* self);
extern __declspec(dllexport) void QPageRanges_Clear(QPageRanges* self);
extern __declspec(dllexport) struct miqt_string QPageRanges_ToString(const QPageRanges* self);
extern __declspec(dllexport) QPageRanges* QPageRanges_FromString(struct miqt_string ranges);
extern __declspec(dllexport) bool QPageRanges_Contains(const QPageRanges* self, int pageNumber);
extern __declspec(dllexport) bool QPageRanges_IsEmpty(const QPageRanges* self);
extern __declspec(dllexport) int QPageRanges_FirstPage(const QPageRanges* self);
extern __declspec(dllexport) int QPageRanges_LastPage(const QPageRanges* self);
extern __declspec(dllexport) void QPageRanges_Detach(QPageRanges* self);
extern __declspec(dllexport) void QPageRanges_Delete(QPageRanges* self, bool isSubclass);

extern __declspec(dllexport) QPageRanges__Range* QPageRanges__Range_new();
extern __declspec(dllexport) QPageRanges__Range* QPageRanges__Range_new2(const Range* param1);
extern __declspec(dllexport) bool QPageRanges__Range_Contains(const QPageRanges__Range* self, int pageNumber);
extern __declspec(dllexport) void QPageRanges__Range_Delete(QPageRanges__Range* self, bool isSubclass);

} 
