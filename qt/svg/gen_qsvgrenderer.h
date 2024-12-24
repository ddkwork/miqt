#pragma once
#ifndef MIQT_QT_SVG_GEN_QSVGRENDERER_H
#define MIQT_QT_SVG_GEN_QSVGRENDERER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChildEvent;
class QEvent;
class QMetaMethod;
class QMetaObject;
class QObject;
class QPainter;
class QRect;
class QRectF;
class QSize;
class QSvgRenderer;
class QTimerEvent;
class QTransform;
class QXmlStreamReader;
class _GUID;
class type_info;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPainter QPainter;
typedef struct QRect QRect;
typedef struct QRectF QRectF;
typedef struct QSize QSize;
typedef struct QSvgRenderer QSvgRenderer;
typedef struct QTimerEvent QTimerEvent;
typedef struct QTransform QTransform;
typedef struct QXmlStreamReader QXmlStreamReader;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSvgRenderer* QSvgRenderer_new();
extern __declspec(dllexport) QSvgRenderer* QSvgRenderer_new2(struct miqt_string filename);
extern __declspec(dllexport) QSvgRenderer* QSvgRenderer_new3(struct miqt_string contents);
extern __declspec(dllexport) QSvgRenderer* QSvgRenderer_new4(QXmlStreamReader* contents);
extern __declspec(dllexport) QSvgRenderer* QSvgRenderer_new5(QObject* parent);
extern __declspec(dllexport) QSvgRenderer* QSvgRenderer_new6(struct miqt_string filename, QObject* parent);
extern __declspec(dllexport) QSvgRenderer* QSvgRenderer_new7(struct miqt_string contents, QObject* parent);
extern __declspec(dllexport) QSvgRenderer* QSvgRenderer_new8(QXmlStreamReader* contents, QObject* parent);
extern __declspec(dllexport) void QSvgRenderer_virtbase(QSvgRenderer* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QSvgRenderer_MetaObject(const QSvgRenderer* self);
extern __declspec(dllexport) void* QSvgRenderer_Metacast(QSvgRenderer* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QSvgRenderer_Tr(const char* s);
extern __declspec(dllexport) bool QSvgRenderer_IsValid(const QSvgRenderer* self);
extern __declspec(dllexport) QSize* QSvgRenderer_DefaultSize(const QSvgRenderer* self);
extern __declspec(dllexport) QRect* QSvgRenderer_ViewBox(const QSvgRenderer* self);
extern __declspec(dllexport) QRectF* QSvgRenderer_ViewBoxF(const QSvgRenderer* self);
extern __declspec(dllexport) void QSvgRenderer_SetViewBox(QSvgRenderer* self, QRect* viewbox);
extern __declspec(dllexport) void QSvgRenderer_SetViewBoxWithViewbox(QSvgRenderer* self, QRectF* viewbox);
extern __declspec(dllexport) int QSvgRenderer_AspectRatioMode(const QSvgRenderer* self);
extern __declspec(dllexport) void QSvgRenderer_SetAspectRatioMode(QSvgRenderer* self, int mode);
extern __declspec(dllexport) int QSvgRenderer_Options(const QSvgRenderer* self);
extern __declspec(dllexport) void QSvgRenderer_SetOptions(QSvgRenderer* self, int flags);
extern __declspec(dllexport) bool QSvgRenderer_Animated(const QSvgRenderer* self);
extern __declspec(dllexport) int QSvgRenderer_FramesPerSecond(const QSvgRenderer* self);
extern __declspec(dllexport) void QSvgRenderer_SetFramesPerSecond(QSvgRenderer* self, int num);
extern __declspec(dllexport) int QSvgRenderer_CurrentFrame(const QSvgRenderer* self);
extern __declspec(dllexport) void QSvgRenderer_SetCurrentFrame(QSvgRenderer* self, int currentFrame);
extern __declspec(dllexport) int QSvgRenderer_AnimationDuration(const QSvgRenderer* self);
extern __declspec(dllexport) bool QSvgRenderer_IsAnimationEnabled(const QSvgRenderer* self);
extern __declspec(dllexport) void QSvgRenderer_SetAnimationEnabled(QSvgRenderer* self, bool enable);
extern __declspec(dllexport) QRectF* QSvgRenderer_BoundsOnElement(const QSvgRenderer* self, struct miqt_string id);
extern __declspec(dllexport) bool QSvgRenderer_ElementExists(const QSvgRenderer* self, struct miqt_string id);
extern __declspec(dllexport) QTransform* QSvgRenderer_TransformForElement(const QSvgRenderer* self, struct miqt_string id);
extern __declspec(dllexport) void QSvgRenderer_SetDefaultOptions(int flags);
extern __declspec(dllexport) bool QSvgRenderer_Load(QSvgRenderer* self, struct miqt_string filename);
extern __declspec(dllexport) bool QSvgRenderer_LoadWithContents(QSvgRenderer* self, struct miqt_string contents);
extern __declspec(dllexport) bool QSvgRenderer_Load2(QSvgRenderer* self, QXmlStreamReader* contents);
extern __declspec(dllexport) void QSvgRenderer_Render(QSvgRenderer* self, QPainter* p);
extern __declspec(dllexport) void QSvgRenderer_Render2(QSvgRenderer* self, QPainter* p, QRectF* bounds);
extern __declspec(dllexport) void QSvgRenderer_Render3(QSvgRenderer* self, QPainter* p, struct miqt_string elementId);
extern __declspec(dllexport) void QSvgRenderer_RepaintNeeded(QSvgRenderer* self);
void QSvgRenderer_connect_RepaintNeeded(QSvgRenderer* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QSvgRenderer_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QSvgRenderer_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QSvgRenderer_Render32(QSvgRenderer* self, QPainter* p, struct miqt_string elementId, QRectF* bounds);
extern __declspec(dllexport) void QSvgRenderer_override_virtual_Event(void* self, intptr_t slot);
bool QSvgRenderer_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QSvgRenderer_override_virtual_EventFilter(void* self, intptr_t slot);
bool QSvgRenderer_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QSvgRenderer_override_virtual_TimerEvent(void* self, intptr_t slot);
void QSvgRenderer_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QSvgRenderer_override_virtual_ChildEvent(void* self, intptr_t slot);
void QSvgRenderer_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QSvgRenderer_override_virtual_CustomEvent(void* self, intptr_t slot);
void QSvgRenderer_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QSvgRenderer_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QSvgRenderer_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QSvgRenderer_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QSvgRenderer_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QSvgRenderer_Delete(QSvgRenderer* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
