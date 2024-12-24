#pragma once
#ifndef MIQT_QT_GEN_QSTANDARDITEMMODEL_H
#define MIQT_QT_GEN_QSTANDARDITEMMODEL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QBrush QBrush;
typedef struct QDataStream QDataStream;
typedef struct QFont QFont;
typedef struct QIcon QIcon;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QModelIndex QModelIndex;
typedef struct QModelRoleDataSpan QModelRoleDataSpan;
typedef struct QObject QObject;
typedef struct QSize QSize;
typedef struct QStandardItem QStandardItem;
typedef struct QStandardItemModel QStandardItemModel;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QStandardItem* QStandardItem_new();
extern __declspec(dllexport) 
QStandardItem* QStandardItem_new2(struct miqt_string text);
extern __declspec(dllexport) 
QStandardItem* QStandardItem_new3(QIcon* icon, struct miqt_string text);
extern __declspec(dllexport) 
QStandardItem* QStandardItem_new4(int rows);
extern __declspec(dllexport) 
QStandardItem* QStandardItem_new5(int rows, int columns);
extern __declspec(dllexport) 
QVariant* QStandardItem_Data(const QStandardItem* self, int role);
extern __declspec(dllexport) 
void QStandardItem_MultiData(const QStandardItem* self, QModelRoleDataSpan* roleDataSpan);
extern __declspec(dllexport) 
void QStandardItem_SetData(QStandardItem* self, QVariant* value, int role);
extern __declspec(dllexport) 
void QStandardItem_ClearData(QStandardItem* self);
extern __declspec(dllexport) 
struct miqt_string QStandardItem_Text(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetText(QStandardItem* self, struct miqt_string text);
extern __declspec(dllexport) 
QIcon* QStandardItem_Icon(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetIcon(QStandardItem* self, QIcon* icon);
extern __declspec(dllexport) 
struct miqt_string QStandardItem_ToolTip(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetToolTip(QStandardItem* self, struct miqt_string toolTip);
extern __declspec(dllexport) 
struct miqt_string QStandardItem_StatusTip(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetStatusTip(QStandardItem* self, struct miqt_string statusTip);
extern __declspec(dllexport) 
struct miqt_string QStandardItem_WhatsThis(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetWhatsThis(QStandardItem* self, struct miqt_string whatsThis);
extern __declspec(dllexport) 
QSize* QStandardItem_SizeHint(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetSizeHint(QStandardItem* self, QSize* sizeHint);
extern __declspec(dllexport) 
QFont* QStandardItem_Font(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetFont(QStandardItem* self, QFont* font);
extern __declspec(dllexport) 
int QStandardItem_TextAlignment(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetTextAlignment(QStandardItem* self, int textAlignment);
extern __declspec(dllexport) 
QBrush* QStandardItem_Background(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetBackground(QStandardItem* self, QBrush* brush);
extern __declspec(dllexport) 
QBrush* QStandardItem_Foreground(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetForeground(QStandardItem* self, QBrush* brush);
extern __declspec(dllexport) 
int QStandardItem_CheckState(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetCheckState(QStandardItem* self, int checkState);
extern __declspec(dllexport) 
struct miqt_string QStandardItem_AccessibleText(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetAccessibleText(QStandardItem* self, struct miqt_string accessibleText);
extern __declspec(dllexport) 
struct miqt_string QStandardItem_AccessibleDescription(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetAccessibleDescription(QStandardItem* self, struct miqt_string accessibleDescription);
extern __declspec(dllexport) 
int QStandardItem_Flags(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetFlags(QStandardItem* self, int flags);
extern __declspec(dllexport) 
bool QStandardItem_IsEnabled(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetEnabled(QStandardItem* self, bool enabled);
extern __declspec(dllexport) 
bool QStandardItem_IsEditable(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetEditable(QStandardItem* self, bool editable);
extern __declspec(dllexport) 
bool QStandardItem_IsSelectable(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetSelectable(QStandardItem* self, bool selectable);
extern __declspec(dllexport) 
bool QStandardItem_IsCheckable(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetCheckable(QStandardItem* self, bool checkable);
extern __declspec(dllexport) 
bool QStandardItem_IsAutoTristate(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetAutoTristate(QStandardItem* self, bool tristate);
extern __declspec(dllexport) 
bool QStandardItem_IsUserTristate(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetUserTristate(QStandardItem* self, bool tristate);
extern __declspec(dllexport) 
bool QStandardItem_IsDragEnabled(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetDragEnabled(QStandardItem* self, bool dragEnabled);
extern __declspec(dllexport) 
bool QStandardItem_IsDropEnabled(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetDropEnabled(QStandardItem* self, bool dropEnabled);
extern __declspec(dllexport) 
QStandardItem* QStandardItem_Parent(const QStandardItem* self);
extern __declspec(dllexport) 
int QStandardItem_Row(const QStandardItem* self);
extern __declspec(dllexport) 
int QStandardItem_Column(const QStandardItem* self);
extern __declspec(dllexport) 
QModelIndex* QStandardItem_Index(const QStandardItem* self);
extern __declspec(dllexport) 
QStandardItemModel* QStandardItem_Model(const QStandardItem* self);
extern __declspec(dllexport) 
int QStandardItem_RowCount(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetRowCount(QStandardItem* self, int rows);
extern __declspec(dllexport) 
int QStandardItem_ColumnCount(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_SetColumnCount(QStandardItem* self, int columns);
extern __declspec(dllexport) 
bool QStandardItem_HasChildren(const QStandardItem* self);
extern __declspec(dllexport) 
QStandardItem* QStandardItem_Child(const QStandardItem* self, int row);
extern __declspec(dllexport) 
void QStandardItem_SetChild(QStandardItem* self, int row, int column, QStandardItem* item);
extern __declspec(dllexport) 
void QStandardItem_SetChild2(QStandardItem* self, int row, QStandardItem* item);
extern __declspec(dllexport) 
void QStandardItem_InsertRow(QStandardItem* self, int row, struct miqt_array /* of QStandardItem* */  items);
extern __declspec(dllexport) 
void QStandardItem_InsertColumn(QStandardItem* self, int column, struct miqt_array /* of QStandardItem* */  items);
extern __declspec(dllexport) 
void QStandardItem_InsertRows(QStandardItem* self, int row, struct miqt_array /* of QStandardItem* */  items);
extern __declspec(dllexport) 
void QStandardItem_InsertRows2(QStandardItem* self, int row, int count);
extern __declspec(dllexport) 
void QStandardItem_InsertColumns(QStandardItem* self, int column, int count);
extern __declspec(dllexport) 
void QStandardItem_RemoveRow(QStandardItem* self, int row);
extern __declspec(dllexport) 
void QStandardItem_RemoveColumn(QStandardItem* self, int column);
extern __declspec(dllexport) 
void QStandardItem_RemoveRows(QStandardItem* self, int row, int count);
extern __declspec(dllexport) 
void QStandardItem_RemoveColumns(QStandardItem* self, int column, int count);
extern __declspec(dllexport) 
void QStandardItem_AppendRow(QStandardItem* self, struct miqt_array /* of QStandardItem* */  items);
extern __declspec(dllexport) 
void QStandardItem_AppendRows(QStandardItem* self, struct miqt_array /* of QStandardItem* */  items);
extern __declspec(dllexport) 
void QStandardItem_AppendColumn(QStandardItem* self, struct miqt_array /* of QStandardItem* */  items);
extern __declspec(dllexport) 
void QStandardItem_InsertRow2(QStandardItem* self, int row, QStandardItem* item);
extern __declspec(dllexport) 
void QStandardItem_AppendRowWithItem(QStandardItem* self, QStandardItem* item);
extern __declspec(dllexport) 
QStandardItem* QStandardItem_TakeChild(QStandardItem* self, int row);
extern __declspec(dllexport) 
struct miqt_array /* of QStandardItem* */  QStandardItem_TakeRow(QStandardItem* self, int row);
extern __declspec(dllexport) 
struct miqt_array /* of QStandardItem* */  QStandardItem_TakeColumn(QStandardItem* self, int column);
extern __declspec(dllexport) 
void QStandardItem_SortChildren(QStandardItem* self, int column);
extern __declspec(dllexport) 
QStandardItem* QStandardItem_Clone(const QStandardItem* self);
extern __declspec(dllexport) 
int QStandardItem_Type(const QStandardItem* self);
extern __declspec(dllexport) 
void QStandardItem_Read(QStandardItem* self, QDataStream* in);
extern __declspec(dllexport) 
void QStandardItem_Write(const QStandardItem* self, QDataStream* out);
extern __declspec(dllexport) 
bool QStandardItem_OperatorLesser(const QStandardItem* self, QStandardItem* other);
extern __declspec(dllexport) 
QStandardItem* QStandardItem_Child2(const QStandardItem* self, int row, int column);
extern __declspec(dllexport) 
QStandardItem* QStandardItem_TakeChild2(QStandardItem* self, int row, int column);
extern __declspec(dllexport) 
void QStandardItem_SortChildren2(QStandardItem* self, int column, int order);
extern __declspec(dllexport) 
void QStandardItem_Delete(QStandardItem* self, bool isSubclass);

extern __declspec(dllexport) 
QStandardItemModel* QStandardItemModel_new();
extern __declspec(dllexport) 
QStandardItemModel* QStandardItemModel_new2(int rows, int columns);
extern __declspec(dllexport) 
QStandardItemModel* QStandardItemModel_new3(QObject* parent);
extern __declspec(dllexport) 
QStandardItemModel* QStandardItemModel_new4(int rows, int columns, QObject* parent);
extern __declspec(dllexport) 
void QStandardItemModel_virtbase(QStandardItemModel* src
, QAbstractItemModel** outptr_QAbstractItemModel
);
extern __declspec(dllexport) 
QMetaObject* QStandardItemModel_MetaObject(const QStandardItemModel* self);
extern __declspec(dllexport) 
void* QStandardItemModel_Metacast(QStandardItemModel* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QStandardItemModel_Tr(const char* s);
extern __declspec(dllexport) 
void QStandardItemModel_SetItemRoleNames(QStandardItemModel* self, struct miqt_map /* of int to struct miqt_string */  roleNames);
extern __declspec(dllexport) 
struct miqt_map /* of int to struct miqt_string */  QStandardItemModel_RoleNames(const QStandardItemModel* self);
extern __declspec(dllexport) 
QModelIndex* QStandardItemModel_Index(const QStandardItemModel* self, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
QModelIndex* QStandardItemModel_Parent(const QStandardItemModel* self, QModelIndex* child);
extern __declspec(dllexport) 
int QStandardItemModel_RowCount(const QStandardItemModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
int QStandardItemModel_ColumnCount(const QStandardItemModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
bool QStandardItemModel_HasChildren(const QStandardItemModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
QVariant* QStandardItemModel_Data(const QStandardItemModel* self, QModelIndex* index, int role);
extern __declspec(dllexport) 
void QStandardItemModel_MultiData(const QStandardItemModel* self, QModelIndex* index, QModelRoleDataSpan* roleDataSpan);
extern __declspec(dllexport) 
bool QStandardItemModel_SetData(QStandardItemModel* self, QModelIndex* index, QVariant* value, int role);
extern __declspec(dllexport) 
bool QStandardItemModel_ClearItemData(QStandardItemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
QVariant* QStandardItemModel_HeaderData(const QStandardItemModel* self, int section, int orientation, int role);
extern __declspec(dllexport) 
bool QStandardItemModel_SetHeaderData(QStandardItemModel* self, int section, int orientation, QVariant* value, int role);
extern __declspec(dllexport) 
bool QStandardItemModel_InsertRows(QStandardItemModel* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QStandardItemModel_InsertColumns(QStandardItemModel* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QStandardItemModel_RemoveRows(QStandardItemModel* self, int row, int count, QModelIndex* parent);
extern __declspec(dllexport) 
bool QStandardItemModel_RemoveColumns(QStandardItemModel* self, int column, int count, QModelIndex* parent);
extern __declspec(dllexport) 
int QStandardItemModel_Flags(const QStandardItemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
int QStandardItemModel_SupportedDropActions(const QStandardItemModel* self);
extern __declspec(dllexport) 
struct miqt_map /* of int to QVariant* */  QStandardItemModel_ItemData(const QStandardItemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
bool QStandardItemModel_SetItemData(QStandardItemModel* self, QModelIndex* index, struct miqt_map /* of int to QVariant* */  roles);
extern __declspec(dllexport) 
void QStandardItemModel_Clear(QStandardItemModel* self);
extern __declspec(dllexport) 
void QStandardItemModel_Sort(QStandardItemModel* self, int column, int order);
extern __declspec(dllexport) 
QStandardItem* QStandardItemModel_ItemFromIndex(const QStandardItemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
QModelIndex* QStandardItemModel_IndexFromItem(const QStandardItemModel* self, QStandardItem* item);
extern __declspec(dllexport) 
QStandardItem* QStandardItemModel_Item(const QStandardItemModel* self, int row);
extern __declspec(dllexport) 
void QStandardItemModel_SetItem(QStandardItemModel* self, int row, int column, QStandardItem* item);
extern __declspec(dllexport) 
void QStandardItemModel_SetItem2(QStandardItemModel* self, int row, QStandardItem* item);
extern __declspec(dllexport) 
QStandardItem* QStandardItemModel_InvisibleRootItem(const QStandardItemModel* self);
extern __declspec(dllexport) 
QStandardItem* QStandardItemModel_HorizontalHeaderItem(const QStandardItemModel* self, int column);
extern __declspec(dllexport) 
void QStandardItemModel_SetHorizontalHeaderItem(QStandardItemModel* self, int column, QStandardItem* item);
extern __declspec(dllexport) 
QStandardItem* QStandardItemModel_VerticalHeaderItem(const QStandardItemModel* self, int row);
extern __declspec(dllexport) 
void QStandardItemModel_SetVerticalHeaderItem(QStandardItemModel* self, int row, QStandardItem* item);
extern __declspec(dllexport) 
void QStandardItemModel_SetHorizontalHeaderLabels(QStandardItemModel* self, struct miqt_array /* of struct miqt_string */  labels);
extern __declspec(dllexport) 
void QStandardItemModel_SetVerticalHeaderLabels(QStandardItemModel* self, struct miqt_array /* of struct miqt_string */  labels);
extern __declspec(dllexport) 
void QStandardItemModel_SetRowCount(QStandardItemModel* self, int rows);
extern __declspec(dllexport) 
void QStandardItemModel_SetColumnCount(QStandardItemModel* self, int columns);
extern __declspec(dllexport) 
void QStandardItemModel_AppendRow(QStandardItemModel* self, struct miqt_array /* of QStandardItem* */  items);
extern __declspec(dllexport) 
void QStandardItemModel_AppendColumn(QStandardItemModel* self, struct miqt_array /* of QStandardItem* */  items);
extern __declspec(dllexport) 
void QStandardItemModel_AppendRowWithItem(QStandardItemModel* self, QStandardItem* item);
extern __declspec(dllexport) 
void QStandardItemModel_InsertRow(QStandardItemModel* self, int row, struct miqt_array /* of QStandardItem* */  items);
extern __declspec(dllexport) 
void QStandardItemModel_InsertColumn(QStandardItemModel* self, int column, struct miqt_array /* of QStandardItem* */  items);
extern __declspec(dllexport) 
void QStandardItemModel_InsertRow2(QStandardItemModel* self, int row, QStandardItem* item);
extern __declspec(dllexport) 
bool QStandardItemModel_InsertRowWithRow(QStandardItemModel* self, int row);
extern __declspec(dllexport) 
bool QStandardItemModel_InsertColumnWithColumn(QStandardItemModel* self, int column);
extern __declspec(dllexport) 
QStandardItem* QStandardItemModel_TakeItem(QStandardItemModel* self, int row);
extern __declspec(dllexport) 
struct miqt_array /* of QStandardItem* */  QStandardItemModel_TakeRow(QStandardItemModel* self, int row);
extern __declspec(dllexport) 
struct miqt_array /* of QStandardItem* */  QStandardItemModel_TakeColumn(QStandardItemModel* self, int column);
extern __declspec(dllexport) 
QStandardItem* QStandardItemModel_TakeHorizontalHeaderItem(QStandardItemModel* self, int column);
extern __declspec(dllexport) 
QStandardItem* QStandardItemModel_TakeVerticalHeaderItem(QStandardItemModel* self, int row);
extern __declspec(dllexport) 
QStandardItem* QStandardItemModel_ItemPrototype(const QStandardItemModel* self);
extern __declspec(dllexport) 
void QStandardItemModel_SetItemPrototype(QStandardItemModel* self, QStandardItem* item);
extern __declspec(dllexport) 
struct miqt_array /* of QStandardItem* */  QStandardItemModel_FindItems(const QStandardItemModel* self, struct miqt_string text);
extern __declspec(dllexport) 
int QStandardItemModel_SortRole(const QStandardItemModel* self);
extern __declspec(dllexport) 
void QStandardItemModel_SetSortRole(QStandardItemModel* self, int role);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QStandardItemModel_MimeTypes(const QStandardItemModel* self);
extern __declspec(dllexport) 
QMimeData* QStandardItemModel_MimeData(const QStandardItemModel* self, struct miqt_array /* of QModelIndex* */  indexes);
extern __declspec(dllexport) 
bool QStandardItemModel_DropMimeData(QStandardItemModel* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
void QStandardItemModel_ItemChanged(QStandardItemModel* self, QStandardItem* item);
void QStandardItemModel_connect_ItemChanged(QStandardItemModel* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QStandardItemModel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QStandardItemModel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
QStandardItem* QStandardItemModel_Item2(const QStandardItemModel* self, int row, int column);
extern __declspec(dllexport) 
bool QStandardItemModel_InsertRow22(QStandardItemModel* self, int row, QModelIndex* parent);
extern __declspec(dllexport) 
bool QStandardItemModel_InsertColumn2(QStandardItemModel* self, int column, QModelIndex* parent);
extern __declspec(dllexport) 
QStandardItem* QStandardItemModel_TakeItem2(QStandardItemModel* self, int row, int column);
extern __declspec(dllexport) 
struct miqt_array /* of QStandardItem* */  QStandardItemModel_FindItems2(const QStandardItemModel* self, struct miqt_string text, int flags);
extern __declspec(dllexport) 
struct miqt_array /* of QStandardItem* */  QStandardItemModel_FindItems3(const QStandardItemModel* self, struct miqt_string text, int flags, int column);
extern __declspec(dllexport) 
void QStandardItemModel_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QStandardItemModel_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QStandardItemModel_override_virtual_Metacast(void* self, intptr_t slot);
void* QStandardItemModel_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QStandardItemModel_Delete(QStandardItemModel* self, bool isSubclass);

}
