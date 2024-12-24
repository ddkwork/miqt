#pragma once
#ifndef MIQT_QT_GEN_QPAINTDEVICEWINDOW_H
#define MIQT_QT_GEN_QPAINTDEVICEWINDOW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QExposeEvent QExposeEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintDeviceWindow QPaintDeviceWindow;
typedef struct QPaintEvent QPaintEvent;
typedef struct QRect QRect;
typedef struct QRegion QRegion;
typedef struct QSurface QSurface;
typedef struct QWindow QWindow;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QPaintDeviceWindow_virtbase(QPaintDeviceWindow* src, QWindow** outptr_QWindow, QPaintDevice** outptr_QPaintDevice);
extern __declspec(dllexport) QMetaObject* QPaintDeviceWindow_MetaObject(const QPaintDeviceWindow* self);
extern __declspec(dllexport) void* QPaintDeviceWindow_Metacast(QPaintDeviceWindow* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QPaintDeviceWindow_Tr(const char* s);
extern __declspec(dllexport) void QPaintDeviceWindow_Update(QPaintDeviceWindow* self, QRect* rect);
extern __declspec(dllexport) void QPaintDeviceWindow_UpdateWithRegion(QPaintDeviceWindow* self, QRegion* region);
extern __declspec(dllexport) void QPaintDeviceWindow_Update2(QPaintDeviceWindow* self);
extern __declspec(dllexport) void QPaintDeviceWindow_ExposeEvent(QPaintDeviceWindow* self, QExposeEvent* param1);
extern __declspec(dllexport) void QPaintDeviceWindow_PaintEvent(QPaintDeviceWindow* self, QPaintEvent* event);
extern __declspec(dllexport) int QPaintDeviceWindow_Metric(const QPaintDeviceWindow* self, PaintDeviceMetric metric);
extern __declspec(dllexport) bool QPaintDeviceWindow_Event(QPaintDeviceWindow* self, QEvent* event);
extern __declspec(dllexport) struct miqt_string QPaintDeviceWindow_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QPaintDeviceWindow_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QPaintDeviceWindow_Delete(QPaintDeviceWindow* self, bool isSubclass);

} 
