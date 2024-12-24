#pragma once
#ifndef MIQT_QT_GEN_QTOOLBAR_H
#define MIQT_QT_GEN_QTOOLBAR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAction QAction;
typedef struct QActionEvent QActionEvent;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QStyleOptionToolBar QStyleOptionToolBar;
typedef struct QToolBar QToolBar;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QToolBar* QToolBar_new(QWidget* parent);
extern __declspec(dllexport) 
QToolBar* QToolBar_new2(struct miqt_string title);
extern __declspec(dllexport) 
QToolBar* QToolBar_new3();
extern __declspec(dllexport) 
QToolBar* QToolBar_new4(struct miqt_string title, QWidget* parent);
extern __declspec(dllexport) 
void QToolBar_virtbase(QToolBar* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QToolBar_MetaObject(const QToolBar* self);
extern __declspec(dllexport) 
void* QToolBar_Metacast(QToolBar* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QToolBar_Tr(const char* s);
extern __declspec(dllexport) 
void QToolBar_SetMovable(QToolBar* self, bool movable);
extern __declspec(dllexport) 
bool QToolBar_IsMovable(const QToolBar* self);
extern __declspec(dllexport) 
void QToolBar_SetAllowedAreas(QToolBar* self, int areas);
extern __declspec(dllexport) 
int QToolBar_AllowedAreas(const QToolBar* self);
extern __declspec(dllexport) 
bool QToolBar_IsAreaAllowed(const QToolBar* self, int area);
extern __declspec(dllexport) 
void QToolBar_SetOrientation(QToolBar* self, int orientation);
extern __declspec(dllexport) 
int QToolBar_Orientation(const QToolBar* self);
extern __declspec(dllexport) 
void QToolBar_Clear(QToolBar* self);
extern __declspec(dllexport) 
QAction* QToolBar_AddSeparator(QToolBar* self);
extern __declspec(dllexport) 
QAction* QToolBar_InsertSeparator(QToolBar* self, QAction* before);
extern __declspec(dllexport) 
QAction* QToolBar_AddWidget(QToolBar* self, QWidget* widget);
extern __declspec(dllexport) 
QAction* QToolBar_InsertWidget(QToolBar* self, QAction* before, QWidget* widget);
extern __declspec(dllexport) 
QRect* QToolBar_ActionGeometry(const QToolBar* self, QAction* action);
extern __declspec(dllexport) 
QAction* QToolBar_ActionAt(const QToolBar* self, QPoint* p);
extern __declspec(dllexport) 
QAction* QToolBar_ActionAt2(const QToolBar* self, int x, int y);
extern __declspec(dllexport) 
QAction* QToolBar_ToggleViewAction(const QToolBar* self);
extern __declspec(dllexport) 
QSize* QToolBar_IconSize(const QToolBar* self);
extern __declspec(dllexport) 
int QToolBar_ToolButtonStyle(const QToolBar* self);
extern __declspec(dllexport) 
QWidget* QToolBar_WidgetForAction(const QToolBar* self, QAction* action);
extern __declspec(dllexport) 
bool QToolBar_IsFloatable(const QToolBar* self);
extern __declspec(dllexport) 
void QToolBar_SetFloatable(QToolBar* self, bool floatable);
extern __declspec(dllexport) 
bool QToolBar_IsFloating(const QToolBar* self);
extern __declspec(dllexport) 
void QToolBar_SetIconSize(QToolBar* self, QSize* iconSize);
extern __declspec(dllexport) 
void QToolBar_SetToolButtonStyle(QToolBar* self, int toolButtonStyle);
extern __declspec(dllexport) 
void QToolBar_ActionTriggered(QToolBar* self, QAction* action);
void QToolBar_connect_ActionTriggered(QToolBar* self, intptr_t slot);
extern __declspec(dllexport) 
void QToolBar_MovableChanged(QToolBar* self, bool movable);
void QToolBar_connect_MovableChanged(QToolBar* self, intptr_t slot);
extern __declspec(dllexport) 
void QToolBar_AllowedAreasChanged(QToolBar* self, int allowedAreas);
void QToolBar_connect_AllowedAreasChanged(QToolBar* self, intptr_t slot);
extern __declspec(dllexport) 
void QToolBar_OrientationChanged(QToolBar* self, int orientation);
void QToolBar_connect_OrientationChanged(QToolBar* self, intptr_t slot);
extern __declspec(dllexport) 
void QToolBar_IconSizeChanged(QToolBar* self, QSize* iconSize);
void QToolBar_connect_IconSizeChanged(QToolBar* self, intptr_t slot);
extern __declspec(dllexport) 
void QToolBar_ToolButtonStyleChanged(QToolBar* self, int toolButtonStyle);
void QToolBar_connect_ToolButtonStyleChanged(QToolBar* self, intptr_t slot);
extern __declspec(dllexport) 
void QToolBar_TopLevelChanged(QToolBar* self, bool topLevel);
void QToolBar_connect_TopLevelChanged(QToolBar* self, intptr_t slot);
extern __declspec(dllexport) 
void QToolBar_VisibilityChanged(QToolBar* self, bool visible);
void QToolBar_connect_VisibilityChanged(QToolBar* self, intptr_t slot);
extern __declspec(dllexport) 
void QToolBar_ActionEvent(QToolBar* self, QActionEvent* event);
extern __declspec(dllexport) 
void QToolBar_ChangeEvent(QToolBar* self, QEvent* event);
extern __declspec(dllexport) 
void QToolBar_PaintEvent(QToolBar* self, QPaintEvent* event);
extern __declspec(dllexport) 
bool QToolBar_Event(QToolBar* self, QEvent* event);
extern __declspec(dllexport) 
void QToolBar_InitStyleOption(const QToolBar* self, QStyleOptionToolBar* option);
extern __declspec(dllexport) 
struct miqt_string QToolBar_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QToolBar_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QToolBar_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QToolBar_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QToolBar_override_virtual_Metacast(void* self, intptr_t slot);
void* QToolBar_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QToolBar_Delete(QToolBar* self, bool isSubclass);

}
