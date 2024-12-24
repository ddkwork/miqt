#pragma once
#ifndef MIQT_QT_GEN_QOFFSCREENSURFACE_H
#define MIQT_QT_GEN_QOFFSCREENSURFACE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChildEvent;
class QEvent;
class QMetaMethod;
class QMetaObject;
class QObject;
class QOffscreenSurface;
class QScreen;
class QSize;
class QSurface;
class QSurfaceFormat;
class QTimerEvent;
class _GUID;
class type_info;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QOffscreenSurface QOffscreenSurface;
typedef struct QScreen QScreen;
typedef struct QSize QSize;
typedef struct QSurface QSurface;
typedef struct QSurfaceFormat QSurfaceFormat;
typedef struct QTimerEvent QTimerEvent;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QOffscreenSurface* QOffscreenSurface_new();
extern __declspec(dllexport) QOffscreenSurface* QOffscreenSurface_new2(QScreen* screen);
extern __declspec(dllexport) QOffscreenSurface* QOffscreenSurface_new3(QScreen* screen, QObject* parent);
extern __declspec(dllexport) void QOffscreenSurface_virtbase(QOffscreenSurface* src, QObject** outptr_QObject, QSurface** outptr_QSurface);
extern __declspec(dllexport) QMetaObject* QOffscreenSurface_MetaObject(const QOffscreenSurface* self);
extern __declspec(dllexport) void* QOffscreenSurface_Metacast(QOffscreenSurface* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QOffscreenSurface_Tr(const char* s);
extern __declspec(dllexport) SurfaceType QOffscreenSurface_SurfaceType(const QOffscreenSurface* self);
extern __declspec(dllexport) void QOffscreenSurface_Create(QOffscreenSurface* self);
extern __declspec(dllexport) void QOffscreenSurface_Destroy(QOffscreenSurface* self);
extern __declspec(dllexport) bool QOffscreenSurface_IsValid(const QOffscreenSurface* self);
extern __declspec(dllexport) void QOffscreenSurface_SetFormat(QOffscreenSurface* self, QSurfaceFormat* format);
extern __declspec(dllexport) QSurfaceFormat* QOffscreenSurface_Format(const QOffscreenSurface* self);
extern __declspec(dllexport) QSurfaceFormat* QOffscreenSurface_RequestedFormat(const QOffscreenSurface* self);
extern __declspec(dllexport) QSize* QOffscreenSurface_Size(const QOffscreenSurface* self);
extern __declspec(dllexport) QScreen* QOffscreenSurface_Screen(const QOffscreenSurface* self);
extern __declspec(dllexport) void QOffscreenSurface_SetScreen(QOffscreenSurface* self, QScreen* screen);
extern __declspec(dllexport) void QOffscreenSurface_ScreenChanged(QOffscreenSurface* self, QScreen* screen);
void QOffscreenSurface_connect_ScreenChanged(QOffscreenSurface* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QOffscreenSurface_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QOffscreenSurface_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QOffscreenSurface_override_virtual_SurfaceType(void* self, intptr_t slot);
SurfaceType QOffscreenSurface_virtualbase_SurfaceType(const void* self);
extern __declspec(dllexport) void QOffscreenSurface_override_virtual_Format(void* self, intptr_t slot);
QSurfaceFormat* QOffscreenSurface_virtualbase_Format(const void* self);
extern __declspec(dllexport) void QOffscreenSurface_override_virtual_Size(void* self, intptr_t slot);
QSize* QOffscreenSurface_virtualbase_Size(const void* self);
extern __declspec(dllexport) void QOffscreenSurface_override_virtual_Event(void* self, intptr_t slot);
bool QOffscreenSurface_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QOffscreenSurface_override_virtual_EventFilter(void* self, intptr_t slot);
bool QOffscreenSurface_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QOffscreenSurface_override_virtual_TimerEvent(void* self, intptr_t slot);
void QOffscreenSurface_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QOffscreenSurface_override_virtual_ChildEvent(void* self, intptr_t slot);
void QOffscreenSurface_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QOffscreenSurface_override_virtual_CustomEvent(void* self, intptr_t slot);
void QOffscreenSurface_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QOffscreenSurface_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QOffscreenSurface_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QOffscreenSurface_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QOffscreenSurface_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QOffscreenSurface_Delete(QOffscreenSurface* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
