#pragma once
#ifndef MIQT_QT_GEN_QIMAGEREADER_H
#define MIQT_QT_GEN_QIMAGEREADER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QColor QColor;
typedef struct QIODevice QIODevice;
typedef struct QImage QImage;
typedef struct QImageReader QImageReader;
typedef struct QRect QRect;
typedef struct QSize QSize;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QImageReader* QImageReader_new();
extern __declspec(dllexport) QImageReader* QImageReader_new2(QIODevice* device);
extern __declspec(dllexport) QImageReader* QImageReader_new3(struct miqt_string fileName);
extern __declspec(dllexport) QImageReader* QImageReader_new4(QIODevice* device, struct miqt_string format);
extern __declspec(dllexport) QImageReader* QImageReader_new5(struct miqt_string fileName, struct miqt_string format);
extern __declspec(dllexport) struct miqt_string QImageReader_Tr(const char* sourceText);
extern __declspec(dllexport) void QImageReader_SetFormat(QImageReader* self, struct miqt_string format);
extern __declspec(dllexport) struct miqt_string QImageReader_Format(const QImageReader* self);
extern __declspec(dllexport) void QImageReader_SetAutoDetectImageFormat(QImageReader* self, bool enabled);
extern __declspec(dllexport) bool QImageReader_AutoDetectImageFormat(const QImageReader* self);
extern __declspec(dllexport) void QImageReader_SetDecideFormatFromContent(QImageReader* self, bool ignored);
extern __declspec(dllexport) bool QImageReader_DecideFormatFromContent(const QImageReader* self);
extern __declspec(dllexport) void QImageReader_SetDevice(QImageReader* self, QIODevice* device);
extern __declspec(dllexport) QIODevice* QImageReader_Device(const QImageReader* self);
extern __declspec(dllexport) void QImageReader_SetFileName(QImageReader* self, struct miqt_string fileName);
extern __declspec(dllexport) struct miqt_string QImageReader_FileName(const QImageReader* self);
extern __declspec(dllexport) QSize* QImageReader_Size(const QImageReader* self);
extern __declspec(dllexport) int QImageReader_ImageFormat(const QImageReader* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QImageReader_TextKeys(const QImageReader* self);
extern __declspec(dllexport) struct miqt_string QImageReader_Text(const QImageReader* self, struct miqt_string key);
extern __declspec(dllexport) void QImageReader_SetClipRect(QImageReader* self, QRect* rect);
extern __declspec(dllexport) QRect* QImageReader_ClipRect(const QImageReader* self);
extern __declspec(dllexport) void QImageReader_SetScaledSize(QImageReader* self, QSize* size);
extern __declspec(dllexport) QSize* QImageReader_ScaledSize(const QImageReader* self);
extern __declspec(dllexport) void QImageReader_SetQuality(QImageReader* self, int quality);
extern __declspec(dllexport) int QImageReader_Quality(const QImageReader* self);
extern __declspec(dllexport) void QImageReader_SetScaledClipRect(QImageReader* self, QRect* rect);
extern __declspec(dllexport) QRect* QImageReader_ScaledClipRect(const QImageReader* self);
extern __declspec(dllexport) void QImageReader_SetBackgroundColor(QImageReader* self, QColor* color);
extern __declspec(dllexport) QColor* QImageReader_BackgroundColor(const QImageReader* self);
extern __declspec(dllexport) bool QImageReader_SupportsAnimation(const QImageReader* self);
extern __declspec(dllexport) int QImageReader_Transformation(const QImageReader* self);
extern __declspec(dllexport) void QImageReader_SetAutoTransform(QImageReader* self, bool enabled);
extern __declspec(dllexport) bool QImageReader_AutoTransform(const QImageReader* self);
extern __declspec(dllexport) struct miqt_string QImageReader_SubType(const QImageReader* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QImageReader_SupportedSubTypes(const QImageReader* self);
extern __declspec(dllexport) bool QImageReader_CanRead(const QImageReader* self);
extern __declspec(dllexport) QImage* QImageReader_Read(QImageReader* self);
extern __declspec(dllexport) bool QImageReader_ReadWithImage(QImageReader* self, QImage* image);
extern __declspec(dllexport) bool QImageReader_JumpToNextImage(QImageReader* self);
extern __declspec(dllexport) bool QImageReader_JumpToImage(QImageReader* self, int imageNumber);
extern __declspec(dllexport) int QImageReader_LoopCount(const QImageReader* self);
extern __declspec(dllexport) int QImageReader_ImageCount(const QImageReader* self);
extern __declspec(dllexport) int QImageReader_NextImageDelay(const QImageReader* self);
extern __declspec(dllexport) int QImageReader_CurrentImageNumber(const QImageReader* self);
extern __declspec(dllexport) QRect* QImageReader_CurrentImageRect(const QImageReader* self);
extern __declspec(dllexport) ImageReaderError QImageReader_Error(const QImageReader* self);
extern __declspec(dllexport) struct miqt_string QImageReader_ErrorString(const QImageReader* self);
extern __declspec(dllexport) bool QImageReader_SupportsOption(const QImageReader* self, int option);
extern __declspec(dllexport) struct miqt_string QImageReader_ImageFormatWithFileName(struct miqt_string fileName);
extern __declspec(dllexport) struct miqt_string QImageReader_ImageFormatWithDevice(QIODevice* device);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QImageReader_SupportedImageFormats();
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QImageReader_SupportedMimeTypes();
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QImageReader_ImageFormatsForMimeType(struct miqt_string mimeType);
extern __declspec(dllexport) int QImageReader_AllocationLimit();
extern __declspec(dllexport) void QImageReader_SetAllocationLimit(int mbLimit);
extern __declspec(dllexport) struct miqt_string QImageReader_Tr2(const char* sourceText, const char* disambiguation);
extern __declspec(dllexport) struct miqt_string QImageReader_Tr3(const char* sourceText, const char* disambiguation, int n);
extern __declspec(dllexport) void QImageReader_Delete(QImageReader* self, bool isSubclass);

} 
