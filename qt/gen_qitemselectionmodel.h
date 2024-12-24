#pragma once
#ifndef MIQT_QT_GEN_QITEMSELECTIONMODEL_H
#define MIQT_QT_GEN_QITEMSELECTIONMODEL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QItemSelection QItemSelection;
typedef struct QItemSelectionModel QItemSelectionModel;
typedef struct QItemSelectionRange QItemSelectionRange;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QPersistentModelIndex QPersistentModelIndex;
typedef struct QTimerEvent QTimerEvent;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QItemSelectionRange* QItemSelectionRange_new();
extern __declspec(dllexport) QItemSelectionRange* QItemSelectionRange_new2(QModelIndex* topL, QModelIndex* bottomR);
extern __declspec(dllexport) QItemSelectionRange* QItemSelectionRange_new3(QModelIndex* index);
extern __declspec(dllexport) void QItemSelectionRange_Swap(QItemSelectionRange* self, QItemSelectionRange* other);
extern __declspec(dllexport) int QItemSelectionRange_Top(const QItemSelectionRange* self);
extern __declspec(dllexport) int QItemSelectionRange_Left(const QItemSelectionRange* self);
extern __declspec(dllexport) int QItemSelectionRange_Bottom(const QItemSelectionRange* self);
extern __declspec(dllexport) int QItemSelectionRange_Right(const QItemSelectionRange* self);
extern __declspec(dllexport) int QItemSelectionRange_Width(const QItemSelectionRange* self);
extern __declspec(dllexport) int QItemSelectionRange_Height(const QItemSelectionRange* self);
extern __declspec(dllexport) QPersistentModelIndex* QItemSelectionRange_TopLeft(const QItemSelectionRange* self);
extern __declspec(dllexport) QPersistentModelIndex* QItemSelectionRange_BottomRight(const QItemSelectionRange* self);
extern __declspec(dllexport) QModelIndex* QItemSelectionRange_Parent(const QItemSelectionRange* self);
extern __declspec(dllexport) QAbstractItemModel* QItemSelectionRange_Model(const QItemSelectionRange* self);
extern __declspec(dllexport) bool QItemSelectionRange_Contains(const QItemSelectionRange* self, QModelIndex* index);
extern __declspec(dllexport) bool QItemSelectionRange_Contains2(const QItemSelectionRange* self, int row, int column, QModelIndex* parentIndex);
extern __declspec(dllexport) bool QItemSelectionRange_Intersects(const QItemSelectionRange* self, QItemSelectionRange* other);
extern __declspec(dllexport) QItemSelectionRange* QItemSelectionRange_Intersected(const QItemSelectionRange* self, QItemSelectionRange* other);
extern __declspec(dllexport) bool QItemSelectionRange_IsValid(const QItemSelectionRange* self);
extern __declspec(dllexport) bool QItemSelectionRange_IsEmpty(const QItemSelectionRange* self);
extern __declspec(dllexport) struct miqt_array /* of QModelIndex* */  QItemSelectionRange_Indexes(const QItemSelectionRange* self);
extern __declspec(dllexport) void QItemSelectionRange_Delete(QItemSelectionRange* self, bool isSubclass);

extern __declspec(dllexport) QItemSelectionModel* QItemSelectionModel_new();
extern __declspec(dllexport) QItemSelectionModel* QItemSelectionModel_new2(QAbstractItemModel* model, QObject* parent);
extern __declspec(dllexport) QItemSelectionModel* QItemSelectionModel_new3(QAbstractItemModel* model);
extern __declspec(dllexport) void QItemSelectionModel_virtbase(QItemSelectionModel* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QItemSelectionModel_MetaObject(const QItemSelectionModel* self);
extern __declspec(dllexport) void* QItemSelectionModel_Metacast(QItemSelectionModel* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QItemSelectionModel_Tr(const char* s);
extern __declspec(dllexport) QModelIndex* QItemSelectionModel_CurrentIndex(const QItemSelectionModel* self);
extern __declspec(dllexport) bool QItemSelectionModel_IsSelected(const QItemSelectionModel* self, QModelIndex* index);
extern __declspec(dllexport) bool QItemSelectionModel_IsRowSelected(const QItemSelectionModel* self, int row);
extern __declspec(dllexport) bool QItemSelectionModel_IsColumnSelected(const QItemSelectionModel* self, int column);
extern __declspec(dllexport) bool QItemSelectionModel_RowIntersectsSelection(const QItemSelectionModel* self, int row);
extern __declspec(dllexport) bool QItemSelectionModel_ColumnIntersectsSelection(const QItemSelectionModel* self, int column);
extern __declspec(dllexport) bool QItemSelectionModel_HasSelection(const QItemSelectionModel* self);
extern __declspec(dllexport) struct miqt_array /* of QModelIndex* */  QItemSelectionModel_SelectedIndexes(const QItemSelectionModel* self);
extern __declspec(dllexport) struct miqt_array /* of QModelIndex* */  QItemSelectionModel_SelectedRows(const QItemSelectionModel* self);
extern __declspec(dllexport) struct miqt_array /* of QModelIndex* */  QItemSelectionModel_SelectedColumns(const QItemSelectionModel* self);
extern __declspec(dllexport) QItemSelection* QItemSelectionModel_Selection(const QItemSelectionModel* self);
extern __declspec(dllexport) QAbstractItemModel* QItemSelectionModel_Model(const QItemSelectionModel* self);
extern __declspec(dllexport) QAbstractItemModel* QItemSelectionModel_Model2(QItemSelectionModel* self);
extern __declspec(dllexport) void QItemSelectionModel_SetModel(QItemSelectionModel* self, QAbstractItemModel* model);
extern __declspec(dllexport) void QItemSelectionModel_SetCurrentIndex(QItemSelectionModel* self, QModelIndex* index, int command);
extern __declspec(dllexport) void QItemSelectionModel_Select(QItemSelectionModel* self, QModelIndex* index, int command);
extern __declspec(dllexport) void QItemSelectionModel_Select2(QItemSelectionModel* self, QItemSelection* selection, int command);
extern __declspec(dllexport) void QItemSelectionModel_Clear(QItemSelectionModel* self);
extern __declspec(dllexport) void QItemSelectionModel_Reset(QItemSelectionModel* self);
extern __declspec(dllexport) void QItemSelectionModel_ClearSelection(QItemSelectionModel* self);
extern __declspec(dllexport) void QItemSelectionModel_ClearCurrentIndex(QItemSelectionModel* self);
extern __declspec(dllexport) void QItemSelectionModel_SelectionChanged(QItemSelectionModel* self, QItemSelection* selected, QItemSelection* deselected);
void QItemSelectionModel_connect_SelectionChanged(QItemSelectionModel* self, intptr_t slot);
extern __declspec(dllexport) void QItemSelectionModel_CurrentChanged(QItemSelectionModel* self, QModelIndex* current, QModelIndex* previous);
void QItemSelectionModel_connect_CurrentChanged(QItemSelectionModel* self, intptr_t slot);
extern __declspec(dllexport) void QItemSelectionModel_CurrentRowChanged(QItemSelectionModel* self, QModelIndex* current, QModelIndex* previous);
void QItemSelectionModel_connect_CurrentRowChanged(QItemSelectionModel* self, intptr_t slot);
extern __declspec(dllexport) void QItemSelectionModel_CurrentColumnChanged(QItemSelectionModel* self, QModelIndex* current, QModelIndex* previous);
void QItemSelectionModel_connect_CurrentColumnChanged(QItemSelectionModel* self, intptr_t slot);
extern __declspec(dllexport) void QItemSelectionModel_ModelChanged(QItemSelectionModel* self, QAbstractItemModel* model);
void QItemSelectionModel_connect_ModelChanged(QItemSelectionModel* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QItemSelectionModel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QItemSelectionModel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) bool QItemSelectionModel_IsRowSelected2(const QItemSelectionModel* self, int row, QModelIndex* parent);
extern __declspec(dllexport) bool QItemSelectionModel_IsColumnSelected2(const QItemSelectionModel* self, int column, QModelIndex* parent);
extern __declspec(dllexport) bool QItemSelectionModel_RowIntersectsSelection2(const QItemSelectionModel* self, int row, QModelIndex* parent);
extern __declspec(dllexport) bool QItemSelectionModel_ColumnIntersectsSelection2(const QItemSelectionModel* self, int column, QModelIndex* parent);
extern __declspec(dllexport) struct miqt_array /* of QModelIndex* */  QItemSelectionModel_SelectedRows1(const QItemSelectionModel* self, int column);
extern __declspec(dllexport) struct miqt_array /* of QModelIndex* */  QItemSelectionModel_SelectedColumns1(const QItemSelectionModel* self, int row);
extern __declspec(dllexport) void QItemSelectionModel_override_virtual_SetCurrentIndex(void* self, intptr_t slot);
void QItemSelectionModel_virtualbase_SetCurrentIndex(void* self, QModelIndex* index, int command);
extern __declspec(dllexport) void QItemSelectionModel_override_virtual_Select(void* self, intptr_t slot);
void QItemSelectionModel_virtualbase_Select(void* self, QModelIndex* index, int command);
extern __declspec(dllexport) void QItemSelectionModel_override_virtual_Select2(void* self, intptr_t slot);
void QItemSelectionModel_virtualbase_Select2(void* self, QItemSelection* selection, int command);
extern __declspec(dllexport) void QItemSelectionModel_override_virtual_Clear(void* self, intptr_t slot);
void QItemSelectionModel_virtualbase_Clear(void* self);
extern __declspec(dllexport) void QItemSelectionModel_override_virtual_Reset(void* self, intptr_t slot);
void QItemSelectionModel_virtualbase_Reset(void* self);
extern __declspec(dllexport) void QItemSelectionModel_override_virtual_ClearCurrentIndex(void* self, intptr_t slot);
void QItemSelectionModel_virtualbase_ClearCurrentIndex(void* self);
extern __declspec(dllexport) void QItemSelectionModel_override_virtual_Event(void* self, intptr_t slot);
bool QItemSelectionModel_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QItemSelectionModel_override_virtual_EventFilter(void* self, intptr_t slot);
bool QItemSelectionModel_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QItemSelectionModel_override_virtual_TimerEvent(void* self, intptr_t slot);
void QItemSelectionModel_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QItemSelectionModel_override_virtual_ChildEvent(void* self, intptr_t slot);
void QItemSelectionModel_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QItemSelectionModel_override_virtual_CustomEvent(void* self, intptr_t slot);
void QItemSelectionModel_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QItemSelectionModel_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QItemSelectionModel_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QItemSelectionModel_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QItemSelectionModel_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QItemSelectionModel_Delete(QItemSelectionModel* self, bool isSubclass);

extern __declspec(dllexport) QItemSelection* QItemSelection_new(QModelIndex* topLeft, QModelIndex* bottomRight);
extern __declspec(dllexport) QItemSelection* QItemSelection_new2();
extern __declspec(dllexport) QItemSelection* QItemSelection_new3(QItemSelection* param1);
extern __declspec(dllexport) void QItemSelection_Select(QItemSelection* self, QModelIndex* topLeft, QModelIndex* bottomRight);
extern __declspec(dllexport) bool QItemSelection_Contains(const QItemSelection* self, QModelIndex* index);
extern __declspec(dllexport) struct miqt_array /* of QModelIndex* */  QItemSelection_Indexes(const QItemSelection* self);
extern __declspec(dllexport) void QItemSelection_Merge(QItemSelection* self, QItemSelection* other, int command);
extern __declspec(dllexport) void QItemSelection_Split(QItemSelectionRange* rangeVal, QItemSelectionRange* other, QItemSelection* result);
extern __declspec(dllexport) void QItemSelection_Delete(QItemSelection* self, bool isSubclass);

} 
