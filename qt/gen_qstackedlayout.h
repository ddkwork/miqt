#pragma once
#ifndef MIQT_QT_GEN_QSTACKEDLAYOUT_H
#define MIQT_QT_GEN_QSTACKEDLAYOUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QLayout QLayout;
typedef struct QLayoutItem QLayoutItem;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QStackedLayout QStackedLayout;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QStackedLayout* QStackedLayout_new(QWidget* parent);
extern __declspec(dllexport) 
QStackedLayout* QStackedLayout_new2();
extern __declspec(dllexport) 
QStackedLayout* QStackedLayout_new3(QLayout* parentLayout);
extern __declspec(dllexport) 
void QStackedLayout_virtbase(QStackedLayout* src
, QLayout** outptr_QLayout
);
extern __declspec(dllexport) 
QMetaObject* QStackedLayout_MetaObject(const QStackedLayout* self);
extern __declspec(dllexport) 
void* QStackedLayout_Metacast(QStackedLayout* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QStackedLayout_Tr(const char* s);
extern __declspec(dllexport) 
int QStackedLayout_AddWidget(QStackedLayout* self, QWidget* w);
extern __declspec(dllexport) 
int QStackedLayout_InsertWidget(QStackedLayout* self, int index, QWidget* w);
extern __declspec(dllexport) 
QWidget* QStackedLayout_CurrentWidget(const QStackedLayout* self);
extern __declspec(dllexport) 
int QStackedLayout_CurrentIndex(const QStackedLayout* self);
extern __declspec(dllexport) 
QWidget* QStackedLayout_Widget(const QStackedLayout* self, int param1);
extern __declspec(dllexport) 
int QStackedLayout_Count(const QStackedLayout* self);
extern __declspec(dllexport) 
StackingMode QStackedLayout_StackingMode(const QStackedLayout* self);
extern __declspec(dllexport) 
void QStackedLayout_SetStackingMode(QStackedLayout* self, StackingMode stackingMode);
extern __declspec(dllexport) 
void QStackedLayout_AddItem(QStackedLayout* self, QLayoutItem* item);
extern __declspec(dllexport) 
QSize* QStackedLayout_SizeHint(const QStackedLayout* self);
extern __declspec(dllexport) 
QSize* QStackedLayout_MinimumSize(const QStackedLayout* self);
extern __declspec(dllexport) 
QLayoutItem* QStackedLayout_ItemAt(const QStackedLayout* self, int param1);
extern __declspec(dllexport) 
QLayoutItem* QStackedLayout_TakeAt(QStackedLayout* self, int param1);
extern __declspec(dllexport) 
void QStackedLayout_SetGeometry(QStackedLayout* self, QRect* rect);
extern __declspec(dllexport) 
bool QStackedLayout_HasHeightForWidth(const QStackedLayout* self);
extern __declspec(dllexport) 
int QStackedLayout_HeightForWidth(const QStackedLayout* self, int width);
extern __declspec(dllexport) 
void QStackedLayout_WidgetRemoved(QStackedLayout* self, int index);
void QStackedLayout_connect_WidgetRemoved(QStackedLayout* self, intptr_t slot);
extern __declspec(dllexport) 
void QStackedLayout_CurrentChanged(QStackedLayout* self, int index);
void QStackedLayout_connect_CurrentChanged(QStackedLayout* self, intptr_t slot);
extern __declspec(dllexport) 
void QStackedLayout_WidgetAdded(QStackedLayout* self, int index);
void QStackedLayout_connect_WidgetAdded(QStackedLayout* self, intptr_t slot);
extern __declspec(dllexport) 
void QStackedLayout_SetCurrentIndex(QStackedLayout* self, int index);
extern __declspec(dllexport) 
void QStackedLayout_SetCurrentWidget(QStackedLayout* self, QWidget* w);
extern __declspec(dllexport) 
struct miqt_string QStackedLayout_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QStackedLayout_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QStackedLayout_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QStackedLayout_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QStackedLayout_override_virtual_Metacast(void* self, intptr_t slot);
void* QStackedLayout_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QStackedLayout_Delete(QStackedLayout* self, bool isSubclass);

}
