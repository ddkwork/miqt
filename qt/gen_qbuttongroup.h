#pragma once
#ifndef MIQT_QT_GEN_QBUTTONGROUP_H
#define MIQT_QT_GEN_QBUTTONGROUP_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractButton QAbstractButton;
typedef struct QButtonGroup QButtonGroup;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QButtonGroup* QButtonGroup_new();
extern __declspec(dllexport) 
QButtonGroup* QButtonGroup_new2(QObject* parent);
extern __declspec(dllexport) 
void QButtonGroup_virtbase(QButtonGroup* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QButtonGroup_MetaObject(const QButtonGroup* self);
extern __declspec(dllexport) 
void* QButtonGroup_Metacast(QButtonGroup* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QButtonGroup_Tr(const char* s);
extern __declspec(dllexport) 
void QButtonGroup_SetExclusive(QButtonGroup* self, bool exclusive);
extern __declspec(dllexport) 
bool QButtonGroup_Exclusive(const QButtonGroup* self);
extern __declspec(dllexport) 
void QButtonGroup_AddButton(QButtonGroup* self, QAbstractButton* param1);
extern __declspec(dllexport) 
void QButtonGroup_RemoveButton(QButtonGroup* self, QAbstractButton* param1);
extern __declspec(dllexport) 
struct miqt_array /* of QAbstractButton* */  QButtonGroup_Buttons(const QButtonGroup* self);
extern __declspec(dllexport) 
QAbstractButton* QButtonGroup_CheckedButton(const QButtonGroup* self);
extern __declspec(dllexport) 
QAbstractButton* QButtonGroup_Button(const QButtonGroup* self, int id);
extern __declspec(dllexport) 
void QButtonGroup_SetId(QButtonGroup* self, QAbstractButton* button, int id);
extern __declspec(dllexport) 
int QButtonGroup_Id(const QButtonGroup* self, QAbstractButton* button);
extern __declspec(dllexport) 
int QButtonGroup_CheckedId(const QButtonGroup* self);
extern __declspec(dllexport) 
void QButtonGroup_ButtonClicked(QButtonGroup* self, QAbstractButton* param1);
void QButtonGroup_connect_ButtonClicked(QButtonGroup* self, intptr_t slot);
extern __declspec(dllexport) 
void QButtonGroup_ButtonPressed(QButtonGroup* self, QAbstractButton* param1);
void QButtonGroup_connect_ButtonPressed(QButtonGroup* self, intptr_t slot);
extern __declspec(dllexport) 
void QButtonGroup_ButtonReleased(QButtonGroup* self, QAbstractButton* param1);
void QButtonGroup_connect_ButtonReleased(QButtonGroup* self, intptr_t slot);
extern __declspec(dllexport) 
void QButtonGroup_ButtonToggled(QButtonGroup* self, QAbstractButton* param1, bool param2);
void QButtonGroup_connect_ButtonToggled(QButtonGroup* self, intptr_t slot);
extern __declspec(dllexport) 
void QButtonGroup_IdClicked(QButtonGroup* self, int param1);
void QButtonGroup_connect_IdClicked(QButtonGroup* self, intptr_t slot);
extern __declspec(dllexport) 
void QButtonGroup_IdPressed(QButtonGroup* self, int param1);
void QButtonGroup_connect_IdPressed(QButtonGroup* self, intptr_t slot);
extern __declspec(dllexport) 
void QButtonGroup_IdReleased(QButtonGroup* self, int param1);
void QButtonGroup_connect_IdReleased(QButtonGroup* self, intptr_t slot);
extern __declspec(dllexport) 
void QButtonGroup_IdToggled(QButtonGroup* self, int param1, bool param2);
void QButtonGroup_connect_IdToggled(QButtonGroup* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QButtonGroup_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QButtonGroup_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QButtonGroup_AddButton2(QButtonGroup* self, QAbstractButton* param1, int id);
extern __declspec(dllexport) 
void QButtonGroup_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QButtonGroup_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QButtonGroup_override_virtual_Metacast(void* self, intptr_t slot);
void* QButtonGroup_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QButtonGroup_Delete(QButtonGroup* self, bool isSubclass);

}
