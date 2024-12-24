#pragma once
#ifndef MIQT_QT_GEN_QSEQUENTIALANIMATIONGROUP_H
#define MIQT_QT_GEN_QSEQUENTIALANIMATIONGROUP_H

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
typedef struct QPauseAnimation QPauseAnimation;
typedef struct QSequentialAnimationGroup QSequentialAnimationGroup;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSequentialAnimationGroup* QSequentialAnimationGroup_new();
extern __declspec(dllexport) QSequentialAnimationGroup* QSequentialAnimationGroup_new2(QObject* parent);
extern __declspec(dllexport) void QSequentialAnimationGroup_virtbase(QSequentialAnimationGroup* src, QAnimationGroup** outptr_QAnimationGroup);
extern __declspec(dllexport) QMetaObject* QSequentialAnimationGroup_MetaObject(const QSequentialAnimationGroup* self);
extern __declspec(dllexport) void* QSequentialAnimationGroup_Metacast(QSequentialAnimationGroup* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QSequentialAnimationGroup_Tr(const char* s);
extern __declspec(dllexport) QPauseAnimation* QSequentialAnimationGroup_AddPause(QSequentialAnimationGroup* self, int msecs);
extern __declspec(dllexport) QPauseAnimation* QSequentialAnimationGroup_InsertPause(QSequentialAnimationGroup* self, int index, int msecs);
extern __declspec(dllexport) QAbstractAnimation* QSequentialAnimationGroup_CurrentAnimation(const QSequentialAnimationGroup* self);
extern __declspec(dllexport) int QSequentialAnimationGroup_Duration(const QSequentialAnimationGroup* self);
extern __declspec(dllexport) void QSequentialAnimationGroup_CurrentAnimationChanged(QSequentialAnimationGroup* self, QAbstractAnimation* current);
void QSequentialAnimationGroup_connect_CurrentAnimationChanged(QSequentialAnimationGroup* self, intptr_t slot);
extern __declspec(dllexport) bool QSequentialAnimationGroup_Event(QSequentialAnimationGroup* self, QEvent* event);
extern __declspec(dllexport) void QSequentialAnimationGroup_UpdateCurrentTime(QSequentialAnimationGroup* self, int param1);
extern __declspec(dllexport) void QSequentialAnimationGroup_UpdateState(QSequentialAnimationGroup* self, int newState, int oldState);
extern __declspec(dllexport) void QSequentialAnimationGroup_UpdateDirection(QSequentialAnimationGroup* self, int direction);
extern __declspec(dllexport) struct miqt_string QSequentialAnimationGroup_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QSequentialAnimationGroup_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QSequentialAnimationGroup_override_virtual_Duration(void* self, intptr_t slot);
int QSequentialAnimationGroup_virtualbase_Duration(const void* self);
extern __declspec(dllexport) void QSequentialAnimationGroup_override_virtual_Event(void* self, intptr_t slot);
bool QSequentialAnimationGroup_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QSequentialAnimationGroup_override_virtual_UpdateCurrentTime(void* self, intptr_t slot);
void QSequentialAnimationGroup_virtualbase_UpdateCurrentTime(void* self, int param1);
extern __declspec(dllexport) void QSequentialAnimationGroup_override_virtual_UpdateState(void* self, intptr_t slot);
void QSequentialAnimationGroup_virtualbase_UpdateState(void* self, int newState, int oldState);
extern __declspec(dllexport) void QSequentialAnimationGroup_override_virtual_UpdateDirection(void* self, intptr_t slot);
void QSequentialAnimationGroup_virtualbase_UpdateDirection(void* self, int direction);
extern __declspec(dllexport) void QSequentialAnimationGroup_Delete(QSequentialAnimationGroup* self, bool isSubclass);

} 
