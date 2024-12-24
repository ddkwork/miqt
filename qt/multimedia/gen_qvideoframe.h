#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QVIDEOFRAME_H
#define MIQT_QT_MULTIMEDIA_GEN_QVIDEOFRAME_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QVideoFrame__PaintOptions)
typedef QVideoFrame::PaintOptions QVideoFrame__PaintOptions;
typedef struct QImage QImage;
typedef struct QPainter QPainter;
typedef struct QRectF QRectF;
typedef struct QSize QSize;
typedef struct QVideoFrame QVideoFrame;
typedef struct QVideoFrame__PaintOptions QVideoFrame__PaintOptions;
typedef struct QVideoFrameFormat QVideoFrameFormat;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QVideoFrame* QVideoFrame_new();
extern __declspec(dllexport) 
QVideoFrame* QVideoFrame_new2(QVideoFrameFormat* format);
extern __declspec(dllexport) 
QVideoFrame* QVideoFrame_new3(QImage* image);
extern __declspec(dllexport) 
QVideoFrame* QVideoFrame_new4(QVideoFrame* other);
extern __declspec(dllexport) 
void QVideoFrame_Swap(QVideoFrame* self, QVideoFrame* other);
extern __declspec(dllexport) 
void QVideoFrame_OperatorAssign(QVideoFrame* self, QVideoFrame* other);
extern __declspec(dllexport) 
bool QVideoFrame_OperatorEqual(const QVideoFrame* self, QVideoFrame* other);
extern __declspec(dllexport) 
bool QVideoFrame_OperatorNotEqual(const QVideoFrame* self, QVideoFrame* other);
extern __declspec(dllexport) 
bool QVideoFrame_IsValid(const QVideoFrame* self);
extern __declspec(dllexport) 
int QVideoFrame_PixelFormat(const QVideoFrame* self);
extern __declspec(dllexport) 
QVideoFrameFormat* QVideoFrame_SurfaceFormat(const QVideoFrame* self);
extern __declspec(dllexport) 
int QVideoFrame_HandleType(const QVideoFrame* self);
extern __declspec(dllexport) 
QSize* QVideoFrame_Size(const QVideoFrame* self);
extern __declspec(dllexport) 
int QVideoFrame_Width(const QVideoFrame* self);
extern __declspec(dllexport) 
int QVideoFrame_Height(const QVideoFrame* self);
extern __declspec(dllexport) 
bool QVideoFrame_IsMapped(const QVideoFrame* self);
extern __declspec(dllexport) 
bool QVideoFrame_IsReadable(const QVideoFrame* self);
extern __declspec(dllexport) 
bool QVideoFrame_IsWritable(const QVideoFrame* self);
extern __declspec(dllexport) 
int QVideoFrame_MapMode(const QVideoFrame* self);
extern __declspec(dllexport) 
bool QVideoFrame_Map(QVideoFrame* self, int mode);
extern __declspec(dllexport) 
void QVideoFrame_Unmap(QVideoFrame* self);
extern __declspec(dllexport) 
int QVideoFrame_BytesPerLine(const QVideoFrame* self, int plane);
extern __declspec(dllexport) 
unsigned char* QVideoFrame_Bits(QVideoFrame* self, int plane);
extern __declspec(dllexport) 
const unsigned char* QVideoFrame_BitsWithPlane(const QVideoFrame* self, int plane);
extern __declspec(dllexport) 
int QVideoFrame_MappedBytes(const QVideoFrame* self, int plane);
extern __declspec(dllexport) 
int QVideoFrame_PlaneCount(const QVideoFrame* self);
extern __declspec(dllexport) 
long long QVideoFrame_StartTime(const QVideoFrame* self);
extern __declspec(dllexport) 
void QVideoFrame_SetStartTime(QVideoFrame* self, long long time);
extern __declspec(dllexport) 
long long QVideoFrame_EndTime(const QVideoFrame* self);
extern __declspec(dllexport) 
void QVideoFrame_SetEndTime(QVideoFrame* self, long long time);
extern __declspec(dllexport) 
void QVideoFrame_SetRotationAngle(QVideoFrame* self, RotationAngle angle);
extern __declspec(dllexport) 
RotationAngle QVideoFrame_RotationAngle(const QVideoFrame* self);
extern __declspec(dllexport) 
void QVideoFrame_SetRotation(QVideoFrame* self, int angle);
extern __declspec(dllexport) 
int QVideoFrame_Rotation(const QVideoFrame* self);
extern __declspec(dllexport) 
void QVideoFrame_SetMirrored(QVideoFrame* self, bool mirrored);
extern __declspec(dllexport) 
bool QVideoFrame_Mirrored(const QVideoFrame* self);
extern __declspec(dllexport) 
void QVideoFrame_SetStreamFrameRate(QVideoFrame* self, double rate);
extern __declspec(dllexport) 
double QVideoFrame_StreamFrameRate(const QVideoFrame* self);
extern __declspec(dllexport) 
QImage* QVideoFrame_ToImage(const QVideoFrame* self);
extern __declspec(dllexport) 
struct miqt_string QVideoFrame_SubtitleText(const QVideoFrame* self);
extern __declspec(dllexport) 
void QVideoFrame_SetSubtitleText(QVideoFrame* self, struct miqt_string text);
extern __declspec(dllexport) 
void QVideoFrame_Paint(QVideoFrame* self, QPainter* painter, QRectF* rect, const PaintOptions* options);
extern __declspec(dllexport) 
void QVideoFrame_Delete(QVideoFrame* self, bool isSubclass);

extern __declspec(dllexport) 
void QVideoFrame__PaintOptions_Delete(QVideoFrame__PaintOptions* self, bool isSubclass);

}
