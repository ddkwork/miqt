#ifndef GEN_QCONCATENATETABLESPROXYMODEL_H
#define GEN_QCONCATENATETABLESPROXYMODEL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractItemModel;
class QConcatenateTablesProxyModel;
class QMetaObject;
class QMimeData;
class QModelIndex;
class QObject;
class QSize;
class QVariant;
#else
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QConcatenateTablesProxyModel QConcatenateTablesProxyModel;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QSize QSize;
typedef struct QVariant QVariant;
#endif

QConcatenateTablesProxyModel* QConcatenateTablesProxyModel_new();
QConcatenateTablesProxyModel* QConcatenateTablesProxyModel_new2(QObject* parent);
QMetaObject* QConcatenateTablesProxyModel_MetaObject(QConcatenateTablesProxyModel* self);
void QConcatenateTablesProxyModel_Tr(const char* s, char** _out, int* _out_Strlen);
void QConcatenateTablesProxyModel_TrUtf8(const char* s, char** _out, int* _out_Strlen);
void QConcatenateTablesProxyModel_SourceModels(QConcatenateTablesProxyModel* self, QAbstractItemModel*** _out, size_t* _out_len);
void QConcatenateTablesProxyModel_AddSourceModel(QConcatenateTablesProxyModel* self, QAbstractItemModel* sourceModel);
void QConcatenateTablesProxyModel_RemoveSourceModel(QConcatenateTablesProxyModel* self, QAbstractItemModel* sourceModel);
QModelIndex* QConcatenateTablesProxyModel_MapFromSource(QConcatenateTablesProxyModel* self, QModelIndex* sourceIndex);
QModelIndex* QConcatenateTablesProxyModel_MapToSource(QConcatenateTablesProxyModel* self, QModelIndex* proxyIndex);
QVariant* QConcatenateTablesProxyModel_Data(QConcatenateTablesProxyModel* self, QModelIndex* index);
bool QConcatenateTablesProxyModel_SetData(QConcatenateTablesProxyModel* self, QModelIndex* index, QVariant* value);
int QConcatenateTablesProxyModel_Flags(QConcatenateTablesProxyModel* self, QModelIndex* index);
QModelIndex* QConcatenateTablesProxyModel_Index(QConcatenateTablesProxyModel* self, int row, int column);
QModelIndex* QConcatenateTablesProxyModel_Parent(QConcatenateTablesProxyModel* self, QModelIndex* index);
int QConcatenateTablesProxyModel_RowCount(QConcatenateTablesProxyModel* self);
QVariant* QConcatenateTablesProxyModel_HeaderData(QConcatenateTablesProxyModel* self, int section, uintptr_t orientation);
int QConcatenateTablesProxyModel_ColumnCount(QConcatenateTablesProxyModel* self);
void QConcatenateTablesProxyModel_MimeTypes(QConcatenateTablesProxyModel* self, char*** _out, int** _out_Lengths, size_t* _out_len);
QMimeData* QConcatenateTablesProxyModel_MimeData(QConcatenateTablesProxyModel* self, QModelIndex** indexes, size_t indexes_len);
bool QConcatenateTablesProxyModel_CanDropMimeData(QConcatenateTablesProxyModel* self, QMimeData* data, uintptr_t action, int row, int column, QModelIndex* parent);
bool QConcatenateTablesProxyModel_DropMimeData(QConcatenateTablesProxyModel* self, QMimeData* data, uintptr_t action, int row, int column, QModelIndex* parent);
QSize* QConcatenateTablesProxyModel_Span(QConcatenateTablesProxyModel* self, QModelIndex* index);
void QConcatenateTablesProxyModel_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen);
void QConcatenateTablesProxyModel_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QConcatenateTablesProxyModel_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen);
void QConcatenateTablesProxyModel_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
QVariant* QConcatenateTablesProxyModel_Data2(QConcatenateTablesProxyModel* self, QModelIndex* index, int role);
bool QConcatenateTablesProxyModel_SetData3(QConcatenateTablesProxyModel* self, QModelIndex* index, QVariant* value, int role);
QModelIndex* QConcatenateTablesProxyModel_Index3(QConcatenateTablesProxyModel* self, int row, int column, QModelIndex* parent);
int QConcatenateTablesProxyModel_RowCount1(QConcatenateTablesProxyModel* self, QModelIndex* parent);
QVariant* QConcatenateTablesProxyModel_HeaderData3(QConcatenateTablesProxyModel* self, int section, uintptr_t orientation, int role);
int QConcatenateTablesProxyModel_ColumnCount1(QConcatenateTablesProxyModel* self, QModelIndex* parent);
void QConcatenateTablesProxyModel_Delete(QConcatenateTablesProxyModel* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif