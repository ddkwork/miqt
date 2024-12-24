#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QIMAGECAPTURE_H
#define MIQT_QT_MULTIMEDIA_GEN_QIMAGECAPTURE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QImage QImage;
typedef struct QImageCapture QImageCapture;
typedef struct QMediaCaptureSession QMediaCaptureSession;
typedef struct QMediaMetaData QMediaMetaData;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSize QSize;
typedef struct QVideoFrame QVideoFrame;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QImageCapture* QImageCapture_new();
extern __declspec(dllexport) 
QImageCapture* QImageCapture_new2(QObject* parent);
extern __declspec(dllexport) 
void QImageCapture_virtbase(QImageCapture* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QImageCapture_MetaObject(const QImageCapture* self);
extern __declspec(dllexport) 
void* QImageCapture_Metacast(QImageCapture* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QImageCapture_Tr(const char* s);
extern __declspec(dllexport) 
bool QImageCapture_IsAvailable(const QImageCapture* self);
extern __declspec(dllexport) 
QMediaCaptureSession* QImageCapture_CaptureSession(const QImageCapture* self);
extern __declspec(dllexport) 
Error QImageCapture_Error(const QImageCapture* self);
extern __declspec(dllexport) 
struct miqt_string QImageCapture_ErrorString(const QImageCapture* self);
extern __declspec(dllexport) 
bool QImageCapture_IsReadyForCapture(const QImageCapture* self);
extern __declspec(dllexport) 
FileFormat QImageCapture_FileFormat(const QImageCapture* self);
extern __declspec(dllexport) 
void QImageCapture_SetFileFormat(QImageCapture* self, FileFormat format);
extern __declspec(dllexport) 
struct miqt_array /* of FileFormat */  QImageCapture_SupportedFormats();
extern __declspec(dllexport) 
struct miqt_string QImageCapture_FileFormatName(FileFormat c);
extern __declspec(dllexport) 
struct miqt_string QImageCapture_FileFormatDescription(FileFormat c);
extern __declspec(dllexport) 
QSize* QImageCapture_Resolution(const QImageCapture* self);
extern __declspec(dllexport) 
void QImageCapture_SetResolution(QImageCapture* self, QSize* resolution);
extern __declspec(dllexport) 
void QImageCapture_SetResolution2(QImageCapture* self, int width, int height);
extern __declspec(dllexport) 
Quality QImageCapture_Quality(const QImageCapture* self);
extern __declspec(dllexport) 
void QImageCapture_SetQuality(QImageCapture* self, Quality quality);
extern __declspec(dllexport) 
QMediaMetaData* QImageCapture_MetaData(const QImageCapture* self);
extern __declspec(dllexport) 
void QImageCapture_SetMetaData(QImageCapture* self, QMediaMetaData* metaData);
extern __declspec(dllexport) 
void QImageCapture_AddMetaData(QImageCapture* self, QMediaMetaData* metaData);
extern __declspec(dllexport) 
int QImageCapture_CaptureToFile(QImageCapture* self);
extern __declspec(dllexport) 
int QImageCapture_Capture(QImageCapture* self);
extern __declspec(dllexport) 
void QImageCapture_ErrorChanged(QImageCapture* self);
void QImageCapture_connect_ErrorChanged(QImageCapture* self, intptr_t slot);
extern __declspec(dllexport) 
void QImageCapture_ErrorOccurred(QImageCapture* self, int id, int error, struct miqt_string errorString);
void QImageCapture_connect_ErrorOccurred(QImageCapture* self, intptr_t slot);
extern __declspec(dllexport) 
void QImageCapture_ReadyForCaptureChanged(QImageCapture* self, bool ready);
void QImageCapture_connect_ReadyForCaptureChanged(QImageCapture* self, intptr_t slot);
extern __declspec(dllexport) 
void QImageCapture_MetaDataChanged(QImageCapture* self);
void QImageCapture_connect_MetaDataChanged(QImageCapture* self, intptr_t slot);
extern __declspec(dllexport) 
void QImageCapture_FileFormatChanged(QImageCapture* self);
void QImageCapture_connect_FileFormatChanged(QImageCapture* self, intptr_t slot);
extern __declspec(dllexport) 
void QImageCapture_QualityChanged(QImageCapture* self);
void QImageCapture_connect_QualityChanged(QImageCapture* self, intptr_t slot);
extern __declspec(dllexport) 
void QImageCapture_ResolutionChanged(QImageCapture* self);
void QImageCapture_connect_ResolutionChanged(QImageCapture* self, intptr_t slot);
extern __declspec(dllexport) 
void QImageCapture_ImageExposed(QImageCapture* self, int id);
void QImageCapture_connect_ImageExposed(QImageCapture* self, intptr_t slot);
extern __declspec(dllexport) 
void QImageCapture_ImageCaptured(QImageCapture* self, int id, QImage* preview);
void QImageCapture_connect_ImageCaptured(QImageCapture* self, intptr_t slot);
extern __declspec(dllexport) 
void QImageCapture_ImageMetadataAvailable(QImageCapture* self, int id, QMediaMetaData* metaData);
void QImageCapture_connect_ImageMetadataAvailable(QImageCapture* self, intptr_t slot);
extern __declspec(dllexport) 
void QImageCapture_ImageAvailable(QImageCapture* self, int id, QVideoFrame* frame);
void QImageCapture_connect_ImageAvailable(QImageCapture* self, intptr_t slot);
extern __declspec(dllexport) 
void QImageCapture_ImageSaved(QImageCapture* self, int id, struct miqt_string fileName);
void QImageCapture_connect_ImageSaved(QImageCapture* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QImageCapture_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QImageCapture_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
int QImageCapture_CaptureToFile1(QImageCapture* self, struct miqt_string location);
extern __declspec(dllexport) 
void QImageCapture_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QImageCapture_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QImageCapture_override_virtual_Metacast(void* self, intptr_t slot);
void* QImageCapture_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QImageCapture_Delete(QImageCapture* self, bool isSubclass);

}
