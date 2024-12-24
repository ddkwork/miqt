#pragma once
#ifndef MIQT_QT_GEN_QITEMDELEGATE_H
#define MIQT_QT_GEN_QITEMDELEGATE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractItemDelegate QAbstractItemDelegate;
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QEvent QEvent;
typedef struct QItemDelegate QItemDelegate;
typedef struct QItemEditorFactory QItemEditorFactory;
typedef struct QMetaObject QMetaObject;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QPainter QPainter;
typedef struct QPixmap QPixmap;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QStyleOptionViewItem QStyleOptionViewItem;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QItemDelegate* QItemDelegate_new();
extern __declspec(dllexport) 
QItemDelegate* QItemDelegate_new2(QObject* parent);
extern __declspec(dllexport) 
void QItemDelegate_virtbase(QItemDelegate* src
, QAbstractItemDelegate** outptr_QAbstractItemDelegate
);
extern __declspec(dllexport) 
QMetaObject* QItemDelegate_MetaObject(const QItemDelegate* self);
extern __declspec(dllexport) 
void* QItemDelegate_Metacast(QItemDelegate* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QItemDelegate_Tr(const char* s);
extern __declspec(dllexport) 
bool QItemDelegate_HasClipping(const QItemDelegate* self);
extern __declspec(dllexport) 
void QItemDelegate_SetClipping(QItemDelegate* self, bool clip);
extern __declspec(dllexport) 
void QItemDelegate_Paint(const QItemDelegate* self, QPainter* painter, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) 
QSize* QItemDelegate_SizeHint(const QItemDelegate* self, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) 
QWidget* QItemDelegate_CreateEditor(const QItemDelegate* self, QWidget* parent, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) 
void QItemDelegate_SetEditorData(const QItemDelegate* self, QWidget* editor, QModelIndex* index);
extern __declspec(dllexport) 
void QItemDelegate_SetModelData(const QItemDelegate* self, QWidget* editor, QAbstractItemModel* model, QModelIndex* index);
extern __declspec(dllexport) 
void QItemDelegate_UpdateEditorGeometry(const QItemDelegate* self, QWidget* editor, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) 
QItemEditorFactory* QItemDelegate_ItemEditorFactory(const QItemDelegate* self);
extern __declspec(dllexport) 
void QItemDelegate_SetItemEditorFactory(QItemDelegate* self, QItemEditorFactory* factory);
extern __declspec(dllexport) 
void QItemDelegate_DrawDisplay(const QItemDelegate* self, QPainter* painter, QStyleOptionViewItem* option, QRect* rect, struct miqt_string text);
extern __declspec(dllexport) 
void QItemDelegate_DrawDecoration(const QItemDelegate* self, QPainter* painter, QStyleOptionViewItem* option, QRect* rect, QPixmap* pixmap);
extern __declspec(dllexport) 
void QItemDelegate_DrawFocus(const QItemDelegate* self, QPainter* painter, QStyleOptionViewItem* option, QRect* rect);
extern __declspec(dllexport) 
void QItemDelegate_DrawCheck(const QItemDelegate* self, QPainter* painter, QStyleOptionViewItem* option, QRect* rect, int state);
extern __declspec(dllexport) 
bool QItemDelegate_EventFilter(QItemDelegate* self, QObject* object, QEvent* event);
extern __declspec(dllexport) 
bool QItemDelegate_EditorEvent(QItemDelegate* self, QEvent* event, QAbstractItemModel* model, QStyleOptionViewItem* option, QModelIndex* index);
extern __declspec(dllexport) 
struct miqt_string QItemDelegate_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QItemDelegate_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QItemDelegate_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QItemDelegate_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QItemDelegate_override_virtual_Metacast(void* self, intptr_t slot);
void* QItemDelegate_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QItemDelegate_Delete(QItemDelegate* self, bool isSubclass);

}
