#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QMEDIAFORMAT_H
#define MIQT_QT_MULTIMEDIA_GEN_QMEDIAFORMAT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMediaFormat QMediaFormat;
typedef struct QMimeType QMimeType;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QMediaFormat* QMediaFormat_new();
extern __declspec(dllexport) 
QMediaFormat* QMediaFormat_new2(QMediaFormat* other);
extern __declspec(dllexport) 
QMediaFormat* QMediaFormat_new3(FileFormat format);
extern __declspec(dllexport) 
void QMediaFormat_OperatorAssign(QMediaFormat* self, QMediaFormat* other);
extern __declspec(dllexport) 
void QMediaFormat_Swap(QMediaFormat* self, QMediaFormat* other);
extern __declspec(dllexport) 
FileFormat QMediaFormat_FileFormat(const QMediaFormat* self);
extern __declspec(dllexport) 
void QMediaFormat_SetFileFormat(QMediaFormat* self, FileFormat f);
extern __declspec(dllexport) 
void QMediaFormat_SetVideoCodec(QMediaFormat* self, VideoCodec codec);
extern __declspec(dllexport) 
VideoCodec QMediaFormat_VideoCodec(const QMediaFormat* self);
extern __declspec(dllexport) 
void QMediaFormat_SetAudioCodec(QMediaFormat* self, AudioCodec codec);
extern __declspec(dllexport) 
AudioCodec QMediaFormat_AudioCodec(const QMediaFormat* self);
extern __declspec(dllexport) 
bool QMediaFormat_IsSupported(const QMediaFormat* self, ConversionMode mode);
extern __declspec(dllexport) 
QMimeType* QMediaFormat_MimeType(const QMediaFormat* self);
extern __declspec(dllexport) 
struct miqt_array /* of FileFormat */  QMediaFormat_SupportedFileFormats(QMediaFormat* self, ConversionMode m);
extern __declspec(dllexport) 
struct miqt_array /* of VideoCodec */  QMediaFormat_SupportedVideoCodecs(QMediaFormat* self, ConversionMode m);
extern __declspec(dllexport) 
struct miqt_array /* of AudioCodec */  QMediaFormat_SupportedAudioCodecs(QMediaFormat* self, ConversionMode m);
extern __declspec(dllexport) 
struct miqt_string QMediaFormat_FileFormatName(FileFormat fileFormat);
extern __declspec(dllexport) 
struct miqt_string QMediaFormat_AudioCodecName(AudioCodec codec);
extern __declspec(dllexport) 
struct miqt_string QMediaFormat_VideoCodecName(VideoCodec codec);
extern __declspec(dllexport) 
struct miqt_string QMediaFormat_FileFormatDescription(int fileFormat);
extern __declspec(dllexport) 
struct miqt_string QMediaFormat_AudioCodecDescription(int codec);
extern __declspec(dllexport) 
struct miqt_string QMediaFormat_VideoCodecDescription(int codec);
extern __declspec(dllexport) 
bool QMediaFormat_OperatorEqual(const QMediaFormat* self, QMediaFormat* other);
extern __declspec(dllexport) 
bool QMediaFormat_OperatorNotEqual(const QMediaFormat* self, QMediaFormat* other);
extern __declspec(dllexport) 
void QMediaFormat_ResolveForEncoding(QMediaFormat* self, ResolveFlags flags);
extern __declspec(dllexport) 
void QMediaFormat_Delete(QMediaFormat* self, bool isSubclass);

}
