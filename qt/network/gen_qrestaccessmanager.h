#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QRESTACCESSMANAGER_H
#define MIQT_QT_NETWORK_GEN_QRESTACCESSMANAGER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QHttpMultiPart QHttpMultiPart;
typedef struct QIODevice QIODevice;
typedef struct QJsonDocument QJsonDocument;
typedef struct QMetaObject QMetaObject;
typedef struct QNetworkAccessManager QNetworkAccessManager;
typedef struct QNetworkReply QNetworkReply;
typedef struct QNetworkRequest QNetworkRequest;
typedef struct QObject QObject;
typedef struct QRestAccessManager QRestAccessManager;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QRestAccessManager* QRestAccessManager_new(QNetworkAccessManager* manager);
extern __declspec(dllexport) 
QRestAccessManager* QRestAccessManager_new2(QNetworkAccessManager* manager, QObject* parent);
extern __declspec(dllexport) 
void QRestAccessManager_virtbase(QRestAccessManager* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QRestAccessManager_MetaObject(const QRestAccessManager* self);
extern __declspec(dllexport) 
void* QRestAccessManager_Metacast(QRestAccessManager* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QRestAccessManager_Tr(const char* s);
extern __declspec(dllexport) 
QNetworkAccessManager* QRestAccessManager_NetworkAccessManager(const QRestAccessManager* self);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_DeleteResource(QRestAccessManager* self, QNetworkRequest* request);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Head(QRestAccessManager* self, QNetworkRequest* request);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Get(QRestAccessManager* self, QNetworkRequest* request);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Get2(QRestAccessManager* self, QNetworkRequest* request, struct miqt_string data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Get3(QRestAccessManager* self, QNetworkRequest* request, QJsonDocument* data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Get4(QRestAccessManager* self, QNetworkRequest* request, QIODevice* data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Post(QRestAccessManager* self, QNetworkRequest* request, QJsonDocument* data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Post2(QRestAccessManager* self, QNetworkRequest* request, struct miqt_map /* of struct miqt_string to QVariant* */  data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Post3(QRestAccessManager* self, QNetworkRequest* request, struct miqt_string data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Post4(QRestAccessManager* self, QNetworkRequest* request, QHttpMultiPart* data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Post5(QRestAccessManager* self, QNetworkRequest* request, QIODevice* data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Put(QRestAccessManager* self, QNetworkRequest* request, QJsonDocument* data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Put2(QRestAccessManager* self, QNetworkRequest* request, struct miqt_map /* of struct miqt_string to QVariant* */  data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Put3(QRestAccessManager* self, QNetworkRequest* request, struct miqt_string data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Put4(QRestAccessManager* self, QNetworkRequest* request, QHttpMultiPart* data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Put5(QRestAccessManager* self, QNetworkRequest* request, QIODevice* data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Patch(QRestAccessManager* self, QNetworkRequest* request, QJsonDocument* data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Patch2(QRestAccessManager* self, QNetworkRequest* request, struct miqt_map /* of struct miqt_string to QVariant* */  data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Patch3(QRestAccessManager* self, QNetworkRequest* request, struct miqt_string data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_Patch4(QRestAccessManager* self, QNetworkRequest* request, QIODevice* data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_SendCustomRequest(QRestAccessManager* self, QNetworkRequest* request, struct miqt_string method, struct miqt_string data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_SendCustomRequest2(QRestAccessManager* self, QNetworkRequest* request, struct miqt_string method, QIODevice* data);
extern __declspec(dllexport) 
QNetworkReply* QRestAccessManager_SendCustomRequest3(QRestAccessManager* self, QNetworkRequest* request, struct miqt_string method, QHttpMultiPart* data);
extern __declspec(dllexport) 
struct miqt_string QRestAccessManager_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QRestAccessManager_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QRestAccessManager_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QRestAccessManager_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QRestAccessManager_override_virtual_Metacast(void* self, intptr_t slot);
void* QRestAccessManager_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QRestAccessManager_Delete(QRestAccessManager* self, bool isSubclass);

}
