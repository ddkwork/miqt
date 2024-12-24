#pragma once
#ifndef MIQT_QT_GEN_QGENERICPLUGIN_H
#define MIQT_QT_GEN_QGENERICPLUGIN_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QGenericPlugin QGenericPlugin;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QGenericPlugin* QGenericPlugin_new();
extern __declspec(dllexport) 
QGenericPlugin* QGenericPlugin_new2(QObject* parent);
extern __declspec(dllexport) 
void QGenericPlugin_virtbase(QGenericPlugin* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QGenericPlugin_MetaObject(const QGenericPlugin* self);
extern __declspec(dllexport) 
void* QGenericPlugin_Metacast(QGenericPlugin* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QGenericPlugin_Tr(const char* s);
extern __declspec(dllexport) 
QObject* QGenericPlugin_Create(QGenericPlugin* self, struct miqt_string name, struct miqt_string spec);
extern __declspec(dllexport) 
struct miqt_string QGenericPlugin_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QGenericPlugin_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QGenericPlugin_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QGenericPlugin_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QGenericPlugin_override_virtual_Metacast(void* self, intptr_t slot);
void* QGenericPlugin_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QGenericPlugin_Delete(QGenericPlugin* self, bool isSubclass);

}
