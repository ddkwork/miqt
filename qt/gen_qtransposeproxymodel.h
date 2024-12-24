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
typedef struct QMetaObject QMetaObject;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QSize QSize;
typedef struct QTransposeProxyModel QTransposeProxyModel;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QTransposeProxyModel* QTransposeProxyModel_new();
extern __declspec(dllexport) 
QTransposeProxyModel* QTransposeProxyModel_new2(QObject* parent);
extern __declspec(dllexport) 
void QTransposeProxyModel_virtbase(QTransposeProxyModel* src
, QAbstractProxyModel** outptr_QAbstractProxyModel
);
extern __declspec(dllexport) 
QMetaObject* QTransposeProxyModel_MetaObject(const QTransposeProxyModel* self);
extern __declspec(dllexport) 
void* QTransposeProxyModel_Metacast(QTransposeProxyModel* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QTransposeProxyModel_Tr(const char* s);
extern __declspec(dllexport) 
void QTransposeProxyModel_SetSourceModel(QTransposeProxyModel* self, QAbstractItemModel* newSourceModel);
extern __declspec(dllexport) 
int QTransposeProxyModel_RowCount(const QTransposeProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
int QTransposeProxyModel_ColumnCount(const QTransposeProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
QVariant* QTransposeProxyModel_HeaderData(const QTransposeProxyModel* self, int section, int orientation, int role);
extern __declspec(dllexport) 
bool QTransposeProxyModel_SetHeaderData(QTransposeProxyModel* self, int section, int orientation, QVariant* value, int role);
extern __declspec(dllexport) 
bool QTransposeProxyModel_SetItemData(QTransposeProxyModel* self, QModelIndex* index, struct miqt_map /* of int to QVariant* */  roles);
extern __declspec(dllexport) 
QSize* QTransposeProxyModel_Span(const QTransposeProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) 
struct miqt_map /* of int to QVariant* */  QTransposeProxyModel_ItemData(const QTransposeProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) 
QModelIndex* QTransposeProxyModel_MapFromSource(const QTransposeProxyModel* self, QModelIndex* sourceIndex);
extern __declspec(dllexport) 
QModelIndex* QTransposeProxyModel_MapToSource(const QTransposeProxyModel* self, QModelIndex* proxyIndex);
extern __declspec(dllexport) 
QModelIndex* QTransposeProxyModel_Parent(const QTransposeProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) 
QModelIndex* QTransposeProxyModel_Index(const QTransposeProxyModel* self, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
bool QTransposeProxyModel_InsertRows(QTransposeProxyModel* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QTransposeProxyModel_RemoveRows(QTransposeProxyModel* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QTransposeProxyModel_MoveRows(QTransposeProxyModel* self, QModelIndex* sourceParent, int sourceRow, int count, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) 
bool QTransposeProxyModel_InsertColumns(QTransposeProxyModel* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QTransposeProxyModel_RemoveColumns(QTransposeProxyModel* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QTransposeProxyModel_MoveColumns(QTransposeProxyModel* self, QModelIndex* sourceParent, int sourceColumn, int count, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) 
void QTransposeProxyModel_Sort(QTransposeProxyModel* self, int column, int order);
extern __declspec(dllexport) 
struct miqt_string QTransposeProxyModel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QTransposeProxyModel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QTransposeProxyModel_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QTransposeProxyModel_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QTransposeProxyModel_override_virtual_Metacast(void* self, intptr_t slot);
void* QTransposeProxyModel_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QTransposeProxyModel_Delete(QTransposeProxyModel* self, bool isSubclass);

}
