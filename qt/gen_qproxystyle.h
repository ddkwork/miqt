#pragma once
#ifndef MIQT_QT_GEN_QPROXYSTYLE_H
#define MIQT_QT_GEN_QPROXYSTYLE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QApplication QApplication;
typedef struct QCommonStyle QCommonStyle;
typedef struct QEvent QEvent;
typedef struct QFontMetrics QFontMetrics;
typedef struct QIcon QIcon;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPainter QPainter;
typedef struct QPalette QPalette;
typedef struct QPixmap QPixmap;
typedef struct QPoint QPoint;
typedef struct QProxyStyle QProxyStyle;
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
QProxyStyle* QProxyStyle_new();
extern __declspec(dllexport) 
QProxyStyle* QProxyStyle_new2(struct miqt_string key);
extern __declspec(dllexport) 
QProxyStyle* QProxyStyle_new3(QStyle* style);
extern __declspec(dllexport) 
void QProxyStyle_virtbase(QProxyStyle* src
, QCommonStyle** outptr_QCommonStyle
);
extern __declspec(dllexport) 
QMetaObject* QProxyStyle_MetaObject(const QProxyStyle* self);
extern __declspec(dllexport) 
void* QProxyStyle_Metacast(QProxyStyle* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QProxyStyle_Tr(const char* s);
extern __declspec(dllexport) 
QStyle* QProxyStyle_BaseStyle(const QProxyStyle* self);
extern __declspec(dllexport) 
void QProxyStyle_SetBaseStyle(QProxyStyle* self, QStyle* style);
extern __declspec(dllexport) 
void QProxyStyle_DrawPrimitive(const QProxyStyle* self, PrimitiveElement element, QStyleOption* option, QPainter* painter, QWidget* widget);
extern __declspec(dllexport) 
void QProxyStyle_DrawControl(const QProxyStyle* self, ControlElement element, QStyleOption* option, QPainter* painter, QWidget* widget);
extern __declspec(dllexport) 
void QProxyStyle_DrawComplexControl(const QProxyStyle* self, ComplexControl control, QStyleOptionComplex* option, QPainter* painter, QWidget* widget);
extern __declspec(dllexport) 
void QProxyStyle_DrawItemText(const QProxyStyle* self, QPainter* painter, QRect* rect, int flags, QPalette* pal, bool enabled, struct miqt_string text, int textRole);
extern __declspec(dllexport) 
void QProxyStyle_DrawItemPixmap(const QProxyStyle* self, QPainter* painter, QRect* rect, int alignment, QPixmap* pixmap);
extern __declspec(dllexport) 
QSize* QProxyStyle_SizeFromContents(const QProxyStyle* self, ContentsType typeVal, QStyleOption* option, QSize* size, QWidget* widget);
extern __declspec(dllexport) 
QRect* QProxyStyle_SubElementRect(const QProxyStyle* self, SubElement element, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) 
QRect* QProxyStyle_SubControlRect(const QProxyStyle* self, ComplexControl cc, QStyleOptionComplex* opt, SubControl sc, QWidget* widget);
extern __declspec(dllexport) 
QRect* QProxyStyle_ItemTextRect(const QProxyStyle* self, QFontMetrics* fm, QRect* r, int flags, bool enabled, struct miqt_string text);
extern __declspec(dllexport) 
QRect* QProxyStyle_ItemPixmapRect(const QProxyStyle* self, QRect* r, int flags, QPixmap* pixmap);
extern __declspec(dllexport) 
SubControl QProxyStyle_HitTestComplexControl(const QProxyStyle* self, ComplexControl control, QStyleOptionComplex* option, QPoint* pos, QWidget* widget);
extern __declspec(dllexport) 
int QProxyStyle_StyleHint(const QProxyStyle* self, StyleHint hint, QStyleOption* option, QWidget* widget, QStyleHintReturn* returnData);
extern __declspec(dllexport) 
int QProxyStyle_PixelMetric(const QProxyStyle* self, PixelMetric metric, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) 
int QProxyStyle_LayoutSpacing(const QProxyStyle* self, int control1, int control2, int orientation, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) 
QIcon* QProxyStyle_StandardIcon(const QProxyStyle* self, StandardPixmap standardIcon, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) 
QPixmap* QProxyStyle_StandardPixmap(const QProxyStyle* self, StandardPixmap standardPixmap, QStyleOption* opt, QWidget* widget);
extern __declspec(dllexport) 
QPixmap* QProxyStyle_GeneratedIconPixmap(const QProxyStyle* self, int iconMode, QPixmap* pixmap, QStyleOption* opt);
extern __declspec(dllexport) 
QPalette* QProxyStyle_StandardPalette(const QProxyStyle* self);
extern __declspec(dllexport) 
void QProxyStyle_Polish(QProxyStyle* self, QWidget* widget);
extern __declspec(dllexport) 
void QProxyStyle_PolishWithPal(QProxyStyle* self, QPalette* pal);
extern __declspec(dllexport) 
void QProxyStyle_PolishWithApp(QProxyStyle* self, QApplication* app);
extern __declspec(dllexport) 
void QProxyStyle_Unpolish(QProxyStyle* self, QWidget* widget);
extern __declspec(dllexport) 
void QProxyStyle_UnpolishWithApp(QProxyStyle* self, QApplication* app);
extern __declspec(dllexport) 
bool QProxyStyle_Event(QProxyStyle* self, QEvent* e);
extern __declspec(dllexport) 
struct miqt_string QProxyStyle_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QProxyStyle_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QProxyStyle_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QProxyStyle_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QProxyStyle_override_virtual_Metacast(void* self, intptr_t slot);
void* QProxyStyle_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QProxyStyle_Delete(QProxyStyle* self, bool isSubclass);

}
