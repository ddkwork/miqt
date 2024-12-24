#pragma once
#ifndef MIQT_QT_GEN_QQUATERNION_H
#define MIQT_QT_GEN_QQUATERNION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QQuaternion QQuaternion;
typedef struct QVector3D QVector3D;
typedef struct QVector4D QVector4D;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QQuaternion* QQuaternion_new();
extern __declspec(dllexport) 
QQuaternion* QQuaternion_new2(int param1);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_new3(float scalar, float xpos, float ypos, float zpos);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_new4(float scalar, QVector3D* vector);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_new5(QVector4D* vector);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_new6(QQuaternion* param1);
extern __declspec(dllexport) 
bool QQuaternion_IsNull(const QQuaternion* self);
extern __declspec(dllexport) 
bool QQuaternion_IsIdentity(const QQuaternion* self);
extern __declspec(dllexport) 
QVector3D* QQuaternion_Vector(const QQuaternion* self);
extern __declspec(dllexport) 
void QQuaternion_SetVector(QQuaternion* self, QVector3D* vector);
extern __declspec(dllexport) 
void QQuaternion_SetVector2(QQuaternion* self, float x, float y, float z);
extern __declspec(dllexport) 
float QQuaternion_X(const QQuaternion* self);
extern __declspec(dllexport) 
float QQuaternion_Y(const QQuaternion* self);
extern __declspec(dllexport) 
float QQuaternion_Z(const QQuaternion* self);
extern __declspec(dllexport) 
float QQuaternion_Scalar(const QQuaternion* self);
extern __declspec(dllexport) 
void QQuaternion_SetX(QQuaternion* self, float x);
extern __declspec(dllexport) 
void QQuaternion_SetY(QQuaternion* self, float y);
extern __declspec(dllexport) 
void QQuaternion_SetZ(QQuaternion* self, float z);
extern __declspec(dllexport) 
void QQuaternion_SetScalar(QQuaternion* self, float scalar);
extern __declspec(dllexport) 
float QQuaternion_DotProduct(QQuaternion* q1, QQuaternion* q2);
extern __declspec(dllexport) 
float QQuaternion_Length(const QQuaternion* self);
extern __declspec(dllexport) 
float QQuaternion_LengthSquared(const QQuaternion* self);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_Normalized(const QQuaternion* self);
extern __declspec(dllexport) 
void QQuaternion_Normalize(QQuaternion* self);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_Inverted(const QQuaternion* self);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_Conjugated(const QQuaternion* self);
extern __declspec(dllexport) 
QVector3D* QQuaternion_RotatedVector(const QQuaternion* self, QVector3D* vector);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_OperatorPlusAssign(QQuaternion* self, QQuaternion* quaternion);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_OperatorMinusAssign(QQuaternion* self, QQuaternion* quaternion);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_OperatorMultiplyAssign(QQuaternion* self, float factor);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_OperatorMultiplyAssignWithQuaternion(QQuaternion* self, QQuaternion* quaternion);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_OperatorDivideAssign(QQuaternion* self, float divisor);
extern __declspec(dllexport) 
QVector4D* QQuaternion_ToVector4D(const QQuaternion* self);
extern __declspec(dllexport) 
void QQuaternion_GetAxisAndAngle(const QQuaternion* self, QVector3D* axis, float* angle);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_FromAxisAndAngle(QVector3D* axis, float angle);
extern __declspec(dllexport) 
void QQuaternion_GetAxisAndAngle2(const QQuaternion* self, float* x, float* y, float* z, float* angle);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_FromAxisAndAngle2(float x, float y, float z, float angle);
extern __declspec(dllexport) 
QVector3D* QQuaternion_ToEulerAngles(const QQuaternion* self);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_FromEulerAngles(QVector3D* eulerAngles);
extern __declspec(dllexport) 
void QQuaternion_GetEulerAngles(const QQuaternion* self, float* pitch, float* yaw, float* roll);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_FromEulerAngles2(float pitch, float yaw, float roll);
extern __declspec(dllexport) 
void QQuaternion_GetAxes(const QQuaternion* self, QVector3D* xAxis, QVector3D* yAxis, QVector3D* zAxis);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_FromAxes(QVector3D* xAxis, QVector3D* yAxis, QVector3D* zAxis);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_FromDirection(QVector3D* direction, QVector3D* up);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_RotationTo(QVector3D* from, QVector3D* to);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_Slerp(QQuaternion* q1, QQuaternion* q2, float t);
extern __declspec(dllexport) 
QQuaternion* QQuaternion_Nlerp(QQuaternion* q1, QQuaternion* q2, float t);
extern __declspec(dllexport) 
void QQuaternion_Delete(QQuaternion* self, bool isSubclass);

}
