#pragma once
#ifndef MIQT_QT_GEN_QAPPLICATION_H
#define MIQT_QT_GEN_QAPPLICATION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
typedef struct QApplication QApplication;
typedef struct QCoreApplication QCoreApplication;
typedef struct QEvent QEvent;
typedef struct QFont QFont;
typedef struct QFontMetrics QFontMetrics;
typedef struct QGuiApplication QGuiApplication;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPalette QPalette;
typedef struct QPoint QPoint;
typedef struct QStyle QStyle;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QApplication* QApplication_new(int* argc, char** argv);
extern __declspec(dllexport) QApplication* QApplication_new2(int* argc, char** argv, int param3);
extern __declspec(dllexport) void QApplication_virtbase(QApplication* src, QGuiApplication** outptr_QGuiApplication);
extern __declspec(dllexport) QMetaObject* QApplication_MetaObject(const QApplication* self);
extern __declspec(dllexport) void* QApplication_Metacast(QApplication* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QApplication_Tr(const char* s);
extern __declspec(dllexport) QStyle* QApplication_Style();
extern __declspec(dllexport) void QApplication_SetStyle(QStyle* style);
extern __declspec(dllexport) QStyle* QApplication_SetStyleWithStyle(struct miqt_string style);
extern __declspec(dllexport) QPalette* QApplication_Palette(QWidget* param1);
extern __declspec(dllexport) QPalette* QApplication_PaletteWithClassName(const char* className);
extern __declspec(dllexport) void QApplication_SetPalette(QPalette* param1);
extern __declspec(dllexport) QFont* QApplication_Font();
extern __declspec(dllexport) QFont* QApplication_FontWithQWidget(QWidget* param1);
extern __declspec(dllexport) QFont* QApplication_FontWithClassName(const char* className);
extern __declspec(dllexport) void QApplication_SetFont(QFont* param1);
extern __declspec(dllexport) QFontMetrics* QApplication_FontMetrics();
extern __declspec(dllexport) struct miqt_array /* of QWidget* */  QApplication_AllWidgets();
extern __declspec(dllexport) struct miqt_array /* of QWidget* */  QApplication_TopLevelWidgets();
extern __declspec(dllexport) QWidget* QApplication_ActivePopupWidget();
extern __declspec(dllexport) QWidget* QApplication_ActiveModalWidget();
extern __declspec(dllexport) QWidget* QApplication_FocusWidget();
extern __declspec(dllexport) QWidget* QApplication_ActiveWindow();
extern __declspec(dllexport) void QApplication_SetActiveWindow(QWidget* act);
extern __declspec(dllexport) QWidget* QApplication_WidgetAt(QPoint* p);
extern __declspec(dllexport) QWidget* QApplication_WidgetAt2(int x, int y);
extern __declspec(dllexport) QWidget* QApplication_TopLevelAt(QPoint* p);
extern __declspec(dllexport) QWidget* QApplication_TopLevelAt2(int x, int y);
extern __declspec(dllexport) void QApplication_Beep();
extern __declspec(dllexport) void QApplication_Alert(QWidget* widget);
extern __declspec(dllexport) void QApplication_SetCursorFlashTime(int cursorFlashTime);
extern __declspec(dllexport) int QApplication_CursorFlashTime();
extern __declspec(dllexport) void QApplication_SetDoubleClickInterval(int doubleClickInterval);
extern __declspec(dllexport) int QApplication_DoubleClickInterval();
extern __declspec(dllexport) void QApplication_SetKeyboardInputInterval(int keyboardInputInterval);
extern __declspec(dllexport) int QApplication_KeyboardInputInterval();
extern __declspec(dllexport) void QApplication_SetWheelScrollLines(int wheelScrollLines);
extern __declspec(dllexport) int QApplication_WheelScrollLines();
extern __declspec(dllexport) void QApplication_SetStartDragTime(int ms);
extern __declspec(dllexport) int QApplication_StartDragTime();
extern __declspec(dllexport) void QApplication_SetStartDragDistance(int l);
extern __declspec(dllexport) int QApplication_StartDragDistance();
extern __declspec(dllexport) bool QApplication_IsEffectEnabled(int param1);
extern __declspec(dllexport) void QApplication_SetEffectEnabled(int param1);
extern __declspec(dllexport) int QApplication_Exec(QApplication * self);
extern __declspec(dllexport) bool QApplication_Notify(QApplication* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) void QApplication_FocusChanged(QApplication* self, QWidget* old, QWidget* now);
void QApplication_connect_FocusChanged(QApplication* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QApplication_StyleSheet(const QApplication* self);
extern __declspec(dllexport) bool QApplication_AutoSipEnabled(const QApplication* self);
extern __declspec(dllexport) void QApplication_SetStyleSheet(QApplication* self, struct miqt_string sheet);
extern __declspec(dllexport) void QApplication_SetAutoSipEnabled(QApplication* self, const bool enabled);
extern __declspec(dllexport) void QApplication_CloseAllWindows();
extern __declspec(dllexport) void QApplication_AboutQt();
extern __declspec(dllexport) bool QApplication_Event(QApplication* self, QEvent* param1);
extern __declspec(dllexport) struct miqt_string QApplication_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QApplication_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QApplication_SetPalette2(QPalette* param1, const char* className);
extern __declspec(dllexport) void QApplication_SetFont2(QFont* param1, const char* className);
extern __declspec(dllexport) void QApplication_Alert2(QWidget* widget, int duration);
extern __declspec(dllexport) void QApplication_SetEffectEnabled2(int param1, bool enable);
extern __declspec(dllexport) void QApplication_override_virtual_Notify(void* self, intptr_t slot);
bool QApplication_virtualbase_Notify(void* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) void QApplication_override_virtual_Event(void* self, intptr_t slot);
bool QApplication_virtualbase_Event(void* self, QEvent* param1);
extern __declspec(dllexport) void QApplication_Delete(QApplication* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
