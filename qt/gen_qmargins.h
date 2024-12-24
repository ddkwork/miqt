#pragma once
#ifndef MIQT_QT_GEN_QMARGINS_H
#define MIQT_QT_GEN_QMARGINS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QMargins;
class QMarginsF;
class _GUID;
class type_info;
#else
typedef struct QMargins QMargins;
typedef struct QMarginsF QMarginsF;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QMargins* QMargins_new();
extern __declspec(dllexport) QMargins* QMargins_new2(int left, int top, int right, int bottom);
extern __declspec(dllexport) QMargins* QMargins_new3(QMargins* param1);
extern __declspec(dllexport) bool QMargins_IsNull(const QMargins* self);
extern __declspec(dllexport) int QMargins_Left(const QMargins* self);
extern __declspec(dllexport) int QMargins_Top(const QMargins* self);
extern __declspec(dllexport) int QMargins_Right(const QMargins* self);
extern __declspec(dllexport) int QMargins_Bottom(const QMargins* self);
extern __declspec(dllexport) void QMargins_SetLeft(QMargins* self, int left);
extern __declspec(dllexport) void QMargins_SetTop(QMargins* self, int top);
extern __declspec(dllexport) void QMargins_SetRight(QMargins* self, int right);
extern __declspec(dllexport) void QMargins_SetBottom(QMargins* self, int bottom);
extern __declspec(dllexport) QMargins* QMargins_OperatorPlusAssign(QMargins* self, QMargins* margins);
extern __declspec(dllexport) QMargins* QMargins_OperatorMinusAssign(QMargins* self, QMargins* margins);
extern __declspec(dllexport) QMargins* QMargins_OperatorPlusAssignWithInt(QMargins* self, int param1);
extern __declspec(dllexport) QMargins* QMargins_OperatorMinusAssignWithInt(QMargins* self, int param1);
extern __declspec(dllexport) QMargins* QMargins_OperatorMultiplyAssign(QMargins* self, int param1);
extern __declspec(dllexport) QMargins* QMargins_OperatorDivideAssign(QMargins* self, int param1);
extern __declspec(dllexport) QMargins* QMargins_OperatorMultiplyAssignWithQreal(QMargins* self, double param1);
extern __declspec(dllexport) QMargins* QMargins_OperatorDivideAssignWithQreal(QMargins* self, double param1);
extern __declspec(dllexport) QMarginsF* QMargins_ToMarginsF(const QMargins* self);
extern __declspec(dllexport) void QMargins_Delete(QMargins* self, bool isSubclass);

extern __declspec(dllexport) QMarginsF* QMarginsF_new();
extern __declspec(dllexport) QMarginsF* QMarginsF_new2(double left, double top, double right, double bottom);
extern __declspec(dllexport) QMarginsF* QMarginsF_new3(QMargins* margins);
extern __declspec(dllexport) QMarginsF* QMarginsF_new4(QMarginsF* param1);
extern __declspec(dllexport) bool QMarginsF_IsNull(const QMarginsF* self);
extern __declspec(dllexport) double QMarginsF_Left(const QMarginsF* self);
extern __declspec(dllexport) double QMarginsF_Top(const QMarginsF* self);
extern __declspec(dllexport) double QMarginsF_Right(const QMarginsF* self);
extern __declspec(dllexport) double QMarginsF_Bottom(const QMarginsF* self);
extern __declspec(dllexport) void QMarginsF_SetLeft(QMarginsF* self, double aleft);
extern __declspec(dllexport) void QMarginsF_SetTop(QMarginsF* self, double atop);
extern __declspec(dllexport) void QMarginsF_SetRight(QMarginsF* self, double aright);
extern __declspec(dllexport) void QMarginsF_SetBottom(QMarginsF* self, double abottom);
extern __declspec(dllexport) QMarginsF* QMarginsF_OperatorPlusAssign(QMarginsF* self, QMarginsF* margins);
extern __declspec(dllexport) QMarginsF* QMarginsF_OperatorMinusAssign(QMarginsF* self, QMarginsF* margins);
extern __declspec(dllexport) QMarginsF* QMarginsF_OperatorPlusAssignWithAddend(QMarginsF* self, double addend);
extern __declspec(dllexport) QMarginsF* QMarginsF_OperatorMinusAssignWithSubtrahend(QMarginsF* self, double subtrahend);
extern __declspec(dllexport) QMarginsF* QMarginsF_OperatorMultiplyAssign(QMarginsF* self, double factor);
extern __declspec(dllexport) QMarginsF* QMarginsF_OperatorDivideAssign(QMarginsF* self, double divisor);
extern __declspec(dllexport) QMargins* QMarginsF_ToMargins(const QMarginsF* self);
extern __declspec(dllexport) void QMarginsF_Delete(QMarginsF* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
