#pragma once
#ifndef MIQT_QT_GEN_QSORTFILTERPROXYMODEL_H
#define MIQT_QT_GEN_QSORTFILTERPROXYMODEL_H

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
typedef struct QRegularExpression QRegularExpression;
typedef struct QSize QSize;
typedef struct QSortFilterProxyModel QSortFilterProxyModel;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QSortFilterProxyModel* QSortFilterProxyModel_new();
extern __declspec(dllexport) 
QSortFilterProxyModel* QSortFilterProxyModel_new2(QObject* parent);
extern __declspec(dllexport) 
void QSortFilterProxyModel_virtbase(QSortFilterProxyModel* src
, QAbstractProxyModel** outptr_QAbstractProxyModel
);
extern __declspec(dllexport) 
QMetaObject* QSortFilterProxyModel_MetaObject(const QSortFilterProxyModel* self);
extern __declspec(dllexport) 
void* QSortFilterProxyModel_Metacast(QSortFilterProxyModel* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QSortFilterProxyModel_Tr(const char* s);
extern __declspec(dllexport) 
void QSortFilterProxyModel_SetSourceModel(QSortFilterProxyModel* self, QAbstractItemModel* sourceModel);
extern __declspec(dllexport) 
QModelIndex* QSortFilterProxyModel_MapToSource(const QSortFilterProxyModel* self, QModelIndex* proxyIndex);
extern __declspec(dllexport) 
QModelIndex* QSortFilterProxyModel_MapFromSource(const QSortFilterProxyModel* self, QModelIndex* sourceIndex);
extern __declspec(dllexport) 
QItemSelection* QSortFilterProxyModel_MapSelectionToSource(const QSortFilterProxyModel* self, QItemSelection* proxySelection);
extern __declspec(dllexport) 
QItemSelection* QSortFilterProxyModel_MapSelectionFromSource(const QSortFilterProxyModel* self, QItemSelection* sourceSelection);
extern __declspec(dllexport) 
QRegularExpression* QSortFilterProxyModel_FilterRegularExpression(const QSortFilterProxyModel* self);
extern __declspec(dllexport) 
int QSortFilterProxyModel_FilterKeyColumn(const QSortFilterProxyModel* self);
extern __declspec(dllexport) 
void QSortFilterProxyModel_SetFilterKeyColumn(QSortFilterProxyModel* self, int column);
extern __declspec(dllexport) 
int QSortFilterProxyModel_FilterCaseSensitivity(const QSortFilterProxyModel* self);
extern __declspec(dllexport) 
void QSortFilterProxyModel_SetFilterCaseSensitivity(QSortFilterProxyModel* self, int cs);
extern __declspec(dllexport) 
int QSortFilterProxyModel_SortCaseSensitivity(const QSortFilterProxyModel* self);
extern __declspec(dllexport) 
void QSortFilterProxyModel_SetSortCaseSensitivity(QSortFilterProxyModel* self, int cs);
extern __declspec(dllexport) 
bool QSortFilterProxyModel_IsSortLocaleAware(const QSortFilterProxyModel* self);
extern __declspec(dllexport) 
void QSortFilterProxyModel_SetSortLocaleAware(QSortFilterProxyModel* self, bool on);
extern __declspec(dllexport) 
int QSortFilterProxyModel_SortColumn(const QSortFilterProxyModel* self);
extern __declspec(dllexport) 
int QSortFilterProxyModel_SortOrder(const QSortFilterProxyModel* self);
extern __declspec(dllexport) 
bool QSortFilterProxyModel_DynamicSortFilter(const QSortFilterProxyModel* self);
extern __declspec(dllexport) 
void QSortFilterProxyModel_SetDynamicSortFilter(QSortFilterProxyModel* self, bool enable);
extern __declspec(dllexport) 
int QSortFilterProxyModel_SortRole(const QSortFilterProxyModel* self);
extern __declspec(dllexport) 
void QSortFilterProxyModel_SetSortRole(QSortFilterProxyModel* self, int role);
extern __declspec(dllexport) 
int QSortFilterProxyModel_FilterRole(const QSortFilterProxyModel* self);
extern __declspec(dllexport) 
void QSortFilterProxyModel_SetFilterRole(QSortFilterProxyModel* self, int role);
extern __declspec(dllexport) 
bool QSortFilterProxyModel_IsRecursiveFilteringEnabled(const QSortFilterProxyModel* self);
extern __declspec(dllexport) 
void QSortFilterProxyModel_SetRecursiveFilteringEnabled(QSortFilterProxyModel* self, bool recursive);
extern __declspec(dllexport) 
bool QSortFilterProxyModel_AutoAcceptChildRows(const QSortFilterProxyModel* self);
extern __declspec(dllexport) 
void QSortFilterProxyModel_SetAutoAcceptChildRows(QSortFilterProxyModel* self, bool accept);
extern __declspec(dllexport) 
void QSortFilterProxyModel_SetFilterRegularExpression(QSortFilterProxyModel* self, struct miqt_string pattern);
extern __declspec(dllexport) 
void QSortFilterProxyModel_SetFilterRegularExpressionWithRegularExpression(QSortFilterProxyModel* self, QRegularExpression* regularExpression);
extern __declspec(dllexport) 
void QSortFilterProxyModel_SetFilterWildcard(QSortFilterProxyModel* self, struct miqt_string pattern);
extern __declspec(dllexport) 
void QSortFilterProxyModel_SetFilterFixedString(QSortFilterProxyModel* self, struct miqt_string pattern);
extern __declspec(dllexport) 
void QSortFilterProxyModel_Invalidate(QSortFilterProxyModel* self);
extern __declspec(dllexport) 
bool QSortFilterProxyModel_FilterAcceptsRow(const QSortFilterProxyModel* self, int source_row, QModelIndex* source_parent);
extern __declspec(dllexport) 
bool QSortFilterProxyModel_FilterAcceptsColumn(const QSortFilterProxyModel* self, int source_column, QModelIndex* source_parent);
extern __declspec(dllexport) 
bool QSortFilterProxyModel_LessThan(const QSortFilterProxyModel* self, QModelIndex* source_left, QModelIndex* source_right);
extern __declspec(dllexport) 
QModelIndex* QSortFilterProxyModel_Index(const QSortFilterProxyModel* self, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
QModelIndex* QSortFilterProxyModel_Parent(const QSortFilterProxyModel* self, QModelIndex* child);
extern __declspec(dllexport) 
QModelIndex* QSortFilterProxyModel_Sibling(const QSortFilterProxyModel* self, int row, int column, QModelIndex* idx);
extern __declspec(dllexport) 
int QSortFilterProxyModel_RowCount(const QSortFilterProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
int QSortFilterProxyModel_ColumnCount(const QSortFilterProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
bool QSortFilterProxyModel_HasChildren(const QSortFilterProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
QVariant* QSortFilterProxyModel_Data(const QSortFilterProxyModel* self, QModelIndex* index, int role);
extern __declspec(dllexport) 
bool QSortFilterProxyModel_SetData(QSortFilterProxyModel* self, QModelIndex* index, QVariant* value, int role);
extern __declspec(dllexport) 
QVariant* QSortFilterProxyModel_HeaderData(const QSortFilterProxyModel* self, int section, int orientation, int role);
extern __declspec(dllexport) 
bool QSortFilterProxyModel_SetHeaderData(QSortFilterProxyModel* self, int section, int orientation, QVariant* value, int role);
extern __declspec(dllexport) 
QMimeData* QSortFilterProxyModel_MimeData(const QSortFilterProxyModel* self, struct miqt_array /* of QModelIndex* */  indexes);
extern __declspec(dllexport) 
bool QSortFilterProxyModel_DropMimeData(QSortFilterProxyModel* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
bool QSortFilterProxyModel_InsertRows(QSortFilterProxyModel* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QSortFilterProxyModel_InsertColumns(QSortFilterProxyModel* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QSortFilterProxyModel_RemoveRows(QSortFilterProxyModel* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QSortFilterProxyModel_RemoveColumns(QSortFilterProxyModel* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) 
void QSortFilterProxyModel_FetchMore(QSortFilterProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
bool QSortFilterProxyModel_CanFetchMore(const QSortFilterProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
int QSortFilterProxyModel_Flags(const QSortFilterProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) 
QModelIndex* QSortFilterProxyModel_Buddy(const QSortFilterProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) 
struct miqt_array /* of QModelIndex* */  QSortFilterProxyModel_Match(const QSortFilterProxyModel* self, QModelIndex* start, int role, QVariant* value, int hits, int flags);
extern __declspec(dllexport) 
QSize* QSortFilterProxyModel_Span(const QSortFilterProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) 
void QSortFilterProxyModel_Sort(QSortFilterProxyModel* self, int column, int order);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QSortFilterProxyModel_MimeTypes(const QSortFilterProxyModel* self);
extern __declspec(dllexport) 
int QSortFilterProxyModel_SupportedDropActions(const QSortFilterProxyModel* self);
extern __declspec(dllexport) 
void QSortFilterProxyModel_DynamicSortFilterChanged(QSortFilterProxyModel* self, bool dynamicSortFilter);
void QSortFilterProxyModel_connect_DynamicSortFilterChanged(QSortFilterProxyModel* self, intptr_t slot);
extern __declspec(dllexport) 
void QSortFilterProxyModel_FilterCaseSensitivityChanged(QSortFilterProxyModel* self, int filterCaseSensitivity);
void QSortFilterProxyModel_connect_FilterCaseSensitivityChanged(QSortFilterProxyModel* self, intptr_t slot);
extern __declspec(dllexport) 
void QSortFilterProxyModel_SortCaseSensitivityChanged(QSortFilterProxyModel* self, int sortCaseSensitivity);
void QSortFilterProxyModel_connect_SortCaseSensitivityChanged(QSortFilterProxyModel* self, intptr_t slot);
extern __declspec(dllexport) 
void QSortFilterProxyModel_SortLocaleAwareChanged(QSortFilterProxyModel* self, bool sortLocaleAware);
void QSortFilterProxyModel_connect_SortLocaleAwareChanged(QSortFilterProxyModel* self, intptr_t slot);
extern __declspec(dllexport) 
void QSortFilterProxyModel_SortRoleChanged(QSortFilterProxyModel* self, int sortRole);
void QSortFilterProxyModel_connect_SortRoleChanged(QSortFilterProxyModel* self, intptr_t slot);
extern __declspec(dllexport) 
void QSortFilterProxyModel_FilterRoleChanged(QSortFilterProxyModel* self, int filterRole);
void QSortFilterProxyModel_connect_FilterRoleChanged(QSortFilterProxyModel* self, intptr_t slot);
extern __declspec(dllexport) 
void QSortFilterProxyModel_RecursiveFilteringEnabledChanged(QSortFilterProxyModel* self, bool recursiveFilteringEnabled);
void QSortFilterProxyModel_connect_RecursiveFilteringEnabledChanged(QSortFilterProxyModel* self, intptr_t slot);
extern __declspec(dllexport) 
void QSortFilterProxyModel_AutoAcceptChildRowsChanged(QSortFilterProxyModel* self, bool autoAcceptChildRows);
void QSortFilterProxyModel_connect_AutoAcceptChildRowsChanged(QSortFilterProxyModel* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QSortFilterProxyModel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QSortFilterProxyModel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QSortFilterProxyModel_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QSortFilterProxyModel_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QSortFilterProxyModel_override_virtual_Metacast(void* self, intptr_t slot);
void* QSortFilterProxyModel_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QSortFilterProxyModel_Delete(QSortFilterProxyModel* self, bool isSubclass);

}
