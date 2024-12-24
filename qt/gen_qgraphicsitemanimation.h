#pragma once
#ifndef MIQT_QT_GEN_QGRAPHICSITEMANIMATION_H
#define MIQT_QT_GEN_QGRAPHICSITEMANIMATION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QGraphicsItem QGraphicsItem;
typedef struct QGraphicsItemAnimation QGraphicsItemAnimation;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPointF QPointF;
typedef struct QTimeLine QTimeLine;
typedef struct QTransform QTransform;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QGraphicsItemAnimation* QGraphicsItemAnimation_new();
extern __declspec(dllexport) 
QGraphicsItemAnimation* QGraphicsItemAnimation_new2(QObject* parent);
extern __declspec(dllexport) 
void QGraphicsItemAnimation_virtbase(QGraphicsItemAnimation* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QGraphicsItemAnimation_MetaObject(const QGraphicsItemAnimation* self);
extern __declspec(dllexport) 
void* QGraphicsItemAnimation_Metacast(QGraphicsItemAnimation* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QGraphicsItemAnimation_Tr(const char* s);
extern __declspec(dllexport) 
QGraphicsItem* QGraphicsItemAnimation_Item(const QGraphicsItemAnimation* self);
extern __declspec(dllexport) 
void QGraphicsItemAnimation_SetItem(QGraphicsItemAnimation* self, QGraphicsItem* item);
extern __declspec(dllexport) 
QTimeLine* QGraphicsItemAnimation_TimeLine(const QGraphicsItemAnimation* self);
extern __declspec(dllexport) 
void QGraphicsItemAnimation_SetTimeLine(QGraphicsItemAnimation* self, QTimeLine* timeLine);
extern __declspec(dllexport) 
QPointF* QGraphicsItemAnimation_PosAt(const QGraphicsItemAnimation* self, double step);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_map  tuple of double and QPointF*   */  QGraphicsItemAnimation_PosList(const QGraphicsItemAnimation* self);
extern __declspec(dllexport) 
void QGraphicsItemAnimation_SetPosAt(QGraphicsItemAnimation* self, double step, QPointF* pos);
extern __declspec(dllexport) 
QTransform* QGraphicsItemAnimation_TransformAt(const QGraphicsItemAnimation* self, double step);
extern __declspec(dllexport) 
double QGraphicsItemAnimation_RotationAt(const QGraphicsItemAnimation* self, double step);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_map  tuple of double and double   */  QGraphicsItemAnimation_RotationList(const QGraphicsItemAnimation* self);
extern __declspec(dllexport) 
void QGraphicsItemAnimation_SetRotationAt(QGraphicsItemAnimation* self, double step, double angle);
extern __declspec(dllexport) 
double QGraphicsItemAnimation_XTranslationAt(const QGraphicsItemAnimation* self, double step);
extern __declspec(dllexport) 
double QGraphicsItemAnimation_YTranslationAt(const QGraphicsItemAnimation* self, double step);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_map  tuple of double and QPointF*   */  QGraphicsItemAnimation_TranslationList(const QGraphicsItemAnimation* self);
extern __declspec(dllexport) 
void QGraphicsItemAnimation_SetTranslationAt(QGraphicsItemAnimation* self, double step, double dx, double dy);
extern __declspec(dllexport) 
double QGraphicsItemAnimation_VerticalScaleAt(const QGraphicsItemAnimation* self, double step);
extern __declspec(dllexport) 
double QGraphicsItemAnimation_HorizontalScaleAt(const QGraphicsItemAnimation* self, double step);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_map  tuple of double and QPointF*   */  QGraphicsItemAnimation_ScaleList(const QGraphicsItemAnimation* self);
extern __declspec(dllexport) 
void QGraphicsItemAnimation_SetScaleAt(QGraphicsItemAnimation* self, double step, double sx, double sy);
extern __declspec(dllexport) 
double QGraphicsItemAnimation_VerticalShearAt(const QGraphicsItemAnimation* self, double step);
extern __declspec(dllexport) 
double QGraphicsItemAnimation_HorizontalShearAt(const QGraphicsItemAnimation* self, double step);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_map  tuple of double and QPointF*   */  QGraphicsItemAnimation_ShearList(const QGraphicsItemAnimation* self);
extern __declspec(dllexport) 
void QGraphicsItemAnimation_SetShearAt(QGraphicsItemAnimation* self, double step, double sh, double sv);
extern __declspec(dllexport) 
void QGraphicsItemAnimation_Clear(QGraphicsItemAnimation* self);
extern __declspec(dllexport) 
void QGraphicsItemAnimation_SetStep(QGraphicsItemAnimation* self, double x);
extern __declspec(dllexport) 
void QGraphicsItemAnimation_BeforeAnimationStep(QGraphicsItemAnimation* self, double step);
extern __declspec(dllexport) 
void QGraphicsItemAnimation_AfterAnimationStep(QGraphicsItemAnimation* self, double step);
extern __declspec(dllexport) 
struct miqt_string QGraphicsItemAnimation_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QGraphicsItemAnimation_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QGraphicsItemAnimation_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QGraphicsItemAnimation_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QGraphicsItemAnimation_override_virtual_Metacast(void* self, intptr_t slot);
void* QGraphicsItemAnimation_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QGraphicsItemAnimation_Delete(QGraphicsItemAnimation* self, bool isSubclass);

}
