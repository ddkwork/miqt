#pragma once
#ifndef MIQT_QT_GEN_QSTYLE_H
#define MIQT_QT_GEN_QSTYLE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QApplication QApplication;
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QFontMetrics QFontMetrics;
typedef struct QIcon QIcon;
typedef struct QMetaMethod QMetaMethod;
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
typedef struct QTimerEvent QTimerEvent;
typedef struct QWidget QWidget;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QStyle* QStyle_new();
extern __declspec(dllexport) void QStyle_virtbase(QStyle* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QStyle_MetaObject(const QStyle* self);
extern __declspec(dllexport) void* QStyle_Metacast(QStyle* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QStyle_Tr(const char* s);
extern __declspec(dllexport) struct miqt_string QStyle_Name(const QStyle* self);
extern __declspec(dllexport) void QStyle_Polish(QStyle* self, QWidget* widget);
extern __declspec(dllexport) void QStyle_Unpolish(QStyle* self, QWidget* widget);
extern __declspec(dllexport) void QStyle_PolishWithApplication(QStyle* self, QApplication* application);
extern __declspec(dllexport) void QStyle_UnpolishWithApplication(QStyle* self, QApplication* application);
extern __declspec(dllexport) void QStyle_PolishWithPalette(QStyle* self, QPalette* palette);
extern __declspec(dllexport) QRect* QStyle_ItemTextRect(const QStyle* self, QFontMetrics* fm, QRect* r, int flags, bool enabled, struct miqt_string text);
extern __declspec(dllexport) QRect* QStyle_ItemPixmapRect(const QStyle* self, QRect* r, int flags, QPixmap* pixmap);
extern __declspec(dllexport) void QStyle_DrawItemText(const QStyle* self, QPainter* painter, QRect* rect, int flags, QPalette* pal, bool enabled, struct miqt_string text, int textRole);
extern __declspec(dllexport) void QStyle_DrawItemPixmap(const QStyle* self, QPainter* painter, QRect* rect, int alignment, QPixmap* pixmap);
extern __declspec(dllexport) QPalette* QStyle_StandardPalette(const QStyle* self);
extern __declspec(dllexport) void QStyle_DrawPrimitive(const QStyle* self, PrimitiveElement pe, QStyleOption* opt, QPainter* p, QWidget* w);
extern __declspec(dllexport) void QStyle_DrawControl(const QStyle* self, ControlElement element, QStyleOption* opt, QPainter* p, QWidget* w);
extern __declspec(dllexport) QRect* QStyle_SubElementRect(const QStyle* self, SubElement subElement, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) void QStyle_DrawComplexControl(const QStyle* self, ComplexControl cc, QStyleOptionComplex* opt, QPainter* p, QWidget* widget);
extern __declspec(dllexport) SubControl QStyle_HitTestComplexControl(const QStyle* self, ComplexControl cc, QStyleOptionComplex* opt, QPoint* pt, QWidget* widget);
extern __declspec(dllexport) QRect* QStyle_SubControlRect(const QStyle* self, ComplexControl cc, QStyleOptionComplex* opt, SubControl sc, QWidget* widget);
extern __declspec(dllexport) int QStyle_PixelMetric(const QStyle* self, PixelMetric metric, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) QSize* QStyle_SizeFromContents(const QStyle* self, ContentsType ct, QStyleOption* opt, QSize* contentsSize, QWidget* w);
extern __declspec(dllexport) int QStyle_StyleHint(const QStyle* self, StyleHint stylehint, QStyleOption* opt, QWidget* widget, QStyleHintReturn* returnData);
extern __declspec(dllexport) QPixmap* QStyle_StandardPixmap(const QStyle* self, StandardPixmap standardPixmap, QStyleOption* opt, QWidget* widget);
extern __declspec(dllexport) QIcon* QStyle_StandardIcon(const QStyle* self, StandardPixmap standardIcon, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) QPixmap* QStyle_GeneratedIconPixmap(const QStyle* self, int iconMode, QPixmap* pixmap, QStyleOption* opt);
extern __declspec(dllexport) QRect* QStyle_VisualRect(int direction, QRect* boundingRect, QRect* logicalRect);
extern __declspec(dllexport) QPoint* QStyle_VisualPos(int direction, QRect* boundingRect, QPoint* logicalPos);
extern __declspec(dllexport) int QStyle_SliderPositionFromValue(int min, int max, int val, int space);
extern __declspec(dllexport) int QStyle_SliderValueFromPosition(int min, int max, int pos, int space);
extern __declspec(dllexport) int QStyle_VisualAlignment(int direction, int alignment);
extern __declspec(dllexport) QRect* QStyle_AlignedRect(int direction, int alignment, QSize* size, QRect* rectangle);
extern __declspec(dllexport) int QStyle_LayoutSpacing(const QStyle* self, int control1, int control2, int orientation, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) int QStyle_CombinedLayoutSpacing(const QStyle* self, int controls1, int controls2, int orientation);
extern __declspec(dllexport) QStyle* QStyle_Proxy(const QStyle* self);
extern __declspec(dllexport) struct miqt_string QStyle_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QStyle_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) int QStyle_SliderPositionFromValue5(int min, int max, int val, int space, bool upsideDown);
extern __declspec(dllexport) int QStyle_SliderValueFromPosition5(int min, int max, int pos, int space, bool upsideDown);
extern __declspec(dllexport) int QStyle_CombinedLayoutSpacing4(const QStyle* self, int controls1, int controls2, int orientation, QStyleOption* option);
extern __declspec(dllexport) int QStyle_CombinedLayoutSpacing5(const QStyle* self, int controls1, int controls2, int orientation, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) void QStyle_override_virtual_Polish(void* self, intptr_t slot);
void QStyle_virtualbase_Polish(void* self, QWidget* widget);
extern __declspec(dllexport) void QStyle_override_virtual_Unpolish(void* self, intptr_t slot);
void QStyle_virtualbase_Unpolish(void* self, QWidget* widget);
extern __declspec(dllexport) void QStyle_override_virtual_PolishWithApplication(void* self, intptr_t slot);
void QStyle_virtualbase_PolishWithApplication(void* self, QApplication* application);
extern __declspec(dllexport) void QStyle_override_virtual_UnpolishWithApplication(void* self, intptr_t slot);
void QStyle_virtualbase_UnpolishWithApplication(void* self, QApplication* application);
extern __declspec(dllexport) void QStyle_override_virtual_PolishWithPalette(void* self, intptr_t slot);
void QStyle_virtualbase_PolishWithPalette(void* self, QPalette* palette);
extern __declspec(dllexport) void QStyle_override_virtual_ItemTextRect(void* self, intptr_t slot);
QRect* QStyle_virtualbase_ItemTextRect(const void* self, QFontMetrics* fm, QRect* r, int flags, bool enabled, struct miqt_string text);
extern __declspec(dllexport) void QStyle_override_virtual_ItemPixmapRect(void* self, intptr_t slot);
QRect* QStyle_virtualbase_ItemPixmapRect(const void* self, QRect* r, int flags, QPixmap* pixmap);
extern __declspec(dllexport) void QStyle_override_virtual_DrawItemText(void* self, intptr_t slot);
void QStyle_virtualbase_DrawItemText(const void* self, QPainter* painter, QRect* rect, int flags, QPalette* pal, bool enabled, struct miqt_string text, int textRole);
extern __declspec(dllexport) void QStyle_override_virtual_DrawItemPixmap(void* self, intptr_t slot);
void QStyle_virtualbase_DrawItemPixmap(const void* self, QPainter* painter, QRect* rect, int alignment, QPixmap* pixmap);
extern __declspec(dllexport) void QStyle_override_virtual_StandardPalette(void* self, intptr_t slot);
QPalette* QStyle_virtualbase_StandardPalette(const void* self);
extern __declspec(dllexport) void QStyle_override_virtual_DrawPrimitive(void* self, intptr_t slot);
void QStyle_virtualbase_DrawPrimitive(const void* self, PrimitiveElement pe, QStyleOption* opt, QPainter* p, QWidget* w);
extern __declspec(dllexport) void QStyle_override_virtual_DrawControl(void* self, intptr_t slot);
void QStyle_virtualbase_DrawControl(const void* self, ControlElement element, QStyleOption* opt, QPainter* p, QWidget* w);
extern __declspec(dllexport) void QStyle_override_virtual_SubElementRect(void* self, intptr_t slot);
QRect* QStyle_virtualbase_SubElementRect(const void* self, SubElement subElement, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) void QStyle_override_virtual_DrawComplexControl(void* self, intptr_t slot);
void QStyle_virtualbase_DrawComplexControl(const void* self, ComplexControl cc, QStyleOptionComplex* opt, QPainter* p, QWidget* widget);
extern __declspec(dllexport) void QStyle_override_virtual_HitTestComplexControl(void* self, intptr_t slot);
SubControl QStyle_virtualbase_HitTestComplexControl(const void* self, ComplexControl cc, QStyleOptionComplex* opt, QPoint* pt, QWidget* widget);
extern __declspec(dllexport) void QStyle_override_virtual_SubControlRect(void* self, intptr_t slot);
QRect* QStyle_virtualbase_SubControlRect(const void* self, ComplexControl cc, QStyleOptionComplex* opt, SubControl sc, QWidget* widget);
extern __declspec(dllexport) void QStyle_override_virtual_PixelMetric(void* self, intptr_t slot);
int QStyle_virtualbase_PixelMetric(const void* self, PixelMetric metric, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) void QStyle_override_virtual_SizeFromContents(void* self, intptr_t slot);
QSize* QStyle_virtualbase_SizeFromContents(const void* self, ContentsType ct, QStyleOption* opt, QSize* contentsSize, QWidget* w);
extern __declspec(dllexport) void QStyle_override_virtual_StyleHint(void* self, intptr_t slot);
int QStyle_virtualbase_StyleHint(const void* self, StyleHint stylehint, QStyleOption* opt, QWidget* widget, QStyleHintReturn* returnData);
extern __declspec(dllexport) void QStyle_override_virtual_StandardPixmap(void* self, intptr_t slot);
QPixmap* QStyle_virtualbase_StandardPixmap(const void* self, StandardPixmap standardPixmap, QStyleOption* opt, QWidget* widget);
extern __declspec(dllexport) void QStyle_override_virtual_StandardIcon(void* self, intptr_t slot);
QIcon* QStyle_virtualbase_StandardIcon(const void* self, StandardPixmap standardIcon, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) void QStyle_override_virtual_GeneratedIconPixmap(void* self, intptr_t slot);
QPixmap* QStyle_virtualbase_GeneratedIconPixmap(const void* self, int iconMode, QPixmap* pixmap, QStyleOption* opt);
extern __declspec(dllexport) void QStyle_override_virtual_LayoutSpacing(void* self, intptr_t slot);
int QStyle_virtualbase_LayoutSpacing(const void* self, int control1, int control2, int orientation, QStyleOption* option, QWidget* widget);
extern __declspec(dllexport) void QStyle_override_virtual_Event(void* self, intptr_t slot);
bool QStyle_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QStyle_override_virtual_EventFilter(void* self, intptr_t slot);
bool QStyle_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QStyle_override_virtual_TimerEvent(void* self, intptr_t slot);
void QStyle_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QStyle_override_virtual_ChildEvent(void* self, intptr_t slot);
void QStyle_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QStyle_override_virtual_CustomEvent(void* self, intptr_t slot);
void QStyle_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QStyle_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QStyle_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QStyle_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QStyle_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QStyle_Delete(QStyle* self, bool isSubclass);

} 
