#pragma once
#ifndef MIQT_QT_GEN_QACTION_H
#define MIQT_QT_GEN_QACTION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAction QAction;
typedef struct QActionGroup QActionGroup;
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QFont QFont;
typedef struct QIcon QIcon;
typedef struct QKeySequence QKeySequence;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;
typedef struct QVariant QVariant;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QAction* QAction_new();
extern __declspec(dllexport) QAction* QAction_new2(struct miqt_string text);
extern __declspec(dllexport) QAction* QAction_new3(QIcon* icon, struct miqt_string text);
extern __declspec(dllexport) QAction* QAction_new4(QObject* parent);
extern __declspec(dllexport) QAction* QAction_new5(struct miqt_string text, QObject* parent);
extern __declspec(dllexport) QAction* QAction_new6(QIcon* icon, struct miqt_string text, QObject* parent);
extern __declspec(dllexport) void QAction_virtbase(QAction* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QAction_MetaObject(const QAction* self);
extern __declspec(dllexport) void* QAction_Metacast(QAction* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QAction_Tr(const char* s);
extern __declspec(dllexport) struct miqt_array /* of QObject* */  QAction_AssociatedObjects(const QAction* self);
extern __declspec(dllexport) void QAction_SetActionGroup(QAction* self, QActionGroup* group);
extern __declspec(dllexport) QActionGroup* QAction_ActionGroup(const QAction* self);
extern __declspec(dllexport) void QAction_SetIcon(QAction* self, QIcon* icon);
extern __declspec(dllexport) QIcon* QAction_Icon(const QAction* self);
extern __declspec(dllexport) void QAction_SetText(QAction* self, struct miqt_string text);
extern __declspec(dllexport) struct miqt_string QAction_Text(const QAction* self);
extern __declspec(dllexport) void QAction_SetIconText(QAction* self, struct miqt_string text);
extern __declspec(dllexport) struct miqt_string QAction_IconText(const QAction* self);
extern __declspec(dllexport) void QAction_SetToolTip(QAction* self, struct miqt_string tip);
extern __declspec(dllexport) struct miqt_string QAction_ToolTip(const QAction* self);
extern __declspec(dllexport) void QAction_SetStatusTip(QAction* self, struct miqt_string statusTip);
extern __declspec(dllexport) struct miqt_string QAction_StatusTip(const QAction* self);
extern __declspec(dllexport) void QAction_SetWhatsThis(QAction* self, struct miqt_string what);
extern __declspec(dllexport) struct miqt_string QAction_WhatsThis(const QAction* self);
extern __declspec(dllexport) void QAction_SetPriority(QAction* self, Priority priority);
extern __declspec(dllexport) Priority QAction_Priority(const QAction* self);
extern __declspec(dllexport) void QAction_SetSeparator(QAction* self, bool b);
extern __declspec(dllexport) bool QAction_IsSeparator(const QAction* self);
extern __declspec(dllexport) void QAction_SetShortcut(QAction* self, QKeySequence* shortcut);
extern __declspec(dllexport) QKeySequence* QAction_Shortcut(const QAction* self);
extern __declspec(dllexport) void QAction_SetShortcuts(QAction* self, struct miqt_array /* of QKeySequence* */  shortcuts);
extern __declspec(dllexport) void QAction_SetShortcutsWithShortcuts(QAction* self, int shortcuts);
extern __declspec(dllexport) struct miqt_array /* of QKeySequence* */  QAction_Shortcuts(const QAction* self);
extern __declspec(dllexport) void QAction_SetShortcutContext(QAction* self, int context);
extern __declspec(dllexport) int QAction_ShortcutContext(const QAction* self);
extern __declspec(dllexport) void QAction_SetAutoRepeat(QAction* self, bool autoRepeat);
extern __declspec(dllexport) bool QAction_AutoRepeat(const QAction* self);
extern __declspec(dllexport) void QAction_SetFont(QAction* self, QFont* font);
extern __declspec(dllexport) QFont* QAction_Font(const QAction* self);
extern __declspec(dllexport) void QAction_SetCheckable(QAction* self, bool checkable);
extern __declspec(dllexport) bool QAction_IsCheckable(const QAction* self);
extern __declspec(dllexport) QVariant* QAction_Data(const QAction* self);
extern __declspec(dllexport) void QAction_SetData(QAction* self, QVariant* varVal);
extern __declspec(dllexport) bool QAction_IsChecked(const QAction* self);
extern __declspec(dllexport) bool QAction_IsEnabled(const QAction* self);
extern __declspec(dllexport) bool QAction_IsVisible(const QAction* self);
extern __declspec(dllexport) void QAction_Activate(QAction* self, ActionEvent event);
extern __declspec(dllexport) void QAction_SetMenuRole(QAction* self, MenuRole menuRole);
extern __declspec(dllexport) MenuRole QAction_MenuRole(const QAction* self);
extern __declspec(dllexport) void QAction_SetIconVisibleInMenu(QAction* self, bool visible);
extern __declspec(dllexport) bool QAction_IsIconVisibleInMenu(const QAction* self);
extern __declspec(dllexport) void QAction_SetShortcutVisibleInContextMenu(QAction* self, bool show);
extern __declspec(dllexport) bool QAction_IsShortcutVisibleInContextMenu(const QAction* self);
extern __declspec(dllexport) bool QAction_ShowStatusText(QAction* self);
extern __declspec(dllexport) bool QAction_Event(QAction* self, QEvent* param1);
extern __declspec(dllexport) void QAction_Trigger(QAction* self);
extern __declspec(dllexport) void QAction_Hover(QAction* self);
extern __declspec(dllexport) void QAction_SetChecked(QAction* self, bool checked);
extern __declspec(dllexport) void QAction_Toggle(QAction* self);
extern __declspec(dllexport) void QAction_SetEnabled(QAction* self, bool enabled);
extern __declspec(dllexport) void QAction_ResetEnabled(QAction* self);
extern __declspec(dllexport) void QAction_SetDisabled(QAction* self, bool b);
extern __declspec(dllexport) void QAction_SetVisible(QAction* self, bool visible);
extern __declspec(dllexport) void QAction_Changed(QAction* self);
void QAction_connect_Changed(QAction* self, intptr_t slot);
extern __declspec(dllexport) void QAction_EnabledChanged(QAction* self, bool enabled);
void QAction_connect_EnabledChanged(QAction* self, intptr_t slot);
extern __declspec(dllexport) void QAction_CheckableChanged(QAction* self, bool checkable);
void QAction_connect_CheckableChanged(QAction* self, intptr_t slot);
extern __declspec(dllexport) void QAction_VisibleChanged(QAction* self);
void QAction_connect_VisibleChanged(QAction* self, intptr_t slot);
extern __declspec(dllexport) void QAction_Triggered(QAction* self);
void QAction_connect_Triggered(QAction* self, intptr_t slot);
extern __declspec(dllexport) void QAction_Hovered(QAction* self);
void QAction_connect_Hovered(QAction* self, intptr_t slot);
extern __declspec(dllexport) void QAction_Toggled(QAction* self, bool param1);
void QAction_connect_Toggled(QAction* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QAction_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QAction_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) bool QAction_ShowStatusText1(QAction* self, QObject* object);
extern __declspec(dllexport) void QAction_Triggered1(QAction* self, bool checked);
void QAction_connect_Triggered1(QAction* self, intptr_t slot);
extern __declspec(dllexport) void QAction_override_virtual_Event(void* self, intptr_t slot);
bool QAction_virtualbase_Event(void* self, QEvent* param1);
extern __declspec(dllexport) void QAction_override_virtual_EventFilter(void* self, intptr_t slot);
bool QAction_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QAction_override_virtual_TimerEvent(void* self, intptr_t slot);
void QAction_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QAction_override_virtual_ChildEvent(void* self, intptr_t slot);
void QAction_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QAction_override_virtual_CustomEvent(void* self, intptr_t slot);
void QAction_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QAction_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QAction_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QAction_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QAction_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QAction_Delete(QAction* self, bool isSubclass);

} 
