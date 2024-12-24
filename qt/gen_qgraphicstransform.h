#pragma once
#ifndef MIQT_QT_GEN_QGRAPHICSTRANSFORM_H
#define MIQT_QT_GEN_QGRAPHICSTRANSFORM_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QGraphicsRotation QGraphicsRotation;
typedef struct QGraphicsScale QGraphicsScale;
typedef struct QGraphicsTransform QGraphicsTransform;
typedef struct QMatrix4x4 QMatrix4x4;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;
typedef struct QVector3D QVector3D;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsTransform* QGraphicsTransform_new();
extern __declspec(dllexport) QGraphicsTransform* QGraphicsTransform_new2(QObject* parent);
extern __declspec(dllexport) void QGraphicsTransform_virtbase(QGraphicsTransform* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QGraphicsTransform_MetaObject(const QGraphicsTransform* self);
extern __declspec(dllexport) void* QGraphicsTransform_Metacast(QGraphicsTransform* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QGraphicsTransform_Tr(const char* s);
extern __declspec(dllexport) void QGraphicsTransform_ApplyTo(const QGraphicsTransform* self, QMatrix4x4* matrix);
extern __declspec(dllexport) struct miqt_string QGraphicsTransform_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QGraphicsTransform_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QGraphicsTransform_override_virtual_ApplyTo(void* self, intptr_t slot);
void QGraphicsTransform_virtualbase_ApplyTo(const void* self, QMatrix4x4* matrix);
extern __declspec(dllexport) void QGraphicsTransform_override_virtual_Event(void* self, intptr_t slot);
bool QGraphicsTransform_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QGraphicsTransform_override_virtual_EventFilter(void* self, intptr_t slot);
bool QGraphicsTransform_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QGraphicsTransform_override_virtual_TimerEvent(void* self, intptr_t slot);
void QGraphicsTransform_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QGraphicsTransform_override_virtual_ChildEvent(void* self, intptr_t slot);
void QGraphicsTransform_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QGraphicsTransform_override_virtual_CustomEvent(void* self, intptr_t slot);
void QGraphicsTransform_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QGraphicsTransform_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QGraphicsTransform_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QGraphicsTransform_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QGraphicsTransform_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QGraphicsTransform_Delete(QGraphicsTransform* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsScale* QGraphicsScale_new();
extern __declspec(dllexport) QGraphicsScale* QGraphicsScale_new2(QObject* parent);
extern __declspec(dllexport) void QGraphicsScale_virtbase(QGraphicsScale* src, QGraphicsTransform** outptr_QGraphicsTransform);
extern __declspec(dllexport) QMetaObject* QGraphicsScale_MetaObject(const QGraphicsScale* self);
extern __declspec(dllexport) void* QGraphicsScale_Metacast(QGraphicsScale* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QGraphicsScale_Tr(const char* s);
extern __declspec(dllexport) QVector3D* QGraphicsScale_Origin(const QGraphicsScale* self);
extern __declspec(dllexport) void QGraphicsScale_SetOrigin(QGraphicsScale* self, QVector3D* point);
extern __declspec(dllexport) double QGraphicsScale_XScale(const QGraphicsScale* self);
extern __declspec(dllexport) void QGraphicsScale_SetXScale(QGraphicsScale* self, double xScale);
extern __declspec(dllexport) double QGraphicsScale_YScale(const QGraphicsScale* self);
extern __declspec(dllexport) void QGraphicsScale_SetYScale(QGraphicsScale* self, double yScale);
extern __declspec(dllexport) double QGraphicsScale_ZScale(const QGraphicsScale* self);
extern __declspec(dllexport) void QGraphicsScale_SetZScale(QGraphicsScale* self, double zScale);
extern __declspec(dllexport) void QGraphicsScale_ApplyTo(const QGraphicsScale* self, QMatrix4x4* matrix);
extern __declspec(dllexport) void QGraphicsScale_OriginChanged(QGraphicsScale* self);
void QGraphicsScale_connect_OriginChanged(QGraphicsScale* self, intptr_t slot);
extern __declspec(dllexport) void QGraphicsScale_XScaleChanged(QGraphicsScale* self);
void QGraphicsScale_connect_XScaleChanged(QGraphicsScale* self, intptr_t slot);
extern __declspec(dllexport) void QGraphicsScale_YScaleChanged(QGraphicsScale* self);
void QGraphicsScale_connect_YScaleChanged(QGraphicsScale* self, intptr_t slot);
extern __declspec(dllexport) void QGraphicsScale_ZScaleChanged(QGraphicsScale* self);
void QGraphicsScale_connect_ZScaleChanged(QGraphicsScale* self, intptr_t slot);
extern __declspec(dllexport) void QGraphicsScale_ScaleChanged(QGraphicsScale* self);
void QGraphicsScale_connect_ScaleChanged(QGraphicsScale* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QGraphicsScale_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QGraphicsScale_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QGraphicsScale_override_virtual_ApplyTo(void* self, intptr_t slot);
void QGraphicsScale_virtualbase_ApplyTo(const void* self, QMatrix4x4* matrix);
extern __declspec(dllexport) void QGraphicsScale_Delete(QGraphicsScale* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsRotation* QGraphicsRotation_new();
extern __declspec(dllexport) QGraphicsRotation* QGraphicsRotation_new2(QObject* parent);
extern __declspec(dllexport) void QGraphicsRotation_virtbase(QGraphicsRotation* src, QGraphicsTransform** outptr_QGraphicsTransform);
extern __declspec(dllexport) QMetaObject* QGraphicsRotation_MetaObject(const QGraphicsRotation* self);
extern __declspec(dllexport) void* QGraphicsRotation_Metacast(QGraphicsRotation* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QGraphicsRotation_Tr(const char* s);
extern __declspec(dllexport) QVector3D* QGraphicsRotation_Origin(const QGraphicsRotation* self);
extern __declspec(dllexport) void QGraphicsRotation_SetOrigin(QGraphicsRotation* self, QVector3D* point);
extern __declspec(dllexport) double QGraphicsRotation_Angle(const QGraphicsRotation* self);
extern __declspec(dllexport) void QGraphicsRotation_SetAngle(QGraphicsRotation* self, double angle);
extern __declspec(dllexport) QVector3D* QGraphicsRotation_Axis(const QGraphicsRotation* self);
extern __declspec(dllexport) void QGraphicsRotation_SetAxis(QGraphicsRotation* self, QVector3D* axis);
extern __declspec(dllexport) void QGraphicsRotation_SetAxisWithAxis(QGraphicsRotation* self, int axis);
extern __declspec(dllexport) void QGraphicsRotation_ApplyTo(const QGraphicsRotation* self, QMatrix4x4* matrix);
extern __declspec(dllexport) void QGraphicsRotation_OriginChanged(QGraphicsRotation* self);
void QGraphicsRotation_connect_OriginChanged(QGraphicsRotation* self, intptr_t slot);
extern __declspec(dllexport) void QGraphicsRotation_AngleChanged(QGraphicsRotation* self);
void QGraphicsRotation_connect_AngleChanged(QGraphicsRotation* self, intptr_t slot);
extern __declspec(dllexport) void QGraphicsRotation_AxisChanged(QGraphicsRotation* self);
void QGraphicsRotation_connect_AxisChanged(QGraphicsRotation* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QGraphicsRotation_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QGraphicsRotation_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QGraphicsRotation_override_virtual_ApplyTo(void* self, intptr_t slot);
void QGraphicsRotation_virtualbase_ApplyTo(const void* self, QMatrix4x4* matrix);
extern __declspec(dllexport) void QGraphicsRotation_Delete(QGraphicsRotation* self, bool isSubclass);

} 
