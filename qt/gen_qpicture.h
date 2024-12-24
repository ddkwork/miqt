#pragma once
#ifndef MIQT_QT_GEN_QPICTURE_H
#define MIQT_QT_GEN_QPICTURE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QIODevice QIODevice;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPainter QPainter;
typedef struct QPicture QPicture;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QPicture* QPicture_new();
extern __declspec(dllexport) QPicture* QPicture_new2(QPicture* param1);
extern __declspec(dllexport) QPicture* QPicture_new3(int formatVersion);
extern __declspec(dllexport) void QPicture_virtbase(QPicture* src, QPaintDevice** outptr_QPaintDevice);
extern __declspec(dllexport) bool QPicture_IsNull(const QPicture* self);
extern __declspec(dllexport) int QPicture_DevType(const QPicture* self);
extern __declspec(dllexport) unsigned int QPicture_Size(const QPicture* self);
extern __declspec(dllexport) const char* QPicture_Data(const QPicture* self);
extern __declspec(dllexport) void QPicture_SetData(QPicture* self, const char* data, unsigned int size);
extern __declspec(dllexport) bool QPicture_Play(QPicture* self, QPainter* p);
extern __declspec(dllexport) bool QPicture_Load(QPicture* self, QIODevice* dev);
extern __declspec(dllexport) bool QPicture_LoadWithFileName(QPicture* self, struct miqt_string fileName);
extern __declspec(dllexport) bool QPicture_Save(QPicture* self, QIODevice* dev);
extern __declspec(dllexport) bool QPicture_SaveWithFileName(QPicture* self, struct miqt_string fileName);
extern __declspec(dllexport) QRect* QPicture_BoundingRect(const QPicture* self);
extern __declspec(dllexport) void QPicture_SetBoundingRect(QPicture* self, QRect* r);
extern __declspec(dllexport) void QPicture_OperatorAssign(QPicture* self, QPicture* p);
extern __declspec(dllexport) void QPicture_Swap(QPicture* self, QPicture* other);
extern __declspec(dllexport) void QPicture_Detach(QPicture* self);
extern __declspec(dllexport) bool QPicture_IsDetached(const QPicture* self);
extern __declspec(dllexport) QPaintEngine* QPicture_PaintEngine(const QPicture* self);
extern __declspec(dllexport) int QPicture_Metric(const QPicture* self, PaintDeviceMetric m);
extern __declspec(dllexport) DataPtr* QPicture_DataPtr(QPicture* self);
extern __declspec(dllexport) void QPicture_override_virtual_DevType(void* self, intptr_t slot);
int QPicture_virtualbase_DevType(const void* self);
extern __declspec(dllexport) void QPicture_override_virtual_SetData(void* self, intptr_t slot);
void QPicture_virtualbase_SetData(void* self, const char* data, unsigned int size);
extern __declspec(dllexport) void QPicture_override_virtual_PaintEngine(void* self, intptr_t slot);
QPaintEngine* QPicture_virtualbase_PaintEngine(const void* self);
extern __declspec(dllexport) void QPicture_override_virtual_Metric(void* self, intptr_t slot);
int QPicture_virtualbase_Metric(const void* self, PaintDeviceMetric m);
extern __declspec(dllexport) void QPicture_override_virtual_InitPainter(void* self, intptr_t slot);
void QPicture_virtualbase_InitPainter(const void* self, QPainter* painter);
extern __declspec(dllexport) void QPicture_override_virtual_Redirected(void* self, intptr_t slot);
QPaintDevice* QPicture_virtualbase_Redirected(const void* self, QPoint* offset);
extern __declspec(dllexport) void QPicture_override_virtual_SharedPainter(void* self, intptr_t slot);
QPainter* QPicture_virtualbase_SharedPainter(const void* self);
extern __declspec(dllexport) void QPicture_Delete(QPicture* self, bool isSubclass);

} 
