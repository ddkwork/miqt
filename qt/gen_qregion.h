#pragma once
#ifndef MIQT_QT_GEN_QREGION_H
#define MIQT_QT_GEN_QREGION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QBitmap QBitmap;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QRegion QRegion;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QRegion* QRegion_new();
extern __declspec(dllexport) QRegion* QRegion_new2(int x, int y, int w, int h);
extern __declspec(dllexport) QRegion* QRegion_new3(QRect* r);
extern __declspec(dllexport) QRegion* QRegion_new4(QRegion* region);
extern __declspec(dllexport) QRegion* QRegion_new5(QBitmap* bitmap);
extern __declspec(dllexport) QRegion* QRegion_new6(int x, int y, int w, int h, RegionType t);
extern __declspec(dllexport) QRegion* QRegion_new7(QRect* r, RegionType t);
extern __declspec(dllexport) void QRegion_OperatorAssign(QRegion* self, QRegion* param1);
extern __declspec(dllexport) void QRegion_Swap(QRegion* self, QRegion* other);
extern __declspec(dllexport) bool QRegion_IsEmpty(const QRegion* self);
extern __declspec(dllexport) bool QRegion_IsNull(const QRegion* self);
extern __declspec(dllexport) const_iterator QRegion_Begin(const QRegion* self);
extern __declspec(dllexport) const_iterator QRegion_Cbegin(const QRegion* self);
extern __declspec(dllexport) const_iterator QRegion_End(const QRegion* self);
extern __declspec(dllexport) const_iterator QRegion_Cend(const QRegion* self);
extern __declspec(dllexport) const_reverse_iterator QRegion_Rbegin(const QRegion* self);
extern __declspec(dllexport) const_reverse_iterator QRegion_Crbegin(const QRegion* self);
extern __declspec(dllexport) const_reverse_iterator QRegion_Rend(const QRegion* self);
extern __declspec(dllexport) const_reverse_iterator QRegion_Crend(const QRegion* self);
extern __declspec(dllexport) bool QRegion_Contains(const QRegion* self, QPoint* p);
extern __declspec(dllexport) bool QRegion_ContainsWithQRect(const QRegion* self, QRect* r);
extern __declspec(dllexport) void QRegion_Translate(QRegion* self, int dx, int dy);
extern __declspec(dllexport) void QRegion_TranslateWithQPoint(QRegion* self, QPoint* p);
extern __declspec(dllexport) QRegion* QRegion_Translated(const QRegion* self, int dx, int dy);
extern __declspec(dllexport) QRegion* QRegion_TranslatedWithQPoint(const QRegion* self, QPoint* p);
extern __declspec(dllexport) QRegion* QRegion_United(const QRegion* self, QRegion* r);
extern __declspec(dllexport) QRegion* QRegion_UnitedWithQRect(const QRegion* self, QRect* r);
extern __declspec(dllexport) QRegion* QRegion_Intersected(const QRegion* self, QRegion* r);
extern __declspec(dllexport) QRegion* QRegion_IntersectedWithQRect(const QRegion* self, QRect* r);
extern __declspec(dllexport) QRegion* QRegion_Subtracted(const QRegion* self, QRegion* r);
extern __declspec(dllexport) QRegion* QRegion_Xored(const QRegion* self, QRegion* r);
extern __declspec(dllexport) bool QRegion_Intersects(const QRegion* self, QRegion* r);
extern __declspec(dllexport) bool QRegion_IntersectsWithQRect(const QRegion* self, QRect* r);
extern __declspec(dllexport) QRect* QRegion_BoundingRect(const QRegion* self);
extern __declspec(dllexport) void QRegion_SetRects(QRegion* self, QRect* rect, int num);
extern __declspec(dllexport) void QRegion_SetRectsWithQSpanLesserconstQRectGreater(QRegion* self, QSpan<const QRect> r);
extern __declspec(dllexport) QSpan<const QRect> QRegion_Rects(const QRegion* self);
extern __declspec(dllexport) int QRegion_RectCount(const QRegion* self);
extern __declspec(dllexport) QRegion* QRegion_OperatorBitwiseOr(const QRegion* self, QRegion* r);
extern __declspec(dllexport) QRegion* QRegion_OperatorPlus(const QRegion* self, QRegion* r);
extern __declspec(dllexport) QRegion* QRegion_OperatorPlusWithQRect(const QRegion* self, QRect* r);
extern __declspec(dllexport) QRegion* QRegion_OperatorBitwiseAnd(const QRegion* self, QRegion* r);
extern __declspec(dllexport) QRegion* QRegion_OperatorBitwiseAndWithQRect(const QRegion* self, QRect* r);
extern __declspec(dllexport) QRegion* QRegion_OperatorMinus(const QRegion* self, QRegion* r);
extern __declspec(dllexport) QRegion* QRegion_OperatorBitwiseNot(const QRegion* self, QRegion* r);
extern __declspec(dllexport) void QRegion_OperatorBitwiseOrAssign(QRegion* self, QRegion* r);
extern __declspec(dllexport) QRegion* QRegion_OperatorPlusAssign(QRegion* self, QRegion* r);
extern __declspec(dllexport) QRegion* QRegion_OperatorPlusAssignWithQRect(QRegion* self, QRect* r);
extern __declspec(dllexport) void QRegion_OperatorBitwiseAndAssign(QRegion* self, QRegion* r);
extern __declspec(dllexport) void QRegion_OperatorBitwiseAndAssignWithQRect(QRegion* self, QRect* r);
extern __declspec(dllexport) QRegion* QRegion_OperatorMinusAssign(QRegion* self, QRegion* r);
extern __declspec(dllexport) void QRegion_OperatorBitwiseNotAssign(QRegion* self, QRegion* r);
extern __declspec(dllexport) bool QRegion_OperatorEqual(const QRegion* self, QRegion* r);
extern __declspec(dllexport) bool QRegion_OperatorNotEqual(const QRegion* self, QRegion* r);
extern __declspec(dllexport) struct HRGN__* QRegion_ToHRGN(const QRegion* self);
extern __declspec(dllexport) QRegion* QRegion_FromHRGN(struct HRGN__* hrgn);
extern __declspec(dllexport) void QRegion_Delete(QRegion* self, bool isSubclass);

} 
