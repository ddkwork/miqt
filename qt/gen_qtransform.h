#pragma once
#ifndef MIQT_QT_GEN_QTRANSFORM_H
#define MIQT_QT_GEN_QTRANSFORM_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QLine QLine;
typedef struct QLineF QLineF;
typedef struct QPainterPath QPainterPath;
typedef struct QPoint QPoint;
typedef struct QPointF QPointF;
typedef struct QRect QRect;
typedef struct QRectF QRectF;
typedef struct QRegion QRegion;
typedef struct QTransform QTransform;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTransform* QTransform_new(int param1);
extern __declspec(dllexport) QTransform* QTransform_new2();
extern __declspec(dllexport) QTransform* QTransform_new3(double h11, double h12, double h13, double h21, double h22, double h23, double h31, double h32, double h33);
extern __declspec(dllexport) QTransform* QTransform_new4(double h11, double h12, double h21, double h22, double dx, double dy);
extern __declspec(dllexport) QTransform* QTransform_new5(QTransform* other);
extern __declspec(dllexport) void QTransform_OperatorAssign(QTransform* self, QTransform* param1);
extern __declspec(dllexport) bool QTransform_IsAffine(const QTransform* self);
extern __declspec(dllexport) bool QTransform_IsIdentity(const QTransform* self);
extern __declspec(dllexport) bool QTransform_IsInvertible(const QTransform* self);
extern __declspec(dllexport) bool QTransform_IsScaling(const QTransform* self);
extern __declspec(dllexport) bool QTransform_IsRotating(const QTransform* self);
extern __declspec(dllexport) bool QTransform_IsTranslating(const QTransform* self);
extern __declspec(dllexport) TransformationType QTransform_Type(const QTransform* self);
extern __declspec(dllexport) double QTransform_Determinant(const QTransform* self);
extern __declspec(dllexport) double QTransform_M11(const QTransform* self);
extern __declspec(dllexport) double QTransform_M12(const QTransform* self);
extern __declspec(dllexport) double QTransform_M13(const QTransform* self);
extern __declspec(dllexport) double QTransform_M21(const QTransform* self);
extern __declspec(dllexport) double QTransform_M22(const QTransform* self);
extern __declspec(dllexport) double QTransform_M23(const QTransform* self);
extern __declspec(dllexport) double QTransform_M31(const QTransform* self);
extern __declspec(dllexport) double QTransform_M32(const QTransform* self);
extern __declspec(dllexport) double QTransform_M33(const QTransform* self);
extern __declspec(dllexport) double QTransform_Dx(const QTransform* self);
extern __declspec(dllexport) double QTransform_Dy(const QTransform* self);
extern __declspec(dllexport) void QTransform_SetMatrix(QTransform* self, double m11, double m12, double m13, double m21, double m22, double m23, double m31, double m32, double m33);
extern __declspec(dllexport) QTransform* QTransform_Inverted(const QTransform* self);
extern __declspec(dllexport) QTransform* QTransform_Adjoint(const QTransform* self);
extern __declspec(dllexport) QTransform* QTransform_Transposed(const QTransform* self);
extern __declspec(dllexport) QTransform* QTransform_Translate(QTransform* self, double dx, double dy);
extern __declspec(dllexport) QTransform* QTransform_Scale(QTransform* self, double sx, double sy);
extern __declspec(dllexport) QTransform* QTransform_Shear(QTransform* self, double sh, double sv);
extern __declspec(dllexport) QTransform* QTransform_Rotate(QTransform* self, double a, int axis, double distanceToPlane);
extern __declspec(dllexport) QTransform* QTransform_RotateWithQreal(QTransform* self, double a);
extern __declspec(dllexport) QTransform* QTransform_RotateRadians(QTransform* self, double a, int axis, double distanceToPlane);
extern __declspec(dllexport) QTransform* QTransform_RotateRadiansWithQreal(QTransform* self, double a);
extern __declspec(dllexport) bool QTransform_OperatorEqual(const QTransform* self, QTransform* param1);
extern __declspec(dllexport) bool QTransform_OperatorNotEqual(const QTransform* self, QTransform* param1);
extern __declspec(dllexport) QTransform* QTransform_OperatorMultiplyAssign(QTransform* self, QTransform* param1);
extern __declspec(dllexport) QTransform* QTransform_OperatorMultiply(const QTransform* self, QTransform* o);
extern __declspec(dllexport) void QTransform_Reset(QTransform* self);
extern __declspec(dllexport) QPoint* QTransform_Map(const QTransform* self, QPoint* p);
extern __declspec(dllexport) QPointF* QTransform_MapWithQPointF(const QTransform* self, QPointF* p);
extern __declspec(dllexport) QLine* QTransform_MapWithQLine(const QTransform* self, QLine* l);
extern __declspec(dllexport) QLineF* QTransform_MapWithQLineF(const QTransform* self, QLineF* l);
extern __declspec(dllexport) QRegion* QTransform_MapWithQRegion(const QTransform* self, QRegion* r);
extern __declspec(dllexport) QPainterPath* QTransform_MapWithQPainterPath(const QTransform* self, QPainterPath* p);
extern __declspec(dllexport) QRect* QTransform_MapRect(const QTransform* self, QRect* param1);
extern __declspec(dllexport) QRectF* QTransform_MapRectWithQRectF(const QTransform* self, QRectF* param1);
extern __declspec(dllexport) void QTransform_Map2(const QTransform* self, int x, int y, int* tx, int* ty);
extern __declspec(dllexport) void QTransform_Map3(const QTransform* self, double x, double y, double* tx, double* ty);
extern __declspec(dllexport) QTransform* QTransform_OperatorMultiplyAssignWithDiv(QTransform* self, double div);
extern __declspec(dllexport) QTransform* QTransform_OperatorDivideAssign(QTransform* self, double div);
extern __declspec(dllexport) QTransform* QTransform_OperatorPlusAssign(QTransform* self, double div);
extern __declspec(dllexport) QTransform* QTransform_OperatorMinusAssign(QTransform* self, double div);
extern __declspec(dllexport) QTransform* QTransform_FromTranslate(double dx, double dy);
extern __declspec(dllexport) QTransform* QTransform_FromScale(double dx, double dy);
extern __declspec(dllexport) Affine QTransform_AsAffineMatrix(QTransform* self);
extern __declspec(dllexport) QTransform* QTransform_Inverted1(const QTransform* self, bool* invertible);
extern __declspec(dllexport) QTransform* QTransform_Rotate2(QTransform* self, double a, int axis);
extern __declspec(dllexport) QTransform* QTransform_RotateRadians2(QTransform* self, double a, int axis);
extern __declspec(dllexport) void QTransform_Delete(QTransform* self, bool isSubclass);

} 
