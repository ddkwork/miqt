#pragma once
#ifndef MIQT_QT_GEN_QBACKINGSTORE_H
#define MIQT_QT_GEN_QBACKINGSTORE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QBackingStore QBackingStore;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPoint QPoint;
typedef struct QRegion QRegion;
typedef struct QSize QSize;
typedef struct QWindow QWindow;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QBackingStore* QBackingStore_new(QWindow* window);
extern __declspec(dllexport) 
QWindow* QBackingStore_Window(const QBackingStore* self);
extern __declspec(dllexport) 
QPaintDevice* QBackingStore_PaintDevice(QBackingStore* self);
extern __declspec(dllexport) 
void QBackingStore_Flush(QBackingStore* self, QRegion* region);
extern __declspec(dllexport) 
void QBackingStore_Resize(QBackingStore* self, QSize* size);
extern __declspec(dllexport) 
QSize* QBackingStore_Size(const QBackingStore* self);
extern __declspec(dllexport) 
bool QBackingStore_Scroll(QBackingStore* self, QRegion* area, int dx, int dy);
extern __declspec(dllexport) 
void QBackingStore_BeginPaint(QBackingStore* self, QRegion* param1);
extern __declspec(dllexport) 
void QBackingStore_EndPaint(QBackingStore* self);
extern __declspec(dllexport) 
void QBackingStore_SetStaticContents(QBackingStore* self, QRegion* region);
extern __declspec(dllexport) 
QRegion* QBackingStore_StaticContents(const QBackingStore* self);
extern __declspec(dllexport) 
bool QBackingStore_HasStaticContents(const QBackingStore* self);
extern __declspec(dllexport) 
void QBackingStore_Flush2(QBackingStore* self, QRegion* region, QWindow* window);
extern __declspec(dllexport) 
void QBackingStore_Flush3(QBackingStore* self, QRegion* region, QWindow* window, QPoint* offset);
extern __declspec(dllexport) 
void QBackingStore_Delete(QBackingStore* self, bool isSubclass);

}
