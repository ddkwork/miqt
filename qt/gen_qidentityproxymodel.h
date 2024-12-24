#pragma once
#ifndef MIQT_QT_GEN_QIDENTITYPROXYMODEL_H
#define MIQT_QT_GEN_QIDENTITYPROXYMODEL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QAbstractProxyModel QAbstractProxyModel;
typedef struct QIdentityProxyModel QIdentityProxyModel;
typedef struct QItemSelection QItemSelection;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QIdentityProxyModel* QIdentityProxyModel_new();
extern __declspec(dllexport) 
QIdentityProxyModel* QIdentityProxyModel_new2(QObject* parent);
extern __declspec(dllexport) 
void QIdentityProxyModel_virtbase(QIdentityProxyModel* src
, QAbstractProxyModel** outptr_QAbstractProxyModel
);
extern __declspec(dllexport) 
QMetaObject* QIdentityProxyModel_MetaObject(const QIdentityProxyModel* self);
extern __declspec(dllexport) 
void* QIdentityProxyModel_Metacast(QIdentityProxyModel* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QIdentityProxyModel_Tr(const char* s);
extern __declspec(dllexport) 
int QIdentityProxyModel_ColumnCount(const QIdentityProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
QModelIndex* QIdentityProxyModel_Index(const QIdentityProxyModel* self, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
QModelIndex* QIdentityProxyModel_MapFromSource(const QIdentityProxyModel* self, QModelIndex* sourceIndex);
extern __declspec(dllexport) 
QModelIndex* QIdentityProxyModel_MapToSource(const QIdentityProxyModel* self, QModelIndex* proxyIndex);
extern __declspec(dllexport) 
QModelIndex* QIdentityProxyModel_Parent(const QIdentityProxyModel* self, QModelIndex* child);
extern __declspec(dllexport) 
int QIdentityProxyModel_RowCount(const QIdentityProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
QVariant* QIdentityProxyModel_HeaderData(const QIdentityProxyModel* self, int section, int orientation, int role);
extern __declspec(dllexport) 
bool QIdentityProxyModel_DropMimeData(QIdentityProxyModel* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
QModelIndex* QIdentityProxyModel_Sibling(const QIdentityProxyModel* self, int row, int column, QModelIndex* idx);
extern __declspec(dllexport) 
QItemSelection* QIdentityProxyModel_MapSelectionFromSource(const QIdentityProxyModel* self, QItemSelection* selection);
extern __declspec(dllexport) 
QItemSelection* QIdentityProxyModel_MapSelectionToSource(const QIdentityProxyModel* self, QItemSelection* selection);
extern __declspec(dllexport) 
struct miqt_array /* of QModelIndex* */  QIdentityProxyModel_Match(const QIdentityProxyModel* self, QModelIndex* start, int role, QVariant* value, int hits, int flags);
extern __declspec(dllexport) 
void QIdentityProxyModel_SetSourceModel(QIdentityProxyModel* self, QAbstractItemModel* sourceModel);
extern __declspec(dllexport) 
bool QIdentityProxyModel_InsertColumns(QIdentityProxyModel* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QIdentityProxyModel_InsertRows(QIdentityProxyModel* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QIdentityProxyModel_RemoveColumns(QIdentityProxyModel* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QIdentityProxyModel_RemoveRows(QIdentityProxyModel* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QIdentityProxyModel_MoveRows(QIdentityProxyModel* self, QModelIndex* sourceParent, int sourceRow, int count, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) 
bool QIdentityProxyModel_MoveColumns(QIdentityProxyModel* self, QModelIndex* sourceParent, int sourceColumn, int count, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) 
bool QIdentityProxyModel_HandleSourceLayoutChanges(const QIdentityProxyModel* self);
extern __declspec(dllexport) 
bool QIdentityProxyModel_HandleSourceDataChanges(const QIdentityProxyModel* self);
extern __declspec(dllexport) 
struct miqt_string QIdentityProxyModel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QIdentityProxyModel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QIdentityProxyModel_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QIdentityProxyModel_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QIdentityProxyModel_override_virtual_Metacast(void* self, intptr_t slot);
void* QIdentityProxyModel_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QIdentityProxyModel_Delete(QIdentityProxyModel* self, bool isSubclass);

}
