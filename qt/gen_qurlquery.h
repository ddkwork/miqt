#pragma once
#ifndef MIQT_QT_GEN_QURLQUERY_H
#define MIQT_QT_GEN_QURLQUERY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChar QChar;
typedef struct QUrl QUrl;
typedef struct QUrlQuery QUrlQuery;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QUrlQuery* QUrlQuery_new();
extern __declspec(dllexport) QUrlQuery* QUrlQuery_new2(QUrl* url);
extern __declspec(dllexport) QUrlQuery* QUrlQuery_new3(struct miqt_string queryString);
extern __declspec(dllexport) QUrlQuery* QUrlQuery_new4(QUrlQuery* other);
extern __declspec(dllexport) void QUrlQuery_OperatorAssign(QUrlQuery* self, QUrlQuery* other);
extern __declspec(dllexport) void QUrlQuery_Swap(QUrlQuery* self, QUrlQuery* other);
extern __declspec(dllexport) bool QUrlQuery_IsEmpty(const QUrlQuery* self);
extern __declspec(dllexport) bool QUrlQuery_IsDetached(const QUrlQuery* self);
extern __declspec(dllexport) void QUrlQuery_Clear(QUrlQuery* self);
extern __declspec(dllexport) struct miqt_string QUrlQuery_Query(const QUrlQuery* self);
extern __declspec(dllexport) void QUrlQuery_SetQuery(QUrlQuery* self, struct miqt_string queryString);
extern __declspec(dllexport) struct miqt_string QUrlQuery_ToString(const QUrlQuery* self);
extern __declspec(dllexport) void QUrlQuery_SetQueryDelimiters(QUrlQuery* self, QChar* valueDelimiter, QChar* pairDelimiter);
extern __declspec(dllexport) QChar* QUrlQuery_QueryValueDelimiter(const QUrlQuery* self);
extern __declspec(dllexport) QChar* QUrlQuery_QueryPairDelimiter(const QUrlQuery* self);
extern __declspec(dllexport) bool QUrlQuery_HasQueryItem(const QUrlQuery* self, struct miqt_string key);
extern __declspec(dllexport) void QUrlQuery_AddQueryItem(QUrlQuery* self, struct miqt_string key, struct miqt_string value);
extern __declspec(dllexport) void QUrlQuery_RemoveQueryItem(QUrlQuery* self, struct miqt_string key);
extern __declspec(dllexport) struct miqt_string QUrlQuery_QueryItemValue(const QUrlQuery* self, struct miqt_string key);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QUrlQuery_AllQueryItemValues(const QUrlQuery* self, struct miqt_string key);
extern __declspec(dllexport) void QUrlQuery_RemoveAllQueryItems(QUrlQuery* self, struct miqt_string key);
extern __declspec(dllexport) DataPtr* QUrlQuery_DataPtr(QUrlQuery* self);
extern __declspec(dllexport) struct miqt_string QUrlQuery_Query1(const QUrlQuery* self, int encoding);
extern __declspec(dllexport) struct miqt_string QUrlQuery_ToString1(const QUrlQuery* self, int encoding);
extern __declspec(dllexport) struct miqt_string QUrlQuery_QueryItemValue2(const QUrlQuery* self, struct miqt_string key, int encoding);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QUrlQuery_AllQueryItemValues2(const QUrlQuery* self, struct miqt_string key, int encoding);
extern __declspec(dllexport) void QUrlQuery_Delete(QUrlQuery* self, bool isSubclass);

} 
