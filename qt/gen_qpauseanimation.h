#pragma once
#ifndef MIQT_QT_GEN_QPAUSEANIMATION_H
#define MIQT_QT_GEN_QPAUSEANIMATION_H

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
class QPauseAnimation;
class _GUID;
class type_info;
#else
typedef struct QAbstractAnimation QAbstractAnimation;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPauseAnimation QPauseAnimation;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QPauseAnimation* QPauseAnimation_new();
extern __declspec(dllexport) QPauseAnimation* QPauseAnimation_new2(int msecs);
extern __declspec(dllexport) QPauseAnimation* QPauseAnimation_new3(QObject* parent);
extern __declspec(dllexport) QPauseAnimation* QPauseAnimation_new4(int msecs, QObject* parent);
extern __declspec(dllexport) void QPauseAnimation_virtbase(QPauseAnimation* src, QAbstractAnimation** outptr_QAbstractAnimation);
extern __declspec(dllexport) QMetaObject* QPauseAnimation_MetaObject(const QPauseAnimation* self);
extern __declspec(dllexport) void* QPauseAnimation_Metacast(QPauseAnimation* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QPauseAnimation_Tr(const char* s);
extern __declspec(dllexport) int QPauseAnimation_Duration(const QPauseAnimation* self);
extern __declspec(dllexport) void QPauseAnimation_SetDuration(QPauseAnimation* self, int msecs);
extern __declspec(dllexport) bool QPauseAnimation_Event(QPauseAnimation* self, QEvent* e);
extern __declspec(dllexport) void QPauseAnimation_UpdateCurrentTime(QPauseAnimation* self, int param1);
extern __declspec(dllexport) struct miqt_string QPauseAnimation_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QPauseAnimation_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QPauseAnimation_override_virtual_Duration(void* self, intptr_t slot);
int QPauseAnimation_virtualbase_Duration(const void* self);
extern __declspec(dllexport) void QPauseAnimation_override_virtual_Event(void* self, intptr_t slot);
bool QPauseAnimation_virtualbase_Event(void* self, QEvent* e);
extern __declspec(dllexport) void QPauseAnimation_override_virtual_UpdateCurrentTime(void* self, intptr_t slot);
void QPauseAnimation_virtualbase_UpdateCurrentTime(void* self, int param1);
extern __declspec(dllexport) void QPauseAnimation_override_virtual_UpdateState(void* self, intptr_t slot);
void QPauseAnimation_virtualbase_UpdateState(void* self, int newState, int oldState);
extern __declspec(dllexport) void QPauseAnimation_override_virtual_UpdateDirection(void* self, intptr_t slot);
void QPauseAnimation_virtualbase_UpdateDirection(void* self, int direction);
extern __declspec(dllexport) void QPauseAnimation_Delete(QPauseAnimation* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
