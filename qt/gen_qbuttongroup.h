#pragma once
#ifndef MIQT_QT_GEN_QBUTTONGROUP_H
#define MIQT_QT_GEN_QBUTTONGROUP_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractButton;
class QButtonGroup;
class QChildEvent;
class QEvent;
class QMetaMethod;
class QMetaObject;
class QObject;
class QTimerEvent;
class _GUID;
class type_info;
#else
typedef struct QAbstractButton QAbstractButton;
typedef struct QButtonGroup QButtonGroup;
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTimerEvent QTimerEvent;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QButtonGroup* QButtonGroup_new();
extern __declspec(dllexport) QButtonGroup* QButtonGroup_new2(QObject* parent);
extern __declspec(dllexport) void QButtonGroup_virtbase(QButtonGroup* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QButtonGroup_MetaObject(const QButtonGroup* self);
extern __declspec(dllexport) void* QButtonGroup_Metacast(QButtonGroup* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QButtonGroup_Tr(const char* s);
extern __declspec(dllexport) void QButtonGroup_SetExclusive(QButtonGroup* self, bool exclusive);
extern __declspec(dllexport) bool QButtonGroup_Exclusive(const QButtonGroup* self);
extern __declspec(dllexport) void QButtonGroup_AddButton(QButtonGroup* self, QAbstractButton* param1);
extern __declspec(dllexport) void QButtonGroup_RemoveButton(QButtonGroup* self, QAbstractButton* param1);
extern __declspec(dllexport) struct miqt_array /* of QAbstractButton* */  QButtonGroup_Buttons(const QButtonGroup* self);
extern __declspec(dllexport) QAbstractButton* QButtonGroup_CheckedButton(const QButtonGroup* self);
extern __declspec(dllexport) QAbstractButton* QButtonGroup_Button(const QButtonGroup* self, int id);
extern __declspec(dllexport) void QButtonGroup_SetId(QButtonGroup* self, QAbstractButton* button, int id);
extern __declspec(dllexport) int QButtonGroup_Id(const QButtonGroup* self, QAbstractButton* button);
extern __declspec(dllexport) int QButtonGroup_CheckedId(const QButtonGroup* self);
extern __declspec(dllexport) void QButtonGroup_ButtonClicked(QButtonGroup* self, QAbstractButton* param1);
void QButtonGroup_connect_ButtonClicked(QButtonGroup* self, intptr_t slot);
extern __declspec(dllexport) void QButtonGroup_ButtonPressed(QButtonGroup* self, QAbstractButton* param1);
void QButtonGroup_connect_ButtonPressed(QButtonGroup* self, intptr_t slot);
extern __declspec(dllexport) void QButtonGroup_ButtonReleased(QButtonGroup* self, QAbstractButton* param1);
void QButtonGroup_connect_ButtonReleased(QButtonGroup* self, intptr_t slot);
extern __declspec(dllexport) void QButtonGroup_ButtonToggled(QButtonGroup* self, QAbstractButton* param1, bool param2);
void QButtonGroup_connect_ButtonToggled(QButtonGroup* self, intptr_t slot);
extern __declspec(dllexport) void QButtonGroup_IdClicked(QButtonGroup* self, int param1);
void QButtonGroup_connect_IdClicked(QButtonGroup* self, intptr_t slot);
extern __declspec(dllexport) void QButtonGroup_IdPressed(QButtonGroup* self, int param1);
void QButtonGroup_connect_IdPressed(QButtonGroup* self, intptr_t slot);
extern __declspec(dllexport) void QButtonGroup_IdReleased(QButtonGroup* self, int param1);
void QButtonGroup_connect_IdReleased(QButtonGroup* self, intptr_t slot);
extern __declspec(dllexport) void QButtonGroup_IdToggled(QButtonGroup* self, int param1, bool param2);
void QButtonGroup_connect_IdToggled(QButtonGroup* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QButtonGroup_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QButtonGroup_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QButtonGroup_AddButton2(QButtonGroup* self, QAbstractButton* param1, int id);
extern __declspec(dllexport) void QButtonGroup_override_virtual_Event(void* self, intptr_t slot);
bool QButtonGroup_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QButtonGroup_override_virtual_EventFilter(void* self, intptr_t slot);
bool QButtonGroup_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QButtonGroup_override_virtual_TimerEvent(void* self, intptr_t slot);
void QButtonGroup_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QButtonGroup_override_virtual_ChildEvent(void* self, intptr_t slot);
void QButtonGroup_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QButtonGroup_override_virtual_CustomEvent(void* self, intptr_t slot);
void QButtonGroup_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QButtonGroup_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QButtonGroup_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QButtonGroup_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QButtonGroup_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QButtonGroup_Delete(QButtonGroup* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
