#ifndef GEN_QJSONDOCUMENT_H
#define GEN_QJSONDOCUMENT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QByteArray;
class QJsonArray;
class QJsonDocument;
class QJsonObject;
class QJsonParseError;
class QJsonValue;
class QVariant;
#else
typedef struct QByteArray QByteArray;
typedef struct QJsonArray QJsonArray;
typedef struct QJsonDocument QJsonDocument;
typedef struct QJsonObject QJsonObject;
typedef struct QJsonParseError QJsonParseError;
typedef struct QJsonValue QJsonValue;
typedef struct QVariant QVariant;
#endif

void QJsonParseError_ErrorString(const QJsonParseError* self, char** _out, int* _out_Strlen);
void QJsonParseError_Delete(QJsonParseError* self);

QJsonDocument* QJsonDocument_new();
QJsonDocument* QJsonDocument_new2(QJsonObject* object);
QJsonDocument* QJsonDocument_new3(QJsonArray* array);
QJsonDocument* QJsonDocument_new4(QJsonDocument* other);
void QJsonDocument_OperatorAssign(QJsonDocument* self, QJsonDocument* other);
void QJsonDocument_Swap(QJsonDocument* self, QJsonDocument* other);
QJsonDocument* QJsonDocument_FromRawData(const char* data, int size);
const char* QJsonDocument_RawData(const QJsonDocument* self, int* size);
QJsonDocument* QJsonDocument_FromBinaryData(QByteArray* data);
QByteArray* QJsonDocument_ToBinaryData(const QJsonDocument* self);
QJsonDocument* QJsonDocument_FromVariant(QVariant* variant);
QVariant* QJsonDocument_ToVariant(const QJsonDocument* self);
QJsonDocument* QJsonDocument_FromJson(QByteArray* json);
QByteArray* QJsonDocument_ToJson(const QJsonDocument* self);
QByteArray* QJsonDocument_ToJsonWithFormat(const QJsonDocument* self, uintptr_t format);
bool QJsonDocument_IsEmpty(const QJsonDocument* self);
bool QJsonDocument_IsArray(const QJsonDocument* self);
bool QJsonDocument_IsObject(const QJsonDocument* self);
QJsonObject* QJsonDocument_Object(const QJsonDocument* self);
QJsonArray* QJsonDocument_Array(const QJsonDocument* self);
void QJsonDocument_SetObject(QJsonDocument* self, QJsonObject* object);
void QJsonDocument_SetArray(QJsonDocument* self, QJsonArray* array);
QJsonValue* QJsonDocument_OperatorSubscript(const QJsonDocument* self, const char* key, size_t key_Strlen);
QJsonValue* QJsonDocument_OperatorSubscriptWithInt(const QJsonDocument* self, int i);
bool QJsonDocument_OperatorEqual(const QJsonDocument* self, QJsonDocument* other);
bool QJsonDocument_OperatorNotEqual(const QJsonDocument* self, QJsonDocument* other);
bool QJsonDocument_IsNull(const QJsonDocument* self);
QJsonDocument* QJsonDocument_FromRawData3(const char* data, int size, uintptr_t validation);
QJsonDocument* QJsonDocument_FromBinaryData2(QByteArray* data, uintptr_t validation);
QJsonDocument* QJsonDocument_FromJson2(QByteArray* json, QJsonParseError* error);
void QJsonDocument_Delete(QJsonDocument* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
