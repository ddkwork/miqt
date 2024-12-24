#pragma once
#ifndef MIQT_QT_GEN_QPAINTERPATH_H
#define MIQT_QT_GEN_QPAINTERPATH_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QPainterPath__Element)
typedef QPainterPath::Element QPainterPath__Element;
typedef struct QFont QFont;
typedef struct QPainterPath QPainterPath;
typedef struct QPainterPath__Element QPainterPath__Element;
typedef struct QPainterPathStroker QPainterPathStroker;
typedef struct QPen QPen;
typedef struct QPointF QPointF;
typedef struct QRectF QRectF;
typedef struct QRegion QRegion;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QPainterPath* QPainterPath_new();
extern __declspec(dllexport) 
QPainterPath* QPainterPath_new2(QPointF* startPoint);
extern __declspec(dllexport) 
QPainterPath* QPainterPath_new3(QPainterPath* other);
extern __declspec(dllexport) 
void QPainterPath_OperatorAssign(QPainterPath* self, QPainterPath* other);
extern __declspec(dllexport) 
void QPainterPath_Swap(QPainterPath* self, QPainterPath* other);
extern __declspec(dllexport) 
void QPainterPath_Clear(QPainterPath* self);
extern __declspec(dllexport) 
void QPainterPath_Reserve(QPainterPath* self, int size);
extern __declspec(dllexport) 
int QPainterPath_Capacity(const QPainterPath* self);
extern __declspec(dllexport) 
void QPainterPath_CloseSubpath(QPainterPath* self);
extern __declspec(dllexport) 
void QPainterPath_MoveTo(QPainterPath* self, QPointF* p);
extern __declspec(dllexport) 
void QPainterPath_MoveTo2(QPainterPath* self, double x, double y);
extern __declspec(dllexport) 
void QPainterPath_LineTo(QPainterPath* self, QPointF* p);
extern __declspec(dllexport) 
void QPainterPath_LineTo2(QPainterPath* self, double x, double y);
extern __declspec(dllexport) 
void QPainterPath_ArcMoveTo(QPainterPath* self, QRectF* rect, double angle);
extern __declspec(dllexport) 
void QPainterPath_ArcMoveTo2(QPainterPath* self, double x, double y, double w, double h, double angle);
extern __declspec(dllexport) 
void QPainterPath_ArcTo(QPainterPath* self, QRectF* rect, double startAngle, double arcLength);
extern __declspec(dllexport) 
void QPainterPath_ArcTo2(QPainterPath* self, double x, double y, double w, double h, double startAngle, double arcLength);
extern __declspec(dllexport) 
void QPainterPath_CubicTo(QPainterPath* self, QPointF* ctrlPt1, QPointF* ctrlPt2, QPointF* endPt);
extern __declspec(dllexport) 
void QPainterPath_CubicTo2(QPainterPath* self, double ctrlPt1x, double ctrlPt1y, double ctrlPt2x, double ctrlPt2y, double endPtx, double endPty);
extern __declspec(dllexport) 
void QPainterPath_QuadTo(QPainterPath* self, QPointF* ctrlPt, QPointF* endPt);
extern __declspec(dllexport) 
void QPainterPath_QuadTo2(QPainterPath* self, double ctrlPtx, double ctrlPty, double endPtx, double endPty);
extern __declspec(dllexport) 
QPointF* QPainterPath_CurrentPosition(const QPainterPath* self);
extern __declspec(dllexport) 
void QPainterPath_AddRect(QPainterPath* self, QRectF* rect);
extern __declspec(dllexport) 
void QPainterPath_AddRect2(QPainterPath* self, double x, double y, double w, double h);
extern __declspec(dllexport) 
void QPainterPath_AddEllipse(QPainterPath* self, QRectF* rect);
extern __declspec(dllexport) 
void QPainterPath_AddEllipse2(QPainterPath* self, double x, double y, double w, double h);
extern __declspec(dllexport) 
void QPainterPath_AddEllipse3(QPainterPath* self, QPointF* center, double rx, double ry);
extern __declspec(dllexport) 
void QPainterPath_AddText(QPainterPath* self, QPointF* point, QFont* f, struct miqt_string text);
extern __declspec(dllexport) 
void QPainterPath_AddText2(QPainterPath* self, double x, double y, QFont* f, struct miqt_string text);
extern __declspec(dllexport) 
void QPainterPath_AddPath(QPainterPath* self, QPainterPath* path);
extern __declspec(dllexport) 
void QPainterPath_AddRegion(QPainterPath* self, QRegion* region);
extern __declspec(dllexport) 
void QPainterPath_AddRoundedRect(QPainterPath* self, QRectF* rect, double xRadius, double yRadius);
extern __declspec(dllexport) 
void QPainterPath_AddRoundedRect2(QPainterPath* self, double x, double y, double w, double h, double xRadius, double yRadius);
extern __declspec(dllexport) 
void QPainterPath_ConnectPath(QPainterPath* self, QPainterPath* path);
extern __declspec(dllexport) 
bool QPainterPath_Contains(const QPainterPath* self, QPointF* pt);
extern __declspec(dllexport) 
bool QPainterPath_ContainsWithRect(const QPainterPath* self, QRectF* rect);
extern __declspec(dllexport) 
bool QPainterPath_Intersects(const QPainterPath* self, QRectF* rect);
extern __declspec(dllexport) 
void QPainterPath_Translate(QPainterPath* self, double dx, double dy);
extern __declspec(dllexport) 
void QPainterPath_TranslateWithOffset(QPainterPath* self, QPointF* offset);
extern __declspec(dllexport) 
QPainterPath* QPainterPath_Translated(const QPainterPath* self, double dx, double dy);
extern __declspec(dllexport) 
QPainterPath* QPainterPath_TranslatedWithOffset(const QPainterPath* self, QPointF* offset);
extern __declspec(dllexport) 
QRectF* QPainterPath_BoundingRect(const QPainterPath* self);
extern __declspec(dllexport) 
QRectF* QPainterPath_ControlPointRect(const QPainterPath* self);
extern __declspec(dllexport) 
int QPainterPath_FillRule(const QPainterPath* self);
extern __declspec(dllexport) 
void QPainterPath_SetFillRule(QPainterPath* self, int fillRule);
extern __declspec(dllexport) 
bool QPainterPath_IsEmpty(const QPainterPath* self);
extern __declspec(dllexport) 
QPainterPath* QPainterPath_ToReversed(const QPainterPath* self);
extern __declspec(dllexport) 
int QPainterPath_ElementCount(const QPainterPath* self);
extern __declspec(dllexport) 
QPainterPath__Element* QPainterPath_ElementAt(const QPainterPath* self, int i);
extern __declspec(dllexport) 
void QPainterPath_SetElementPositionAt(QPainterPath* self, int i, double x, double y);
extern __declspec(dllexport) 
double QPainterPath_Length(const QPainterPath* self);
extern __declspec(dllexport) 
double QPainterPath_PercentAtLength(const QPainterPath* self, double t);
extern __declspec(dllexport) 
QPointF* QPainterPath_PointAtPercent(const QPainterPath* self, double t);
extern __declspec(dllexport) 
double QPainterPath_AngleAtPercent(const QPainterPath* self, double t);
extern __declspec(dllexport) 
double QPainterPath_SlopeAtPercent(const QPainterPath* self, double t);
extern __declspec(dllexport) 
bool QPainterPath_IntersectsWithQPainterPath(const QPainterPath* self, QPainterPath* p);
extern __declspec(dllexport) 
bool QPainterPath_ContainsWithQPainterPath(const QPainterPath* self, QPainterPath* p);
extern __declspec(dllexport) 
QPainterPath* QPainterPath_United(const QPainterPath* self, QPainterPath* r);
extern __declspec(dllexport) 
QPainterPath* QPainterPath_Intersected(const QPainterPath* self, QPainterPath* r);
extern __declspec(dllexport) 
QPainterPath* QPainterPath_Subtracted(const QPainterPath* self, QPainterPath* r);
extern __declspec(dllexport) 
QPainterPath* QPainterPath_Simplified(const QPainterPath* self);
extern __declspec(dllexport) 
bool QPainterPath_OperatorEqual(const QPainterPath* self, QPainterPath* other);
extern __declspec(dllexport) 
bool QPainterPath_OperatorNotEqual(const QPainterPath* self, QPainterPath* other);
extern __declspec(dllexport) 
QPainterPath* QPainterPath_OperatorBitwiseAnd(const QPainterPath* self, QPainterPath* other);
extern __declspec(dllexport) 
QPainterPath* QPainterPath_OperatorBitwiseOr(const QPainterPath* self, QPainterPath* other);
extern __declspec(dllexport) 
QPainterPath* QPainterPath_OperatorPlus(const QPainterPath* self, QPainterPath* other);
extern __declspec(dllexport) 
QPainterPath* QPainterPath_OperatorMinus(const QPainterPath* self, QPainterPath* other);
extern __declspec(dllexport) 
void QPainterPath_OperatorBitwiseAndAssign(QPainterPath* self, QPainterPath* other);
extern __declspec(dllexport) 
void QPainterPath_OperatorBitwiseOrAssign(QPainterPath* self, QPainterPath* other);
extern __declspec(dllexport) 
QPainterPath* QPainterPath_OperatorPlusAssign(QPainterPath* self, QPainterPath* other);
extern __declspec(dllexport) 
QPainterPath* QPainterPath_OperatorMinusAssign(QPainterPath* self, QPainterPath* other);
extern __declspec(dllexport) 
void QPainterPath_AddRoundedRect4(QPainterPath* self, QRectF* rect, double xRadius, double yRadius, int mode);
extern __declspec(dllexport) 
void QPainterPath_AddRoundedRect7(QPainterPath* self, double x, double y, double w, double h, double xRadius, double yRadius, int mode);
extern __declspec(dllexport) 
void QPainterPath_Delete(QPainterPath* self, bool isSubclass);

extern __declspec(dllexport) 
QPainterPathStroker* QPainterPathStroker_new();
extern __declspec(dllexport) 
QPainterPathStroker* QPainterPathStroker_new2(QPen* pen);
extern __declspec(dllexport) 
void QPainterPathStroker_SetWidth(QPainterPathStroker* self, double width);
extern __declspec(dllexport) 
double QPainterPathStroker_Width(const QPainterPathStroker* self);
extern __declspec(dllexport) 
void QPainterPathStroker_SetCapStyle(QPainterPathStroker* self, int style);
extern __declspec(dllexport) 
int QPainterPathStroker_CapStyle(const QPainterPathStroker* self);
extern __declspec(dllexport) 
void QPainterPathStroker_SetJoinStyle(QPainterPathStroker* self, int style);
extern __declspec(dllexport) 
int QPainterPathStroker_JoinStyle(const QPainterPathStroker* self);
extern __declspec(dllexport) 
void QPainterPathStroker_SetMiterLimit(QPainterPathStroker* self, double length);
extern __declspec(dllexport) 
double QPainterPathStroker_MiterLimit(const QPainterPathStroker* self);
extern __declspec(dllexport) 
void QPainterPathStroker_SetCurveThreshold(QPainterPathStroker* self, double threshold);
extern __declspec(dllexport) 
double QPainterPathStroker_CurveThreshold(const QPainterPathStroker* self);
extern __declspec(dllexport) 
void QPainterPathStroker_SetDashPattern(QPainterPathStroker* self, int dashPattern);
extern __declspec(dllexport) 
void QPainterPathStroker_SetDashPatternWithDashPattern(QPainterPathStroker* self, struct miqt_array /* of double */  dashPattern);
extern __declspec(dllexport) 
struct miqt_array /* of double */  QPainterPathStroker_DashPattern(const QPainterPathStroker* self);
extern __declspec(dllexport) 
void QPainterPathStroker_SetDashOffset(QPainterPathStroker* self, double offset);
extern __declspec(dllexport) 
double QPainterPathStroker_DashOffset(const QPainterPathStroker* self);
extern __declspec(dllexport) 
QPainterPath* QPainterPathStroker_CreateStroke(const QPainterPathStroker* self, QPainterPath* path);
extern __declspec(dllexport) 
void QPainterPathStroker_Delete(QPainterPathStroker* self, bool isSubclass);

extern __declspec(dllexport) 
QPainterPath__Element* QPainterPath__Element_new();
extern __declspec(dllexport) 
QPainterPath__Element* QPainterPath__Element_new2(const Element* param1);
extern __declspec(dllexport) 
bool QPainterPath__Element_IsMoveTo(const QPainterPath__Element* self);
extern __declspec(dllexport) 
bool QPainterPath__Element_IsLineTo(const QPainterPath__Element* self);
extern __declspec(dllexport) 
bool QPainterPath__Element_IsCurveTo(const QPainterPath__Element* self);
extern __declspec(dllexport) 
bool QPainterPath__Element_OperatorEqual(const QPainterPath__Element* self, const Element* e);
extern __declspec(dllexport) 
bool QPainterPath__Element_OperatorNotEqual(const QPainterPath__Element* self, const Element* e);
extern __declspec(dllexport) 
void QPainterPath__Element_Delete(QPainterPath__Element* self, bool isSubclass);

}
