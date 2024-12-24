#pragma once
#ifndef MIQT_QT_GEN_QCOMPLETER_H
#define MIQT_QT_GEN_QCOMPLETER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QAbstractItemView QAbstractItemView;
typedef struct QChildEvent QChildEvent;
typedef struct QCompleter QCompleter;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QRect QRect;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QCompleter* QCompleter_new();
extern __declspec(dllexport) QCompleter* QCompleter_new2(QAbstractItemModel* model);
extern __declspec(dllexport) QCompleter* QCompleter_new3(struct miqt_array /* of struct miqt_string */  completions);
extern __declspec(dllexport) QCompleter* QCompleter_new4(QObject* parent);
extern __declspec(dllexport) QCompleter* QCompleter_new5(QAbstractItemModel* model, QObject* parent);
extern __declspec(dllexport) QCompleter* QCompleter_new6(struct miqt_array /* of struct miqt_string */  completions, QObject* parent);
extern __declspec(dllexport) void QCompleter_virtbase(QCompleter* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QCompleter_MetaObject(const QCompleter* self);
extern __declspec(dllexport) void* QCompleter_Metacast(QCompleter* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QCompleter_Tr(const char* s);
extern __declspec(dllexport) void QCompleter_SetWidget(QCompleter* self, QWidget* widget);
extern __declspec(dllexport) QWidget* QCompleter_Widget(const QCompleter* self);
extern __declspec(dllexport) void QCompleter_SetModel(QCompleter* self, QAbstractItemModel* c);
extern __declspec(dllexport) QAbstractItemModel* QCompleter_Model(const QCompleter* self);
extern __declspec(dllexport) void QCompleter_SetCompletionMode(QCompleter* self, CompletionMode mode);
extern __declspec(dllexport) CompletionMode QCompleter_CompletionMode(const QCompleter* self);
extern __declspec(dllexport) void QCompleter_SetFilterMode(QCompleter* self, int filterMode);
extern __declspec(dllexport) int QCompleter_FilterMode(const QCompleter* self);
extern __declspec(dllexport) QAbstractItemView* QCompleter_Popup(const QCompleter* self);
extern __declspec(dllexport) void QCompleter_SetPopup(QCompleter* self, QAbstractItemView* popup);
extern __declspec(dllexport) void QCompleter_SetCaseSensitivity(QCompleter* self, int caseSensitivity);
extern __declspec(dllexport) int QCompleter_CaseSensitivity(const QCompleter* self);
extern __declspec(dllexport) void QCompleter_SetModelSorting(QCompleter* self, ModelSorting sorting);
extern __declspec(dllexport) ModelSorting QCompleter_ModelSorting(const QCompleter* self);
extern __declspec(dllexport) void QCompleter_SetCompletionColumn(QCompleter* self, int column);
extern __declspec(dllexport) int QCompleter_CompletionColumn(const QCompleter* self);
extern __declspec(dllexport) void QCompleter_SetCompletionRole(QCompleter* self, int role);
extern __declspec(dllexport) int QCompleter_CompletionRole(const QCompleter* self);
extern __declspec(dllexport) bool QCompleter_WrapAround(const QCompleter* self);
extern __declspec(dllexport) int QCompleter_MaxVisibleItems(const QCompleter* self);
extern __declspec(dllexport) void QCompleter_SetMaxVisibleItems(QCompleter* self, int maxItems);
extern __declspec(dllexport) int QCompleter_CompletionCount(const QCompleter* self);
extern __declspec(dllexport) bool QCompleter_SetCurrentRow(QCompleter* self, int row);
extern __declspec(dllexport) int QCompleter_CurrentRow(const QCompleter* self);
extern __declspec(dllexport) QModelIndex* QCompleter_CurrentIndex(const QCompleter* self);
extern __declspec(dllexport) struct miqt_string QCompleter_CurrentCompletion(const QCompleter* self);
extern __declspec(dllexport) QAbstractItemModel* QCompleter_CompletionModel(const QCompleter* self);
extern __declspec(dllexport) struct miqt_string QCompleter_CompletionPrefix(const QCompleter* self);
extern __declspec(dllexport) void QCompleter_SetCompletionPrefix(QCompleter* self, struct miqt_string prefix);
extern __declspec(dllexport) void QCompleter_Complete(QCompleter* self);
extern __declspec(dllexport) void QCompleter_SetWrapAround(QCompleter* self, bool wrap);
extern __declspec(dllexport) struct miqt_string QCompleter_PathFromIndex(const QCompleter* self, QModelIndex* index);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QCompleter_SplitPath(const QCompleter* self, struct miqt_string path);
extern __declspec(dllexport) bool QCompleter_EventFilter(QCompleter* self, QObject* o, QEvent* e);
extern __declspec(dllexport) bool QCompleter_Event(QCompleter* self, QEvent* param1);
extern __declspec(dllexport) void QCompleter_Activated(QCompleter* self, struct miqt_string text);
void QCompleter_connect_Activated(QCompleter* self, intptr_t slot);
extern __declspec(dllexport) void QCompleter_ActivatedWithIndex(QCompleter* self, QModelIndex* index);
void QCompleter_connect_ActivatedWithIndex(QCompleter* self, intptr_t slot);
extern __declspec(dllexport) void QCompleter_Highlighted(QCompleter* self, struct miqt_string text);
void QCompleter_connect_Highlighted(QCompleter* self, intptr_t slot);
extern __declspec(dllexport) void QCompleter_HighlightedWithIndex(QCompleter* self, QModelIndex* index);
void QCompleter_connect_HighlightedWithIndex(QCompleter* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QCompleter_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QCompleter_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QCompleter_Complete1(QCompleter* self, QRect* rect);
extern __declspec(dllexport) void QCompleter_override_virtual_PathFromIndex(void* self, intptr_t slot);
struct miqt_string QCompleter_virtualbase_PathFromIndex(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QCompleter_override_virtual_SplitPath(void* self, intptr_t slot);
struct miqt_array /* of struct miqt_string */  QCompleter_virtualbase_SplitPath(const void* self, struct miqt_string path);
extern __declspec(dllexport) void QCompleter_override_virtual_EventFilter(void* self, intptr_t slot);
bool QCompleter_virtualbase_EventFilter(void* self, QObject* o, QEvent* e);
extern __declspec(dllexport) void QCompleter_override_virtual_Event(void* self, intptr_t slot);
bool QCompleter_virtualbase_Event(void* self, QEvent* param1);
extern __declspec(dllexport) void QCompleter_override_virtual_TimerEvent(void* self, intptr_t slot);
void QCompleter_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QCompleter_override_virtual_ChildEvent(void* self, intptr_t slot);
void QCompleter_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QCompleter_override_virtual_CustomEvent(void* self, intptr_t slot);
void QCompleter_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QCompleter_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QCompleter_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QCompleter_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QCompleter_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QCompleter_Delete(QCompleter* self, bool isSubclass);

} 
