#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QHTTPHEADERS_H
#define MIQT_QT_NETWORK_GEN_QHTTPHEADERS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAnyStringView QAnyStringView;
typedef struct QByteArrayView QByteArrayView;
typedef struct QHttpHeaders QHttpHeaders;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QHttpHeaders* QHttpHeaders_new();
extern __declspec(dllexport) 
QHttpHeaders* QHttpHeaders_new2(QHttpHeaders* other);
extern __declspec(dllexport) 
void QHttpHeaders_OperatorAssign(QHttpHeaders* self, QHttpHeaders* other);
extern __declspec(dllexport) 
void QHttpHeaders_Swap(QHttpHeaders* self, QHttpHeaders* other);
extern __declspec(dllexport) 
bool QHttpHeaders_Append(QHttpHeaders* self, QAnyStringView* name, QAnyStringView* value);
extern __declspec(dllexport) 
bool QHttpHeaders_Append2(QHttpHeaders* self, WellKnownHeader name, QAnyStringView* value);
extern __declspec(dllexport) 
bool QHttpHeaders_Insert(QHttpHeaders* self, ptrdiff_t i, QAnyStringView* name, QAnyStringView* value);
extern __declspec(dllexport) 
bool QHttpHeaders_Insert2(QHttpHeaders* self, ptrdiff_t i, WellKnownHeader name, QAnyStringView* value);
extern __declspec(dllexport) 
bool QHttpHeaders_Replace(QHttpHeaders* self, ptrdiff_t i, QAnyStringView* name, QAnyStringView* newValue);
extern __declspec(dllexport) 
bool QHttpHeaders_Replace2(QHttpHeaders* self, ptrdiff_t i, WellKnownHeader name, QAnyStringView* newValue);
extern __declspec(dllexport) 
bool QHttpHeaders_ReplaceOrAppend(QHttpHeaders* self, QAnyStringView* name, QAnyStringView* newValue);
extern __declspec(dllexport) 
bool QHttpHeaders_ReplaceOrAppend2(QHttpHeaders* self, WellKnownHeader name, QAnyStringView* newValue);
extern __declspec(dllexport) 
bool QHttpHeaders_Contains(const QHttpHeaders* self, QAnyStringView* name);
extern __declspec(dllexport) 
bool QHttpHeaders_ContainsWithName(const QHttpHeaders* self, WellKnownHeader name);
extern __declspec(dllexport) 
void QHttpHeaders_Clear(QHttpHeaders* self);
extern __declspec(dllexport) 
void QHttpHeaders_RemoveAll(QHttpHeaders* self, QAnyStringView* name);
extern __declspec(dllexport) 
void QHttpHeaders_RemoveAllWithName(QHttpHeaders* self, WellKnownHeader name);
extern __declspec(dllexport) 
void QHttpHeaders_RemoveAt(QHttpHeaders* self, ptrdiff_t i);
extern __declspec(dllexport) 
QByteArrayView* QHttpHeaders_Value(const QHttpHeaders* self, QAnyStringView* name);
extern __declspec(dllexport) 
QByteArrayView* QHttpHeaders_ValueWithName(const QHttpHeaders* self, WellKnownHeader name);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QHttpHeaders_Values(const QHttpHeaders* self, QAnyStringView* name);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QHttpHeaders_ValuesWithName(const QHttpHeaders* self, WellKnownHeader name);
extern __declspec(dllexport) 
QByteArrayView* QHttpHeaders_ValueAt(const QHttpHeaders* self, ptrdiff_t i);
extern __declspec(dllexport) 
struct miqt_string QHttpHeaders_CombinedValue(const QHttpHeaders* self, QAnyStringView* name);
extern __declspec(dllexport) 
struct miqt_string QHttpHeaders_CombinedValueWithName(const QHttpHeaders* self, WellKnownHeader name);
extern __declspec(dllexport) 
ptrdiff_t QHttpHeaders_Size(const QHttpHeaders* self);
extern __declspec(dllexport) 
void QHttpHeaders_Reserve(QHttpHeaders* self, ptrdiff_t size);
extern __declspec(dllexport) 
bool QHttpHeaders_IsEmpty(const QHttpHeaders* self);
extern __declspec(dllexport) 
QByteArrayView* QHttpHeaders_WellKnownHeaderName(WellKnownHeader name);
extern __declspec(dllexport) 
QByteArrayView* QHttpHeaders_Value2(const QHttpHeaders* self, QAnyStringView* name, QByteArrayView* defaultValue);
extern __declspec(dllexport) 
QByteArrayView* QHttpHeaders_Value22(const QHttpHeaders* self, WellKnownHeader name, QByteArrayView* defaultValue);
extern __declspec(dllexport) 
void QHttpHeaders_Delete(QHttpHeaders* self, bool isSubclass);

}
