#pragma once
#ifndef MIQT_QT_GEN_QLINE_H
#define MIQT_QT_GEN_QLINE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QLine;
class QLineF;
class QPoint;
class QPointF;
class _GUID;
class type_info;
#else
typedef struct QLine QLine;
typedef struct QLineF QLineF;
typedef struct QPoint QPoint;
typedef struct QPointF QPointF;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QLine* QLine_new();
extern __declspec(dllexport) QLine* QLine_new2(QPoint* pt1, QPoint* pt2);
extern __declspec(dllexport) QLine* QLine_new3(int x1, int y1, int x2, int y2);
extern __declspec(dllexport) QLine* QLine_new4(QLine* param1);
extern __declspec(dllexport) bool QLine_IsNull(const QLine* self);
extern __declspec(dllexport) QPoint* QLine_P1(const QLine* self);
extern __declspec(dllexport) QPoint* QLine_P2(const QLine* self);
extern __declspec(dllexport) int QLine_X1(const QLine* self);
extern __declspec(dllexport) int QLine_Y1(const QLine* self);
extern __declspec(dllexport) int QLine_X2(const QLine* self);
extern __declspec(dllexport) int QLine_Y2(const QLine* self);
extern __declspec(dllexport) int QLine_Dx(const QLine* self);
extern __declspec(dllexport) int QLine_Dy(const QLine* self);
extern __declspec(dllexport) void QLine_Translate(QLine* self, QPoint* p);
extern __declspec(dllexport) void QLine_Translate2(QLine* self, int dx, int dy);
extern __declspec(dllexport) QLine* QLine_Translated(const QLine* self, QPoint* p);
extern __declspec(dllexport) QLine* QLine_Translated2(const QLine* self, int dx, int dy);
extern __declspec(dllexport) QPoint* QLine_Center(const QLine* self);
extern __declspec(dllexport) void QLine_SetP1(QLine* self, QPoint* p1);
extern __declspec(dllexport) void QLine_SetP2(QLine* self, QPoint* p2);
extern __declspec(dllexport) void QLine_SetPoints(QLine* self, QPoint* p1, QPoint* p2);
extern __declspec(dllexport) void QLine_SetLine(QLine* self, int x1, int y1, int x2, int y2);
extern __declspec(dllexport) QLineF* QLine_ToLineF(const QLine* self);
extern __declspec(dllexport) void QLine_Delete(QLine* self, bool isSubclass);

extern __declspec(dllexport) QLineF* QLineF_new();
extern __declspec(dllexport) QLineF* QLineF_new2(QPointF* pt1, QPointF* pt2);
extern __declspec(dllexport) QLineF* QLineF_new3(double x1, double y1, double x2, double y2);
extern __declspec(dllexport) QLineF* QLineF_new4(QLine* line);
extern __declspec(dllexport) QLineF* QLineF_new5(QLineF* param1);
extern __declspec(dllexport) QLineF* QLineF_FromPolar(double length, double angle);
extern __declspec(dllexport) bool QLineF_IsNull(const QLineF* self);
extern __declspec(dllexport) QPointF* QLineF_P1(const QLineF* self);
extern __declspec(dllexport) QPointF* QLineF_P2(const QLineF* self);
extern __declspec(dllexport) double QLineF_X1(const QLineF* self);
extern __declspec(dllexport) double QLineF_Y1(const QLineF* self);
extern __declspec(dllexport) double QLineF_X2(const QLineF* self);
extern __declspec(dllexport) double QLineF_Y2(const QLineF* self);
extern __declspec(dllexport) double QLineF_Dx(const QLineF* self);
extern __declspec(dllexport) double QLineF_Dy(const QLineF* self);
extern __declspec(dllexport) double QLineF_Length(const QLineF* self);
extern __declspec(dllexport) void QLineF_SetLength(QLineF* self, double lenVal);
extern __declspec(dllexport) double QLineF_Angle(const QLineF* self);
extern __declspec(dllexport) void QLineF_SetAngle(QLineF* self, double angle);
extern __declspec(dllexport) double QLineF_AngleTo(const QLineF* self, QLineF* l);
extern __declspec(dllexport) QLineF* QLineF_UnitVector(const QLineF* self);
extern __declspec(dllexport) QLineF* QLineF_NormalVector(const QLineF* self);
extern __declspec(dllexport) IntersectionType QLineF_Intersects(const QLineF* self, QLineF* l);
extern __declspec(dllexport) QPointF* QLineF_PointAt(const QLineF* self, double t);
extern __declspec(dllexport) void QLineF_Translate(QLineF* self, QPointF* p);
extern __declspec(dllexport) void QLineF_Translate2(QLineF* self, double dx, double dy);
extern __declspec(dllexport) QLineF* QLineF_Translated(const QLineF* self, QPointF* p);
extern __declspec(dllexport) QLineF* QLineF_Translated2(const QLineF* self, double dx, double dy);
extern __declspec(dllexport) QPointF* QLineF_Center(const QLineF* self);
extern __declspec(dllexport) void QLineF_SetP1(QLineF* self, QPointF* p1);
extern __declspec(dllexport) void QLineF_SetP2(QLineF* self, QPointF* p2);
extern __declspec(dllexport) void QLineF_SetPoints(QLineF* self, QPointF* p1, QPointF* p2);
extern __declspec(dllexport) void QLineF_SetLine(QLineF* self, double x1, double y1, double x2, double y2);
extern __declspec(dllexport) QLine* QLineF_ToLine(const QLineF* self);
extern __declspec(dllexport) IntersectionType QLineF_Intersects2(const QLineF* self, QLineF* l, QPointF* intersectionPoint);
extern __declspec(dllexport) void QLineF_Delete(QLineF* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
