#ifndef GEN_QSTANDARDITEMMODEL_H
#define GEN_QSTANDARDITEMMODEL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QBrush;
class QDataStream;
class QFont;
class QIcon;
class QMetaObject;
class QMimeData;
class QModelIndex;
class QObject;
class QSize;
class QStandardItem;
class QStandardItemModel;
class QVariant;
#else
typedef struct QBrush QBrush;
typedef struct QDataStream QDataStream;
typedef struct QFont QFont;
typedef struct QIcon QIcon;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QSize QSize;
typedef struct QStandardItem QStandardItem;
typedef struct QStandardItemModel QStandardItemModel;
typedef struct QVariant QVariant;
#endif

QStandardItem* QStandardItem_new();
QStandardItem* QStandardItem_new2(const char* text, size_t text_Strlen);
QStandardItem* QStandardItem_new3(QIcon* icon, const char* text, size_t text_Strlen);
QStandardItem* QStandardItem_new4(int rows);
QStandardItem* QStandardItem_new5(int rows, int columns);
QVariant* QStandardItem_Data(const QStandardItem* self);
void QStandardItem_SetData(QStandardItem* self, QVariant* value);
void QStandardItem_ClearData(QStandardItem* self);
void QStandardItem_Text(const QStandardItem* self, char** _out, int* _out_Strlen);
void QStandardItem_SetText(QStandardItem* self, const char* text, size_t text_Strlen);
QIcon* QStandardItem_Icon(const QStandardItem* self);
void QStandardItem_SetIcon(QStandardItem* self, QIcon* icon);
void QStandardItem_ToolTip(const QStandardItem* self, char** _out, int* _out_Strlen);
void QStandardItem_SetToolTip(QStandardItem* self, const char* toolTip, size_t toolTip_Strlen);
void QStandardItem_StatusTip(const QStandardItem* self, char** _out, int* _out_Strlen);
void QStandardItem_SetStatusTip(QStandardItem* self, const char* statusTip, size_t statusTip_Strlen);
void QStandardItem_WhatsThis(const QStandardItem* self, char** _out, int* _out_Strlen);
void QStandardItem_SetWhatsThis(QStandardItem* self, const char* whatsThis, size_t whatsThis_Strlen);
QSize* QStandardItem_SizeHint(const QStandardItem* self);
void QStandardItem_SetSizeHint(QStandardItem* self, QSize* sizeHint);
QFont* QStandardItem_Font(const QStandardItem* self);
void QStandardItem_SetFont(QStandardItem* self, QFont* font);
int QStandardItem_TextAlignment(const QStandardItem* self);
void QStandardItem_SetTextAlignment(QStandardItem* self, int textAlignment);
QBrush* QStandardItem_Background(const QStandardItem* self);
void QStandardItem_SetBackground(QStandardItem* self, QBrush* brush);
QBrush* QStandardItem_Foreground(const QStandardItem* self);
void QStandardItem_SetForeground(QStandardItem* self, QBrush* brush);
uintptr_t QStandardItem_CheckState(const QStandardItem* self);
void QStandardItem_SetCheckState(QStandardItem* self, uintptr_t checkState);
void QStandardItem_AccessibleText(const QStandardItem* self, char** _out, int* _out_Strlen);
void QStandardItem_SetAccessibleText(QStandardItem* self, const char* accessibleText, size_t accessibleText_Strlen);
void QStandardItem_AccessibleDescription(const QStandardItem* self, char** _out, int* _out_Strlen);
void QStandardItem_SetAccessibleDescription(QStandardItem* self, const char* accessibleDescription, size_t accessibleDescription_Strlen);
int QStandardItem_Flags(const QStandardItem* self);
void QStandardItem_SetFlags(QStandardItem* self, int flags);
bool QStandardItem_IsEnabled(const QStandardItem* self);
void QStandardItem_SetEnabled(QStandardItem* self, bool enabled);
bool QStandardItem_IsEditable(const QStandardItem* self);
void QStandardItem_SetEditable(QStandardItem* self, bool editable);
bool QStandardItem_IsSelectable(const QStandardItem* self);
void QStandardItem_SetSelectable(QStandardItem* self, bool selectable);
bool QStandardItem_IsCheckable(const QStandardItem* self);
void QStandardItem_SetCheckable(QStandardItem* self, bool checkable);
bool QStandardItem_IsAutoTristate(const QStandardItem* self);
void QStandardItem_SetAutoTristate(QStandardItem* self, bool tristate);
bool QStandardItem_IsUserTristate(const QStandardItem* self);
void QStandardItem_SetUserTristate(QStandardItem* self, bool tristate);
bool QStandardItem_IsTristate(const QStandardItem* self);
void QStandardItem_SetTristate(QStandardItem* self, bool tristate);
bool QStandardItem_IsDragEnabled(const QStandardItem* self);
void QStandardItem_SetDragEnabled(QStandardItem* self, bool dragEnabled);
bool QStandardItem_IsDropEnabled(const QStandardItem* self);
void QStandardItem_SetDropEnabled(QStandardItem* self, bool dropEnabled);
QStandardItem* QStandardItem_Parent(const QStandardItem* self);
int QStandardItem_Row(const QStandardItem* self);
int QStandardItem_Column(const QStandardItem* self);
QModelIndex* QStandardItem_Index(const QStandardItem* self);
QStandardItemModel* QStandardItem_Model(const QStandardItem* self);
int QStandardItem_RowCount(const QStandardItem* self);
void QStandardItem_SetRowCount(QStandardItem* self, int rows);
int QStandardItem_ColumnCount(const QStandardItem* self);
void QStandardItem_SetColumnCount(QStandardItem* self, int columns);
bool QStandardItem_HasChildren(const QStandardItem* self);
QStandardItem* QStandardItem_Child(const QStandardItem* self, int row);
void QStandardItem_SetChild(QStandardItem* self, int row, int column, QStandardItem* item);
void QStandardItem_SetChild2(QStandardItem* self, int row, QStandardItem* item);
void QStandardItem_InsertRow(QStandardItem* self, int row, QStandardItem** items, size_t items_len);
void QStandardItem_InsertColumn(QStandardItem* self, int column, QStandardItem** items, size_t items_len);
void QStandardItem_InsertRows(QStandardItem* self, int row, QStandardItem** items, size_t items_len);
void QStandardItem_InsertRows2(QStandardItem* self, int row, int count);
void QStandardItem_InsertColumns(QStandardItem* self, int column, int count);
void QStandardItem_RemoveRow(QStandardItem* self, int row);
void QStandardItem_RemoveColumn(QStandardItem* self, int column);
void QStandardItem_RemoveRows(QStandardItem* self, int row, int count);
void QStandardItem_RemoveColumns(QStandardItem* self, int column, int count);
void QStandardItem_AppendRow(QStandardItem* self, QStandardItem** items, size_t items_len);
void QStandardItem_AppendRows(QStandardItem* self, QStandardItem** items, size_t items_len);
void QStandardItem_AppendColumn(QStandardItem* self, QStandardItem** items, size_t items_len);
void QStandardItem_InsertRow2(QStandardItem* self, int row, QStandardItem* item);
void QStandardItem_AppendRowWithItem(QStandardItem* self, QStandardItem* item);
QStandardItem* QStandardItem_TakeChild(QStandardItem* self, int row);
void QStandardItem_TakeRow(QStandardItem* self, int row, QStandardItem*** _out, size_t* _out_len);
void QStandardItem_TakeColumn(QStandardItem* self, int column, QStandardItem*** _out, size_t* _out_len);
void QStandardItem_SortChildren(QStandardItem* self, int column);
QStandardItem* QStandardItem_Clone(const QStandardItem* self);
int QStandardItem_Type(const QStandardItem* self);
void QStandardItem_Read(QStandardItem* self, QDataStream* in);
void QStandardItem_Write(const QStandardItem* self, QDataStream* out);
bool QStandardItem_OperatorLesser(const QStandardItem* self, QStandardItem* other);
QVariant* QStandardItem_Data1(const QStandardItem* self, int role);
void QStandardItem_SetData2(QStandardItem* self, QVariant* value, int role);
QStandardItem* QStandardItem_Child2(const QStandardItem* self, int row, int column);
QStandardItem* QStandardItem_TakeChild2(QStandardItem* self, int row, int column);
void QStandardItem_SortChildren2(QStandardItem* self, int column, uintptr_t order);
void QStandardItem_Delete(QStandardItem* self);

QStandardItemModel* QStandardItemModel_new();
QStandardItemModel* QStandardItemModel_new2(int rows, int columns);
QStandardItemModel* QStandardItemModel_new3(QObject* parent);
QStandardItemModel* QStandardItemModel_new4(int rows, int columns, QObject* parent);
QMetaObject* QStandardItemModel_MetaObject(const QStandardItemModel* self);
void QStandardItemModel_Tr(const char* s, char** _out, int* _out_Strlen);
void QStandardItemModel_TrUtf8(const char* s, char** _out, int* _out_Strlen);
QModelIndex* QStandardItemModel_Index(const QStandardItemModel* self, int row, int column);
QModelIndex* QStandardItemModel_Parent(const QStandardItemModel* self, QModelIndex* child);
int QStandardItemModel_RowCount(const QStandardItemModel* self);
int QStandardItemModel_ColumnCount(const QStandardItemModel* self);
bool QStandardItemModel_HasChildren(const QStandardItemModel* self);
QModelIndex* QStandardItemModel_Sibling(const QStandardItemModel* self, int row, int column, QModelIndex* idx);
QVariant* QStandardItemModel_Data(const QStandardItemModel* self, QModelIndex* index);
bool QStandardItemModel_SetData(QStandardItemModel* self, QModelIndex* index, QVariant* value);
bool QStandardItemModel_ClearItemData(QStandardItemModel* self, QModelIndex* index);
QVariant* QStandardItemModel_HeaderData(const QStandardItemModel* self, int section, uintptr_t orientation);
bool QStandardItemModel_SetHeaderData(QStandardItemModel* self, int section, uintptr_t orientation, QVariant* value);
bool QStandardItemModel_InsertRows(QStandardItemModel* self, int row, int count);
bool QStandardItemModel_InsertColumns(QStandardItemModel* self, int column, int count);
bool QStandardItemModel_RemoveRows(QStandardItemModel* self, int row, int count);
bool QStandardItemModel_RemoveColumns(QStandardItemModel* self, int column, int count);
int QStandardItemModel_Flags(const QStandardItemModel* self, QModelIndex* index);
int QStandardItemModel_SupportedDropActions(const QStandardItemModel* self);
void QStandardItemModel_Clear(QStandardItemModel* self);
void QStandardItemModel_Sort(QStandardItemModel* self, int column);
QStandardItem* QStandardItemModel_ItemFromIndex(const QStandardItemModel* self, QModelIndex* index);
QModelIndex* QStandardItemModel_IndexFromItem(const QStandardItemModel* self, QStandardItem* item);
QStandardItem* QStandardItemModel_Item(const QStandardItemModel* self, int row);
void QStandardItemModel_SetItem(QStandardItemModel* self, int row, int column, QStandardItem* item);
void QStandardItemModel_SetItem2(QStandardItemModel* self, int row, QStandardItem* item);
QStandardItem* QStandardItemModel_InvisibleRootItem(const QStandardItemModel* self);
QStandardItem* QStandardItemModel_HorizontalHeaderItem(const QStandardItemModel* self, int column);
void QStandardItemModel_SetHorizontalHeaderItem(QStandardItemModel* self, int column, QStandardItem* item);
QStandardItem* QStandardItemModel_VerticalHeaderItem(const QStandardItemModel* self, int row);
void QStandardItemModel_SetVerticalHeaderItem(QStandardItemModel* self, int row, QStandardItem* item);
void QStandardItemModel_SetHorizontalHeaderLabels(QStandardItemModel* self, char** labels, uint64_t* labels_Lengths, size_t labels_len);
void QStandardItemModel_SetVerticalHeaderLabels(QStandardItemModel* self, char** labels, uint64_t* labels_Lengths, size_t labels_len);
void QStandardItemModel_SetRowCount(QStandardItemModel* self, int rows);
void QStandardItemModel_SetColumnCount(QStandardItemModel* self, int columns);
void QStandardItemModel_AppendRow(QStandardItemModel* self, QStandardItem** items, size_t items_len);
void QStandardItemModel_AppendColumn(QStandardItemModel* self, QStandardItem** items, size_t items_len);
void QStandardItemModel_AppendRowWithItem(QStandardItemModel* self, QStandardItem* item);
void QStandardItemModel_InsertRow(QStandardItemModel* self, int row, QStandardItem** items, size_t items_len);
void QStandardItemModel_InsertColumn(QStandardItemModel* self, int column, QStandardItem** items, size_t items_len);
void QStandardItemModel_InsertRow2(QStandardItemModel* self, int row, QStandardItem* item);
bool QStandardItemModel_InsertRowWithRow(QStandardItemModel* self, int row);
bool QStandardItemModel_InsertColumnWithColumn(QStandardItemModel* self, int column);
QStandardItem* QStandardItemModel_TakeItem(QStandardItemModel* self, int row);
void QStandardItemModel_TakeRow(QStandardItemModel* self, int row, QStandardItem*** _out, size_t* _out_len);
void QStandardItemModel_TakeColumn(QStandardItemModel* self, int column, QStandardItem*** _out, size_t* _out_len);
QStandardItem* QStandardItemModel_TakeHorizontalHeaderItem(QStandardItemModel* self, int column);
QStandardItem* QStandardItemModel_TakeVerticalHeaderItem(QStandardItemModel* self, int row);
QStandardItem* QStandardItemModel_ItemPrototype(const QStandardItemModel* self);
void QStandardItemModel_SetItemPrototype(QStandardItemModel* self, QStandardItem* item);
void QStandardItemModel_FindItems(const QStandardItemModel* self, const char* text, size_t text_Strlen, QStandardItem*** _out, size_t* _out_len);
int QStandardItemModel_SortRole(const QStandardItemModel* self);
void QStandardItemModel_SetSortRole(QStandardItemModel* self, int role);
void QStandardItemModel_MimeTypes(const QStandardItemModel* self, char*** _out, int** _out_Lengths, size_t* _out_len);
QMimeData* QStandardItemModel_MimeData(const QStandardItemModel* self, QModelIndex** indexes, size_t indexes_len);
bool QStandardItemModel_DropMimeData(QStandardItemModel* self, QMimeData* data, uintptr_t action, int row, int column, QModelIndex* parent);
void QStandardItemModel_ItemChanged(QStandardItemModel* self, QStandardItem* item);
void QStandardItemModel_connect_ItemChanged(QStandardItemModel* self, void* slot);
void QStandardItemModel_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen);
void QStandardItemModel_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QStandardItemModel_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen);
void QStandardItemModel_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
QModelIndex* QStandardItemModel_Index3(const QStandardItemModel* self, int row, int column, QModelIndex* parent);
int QStandardItemModel_RowCount1(const QStandardItemModel* self, QModelIndex* parent);
int QStandardItemModel_ColumnCount1(const QStandardItemModel* self, QModelIndex* parent);
bool QStandardItemModel_HasChildren1(const QStandardItemModel* self, QModelIndex* parent);
QVariant* QStandardItemModel_Data2(const QStandardItemModel* self, QModelIndex* index, int role);
bool QStandardItemModel_SetData3(QStandardItemModel* self, QModelIndex* index, QVariant* value, int role);
QVariant* QStandardItemModel_HeaderData3(const QStandardItemModel* self, int section, uintptr_t orientation, int role);
bool QStandardItemModel_SetHeaderData4(QStandardItemModel* self, int section, uintptr_t orientation, QVariant* value, int role);
bool QStandardItemModel_InsertRows3(QStandardItemModel* self, int row, int count, QModelIndex* parent);
bool QStandardItemModel_InsertColumns3(QStandardItemModel* self, int column, int count, QModelIndex* parent);
bool QStandardItemModel_RemoveRows3(QStandardItemModel* self, int row, int count, QModelIndex* parent);
bool QStandardItemModel_RemoveColumns3(QStandardItemModel* self, int column, int count, QModelIndex* parent);
void QStandardItemModel_Sort2(QStandardItemModel* self, int column, uintptr_t order);
QStandardItem* QStandardItemModel_Item2(const QStandardItemModel* self, int row, int column);
bool QStandardItemModel_InsertRow22(QStandardItemModel* self, int row, QModelIndex* parent);
bool QStandardItemModel_InsertColumn2(QStandardItemModel* self, int column, QModelIndex* parent);
QStandardItem* QStandardItemModel_TakeItem2(QStandardItemModel* self, int row, int column);
void QStandardItemModel_FindItems2(const QStandardItemModel* self, const char* text, size_t text_Strlen, int flags, QStandardItem*** _out, size_t* _out_len);
void QStandardItemModel_FindItems3(const QStandardItemModel* self, const char* text, size_t text_Strlen, int flags, int column, QStandardItem*** _out, size_t* _out_len);
void QStandardItemModel_Delete(QStandardItemModel* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
