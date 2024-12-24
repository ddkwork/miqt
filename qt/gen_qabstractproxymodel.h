#pragma once
#ifndef MIQT_QT_GEN_QABSTRACTPROXYMODEL_H
#define MIQT_QT_GEN_QABSTRACTPROXYMODEL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QAbstractProxyModel QAbstractProxyModel;
typedef struct QItemSelection QItemSelection;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QModelIndex QModelIndex;
typedef struct QModelRoleDataSpan QModelRoleDataSpan;
typedef struct QObject QObject;
typedef struct QSize QSize;
typedef struct QVariant QVariant;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QAbstractProxyModel* QAbstractProxyModel_new();
extern __declspec(dllexport) QAbstractProxyModel* QAbstractProxyModel_new2(QObject* parent);
extern __declspec(dllexport) void QAbstractProxyModel_virtbase(QAbstractProxyModel* src, QAbstractItemModel** outptr_QAbstractItemModel);
extern __declspec(dllexport) QMetaObject* QAbstractProxyModel_MetaObject(const QAbstractProxyModel* self);
extern __declspec(dllexport) void* QAbstractProxyModel_Metacast(QAbstractProxyModel* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QAbstractProxyModel_Tr(const char* s);
extern __declspec(dllexport) void QAbstractProxyModel_SetSourceModel(QAbstractProxyModel* self, QAbstractItemModel* sourceModel);
extern __declspec(dllexport) QAbstractItemModel* QAbstractProxyModel_SourceModel(const QAbstractProxyModel* self);
extern __declspec(dllexport) QModelIndex* QAbstractProxyModel_MapToSource(const QAbstractProxyModel* self, QModelIndex* proxyIndex);
extern __declspec(dllexport) QModelIndex* QAbstractProxyModel_MapFromSource(const QAbstractProxyModel* self, QModelIndex* sourceIndex);
extern __declspec(dllexport) QItemSelection* QAbstractProxyModel_MapSelectionToSource(const QAbstractProxyModel* self, QItemSelection* selection);
extern __declspec(dllexport) QItemSelection* QAbstractProxyModel_MapSelectionFromSource(const QAbstractProxyModel* self, QItemSelection* selection);
extern __declspec(dllexport) bool QAbstractProxyModel_Submit(QAbstractProxyModel* self);
extern __declspec(dllexport) void QAbstractProxyModel_Revert(QAbstractProxyModel* self);
extern __declspec(dllexport) QVariant* QAbstractProxyModel_Data(const QAbstractProxyModel* self, QModelIndex* proxyIndex, int role);
extern __declspec(dllexport) QVariant* QAbstractProxyModel_HeaderData(const QAbstractProxyModel* self, int section, int orientation, int role);
extern __declspec(dllexport) struct miqt_map /* of int to QVariant* */  QAbstractProxyModel_ItemData(const QAbstractProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) int QAbstractProxyModel_Flags(const QAbstractProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) bool QAbstractProxyModel_SetData(QAbstractProxyModel* self, QModelIndex* index, QVariant* value, int role);
extern __declspec(dllexport) bool QAbstractProxyModel_SetItemData(QAbstractProxyModel* self, QModelIndex* index, struct miqt_map /* of int to QVariant* */  roles);
extern __declspec(dllexport) bool QAbstractProxyModel_SetHeaderData(QAbstractProxyModel* self, int section, int orientation, QVariant* value, int role);
extern __declspec(dllexport) bool QAbstractProxyModel_ClearItemData(QAbstractProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) QModelIndex* QAbstractProxyModel_Buddy(const QAbstractProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) bool QAbstractProxyModel_CanFetchMore(const QAbstractProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) void QAbstractProxyModel_FetchMore(QAbstractProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) void QAbstractProxyModel_Sort(QAbstractProxyModel* self, int column, int order);
extern __declspec(dllexport) QSize* QAbstractProxyModel_Span(const QAbstractProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) bool QAbstractProxyModel_HasChildren(const QAbstractProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) QModelIndex* QAbstractProxyModel_Sibling(const QAbstractProxyModel* self, int row, int column, QModelIndex* idx);
extern __declspec(dllexport) QMimeData* QAbstractProxyModel_MimeData(const QAbstractProxyModel* self, struct miqt_array /* of QModelIndex* */  indexes);
extern __declspec(dllexport) bool QAbstractProxyModel_CanDropMimeData(const QAbstractProxyModel* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) bool QAbstractProxyModel_DropMimeData(QAbstractProxyModel* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QAbstractProxyModel_MimeTypes(const QAbstractProxyModel* self);
extern __declspec(dllexport) int QAbstractProxyModel_SupportedDragActions(const QAbstractProxyModel* self);
extern __declspec(dllexport) int QAbstractProxyModel_SupportedDropActions(const QAbstractProxyModel* self);
extern __declspec(dllexport) struct miqt_map /* of int to struct miqt_string */  QAbstractProxyModel_RoleNames(const QAbstractProxyModel* self);
extern __declspec(dllexport) struct miqt_string QAbstractProxyModel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QAbstractProxyModel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_SetSourceModel(void* self, intptr_t slot);
void QAbstractProxyModel_virtualbase_SetSourceModel(void* self, QAbstractItemModel* sourceModel);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_MapToSource(void* self, intptr_t slot);
QModelIndex* QAbstractProxyModel_virtualbase_MapToSource(const void* self, QModelIndex* proxyIndex);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_MapFromSource(void* self, intptr_t slot);
QModelIndex* QAbstractProxyModel_virtualbase_MapFromSource(const void* self, QModelIndex* sourceIndex);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_MapSelectionToSource(void* self, intptr_t slot);
QItemSelection* QAbstractProxyModel_virtualbase_MapSelectionToSource(const void* self, QItemSelection* selection);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_MapSelectionFromSource(void* self, intptr_t slot);
QItemSelection* QAbstractProxyModel_virtualbase_MapSelectionFromSource(const void* self, QItemSelection* selection);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_Submit(void* self, intptr_t slot);
bool QAbstractProxyModel_virtualbase_Submit(void* self);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_Revert(void* self, intptr_t slot);
void QAbstractProxyModel_virtualbase_Revert(void* self);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_Data(void* self, intptr_t slot);
QVariant* QAbstractProxyModel_virtualbase_Data(const void* self, QModelIndex* proxyIndex, int role);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_HeaderData(void* self, intptr_t slot);
QVariant* QAbstractProxyModel_virtualbase_HeaderData(const void* self, int section, int orientation, int role);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_ItemData(void* self, intptr_t slot);
struct miqt_map /* of int to QVariant* */  QAbstractProxyModel_virtualbase_ItemData(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_Flags(void* self, intptr_t slot);
int QAbstractProxyModel_virtualbase_Flags(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_SetData(void* self, intptr_t slot);
bool QAbstractProxyModel_virtualbase_SetData(void* self, QModelIndex* index, QVariant* value, int role);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_SetItemData(void* self, intptr_t slot);
bool QAbstractProxyModel_virtualbase_SetItemData(void* self, QModelIndex* index, struct miqt_map /* of int to QVariant* */  roles);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_SetHeaderData(void* self, intptr_t slot);
bool QAbstractProxyModel_virtualbase_SetHeaderData(void* self, int section, int orientation, QVariant* value, int role);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_ClearItemData(void* self, intptr_t slot);
bool QAbstractProxyModel_virtualbase_ClearItemData(void* self, QModelIndex* index);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_Buddy(void* self, intptr_t slot);
QModelIndex* QAbstractProxyModel_virtualbase_Buddy(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_CanFetchMore(void* self, intptr_t slot);
bool QAbstractProxyModel_virtualbase_CanFetchMore(const void* self, QModelIndex* parent);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_FetchMore(void* self, intptr_t slot);
void QAbstractProxyModel_virtualbase_FetchMore(void* self, QModelIndex* parent);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_Sort(void* self, intptr_t slot);
void QAbstractProxyModel_virtualbase_Sort(void* self, int column, int order);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_Span(void* self, intptr_t slot);
QSize* QAbstractProxyModel_virtualbase_Span(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_HasChildren(void* self, intptr_t slot);
bool QAbstractProxyModel_virtualbase_HasChildren(const void* self, QModelIndex* parent);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_Sibling(void* self, intptr_t slot);
QModelIndex* QAbstractProxyModel_virtualbase_Sibling(const void* self, int row, int column, QModelIndex* idx);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_MimeData(void* self, intptr_t slot);
QMimeData* QAbstractProxyModel_virtualbase_MimeData(const void* self, struct miqt_array /* of QModelIndex* */  indexes);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_CanDropMimeData(void* self, intptr_t slot);
bool QAbstractProxyModel_virtualbase_CanDropMimeData(const void* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_DropMimeData(void* self, intptr_t slot);
bool QAbstractProxyModel_virtualbase_DropMimeData(void* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_MimeTypes(void* self, intptr_t slot);
struct miqt_array /* of struct miqt_string */  QAbstractProxyModel_virtualbase_MimeTypes(const void* self);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_SupportedDragActions(void* self, intptr_t slot);
int QAbstractProxyModel_virtualbase_SupportedDragActions(const void* self);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_SupportedDropActions(void* self, intptr_t slot);
int QAbstractProxyModel_virtualbase_SupportedDropActions(const void* self);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_RoleNames(void* self, intptr_t slot);
struct miqt_map /* of int to struct miqt_string */  QAbstractProxyModel_virtualbase_RoleNames(const void* self);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_Index(void* self, intptr_t slot);
QModelIndex* QAbstractProxyModel_virtualbase_Index(const void* self, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_Parent(void* self, intptr_t slot);
QModelIndex* QAbstractProxyModel_virtualbase_Parent(const void* self, QModelIndex* child);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_RowCount(void* self, intptr_t slot);
int QAbstractProxyModel_virtualbase_RowCount(const void* self, QModelIndex* parent);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_ColumnCount(void* self, intptr_t slot);
int QAbstractProxyModel_virtualbase_ColumnCount(const void* self, QModelIndex* parent);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_InsertRows(void* self, intptr_t slot);
bool QAbstractProxyModel_virtualbase_InsertRows(void* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_InsertColumns(void* self, intptr_t slot);
bool QAbstractProxyModel_virtualbase_InsertColumns(void* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_RemoveRows(void* self, intptr_t slot);
bool QAbstractProxyModel_virtualbase_RemoveRows(void* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_RemoveColumns(void* self, intptr_t slot);
bool QAbstractProxyModel_virtualbase_RemoveColumns(void* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_MoveRows(void* self, intptr_t slot);
bool QAbstractProxyModel_virtualbase_MoveRows(void* self, QModelIndex* sourceParent, int sourceRow, int count, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_MoveColumns(void* self, intptr_t slot);
bool QAbstractProxyModel_virtualbase_MoveColumns(void* self, QModelIndex* sourceParent, int sourceColumn, int count, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_Match(void* self, intptr_t slot);
struct miqt_array /* of QModelIndex* */  QAbstractProxyModel_virtualbase_Match(const void* self, QModelIndex* start, int role, QVariant* value, int hits, int flags);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_MultiData(void* self, intptr_t slot);
void QAbstractProxyModel_virtualbase_MultiData(const void* self, QModelIndex* index, QModelRoleDataSpan* roleDataSpan);
extern __declspec(dllexport) void QAbstractProxyModel_override_virtual_ResetInternalData(void* self, intptr_t slot);
void QAbstractProxyModel_virtualbase_ResetInternalData(void* self);
extern __declspec(dllexport) void QAbstractProxyModel_Delete(QAbstractProxyModel* self, bool isSubclass);

} 
