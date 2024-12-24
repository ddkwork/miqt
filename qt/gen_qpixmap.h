#pragma once
#ifndef MIQT_QT_GEN_QPIXMAP_H
#define MIQT_QT_GEN_QPIXMAP_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QBitmap QBitmap;
typedef struct QColor QColor;
typedef struct QIODevice QIODevice;
typedef struct QImage QImage;
typedef struct QImageReader QImageReader;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPainter QPainter;
typedef struct QPixmap QPixmap;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QRegion QRegion;
typedef struct QSize QSize;
typedef struct QSizeF QSizeF;
typedef struct QTransform QTransform;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QPixmap* QPixmap_new();
extern __declspec(dllexport) QPixmap* QPixmap_new2(int w, int h);
extern __declspec(dllexport) QPixmap* QPixmap_new3(QSize* param1);
extern __declspec(dllexport) QPixmap* QPixmap_new4(struct miqt_string fileName);
extern __declspec(dllexport) QPixmap* QPixmap_new5(QPixmap* param1);
extern __declspec(dllexport) QPixmap* QPixmap_new6(struct miqt_string fileName, const char* format);
extern __declspec(dllexport) QPixmap* QPixmap_new7(struct miqt_string fileName, const char* format, int flags);
extern __declspec(dllexport) void QPixmap_virtbase(QPixmap* src, QPaintDevice** outptr_QPaintDevice);
extern __declspec(dllexport) void QPixmap_OperatorAssign(QPixmap* self, QPixmap* param1);
extern __declspec(dllexport) void QPixmap_Swap(QPixmap* self, QPixmap* other);
extern __declspec(dllexport) bool QPixmap_IsNull(const QPixmap* self);
extern __declspec(dllexport) int QPixmap_DevType(const QPixmap* self);
extern __declspec(dllexport) int QPixmap_Width(const QPixmap* self);
extern __declspec(dllexport) int QPixmap_Height(const QPixmap* self);
extern __declspec(dllexport) QSize* QPixmap_Size(const QPixmap* self);
extern __declspec(dllexport) QRect* QPixmap_Rect(const QPixmap* self);
extern __declspec(dllexport) int QPixmap_Depth(const QPixmap* self);
extern __declspec(dllexport) int QPixmap_DefaultDepth();
extern __declspec(dllexport) void QPixmap_Fill(QPixmap* self);
extern __declspec(dllexport) QBitmap* QPixmap_Mask(const QPixmap* self);
extern __declspec(dllexport) void QPixmap_SetMask(QPixmap* self, QBitmap* mask);
extern __declspec(dllexport) double QPixmap_DevicePixelRatio(const QPixmap* self);
extern __declspec(dllexport) void QPixmap_SetDevicePixelRatio(QPixmap* self, double scaleFactor);
extern __declspec(dllexport) QSizeF* QPixmap_DeviceIndependentSize(const QPixmap* self);
extern __declspec(dllexport) bool QPixmap_HasAlpha(const QPixmap* self);
extern __declspec(dllexport) bool QPixmap_HasAlphaChannel(const QPixmap* self);
extern __declspec(dllexport) QBitmap* QPixmap_CreateHeuristicMask(const QPixmap* self);
extern __declspec(dllexport) QBitmap* QPixmap_CreateMaskFromColor(const QPixmap* self, QColor* maskColor);
extern __declspec(dllexport) QPixmap* QPixmap_Scaled(const QPixmap* self, int w, int h);
extern __declspec(dllexport) QPixmap* QPixmap_ScaledWithQSize(const QPixmap* self, QSize* s);
extern __declspec(dllexport) QPixmap* QPixmap_ScaledToWidth(const QPixmap* self, int w);
extern __declspec(dllexport) QPixmap* QPixmap_ScaledToHeight(const QPixmap* self, int h);
extern __declspec(dllexport) QPixmap* QPixmap_Transformed(const QPixmap* self, QTransform* param1);
extern __declspec(dllexport) QTransform* QPixmap_TrueMatrix(QTransform* m, int w, int h);
extern __declspec(dllexport) QImage* QPixmap_ToImage(const QPixmap* self);
extern __declspec(dllexport) QPixmap* QPixmap_FromImage(QImage* image);
extern __declspec(dllexport) QPixmap* QPixmap_FromImageReader(QImageReader* imageReader);
extern __declspec(dllexport) bool QPixmap_Load(QPixmap* self, struct miqt_string fileName);
extern __declspec(dllexport) bool QPixmap_LoadFromData(QPixmap* self, const unsigned char* buf, unsigned int lenVal);
extern __declspec(dllexport) bool QPixmap_LoadFromDataWithData(QPixmap* self, struct miqt_string data);
extern __declspec(dllexport) bool QPixmap_Save(const QPixmap* self, struct miqt_string fileName);
extern __declspec(dllexport) bool QPixmap_SaveWithDevice(const QPixmap* self, QIODevice* device);
extern __declspec(dllexport) bool QPixmap_ConvertFromImage(QPixmap* self, QImage* img);
extern __declspec(dllexport) QPixmap* QPixmap_Copy(const QPixmap* self, int x, int y, int width, int height);
extern __declspec(dllexport) QPixmap* QPixmap_Copy2(const QPixmap* self);
extern __declspec(dllexport) void QPixmap_Scroll(QPixmap* self, int dx, int dy, int x, int y, int width, int height);
extern __declspec(dllexport) void QPixmap_Scroll2(QPixmap* self, int dx, int dy, QRect* rect);
extern __declspec(dllexport) long long QPixmap_CacheKey(const QPixmap* self);
extern __declspec(dllexport) bool QPixmap_IsDetached(const QPixmap* self);
extern __declspec(dllexport) void QPixmap_Detach(QPixmap* self);
extern __declspec(dllexport) bool QPixmap_IsQBitmap(const QPixmap* self);
extern __declspec(dllexport) QPaintEngine* QPixmap_PaintEngine(const QPixmap* self);
extern __declspec(dllexport) bool QPixmap_OperatorNot(const QPixmap* self);
extern __declspec(dllexport) int QPixmap_Metric(const QPixmap* self, PaintDeviceMetric param1);
extern __declspec(dllexport) DataPtr* QPixmap_DataPtr(QPixmap* self);
extern __declspec(dllexport) void QPixmap_Fill1(QPixmap* self, QColor* fillColor);
extern __declspec(dllexport) QBitmap* QPixmap_CreateHeuristicMask1(const QPixmap* self, bool clipTight);
extern __declspec(dllexport) QBitmap* QPixmap_CreateMaskFromColor2(const QPixmap* self, QColor* maskColor, int mode);
extern __declspec(dllexport) QPixmap* QPixmap_Scaled3(const QPixmap* self, int w, int h, int aspectMode);
extern __declspec(dllexport) QPixmap* QPixmap_Scaled4(const QPixmap* self, int w, int h, int aspectMode, int mode);
extern __declspec(dllexport) QPixmap* QPixmap_Scaled2(const QPixmap* self, QSize* s, int aspectMode);
extern __declspec(dllexport) QPixmap* QPixmap_Scaled32(const QPixmap* self, QSize* s, int aspectMode, int mode);
extern __declspec(dllexport) QPixmap* QPixmap_ScaledToWidth2(const QPixmap* self, int w, int mode);
extern __declspec(dllexport) QPixmap* QPixmap_ScaledToHeight2(const QPixmap* self, int h, int mode);
extern __declspec(dllexport) QPixmap* QPixmap_Transformed2(const QPixmap* self, QTransform* param1, int mode);
extern __declspec(dllexport) QPixmap* QPixmap_FromImage2(QImage* image, int flags);
extern __declspec(dllexport) QPixmap* QPixmap_FromImageReader2(QImageReader* imageReader, int flags);
extern __declspec(dllexport) bool QPixmap_Load2(QPixmap* self, struct miqt_string fileName, const char* format);
extern __declspec(dllexport) bool QPixmap_Load3(QPixmap* self, struct miqt_string fileName, const char* format, int flags);
extern __declspec(dllexport) bool QPixmap_LoadFromData3(QPixmap* self, const unsigned char* buf, unsigned int lenVal, const char* format);
extern __declspec(dllexport) bool QPixmap_LoadFromData4(QPixmap* self, const unsigned char* buf, unsigned int lenVal, const char* format, int flags);
extern __declspec(dllexport) bool QPixmap_LoadFromData2(QPixmap* self, struct miqt_string data, const char* format);
extern __declspec(dllexport) bool QPixmap_LoadFromData32(QPixmap* self, struct miqt_string data, const char* format, int flags);
extern __declspec(dllexport) bool QPixmap_Save2(const QPixmap* self, struct miqt_string fileName, const char* format);
extern __declspec(dllexport) bool QPixmap_Save3(const QPixmap* self, struct miqt_string fileName, const char* format, int quality);
extern __declspec(dllexport) bool QPixmap_Save22(const QPixmap* self, QIODevice* device, const char* format);
extern __declspec(dllexport) bool QPixmap_Save32(const QPixmap* self, QIODevice* device, const char* format, int quality);
extern __declspec(dllexport) bool QPixmap_ConvertFromImage2(QPixmap* self, QImage* img, int flags);
extern __declspec(dllexport) QPixmap* QPixmap_Copy1(const QPixmap* self, QRect* rect);
extern __declspec(dllexport) void QPixmap_Scroll7(QPixmap* self, int dx, int dy, int x, int y, int width, int height, QRegion* exposed);
extern __declspec(dllexport) void QPixmap_Scroll4(QPixmap* self, int dx, int dy, QRect* rect, QRegion* exposed);
extern __declspec(dllexport) void QPixmap_override_virtual_DevType(void* self, intptr_t slot);
int QPixmap_virtualbase_DevType(const void* self);
extern __declspec(dllexport) void QPixmap_override_virtual_PaintEngine(void* self, intptr_t slot);
QPaintEngine* QPixmap_virtualbase_PaintEngine(const void* self);
extern __declspec(dllexport) void QPixmap_override_virtual_Metric(void* self, intptr_t slot);
int QPixmap_virtualbase_Metric(const void* self, PaintDeviceMetric param1);
extern __declspec(dllexport) void QPixmap_override_virtual_InitPainter(void* self, intptr_t slot);
void QPixmap_virtualbase_InitPainter(const void* self, QPainter* painter);
extern __declspec(dllexport) void QPixmap_override_virtual_Redirected(void* self, intptr_t slot);
QPaintDevice* QPixmap_virtualbase_Redirected(const void* self, QPoint* offset);
extern __declspec(dllexport) void QPixmap_override_virtual_SharedPainter(void* self, intptr_t slot);
QPainter* QPixmap_virtualbase_SharedPainter(const void* self);
extern __declspec(dllexport) void QPixmap_Delete(QPixmap* self, bool isSubclass);

} 
