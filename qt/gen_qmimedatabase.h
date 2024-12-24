#pragma once
#ifndef MIQT_QT_GEN_QMIMEDATABASE_H
#define MIQT_QT_GEN_QMIMEDATABASE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QFileInfo QFileInfo;
typedef struct QIODevice QIODevice;
typedef struct QMimeDatabase QMimeDatabase;
typedef struct QMimeType QMimeType;
typedef struct QUrl QUrl;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QMimeDatabase* QMimeDatabase_new();
extern __declspec(dllexport) 
QMimeType* QMimeDatabase_MimeTypeForName(const QMimeDatabase* self, struct miqt_string nameOrAlias);
extern __declspec(dllexport) 
QMimeType* QMimeDatabase_MimeTypeForFile(const QMimeDatabase* self, struct miqt_string fileName);
extern __declspec(dllexport) 
QMimeType* QMimeDatabase_MimeTypeForFileWithFileInfo(const QMimeDatabase* self, QFileInfo* fileInfo);
extern __declspec(dllexport) 
struct miqt_array /* of QMimeType* */  QMimeDatabase_MimeTypesForFileName(const QMimeDatabase* self, struct miqt_string fileName);
extern __declspec(dllexport) 
QMimeType* QMimeDatabase_MimeTypeForData(const QMimeDatabase* self, struct miqt_string data);
extern __declspec(dllexport) 
QMimeType* QMimeDatabase_MimeTypeForDataWithDevice(const QMimeDatabase* self, QIODevice* device);
extern __declspec(dllexport) 
QMimeType* QMimeDatabase_MimeTypeForUrl(const QMimeDatabase* self, QUrl* url);
extern __declspec(dllexport) 
QMimeType* QMimeDatabase_MimeTypeForFileNameAndData(const QMimeDatabase* self, struct miqt_string fileName, QIODevice* device);
extern __declspec(dllexport) 
QMimeType* QMimeDatabase_MimeTypeForFileNameAndData2(const QMimeDatabase* self, struct miqt_string fileName, struct miqt_string data);
extern __declspec(dllexport) 
struct miqt_string QMimeDatabase_SuffixForFileName(const QMimeDatabase* self, struct miqt_string fileName);
extern __declspec(dllexport) 
struct miqt_array /* of QMimeType* */  QMimeDatabase_AllMimeTypes(const QMimeDatabase* self);
extern __declspec(dllexport) 
QMimeType* QMimeDatabase_MimeTypeForFile2(const QMimeDatabase* self, struct miqt_string fileName, MatchMode mode);
extern __declspec(dllexport) 
QMimeType* QMimeDatabase_MimeTypeForFile22(const QMimeDatabase* self, QFileInfo* fileInfo, MatchMode mode);
extern __declspec(dllexport) 
void QMimeDatabase_Delete(QMimeDatabase* self, bool isSubclass);

}
