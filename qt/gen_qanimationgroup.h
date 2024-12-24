#pragma once
#ifndef MIQT_QT_GEN_QANIMATIONGROUP_H
#define MIQT_QT_GEN_QANIMATIONGROUP_H

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

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAnimationGroup* QAnimationGroup_new();
extern __declspec(dllexport) 
QAnimationGroup* QAnimationGroup_new2(QObject* parent);
extern __declspec(dllexport) 
void QAnimationGroup_virtbase(QAnimationGroup* src
, QAbstractAnimation** outptr_QAbstractAnimation
);
extern __declspec(dllexport) 
QMetaObject* QAnimationGroup_MetaObject(const QAnimationGroup* self);
extern __declspec(dllexport) 
void* QAnimationGroup_Metacast(QAnimationGroup* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAnimationGroup_Tr(const char* s);
extern __declspec(dllexport) 
QAbstractAnimation* QAnimationGroup_AnimationAt(const QAnimationGroup* self, int index);
extern __declspec(dllexport) 
int QAnimationGroup_AnimationCount(const QAnimationGroup* self);
extern __declspec(dllexport) 
int QAnimationGroup_IndexOfAnimation(const QAnimationGroup* self, QAbstractAnimation* animation);
extern __declspec(dllexport) 
void QAnimationGroup_AddAnimation(QAnimationGroup* self, QAbstractAnimation* animation);
extern __declspec(dllexport) 
void QAnimationGroup_InsertAnimation(QAnimationGroup* self, int index, QAbstractAnimation* animation);
extern __declspec(dllexport) 
void QAnimationGroup_RemoveAnimation(QAnimationGroup* self, QAbstractAnimation* animation);
extern __declspec(dllexport) 
QAbstractAnimation* QAnimationGroup_TakeAnimation(QAnimationGroup* self, int index);
extern __declspec(dllexport) 
void QAnimationGroup_Clear(QAnimationGroup* self);
extern __declspec(dllexport) 
bool QAnimationGroup_Event(QAnimationGroup* self, QEvent* event);
extern __declspec(dllexport) 
struct miqt_string QAnimationGroup_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAnimationGroup_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAnimationGroup_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAnimationGroup_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAnimationGroup_override_virtual_Metacast(void* self, intptr_t slot);
void* QAnimationGroup_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAnimationGroup_Delete(QAnimationGroup* self, bool isSubclass);

}
