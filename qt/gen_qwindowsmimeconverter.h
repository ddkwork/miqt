#pragma once
#ifndef MIQT_QT_GEN_QWINDOWSMIMECONVERTER_H
#define MIQT_QT_GEN_QWINDOWSMIMECONVERTER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMetaType QMetaType;
typedef struct QMimeData QMimeData;
typedef struct QVariant QVariant;
typedef struct QWindowsMimeConverter QWindowsMimeConverter;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QWindowsMimeConverter* QWindowsMimeConverter_new();
extern __declspec(dllexport) 
int QWindowsMimeConverter_RegisterMimeType(struct miqt_string mimeType);
extern __declspec(dllexport) 
bool QWindowsMimeConverter_CanConvertFromMime(const QWindowsMimeConverter* self, const tagFORMATETC* formatetc, QMimeData* mimeData);
extern __declspec(dllexport) 
bool QWindowsMimeConverter_ConvertFromMime(const QWindowsMimeConverter* self, const tagFORMATETC* formatetc, QMimeData* mimeData, tagSTGMEDIUM* pmedium);
extern __declspec(dllexport) 
struct miqt_array /* of tagFORMATETC */  QWindowsMimeConverter_FormatsForMime(const QWindowsMimeConverter* self, struct miqt_string mimeType, QMimeData* mimeData);
extern __declspec(dllexport) 
bool QWindowsMimeConverter_CanConvertToMime(const QWindowsMimeConverter* self, struct miqt_string mimeType, IDataObject* pDataObj);
extern __declspec(dllexport) 
QVariant* QWindowsMimeConverter_ConvertToMime(const QWindowsMimeConverter* self, struct miqt_string mimeType, IDataObject* pDataObj, QMetaType* preferredType);
extern __declspec(dllexport) 
struct miqt_string QWindowsMimeConverter_MimeForFormat(const QWindowsMimeConverter* self, const tagFORMATETC* formatetc);
extern __declspec(dllexport) 
void QWindowsMimeConverter_Delete(QWindowsMimeConverter* self, bool isSubclass);

}
