#ifndef GEN_QABSTRACTPROXYMODEL_H
#define GEN_QABSTRACTPROXYMODEL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractItemModel;
class QAbstractProxyModel;
class QMetaObject;
class QMimeData;
class QModelIndex;
class QSize;
class QVariant;
#else
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QAbstractProxyModel QAbstractProxyModel;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QModelIndex QModelIndex;
typedef struct QSize QSize;
typedef struct QVariant QVariant;
#endif

QMetaObject* QAbstractProxyModel_MetaObject(QAbstractProxyModel* self);
void QAbstractProxyModel_Tr(const char* s, char** _out, int* _out_Strlen);
void QAbstractProxyModel_TrUtf8(const char* s, char** _out, int* _out_Strlen);
void QAbstractProxyModel_SetSourceModel(QAbstractProxyModel* self, QAbstractItemModel* sourceModel);
QAbstractItemModel* QAbstractProxyModel_SourceModel(QAbstractProxyModel* self);
QModelIndex* QAbstractProxyModel_MapToSource(QAbstractProxyModel* self, QModelIndex* proxyIndex);
QModelIndex* QAbstractProxyModel_MapFromSource(QAbstractProxyModel* self, QModelIndex* sourceIndex);
bool QAbstractProxyModel_Submit(QAbstractProxyModel* self);
void QAbstractProxyModel_Revert(QAbstractProxyModel* self);
QVariant* QAbstractProxyModel_Data(QAbstractProxyModel* self, QModelIndex* proxyIndex);
QVariant* QAbstractProxyModel_HeaderData(QAbstractProxyModel* self, int section, uintptr_t orientation);
int QAbstractProxyModel_Flags(QAbstractProxyModel* self, QModelIndex* index);
bool QAbstractProxyModel_SetData(QAbstractProxyModel* self, QModelIndex* index, QVariant* value);
bool QAbstractProxyModel_SetHeaderData(QAbstractProxyModel* self, int section, uintptr_t orientation, QVariant* value);
QModelIndex* QAbstractProxyModel_Buddy(QAbstractProxyModel* self, QModelIndex* index);
bool QAbstractProxyModel_CanFetchMore(QAbstractProxyModel* self, QModelIndex* parent);
void QAbstractProxyModel_FetchMore(QAbstractProxyModel* self, QModelIndex* parent);
void QAbstractProxyModel_Sort(QAbstractProxyModel* self, int column);
QSize* QAbstractProxyModel_Span(QAbstractProxyModel* self, QModelIndex* index);
bool QAbstractProxyModel_HasChildren(QAbstractProxyModel* self);
QModelIndex* QAbstractProxyModel_Sibling(QAbstractProxyModel* self, int row, int column, QModelIndex* idx);
QMimeData* QAbstractProxyModel_MimeData(QAbstractProxyModel* self, QModelIndex** indexes, size_t indexes_len);
bool QAbstractProxyModel_CanDropMimeData(QAbstractProxyModel* self, QMimeData* data, uintptr_t action, int row, int column, QModelIndex* parent);
bool QAbstractProxyModel_DropMimeData(QAbstractProxyModel* self, QMimeData* data, uintptr_t action, int row, int column, QModelIndex* parent);
void QAbstractProxyModel_MimeTypes(QAbstractProxyModel* self, char*** _out, int** _out_Lengths, size_t* _out_len);
int QAbstractProxyModel_SupportedDragActions(QAbstractProxyModel* self);
int QAbstractProxyModel_SupportedDropActions(QAbstractProxyModel* self);
void QAbstractProxyModel_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen);
void QAbstractProxyModel_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QAbstractProxyModel_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen);
void QAbstractProxyModel_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
QVariant* QAbstractProxyModel_Data2(QAbstractProxyModel* self, QModelIndex* proxyIndex, int role);
QVariant* QAbstractProxyModel_HeaderData3(QAbstractProxyModel* self, int section, uintptr_t orientation, int role);
bool QAbstractProxyModel_SetData3(QAbstractProxyModel* self, QModelIndex* index, QVariant* value, int role);
bool QAbstractProxyModel_SetHeaderData4(QAbstractProxyModel* self, int section, uintptr_t orientation, QVariant* value, int role);
void QAbstractProxyModel_Sort2(QAbstractProxyModel* self, int column, uintptr_t order);
bool QAbstractProxyModel_HasChildren1(QAbstractProxyModel* self, QModelIndex* parent);
void QAbstractProxyModel_Delete(QAbstractProxyModel* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif