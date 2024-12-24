#pragma once
#ifndef MIQT_QT_GEN_QCONCATENATETABLESPROXYMODEL_H
#define MIQT_QT_GEN_QCONCATENATETABLESPROXYMODEL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QConcatenateTablesProxyModel QConcatenateTablesProxyModel;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QModelIndex QModelIndex;
typedef struct QModelRoleDataSpan QModelRoleDataSpan;
typedef struct QObject QObject;
typedef struct QSize QSize;
typedef struct QVariant QVariant;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QConcatenateTablesProxyModel* QConcatenateTablesProxyModel_new();
extern __declspec(dllexport) QConcatenateTablesProxyModel* QConcatenateTablesProxyModel_new2(QObject* parent);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_virtbase(QConcatenateTablesProxyModel* src, QAbstractItemModel** outptr_QAbstractItemModel);
extern __declspec(dllexport) QMetaObject* QConcatenateTablesProxyModel_MetaObject(const QConcatenateTablesProxyModel* self);
extern __declspec(dllexport) void* QConcatenateTablesProxyModel_Metacast(QConcatenateTablesProxyModel* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QConcatenateTablesProxyModel_Tr(const char* s);
extern __declspec(dllexport) struct miqt_array /* of QAbstractItemModel* */  QConcatenateTablesProxyModel_SourceModels(const QConcatenateTablesProxyModel* self);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_AddSourceModel(QConcatenateTablesProxyModel* self, QAbstractItemModel* sourceModel);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_RemoveSourceModel(QConcatenateTablesProxyModel* self, QAbstractItemModel* sourceModel);
extern __declspec(dllexport) QModelIndex* QConcatenateTablesProxyModel_MapFromSource(const QConcatenateTablesProxyModel* self, QModelIndex* sourceIndex);
extern __declspec(dllexport) QModelIndex* QConcatenateTablesProxyModel_MapToSource(const QConcatenateTablesProxyModel* self, QModelIndex* proxyIndex);
extern __declspec(dllexport) QVariant* QConcatenateTablesProxyModel_Data(const QConcatenateTablesProxyModel* self, QModelIndex* index, int role);
extern __declspec(dllexport) bool QConcatenateTablesProxyModel_SetData(QConcatenateTablesProxyModel* self, QModelIndex* index, QVariant* value, int role);
extern __declspec(dllexport) struct miqt_map /* of int to QVariant* */  QConcatenateTablesProxyModel_ItemData(const QConcatenateTablesProxyModel* self, QModelIndex* proxyIndex);
extern __declspec(dllexport) bool QConcatenateTablesProxyModel_SetItemData(QConcatenateTablesProxyModel* self, QModelIndex* index, struct miqt_map /* of int to QVariant* */  roles);
extern __declspec(dllexport) int QConcatenateTablesProxyModel_Flags(const QConcatenateTablesProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) QModelIndex* QConcatenateTablesProxyModel_Index(const QConcatenateTablesProxyModel* self, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) QModelIndex* QConcatenateTablesProxyModel_Parent(const QConcatenateTablesProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) int QConcatenateTablesProxyModel_RowCount(const QConcatenateTablesProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) QVariant* QConcatenateTablesProxyModel_HeaderData(const QConcatenateTablesProxyModel* self, int section, int orientation, int role);
extern __declspec(dllexport) int QConcatenateTablesProxyModel_ColumnCount(const QConcatenateTablesProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QConcatenateTablesProxyModel_MimeTypes(const QConcatenateTablesProxyModel* self);
extern __declspec(dllexport) QMimeData* QConcatenateTablesProxyModel_MimeData(const QConcatenateTablesProxyModel* self, struct miqt_array /* of QModelIndex* */  indexes);
extern __declspec(dllexport) bool QConcatenateTablesProxyModel_CanDropMimeData(const QConcatenateTablesProxyModel* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) bool QConcatenateTablesProxyModel_DropMimeData(QConcatenateTablesProxyModel* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) QSize* QConcatenateTablesProxyModel_Span(const QConcatenateTablesProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) struct miqt_string QConcatenateTablesProxyModel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QConcatenateTablesProxyModel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_Data(void* self, intptr_t slot);
QVariant* QConcatenateTablesProxyModel_virtualbase_Data(const void* self, QModelIndex* index, int role);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_SetData(void* self, intptr_t slot);
bool QConcatenateTablesProxyModel_virtualbase_SetData(void* self, QModelIndex* index, QVariant* value, int role);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_ItemData(void* self, intptr_t slot);
struct miqt_map /* of int to QVariant* */  QConcatenateTablesProxyModel_virtualbase_ItemData(const void* self, QModelIndex* proxyIndex);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_SetItemData(void* self, intptr_t slot);
bool QConcatenateTablesProxyModel_virtualbase_SetItemData(void* self, QModelIndex* index, struct miqt_map /* of int to QVariant* */  roles);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_Flags(void* self, intptr_t slot);
int QConcatenateTablesProxyModel_virtualbase_Flags(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_Index(void* self, intptr_t slot);
QModelIndex* QConcatenateTablesProxyModel_virtualbase_Index(const void* self, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_Parent(void* self, intptr_t slot);
QModelIndex* QConcatenateTablesProxyModel_virtualbase_Parent(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_RowCount(void* self, intptr_t slot);
int QConcatenateTablesProxyModel_virtualbase_RowCount(const void* self, QModelIndex* parent);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_HeaderData(void* self, intptr_t slot);
QVariant* QConcatenateTablesProxyModel_virtualbase_HeaderData(const void* self, int section, int orientation, int role);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_ColumnCount(void* self, intptr_t slot);
int QConcatenateTablesProxyModel_virtualbase_ColumnCount(const void* self, QModelIndex* parent);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_MimeTypes(void* self, intptr_t slot);
struct miqt_array /* of struct miqt_string */  QConcatenateTablesProxyModel_virtualbase_MimeTypes(const void* self);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_MimeData(void* self, intptr_t slot);
QMimeData* QConcatenateTablesProxyModel_virtualbase_MimeData(const void* self, struct miqt_array /* of QModelIndex* */  indexes);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_CanDropMimeData(void* self, intptr_t slot);
bool QConcatenateTablesProxyModel_virtualbase_CanDropMimeData(const void* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_DropMimeData(void* self, intptr_t slot);
bool QConcatenateTablesProxyModel_virtualbase_DropMimeData(void* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_Span(void* self, intptr_t slot);
QSize* QConcatenateTablesProxyModel_virtualbase_Span(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_Sibling(void* self, intptr_t slot);
QModelIndex* QConcatenateTablesProxyModel_virtualbase_Sibling(const void* self, int row, int column, QModelIndex* idx);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_HasChildren(void* self, intptr_t slot);
bool QConcatenateTablesProxyModel_virtualbase_HasChildren(const void* self, QModelIndex* parent);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_SetHeaderData(void* self, intptr_t slot);
bool QConcatenateTablesProxyModel_virtualbase_SetHeaderData(void* self, int section, int orientation, QVariant* value, int role);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_ClearItemData(void* self, intptr_t slot);
bool QConcatenateTablesProxyModel_virtualbase_ClearItemData(void* self, QModelIndex* index);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_SupportedDropActions(void* self, intptr_t slot);
int QConcatenateTablesProxyModel_virtualbase_SupportedDropActions(const void* self);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_SupportedDragActions(void* self, intptr_t slot);
int QConcatenateTablesProxyModel_virtualbase_SupportedDragActions(const void* self);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_InsertRows(void* self, intptr_t slot);
bool QConcatenateTablesProxyModel_virtualbase_InsertRows(void* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_InsertColumns(void* self, intptr_t slot);
bool QConcatenateTablesProxyModel_virtualbase_InsertColumns(void* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_RemoveRows(void* self, intptr_t slot);
bool QConcatenateTablesProxyModel_virtualbase_RemoveRows(void* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_RemoveColumns(void* self, intptr_t slot);
bool QConcatenateTablesProxyModel_virtualbase_RemoveColumns(void* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_MoveRows(void* self, intptr_t slot);
bool QConcatenateTablesProxyModel_virtualbase_MoveRows(void* self, QModelIndex* sourceParent, int sourceRow, int count, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_MoveColumns(void* self, intptr_t slot);
bool QConcatenateTablesProxyModel_virtualbase_MoveColumns(void* self, QModelIndex* sourceParent, int sourceColumn, int count, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_FetchMore(void* self, intptr_t slot);
void QConcatenateTablesProxyModel_virtualbase_FetchMore(void* self, QModelIndex* parent);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_CanFetchMore(void* self, intptr_t slot);
bool QConcatenateTablesProxyModel_virtualbase_CanFetchMore(const void* self, QModelIndex* parent);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_Sort(void* self, intptr_t slot);
void QConcatenateTablesProxyModel_virtualbase_Sort(void* self, int column, int order);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_Buddy(void* self, intptr_t slot);
QModelIndex* QConcatenateTablesProxyModel_virtualbase_Buddy(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_Match(void* self, intptr_t slot);
struct miqt_array /* of QModelIndex* */  QConcatenateTablesProxyModel_virtualbase_Match(const void* self, QModelIndex* start, int role, QVariant* value, int hits, int flags);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_RoleNames(void* self, intptr_t slot);
struct miqt_map /* of int to struct miqt_string */  QConcatenateTablesProxyModel_virtualbase_RoleNames(const void* self);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_MultiData(void* self, intptr_t slot);
void QConcatenateTablesProxyModel_virtualbase_MultiData(const void* self, QModelIndex* index, QModelRoleDataSpan* roleDataSpan);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_Submit(void* self, intptr_t slot);
bool QConcatenateTablesProxyModel_virtualbase_Submit(void* self);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_Revert(void* self, intptr_t slot);
void QConcatenateTablesProxyModel_virtualbase_Revert(void* self);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_override_virtual_ResetInternalData(void* self, intptr_t slot);
void QConcatenateTablesProxyModel_virtualbase_ResetInternalData(void* self);
extern __declspec(dllexport) void QConcatenateTablesProxyModel_Delete(QConcatenateTablesProxyModel* self, bool isSubclass);

} 
