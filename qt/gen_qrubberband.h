#pragma once
#ifndef MIQT_QT_GEN_QRUBBERBAND_H
#define MIQT_QT_GEN_QRUBBERBAND_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMoveEvent QMoveEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QResizeEvent QResizeEvent;
typedef struct QRubberBand QRubberBand;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QStyleOptionRubberBand QStyleOptionRubberBand;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QRubberBand* QRubberBand_new(Shape param1);
extern __declspec(dllexport) 
QRubberBand* QRubberBand_new2(Shape param1, QWidget* param2);
extern __declspec(dllexport) 
void QRubberBand_virtbase(QRubberBand* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QRubberBand_MetaObject(const QRubberBand* self);
extern __declspec(dllexport) 
void* QRubberBand_Metacast(QRubberBand* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QRubberBand_Tr(const char* s);
extern __declspec(dllexport) 
Shape QRubberBand_Shape(const QRubberBand* self);
extern __declspec(dllexport) 
void QRubberBand_SetGeometry(QRubberBand* self, QRect* r);
extern __declspec(dllexport) 
void QRubberBand_SetGeometry2(QRubberBand* self, int x, int y, int w, int h);
extern __declspec(dllexport) 
void QRubberBand_Move(QRubberBand* self, int x, int y);
extern __declspec(dllexport) 
void QRubberBand_MoveWithQPoint(QRubberBand* self, QPoint* p);
extern __declspec(dllexport) 
void QRubberBand_Resize(QRubberBand* self, int w, int h);
extern __declspec(dllexport) 
void QRubberBand_ResizeWithQSize(QRubberBand* self, QSize* s);
extern __declspec(dllexport) 
bool QRubberBand_Event(QRubberBand* self, QEvent* e);
extern __declspec(dllexport) 
void QRubberBand_PaintEvent(QRubberBand* self, QPaintEvent* param1);
extern __declspec(dllexport) 
void QRubberBand_ChangeEvent(QRubberBand* self, QEvent* param1);
extern __declspec(dllexport) 
void QRubberBand_ShowEvent(QRubberBand* self, QShowEvent* param1);
extern __declspec(dllexport) 
void QRubberBand_ResizeEvent(QRubberBand* self, QResizeEvent* param1);
extern __declspec(dllexport) 
void QRubberBand_MoveEvent(QRubberBand* self, QMoveEvent* param1);
extern __declspec(dllexport) 
void QRubberBand_InitStyleOption(const QRubberBand* self, QStyleOptionRubberBand* option);
extern __declspec(dllexport) 
struct miqt_string QRubberBand_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QRubberBand_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QRubberBand_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QRubberBand_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QRubberBand_override_virtual_Metacast(void* self, intptr_t slot);
void* QRubberBand_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QRubberBand_Delete(QRubberBand* self, bool isSubclass);

}
