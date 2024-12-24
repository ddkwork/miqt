#pragma once
#ifndef MIQT_QT_GEN_QPEN_H
#define MIQT_QT_GEN_QPEN_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QBrush;
class QColor;
class QPen;
class _GUID;
class type_info;
#else
typedef struct QBrush QBrush;
typedef struct QColor QColor;
typedef struct QPen QPen;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QPen* QPen_new();
extern __declspec(dllexport) QPen* QPen_new2(int param1);
extern __declspec(dllexport) QPen* QPen_new3(QColor* color);
extern __declspec(dllexport) QPen* QPen_new4(QBrush* brush, double width);
extern __declspec(dllexport) QPen* QPen_new5(QPen* pen);
extern __declspec(dllexport) QPen* QPen_new6(QBrush* brush, double width, int s);
extern __declspec(dllexport) QPen* QPen_new7(QBrush* brush, double width, int s, int c);
extern __declspec(dllexport) QPen* QPen_new8(QBrush* brush, double width, int s, int c, int j);
extern __declspec(dllexport) void QPen_OperatorAssign(QPen* self, QPen* pen);
extern __declspec(dllexport) void QPen_Swap(QPen* self, QPen* other);
extern __declspec(dllexport) void QPen_OperatorAssignWithColor(QPen* self, QColor* color);
extern __declspec(dllexport) void QPen_OperatorAssignWithStyle(QPen* self, int style);
extern __declspec(dllexport) int QPen_Style(const QPen* self);
extern __declspec(dllexport) void QPen_SetStyle(QPen* self, int style);
extern __declspec(dllexport) struct miqt_array /* of double */  QPen_DashPattern(const QPen* self);
extern __declspec(dllexport) void QPen_SetDashPattern(QPen* self, struct miqt_array /* of double */  pattern);
extern __declspec(dllexport) double QPen_DashOffset(const QPen* self);
extern __declspec(dllexport) void QPen_SetDashOffset(QPen* self, double doffset);
extern __declspec(dllexport) double QPen_MiterLimit(const QPen* self);
extern __declspec(dllexport) void QPen_SetMiterLimit(QPen* self, double limit);
extern __declspec(dllexport) double QPen_WidthF(const QPen* self);
extern __declspec(dllexport) void QPen_SetWidthF(QPen* self, double width);
extern __declspec(dllexport) int QPen_Width(const QPen* self);
extern __declspec(dllexport) void QPen_SetWidth(QPen* self, int width);
extern __declspec(dllexport) QColor* QPen_Color(const QPen* self);
extern __declspec(dllexport) void QPen_SetColor(QPen* self, QColor* color);
extern __declspec(dllexport) QBrush* QPen_Brush(const QPen* self);
extern __declspec(dllexport) void QPen_SetBrush(QPen* self, QBrush* brush);
extern __declspec(dllexport) bool QPen_IsSolid(const QPen* self);
extern __declspec(dllexport) int QPen_CapStyle(const QPen* self);
extern __declspec(dllexport) void QPen_SetCapStyle(QPen* self, int pcs);
extern __declspec(dllexport) int QPen_JoinStyle(const QPen* self);
extern __declspec(dllexport) void QPen_SetJoinStyle(QPen* self, int pcs);
extern __declspec(dllexport) bool QPen_IsCosmetic(const QPen* self);
extern __declspec(dllexport) void QPen_SetCosmetic(QPen* self, bool cosmetic);
extern __declspec(dllexport) bool QPen_OperatorEqual(const QPen* self, QPen* p);
extern __declspec(dllexport) bool QPen_OperatorNotEqual(const QPen* self, QPen* p);
extern __declspec(dllexport) bool QPen_IsDetached(QPen* self);
extern __declspec(dllexport) DataPtr* QPen_DataPtr(QPen* self);
extern __declspec(dllexport) void QPen_Delete(QPen* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
