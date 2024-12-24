#pragma once
#ifndef MIQT_QT_GEN_QBITMAP_H
#define MIQT_QT_GEN_QBITMAP_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QBitmap QBitmap;
typedef struct QImage QImage;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPixmap QPixmap;
typedef struct QSize QSize;
typedef struct QTransform QTransform;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QBitmap* QBitmap_new();
extern __declspec(dllexport) 
QBitmap* QBitmap_new2(QPixmap* param1);
extern __declspec(dllexport) 
QBitmap* QBitmap_new3(int w, int h);
extern __declspec(dllexport) 
QBitmap* QBitmap_new4(QSize* param1);
extern __declspec(dllexport) 
QBitmap* QBitmap_new5(struct miqt_string fileName);
extern __declspec(dllexport) 
QBitmap* QBitmap_new6(struct miqt_string fileName, const char* format);
extern __declspec(dllexport) 
void QBitmap_virtbase(QBitmap* src
, QPixmap** outptr_QPixmap
);
extern __declspec(dllexport) 
void QBitmap_OperatorAssign(QBitmap* self, QPixmap* param1);
extern __declspec(dllexport) 
void QBitmap_Swap(QBitmap* self, QBitmap* other);
extern __declspec(dllexport) 
void QBitmap_Clear(QBitmap* self);
extern __declspec(dllexport) 
QBitmap* QBitmap_FromImage(QImage* image);
extern __declspec(dllexport) 
QBitmap* QBitmap_FromData(QSize* size, const unsigned char* bits);
extern __declspec(dllexport) 
QBitmap* QBitmap_FromPixmap(QPixmap* pixmap);
extern __declspec(dllexport) 
QBitmap* QBitmap_Transformed(const QBitmap* self, QTransform* matrix);
extern __declspec(dllexport) 
void QBitmap_OperatorAssignWithQBitmap(QBitmap* self, QBitmap* param1);
extern __declspec(dllexport) 
QBitmap* QBitmap_FromImage2(QImage* image, int flags);
extern __declspec(dllexport) 
QBitmap* QBitmap_FromData3(QSize* size, const unsigned char* bits, int monoFormat);
extern __declspec(dllexport) 
void QBitmap_Delete(QBitmap* self, bool isSubclass);

}
