#pragma once
#ifndef MIQT_QT_GEN_QSURFACE_H
#define MIQT_QT_GEN_QSURFACE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QSize;
class QSurface;
class QSurfaceFormat;
class _GUID;
class type_info;
#else
typedef struct QSize QSize;
typedef struct QSurface QSurface;
typedef struct QSurfaceFormat QSurfaceFormat;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) SurfaceClass QSurface_SurfaceClass(const QSurface* self);
extern __declspec(dllexport) QSurfaceFormat* QSurface_Format(const QSurface* self);
extern __declspec(dllexport) SurfaceType QSurface_SurfaceType(const QSurface* self);
extern __declspec(dllexport) bool QSurface_SupportsOpenGL(const QSurface* self);
extern __declspec(dllexport) QSize* QSurface_Size(const QSurface* self);
extern __declspec(dllexport) void QSurface_Delete(QSurface* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
