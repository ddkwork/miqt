#pragma once
#ifndef MIQT_QT_GEN_QTREEWIDGET_H
#define MIQT_QT_GEN_QTREEWIDGET_H

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
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QTreeView QTreeView;
typedef struct QTreeWidget QTreeWidget;
typedef struct QTreeWidgetItem QTreeWidgetItem;
typedef struct QVariant QVariant;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_new();
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_new2(struct miqt_array /* of struct miqt_string */  strings);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_new3(QTreeWidget* treeview);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_new4(QTreeWidget* treeview, struct miqt_array /* of struct miqt_string */  strings);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_new5(QTreeWidget* treeview, QTreeWidgetItem* after);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_new6(QTreeWidgetItem* parent);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_new7(QTreeWidgetItem* parent, struct miqt_array /* of struct miqt_string */  strings);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_new8(QTreeWidgetItem* parent, QTreeWidgetItem* after);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_new9(QTreeWidgetItem* other);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_new10(int typeVal);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_new11(struct miqt_array /* of struct miqt_string */  strings, int typeVal);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_new12(QTreeWidget* treeview, int typeVal);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_new13(QTreeWidget* treeview, struct miqt_array /* of struct miqt_string */  strings, int typeVal);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_new14(QTreeWidget* treeview, QTreeWidgetItem* after, int typeVal);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_new15(QTreeWidgetItem* parent, int typeVal);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_new16(QTreeWidgetItem* parent, struct miqt_array /* of struct miqt_string */  strings, int typeVal);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_new17(QTreeWidgetItem* parent, QTreeWidgetItem* after, int typeVal);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_Clone(const QTreeWidgetItem* self);
extern __declspec(dllexport) 
QTreeWidget* QTreeWidgetItem_TreeWidget(const QTreeWidgetItem* self);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetSelected(QTreeWidgetItem* self, bool selectVal);
extern __declspec(dllexport) 
bool QTreeWidgetItem_IsSelected(const QTreeWidgetItem* self);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetHidden(QTreeWidgetItem* self, bool hide);
extern __declspec(dllexport) 
bool QTreeWidgetItem_IsHidden(const QTreeWidgetItem* self);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetExpanded(QTreeWidgetItem* self, bool expand);
extern __declspec(dllexport) 
bool QTreeWidgetItem_IsExpanded(const QTreeWidgetItem* self);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetFirstColumnSpanned(QTreeWidgetItem* self, bool span);
extern __declspec(dllexport) 
bool QTreeWidgetItem_IsFirstColumnSpanned(const QTreeWidgetItem* self);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetDisabled(QTreeWidgetItem* self, bool disabled);
extern __declspec(dllexport) 
bool QTreeWidgetItem_IsDisabled(const QTreeWidgetItem* self);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetChildIndicatorPolicy(QTreeWidgetItem* self, int policy);
extern __declspec(dllexport) 
int QTreeWidgetItem_ChildIndicatorPolicy(const QTreeWidgetItem* self);
extern __declspec(dllexport) 
int QTreeWidgetItem_Flags(const QTreeWidgetItem* self);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetFlags(QTreeWidgetItem* self, int flags);
extern __declspec(dllexport) 
struct miqt_string QTreeWidgetItem_Text(const QTreeWidgetItem* self, int column);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetText(QTreeWidgetItem* self, int column, struct miqt_string text);
extern __declspec(dllexport) 
QIcon* QTreeWidgetItem_Icon(const QTreeWidgetItem* self, int column);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetIcon(QTreeWidgetItem* self, int column, QIcon* icon);
extern __declspec(dllexport) 
struct miqt_string QTreeWidgetItem_StatusTip(const QTreeWidgetItem* self, int column);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetStatusTip(QTreeWidgetItem* self, int column, struct miqt_string statusTip);
extern __declspec(dllexport) 
struct miqt_string QTreeWidgetItem_ToolTip(const QTreeWidgetItem* self, int column);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetToolTip(QTreeWidgetItem* self, int column, struct miqt_string toolTip);
extern __declspec(dllexport) 
struct miqt_string QTreeWidgetItem_WhatsThis(const QTreeWidgetItem* self, int column);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetWhatsThis(QTreeWidgetItem* self, int column, struct miqt_string whatsThis);
extern __declspec(dllexport) 
QFont* QTreeWidgetItem_Font(const QTreeWidgetItem* self, int column);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetFont(QTreeWidgetItem* self, int column, QFont* font);
extern __declspec(dllexport) 
int QTreeWidgetItem_TextAlignment(const QTreeWidgetItem* self, int column);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetTextAlignment(QTreeWidgetItem* self, int column, int alignment);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetTextAlignment2(QTreeWidgetItem* self, int column, int alignment);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetTextAlignment3(QTreeWidgetItem* self, int column, int alignment);
extern __declspec(dllexport) 
QBrush* QTreeWidgetItem_Background(const QTreeWidgetItem* self, int column);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetBackground(QTreeWidgetItem* self, int column, QBrush* brush);
extern __declspec(dllexport) 
QBrush* QTreeWidgetItem_Foreground(const QTreeWidgetItem* self, int column);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetForeground(QTreeWidgetItem* self, int column, QBrush* brush);
extern __declspec(dllexport) 
int QTreeWidgetItem_CheckState(const QTreeWidgetItem* self, int column);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetCheckState(QTreeWidgetItem* self, int column, int state);
extern __declspec(dllexport) 
QSize* QTreeWidgetItem_SizeHint(const QTreeWidgetItem* self, int column);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetSizeHint(QTreeWidgetItem* self, int column, QSize* size);
extern __declspec(dllexport) 
QVariant* QTreeWidgetItem_Data(const QTreeWidgetItem* self, int column, int role);
extern __declspec(dllexport) 
void QTreeWidgetItem_SetData(QTreeWidgetItem* self, int column, int role, QVariant* value);
extern __declspec(dllexport) 
bool QTreeWidgetItem_OperatorLesser(const QTreeWidgetItem* self, QTreeWidgetItem* other);
extern __declspec(dllexport) 
void QTreeWidgetItem_Read(QTreeWidgetItem* self, QDataStream* in);
extern __declspec(dllexport) 
void QTreeWidgetItem_Write(const QTreeWidgetItem* self, QDataStream* out);
extern __declspec(dllexport) 
void QTreeWidgetItem_OperatorAssign(QTreeWidgetItem* self, QTreeWidgetItem* other);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_Parent(const QTreeWidgetItem* self);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_Child(const QTreeWidgetItem* self, int index);
extern __declspec(dllexport) 
int QTreeWidgetItem_ChildCount(const QTreeWidgetItem* self);
extern __declspec(dllexport) 
int QTreeWidgetItem_ColumnCount(const QTreeWidgetItem* self);
extern __declspec(dllexport) 
int QTreeWidgetItem_IndexOfChild(const QTreeWidgetItem* self, QTreeWidgetItem* child);
extern __declspec(dllexport) 
void QTreeWidgetItem_AddChild(QTreeWidgetItem* self, QTreeWidgetItem* child);
extern __declspec(dllexport) 
void QTreeWidgetItem_InsertChild(QTreeWidgetItem* self, int index, QTreeWidgetItem* child);
extern __declspec(dllexport) 
void QTreeWidgetItem_RemoveChild(QTreeWidgetItem* self, QTreeWidgetItem* child);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidgetItem_TakeChild(QTreeWidgetItem* self, int index);
extern __declspec(dllexport) 
void QTreeWidgetItem_AddChildren(QTreeWidgetItem* self, struct miqt_array /* of QTreeWidgetItem* */  children);
extern __declspec(dllexport) 
void QTreeWidgetItem_InsertChildren(QTreeWidgetItem* self, int index, struct miqt_array /* of QTreeWidgetItem* */  children);
extern __declspec(dllexport) 
struct miqt_array /* of QTreeWidgetItem* */  QTreeWidgetItem_TakeChildren(QTreeWidgetItem* self);
extern __declspec(dllexport) 
int QTreeWidgetItem_Type(const QTreeWidgetItem* self);
extern __declspec(dllexport) 
void QTreeWidgetItem_SortChildren(QTreeWidgetItem* self, int column, int order);
extern __declspec(dllexport) 
void QTreeWidgetItem_Delete(QTreeWidgetItem* self, bool isSubclass);

extern __declspec(dllexport) 
QTreeWidget* QTreeWidget_new(QWidget* parent);
extern __declspec(dllexport) 
QTreeWidget* QTreeWidget_new2();
extern __declspec(dllexport) 
void QTreeWidget_virtbase(QTreeWidget* src
, QTreeView** outptr_QTreeView
);
extern __declspec(dllexport) 
QMetaObject* QTreeWidget_MetaObject(const QTreeWidget* self);
extern __declspec(dllexport) 
void* QTreeWidget_Metacast(QTreeWidget* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QTreeWidget_Tr(const char* s);
extern __declspec(dllexport) 
int QTreeWidget_ColumnCount(const QTreeWidget* self);
extern __declspec(dllexport) 
void QTreeWidget_SetColumnCount(QTreeWidget* self, int columns);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidget_InvisibleRootItem(const QTreeWidget* self);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidget_TopLevelItem(const QTreeWidget* self, int index);
extern __declspec(dllexport) 
int QTreeWidget_TopLevelItemCount(const QTreeWidget* self);
extern __declspec(dllexport) 
void QTreeWidget_InsertTopLevelItem(QTreeWidget* self, int index, QTreeWidgetItem* item);
extern __declspec(dllexport) 
void QTreeWidget_AddTopLevelItem(QTreeWidget* self, QTreeWidgetItem* item);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidget_TakeTopLevelItem(QTreeWidget* self, int index);
extern __declspec(dllexport) 
int QTreeWidget_IndexOfTopLevelItem(const QTreeWidget* self, QTreeWidgetItem* item);
extern __declspec(dllexport) 
void QTreeWidget_InsertTopLevelItems(QTreeWidget* self, int index, struct miqt_array /* of QTreeWidgetItem* */  items);
extern __declspec(dllexport) 
void QTreeWidget_AddTopLevelItems(QTreeWidget* self, struct miqt_array /* of QTreeWidgetItem* */  items);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidget_HeaderItem(const QTreeWidget* self);
extern __declspec(dllexport) 
void QTreeWidget_SetHeaderItem(QTreeWidget* self, QTreeWidgetItem* item);
extern __declspec(dllexport) 
void QTreeWidget_SetHeaderLabels(QTreeWidget* self, struct miqt_array /* of struct miqt_string */  labels);
extern __declspec(dllexport) 
void QTreeWidget_SetHeaderLabel(QTreeWidget* self, struct miqt_string label);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidget_CurrentItem(const QTreeWidget* self);
extern __declspec(dllexport) 
int QTreeWidget_CurrentColumn(const QTreeWidget* self);
extern __declspec(dllexport) 
void QTreeWidget_SetCurrentItem(QTreeWidget* self, QTreeWidgetItem* item);
extern __declspec(dllexport) 
void QTreeWidget_SetCurrentItem2(QTreeWidget* self, QTreeWidgetItem* item, int column);
extern __declspec(dllexport) 
void QTreeWidget_SetCurrentItem3(QTreeWidget* self, QTreeWidgetItem* item, int column, int command);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidget_ItemAt(const QTreeWidget* self, QPoint* p);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidget_ItemAt2(const QTreeWidget* self, int x, int y);
extern __declspec(dllexport) 
QRect* QTreeWidget_VisualItemRect(const QTreeWidget* self, QTreeWidgetItem* item);
extern __declspec(dllexport) 
int QTreeWidget_SortColumn(const QTreeWidget* self);
extern __declspec(dllexport) 
void QTreeWidget_SortItems(QTreeWidget* self, int column, int order);
extern __declspec(dllexport) 
void QTreeWidget_EditItem(QTreeWidget* self, QTreeWidgetItem* item);
extern __declspec(dllexport) 
void QTreeWidget_OpenPersistentEditor(QTreeWidget* self, QTreeWidgetItem* item);
extern __declspec(dllexport) 
void QTreeWidget_ClosePersistentEditor(QTreeWidget* self, QTreeWidgetItem* item);
extern __declspec(dllexport) 
bool QTreeWidget_IsPersistentEditorOpen(const QTreeWidget* self, QTreeWidgetItem* item);
extern __declspec(dllexport) 
QWidget* QTreeWidget_ItemWidget(const QTreeWidget* self, QTreeWidgetItem* item, int column);
extern __declspec(dllexport) 
void QTreeWidget_SetItemWidget(QTreeWidget* self, QTreeWidgetItem* item, int column, QWidget* widget);
extern __declspec(dllexport) 
void QTreeWidget_RemoveItemWidget(QTreeWidget* self, QTreeWidgetItem* item, int column);
extern __declspec(dllexport) 
struct miqt_array /* of QTreeWidgetItem* */  QTreeWidget_SelectedItems(const QTreeWidget* self);
extern __declspec(dllexport) 
struct miqt_array /* of QTreeWidgetItem* */  QTreeWidget_FindItems(const QTreeWidget* self, struct miqt_string text, int flags);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidget_ItemAbove(const QTreeWidget* self, QTreeWidgetItem* item);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidget_ItemBelow(const QTreeWidget* self, QTreeWidgetItem* item);
extern __declspec(dllexport) 
QModelIndex* QTreeWidget_IndexFromItem(const QTreeWidget* self, QTreeWidgetItem* item);
extern __declspec(dllexport) 
QTreeWidgetItem* QTreeWidget_ItemFromIndex(const QTreeWidget* self, QModelIndex* index);
extern __declspec(dllexport) 
void QTreeWidget_SetSelectionModel(QTreeWidget* self, QItemSelectionModel* selectionModel);
extern __declspec(dllexport) 
void QTreeWidget_ScrollToItem(QTreeWidget* self, QTreeWidgetItem* item);
extern __declspec(dllexport) 
void QTreeWidget_ExpandItem(QTreeWidget* self, QTreeWidgetItem* item);
extern __declspec(dllexport) 
void QTreeWidget_CollapseItem(QTreeWidget* self, QTreeWidgetItem* item);
extern __declspec(dllexport) 
void QTreeWidget_Clear(QTreeWidget* self);
extern __declspec(dllexport) 
void QTreeWidget_ItemPressed(QTreeWidget* self, QTreeWidgetItem* item, int column);
void QTreeWidget_connect_ItemPressed(QTreeWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QTreeWidget_ItemClicked(QTreeWidget* self, QTreeWidgetItem* item, int column);
void QTreeWidget_connect_ItemClicked(QTreeWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QTreeWidget_ItemDoubleClicked(QTreeWidget* self, QTreeWidgetItem* item, int column);
void QTreeWidget_connect_ItemDoubleClicked(QTreeWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QTreeWidget_ItemActivated(QTreeWidget* self, QTreeWidgetItem* item, int column);
void QTreeWidget_connect_ItemActivated(QTreeWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QTreeWidget_ItemEntered(QTreeWidget* self, QTreeWidgetItem* item, int column);
void QTreeWidget_connect_ItemEntered(QTreeWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QTreeWidget_ItemChanged(QTreeWidget* self, QTreeWidgetItem* item, int column);
void QTreeWidget_connect_ItemChanged(QTreeWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QTreeWidget_ItemExpanded(QTreeWidget* self, QTreeWidgetItem* item);
void QTreeWidget_connect_ItemExpanded(QTreeWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QTreeWidget_ItemCollapsed(QTreeWidget* self, QTreeWidgetItem* item);
void QTreeWidget_connect_ItemCollapsed(QTreeWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QTreeWidget_CurrentItemChanged(QTreeWidget* self, QTreeWidgetItem* current, QTreeWidgetItem* previous);
void QTreeWidget_connect_CurrentItemChanged(QTreeWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QTreeWidget_ItemSelectionChanged(QTreeWidget* self);
void QTreeWidget_connect_ItemSelectionChanged(QTreeWidget* self, intptr_t slot);
extern __declspec(dllexport) 
bool QTreeWidget_Event(QTreeWidget* self, QEvent* e);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QTreeWidget_MimeTypes(const QTreeWidget* self);
extern __declspec(dllexport) 
QMimeData* QTreeWidget_MimeData(const QTreeWidget* self, struct miqt_array /* of QTreeWidgetItem* */  items);
extern __declspec(dllexport) 
bool QTreeWidget_DropMimeData(QTreeWidget* self, QTreeWidgetItem* parent, int index, QMimeData* data, int action);
extern __declspec(dllexport) 
int QTreeWidget_SupportedDropActions(const QTreeWidget* self);
extern __declspec(dllexport) 
void QTreeWidget_DropEvent(QTreeWidget* self, QDropEvent* event);
extern __declspec(dllexport) 
struct miqt_string QTreeWidget_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QTreeWidget_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QTreeWidget_EditItem2(QTreeWidget* self, QTreeWidgetItem* item, int column);
extern __declspec(dllexport) 
void QTreeWidget_OpenPersistentEditor2(QTreeWidget* self, QTreeWidgetItem* item, int column);
extern __declspec(dllexport) 
void QTreeWidget_ClosePersistentEditor2(QTreeWidget* self, QTreeWidgetItem* item, int column);
extern __declspec(dllexport) 
bool QTreeWidget_IsPersistentEditorOpen2(const QTreeWidget* self, QTreeWidgetItem* item, int column);
extern __declspec(dllexport) 
struct miqt_array /* of QTreeWidgetItem* */  QTreeWidget_FindItems3(const QTreeWidget* self, struct miqt_string text, int flags, int column);
extern __declspec(dllexport) 
QModelIndex* QTreeWidget_IndexFromItem2(const QTreeWidget* self, QTreeWidgetItem* item, int column);
extern __declspec(dllexport) 
void QTreeWidget_ScrollToItem2(QTreeWidget* self, QTreeWidgetItem* item, int hint);
extern __declspec(dllexport) 
void QTreeWidget_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QTreeWidget_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QTreeWidget_override_virtual_Metacast(void* self, intptr_t slot);
void* QTreeWidget_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QTreeWidget_Delete(QTreeWidget* self, bool isSubclass);

}
