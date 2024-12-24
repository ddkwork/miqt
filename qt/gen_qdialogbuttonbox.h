#pragma once
#ifndef MIQT_QT_GEN_QDIALOGBUTTONBOX_H
#define MIQT_QT_GEN_QDIALOGBUTTONBOX_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractButton QAbstractButton;
typedef struct QDialogButtonBox QDialogButtonBox;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPushButton QPushButton;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QDialogButtonBox* QDialogButtonBox_new(QWidget* parent);
extern __declspec(dllexport) 
QDialogButtonBox* QDialogButtonBox_new2();
extern __declspec(dllexport) 
QDialogButtonBox* QDialogButtonBox_new3(int orientation);
extern __declspec(dllexport) 
QDialogButtonBox* QDialogButtonBox_new4(StandardButtons buttons);
extern __declspec(dllexport) 
QDialogButtonBox* QDialogButtonBox_new5(StandardButtons buttons, int orientation);
extern __declspec(dllexport) 
QDialogButtonBox* QDialogButtonBox_new6(int orientation, QWidget* parent);
extern __declspec(dllexport) 
QDialogButtonBox* QDialogButtonBox_new7(StandardButtons buttons, QWidget* parent);
extern __declspec(dllexport) 
QDialogButtonBox* QDialogButtonBox_new8(StandardButtons buttons, int orientation, QWidget* parent);
extern __declspec(dllexport) 
void QDialogButtonBox_virtbase(QDialogButtonBox* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QDialogButtonBox_MetaObject(const QDialogButtonBox* self);
extern __declspec(dllexport) 
void* QDialogButtonBox_Metacast(QDialogButtonBox* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QDialogButtonBox_Tr(const char* s);
extern __declspec(dllexport) 
void QDialogButtonBox_SetOrientation(QDialogButtonBox* self, int orientation);
extern __declspec(dllexport) 
int QDialogButtonBox_Orientation(const QDialogButtonBox* self);
extern __declspec(dllexport) 
void QDialogButtonBox_AddButton(QDialogButtonBox* self, QAbstractButton* button, ButtonRole role);
extern __declspec(dllexport) 
QPushButton* QDialogButtonBox_AddButton2(QDialogButtonBox* self, struct miqt_string text, ButtonRole role);
extern __declspec(dllexport) 
QPushButton* QDialogButtonBox_AddButtonWithButton(QDialogButtonBox* self, StandardButton button);
extern __declspec(dllexport) 
void QDialogButtonBox_RemoveButton(QDialogButtonBox* self, QAbstractButton* button);
extern __declspec(dllexport) 
void QDialogButtonBox_Clear(QDialogButtonBox* self);
extern __declspec(dllexport) 
struct miqt_array /* of QAbstractButton* */  QDialogButtonBox_Buttons(const QDialogButtonBox* self);
extern __declspec(dllexport) 
ButtonRole QDialogButtonBox_ButtonRole(const QDialogButtonBox* self, QAbstractButton* button);
extern __declspec(dllexport) 
void QDialogButtonBox_SetStandardButtons(QDialogButtonBox* self, StandardButtons buttons);
extern __declspec(dllexport) 
StandardButtons QDialogButtonBox_StandardButtons(const QDialogButtonBox* self);
extern __declspec(dllexport) 
StandardButton QDialogButtonBox_StandardButton(const QDialogButtonBox* self, QAbstractButton* button);
extern __declspec(dllexport) 
QPushButton* QDialogButtonBox_Button(const QDialogButtonBox* self, StandardButton which);
extern __declspec(dllexport) 
void QDialogButtonBox_SetCenterButtons(QDialogButtonBox* self, bool center);
extern __declspec(dllexport) 
bool QDialogButtonBox_CenterButtons(const QDialogButtonBox* self);
extern __declspec(dllexport) 
void QDialogButtonBox_Clicked(QDialogButtonBox* self, QAbstractButton* button);
void QDialogButtonBox_connect_Clicked(QDialogButtonBox* self, intptr_t slot);
extern __declspec(dllexport) 
void QDialogButtonBox_Accepted(QDialogButtonBox* self);
void QDialogButtonBox_connect_Accepted(QDialogButtonBox* self, intptr_t slot);
extern __declspec(dllexport) 
void QDialogButtonBox_HelpRequested(QDialogButtonBox* self);
void QDialogButtonBox_connect_HelpRequested(QDialogButtonBox* self, intptr_t slot);
extern __declspec(dllexport) 
void QDialogButtonBox_Rejected(QDialogButtonBox* self);
void QDialogButtonBox_connect_Rejected(QDialogButtonBox* self, intptr_t slot);
extern __declspec(dllexport) 
void QDialogButtonBox_ChangeEvent(QDialogButtonBox* self, QEvent* event);
extern __declspec(dllexport) 
bool QDialogButtonBox_Event(QDialogButtonBox* self, QEvent* event);
extern __declspec(dllexport) 
struct miqt_string QDialogButtonBox_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QDialogButtonBox_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QDialogButtonBox_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QDialogButtonBox_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QDialogButtonBox_override_virtual_Metacast(void* self, intptr_t slot);
void* QDialogButtonBox_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QDialogButtonBox_Delete(QDialogButtonBox* self, bool isSubclass);

}
