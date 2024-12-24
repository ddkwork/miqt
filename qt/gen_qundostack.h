#pragma once
#ifndef MIQT_QT_GEN_QUNDOSTACK_H
#define MIQT_QT_GEN_QUNDOSTACK_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAction QAction;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QUndoCommand QUndoCommand;
typedef struct QUndoStack QUndoStack;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QUndoCommand* QUndoCommand_new();
extern __declspec(dllexport) 
QUndoCommand* QUndoCommand_new2(struct miqt_string text);
extern __declspec(dllexport) 
QUndoCommand* QUndoCommand_new3(QUndoCommand* parent);
extern __declspec(dllexport) 
QUndoCommand* QUndoCommand_new4(struct miqt_string text, QUndoCommand* parent);
extern __declspec(dllexport) 
void QUndoCommand_Undo(QUndoCommand* self);
extern __declspec(dllexport) 
void QUndoCommand_Redo(QUndoCommand* self);
extern __declspec(dllexport) 
struct miqt_string QUndoCommand_Text(const QUndoCommand* self);
extern __declspec(dllexport) 
struct miqt_string QUndoCommand_ActionText(const QUndoCommand* self);
extern __declspec(dllexport) 
void QUndoCommand_SetText(QUndoCommand* self, struct miqt_string text);
extern __declspec(dllexport) 
bool QUndoCommand_IsObsolete(const QUndoCommand* self);
extern __declspec(dllexport) 
void QUndoCommand_SetObsolete(QUndoCommand* self, bool obsolete);
extern __declspec(dllexport) 
int QUndoCommand_Id(const QUndoCommand* self);
extern __declspec(dllexport) 
bool QUndoCommand_MergeWith(QUndoCommand* self, QUndoCommand* other);
extern __declspec(dllexport) 
int QUndoCommand_ChildCount(const QUndoCommand* self);
extern __declspec(dllexport) 
QUndoCommand* QUndoCommand_Child(const QUndoCommand* self, int index);
extern __declspec(dllexport) 
void QUndoCommand_Delete(QUndoCommand* self, bool isSubclass);

extern __declspec(dllexport) 
QUndoStack* QUndoStack_new();
extern __declspec(dllexport) 
QUndoStack* QUndoStack_new2(QObject* parent);
extern __declspec(dllexport) 
void QUndoStack_virtbase(QUndoStack* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QUndoStack_MetaObject(const QUndoStack* self);
extern __declspec(dllexport) 
void* QUndoStack_Metacast(QUndoStack* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QUndoStack_Tr(const char* s);
extern __declspec(dllexport) 
void QUndoStack_Clear(QUndoStack* self);
extern __declspec(dllexport) 
void QUndoStack_Push(QUndoStack* self, QUndoCommand* cmd);
extern __declspec(dllexport) 
bool QUndoStack_CanUndo(const QUndoStack* self);
extern __declspec(dllexport) 
bool QUndoStack_CanRedo(const QUndoStack* self);
extern __declspec(dllexport) 
struct miqt_string QUndoStack_UndoText(const QUndoStack* self);
extern __declspec(dllexport) 
struct miqt_string QUndoStack_RedoText(const QUndoStack* self);
extern __declspec(dllexport) 
int QUndoStack_Count(const QUndoStack* self);
extern __declspec(dllexport) 
int QUndoStack_Index(const QUndoStack* self);
extern __declspec(dllexport) 
struct miqt_string QUndoStack_Text(const QUndoStack* self, int idx);
extern __declspec(dllexport) 
QAction* QUndoStack_CreateUndoAction(const QUndoStack* self, QObject* parent);
extern __declspec(dllexport) 
QAction* QUndoStack_CreateRedoAction(const QUndoStack* self, QObject* parent);
extern __declspec(dllexport) 
bool QUndoStack_IsActive(const QUndoStack* self);
extern __declspec(dllexport) 
bool QUndoStack_IsClean(const QUndoStack* self);
extern __declspec(dllexport) 
int QUndoStack_CleanIndex(const QUndoStack* self);
extern __declspec(dllexport) 
void QUndoStack_BeginMacro(QUndoStack* self, struct miqt_string text);
extern __declspec(dllexport) 
void QUndoStack_EndMacro(QUndoStack* self);
extern __declspec(dllexport) 
void QUndoStack_SetUndoLimit(QUndoStack* self, int limit);
extern __declspec(dllexport) 
int QUndoStack_UndoLimit(const QUndoStack* self);
extern __declspec(dllexport) 
QUndoCommand* QUndoStack_Command(const QUndoStack* self, int index);
extern __declspec(dllexport) 
void QUndoStack_SetClean(QUndoStack* self);
extern __declspec(dllexport) 
void QUndoStack_ResetClean(QUndoStack* self);
extern __declspec(dllexport) 
void QUndoStack_SetIndex(QUndoStack* self, int idx);
extern __declspec(dllexport) 
void QUndoStack_Undo(QUndoStack* self);
extern __declspec(dllexport) 
void QUndoStack_Redo(QUndoStack* self);
extern __declspec(dllexport) 
void QUndoStack_SetActive(QUndoStack* self);
extern __declspec(dllexport) 
void QUndoStack_IndexChanged(QUndoStack* self, int idx);
void QUndoStack_connect_IndexChanged(QUndoStack* self, intptr_t slot);
extern __declspec(dllexport) 
void QUndoStack_CleanChanged(QUndoStack* self, bool clean);
void QUndoStack_connect_CleanChanged(QUndoStack* self, intptr_t slot);
extern __declspec(dllexport) 
void QUndoStack_CanUndoChanged(QUndoStack* self, bool canUndo);
void QUndoStack_connect_CanUndoChanged(QUndoStack* self, intptr_t slot);
extern __declspec(dllexport) 
void QUndoStack_CanRedoChanged(QUndoStack* self, bool canRedo);
void QUndoStack_connect_CanRedoChanged(QUndoStack* self, intptr_t slot);
extern __declspec(dllexport) 
void QUndoStack_UndoTextChanged(QUndoStack* self, struct miqt_string undoText);
void QUndoStack_connect_UndoTextChanged(QUndoStack* self, intptr_t slot);
extern __declspec(dllexport) 
void QUndoStack_RedoTextChanged(QUndoStack* self, struct miqt_string redoText);
void QUndoStack_connect_RedoTextChanged(QUndoStack* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QUndoStack_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QUndoStack_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
QAction* QUndoStack_CreateUndoAction2(const QUndoStack* self, QObject* parent, struct miqt_string prefix);
extern __declspec(dllexport) 
QAction* QUndoStack_CreateRedoAction2(const QUndoStack* self, QObject* parent, struct miqt_string prefix);
extern __declspec(dllexport) 
void QUndoStack_SetActive1(QUndoStack* self, bool active);
extern __declspec(dllexport) 
void QUndoStack_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QUndoStack_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QUndoStack_override_virtual_Metacast(void* self, intptr_t slot);
void* QUndoStack_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QUndoStack_Delete(QUndoStack* self, bool isSubclass);

}
