#pragma once
#ifndef MIQT_QT_GEN_QBOXLAYOUT_H
#define MIQT_QT_GEN_QBOXLAYOUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QBoxLayout QBoxLayout;
typedef struct QHBoxLayout QHBoxLayout;
typedef struct QLayout QLayout;
typedef struct QLayoutItem QLayoutItem;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QSpacerItem QSpacerItem;
typedef struct QVBoxLayout QVBoxLayout;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QBoxLayout* QBoxLayout_new(Direction param1);
extern __declspec(dllexport) 
QBoxLayout* QBoxLayout_new2(Direction param1, QWidget* parent);
extern __declspec(dllexport) 
void QBoxLayout_virtbase(QBoxLayout* src
, QLayout** outptr_QLayout
);
extern __declspec(dllexport) 
QMetaObject* QBoxLayout_MetaObject(const QBoxLayout* self);
extern __declspec(dllexport) 
void* QBoxLayout_Metacast(QBoxLayout* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QBoxLayout_Tr(const char* s);
extern __declspec(dllexport) 
Direction QBoxLayout_Direction(const QBoxLayout* self);
extern __declspec(dllexport) 
void QBoxLayout_SetDirection(QBoxLayout* self, Direction direction);
extern __declspec(dllexport) 
void QBoxLayout_AddSpacing(QBoxLayout* self, int size);
extern __declspec(dllexport) 
void QBoxLayout_AddStretch(QBoxLayout* self);
extern __declspec(dllexport) 
void QBoxLayout_AddSpacerItem(QBoxLayout* self, QSpacerItem* spacerItem);
extern __declspec(dllexport) 
void QBoxLayout_AddWidget(QBoxLayout* self, QWidget* param1);
extern __declspec(dllexport) 
void QBoxLayout_AddLayout(QBoxLayout* self, QLayout* layout);
extern __declspec(dllexport) 
void QBoxLayout_AddStrut(QBoxLayout* self, int param1);
extern __declspec(dllexport) 
void QBoxLayout_AddItem(QBoxLayout* self, QLayoutItem* param1);
extern __declspec(dllexport) 
void QBoxLayout_InsertSpacing(QBoxLayout* self, int index, int size);
extern __declspec(dllexport) 
void QBoxLayout_InsertStretch(QBoxLayout* self, int index);
extern __declspec(dllexport) 
void QBoxLayout_InsertSpacerItem(QBoxLayout* self, int index, QSpacerItem* spacerItem);
extern __declspec(dllexport) 
void QBoxLayout_InsertWidget(QBoxLayout* self, int index, QWidget* widget);
extern __declspec(dllexport) 
void QBoxLayout_InsertLayout(QBoxLayout* self, int index, QLayout* layout);
extern __declspec(dllexport) 
void QBoxLayout_InsertItem(QBoxLayout* self, int index, QLayoutItem* param2);
extern __declspec(dllexport) 
int QBoxLayout_Spacing(const QBoxLayout* self);
extern __declspec(dllexport) 
void QBoxLayout_SetSpacing(QBoxLayout* self, int spacing);
extern __declspec(dllexport) 
bool QBoxLayout_SetStretchFactor(QBoxLayout* self, QWidget* w, int stretch);
extern __declspec(dllexport) 
bool QBoxLayout_SetStretchFactor2(QBoxLayout* self, QLayout* l, int stretch);
extern __declspec(dllexport) 
void QBoxLayout_SetStretch(QBoxLayout* self, int index, int stretch);
extern __declspec(dllexport) 
int QBoxLayout_Stretch(const QBoxLayout* self, int index);
extern __declspec(dllexport) 
QSize* QBoxLayout_SizeHint(const QBoxLayout* self);
extern __declspec(dllexport) 
QSize* QBoxLayout_MinimumSize(const QBoxLayout* self);
extern __declspec(dllexport) 
QSize* QBoxLayout_MaximumSize(const QBoxLayout* self);
extern __declspec(dllexport) 
bool QBoxLayout_HasHeightForWidth(const QBoxLayout* self);
extern __declspec(dllexport) 
int QBoxLayout_HeightForWidth(const QBoxLayout* self, int param1);
extern __declspec(dllexport) 
int QBoxLayout_MinimumHeightForWidth(const QBoxLayout* self, int param1);
extern __declspec(dllexport) 
int QBoxLayout_ExpandingDirections(const QBoxLayout* self);
extern __declspec(dllexport) 
void QBoxLayout_Invalidate(QBoxLayout* self);
extern __declspec(dllexport) 
QLayoutItem* QBoxLayout_ItemAt(const QBoxLayout* self, int param1);
extern __declspec(dllexport) 
QLayoutItem* QBoxLayout_TakeAt(QBoxLayout* self, int param1);
extern __declspec(dllexport) 
int QBoxLayout_Count(const QBoxLayout* self);
extern __declspec(dllexport) 
void QBoxLayout_SetGeometry(QBoxLayout* self, QRect* geometry);
extern __declspec(dllexport) 
struct miqt_string QBoxLayout_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QBoxLayout_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QBoxLayout_AddStretch1(QBoxLayout* self, int stretch);
extern __declspec(dllexport) 
void QBoxLayout_AddWidget2(QBoxLayout* self, QWidget* param1, int stretch);
extern __declspec(dllexport) 
void QBoxLayout_AddWidget3(QBoxLayout* self, QWidget* param1, int stretch, int alignment);
extern __declspec(dllexport) 
void QBoxLayout_AddLayout2(QBoxLayout* self, QLayout* layout, int stretch);
extern __declspec(dllexport) 
void QBoxLayout_InsertStretch2(QBoxLayout* self, int index, int stretch);
extern __declspec(dllexport) 
void QBoxLayout_InsertWidget3(QBoxLayout* self, int index, QWidget* widget, int stretch);
extern __declspec(dllexport) 
void QBoxLayout_InsertWidget4(QBoxLayout* self, int index, QWidget* widget, int stretch, int alignment);
extern __declspec(dllexport) 
void QBoxLayout_InsertLayout3(QBoxLayout* self, int index, QLayout* layout, int stretch);
extern __declspec(dllexport) 
void QBoxLayout_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QBoxLayout_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QBoxLayout_override_virtual_Metacast(void* self, intptr_t slot);
void* QBoxLayout_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QBoxLayout_Delete(QBoxLayout* self, bool isSubclass);

extern __declspec(dllexport) 
QHBoxLayout* QHBoxLayout_new(QWidget* parent);
extern __declspec(dllexport) 
QHBoxLayout* QHBoxLayout_new2();
extern __declspec(dllexport) 
void QHBoxLayout_virtbase(QHBoxLayout* src
, QBoxLayout** outptr_QBoxLayout
);
extern __declspec(dllexport) 
QMetaObject* QHBoxLayout_MetaObject(const QHBoxLayout* self);
extern __declspec(dllexport) 
void* QHBoxLayout_Metacast(QHBoxLayout* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QHBoxLayout_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_string QHBoxLayout_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QHBoxLayout_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QHBoxLayout_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QHBoxLayout_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QHBoxLayout_override_virtual_Metacast(void* self, intptr_t slot);
void* QHBoxLayout_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QHBoxLayout_Delete(QHBoxLayout* self, bool isSubclass);

extern __declspec(dllexport) 
QVBoxLayout* QVBoxLayout_new(QWidget* parent);
extern __declspec(dllexport) 
QVBoxLayout* QVBoxLayout_new2();
extern __declspec(dllexport) 
void QVBoxLayout_virtbase(QVBoxLayout* src
, QBoxLayout** outptr_QBoxLayout
);
extern __declspec(dllexport) 
QMetaObject* QVBoxLayout_MetaObject(const QVBoxLayout* self);
extern __declspec(dllexport) 
void* QVBoxLayout_Metacast(QVBoxLayout* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QVBoxLayout_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_string QVBoxLayout_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QVBoxLayout_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QVBoxLayout_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QVBoxLayout_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QVBoxLayout_override_virtual_Metacast(void* self, intptr_t slot);
void* QVBoxLayout_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QVBoxLayout_Delete(QVBoxLayout* self, bool isSubclass);

}
