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
typedef struct QObject QObject;
typedef struct QSize QSize;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAbstractProxyModel* QAbstractProxyModel_new();
extern __declspec(dllexport) 
QAbstractProxyModel* QAbstractProxyModel_new2(QObject* parent);
extern __declspec(dllexport) 
void QAbstractProxyModel_virtbase(QAbstractProxyModel* src
, QAbstractItemModel** outptr_QAbstractItemModel
);
extern __declspec(dllexport) 
QMetaObject* QAbstractProxyModel_MetaObject(const QAbstractProxyModel* self);
extern __declspec(dllexport) 
void* QAbstractProxyModel_Metacast(QAbstractProxyModel* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAbstractProxyModel_Tr(const char* s);
extern __declspec(dllexport) 
void QAbstractProxyModel_SetSourceModel(QAbstractProxyModel* self, QAbstractItemModel* sourceModel);
extern __declspec(dllexport) 
QAbstractItemModel* QAbstractProxyModel_SourceModel(const QAbstractProxyModel* self);
extern __declspec(dllexport) 
QModelIndex* QAbstractProxyModel_MapToSource(const QAbstractProxyModel* self, QModelIndex* proxyIndex);
extern __declspec(dllexport) 
QModelIndex* QAbstractProxyModel_MapFromSource(const QAbstractProxyModel* self, QModelIndex* sourceIndex);
extern __declspec(dllexport) 
QItemSelection* QAbstractProxyModel_MapSelectionToSource(const QAbstractProxyModel* self, QItemSelection* selection);
extern __declspec(dllexport) 
QItemSelection* QAbstractProxyModel_MapSelectionFromSource(const QAbstractProxyModel* self, QItemSelection* selection);
extern __declspec(dllexport) 
bool QAbstractProxyModel_Submit(QAbstractProxyModel* self);
extern __declspec(dllexport) 
void QAbstractProxyModel_Revert(QAbstractProxyModel* self);
extern __declspec(dllexport) 
QVariant* QAbstractProxyModel_Data(const QAbstractProxyModel* self, QModelIndex* proxyIndex, int role);
extern __declspec(dllexport) 
QVariant* QAbstractProxyModel_HeaderData(const QAbstractProxyModel* self, int section, int orientation, int role);
extern __declspec(dllexport) 
struct miqt_map /* of int to QVariant* */  QAbstractProxyModel_ItemData(const QAbstractProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) 
int QAbstractProxyModel_Flags(const QAbstractProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) 
bool QAbstractProxyModel_SetData(QAbstractProxyModel* self, QModelIndex* index, QVariant* value, int role);
extern __declspec(dllexport) 
bool QAbstractProxyModel_SetItemData(QAbstractProxyModel* self, QModelIndex* index, struct miqt_map /* of int to QVariant* */  roles);
extern __declspec(dllexport) 
bool QAbstractProxyModel_SetHeaderData(QAbstractProxyModel* self, int section, int orientation, QVariant* value, int role);
extern __declspec(dllexport) 
bool QAbstractProxyModel_ClearItemData(QAbstractProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) 
QModelIndex* QAbstractProxyModel_Buddy(const QAbstractProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) 
bool QAbstractProxyModel_CanFetchMore(const QAbstractProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
void QAbstractProxyModel_FetchMore(QAbstractProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
void QAbstractProxyModel_Sort(QAbstractProxyModel* self, int column, int order);
extern __declspec(dllexport) 
QSize* QAbstractProxyModel_Span(const QAbstractProxyModel* self, QModelIndex* index);
extern __declspec(dllexport) 
bool QAbstractProxyModel_HasChildren(const QAbstractProxyModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
QModelIndex* QAbstractProxyModel_Sibling(const QAbstractProxyModel* self, int row, int column, QModelIndex* idx);
extern __declspec(dllexport) 
QMimeData* QAbstractProxyModel_MimeData(const QAbstractProxyModel* self, struct miqt_array /* of QModelIndex* */  indexes);
extern __declspec(dllexport) 
bool QAbstractProxyModel_CanDropMimeData(const QAbstractProxyModel* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
bool QAbstractProxyModel_DropMimeData(QAbstractProxyModel* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QAbstractProxyModel_MimeTypes(const QAbstractProxyModel* self);
extern __declspec(dllexport) 
int QAbstractProxyModel_SupportedDragActions(const QAbstractProxyModel* self);
extern __declspec(dllexport) 
int QAbstractProxyModel_SupportedDropActions(const QAbstractProxyModel* self);
extern __declspec(dllexport) 
struct miqt_map /* of int to struct miqt_string */  QAbstractProxyModel_RoleNames(const QAbstractProxyModel* self);
extern __declspec(dllexport) 
struct miqt_string QAbstractProxyModel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAbstractProxyModel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAbstractProxyModel_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAbstractProxyModel_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAbstractProxyModel_override_virtual_Metacast(void* self, intptr_t slot);
void* QAbstractProxyModel_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAbstractProxyModel_Delete(QAbstractProxyModel* self, bool isSubclass);

}
