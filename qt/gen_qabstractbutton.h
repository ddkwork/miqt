#pragma once
#ifndef MIQT_QT_GEN_QABSTRACTBUTTON_H
#define MIQT_QT_GEN_QABSTRACTBUTTON_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractButton QAbstractButton;
typedef struct QButtonGroup QButtonGroup;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QIcon QIcon;
typedef struct QKeyEvent QKeyEvent;
typedef struct QKeySequence QKeySequence;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPoint QPoint;
typedef struct QSize QSize;
typedef struct QTimerEvent QTimerEvent;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAbstractButton* QAbstractButton_new(QWidget* parent);
extern __declspec(dllexport) 
QAbstractButton* QAbstractButton_new2();
extern __declspec(dllexport) 
void QAbstractButton_virtbase(QAbstractButton* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QAbstractButton_MetaObject(const QAbstractButton* self);
extern __declspec(dllexport) 
void* QAbstractButton_Metacast(QAbstractButton* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAbstractButton_Tr(const char* s);
extern __declspec(dllexport) 
void QAbstractButton_SetText(QAbstractButton* self, struct miqt_string text);
extern __declspec(dllexport) 
struct miqt_string QAbstractButton_Text(const QAbstractButton* self);
extern __declspec(dllexport) 
void QAbstractButton_SetIcon(QAbstractButton* self, QIcon* icon);
extern __declspec(dllexport) 
QIcon* QAbstractButton_Icon(const QAbstractButton* self);
extern __declspec(dllexport) 
QSize* QAbstractButton_IconSize(const QAbstractButton* self);
extern __declspec(dllexport) 
void QAbstractButton_SetShortcut(QAbstractButton* self, QKeySequence* key);
extern __declspec(dllexport) 
QKeySequence* QAbstractButton_Shortcut(const QAbstractButton* self);
extern __declspec(dllexport) 
void QAbstractButton_SetCheckable(QAbstractButton* self, bool checkable);
extern __declspec(dllexport) 
bool QAbstractButton_IsCheckable(const QAbstractButton* self);
extern __declspec(dllexport) 
bool QAbstractButton_IsChecked(const QAbstractButton* self);
extern __declspec(dllexport) 
void QAbstractButton_SetDown(QAbstractButton* self, bool down);
extern __declspec(dllexport) 
bool QAbstractButton_IsDown(const QAbstractButton* self);
extern __declspec(dllexport) 
void QAbstractButton_SetAutoRepeat(QAbstractButton* self, bool autoRepeat);
extern __declspec(dllexport) 
bool QAbstractButton_AutoRepeat(const QAbstractButton* self);
extern __declspec(dllexport) 
void QAbstractButton_SetAutoRepeatDelay(QAbstractButton* self, int autoRepeatDelay);
extern __declspec(dllexport) 
int QAbstractButton_AutoRepeatDelay(const QAbstractButton* self);
extern __declspec(dllexport) 
void QAbstractButton_SetAutoRepeatInterval(QAbstractButton* self, int autoRepeatInterval);
extern __declspec(dllexport) 
int QAbstractButton_AutoRepeatInterval(const QAbstractButton* self);
extern __declspec(dllexport) 
void QAbstractButton_SetAutoExclusive(QAbstractButton* self, bool autoExclusive);
extern __declspec(dllexport) 
bool QAbstractButton_AutoExclusive(const QAbstractButton* self);
extern __declspec(dllexport) 
QButtonGroup* QAbstractButton_Group(const QAbstractButton* self);
extern __declspec(dllexport) 
void QAbstractButton_SetIconSize(QAbstractButton* self, QSize* size);
extern __declspec(dllexport) 
void QAbstractButton_AnimateClick(QAbstractButton* self);
extern __declspec(dllexport) 
void QAbstractButton_Click(QAbstractButton* self);
extern __declspec(dllexport) 
void QAbstractButton_Toggle(QAbstractButton* self);
extern __declspec(dllexport) 
void QAbstractButton_SetChecked(QAbstractButton* self, bool checked);
extern __declspec(dllexport) 
void QAbstractButton_Pressed(QAbstractButton* self);
void QAbstractButton_connect_Pressed(QAbstractButton* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractButton_Released(QAbstractButton* self);
void QAbstractButton_connect_Released(QAbstractButton* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractButton_Clicked(QAbstractButton* self);
void QAbstractButton_connect_Clicked(QAbstractButton* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractButton_Toggled(QAbstractButton* self, bool checked);
void QAbstractButton_connect_Toggled(QAbstractButton* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractButton_PaintEvent(QAbstractButton* self, QPaintEvent* e);
extern __declspec(dllexport) 
bool QAbstractButton_HitButton(const QAbstractButton* self, QPoint* pos);
extern __declspec(dllexport) 
void QAbstractButton_CheckStateSet(QAbstractButton* self);
extern __declspec(dllexport) 
void QAbstractButton_NextCheckState(QAbstractButton* self);
extern __declspec(dllexport) 
bool QAbstractButton_Event(QAbstractButton* self, QEvent* e);
extern __declspec(dllexport) 
void QAbstractButton_KeyPressEvent(QAbstractButton* self, QKeyEvent* e);
extern __declspec(dllexport) 
void QAbstractButton_KeyReleaseEvent(QAbstractButton* self, QKeyEvent* e);
extern __declspec(dllexport) 
void QAbstractButton_MousePressEvent(QAbstractButton* self, QMouseEvent* e);
extern __declspec(dllexport) 
void QAbstractButton_MouseReleaseEvent(QAbstractButton* self, QMouseEvent* e);
extern __declspec(dllexport) 
void QAbstractButton_MouseMoveEvent(QAbstractButton* self, QMouseEvent* e);
extern __declspec(dllexport) 
void QAbstractButton_FocusInEvent(QAbstractButton* self, QFocusEvent* e);
extern __declspec(dllexport) 
void QAbstractButton_FocusOutEvent(QAbstractButton* self, QFocusEvent* e);
extern __declspec(dllexport) 
void QAbstractButton_ChangeEvent(QAbstractButton* self, QEvent* e);
extern __declspec(dllexport) 
void QAbstractButton_TimerEvent(QAbstractButton* self, QTimerEvent* e);
extern __declspec(dllexport) 
struct miqt_string QAbstractButton_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAbstractButton_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAbstractButton_Clicked1(QAbstractButton* self, bool checked);
void QAbstractButton_connect_Clicked1(QAbstractButton* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractButton_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAbstractButton_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAbstractButton_override_virtual_Metacast(void* self, intptr_t slot);
void* QAbstractButton_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAbstractButton_Delete(QAbstractButton* self, bool isSubclass);

}
