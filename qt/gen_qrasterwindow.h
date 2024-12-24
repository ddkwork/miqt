#pragma once
#ifndef MIQT_QT_GEN_QRASTERWINDOW_H
#define MIQT_QT_GEN_QRASTERWINDOW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintDeviceWindow QPaintDeviceWindow;
typedef struct QPoint QPoint;
typedef struct QRasterWindow QRasterWindow;
typedef struct QResizeEvent QResizeEvent;
typedef struct QSurface QSurface;
typedef struct QWindow QWindow;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QRasterWindow* QRasterWindow_new();
extern __declspec(dllexport) 
QRasterWindow* QRasterWindow_new2(QWindow* parent);
extern __declspec(dllexport) 
void QRasterWindow_virtbase(QRasterWindow* src
, QPaintDeviceWindow** outptr_QPaintDeviceWindow
);
extern __declspec(dllexport) 
QMetaObject* QRasterWindow_MetaObject(const QRasterWindow* self);
extern __declspec(dllexport) 
void* QRasterWindow_Metacast(QRasterWindow* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QRasterWindow_Tr(const char* s);
extern __declspec(dllexport) 
int QRasterWindow_Metric(const QRasterWindow* self, PaintDeviceMetric metric);
extern __declspec(dllexport) 
QPaintDevice* QRasterWindow_Redirected(const QRasterWindow* self, QPoint* param1);
extern __declspec(dllexport) 
void QRasterWindow_ResizeEvent(QRasterWindow* self, QResizeEvent* event);
extern __declspec(dllexport) 
struct miqt_string QRasterWindow_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QRasterWindow_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QRasterWindow_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QRasterWindow_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QRasterWindow_override_virtual_Metacast(void* self, intptr_t slot);
void* QRasterWindow_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QRasterWindow_Delete(QRasterWindow* self, bool isSubclass);

}
