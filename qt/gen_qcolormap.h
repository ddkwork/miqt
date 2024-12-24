#pragma once
#ifndef MIQT_QT_GEN_QCOLORMAP_H
#define MIQT_QT_GEN_QCOLORMAP_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QColor QColor;
typedef struct QColormap QColormap;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QColormap* QColormap_new(QColormap* colormap);
extern __declspec(dllexport) void QColormap_Initialize();
extern __declspec(dllexport) void QColormap_Cleanup();
extern __declspec(dllexport) QColormap* QColormap_Instance();
extern __declspec(dllexport) void QColormap_OperatorAssign(QColormap* self, QColormap* colormap);
extern __declspec(dllexport) Mode QColormap_Mode(const QColormap* self);
extern __declspec(dllexport) int QColormap_Depth(const QColormap* self);
extern __declspec(dllexport) int QColormap_Size(const QColormap* self);
extern __declspec(dllexport) unsigned int QColormap_Pixel(const QColormap* self, QColor* color);
extern __declspec(dllexport) QColor* QColormap_ColorAt(const QColormap* self, unsigned int pixel);
extern __declspec(dllexport) struct miqt_array /* of QColor* */  QColormap_Colormap(const QColormap* self);
extern __declspec(dllexport) QColormap* QColormap_Instance1(int screen);
extern __declspec(dllexport) void QColormap_Delete(QColormap* self, bool isSubclass);

} 
