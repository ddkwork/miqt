#pragma once
#ifndef MIQT_QT_GEN_QJSONDOCUMENT_H
#define MIQT_QT_GEN_QJSONDOCUMENT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QJsonArray QJsonArray;
typedef struct QJsonDocument QJsonDocument;
typedef struct QJsonObject QJsonObject;
typedef struct QJsonParseError QJsonParseError;
typedef struct QJsonValue QJsonValue;
typedef struct QVariant QVariant;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QJsonDocument* QJsonDocument_new();
extern __declspec(dllexport) QJsonDocument* QJsonDocument_new2(QJsonObject* object);
extern __declspec(dllexport) QJsonDocument* QJsonDocument_new3(QJsonArray* array);
extern __declspec(dllexport) QJsonDocument* QJsonDocument_new4(QJsonDocument* other);
extern __declspec(dllexport) void QJsonDocument_OperatorAssign(QJsonDocument* self, QJsonDocument* other);
extern __declspec(dllexport) void QJsonDocument_Swap(QJsonDocument* self, QJsonDocument* other);
extern __declspec(dllexport) QJsonDocument* QJsonDocument_FromVariant(QVariant* variant);
extern __declspec(dllexport) QVariant* QJsonDocument_ToVariant(const QJsonDocument* self);
extern __declspec(dllexport) QJsonDocument* QJsonDocument_FromJson(struct miqt_string json);
extern __declspec(dllexport) struct miqt_string QJsonDocument_ToJson(const QJsonDocument* self);
extern __declspec(dllexport) bool QJsonDocument_IsEmpty(const QJsonDocument* self);
extern __declspec(dllexport) bool QJsonDocument_IsArray(const QJsonDocument* self);
extern __declspec(dllexport) bool QJsonDocument_IsObject(const QJsonDocument* self);
extern __declspec(dllexport) QJsonObject* QJsonDocument_Object(const QJsonDocument* self);
extern __declspec(dllexport) QJsonArray* QJsonDocument_Array(const QJsonDocument* self);
extern __declspec(dllexport) void QJsonDocument_SetObject(QJsonDocument* self, QJsonObject* object);
extern __declspec(dllexport) void QJsonDocument_SetArray(QJsonDocument* self, QJsonArray* array);
extern __declspec(dllexport) QJsonValue* QJsonDocument_OperatorSubscript(const QJsonDocument* self, struct miqt_string key);
extern __declspec(dllexport) QJsonValue* QJsonDocument_OperatorSubscriptWithQsizetype(const QJsonDocument* self, ptrdiff_t i);
extern __declspec(dllexport) bool QJsonDocument_IsNull(const QJsonDocument* self);
extern __declspec(dllexport) QJsonDocument* QJsonDocument_FromJson2(struct miqt_string json, QJsonParseError* error);
extern __declspec(dllexport) struct miqt_string QJsonDocument_ToJson1(const QJsonDocument* self, JsonFormat format);
extern __declspec(dllexport) void QJsonDocument_Delete(QJsonDocument* self, bool isSubclass);

} 
