#pragma once
#ifndef MIQT_QT_GEN_QMIMEDATA_H
#define MIQT_QT_GEN_QMIMEDATA_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMetaObject QMetaObject;
typedef struct QMetaType QMetaType;
typedef struct QMimeData QMimeData;
typedef struct QObject QObject;
typedef struct QUrl QUrl;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QMimeData* QMimeData_new();
extern __declspec(dllexport) 
void QMimeData_virtbase(QMimeData* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QMimeData_MetaObject(const QMimeData* self);
extern __declspec(dllexport) 
void* QMimeData_Metacast(QMimeData* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QMimeData_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_array /* of QUrl* */  QMimeData_Urls(const QMimeData* self);
extern __declspec(dllexport) 
void QMimeData_SetUrls(QMimeData* self, struct miqt_array /* of QUrl* */  urls);
extern __declspec(dllexport) 
bool QMimeData_HasUrls(const QMimeData* self);
extern __declspec(dllexport) 
struct miqt_string QMimeData_Text(const QMimeData* self);
extern __declspec(dllexport) 
void QMimeData_SetText(QMimeData* self, struct miqt_string text);
extern __declspec(dllexport) 
bool QMimeData_HasText(const QMimeData* self);
extern __declspec(dllexport) 
struct miqt_string QMimeData_Html(const QMimeData* self);
extern __declspec(dllexport) 
void QMimeData_SetHtml(QMimeData* self, struct miqt_string html);
extern __declspec(dllexport) 
bool QMimeData_HasHtml(const QMimeData* self);
extern __declspec(dllexport) 
QVariant* QMimeData_ImageData(const QMimeData* self);
extern __declspec(dllexport) 
void QMimeData_SetImageData(QMimeData* self, QVariant* image);
extern __declspec(dllexport) 
bool QMimeData_HasImage(const QMimeData* self);
extern __declspec(dllexport) 
QVariant* QMimeData_ColorData(const QMimeData* self);
extern __declspec(dllexport) 
void QMimeData_SetColorData(QMimeData* self, QVariant* color);
extern __declspec(dllexport) 
bool QMimeData_HasColor(const QMimeData* self);
extern __declspec(dllexport) 
struct miqt_string QMimeData_Data(const QMimeData* self, struct miqt_string mimetype);
extern __declspec(dllexport) 
void QMimeData_SetData(QMimeData* self, struct miqt_string mimetype, struct miqt_string data);
extern __declspec(dllexport) 
void QMimeData_RemoveFormat(QMimeData* self, struct miqt_string mimetype);
extern __declspec(dllexport) 
bool QMimeData_HasFormat(const QMimeData* self, struct miqt_string mimetype);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QMimeData_Formats(const QMimeData* self);
extern __declspec(dllexport) 
void QMimeData_Clear(QMimeData* self);
extern __declspec(dllexport) 
QVariant* QMimeData_RetrieveData(const QMimeData* self, struct miqt_string mimetype, QMetaType* preferredType);
extern __declspec(dllexport) 
struct miqt_string QMimeData_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QMimeData_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QMimeData_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QMimeData_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QMimeData_override_virtual_Metacast(void* self, intptr_t slot);
void* QMimeData_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QMimeData_Delete(QMimeData* self, bool isSubclass);

}
