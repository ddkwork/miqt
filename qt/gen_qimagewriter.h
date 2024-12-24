#pragma once
#ifndef MIQT_QT_GEN_QIMAGEWRITER_H
#define MIQT_QT_GEN_QIMAGEWRITER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QIODevice QIODevice;
typedef struct QImage QImage;
typedef struct QImageWriter QImageWriter;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QImageWriter* QImageWriter_new();
extern __declspec(dllexport) QImageWriter* QImageWriter_new2(QIODevice* device, struct miqt_string format);
extern __declspec(dllexport) QImageWriter* QImageWriter_new3(struct miqt_string fileName);
extern __declspec(dllexport) QImageWriter* QImageWriter_new4(struct miqt_string fileName, struct miqt_string format);
extern __declspec(dllexport) struct miqt_string QImageWriter_Tr(const char* sourceText);
extern __declspec(dllexport) void QImageWriter_SetFormat(QImageWriter* self, struct miqt_string format);
extern __declspec(dllexport) struct miqt_string QImageWriter_Format(const QImageWriter* self);
extern __declspec(dllexport) void QImageWriter_SetDevice(QImageWriter* self, QIODevice* device);
extern __declspec(dllexport) QIODevice* QImageWriter_Device(const QImageWriter* self);
extern __declspec(dllexport) void QImageWriter_SetFileName(QImageWriter* self, struct miqt_string fileName);
extern __declspec(dllexport) struct miqt_string QImageWriter_FileName(const QImageWriter* self);
extern __declspec(dllexport) void QImageWriter_SetQuality(QImageWriter* self, int quality);
extern __declspec(dllexport) int QImageWriter_Quality(const QImageWriter* self);
extern __declspec(dllexport) void QImageWriter_SetCompression(QImageWriter* self, int compression);
extern __declspec(dllexport) int QImageWriter_Compression(const QImageWriter* self);
extern __declspec(dllexport) void QImageWriter_SetSubType(QImageWriter* self, struct miqt_string typeVal);
extern __declspec(dllexport) struct miqt_string QImageWriter_SubType(const QImageWriter* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QImageWriter_SupportedSubTypes(const QImageWriter* self);
extern __declspec(dllexport) void QImageWriter_SetOptimizedWrite(QImageWriter* self, bool optimize);
extern __declspec(dllexport) bool QImageWriter_OptimizedWrite(const QImageWriter* self);
extern __declspec(dllexport) void QImageWriter_SetProgressiveScanWrite(QImageWriter* self, bool progressive);
extern __declspec(dllexport) bool QImageWriter_ProgressiveScanWrite(const QImageWriter* self);
extern __declspec(dllexport) int QImageWriter_Transformation(const QImageWriter* self);
extern __declspec(dllexport) void QImageWriter_SetTransformation(QImageWriter* self, int orientation);
extern __declspec(dllexport) void QImageWriter_SetText(QImageWriter* self, struct miqt_string key, struct miqt_string text);
extern __declspec(dllexport) bool QImageWriter_CanWrite(const QImageWriter* self);
extern __declspec(dllexport) bool QImageWriter_Write(QImageWriter* self, QImage* image);
extern __declspec(dllexport) ImageWriterError QImageWriter_Error(const QImageWriter* self);
extern __declspec(dllexport) struct miqt_string QImageWriter_ErrorString(const QImageWriter* self);
extern __declspec(dllexport) bool QImageWriter_SupportsOption(const QImageWriter* self, int option);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QImageWriter_SupportedImageFormats();
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QImageWriter_SupportedMimeTypes();
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QImageWriter_ImageFormatsForMimeType(struct miqt_string mimeType);
extern __declspec(dllexport) struct miqt_string QImageWriter_Tr2(const char* sourceText, const char* disambiguation);
extern __declspec(dllexport) struct miqt_string QImageWriter_Tr3(const char* sourceText, const char* disambiguation, int n);
extern __declspec(dllexport) void QImageWriter_Delete(QImageWriter* self, bool isSubclass);

} 
