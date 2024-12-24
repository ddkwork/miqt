#pragma once
#ifndef MIQT_QT_GEN_QPOINT_H
#define MIQT_QT_GEN_QPOINT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QPoint;
class QPointF;
class _GUID;
class type_info;
#else
typedef struct QPoint QPoint;
typedef struct QPointF QPointF;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QPoint* QPoint_new();
extern __declspec(dllexport) QPoint* QPoint_new2(int xpos, int ypos);
extern __declspec(dllexport) QPoint* QPoint_new3(QPoint* param1);
extern __declspec(dllexport) bool QPoint_IsNull(const QPoint* self);
extern __declspec(dllexport) int QPoint_X(const QPoint* self);
extern __declspec(dllexport) int QPoint_Y(const QPoint* self);
extern __declspec(dllexport) void QPoint_SetX(QPoint* self, int x);
extern __declspec(dllexport) void QPoint_SetY(QPoint* self, int y);
extern __declspec(dllexport) int QPoint_ManhattanLength(const QPoint* self);
extern __declspec(dllexport) QPoint* QPoint_Transposed(const QPoint* self);
extern __declspec(dllexport) QPoint* QPoint_OperatorPlusAssign(QPoint* self, QPoint* p);
extern __declspec(dllexport) QPoint* QPoint_OperatorMinusAssign(QPoint* self, QPoint* p);
extern __declspec(dllexport) QPoint* QPoint_OperatorMultiplyAssign(QPoint* self, float factor);
extern __declspec(dllexport) QPoint* QPoint_OperatorMultiplyAssignWithFactor(QPoint* self, double factor);
extern __declspec(dllexport) QPoint* QPoint_OperatorMultiplyAssign2(QPoint* self, int factor);
extern __declspec(dllexport) QPoint* QPoint_OperatorDivideAssign(QPoint* self, double divisor);
extern __declspec(dllexport) int QPoint_DotProduct(QPoint* p1, QPoint* p2);
extern __declspec(dllexport) QPointF* QPoint_ToPointF(const QPoint* self);
extern __declspec(dllexport) void QPoint_Delete(QPoint* self, bool isSubclass);

extern __declspec(dllexport) QPointF* QPointF_new();
extern __declspec(dllexport) QPointF* QPointF_new2(QPoint* p);
extern __declspec(dllexport) QPointF* QPointF_new3(double xpos, double ypos);
extern __declspec(dllexport) QPointF* QPointF_new4(QPointF* param1);
extern __declspec(dllexport) double QPointF_ManhattanLength(const QPointF* self);
extern __declspec(dllexport) bool QPointF_IsNull(const QPointF* self);
extern __declspec(dllexport) double QPointF_X(const QPointF* self);
extern __declspec(dllexport) double QPointF_Y(const QPointF* self);
extern __declspec(dllexport) void QPointF_SetX(QPointF* self, double x);
extern __declspec(dllexport) void QPointF_SetY(QPointF* self, double y);
extern __declspec(dllexport) QPointF* QPointF_Transposed(const QPointF* self);
extern __declspec(dllexport) QPointF* QPointF_OperatorPlusAssign(QPointF* self, QPointF* p);
extern __declspec(dllexport) QPointF* QPointF_OperatorMinusAssign(QPointF* self, QPointF* p);
extern __declspec(dllexport) QPointF* QPointF_OperatorMultiplyAssign(QPointF* self, double c);
extern __declspec(dllexport) QPointF* QPointF_OperatorDivideAssign(QPointF* self, double c);
extern __declspec(dllexport) double QPointF_DotProduct(QPointF* p1, QPointF* p2);
extern __declspec(dllexport) QPoint* QPointF_ToPoint(const QPointF* self);
extern __declspec(dllexport) void QPointF_Delete(QPointF* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
