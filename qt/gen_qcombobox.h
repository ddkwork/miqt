#pragma once
#ifndef MIQT_QT_GEN_QCOMBOBOX_H
#define MIQT_QT_GEN_QCOMBOBOX_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemDelegate QAbstractItemDelegate;
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QAbstractItemView QAbstractItemView;
typedef struct QComboBox QComboBox;
typedef struct QCompleter QCompleter;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QHideEvent QHideEvent;
typedef struct QIcon QIcon;
typedef struct QInputMethodEvent QInputMethodEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QLineEdit QLineEdit;
typedef struct QMetaObject QMetaObject;
typedef struct QModelIndex QModelIndex;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QStyleOptionComboBox QStyleOptionComboBox;
typedef struct QValidator QValidator;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QComboBox* QComboBox_new(QWidget* parent);
extern __declspec(dllexport) 
QComboBox* QComboBox_new2();
extern __declspec(dllexport) 
void QComboBox_virtbase(QComboBox* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QComboBox_MetaObject(const QComboBox* self);
extern __declspec(dllexport) 
void* QComboBox_Metacast(QComboBox* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QComboBox_Tr(const char* s);
extern __declspec(dllexport) 
int QComboBox_MaxVisibleItems(const QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_SetMaxVisibleItems(QComboBox* self, int maxItems);
extern __declspec(dllexport) 
int QComboBox_Count(const QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_SetMaxCount(QComboBox* self, int max);
extern __declspec(dllexport) 
int QComboBox_MaxCount(const QComboBox* self);
extern __declspec(dllexport) 
bool QComboBox_DuplicatesEnabled(const QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_SetDuplicatesEnabled(QComboBox* self, bool enable);
extern __declspec(dllexport) 
void QComboBox_SetFrame(QComboBox* self, bool frame);
extern __declspec(dllexport) 
bool QComboBox_HasFrame(const QComboBox* self);
extern __declspec(dllexport) 
int QComboBox_FindText(const QComboBox* self, struct miqt_string text);
extern __declspec(dllexport) 
int QComboBox_FindData(const QComboBox* self, QVariant* data);
extern __declspec(dllexport) 
InsertPolicy QComboBox_InsertPolicy(const QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_SetInsertPolicy(QComboBox* self, InsertPolicy policy);
extern __declspec(dllexport) 
SizeAdjustPolicy QComboBox_SizeAdjustPolicy(const QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_SetSizeAdjustPolicy(QComboBox* self, SizeAdjustPolicy policy);
extern __declspec(dllexport) 
int QComboBox_MinimumContentsLength(const QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_SetMinimumContentsLength(QComboBox* self, int characters);
extern __declspec(dllexport) 
QSize* QComboBox_IconSize(const QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_SetIconSize(QComboBox* self, QSize* size);
extern __declspec(dllexport) 
void QComboBox_SetPlaceholderText(QComboBox* self, struct miqt_string placeholderText);
extern __declspec(dllexport) 
struct miqt_string QComboBox_PlaceholderText(const QComboBox* self);
extern __declspec(dllexport) 
bool QComboBox_IsEditable(const QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_SetEditable(QComboBox* self, bool editable);
extern __declspec(dllexport) 
void QComboBox_SetLineEdit(QComboBox* self, QLineEdit* edit);
extern __declspec(dllexport) 
QLineEdit* QComboBox_LineEdit(const QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_SetValidator(QComboBox* self, QValidator* v);
extern __declspec(dllexport) 
QValidator* QComboBox_Validator(const QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_SetCompleter(QComboBox* self, QCompleter* c);
extern __declspec(dllexport) 
QCompleter* QComboBox_Completer(const QComboBox* self);
extern __declspec(dllexport) 
QAbstractItemDelegate* QComboBox_ItemDelegate(const QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_SetItemDelegate(QComboBox* self, QAbstractItemDelegate* delegate);
extern __declspec(dllexport) 
QAbstractItemModel* QComboBox_Model(const QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_SetModel(QComboBox* self, QAbstractItemModel* model);
extern __declspec(dllexport) 
QModelIndex* QComboBox_RootModelIndex(const QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_SetRootModelIndex(QComboBox* self, QModelIndex* index);
extern __declspec(dllexport) 
int QComboBox_ModelColumn(const QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_SetModelColumn(QComboBox* self, int visibleColumn);
extern __declspec(dllexport) 
LabelDrawingMode QComboBox_LabelDrawingMode(const QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_SetLabelDrawingMode(QComboBox* self, LabelDrawingMode labelDrawing);
extern __declspec(dllexport) 
int QComboBox_CurrentIndex(const QComboBox* self);
extern __declspec(dllexport) 
struct miqt_string QComboBox_CurrentText(const QComboBox* self);
extern __declspec(dllexport) 
QVariant* QComboBox_CurrentData(const QComboBox* self);
extern __declspec(dllexport) 
struct miqt_string QComboBox_ItemText(const QComboBox* self, int index);
extern __declspec(dllexport) 
QIcon* QComboBox_ItemIcon(const QComboBox* self, int index);
extern __declspec(dllexport) 
QVariant* QComboBox_ItemData(const QComboBox* self, int index);
extern __declspec(dllexport) 
void QComboBox_AddItem(QComboBox* self, struct miqt_string text);
extern __declspec(dllexport) 
void QComboBox_AddItem2(QComboBox* self, QIcon* icon, struct miqt_string text);
extern __declspec(dllexport) 
void QComboBox_AddItems(QComboBox* self, struct miqt_array /* of struct miqt_string */  texts);
extern __declspec(dllexport) 
void QComboBox_InsertItem(QComboBox* self, int index, struct miqt_string text);
extern __declspec(dllexport) 
void QComboBox_InsertItem2(QComboBox* self, int index, QIcon* icon, struct miqt_string text);
extern __declspec(dllexport) 
void QComboBox_InsertItems(QComboBox* self, int index, struct miqt_array /* of struct miqt_string */  texts);
extern __declspec(dllexport) 
void QComboBox_InsertSeparator(QComboBox* self, int index);
extern __declspec(dllexport) 
void QComboBox_RemoveItem(QComboBox* self, int index);
extern __declspec(dllexport) 
void QComboBox_SetItemText(QComboBox* self, int index, struct miqt_string text);
extern __declspec(dllexport) 
void QComboBox_SetItemIcon(QComboBox* self, int index, QIcon* icon);
extern __declspec(dllexport) 
void QComboBox_SetItemData(QComboBox* self, int index, QVariant* value);
extern __declspec(dllexport) 
QAbstractItemView* QComboBox_View(const QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_SetView(QComboBox* self, QAbstractItemView* itemView);
extern __declspec(dllexport) 
QSize* QComboBox_SizeHint(const QComboBox* self);
extern __declspec(dllexport) 
QSize* QComboBox_MinimumSizeHint(const QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_ShowPopup(QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_HidePopup(QComboBox* self);
extern __declspec(dllexport) 
bool QComboBox_Event(QComboBox* self, QEvent* event);
extern __declspec(dllexport) 
QVariant* QComboBox_InputMethodQuery(const QComboBox* self, int param1);
extern __declspec(dllexport) 
QVariant* QComboBox_InputMethodQuery2(const QComboBox* self, int query, QVariant* argument);
extern __declspec(dllexport) 
void QComboBox_Clear(QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_ClearEditText(QComboBox* self);
extern __declspec(dllexport) 
void QComboBox_SetEditText(QComboBox* self, struct miqt_string text);
extern __declspec(dllexport) 
void QComboBox_SetCurrentIndex(QComboBox* self, int index);
extern __declspec(dllexport) 
void QComboBox_SetCurrentText(QComboBox* self, struct miqt_string text);
extern __declspec(dllexport) 
void QComboBox_EditTextChanged(QComboBox* self, struct miqt_string param1);
void QComboBox_connect_EditTextChanged(QComboBox* self, intptr_t slot);
extern __declspec(dllexport) 
void QComboBox_Activated(QComboBox* self, int index);
void QComboBox_connect_Activated(QComboBox* self, intptr_t slot);
extern __declspec(dllexport) 
void QComboBox_TextActivated(QComboBox* self, struct miqt_string param1);
void QComboBox_connect_TextActivated(QComboBox* self, intptr_t slot);
extern __declspec(dllexport) 
void QComboBox_Highlighted(QComboBox* self, int index);
void QComboBox_connect_Highlighted(QComboBox* self, intptr_t slot);
extern __declspec(dllexport) 
void QComboBox_TextHighlighted(QComboBox* self, struct miqt_string param1);
void QComboBox_connect_TextHighlighted(QComboBox* self, intptr_t slot);
extern __declspec(dllexport) 
void QComboBox_CurrentIndexChanged(QComboBox* self, int index);
void QComboBox_connect_CurrentIndexChanged(QComboBox* self, intptr_t slot);
extern __declspec(dllexport) 
void QComboBox_CurrentTextChanged(QComboBox* self, struct miqt_string param1);
void QComboBox_connect_CurrentTextChanged(QComboBox* self, intptr_t slot);
extern __declspec(dllexport) 
void QComboBox_FocusInEvent(QComboBox* self, QFocusEvent* e);
extern __declspec(dllexport) 
void QComboBox_FocusOutEvent(QComboBox* self, QFocusEvent* e);
extern __declspec(dllexport) 
void QComboBox_ChangeEvent(QComboBox* self, QEvent* e);
extern __declspec(dllexport) 
void QComboBox_ResizeEvent(QComboBox* self, QResizeEvent* e);
extern __declspec(dllexport) 
void QComboBox_PaintEvent(QComboBox* self, QPaintEvent* e);
extern __declspec(dllexport) 
void QComboBox_ShowEvent(QComboBox* self, QShowEvent* e);
extern __declspec(dllexport) 
void QComboBox_HideEvent(QComboBox* self, QHideEvent* e);
extern __declspec(dllexport) 
void QComboBox_MousePressEvent(QComboBox* self, QMouseEvent* e);
extern __declspec(dllexport) 
void QComboBox_MouseReleaseEvent(QComboBox* self, QMouseEvent* e);
extern __declspec(dllexport) 
void QComboBox_KeyPressEvent(QComboBox* self, QKeyEvent* e);
extern __declspec(dllexport) 
void QComboBox_KeyReleaseEvent(QComboBox* self, QKeyEvent* e);
extern __declspec(dllexport) 
void QComboBox_WheelEvent(QComboBox* self, QWheelEvent* e);
extern __declspec(dllexport) 
void QComboBox_ContextMenuEvent(QComboBox* self, QContextMenuEvent* e);
extern __declspec(dllexport) 
void QComboBox_InputMethodEvent(QComboBox* self, QInputMethodEvent* param1);
extern __declspec(dllexport) 
void QComboBox_InitStyleOption(const QComboBox* self, QStyleOptionComboBox* option);
extern __declspec(dllexport) 
struct miqt_string QComboBox_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QComboBox_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
int QComboBox_FindText2(const QComboBox* self, struct miqt_string text, int flags);
extern __declspec(dllexport) 
int QComboBox_FindData2(const QComboBox* self, QVariant* data, int role);
extern __declspec(dllexport) 
int QComboBox_FindData3(const QComboBox* self, QVariant* data, int role, int flags);
extern __declspec(dllexport) 
QVariant* QComboBox_CurrentData1(const QComboBox* self, int role);
extern __declspec(dllexport) 
QVariant* QComboBox_ItemData2(const QComboBox* self, int index, int role);
extern __declspec(dllexport) 
void QComboBox_AddItem22(QComboBox* self, struct miqt_string text, QVariant* userData);
extern __declspec(dllexport) 
void QComboBox_AddItem3(QComboBox* self, QIcon* icon, struct miqt_string text, QVariant* userData);
extern __declspec(dllexport) 
void QComboBox_InsertItem3(QComboBox* self, int index, struct miqt_string text, QVariant* userData);
extern __declspec(dllexport) 
void QComboBox_InsertItem4(QComboBox* self, int index, QIcon* icon, struct miqt_string text, QVariant* userData);
extern __declspec(dllexport) 
void QComboBox_SetItemData3(QComboBox* self, int index, QVariant* value, int role);
extern __declspec(dllexport) 
void QComboBox_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QComboBox_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QComboBox_override_virtual_Metacast(void* self, intptr_t slot);
void* QComboBox_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QComboBox_Delete(QComboBox* self, bool isSubclass);

}
