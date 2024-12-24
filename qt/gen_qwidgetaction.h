#pragma once
#ifndef MIQT_QT_GEN_QWIDGETACTION_H
#define MIQT_QT_GEN_QWIDGETACTION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAction QAction;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QWidget QWidget;
typedef struct QWidgetAction QWidgetAction;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QWidgetAction* QWidgetAction_new(QObject* parent);
extern __declspec(dllexport) 
void QWidgetAction_virtbase(QWidgetAction* src
, QAction** outptr_QAction
);
extern __declspec(dllexport) 
QMetaObject* QWidgetAction_MetaObject(const QWidgetAction* self);
extern __declspec(dllexport) 
void* QWidgetAction_Metacast(QWidgetAction* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QWidgetAction_Tr(const char* s);
extern __declspec(dllexport) 
void QWidgetAction_SetDefaultWidget(QWidgetAction* self, QWidget* w);
extern __declspec(dllexport) 
QWidget* QWidgetAction_DefaultWidget(const QWidgetAction* self);
extern __declspec(dllexport) 
QWidget* QWidgetAction_RequestWidget(QWidgetAction* self, QWidget* parent);
extern __declspec(dllexport) 
void QWidgetAction_ReleaseWidget(QWidgetAction* self, QWidget* widget);
extern __declspec(dllexport) 
bool QWidgetAction_Event(QWidgetAction* self, QEvent* param1);
extern __declspec(dllexport) 
bool QWidgetAction_EventFilter(QWidgetAction* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) 
QWidget* QWidgetAction_CreateWidget(QWidgetAction* self, QWidget* parent);
extern __declspec(dllexport) 
void QWidgetAction_DeleteWidget(QWidgetAction* self, QWidget* widget);
extern __declspec(dllexport) 
struct miqt_string QWidgetAction_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QWidgetAction_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QWidgetAction_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QWidgetAction_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QWidgetAction_override_virtual_Metacast(void* self, intptr_t slot);
void* QWidgetAction_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QWidgetAction_Delete(QWidgetAction* self, bool isSubclass);

}
