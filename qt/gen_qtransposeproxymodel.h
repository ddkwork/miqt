#pragma once
#ifndef MIQT_QT_GEN_QTRANSPOSEPROXYMODEL_H
#define MIQT_QT_GEN_QTRANSPOSEPROXYMODEL_H

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
typedef struct QObject QObject;
typedef struct QSize QSize;
typedef struct QTransposeProxyModel QTransposeProxyModel;
typedef struct QVariant QVariant;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTransposeProxyModel* QTransposeProxyModel_new();
extern __declspec(dllexport) QTransposeProxyModel* QTransposeProxyModel_new2(QObject* parent);
extern __declspec(dllexport) void QTransposeProxyModel_virtbase(QTransposeProxyModel* src, QAbstractProxyModel** outptr_QAbstractProxyModel);
extern __declspec(dllexport) QMetaObject* QTransposeProxyModel_MetaObject(const QTransposeProxyModel* self);
extern __declspec(dllexport) void* QTransposeProxyModel_Metacast(QTransposeProxyModel* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QTransposeProxyModel_Tr(const char* s);
extern __declspec(dllexport) void QTransposeProxyModel_SetSourceModel(QTransposeProxyModel* self, QAbstractItemModel* newSourceModel);
extern __declspec(dllexport) int QTransposeProxyModel_RowCount(const QTransposeProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) int QTransposeProxyModel_ColumnCount(const QTransposeProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) QVariant* QTransposeProxyModel_HeaderData(const QTransposeProxyModel* self, int section, int orientation, int role);
extern __declspec(dllexport) bool QTransposeProxyModel_SetHeaderData(QTransposeProxyModel* self, int section, int orientation, QVariant* value, int role);
extern __declspec(dllexport) bool QTransposeProxyModel_SetItemData(QTransposeProxyModel* self, QModelIndex* index, struct miqt_map /* of int to QVariant* */  roles);
extern __declspec(dllexport) QSize* QTransposeProxyModel_Span(const QTransposeProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) struct miqt_map /* of int to QVariant* */  QTransposeProxyModel_ItemData(const QTransposeProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) QModelIndex* QTransposeProxyModel_MapFromSource(const QTransposeProxyModel* self, QModelIndex* sourceIndex);
extern __declspec(dllexport) QModelIndex* QTransposeProxyModel_MapToSource(const QTransposeProxyModel* self, QModelIndex* proxyIndex);
extern __declspec(dllexport) QModelIndex* QTransposeProxyModel_Parent(const QTransposeProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) QModelIndex* QTransposeProxyModel_Index(const QTransposeProxyModel* self, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) bool QTransposeProxyModel_InsertRows(QTransposeProxyModel* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) bool QTransposeProxyModel_RemoveRows(QTransposeProxyModel* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) bool QTransposeProxyModel_MoveRows(QTransposeProxyModel* self, QModelIndex* sourceParent, int sourceRow, int count, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) bool QTransposeProxyModel_InsertColumns(QTransposeProxyModel* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) bool QTransposeProxyModel_RemoveColumns(QTransposeProxyModel* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) bool QTransposeProxyModel_MoveColumns(QTransposeProxyModel* self, QModelIndex* sourceParent, int sourceColumn, int count, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) void QTransposeProxyModel_Sort(QTransposeProxyModel* self, int column, int order);
extern __declspec(dllexport) struct miqt_string QTransposeProxyModel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QTransposeProxyModel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_SetSourceModel(void* self, intptr_t slot);
void QTransposeProxyModel_virtualbase_SetSourceModel(void* self, QAbstractItemModel* newSourceModel);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_RowCount(void* self, intptr_t slot);
int QTransposeProxyModel_virtualbase_RowCount(const void* self, QModelIndex* parent);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_ColumnCount(void* self, intptr_t slot);
int QTransposeProxyModel_virtualbase_ColumnCount(const void* self, QModelIndex* parent);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_HeaderData(void* self, intptr_t slot);
QVariant* QTransposeProxyModel_virtualbase_HeaderData(const void* self, int section, int orientation, int role);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_SetHeaderData(void* self, intptr_t slot);
bool QTransposeProxyModel_virtualbase_SetHeaderData(void* self, int section, int orientation, QVariant* value, int role);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_SetItemData(void* self, intptr_t slot);
bool QTransposeProxyModel_virtualbase_SetItemData(void* self, QModelIndex* index, struct miqt_map /* of int to QVariant* */  roles);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_Span(void* self, intptr_t slot);
QSize* QTransposeProxyModel_virtualbase_Span(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_ItemData(void* self, intptr_t slot);
struct miqt_map /* of int to QVariant* */  QTransposeProxyModel_virtualbase_ItemData(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_MapFromSource(void* self, intptr_t slot);
QModelIndex* QTransposeProxyModel_virtualbase_MapFromSource(const void* self, QModelIndex* sourceIndex);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_MapToSource(void* self, intptr_t slot);
QModelIndex* QTransposeProxyModel_virtualbase_MapToSource(const void* self, QModelIndex* proxyIndex);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_Parent(void* self, intptr_t slot);
QModelIndex* QTransposeProxyModel_virtualbase_Parent(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_Index(void* self, intptr_t slot);
QModelIndex* QTransposeProxyModel_virtualbase_Index(const void* self, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_InsertRows(void* self, intptr_t slot);
bool QTransposeProxyModel_virtualbase_InsertRows(void* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_RemoveRows(void* self, intptr_t slot);
bool QTransposeProxyModel_virtualbase_RemoveRows(void* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_MoveRows(void* self, intptr_t slot);
bool QTransposeProxyModel_virtualbase_MoveRows(void* self, QModelIndex* sourceParent, int sourceRow, int count, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_InsertColumns(void* self, intptr_t slot);
bool QTransposeProxyModel_virtualbase_InsertColumns(void* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_RemoveColumns(void* self, intptr_t slot);
bool QTransposeProxyModel_virtualbase_RemoveColumns(void* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_MoveColumns(void* self, intptr_t slot);
bool QTransposeProxyModel_virtualbase_MoveColumns(void* self, QModelIndex* sourceParent, int sourceColumn, int count, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_Sort(void* self, intptr_t slot);
void QTransposeProxyModel_virtualbase_Sort(void* self, int column, int order);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_MapSelectionToSource(void* self, intptr_t slot);
QItemSelection* QTransposeProxyModel_virtualbase_MapSelectionToSource(const void* self, QItemSelection* selection);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_MapSelectionFromSource(void* self, intptr_t slot);
QItemSelection* QTransposeProxyModel_virtualbase_MapSelectionFromSource(const void* self, QItemSelection* selection);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_Submit(void* self, intptr_t slot);
bool QTransposeProxyModel_virtualbase_Submit(void* self);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_Revert(void* self, intptr_t slot);
void QTransposeProxyModel_virtualbase_Revert(void* self);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_Data(void* self, intptr_t slot);
QVariant* QTransposeProxyModel_virtualbase_Data(const void* self, QModelIndex* proxyIndex, int role);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_Flags(void* self, intptr_t slot);
int QTransposeProxyModel_virtualbase_Flags(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_SetData(void* self, intptr_t slot);
bool QTransposeProxyModel_virtualbase_SetData(void* self, QModelIndex* index, QVariant* value, int role);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_ClearItemData(void* self, intptr_t slot);
bool QTransposeProxyModel_virtualbase_ClearItemData(void* self, QModelIndex* index);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_Buddy(void* self, intptr_t slot);
QModelIndex* QTransposeProxyModel_virtualbase_Buddy(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_CanFetchMore(void* self, intptr_t slot);
bool QTransposeProxyModel_virtualbase_CanFetchMore(const void* self, QModelIndex* parent);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_FetchMore(void* self, intptr_t slot);
void QTransposeProxyModel_virtualbase_FetchMore(void* self, QModelIndex* parent);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_HasChildren(void* self, intptr_t slot);
bool QTransposeProxyModel_virtualbase_HasChildren(const void* self, QModelIndex* parent);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_Sibling(void* self, intptr_t slot);
QModelIndex* QTransposeProxyModel_virtualbase_Sibling(const void* self, int row, int column, QModelIndex* idx);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_MimeData(void* self, intptr_t slot);
QMimeData* QTransposeProxyModel_virtualbase_MimeData(const void* self, struct miqt_array /* of QModelIndex* */  indexes);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_CanDropMimeData(void* self, intptr_t slot);
bool QTransposeProxyModel_virtualbase_CanDropMimeData(const void* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_DropMimeData(void* self, intptr_t slot);
bool QTransposeProxyModel_virtualbase_DropMimeData(void* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_MimeTypes(void* self, intptr_t slot);
struct miqt_array /* of struct miqt_string */  QTransposeProxyModel_virtualbase_MimeTypes(const void* self);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_SupportedDragActions(void* self, intptr_t slot);
int QTransposeProxyModel_virtualbase_SupportedDragActions(const void* self);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_SupportedDropActions(void* self, intptr_t slot);
int QTransposeProxyModel_virtualbase_SupportedDropActions(const void* self);
extern __declspec(dllexport) void QTransposeProxyModel_override_virtual_RoleNames(void* self, intptr_t slot);
struct miqt_map /* of int to struct miqt_string */  QTransposeProxyModel_virtualbase_RoleNames(const void* self);
extern __declspec(dllexport) void QTransposeProxyModel_Delete(QTransposeProxyModel* self, bool isSubclass);

} 
