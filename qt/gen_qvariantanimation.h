#pragma once
#ifndef MIQT_QT_GEN_QVARIANTANIMATION_H
#define MIQT_QT_GEN_QVARIANTANIMATION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractAnimation;
class QEasingCurve;
class QEvent;
class QMetaObject;
class QObject;
class QVariant;
class QVariantAnimation;
class _GUID;
class type_info;
#else
typedef struct QAbstractAnimation QAbstractAnimation;
typedef struct QEasingCurve QEasingCurve;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QVariant QVariant;
typedef struct QVariantAnimation QVariantAnimation;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QVariantAnimation* QVariantAnimation_new();
extern __declspec(dllexport) QVariantAnimation* QVariantAnimation_new2(QObject* parent);
extern __declspec(dllexport) void QVariantAnimation_virtbase(QVariantAnimation* src, QAbstractAnimation** outptr_QAbstractAnimation);
extern __declspec(dllexport) QMetaObject* QVariantAnimation_MetaObject(const QVariantAnimation* self);
extern __declspec(dllexport) void* QVariantAnimation_Metacast(QVariantAnimation* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QVariantAnimation_Tr(const char* s);
extern __declspec(dllexport) QVariant* QVariantAnimation_StartValue(const QVariantAnimation* self);
extern __declspec(dllexport) void QVariantAnimation_SetStartValue(QVariantAnimation* self, QVariant* value);
extern __declspec(dllexport) QVariant* QVariantAnimation_EndValue(const QVariantAnimation* self);
extern __declspec(dllexport) void QVariantAnimation_SetEndValue(QVariantAnimation* self, QVariant* value);
extern __declspec(dllexport) QVariant* QVariantAnimation_KeyValueAt(const QVariantAnimation* self, double step);
extern __declspec(dllexport) void QVariantAnimation_SetKeyValueAt(QVariantAnimation* self, double step, QVariant* value);
extern __declspec(dllexport) KeyValues QVariantAnimation_KeyValues(const QVariantAnimation* self);
extern __declspec(dllexport) void QVariantAnimation_SetKeyValues(QVariantAnimation* self, const KeyValues* values);
extern __declspec(dllexport) QVariant* QVariantAnimation_CurrentValue(const QVariantAnimation* self);
extern __declspec(dllexport) int QVariantAnimation_Duration(const QVariantAnimation* self);
extern __declspec(dllexport) void QVariantAnimation_SetDuration(QVariantAnimation* self, int msecs);
extern __declspec(dllexport) QEasingCurve* QVariantAnimation_EasingCurve(const QVariantAnimation* self);
extern __declspec(dllexport) void QVariantAnimation_SetEasingCurve(QVariantAnimation* self, QEasingCurve* easing);
extern __declspec(dllexport) void QVariantAnimation_ValueChanged(QVariantAnimation* self, QVariant* value);
void QVariantAnimation_connect_ValueChanged(QVariantAnimation* self, intptr_t slot);
extern __declspec(dllexport) bool QVariantAnimation_Event(QVariantAnimation* self, QEvent* event);
extern __declspec(dllexport) void QVariantAnimation_UpdateCurrentTime(QVariantAnimation* self, int param1);
extern __declspec(dllexport) void QVariantAnimation_UpdateState(QVariantAnimation* self, int newState, int oldState);
extern __declspec(dllexport) void QVariantAnimation_UpdateCurrentValue(QVariantAnimation* self, QVariant* value);
extern __declspec(dllexport) QVariant* QVariantAnimation_Interpolated(const QVariantAnimation* self, QVariant* from, QVariant* to, double progress);
extern __declspec(dllexport) struct miqt_string QVariantAnimation_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QVariantAnimation_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QVariantAnimation_override_virtual_Duration(void* self, intptr_t slot);
int QVariantAnimation_virtualbase_Duration(const void* self);
extern __declspec(dllexport) void QVariantAnimation_override_virtual_Event(void* self, intptr_t slot);
bool QVariantAnimation_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QVariantAnimation_override_virtual_UpdateCurrentTime(void* self, intptr_t slot);
void QVariantAnimation_virtualbase_UpdateCurrentTime(void* self, int param1);
extern __declspec(dllexport) void QVariantAnimation_override_virtual_UpdateState(void* self, intptr_t slot);
void QVariantAnimation_virtualbase_UpdateState(void* self, int newState, int oldState);
extern __declspec(dllexport) void QVariantAnimation_override_virtual_UpdateCurrentValue(void* self, intptr_t slot);
void QVariantAnimation_virtualbase_UpdateCurrentValue(void* self, QVariant* value);
extern __declspec(dllexport) void QVariantAnimation_override_virtual_Interpolated(void* self, intptr_t slot);
QVariant* QVariantAnimation_virtualbase_Interpolated(const void* self, QVariant* from, QVariant* to, double progress);
extern __declspec(dllexport) void QVariantAnimation_override_virtual_UpdateDirection(void* self, intptr_t slot);
void QVariantAnimation_virtualbase_UpdateDirection(void* self, int direction);
extern __declspec(dllexport) void QVariantAnimation_Delete(QVariantAnimation* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
