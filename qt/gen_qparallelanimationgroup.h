#pragma once
#ifndef MIQT_QT_GEN_QPARALLELANIMATIONGROUP_H
#define MIQT_QT_GEN_QPARALLELANIMATIONGROUP_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractAnimation QAbstractAnimation;
typedef struct QAnimationGroup QAnimationGroup;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QParallelAnimationGroup QParallelAnimationGroup;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QParallelAnimationGroup* QParallelAnimationGroup_new();
extern __declspec(dllexport) QParallelAnimationGroup* QParallelAnimationGroup_new2(QObject* parent);
extern __declspec(dllexport) void QParallelAnimationGroup_virtbase(QParallelAnimationGroup* src, QAnimationGroup** outptr_QAnimationGroup);
extern __declspec(dllexport) QMetaObject* QParallelAnimationGroup_MetaObject(const QParallelAnimationGroup* self);
extern __declspec(dllexport) void* QParallelAnimationGroup_Metacast(QParallelAnimationGroup* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QParallelAnimationGroup_Tr(const char* s);
extern __declspec(dllexport) int QParallelAnimationGroup_Duration(const QParallelAnimationGroup* self);
extern __declspec(dllexport) bool QParallelAnimationGroup_Event(QParallelAnimationGroup* self, QEvent* event);
extern __declspec(dllexport) void QParallelAnimationGroup_UpdateCurrentTime(QParallelAnimationGroup* self, int currentTime);
extern __declspec(dllexport) void QParallelAnimationGroup_UpdateState(QParallelAnimationGroup* self, int newState, int oldState);
extern __declspec(dllexport) void QParallelAnimationGroup_UpdateDirection(QParallelAnimationGroup* self, int direction);
extern __declspec(dllexport) struct miqt_string QParallelAnimationGroup_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QParallelAnimationGroup_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QParallelAnimationGroup_override_virtual_Duration(void* self, intptr_t slot);
int QParallelAnimationGroup_virtualbase_Duration(const void* self);
extern __declspec(dllexport) void QParallelAnimationGroup_override_virtual_Event(void* self, intptr_t slot);
bool QParallelAnimationGroup_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QParallelAnimationGroup_override_virtual_UpdateCurrentTime(void* self, intptr_t slot);
void QParallelAnimationGroup_virtualbase_UpdateCurrentTime(void* self, int currentTime);
extern __declspec(dllexport) void QParallelAnimationGroup_override_virtual_UpdateState(void* self, intptr_t slot);
void QParallelAnimationGroup_virtualbase_UpdateState(void* self, int newState, int oldState);
extern __declspec(dllexport) void QParallelAnimationGroup_override_virtual_UpdateDirection(void* self, intptr_t slot);
void QParallelAnimationGroup_virtualbase_UpdateDirection(void* self, int direction);
extern __declspec(dllexport) void QParallelAnimationGroup_Delete(QParallelAnimationGroup* self, bool isSubclass);

} 
