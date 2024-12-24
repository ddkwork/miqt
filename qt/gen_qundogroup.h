#pragma once
#ifndef MIQT_QT_GEN_QUNDOGROUP_H
#define MIQT_QT_GEN_QUNDOGROUP_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAction QAction;
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;
typedef struct QUndoGroup QUndoGroup;
typedef struct QUndoStack QUndoStack;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QUndoGroup* QUndoGroup_new();
extern __declspec(dllexport) QUndoGroup* QUndoGroup_new2(QObject* parent);
extern __declspec(dllexport) void QUndoGroup_virtbase(QUndoGroup* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QUndoGroup_MetaObject(const QUndoGroup* self);
extern __declspec(dllexport) void* QUndoGroup_Metacast(QUndoGroup* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QUndoGroup_Tr(const char* s);
extern __declspec(dllexport) void QUndoGroup_AddStack(QUndoGroup* self, QUndoStack* stack);
extern __declspec(dllexport) void QUndoGroup_RemoveStack(QUndoGroup* self, QUndoStack* stack);
extern __declspec(dllexport) struct miqt_array /* of QUndoStack* */  QUndoGroup_Stacks(const QUndoGroup* self);
extern __declspec(dllexport) QUndoStack* QUndoGroup_ActiveStack(const QUndoGroup* self);
extern __declspec(dllexport) QAction* QUndoGroup_CreateUndoAction(const QUndoGroup* self, QObject* parent);
extern __declspec(dllexport) QAction* QUndoGroup_CreateRedoAction(const QUndoGroup* self, QObject* parent);
extern __declspec(dllexport) bool QUndoGroup_CanUndo(const QUndoGroup* self);
extern __declspec(dllexport) bool QUndoGroup_CanRedo(const QUndoGroup* self);
extern __declspec(dllexport) struct miqt_string QUndoGroup_UndoText(const QUndoGroup* self);
extern __declspec(dllexport) struct miqt_string QUndoGroup_RedoText(const QUndoGroup* self);
extern __declspec(dllexport) bool QUndoGroup_IsClean(const QUndoGroup* self);
extern __declspec(dllexport) void QUndoGroup_Undo(QUndoGroup* self);
extern __declspec(dllexport) void QUndoGroup_Redo(QUndoGroup* self);
extern __declspec(dllexport) void QUndoGroup_SetActiveStack(QUndoGroup* self, QUndoStack* stack);
extern __declspec(dllexport) void QUndoGroup_ActiveStackChanged(QUndoGroup* self, QUndoStack* stack);
void QUndoGroup_connect_ActiveStackChanged(QUndoGroup* self, intptr_t slot);
extern __declspec(dllexport) void QUndoGroup_IndexChanged(QUndoGroup* self, int idx);
void QUndoGroup_connect_IndexChanged(QUndoGroup* self, intptr_t slot);
extern __declspec(dllexport) void QUndoGroup_CleanChanged(QUndoGroup* self, bool clean);
void QUndoGroup_connect_CleanChanged(QUndoGroup* self, intptr_t slot);
extern __declspec(dllexport) void QUndoGroup_CanUndoChanged(QUndoGroup* self, bool canUndo);
void QUndoGroup_connect_CanUndoChanged(QUndoGroup* self, intptr_t slot);
extern __declspec(dllexport) void QUndoGroup_CanRedoChanged(QUndoGroup* self, bool canRedo);
void QUndoGroup_connect_CanRedoChanged(QUndoGroup* self, intptr_t slot);
extern __declspec(dllexport) void QUndoGroup_UndoTextChanged(QUndoGroup* self, struct miqt_string undoText);
void QUndoGroup_connect_UndoTextChanged(QUndoGroup* self, intptr_t slot);
extern __declspec(dllexport) void QUndoGroup_RedoTextChanged(QUndoGroup* self, struct miqt_string redoText);
void QUndoGroup_connect_RedoTextChanged(QUndoGroup* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QUndoGroup_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QUndoGroup_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) QAction* QUndoGroup_CreateUndoAction2(const QUndoGroup* self, QObject* parent, struct miqt_string prefix);
extern __declspec(dllexport) QAction* QUndoGroup_CreateRedoAction2(const QUndoGroup* self, QObject* parent, struct miqt_string prefix);
extern __declspec(dllexport) void QUndoGroup_override_virtual_Event(void* self, intptr_t slot);
bool QUndoGroup_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QUndoGroup_override_virtual_EventFilter(void* self, intptr_t slot);
bool QUndoGroup_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QUndoGroup_override_virtual_TimerEvent(void* self, intptr_t slot);
void QUndoGroup_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QUndoGroup_override_virtual_ChildEvent(void* self, intptr_t slot);
void QUndoGroup_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QUndoGroup_override_virtual_CustomEvent(void* self, intptr_t slot);
void QUndoGroup_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QUndoGroup_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QUndoGroup_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QUndoGroup_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QUndoGroup_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QUndoGroup_Delete(QUndoGroup* self, bool isSubclass);

} 
