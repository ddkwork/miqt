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
typedef struct QObject QObject;
typedef struct QSize QSize;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QConcatenateTablesProxyModel* QConcatenateTablesProxyModel_new();
extern __declspec(dllexport) 
QConcatenateTablesProxyModel* QConcatenateTablesProxyModel_new2(QObject* parent);
extern __declspec(dllexport) 
void QConcatenateTablesProxyModel_virtbase(QConcatenateTablesProxyModel* src
, QAbstractItemModel** outptr_QAbstractItemModel
);
extern __declspec(dllexport) 
QMetaObject* QConcatenateTablesProxyModel_MetaObject(const QConcatenateTablesProxyModel* self);
extern __declspec(dllexport) 
void* QConcatenateTablesProxyModel_Metacast(QConcatenateTablesProxyModel* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QConcatenateTablesProxyModel_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_array /* of QAbstractItemModel* */  QConcatenateTablesProxyModel_SourceModels(const QConcatenateTablesProxyModel* self);
extern __declspec(dllexport) 
void QConcatenateTablesProxyModel_AddSourceModel(QConcatenateTablesProxyModel* self, QAbstractItemModel* sourceModel);
extern __declspec(dllexport) 
void QConcatenateTablesProxyModel_RemoveSourceModel(QConcatenateTablesProxyModel* self, QAbstractItemModel* sourceModel);
extern __declspec(dllexport) 
QModelIndex* QConcatenateTablesProxyModel_MapFromSource(const QConcatenateTablesProxyModel* self, QModelIndex* sourceIndex);
extern __declspec(dllexport) 
QModelIndex* QConcatenateTablesProxyModel_MapToSource(const QConcatenateTablesProxyModel* self, QModelIndex* proxyIndex);
extern __declspec(dllexport) 
QVariant* QConcatenateTablesProxyModel_Data(const QConcatenateTablesProxyModel* self, QModelIndex* index, int role);
extern __declspec(dllexport) 
bool QConcatenateTablesProxyModel_SetData(QConcatenateTablesProxyModel* self, QModelIndex* index, QVariant* value, int role);
extern __declspec(dllexport) 
struct miqt_map /* of int to QVariant* */  QConcatenateTablesProxyModel_ItemData(const QConcatenateTablesProxyModel* self, QModelIndex* proxyIndex);
extern __declspec(dllexport) 
bool QConcatenateTablesProxyModel_SetItemData(QConcatenateTablesProxyModel* self, QModelIndex* index, struct miqt_map /* of int to QVariant* */  roles);
extern __declspec(dllexport) 
int QConcatenateTablesProxyModel_Flags(const QConcatenateTablesProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) 
QModelIndex* QConcatenateTablesProxyModel_Index(const QConcatenateTablesProxyModel* self, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
QModelIndex* QConcatenateTablesProxyModel_Parent(const QConcatenateTablesProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) 
int QConcatenateTablesProxyModel_RowCount(const QConcatenateTablesProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
QVariant* QConcatenateTablesProxyModel_HeaderData(const QConcatenateTablesProxyModel* self, int section, int orientation, int role);
extern __declspec(dllexport) 
int QConcatenateTablesProxyModel_ColumnCount(const QConcatenateTablesProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QConcatenateTablesProxyModel_MimeTypes(const QConcatenateTablesProxyModel* self);
extern __declspec(dllexport) 
QMimeData* QConcatenateTablesProxyModel_MimeData(const QConcatenateTablesProxyModel* self, struct miqt_array /* of QModelIndex* */  indexes);
extern __declspec(dllexport) 
bool QConcatenateTablesProxyModel_CanDropMimeData(const QConcatenateTablesProxyModel* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
bool QConcatenateTablesProxyModel_DropMimeData(QConcatenateTablesProxyModel* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
QSize* QConcatenateTablesProxyModel_Span(const QConcatenateTablesProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) 
struct miqt_string QConcatenateTablesProxyModel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QConcatenateTablesProxyModel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QConcatenateTablesProxyModel_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QConcatenateTablesProxyModel_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QConcatenateTablesProxyModel_override_virtual_Metacast(void* self, intptr_t slot);
void* QConcatenateTablesProxyModel_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QConcatenateTablesProxyModel_Delete(QConcatenateTablesProxyModel* self, bool isSubclass);

}
