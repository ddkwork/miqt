#pragma once
#ifndef MIQT_QT_GEN_QEASINGCURVE_H
#define MIQT_QT_GEN_QEASINGCURVE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEasingCurve QEasingCurve;
typedef struct QPointF QPointF;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QEasingCurve* QEasingCurve_new();
extern __declspec(dllexport) QEasingCurve* QEasingCurve_new2(QEasingCurve* other);
extern __declspec(dllexport) QEasingCurve* QEasingCurve_new3(Type typeVal);
extern __declspec(dllexport) void QEasingCurve_OperatorAssign(QEasingCurve* self, QEasingCurve* other);
extern __declspec(dllexport) void QEasingCurve_Swap(QEasingCurve* self, QEasingCurve* other);
extern __declspec(dllexport) double QEasingCurve_Amplitude(const QEasingCurve* self);
extern __declspec(dllexport) void QEasingCurve_SetAmplitude(QEasingCurve* self, double amplitude);
extern __declspec(dllexport) double QEasingCurve_Period(const QEasingCurve* self);
extern __declspec(dllexport) void QEasingCurve_SetPeriod(QEasingCurve* self, double period);
extern __declspec(dllexport) double QEasingCurve_Overshoot(const QEasingCurve* self);
extern __declspec(dllexport) void QEasingCurve_SetOvershoot(QEasingCurve* self, double overshoot);
extern __declspec(dllexport) void QEasingCurve_AddCubicBezierSegment(QEasingCurve* self, QPointF* c1, QPointF* c2, QPointF* endPoint);
extern __declspec(dllexport) void QEasingCurve_AddTCBSegment(QEasingCurve* self, QPointF* nextPoint, double t, double c, double b);
extern __declspec(dllexport) struct miqt_array /* of QPointF* */  QEasingCurve_ToCubicSpline(const QEasingCurve* self);
extern __declspec(dllexport) Type QEasingCurve_Type(const QEasingCurve* self);
extern __declspec(dllexport) void QEasingCurve_SetType(QEasingCurve* self, Type typeVal);
extern __declspec(dllexport) void QEasingCurve_SetCustomType(QEasingCurve* self, EasingFunction funcVal);
extern __declspec(dllexport) EasingFunction QEasingCurve_CustomType(const QEasingCurve* self);
extern __declspec(dllexport) double QEasingCurve_ValueForProgress(const QEasingCurve* self, double progress);
extern __declspec(dllexport) void QEasingCurve_Delete(QEasingCurve* self, bool isSubclass);

} 
