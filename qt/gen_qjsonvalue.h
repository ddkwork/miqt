#pragma once
#ifndef MIQT_QT_GEN_QJSONVALUE_H
#define MIQT_QT_GEN_QJSONVALUE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QByteArrayView QByteArrayView;
typedef struct QJsonArray QJsonArray;
typedef struct QJsonObject QJsonObject;
typedef struct QJsonParseError QJsonParseError;
typedef struct QJsonValue QJsonValue;
typedef struct QJsonValueConstRef QJsonValueConstRef;
typedef struct QJsonValueRef QJsonValueRef;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QJsonValue* QJsonValue_new();
extern __declspec(dllexport) 
QJsonValue* QJsonValue_new2(bool b);
extern __declspec(dllexport) 
QJsonValue* QJsonValue_new3(double n);
extern __declspec(dllexport) 
QJsonValue* QJsonValue_new4(int n);
extern __declspec(dllexport) 
QJsonValue* QJsonValue_new5(long long v);
extern __declspec(dllexport) 
QJsonValue* QJsonValue_new6(struct miqt_string s);
extern __declspec(dllexport) 
QJsonValue* QJsonValue_new7(const char* s);
extern __declspec(dllexport) 
QJsonValue* QJsonValue_new8(QJsonArray* a);
extern __declspec(dllexport) 
QJsonValue* QJsonValue_new9(QJsonObject* o);
extern __declspec(dllexport) 
QJsonValue* QJsonValue_new10(QJsonValue* other);
extern __declspec(dllexport) 
QJsonValue* QJsonValue_new11(Type param1);
extern __declspec(dllexport) 
void QJsonValue_OperatorAssign(QJsonValue* self, QJsonValue* other);
extern __declspec(dllexport) 
void QJsonValue_Swap(QJsonValue* self, QJsonValue* other);
extern __declspec(dllexport) 
QJsonValue* QJsonValue_FromVariant(QVariant* variant);
extern __declspec(dllexport) 
QVariant* QJsonValue_ToVariant(const QJsonValue* self);
extern __declspec(dllexport) 
QJsonValue* QJsonValue_FromJson(QByteArrayView* json);
extern __declspec(dllexport) 
struct miqt_string QJsonValue_ToJson(const QJsonValue* self);
extern __declspec(dllexport) 
Type QJsonValue_Type(const QJsonValue* self);
extern __declspec(dllexport) 
bool QJsonValue_IsNull(const QJsonValue* self);
extern __declspec(dllexport) 
bool QJsonValue_IsBool(const QJsonValue* self);
extern __declspec(dllexport) 
bool QJsonValue_IsDouble(const QJsonValue* self);
extern __declspec(dllexport) 
bool QJsonValue_IsString(const QJsonValue* self);
extern __declspec(dllexport) 
bool QJsonValue_IsArray(const QJsonValue* self);
extern __declspec(dllexport) 
bool QJsonValue_IsObject(const QJsonValue* self);
extern __declspec(dllexport) 
bool QJsonValue_IsUndefined(const QJsonValue* self);
extern __declspec(dllexport) 
bool QJsonValue_ToBool(const QJsonValue* self);
extern __declspec(dllexport) 
int QJsonValue_ToInt(const QJsonValue* self);
extern __declspec(dllexport) 
long long QJsonValue_ToInteger(const QJsonValue* self);
extern __declspec(dllexport) 
double QJsonValue_ToDouble(const QJsonValue* self);
extern __declspec(dllexport) 
struct miqt_string QJsonValue_ToString(const QJsonValue* self);
extern __declspec(dllexport) 
struct miqt_string QJsonValue_ToStringWithDefaultValue(const QJsonValue* self, struct miqt_string defaultValue);
extern __declspec(dllexport) 
QJsonArray* QJsonValue_ToArray(const QJsonValue* self);
extern __declspec(dllexport) 
QJsonArray* QJsonValue_ToArrayWithDefaultValue(const QJsonValue* self, QJsonArray* defaultValue);
extern __declspec(dllexport) 
QJsonObject* QJsonValue_ToObject(const QJsonValue* self);
extern __declspec(dllexport) 
QJsonObject* QJsonValue_ToObjectWithDefaultValue(const QJsonValue* self, QJsonObject* defaultValue);
extern __declspec(dllexport) 
QJsonValue* QJsonValue_OperatorSubscript(const QJsonValue* self, struct miqt_string key);
extern __declspec(dllexport) 
QJsonValue* QJsonValue_OperatorSubscriptWithQsizetype(const QJsonValue* self, ptrdiff_t i);
extern __declspec(dllexport) 
QJsonValue* QJsonValue_FromJson2(QByteArrayView* json, QJsonParseError* error);
extern __declspec(dllexport) 
struct miqt_string QJsonValue_ToJson1(const QJsonValue* self, JsonFormat format);
extern __declspec(dllexport) 
bool QJsonValue_ToBool1(const QJsonValue* self, bool defaultValue);
extern __declspec(dllexport) 
int QJsonValue_ToInt1(const QJsonValue* self, int defaultValue);
extern __declspec(dllexport) 
long long QJsonValue_ToInteger1(const QJsonValue* self, long long defaultValue);
extern __declspec(dllexport) 
double QJsonValue_ToDouble1(const QJsonValue* self, double defaultValue);
extern __declspec(dllexport) 
void QJsonValue_Delete(QJsonValue* self, bool isSubclass);

extern __declspec(dllexport) 
QJsonValueConstRef* QJsonValueConstRef_new(QJsonValueConstRef* param1);
extern __declspec(dllexport) 
QVariant* QJsonValueConstRef_ToVariant(const QJsonValueConstRef* self);
extern __declspec(dllexport) 
int QJsonValueConstRef_Type(const QJsonValueConstRef* self);
extern __declspec(dllexport) 
bool QJsonValueConstRef_IsNull(const QJsonValueConstRef* self);
extern __declspec(dllexport) 
bool QJsonValueConstRef_IsBool(const QJsonValueConstRef* self);
extern __declspec(dllexport) 
bool QJsonValueConstRef_IsDouble(const QJsonValueConstRef* self);
extern __declspec(dllexport) 
bool QJsonValueConstRef_IsString(const QJsonValueConstRef* self);
extern __declspec(dllexport) 
bool QJsonValueConstRef_IsArray(const QJsonValueConstRef* self);
extern __declspec(dllexport) 
bool QJsonValueConstRef_IsObject(const QJsonValueConstRef* self);
extern __declspec(dllexport) 
bool QJsonValueConstRef_IsUndefined(const QJsonValueConstRef* self);
extern __declspec(dllexport) 
bool QJsonValueConstRef_ToBool(const QJsonValueConstRef* self);
extern __declspec(dllexport) 
int QJsonValueConstRef_ToInt(const QJsonValueConstRef* self);
extern __declspec(dllexport) 
long long QJsonValueConstRef_ToInteger(const QJsonValueConstRef* self);
extern __declspec(dllexport) 
double QJsonValueConstRef_ToDouble(const QJsonValueConstRef* self);
extern __declspec(dllexport) 
struct miqt_string QJsonValueConstRef_ToString(const QJsonValueConstRef* self);
extern __declspec(dllexport) 
QJsonArray* QJsonValueConstRef_ToArray(const QJsonValueConstRef* self);
extern __declspec(dllexport) 
QJsonObject* QJsonValueConstRef_ToObject(const QJsonValueConstRef* self);
extern __declspec(dllexport) 
QJsonValue* QJsonValueConstRef_OperatorSubscriptWithQsizetype(const QJsonValueConstRef* self, ptrdiff_t i);
extern __declspec(dllexport) 
bool QJsonValueConstRef_ToBool1(const QJsonValueConstRef* self, bool defaultValue);
extern __declspec(dllexport) 
int QJsonValueConstRef_ToInt1(const QJsonValueConstRef* self, int defaultValue);
extern __declspec(dllexport) 
long long QJsonValueConstRef_ToInteger1(const QJsonValueConstRef* self, long long defaultValue);
extern __declspec(dllexport) 
double QJsonValueConstRef_ToDouble1(const QJsonValueConstRef* self, double defaultValue);
extern __declspec(dllexport) 
struct miqt_string QJsonValueConstRef_ToString1(const QJsonValueConstRef* self, struct miqt_string defaultValue);
extern __declspec(dllexport) 
void QJsonValueConstRef_Delete(QJsonValueConstRef* self, bool isSubclass);

extern __declspec(dllexport) 
QJsonValueRef* QJsonValueRef_new(QJsonValueRef* param1);
extern __declspec(dllexport) 
QJsonValueRef* QJsonValueRef_new2(QJsonArray* array, ptrdiff_t idx);
extern __declspec(dllexport) 
QJsonValueRef* QJsonValueRef_new3(QJsonObject* object, ptrdiff_t idx);
extern __declspec(dllexport) 
void QJsonValueRef_virtbase(QJsonValueRef* src
, QJsonValueConstRef** outptr_QJsonValueConstRef
);
extern __declspec(dllexport) 
void QJsonValueRef_OperatorAssign(QJsonValueRef* self, QJsonValue* val);
extern __declspec(dllexport) 
void QJsonValueRef_OperatorAssignWithVal(QJsonValueRef* self, QJsonValueRef* val);
extern __declspec(dllexport) 
QVariant* QJsonValueRef_ToVariant(const QJsonValueRef* self);
extern __declspec(dllexport) 
int QJsonValueRef_Type(const QJsonValueRef* self);
extern __declspec(dllexport) 
bool QJsonValueRef_IsNull(const QJsonValueRef* self);
extern __declspec(dllexport) 
bool QJsonValueRef_IsBool(const QJsonValueRef* self);
extern __declspec(dllexport) 
bool QJsonValueRef_IsDouble(const QJsonValueRef* self);
extern __declspec(dllexport) 
bool QJsonValueRef_IsString(const QJsonValueRef* self);
extern __declspec(dllexport) 
bool QJsonValueRef_IsArray(const QJsonValueRef* self);
extern __declspec(dllexport) 
bool QJsonValueRef_IsObject(const QJsonValueRef* self);
extern __declspec(dllexport) 
bool QJsonValueRef_IsUndefined(const QJsonValueRef* self);
extern __declspec(dllexport) 
bool QJsonValueRef_ToBool(const QJsonValueRef* self);
extern __declspec(dllexport) 
int QJsonValueRef_ToInt(const QJsonValueRef* self);
extern __declspec(dllexport) 
long long QJsonValueRef_ToInteger(const QJsonValueRef* self);
extern __declspec(dllexport) 
double QJsonValueRef_ToDouble(const QJsonValueRef* self);
extern __declspec(dllexport) 
struct miqt_string QJsonValueRef_ToString(const QJsonValueRef* self);
extern __declspec(dllexport) 
QJsonArray* QJsonValueRef_ToArray(const QJsonValueRef* self);
extern __declspec(dllexport) 
QJsonObject* QJsonValueRef_ToObject(const QJsonValueRef* self);
extern __declspec(dllexport) 
QJsonValue* QJsonValueRef_OperatorSubscriptWithQsizetype(const QJsonValueRef* self, ptrdiff_t i);
extern __declspec(dllexport) 
bool QJsonValueRef_ToBool1(const QJsonValueRef* self, bool defaultValue);
extern __declspec(dllexport) 
int QJsonValueRef_ToInt1(const QJsonValueRef* self, int defaultValue);
extern __declspec(dllexport) 
long long QJsonValueRef_ToInteger1(const QJsonValueRef* self, long long defaultValue);
extern __declspec(dllexport) 
double QJsonValueRef_ToDouble1(const QJsonValueRef* self, double defaultValue);
extern __declspec(dllexport) 
struct miqt_string QJsonValueRef_ToString1(const QJsonValueRef* self, struct miqt_string defaultValue);
extern __declspec(dllexport) 
void QJsonValueRef_Delete(QJsonValueRef* self, bool isSubclass);

}
