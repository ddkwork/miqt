#pragma once
#ifndef MIQT_QT_GEN_QEVENTPOINT_H
#define MIQT_QT_GEN_QEVENTPOINT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEventPoint QEventPoint;
typedef struct QPointF QPointF;
typedef struct QPointingDevice QPointingDevice;
typedef struct QPointingDeviceUniqueId QPointingDeviceUniqueId;
typedef struct QSizeF QSizeF;
typedef struct QVector2D QVector2D;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QEventPoint* QEventPoint_new();
extern __declspec(dllexport) QEventPoint* QEventPoint_new2(int pointId, State state, QPointF* scenePosition, QPointF* globalPosition);
extern __declspec(dllexport) QEventPoint* QEventPoint_new3(QEventPoint* other);
extern __declspec(dllexport) QEventPoint* QEventPoint_new4(int id);
extern __declspec(dllexport) QEventPoint* QEventPoint_new5(int id, QPointingDevice* device);
extern __declspec(dllexport) void QEventPoint_OperatorAssign(QEventPoint* self, QEventPoint* other);
extern __declspec(dllexport) bool QEventPoint_OperatorEqual(const QEventPoint* self, QEventPoint* other);
extern __declspec(dllexport) bool QEventPoint_OperatorNotEqual(const QEventPoint* self, QEventPoint* other);
extern __declspec(dllexport) void QEventPoint_Swap(QEventPoint* self, QEventPoint* other);
extern __declspec(dllexport) QPointF* QEventPoint_Position(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_PressPosition(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_GrabPosition(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_LastPosition(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_ScenePosition(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_ScenePressPosition(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_SceneGrabPosition(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_SceneLastPosition(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_GlobalPosition(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_GlobalPressPosition(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_GlobalGrabPosition(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_GlobalLastPosition(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_NormalizedPosition(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_Pos(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_StartPos(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_ScenePos(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_StartScenePos(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_ScreenPos(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_StartScreenPos(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_StartNormalizedPos(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_NormalizedPos(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_LastPos(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_LastScenePos(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_LastScreenPos(const QEventPoint* self);
extern __declspec(dllexport) QPointF* QEventPoint_LastNormalizedPos(const QEventPoint* self);
extern __declspec(dllexport) QVector2D* QEventPoint_Velocity(const QEventPoint* self);
extern __declspec(dllexport) State QEventPoint_State(const QEventPoint* self);
extern __declspec(dllexport) QPointingDevice* QEventPoint_Device(const QEventPoint* self);
extern __declspec(dllexport) int QEventPoint_Id(const QEventPoint* self);
extern __declspec(dllexport) QPointingDeviceUniqueId* QEventPoint_UniqueId(const QEventPoint* self);
extern __declspec(dllexport) unsigned long QEventPoint_Timestamp(const QEventPoint* self);
extern __declspec(dllexport) unsigned long QEventPoint_LastTimestamp(const QEventPoint* self);
extern __declspec(dllexport) unsigned long QEventPoint_PressTimestamp(const QEventPoint* self);
extern __declspec(dllexport) double QEventPoint_TimeHeld(const QEventPoint* self);
extern __declspec(dllexport) double QEventPoint_Pressure(const QEventPoint* self);
extern __declspec(dllexport) double QEventPoint_Rotation(const QEventPoint* self);
extern __declspec(dllexport) QSizeF* QEventPoint_EllipseDiameters(const QEventPoint* self);
extern __declspec(dllexport) bool QEventPoint_IsAccepted(const QEventPoint* self);
extern __declspec(dllexport) void QEventPoint_SetAccepted(QEventPoint* self);
extern __declspec(dllexport) void QEventPoint_SetAccepted1(QEventPoint* self, bool accepted);
extern __declspec(dllexport) void QEventPoint_Delete(QEventPoint* self, bool isSubclass);

} 
