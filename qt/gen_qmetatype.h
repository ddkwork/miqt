#pragma once
#ifndef MIQT_QT_GEN_QMETATYPE_H
#define MIQT_QT_GEN_QMETATYPE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_Disambiguated_t)
typedef Qt::Disambiguated_t Disambiguated_t;
typedef struct QByteArrayView QByteArrayView;
typedef struct QDataStream QDataStream;
typedef struct QDebug QDebug;
typedef struct QMetaObject QMetaObject;
typedef struct QMetaType QMetaType;
typedef struct QPartialOrdering QPartialOrdering;
typedef struct Disambiguated_t Disambiguated_t;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QMetaType* QMetaType_new(int typeVal);
extern __declspec(dllexport) QMetaType* QMetaType_new2();
extern __declspec(dllexport) QMetaType* QMetaType_new3(QMetaType* param1);
extern __declspec(dllexport) void QMetaType_RegisterNormalizedTypedef(struct miqt_string normalizedTypeName, QMetaType* typeVal);
extern __declspec(dllexport) int QMetaType_Type(const char* typeName);
extern __declspec(dllexport) int QMetaType_TypeWithTypeName(struct miqt_string typeName);
extern __declspec(dllexport) const char* QMetaType_TypeName(int typeVal);
extern __declspec(dllexport) int QMetaType_SizeOf(int typeVal);
extern __declspec(dllexport) TypeFlags QMetaType_TypeFlags(int typeVal);
extern __declspec(dllexport) QMetaObject* QMetaType_MetaObjectForType(int typeVal);
extern __declspec(dllexport) void* QMetaType_Create(int typeVal);
extern __declspec(dllexport) void QMetaType_Destroy(int typeVal, void* data);
extern __declspec(dllexport) void* QMetaType_Construct(int typeVal, void* where, const void* copyVal);
extern __declspec(dllexport) void QMetaType_Destruct(int typeVal, void* where);
extern __declspec(dllexport) bool QMetaType_IsRegistered(int typeVal);
extern __declspec(dllexport) bool QMetaType_IsValid(const QMetaType* self);
extern __declspec(dllexport) bool QMetaType_IsRegistered2(const QMetaType* self);
extern __declspec(dllexport) void QMetaType_RegisterType(const QMetaType* self);
extern __declspec(dllexport) int QMetaType_Id(const QMetaType* self);
extern __declspec(dllexport) ptrdiff_t QMetaType_SizeOf2(const QMetaType* self);
extern __declspec(dllexport) ptrdiff_t QMetaType_AlignOf(const QMetaType* self);
extern __declspec(dllexport) TypeFlags QMetaType_Flags(const QMetaType* self);
extern __declspec(dllexport) QMetaObject* QMetaType_MetaObject(const QMetaType* self);
extern __declspec(dllexport) const char* QMetaType_Name(const QMetaType* self);
extern __declspec(dllexport) void* QMetaType_Create2(const QMetaType* self);
extern __declspec(dllexport) void QMetaType_DestroyWithData(const QMetaType* self, void* data);
extern __declspec(dllexport) void* QMetaType_ConstructWithWhere(const QMetaType* self, void* where);
extern __declspec(dllexport) void QMetaType_DestructWithData(const QMetaType* self, void* data);
extern __declspec(dllexport) QPartialOrdering* QMetaType_Compare(const QMetaType* self, const void* lhs, const void* rhs);
extern __declspec(dllexport) bool QMetaType_Equals(const QMetaType* self, const void* lhs, const void* rhs);
extern __declspec(dllexport) bool QMetaType_IsDefaultConstructible(const QMetaType* self);
extern __declspec(dllexport) bool QMetaType_IsCopyConstructible(const QMetaType* self);
extern __declspec(dllexport) bool QMetaType_IsMoveConstructible(const QMetaType* self);
extern __declspec(dllexport) bool QMetaType_IsDestructible(const QMetaType* self);
extern __declspec(dllexport) bool QMetaType_IsEqualityComparable(const QMetaType* self);
extern __declspec(dllexport) bool QMetaType_IsOrdered(const QMetaType* self);
extern __declspec(dllexport) bool QMetaType_Save(const QMetaType* self, QDataStream* stream, const void* data);
extern __declspec(dllexport) bool QMetaType_Load(const QMetaType* self, QDataStream* stream, void* data);
extern __declspec(dllexport) bool QMetaType_HasRegisteredDataStreamOperators(const QMetaType* self);
extern __declspec(dllexport) bool QMetaType_Save2(QDataStream* stream, int typeVal, const void* data);
extern __declspec(dllexport) bool QMetaType_Load2(QDataStream* stream, int typeVal, void* data);
extern __declspec(dllexport) QMetaType* QMetaType_UnderlyingType(const QMetaType* self);
extern __declspec(dllexport) QMetaType* QMetaType_FromName(QByteArrayView* name);
extern __declspec(dllexport) bool QMetaType_DebugStream(QMetaType* self, QDebug* dbg, const void* rhs);
extern __declspec(dllexport) bool QMetaType_HasRegisteredDebugStreamOperator(const QMetaType* self);
extern __declspec(dllexport) bool QMetaType_DebugStream2(QDebug* dbg, const void* rhs, int typeId);
extern __declspec(dllexport) bool QMetaType_HasRegisteredDebugStreamOperatorWithTypeId(int typeId);
extern __declspec(dllexport) bool QMetaType_Convert(QMetaType* fromType, const void* from, QMetaType* toType, void* to);
extern __declspec(dllexport) bool QMetaType_CanConvert(QMetaType* fromType, QMetaType* toType);
extern __declspec(dllexport) bool QMetaType_View(QMetaType* fromType, void* from, QMetaType* toType, void* to);
extern __declspec(dllexport) bool QMetaType_CanView(QMetaType* fromType, QMetaType* toType);
extern __declspec(dllexport) bool QMetaType_Convert2(const void* from, int fromTypeId, void* to, int toTypeId);
extern __declspec(dllexport) bool QMetaType_Compare2(const void* lhs, const void* rhs, int typeId, int* result);
extern __declspec(dllexport) bool QMetaType_Equals2(const void* lhs, const void* rhs, int typeId, int* result);
extern __declspec(dllexport) bool QMetaType_HasRegisteredConverterFunction(QMetaType* fromType, QMetaType* toType);
extern __declspec(dllexport) bool QMetaType_HasRegisteredMutableViewFunction(QMetaType* fromType, QMetaType* toType);
extern __declspec(dllexport) bool QMetaType_RegisterConverterFunction(const ConverterFunction* f, QMetaType* from, QMetaType* to);
extern __declspec(dllexport) void QMetaType_UnregisterConverterFunction(QMetaType* from, QMetaType* to);
extern __declspec(dllexport) bool QMetaType_RegisterMutableViewFunction(const MutableViewFunction* f, QMetaType* from, QMetaType* to);
extern __declspec(dllexport) void QMetaType_UnregisterMutableViewFunction(QMetaType* from, QMetaType* to);
extern __declspec(dllexport) void QMetaType_UnregisterMetaType(QMetaType* typeVal);
extern __declspec(dllexport) void* QMetaType_Create22(int typeVal, const void* copyVal);
extern __declspec(dllexport) bool QMetaType_IsValid1(const QMetaType* self, Disambiguated_t* param1);
extern __declspec(dllexport) bool QMetaType_IsRegistered1(const QMetaType* self, Disambiguated_t* param1);
extern __declspec(dllexport) int QMetaType_Id1(const QMetaType* self, int param1);
extern __declspec(dllexport) void* QMetaType_Create1(const QMetaType* self, const void* copyVal);
extern __declspec(dllexport) void* QMetaType_Construct2(const QMetaType* self, void* where, const void* copyVal);
extern __declspec(dllexport) void QMetaType_Delete(QMetaType* self, bool isSubclass);

} 
