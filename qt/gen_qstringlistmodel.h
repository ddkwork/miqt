#pragma once
#ifndef MIQT_QT_GEN_QSTRINGLISTMODEL_H
#define MIQT_QT_GEN_QSTRINGLISTMODEL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractItemModel;
class QAbstractListModel;
class QMetaObject;
class QMimeData;
class QModelIndex;
class QObject;
class QStringListModel;
class QVariant;
class _GUID;
class type_info;
#else
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QAbstractListModel QAbstractListModel;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QStringListModel QStringListModel;
typedef struct QVariant QVariant;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QStringListModel* QStringListModel_new();
extern __declspec(dllexport) QStringListModel* QStringListModel_new2(struct miqt_array /* of struct miqt_string */  strings);
extern __declspec(dllexport) QStringListModel* QStringListModel_new3(QObject* parent);
extern __declspec(dllexport) QStringListModel* QStringListModel_new4(struct miqt_array /* of struct miqt_string */  strings, QObject* parent);
extern __declspec(dllexport) void QStringListModel_virtbase(QStringListModel* src, QAbstractListModel** outptr_QAbstractListModel);
extern __declspec(dllexport) QMetaObject* QStringListModel_MetaObject(const QStringListModel* self);
extern __declspec(dllexport) void* QStringListModel_Metacast(QStringListModel* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QStringListModel_Tr(const char* s);
extern __declspec(dllexport) int QStringListModel_RowCount(const QStringListModel* self, QModelIndex* parent);
extern __declspec(dllexport) QModelIndex* QStringListModel_Sibling(const QStringListModel* self, int row, int column, QModelIndex* idx);
extern __declspec(dllexport) QVariant* QStringListModel_Data(const QStringListModel* self, QModelIndex* index, int role);
extern __declspec(dllexport) bool QStringListModel_SetData(QStringListModel* self, QModelIndex* index, QVariant* value, int role);
extern __declspec(dllexport) bool QStringListModel_ClearItemData(QStringListModel* self, QModelIndex* index);
extern __declspec(dllexport) int QStringListModel_Flags(const QStringListModel* self, QModelIndex* index);
extern __declspec(dllexport) bool QStringListModel_InsertRows(QStringListModel* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) bool QStringListModel_RemoveRows(QStringListModel* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) bool QStringListModel_MoveRows(QStringListModel* self, QModelIndex* sourceParent, int sourceRow, int count, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) struct miqt_map /* of int to QVariant* */  QStringListModel_ItemData(const QStringListModel* self, QModelIndex* index);
extern __declspec(dllexport) bool QStringListModel_SetItemData(QStringListModel* self, QModelIndex* index, struct miqt_map /* of int to QVariant* */  roles);
extern __declspec(dllexport) void QStringListModel_Sort(QStringListModel* self, int column, int order);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QStringListModel_StringList(const QStringListModel* self);
extern __declspec(dllexport) void QStringListModel_SetStringList(QStringListModel* self, struct miqt_array /* of struct miqt_string */  strings);
extern __declspec(dllexport) int QStringListModel_SupportedDropActions(const QStringListModel* self);
extern __declspec(dllexport) struct miqt_string QStringListModel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QStringListModel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QStringListModel_override_virtual_RowCount(void* self, intptr_t slot);
int QStringListModel_virtualbase_RowCount(const void* self, QModelIndex* parent);
extern __declspec(dllexport) void QStringListModel_override_virtual_Sibling(void* self, intptr_t slot);
QModelIndex* QStringListModel_virtualbase_Sibling(const void* self, int row, int column, QModelIndex* idx);
extern __declspec(dllexport) void QStringListModel_override_virtual_Data(void* self, intptr_t slot);
QVariant* QStringListModel_virtualbase_Data(const void* self, QModelIndex* index, int role);
extern __declspec(dllexport) void QStringListModel_override_virtual_SetData(void* self, intptr_t slot);
bool QStringListModel_virtualbase_SetData(void* self, QModelIndex* index, QVariant* value, int role);
extern __declspec(dllexport) void QStringListModel_override_virtual_ClearItemData(void* self, intptr_t slot);
bool QStringListModel_virtualbase_ClearItemData(void* self, QModelIndex* index);
extern __declspec(dllexport) void QStringListModel_override_virtual_Flags(void* self, intptr_t slot);
int QStringListModel_virtualbase_Flags(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QStringListModel_override_virtual_InsertRows(void* self, intptr_t slot);
bool QStringListModel_virtualbase_InsertRows(void* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) void QStringListModel_override_virtual_RemoveRows(void* self, intptr_t slot);
bool QStringListModel_virtualbase_RemoveRows(void* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) void QStringListModel_override_virtual_MoveRows(void* self, intptr_t slot);
bool QStringListModel_virtualbase_MoveRows(void* self, QModelIndex* sourceParent, int sourceRow, int count, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) void QStringListModel_override_virtual_ItemData(void* self, intptr_t slot);
struct miqt_map /* of int to QVariant* */  QStringListModel_virtualbase_ItemData(const void* self, QModelIndex* index);
extern __declspec(dllexport) void QStringListModel_override_virtual_SetItemData(void* self, intptr_t slot);
bool QStringListModel_virtualbase_SetItemData(void* self, QModelIndex* index, struct miqt_map /* of int to QVariant* */  roles);
extern __declspec(dllexport) void QStringListModel_override_virtual_Sort(void* self, intptr_t slot);
void QStringListModel_virtualbase_Sort(void* self, int column, int order);
extern __declspec(dllexport) void QStringListModel_override_virtual_SupportedDropActions(void* self, intptr_t slot);
int QStringListModel_virtualbase_SupportedDropActions(const void* self);
extern __declspec(dllexport) void QStringListModel_override_virtual_Index(void* self, intptr_t slot);
QModelIndex* QStringListModel_virtualbase_Index(const void* self, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) void QStringListModel_override_virtual_DropMimeData(void* self, intptr_t slot);
bool QStringListModel_virtualbase_DropMimeData(void* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) void QStringListModel_Delete(QStringListModel* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
