#pragma once
#ifndef MIQT_QT_GEN_QCOLORTRANSFORM_H
#define MIQT_QT_GEN_QCOLORTRANSFORM_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QColor;
class QColorTransform;
class QRgba64;
class _GUID;
class type_info;
#else
typedef struct QColor QColor;
typedef struct QColorTransform QColorTransform;
typedef struct QRgba64 QRgba64;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QColorTransform* QColorTransform_new();
extern __declspec(dllexport) QColorTransform* QColorTransform_new2(QColorTransform* colorTransform);
extern __declspec(dllexport) void QColorTransform_OperatorAssign(QColorTransform* self, QColorTransform* other);
extern __declspec(dllexport) void QColorTransform_Swap(QColorTransform* self, QColorTransform* other);
extern __declspec(dllexport) bool QColorTransform_IsIdentity(const QColorTransform* self);
extern __declspec(dllexport) unsigned int QColorTransform_Map(const QColorTransform* self, unsigned int argb);
extern __declspec(dllexport) QRgba64* QColorTransform_MapWithRgba64(const QColorTransform* self, QRgba64* rgba64);
extern __declspec(dllexport) QColor* QColorTransform_MapWithColor(const QColorTransform* self, QColor* color);
extern __declspec(dllexport) void QColorTransform_Delete(QColorTransform* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
