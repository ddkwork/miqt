#pragma once
#ifndef MIQT_QT_GEN_QMAINWINDOW_H
#define MIQT_QT_GEN_QMAINWINDOW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDockWidget QDockWidget;
typedef struct QEvent QEvent;
typedef struct QMainWindow QMainWindow;
typedef struct QMenu QMenu;
typedef struct QMenuBar QMenuBar;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPoint QPoint;
typedef struct QSize QSize;
typedef struct QStatusBar QStatusBar;
typedef struct QToolBar QToolBar;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QMainWindow* QMainWindow_new(QWidget* parent);
extern __declspec(dllexport) 
QMainWindow* QMainWindow_new2();
extern __declspec(dllexport) 
QMainWindow* QMainWindow_new3(QWidget* parent, int flags);
extern __declspec(dllexport) 
void QMainWindow_virtbase(QMainWindow* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QMainWindow_MetaObject(const QMainWindow* self);
extern __declspec(dllexport) 
void* QMainWindow_Metacast(QMainWindow* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QMainWindow_Tr(const char* s);
extern __declspec(dllexport) 
QSize* QMainWindow_IconSize(const QMainWindow* self);
extern __declspec(dllexport) 
void QMainWindow_SetIconSize(QMainWindow* self, QSize* iconSize);
extern __declspec(dllexport) 
int QMainWindow_ToolButtonStyle(const QMainWindow* self);
extern __declspec(dllexport) 
void QMainWindow_SetToolButtonStyle(QMainWindow* self, int toolButtonStyle);
extern __declspec(dllexport) 
bool QMainWindow_IsAnimated(const QMainWindow* self);
extern __declspec(dllexport) 
bool QMainWindow_IsDockNestingEnabled(const QMainWindow* self);
extern __declspec(dllexport) 
bool QMainWindow_DocumentMode(const QMainWindow* self);
extern __declspec(dllexport) 
void QMainWindow_SetDocumentMode(QMainWindow* self, bool enabled);
extern __declspec(dllexport) 
int QMainWindow_TabShape(const QMainWindow* self);
extern __declspec(dllexport) 
void QMainWindow_SetTabShape(QMainWindow* self, int tabShape);
extern __declspec(dllexport) 
int QMainWindow_TabPosition(const QMainWindow* self, int area);
extern __declspec(dllexport) 
void QMainWindow_SetTabPosition(QMainWindow* self, int areas, int tabPosition);
extern __declspec(dllexport) 
void QMainWindow_SetDockOptions(QMainWindow* self, DockOptions options);
extern __declspec(dllexport) 
DockOptions QMainWindow_DockOptions(const QMainWindow* self);
extern __declspec(dllexport) 
bool QMainWindow_IsSeparator(const QMainWindow* self, QPoint* pos);
extern __declspec(dllexport) 
QMenuBar* QMainWindow_MenuBar(const QMainWindow* self);
extern __declspec(dllexport) 
void QMainWindow_SetMenuBar(QMainWindow* self, QMenuBar* menubar);
extern __declspec(dllexport) 
QWidget* QMainWindow_MenuWidget(const QMainWindow* self);
extern __declspec(dllexport) 
void QMainWindow_SetMenuWidget(QMainWindow* self, QWidget* menubar);
extern __declspec(dllexport) 
QStatusBar* QMainWindow_StatusBar(const QMainWindow* self);
extern __declspec(dllexport) 
void QMainWindow_SetStatusBar(QMainWindow* self, QStatusBar* statusbar);
extern __declspec(dllexport) 
QWidget* QMainWindow_CentralWidget(const QMainWindow* self);
extern __declspec(dllexport) 
void QMainWindow_SetCentralWidget(QMainWindow* self, QWidget* widget);
extern __declspec(dllexport) 
QWidget* QMainWindow_TakeCentralWidget(QMainWindow* self);
extern __declspec(dllexport) 
void QMainWindow_SetCorner(QMainWindow* self, int corner, int area);
extern __declspec(dllexport) 
int QMainWindow_Corner(const QMainWindow* self, int corner);
extern __declspec(dllexport) 
void QMainWindow_AddToolBarBreak(QMainWindow* self);
extern __declspec(dllexport) 
void QMainWindow_InsertToolBarBreak(QMainWindow* self, QToolBar* before);
extern __declspec(dllexport) 
void QMainWindow_AddToolBar(QMainWindow* self, int area, QToolBar* toolbar);
extern __declspec(dllexport) 
void QMainWindow_AddToolBarWithToolbar(QMainWindow* self, QToolBar* toolbar);
extern __declspec(dllexport) 
QToolBar* QMainWindow_AddToolBarWithTitle(QMainWindow* self, struct miqt_string title);
extern __declspec(dllexport) 
void QMainWindow_InsertToolBar(QMainWindow* self, QToolBar* before, QToolBar* toolbar);
extern __declspec(dllexport) 
void QMainWindow_RemoveToolBar(QMainWindow* self, QToolBar* toolbar);
extern __declspec(dllexport) 
void QMainWindow_RemoveToolBarBreak(QMainWindow* self, QToolBar* before);
extern __declspec(dllexport) 
bool QMainWindow_UnifiedTitleAndToolBarOnMac(const QMainWindow* self);
extern __declspec(dllexport) 
int QMainWindow_ToolBarArea(const QMainWindow* self, QToolBar* toolbar);
extern __declspec(dllexport) 
bool QMainWindow_ToolBarBreak(const QMainWindow* self, QToolBar* toolbar);
extern __declspec(dllexport) 
void QMainWindow_AddDockWidget(QMainWindow* self, int area, QDockWidget* dockwidget);
extern __declspec(dllexport) 
void QMainWindow_AddDockWidget2(QMainWindow* self, int area, QDockWidget* dockwidget, int orientation);
extern __declspec(dllexport) 
void QMainWindow_SplitDockWidget(QMainWindow* self, QDockWidget* after, QDockWidget* dockwidget, int orientation);
extern __declspec(dllexport) 
void QMainWindow_TabifyDockWidget(QMainWindow* self, QDockWidget* first, QDockWidget* second);
extern __declspec(dllexport) 
struct miqt_array /* of QDockWidget* */  QMainWindow_TabifiedDockWidgets(const QMainWindow* self, QDockWidget* dockwidget);
extern __declspec(dllexport) 
void QMainWindow_RemoveDockWidget(QMainWindow* self, QDockWidget* dockwidget);
extern __declspec(dllexport) 
bool QMainWindow_RestoreDockWidget(QMainWindow* self, QDockWidget* dockwidget);
extern __declspec(dllexport) 
int QMainWindow_DockWidgetArea(const QMainWindow* self, QDockWidget* dockwidget);
extern __declspec(dllexport) 
void QMainWindow_ResizeDocks(QMainWindow* self, struct miqt_array /* of QDockWidget* */  docks, struct miqt_array /* of int */  sizes, int orientation);
extern __declspec(dllexport) 
struct miqt_string QMainWindow_SaveState(const QMainWindow* self);
extern __declspec(dllexport) 
bool QMainWindow_RestoreState(QMainWindow* self, struct miqt_string state);
extern __declspec(dllexport) 
QMenu* QMainWindow_CreatePopupMenu(QMainWindow* self);
extern __declspec(dllexport) 
void QMainWindow_SetAnimated(QMainWindow* self, bool enabled);
extern __declspec(dllexport) 
void QMainWindow_SetDockNestingEnabled(QMainWindow* self, bool enabled);
extern __declspec(dllexport) 
void QMainWindow_SetUnifiedTitleAndToolBarOnMac(QMainWindow* self, bool set);
extern __declspec(dllexport) 
void QMainWindow_IconSizeChanged(QMainWindow* self, QSize* iconSize);
void QMainWindow_connect_IconSizeChanged(QMainWindow* self, intptr_t slot);
extern __declspec(dllexport) 
void QMainWindow_ToolButtonStyleChanged(QMainWindow* self, int toolButtonStyle);
void QMainWindow_connect_ToolButtonStyleChanged(QMainWindow* self, intptr_t slot);
extern __declspec(dllexport) 
void QMainWindow_TabifiedDockWidgetActivated(QMainWindow* self, QDockWidget* dockWidget);
void QMainWindow_connect_TabifiedDockWidgetActivated(QMainWindow* self, intptr_t slot);
extern __declspec(dllexport) 
void QMainWindow_ContextMenuEvent(QMainWindow* self, QContextMenuEvent* event);
extern __declspec(dllexport) 
bool QMainWindow_Event(QMainWindow* self, QEvent* event);
extern __declspec(dllexport) 
struct miqt_string QMainWindow_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QMainWindow_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QMainWindow_AddToolBarBreak1(QMainWindow* self, int area);
extern __declspec(dllexport) 
struct miqt_string QMainWindow_SaveState1(const QMainWindow* self, int version);
extern __declspec(dllexport) 
bool QMainWindow_RestoreState2(QMainWindow* self, struct miqt_string state, int version);
extern __declspec(dllexport) 
void QMainWindow_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QMainWindow_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QMainWindow_override_virtual_Metacast(void* self, intptr_t slot);
void* QMainWindow_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QMainWindow_Delete(QMainWindow* self, bool isSubclass);

}
