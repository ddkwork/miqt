#pragma once
#ifndef MIQT_QT_GEN_QGRAPHICSLINEARLAYOUT_H
#define MIQT_QT_GEN_QGRAPHICSLINEARLAYOUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QGraphicsLayout QGraphicsLayout;
typedef struct QGraphicsLayoutItem QGraphicsLayoutItem;
typedef struct QGraphicsLinearLayout QGraphicsLinearLayout;
typedef struct QRectF QRectF;
typedef struct QSizeF QSizeF;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QGraphicsLinearLayout* QGraphicsLinearLayout_new();
extern __declspec(dllexport) 
QGraphicsLinearLayout* QGraphicsLinearLayout_new2(int orientation);
extern __declspec(dllexport) 
QGraphicsLinearLayout* QGraphicsLinearLayout_new3(QGraphicsLayoutItem* parent);
extern __declspec(dllexport) 
QGraphicsLinearLayout* QGraphicsLinearLayout_new4(int orientation, QGraphicsLayoutItem* parent);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_virtbase(QGraphicsLinearLayout* src
, QGraphicsLayout** outptr_QGraphicsLayout
);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_SetOrientation(QGraphicsLinearLayout* self, int orientation);
extern __declspec(dllexport) 
int QGraphicsLinearLayout_Orientation(const QGraphicsLinearLayout* self);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_AddItem(QGraphicsLinearLayout* self, QGraphicsLayoutItem* item);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_AddStretch(QGraphicsLinearLayout* self);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_InsertItem(QGraphicsLinearLayout* self, int index, QGraphicsLayoutItem* item);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_InsertStretch(QGraphicsLinearLayout* self, int index);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_RemoveItem(QGraphicsLinearLayout* self, QGraphicsLayoutItem* item);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_RemoveAt(QGraphicsLinearLayout* self, int index);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_SetSpacing(QGraphicsLinearLayout* self, double spacing);
extern __declspec(dllexport) 
double QGraphicsLinearLayout_Spacing(const QGraphicsLinearLayout* self);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_SetItemSpacing(QGraphicsLinearLayout* self, int index, double spacing);
extern __declspec(dllexport) 
double QGraphicsLinearLayout_ItemSpacing(const QGraphicsLinearLayout* self, int index);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_SetStretchFactor(QGraphicsLinearLayout* self, QGraphicsLayoutItem* item, int stretch);
extern __declspec(dllexport) 
int QGraphicsLinearLayout_StretchFactor(const QGraphicsLinearLayout* self, QGraphicsLayoutItem* item);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_SetAlignment(QGraphicsLinearLayout* self, QGraphicsLayoutItem* item, int alignment);
extern __declspec(dllexport) 
int QGraphicsLinearLayout_Alignment(const QGraphicsLinearLayout* self, QGraphicsLayoutItem* item);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_SetGeometry(QGraphicsLinearLayout* self, QRectF* rect);
extern __declspec(dllexport) 
int QGraphicsLinearLayout_Count(const QGraphicsLinearLayout* self);
extern __declspec(dllexport) 
QGraphicsLayoutItem* QGraphicsLinearLayout_ItemAt(const QGraphicsLinearLayout* self, int index);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_Invalidate(QGraphicsLinearLayout* self);
extern __declspec(dllexport) 
QSizeF* QGraphicsLinearLayout_SizeHint(const QGraphicsLinearLayout* self, int which, QSizeF* constraint);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_Dump(const QGraphicsLinearLayout* self);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_AddStretch1(QGraphicsLinearLayout* self, int stretch);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_InsertStretch2(QGraphicsLinearLayout* self, int index, int stretch);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_Dump1(const QGraphicsLinearLayout* self, int indent);
extern __declspec(dllexport) 
void QGraphicsLinearLayout_Delete(QGraphicsLinearLayout* self, bool isSubclass);

}
