#pragma once
#ifndef MIQT_QT_GEN_QSIZEPOLICY_H
#define MIQT_QT_GEN_QSIZEPOLICY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QSizePolicy QSizePolicy;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSizePolicy* QSizePolicy_new();
extern __declspec(dllexport) QSizePolicy* QSizePolicy_new2(Policy horizontal, Policy vertical);
extern __declspec(dllexport) QSizePolicy* QSizePolicy_new3(QSizePolicy* param1);
extern __declspec(dllexport) QSizePolicy* QSizePolicy_new4(Policy horizontal, Policy vertical, ControlType typeVal);
extern __declspec(dllexport) Policy QSizePolicy_HorizontalPolicy(const QSizePolicy* self);
extern __declspec(dllexport) Policy QSizePolicy_VerticalPolicy(const QSizePolicy* self);
extern __declspec(dllexport) ControlType QSizePolicy_ControlType(const QSizePolicy* self);
extern __declspec(dllexport) void QSizePolicy_SetHorizontalPolicy(QSizePolicy* self, Policy d);
extern __declspec(dllexport) void QSizePolicy_SetVerticalPolicy(QSizePolicy* self, Policy d);
extern __declspec(dllexport) void QSizePolicy_SetControlType(QSizePolicy* self, ControlType typeVal);
extern __declspec(dllexport) int QSizePolicy_ExpandingDirections(const QSizePolicy* self);
extern __declspec(dllexport) void QSizePolicy_SetHeightForWidth(QSizePolicy* self, bool b);
extern __declspec(dllexport) bool QSizePolicy_HasHeightForWidth(const QSizePolicy* self);
extern __declspec(dllexport) void QSizePolicy_SetWidthForHeight(QSizePolicy* self, bool b);
extern __declspec(dllexport) bool QSizePolicy_HasWidthForHeight(const QSizePolicy* self);
extern __declspec(dllexport) bool QSizePolicy_OperatorEqual(const QSizePolicy* self, QSizePolicy* s);
extern __declspec(dllexport) bool QSizePolicy_OperatorNotEqual(const QSizePolicy* self, QSizePolicy* s);
extern __declspec(dllexport) int QSizePolicy_HorizontalStretch(const QSizePolicy* self);
extern __declspec(dllexport) int QSizePolicy_VerticalStretch(const QSizePolicy* self);
extern __declspec(dllexport) void QSizePolicy_SetHorizontalStretch(QSizePolicy* self, int stretchFactor);
extern __declspec(dllexport) void QSizePolicy_SetVerticalStretch(QSizePolicy* self, int stretchFactor);
extern __declspec(dllexport) bool QSizePolicy_RetainSizeWhenHidden(const QSizePolicy* self);
extern __declspec(dllexport) void QSizePolicy_SetRetainSizeWhenHidden(QSizePolicy* self, bool retainSize);
extern __declspec(dllexport) void QSizePolicy_Transpose(QSizePolicy* self);
extern __declspec(dllexport) QSizePolicy* QSizePolicy_Transposed(const QSizePolicy* self);
extern __declspec(dllexport) void QSizePolicy_Delete(QSizePolicy* self, bool isSubclass);

} 
