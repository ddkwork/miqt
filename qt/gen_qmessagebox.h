#pragma once
#ifndef MIQT_QT_GEN_QMESSAGEBOX_H
#define MIQT_QT_GEN_QMESSAGEBOX_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractButton QAbstractButton;
typedef struct QCheckBox QCheckBox;
typedef struct QCloseEvent QCloseEvent;
typedef struct QDialog QDialog;
typedef struct QEvent QEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMessageBox QMessageBox;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPixmap QPixmap;
typedef struct QPushButton QPushButton;
typedef struct QResizeEvent QResizeEvent;
typedef struct QShowEvent QShowEvent;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QMessageBox* QMessageBox_new(QWidget* parent);
extern __declspec(dllexport) 
QMessageBox* QMessageBox_new2();
extern __declspec(dllexport) 
QMessageBox* QMessageBox_new3(Icon icon, struct miqt_string title, struct miqt_string text);
extern __declspec(dllexport) 
QMessageBox* QMessageBox_new4(struct miqt_string title, struct miqt_string text, Icon icon, int button0, int button1, int button2);
extern __declspec(dllexport) 
QMessageBox* QMessageBox_new5(Icon icon, struct miqt_string title, struct miqt_string text, StandardButtons buttons);
extern __declspec(dllexport) 
QMessageBox* QMessageBox_new6(Icon icon, struct miqt_string title, struct miqt_string text, StandardButtons buttons, QWidget* parent);
extern __declspec(dllexport) 
QMessageBox* QMessageBox_new7(Icon icon, struct miqt_string title, struct miqt_string text, StandardButtons buttons, QWidget* parent, int flags);
extern __declspec(dllexport) 
QMessageBox* QMessageBox_new8(struct miqt_string title, struct miqt_string text, Icon icon, int button0, int button1, int button2, QWidget* parent);
extern __declspec(dllexport) 
QMessageBox* QMessageBox_new9(struct miqt_string title, struct miqt_string text, Icon icon, int button0, int button1, int button2, QWidget* parent, int f);
extern __declspec(dllexport) 
void QMessageBox_virtbase(QMessageBox* src
, QDialog** outptr_QDialog
);
extern __declspec(dllexport) 
QMetaObject* QMessageBox_MetaObject(const QMessageBox* self);
extern __declspec(dllexport) 
void* QMessageBox_Metacast(QMessageBox* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QMessageBox_Tr(const char* s);
extern __declspec(dllexport) 
void QMessageBox_AddButton(QMessageBox* self, QAbstractButton* button, ButtonRole role);
extern __declspec(dllexport) 
QPushButton* QMessageBox_AddButton2(QMessageBox* self, struct miqt_string text, ButtonRole role);
extern __declspec(dllexport) 
QPushButton* QMessageBox_AddButtonWithButton(QMessageBox* self, StandardButton button);
extern __declspec(dllexport) 
void QMessageBox_RemoveButton(QMessageBox* self, QAbstractButton* button);
extern __declspec(dllexport) 
struct miqt_array /* of QAbstractButton* */  QMessageBox_Buttons(const QMessageBox* self);
extern __declspec(dllexport) 
ButtonRole QMessageBox_ButtonRole(const QMessageBox* self, QAbstractButton* button);
extern __declspec(dllexport) 
void QMessageBox_SetStandardButtons(QMessageBox* self, StandardButtons buttons);
extern __declspec(dllexport) 
StandardButtons QMessageBox_StandardButtons(const QMessageBox* self);
extern __declspec(dllexport) 
StandardButton QMessageBox_StandardButton(const QMessageBox* self, QAbstractButton* button);
extern __declspec(dllexport) 
QAbstractButton* QMessageBox_Button(const QMessageBox* self, StandardButton which);
extern __declspec(dllexport) 
QPushButton* QMessageBox_DefaultButton(const QMessageBox* self);
extern __declspec(dllexport) 
void QMessageBox_SetDefaultButton(QMessageBox* self, QPushButton* button);
extern __declspec(dllexport) 
void QMessageBox_SetDefaultButtonWithButton(QMessageBox* self, StandardButton button);
extern __declspec(dllexport) 
QAbstractButton* QMessageBox_EscapeButton(const QMessageBox* self);
extern __declspec(dllexport) 
void QMessageBox_SetEscapeButton(QMessageBox* self, QAbstractButton* button);
extern __declspec(dllexport) 
void QMessageBox_SetEscapeButtonWithButton(QMessageBox* self, StandardButton button);
extern __declspec(dllexport) 
QAbstractButton* QMessageBox_ClickedButton(const QMessageBox* self);
extern __declspec(dllexport) 
struct miqt_string QMessageBox_Text(const QMessageBox* self);
extern __declspec(dllexport) 
void QMessageBox_SetText(QMessageBox* self, struct miqt_string text);
extern __declspec(dllexport) 
Icon QMessageBox_Icon(const QMessageBox* self);
extern __declspec(dllexport) 
void QMessageBox_SetIcon(QMessageBox* self, Icon icon);
extern __declspec(dllexport) 
QPixmap* QMessageBox_IconPixmap(const QMessageBox* self);
extern __declspec(dllexport) 
void QMessageBox_SetIconPixmap(QMessageBox* self, QPixmap* pixmap);
extern __declspec(dllexport) 
int QMessageBox_TextFormat(const QMessageBox* self);
extern __declspec(dllexport) 
void QMessageBox_SetTextFormat(QMessageBox* self, int format);
extern __declspec(dllexport) 
void QMessageBox_SetTextInteractionFlags(QMessageBox* self, int flags);
extern __declspec(dllexport) 
int QMessageBox_TextInteractionFlags(const QMessageBox* self);
extern __declspec(dllexport) 
void QMessageBox_SetCheckBox(QMessageBox* self, QCheckBox* cb);
extern __declspec(dllexport) 
QCheckBox* QMessageBox_CheckBox(const QMessageBox* self);
extern __declspec(dllexport) 
void QMessageBox_SetOption(QMessageBox* self, Option option);
extern __declspec(dllexport) 
bool QMessageBox_TestOption(const QMessageBox* self, Option option);
extern __declspec(dllexport) 
void QMessageBox_SetOptions(QMessageBox* self, Options options);
extern __declspec(dllexport) 
Options QMessageBox_Options(const QMessageBox* self);
extern __declspec(dllexport) 
StandardButton QMessageBox_Information(QWidget* parent, struct miqt_string title, struct miqt_string text);
extern __declspec(dllexport) 
StandardButton QMessageBox_Information2(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButton button0);
extern __declspec(dllexport) 
StandardButton QMessageBox_Question(QWidget* parent, struct miqt_string title, struct miqt_string text);
extern __declspec(dllexport) 
int QMessageBox_Question2(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButton button0, StandardButton button1);
extern __declspec(dllexport) 
StandardButton QMessageBox_Warning(QWidget* parent, struct miqt_string title, struct miqt_string text);
extern __declspec(dllexport) 
int QMessageBox_Warning2(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButton button0, StandardButton button1);
extern __declspec(dllexport) 
StandardButton QMessageBox_Critical(QWidget* parent, struct miqt_string title, struct miqt_string text);
extern __declspec(dllexport) 
int QMessageBox_Critical2(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButton button0, StandardButton button1);
extern __declspec(dllexport) 
void QMessageBox_About(QWidget* parent, struct miqt_string title, struct miqt_string text);
extern __declspec(dllexport) 
void QMessageBox_AboutQt(QWidget* parent);
extern __declspec(dllexport) 
int QMessageBox_Information3(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0);
extern __declspec(dllexport) 
int QMessageBox_Information4(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text);
extern __declspec(dllexport) 
int QMessageBox_Question3(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0);
extern __declspec(dllexport) 
int QMessageBox_Question4(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text);
extern __declspec(dllexport) 
int QMessageBox_Warning3(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0, int button1);
extern __declspec(dllexport) 
int QMessageBox_Warning4(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text);
extern __declspec(dllexport) 
int QMessageBox_Critical3(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0, int button1);
extern __declspec(dllexport) 
int QMessageBox_Critical4(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text);
extern __declspec(dllexport) 
struct miqt_string QMessageBox_ButtonText(const QMessageBox* self, int button);
extern __declspec(dllexport) 
void QMessageBox_SetButtonText(QMessageBox* self, int button, struct miqt_string text);
extern __declspec(dllexport) 
struct miqt_string QMessageBox_InformativeText(const QMessageBox* self);
extern __declspec(dllexport) 
void QMessageBox_SetInformativeText(QMessageBox* self, struct miqt_string text);
extern __declspec(dllexport) 
struct miqt_string QMessageBox_DetailedText(const QMessageBox* self);
extern __declspec(dllexport) 
void QMessageBox_SetDetailedText(QMessageBox* self, struct miqt_string text);
extern __declspec(dllexport) 
void QMessageBox_SetWindowTitle(QMessageBox* self, struct miqt_string title);
extern __declspec(dllexport) 
void QMessageBox_SetWindowModality(QMessageBox* self, int windowModality);
extern __declspec(dllexport) 
QPixmap* QMessageBox_StandardIcon(Icon icon);
extern __declspec(dllexport) 
void QMessageBox_ButtonClicked(QMessageBox* self, QAbstractButton* button);
void QMessageBox_connect_ButtonClicked(QMessageBox* self, intptr_t slot);
extern __declspec(dllexport) 
bool QMessageBox_Event(QMessageBox* self, QEvent* e);
extern __declspec(dllexport) 
void QMessageBox_ResizeEvent(QMessageBox* self, QResizeEvent* event);
extern __declspec(dllexport) 
void QMessageBox_ShowEvent(QMessageBox* self, QShowEvent* event);
extern __declspec(dllexport) 
void QMessageBox_CloseEvent(QMessageBox* self, QCloseEvent* event);
extern __declspec(dllexport) 
void QMessageBox_KeyPressEvent(QMessageBox* self, QKeyEvent* event);
extern __declspec(dllexport) 
void QMessageBox_ChangeEvent(QMessageBox* self, QEvent* event);
extern __declspec(dllexport) 
struct miqt_string QMessageBox_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QMessageBox_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QMessageBox_SetOption2(QMessageBox* self, Option option, bool on);
extern __declspec(dllexport) 
StandardButton QMessageBox_Information42(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButtons buttons);
extern __declspec(dllexport) 
StandardButton QMessageBox_Information5(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButtons buttons, StandardButton defaultButton);
extern __declspec(dllexport) 
StandardButton QMessageBox_Information52(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButton button0, StandardButton button1);
extern __declspec(dllexport) 
StandardButton QMessageBox_Question42(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButtons buttons);
extern __declspec(dllexport) 
StandardButton QMessageBox_Question5(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButtons buttons, StandardButton defaultButton);
extern __declspec(dllexport) 
StandardButton QMessageBox_Warning42(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButtons buttons);
extern __declspec(dllexport) 
StandardButton QMessageBox_Warning5(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButtons buttons, StandardButton defaultButton);
extern __declspec(dllexport) 
StandardButton QMessageBox_Critical42(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButtons buttons);
extern __declspec(dllexport) 
StandardButton QMessageBox_Critical5(QWidget* parent, struct miqt_string title, struct miqt_string text, StandardButtons buttons, StandardButton defaultButton);
extern __declspec(dllexport) 
void QMessageBox_AboutQt2(QWidget* parent, struct miqt_string title);
extern __declspec(dllexport) 
int QMessageBox_Information53(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0, int button1);
extern __declspec(dllexport) 
int QMessageBox_Information6(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0, int button1, int button2);
extern __declspec(dllexport) 
int QMessageBox_Information54(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text);
extern __declspec(dllexport) 
int QMessageBox_Information62(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text);
extern __declspec(dllexport) 
int QMessageBox_Information7(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text, int defaultButtonNumber);
extern __declspec(dllexport) 
int QMessageBox_Information8(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text, int defaultButtonNumber, int escapeButtonNumber);
extern __declspec(dllexport) 
int QMessageBox_Question52(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0, int button1);
extern __declspec(dllexport) 
int QMessageBox_Question6(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0, int button1, int button2);
extern __declspec(dllexport) 
int QMessageBox_Question53(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text);
extern __declspec(dllexport) 
int QMessageBox_Question62(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text);
extern __declspec(dllexport) 
int QMessageBox_Question7(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text, int defaultButtonNumber);
extern __declspec(dllexport) 
int QMessageBox_Question8(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text, int defaultButtonNumber, int escapeButtonNumber);
extern __declspec(dllexport) 
int QMessageBox_Warning6(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0, int button1, int button2);
extern __declspec(dllexport) 
int QMessageBox_Warning52(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text);
extern __declspec(dllexport) 
int QMessageBox_Warning62(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text);
extern __declspec(dllexport) 
int QMessageBox_Warning7(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text, int defaultButtonNumber);
extern __declspec(dllexport) 
int QMessageBox_Warning8(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text, int defaultButtonNumber, int escapeButtonNumber);
extern __declspec(dllexport) 
int QMessageBox_Critical6(QWidget* parent, struct miqt_string title, struct miqt_string text, int button0, int button1, int button2);
extern __declspec(dllexport) 
int QMessageBox_Critical52(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text);
extern __declspec(dllexport) 
int QMessageBox_Critical62(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text);
extern __declspec(dllexport) 
int QMessageBox_Critical7(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text, int defaultButtonNumber);
extern __declspec(dllexport) 
int QMessageBox_Critical8(QWidget* parent, struct miqt_string title, struct miqt_string text, struct miqt_string button0Text, struct miqt_string button1Text, struct miqt_string button2Text, int defaultButtonNumber, int escapeButtonNumber);
extern __declspec(dllexport) 
void QMessageBox_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QMessageBox_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QMessageBox_override_virtual_Metacast(void* self, intptr_t slot);
void* QMessageBox_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QMessageBox_Delete(QMessageBox* self, bool isSubclass);

}
