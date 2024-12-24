#pragma once
#ifndef MIQT_QT_GEN_QSURFACEFORMAT_H
#define MIQT_QT_GEN_QSURFACEFORMAT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QColorSpace;
class QSurfaceFormat;
class _GUID;
class type_info;
#else
typedef struct QColorSpace QColorSpace;
typedef struct QSurfaceFormat QSurfaceFormat;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSurfaceFormat* QSurfaceFormat_new();
extern __declspec(dllexport) QSurfaceFormat* QSurfaceFormat_new2(FormatOptions options);
extern __declspec(dllexport) QSurfaceFormat* QSurfaceFormat_new3(QSurfaceFormat* other);
extern __declspec(dllexport) void QSurfaceFormat_OperatorAssign(QSurfaceFormat* self, QSurfaceFormat* other);
extern __declspec(dllexport) void QSurfaceFormat_SetDepthBufferSize(QSurfaceFormat* self, int size);
extern __declspec(dllexport) int QSurfaceFormat_DepthBufferSize(const QSurfaceFormat* self);
extern __declspec(dllexport) void QSurfaceFormat_SetStencilBufferSize(QSurfaceFormat* self, int size);
extern __declspec(dllexport) int QSurfaceFormat_StencilBufferSize(const QSurfaceFormat* self);
extern __declspec(dllexport) void QSurfaceFormat_SetRedBufferSize(QSurfaceFormat* self, int size);
extern __declspec(dllexport) int QSurfaceFormat_RedBufferSize(const QSurfaceFormat* self);
extern __declspec(dllexport) void QSurfaceFormat_SetGreenBufferSize(QSurfaceFormat* self, int size);
extern __declspec(dllexport) int QSurfaceFormat_GreenBufferSize(const QSurfaceFormat* self);
extern __declspec(dllexport) void QSurfaceFormat_SetBlueBufferSize(QSurfaceFormat* self, int size);
extern __declspec(dllexport) int QSurfaceFormat_BlueBufferSize(const QSurfaceFormat* self);
extern __declspec(dllexport) void QSurfaceFormat_SetAlphaBufferSize(QSurfaceFormat* self, int size);
extern __declspec(dllexport) int QSurfaceFormat_AlphaBufferSize(const QSurfaceFormat* self);
extern __declspec(dllexport) void QSurfaceFormat_SetSamples(QSurfaceFormat* self, int numSamples);
extern __declspec(dllexport) int QSurfaceFormat_Samples(const QSurfaceFormat* self);
extern __declspec(dllexport) void QSurfaceFormat_SetSwapBehavior(QSurfaceFormat* self, SwapBehavior behavior);
extern __declspec(dllexport) SwapBehavior QSurfaceFormat_SwapBehavior(const QSurfaceFormat* self);
extern __declspec(dllexport) bool QSurfaceFormat_HasAlpha(const QSurfaceFormat* self);
extern __declspec(dllexport) void QSurfaceFormat_SetProfile(QSurfaceFormat* self, OpenGLContextProfile profile);
extern __declspec(dllexport) OpenGLContextProfile QSurfaceFormat_Profile(const QSurfaceFormat* self);
extern __declspec(dllexport) void QSurfaceFormat_SetRenderableType(QSurfaceFormat* self, RenderableType typeVal);
extern __declspec(dllexport) RenderableType QSurfaceFormat_RenderableType(const QSurfaceFormat* self);
extern __declspec(dllexport) void QSurfaceFormat_SetMajorVersion(QSurfaceFormat* self, int majorVersion);
extern __declspec(dllexport) int QSurfaceFormat_MajorVersion(const QSurfaceFormat* self);
extern __declspec(dllexport) void QSurfaceFormat_SetMinorVersion(QSurfaceFormat* self, int minorVersion);
extern __declspec(dllexport) int QSurfaceFormat_MinorVersion(const QSurfaceFormat* self);
extern __declspec(dllexport) struct miqt_map /* tuple of int and int */  QSurfaceFormat_Version(const QSurfaceFormat* self);
extern __declspec(dllexport) void QSurfaceFormat_SetVersion(QSurfaceFormat* self, int major, int minor);
extern __declspec(dllexport) bool QSurfaceFormat_Stereo(const QSurfaceFormat* self);
extern __declspec(dllexport) void QSurfaceFormat_SetStereo(QSurfaceFormat* self, bool enable);
extern __declspec(dllexport) void QSurfaceFormat_SetOptions(QSurfaceFormat* self, int options);
extern __declspec(dllexport) void QSurfaceFormat_SetOption(QSurfaceFormat* self, FormatOption option);
extern __declspec(dllexport) bool QSurfaceFormat_TestOption(const QSurfaceFormat* self, FormatOption option);
extern __declspec(dllexport) int QSurfaceFormat_Options(const QSurfaceFormat* self);
extern __declspec(dllexport) int QSurfaceFormat_SwapInterval(const QSurfaceFormat* self);
extern __declspec(dllexport) void QSurfaceFormat_SetSwapInterval(QSurfaceFormat* self, int interval);
extern __declspec(dllexport) QColorSpace* QSurfaceFormat_ColorSpace(const QSurfaceFormat* self);
extern __declspec(dllexport) void QSurfaceFormat_SetColorSpace(QSurfaceFormat* self, QColorSpace* colorSpace);
extern __declspec(dllexport) void QSurfaceFormat_SetColorSpaceWithColorSpace(QSurfaceFormat* self, ColorSpace colorSpace);
extern __declspec(dllexport) void QSurfaceFormat_SetDefaultFormat(QSurfaceFormat* format);
extern __declspec(dllexport) QSurfaceFormat* QSurfaceFormat_DefaultFormat();
extern __declspec(dllexport) void QSurfaceFormat_SetOption2(QSurfaceFormat* self, FormatOption option, bool on);
extern __declspec(dllexport) void QSurfaceFormat_Delete(QSurfaceFormat* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
