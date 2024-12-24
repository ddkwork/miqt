#pragma once
#ifndef MIQT_QT_GEN_QGRAPHICSEFFECT_H
#define MIQT_QT_GEN_QGRAPHICSEFFECT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QBrush QBrush;
typedef struct QChildEvent QChildEvent;
typedef struct QColor QColor;
typedef struct QEvent QEvent;
typedef struct QGraphicsBlurEffect QGraphicsBlurEffect;
typedef struct QGraphicsColorizeEffect QGraphicsColorizeEffect;
typedef struct QGraphicsDropShadowEffect QGraphicsDropShadowEffect;
typedef struct QGraphicsEffect QGraphicsEffect;
typedef struct QGraphicsOpacityEffect QGraphicsOpacityEffect;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPainter QPainter;
typedef struct QPointF QPointF;
typedef struct QRectF QRectF;
typedef struct QTimerEvent QTimerEvent;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsEffect* QGraphicsEffect_new();
extern __declspec(dllexport) QGraphicsEffect* QGraphicsEffect_new2(QObject* parent);
extern __declspec(dllexport) void QGraphicsEffect_virtbase(QGraphicsEffect* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QGraphicsEffect_MetaObject(const QGraphicsEffect* self);
extern __declspec(dllexport) void* QGraphicsEffect_Metacast(QGraphicsEffect* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QGraphicsEffect_Tr(const char* s);
extern __declspec(dllexport) QRectF* QGraphicsEffect_BoundingRectFor(const QGraphicsEffect* self, QRectF* sourceRect);
extern __declspec(dllexport) QRectF* QGraphicsEffect_BoundingRect(const QGraphicsEffect* self);
extern __declspec(dllexport) bool QGraphicsEffect_IsEnabled(const QGraphicsEffect* self);
extern __declspec(dllexport) void QGraphicsEffect_SetEnabled(QGraphicsEffect* self, bool enable);
extern __declspec(dllexport) void QGraphicsEffect_Update(QGraphicsEffect* self);
extern __declspec(dllexport) void QGraphicsEffect_EnabledChanged(QGraphicsEffect* self, bool enabled);
void QGraphicsEffect_connect_EnabledChanged(QGraphicsEffect* self, intptr_t slot);
extern __declspec(dllexport) void QGraphicsEffect_Draw(QGraphicsEffect* self, QPainter* painter);
extern __declspec(dllexport) void QGraphicsEffect_SourceChanged(QGraphicsEffect* self, ChangeFlags flags);
extern __declspec(dllexport) struct miqt_string QGraphicsEffect_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QGraphicsEffect_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QGraphicsEffect_override_virtual_BoundingRectFor(void* self, intptr_t slot);
QRectF* QGraphicsEffect_virtualbase_BoundingRectFor(const void* self, QRectF* sourceRect);
extern __declspec(dllexport) void QGraphicsEffect_override_virtual_Draw(void* self, intptr_t slot);
void QGraphicsEffect_virtualbase_Draw(void* self, QPainter* painter);
extern __declspec(dllexport) void QGraphicsEffect_override_virtual_SourceChanged(void* self, intptr_t slot);
void QGraphicsEffect_virtualbase_SourceChanged(void* self, ChangeFlags flags);
extern __declspec(dllexport) void QGraphicsEffect_override_virtual_Event(void* self, intptr_t slot);
bool QGraphicsEffect_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QGraphicsEffect_override_virtual_EventFilter(void* self, intptr_t slot);
bool QGraphicsEffect_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QGraphicsEffect_override_virtual_TimerEvent(void* self, intptr_t slot);
void QGraphicsEffect_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QGraphicsEffect_override_virtual_ChildEvent(void* self, intptr_t slot);
void QGraphicsEffect_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QGraphicsEffect_override_virtual_CustomEvent(void* self, intptr_t slot);
void QGraphicsEffect_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QGraphicsEffect_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QGraphicsEffect_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QGraphicsEffect_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QGraphicsEffect_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QGraphicsEffect_Delete(QGraphicsEffect* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsColorizeEffect* QGraphicsColorizeEffect_new();
extern __declspec(dllexport) QGraphicsColorizeEffect* QGraphicsColorizeEffect_new2(QObject* parent);
extern __declspec(dllexport) void QGraphicsColorizeEffect_virtbase(QGraphicsColorizeEffect* src, QGraphicsEffect** outptr_QGraphicsEffect);
extern __declspec(dllexport) QMetaObject* QGraphicsColorizeEffect_MetaObject(const QGraphicsColorizeEffect* self);
extern __declspec(dllexport) void* QGraphicsColorizeEffect_Metacast(QGraphicsColorizeEffect* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QGraphicsColorizeEffect_Tr(const char* s);
extern __declspec(dllexport) QColor* QGraphicsColorizeEffect_Color(const QGraphicsColorizeEffect* self);
extern __declspec(dllexport) double QGraphicsColorizeEffect_Strength(const QGraphicsColorizeEffect* self);
extern __declspec(dllexport) void QGraphicsColorizeEffect_SetColor(QGraphicsColorizeEffect* self, QColor* c);
extern __declspec(dllexport) void QGraphicsColorizeEffect_SetStrength(QGraphicsColorizeEffect* self, double strength);
extern __declspec(dllexport) void QGraphicsColorizeEffect_ColorChanged(QGraphicsColorizeEffect* self, QColor* color);
void QGraphicsColorizeEffect_connect_ColorChanged(QGraphicsColorizeEffect* self, intptr_t slot);
extern __declspec(dllexport) void QGraphicsColorizeEffect_StrengthChanged(QGraphicsColorizeEffect* self, double strength);
void QGraphicsColorizeEffect_connect_StrengthChanged(QGraphicsColorizeEffect* self, intptr_t slot);
extern __declspec(dllexport) void QGraphicsColorizeEffect_Draw(QGraphicsColorizeEffect* self, QPainter* painter);
extern __declspec(dllexport) struct miqt_string QGraphicsColorizeEffect_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QGraphicsColorizeEffect_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QGraphicsColorizeEffect_override_virtual_Draw(void* self, intptr_t slot);
void QGraphicsColorizeEffect_virtualbase_Draw(void* self, QPainter* painter);
extern __declspec(dllexport) void QGraphicsColorizeEffect_override_virtual_BoundingRectFor(void* self, intptr_t slot);
QRectF* QGraphicsColorizeEffect_virtualbase_BoundingRectFor(const void* self, QRectF* sourceRect);
extern __declspec(dllexport) void QGraphicsColorizeEffect_override_virtual_SourceChanged(void* self, intptr_t slot);
void QGraphicsColorizeEffect_virtualbase_SourceChanged(void* self, ChangeFlags flags);
extern __declspec(dllexport) void QGraphicsColorizeEffect_Delete(QGraphicsColorizeEffect* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsBlurEffect* QGraphicsBlurEffect_new();
extern __declspec(dllexport) QGraphicsBlurEffect* QGraphicsBlurEffect_new2(QObject* parent);
extern __declspec(dllexport) void QGraphicsBlurEffect_virtbase(QGraphicsBlurEffect* src, QGraphicsEffect** outptr_QGraphicsEffect);
extern __declspec(dllexport) QMetaObject* QGraphicsBlurEffect_MetaObject(const QGraphicsBlurEffect* self);
extern __declspec(dllexport) void* QGraphicsBlurEffect_Metacast(QGraphicsBlurEffect* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QGraphicsBlurEffect_Tr(const char* s);
extern __declspec(dllexport) QRectF* QGraphicsBlurEffect_BoundingRectFor(const QGraphicsBlurEffect* self, QRectF* rect);
extern __declspec(dllexport) double QGraphicsBlurEffect_BlurRadius(const QGraphicsBlurEffect* self);
extern __declspec(dllexport) BlurHints QGraphicsBlurEffect_BlurHints(const QGraphicsBlurEffect* self);
extern __declspec(dllexport) void QGraphicsBlurEffect_SetBlurRadius(QGraphicsBlurEffect* self, double blurRadius);
extern __declspec(dllexport) void QGraphicsBlurEffect_SetBlurHints(QGraphicsBlurEffect* self, BlurHints hints);
extern __declspec(dllexport) void QGraphicsBlurEffect_BlurRadiusChanged(QGraphicsBlurEffect* self, double blurRadius);
void QGraphicsBlurEffect_connect_BlurRadiusChanged(QGraphicsBlurEffect* self, intptr_t slot);
extern __declspec(dllexport) void QGraphicsBlurEffect_BlurHintsChanged(QGraphicsBlurEffect* self, BlurHints hints);
void QGraphicsBlurEffect_connect_BlurHintsChanged(QGraphicsBlurEffect* self, intptr_t slot);
extern __declspec(dllexport) void QGraphicsBlurEffect_Draw(QGraphicsBlurEffect* self, QPainter* painter);
extern __declspec(dllexport) struct miqt_string QGraphicsBlurEffect_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QGraphicsBlurEffect_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QGraphicsBlurEffect_override_virtual_BoundingRectFor(void* self, intptr_t slot);
QRectF* QGraphicsBlurEffect_virtualbase_BoundingRectFor(const void* self, QRectF* rect);
extern __declspec(dllexport) void QGraphicsBlurEffect_override_virtual_Draw(void* self, intptr_t slot);
void QGraphicsBlurEffect_virtualbase_Draw(void* self, QPainter* painter);
extern __declspec(dllexport) void QGraphicsBlurEffect_override_virtual_SourceChanged(void* self, intptr_t slot);
void QGraphicsBlurEffect_virtualbase_SourceChanged(void* self, ChangeFlags flags);
extern __declspec(dllexport) void QGraphicsBlurEffect_Delete(QGraphicsBlurEffect* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsDropShadowEffect* QGraphicsDropShadowEffect_new();
extern __declspec(dllexport) QGraphicsDropShadowEffect* QGraphicsDropShadowEffect_new2(QObject* parent);
extern __declspec(dllexport) void QGraphicsDropShadowEffect_virtbase(QGraphicsDropShadowEffect* src, QGraphicsEffect** outptr_QGraphicsEffect);
extern __declspec(dllexport) QMetaObject* QGraphicsDropShadowEffect_MetaObject(const QGraphicsDropShadowEffect* self);
extern __declspec(dllexport) void* QGraphicsDropShadowEffect_Metacast(QGraphicsDropShadowEffect* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QGraphicsDropShadowEffect_Tr(const char* s);
extern __declspec(dllexport) QRectF* QGraphicsDropShadowEffect_BoundingRectFor(const QGraphicsDropShadowEffect* self, QRectF* rect);
extern __declspec(dllexport) QPointF* QGraphicsDropShadowEffect_Offset(const QGraphicsDropShadowEffect* self);
extern __declspec(dllexport) double QGraphicsDropShadowEffect_XOffset(const QGraphicsDropShadowEffect* self);
extern __declspec(dllexport) double QGraphicsDropShadowEffect_YOffset(const QGraphicsDropShadowEffect* self);
extern __declspec(dllexport) double QGraphicsDropShadowEffect_BlurRadius(const QGraphicsDropShadowEffect* self);
extern __declspec(dllexport) QColor* QGraphicsDropShadowEffect_Color(const QGraphicsDropShadowEffect* self);
extern __declspec(dllexport) void QGraphicsDropShadowEffect_SetOffset(QGraphicsDropShadowEffect* self, QPointF* ofs);
extern __declspec(dllexport) void QGraphicsDropShadowEffect_SetOffset2(QGraphicsDropShadowEffect* self, double dx, double dy);
extern __declspec(dllexport) void QGraphicsDropShadowEffect_SetOffsetWithQreal(QGraphicsDropShadowEffect* self, double d);
extern __declspec(dllexport) void QGraphicsDropShadowEffect_SetXOffset(QGraphicsDropShadowEffect* self, double dx);
extern __declspec(dllexport) void QGraphicsDropShadowEffect_SetYOffset(QGraphicsDropShadowEffect* self, double dy);
extern __declspec(dllexport) void QGraphicsDropShadowEffect_SetBlurRadius(QGraphicsDropShadowEffect* self, double blurRadius);
extern __declspec(dllexport) void QGraphicsDropShadowEffect_SetColor(QGraphicsDropShadowEffect* self, QColor* color);
extern __declspec(dllexport) void QGraphicsDropShadowEffect_OffsetChanged(QGraphicsDropShadowEffect* self, QPointF* offset);
void QGraphicsDropShadowEffect_connect_OffsetChanged(QGraphicsDropShadowEffect* self, intptr_t slot);
extern __declspec(dllexport) void QGraphicsDropShadowEffect_BlurRadiusChanged(QGraphicsDropShadowEffect* self, double blurRadius);
void QGraphicsDropShadowEffect_connect_BlurRadiusChanged(QGraphicsDropShadowEffect* self, intptr_t slot);
extern __declspec(dllexport) void QGraphicsDropShadowEffect_ColorChanged(QGraphicsDropShadowEffect* self, QColor* color);
void QGraphicsDropShadowEffect_connect_ColorChanged(QGraphicsDropShadowEffect* self, intptr_t slot);
extern __declspec(dllexport) void QGraphicsDropShadowEffect_Draw(QGraphicsDropShadowEffect* self, QPainter* painter);
extern __declspec(dllexport) struct miqt_string QGraphicsDropShadowEffect_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QGraphicsDropShadowEffect_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QGraphicsDropShadowEffect_override_virtual_BoundingRectFor(void* self, intptr_t slot);
QRectF* QGraphicsDropShadowEffect_virtualbase_BoundingRectFor(const void* self, QRectF* rect);
extern __declspec(dllexport) void QGraphicsDropShadowEffect_override_virtual_Draw(void* self, intptr_t slot);
void QGraphicsDropShadowEffect_virtualbase_Draw(void* self, QPainter* painter);
extern __declspec(dllexport) void QGraphicsDropShadowEffect_override_virtual_SourceChanged(void* self, intptr_t slot);
void QGraphicsDropShadowEffect_virtualbase_SourceChanged(void* self, ChangeFlags flags);
extern __declspec(dllexport) void QGraphicsDropShadowEffect_Delete(QGraphicsDropShadowEffect* self, bool isSubclass);

extern __declspec(dllexport) QGraphicsOpacityEffect* QGraphicsOpacityEffect_new();
extern __declspec(dllexport) QGraphicsOpacityEffect* QGraphicsOpacityEffect_new2(QObject* parent);
extern __declspec(dllexport) void QGraphicsOpacityEffect_virtbase(QGraphicsOpacityEffect* src, QGraphicsEffect** outptr_QGraphicsEffect);
extern __declspec(dllexport) QMetaObject* QGraphicsOpacityEffect_MetaObject(const QGraphicsOpacityEffect* self);
extern __declspec(dllexport) void* QGraphicsOpacityEffect_Metacast(QGraphicsOpacityEffect* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QGraphicsOpacityEffect_Tr(const char* s);
extern __declspec(dllexport) double QGraphicsOpacityEffect_Opacity(const QGraphicsOpacityEffect* self);
extern __declspec(dllexport) QBrush* QGraphicsOpacityEffect_OpacityMask(const QGraphicsOpacityEffect* self);
extern __declspec(dllexport) void QGraphicsOpacityEffect_SetOpacity(QGraphicsOpacityEffect* self, double opacity);
extern __declspec(dllexport) void QGraphicsOpacityEffect_SetOpacityMask(QGraphicsOpacityEffect* self, QBrush* mask);
extern __declspec(dllexport) void QGraphicsOpacityEffect_OpacityChanged(QGraphicsOpacityEffect* self, double opacity);
void QGraphicsOpacityEffect_connect_OpacityChanged(QGraphicsOpacityEffect* self, intptr_t slot);
extern __declspec(dllexport) void QGraphicsOpacityEffect_OpacityMaskChanged(QGraphicsOpacityEffect* self, QBrush* mask);
void QGraphicsOpacityEffect_connect_OpacityMaskChanged(QGraphicsOpacityEffect* self, intptr_t slot);
extern __declspec(dllexport) void QGraphicsOpacityEffect_Draw(QGraphicsOpacityEffect* self, QPainter* painter);
extern __declspec(dllexport) struct miqt_string QGraphicsOpacityEffect_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QGraphicsOpacityEffect_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QGraphicsOpacityEffect_override_virtual_Draw(void* self, intptr_t slot);
void QGraphicsOpacityEffect_virtualbase_Draw(void* self, QPainter* painter);
extern __declspec(dllexport) void QGraphicsOpacityEffect_override_virtual_BoundingRectFor(void* self, intptr_t slot);
QRectF* QGraphicsOpacityEffect_virtualbase_BoundingRectFor(const void* self, QRectF* sourceRect);
extern __declspec(dllexport) void QGraphicsOpacityEffect_override_virtual_SourceChanged(void* self, intptr_t slot);
void QGraphicsOpacityEffect_virtualbase_SourceChanged(void* self, ChangeFlags flags);
extern __declspec(dllexport) void QGraphicsOpacityEffect_Delete(QGraphicsOpacityEffect* self, bool isSubclass);

} 
