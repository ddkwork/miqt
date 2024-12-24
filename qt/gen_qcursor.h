#pragma once
#ifndef MIQT_QT_GEN_QCURSOR_H
#define MIQT_QT_GEN_QCURSOR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QBitmap;
class QCursor;
class QPixmap;
class QPoint;
class QScreen;
class _GUID;
class type_info;
#else
typedef struct QBitmap QBitmap;
typedef struct QCursor QCursor;
typedef struct QPixmap QPixmap;
typedef struct QPoint QPoint;
typedef struct QScreen QScreen;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QCursor* QCursor_new();
extern __declspec(dllexport) QCursor* QCursor_new2(int shape);
extern __declspec(dllexport) QCursor* QCursor_new3(QBitmap* bitmap, QBitmap* mask);
extern __declspec(dllexport) QCursor* QCursor_new4(QPixmap* pixmap);
extern __declspec(dllexport) QCursor* QCursor_new5(QCursor* cursor);
extern __declspec(dllexport) QCursor* QCursor_new6(QBitmap* bitmap, QBitmap* mask, int hotX);
extern __declspec(dllexport) QCursor* QCursor_new7(QBitmap* bitmap, QBitmap* mask, int hotX, int hotY);
extern __declspec(dllexport) QCursor* QCursor_new8(QPixmap* pixmap, int hotX);
extern __declspec(dllexport) QCursor* QCursor_new9(QPixmap* pixmap, int hotX, int hotY);
extern __declspec(dllexport) void QCursor_OperatorAssign(QCursor* self, QCursor* cursor);
extern __declspec(dllexport) void QCursor_Swap(QCursor* self, QCursor* other);
extern __declspec(dllexport) int QCursor_Shape(const QCursor* self);
extern __declspec(dllexport) void QCursor_SetShape(QCursor* self, int newShape);
extern __declspec(dllexport) QBitmap* QCursor_Bitmap(const QCursor* self, int param1);
extern __declspec(dllexport) QBitmap* QCursor_Mask(const QCursor* self, int param1);
extern __declspec(dllexport) QBitmap* QCursor_Bitmap2(const QCursor* self);
extern __declspec(dllexport) QBitmap* QCursor_Mask2(const QCursor* self);
extern __declspec(dllexport) QPixmap* QCursor_Pixmap(const QCursor* self);
extern __declspec(dllexport) QPoint* QCursor_HotSpot(const QCursor* self);
extern __declspec(dllexport) QPoint* QCursor_Pos();
extern __declspec(dllexport) QPoint* QCursor_PosWithScreen(QScreen* screen);
extern __declspec(dllexport) void QCursor_SetPos(int x, int y);
extern __declspec(dllexport) void QCursor_SetPos2(QScreen* screen, int x, int y);
extern __declspec(dllexport) void QCursor_SetPosWithQPoint(QPoint* p);
extern __declspec(dllexport) void QCursor_SetPos3(QScreen* screen, QPoint* p);
extern __declspec(dllexport) void QCursor_Delete(QCursor* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
