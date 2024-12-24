#pragma once
#ifndef MIQT_QT_GEN_QABSTRACTITEMMODEL_H
#define MIQT_QT_GEN_QABSTRACTITEMMODEL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QAbstractListModel QAbstractListModel;
typedef struct QAbstractTableModel QAbstractTableModel;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QModelIndex QModelIndex;
typedef struct QModelRoleData QModelRoleData;
typedef struct QModelRoleDataSpan QModelRoleDataSpan;
typedef struct QObject QObject;
typedef struct QPersistentModelIndex QPersistentModelIndex;
typedef struct QSize QSize;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QModelRoleData* QModelRoleData_new(int role);
extern __declspec(dllexport) 
int QModelRoleData_Role(const QModelRoleData* self);
extern __declspec(dllexport) 
QVariant* QModelRoleData_Data(QModelRoleData* self);
extern __declspec(dllexport) 
QVariant* QModelRoleData_Data2(const QModelRoleData* self);
extern __declspec(dllexport) 
void QModelRoleData_ClearData(QModelRoleData* self);
extern __declspec(dllexport) 
void QModelRoleData_OperatorAssign(QModelRoleData* self, QModelRoleData* param1);
extern __declspec(dllexport) 
void QModelRoleData_Delete(QModelRoleData* self, bool isSubclass);

extern __declspec(dllexport) 
QModelRoleDataSpan* QModelRoleDataSpan_new();
extern __declspec(dllexport) 
QModelRoleDataSpan* QModelRoleDataSpan_new2(QModelRoleData* modelRoleData);
extern __declspec(dllexport) 
QModelRoleDataSpan* QModelRoleDataSpan_new3(QModelRoleData* modelRoleData, ptrdiff_t lenVal);
extern __declspec(dllexport) 
QModelRoleDataSpan* QModelRoleDataSpan_new4(QModelRoleDataSpan* param1);
extern __declspec(dllexport) 
ptrdiff_t QModelRoleDataSpan_Size(const QModelRoleDataSpan* self);
extern __declspec(dllexport) 
ptrdiff_t QModelRoleDataSpan_Length(const QModelRoleDataSpan* self);
extern __declspec(dllexport) 
QModelRoleData* QModelRoleDataSpan_Data(const QModelRoleDataSpan* self);
extern __declspec(dllexport) 
QModelRoleData* QModelRoleDataSpan_Begin(const QModelRoleDataSpan* self);
extern __declspec(dllexport) 
QModelRoleData* QModelRoleDataSpan_End(const QModelRoleDataSpan* self);
extern __declspec(dllexport) 
QModelRoleData* QModelRoleDataSpan_OperatorSubscript(const QModelRoleDataSpan* self, ptrdiff_t index);
extern __declspec(dllexport) 
QVariant* QModelRoleDataSpan_DataForRole(const QModelRoleDataSpan* self, int role);
extern __declspec(dllexport) 
void QModelRoleDataSpan_Delete(QModelRoleDataSpan* self, bool isSubclass);

extern __declspec(dllexport) 
QModelIndex* QModelIndex_new();
extern __declspec(dllexport) 
QModelIndex* QModelIndex_new2(QModelIndex* param1);
extern __declspec(dllexport) 
int QModelIndex_Row(const QModelIndex* self);
extern __declspec(dllexport) 
int QModelIndex_Column(const QModelIndex* self);
extern __declspec(dllexport) 
uintptr_t QModelIndex_InternalId(const QModelIndex* self);
extern __declspec(dllexport) 
void* QModelIndex_InternalPointer(const QModelIndex* self);
extern __declspec(dllexport) 
const void* QModelIndex_ConstInternalPointer(const QModelIndex* self);
extern __declspec(dllexport) 
QModelIndex* QModelIndex_Parent(const QModelIndex* self);
extern __declspec(dllexport) 
QModelIndex* QModelIndex_Sibling(const QModelIndex* self, int row, int column);
extern __declspec(dllexport) 
QModelIndex* QModelIndex_SiblingAtColumn(const QModelIndex* self, int column);
extern __declspec(dllexport) 
QModelIndex* QModelIndex_SiblingAtRow(const QModelIndex* self, int row);
extern __declspec(dllexport) 
QVariant* QModelIndex_Data(const QModelIndex* self);
extern __declspec(dllexport) 
void QModelIndex_MultiData(const QModelIndex* self, QModelRoleDataSpan* roleDataSpan);
extern __declspec(dllexport) 
int QModelIndex_Flags(const QModelIndex* self);
extern __declspec(dllexport) 
QAbstractItemModel* QModelIndex_Model(const QModelIndex* self);
extern __declspec(dllexport) 
bool QModelIndex_IsValid(const QModelIndex* self);
extern __declspec(dllexport) 
QVariant* QModelIndex_Data1(const QModelIndex* self, int role);
extern __declspec(dllexport) 
void QModelIndex_Delete(QModelIndex* self, bool isSubclass);

extern __declspec(dllexport) 
QPersistentModelIndex* QPersistentModelIndex_new();
extern __declspec(dllexport) 
QPersistentModelIndex* QPersistentModelIndex_new2(QModelIndex* index);
extern __declspec(dllexport) 
QPersistentModelIndex* QPersistentModelIndex_new3(QPersistentModelIndex* other);
extern __declspec(dllexport) 
void QPersistentModelIndex_OperatorAssign(QPersistentModelIndex* self, QPersistentModelIndex* other);
extern __declspec(dllexport) 
void QPersistentModelIndex_Swap(QPersistentModelIndex* self, QPersistentModelIndex* other);
extern __declspec(dllexport) 
void QPersistentModelIndex_OperatorAssignWithOther(QPersistentModelIndex* self, QModelIndex* other);
extern __declspec(dllexport) 
int QPersistentModelIndex_Row(const QPersistentModelIndex* self);
extern __declspec(dllexport) 
int QPersistentModelIndex_Column(const QPersistentModelIndex* self);
extern __declspec(dllexport) 
void* QPersistentModelIndex_InternalPointer(const QPersistentModelIndex* self);
extern __declspec(dllexport) 
const void* QPersistentModelIndex_ConstInternalPointer(const QPersistentModelIndex* self);
extern __declspec(dllexport) 
uintptr_t QPersistentModelIndex_InternalId(const QPersistentModelIndex* self);
extern __declspec(dllexport) 
QModelIndex* QPersistentModelIndex_Parent(const QPersistentModelIndex* self);
extern __declspec(dllexport) 
QModelIndex* QPersistentModelIndex_Sibling(const QPersistentModelIndex* self, int row, int column);
extern __declspec(dllexport) 
QVariant* QPersistentModelIndex_Data(const QPersistentModelIndex* self);
extern __declspec(dllexport) 
void QPersistentModelIndex_MultiData(const QPersistentModelIndex* self, QModelRoleDataSpan* roleDataSpan);
extern __declspec(dllexport) 
int QPersistentModelIndex_Flags(const QPersistentModelIndex* self);
extern __declspec(dllexport) 
QAbstractItemModel* QPersistentModelIndex_Model(const QPersistentModelIndex* self);
extern __declspec(dllexport) 
bool QPersistentModelIndex_IsValid(const QPersistentModelIndex* self);
extern __declspec(dllexport) 
QVariant* QPersistentModelIndex_Data1(const QPersistentModelIndex* self, int role);
extern __declspec(dllexport) 
void QPersistentModelIndex_Delete(QPersistentModelIndex* self, bool isSubclass);

extern __declspec(dllexport) 
QAbstractItemModel* QAbstractItemModel_new();
extern __declspec(dllexport) 
QAbstractItemModel* QAbstractItemModel_new2(QObject* parent);
extern __declspec(dllexport) 
void QAbstractItemModel_virtbase(QAbstractItemModel* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QAbstractItemModel_MetaObject(const QAbstractItemModel* self);
extern __declspec(dllexport) 
void* QAbstractItemModel_Metacast(QAbstractItemModel* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAbstractItemModel_Tr(const char* s);
extern __declspec(dllexport) 
bool QAbstractItemModel_HasIndex(const QAbstractItemModel* self, int row, int column);
extern __declspec(dllexport) 
QModelIndex* QAbstractItemModel_Index(const QAbstractItemModel* self, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
QModelIndex* QAbstractItemModel_Parent(const QAbstractItemModel* self, QModelIndex* child);
extern __declspec(dllexport) 
QModelIndex* QAbstractItemModel_Sibling(const QAbstractItemModel* self, int row, int column, QModelIndex* idx);
extern __declspec(dllexport) 
int QAbstractItemModel_RowCount(const QAbstractItemModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
int QAbstractItemModel_ColumnCount(const QAbstractItemModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
bool QAbstractItemModel_HasChildren(const QAbstractItemModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
QVariant* QAbstractItemModel_Data(const QAbstractItemModel* self, QModelIndex* index, int role);
extern __declspec(dllexport) 
bool QAbstractItemModel_SetData(QAbstractItemModel* self, QModelIndex* index, QVariant* value, int role);
extern __declspec(dllexport) 
QVariant* QAbstractItemModel_HeaderData(const QAbstractItemModel* self, int section, int orientation, int role);
extern __declspec(dllexport) 
bool QAbstractItemModel_SetHeaderData(QAbstractItemModel* self, int section, int orientation, QVariant* value, int role);
extern __declspec(dllexport) 
struct miqt_map /* of int to QVariant* */  QAbstractItemModel_ItemData(const QAbstractItemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
bool QAbstractItemModel_SetItemData(QAbstractItemModel* self, QModelIndex* index, struct miqt_map /* of int to QVariant* */  roles);
extern __declspec(dllexport) 
bool QAbstractItemModel_ClearItemData(QAbstractItemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QAbstractItemModel_MimeTypes(const QAbstractItemModel* self);
extern __declspec(dllexport) 
QMimeData* QAbstractItemModel_MimeData(const QAbstractItemModel* self, struct miqt_array /* of QModelIndex* */  indexes);
extern __declspec(dllexport) 
bool QAbstractItemModel_CanDropMimeData(const QAbstractItemModel* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
bool QAbstractItemModel_DropMimeData(QAbstractItemModel* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
int QAbstractItemModel_SupportedDropActions(const QAbstractItemModel* self);
extern __declspec(dllexport) 
int QAbstractItemModel_SupportedDragActions(const QAbstractItemModel* self);
extern __declspec(dllexport) 
bool QAbstractItemModel_InsertRows(QAbstractItemModel* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QAbstractItemModel_InsertColumns(QAbstractItemModel* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QAbstractItemModel_RemoveRows(QAbstractItemModel* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QAbstractItemModel_RemoveColumns(QAbstractItemModel* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QAbstractItemModel_MoveRows(QAbstractItemModel* self, QModelIndex* sourceParent, int sourceRow, int count, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) 
bool QAbstractItemModel_MoveColumns(QAbstractItemModel* self, QModelIndex* sourceParent, int sourceColumn, int count, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) 
bool QAbstractItemModel_InsertRow(QAbstractItemModel* self, int row);
extern __declspec(dllexport) 
bool QAbstractItemModel_InsertColumn(QAbstractItemModel* self, int column);
extern __declspec(dllexport) 
bool QAbstractItemModel_RemoveRow(QAbstractItemModel* self, int row);
extern __declspec(dllexport) 
bool QAbstractItemModel_RemoveColumn(QAbstractItemModel* self, int column);
extern __declspec(dllexport) 
bool QAbstractItemModel_MoveRow(QAbstractItemModel* self, QModelIndex* sourceParent, int sourceRow, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) 
bool QAbstractItemModel_MoveColumn(QAbstractItemModel* self, QModelIndex* sourceParent, int sourceColumn, QModelIndex* destinationParent, int destinationChild);
extern __declspec(dllexport) 
void QAbstractItemModel_FetchMore(QAbstractItemModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
bool QAbstractItemModel_CanFetchMore(const QAbstractItemModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
int QAbstractItemModel_Flags(const QAbstractItemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
void QAbstractItemModel_Sort(QAbstractItemModel* self, int column, int order);
extern __declspec(dllexport) 
QModelIndex* QAbstractItemModel_Buddy(const QAbstractItemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
struct miqt_array /* of QModelIndex* */  QAbstractItemModel_Match(const QAbstractItemModel* self, QModelIndex* start, int role, QVariant* value, int hits, int flags);
extern __declspec(dllexport) 
QSize* QAbstractItemModel_Span(const QAbstractItemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
struct miqt_map /* of int to struct miqt_string */  QAbstractItemModel_RoleNames(const QAbstractItemModel* self);
extern __declspec(dllexport) 
bool QAbstractItemModel_CheckIndex(const QAbstractItemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
void QAbstractItemModel_MultiData(const QAbstractItemModel* self, QModelIndex* index, QModelRoleDataSpan* roleDataSpan);
extern __declspec(dllexport) 
void QAbstractItemModel_DataChanged(QAbstractItemModel* self, QModelIndex* topLeft, QModelIndex* bottomRight);
void QAbstractItemModel_connect_DataChanged(QAbstractItemModel* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractItemModel_HeaderDataChanged(QAbstractItemModel* self, int orientation, int first, int last);
void QAbstractItemModel_connect_HeaderDataChanged(QAbstractItemModel* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractItemModel_LayoutChanged(QAbstractItemModel* self);
void QAbstractItemModel_connect_LayoutChanged(QAbstractItemModel* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractItemModel_LayoutAboutToBeChanged(QAbstractItemModel* self);
void QAbstractItemModel_connect_LayoutAboutToBeChanged(QAbstractItemModel* self, intptr_t slot);
extern __declspec(dllexport) 
bool QAbstractItemModel_Submit(QAbstractItemModel* self);
extern __declspec(dllexport) 
void QAbstractItemModel_Revert(QAbstractItemModel* self);
extern __declspec(dllexport) 
void QAbstractItemModel_ResetInternalData(QAbstractItemModel* self);
extern __declspec(dllexport) 
struct miqt_string QAbstractItemModel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAbstractItemModel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
bool QAbstractItemModel_HasIndex3(const QAbstractItemModel* self, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
bool QAbstractItemModel_InsertRow2(QAbstractItemModel* self, int row, QModelIndex* parent);
extern __declspec(dllexport) 
bool QAbstractItemModel_InsertColumn2(QAbstractItemModel* self, int column, QModelIndex* parent);
extern __declspec(dllexport) 
bool QAbstractItemModel_RemoveRow2(QAbstractItemModel* self, int row, QModelIndex* parent);
extern __declspec(dllexport) 
bool QAbstractItemModel_RemoveColumn2(QAbstractItemModel* self, int column, QModelIndex* parent);
extern __declspec(dllexport) 
bool QAbstractItemModel_CheckIndex2(const QAbstractItemModel* self, QModelIndex* index, CheckIndexOptions options);
extern __declspec(dllexport) 
void QAbstractItemModel_DataChanged3(QAbstractItemModel* self, QModelIndex* topLeft, QModelIndex* bottomRight, struct miqt_array /* of int */  roles);
void QAbstractItemModel_connect_DataChanged3(QAbstractItemModel* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractItemModel_LayoutChanged1(QAbstractItemModel* self, struct miqt_array /* of QPersistentModelIndex* */  parents);
void QAbstractItemModel_connect_LayoutChanged1(QAbstractItemModel* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractItemModel_LayoutChanged2(QAbstractItemModel* self, struct miqt_array /* of QPersistentModelIndex* */  parents, int hint);
void QAbstractItemModel_connect_LayoutChanged2(QAbstractItemModel* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractItemModel_LayoutAboutToBeChanged1(QAbstractItemModel* self, struct miqt_array /* of QPersistentModelIndex* */  parents);
void QAbstractItemModel_connect_LayoutAboutToBeChanged1(QAbstractItemModel* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractItemModel_LayoutAboutToBeChanged2(QAbstractItemModel* self, struct miqt_array /* of QPersistentModelIndex* */  parents, int hint);
void QAbstractItemModel_connect_LayoutAboutToBeChanged2(QAbstractItemModel* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractItemModel_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAbstractItemModel_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAbstractItemModel_override_virtual_Metacast(void* self, intptr_t slot);
void* QAbstractItemModel_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAbstractItemModel_Delete(QAbstractItemModel* self, bool isSubclass);

extern __declspec(dllexport) 
QAbstractTableModel* QAbstractTableModel_new();
extern __declspec(dllexport) 
QAbstractTableModel* QAbstractTableModel_new2(QObject* parent);
extern __declspec(dllexport) 
void QAbstractTableModel_virtbase(QAbstractTableModel* src
, QAbstractItemModel** outptr_QAbstractItemModel
);
extern __declspec(dllexport) 
QMetaObject* QAbstractTableModel_MetaObject(const QAbstractTableModel* self);
extern __declspec(dllexport) 
void* QAbstractTableModel_Metacast(QAbstractTableModel* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAbstractTableModel_Tr(const char* s);
extern __declspec(dllexport) 
QModelIndex* QAbstractTableModel_Index(const QAbstractTableModel* self, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
QModelIndex* QAbstractTableModel_Sibling(const QAbstractTableModel* self, int row, int column, QModelIndex* idx);
extern __declspec(dllexport) 
bool QAbstractTableModel_DropMimeData(QAbstractTableModel* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
int QAbstractTableModel_Flags(const QAbstractTableModel* self, QModelIndex* index);
extern __declspec(dllexport) 
struct miqt_string QAbstractTableModel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAbstractTableModel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAbstractTableModel_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAbstractTableModel_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAbstractTableModel_override_virtual_Metacast(void* self, intptr_t slot);
void* QAbstractTableModel_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAbstractTableModel_Delete(QAbstractTableModel* self, bool isSubclass);

extern __declspec(dllexport) 
QAbstractListModel* QAbstractListModel_new();
extern __declspec(dllexport) 
QAbstractListModel* QAbstractListModel_new2(QObject* parent);
extern __declspec(dllexport) 
void QAbstractListModel_virtbase(QAbstractListModel* src
, QAbstractItemModel** outptr_QAbstractItemModel
);
extern __declspec(dllexport) 
QMetaObject* QAbstractListModel_MetaObject(const QAbstractListModel* self);
extern __declspec(dllexport) 
void* QAbstractListModel_Metacast(QAbstractListModel* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAbstractListModel_Tr(const char* s);
extern __declspec(dllexport) 
QModelIndex* QAbstractListModel_Index(const QAbstractListModel* self, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
QModelIndex* QAbstractListModel_Sibling(const QAbstractListModel* self, int row, int column, QModelIndex* idx);
extern __declspec(dllexport) 
bool QAbstractListModel_DropMimeData(QAbstractListModel* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
int QAbstractListModel_Flags(const QAbstractListModel* self, QModelIndex* index);
extern __declspec(dllexport) 
struct miqt_string QAbstractListModel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAbstractListModel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAbstractListModel_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAbstractListModel_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAbstractListModel_override_virtual_Metacast(void* self, intptr_t slot);
void* QAbstractListModel_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAbstractListModel_Delete(QAbstractListModel* self, bool isSubclass);

}
