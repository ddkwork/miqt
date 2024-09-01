#ifndef GEN_QSORTFILTERPROXYMODEL_H
#define GEN_QSORTFILTERPROXYMODEL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractItemModel;
class QMetaObject;
class QMimeData;
class QModelIndex;
class QObject;
class QRegExp;
class QRegularExpression;
class QSize;
class QSortFilterProxyModel;
class QVariant;
#else
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QRegExp QRegExp;
typedef struct QRegularExpression QRegularExpression;
typedef struct QSize QSize;
typedef struct QSortFilterProxyModel QSortFilterProxyModel;
typedef struct QVariant QVariant;
#endif

QSortFilterProxyModel* QSortFilterProxyModel_new();
QSortFilterProxyModel* QSortFilterProxyModel_new2(QObject* parent);
QMetaObject* QSortFilterProxyModel_MetaObject(QSortFilterProxyModel* self);
void QSortFilterProxyModel_Tr(const char* s, char** _out, int* _out_Strlen);
void QSortFilterProxyModel_TrUtf8(const char* s, char** _out, int* _out_Strlen);
void QSortFilterProxyModel_SetSourceModel(QSortFilterProxyModel* self, QAbstractItemModel* sourceModel);
QModelIndex* QSortFilterProxyModel_MapToSource(QSortFilterProxyModel* self, QModelIndex* proxyIndex);
QModelIndex* QSortFilterProxyModel_MapFromSource(QSortFilterProxyModel* self, QModelIndex* sourceIndex);
QRegExp* QSortFilterProxyModel_FilterRegExp(QSortFilterProxyModel* self);
QRegularExpression* QSortFilterProxyModel_FilterRegularExpression(QSortFilterProxyModel* self);
int QSortFilterProxyModel_FilterKeyColumn(QSortFilterProxyModel* self);
void QSortFilterProxyModel_SetFilterKeyColumn(QSortFilterProxyModel* self, int column);
uintptr_t QSortFilterProxyModel_FilterCaseSensitivity(QSortFilterProxyModel* self);
void QSortFilterProxyModel_SetFilterCaseSensitivity(QSortFilterProxyModel* self, uintptr_t cs);
uintptr_t QSortFilterProxyModel_SortCaseSensitivity(QSortFilterProxyModel* self);
void QSortFilterProxyModel_SetSortCaseSensitivity(QSortFilterProxyModel* self, uintptr_t cs);
bool QSortFilterProxyModel_IsSortLocaleAware(QSortFilterProxyModel* self);
void QSortFilterProxyModel_SetSortLocaleAware(QSortFilterProxyModel* self, bool on);
int QSortFilterProxyModel_SortColumn(QSortFilterProxyModel* self);
uintptr_t QSortFilterProxyModel_SortOrder(QSortFilterProxyModel* self);
bool QSortFilterProxyModel_DynamicSortFilter(QSortFilterProxyModel* self);
void QSortFilterProxyModel_SetDynamicSortFilter(QSortFilterProxyModel* self, bool enable);
int QSortFilterProxyModel_SortRole(QSortFilterProxyModel* self);
void QSortFilterProxyModel_SetSortRole(QSortFilterProxyModel* self, int role);
int QSortFilterProxyModel_FilterRole(QSortFilterProxyModel* self);
void QSortFilterProxyModel_SetFilterRole(QSortFilterProxyModel* self, int role);
bool QSortFilterProxyModel_IsRecursiveFilteringEnabled(QSortFilterProxyModel* self);
void QSortFilterProxyModel_SetRecursiveFilteringEnabled(QSortFilterProxyModel* self, bool recursive);
void QSortFilterProxyModel_SetFilterRegExp(QSortFilterProxyModel* self, const char* pattern, size_t pattern_Strlen);
void QSortFilterProxyModel_SetFilterRegExpWithRegExp(QSortFilterProxyModel* self, QRegExp* regExp);
void QSortFilterProxyModel_SetFilterRegularExpression(QSortFilterProxyModel* self, const char* pattern, size_t pattern_Strlen);
void QSortFilterProxyModel_SetFilterRegularExpressionWithRegularExpression(QSortFilterProxyModel* self, QRegularExpression* regularExpression);
void QSortFilterProxyModel_SetFilterWildcard(QSortFilterProxyModel* self, const char* pattern, size_t pattern_Strlen);
void QSortFilterProxyModel_SetFilterFixedString(QSortFilterProxyModel* self, const char* pattern, size_t pattern_Strlen);
void QSortFilterProxyModel_Clear(QSortFilterProxyModel* self);
void QSortFilterProxyModel_Invalidate(QSortFilterProxyModel* self);
QModelIndex* QSortFilterProxyModel_Index(QSortFilterProxyModel* self, int row, int column);
QModelIndex* QSortFilterProxyModel_Parent(QSortFilterProxyModel* self, QModelIndex* child);
QModelIndex* QSortFilterProxyModel_Sibling(QSortFilterProxyModel* self, int row, int column, QModelIndex* idx);
int QSortFilterProxyModel_RowCount(QSortFilterProxyModel* self);
int QSortFilterProxyModel_ColumnCount(QSortFilterProxyModel* self);
bool QSortFilterProxyModel_HasChildren(QSortFilterProxyModel* self);
QVariant* QSortFilterProxyModel_Data(QSortFilterProxyModel* self, QModelIndex* index);
bool QSortFilterProxyModel_SetData(QSortFilterProxyModel* self, QModelIndex* index, QVariant* value);
QVariant* QSortFilterProxyModel_HeaderData(QSortFilterProxyModel* self, int section, uintptr_t orientation);
bool QSortFilterProxyModel_SetHeaderData(QSortFilterProxyModel* self, int section, uintptr_t orientation, QVariant* value);
QMimeData* QSortFilterProxyModel_MimeData(QSortFilterProxyModel* self, QModelIndex** indexes, size_t indexes_len);
bool QSortFilterProxyModel_DropMimeData(QSortFilterProxyModel* self, QMimeData* data, uintptr_t action, int row, int column, QModelIndex* parent);
bool QSortFilterProxyModel_InsertRows(QSortFilterProxyModel* self, int row, int count);
bool QSortFilterProxyModel_InsertColumns(QSortFilterProxyModel* self, int column, int count);
bool QSortFilterProxyModel_RemoveRows(QSortFilterProxyModel* self, int row, int count);
bool QSortFilterProxyModel_RemoveColumns(QSortFilterProxyModel* self, int column, int count);
void QSortFilterProxyModel_FetchMore(QSortFilterProxyModel* self, QModelIndex* parent);
bool QSortFilterProxyModel_CanFetchMore(QSortFilterProxyModel* self, QModelIndex* parent);
int QSortFilterProxyModel_Flags(QSortFilterProxyModel* self, QModelIndex* index);
QModelIndex* QSortFilterProxyModel_Buddy(QSortFilterProxyModel* self, QModelIndex* index);
void QSortFilterProxyModel_Match(QSortFilterProxyModel* self, QModelIndex* start, int role, QVariant* value, QModelIndex*** _out, size_t* _out_len);
QSize* QSortFilterProxyModel_Span(QSortFilterProxyModel* self, QModelIndex* index);
void QSortFilterProxyModel_Sort(QSortFilterProxyModel* self, int column);
void QSortFilterProxyModel_MimeTypes(QSortFilterProxyModel* self, char*** _out, int** _out_Lengths, size_t* _out_len);
int QSortFilterProxyModel_SupportedDropActions(QSortFilterProxyModel* self);
void QSortFilterProxyModel_DynamicSortFilterChanged(QSortFilterProxyModel* self, bool dynamicSortFilter);
void QSortFilterProxyModel_connect_DynamicSortFilterChanged(QSortFilterProxyModel* self, void* slot);
void QSortFilterProxyModel_FilterCaseSensitivityChanged(QSortFilterProxyModel* self, uintptr_t filterCaseSensitivity);
void QSortFilterProxyModel_connect_FilterCaseSensitivityChanged(QSortFilterProxyModel* self, void* slot);
void QSortFilterProxyModel_SortCaseSensitivityChanged(QSortFilterProxyModel* self, uintptr_t sortCaseSensitivity);
void QSortFilterProxyModel_connect_SortCaseSensitivityChanged(QSortFilterProxyModel* self, void* slot);
void QSortFilterProxyModel_SortLocaleAwareChanged(QSortFilterProxyModel* self, bool sortLocaleAware);
void QSortFilterProxyModel_connect_SortLocaleAwareChanged(QSortFilterProxyModel* self, void* slot);
void QSortFilterProxyModel_SortRoleChanged(QSortFilterProxyModel* self, int sortRole);
void QSortFilterProxyModel_connect_SortRoleChanged(QSortFilterProxyModel* self, void* slot);
void QSortFilterProxyModel_FilterRoleChanged(QSortFilterProxyModel* self, int filterRole);
void QSortFilterProxyModel_connect_FilterRoleChanged(QSortFilterProxyModel* self, void* slot);
void QSortFilterProxyModel_RecursiveFilteringEnabledChanged(QSortFilterProxyModel* self, bool recursiveFilteringEnabled);
void QSortFilterProxyModel_connect_RecursiveFilteringEnabledChanged(QSortFilterProxyModel* self, void* slot);
void QSortFilterProxyModel_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen);
void QSortFilterProxyModel_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QSortFilterProxyModel_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen);
void QSortFilterProxyModel_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
QModelIndex* QSortFilterProxyModel_Index3(QSortFilterProxyModel* self, int row, int column, QModelIndex* parent);
int QSortFilterProxyModel_RowCount1(QSortFilterProxyModel* self, QModelIndex* parent);
int QSortFilterProxyModel_ColumnCount1(QSortFilterProxyModel* self, QModelIndex* parent);
bool QSortFilterProxyModel_HasChildren1(QSortFilterProxyModel* self, QModelIndex* parent);
QVariant* QSortFilterProxyModel_Data2(QSortFilterProxyModel* self, QModelIndex* index, int role);
bool QSortFilterProxyModel_SetData3(QSortFilterProxyModel* self, QModelIndex* index, QVariant* value, int role);
QVariant* QSortFilterProxyModel_HeaderData3(QSortFilterProxyModel* self, int section, uintptr_t orientation, int role);
bool QSortFilterProxyModel_SetHeaderData4(QSortFilterProxyModel* self, int section, uintptr_t orientation, QVariant* value, int role);
bool QSortFilterProxyModel_InsertRows3(QSortFilterProxyModel* self, int row, int count, QModelIndex* parent);
bool QSortFilterProxyModel_InsertColumns3(QSortFilterProxyModel* self, int column, int count, QModelIndex* parent);
bool QSortFilterProxyModel_RemoveRows3(QSortFilterProxyModel* self, int row, int count, QModelIndex* parent);
bool QSortFilterProxyModel_RemoveColumns3(QSortFilterProxyModel* self, int column, int count, QModelIndex* parent);
void QSortFilterProxyModel_Match4(QSortFilterProxyModel* self, QModelIndex* start, int role, QVariant* value, int hits, QModelIndex*** _out, size_t* _out_len);
void QSortFilterProxyModel_Match5(QSortFilterProxyModel* self, QModelIndex* start, int role, QVariant* value, int hits, int flags, QModelIndex*** _out, size_t* _out_len);
void QSortFilterProxyModel_Sort2(QSortFilterProxyModel* self, int column, uintptr_t order);
void QSortFilterProxyModel_Delete(QSortFilterProxyModel* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif