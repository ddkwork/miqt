#pragma once
#ifndef MIQT_QT_GEN_QGRAPHICSLAYOUT_H
#define MIQT_QT_GEN_QGRAPHICSLAYOUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QGraphicsLayout QGraphicsLayout;
typedef struct QGraphicsLayoutItem QGraphicsLayoutItem;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QGraphicsLayout* QGraphicsLayout_new();
extern __declspec(dllexport) 
QGraphicsLayout* QGraphicsLayout_new2(QGraphicsLayoutItem* parent);
extern __declspec(dllexport) 
void QGraphicsLayout_virtbase(QGraphicsLayout* src
, QGraphicsLayoutItem** outptr_QGraphicsLayoutItem
);
extern __declspec(dllexport) 
void QGraphicsLayout_SetContentsMargins(QGraphicsLayout* self, double left, double top, double right, double bottom);
extern __declspec(dllexport) 
void QGraphicsLayout_GetContentsMargins(const QGraphicsLayout* self, double* left, double* top, double* right, double* bottom);
extern __declspec(dllexport) 
void QGraphicsLayout_Activate(QGraphicsLayout* self);
extern __declspec(dllexport) 
bool QGraphicsLayout_IsActivated(const QGraphicsLayout* self);
extern __declspec(dllexport) 
void QGraphicsLayout_Invalidate(QGraphicsLayout* self);
extern __declspec(dllexport) 
void QGraphicsLayout_UpdateGeometry(QGraphicsLayout* self);
extern __declspec(dllexport) 
void QGraphicsLayout_WidgetEvent(QGraphicsLayout* self, QEvent* e);
extern __declspec(dllexport) 
int QGraphicsLayout_Count(const QGraphicsLayout* self);
extern __declspec(dllexport) 
QGraphicsLayoutItem* QGraphicsLayout_ItemAt(const QGraphicsLayout* self, int i);
extern __declspec(dllexport) 
void QGraphicsLayout_RemoveAt(QGraphicsLayout* self, int index);
extern __declspec(dllexport) 
void QGraphicsLayout_SetInstantInvalidatePropagation(bool enable);
extern __declspec(dllexport) 
bool QGraphicsLayout_InstantInvalidatePropagation();
extern __declspec(dllexport) 
void QGraphicsLayout_Delete(QGraphicsLayout* self, bool isSubclass);

}
