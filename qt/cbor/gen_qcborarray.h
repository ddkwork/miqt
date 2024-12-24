#pragma once
#ifndef MIQT_QT_CBOR_GEN_QCBORARRAY_H
#define MIQT_QT_CBOR_GEN_QCBORARRAY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QCborArray__ConstIterator)
typedef QCborArray::ConstIterator QCborArray__ConstIterator;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QCborArray__Iterator)
typedef QCborArray::Iterator QCborArray__Iterator;
typedef struct QCborArray QCborArray;
typedef struct QCborArray__ConstIterator QCborArray__ConstIterator;
typedef struct QCborArray__Iterator QCborArray__Iterator;
typedef struct QCborValue QCborValue;
typedef struct QCborValueConstRef QCborValueConstRef;
typedef struct QCborValueRef QCborValueRef;
typedef struct QJsonArray QJsonArray;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QCborArray* QCborArray_new();
extern __declspec(dllexport) QCborArray* QCborArray_new2(QCborArray* other);
extern __declspec(dllexport) void QCborArray_OperatorAssign(QCborArray* self, QCborArray* other);
extern __declspec(dllexport) void QCborArray_Swap(QCborArray* self, QCborArray* other);
extern __declspec(dllexport) QCborValue* QCborArray_ToCborValue(const QCborArray* self);
extern __declspec(dllexport) ptrdiff_t QCborArray_Size(const QCborArray* self);
extern __declspec(dllexport) bool QCborArray_IsEmpty(const QCborArray* self);
extern __declspec(dllexport) void QCborArray_Clear(QCborArray* self);
extern __declspec(dllexport) QCborValue* QCborArray_At(const QCborArray* self, ptrdiff_t i);
extern __declspec(dllexport) QCborValue* QCborArray_First(const QCborArray* self);
extern __declspec(dllexport) QCborValue* QCborArray_Last(const QCborArray* self);
extern __declspec(dllexport) QCborValue* QCborArray_OperatorSubscript(const QCborArray* self, ptrdiff_t i);
extern __declspec(dllexport) QCborValueRef* QCborArray_First2(QCborArray* self);
extern __declspec(dllexport) QCborValueRef* QCborArray_Last2(QCborArray* self);
extern __declspec(dllexport) QCborValueRef* QCborArray_OperatorSubscriptWithQsizetype(QCborArray* self, ptrdiff_t i);
extern __declspec(dllexport) void QCborArray_Insert(QCborArray* self, ptrdiff_t i, QCborValue* value);
extern __declspec(dllexport) void QCborArray_Prepend(QCborArray* self, QCborValue* value);
extern __declspec(dllexport) void QCborArray_Append(QCborArray* self, QCborValue* value);
extern __declspec(dllexport) QCborValue* QCborArray_Extract(QCborArray* self, ConstIterator it);
extern __declspec(dllexport) QCborValue* QCborArray_ExtractWithIt(QCborArray* self, Iterator it);
extern __declspec(dllexport) void QCborArray_RemoveAt(QCborArray* self, ptrdiff_t i);
extern __declspec(dllexport) QCborValue* QCborArray_TakeAt(QCborArray* self, ptrdiff_t i);
extern __declspec(dllexport) void QCborArray_RemoveFirst(QCborArray* self);
extern __declspec(dllexport) void QCborArray_RemoveLast(QCborArray* self);
extern __declspec(dllexport) QCborValue* QCborArray_TakeFirst(QCborArray* self);
extern __declspec(dllexport) QCborValue* QCborArray_TakeLast(QCborArray* self);
extern __declspec(dllexport) bool QCborArray_Contains(const QCborArray* self, QCborValue* value);
extern __declspec(dllexport) int QCborArray_Compare(const QCborArray* self, QCborArray* other);
extern __declspec(dllexport) iterator QCborArray_Begin(QCborArray* self);
extern __declspec(dllexport) const_iterator QCborArray_ConstBegin(const QCborArray* self);
extern __declspec(dllexport) const_iterator QCborArray_Begin2(const QCborArray* self);
extern __declspec(dllexport) const_iterator QCborArray_Cbegin(const QCborArray* self);
extern __declspec(dllexport) iterator QCborArray_End(QCborArray* self);
extern __declspec(dllexport) const_iterator QCborArray_ConstEnd(const QCborArray* self);
extern __declspec(dllexport) const_iterator QCborArray_End2(const QCborArray* self);
extern __declspec(dllexport) const_iterator QCborArray_Cend(const QCborArray* self);
extern __declspec(dllexport) iterator QCborArray_Insert2(QCborArray* self, iterator before, QCborValue* value);
extern __declspec(dllexport) iterator QCborArray_Insert3(QCborArray* self, const_iterator before, QCborValue* value);
extern __declspec(dllexport) iterator QCborArray_Erase(QCborArray* self, iterator it);
extern __declspec(dllexport) iterator QCborArray_EraseWithIt(QCborArray* self, const_iterator it);
extern __declspec(dllexport) void QCborArray_PushBack(QCborArray* self, QCborValue* t);
extern __declspec(dllexport) void QCborArray_PushFront(QCborArray* self, QCborValue* t);
extern __declspec(dllexport) void QCborArray_PopFront(QCborArray* self);
extern __declspec(dllexport) void QCborArray_PopBack(QCborArray* self);
extern __declspec(dllexport) bool QCborArray_Empty(const QCborArray* self);
extern __declspec(dllexport) QCborArray* QCborArray_OperatorPlus(const QCborArray* self, QCborValue* v);
extern __declspec(dllexport) QCborArray* QCborArray_OperatorPlusAssign(QCborArray* self, QCborValue* v);
extern __declspec(dllexport) QCborArray* QCborArray_OperatorShiftLeft(QCborArray* self, QCborValue* v);
extern __declspec(dllexport) QCborArray* QCborArray_FromStringList(struct miqt_array /* of struct miqt_string */  list);
extern __declspec(dllexport) QCborArray* QCborArray_FromJsonArray(QJsonArray* array);
extern __declspec(dllexport) QJsonArray* QCborArray_ToJsonArray(const QCborArray* self);
extern __declspec(dllexport) void QCborArray_Delete(QCborArray* self, bool isSubclass);

extern __declspec(dllexport) QCborArray__Iterator* QCborArray__Iterator_new();
extern __declspec(dllexport) QCborArray__Iterator* QCborArray__Iterator_new2(const Iterator* param1);
extern __declspec(dllexport) void QCborArray__Iterator_OperatorAssign(QCborArray__Iterator* self, const Iterator* other);
extern __declspec(dllexport) QCborValueRef* QCborArray__Iterator_OperatorMultiply(const QCborArray__Iterator* self);
extern __declspec(dllexport) QCborValueRef* QCborArray__Iterator_OperatorMinusGreater(QCborArray__Iterator* self);
extern __declspec(dllexport) QCborValueConstRef* QCborArray__Iterator_OperatorMinusGreater2(const QCborArray__Iterator* self);
extern __declspec(dllexport) QCborValueRef* QCborArray__Iterator_OperatorSubscript(const QCborArray__Iterator* self, ptrdiff_t j);
extern __declspec(dllexport) Iterator* QCborArray__Iterator_OperatorPlusPlus(QCborArray__Iterator* self);
extern __declspec(dllexport) Iterator QCborArray__Iterator_OperatorPlusPlusWithInt(QCborArray__Iterator* self, int param1);
extern __declspec(dllexport) Iterator* QCborArray__Iterator_OperatorMinusMinus(QCborArray__Iterator* self);
extern __declspec(dllexport) Iterator QCborArray__Iterator_OperatorMinusMinusWithInt(QCborArray__Iterator* self, int param1);
extern __declspec(dllexport) Iterator* QCborArray__Iterator_OperatorPlusAssign(QCborArray__Iterator* self, ptrdiff_t j);
extern __declspec(dllexport) Iterator* QCborArray__Iterator_OperatorMinusAssign(QCborArray__Iterator* self, ptrdiff_t j);
extern __declspec(dllexport) Iterator QCborArray__Iterator_OperatorPlus(const QCborArray__Iterator* self, ptrdiff_t j);
extern __declspec(dllexport) Iterator QCborArray__Iterator_OperatorMinus(const QCborArray__Iterator* self, ptrdiff_t j);
extern __declspec(dllexport) ptrdiff_t QCborArray__Iterator_OperatorMinusWithIterator(const QCborArray__Iterator* self, Iterator j);
extern __declspec(dllexport) void QCborArray__Iterator_Delete(QCborArray__Iterator* self, bool isSubclass);

extern __declspec(dllexport) QCborArray__ConstIterator* QCborArray__ConstIterator_new();
extern __declspec(dllexport) QCborArray__ConstIterator* QCborArray__ConstIterator_new2(const ConstIterator* param1);
extern __declspec(dllexport) void QCborArray__ConstIterator_OperatorAssign(QCborArray__ConstIterator* self, const ConstIterator* other);
extern __declspec(dllexport) QCborValueConstRef* QCborArray__ConstIterator_OperatorMultiply(const QCborArray__ConstIterator* self);
extern __declspec(dllexport) QCborValueConstRef* QCborArray__ConstIterator_OperatorMinusGreater(const QCborArray__ConstIterator* self);
extern __declspec(dllexport) QCborValueConstRef* QCborArray__ConstIterator_OperatorSubscript(const QCborArray__ConstIterator* self, ptrdiff_t j);
extern __declspec(dllexport) ConstIterator* QCborArray__ConstIterator_OperatorPlusPlus(QCborArray__ConstIterator* self);
extern __declspec(dllexport) ConstIterator QCborArray__ConstIterator_OperatorPlusPlusWithInt(QCborArray__ConstIterator* self, int param1);
extern __declspec(dllexport) ConstIterator* QCborArray__ConstIterator_OperatorMinusMinus(QCborArray__ConstIterator* self);
extern __declspec(dllexport) ConstIterator QCborArray__ConstIterator_OperatorMinusMinusWithInt(QCborArray__ConstIterator* self, int param1);
extern __declspec(dllexport) ConstIterator* QCborArray__ConstIterator_OperatorPlusAssign(QCborArray__ConstIterator* self, ptrdiff_t j);
extern __declspec(dllexport) ConstIterator* QCborArray__ConstIterator_OperatorMinusAssign(QCborArray__ConstIterator* self, ptrdiff_t j);
extern __declspec(dllexport) ConstIterator QCborArray__ConstIterator_OperatorPlus(const QCborArray__ConstIterator* self, ptrdiff_t j);
extern __declspec(dllexport) ConstIterator QCborArray__ConstIterator_OperatorMinus(const QCborArray__ConstIterator* self, ptrdiff_t j);
extern __declspec(dllexport) ptrdiff_t QCborArray__ConstIterator_OperatorMinusWithConstIterator(const QCborArray__ConstIterator* self, ConstIterator j);
extern __declspec(dllexport) void QCborArray__ConstIterator_Delete(QCborArray__ConstIterator* self, bool isSubclass);

} 
