#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QVIDEOFRAMEFORMAT_H
#define MIQT_QT_MULTIMEDIA_GEN_QVIDEOFRAMEFORMAT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QVideoFrameFormat QVideoFrameFormat;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QVideoFrameFormat* QVideoFrameFormat_new();
extern __declspec(dllexport) 
QVideoFrameFormat* QVideoFrameFormat_new2(QSize* size, PixelFormat pixelFormat);
extern __declspec(dllexport) 
QVideoFrameFormat* QVideoFrameFormat_new3(QVideoFrameFormat* format);
extern __declspec(dllexport) 
void QVideoFrameFormat_Swap(QVideoFrameFormat* self, QVideoFrameFormat* other);
extern __declspec(dllexport) 
void QVideoFrameFormat_Detach(QVideoFrameFormat* self);
extern __declspec(dllexport) 
void QVideoFrameFormat_OperatorAssign(QVideoFrameFormat* self, QVideoFrameFormat* format);
extern __declspec(dllexport) 
bool QVideoFrameFormat_OperatorEqual(const QVideoFrameFormat* self, QVideoFrameFormat* format);
extern __declspec(dllexport) 
bool QVideoFrameFormat_OperatorNotEqual(const QVideoFrameFormat* self, QVideoFrameFormat* format);
extern __declspec(dllexport) 
bool QVideoFrameFormat_IsValid(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
int QVideoFrameFormat_PixelFormat(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
QSize* QVideoFrameFormat_FrameSize(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
void QVideoFrameFormat_SetFrameSize(QVideoFrameFormat* self, QSize* size);
extern __declspec(dllexport) 
void QVideoFrameFormat_SetFrameSize2(QVideoFrameFormat* self, int width, int height);
extern __declspec(dllexport) 
int QVideoFrameFormat_FrameWidth(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
int QVideoFrameFormat_FrameHeight(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
int QVideoFrameFormat_PlaneCount(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
QRect* QVideoFrameFormat_Viewport(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
void QVideoFrameFormat_SetViewport(QVideoFrameFormat* self, QRect* viewport);
extern __declspec(dllexport) 
Direction QVideoFrameFormat_ScanLineDirection(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
void QVideoFrameFormat_SetScanLineDirection(QVideoFrameFormat* self, Direction direction);
extern __declspec(dllexport) 
double QVideoFrameFormat_FrameRate(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
void QVideoFrameFormat_SetFrameRate(QVideoFrameFormat* self, double rate);
extern __declspec(dllexport) 
double QVideoFrameFormat_StreamFrameRate(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
void QVideoFrameFormat_SetStreamFrameRate(QVideoFrameFormat* self, double rate);
extern __declspec(dllexport) 
YCbCrColorSpace QVideoFrameFormat_YCbCrColorSpace(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
void QVideoFrameFormat_SetYCbCrColorSpace(QVideoFrameFormat* self, YCbCrColorSpace colorSpace);
extern __declspec(dllexport) 
ColorSpace QVideoFrameFormat_ColorSpace(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
void QVideoFrameFormat_SetColorSpace(QVideoFrameFormat* self, ColorSpace colorSpace);
extern __declspec(dllexport) 
ColorTransfer QVideoFrameFormat_ColorTransfer(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
void QVideoFrameFormat_SetColorTransfer(QVideoFrameFormat* self, ColorTransfer colorTransfer);
extern __declspec(dllexport) 
ColorRange QVideoFrameFormat_ColorRange(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
void QVideoFrameFormat_SetColorRange(QVideoFrameFormat* self, ColorRange rangeVal);
extern __declspec(dllexport) 
bool QVideoFrameFormat_IsMirrored(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
void QVideoFrameFormat_SetMirrored(QVideoFrameFormat* self, bool mirrored);
extern __declspec(dllexport) 
int QVideoFrameFormat_Rotation(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
void QVideoFrameFormat_SetRotation(QVideoFrameFormat* self, int rotation);
extern __declspec(dllexport) 
struct miqt_string QVideoFrameFormat_VertexShaderFileName(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
struct miqt_string QVideoFrameFormat_FragmentShaderFileName(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
float QVideoFrameFormat_MaxLuminance(const QVideoFrameFormat* self);
extern __declspec(dllexport) 
void QVideoFrameFormat_SetMaxLuminance(QVideoFrameFormat* self, float lum);
extern __declspec(dllexport) 
PixelFormat QVideoFrameFormat_PixelFormatFromImageFormat(int format);
extern __declspec(dllexport) 
int QVideoFrameFormat_ImageFormatFromPixelFormat(PixelFormat format);
extern __declspec(dllexport) 
struct miqt_string QVideoFrameFormat_PixelFormatToString(int pixelFormat);
extern __declspec(dllexport) 
void QVideoFrameFormat_Delete(QVideoFrameFormat* self, bool isSubclass);

}
