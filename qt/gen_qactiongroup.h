#pragma once
#ifndef MIQT_QT_GEN_QACTIONGROUP_H
#define MIQT_QT_GEN_QACTIONGROUP_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAction QAction;
typedef struct QActionGroup QActionGroup;
typedef struct QIcon QIcon;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QActionGroup* QActionGroup_new(QObject* parent);
extern __declspec(dllexport) 
void QActionGroup_virtbase(QActionGroup* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QActionGroup_MetaObject(const QActionGroup* self);
extern __declspec(dllexport) 
void* QActionGroup_Metacast(QActionGroup* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QActionGroup_Tr(const char* s);
extern __declspec(dllexport) 
QAction* QActionGroup_AddAction(QActionGroup* self, QAction* a);
extern __declspec(dllexport) 
QAction* QActionGroup_AddActionWithText(QActionGroup* self, struct miqt_string text);
extern __declspec(dllexport) 
QAction* QActionGroup_AddAction2(QActionGroup* self, QIcon* icon, struct miqt_string text);
extern __declspec(dllexport) 
void QActionGroup_RemoveAction(QActionGroup* self, QAction* a);
extern __declspec(dllexport) 
struct miqt_array /* of QAction* */  QActionGroup_Actions(const QActionGroup* self);
extern __declspec(dllexport) 
QAction* QActionGroup_CheckedAction(const QActionGroup* self);
extern __declspec(dllexport) 
bool QActionGroup_IsExclusive(const QActionGroup* self);
extern __declspec(dllexport) 
bool QActionGroup_IsEnabled(const QActionGroup* self);
extern __declspec(dllexport) 
bool QActionGroup_IsVisible(const QActionGroup* self);
extern __declspec(dllexport) 
ExclusionPolicy QActionGroup_ExclusionPolicy(const QActionGroup* self);
extern __declspec(dllexport) 
void QActionGroup_SetEnabled(QActionGroup* self, bool enabled);
extern __declspec(dllexport) 
void QActionGroup_SetDisabled(QActionGroup* self, bool b);
extern __declspec(dllexport) 
void QActionGroup_SetVisible(QActionGroup* self, bool visible);
extern __declspec(dllexport) 
void QActionGroup_SetExclusive(QActionGroup* self, bool exclusive);
extern __declspec(dllexport) 
void QActionGroup_SetExclusionPolicy(QActionGroup* self, ExclusionPolicy policy);
extern __declspec(dllexport) 
void QActionGroup_Triggered(QActionGroup* self, QAction* param1);
void QActionGroup_connect_Triggered(QActionGroup* self, intptr_t slot);
extern __declspec(dllexport) 
void QActionGroup_Hovered(QActionGroup* self, QAction* param1);
void QActionGroup_connect_Hovered(QActionGroup* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QActionGroup_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QActionGroup_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QActionGroup_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QActionGroup_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QActionGroup_override_virtual_Metacast(void* self, intptr_t slot);
void* QActionGroup_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QActionGroup_Delete(QActionGroup* self, bool isSubclass);

}
