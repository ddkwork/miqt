#pragma once
#ifndef MIQT_QT_GEN_QSCREEN_H
#define MIQT_QT_GEN_QSCREEN_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPixmap QPixmap;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QScreen QScreen;
typedef struct QSize QSize;
typedef struct QSizeF QSizeF;
typedef struct QTransform QTransform;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
void QScreen_virtbase(QScreen* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QScreen_MetaObject(const QScreen* self);
extern __declspec(dllexport) 
void* QScreen_Metacast(QScreen* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QScreen_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_string QScreen_Name(const QScreen* self);
extern __declspec(dllexport) 
struct miqt_string QScreen_Manufacturer(const QScreen* self);
extern __declspec(dllexport) 
struct miqt_string QScreen_Model(const QScreen* self);
extern __declspec(dllexport) 
struct miqt_string QScreen_SerialNumber(const QScreen* self);
extern __declspec(dllexport) 
int QScreen_Depth(const QScreen* self);
extern __declspec(dllexport) 
QSize* QScreen_Size(const QScreen* self);
extern __declspec(dllexport) 
QRect* QScreen_Geometry(const QScreen* self);
extern __declspec(dllexport) 
QSizeF* QScreen_PhysicalSize(const QScreen* self);
extern __declspec(dllexport) 
double QScreen_PhysicalDotsPerInchX(const QScreen* self);
extern __declspec(dllexport) 
double QScreen_PhysicalDotsPerInchY(const QScreen* self);
extern __declspec(dllexport) 
double QScreen_PhysicalDotsPerInch(const QScreen* self);
extern __declspec(dllexport) 
double QScreen_LogicalDotsPerInchX(const QScreen* self);
extern __declspec(dllexport) 
double QScreen_LogicalDotsPerInchY(const QScreen* self);
extern __declspec(dllexport) 
double QScreen_LogicalDotsPerInch(const QScreen* self);
extern __declspec(dllexport) 
double QScreen_DevicePixelRatio(const QScreen* self);
extern __declspec(dllexport) 
QSize* QScreen_AvailableSize(const QScreen* self);
extern __declspec(dllexport) 
QRect* QScreen_AvailableGeometry(const QScreen* self);
extern __declspec(dllexport) 
struct miqt_array /* of QScreen* */  QScreen_VirtualSiblings(const QScreen* self);
extern __declspec(dllexport) 
QScreen* QScreen_VirtualSiblingAt(QScreen* self, QPoint* point);
extern __declspec(dllexport) 
QSize* QScreen_VirtualSize(const QScreen* self);
extern __declspec(dllexport) 
QRect* QScreen_VirtualGeometry(const QScreen* self);
extern __declspec(dllexport) 
QSize* QScreen_AvailableVirtualSize(const QScreen* self);
extern __declspec(dllexport) 
QRect* QScreen_AvailableVirtualGeometry(const QScreen* self);
extern __declspec(dllexport) 
int QScreen_PrimaryOrientation(const QScreen* self);
extern __declspec(dllexport) 
int QScreen_Orientation(const QScreen* self);
extern __declspec(dllexport) 
int QScreen_NativeOrientation(const QScreen* self);
extern __declspec(dllexport) 
int QScreen_AngleBetween(const QScreen* self, int a, int b);
extern __declspec(dllexport) 
QTransform* QScreen_TransformBetween(const QScreen* self, int a, int b, QRect* target);
extern __declspec(dllexport) 
QRect* QScreen_MapBetween(const QScreen* self, int a, int b, QRect* rect);
extern __declspec(dllexport) 
bool QScreen_IsPortrait(const QScreen* self, int orientation);
extern __declspec(dllexport) 
bool QScreen_IsLandscape(const QScreen* self, int orientation);
extern __declspec(dllexport) 
QPixmap* QScreen_GrabWindow(QScreen* self);
extern __declspec(dllexport) 
double QScreen_RefreshRate(const QScreen* self);
extern __declspec(dllexport) 
void QScreen_GeometryChanged(QScreen* self, QRect* geometry);
void QScreen_connect_GeometryChanged(QScreen* self, intptr_t slot);
extern __declspec(dllexport) 
void QScreen_AvailableGeometryChanged(QScreen* self, QRect* geometry);
void QScreen_connect_AvailableGeometryChanged(QScreen* self, intptr_t slot);
extern __declspec(dllexport) 
void QScreen_PhysicalSizeChanged(QScreen* self, QSizeF* size);
void QScreen_connect_PhysicalSizeChanged(QScreen* self, intptr_t slot);
extern __declspec(dllexport) 
void QScreen_PhysicalDotsPerInchChanged(QScreen* self, double dpi);
void QScreen_connect_PhysicalDotsPerInchChanged(QScreen* self, intptr_t slot);
extern __declspec(dllexport) 
void QScreen_LogicalDotsPerInchChanged(QScreen* self, double dpi);
void QScreen_connect_LogicalDotsPerInchChanged(QScreen* self, intptr_t slot);
extern __declspec(dllexport) 
void QScreen_VirtualGeometryChanged(QScreen* self, QRect* rect);
void QScreen_connect_VirtualGeometryChanged(QScreen* self, intptr_t slot);
extern __declspec(dllexport) 
void QScreen_PrimaryOrientationChanged(QScreen* self, int orientation);
void QScreen_connect_PrimaryOrientationChanged(QScreen* self, intptr_t slot);
extern __declspec(dllexport) 
void QScreen_OrientationChanged(QScreen* self, int orientation);
void QScreen_connect_OrientationChanged(QScreen* self, intptr_t slot);
extern __declspec(dllexport) 
void QScreen_RefreshRateChanged(QScreen* self, double refreshRate);
void QScreen_connect_RefreshRateChanged(QScreen* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QScreen_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QScreen_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
QPixmap* QScreen_GrabWindow1(QScreen* self, uintptr_t window);
extern __declspec(dllexport) 
QPixmap* QScreen_GrabWindow2(QScreen* self, uintptr_t window, int x);
extern __declspec(dllexport) 
QPixmap* QScreen_GrabWindow3(QScreen* self, uintptr_t window, int x, int y);
extern __declspec(dllexport) 
QPixmap* QScreen_GrabWindow4(QScreen* self, uintptr_t window, int x, int y, int w);
extern __declspec(dllexport) 
QPixmap* QScreen_GrabWindow5(QScreen* self, uintptr_t window, int x, int y, int w, int h);
extern __declspec(dllexport) 
void QScreen_Delete(QScreen* self, bool isSubclass);

}
