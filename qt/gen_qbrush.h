#pragma once
#ifndef MIQT_QT_GEN_QBRUSH_H
#define MIQT_QT_GEN_QBRUSH_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QGradient__QGradientData)
typedef QGradient::QGradientData QGradient__QGradientData;
typedef struct QBrush QBrush;
typedef struct QBrushData QBrushData;
typedef struct QColor QColor;
typedef struct QConicalGradient QConicalGradient;
typedef struct QGradient QGradient;
typedef struct QGradient__QGradientData QGradient__QGradientData;
typedef struct QImage QImage;
typedef struct QLinearGradient QLinearGradient;
typedef struct QPixmap QPixmap;
typedef struct QPointF QPointF;
typedef struct QRadialGradient QRadialGradient;
typedef struct QTransform QTransform;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QBrush* QBrush_new();
extern __declspec(dllexport) 
QBrush* QBrush_new2(int bs);
extern __declspec(dllexport) 
QBrush* QBrush_new3(QColor* color);
extern __declspec(dllexport) 
QBrush* QBrush_new4(int color);
extern __declspec(dllexport) 
QBrush* QBrush_new5(QColor* color, QPixmap* pixmap);
extern __declspec(dllexport) 
QBrush* QBrush_new6(int color, QPixmap* pixmap);
extern __declspec(dllexport) 
QBrush* QBrush_new7(QPixmap* pixmap);
extern __declspec(dllexport) 
QBrush* QBrush_new8(QImage* image);
extern __declspec(dllexport) 
QBrush* QBrush_new9(QBrush* brush);
extern __declspec(dllexport) 
QBrush* QBrush_new10(QGradient* gradient);
extern __declspec(dllexport) 
QBrush* QBrush_new11(QColor* color, int bs);
extern __declspec(dllexport) 
QBrush* QBrush_new12(int color, int bs);
extern __declspec(dllexport) 
void QBrush_OperatorAssign(QBrush* self, QBrush* brush);
extern __declspec(dllexport) 
void QBrush_Swap(QBrush* self, QBrush* other);
extern __declspec(dllexport) 
void QBrush_OperatorAssignWithStyle(QBrush* self, int style);
extern __declspec(dllexport) 
void QBrush_OperatorAssignWithColor(QBrush* self, QColor* color);
extern __declspec(dllexport) 
void QBrush_OperatorAssign2(QBrush* self, int color);
extern __declspec(dllexport) 
int QBrush_Style(const QBrush* self);
extern __declspec(dllexport) 
void QBrush_SetStyle(QBrush* self, int style);
extern __declspec(dllexport) 
QTransform* QBrush_Transform(const QBrush* self);
extern __declspec(dllexport) 
void QBrush_SetTransform(QBrush* self, QTransform* transform);
extern __declspec(dllexport) 
QPixmap* QBrush_Texture(const QBrush* self);
extern __declspec(dllexport) 
void QBrush_SetTexture(QBrush* self, QPixmap* pixmap);
extern __declspec(dllexport) 
QImage* QBrush_TextureImage(const QBrush* self);
extern __declspec(dllexport) 
void QBrush_SetTextureImage(QBrush* self, QImage* image);
extern __declspec(dllexport) 
QColor* QBrush_Color(const QBrush* self);
extern __declspec(dllexport) 
void QBrush_SetColor(QBrush* self, QColor* color);
extern __declspec(dllexport) 
void QBrush_SetColorWithColor(QBrush* self, int color);
extern __declspec(dllexport) 
QGradient* QBrush_Gradient(const QBrush* self);
extern __declspec(dllexport) 
bool QBrush_IsOpaque(const QBrush* self);
extern __declspec(dllexport) 
bool QBrush_OperatorEqual(const QBrush* self, QBrush* b);
extern __declspec(dllexport) 
bool QBrush_OperatorNotEqual(const QBrush* self, QBrush* b);
extern __declspec(dllexport) 
bool QBrush_IsDetached(const QBrush* self);
extern __declspec(dllexport) 
DataPtr* QBrush_DataPtr(QBrush* self);
extern __declspec(dllexport) 
void QBrush_Delete(QBrush* self, bool isSubclass);

extern __declspec(dllexport) 
QBrushData* QBrushData_new(QBrushData* param1);
extern __declspec(dllexport) 
void QBrushData_OperatorAssign(QBrushData* self, QBrushData* param1);
extern __declspec(dllexport) 
void QBrushData_Delete(QBrushData* self, bool isSubclass);

extern __declspec(dllexport) 
QGradient* QGradient_new();
extern __declspec(dllexport) 
QGradient* QGradient_new2(Preset param1);
extern __declspec(dllexport) 
QGradient* QGradient_new3(QGradient* param1);
extern __declspec(dllexport) 
Type QGradient_Type(const QGradient* self);
extern __declspec(dllexport) 
void QGradient_SetSpread(QGradient* self, Spread spread);
extern __declspec(dllexport) 
Spread QGradient_Spread(const QGradient* self);
extern __declspec(dllexport) 
void QGradient_SetColorAt(QGradient* self, double pos, QColor* color);
extern __declspec(dllexport) 
void QGradient_SetStops(QGradient* self, struct miqt_array /* of struct miqt_map  tuple of double and QColor*   */  stops);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_map  tuple of double and QColor*   */  QGradient_Stops(const QGradient* self);
extern __declspec(dllexport) 
CoordinateMode QGradient_CoordinateMode(const QGradient* self);
extern __declspec(dllexport) 
void QGradient_SetCoordinateMode(QGradient* self, CoordinateMode mode);
extern __declspec(dllexport) 
InterpolationMode QGradient_InterpolationMode(const QGradient* self);
extern __declspec(dllexport) 
void QGradient_SetInterpolationMode(QGradient* self, InterpolationMode mode);
extern __declspec(dllexport) 
bool QGradient_OperatorEqual(const QGradient* self, QGradient* gradient);
extern __declspec(dllexport) 
bool QGradient_OperatorNotEqual(const QGradient* self, QGradient* other);
extern __declspec(dllexport) 
void QGradient_Delete(QGradient* self, bool isSubclass);

extern __declspec(dllexport) 
QLinearGradient* QLinearGradient_new();
extern __declspec(dllexport) 
QLinearGradient* QLinearGradient_new2(QPointF* start, QPointF* finalStop);
extern __declspec(dllexport) 
QLinearGradient* QLinearGradient_new3(double xStart, double yStart, double xFinalStop, double yFinalStop);
extern __declspec(dllexport) 
QLinearGradient* QLinearGradient_new4(QLinearGradient* param1);
extern __declspec(dllexport) 
void QLinearGradient_virtbase(QLinearGradient* src
, QGradient** outptr_QGradient
);
extern __declspec(dllexport) 
QPointF* QLinearGradient_Start(const QLinearGradient* self);
extern __declspec(dllexport) 
void QLinearGradient_SetStart(QLinearGradient* self, QPointF* start);
extern __declspec(dllexport) 
void QLinearGradient_SetStart2(QLinearGradient* self, double x, double y);
extern __declspec(dllexport) 
QPointF* QLinearGradient_FinalStop(const QLinearGradient* self);
extern __declspec(dllexport) 
void QLinearGradient_SetFinalStop(QLinearGradient* self, QPointF* stop);
extern __declspec(dllexport) 
void QLinearGradient_SetFinalStop2(QLinearGradient* self, double x, double y);
extern __declspec(dllexport) 
void QLinearGradient_Delete(QLinearGradient* self, bool isSubclass);

extern __declspec(dllexport) 
QRadialGradient* QRadialGradient_new();
extern __declspec(dllexport) 
QRadialGradient* QRadialGradient_new2(QPointF* center, double radius, QPointF* focalPoint);
extern __declspec(dllexport) 
QRadialGradient* QRadialGradient_new3(double cx, double cy, double radius, double fx, double fy);
extern __declspec(dllexport) 
QRadialGradient* QRadialGradient_new4(QPointF* center, double radius);
extern __declspec(dllexport) 
QRadialGradient* QRadialGradient_new5(double cx, double cy, double radius);
extern __declspec(dllexport) 
QRadialGradient* QRadialGradient_new6(QPointF* center, double centerRadius, QPointF* focalPoint, double focalRadius);
extern __declspec(dllexport) 
QRadialGradient* QRadialGradient_new7(double cx, double cy, double centerRadius, double fx, double fy, double focalRadius);
extern __declspec(dllexport) 
QRadialGradient* QRadialGradient_new8(QRadialGradient* param1);
extern __declspec(dllexport) 
void QRadialGradient_virtbase(QRadialGradient* src
, QGradient** outptr_QGradient
);
extern __declspec(dllexport) 
QPointF* QRadialGradient_Center(const QRadialGradient* self);
extern __declspec(dllexport) 
void QRadialGradient_SetCenter(QRadialGradient* self, QPointF* center);
extern __declspec(dllexport) 
void QRadialGradient_SetCenter2(QRadialGradient* self, double x, double y);
extern __declspec(dllexport) 
QPointF* QRadialGradient_FocalPoint(const QRadialGradient* self);
extern __declspec(dllexport) 
void QRadialGradient_SetFocalPoint(QRadialGradient* self, QPointF* focalPoint);
extern __declspec(dllexport) 
void QRadialGradient_SetFocalPoint2(QRadialGradient* self, double x, double y);
extern __declspec(dllexport) 
double QRadialGradient_Radius(const QRadialGradient* self);
extern __declspec(dllexport) 
void QRadialGradient_SetRadius(QRadialGradient* self, double radius);
extern __declspec(dllexport) 
double QRadialGradient_CenterRadius(const QRadialGradient* self);
extern __declspec(dllexport) 
void QRadialGradient_SetCenterRadius(QRadialGradient* self, double radius);
extern __declspec(dllexport) 
double QRadialGradient_FocalRadius(const QRadialGradient* self);
extern __declspec(dllexport) 
void QRadialGradient_SetFocalRadius(QRadialGradient* self, double radius);
extern __declspec(dllexport) 
void QRadialGradient_Delete(QRadialGradient* self, bool isSubclass);

extern __declspec(dllexport) 
QConicalGradient* QConicalGradient_new();
extern __declspec(dllexport) 
QConicalGradient* QConicalGradient_new2(QPointF* center, double startAngle);
extern __declspec(dllexport) 
QConicalGradient* QConicalGradient_new3(double cx, double cy, double startAngle);
extern __declspec(dllexport) 
QConicalGradient* QConicalGradient_new4(QConicalGradient* param1);
extern __declspec(dllexport) 
void QConicalGradient_virtbase(QConicalGradient* src
, QGradient** outptr_QGradient
);
extern __declspec(dllexport) 
QPointF* QConicalGradient_Center(const QConicalGradient* self);
extern __declspec(dllexport) 
void QConicalGradient_SetCenter(QConicalGradient* self, QPointF* center);
extern __declspec(dllexport) 
void QConicalGradient_SetCenter2(QConicalGradient* self, double x, double y);
extern __declspec(dllexport) 
double QConicalGradient_Angle(const QConicalGradient* self);
extern __declspec(dllexport) 
void QConicalGradient_SetAngle(QConicalGradient* self, double angle);
extern __declspec(dllexport) 
void QConicalGradient_Delete(QConicalGradient* self, bool isSubclass);

extern __declspec(dllexport) 
QGradient__QGradientData* QGradient__QGradientData_new(const QGradientData* param1);
extern __declspec(dllexport) 
void QGradient__QGradientData_Delete(QGradient__QGradientData* self, bool isSubclass);

}
