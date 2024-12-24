#pragma once
#ifndef MIQT_QT_GEN_QSTYLEDITEMDELEGATE_H
#define MIQT_QT_GEN_QSTYLEDITEMDELEGATE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemDelegate QAbstractItemDelegate;
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QEvent QEvent;
typedef struct QItemEditorFactory QItemEditorFactory;
typedef struct QLocale QLocale;
typedef struct QMetaObject QMetaObject;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QPainter QPainter;
typedef struct QSize QSize;
typedef struct QStyleOptionViewItem QStyleOptionViewItem;
typedef struct QStyledItemDelegate QStyledItemDelegate;
typedef struct QVariant QVariant;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QStyledItemDelegate* QStyledItemDelegate_new();
extern __declspec(dllexport) 
QStyledItemDelegate* QStyledItemDelegate_new2(QObject* parent);
extern __declspec(dllexport) 
void QStyledItemDelegate_virtbase(QStyledItemDelegate* src
, QAbstractItemDelegate** outptr_QAbstractItemDelegate
);
extern __declspec(dllexport) 
QMetaObject* QStyledItemDelegate_MetaObject(const QStyledItemDelegate* self);
extern __declspec(dllexport) 
void* QStyledItemDelegate_Metacast(QStyledItemDelegate* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QStyledItemDelegate_Tr(const char* s);
extern __declspec(dllexport) 
void QStyledItemDelegate_Paint(const QStyledItemDelegate* self, QPainter* painter, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) 
QSize* QStyledItemDelegate_SizeHint(const QStyledItemDelegate* self, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) 
QWidget* QStyledItemDelegate_CreateEditor(const QStyledItemDelegate* self, QWidget* parent, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) 
void QStyledItemDelegate_SetEditorData(const QStyledItemDelegate* self, QWidget* editor, QModelIndex* index);
extern __declspec(dllexport) 
void QStyledItemDelegate_SetModelData(const QStyledItemDelegate* self, QWidget* editor, QAbstractItemModel* model, QModelIndex* index);
extern __declspec(dllexport) 
void QStyledItemDelegate_UpdateEditorGeometry(const QStyledItemDelegate* self, QWidget* editor, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) 
QItemEditorFactory* QStyledItemDelegate_ItemEditorFactory(const QStyledItemDelegate* self);
extern __declspec(dllexport) 
void QStyledItemDelegate_SetItemEditorFactory(QStyledItemDelegate* self, QItemEditorFactory* factory);
extern __declspec(dllexport) 
struct miqt_string QStyledItemDelegate_DisplayText(const QStyledItemDelegate* self, QVariant* value, QLocale* locale);
extern __declspec(dllexport) 
void QStyledItemDelegate_InitStyleOption(const QStyledItemDelegate* self, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) 
bool QStyledItemDelegate_EventFilter(QStyledItemDelegate* self, QObject* object, QEvent* event);
extern __declspec(dllexport) 
bool QStyledItemDelegate_EditorEvent(QStyledItemDelegate* self, QEvent* event, QAbstractItemModel* model, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) 
struct miqt_string QStyledItemDelegate_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QStyledItemDelegate_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QStyledItemDelegate_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QStyledItemDelegate_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QStyledItemDelegate_override_virtual_Metacast(void* self, intptr_t slot);
void* QStyledItemDelegate_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QStyledItemDelegate_Delete(QStyledItemDelegate* self, bool isSubclass);

}
