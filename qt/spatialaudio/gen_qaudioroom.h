#pragma once
#ifndef MIQT_QT_SPATIALAUDIO_GEN_QAUDIOROOM_H
#define MIQT_QT_SPATIALAUDIO_GEN_QAUDIOROOM_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioEngine QAudioEngine;
typedef struct QAudioRoom QAudioRoom;
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QQuaternion QQuaternion;
typedef struct QTimerEvent QTimerEvent;
typedef struct QVector3D QVector3D;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QAudioRoom* QAudioRoom_new(QAudioEngine* engine);
extern __declspec(dllexport) void QAudioRoom_virtbase(QAudioRoom* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QAudioRoom_MetaObject(const QAudioRoom* self);
extern __declspec(dllexport) void* QAudioRoom_Metacast(QAudioRoom* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QAudioRoom_Tr(const char* s);
extern __declspec(dllexport) void QAudioRoom_SetPosition(QAudioRoom* self, QVector3D* pos);
extern __declspec(dllexport) QVector3D* QAudioRoom_Position(const QAudioRoom* self);
extern __declspec(dllexport) void QAudioRoom_SetDimensions(QAudioRoom* self, QVector3D* dim);
extern __declspec(dllexport) QVector3D* QAudioRoom_Dimensions(const QAudioRoom* self);
extern __declspec(dllexport) void QAudioRoom_SetRotation(QAudioRoom* self, QQuaternion* q);
extern __declspec(dllexport) QQuaternion* QAudioRoom_Rotation(const QAudioRoom* self);
extern __declspec(dllexport) void QAudioRoom_SetWallMaterial(QAudioRoom* self, Wall wall, Material material);
extern __declspec(dllexport) Material QAudioRoom_WallMaterial(const QAudioRoom* self, Wall wall);
extern __declspec(dllexport) void QAudioRoom_SetReflectionGain(QAudioRoom* self, float factor);
extern __declspec(dllexport) float QAudioRoom_ReflectionGain(const QAudioRoom* self);
extern __declspec(dllexport) void QAudioRoom_SetReverbGain(QAudioRoom* self, float factor);
extern __declspec(dllexport) float QAudioRoom_ReverbGain(const QAudioRoom* self);
extern __declspec(dllexport) void QAudioRoom_SetReverbTime(QAudioRoom* self, float factor);
extern __declspec(dllexport) float QAudioRoom_ReverbTime(const QAudioRoom* self);
extern __declspec(dllexport) void QAudioRoom_SetReverbBrightness(QAudioRoom* self, float factor);
extern __declspec(dllexport) float QAudioRoom_ReverbBrightness(const QAudioRoom* self);
extern __declspec(dllexport) void QAudioRoom_PositionChanged(QAudioRoom* self);
void QAudioRoom_connect_PositionChanged(QAudioRoom* self, intptr_t slot);
extern __declspec(dllexport) void QAudioRoom_DimensionsChanged(QAudioRoom* self);
void QAudioRoom_connect_DimensionsChanged(QAudioRoom* self, intptr_t slot);
extern __declspec(dllexport) void QAudioRoom_RotationChanged(QAudioRoom* self);
void QAudioRoom_connect_RotationChanged(QAudioRoom* self, intptr_t slot);
extern __declspec(dllexport) void QAudioRoom_WallsChanged(QAudioRoom* self);
void QAudioRoom_connect_WallsChanged(QAudioRoom* self, intptr_t slot);
extern __declspec(dllexport) void QAudioRoom_ReflectionGainChanged(QAudioRoom* self);
void QAudioRoom_connect_ReflectionGainChanged(QAudioRoom* self, intptr_t slot);
extern __declspec(dllexport) void QAudioRoom_ReverbGainChanged(QAudioRoom* self);
void QAudioRoom_connect_ReverbGainChanged(QAudioRoom* self, intptr_t slot);
extern __declspec(dllexport) void QAudioRoom_ReverbTimeChanged(QAudioRoom* self);
void QAudioRoom_connect_ReverbTimeChanged(QAudioRoom* self, intptr_t slot);
extern __declspec(dllexport) void QAudioRoom_ReverbBrightnessChanged(QAudioRoom* self);
void QAudioRoom_connect_ReverbBrightnessChanged(QAudioRoom* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QAudioRoom_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QAudioRoom_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QAudioRoom_override_virtual_Event(void* self, intptr_t slot);
bool QAudioRoom_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QAudioRoom_override_virtual_EventFilter(void* self, intptr_t slot);
bool QAudioRoom_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QAudioRoom_override_virtual_TimerEvent(void* self, intptr_t slot);
void QAudioRoom_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QAudioRoom_override_virtual_ChildEvent(void* self, intptr_t slot);
void QAudioRoom_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QAudioRoom_override_virtual_CustomEvent(void* self, intptr_t slot);
void QAudioRoom_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QAudioRoom_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QAudioRoom_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QAudioRoom_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QAudioRoom_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QAudioRoom_Delete(QAudioRoom* self, bool isSubclass);

} 
