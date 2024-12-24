#pragma once
#ifndef MIQT_QT_GEN_QSTYLE_H
#define MIQT_QT_GEN_QSTYLE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QApplication QApplication;
typedef struct QFontMetrics QFontMetrics;
typedef struct QIcon QIcon;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPainter QPainter;
typedef struct QPalette QPalette;
typedef struct QPixmap QPixmap;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QStyle QStyle;
typedef struct QStyleHintReturn QStyleHintReturn;
typedef struct QStyleOption QStyleOption;
typedef struct QStyleOptionComplex QStyleOptionComplex;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QStyle* QStyle_new();
extern __declspec(dllexport) 
void QStyle_virtbase(QStyle* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QStyle_MetaObject(const QStyle* self);
extern __declspec(dllexport) 
void* QStyle_Metacast(QStyle* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QStyle_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_string QStyle_Name(const QStyle* self);
extern __declspec(dllexport) 
void QStyle_Polish(QStyle* self, QWidget* widget);
extern __declspec(dllexport) 
void QStyle_Unpolish(QStyle* self, QWidget* widget);
extern __declspec(dllexport) 
void QStyle_PolishWithApplication(QStyle* self, QApplication* application);
extern __declspec(dllexport) 
void QStyle_UnpolishWithApplication(QStyle* self, QApplication* application);
extern __declspec(dllexport) 
void QStyle_PolishWithPalette(QStyle* self, QPalette* palette);
extern __declspec(dllexport) 
QRect* QStyle_ItemTextRect(const QStyle* self, QFontMetrics* fm, QRect* r, int flags, bool enabled, struct miqt_string text);
extern __declspec(dllexport) 
QRect* QStyle_ItemPixmapRect(const QStyle* self, QRect* r, int flags, QPixmap* pixmap);
extern __declspec(dllexport) 
void QStyle_DrawItemText(const QStyle* self, QPainter* painter, QRect* rect, int flags, QPalette* pal, bool enabled, struct miqt_string text, int textRole);
extern __declspec(dllexport) 
void QStyle_DrawItemPixmap(const QStyle* self, QPainter* painter, QRect* rect, int alignment, QPixmap* pixmap);
extern __declspec(dllexport) 
QPalette* QStyle_StandardPalette(const QStyle* self);
extern __declspec(dllexport) 
void QStyle_DrawPrimitive(const QStyle* self, PrimitiveElement pe, QStyleOption* opt, QPainter* p, QWidget* w);
extern __declspec(dllexport) 
void QStyle_DrawControl(const QStyle* self, ControlElement element, QStyleOption* opt, QPainter* p, QWidget* w);
extern __declspec(dllexport) 
QRect* QStyle_SubElementRect(const QStyle* self, SubElement subElement, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) 
void QStyle_DrawComplexControl(const QStyle* self, ComplexControl cc, QStyleOptionComplex* opt, QPainter* p, QWidget* widget);
extern __declspec(dllexport) 
SubControl QStyle_HitTestComplexControl(const QStyle* self, ComplexControl cc, QStyleOptionComplex* opt, QPoint* pt, QWidget* widget);
extern __declspec(dllexport) 
QRect* QStyle_SubControlRect(const QStyle* self, ComplexControl cc, QStyleOptionComplex* opt, SubControl sc, QWidget* widget);
extern __declspec(dllexport) 
int QStyle_PixelMetric(const QStyle* self, PixelMetric metric, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) 
QSize* QStyle_SizeFromContents(const QStyle* self, ContentsType ct, QStyleOption* opt, QSize* contentsSize, QWidget* w);
extern __declspec(dllexport) 
int QStyle_StyleHint(const QStyle* self, StyleHint stylehint, QStyleOption* opt, QWidget* widget, QStyleHintReturn* returnData);
extern __declspec(dllexport) 
QPixmap* QStyle_StandardPixmap(const QStyle* self, StandardPixmap standardPixmap, QStyleOption* opt, QWidget* widget);
extern __declspec(dllexport) 
QIcon* QStyle_StandardIcon(const QStyle* self, StandardPixmap standardIcon, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) 
QPixmap* QStyle_GeneratedIconPixmap(const QStyle* self, int iconMode, QPixmap* pixmap, QStyleOption* opt);
extern __declspec(dllexport) 
QRect* QStyle_VisualRect(int direction, QRect* boundingRect, QRect* logicalRect);
extern __declspec(dllexport) 
QPoint* QStyle_VisualPos(int direction, QRect* boundingRect, QPoint* logicalPos);
extern __declspec(dllexport) 
int QStyle_SliderPositionFromValue(int min, int max, int val, int space);
extern __declspec(dllexport) 
int QStyle_SliderValueFromPosition(int min, int max, int pos, int space);
extern __declspec(dllexport) 
int QStyle_VisualAlignment(int direction, int alignment);
extern __declspec(dllexport) 
QRect* QStyle_AlignedRect(int direction, int alignment, QSize* size, QRect* rectangle);
extern __declspec(dllexport) 
int QStyle_LayoutSpacing(const QStyle* self, int control1, int control2, int orientation, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) 
int QStyle_CombinedLayoutSpacing(const QStyle* self, int controls1, int controls2, int orientation);
extern __declspec(dllexport) 
QStyle* QStyle_Proxy(const QStyle* self);
extern __declspec(dllexport) 
struct miqt_string QStyle_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QStyle_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
int QStyle_SliderPositionFromValue5(int min, int max, int val, int space, bool upsideDown);
extern __declspec(dllexport) 
int QStyle_SliderValueFromPosition5(int min, int max, int pos, int space, bool upsideDown);
extern __declspec(dllexport) 
int QStyle_CombinedLayoutSpacing4(const QStyle* self, int controls1, int controls2, int orientation, QStyleOption* option);
extern __declspec(dllexport) 
int QStyle_CombinedLayoutSpacing5(const QStyle* self, int controls1, int controls2, int orientation, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) 
void QStyle_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QStyle_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QStyle_override_virtual_Metacast(void* self, intptr_t slot);
void* QStyle_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QStyle_Delete(QStyle* self, bool isSubclass);

}
