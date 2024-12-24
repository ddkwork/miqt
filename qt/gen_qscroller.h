#pragma once
#ifndef MIQT_QT_GEN_QSCROLLER_H
#define MIQT_QT_GEN_QSCROLLER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPointF QPointF;
typedef struct QRectF QRectF;
typedef struct QScroller QScroller;
typedef struct QScrollerProperties QScrollerProperties;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QScroller_virtbase(QScroller* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QScroller_MetaObject(const QScroller* self);
extern __declspec(dllexport) void* QScroller_Metacast(QScroller* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QScroller_Tr(const char* s);
extern __declspec(dllexport) bool QScroller_HasScroller(QObject* target);
extern __declspec(dllexport) QScroller* QScroller_Scroller(QObject* target);
extern __declspec(dllexport) QScroller* QScroller_ScrollerWithTarget(QObject* target);
extern __declspec(dllexport) int QScroller_GrabGesture(QObject* target);
extern __declspec(dllexport) int QScroller_GrabbedGesture(QObject* target);
extern __declspec(dllexport) void QScroller_UngrabGesture(QObject* target);
extern __declspec(dllexport) struct miqt_array /* of QScroller* */  QScroller_ActiveScrollers();
extern __declspec(dllexport) QObject* QScroller_Target(const QScroller* self);
extern __declspec(dllexport) State QScroller_State(const QScroller* self);
extern __declspec(dllexport) bool QScroller_HandleInput(QScroller* self, Input input, QPointF* position);
extern __declspec(dllexport) void QScroller_Stop(QScroller* self);
extern __declspec(dllexport) QPointF* QScroller_Velocity(const QScroller* self);
extern __declspec(dllexport) QPointF* QScroller_FinalPosition(const QScroller* self);
extern __declspec(dllexport) QPointF* QScroller_PixelPerMeter(const QScroller* self);
extern __declspec(dllexport) QScrollerProperties* QScroller_ScrollerProperties(const QScroller* self);
extern __declspec(dllexport) void QScroller_SetSnapPositionsX(QScroller* self, struct miqt_array /* of double */  positions);
extern __declspec(dllexport) void QScroller_SetSnapPositionsX2(QScroller* self, double first, double interval);
extern __declspec(dllexport) void QScroller_SetSnapPositionsY(QScroller* self, struct miqt_array /* of double */  positions);
extern __declspec(dllexport) void QScroller_SetSnapPositionsY2(QScroller* self, double first, double interval);
extern __declspec(dllexport) void QScroller_SetScrollerProperties(QScroller* self, QScrollerProperties* prop);
extern __declspec(dllexport) void QScroller_ScrollTo(QScroller* self, QPointF* pos);
extern __declspec(dllexport) void QScroller_ScrollTo2(QScroller* self, QPointF* pos, int scrollTime);
extern __declspec(dllexport) void QScroller_EnsureVisible(QScroller* self, QRectF* rect, double xmargin, double ymargin);
extern __declspec(dllexport) void QScroller_EnsureVisible2(QScroller* self, QRectF* rect, double xmargin, double ymargin, int scrollTime);
extern __declspec(dllexport) void QScroller_ResendPrepareEvent(QScroller* self);
extern __declspec(dllexport) void QScroller_StateChanged(QScroller* self, int newstate);
void QScroller_connect_StateChanged(QScroller* self, intptr_t slot);
extern __declspec(dllexport) void QScroller_ScrollerPropertiesChanged(QScroller* self, QScrollerProperties* param1);
void QScroller_connect_ScrollerPropertiesChanged(QScroller* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QScroller_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QScroller_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) int QScroller_GrabGesture2(QObject* target, ScrollerGestureType gestureType);
extern __declspec(dllexport) bool QScroller_HandleInput3(QScroller* self, Input input, QPointF* position, long long timestamp);

} 
