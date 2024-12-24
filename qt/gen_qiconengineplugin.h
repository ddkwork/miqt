#pragma once
#ifndef MIQT_QT_GEN_QICONENGINEPLUGIN_H
#define MIQT_QT_GEN_QICONENGINEPLUGIN_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QIconEngine QIconEngine;
typedef struct QIconEnginePlugin QIconEnginePlugin;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QIconEnginePlugin* QIconEnginePlugin_new();
extern __declspec(dllexport) 
QIconEnginePlugin* QIconEnginePlugin_new2(QObject* parent);
extern __declspec(dllexport) 
void QIconEnginePlugin_virtbase(QIconEnginePlugin* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QIconEnginePlugin_MetaObject(const QIconEnginePlugin* self);
extern __declspec(dllexport) 
void* QIconEnginePlugin_Metacast(QIconEnginePlugin* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QIconEnginePlugin_Tr(const char* s);
extern __declspec(dllexport) 
QIconEngine* QIconEnginePlugin_Create(QIconEnginePlugin* self, struct miqt_string filename);
extern __declspec(dllexport) 
struct miqt_string QIconEnginePlugin_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QIconEnginePlugin_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QIconEnginePlugin_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QIconEnginePlugin_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QIconEnginePlugin_override_virtual_Metacast(void* self, intptr_t slot);
void* QIconEnginePlugin_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QIconEnginePlugin_Delete(QIconEnginePlugin* self, bool isSubclass);

}
