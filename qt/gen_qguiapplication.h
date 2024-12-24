#pragma once
#ifndef MIQT_QT_GEN_QGUIAPPLICATION_H
#define MIQT_QT_GEN_QGUIAPPLICATION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QClipboard;
class QCoreApplication;
class QCursor;
class QEvent;
class QFont;
class QGuiApplication;
class QIcon;
class QInputMethod;
class QMetaObject;
class QObject;
class QPalette;
class QPoint;
class QScreen;
class QSessionManager;
class QStyleHints;
class QWindow;
class _GUID;
class type_info;
#else
typedef struct QClipboard QClipboard;
typedef struct QCoreApplication QCoreApplication;
typedef struct QCursor QCursor;
typedef struct QEvent QEvent;
typedef struct QFont QFont;
typedef struct QGuiApplication QGuiApplication;
typedef struct QIcon QIcon;
typedef struct QInputMethod QInputMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPalette QPalette;
typedef struct QPoint QPoint;
typedef struct QScreen QScreen;
typedef struct QSessionManager QSessionManager;
typedef struct QStyleHints QStyleHints;
typedef struct QWindow QWindow;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QGuiApplication* QGuiApplication_new(int* argc, char** argv);
extern __declspec(dllexport) QGuiApplication* QGuiApplication_new2(int* argc, char** argv, int param3);
extern __declspec(dllexport) void QGuiApplication_virtbase(QGuiApplication* src, QCoreApplication** outptr_QCoreApplication);
extern __declspec(dllexport) QMetaObject* QGuiApplication_MetaObject(const QGuiApplication* self);
extern __declspec(dllexport) void* QGuiApplication_Metacast(QGuiApplication* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QGuiApplication_Tr(const char* s);
extern __declspec(dllexport) void QGuiApplication_SetApplicationDisplayName(struct miqt_string name);
extern __declspec(dllexport) struct miqt_string QGuiApplication_ApplicationDisplayName();
extern __declspec(dllexport) void QGuiApplication_SetBadgeNumber(QGuiApplication* self, long long number);
extern __declspec(dllexport) void QGuiApplication_SetDesktopFileName(struct miqt_string name);
extern __declspec(dllexport) struct miqt_string QGuiApplication_DesktopFileName();
extern __declspec(dllexport) struct miqt_array /* of QWindow* */  QGuiApplication_AllWindows();
extern __declspec(dllexport) struct miqt_array /* of QWindow* */  QGuiApplication_TopLevelWindows();
extern __declspec(dllexport) QWindow* QGuiApplication_TopLevelAt(QPoint* pos);
extern __declspec(dllexport) void QGuiApplication_SetWindowIcon(QIcon* icon);
extern __declspec(dllexport) QIcon* QGuiApplication_WindowIcon();
extern __declspec(dllexport) struct miqt_string QGuiApplication_PlatformName();
extern __declspec(dllexport) QWindow* QGuiApplication_ModalWindow();
extern __declspec(dllexport) QWindow* QGuiApplication_FocusWindow();
extern __declspec(dllexport) QObject* QGuiApplication_FocusObject();
extern __declspec(dllexport) QScreen* QGuiApplication_PrimaryScreen();
extern __declspec(dllexport) struct miqt_array /* of QScreen* */  QGuiApplication_Screens();
extern __declspec(dllexport) QScreen* QGuiApplication_ScreenAt(QPoint* point);
extern __declspec(dllexport) double QGuiApplication_DevicePixelRatio(const QGuiApplication* self);
extern __declspec(dllexport) QCursor* QGuiApplication_OverrideCursor();
extern __declspec(dllexport) void QGuiApplication_SetOverrideCursor(QCursor* overrideCursor);
extern __declspec(dllexport) void QGuiApplication_ChangeOverrideCursor(QCursor* param1);
extern __declspec(dllexport) void QGuiApplication_RestoreOverrideCursor();
extern __declspec(dllexport) QFont* QGuiApplication_Font();
extern __declspec(dllexport) void QGuiApplication_SetFont(QFont* font);
extern __declspec(dllexport) QClipboard* QGuiApplication_Clipboard();
extern __declspec(dllexport) QPalette* QGuiApplication_Palette();
extern __declspec(dllexport) void QGuiApplication_SetPalette(QPalette* pal);
extern __declspec(dllexport) int QGuiApplication_KeyboardModifiers();
extern __declspec(dllexport) int QGuiApplication_QueryKeyboardModifiers();
extern __declspec(dllexport) int QGuiApplication_MouseButtons();
extern __declspec(dllexport) void QGuiApplication_SetLayoutDirection(int direction);
extern __declspec(dllexport) int QGuiApplication_LayoutDirection();
extern __declspec(dllexport) bool QGuiApplication_IsRightToLeft();
extern __declspec(dllexport) bool QGuiApplication_IsLeftToRight();
extern __declspec(dllexport) QStyleHints* QGuiApplication_StyleHints();
extern __declspec(dllexport) void QGuiApplication_SetDesktopSettingsAware(bool on);
extern __declspec(dllexport) bool QGuiApplication_DesktopSettingsAware();
extern __declspec(dllexport) QInputMethod* QGuiApplication_InputMethod();
extern __declspec(dllexport) void QGuiApplication_SetQuitOnLastWindowClosed(bool quit);
extern __declspec(dllexport) bool QGuiApplication_QuitOnLastWindowClosed();
extern __declspec(dllexport) int QGuiApplication_ApplicationState();
extern __declspec(dllexport) void QGuiApplication_SetHighDpiScaleFactorRoundingPolicy(int policy);
extern __declspec(dllexport) int QGuiApplication_HighDpiScaleFactorRoundingPolicy();
extern __declspec(dllexport) int QGuiApplication_Exec();
extern __declspec(dllexport) bool QGuiApplication_Notify(QGuiApplication* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) bool QGuiApplication_IsSessionRestored(const QGuiApplication* self);
extern __declspec(dllexport) struct miqt_string QGuiApplication_SessionId(const QGuiApplication* self);
extern __declspec(dllexport) struct miqt_string QGuiApplication_SessionKey(const QGuiApplication* self);
extern __declspec(dllexport) bool QGuiApplication_IsSavingSession(const QGuiApplication* self);
extern __declspec(dllexport) void QGuiApplication_Sync();
extern __declspec(dllexport) void QGuiApplication_FontDatabaseChanged(QGuiApplication* self);
void QGuiApplication_connect_FontDatabaseChanged(QGuiApplication* self, intptr_t slot);
extern __declspec(dllexport) void QGuiApplication_ScreenAdded(QGuiApplication* self, QScreen* screen);
void QGuiApplication_connect_ScreenAdded(QGuiApplication* self, intptr_t slot);
extern __declspec(dllexport) void QGuiApplication_ScreenRemoved(QGuiApplication* self, QScreen* screen);
void QGuiApplication_connect_ScreenRemoved(QGuiApplication* self, intptr_t slot);
extern __declspec(dllexport) void QGuiApplication_PrimaryScreenChanged(QGuiApplication* self, QScreen* screen);
void QGuiApplication_connect_PrimaryScreenChanged(QGuiApplication* self, intptr_t slot);
extern __declspec(dllexport) void QGuiApplication_LastWindowClosed(QGuiApplication* self);
void QGuiApplication_connect_LastWindowClosed(QGuiApplication* self, intptr_t slot);
extern __declspec(dllexport) void QGuiApplication_FocusObjectChanged(QGuiApplication* self, QObject* focusObject);
void QGuiApplication_connect_FocusObjectChanged(QGuiApplication* self, intptr_t slot);
extern __declspec(dllexport) void QGuiApplication_FocusWindowChanged(QGuiApplication* self, QWindow* focusWindow);
void QGuiApplication_connect_FocusWindowChanged(QGuiApplication* self, intptr_t slot);
extern __declspec(dllexport) void QGuiApplication_ApplicationStateChanged(QGuiApplication* self, int state);
void QGuiApplication_connect_ApplicationStateChanged(QGuiApplication* self, intptr_t slot);
extern __declspec(dllexport) void QGuiApplication_LayoutDirectionChanged(QGuiApplication* self, int direction);
void QGuiApplication_connect_LayoutDirectionChanged(QGuiApplication* self, intptr_t slot);
extern __declspec(dllexport) void QGuiApplication_CommitDataRequest(QGuiApplication* self, QSessionManager* sessionManager);
void QGuiApplication_connect_CommitDataRequest(QGuiApplication* self, intptr_t slot);
extern __declspec(dllexport) void QGuiApplication_SaveStateRequest(QGuiApplication* self, QSessionManager* sessionManager);
void QGuiApplication_connect_SaveStateRequest(QGuiApplication* self, intptr_t slot);
extern __declspec(dllexport) void QGuiApplication_ApplicationDisplayNameChanged(QGuiApplication* self);
void QGuiApplication_connect_ApplicationDisplayNameChanged(QGuiApplication* self, intptr_t slot);
extern __declspec(dllexport) void QGuiApplication_PaletteChanged(QGuiApplication* self, QPalette* pal);
void QGuiApplication_connect_PaletteChanged(QGuiApplication* self, intptr_t slot);
extern __declspec(dllexport) void QGuiApplication_FontChanged(QGuiApplication* self, QFont* font);
void QGuiApplication_connect_FontChanged(QGuiApplication* self, intptr_t slot);
extern __declspec(dllexport) bool QGuiApplication_Event(QGuiApplication* self, QEvent* param1);
extern __declspec(dllexport) struct miqt_string QGuiApplication_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QGuiApplication_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QGuiApplication_override_virtual_Notify(void* self, intptr_t slot);
bool QGuiApplication_virtualbase_Notify(void* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) void QGuiApplication_override_virtual_Event(void* self, intptr_t slot);
bool QGuiApplication_virtualbase_Event(void* self, QEvent* param1);
extern __declspec(dllexport) void QGuiApplication_Delete(QGuiApplication* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
