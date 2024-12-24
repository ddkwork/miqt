#pragma once
#ifndef MIQT_QT_GEN_QJSONARRAY_H
#define MIQT_QT_GEN_QJSONARRAY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QJsonArray__const_iterator)
typedef QJsonArray::const_iterator QJsonArray__const_iterator;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QJsonArray__iterator)
typedef QJsonArray::iterator QJsonArray__iterator;
typedef struct QJsonArray QJsonArray;
typedef struct QJsonArray__const_iterator QJsonArray__const_iterator;
typedef struct QJsonArray__iterator QJsonArray__iterator;
typedef struct QJsonValue QJsonValue;
typedef struct QJsonValueConstRef QJsonValueConstRef;
typedef struct QJsonValueRef QJsonValueRef;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QJsonArray* QJsonArray_new();
extern __declspec(dllexport) QJsonArray* QJsonArray_new2(QJsonArray* other);
extern __declspec(dllexport) void QJsonArray_OperatorAssign(QJsonArray* self, QJsonArray* other);
extern __declspec(dllexport) QJsonArray* QJsonArray_FromStringList(struct miqt_array /* of struct miqt_string */  list);
extern __declspec(dllexport) ptrdiff_t QJsonArray_Size(const QJsonArray* self);
extern __declspec(dllexport) ptrdiff_t QJsonArray_Count(const QJsonArray* self);
extern __declspec(dllexport) bool QJsonArray_IsEmpty(const QJsonArray* self);
extern __declspec(dllexport) QJsonValue* QJsonArray_At(const QJsonArray* self, ptrdiff_t i);
extern __declspec(dllexport) QJsonValue* QJsonArray_First(const QJsonArray* self);
extern __declspec(dllexport) QJsonValue* QJsonArray_Last(const QJsonArray* self);
extern __declspec(dllexport) void QJsonArray_Prepend(QJsonArray* self, QJsonValue* value);
extern __declspec(dllexport) void QJsonArray_Append(QJsonArray* self, QJsonValue* value);
extern __declspec(dllexport) void QJsonArray_RemoveAt(QJsonArray* self, ptrdiff_t i);
extern __declspec(dllexport) QJsonValue* QJsonArray_TakeAt(QJsonArray* self, ptrdiff_t i);
extern __declspec(dllexport) void QJsonArray_RemoveFirst(QJsonArray* self);
extern __declspec(dllexport) void QJsonArray_RemoveLast(QJsonArray* self);
extern __declspec(dllexport) void QJsonArray_Insert(QJsonArray* self, ptrdiff_t i, QJsonValue* value);
extern __declspec(dllexport) void QJsonArray_Replace(QJsonArray* self, ptrdiff_t i, QJsonValue* value);
extern __declspec(dllexport) bool QJsonArray_Contains(const QJsonArray* self, QJsonValue* element);
extern __declspec(dllexport) QJsonValueRef* QJsonArray_OperatorSubscript(QJsonArray* self, ptrdiff_t i);
extern __declspec(dllexport) QJsonValue* QJsonArray_OperatorSubscriptWithQsizetype(const QJsonArray* self, ptrdiff_t i);
extern __declspec(dllexport) void QJsonArray_Swap(QJsonArray* self, QJsonArray* other);
extern __declspec(dllexport) iterator QJsonArray_Begin(QJsonArray* self);
extern __declspec(dllexport) const_iterator QJsonArray_Begin2(const QJsonArray* self);
extern __declspec(dllexport) const_iterator QJsonArray_ConstBegin(const QJsonArray* self);
extern __declspec(dllexport) const_iterator QJsonArray_Cbegin(const QJsonArray* self);
extern __declspec(dllexport) iterator QJsonArray_End(QJsonArray* self);
extern __declspec(dllexport) const_iterator QJsonArray_End2(const QJsonArray* self);
extern __declspec(dllexport) const_iterator QJsonArray_ConstEnd(const QJsonArray* self);
extern __declspec(dllexport) const_iterator QJsonArray_Cend(const QJsonArray* self);
extern __declspec(dllexport) iterator QJsonArray_Insert2(QJsonArray* self, iterator before, QJsonValue* value);
extern __declspec(dllexport) iterator QJsonArray_Erase(QJsonArray* self, iterator it);
extern __declspec(dllexport) QJsonArray* QJsonArray_OperatorPlus(const QJsonArray* self, QJsonValue* v);
extern __declspec(dllexport) QJsonArray* QJsonArray_OperatorPlusAssign(QJsonArray* self, QJsonValue* v);
extern __declspec(dllexport) QJsonArray* QJsonArray_OperatorShiftLeft(QJsonArray* self, QJsonValue* v);
extern __declspec(dllexport) void QJsonArray_PushBack(QJsonArray* self, QJsonValue* t);
extern __declspec(dllexport) void QJsonArray_PushFront(QJsonArray* self, QJsonValue* t);
extern __declspec(dllexport) void QJsonArray_PopFront(QJsonArray* self);
extern __declspec(dllexport) void QJsonArray_PopBack(QJsonArray* self);
extern __declspec(dllexport) bool QJsonArray_Empty(const QJsonArray* self);
extern __declspec(dllexport) void QJsonArray_Delete(QJsonArray* self, bool isSubclass);

extern __declspec(dllexport) QJsonArray__iterator* QJsonArray__iterator_new();
extern __declspec(dllexport) QJsonArray__iterator* QJsonArray__iterator_new2(QJsonArray* array, ptrdiff_t index);
extern __declspec(dllexport) QJsonArray__iterator* QJsonArray__iterator_new3(const iterator* other);
extern __declspec(dllexport) void QJsonArray__iterator_OperatorAssign(QJsonArray__iterator* self, const iterator* other);
extern __declspec(dllexport) QJsonValueRef* QJsonArray__iterator_OperatorMultiply(const QJsonArray__iterator* self);
extern __declspec(dllexport) QJsonValueConstRef* QJsonArray__iterator_OperatorMinusGreater(const QJsonArray__iterator* self);
extern __declspec(dllexport) QJsonValueRef* QJsonArray__iterator_OperatorMinusGreater2(QJsonArray__iterator* self);
extern __declspec(dllexport) QJsonValueRef* QJsonArray__iterator_OperatorSubscript(const QJsonArray__iterator* self, ptrdiff_t j);
extern __declspec(dllexport) iterator* QJsonArray__iterator_OperatorPlusPlus(QJsonArray__iterator* self);
extern __declspec(dllexport) iterator QJsonArray__iterator_OperatorPlusPlusWithInt(QJsonArray__iterator* self, int param1);
extern __declspec(dllexport) iterator* QJsonArray__iterator_OperatorMinusMinus(QJsonArray__iterator* self);
extern __declspec(dllexport) iterator QJsonArray__iterator_OperatorMinusMinusWithInt(QJsonArray__iterator* self, int param1);
extern __declspec(dllexport) iterator* QJsonArray__iterator_OperatorPlusAssign(QJsonArray__iterator* self, ptrdiff_t j);
extern __declspec(dllexport) iterator* QJsonArray__iterator_OperatorMinusAssign(QJsonArray__iterator* self, ptrdiff_t j);
extern __declspec(dllexport) iterator QJsonArray__iterator_OperatorPlus(const QJsonArray__iterator* self, ptrdiff_t j);
extern __declspec(dllexport) iterator QJsonArray__iterator_OperatorMinus(const QJsonArray__iterator* self, ptrdiff_t j);
extern __declspec(dllexport) ptrdiff_t QJsonArray__iterator_OperatorMinusWithIterator(const QJsonArray__iterator* self, iterator j);
extern __declspec(dllexport) void QJsonArray__iterator_Delete(QJsonArray__iterator* self, bool isSubclass);

extern __declspec(dllexport) QJsonArray__const_iterator* QJsonArray__const_iterator_new();
extern __declspec(dllexport) QJsonArray__const_iterator* QJsonArray__const_iterator_new2(QJsonArray* array, ptrdiff_t index);
extern __declspec(dllexport) QJsonArray__const_iterator* QJsonArray__const_iterator_new3(const iterator* o);
extern __declspec(dllexport) QJsonArray__const_iterator* QJsonArray__const_iterator_new4(const const_iterator* other);
extern __declspec(dllexport) void QJsonArray__const_iterator_OperatorAssign(QJsonArray__const_iterator* self, const const_iterator* other);
extern __declspec(dllexport) QJsonValueConstRef* QJsonArray__const_iterator_OperatorMultiply(const QJsonArray__const_iterator* self);
extern __declspec(dllexport) QJsonValueConstRef* QJsonArray__const_iterator_OperatorMinusGreater(const QJsonArray__const_iterator* self);
extern __declspec(dllexport) QJsonValueConstRef* QJsonArray__const_iterator_OperatorSubscript(const QJsonArray__const_iterator* self, ptrdiff_t j);
extern __declspec(dllexport) const_iterator* QJsonArray__const_iterator_OperatorPlusPlus(QJsonArray__const_iterator* self);
extern __declspec(dllexport) const_iterator QJsonArray__const_iterator_OperatorPlusPlusWithInt(QJsonArray__const_iterator* self, int param1);
extern __declspec(dllexport) const_iterator* QJsonArray__const_iterator_OperatorMinusMinus(QJsonArray__const_iterator* self);
extern __declspec(dllexport) const_iterator QJsonArray__const_iterator_OperatorMinusMinusWithInt(QJsonArray__const_iterator* self, int param1);
extern __declspec(dllexport) const_iterator* QJsonArray__const_iterator_OperatorPlusAssign(QJsonArray__const_iterator* self, ptrdiff_t j);
extern __declspec(dllexport) const_iterator* QJsonArray__const_iterator_OperatorMinusAssign(QJsonArray__const_iterator* self, ptrdiff_t j);
extern __declspec(dllexport) const_iterator QJsonArray__const_iterator_OperatorPlus(const QJsonArray__const_iterator* self, ptrdiff_t j);
extern __declspec(dllexport) const_iterator QJsonArray__const_iterator_OperatorMinus(const QJsonArray__const_iterator* self, ptrdiff_t j);
extern __declspec(dllexport) ptrdiff_t QJsonArray__const_iterator_OperatorMinusWithConstIterator(const QJsonArray__const_iterator* self, const_iterator j);
extern __declspec(dllexport) void QJsonArray__const_iterator_Delete(QJsonArray__const_iterator* self, bool isSubclass);

} 
