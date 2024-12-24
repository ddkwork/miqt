#pragma once
#ifndef MIQT_QT_GEN_QGESTURE_H
#define MIQT_QT_GEN_QGESTURE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChildEvent;
class QEvent;
class QGesture;
class QGestureEvent;
class QMetaMethod;
class QMetaObject;
class QObject;
class QPanGesture;
class QPinchGesture;
class QPointF;
class QSwipeGesture;
class QTapAndHoldGesture;
class QTapGesture;
class QTimerEvent;
class QWidget;
class _GUID;
class type_info;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QGesture QGesture;
typedef struct QGestureEvent QGestureEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPanGesture QPanGesture;
typedef struct QPinchGesture QPinchGesture;
typedef struct QPointF QPointF;
typedef struct QSwipeGesture QSwipeGesture;
typedef struct QTapAndHoldGesture QTapAndHoldGesture;
typedef struct QTapGesture QTapGesture;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QGesture* QGesture_new();
extern __declspec(dllexport) QGesture* QGesture_new2(QObject* parent);
extern __declspec(dllexport) void QGesture_virtbase(QGesture* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QGesture_MetaObject(const QGesture* self);
extern __declspec(dllexport) void* QGesture_Metacast(QGesture* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QGesture_Tr(const char* s);
extern __declspec(dllexport) int QGesture_GestureType(const QGesture* self);
extern __declspec(dllexport) int QGesture_State(const QGesture* self);
extern __declspec(dllexport) QPointF* QGesture_HotSpot(const QGesture* self);
extern __declspec(dllexport) void QGesture_SetHotSpot(QGesture* self, QPointF* value);
extern __declspec(dllexport) bool QGesture_HasHotSpot(const QGesture* self);
extern __declspec(dllexport) void QGesture_UnsetHotSpot(QGesture* self);
extern __declspec(dllexport) void QGesture_SetGestureCancelPolicy(QGesture* self, GestureCancelPolicy policy);
extern __declspec(dllexport) GestureCancelPolicy QGesture_GestureCancelPolicy(const QGesture* self);
extern __declspec(dllexport) struct miqt_string QGesture_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QGesture_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QGesture_override_virtual_Event(void* self, intptr_t slot);
bool QGesture_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QGesture_override_virtual_EventFilter(void* self, intptr_t slot);
bool QGesture_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QGesture_override_virtual_TimerEvent(void* self, intptr_t slot);
void QGesture_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QGesture_override_virtual_ChildEvent(void* self, intptr_t slot);
void QGesture_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QGesture_override_virtual_CustomEvent(void* self, intptr_t slot);
void QGesture_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QGesture_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QGesture_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QGesture_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QGesture_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QGesture_Delete(QGesture* self, bool isSubclass);

extern __declspec(dllexport) QPanGesture* QPanGesture_new();
extern __declspec(dllexport) QPanGesture* QPanGesture_new2(QObject* parent);
extern __declspec(dllexport) void QPanGesture_virtbase(QPanGesture* src, QGesture** outptr_QGesture);
extern __declspec(dllexport) QMetaObject* QPanGesture_MetaObject(const QPanGesture* self);
extern __declspec(dllexport) void* QPanGesture_Metacast(QPanGesture* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QPanGesture_Tr(const char* s);
extern __declspec(dllexport) QPointF* QPanGesture_LastOffset(const QPanGesture* self);
extern __declspec(dllexport) QPointF* QPanGesture_Offset(const QPanGesture* self);
extern __declspec(dllexport) QPointF* QPanGesture_Delta(const QPanGesture* self);
extern __declspec(dllexport) double QPanGesture_Acceleration(const QPanGesture* self);
extern __declspec(dllexport) void QPanGesture_SetLastOffset(QPanGesture* self, QPointF* value);
extern __declspec(dllexport) void QPanGesture_SetOffset(QPanGesture* self, QPointF* value);
extern __declspec(dllexport) void QPanGesture_SetAcceleration(QPanGesture* self, double value);
extern __declspec(dllexport) struct miqt_string QPanGesture_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QPanGesture_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QPanGesture_Delete(QPanGesture* self, bool isSubclass);

extern __declspec(dllexport) QPinchGesture* QPinchGesture_new();
extern __declspec(dllexport) QPinchGesture* QPinchGesture_new2(QObject* parent);
extern __declspec(dllexport) void QPinchGesture_virtbase(QPinchGesture* src, QGesture** outptr_QGesture);
extern __declspec(dllexport) QMetaObject* QPinchGesture_MetaObject(const QPinchGesture* self);
extern __declspec(dllexport) void* QPinchGesture_Metacast(QPinchGesture* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QPinchGesture_Tr(const char* s);
extern __declspec(dllexport) ChangeFlags QPinchGesture_TotalChangeFlags(const QPinchGesture* self);
extern __declspec(dllexport) void QPinchGesture_SetTotalChangeFlags(QPinchGesture* self, ChangeFlags value);
extern __declspec(dllexport) ChangeFlags QPinchGesture_ChangeFlags(const QPinchGesture* self);
extern __declspec(dllexport) void QPinchGesture_SetChangeFlags(QPinchGesture* self, ChangeFlags value);
extern __declspec(dllexport) QPointF* QPinchGesture_StartCenterPoint(const QPinchGesture* self);
extern __declspec(dllexport) QPointF* QPinchGesture_LastCenterPoint(const QPinchGesture* self);
extern __declspec(dllexport) QPointF* QPinchGesture_CenterPoint(const QPinchGesture* self);
extern __declspec(dllexport) void QPinchGesture_SetStartCenterPoint(QPinchGesture* self, QPointF* value);
extern __declspec(dllexport) void QPinchGesture_SetLastCenterPoint(QPinchGesture* self, QPointF* value);
extern __declspec(dllexport) void QPinchGesture_SetCenterPoint(QPinchGesture* self, QPointF* value);
extern __declspec(dllexport) double QPinchGesture_TotalScaleFactor(const QPinchGesture* self);
extern __declspec(dllexport) double QPinchGesture_LastScaleFactor(const QPinchGesture* self);
extern __declspec(dllexport) double QPinchGesture_ScaleFactor(const QPinchGesture* self);
extern __declspec(dllexport) void QPinchGesture_SetTotalScaleFactor(QPinchGesture* self, double value);
extern __declspec(dllexport) void QPinchGesture_SetLastScaleFactor(QPinchGesture* self, double value);
extern __declspec(dllexport) void QPinchGesture_SetScaleFactor(QPinchGesture* self, double value);
extern __declspec(dllexport) double QPinchGesture_TotalRotationAngle(const QPinchGesture* self);
extern __declspec(dllexport) double QPinchGesture_LastRotationAngle(const QPinchGesture* self);
extern __declspec(dllexport) double QPinchGesture_RotationAngle(const QPinchGesture* self);
extern __declspec(dllexport) void QPinchGesture_SetTotalRotationAngle(QPinchGesture* self, double value);
extern __declspec(dllexport) void QPinchGesture_SetLastRotationAngle(QPinchGesture* self, double value);
extern __declspec(dllexport) void QPinchGesture_SetRotationAngle(QPinchGesture* self, double value);
extern __declspec(dllexport) struct miqt_string QPinchGesture_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QPinchGesture_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QPinchGesture_Delete(QPinchGesture* self, bool isSubclass);

extern __declspec(dllexport) QSwipeGesture* QSwipeGesture_new();
extern __declspec(dllexport) QSwipeGesture* QSwipeGesture_new2(QObject* parent);
extern __declspec(dllexport) void QSwipeGesture_virtbase(QSwipeGesture* src, QGesture** outptr_QGesture);
extern __declspec(dllexport) QMetaObject* QSwipeGesture_MetaObject(const QSwipeGesture* self);
extern __declspec(dllexport) void* QSwipeGesture_Metacast(QSwipeGesture* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QSwipeGesture_Tr(const char* s);
extern __declspec(dllexport) SwipeDirection QSwipeGesture_HorizontalDirection(const QSwipeGesture* self);
extern __declspec(dllexport) SwipeDirection QSwipeGesture_VerticalDirection(const QSwipeGesture* self);
extern __declspec(dllexport) double QSwipeGesture_SwipeAngle(const QSwipeGesture* self);
extern __declspec(dllexport) void QSwipeGesture_SetSwipeAngle(QSwipeGesture* self, double value);
extern __declspec(dllexport) struct miqt_string QSwipeGesture_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QSwipeGesture_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QSwipeGesture_Delete(QSwipeGesture* self, bool isSubclass);

extern __declspec(dllexport) QTapGesture* QTapGesture_new();
extern __declspec(dllexport) QTapGesture* QTapGesture_new2(QObject* parent);
extern __declspec(dllexport) void QTapGesture_virtbase(QTapGesture* src, QGesture** outptr_QGesture);
extern __declspec(dllexport) QMetaObject* QTapGesture_MetaObject(const QTapGesture* self);
extern __declspec(dllexport) void* QTapGesture_Metacast(QTapGesture* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QTapGesture_Tr(const char* s);
extern __declspec(dllexport) QPointF* QTapGesture_Position(const QTapGesture* self);
extern __declspec(dllexport) void QTapGesture_SetPosition(QTapGesture* self, QPointF* pos);
extern __declspec(dllexport) struct miqt_string QTapGesture_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QTapGesture_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QTapGesture_Delete(QTapGesture* self, bool isSubclass);

extern __declspec(dllexport) QTapAndHoldGesture* QTapAndHoldGesture_new();
extern __declspec(dllexport) QTapAndHoldGesture* QTapAndHoldGesture_new2(QObject* parent);
extern __declspec(dllexport) void QTapAndHoldGesture_virtbase(QTapAndHoldGesture* src, QGesture** outptr_QGesture);
extern __declspec(dllexport) QMetaObject* QTapAndHoldGesture_MetaObject(const QTapAndHoldGesture* self);
extern __declspec(dllexport) void* QTapAndHoldGesture_Metacast(QTapAndHoldGesture* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QTapAndHoldGesture_Tr(const char* s);
extern __declspec(dllexport) QPointF* QTapAndHoldGesture_Position(const QTapAndHoldGesture* self);
extern __declspec(dllexport) void QTapAndHoldGesture_SetPosition(QTapAndHoldGesture* self, QPointF* pos);
extern __declspec(dllexport) void QTapAndHoldGesture_SetTimeout(int msecs);
extern __declspec(dllexport) int QTapAndHoldGesture_Timeout();
extern __declspec(dllexport) struct miqt_string QTapAndHoldGesture_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QTapAndHoldGesture_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QTapAndHoldGesture_Delete(QTapAndHoldGesture* self, bool isSubclass);

extern __declspec(dllexport) QGestureEvent* QGestureEvent_new(struct miqt_array /* of QGesture* */  gestures);
extern __declspec(dllexport) QGestureEvent* QGestureEvent_new2(QGestureEvent* param1);
extern __declspec(dllexport) void QGestureEvent_virtbase(QGestureEvent* src, QEvent** outptr_QEvent);
extern __declspec(dllexport) struct miqt_array /* of QGesture* */  QGestureEvent_Gestures(const QGestureEvent* self);
extern __declspec(dllexport) QGesture* QGestureEvent_Gesture(const QGestureEvent* self, int typeVal);
extern __declspec(dllexport) struct miqt_array /* of QGesture* */  QGestureEvent_ActiveGestures(const QGestureEvent* self);
extern __declspec(dllexport) struct miqt_array /* of QGesture* */  QGestureEvent_CanceledGestures(const QGestureEvent* self);
extern __declspec(dllexport) void QGestureEvent_SetAccepted(QGestureEvent* self, QGesture* param1, bool param2);
extern __declspec(dllexport) void QGestureEvent_Accept(QGestureEvent* self, QGesture* param1);
extern __declspec(dllexport) void QGestureEvent_Ignore(QGestureEvent* self, QGesture* param1);
extern __declspec(dllexport) bool QGestureEvent_IsAccepted(const QGestureEvent* self, QGesture* param1);
extern __declspec(dllexport) void QGestureEvent_SetAccepted2(QGestureEvent* self, int param1, bool param2);
extern __declspec(dllexport) void QGestureEvent_AcceptWithQtGestureType(QGestureEvent* self, int param1);
extern __declspec(dllexport) void QGestureEvent_IgnoreWithQtGestureType(QGestureEvent* self, int param1);
extern __declspec(dllexport) bool QGestureEvent_IsAcceptedWithQtGestureType(const QGestureEvent* self, int param1);
extern __declspec(dllexport) void QGestureEvent_SetWidget(QGestureEvent* self, QWidget* widget);
extern __declspec(dllexport) QWidget* QGestureEvent_Widget(const QGestureEvent* self);
extern __declspec(dllexport) QPointF* QGestureEvent_MapToGraphicsScene(const QGestureEvent* self, QPointF* gesturePoint);
extern __declspec(dllexport) void QGestureEvent_override_virtual_SetAccepted(void* self, intptr_t slot);
void QGestureEvent_virtualbase_SetAccepted(void* self, bool accepted);
extern __declspec(dllexport) void QGestureEvent_override_virtual_Clone(void* self, intptr_t slot);
QEvent* QGestureEvent_virtualbase_Clone(const void* self);
extern __declspec(dllexport) void QGestureEvent_Delete(QGestureEvent* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
