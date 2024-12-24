#pragma once
#ifndef MIQT_QT_SPATIALAUDIO_GEN_QSPATIALSOUND_H
#define MIQT_QT_SPATIALAUDIO_GEN_QSPATIALSOUND_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioEngine QAudioEngine;
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QQuaternion QQuaternion;
typedef struct QSpatialSound QSpatialSound;
typedef struct QTimerEvent QTimerEvent;
typedef struct QUrl QUrl;
typedef struct QVector3D QVector3D;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSpatialSound* QSpatialSound_new(QAudioEngine* engine);
extern __declspec(dllexport) void QSpatialSound_virtbase(QSpatialSound* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QSpatialSound_MetaObject(const QSpatialSound* self);
extern __declspec(dllexport) void* QSpatialSound_Metacast(QSpatialSound* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QSpatialSound_Tr(const char* s);
extern __declspec(dllexport) void QSpatialSound_SetSource(QSpatialSound* self, QUrl* url);
extern __declspec(dllexport) QUrl* QSpatialSound_Source(const QSpatialSound* self);
extern __declspec(dllexport) int QSpatialSound_Loops(const QSpatialSound* self);
extern __declspec(dllexport) void QSpatialSound_SetLoops(QSpatialSound* self, int loops);
extern __declspec(dllexport) bool QSpatialSound_AutoPlay(const QSpatialSound* self);
extern __declspec(dllexport) void QSpatialSound_SetAutoPlay(QSpatialSound* self, bool autoPlay);
extern __declspec(dllexport) void QSpatialSound_SetPosition(QSpatialSound* self, QVector3D* pos);
extern __declspec(dllexport) QVector3D* QSpatialSound_Position(const QSpatialSound* self);
extern __declspec(dllexport) void QSpatialSound_SetRotation(QSpatialSound* self, QQuaternion* q);
extern __declspec(dllexport) QQuaternion* QSpatialSound_Rotation(const QSpatialSound* self);
extern __declspec(dllexport) void QSpatialSound_SetVolume(QSpatialSound* self, float volume);
extern __declspec(dllexport) float QSpatialSound_Volume(const QSpatialSound* self);
extern __declspec(dllexport) void QSpatialSound_SetDistanceModel(QSpatialSound* self, DistanceModel model);
extern __declspec(dllexport) DistanceModel QSpatialSound_DistanceModel(const QSpatialSound* self);
extern __declspec(dllexport) void QSpatialSound_SetSize(QSpatialSound* self, float size);
extern __declspec(dllexport) float QSpatialSound_Size(const QSpatialSound* self);
extern __declspec(dllexport) void QSpatialSound_SetDistanceCutoff(QSpatialSound* self, float cutoff);
extern __declspec(dllexport) float QSpatialSound_DistanceCutoff(const QSpatialSound* self);
extern __declspec(dllexport) void QSpatialSound_SetManualAttenuation(QSpatialSound* self, float attenuation);
extern __declspec(dllexport) float QSpatialSound_ManualAttenuation(const QSpatialSound* self);
extern __declspec(dllexport) void QSpatialSound_SetOcclusionIntensity(QSpatialSound* self, float occlusion);
extern __declspec(dllexport) float QSpatialSound_OcclusionIntensity(const QSpatialSound* self);
extern __declspec(dllexport) void QSpatialSound_SetDirectivity(QSpatialSound* self, float alpha);
extern __declspec(dllexport) float QSpatialSound_Directivity(const QSpatialSound* self);
extern __declspec(dllexport) void QSpatialSound_SetDirectivityOrder(QSpatialSound* self, float alpha);
extern __declspec(dllexport) float QSpatialSound_DirectivityOrder(const QSpatialSound* self);
extern __declspec(dllexport) void QSpatialSound_SetNearFieldGain(QSpatialSound* self, float gain);
extern __declspec(dllexport) float QSpatialSound_NearFieldGain(const QSpatialSound* self);
extern __declspec(dllexport) QAudioEngine* QSpatialSound_Engine(const QSpatialSound* self);
extern __declspec(dllexport) void QSpatialSound_SourceChanged(QSpatialSound* self);
void QSpatialSound_connect_SourceChanged(QSpatialSound* self, intptr_t slot);
extern __declspec(dllexport) void QSpatialSound_LoopsChanged(QSpatialSound* self);
void QSpatialSound_connect_LoopsChanged(QSpatialSound* self, intptr_t slot);
extern __declspec(dllexport) void QSpatialSound_AutoPlayChanged(QSpatialSound* self);
void QSpatialSound_connect_AutoPlayChanged(QSpatialSound* self, intptr_t slot);
extern __declspec(dllexport) void QSpatialSound_PositionChanged(QSpatialSound* self);
void QSpatialSound_connect_PositionChanged(QSpatialSound* self, intptr_t slot);
extern __declspec(dllexport) void QSpatialSound_RotationChanged(QSpatialSound* self);
void QSpatialSound_connect_RotationChanged(QSpatialSound* self, intptr_t slot);
extern __declspec(dllexport) void QSpatialSound_VolumeChanged(QSpatialSound* self);
void QSpatialSound_connect_VolumeChanged(QSpatialSound* self, intptr_t slot);
extern __declspec(dllexport) void QSpatialSound_DistanceModelChanged(QSpatialSound* self);
void QSpatialSound_connect_DistanceModelChanged(QSpatialSound* self, intptr_t slot);
extern __declspec(dllexport) void QSpatialSound_SizeChanged(QSpatialSound* self);
void QSpatialSound_connect_SizeChanged(QSpatialSound* self, intptr_t slot);
extern __declspec(dllexport) void QSpatialSound_DistanceCutoffChanged(QSpatialSound* self);
void QSpatialSound_connect_DistanceCutoffChanged(QSpatialSound* self, intptr_t slot);
extern __declspec(dllexport) void QSpatialSound_ManualAttenuationChanged(QSpatialSound* self);
void QSpatialSound_connect_ManualAttenuationChanged(QSpatialSound* self, intptr_t slot);
extern __declspec(dllexport) void QSpatialSound_OcclusionIntensityChanged(QSpatialSound* self);
void QSpatialSound_connect_OcclusionIntensityChanged(QSpatialSound* self, intptr_t slot);
extern __declspec(dllexport) void QSpatialSound_DirectivityChanged(QSpatialSound* self);
void QSpatialSound_connect_DirectivityChanged(QSpatialSound* self, intptr_t slot);
extern __declspec(dllexport) void QSpatialSound_DirectivityOrderChanged(QSpatialSound* self);
void QSpatialSound_connect_DirectivityOrderChanged(QSpatialSound* self, intptr_t slot);
extern __declspec(dllexport) void QSpatialSound_NearFieldGainChanged(QSpatialSound* self);
void QSpatialSound_connect_NearFieldGainChanged(QSpatialSound* self, intptr_t slot);
extern __declspec(dllexport) void QSpatialSound_Play(QSpatialSound* self);
extern __declspec(dllexport) void QSpatialSound_Pause(QSpatialSound* self);
extern __declspec(dllexport) void QSpatialSound_Stop(QSpatialSound* self);
extern __declspec(dllexport) struct miqt_string QSpatialSound_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QSpatialSound_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QSpatialSound_override_virtual_Event(void* self, intptr_t slot);
bool QSpatialSound_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QSpatialSound_override_virtual_EventFilter(void* self, intptr_t slot);
bool QSpatialSound_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QSpatialSound_override_virtual_TimerEvent(void* self, intptr_t slot);
void QSpatialSound_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QSpatialSound_override_virtual_ChildEvent(void* self, intptr_t slot);
void QSpatialSound_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QSpatialSound_override_virtual_CustomEvent(void* self, intptr_t slot);
void QSpatialSound_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QSpatialSound_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QSpatialSound_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QSpatialSound_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QSpatialSound_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QSpatialSound_Delete(QSpatialSound* self, bool isSubclass);

} 
