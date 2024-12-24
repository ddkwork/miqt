#pragma once
#ifndef MIQT_QT_GEN_QLAYOUTITEM_H
#define MIQT_QT_GEN_QLAYOUTITEM_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QLayout QLayout;
typedef struct QLayoutItem QLayoutItem;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QSizePolicy QSizePolicy;
typedef struct QSpacerItem QSpacerItem;
typedef struct QWidget QWidget;
typedef struct QWidgetItem QWidgetItem;
typedef struct QWidgetItemV2 QWidgetItemV2;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QLayoutItem* QLayoutItem_new();
extern __declspec(dllexport) 
QLayoutItem* QLayoutItem_new2(QLayoutItem* param1);
extern __declspec(dllexport) 
QLayoutItem* QLayoutItem_new3(int alignment);
extern __declspec(dllexport) 
QSize* QLayoutItem_SizeHint(const QLayoutItem* self);
extern __declspec(dllexport) 
QSize* QLayoutItem_MinimumSize(const QLayoutItem* self);
extern __declspec(dllexport) 
QSize* QLayoutItem_MaximumSize(const QLayoutItem* self);
extern __declspec(dllexport) 
int QLayoutItem_ExpandingDirections(const QLayoutItem* self);
extern __declspec(dllexport) 
void QLayoutItem_SetGeometry(QLayoutItem* self, QRect* geometry);
extern __declspec(dllexport) 
QRect* QLayoutItem_Geometry(const QLayoutItem* self);
extern __declspec(dllexport) 
bool QLayoutItem_IsEmpty(const QLayoutItem* self);
extern __declspec(dllexport) 
bool QLayoutItem_HasHeightForWidth(const QLayoutItem* self);
extern __declspec(dllexport) 
int QLayoutItem_HeightForWidth(const QLayoutItem* self, int param1);
extern __declspec(dllexport) 
int QLayoutItem_MinimumHeightForWidth(const QLayoutItem* self, int param1);
extern __declspec(dllexport) 
void QLayoutItem_Invalidate(QLayoutItem* self);
extern __declspec(dllexport) 
QWidget* QLayoutItem_Widget(const QLayoutItem* self);
extern __declspec(dllexport) 
QLayout* QLayoutItem_Layout(QLayoutItem* self);
extern __declspec(dllexport) 
QSpacerItem* QLayoutItem_SpacerItem(QLayoutItem* self);
extern __declspec(dllexport) 
int QLayoutItem_Alignment(const QLayoutItem* self);
extern __declspec(dllexport) 
void QLayoutItem_SetAlignment(QLayoutItem* self, int a);
extern __declspec(dllexport) 
int QLayoutItem_ControlTypes(const QLayoutItem* self);
extern __declspec(dllexport) 
void QLayoutItem_Delete(QLayoutItem* self, bool isSubclass);

extern __declspec(dllexport) 
QSpacerItem* QSpacerItem_new(int w, int h);
extern __declspec(dllexport) 
QSpacerItem* QSpacerItem_new2(QSpacerItem* param1);
extern __declspec(dllexport) 
QSpacerItem* QSpacerItem_new3(int w, int h, int hData);
extern __declspec(dllexport) 
QSpacerItem* QSpacerItem_new4(int w, int h, int hData, int vData);
extern __declspec(dllexport) 
void QSpacerItem_virtbase(QSpacerItem* src
, QLayoutItem** outptr_QLayoutItem
);
extern __declspec(dllexport) 
void QSpacerItem_ChangeSize(QSpacerItem* self, int w, int h);
extern __declspec(dllexport) 
QSize* QSpacerItem_SizeHint(const QSpacerItem* self);
extern __declspec(dllexport) 
QSize* QSpacerItem_MinimumSize(const QSpacerItem* self);
extern __declspec(dllexport) 
QSize* QSpacerItem_MaximumSize(const QSpacerItem* self);
extern __declspec(dllexport) 
int QSpacerItem_ExpandingDirections(const QSpacerItem* self);
extern __declspec(dllexport) 
bool QSpacerItem_IsEmpty(const QSpacerItem* self);
extern __declspec(dllexport) 
void QSpacerItem_SetGeometry(QSpacerItem* self, QRect* geometry);
extern __declspec(dllexport) 
QRect* QSpacerItem_Geometry(const QSpacerItem* self);
extern __declspec(dllexport) 
QSpacerItem* QSpacerItem_SpacerItem(QSpacerItem* self);
extern __declspec(dllexport) 
QSizePolicy* QSpacerItem_SizePolicy(const QSpacerItem* self);
extern __declspec(dllexport) 
void QSpacerItem_ChangeSize3(QSpacerItem* self, int w, int h, int hData);
extern __declspec(dllexport) 
void QSpacerItem_ChangeSize4(QSpacerItem* self, int w, int h, int hData, int vData);
extern __declspec(dllexport) 
void QSpacerItem_Delete(QSpacerItem* self, bool isSubclass);

extern __declspec(dllexport) 
QWidgetItem* QWidgetItem_new(QWidget* w);
extern __declspec(dllexport) 
void QWidgetItem_virtbase(QWidgetItem* src
, QLayoutItem** outptr_QLayoutItem
);
extern __declspec(dllexport) 
QSize* QWidgetItem_SizeHint(const QWidgetItem* self);
extern __declspec(dllexport) 
QSize* QWidgetItem_MinimumSize(const QWidgetItem* self);
extern __declspec(dllexport) 
QSize* QWidgetItem_MaximumSize(const QWidgetItem* self);
extern __declspec(dllexport) 
int QWidgetItem_ExpandingDirections(const QWidgetItem* self);
extern __declspec(dllexport) 
bool QWidgetItem_IsEmpty(const QWidgetItem* self);
extern __declspec(dllexport) 
void QWidgetItem_SetGeometry(QWidgetItem* self, QRect* geometry);
extern __declspec(dllexport) 
QRect* QWidgetItem_Geometry(const QWidgetItem* self);
extern __declspec(dllexport) 
QWidget* QWidgetItem_Widget(const QWidgetItem* self);
extern __declspec(dllexport) 
bool QWidgetItem_HasHeightForWidth(const QWidgetItem* self);
extern __declspec(dllexport) 
int QWidgetItem_HeightForWidth(const QWidgetItem* self, int param1);
extern __declspec(dllexport) 
int QWidgetItem_MinimumHeightForWidth(const QWidgetItem* self, int param1);
extern __declspec(dllexport) 
int QWidgetItem_ControlTypes(const QWidgetItem* self);
extern __declspec(dllexport) 
void QWidgetItem_Delete(QWidgetItem* self, bool isSubclass);

extern __declspec(dllexport) 
QWidgetItemV2* QWidgetItemV2_new(QWidget* widget);
extern __declspec(dllexport) 
void QWidgetItemV2_virtbase(QWidgetItemV2* src
, QWidgetItem** outptr_QWidgetItem
);
extern __declspec(dllexport) 
QSize* QWidgetItemV2_SizeHint(const QWidgetItemV2* self);
extern __declspec(dllexport) 
QSize* QWidgetItemV2_MinimumSize(const QWidgetItemV2* self);
extern __declspec(dllexport) 
QSize* QWidgetItemV2_MaximumSize(const QWidgetItemV2* self);
extern __declspec(dllexport) 
int QWidgetItemV2_HeightForWidth(const QWidgetItemV2* self, int width);
extern __declspec(dllexport) 
void QWidgetItemV2_Delete(QWidgetItemV2* self, bool isSubclass);

}
