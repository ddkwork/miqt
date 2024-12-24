#pragma once
#ifndef MIQT_QT_GEN_QCOMMONSTYLE_H
#define MIQT_QT_GEN_QCOMMONSTYLE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QApplication QApplication;
typedef struct QCommonStyle QCommonStyle;
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
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QCommonStyle* QCommonStyle_new();
extern __declspec(dllexport) void QCommonStyle_virtbase(QCommonStyle* src, QStyle** outptr_QStyle);
extern __declspec(dllexport) QMetaObject* QCommonStyle_MetaObject(const QCommonStyle* self);
extern __declspec(dllexport) void* QCommonStyle_Metacast(QCommonStyle* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QCommonStyle_Tr(const char* s);
extern __declspec(dllexport) void QCommonStyle_DrawPrimitive(const QCommonStyle* self, PrimitiveElement pe, QStyleOption* opt, QPainter* p, QWidget* w);
extern __declspec(dllexport) void QCommonStyle_DrawControl(const QCommonStyle* self, ControlElement element, QStyleOption* opt, QPainter* p, QWidget* w);
extern __declspec(dllexport) QRect* QCommonStyle_SubElementRect(const QCommonStyle* self, SubElement r, QStyleOption* opt, QWidget* widget);
extern __declspec(dllexport) void QCommonStyle_DrawComplexControl(const QCommonStyle* self, ComplexControl cc, QStyleOptionComplex* opt, QPainter* p, QWidget* w);
extern __declspec(dllexport) SubControl QCommonStyle_HitTestComplexControl(const QCommonStyle* self, ComplexControl cc, QStyleOptionComplex* opt, QPoint* pt, QWidget* w);
extern __declspec(dllexport) QRect* QCommonStyle_SubControlRect(const QCommonStyle* self, ComplexControl cc, QStyleOptionComplex* opt, SubControl sc, QWidget* w);
extern __declspec(dllexport) QSize* QCommonStyle_SizeFromContents(const QCommonStyle* self, ContentsType ct, QStyleOption* opt, QSize* contentsSize, QWidget* widget);
extern __declspec(dllexport) int QCommonStyle_PixelMetric(const QCommonStyle* self, PixelMetric m, QStyleOption* opt, QWidget* widget);
extern __declspec(dllexport) int QCommonStyle_StyleHint(const QCommonStyle* self, StyleHint sh, QStyleOption* opt, QWidget* w, QStyleHintReturn* shret);
extern __declspec(dllexport) QIcon* QCommonStyle_StandardIcon(const QCommonStyle* self, StandardPixmap standardIcon, QStyleOption* opt, QWidget* widget);
extern __declspec(dllexport) QPixmap* QCommonStyle_StandardPixmap(const QCommonStyle* self, StandardPixmap sp, QStyleOption* opt, QWidget* widget);
extern __declspec(dllexport) QPixmap* QCommonStyle_GeneratedIconPixmap(const QCommonStyle* self, int iconMode, QPixmap* pixmap, QStyleOption* opt);
extern __declspec(dllexport) int QCommonStyle_LayoutSpacing(const QCommonStyle* self, int control1, int control2, int orientation, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) void QCommonStyle_Polish(QCommonStyle* self, QPalette* param1);
extern __declspec(dllexport) void QCommonStyle_PolishWithApp(QCommonStyle* self, QApplication* app);
extern __declspec(dllexport) void QCommonStyle_PolishWithWidget(QCommonStyle* self, QWidget* widget);
extern __declspec(dllexport) void QCommonStyle_Unpolish(QCommonStyle* self, QWidget* widget);
extern __declspec(dllexport) void QCommonStyle_UnpolishWithApplication(QCommonStyle* self, QApplication* application);
extern __declspec(dllexport) struct miqt_string QCommonStyle_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QCommonStyle_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QCommonStyle_override_virtual_DrawPrimitive(void* self, intptr_t slot);
void QCommonStyle_virtualbase_DrawPrimitive(const void* self, PrimitiveElement pe, QStyleOption* opt, QPainter* p, QWidget* w);
extern __declspec(dllexport) void QCommonStyle_override_virtual_DrawControl(void* self, intptr_t slot);
void QCommonStyle_virtualbase_DrawControl(const void* self, ControlElement element, QStyleOption* opt, QPainter* p, QWidget* w);
extern __declspec(dllexport) void QCommonStyle_override_virtual_SubElementRect(void* self, intptr_t slot);
QRect* QCommonStyle_virtualbase_SubElementRect(const void* self, SubElement r, QStyleOption* opt, QWidget* widget);
extern __declspec(dllexport) void QCommonStyle_override_virtual_DrawComplexControl(void* self, intptr_t slot);
void QCommonStyle_virtualbase_DrawComplexControl(const void* self, ComplexControl cc, QStyleOptionComplex* opt, QPainter* p, QWidget* w);
extern __declspec(dllexport) void QCommonStyle_override_virtual_HitTestComplexControl(void* self, intptr_t slot);
SubControl QCommonStyle_virtualbase_HitTestComplexControl(const void* self, ComplexControl cc, QStyleOptionComplex* opt, QPoint* pt, QWidget* w);
extern __declspec(dllexport) void QCommonStyle_override_virtual_SubControlRect(void* self, intptr_t slot);
QRect* QCommonStyle_virtualbase_SubControlRect(const void* self, ComplexControl cc, QStyleOptionComplex* opt, SubControl sc, QWidget* w);
extern __declspec(dllexport) void QCommonStyle_override_virtual_SizeFromContents(void* self, intptr_t slot);
QSize* QCommonStyle_virtualbase_SizeFromContents(const void* self, ContentsType ct, QStyleOption* opt, QSize* contentsSize, QWidget* widget);
extern __declspec(dllexport) void QCommonStyle_override_virtual_PixelMetric(void* self, intptr_t slot);
int QCommonStyle_virtualbase_PixelMetric(const void* self, PixelMetric m, QStyleOption* opt, QWidget* widget);
extern __declspec(dllexport) void QCommonStyle_override_virtual_StyleHint(void* self, intptr_t slot);
int QCommonStyle_virtualbase_StyleHint(const void* self, StyleHint sh, QStyleOption* opt, QWidget* w, QStyleHintReturn* shret);
extern __declspec(dllexport) void QCommonStyle_override_virtual_StandardIcon(void* self, intptr_t slot);
QIcon* QCommonStyle_virtualbase_StandardIcon(const void* self, StandardPixmap standardIcon, QStyleOption* opt, QWidget* widget);
extern __declspec(dllexport) void QCommonStyle_override_virtual_StandardPixmap(void* self, intptr_t slot);
QPixmap* QCommonStyle_virtualbase_StandardPixmap(const void* self, StandardPixmap sp, QStyleOption* opt, QWidget* widget);
extern __declspec(dllexport) void QCommonStyle_override_virtual_GeneratedIconPixmap(void* self, intptr_t slot);
QPixmap* QCommonStyle_virtualbase_GeneratedIconPixmap(const void* self, int iconMode, QPixmap* pixmap, QStyleOption* opt);
extern __declspec(dllexport) void QCommonStyle_override_virtual_LayoutSpacing(void* self, intptr_t slot);
int QCommonStyle_virtualbase_LayoutSpacing(const void* self, int control1, int control2, int orientation, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) void QCommonStyle_override_virtual_Polish(void* self, intptr_t slot);
void QCommonStyle_virtualbase_Polish(void* self, QPalette* param1);
extern __declspec(dllexport) void QCommonStyle_override_virtual_PolishWithApp(void* self, intptr_t slot);
void QCommonStyle_virtualbase_PolishWithApp(void* self, QApplication* app);
extern __declspec(dllexport) void QCommonStyle_override_virtual_PolishWithWidget(void* self, intptr_t slot);
void QCommonStyle_virtualbase_PolishWithWidget(void* self, QWidget* widget);
extern __declspec(dllexport) void QCommonStyle_override_virtual_Unpolish(void* self, intptr_t slot);
void QCommonStyle_virtualbase_Unpolish(void* self, QWidget* widget);
extern __declspec(dllexport) void QCommonStyle_override_virtual_UnpolishWithApplication(void* self, intptr_t slot);
void QCommonStyle_virtualbase_UnpolishWithApplication(void* self, QApplication* application);
extern __declspec(dllexport) void QCommonStyle_override_virtual_ItemTextRect(void* self, intptr_t slot);
QRect* QCommonStyle_virtualbase_ItemTextRect(const void* self, QFontMetrics* fm, QRect* r, int flags, bool enabled, struct miqt_string text);
extern __declspec(dllexport) void QCommonStyle_override_virtual_ItemPixmapRect(void* self, intptr_t slot);
QRect* QCommonStyle_virtualbase_ItemPixmapRect(const void* self, QRect* r, int flags, QPixmap* pixmap);
extern __declspec(dllexport) void QCommonStyle_override_virtual_DrawItemText(void* self, intptr_t slot);
void QCommonStyle_virtualbase_DrawItemText(const void* self, QPainter* painter, QRect* rect, int flags, QPalette* pal, bool enabled, struct miqt_string text, int textRole);
extern __declspec(dllexport) void QCommonStyle_override_virtual_DrawItemPixmap(void* self, intptr_t slot);
void QCommonStyle_virtualbase_DrawItemPixmap(const void* self, QPainter* painter, QRect* rect, int alignment, QPixmap* pixmap);
extern __declspec(dllexport) void QCommonStyle_override_virtual_StandardPalette(void* self, intptr_t slot);
QPalette* QCommonStyle_virtualbase_StandardPalette(const void* self);
extern __declspec(dllexport) void QCommonStyle_Delete(QCommonStyle* self, bool isSubclass);

} 
