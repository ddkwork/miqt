#pragma once
#ifndef MIQT_QT_GEN_QITEMEDITORFACTORY_H
#define MIQT_QT_GEN_QITEMEDITORFACTORY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QItemEditorCreatorBase QItemEditorCreatorBase;
typedef struct QItemEditorFactory QItemEditorFactory;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QWidget* QItemEditorCreatorBase_CreateWidget(const QItemEditorCreatorBase* self, QWidget* parent);
extern __declspec(dllexport) struct miqt_string QItemEditorCreatorBase_ValuePropertyName(const QItemEditorCreatorBase* self);
extern __declspec(dllexport) void QItemEditorCreatorBase_OperatorAssign(QItemEditorCreatorBase* self, QItemEditorCreatorBase* param1);
extern __declspec(dllexport) void QItemEditorCreatorBase_Delete(QItemEditorCreatorBase* self, bool isSubclass);

extern __declspec(dllexport) QItemEditorFactory* QItemEditorFactory_new();
extern __declspec(dllexport) QItemEditorFactory* QItemEditorFactory_new2(QItemEditorFactory* param1);
extern __declspec(dllexport) QWidget* QItemEditorFactory_CreateEditor(const QItemEditorFactory* self, int userType, QWidget* parent);
extern __declspec(dllexport) struct miqt_string QItemEditorFactory_ValuePropertyName(const QItemEditorFactory* self, int userType);
extern __declspec(dllexport) void QItemEditorFactory_RegisterEditor(QItemEditorFactory* self, int userType, QItemEditorCreatorBase* creator);
extern __declspec(dllexport) QItemEditorFactory* QItemEditorFactory_DefaultFactory();
extern __declspec(dllexport) void QItemEditorFactory_SetDefaultFactory(QItemEditorFactory* factory);
extern __declspec(dllexport) void QItemEditorFactory_override_virtual_CreateEditor(void* self, intptr_t slot);
QWidget* QItemEditorFactory_virtualbase_CreateEditor(const void* self, int userType, QWidget* parent);
extern __declspec(dllexport) void QItemEditorFactory_override_virtual_ValuePropertyName(void* self, intptr_t slot);
struct miqt_string QItemEditorFactory_virtualbase_ValuePropertyName(const void* self, int userType);
extern __declspec(dllexport) void QItemEditorFactory_Delete(QItemEditorFactory* self, bool isSubclass);

} 
