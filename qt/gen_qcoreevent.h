#pragma once
#ifndef MIQT_QT_GEN_QCOREEVENT_H
#define MIQT_QT_GEN_QCOREEVENT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QBasicTimer QBasicTimer;
typedef struct QChildEvent QChildEvent;
typedef struct QDynamicPropertyChangeEvent QDynamicPropertyChangeEvent;
typedef struct QEvent QEvent;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QEvent* QEvent_new(Type typeVal);
extern __declspec(dllexport) Type QEvent_Type(const QEvent* self);
extern __declspec(dllexport) bool QEvent_Spontaneous(const QEvent* self);
extern __declspec(dllexport) void QEvent_SetAccepted(QEvent* self, bool accepted);
extern __declspec(dllexport) bool QEvent_IsAccepted(const QEvent* self);
extern __declspec(dllexport) void QEvent_Accept(QEvent* self);
extern __declspec(dllexport) void QEvent_Ignore(QEvent* self);
extern __declspec(dllexport) bool QEvent_IsInputEvent(const QEvent* self);
extern __declspec(dllexport) bool QEvent_IsPointerEvent(const QEvent* self);
extern __declspec(dllexport) bool QEvent_IsSinglePointEvent(const QEvent* self);
extern __declspec(dllexport) int QEvent_RegisterEventType();
extern __declspec(dllexport) QEvent* QEvent_Clone(const QEvent* self);
extern __declspec(dllexport) int QEvent_RegisterEventType1(int hint);
extern __declspec(dllexport) void QEvent_override_virtual_SetAccepted(void* self, intptr_t slot);
void QEvent_virtualbase_SetAccepted(void* self, bool accepted);
extern __declspec(dllexport) void QEvent_override_virtual_Clone(void* self, intptr_t slot);
QEvent* QEvent_virtualbase_Clone(const void* self);
extern __declspec(dllexport) void QEvent_Delete(QEvent* self, bool isSubclass);

extern __declspec(dllexport) QTimerEvent* QTimerEvent_new(int timerId);
extern __declspec(dllexport) QTimerEvent* QTimerEvent_new2(int timerId);
extern __declspec(dllexport) void QTimerEvent_virtbase(QTimerEvent* src, QEvent** outptr_QEvent);
extern __declspec(dllexport) QTimerEvent* QTimerEvent_Clone(const QTimerEvent* self);
extern __declspec(dllexport) int QTimerEvent_TimerId(const QTimerEvent* self);
extern __declspec(dllexport) int QTimerEvent_Id(const QTimerEvent* self);
extern __declspec(dllexport) bool QTimerEvent_Matches(const QTimerEvent* self, QBasicTimer* timer);
extern __declspec(dllexport) void QTimerEvent_override_virtual_Clone(void* self, intptr_t slot);
QTimerEvent* QTimerEvent_virtualbase_Clone(const void* self);
extern __declspec(dllexport) void QTimerEvent_override_virtual_SetAccepted(void* self, intptr_t slot);
void QTimerEvent_virtualbase_SetAccepted(void* self, bool accepted);
extern __declspec(dllexport) void QTimerEvent_Delete(QTimerEvent* self, bool isSubclass);

extern __declspec(dllexport) QChildEvent* QChildEvent_new(Type typeVal, QObject* child);
extern __declspec(dllexport) void QChildEvent_virtbase(QChildEvent* src, QEvent** outptr_QEvent);
extern __declspec(dllexport) QChildEvent* QChildEvent_Clone(const QChildEvent* self);
extern __declspec(dllexport) QObject* QChildEvent_Child(const QChildEvent* self);
extern __declspec(dllexport) bool QChildEvent_Added(const QChildEvent* self);
extern __declspec(dllexport) bool QChildEvent_Polished(const QChildEvent* self);
extern __declspec(dllexport) bool QChildEvent_Removed(const QChildEvent* self);
extern __declspec(dllexport) void QChildEvent_override_virtual_Clone(void* self, intptr_t slot);
QChildEvent* QChildEvent_virtualbase_Clone(const void* self);
extern __declspec(dllexport) void QChildEvent_override_virtual_SetAccepted(void* self, intptr_t slot);
void QChildEvent_virtualbase_SetAccepted(void* self, bool accepted);
extern __declspec(dllexport) void QChildEvent_Delete(QChildEvent* self, bool isSubclass);

extern __declspec(dllexport) QDynamicPropertyChangeEvent* QDynamicPropertyChangeEvent_new(struct miqt_string name);
extern __declspec(dllexport) void QDynamicPropertyChangeEvent_virtbase(QDynamicPropertyChangeEvent* src, QEvent** outptr_QEvent);
extern __declspec(dllexport) QDynamicPropertyChangeEvent* QDynamicPropertyChangeEvent_Clone(const QDynamicPropertyChangeEvent* self);
extern __declspec(dllexport) struct miqt_string QDynamicPropertyChangeEvent_PropertyName(const QDynamicPropertyChangeEvent* self);
extern __declspec(dllexport) void QDynamicPropertyChangeEvent_override_virtual_Clone(void* self, intptr_t slot);
QDynamicPropertyChangeEvent* QDynamicPropertyChangeEvent_virtualbase_Clone(const void* self);
extern __declspec(dllexport) void QDynamicPropertyChangeEvent_override_virtual_SetAccepted(void* self, intptr_t slot);
void QDynamicPropertyChangeEvent_virtualbase_SetAccepted(void* self, bool accepted);
extern __declspec(dllexport) void QDynamicPropertyChangeEvent_Delete(QDynamicPropertyChangeEvent* self, bool isSubclass);

} 
