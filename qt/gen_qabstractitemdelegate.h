#pragma once
#ifndef MIQT_QT_GEN_QABSTRACTITEMDELEGATE_H
#define MIQT_QT_GEN_QABSTRACTITEMDELEGATE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemDelegate QAbstractItemDelegate;
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QAbstractItemView QAbstractItemView;
typedef struct QEvent QEvent;
typedef struct QHelpEvent QHelpEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QPainter QPainter;
typedef struct QSize QSize;
typedef struct QStyleOptionViewItem QStyleOptionViewItem;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAbstractItemDelegate* QAbstractItemDelegate_new();
extern __declspec(dllexport) 
QAbstractItemDelegate* QAbstractItemDelegate_new2(QObject* parent);
extern __declspec(dllexport) 
void QAbstractItemDelegate_virtbase(QAbstractItemDelegate* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QAbstractItemDelegate_MetaObject(const QAbstractItemDelegate* self);
extern __declspec(dllexport) 
void* QAbstractItemDelegate_Metacast(QAbstractItemDelegate* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAbstractItemDelegate_Tr(const char* s);
extern __declspec(dllexport) 
void QAbstractItemDelegate_Paint(const QAbstractItemDelegate* self, QPainter* painter, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) 
QSize* QAbstractItemDelegate_SizeHint(const QAbstractItemDelegate* self, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) 
QWidget* QAbstractItemDelegate_CreateEditor(const QAbstractItemDelegate* self, QWidget* parent, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) 
void QAbstractItemDelegate_DestroyEditor(const QAbstractItemDelegate* self, QWidget* editor, QModelIndex* index);
extern __declspec(dllexport) 
void QAbstractItemDelegate_SetEditorData(const QAbstractItemDelegate* self, QWidget* editor, QModelIndex* index);
extern __declspec(dllexport) 
void QAbstractItemDelegate_SetModelData(const QAbstractItemDelegate* self, QWidget* editor, QAbstractItemModel* model, QModelIndex* index);
extern __declspec(dllexport) 
void QAbstractItemDelegate_UpdateEditorGeometry(const QAbstractItemDelegate* self, QWidget* editor, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) 
bool QAbstractItemDelegate_EditorEvent(QAbstractItemDelegate* self, QEvent* event, QAbstractItemModel* model, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) 
bool QAbstractItemDelegate_HelpEvent(QAbstractItemDelegate* self, QHelpEvent* event, QAbstractItemView* view, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) 
struct miqt_array /* of int */  QAbstractItemDelegate_PaintingRoles(const QAbstractItemDelegate* self);
extern __declspec(dllexport) 
void QAbstractItemDelegate_CommitData(QAbstractItemDelegate* self, QWidget* editor);
void QAbstractItemDelegate_connect_CommitData(QAbstractItemDelegate* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractItemDelegate_CloseEditor(QAbstractItemDelegate* self, QWidget* editor);
void QAbstractItemDelegate_connect_CloseEditor(QAbstractItemDelegate* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractItemDelegate_SizeHintChanged(QAbstractItemDelegate* self, QModelIndex* param1);
void QAbstractItemDelegate_connect_SizeHintChanged(QAbstractItemDelegate* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QAbstractItemDelegate_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAbstractItemDelegate_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAbstractItemDelegate_CloseEditor2(QAbstractItemDelegate* self, QWidget* editor, int hint);
void QAbstractItemDelegate_connect_CloseEditor2(QAbstractItemDelegate* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractItemDelegate_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAbstractItemDelegate_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAbstractItemDelegate_override_virtual_Metacast(void* self, intptr_t slot);
void* QAbstractItemDelegate_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAbstractItemDelegate_Delete(QAbstractItemDelegate* self, bool isSubclass);

}
