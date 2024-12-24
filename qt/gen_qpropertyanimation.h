#pragma once
#ifndef MIQT_QT_GEN_QPROPERTYANIMATION_H
#define MIQT_QT_GEN_QPROPERTYANIMATION_H

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
class QEvent;
class QMetaObject;
class QObject;
class QPropertyAnimation;
class QVariant;
class QVariantAnimation;
class _GUID;
class type_info;
#else
typedef struct QAbstractAnimation QAbstractAnimation;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPropertyAnimation QPropertyAnimation;
typedef struct QVariant QVariant;
typedef struct QVariantAnimation QVariantAnimation;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QPropertyAnimation* QPropertyAnimation_new();
extern __declspec(dllexport) QPropertyAnimation* QPropertyAnimation_new2(QObject* target, struct miqt_string propertyName);
extern __declspec(dllexport) QPropertyAnimation* QPropertyAnimation_new3(QObject* parent);
extern __declspec(dllexport) QPropertyAnimation* QPropertyAnimation_new4(QObject* target, struct miqt_string propertyName, QObject* parent);
extern __declspec(dllexport) void QPropertyAnimation_virtbase(QPropertyAnimation* src, QVariantAnimation** outptr_QVariantAnimation);
extern __declspec(dllexport) QMetaObject* QPropertyAnimation_MetaObject(const QPropertyAnimation* self);
extern __declspec(dllexport) void* QPropertyAnimation_Metacast(QPropertyAnimation* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QPropertyAnimation_Tr(const char* s);
extern __declspec(dllexport) QObject* QPropertyAnimation_TargetObject(const QPropertyAnimation* self);
extern __declspec(dllexport) void QPropertyAnimation_SetTargetObject(QPropertyAnimation* self, QObject* target);
extern __declspec(dllexport) struct miqt_string QPropertyAnimation_PropertyName(const QPropertyAnimation* self);
extern __declspec(dllexport) void QPropertyAnimation_SetPropertyName(QPropertyAnimation* self, struct miqt_string propertyName);
extern __declspec(dllexport) bool QPropertyAnimation_Event(QPropertyAnimation* self, QEvent* event);
extern __declspec(dllexport) void QPropertyAnimation_UpdateCurrentValue(QPropertyAnimation* self, QVariant* value);
extern __declspec(dllexport) void QPropertyAnimation_UpdateState(QPropertyAnimation* self, int newState, int oldState);
extern __declspec(dllexport) struct miqt_string QPropertyAnimation_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QPropertyAnimation_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QPropertyAnimation_override_virtual_Event(void* self, intptr_t slot);
bool QPropertyAnimation_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QPropertyAnimation_override_virtual_UpdateCurrentValue(void* self, intptr_t slot);
void QPropertyAnimation_virtualbase_UpdateCurrentValue(void* self, QVariant* value);
extern __declspec(dllexport) void QPropertyAnimation_override_virtual_UpdateState(void* self, intptr_t slot);
void QPropertyAnimation_virtualbase_UpdateState(void* self, int newState, int oldState);
extern __declspec(dllexport) void QPropertyAnimation_override_virtual_Duration(void* self, intptr_t slot);
int QPropertyAnimation_virtualbase_Duration(const void* self);
extern __declspec(dllexport) void QPropertyAnimation_override_virtual_UpdateCurrentTime(void* self, intptr_t slot);
void QPropertyAnimation_virtualbase_UpdateCurrentTime(void* self, int param1);
extern __declspec(dllexport) void QPropertyAnimation_override_virtual_Interpolated(void* self, intptr_t slot);
QVariant* QPropertyAnimation_virtualbase_Interpolated(const void* self, QVariant* from, QVariant* to, double progress);
extern __declspec(dllexport) void QPropertyAnimation_Delete(QPropertyAnimation* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
