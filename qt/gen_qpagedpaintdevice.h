#pragma once
#ifndef MIQT_QT_GEN_QPAGEDPAINTDEVICE_H
#define MIQT_QT_GEN_QPAGEDPAINTDEVICE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMarginsF QMarginsF;
typedef struct QPageLayout QPageLayout;
typedef struct QPageRanges QPageRanges;
typedef struct QPageSize QPageSize;
typedef struct QPagedPaintDevice QPagedPaintDevice;
typedef struct QPaintDevice QPaintDevice;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QPagedPaintDevice_virtbase(QPagedPaintDevice* src, QPaintDevice** outptr_QPaintDevice);
extern __declspec(dllexport) bool QPagedPaintDevice_NewPage(QPagedPaintDevice* self);
extern __declspec(dllexport) bool QPagedPaintDevice_SetPageLayout(QPagedPaintDevice* self, QPageLayout* pageLayout);
extern __declspec(dllexport) bool QPagedPaintDevice_SetPageSize(QPagedPaintDevice* self, QPageSize* pageSize);
extern __declspec(dllexport) bool QPagedPaintDevice_SetPageOrientation(QPagedPaintDevice* self, int orientation);
extern __declspec(dllexport) bool QPagedPaintDevice_SetPageMargins(QPagedPaintDevice* self, QMarginsF* margins, int units);
extern __declspec(dllexport) QPageLayout* QPagedPaintDevice_PageLayout(const QPagedPaintDevice* self);
extern __declspec(dllexport) void QPagedPaintDevice_SetPageRanges(QPagedPaintDevice* self, QPageRanges* ranges);
extern __declspec(dllexport) QPageRanges* QPagedPaintDevice_PageRanges(const QPagedPaintDevice* self);
extern __declspec(dllexport) void QPagedPaintDevice_Delete(QPagedPaintDevice* self, bool isSubclass);

} 
