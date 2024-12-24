#pragma once
#ifndef MIQT_QT_CBOR_GEN_QCBORVALUE_H
#define MIQT_QT_CBOR_GEN_QCBORVALUE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QCborArray;
class QCborMap;
class QCborParserError;
class QCborStreamReader;
class QCborStreamWriter;
class QCborValue;
class QCborValueConstRef;
class QCborValueRef;
class QDateTime;
class QJsonValue;
class QRegularExpression;
class QUrl;
class QUuid;
class QVariant;
class _GUID;
class type_info;
#else
typedef struct QCborArray QCborArray;
typedef struct QCborMap QCborMap;
typedef struct QCborParserError QCborParserError;
typedef struct QCborStreamReader QCborStreamReader;
typedef struct QCborStreamWriter QCborStreamWriter;
typedef struct QCborValue QCborValue;
typedef struct QCborValueConstRef QCborValueConstRef;
typedef struct QCborValueRef QCborValueRef;
typedef struct QDateTime QDateTime;
typedef struct QJsonValue QJsonValue;
typedef struct QRegularExpression QRegularExpression;
typedef struct QUrl QUrl;
typedef struct QUuid QUuid;
typedef struct QVariant QVariant;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) struct miqt_string QCborParserError_ErrorString(const QCborParserError* self);
extern __declspec(dllexport) void QCborParserError_Delete(QCborParserError* self, bool isSubclass);

extern __declspec(dllexport) QCborValue* QCborValue_new();
extern __declspec(dllexport) QCborValue* QCborValue_new2(Type t_);
extern __declspec(dllexport) QCborValue* QCborValue_new3(bool b_);
extern __declspec(dllexport) QCborValue* QCborValue_new4(int i);
extern __declspec(dllexport) QCborValue* QCborValue_new5(unsigned int u);
extern __declspec(dllexport) QCborValue* QCborValue_new6(long long i);
extern __declspec(dllexport) QCborValue* QCborValue_new7(double v);
extern __declspec(dllexport) QCborValue* QCborValue_new8(uint8_t st);
extern __declspec(dllexport) QCborValue* QCborValue_new9(struct miqt_string ba);
extern __declspec(dllexport) QCborValue* QCborValue_new10(struct miqt_string s);
extern __declspec(dllexport) QCborValue* QCborValue_new11(const char* s);
extern __declspec(dllexport) QCborValue* QCborValue_new12(QCborArray* a);
extern __declspec(dllexport) QCborValue* QCborValue_new13(QCborMap* m);
extern __declspec(dllexport) QCborValue* QCborValue_new14(uint64_t tag);
extern __declspec(dllexport) QCborValue* QCborValue_new15(int t_);
extern __declspec(dllexport) QCborValue* QCborValue_new16(QDateTime* dt);
extern __declspec(dllexport) QCborValue* QCborValue_new17(QUrl* url);
extern __declspec(dllexport) QCborValue* QCborValue_new18(QRegularExpression* rx);
extern __declspec(dllexport) QCborValue* QCborValue_new19(QUuid* uuid);
extern __declspec(dllexport) QCborValue* QCborValue_new20(QCborValue* other);
extern __declspec(dllexport) QCborValue* QCborValue_new21(uint64_t tag, QCborValue* taggedValue);
extern __declspec(dllexport) QCborValue* QCborValue_new22(int t_, QCborValue* tv);
extern __declspec(dllexport) void QCborValue_OperatorAssign(QCborValue* self, QCborValue* other);
extern __declspec(dllexport) void QCborValue_Swap(QCborValue* self, QCborValue* other);
extern __declspec(dllexport) Type QCborValue_Type(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsInteger(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsByteArray(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsString(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsArray(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsMap(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsTag(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsFalse(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsTrue(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsBool(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsNull(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsUndefined(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsDouble(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsDateTime(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsUrl(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsRegularExpression(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsUuid(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsInvalid(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsContainer(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsSimpleType(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_IsSimpleTypeWithSt(const QCborValue* self, uint8_t st);
extern __declspec(dllexport) uint8_t QCborValue_ToSimpleType(const QCborValue* self);
extern __declspec(dllexport) long long QCborValue_ToInteger(const QCborValue* self);
extern __declspec(dllexport) bool QCborValue_ToBool(const QCborValue* self);
extern __declspec(dllexport) double QCborValue_ToDouble(const QCborValue* self);
extern __declspec(dllexport) uint64_t QCborValue_Tag(const QCborValue* self);
extern __declspec(dllexport) QCborValue* QCborValue_TaggedValue(const QCborValue* self);
extern __declspec(dllexport) struct miqt_string QCborValue_ToByteArray(const QCborValue* self);
extern __declspec(dllexport) struct miqt_string QCborValue_ToString(const QCborValue* self);
extern __declspec(dllexport) QDateTime* QCborValue_ToDateTime(const QCborValue* self);
extern __declspec(dllexport) QUrl* QCborValue_ToUrl(const QCborValue* self);
extern __declspec(dllexport) QRegularExpression* QCborValue_ToRegularExpression(const QCborValue* self);
extern __declspec(dllexport) QUuid* QCborValue_ToUuid(const QCborValue* self);
extern __declspec(dllexport) QCborArray* QCborValue_ToArray(const QCborValue* self);
extern __declspec(dllexport) QCborArray* QCborValue_ToArrayWithDefaultValue(const QCborValue* self, QCborArray* defaultValue);
extern __declspec(dllexport) QCborMap* QCborValue_ToMap(const QCborValue* self);
extern __declspec(dllexport) QCborMap* QCborValue_ToMapWithDefaultValue(const QCborValue* self, QCborMap* defaultValue);
extern __declspec(dllexport) QCborValue* QCborValue_OperatorSubscript(const QCborValue* self, struct miqt_string key);
extern __declspec(dllexport) QCborValue* QCborValue_OperatorSubscript2(const QCborValue* self, long long key);
extern __declspec(dllexport) QCborValueRef* QCborValue_OperatorSubscript3(QCborValue* self, long long key);
extern __declspec(dllexport) QCborValueRef* QCborValue_OperatorSubscript5(QCborValue* self, struct miqt_string key);
extern __declspec(dllexport) int QCborValue_Compare(const QCborValue* self, QCborValue* other);
extern __declspec(dllexport) QCborValue* QCborValue_FromVariant(QVariant* variant);
extern __declspec(dllexport) QVariant* QCborValue_ToVariant(const QCborValue* self);
extern __declspec(dllexport) QCborValue* QCborValue_FromJsonValue(QJsonValue* v);
extern __declspec(dllexport) QJsonValue* QCborValue_ToJsonValue(const QCborValue* self);
extern __declspec(dllexport) QCborValue* QCborValue_FromCbor(QCborStreamReader* reader);
extern __declspec(dllexport) QCborValue* QCborValue_FromCborWithBa(struct miqt_string ba);
extern __declspec(dllexport) QCborValue* QCborValue_FromCbor2(const char* data, ptrdiff_t lenVal);
extern __declspec(dllexport) QCborValue* QCborValue_FromCbor3(const unsigned char* data, ptrdiff_t lenVal);
extern __declspec(dllexport) struct miqt_string QCborValue_ToCbor(const QCborValue* self);
extern __declspec(dllexport) void QCborValue_ToCborWithWriter(const QCborValue* self, QCborStreamWriter* writer);
extern __declspec(dllexport) struct miqt_string QCborValue_ToDiagnosticNotation(const QCborValue* self);
extern __declspec(dllexport) uint8_t QCborValue_ToSimpleType1(const QCborValue* self, uint8_t defaultValue);
extern __declspec(dllexport) long long QCborValue_ToInteger1(const QCborValue* self, long long defaultValue);
extern __declspec(dllexport) bool QCborValue_ToBool1(const QCborValue* self, bool defaultValue);
extern __declspec(dllexport) double QCborValue_ToDouble1(const QCborValue* self, double defaultValue);
extern __declspec(dllexport) uint64_t QCborValue_Tag1(const QCborValue* self, uint64_t defaultValue);
extern __declspec(dllexport) QCborValue* QCborValue_TaggedValue1(const QCborValue* self, QCborValue* defaultValue);
extern __declspec(dllexport) struct miqt_string QCborValue_ToByteArray1(const QCborValue* self, struct miqt_string defaultValue);
extern __declspec(dllexport) struct miqt_string QCborValue_ToString1(const QCborValue* self, struct miqt_string defaultValue);
extern __declspec(dllexport) QDateTime* QCborValue_ToDateTime1(const QCborValue* self, QDateTime* defaultValue);
extern __declspec(dllexport) QUrl* QCborValue_ToUrl1(const QCborValue* self, QUrl* defaultValue);
extern __declspec(dllexport) QRegularExpression* QCborValue_ToRegularExpression1(const QCborValue* self, QRegularExpression* defaultValue);
extern __declspec(dllexport) QUuid* QCborValue_ToUuid1(const QCborValue* self, QUuid* defaultValue);
extern __declspec(dllexport) QCborValue* QCborValue_FromCbor22(struct miqt_string ba, QCborParserError* error);
extern __declspec(dllexport) QCborValue* QCborValue_FromCbor32(const char* data, ptrdiff_t lenVal, QCborParserError* error);
extern __declspec(dllexport) QCborValue* QCborValue_FromCbor33(const unsigned char* data, ptrdiff_t lenVal, QCborParserError* error);
extern __declspec(dllexport) struct miqt_string QCborValue_ToCbor1(const QCborValue* self, EncodingOptions opt);
extern __declspec(dllexport) void QCborValue_ToCbor2(const QCborValue* self, QCborStreamWriter* writer, EncodingOptions opt);
extern __declspec(dllexport) struct miqt_string QCborValue_ToDiagnosticNotation1(const QCborValue* self, DiagnosticNotationOptions opts);
extern __declspec(dllexport) void QCborValue_Delete(QCborValue* self, bool isSubclass);

extern __declspec(dllexport) QCborValueConstRef* QCborValueConstRef_new(QCborValueConstRef* param1);
extern __declspec(dllexport) int QCborValueConstRef_Type(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsInteger(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsByteArray(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsString(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsArray(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsMap(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsTag(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsFalse(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsTrue(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsBool(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsNull(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsUndefined(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsDouble(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsDateTime(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsUrl(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsRegularExpression(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsUuid(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsInvalid(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsContainer(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsSimpleType(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_IsSimpleTypeWithSt(const QCborValueConstRef* self, uint8_t st);
extern __declspec(dllexport) uint8_t QCborValueConstRef_ToSimpleType(const QCborValueConstRef* self);
extern __declspec(dllexport) uint64_t QCborValueConstRef_Tag(const QCborValueConstRef* self);
extern __declspec(dllexport) QCborValue* QCborValueConstRef_TaggedValue(const QCborValueConstRef* self);
extern __declspec(dllexport) long long QCborValueConstRef_ToInteger(const QCborValueConstRef* self);
extern __declspec(dllexport) bool QCborValueConstRef_ToBool(const QCborValueConstRef* self);
extern __declspec(dllexport) double QCborValueConstRef_ToDouble(const QCborValueConstRef* self);
extern __declspec(dllexport) struct miqt_string QCborValueConstRef_ToByteArray(const QCborValueConstRef* self);
extern __declspec(dllexport) struct miqt_string QCborValueConstRef_ToString(const QCborValueConstRef* self);
extern __declspec(dllexport) QDateTime* QCborValueConstRef_ToDateTime(const QCborValueConstRef* self);
extern __declspec(dllexport) QUrl* QCborValueConstRef_ToUrl(const QCborValueConstRef* self);
extern __declspec(dllexport) QRegularExpression* QCborValueConstRef_ToRegularExpression(const QCborValueConstRef* self);
extern __declspec(dllexport) QUuid* QCborValueConstRef_ToUuid(const QCborValueConstRef* self);
extern __declspec(dllexport) QCborArray* QCborValueConstRef_ToArray(const QCborValueConstRef* self);
extern __declspec(dllexport) QCborArray* QCborValueConstRef_ToArrayWithQCborArray(const QCborValueConstRef* self, QCborArray* a);
extern __declspec(dllexport) QCborMap* QCborValueConstRef_ToMap(const QCborValueConstRef* self);
extern __declspec(dllexport) QCborMap* QCborValueConstRef_ToMapWithQCborMap(const QCborValueConstRef* self, QCborMap* m);
extern __declspec(dllexport) QCborValue* QCborValueConstRef_OperatorSubscript(const QCborValueConstRef* self, struct miqt_string key);
extern __declspec(dllexport) QCborValue* QCborValueConstRef_OperatorSubscript2(const QCborValueConstRef* self, long long key);
extern __declspec(dllexport) int QCborValueConstRef_Compare(const QCborValueConstRef* self, QCborValue* other);
extern __declspec(dllexport) QVariant* QCborValueConstRef_ToVariant(const QCborValueConstRef* self);
extern __declspec(dllexport) QJsonValue* QCborValueConstRef_ToJsonValue(const QCborValueConstRef* self);
extern __declspec(dllexport) struct miqt_string QCborValueConstRef_ToCbor(const QCborValueConstRef* self);
extern __declspec(dllexport) void QCborValueConstRef_ToCborWithWriter(const QCborValueConstRef* self, QCborStreamWriter* writer);
extern __declspec(dllexport) struct miqt_string QCborValueConstRef_ToDiagnosticNotation(const QCborValueConstRef* self);
extern __declspec(dllexport) uint8_t QCborValueConstRef_ToSimpleType1(const QCborValueConstRef* self, uint8_t defaultValue);
extern __declspec(dllexport) uint64_t QCborValueConstRef_Tag1(const QCborValueConstRef* self, uint64_t defaultValue);
extern __declspec(dllexport) QCborValue* QCborValueConstRef_TaggedValue1(const QCborValueConstRef* self, QCborValue* defaultValue);
extern __declspec(dllexport) long long QCborValueConstRef_ToInteger1(const QCborValueConstRef* self, long long defaultValue);
extern __declspec(dllexport) bool QCborValueConstRef_ToBool1(const QCborValueConstRef* self, bool defaultValue);
extern __declspec(dllexport) double QCborValueConstRef_ToDouble1(const QCborValueConstRef* self, double defaultValue);
extern __declspec(dllexport) struct miqt_string QCborValueConstRef_ToByteArray1(const QCborValueConstRef* self, struct miqt_string defaultValue);
extern __declspec(dllexport) struct miqt_string QCborValueConstRef_ToString1(const QCborValueConstRef* self, struct miqt_string defaultValue);
extern __declspec(dllexport) QDateTime* QCborValueConstRef_ToDateTime1(const QCborValueConstRef* self, QDateTime* defaultValue);
extern __declspec(dllexport) QUrl* QCborValueConstRef_ToUrl1(const QCborValueConstRef* self, QUrl* defaultValue);
extern __declspec(dllexport) QRegularExpression* QCborValueConstRef_ToRegularExpression1(const QCborValueConstRef* self, QRegularExpression* defaultValue);
extern __declspec(dllexport) QUuid* QCborValueConstRef_ToUuid1(const QCborValueConstRef* self, QUuid* defaultValue);
extern __declspec(dllexport) struct miqt_string QCborValueConstRef_ToCbor1(const QCborValueConstRef* self, int opt);
extern __declspec(dllexport) void QCborValueConstRef_ToCbor2(const QCborValueConstRef* self, QCborStreamWriter* writer, int opt);
extern __declspec(dllexport) struct miqt_string QCborValueConstRef_ToDiagnosticNotation1(const QCborValueConstRef* self, int opt);
extern __declspec(dllexport) void QCborValueConstRef_Delete(QCborValueConstRef* self, bool isSubclass);

extern __declspec(dllexport) QCborValueRef* QCborValueRef_new(QCborValueRef* param1);
extern __declspec(dllexport) void QCborValueRef_virtbase(QCborValueRef* src, QCborValueConstRef** outptr_QCborValueConstRef);
extern __declspec(dllexport) void QCborValueRef_OperatorAssign(QCborValueRef* self, QCborValue* other);
extern __declspec(dllexport) void QCborValueRef_OperatorAssignWithOther(QCborValueRef* self, QCborValueRef* other);
extern __declspec(dllexport) QCborValueRef* QCborValueRef_OperatorSubscript(QCborValueRef* self, long long key);
extern __declspec(dllexport) QCborValueRef* QCborValueRef_OperatorSubscript2(QCborValueRef* self, struct miqt_string key);
extern __declspec(dllexport) int QCborValueRef_Type(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsInteger(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsByteArray(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsString(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsArray(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsMap(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsTag(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsFalse(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsTrue(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsBool(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsNull(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsUndefined(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsDouble(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsDateTime(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsUrl(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsRegularExpression(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsUuid(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsInvalid(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsContainer(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsSimpleType(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_IsSimpleTypeWithSt(const QCborValueRef* self, uint8_t st);
extern __declspec(dllexport) uint8_t QCborValueRef_ToSimpleType(const QCborValueRef* self);
extern __declspec(dllexport) uint64_t QCborValueRef_Tag(const QCborValueRef* self);
extern __declspec(dllexport) QCborValue* QCborValueRef_TaggedValue(const QCborValueRef* self);
extern __declspec(dllexport) long long QCborValueRef_ToInteger(const QCborValueRef* self);
extern __declspec(dllexport) bool QCborValueRef_ToBool(const QCborValueRef* self);
extern __declspec(dllexport) double QCborValueRef_ToDouble(const QCborValueRef* self);
extern __declspec(dllexport) struct miqt_string QCborValueRef_ToByteArray(const QCborValueRef* self);
extern __declspec(dllexport) struct miqt_string QCborValueRef_ToString(const QCborValueRef* self);
extern __declspec(dllexport) QDateTime* QCborValueRef_ToDateTime(const QCborValueRef* self);
extern __declspec(dllexport) QUrl* QCborValueRef_ToUrl(const QCborValueRef* self);
extern __declspec(dllexport) QRegularExpression* QCborValueRef_ToRegularExpression(const QCborValueRef* self);
extern __declspec(dllexport) QUuid* QCborValueRef_ToUuid(const QCborValueRef* self);
extern __declspec(dllexport) QCborArray* QCborValueRef_ToArray(const QCborValueRef* self);
extern __declspec(dllexport) QCborArray* QCborValueRef_ToArrayWithQCborArray(const QCborValueRef* self, QCborArray* a);
extern __declspec(dllexport) QCborMap* QCborValueRef_ToMap(const QCborValueRef* self);
extern __declspec(dllexport) QCborMap* QCborValueRef_ToMapWithQCborMap(const QCborValueRef* self, QCborMap* m);
extern __declspec(dllexport) QCborValue* QCborValueRef_OperatorSubscript3(const QCborValueRef* self, struct miqt_string key);
extern __declspec(dllexport) QCborValue* QCborValueRef_OperatorSubscript5(const QCborValueRef* self, long long key);
extern __declspec(dllexport) int QCborValueRef_Compare(const QCborValueRef* self, QCborValue* other);
extern __declspec(dllexport) QVariant* QCborValueRef_ToVariant(const QCborValueRef* self);
extern __declspec(dllexport) QJsonValue* QCborValueRef_ToJsonValue(const QCborValueRef* self);
extern __declspec(dllexport) struct miqt_string QCborValueRef_ToCbor(QCborValueRef* self);
extern __declspec(dllexport) void QCborValueRef_ToCborWithWriter(QCborValueRef* self, QCborStreamWriter* writer);
extern __declspec(dllexport) struct miqt_string QCborValueRef_ToDiagnosticNotation(QCborValueRef* self);
extern __declspec(dllexport) uint8_t QCborValueRef_ToSimpleType1(const QCborValueRef* self, uint8_t defaultValue);
extern __declspec(dllexport) uint64_t QCborValueRef_Tag1(const QCborValueRef* self, uint64_t defaultValue);
extern __declspec(dllexport) QCborValue* QCborValueRef_TaggedValue1(const QCborValueRef* self, QCborValue* defaultValue);
extern __declspec(dllexport) long long QCborValueRef_ToInteger1(const QCborValueRef* self, long long defaultValue);
extern __declspec(dllexport) bool QCborValueRef_ToBool1(const QCborValueRef* self, bool defaultValue);
extern __declspec(dllexport) double QCborValueRef_ToDouble1(const QCborValueRef* self, double defaultValue);
extern __declspec(dllexport) struct miqt_string QCborValueRef_ToByteArray1(const QCborValueRef* self, struct miqt_string defaultValue);
extern __declspec(dllexport) struct miqt_string QCborValueRef_ToString1(const QCborValueRef* self, struct miqt_string defaultValue);
extern __declspec(dllexport) QDateTime* QCborValueRef_ToDateTime1(const QCborValueRef* self, QDateTime* defaultValue);
extern __declspec(dllexport) QUrl* QCborValueRef_ToUrl1(const QCborValueRef* self, QUrl* defaultValue);
extern __declspec(dllexport) QRegularExpression* QCborValueRef_ToRegularExpression1(const QCborValueRef* self, QRegularExpression* defaultValue);
extern __declspec(dllexport) QUuid* QCborValueRef_ToUuid1(const QCborValueRef* self, QUuid* defaultValue);
extern __declspec(dllexport) struct miqt_string QCborValueRef_ToCbor1(QCborValueRef* self, int opt);
extern __declspec(dllexport) void QCborValueRef_ToCbor2(QCborValueRef* self, QCborStreamWriter* writer, int opt);
extern __declspec(dllexport) struct miqt_string QCborValueRef_ToDiagnosticNotation1(QCborValueRef* self, int opt);
extern __declspec(dllexport) void QCborValueRef_Delete(QCborValueRef* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
