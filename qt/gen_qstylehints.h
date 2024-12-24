#pragma once
#ifndef MIQT_QT_GEN_QSTYLEHINTS_H
#define MIQT_QT_GEN_QSTYLEHINTS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChar;
class QMetaObject;
class QObject;
class QStyleHints;
class _GUID;
class type_info;
#else
typedef struct QChar QChar;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QStyleHints QStyleHints;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QStyleHints_virtbase(QStyleHints* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QStyleHints_MetaObject(const QStyleHints* self);
extern __declspec(dllexport) void* QStyleHints_Metacast(QStyleHints* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QStyleHints_Tr(const char* s);
extern __declspec(dllexport) void QStyleHints_SetMouseDoubleClickInterval(QStyleHints* self, int mouseDoubleClickInterval);
extern __declspec(dllexport) int QStyleHints_MouseDoubleClickInterval(const QStyleHints* self);
extern __declspec(dllexport) int QStyleHints_MouseDoubleClickDistance(const QStyleHints* self);
extern __declspec(dllexport) int QStyleHints_TouchDoubleTapDistance(const QStyleHints* self);
extern __declspec(dllexport) void QStyleHints_SetMousePressAndHoldInterval(QStyleHints* self, int mousePressAndHoldInterval);
extern __declspec(dllexport) int QStyleHints_MousePressAndHoldInterval(const QStyleHints* self);
extern __declspec(dllexport) void QStyleHints_SetStartDragDistance(QStyleHints* self, int startDragDistance);
extern __declspec(dllexport) int QStyleHints_StartDragDistance(const QStyleHints* self);
extern __declspec(dllexport) void QStyleHints_SetStartDragTime(QStyleHints* self, int startDragTime);
extern __declspec(dllexport) int QStyleHints_StartDragTime(const QStyleHints* self);
extern __declspec(dllexport) int QStyleHints_StartDragVelocity(const QStyleHints* self);
extern __declspec(dllexport) void QStyleHints_SetKeyboardInputInterval(QStyleHints* self, int keyboardInputInterval);
extern __declspec(dllexport) int QStyleHints_KeyboardInputInterval(const QStyleHints* self);
extern __declspec(dllexport) int QStyleHints_KeyboardAutoRepeatRate(const QStyleHints* self);
extern __declspec(dllexport) double QStyleHints_KeyboardAutoRepeatRateF(const QStyleHints* self);
extern __declspec(dllexport) void QStyleHints_SetCursorFlashTime(QStyleHints* self, int cursorFlashTime);
extern __declspec(dllexport) int QStyleHints_CursorFlashTime(const QStyleHints* self);
extern __declspec(dllexport) bool QStyleHints_ShowIsFullScreen(const QStyleHints* self);
extern __declspec(dllexport) bool QStyleHints_ShowIsMaximized(const QStyleHints* self);
extern __declspec(dllexport) bool QStyleHints_ShowShortcutsInContextMenus(const QStyleHints* self);
extern __declspec(dllexport) void QStyleHints_SetShowShortcutsInContextMenus(QStyleHints* self, bool showShortcutsInContextMenus);
extern __declspec(dllexport) int QStyleHints_ContextMenuTrigger(const QStyleHints* self);
extern __declspec(dllexport) void QStyleHints_SetContextMenuTrigger(QStyleHints* self, int contextMenuTrigger);
extern __declspec(dllexport) int QStyleHints_PasswordMaskDelay(const QStyleHints* self);
extern __declspec(dllexport) QChar* QStyleHints_PasswordMaskCharacter(const QStyleHints* self);
extern __declspec(dllexport) double QStyleHints_FontSmoothingGamma(const QStyleHints* self);
extern __declspec(dllexport) bool QStyleHints_UseRtlExtensions(const QStyleHints* self);
extern __declspec(dllexport) bool QStyleHints_SetFocusOnTouchRelease(const QStyleHints* self);
extern __declspec(dllexport) int QStyleHints_TabFocusBehavior(const QStyleHints* self);
extern __declspec(dllexport) void QStyleHints_SetTabFocusBehavior(QStyleHints* self, int tabFocusBehavior);
extern __declspec(dllexport) bool QStyleHints_SingleClickActivation(const QStyleHints* self);
extern __declspec(dllexport) bool QStyleHints_UseHoverEffects(const QStyleHints* self);
extern __declspec(dllexport) void QStyleHints_SetUseHoverEffects(QStyleHints* self, bool useHoverEffects);
extern __declspec(dllexport) int QStyleHints_WheelScrollLines(const QStyleHints* self);
extern __declspec(dllexport) void QStyleHints_SetWheelScrollLines(QStyleHints* self, int scrollLines);
extern __declspec(dllexport) void QStyleHints_SetMouseQuickSelectionThreshold(QStyleHints* self, int threshold);
extern __declspec(dllexport) int QStyleHints_MouseQuickSelectionThreshold(const QStyleHints* self);
extern __declspec(dllexport) int QStyleHints_ColorScheme(const QStyleHints* self);
extern __declspec(dllexport) void QStyleHints_SetColorScheme(QStyleHints* self, int scheme);
extern __declspec(dllexport) void QStyleHints_UnsetColorScheme(QStyleHints* self);
extern __declspec(dllexport) void QStyleHints_CursorFlashTimeChanged(QStyleHints* self, int cursorFlashTime);
void QStyleHints_connect_CursorFlashTimeChanged(QStyleHints* self, intptr_t slot);
extern __declspec(dllexport) void QStyleHints_KeyboardInputIntervalChanged(QStyleHints* self, int keyboardInputInterval);
void QStyleHints_connect_KeyboardInputIntervalChanged(QStyleHints* self, intptr_t slot);
extern __declspec(dllexport) void QStyleHints_MouseDoubleClickIntervalChanged(QStyleHints* self, int mouseDoubleClickInterval);
void QStyleHints_connect_MouseDoubleClickIntervalChanged(QStyleHints* self, intptr_t slot);
extern __declspec(dllexport) void QStyleHints_MousePressAndHoldIntervalChanged(QStyleHints* self, int mousePressAndHoldInterval);
void QStyleHints_connect_MousePressAndHoldIntervalChanged(QStyleHints* self, intptr_t slot);
extern __declspec(dllexport) void QStyleHints_StartDragDistanceChanged(QStyleHints* self, int startDragDistance);
void QStyleHints_connect_StartDragDistanceChanged(QStyleHints* self, intptr_t slot);
extern __declspec(dllexport) void QStyleHints_StartDragTimeChanged(QStyleHints* self, int startDragTime);
void QStyleHints_connect_StartDragTimeChanged(QStyleHints* self, intptr_t slot);
extern __declspec(dllexport) void QStyleHints_TabFocusBehaviorChanged(QStyleHints* self, int tabFocusBehavior);
void QStyleHints_connect_TabFocusBehaviorChanged(QStyleHints* self, intptr_t slot);
extern __declspec(dllexport) void QStyleHints_UseHoverEffectsChanged(QStyleHints* self, bool useHoverEffects);
void QStyleHints_connect_UseHoverEffectsChanged(QStyleHints* self, intptr_t slot);
extern __declspec(dllexport) void QStyleHints_ShowShortcutsInContextMenusChanged(QStyleHints* self, bool param1);
void QStyleHints_connect_ShowShortcutsInContextMenusChanged(QStyleHints* self, intptr_t slot);
extern __declspec(dllexport) void QStyleHints_ContextMenuTriggerChanged(QStyleHints* self, int contextMenuTrigger);
void QStyleHints_connect_ContextMenuTriggerChanged(QStyleHints* self, intptr_t slot);
extern __declspec(dllexport) void QStyleHints_WheelScrollLinesChanged(QStyleHints* self, int scrollLines);
void QStyleHints_connect_WheelScrollLinesChanged(QStyleHints* self, intptr_t slot);
extern __declspec(dllexport) void QStyleHints_MouseQuickSelectionThresholdChanged(QStyleHints* self, int threshold);
void QStyleHints_connect_MouseQuickSelectionThresholdChanged(QStyleHints* self, intptr_t slot);
extern __declspec(dllexport) void QStyleHints_ColorSchemeChanged(QStyleHints* self, int colorScheme);
void QStyleHints_connect_ColorSchemeChanged(QStyleHints* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QStyleHints_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QStyleHints_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QStyleHints_Delete(QStyleHints* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
