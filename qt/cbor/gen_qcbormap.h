#pragma once
#ifndef MIQT_QT_CBOR_GEN_QCBORMAP_H
#define MIQT_QT_CBOR_GEN_QCBORMAP_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QCborMap__ConstIterator)
typedef QCborMap::ConstIterator QCborMap__ConstIterator;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QCborMap__Iterator)
typedef QCborMap::Iterator QCborMap__Iterator;
typedef struct QCborMap QCborMap;
typedef struct QCborMap__ConstIterator QCborMap__ConstIterator;
typedef struct QCborMap__Iterator QCborMap__Iterator;
typedef struct QCborValue QCborValue;
typedef struct QCborValueConstRef QCborValueConstRef;
typedef struct QCborValueRef QCborValueRef;
typedef struct QJsonObject QJsonObject;
typedef struct QVariant QVariant;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QCborMap* QCborMap_new();
extern __declspec(dllexport) QCborMap* QCborMap_new2(QCborMap* other);
extern __declspec(dllexport) void QCborMap_OperatorAssign(QCborMap* self, QCborMap* other);
extern __declspec(dllexport) void QCborMap_Swap(QCborMap* self, QCborMap* other);
extern __declspec(dllexport) QCborValue* QCborMap_ToCborValue(const QCborMap* self);
extern __declspec(dllexport) ptrdiff_t QCborMap_Size(const QCborMap* self);
extern __declspec(dllexport) bool QCborMap_IsEmpty(const QCborMap* self);
extern __declspec(dllexport) void QCborMap_Clear(QCborMap* self);
extern __declspec(dllexport) struct miqt_array /* of QCborValue* */  QCborMap_Keys(const QCborMap* self);
extern __declspec(dllexport) QCborValue* QCborMap_Value(const QCborMap* self, long long key);
extern __declspec(dllexport) QCborValue* QCborMap_Value2(const QCborMap* self, struct miqt_string key);
extern __declspec(dllexport) QCborValue* QCborMap_Value3(const QCborMap* self, QCborValue* key);
extern __declspec(dllexport) QCborValue* QCborMap_OperatorSubscript(const QCborMap* self, long long key);
extern __declspec(dllexport) QCborValue* QCborMap_OperatorSubscript2(const QCborMap* self, struct miqt_string key);
extern __declspec(dllexport) QCborValue* QCborMap_OperatorSubscript3(const QCborMap* self, QCborValue* key);
extern __declspec(dllexport) QCborValueRef* QCborMap_OperatorSubscript4(QCborMap* self, long long key);
extern __declspec(dllexport) QCborValueRef* QCborMap_OperatorSubscript6(QCborMap* self, struct miqt_string key);
extern __declspec(dllexport) QCborValueRef* QCborMap_OperatorSubscript7(QCborMap* self, QCborValue* key);
extern __declspec(dllexport) QCborValue* QCborMap_Take(QCborMap* self, long long key);
extern __declspec(dllexport) QCborValue* QCborMap_Take2(QCborMap* self, struct miqt_string key);
extern __declspec(dllexport) QCborValue* QCborMap_Take3(QCborMap* self, QCborValue* key);
extern __declspec(dllexport) void QCborMap_Remove(QCborMap* self, long long key);
extern __declspec(dllexport) void QCborMap_Remove2(QCborMap* self, struct miqt_string key);
extern __declspec(dllexport) void QCborMap_Remove3(QCborMap* self, QCborValue* key);
extern __declspec(dllexport) bool QCborMap_Contains(const QCborMap* self, long long key);
extern __declspec(dllexport) bool QCborMap_Contains2(const QCborMap* self, struct miqt_string key);
extern __declspec(dllexport) bool QCborMap_Contains3(const QCborMap* self, QCborValue* key);
extern __declspec(dllexport) int QCborMap_Compare(const QCborMap* self, QCborMap* other);
extern __declspec(dllexport) iterator QCborMap_Begin(QCborMap* self);
extern __declspec(dllexport) const_iterator QCborMap_ConstBegin(const QCborMap* self);
extern __declspec(dllexport) const_iterator QCborMap_Begin2(const QCborMap* self);
extern __declspec(dllexport) const_iterator QCborMap_Cbegin(const QCborMap* self);
extern __declspec(dllexport) iterator QCborMap_End(QCborMap* self);
extern __declspec(dllexport) const_iterator QCborMap_ConstEnd(const QCborMap* self);
extern __declspec(dllexport) const_iterator QCborMap_End2(const QCborMap* self);
extern __declspec(dllexport) const_iterator QCborMap_Cend(const QCborMap* self);
extern __declspec(dllexport) iterator QCborMap_Erase(QCborMap* self, iterator it);
extern __declspec(dllexport) iterator QCborMap_EraseWithIt(QCborMap* self, const_iterator it);
extern __declspec(dllexport) QCborValue* QCborMap_Extract(QCborMap* self, iterator it);
extern __declspec(dllexport) QCborValue* QCborMap_ExtractWithIt(QCborMap* self, const_iterator it);
extern __declspec(dllexport) bool QCborMap_Empty(const QCborMap* self);
extern __declspec(dllexport) iterator QCborMap_Find(QCborMap* self, long long key);
extern __declspec(dllexport) iterator QCborMap_Find2(QCborMap* self, struct miqt_string key);
extern __declspec(dllexport) iterator QCborMap_Find3(QCborMap* self, QCborValue* key);
extern __declspec(dllexport) const_iterator QCborMap_ConstFind(const QCborMap* self, long long key);
extern __declspec(dllexport) const_iterator QCborMap_ConstFind2(const QCborMap* self, struct miqt_string key);
extern __declspec(dllexport) const_iterator QCborMap_ConstFind3(const QCborMap* self, QCborValue* key);
extern __declspec(dllexport) const_iterator QCborMap_Find4(const QCborMap* self, long long key);
extern __declspec(dllexport) const_iterator QCborMap_Find6(const QCborMap* self, struct miqt_string key);
extern __declspec(dllexport) const_iterator QCborMap_Find7(const QCborMap* self, QCborValue* key);
extern __declspec(dllexport) iterator QCborMap_Insert(QCborMap* self, long long key, QCborValue* value_);
extern __declspec(dllexport) iterator QCborMap_Insert3(QCborMap* self, struct miqt_string key, QCborValue* value_);
extern __declspec(dllexport) iterator QCborMap_Insert4(QCborMap* self, QCborValue* key, QCborValue* value_);
extern __declspec(dllexport) iterator QCborMap_InsertWithValueType(QCborMap* self, value_type v);
extern __declspec(dllexport) QCborMap* QCborMap_FromVariantMap(struct miqt_map /* of struct miqt_string to QVariant* */  mapVal);
extern __declspec(dllexport) QCborMap* QCborMap_FromVariantHash(struct miqt_map /* of struct miqt_string to QVariant* */  hash);
extern __declspec(dllexport) QCborMap* QCborMap_FromJsonObject(QJsonObject* o);
extern __declspec(dllexport) struct miqt_map /* of struct miqt_string to QVariant* */  QCborMap_ToVariantMap(const QCborMap* self);
extern __declspec(dllexport) struct miqt_map /* of struct miqt_string to QVariant* */  QCborMap_ToVariantHash(const QCborMap* self);
extern __declspec(dllexport) QJsonObject* QCborMap_ToJsonObject(const QCborMap* self);
extern __declspec(dllexport) void QCborMap_Delete(QCborMap* self, bool isSubclass);

extern __declspec(dllexport) QCborMap__Iterator* QCborMap__Iterator_new();
extern __declspec(dllexport) QCborMap__Iterator* QCborMap__Iterator_new2(const Iterator* param1);
extern __declspec(dllexport) void QCborMap__Iterator_OperatorAssign(QCborMap__Iterator* self, const Iterator* other);
extern __declspec(dllexport) value_type QCborMap__Iterator_OperatorMultiply(const QCborMap__Iterator* self);
extern __declspec(dllexport) value_type QCborMap__Iterator_OperatorSubscript(const QCborMap__Iterator* self, ptrdiff_t j);
extern __declspec(dllexport) QCborValueRef* QCborMap__Iterator_OperatorMinusGreater(QCborMap__Iterator* self);
extern __declspec(dllexport) QCborValueConstRef* QCborMap__Iterator_OperatorMinusGreater2(const QCborMap__Iterator* self);
extern __declspec(dllexport) QCborValue* QCborMap__Iterator_Key(const QCborMap__Iterator* self);
extern __declspec(dllexport) QCborValueRef* QCborMap__Iterator_Value(const QCborMap__Iterator* self);
extern __declspec(dllexport) Iterator* QCborMap__Iterator_OperatorPlusPlus(QCborMap__Iterator* self);
extern __declspec(dllexport) Iterator QCborMap__Iterator_OperatorPlusPlusWithInt(QCborMap__Iterator* self, int param1);
extern __declspec(dllexport) Iterator* QCborMap__Iterator_OperatorMinusMinus(QCborMap__Iterator* self);
extern __declspec(dllexport) Iterator QCborMap__Iterator_OperatorMinusMinusWithInt(QCborMap__Iterator* self, int param1);
extern __declspec(dllexport) Iterator* QCborMap__Iterator_OperatorPlusAssign(QCborMap__Iterator* self, ptrdiff_t j);
extern __declspec(dllexport) Iterator* QCborMap__Iterator_OperatorMinusAssign(QCborMap__Iterator* self, ptrdiff_t j);
extern __declspec(dllexport) Iterator QCborMap__Iterator_OperatorPlus(const QCborMap__Iterator* self, ptrdiff_t j);
extern __declspec(dllexport) Iterator QCborMap__Iterator_OperatorMinus(const QCborMap__Iterator* self, ptrdiff_t j);
extern __declspec(dllexport) ptrdiff_t QCborMap__Iterator_OperatorMinusWithIterator(const QCborMap__Iterator* self, Iterator j);
extern __declspec(dllexport) void QCborMap__Iterator_Delete(QCborMap__Iterator* self, bool isSubclass);

extern __declspec(dllexport) QCborMap__ConstIterator* QCborMap__ConstIterator_new();
extern __declspec(dllexport) QCborMap__ConstIterator* QCborMap__ConstIterator_new2(const ConstIterator* param1);
extern __declspec(dllexport) void QCborMap__ConstIterator_OperatorAssign(QCborMap__ConstIterator* self, const ConstIterator* other);
extern __declspec(dllexport) value_type QCborMap__ConstIterator_OperatorMultiply(const QCborMap__ConstIterator* self);
extern __declspec(dllexport) value_type QCborMap__ConstIterator_OperatorSubscript(const QCborMap__ConstIterator* self, ptrdiff_t j);
extern __declspec(dllexport) QCborValueConstRef* QCborMap__ConstIterator_OperatorMinusGreater(const QCborMap__ConstIterator* self);
extern __declspec(dllexport) QCborValue* QCborMap__ConstIterator_Key(const QCborMap__ConstIterator* self);
extern __declspec(dllexport) QCborValueConstRef* QCborMap__ConstIterator_Value(const QCborMap__ConstIterator* self);
extern __declspec(dllexport) ConstIterator* QCborMap__ConstIterator_OperatorPlusPlus(QCborMap__ConstIterator* self);
extern __declspec(dllexport) ConstIterator QCborMap__ConstIterator_OperatorPlusPlusWithInt(QCborMap__ConstIterator* self, int param1);
extern __declspec(dllexport) ConstIterator* QCborMap__ConstIterator_OperatorMinusMinus(QCborMap__ConstIterator* self);
extern __declspec(dllexport) ConstIterator QCborMap__ConstIterator_OperatorMinusMinusWithInt(QCborMap__ConstIterator* self, int param1);
extern __declspec(dllexport) ConstIterator* QCborMap__ConstIterator_OperatorPlusAssign(QCborMap__ConstIterator* self, ptrdiff_t j);
extern __declspec(dllexport) ConstIterator* QCborMap__ConstIterator_OperatorMinusAssign(QCborMap__ConstIterator* self, ptrdiff_t j);
extern __declspec(dllexport) ConstIterator QCborMap__ConstIterator_OperatorPlus(const QCborMap__ConstIterator* self, ptrdiff_t j);
extern __declspec(dllexport) ConstIterator QCborMap__ConstIterator_OperatorMinus(const QCborMap__ConstIterator* self, ptrdiff_t j);
extern __declspec(dllexport) ptrdiff_t QCborMap__ConstIterator_OperatorMinusWithConstIterator(const QCborMap__ConstIterator* self, ConstIterator j);
extern __declspec(dllexport) void QCborMap__ConstIterator_Delete(QCborMap__ConstIterator* self, bool isSubclass);

} 
