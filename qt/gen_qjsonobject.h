#pragma once
#ifndef MIQT_QT_GEN_QJSONOBJECT_H
#define MIQT_QT_GEN_QJSONOBJECT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QJsonObject__const_iterator)
typedef QJsonObject::const_iterator QJsonObject__const_iterator;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QJsonObject__iterator)
typedef QJsonObject::iterator QJsonObject__iterator;
typedef struct QJsonObject QJsonObject;
typedef struct QJsonObject__const_iterator QJsonObject__const_iterator;
typedef struct QJsonObject__iterator QJsonObject__iterator;
typedef struct QJsonValue QJsonValue;
typedef struct QJsonValueConstRef QJsonValueConstRef;
typedef struct QJsonValueRef QJsonValueRef;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QJsonObject* QJsonObject_new();
extern __declspec(dllexport) 
QJsonObject* QJsonObject_new2(QJsonObject* other);
extern __declspec(dllexport) 
void QJsonObject_OperatorAssign(QJsonObject* self, QJsonObject* other);
extern __declspec(dllexport) 
void QJsonObject_Swap(QJsonObject* self, QJsonObject* other);
extern __declspec(dllexport) 
QJsonObject* QJsonObject_FromVariantMap(struct miqt_map /* of struct miqt_string to QVariant* */  mapVal);
extern __declspec(dllexport) 
struct miqt_map /* of struct miqt_string to QVariant* */  QJsonObject_ToVariantMap(const QJsonObject* self);
extern __declspec(dllexport) 
QJsonObject* QJsonObject_FromVariantHash(struct miqt_map /* of struct miqt_string to QVariant* */  mapVal);
extern __declspec(dllexport) 
struct miqt_map /* of struct miqt_string to QVariant* */  QJsonObject_ToVariantHash(const QJsonObject* self);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QJsonObject_Keys(const QJsonObject* self);
extern __declspec(dllexport) 
ptrdiff_t QJsonObject_Size(const QJsonObject* self);
extern __declspec(dllexport) 
ptrdiff_t QJsonObject_Count(const QJsonObject* self);
extern __declspec(dllexport) 
ptrdiff_t QJsonObject_Length(const QJsonObject* self);
extern __declspec(dllexport) 
bool QJsonObject_IsEmpty(const QJsonObject* self);
extern __declspec(dllexport) 
QJsonValue* QJsonObject_Value(const QJsonObject* self, struct miqt_string key);
extern __declspec(dllexport) 
QJsonValue* QJsonObject_OperatorSubscript(const QJsonObject* self, struct miqt_string key);
extern __declspec(dllexport) 
QJsonValueRef* QJsonObject_OperatorSubscriptWithKey(QJsonObject* self, struct miqt_string key);
extern __declspec(dllexport) 
void QJsonObject_Remove(QJsonObject* self, struct miqt_string key);
extern __declspec(dllexport) 
QJsonValue* QJsonObject_Take(QJsonObject* self, struct miqt_string key);
extern __declspec(dllexport) 
bool QJsonObject_Contains(const QJsonObject* self, struct miqt_string key);
extern __declspec(dllexport) 
iterator QJsonObject_Begin(QJsonObject* self);
extern __declspec(dllexport) 
const_iterator QJsonObject_Begin2(const QJsonObject* self);
extern __declspec(dllexport) 
const_iterator QJsonObject_ConstBegin(const QJsonObject* self);
extern __declspec(dllexport) 
iterator QJsonObject_End(QJsonObject* self);
extern __declspec(dllexport) 
const_iterator QJsonObject_End2(const QJsonObject* self);
extern __declspec(dllexport) 
const_iterator QJsonObject_ConstEnd(const QJsonObject* self);
extern __declspec(dllexport) 
iterator QJsonObject_Erase(QJsonObject* self, iterator it);
extern __declspec(dllexport) 
iterator QJsonObject_Find(QJsonObject* self, struct miqt_string key);
extern __declspec(dllexport) 
const_iterator QJsonObject_FindWithKey(const QJsonObject* self, struct miqt_string key);
extern __declspec(dllexport) 
const_iterator QJsonObject_ConstFind(const QJsonObject* self, struct miqt_string key);
extern __declspec(dllexport) 
iterator QJsonObject_Insert(QJsonObject* self, struct miqt_string key, QJsonValue* value);
extern __declspec(dllexport) 
bool QJsonObject_Empty(const QJsonObject* self);
extern __declspec(dllexport) 
void QJsonObject_Delete(QJsonObject* self, bool isSubclass);

extern __declspec(dllexport) 
QJsonObject__iterator* QJsonObject__iterator_new();
extern __declspec(dllexport) 
QJsonObject__iterator* QJsonObject__iterator_new2(QJsonObject* obj, ptrdiff_t index);
extern __declspec(dllexport) 
QJsonObject__iterator* QJsonObject__iterator_new3(const iterator* other);
extern __declspec(dllexport) 
void QJsonObject__iterator_OperatorAssign(QJsonObject__iterator* self, const iterator* other);
extern __declspec(dllexport) 
struct miqt_string QJsonObject__iterator_Key(const QJsonObject__iterator* self);
extern __declspec(dllexport) 
QJsonValueRef* QJsonObject__iterator_Value(const QJsonObject__iterator* self);
extern __declspec(dllexport) 
QJsonValueRef* QJsonObject__iterator_OperatorMultiply(const QJsonObject__iterator* self);
extern __declspec(dllexport) 
QJsonValueConstRef* QJsonObject__iterator_OperatorMinusGreater(const QJsonObject__iterator* self);
extern __declspec(dllexport) 
QJsonValueRef* QJsonObject__iterator_OperatorMinusGreater2(QJsonObject__iterator* self);
extern __declspec(dllexport) 
QJsonValueRef* QJsonObject__iterator_OperatorSubscript(const QJsonObject__iterator* self, ptrdiff_t j);
extern __declspec(dllexport) 
iterator* QJsonObject__iterator_OperatorPlusPlus(QJsonObject__iterator* self);
extern __declspec(dllexport) 
iterator QJsonObject__iterator_OperatorPlusPlusWithInt(QJsonObject__iterator* self, int param1);
extern __declspec(dllexport) 
iterator* QJsonObject__iterator_OperatorMinusMinus(QJsonObject__iterator* self);
extern __declspec(dllexport) 
iterator QJsonObject__iterator_OperatorMinusMinusWithInt(QJsonObject__iterator* self, int param1);
extern __declspec(dllexport) 
iterator QJsonObject__iterator_OperatorPlus(const QJsonObject__iterator* self, ptrdiff_t j);
extern __declspec(dllexport) 
iterator QJsonObject__iterator_OperatorMinus(const QJsonObject__iterator* self, ptrdiff_t j);
extern __declspec(dllexport) 
iterator* QJsonObject__iterator_OperatorPlusAssign(QJsonObject__iterator* self, ptrdiff_t j);
extern __declspec(dllexport) 
iterator* QJsonObject__iterator_OperatorMinusAssign(QJsonObject__iterator* self, ptrdiff_t j);
extern __declspec(dllexport) 
ptrdiff_t QJsonObject__iterator_OperatorMinusWithIterator(const QJsonObject__iterator* self, iterator j);
extern __declspec(dllexport) 
void QJsonObject__iterator_Delete(QJsonObject__iterator* self, bool isSubclass);

extern __declspec(dllexport) 
QJsonObject__const_iterator* QJsonObject__const_iterator_new();
extern __declspec(dllexport) 
QJsonObject__const_iterator* QJsonObject__const_iterator_new2(QJsonObject* obj, ptrdiff_t index);
extern __declspec(dllexport) 
QJsonObject__const_iterator* QJsonObject__const_iterator_new3(const iterator* other);
extern __declspec(dllexport) 
QJsonObject__const_iterator* QJsonObject__const_iterator_new4(const const_iterator* other);
extern __declspec(dllexport) 
void QJsonObject__const_iterator_OperatorAssign(QJsonObject__const_iterator* self, const const_iterator* other);
extern __declspec(dllexport) 
struct miqt_string QJsonObject__const_iterator_Key(const QJsonObject__const_iterator* self);
extern __declspec(dllexport) 
QJsonValueConstRef* QJsonObject__const_iterator_Value(const QJsonObject__const_iterator* self);
extern __declspec(dllexport) 
QJsonValueConstRef* QJsonObject__const_iterator_OperatorMultiply(const QJsonObject__const_iterator* self);
extern __declspec(dllexport) 
QJsonValueConstRef* QJsonObject__const_iterator_OperatorMinusGreater(const QJsonObject__const_iterator* self);
extern __declspec(dllexport) 
QJsonValueConstRef* QJsonObject__const_iterator_OperatorSubscript(const QJsonObject__const_iterator* self, ptrdiff_t j);
extern __declspec(dllexport) 
const_iterator* QJsonObject__const_iterator_OperatorPlusPlus(QJsonObject__const_iterator* self);
extern __declspec(dllexport) 
const_iterator QJsonObject__const_iterator_OperatorPlusPlusWithInt(QJsonObject__const_iterator* self, int param1);
extern __declspec(dllexport) 
const_iterator* QJsonObject__const_iterator_OperatorMinusMinus(QJsonObject__const_iterator* self);
extern __declspec(dllexport) 
const_iterator QJsonObject__const_iterator_OperatorMinusMinusWithInt(QJsonObject__const_iterator* self, int param1);
extern __declspec(dllexport) 
const_iterator QJsonObject__const_iterator_OperatorPlus(const QJsonObject__const_iterator* self, ptrdiff_t j);
extern __declspec(dllexport) 
const_iterator QJsonObject__const_iterator_OperatorMinus(const QJsonObject__const_iterator* self, ptrdiff_t j);
extern __declspec(dllexport) 
const_iterator* QJsonObject__const_iterator_OperatorPlusAssign(QJsonObject__const_iterator* self, ptrdiff_t j);
extern __declspec(dllexport) 
const_iterator* QJsonObject__const_iterator_OperatorMinusAssign(QJsonObject__const_iterator* self, ptrdiff_t j);
extern __declspec(dllexport) 
ptrdiff_t QJsonObject__const_iterator_OperatorMinusWithConstIterator(const QJsonObject__const_iterator* self, const_iterator j);
extern __declspec(dllexport) 
void QJsonObject__const_iterator_Delete(QJsonObject__const_iterator* self, bool isSubclass);

}
