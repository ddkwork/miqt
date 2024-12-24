#pragma once
#ifndef MIQT_QT_GEN_QWIZARD_H
#define MIQT_QT_GEN_QWIZARD_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractButton QAbstractButton;
typedef struct QDialog QDialog;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPixmap QPixmap;
typedef struct QResizeEvent QResizeEvent;
typedef struct QSize QSize;
typedef struct QVariant QVariant;
typedef struct QWidget QWidget;
typedef struct QWizard QWizard;
typedef struct QWizardPage QWizardPage;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QWizard* QWizard_new(QWidget* parent);
extern __declspec(dllexport) 
QWizard* QWizard_new2();
extern __declspec(dllexport) 
QWizard* QWizard_new3(QWidget* parent, int flags);
extern __declspec(dllexport) 
void QWizard_virtbase(QWizard* src
, QDialog** outptr_QDialog
);
extern __declspec(dllexport) 
QMetaObject* QWizard_MetaObject(const QWizard* self);
extern __declspec(dllexport) 
void* QWizard_Metacast(QWizard* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QWizard_Tr(const char* s);
extern __declspec(dllexport) 
int QWizard_AddPage(QWizard* self, QWizardPage* page);
extern __declspec(dllexport) 
void QWizard_SetPage(QWizard* self, int id, QWizardPage* page);
extern __declspec(dllexport) 
void QWizard_RemovePage(QWizard* self, int id);
extern __declspec(dllexport) 
QWizardPage* QWizard_Page(const QWizard* self, int id);
extern __declspec(dllexport) 
bool QWizard_HasVisitedPage(const QWizard* self, int id);
extern __declspec(dllexport) 
struct miqt_array /* of int */  QWizard_VisitedIds(const QWizard* self);
extern __declspec(dllexport) 
struct miqt_array /* of int */  QWizard_PageIds(const QWizard* self);
extern __declspec(dllexport) 
void QWizard_SetStartId(QWizard* self, int id);
extern __declspec(dllexport) 
int QWizard_StartId(const QWizard* self);
extern __declspec(dllexport) 
QWizardPage* QWizard_CurrentPage(const QWizard* self);
extern __declspec(dllexport) 
int QWizard_CurrentId(const QWizard* self);
extern __declspec(dllexport) 
bool QWizard_ValidateCurrentPage(QWizard* self);
extern __declspec(dllexport) 
int QWizard_NextId(const QWizard* self);
extern __declspec(dllexport) 
void QWizard_SetField(QWizard* self, struct miqt_string name, QVariant* value);
extern __declspec(dllexport) 
QVariant* QWizard_Field(const QWizard* self, struct miqt_string name);
extern __declspec(dllexport) 
void QWizard_SetWizardStyle(QWizard* self, WizardStyle style);
extern __declspec(dllexport) 
WizardStyle QWizard_WizardStyle(const QWizard* self);
extern __declspec(dllexport) 
void QWizard_SetOption(QWizard* self, WizardOption option);
extern __declspec(dllexport) 
bool QWizard_TestOption(const QWizard* self, WizardOption option);
extern __declspec(dllexport) 
void QWizard_SetOptions(QWizard* self, WizardOptions options);
extern __declspec(dllexport) 
WizardOptions QWizard_Options(const QWizard* self);
extern __declspec(dllexport) 
void QWizard_SetButtonText(QWizard* self, WizardButton which, struct miqt_string text);
extern __declspec(dllexport) 
struct miqt_string QWizard_ButtonText(const QWizard* self, WizardButton which);
extern __declspec(dllexport) 
void QWizard_SetButtonLayout(QWizard* self, struct miqt_array /* of WizardButton */  layout);
extern __declspec(dllexport) 
void QWizard_SetButton(QWizard* self, WizardButton which, QAbstractButton* button);
extern __declspec(dllexport) 
QAbstractButton* QWizard_Button(const QWizard* self, WizardButton which);
extern __declspec(dllexport) 
void QWizard_SetTitleFormat(QWizard* self, int format);
extern __declspec(dllexport) 
int QWizard_TitleFormat(const QWizard* self);
extern __declspec(dllexport) 
void QWizard_SetSubTitleFormat(QWizard* self, int format);
extern __declspec(dllexport) 
int QWizard_SubTitleFormat(const QWizard* self);
extern __declspec(dllexport) 
void QWizard_SetPixmap(QWizard* self, WizardPixmap which, QPixmap* pixmap);
extern __declspec(dllexport) 
QPixmap* QWizard_Pixmap(const QWizard* self, WizardPixmap which);
extern __declspec(dllexport) 
void QWizard_SetSideWidget(QWizard* self, QWidget* widget);
extern __declspec(dllexport) 
QWidget* QWizard_SideWidget(const QWizard* self);
extern __declspec(dllexport) 
void QWizard_SetDefaultProperty(QWizard* self, const char* className, const char* property, const char* changedSignal);
extern __declspec(dllexport) 
void QWizard_SetVisible(QWizard* self, bool visible);
extern __declspec(dllexport) 
QSize* QWizard_SizeHint(const QWizard* self);
extern __declspec(dllexport) 
void QWizard_CurrentIdChanged(QWizard* self, int id);
void QWizard_connect_CurrentIdChanged(QWizard* self, intptr_t slot);
extern __declspec(dllexport) 
void QWizard_HelpRequested(QWizard* self);
void QWizard_connect_HelpRequested(QWizard* self, intptr_t slot);
extern __declspec(dllexport) 
void QWizard_CustomButtonClicked(QWizard* self, int which);
void QWizard_connect_CustomButtonClicked(QWizard* self, intptr_t slot);
extern __declspec(dllexport) 
void QWizard_PageAdded(QWizard* self, int id);
void QWizard_connect_PageAdded(QWizard* self, intptr_t slot);
extern __declspec(dllexport) 
void QWizard_PageRemoved(QWizard* self, int id);
void QWizard_connect_PageRemoved(QWizard* self, intptr_t slot);
extern __declspec(dllexport) 
void QWizard_Back(QWizard* self);
extern __declspec(dllexport) 
void QWizard_Next(QWizard* self);
extern __declspec(dllexport) 
void QWizard_SetCurrentId(QWizard* self, int id);
extern __declspec(dllexport) 
void QWizard_Restart(QWizard* self);
extern __declspec(dllexport) 
bool QWizard_Event(QWizard* self, QEvent* event);
extern __declspec(dllexport) 
void QWizard_ResizeEvent(QWizard* self, QResizeEvent* event);
extern __declspec(dllexport) 
void QWizard_PaintEvent(QWizard* self, QPaintEvent* event);
extern __declspec(dllexport) 
bool QWizard_NativeEvent(QWizard* self, struct miqt_string eventType, void* message, intptr_t* result);
extern __declspec(dllexport) 
void QWizard_Done(QWizard* self, int result);
extern __declspec(dllexport) 
void QWizard_InitializePage(QWizard* self, int id);
extern __declspec(dllexport) 
void QWizard_CleanupPage(QWizard* self, int id);
extern __declspec(dllexport) 
struct miqt_string QWizard_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QWizard_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QWizard_SetOption2(QWizard* self, WizardOption option, bool on);
extern __declspec(dllexport) 
void QWizard_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QWizard_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QWizard_override_virtual_Metacast(void* self, intptr_t slot);
void* QWizard_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QWizard_Delete(QWizard* self, bool isSubclass);

extern __declspec(dllexport) 
QWizardPage* QWizardPage_new(QWidget* parent);
extern __declspec(dllexport) 
QWizardPage* QWizardPage_new2();
extern __declspec(dllexport) 
void QWizardPage_virtbase(QWizardPage* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QWizardPage_MetaObject(const QWizardPage* self);
extern __declspec(dllexport) 
void* QWizardPage_Metacast(QWizardPage* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QWizardPage_Tr(const char* s);
extern __declspec(dllexport) 
void QWizardPage_SetTitle(QWizardPage* self, struct miqt_string title);
extern __declspec(dllexport) 
struct miqt_string QWizardPage_Title(const QWizardPage* self);
extern __declspec(dllexport) 
void QWizardPage_SetSubTitle(QWizardPage* self, struct miqt_string subTitle);
extern __declspec(dllexport) 
struct miqt_string QWizardPage_SubTitle(const QWizardPage* self);
extern __declspec(dllexport) 
void QWizardPage_SetPixmap(QWizardPage* self, int which, QPixmap* pixmap);
extern __declspec(dllexport) 
QPixmap* QWizardPage_Pixmap(const QWizardPage* self, int which);
extern __declspec(dllexport) 
void QWizardPage_SetFinalPage(QWizardPage* self, bool finalPage);
extern __declspec(dllexport) 
bool QWizardPage_IsFinalPage(const QWizardPage* self);
extern __declspec(dllexport) 
void QWizardPage_SetCommitPage(QWizardPage* self, bool commitPage);
extern __declspec(dllexport) 
bool QWizardPage_IsCommitPage(const QWizardPage* self);
extern __declspec(dllexport) 
void QWizardPage_SetButtonText(QWizardPage* self, int which, struct miqt_string text);
extern __declspec(dllexport) 
struct miqt_string QWizardPage_ButtonText(const QWizardPage* self, int which);
extern __declspec(dllexport) 
void QWizardPage_InitializePage(QWizardPage* self);
extern __declspec(dllexport) 
void QWizardPage_CleanupPage(QWizardPage* self);
extern __declspec(dllexport) 
bool QWizardPage_ValidatePage(QWizardPage* self);
extern __declspec(dllexport) 
bool QWizardPage_IsComplete(const QWizardPage* self);
extern __declspec(dllexport) 
int QWizardPage_NextId(const QWizardPage* self);
extern __declspec(dllexport) 
void QWizardPage_CompleteChanged(QWizardPage* self);
void QWizardPage_connect_CompleteChanged(QWizardPage* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QWizardPage_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QWizardPage_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QWizardPage_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QWizardPage_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QWizardPage_override_virtual_Metacast(void* self, intptr_t slot);
void* QWizardPage_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QWizardPage_Delete(QWizardPage* self, bool isSubclass);

}
