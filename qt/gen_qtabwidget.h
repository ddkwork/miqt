#pragma once
#ifndef MIQT_QT_GEN_QTABWIDGET_H
#define MIQT_QT_GEN_QTABWIDGET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QIcon QIcon;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QStyleOptionTabWidgetFrame QStyleOptionTabWidgetFrame;
typedef struct QTabBar QTabBar;
typedef struct QTabWidget QTabWidget;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QTabWidget* QTabWidget_new(QWidget* parent);
extern __declspec(dllexport) 
QTabWidget* QTabWidget_new2();
extern __declspec(dllexport) 
void QTabWidget_virtbase(QTabWidget* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QTabWidget_MetaObject(const QTabWidget* self);
extern __declspec(dllexport) 
void* QTabWidget_Metacast(QTabWidget* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QTabWidget_Tr(const char* s);
extern __declspec(dllexport) 
int QTabWidget_AddTab(QTabWidget* self, QWidget* widget, struct miqt_string param2);
extern __declspec(dllexport) 
int QTabWidget_AddTab2(QTabWidget* self, QWidget* widget, QIcon* icon, struct miqt_string label);
extern __declspec(dllexport) 
int QTabWidget_InsertTab(QTabWidget* self, int index, QWidget* widget, struct miqt_string param3);
extern __declspec(dllexport) 
int QTabWidget_InsertTab2(QTabWidget* self, int index, QWidget* widget, QIcon* icon, struct miqt_string label);
extern __declspec(dllexport) 
void QTabWidget_RemoveTab(QTabWidget* self, int index);
extern __declspec(dllexport) 
bool QTabWidget_IsTabEnabled(const QTabWidget* self, int index);
extern __declspec(dllexport) 
void QTabWidget_SetTabEnabled(QTabWidget* self, int index, bool enabled);
extern __declspec(dllexport) 
bool QTabWidget_IsTabVisible(const QTabWidget* self, int index);
extern __declspec(dllexport) 
void QTabWidget_SetTabVisible(QTabWidget* self, int index, bool visible);
extern __declspec(dllexport) 
struct miqt_string QTabWidget_TabText(const QTabWidget* self, int index);
extern __declspec(dllexport) 
void QTabWidget_SetTabText(QTabWidget* self, int index, struct miqt_string text);
extern __declspec(dllexport) 
QIcon* QTabWidget_TabIcon(const QTabWidget* self, int index);
extern __declspec(dllexport) 
void QTabWidget_SetTabIcon(QTabWidget* self, int index, QIcon* icon);
extern __declspec(dllexport) 
void QTabWidget_SetTabToolTip(QTabWidget* self, int index, struct miqt_string tip);
extern __declspec(dllexport) 
struct miqt_string QTabWidget_TabToolTip(const QTabWidget* self, int index);
extern __declspec(dllexport) 
void QTabWidget_SetTabWhatsThis(QTabWidget* self, int index, struct miqt_string text);
extern __declspec(dllexport) 
struct miqt_string QTabWidget_TabWhatsThis(const QTabWidget* self, int index);
extern __declspec(dllexport) 
int QTabWidget_CurrentIndex(const QTabWidget* self);
extern __declspec(dllexport) 
QWidget* QTabWidget_CurrentWidget(const QTabWidget* self);
extern __declspec(dllexport) 
QWidget* QTabWidget_Widget(const QTabWidget* self, int index);
extern __declspec(dllexport) 
int QTabWidget_IndexOf(const QTabWidget* self, QWidget* widget);
extern __declspec(dllexport) 
int QTabWidget_Count(const QTabWidget* self);
extern __declspec(dllexport) 
TabPosition QTabWidget_TabPosition(const QTabWidget* self);
extern __declspec(dllexport) 
void QTabWidget_SetTabPosition(QTabWidget* self, TabPosition position);
extern __declspec(dllexport) 
bool QTabWidget_TabsClosable(const QTabWidget* self);
extern __declspec(dllexport) 
void QTabWidget_SetTabsClosable(QTabWidget* self, bool closeable);
extern __declspec(dllexport) 
bool QTabWidget_IsMovable(const QTabWidget* self);
extern __declspec(dllexport) 
void QTabWidget_SetMovable(QTabWidget* self, bool movable);
extern __declspec(dllexport) 
TabShape QTabWidget_TabShape(const QTabWidget* self);
extern __declspec(dllexport) 
void QTabWidget_SetTabShape(QTabWidget* self, TabShape s);
extern __declspec(dllexport) 
QSize* QTabWidget_SizeHint(const QTabWidget* self);
extern __declspec(dllexport) 
QSize* QTabWidget_MinimumSizeHint(const QTabWidget* self);
extern __declspec(dllexport) 
int QTabWidget_HeightForWidth(const QTabWidget* self, int width);
extern __declspec(dllexport) 
bool QTabWidget_HasHeightForWidth(const QTabWidget* self);
extern __declspec(dllexport) 
void QTabWidget_SetCornerWidget(QTabWidget* self, QWidget* w);
extern __declspec(dllexport) 
QWidget* QTabWidget_CornerWidget(const QTabWidget* self);
extern __declspec(dllexport) 
int QTabWidget_ElideMode(const QTabWidget* self);
extern __declspec(dllexport) 
void QTabWidget_SetElideMode(QTabWidget* self, int mode);
extern __declspec(dllexport) 
QSize* QTabWidget_IconSize(const QTabWidget* self);
extern __declspec(dllexport) 
void QTabWidget_SetIconSize(QTabWidget* self, QSize* size);
extern __declspec(dllexport) 
bool QTabWidget_UsesScrollButtons(const QTabWidget* self);
extern __declspec(dllexport) 
void QTabWidget_SetUsesScrollButtons(QTabWidget* self, bool useButtons);
extern __declspec(dllexport) 
bool QTabWidget_DocumentMode(const QTabWidget* self);
extern __declspec(dllexport) 
void QTabWidget_SetDocumentMode(QTabWidget* self, bool set);
extern __declspec(dllexport) 
bool QTabWidget_TabBarAutoHide(const QTabWidget* self);
extern __declspec(dllexport) 
void QTabWidget_SetTabBarAutoHide(QTabWidget* self, bool enabled);
extern __declspec(dllexport) 
void QTabWidget_Clear(QTabWidget* self);
extern __declspec(dllexport) 
QTabBar* QTabWidget_TabBar(const QTabWidget* self);
extern __declspec(dllexport) 
void QTabWidget_SetCurrentIndex(QTabWidget* self, int index);
extern __declspec(dllexport) 
void QTabWidget_SetCurrentWidget(QTabWidget* self, QWidget* widget);
extern __declspec(dllexport) 
void QTabWidget_CurrentChanged(QTabWidget* self, int index);
void QTabWidget_connect_CurrentChanged(QTabWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QTabWidget_TabCloseRequested(QTabWidget* self, int index);
void QTabWidget_connect_TabCloseRequested(QTabWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QTabWidget_TabBarClicked(QTabWidget* self, int index);
void QTabWidget_connect_TabBarClicked(QTabWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QTabWidget_TabBarDoubleClicked(QTabWidget* self, int index);
void QTabWidget_connect_TabBarDoubleClicked(QTabWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QTabWidget_TabInserted(QTabWidget* self, int index);
extern __declspec(dllexport) 
void QTabWidget_TabRemoved(QTabWidget* self, int index);
extern __declspec(dllexport) 
void QTabWidget_ShowEvent(QTabWidget* self, QShowEvent* param1);
extern __declspec(dllexport) 
void QTabWidget_ResizeEvent(QTabWidget* self, QResizeEvent* param1);
extern __declspec(dllexport) 
void QTabWidget_KeyPressEvent(QTabWidget* self, QKeyEvent* param1);
extern __declspec(dllexport) 
void QTabWidget_PaintEvent(QTabWidget* self, QPaintEvent* param1);
extern __declspec(dllexport) 
void QTabWidget_ChangeEvent(QTabWidget* self, QEvent* param1);
extern __declspec(dllexport) 
bool QTabWidget_Event(QTabWidget* self, QEvent* param1);
extern __declspec(dllexport) 
void QTabWidget_InitStyleOption(const QTabWidget* self, QStyleOptionTabWidgetFrame* option);
extern __declspec(dllexport) 
struct miqt_string QTabWidget_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QTabWidget_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QTabWidget_SetCornerWidget2(QTabWidget* self, QWidget* w, int corner);
extern __declspec(dllexport) 
QWidget* QTabWidget_CornerWidget1(const QTabWidget* self, int corner);
extern __declspec(dllexport) 
void QTabWidget_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QTabWidget_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QTabWidget_override_virtual_Metacast(void* self, intptr_t slot);
void* QTabWidget_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QTabWidget_Delete(QTabWidget* self, bool isSubclass);

}
