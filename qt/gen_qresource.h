#pragma once
#ifndef MIQT_QT_GEN_QRESOURCE_H
#define MIQT_QT_GEN_QRESOURCE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QDateTime;
class QLocale;
class QResource;
class _GUID;
class type_info;
#else
typedef struct QDateTime QDateTime;
typedef struct QLocale QLocale;
typedef struct QResource QResource;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QResource* QResource_new();
extern __declspec(dllexport) QResource* QResource_new2(struct miqt_string file);
extern __declspec(dllexport) QResource* QResource_new3(struct miqt_string file, QLocale* locale);
extern __declspec(dllexport) void QResource_SetFileName(QResource* self, struct miqt_string file);
extern __declspec(dllexport) struct miqt_string QResource_FileName(const QResource* self);
extern __declspec(dllexport) struct miqt_string QResource_AbsoluteFilePath(const QResource* self);
extern __declspec(dllexport) void QResource_SetLocale(QResource* self, QLocale* locale);
extern __declspec(dllexport) QLocale* QResource_Locale(const QResource* self);
extern __declspec(dllexport) bool QResource_IsValid(const QResource* self);
extern __declspec(dllexport) Compression QResource_CompressionAlgorithm(const QResource* self);
extern __declspec(dllexport) long long QResource_Size(const QResource* self);
extern __declspec(dllexport) const unsigned char* QResource_Data(const QResource* self);
extern __declspec(dllexport) long long QResource_UncompressedSize(const QResource* self);
extern __declspec(dllexport) struct miqt_string QResource_UncompressedData(const QResource* self);
extern __declspec(dllexport) QDateTime* QResource_LastModified(const QResource* self);
extern __declspec(dllexport) bool QResource_RegisterResource(struct miqt_string rccFilename);
extern __declspec(dllexport) bool QResource_UnregisterResource(struct miqt_string rccFilename);
extern __declspec(dllexport) bool QResource_RegisterResourceWithRccData(const unsigned char* rccData);
extern __declspec(dllexport) bool QResource_UnregisterResourceWithRccData(const unsigned char* rccData);
extern __declspec(dllexport) bool QResource_RegisterResource2(struct miqt_string rccFilename, struct miqt_string resourceRoot);
extern __declspec(dllexport) bool QResource_UnregisterResource2(struct miqt_string rccFilename, struct miqt_string resourceRoot);
extern __declspec(dllexport) bool QResource_RegisterResource22(const unsigned char* rccData, struct miqt_string resourceRoot);
extern __declspec(dllexport) bool QResource_UnregisterResource22(const unsigned char* rccData, struct miqt_string resourceRoot);
extern __declspec(dllexport) void QResource_Delete(QResource* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
