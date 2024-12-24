#pragma once
#ifndef MIQT_QT_GEN_QSTYLEPLUGIN_H
#define MIQT_QT_GEN_QSTYLEPLUGIN_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QStyle QStyle;
typedef struct QStylePlugin QStylePlugin;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QStylePlugin* QStylePlugin_new();
extern __declspec(dllexport) 
QStylePlugin* QStylePlugin_new2(QObject* parent);
extern __declspec(dllexport) 
void QStylePlugin_virtbase(QStylePlugin* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QStylePlugin_MetaObject(const QStylePlugin* self);
extern __declspec(dllexport) 
void* QStylePlugin_Metacast(QStylePlugin* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QStylePlugin_Tr(const char* s);
extern __declspec(dllexport) 
QStyle* QStylePlugin_Create(QStylePlugin* self, struct miqt_string key);
extern __declspec(dllexport) 
struct miqt_string QStylePlugin_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QStylePlugin_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QStylePlugin_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QStylePlugin_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QStylePlugin_override_virtual_Metacast(void* self, intptr_t slot);
void* QStylePlugin_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QStylePlugin_Delete(QStylePlugin* self, bool isSubclass);

}
