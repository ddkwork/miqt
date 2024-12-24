#pragma once
#ifndef MIQT_QT_GEN_QSTYLEDITEMDELEGATE_H
#define MIQT_QT_GEN_QSTYLEDITEMDELEGATE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractItemDelegate;
class QAbstractItemModel;
class QAbstractItemView;
class QEvent;
class QHelpEvent;
class QItemEditorFactory;
class QLocale;
class QMetaObject;
class QModelIndex;
class QObject;
class QPainter;
class QSize;
class QStyleOptionViewItem;
class QStyledItemDelegate;
class QVariant;
class QWidget;
class _GUID;
class type_info;
#else
typedef struct QAbstractItemDelegate QAbstractItemDelegate;
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QAbstractItemView QAbstractItemView;
typedef struct QEvent QEvent;
typedef struct QHelpEvent QHelpEvent;
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
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QStyledItemDelegate* QStyledItemDelegate_new();
extern __declspec(dllexport) QStyledItemDelegate* QStyledItemDelegate_new2(QObject* parent);
extern __declspec(dllexport) void QStyledItemDelegate_virtbase(QStyledItemDelegate* src, QAbstractItemDelegate** outptr_QAbstractItemDelegate);
extern __declspec(dllexport) QMetaObject* QStyledItemDelegate_MetaObject(const QStyledItemDelegate* self);
extern __declspec(dllexport) void* QStyledItemDelegate_Metacast(QStyledItemDelegate* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QStyledItemDelegate_Tr(const char* s);
extern __declspec(dllexport) void QStyledItemDelegate_Paint(const QStyledItemDelegate* self, QPainter* painter, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) QSize* QStyledItemDelegate_SizeHint(const QStyledItemDelegate* self, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) QWidget* QStyledItemDelegate_CreateEditor(const QStyledItemDelegate* self, QWidget* parent, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) void QStyledItemDelegate_SetEditorData(const QStyledItemDelegate* self, QWidget* editor, QModelIndex* index);
extern __declspec(dllexport) void QStyledItemDelegate_SetModelData(const QStyledItemDelegate* self, QWidget* editor, QAbstractItemModel* model, QModelIndex* index);
extern __declspec(dllexport) void QStyledItemDelegate_UpdateEditorGeometry(const QStyledItemDelegate* self, QWidget* editor, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) QItemEditorFactory* QStyledItemDelegate_ItemEditorFactory(const QStyledItemDelegate* self);
extern __declspec(dllexport) void QStyledItemDelegate_SetItemEditorFactory(QStyledItemDelegate* self, QItemEditorFactory* factory);
extern __declspec(dllexport) struct miqt_string QStyledItemDelegate_DisplayText(const QStyledItemDelegate* self, QVariant* value, QLocale* locale);
extern __declspec(dllexport) void QStyledItemDelegate_InitStyleOption(const QStyledItemDelegate* self, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) bool QStyledItemDelegate_EventFilter(QStyledItemDelegate* self, QObject* object, QEvent* event);
extern __declspec(dllexport) bool QStyledItemDelegate_EditorEvent(QStyledItemDelegate* self, QEvent* event, QAbstractItemModel* model, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) struct miqt_string QStyledItemDelegate_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QStyledItemDelegate_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QStyledItemDelegate_override_virtual_Paint(void* self, intptr_t slot);
void QStyledItemDelegate_virtualbase_Paint(const void* self, QPainter* painter, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) void QStyledItemDelegate_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QStyledItemDelegate_virtualbase_SizeHint(const void* self, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) void QStyledItemDelegate_override_virtual_CreateEditor(void* self, intptr_t slot);
QWidget* QStyledItemDelegate_virtualbase_CreateEditor(const void* self, QWidget* parent, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) void QStyledItemDelegate_override_virtual_SetEditorData(void* self, intptr_t slot);
void QStyledItemDelegate_virtualbase_SetEditorData(const void* self, QWidget* editor, QModelIndex* index);
extern __declspec(dllexport) void QStyledItemDelegate_override_virtual_SetModelData(void* self, intptr_t slot);
void QStyledItemDelegate_virtualbase_SetModelData(const void* self, QWidget* editor, QAbstractItemModel* model, QModelIndex* index);
extern __declspec(dllexport) void QStyledItemDelegate_override_virtual_UpdateEditorGeometry(void* self, intptr_t slot);
void QStyledItemDelegate_virtualbase_UpdateEditorGeometry(const void* self, QWidget* editor, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) void QStyledItemDelegate_override_virtual_DisplayText(void* self, intptr_t slot);
struct miqt_string QStyledItemDelegate_virtualbase_DisplayText(const void* self, QVariant* value, QLocale* locale);
extern __declspec(dllexport) void QStyledItemDelegate_override_virtual_InitStyleOption(void* self, intptr_t slot);
void QStyledItemDelegate_virtualbase_InitStyleOption(const void* self, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) void QStyledItemDelegate_override_virtual_EventFilter(void* self, intptr_t slot);
bool QStyledItemDelegate_virtualbase_EventFilter(void* self, QObject* object, QEvent* event);
extern __declspec(dllexport) void QStyledItemDelegate_override_virtual_EditorEvent(void* self, intptr_t slot);
bool QStyledItemDelegate_virtualbase_EditorEvent(void* self, QEvent* event, QAbstractItemModel* model, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) void QStyledItemDelegate_override_virtual_DestroyEditor(void* self, intptr_t slot);
void QStyledItemDelegate_virtualbase_DestroyEditor(const void* self, QWidget* editor, QModelIndex* index);
extern __declspec(dllexport) void QStyledItemDelegate_override_virtual_HelpEvent(void* self, intptr_t slot);
bool QStyledItemDelegate_virtualbase_HelpEvent(void* self, QHelpEvent* event, QAbstractItemView* view, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) void QStyledItemDelegate_override_virtual_PaintingRoles(void* self, intptr_t slot);
struct miqt_array /* of int */  QStyledItemDelegate_virtualbase_PaintingRoles(const void* self);
extern __declspec(dllexport) void QStyledItemDelegate_Delete(QStyledItemDelegate* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
