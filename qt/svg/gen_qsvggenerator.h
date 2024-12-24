#pragma once
#ifndef MIQT_QT_SVG_GEN_QSVGGENERATOR_H
#define MIQT_QT_SVG_GEN_QSVGGENERATOR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QIODevice QIODevice;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPainter QPainter;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QRectF QRectF;
typedef struct QSize QSize;
typedef struct QSvgGenerator QSvgGenerator;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSvgGenerator* QSvgGenerator_new();
extern __declspec(dllexport) QSvgGenerator* QSvgGenerator_new2(SvgVersion version);
extern __declspec(dllexport) void QSvgGenerator_virtbase(QSvgGenerator* src, QPaintDevice** outptr_QPaintDevice);
extern __declspec(dllexport) struct miqt_string QSvgGenerator_Title(const QSvgGenerator* self);
extern __declspec(dllexport) void QSvgGenerator_SetTitle(QSvgGenerator* self, struct miqt_string title);
extern __declspec(dllexport) struct miqt_string QSvgGenerator_Description(const QSvgGenerator* self);
extern __declspec(dllexport) void QSvgGenerator_SetDescription(QSvgGenerator* self, struct miqt_string description);
extern __declspec(dllexport) QSize* QSvgGenerator_Size(const QSvgGenerator* self);
extern __declspec(dllexport) void QSvgGenerator_SetSize(QSvgGenerator* self, QSize* size);
extern __declspec(dllexport) QRect* QSvgGenerator_ViewBox(const QSvgGenerator* self);
extern __declspec(dllexport) QRectF* QSvgGenerator_ViewBoxF(const QSvgGenerator* self);
extern __declspec(dllexport) void QSvgGenerator_SetViewBox(QSvgGenerator* self, QRect* viewBox);
extern __declspec(dllexport) void QSvgGenerator_SetViewBoxWithViewBox(QSvgGenerator* self, QRectF* viewBox);
extern __declspec(dllexport) struct miqt_string QSvgGenerator_FileName(const QSvgGenerator* self);
extern __declspec(dllexport) void QSvgGenerator_SetFileName(QSvgGenerator* self, struct miqt_string fileName);
extern __declspec(dllexport) QIODevice* QSvgGenerator_OutputDevice(const QSvgGenerator* self);
extern __declspec(dllexport) void QSvgGenerator_SetOutputDevice(QSvgGenerator* self, QIODevice* outputDevice);
extern __declspec(dllexport) void QSvgGenerator_SetResolution(QSvgGenerator* self, int dpi);
extern __declspec(dllexport) int QSvgGenerator_Resolution(const QSvgGenerator* self);
extern __declspec(dllexport) SvgVersion QSvgGenerator_SvgVersion(const QSvgGenerator* self);
extern __declspec(dllexport) QPaintEngine* QSvgGenerator_PaintEngine(const QSvgGenerator* self);
extern __declspec(dllexport) int QSvgGenerator_Metric(const QSvgGenerator* self, int metric);
extern __declspec(dllexport) void QSvgGenerator_override_virtual_PaintEngine(void* self, intptr_t slot);
QPaintEngine* QSvgGenerator_virtualbase_PaintEngine(const void* self);
extern __declspec(dllexport) void QSvgGenerator_override_virtual_Metric(void* self, intptr_t slot);
int QSvgGenerator_virtualbase_Metric(const void* self, int metric);
extern __declspec(dllexport) void QSvgGenerator_override_virtual_DevType(void* self, intptr_t slot);
int QSvgGenerator_virtualbase_DevType(const void* self);
extern __declspec(dllexport) void QSvgGenerator_override_virtual_InitPainter(void* self, intptr_t slot);
void QSvgGenerator_virtualbase_InitPainter(const void* self, QPainter* painter);
extern __declspec(dllexport) void QSvgGenerator_override_virtual_Redirected(void* self, intptr_t slot);
QPaintDevice* QSvgGenerator_virtualbase_Redirected(const void* self, QPoint* offset);
extern __declspec(dllexport) void QSvgGenerator_override_virtual_SharedPainter(void* self, intptr_t slot);
QPainter* QSvgGenerator_virtualbase_SharedPainter(const void* self);
extern __declspec(dllexport) void QSvgGenerator_Delete(QSvgGenerator* self, bool isSubclass);

} 
