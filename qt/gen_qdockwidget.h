#pragma once
#ifndef MIQT_QT_GEN_QDOCKWIDGET_H
#define MIQT_QT_GEN_QDOCKWIDGET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAction QAction;
typedef struct QCloseEvent QCloseEvent;
typedef struct QDockWidget QDockWidget;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QStyleOptionDockWidget QStyleOptionDockWidget;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QDockWidget* QDockWidget_new(QWidget* parent);
extern __declspec(dllexport) 
QDockWidget* QDockWidget_new2(struct miqt_string title);
extern __declspec(dllexport) 
QDockWidget* QDockWidget_new3();
extern __declspec(dllexport) 
QDockWidget* QDockWidget_new4(struct miqt_string title, QWidget* parent);
extern __declspec(dllexport) 
QDockWidget* QDockWidget_new5(struct miqt_string title, QWidget* parent, int flags);
extern __declspec(dllexport) 
QDockWidget* QDockWidget_new6(QWidget* parent, int flags);
extern __declspec(dllexport) 
void QDockWidget_virtbase(QDockWidget* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QDockWidget_MetaObject(const QDockWidget* self);
extern __declspec(dllexport) 
void* QDockWidget_Metacast(QDockWidget* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QDockWidget_Tr(const char* s);
extern __declspec(dllexport) 
QWidget* QDockWidget_Widget(const QDockWidget* self);
extern __declspec(dllexport) 
void QDockWidget_SetWidget(QDockWidget* self, QWidget* widget);
extern __declspec(dllexport) 
void QDockWidget_SetFeatures(QDockWidget* self, DockWidgetFeatures features);
extern __declspec(dllexport) 
DockWidgetFeatures QDockWidget_Features(const QDockWidget* self);
extern __declspec(dllexport) 
void QDockWidget_SetFloating(QDockWidget* self, bool floating);
extern __declspec(dllexport) 
bool QDockWidget_IsFloating(const QDockWidget* self);
extern __declspec(dllexport) 
void QDockWidget_SetAllowedAreas(QDockWidget* self, int areas);
extern __declspec(dllexport) 
int QDockWidget_AllowedAreas(const QDockWidget* self);
extern __declspec(dllexport) 
void QDockWidget_SetTitleBarWidget(QDockWidget* self, QWidget* widget);
extern __declspec(dllexport) 
QWidget* QDockWidget_TitleBarWidget(const QDockWidget* self);
extern __declspec(dllexport) 
void QDockWidget_SetDockLocation(QDockWidget* self, int area);
extern __declspec(dllexport) 
int QDockWidget_DockLocation(const QDockWidget* self);
extern __declspec(dllexport) 
bool QDockWidget_IsAreaAllowed(const QDockWidget* self, int area);
extern __declspec(dllexport) 
QAction* QDockWidget_ToggleViewAction(const QDockWidget* self);
extern __declspec(dllexport) 
void QDockWidget_FeaturesChanged(QDockWidget* self, int features);
void QDockWidget_connect_FeaturesChanged(QDockWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QDockWidget_TopLevelChanged(QDockWidget* self, bool topLevel);
void QDockWidget_connect_TopLevelChanged(QDockWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QDockWidget_AllowedAreasChanged(QDockWidget* self, int allowedAreas);
void QDockWidget_connect_AllowedAreasChanged(QDockWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QDockWidget_VisibilityChanged(QDockWidget* self, bool visible);
void QDockWidget_connect_VisibilityChanged(QDockWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QDockWidget_DockLocationChanged(QDockWidget* self, int area);
void QDockWidget_connect_DockLocationChanged(QDockWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QDockWidget_ChangeEvent(QDockWidget* self, QEvent* event);
extern __declspec(dllexport) 
void QDockWidget_CloseEvent(QDockWidget* self, QCloseEvent* event);
extern __declspec(dllexport) 
void QDockWidget_PaintEvent(QDockWidget* self, QPaintEvent* event);
extern __declspec(dllexport) 
bool QDockWidget_Event(QDockWidget* self, QEvent* event);
extern __declspec(dllexport) 
void QDockWidget_InitStyleOption(const QDockWidget* self, QStyleOptionDockWidget* option);
extern __declspec(dllexport) 
struct miqt_string QDockWidget_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QDockWidget_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QDockWidget_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QDockWidget_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QDockWidget_override_virtual_Metacast(void* self, intptr_t slot);
void* QDockWidget_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QDockWidget_Delete(QDockWidget* self, bool isSubclass);

}
