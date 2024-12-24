#pragma once
#ifndef MIQT_QT_GEN_QPAGESIZE_H
#define MIQT_QT_GEN_QPAGESIZE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QPageSize QPageSize;
typedef struct QRect QRect;
typedef struct QRectF QRectF;
typedef struct QSize QSize;
typedef struct QSizeF QSizeF;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QPageSize* QPageSize_new();
extern __declspec(dllexport) 
QPageSize* QPageSize_new2(PageSizeId pageSizeId);
extern __declspec(dllexport) 
QPageSize* QPageSize_new3(QSize* pointSize);
extern __declspec(dllexport) 
QPageSize* QPageSize_new4(QSizeF* size, Unit units);
extern __declspec(dllexport) 
QPageSize* QPageSize_new5(QPageSize* other);
extern __declspec(dllexport) 
QPageSize* QPageSize_new6(QSize* pointSize, struct miqt_string name);
extern __declspec(dllexport) 
QPageSize* QPageSize_new7(QSize* pointSize, struct miqt_string name, SizeMatchPolicy matchPolicy);
extern __declspec(dllexport) 
QPageSize* QPageSize_new8(QSizeF* size, Unit units, struct miqt_string name);
extern __declspec(dllexport) 
QPageSize* QPageSize_new9(QSizeF* size, Unit units, struct miqt_string name, SizeMatchPolicy matchPolicy);
extern __declspec(dllexport) 
void QPageSize_OperatorAssign(QPageSize* self, QPageSize* other);
extern __declspec(dllexport) 
void QPageSize_Swap(QPageSize* self, QPageSize* other);
extern __declspec(dllexport) 
bool QPageSize_IsEquivalentTo(const QPageSize* self, QPageSize* other);
extern __declspec(dllexport) 
bool QPageSize_IsValid(const QPageSize* self);
extern __declspec(dllexport) 
struct miqt_string QPageSize_Key(const QPageSize* self);
extern __declspec(dllexport) 
struct miqt_string QPageSize_Name(const QPageSize* self);
extern __declspec(dllexport) 
PageSizeId QPageSize_Id(const QPageSize* self);
extern __declspec(dllexport) 
int QPageSize_WindowsId(const QPageSize* self);
extern __declspec(dllexport) 
QSizeF* QPageSize_DefinitionSize(const QPageSize* self);
extern __declspec(dllexport) 
Unit QPageSize_DefinitionUnits(const QPageSize* self);
extern __declspec(dllexport) 
QSizeF* QPageSize_Size(const QPageSize* self, Unit units);
extern __declspec(dllexport) 
QSize* QPageSize_SizePoints(const QPageSize* self);
extern __declspec(dllexport) 
QSize* QPageSize_SizePixels(const QPageSize* self, int resolution);
extern __declspec(dllexport) 
QRectF* QPageSize_Rect(const QPageSize* self, Unit units);
extern __declspec(dllexport) 
QRect* QPageSize_RectPoints(const QPageSize* self);
extern __declspec(dllexport) 
QRect* QPageSize_RectPixels(const QPageSize* self, int resolution);
extern __declspec(dllexport) 
struct miqt_string QPageSize_KeyWithPageSizeId(PageSizeId pageSizeId);
extern __declspec(dllexport) 
struct miqt_string QPageSize_NameWithPageSizeId(PageSizeId pageSizeId);
extern __declspec(dllexport) 
PageSizeId QPageSize_IdWithPointSize(QSize* pointSize);
extern __declspec(dllexport) 
PageSizeId QPageSize_Id2(QSizeF* size, Unit units);
extern __declspec(dllexport) 
PageSizeId QPageSize_IdWithWindowsId(int windowsId);
extern __declspec(dllexport) 
int QPageSize_WindowsIdWithPageSizeId(PageSizeId pageSizeId);
extern __declspec(dllexport) 
QSizeF* QPageSize_DefinitionSizeWithPageSizeId(PageSizeId pageSizeId);
extern __declspec(dllexport) 
Unit QPageSize_DefinitionUnitsWithPageSizeId(PageSizeId pageSizeId);
extern __declspec(dllexport) 
QSizeF* QPageSize_Size2(PageSizeId pageSizeId, Unit units);
extern __declspec(dllexport) 
QSize* QPageSize_SizePointsWithPageSizeId(PageSizeId pageSizeId);
extern __declspec(dllexport) 
QSize* QPageSize_SizePixels2(PageSizeId pageSizeId, int resolution);
extern __declspec(dllexport) 
PageSizeId QPageSize_Id22(QSize* pointSize, SizeMatchPolicy matchPolicy);
extern __declspec(dllexport) 
PageSizeId QPageSize_Id3(QSizeF* size, Unit units, SizeMatchPolicy matchPolicy);
extern __declspec(dllexport) 
void QPageSize_Delete(QPageSize* self, bool isSubclass);

}
