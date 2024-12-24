#pragma once
#ifndef MIQT_QT_GEN_QLISTWIDGET_H
#define MIQT_QT_GEN_QLISTWIDGET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemView QAbstractItemView;
typedef struct QAbstractScrollArea QAbstractScrollArea;
typedef struct QBrush QBrush;
typedef struct QDataStream QDataStream;
typedef struct QDropEvent QDropEvent;
typedef struct QEvent QEvent;
typedef struct QFont QFont;
typedef struct QFrame QFrame;
typedef struct QIcon QIcon;
typedef struct QItemSelectionModel QItemSelectionModel;
typedef struct QListView QListView;
typedef struct QListWidget QListWidget;
typedef struct QListWidgetItem QListWidgetItem;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QVariant QVariant;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QListWidgetItem* QListWidgetItem_new();
extern __declspec(dllexport) 
QListWidgetItem* QListWidgetItem_new2(struct miqt_string text);
extern __declspec(dllexport) 
QListWidgetItem* QListWidgetItem_new3(QIcon* icon, struct miqt_string text);
extern __declspec(dllexport) 
QListWidgetItem* QListWidgetItem_new4(QListWidgetItem* other);
extern __declspec(dllexport) 
QListWidgetItem* QListWidgetItem_new5(QListWidget* listview);
extern __declspec(dllexport) 
QListWidgetItem* QListWidgetItem_new6(QListWidget* listview, int typeVal);
extern __declspec(dllexport) 
QListWidgetItem* QListWidgetItem_new7(struct miqt_string text, QListWidget* listview);
extern __declspec(dllexport) 
QListWidgetItem* QListWidgetItem_new8(struct miqt_string text, QListWidget* listview, int typeVal);
extern __declspec(dllexport) 
QListWidgetItem* QListWidgetItem_new9(QIcon* icon, struct miqt_string text, QListWidget* listview);
extern __declspec(dllexport) 
QListWidgetItem* QListWidgetItem_new10(QIcon* icon, struct miqt_string text, QListWidget* listview, int typeVal);
extern __declspec(dllexport) 
QListWidgetItem* QListWidgetItem_Clone(const QListWidgetItem* self);
extern __declspec(dllexport) 
QListWidget* QListWidgetItem_ListWidget(const QListWidgetItem* self);
extern __declspec(dllexport) 
void QListWidgetItem_SetSelected(QListWidgetItem* self, bool selectVal);
extern __declspec(dllexport) 
bool QListWidgetItem_IsSelected(const QListWidgetItem* self);
extern __declspec(dllexport) 
void QListWidgetItem_SetHidden(QListWidgetItem* self, bool hide);
extern __declspec(dllexport) 
bool QListWidgetItem_IsHidden(const QListWidgetItem* self);
extern __declspec(dllexport) 
int QListWidgetItem_Flags(const QListWidgetItem* self);
extern __declspec(dllexport) 
void QListWidgetItem_SetFlags(QListWidgetItem* self, int flags);
extern __declspec(dllexport) 
struct miqt_string QListWidgetItem_Text(const QListWidgetItem* self);
extern __declspec(dllexport) 
void QListWidgetItem_SetText(QListWidgetItem* self, struct miqt_string text);
extern __declspec(dllexport) 
QIcon* QListWidgetItem_Icon(const QListWidgetItem* self);
extern __declspec(dllexport) 
void QListWidgetItem_SetIcon(QListWidgetItem* self, QIcon* icon);
extern __declspec(dllexport) 
struct miqt_string QListWidgetItem_StatusTip(const QListWidgetItem* self);
extern __declspec(dllexport) 
void QListWidgetItem_SetStatusTip(QListWidgetItem* self, struct miqt_string statusTip);
extern __declspec(dllexport) 
struct miqt_string QListWidgetItem_ToolTip(const QListWidgetItem* self);
extern __declspec(dllexport) 
void QListWidgetItem_SetToolTip(QListWidgetItem* self, struct miqt_string toolTip);
extern __declspec(dllexport) 
struct miqt_string QListWidgetItem_WhatsThis(const QListWidgetItem* self);
extern __declspec(dllexport) 
void QListWidgetItem_SetWhatsThis(QListWidgetItem* self, struct miqt_string whatsThis);
extern __declspec(dllexport) 
QFont* QListWidgetItem_Font(const QListWidgetItem* self);
extern __declspec(dllexport) 
void QListWidgetItem_SetFont(QListWidgetItem* self, QFont* font);
extern __declspec(dllexport) 
int QListWidgetItem_TextAlignment(const QListWidgetItem* self);
extern __declspec(dllexport) 
void QListWidgetItem_SetTextAlignment(QListWidgetItem* self, int alignment);
extern __declspec(dllexport) 
void QListWidgetItem_SetTextAlignmentWithAlignment(QListWidgetItem* self, int alignment);
extern __declspec(dllexport) 
void QListWidgetItem_SetTextAlignment2(QListWidgetItem* self, int alignment);
extern __declspec(dllexport) 
QBrush* QListWidgetItem_Background(const QListWidgetItem* self);
extern __declspec(dllexport) 
void QListWidgetItem_SetBackground(QListWidgetItem* self, QBrush* brush);
extern __declspec(dllexport) 
QBrush* QListWidgetItem_Foreground(const QListWidgetItem* self);
extern __declspec(dllexport) 
void QListWidgetItem_SetForeground(QListWidgetItem* self, QBrush* brush);
extern __declspec(dllexport) 
int QListWidgetItem_CheckState(const QListWidgetItem* self);
extern __declspec(dllexport) 
void QListWidgetItem_SetCheckState(QListWidgetItem* self, int state);
extern __declspec(dllexport) 
QSize* QListWidgetItem_SizeHint(const QListWidgetItem* self);
extern __declspec(dllexport) 
void QListWidgetItem_SetSizeHint(QListWidgetItem* self, QSize* size);
extern __declspec(dllexport) 
QVariant* QListWidgetItem_Data(const QListWidgetItem* self, int role);
extern __declspec(dllexport) 
void QListWidgetItem_SetData(QListWidgetItem* self, int role, QVariant* value);
extern __declspec(dllexport) 
bool QListWidgetItem_OperatorLesser(const QListWidgetItem* self, QListWidgetItem* other);
extern __declspec(dllexport) 
void QListWidgetItem_Read(QListWidgetItem* self, QDataStream* in);
extern __declspec(dllexport) 
void QListWidgetItem_Write(const QListWidgetItem* self, QDataStream* out);
extern __declspec(dllexport) 
void QListWidgetItem_OperatorAssign(QListWidgetItem* self, QListWidgetItem* other);
extern __declspec(dllexport) 
int QListWidgetItem_Type(const QListWidgetItem* self);
extern __declspec(dllexport) 
void QListWidgetItem_Delete(QListWidgetItem* self, bool isSubclass);

extern __declspec(dllexport) 
QListWidget* QListWidget_new(QWidget* parent);
extern __declspec(dllexport) 
QListWidget* QListWidget_new2();
extern __declspec(dllexport) 
void QListWidget_virtbase(QListWidget* src
, QListView** outptr_QListView
);
extern __declspec(dllexport) 
QMetaObject* QListWidget_MetaObject(const QListWidget* self);
extern __declspec(dllexport) 
void* QListWidget_Metacast(QListWidget* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QListWidget_Tr(const char* s);
extern __declspec(dllexport) 
void QListWidget_SetSelectionModel(QListWidget* self, QItemSelectionModel* selectionModel);
extern __declspec(dllexport) 
QListWidgetItem* QListWidget_Item(const QListWidget* self, int row);
extern __declspec(dllexport) 
int QListWidget_Row(const QListWidget* self, QListWidgetItem* item);
extern __declspec(dllexport) 
void QListWidget_InsertItem(QListWidget* self, int row, QListWidgetItem* item);
extern __declspec(dllexport) 
void QListWidget_InsertItem2(QListWidget* self, int row, struct miqt_string label);
extern __declspec(dllexport) 
void QListWidget_InsertItems(QListWidget* self, int row, struct miqt_array /* of struct miqt_string */  labels);
extern __declspec(dllexport) 
void QListWidget_AddItem(QListWidget* self, struct miqt_string label);
extern __declspec(dllexport) 
void QListWidget_AddItemWithItem(QListWidget* self, QListWidgetItem* item);
extern __declspec(dllexport) 
void QListWidget_AddItems(QListWidget* self, struct miqt_array /* of struct miqt_string */  labels);
extern __declspec(dllexport) 
QListWidgetItem* QListWidget_TakeItem(QListWidget* self, int row);
extern __declspec(dllexport) 
int QListWidget_Count(const QListWidget* self);
extern __declspec(dllexport) 
QListWidgetItem* QListWidget_CurrentItem(const QListWidget* self);
extern __declspec(dllexport) 
void QListWidget_SetCurrentItem(QListWidget* self, QListWidgetItem* item);
extern __declspec(dllexport) 
void QListWidget_SetCurrentItem2(QListWidget* self, QListWidgetItem* item, int command);
extern __declspec(dllexport) 
int QListWidget_CurrentRow(const QListWidget* self);
extern __declspec(dllexport) 
void QListWidget_SetCurrentRow(QListWidget* self, int row);
extern __declspec(dllexport) 
void QListWidget_SetCurrentRow2(QListWidget* self, int row, int command);
extern __declspec(dllexport) 
QListWidgetItem* QListWidget_ItemAt(const QListWidget* self, QPoint* p);
extern __declspec(dllexport) 
QListWidgetItem* QListWidget_ItemAt2(const QListWidget* self, int x, int y);
extern __declspec(dllexport) 
QRect* QListWidget_VisualItemRect(const QListWidget* self, QListWidgetItem* item);
extern __declspec(dllexport) 
void QListWidget_SortItems(QListWidget* self);
extern __declspec(dllexport) 
void QListWidget_SetSortingEnabled(QListWidget* self, bool enable);
extern __declspec(dllexport) 
bool QListWidget_IsSortingEnabled(const QListWidget* self);
extern __declspec(dllexport) 
void QListWidget_EditItem(QListWidget* self, QListWidgetItem* item);
extern __declspec(dllexport) 
void QListWidget_OpenPersistentEditor(QListWidget* self, QListWidgetItem* item);
extern __declspec(dllexport) 
void QListWidget_ClosePersistentEditor(QListWidget* self, QListWidgetItem* item);
extern __declspec(dllexport) 
bool QListWidget_IsPersistentEditorOpen(const QListWidget* self, QListWidgetItem* item);
extern __declspec(dllexport) 
QWidget* QListWidget_ItemWidget(const QListWidget* self, QListWidgetItem* item);
extern __declspec(dllexport) 
void QListWidget_SetItemWidget(QListWidget* self, QListWidgetItem* item, QWidget* widget);
extern __declspec(dllexport) 
void QListWidget_RemoveItemWidget(QListWidget* self, QListWidgetItem* item);
extern __declspec(dllexport) 
struct miqt_array /* of QListWidgetItem* */  QListWidget_SelectedItems(const QListWidget* self);
extern __declspec(dllexport) 
struct miqt_array /* of QListWidgetItem* */  QListWidget_FindItems(const QListWidget* self, struct miqt_string text, int flags);
extern __declspec(dllexport) 
struct miqt_array /* of QListWidgetItem* */  QListWidget_Items(const QListWidget* self, QMimeData* data);
extern __declspec(dllexport) 
QModelIndex* QListWidget_IndexFromItem(const QListWidget* self, QListWidgetItem* item);
extern __declspec(dllexport) 
QListWidgetItem* QListWidget_ItemFromIndex(const QListWidget* self, QModelIndex* index);
extern __declspec(dllexport) 
void QListWidget_DropEvent(QListWidget* self, QDropEvent* event);
extern __declspec(dllexport) 
void QListWidget_ScrollToItem(QListWidget* self, QListWidgetItem* item);
extern __declspec(dllexport) 
void QListWidget_Clear(QListWidget* self);
extern __declspec(dllexport) 
void QListWidget_ItemPressed(QListWidget* self, QListWidgetItem* item);
void QListWidget_connect_ItemPressed(QListWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QListWidget_ItemClicked(QListWidget* self, QListWidgetItem* item);
void QListWidget_connect_ItemClicked(QListWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QListWidget_ItemDoubleClicked(QListWidget* self, QListWidgetItem* item);
void QListWidget_connect_ItemDoubleClicked(QListWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QListWidget_ItemActivated(QListWidget* self, QListWidgetItem* item);
void QListWidget_connect_ItemActivated(QListWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QListWidget_ItemEntered(QListWidget* self, QListWidgetItem* item);
void QListWidget_connect_ItemEntered(QListWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QListWidget_ItemChanged(QListWidget* self, QListWidgetItem* item);
void QListWidget_connect_ItemChanged(QListWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QListWidget_CurrentItemChanged(QListWidget* self, QListWidgetItem* current, QListWidgetItem* previous);
void QListWidget_connect_CurrentItemChanged(QListWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QListWidget_CurrentTextChanged(QListWidget* self, struct miqt_string currentText);
void QListWidget_connect_CurrentTextChanged(QListWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QListWidget_CurrentRowChanged(QListWidget* self, int currentRow);
void QListWidget_connect_CurrentRowChanged(QListWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QListWidget_ItemSelectionChanged(QListWidget* self);
void QListWidget_connect_ItemSelectionChanged(QListWidget* self, intptr_t slot);
extern __declspec(dllexport) 
bool QListWidget_Event(QListWidget* self, QEvent* e);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QListWidget_MimeTypes(const QListWidget* self);
extern __declspec(dllexport) 
QMimeData* QListWidget_MimeData(const QListWidget* self, struct miqt_array /* of QListWidgetItem* */  items);
extern __declspec(dllexport) 
bool QListWidget_DropMimeData(QListWidget* self, int index, QMimeData* data, int action);
extern __declspec(dllexport) 
int QListWidget_SupportedDropActions(const QListWidget* self);
extern __declspec(dllexport) 
struct miqt_string QListWidget_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QListWidget_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QListWidget_SortItems1(QListWidget* self, int order);
extern __declspec(dllexport) 
void QListWidget_ScrollToItem2(QListWidget* self, QListWidgetItem* item, int hint);
extern __declspec(dllexport) 
void QListWidget_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QListWidget_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QListWidget_override_virtual_Metacast(void* self, intptr_t slot);
void* QListWidget_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QListWidget_Delete(QListWidget* self, bool isSubclass);

}
