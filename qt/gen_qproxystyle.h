#pragma once
#ifndef MIQT_QT_GEN_QPROXYSTYLE_H
#define MIQT_QT_GEN_QPROXYSTYLE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QApplication;
class QCommonStyle;
class QEvent;
class QFontMetrics;
class QIcon;
class QMetaObject;
class QObject;
class QPainter;
class QPalette;
class QPixmap;
class QPoint;
class QProxyStyle;
class QRect;
class QSize;
class QStyle;
class QStyleHintReturn;
class QStyleOption;
class QStyleOptionComplex;
class QWidget;
class _GUID;
class type_info;
#else
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
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QProxyStyle* QProxyStyle_new();
extern __declspec(dllexport) QProxyStyle* QProxyStyle_new2(struct miqt_string key);
extern __declspec(dllexport) QProxyStyle* QProxyStyle_new3(QStyle* style);
extern __declspec(dllexport) void QProxyStyle_virtbase(QProxyStyle* src, QCommonStyle** outptr_QCommonStyle);
extern __declspec(dllexport) QMetaObject* QProxyStyle_MetaObject(const QProxyStyle* self);
extern __declspec(dllexport) void* QProxyStyle_Metacast(QProxyStyle* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QProxyStyle_Tr(const char* s);
extern __declspec(dllexport) QStyle* QProxyStyle_BaseStyle(const QProxyStyle* self);
extern __declspec(dllexport) void QProxyStyle_SetBaseStyle(QProxyStyle* self, QStyle* style);
extern __declspec(dllexport) void QProxyStyle_DrawPrimitive(const QProxyStyle* self, PrimitiveElement element, QStyleOption* option, QPainter* painter, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_DrawControl(const QProxyStyle* self, ControlElement element, QStyleOption* option, QPainter* painter, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_DrawComplexControl(const QProxyStyle* self, ComplexControl control, QStyleOptionComplex* option, QPainter* painter, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_DrawItemText(const QProxyStyle* self, QPainter* painter, QRect* rect, int flags, QPalette* pal, bool enabled, struct miqt_string text, int textRole);
extern __declspec(dllexport) void QProxyStyle_DrawItemPixmap(const QProxyStyle* self, QPainter* painter, QRect* rect, int alignment, QPixmap* pixmap);
extern __declspec(dllexport) QSize* QProxyStyle_SizeFromContents(const QProxyStyle* self, ContentsType typeVal, QStyleOption* option, QSize* size, QWidget* widget);
extern __declspec(dllexport) QRect* QProxyStyle_SubElementRect(const QProxyStyle* self, SubElement element, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) QRect* QProxyStyle_SubControlRect(const QProxyStyle* self, ComplexControl cc, QStyleOptionComplex* opt, SubControl sc, QWidget* widget);
extern __declspec(dllexport) QRect* QProxyStyle_ItemTextRect(const QProxyStyle* self, QFontMetrics* fm, QRect* r, int flags, bool enabled, struct miqt_string text);
extern __declspec(dllexport) QRect* QProxyStyle_ItemPixmapRect(const QProxyStyle* self, QRect* r, int flags, QPixmap* pixmap);
extern __declspec(dllexport) SubControl QProxyStyle_HitTestComplexControl(const QProxyStyle* self, ComplexControl control, QStyleOptionComplex* option, QPoint* pos, QWidget* widget);
extern __declspec(dllexport) int QProxyStyle_StyleHint(const QProxyStyle* self, StyleHint hint, QStyleOption* option, QWidget* widget, QStyleHintReturn* returnData);
extern __declspec(dllexport) int QProxyStyle_PixelMetric(const QProxyStyle* self, PixelMetric metric, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) int QProxyStyle_LayoutSpacing(const QProxyStyle* self, int control1, int control2, int orientation, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) QIcon* QProxyStyle_StandardIcon(const QProxyStyle* self, StandardPixmap standardIcon, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) QPixmap* QProxyStyle_StandardPixmap(const QProxyStyle* self, StandardPixmap standardPixmap, QStyleOption* opt, QWidget* widget);
extern __declspec(dllexport) QPixmap* QProxyStyle_GeneratedIconPixmap(const QProxyStyle* self, int iconMode, QPixmap* pixmap, QStyleOption* opt);
extern __declspec(dllexport) QPalette* QProxyStyle_StandardPalette(const QProxyStyle* self);
extern __declspec(dllexport) void QProxyStyle_Polish(QProxyStyle* self, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_PolishWithPal(QProxyStyle* self, QPalette* pal);
extern __declspec(dllexport) void QProxyStyle_PolishWithApp(QProxyStyle* self, QApplication* app);
extern __declspec(dllexport) void QProxyStyle_Unpolish(QProxyStyle* self, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_UnpolishWithApp(QProxyStyle* self, QApplication* app);
extern __declspec(dllexport) bool QProxyStyle_Event(QProxyStyle* self, QEvent* e);
extern __declspec(dllexport) struct miqt_string QProxyStyle_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QProxyStyle_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QProxyStyle_override_virtual_DrawPrimitive(void* self, intptr_t slot);
void QProxyStyle_virtualbase_DrawPrimitive(const void* self, PrimitiveElement element, QStyleOption* option, QPainter* painter, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_override_virtual_DrawControl(void* self, intptr_t slot);
void QProxyStyle_virtualbase_DrawControl(const void* self, ControlElement element, QStyleOption* option, QPainter* painter, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_override_virtual_DrawComplexControl(void* self, intptr_t slot);
void QProxyStyle_virtualbase_DrawComplexControl(const void* self, ComplexControl control, QStyleOptionComplex* option, QPainter* painter, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_override_virtual_DrawItemText(void* self, intptr_t slot);
void QProxyStyle_virtualbase_DrawItemText(const void* self, QPainter* painter, QRect* rect, int flags, QPalette* pal, bool enabled, struct miqt_string text, int textRole);
extern __declspec(dllexport) void QProxyStyle_override_virtual_DrawItemPixmap(void* self, intptr_t slot);
void QProxyStyle_virtualbase_DrawItemPixmap(const void* self, QPainter* painter, QRect* rect, int alignment, QPixmap* pixmap);
extern __declspec(dllexport) void QProxyStyle_override_virtual_SizeFromContents(void* self, intptr_t slot);
QSize* QProxyStyle_virtualbase_SizeFromContents(const void* self, ContentsType typeVal, QStyleOption* option, QSize* size, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_override_virtual_SubElementRect(void* self, intptr_t slot);
QRect* QProxyStyle_virtualbase_SubElementRect(const void* self, SubElement element, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_override_virtual_SubControlRect(void* self, intptr_t slot);
QRect* QProxyStyle_virtualbase_SubControlRect(const void* self, ComplexControl cc, QStyleOptionComplex* opt, SubControl sc, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_override_virtual_ItemTextRect(void* self, intptr_t slot);
QRect* QProxyStyle_virtualbase_ItemTextRect(const void* self, QFontMetrics* fm, QRect* r, int flags, bool enabled, struct miqt_string text);
extern __declspec(dllexport) void QProxyStyle_override_virtual_ItemPixmapRect(void* self, intptr_t slot);
QRect* QProxyStyle_virtualbase_ItemPixmapRect(const void* self, QRect* r, int flags, QPixmap* pixmap);
extern __declspec(dllexport) void QProxyStyle_override_virtual_HitTestComplexControl(void* self, intptr_t slot);
SubControl QProxyStyle_virtualbase_HitTestComplexControl(const void* self, ComplexControl control, QStyleOptionComplex* option, QPoint* pos, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_override_virtual_StyleHint(void* self, intptr_t slot);
int QProxyStyle_virtualbase_StyleHint(const void* self, StyleHint hint, QStyleOption* option, QWidget* widget, QStyleHintReturn* returnData);
extern __declspec(dllexport) void QProxyStyle_override_virtual_PixelMetric(void* self, intptr_t slot);
int QProxyStyle_virtualbase_PixelMetric(const void* self, PixelMetric metric, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_override_virtual_LayoutSpacing(void* self, intptr_t slot);
int QProxyStyle_virtualbase_LayoutSpacing(const void* self, int control1, int control2, int orientation, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_override_virtual_StandardIcon(void* self, intptr_t slot);
QIcon* QProxyStyle_virtualbase_StandardIcon(const void* self, StandardPixmap standardIcon, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_override_virtual_StandardPixmap(void* self, intptr_t slot);
QPixmap* QProxyStyle_virtualbase_StandardPixmap(const void* self, StandardPixmap standardPixmap, QStyleOption* opt, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_override_virtual_GeneratedIconPixmap(void* self, intptr_t slot);
QPixmap* QProxyStyle_virtualbase_GeneratedIconPixmap(const void* self, int iconMode, QPixmap* pixmap, QStyleOption* opt);
extern __declspec(dllexport) void QProxyStyle_override_virtual_StandardPalette(void* self, intptr_t slot);
QPalette* QProxyStyle_virtualbase_StandardPalette(const void* self);
extern __declspec(dllexport) void QProxyStyle_override_virtual_Polish(void* self, intptr_t slot);
void QProxyStyle_virtualbase_Polish(void* self, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_override_virtual_PolishWithPal(void* self, intptr_t slot);
void QProxyStyle_virtualbase_PolishWithPal(void* self, QPalette* pal);
extern __declspec(dllexport) void QProxyStyle_override_virtual_PolishWithApp(void* self, intptr_t slot);
void QProxyStyle_virtualbase_PolishWithApp(void* self, QApplication* app);
extern __declspec(dllexport) void QProxyStyle_override_virtual_Unpolish(void* self, intptr_t slot);
void QProxyStyle_virtualbase_Unpolish(void* self, QWidget* widget);
extern __declspec(dllexport) void QProxyStyle_override_virtual_UnpolishWithApp(void* self, intptr_t slot);
void QProxyStyle_virtualbase_UnpolishWithApp(void* self, QApplication* app);
extern __declspec(dllexport) void QProxyStyle_override_virtual_Event(void* self, intptr_t slot);
bool QProxyStyle_virtualbase_Event(void* self, QEvent* e);
extern __declspec(dllexport) void QProxyStyle_Delete(QProxyStyle* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
