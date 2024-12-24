#pragma once
#ifndef MIQT_QT_GEN_QPAGELAYOUT_H
#define MIQT_QT_GEN_QPAGELAYOUT_H

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
class QPageLayout;
class QPageSize;
class QRect;
class QRectF;
class _GUID;
class type_info;
#else
typedef struct QMargins QMargins;
typedef struct QMarginsF QMarginsF;
typedef struct QPageLayout QPageLayout;
typedef struct QPageSize QPageSize;
typedef struct QRect QRect;
typedef struct QRectF QRectF;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QPageLayout* QPageLayout_new();
extern __declspec(dllexport) QPageLayout* QPageLayout_new2(QPageSize* pageSize, Orientation orientation, QMarginsF* margins);
extern __declspec(dllexport) QPageLayout* QPageLayout_new3(QPageLayout* other);
extern __declspec(dllexport) QPageLayout* QPageLayout_new4(QPageSize* pageSize, Orientation orientation, QMarginsF* margins, Unit units);
extern __declspec(dllexport) QPageLayout* QPageLayout_new5(QPageSize* pageSize, Orientation orientation, QMarginsF* margins, Unit units, QMarginsF* minMargins);
extern __declspec(dllexport) void QPageLayout_OperatorAssign(QPageLayout* self, QPageLayout* other);
extern __declspec(dllexport) void QPageLayout_Swap(QPageLayout* self, QPageLayout* other);
extern __declspec(dllexport) bool QPageLayout_IsEquivalentTo(const QPageLayout* self, QPageLayout* other);
extern __declspec(dllexport) bool QPageLayout_IsValid(const QPageLayout* self);
extern __declspec(dllexport) void QPageLayout_SetMode(QPageLayout* self, Mode mode);
extern __declspec(dllexport) Mode QPageLayout_Mode(const QPageLayout* self);
extern __declspec(dllexport) void QPageLayout_SetPageSize(QPageLayout* self, QPageSize* pageSize);
extern __declspec(dllexport) QPageSize* QPageLayout_PageSize(const QPageLayout* self);
extern __declspec(dllexport) void QPageLayout_SetOrientation(QPageLayout* self, Orientation orientation);
extern __declspec(dllexport) Orientation QPageLayout_Orientation(const QPageLayout* self);
extern __declspec(dllexport) void QPageLayout_SetUnits(QPageLayout* self, Unit units);
extern __declspec(dllexport) Unit QPageLayout_Units(const QPageLayout* self);
extern __declspec(dllexport) bool QPageLayout_SetMargins(QPageLayout* self, QMarginsF* margins);
extern __declspec(dllexport) bool QPageLayout_SetLeftMargin(QPageLayout* self, double leftMargin);
extern __declspec(dllexport) bool QPageLayout_SetRightMargin(QPageLayout* self, double rightMargin);
extern __declspec(dllexport) bool QPageLayout_SetTopMargin(QPageLayout* self, double topMargin);
extern __declspec(dllexport) bool QPageLayout_SetBottomMargin(QPageLayout* self, double bottomMargin);
extern __declspec(dllexport) QMarginsF* QPageLayout_Margins(const QPageLayout* self);
extern __declspec(dllexport) QMarginsF* QPageLayout_MarginsWithUnits(const QPageLayout* self, Unit units);
extern __declspec(dllexport) QMargins* QPageLayout_MarginsPoints(const QPageLayout* self);
extern __declspec(dllexport) QMargins* QPageLayout_MarginsPixels(const QPageLayout* self, int resolution);
extern __declspec(dllexport) void QPageLayout_SetMinimumMargins(QPageLayout* self, QMarginsF* minMargins);
extern __declspec(dllexport) QMarginsF* QPageLayout_MinimumMargins(const QPageLayout* self);
extern __declspec(dllexport) QMarginsF* QPageLayout_MaximumMargins(const QPageLayout* self);
extern __declspec(dllexport) QRectF* QPageLayout_FullRect(const QPageLayout* self);
extern __declspec(dllexport) QRectF* QPageLayout_FullRectWithUnits(const QPageLayout* self, Unit units);
extern __declspec(dllexport) QRect* QPageLayout_FullRectPoints(const QPageLayout* self);
extern __declspec(dllexport) QRect* QPageLayout_FullRectPixels(const QPageLayout* self, int resolution);
extern __declspec(dllexport) QRectF* QPageLayout_PaintRect(const QPageLayout* self);
extern __declspec(dllexport) QRectF* QPageLayout_PaintRectWithUnits(const QPageLayout* self, Unit units);
extern __declspec(dllexport) QRect* QPageLayout_PaintRectPoints(const QPageLayout* self);
extern __declspec(dllexport) QRect* QPageLayout_PaintRectPixels(const QPageLayout* self, int resolution);
extern __declspec(dllexport) void QPageLayout_SetPageSize2(QPageLayout* self, QPageSize* pageSize, QMarginsF* minMargins);
extern __declspec(dllexport) bool QPageLayout_SetMargins2(QPageLayout* self, QMarginsF* margins, OutOfBoundsPolicy outOfBoundsPolicy);
extern __declspec(dllexport) bool QPageLayout_SetLeftMargin2(QPageLayout* self, double leftMargin, OutOfBoundsPolicy outOfBoundsPolicy);
extern __declspec(dllexport) bool QPageLayout_SetRightMargin2(QPageLayout* self, double rightMargin, OutOfBoundsPolicy outOfBoundsPolicy);
extern __declspec(dllexport) bool QPageLayout_SetTopMargin2(QPageLayout* self, double topMargin, OutOfBoundsPolicy outOfBoundsPolicy);
extern __declspec(dllexport) bool QPageLayout_SetBottomMargin2(QPageLayout* self, double bottomMargin, OutOfBoundsPolicy outOfBoundsPolicy);
extern __declspec(dllexport) void QPageLayout_Delete(QPageLayout* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
