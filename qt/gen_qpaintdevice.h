#pragma once
#ifndef MIQT_QT_GEN_QPAINTDEVICE_H
#define MIQT_QT_GEN_QPAINTDEVICE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPainter QPainter;
typedef struct QPoint QPoint;

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

} 
