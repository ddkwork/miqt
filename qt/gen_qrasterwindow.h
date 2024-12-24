#pragma once
#ifndef MIQT_QT_GEN_QRASTERWINDOW_H
#define MIQT_QT_GEN_QRASTERWINDOW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QEvent;
class QExposeEvent;
class QMetaObject;
class QObject;
class QPaintDevice;
class QPaintDeviceWindow;
class QPaintEvent;
class QPoint;
class QRasterWindow;
class QResizeEvent;
class QSurface;
class QWindow;
class _GUID;
class type_info;
#else
typedef struct QEvent QEvent;
typedef struct QExposeEvent QExposeEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintDeviceWindow QPaintDeviceWindow;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPoint QPoint;
typedef struct QRasterWindow QRasterWindow;
typedef struct QResizeEvent QResizeEvent;
typedef struct QSurface QSurface;
typedef struct QWindow QWindow;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QRasterWindow* QRasterWindow_new();
extern __declspec(dllexport) QRasterWindow* QRasterWindow_new2(QWindow* parent);
extern __declspec(dllexport) void QRasterWindow_virtbase(QRasterWindow* src, QPaintDeviceWindow** outptr_QPaintDeviceWindow);
extern __declspec(dllexport) QMetaObject* QRasterWindow_MetaObject(const QRasterWindow* self);
extern __declspec(dllexport) void* QRasterWindow_Metacast(QRasterWindow* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QRasterWindow_Tr(const char* s);
extern __declspec(dllexport) int QRasterWindow_Metric(const QRasterWindow* self, PaintDeviceMetric metric);
extern __declspec(dllexport) QPaintDevice* QRasterWindow_Redirected(const QRasterWindow* self, QPoint* param1);
extern __declspec(dllexport) void QRasterWindow_ResizeEvent(QRasterWindow* self, QResizeEvent* event);
extern __declspec(dllexport) struct miqt_string QRasterWindow_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QRasterWindow_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QRasterWindow_override_virtual_Metric(void* self, intptr_t slot);
int QRasterWindow_virtualbase_Metric(const void* self, PaintDeviceMetric metric);
extern __declspec(dllexport) void QRasterWindow_override_virtual_Redirected(void* self, intptr_t slot);
QPaintDevice* QRasterWindow_virtualbase_Redirected(const void* self, QPoint* param1);
extern __declspec(dllexport) void QRasterWindow_override_virtual_ResizeEvent(void* self, intptr_t slot);
void QRasterWindow_virtualbase_ResizeEvent(void* self, QResizeEvent* event);
extern __declspec(dllexport) void QRasterWindow_override_virtual_ExposeEvent(void* self, intptr_t slot);
void QRasterWindow_virtualbase_ExposeEvent(void* self, QExposeEvent* param1);
extern __declspec(dllexport) void QRasterWindow_override_virtual_PaintEvent(void* self, intptr_t slot);
void QRasterWindow_virtualbase_PaintEvent(void* self, QPaintEvent* event);
extern __declspec(dllexport) void QRasterWindow_override_virtual_Event(void* self, intptr_t slot);
bool QRasterWindow_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QRasterWindow_Delete(QRasterWindow* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
