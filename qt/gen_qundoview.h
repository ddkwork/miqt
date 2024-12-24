#pragma once
#ifndef MIQT_QT_GEN_QUNDOVIEW_H
#define MIQT_QT_GEN_QUNDOVIEW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemView QAbstractItemView;
typedef struct QAbstractScrollArea QAbstractScrollArea;
typedef struct QFrame QFrame;
typedef struct QIcon QIcon;
typedef struct QListView QListView;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QUndoGroup QUndoGroup;
typedef struct QUndoStack QUndoStack;
typedef struct QUndoView QUndoView;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QUndoView* QUndoView_new(QWidget* parent);
extern __declspec(dllexport) 
QUndoView* QUndoView_new2();
extern __declspec(dllexport) 
QUndoView* QUndoView_new3(QUndoStack* stack);
extern __declspec(dllexport) 
QUndoView* QUndoView_new4(QUndoGroup* group);
extern __declspec(dllexport) 
QUndoView* QUndoView_new5(QUndoStack* stack, QWidget* parent);
extern __declspec(dllexport) 
QUndoView* QUndoView_new6(QUndoGroup* group, QWidget* parent);
extern __declspec(dllexport) 
void QUndoView_virtbase(QUndoView* src
, QListView** outptr_QListView
);
extern __declspec(dllexport) 
QMetaObject* QUndoView_MetaObject(const QUndoView* self);
extern __declspec(dllexport) 
void* QUndoView_Metacast(QUndoView* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QUndoView_Tr(const char* s);
extern __declspec(dllexport) 
QUndoStack* QUndoView_Stack(const QUndoView* self);
extern __declspec(dllexport) 
QUndoGroup* QUndoView_Group(const QUndoView* self);
extern __declspec(dllexport) 
void QUndoView_SetEmptyLabel(QUndoView* self, struct miqt_string label);
extern __declspec(dllexport) 
struct miqt_string QUndoView_EmptyLabel(const QUndoView* self);
extern __declspec(dllexport) 
void QUndoView_SetCleanIcon(QUndoView* self, QIcon* icon);
extern __declspec(dllexport) 
QIcon* QUndoView_CleanIcon(const QUndoView* self);
extern __declspec(dllexport) 
void QUndoView_SetStack(QUndoView* self, QUndoStack* stack);
extern __declspec(dllexport) 
void QUndoView_SetGroup(QUndoView* self, QUndoGroup* group);
extern __declspec(dllexport) 
struct miqt_string QUndoView_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QUndoView_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QUndoView_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QUndoView_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QUndoView_override_virtual_Metacast(void* self, intptr_t slot);
void* QUndoView_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QUndoView_Delete(QUndoView* self, bool isSubclass);

}
