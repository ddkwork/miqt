#pragma once
#ifndef MIQT_QT_GEN_QCOLLATOR_H
#define MIQT_QT_GEN_QCOLLATOR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChar QChar;
typedef struct QCollator QCollator;
typedef struct QCollatorSortKey QCollatorSortKey;
typedef struct QLocale QLocale;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QCollatorSortKey* QCollatorSortKey_new(QCollatorSortKey* other);
extern __declspec(dllexport) void QCollatorSortKey_OperatorAssign(QCollatorSortKey* self, QCollatorSortKey* other);
extern __declspec(dllexport) void QCollatorSortKey_Swap(QCollatorSortKey* self, QCollatorSortKey* other);
extern __declspec(dllexport) int QCollatorSortKey_Compare(const QCollatorSortKey* self, QCollatorSortKey* key);
extern __declspec(dllexport) void QCollatorSortKey_Delete(QCollatorSortKey* self, bool isSubclass);

extern __declspec(dllexport) QCollator* QCollator_new();
extern __declspec(dllexport) QCollator* QCollator_new2(QLocale* locale);
extern __declspec(dllexport) QCollator* QCollator_new3(QCollator* param1);
extern __declspec(dllexport) void QCollator_OperatorAssign(QCollator* self, QCollator* param1);
extern __declspec(dllexport) void QCollator_Swap(QCollator* self, QCollator* other);
extern __declspec(dllexport) void QCollator_SetLocale(QCollator* self, QLocale* locale);
extern __declspec(dllexport) QLocale* QCollator_Locale(const QCollator* self);
extern __declspec(dllexport) int QCollator_CaseSensitivity(const QCollator* self);
extern __declspec(dllexport) void QCollator_SetCaseSensitivity(QCollator* self, int cs);
extern __declspec(dllexport) void QCollator_SetNumericMode(QCollator* self, bool on);
extern __declspec(dllexport) bool QCollator_NumericMode(const QCollator* self);
extern __declspec(dllexport) void QCollator_SetIgnorePunctuation(QCollator* self, bool on);
extern __declspec(dllexport) bool QCollator_IgnorePunctuation(const QCollator* self);
extern __declspec(dllexport) int QCollator_Compare(const QCollator* self, struct miqt_string s1, struct miqt_string s2);
extern __declspec(dllexport) int QCollator_Compare2(const QCollator* self, QChar* s1, ptrdiff_t len1, QChar* s2, ptrdiff_t len2);
extern __declspec(dllexport) bool QCollator_OperatorCall(const QCollator* self, struct miqt_string s1, struct miqt_string s2);
extern __declspec(dllexport) QCollatorSortKey* QCollator_SortKey(const QCollator* self, struct miqt_string stringVal);
extern __declspec(dllexport) void QCollator_Delete(QCollator* self, bool isSubclass);

} 
