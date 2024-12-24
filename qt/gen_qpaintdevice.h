#pragma once
#ifndef MIQT_QT_GEN_QPAINTDEVICE_H
#define MIQT_QT_GEN_QPAINTDEVICE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QPaintDevice;
class QPaintEngine;
class QPainter;
class QPoint;
class _GUID;
class type_info;
#else
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPainter QPainter;
typedef struct QPoint QPoint;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) int QPaintDevice_DevType(const QPaintDevice* self);
extern __declspec(dllexport) bool QPaintDevice_PaintingActive(const QPaintDevice* self);
extern __declspec(dllexport) QPaintEngine* QPaintDevice_PaintEngine(const QPaintDevice* self);
extern __declspec(dllexport) int QPaintDevice_Width(const QPaintDevice* self);
extern __declspec(dllexport) int QPaintDevice_Height(const QPaintDevice* self);
extern __declspec(dllexport) int QPaintDevice_WidthMM(const QPaintDevice* self);
extern __declspec(dllexport) int QPaintDevice_HeightMM(const QPaintDevice* self);
extern __declspec(dllexport) int QPaintDevice_LogicalDpiX(const QPaintDevice* self);
extern __declspec(dllexport) int QPaintDevice_LogicalDpiY(const QPaintDevice* self);
extern __declspec(dllexport) int QPaintDevice_PhysicalDpiX(const QPaintDevice* self);
extern __declspec(dllexport) int QPaintDevice_PhysicalDpiY(const QPaintDevice* self);
extern __declspec(dllexport) double QPaintDevice_DevicePixelRatio(const QPaintDevice* self);
extern __declspec(dllexport) double QPaintDevice_DevicePixelRatioF(const QPaintDevice* self);
extern __declspec(dllexport) int QPaintDevice_ColorCount(const QPaintDevice* self);
extern __declspec(dllexport) int QPaintDevice_Depth(const QPaintDevice* self);
extern __declspec(dllexport) double QPaintDevice_DevicePixelRatioFScale();
extern __declspec(dllexport) int QPaintDevice_EncodeMetricF(PaintDeviceMetric metric, double value);
extern __declspec(dllexport) int QPaintDevice_Metric(const QPaintDevice* self, PaintDeviceMetric metric);
extern __declspec(dllexport) void QPaintDevice_InitPainter(const QPaintDevice* self, QPainter* painter);
extern __declspec(dllexport) QPaintDevice* QPaintDevice_Redirected(const QPaintDevice* self, QPoint* offset);
extern __declspec(dllexport) QPainter* QPaintDevice_SharedPainter(const QPaintDevice* self);
extern __declspec(dllexport) void QPaintDevice_Delete(QPaintDevice* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
