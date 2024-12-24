#pragma once
#ifndef MIQT_QT_GEN_QSTYLEPAINTER_H
#define MIQT_QT_GEN_QSTYLEPAINTER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QPaintDevice QPaintDevice;
typedef struct QPainter QPainter;
typedef struct QPalette QPalette;
typedef struct QPixmap QPixmap;
typedef struct QRect QRect;
typedef struct QStyle QStyle;
typedef struct QStyleOption QStyleOption;
typedef struct QStyleOptionComplex QStyleOptionComplex;
typedef struct QStylePainter QStylePainter;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QStylePainter* QStylePainter_new(QWidget* w);
extern __declspec(dllexport) QStylePainter* QStylePainter_new2();
extern __declspec(dllexport) QStylePainter* QStylePainter_new3(QPaintDevice* pd, QWidget* w);
extern __declspec(dllexport) void QStylePainter_virtbase(QStylePainter* src, QPainter** outptr_QPainter);
extern __declspec(dllexport) bool QStylePainter_Begin(QStylePainter* self, QWidget* w);
extern __declspec(dllexport) bool QStylePainter_Begin2(QStylePainter* self, QPaintDevice* pd, QWidget* w);
extern __declspec(dllexport) void QStylePainter_DrawPrimitive(QStylePainter* self, int pe, QStyleOption* opt);
extern __declspec(dllexport) void QStylePainter_DrawControl(QStylePainter* self, int ce, QStyleOption* opt);
extern __declspec(dllexport) void QStylePainter_DrawComplexControl(QStylePainter* self, int cc, QStyleOptionComplex* opt);
extern __declspec(dllexport) void QStylePainter_DrawItemText(QStylePainter* self, QRect* r, int flags, QPalette* pal, bool enabled, struct miqt_string text);
extern __declspec(dllexport) void QStylePainter_DrawItemPixmap(QStylePainter* self, QRect* r, int flags, QPixmap* pixmap);
extern __declspec(dllexport) QStyle* QStylePainter_Style(const QStylePainter* self);
extern __declspec(dllexport) void QStylePainter_DrawItemText6(QStylePainter* self, QRect* r, int flags, QPalette* pal, bool enabled, struct miqt_string text, int textRole);
extern __declspec(dllexport) void QStylePainter_Delete(QStylePainter* self, bool isSubclass);

} 
