#pragma once
#ifndef MIQT_QT_GEN_QSYSTEMTRAYICON_H
#define MIQT_QT_GEN_QSYSTEMTRAYICON_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QIcon QIcon;
typedef struct QMenu QMenu;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QRect QRect;
typedef struct QSystemTrayIcon QSystemTrayIcon;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QSystemTrayIcon* QSystemTrayIcon_new();
extern __declspec(dllexport) 
QSystemTrayIcon* QSystemTrayIcon_new2(QIcon* icon);
extern __declspec(dllexport) 
QSystemTrayIcon* QSystemTrayIcon_new3(QObject* parent);
extern __declspec(dllexport) 
QSystemTrayIcon* QSystemTrayIcon_new4(QIcon* icon, QObject* parent);
extern __declspec(dllexport) 
void QSystemTrayIcon_virtbase(QSystemTrayIcon* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QSystemTrayIcon_MetaObject(const QSystemTrayIcon* self);
extern __declspec(dllexport) 
void* QSystemTrayIcon_Metacast(QSystemTrayIcon* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QSystemTrayIcon_Tr(const char* s);
extern __declspec(dllexport) 
void QSystemTrayIcon_SetContextMenu(QSystemTrayIcon* self, QMenu* menu);
extern __declspec(dllexport) 
QMenu* QSystemTrayIcon_ContextMenu(const QSystemTrayIcon* self);
extern __declspec(dllexport) 
QIcon* QSystemTrayIcon_Icon(const QSystemTrayIcon* self);
extern __declspec(dllexport) 
void QSystemTrayIcon_SetIcon(QSystemTrayIcon* self, QIcon* icon);
extern __declspec(dllexport) 
struct miqt_string QSystemTrayIcon_ToolTip(const QSystemTrayIcon* self);
extern __declspec(dllexport) 
void QSystemTrayIcon_SetToolTip(QSystemTrayIcon* self, struct miqt_string tip);
extern __declspec(dllexport) 
bool QSystemTrayIcon_IsSystemTrayAvailable();
extern __declspec(dllexport) 
bool QSystemTrayIcon_SupportsMessages();
extern __declspec(dllexport) 
QRect* QSystemTrayIcon_Geometry(const QSystemTrayIcon* self);
extern __declspec(dllexport) 
bool QSystemTrayIcon_IsVisible(const QSystemTrayIcon* self);
extern __declspec(dllexport) 
void QSystemTrayIcon_SetVisible(QSystemTrayIcon* self, bool visible);
extern __declspec(dllexport) 
void QSystemTrayIcon_Show(QSystemTrayIcon* self);
extern __declspec(dllexport) 
void QSystemTrayIcon_Hide(QSystemTrayIcon* self);
extern __declspec(dllexport) 
void QSystemTrayIcon_ShowMessage(QSystemTrayIcon* self, struct miqt_string title, struct miqt_string msg, QIcon* icon);
extern __declspec(dllexport) 
void QSystemTrayIcon_ShowMessage2(QSystemTrayIcon* self, struct miqt_string title, struct miqt_string msg);
extern __declspec(dllexport) 
void QSystemTrayIcon_Activated(QSystemTrayIcon* self, int reason);
void QSystemTrayIcon_connect_Activated(QSystemTrayIcon* self, intptr_t slot);
extern __declspec(dllexport) 
void QSystemTrayIcon_MessageClicked(QSystemTrayIcon* self);
void QSystemTrayIcon_connect_MessageClicked(QSystemTrayIcon* self, intptr_t slot);
extern __declspec(dllexport) 
bool QSystemTrayIcon_Event(QSystemTrayIcon* self, QEvent* event);
extern __declspec(dllexport) 
struct miqt_string QSystemTrayIcon_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QSystemTrayIcon_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QSystemTrayIcon_ShowMessage4(QSystemTrayIcon* self, struct miqt_string title, struct miqt_string msg, QIcon* icon, int msecs);
extern __declspec(dllexport) 
void QSystemTrayIcon_ShowMessage3(QSystemTrayIcon* self, struct miqt_string title, struct miqt_string msg, int icon);
extern __declspec(dllexport) 
void QSystemTrayIcon_ShowMessage42(QSystemTrayIcon* self, struct miqt_string title, struct miqt_string msg, int icon, int msecs);
extern __declspec(dllexport) 
void QSystemTrayIcon_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QSystemTrayIcon_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QSystemTrayIcon_override_virtual_Metacast(void* self, intptr_t slot);
void* QSystemTrayIcon_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QSystemTrayIcon_Delete(QSystemTrayIcon* self, bool isSubclass);

}
