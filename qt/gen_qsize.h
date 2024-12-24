#pragma once
#ifndef MIQT_QT_GEN_QSIZE_H
#define MIQT_QT_GEN_QSIZE_H

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
class QSize;
class QSizeF;
class _GUID;
class type_info;
#else
typedef struct QMargins QMargins;
typedef struct QMarginsF QMarginsF;
typedef struct QSize QSize;
typedef struct QSizeF QSizeF;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSize* QSize_new();
extern __declspec(dllexport) QSize* QSize_new2(int w, int h);
extern __declspec(dllexport) QSize* QSize_new3(QSize* param1);
extern __declspec(dllexport) bool QSize_IsNull(const QSize* self);
extern __declspec(dllexport) bool QSize_IsEmpty(const QSize* self);
extern __declspec(dllexport) bool QSize_IsValid(const QSize* self);
extern __declspec(dllexport) int QSize_Width(const QSize* self);
extern __declspec(dllexport) int QSize_Height(const QSize* self);
extern __declspec(dllexport) void QSize_SetWidth(QSize* self, int w);
extern __declspec(dllexport) void QSize_SetHeight(QSize* self, int h);
extern __declspec(dllexport) void QSize_Transpose(QSize* self);
extern __declspec(dllexport) QSize* QSize_Transposed(const QSize* self);
extern __declspec(dllexport) void QSize_Scale(QSize* self, int w, int h, int mode);
extern __declspec(dllexport) void QSize_Scale2(QSize* self, QSize* s, int mode);
extern __declspec(dllexport) QSize* QSize_Scaled(const QSize* self, int w, int h, int mode);
extern __declspec(dllexport) QSize* QSize_Scaled2(const QSize* self, QSize* s, int mode);
extern __declspec(dllexport) QSize* QSize_ExpandedTo(const QSize* self, QSize* param1);
extern __declspec(dllexport) QSize* QSize_BoundedTo(const QSize* self, QSize* param1);
extern __declspec(dllexport) QSize* QSize_GrownBy(const QSize* self, QMargins* m);
extern __declspec(dllexport) QSize* QSize_ShrunkBy(const QSize* self, QMargins* m);
extern __declspec(dllexport) QSize* QSize_OperatorPlusAssign(QSize* self, QSize* param1);
extern __declspec(dllexport) QSize* QSize_OperatorMinusAssign(QSize* self, QSize* param1);
extern __declspec(dllexport) QSize* QSize_OperatorMultiplyAssign(QSize* self, double c);
extern __declspec(dllexport) QSize* QSize_OperatorDivideAssign(QSize* self, double c);
extern __declspec(dllexport) QSizeF* QSize_ToSizeF(const QSize* self);
extern __declspec(dllexport) void QSize_Delete(QSize* self, bool isSubclass);

extern __declspec(dllexport) QSizeF* QSizeF_new();
extern __declspec(dllexport) QSizeF* QSizeF_new2(QSize* sz);
extern __declspec(dllexport) QSizeF* QSizeF_new3(double w, double h);
extern __declspec(dllexport) QSizeF* QSizeF_new4(QSizeF* param1);
extern __declspec(dllexport) bool QSizeF_IsNull(const QSizeF* self);
extern __declspec(dllexport) bool QSizeF_IsEmpty(const QSizeF* self);
extern __declspec(dllexport) bool QSizeF_IsValid(const QSizeF* self);
extern __declspec(dllexport) double QSizeF_Width(const QSizeF* self);
extern __declspec(dllexport) double QSizeF_Height(const QSizeF* self);
extern __declspec(dllexport) void QSizeF_SetWidth(QSizeF* self, double w);
extern __declspec(dllexport) void QSizeF_SetHeight(QSizeF* self, double h);
extern __declspec(dllexport) void QSizeF_Transpose(QSizeF* self);
extern __declspec(dllexport) QSizeF* QSizeF_Transposed(const QSizeF* self);
extern __declspec(dllexport) void QSizeF_Scale(QSizeF* self, double w, double h, int mode);
extern __declspec(dllexport) void QSizeF_Scale2(QSizeF* self, QSizeF* s, int mode);
extern __declspec(dllexport) QSizeF* QSizeF_Scaled(const QSizeF* self, double w, double h, int mode);
extern __declspec(dllexport) QSizeF* QSizeF_Scaled2(const QSizeF* self, QSizeF* s, int mode);
extern __declspec(dllexport) QSizeF* QSizeF_ExpandedTo(const QSizeF* self, QSizeF* param1);
extern __declspec(dllexport) QSizeF* QSizeF_BoundedTo(const QSizeF* self, QSizeF* param1);
extern __declspec(dllexport) QSizeF* QSizeF_GrownBy(const QSizeF* self, QMarginsF* m);
extern __declspec(dllexport) QSizeF* QSizeF_ShrunkBy(const QSizeF* self, QMarginsF* m);
extern __declspec(dllexport) QSizeF* QSizeF_OperatorPlusAssign(QSizeF* self, QSizeF* param1);
extern __declspec(dllexport) QSizeF* QSizeF_OperatorMinusAssign(QSizeF* self, QSizeF* param1);
extern __declspec(dllexport) QSizeF* QSizeF_OperatorMultiplyAssign(QSizeF* self, double c);
extern __declspec(dllexport) QSizeF* QSizeF_OperatorDivideAssign(QSizeF* self, double c);
extern __declspec(dllexport) QSize* QSizeF_ToSize(const QSizeF* self);
extern __declspec(dllexport) void QSizeF_Delete(QSizeF* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
