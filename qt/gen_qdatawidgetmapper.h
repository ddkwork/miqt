#pragma once
#ifndef MIQT_QT_GEN_QDATAWIDGETMAPPER_H
#define MIQT_QT_GEN_QDATAWIDGETMAPPER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemDelegate QAbstractItemDelegate;
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QChildEvent QChildEvent;
typedef struct QDataWidgetMapper QDataWidgetMapper;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QDataWidgetMapper* QDataWidgetMapper_new();
extern __declspec(dllexport) QDataWidgetMapper* QDataWidgetMapper_new2(QObject* parent);
extern __declspec(dllexport) void QDataWidgetMapper_virtbase(QDataWidgetMapper* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QDataWidgetMapper_MetaObject(const QDataWidgetMapper* self);
extern __declspec(dllexport) void* QDataWidgetMapper_Metacast(QDataWidgetMapper* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QDataWidgetMapper_Tr(const char* s);
extern __declspec(dllexport) void QDataWidgetMapper_SetModel(QDataWidgetMapper* self, QAbstractItemModel* model);
extern __declspec(dllexport) QAbstractItemModel* QDataWidgetMapper_Model(const QDataWidgetMapper* self);
extern __declspec(dllexport) void QDataWidgetMapper_SetItemDelegate(QDataWidgetMapper* self, QAbstractItemDelegate* delegate);
extern __declspec(dllexport) QAbstractItemDelegate* QDataWidgetMapper_ItemDelegate(const QDataWidgetMapper* self);
extern __declspec(dllexport) void QDataWidgetMapper_SetRootIndex(QDataWidgetMapper* self, QModelIndex* index);
extern __declspec(dllexport) QModelIndex* QDataWidgetMapper_RootIndex(const QDataWidgetMapper* self);
extern __declspec(dllexport) void QDataWidgetMapper_SetOrientation(QDataWidgetMapper* self, int aOrientation);
extern __declspec(dllexport) int QDataWidgetMapper_Orientation(const QDataWidgetMapper* self);
extern __declspec(dllexport) void QDataWidgetMapper_SetSubmitPolicy(QDataWidgetMapper* self, SubmitPolicy policy);
extern __declspec(dllexport) SubmitPolicy QDataWidgetMapper_SubmitPolicy(const QDataWidgetMapper* self);
extern __declspec(dllexport) void QDataWidgetMapper_AddMapping(QDataWidgetMapper* self, QWidget* widget, int section);
extern __declspec(dllexport) void QDataWidgetMapper_AddMapping2(QDataWidgetMapper* self, QWidget* widget, int section, struct miqt_string propertyName);
extern __declspec(dllexport) void QDataWidgetMapper_RemoveMapping(QDataWidgetMapper* self, QWidget* widget);
extern __declspec(dllexport) int QDataWidgetMapper_MappedSection(const QDataWidgetMapper* self, QWidget* widget);
extern __declspec(dllexport) struct miqt_string QDataWidgetMapper_MappedPropertyName(const QDataWidgetMapper* self, QWidget* widget);
extern __declspec(dllexport) QWidget* QDataWidgetMapper_MappedWidgetAt(const QDataWidgetMapper* self, int section);
extern __declspec(dllexport) void QDataWidgetMapper_ClearMapping(QDataWidgetMapper* self);
extern __declspec(dllexport) int QDataWidgetMapper_CurrentIndex(const QDataWidgetMapper* self);
extern __declspec(dllexport) void QDataWidgetMapper_Revert(QDataWidgetMapper* self);
extern __declspec(dllexport) bool QDataWidgetMapper_Submit(QDataWidgetMapper* self);
extern __declspec(dllexport) void QDataWidgetMapper_ToFirst(QDataWidgetMapper* self);
extern __declspec(dllexport) void QDataWidgetMapper_ToLast(QDataWidgetMapper* self);
extern __declspec(dllexport) void QDataWidgetMapper_ToNext(QDataWidgetMapper* self);
extern __declspec(dllexport) void QDataWidgetMapper_ToPrevious(QDataWidgetMapper* self);
extern __declspec(dllexport) void QDataWidgetMapper_SetCurrentIndex(QDataWidgetMapper* self, int index);
extern __declspec(dllexport) void QDataWidgetMapper_SetCurrentModelIndex(QDataWidgetMapper* self, QModelIndex* index);
extern __declspec(dllexport) void QDataWidgetMapper_CurrentIndexChanged(QDataWidgetMapper* self, int index);
void QDataWidgetMapper_connect_CurrentIndexChanged(QDataWidgetMapper* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QDataWidgetMapper_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QDataWidgetMapper_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QDataWidgetMapper_override_virtual_SetCurrentIndex(void* self, intptr_t slot);
void QDataWidgetMapper_virtualbase_SetCurrentIndex(void* self, int index);
extern __declspec(dllexport) void QDataWidgetMapper_override_virtual_Event(void* self, intptr_t slot);
bool QDataWidgetMapper_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QDataWidgetMapper_override_virtual_EventFilter(void* self, intptr_t slot);
bool QDataWidgetMapper_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QDataWidgetMapper_override_virtual_TimerEvent(void* self, intptr_t slot);
void QDataWidgetMapper_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QDataWidgetMapper_override_virtual_ChildEvent(void* self, intptr_t slot);
void QDataWidgetMapper_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QDataWidgetMapper_override_virtual_CustomEvent(void* self, intptr_t slot);
void QDataWidgetMapper_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QDataWidgetMapper_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QDataWidgetMapper_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QDataWidgetMapper_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QDataWidgetMapper_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QDataWidgetMapper_Delete(QDataWidgetMapper* self, bool isSubclass);

} 
