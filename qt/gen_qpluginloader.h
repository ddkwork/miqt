#pragma once
#ifndef MIQT_QT_GEN_QPLUGINLOADER_H
#define MIQT_QT_GEN_QPLUGINLOADER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QJsonObject QJsonObject;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPluginLoader QPluginLoader;
typedef struct QStaticPlugin QStaticPlugin;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QPluginLoader* QPluginLoader_new();
extern __declspec(dllexport) 
QPluginLoader* QPluginLoader_new2(struct miqt_string fileName);
extern __declspec(dllexport) 
QPluginLoader* QPluginLoader_new3(QObject* parent);
extern __declspec(dllexport) 
QPluginLoader* QPluginLoader_new4(struct miqt_string fileName, QObject* parent);
extern __declspec(dllexport) 
void QPluginLoader_virtbase(QPluginLoader* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QPluginLoader_MetaObject(const QPluginLoader* self);
extern __declspec(dllexport) 
void* QPluginLoader_Metacast(QPluginLoader* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QPluginLoader_Tr(const char* s);
extern __declspec(dllexport) 
QObject* QPluginLoader_Instance(QPluginLoader* self);
extern __declspec(dllexport) 
QJsonObject* QPluginLoader_MetaData(const QPluginLoader* self);
extern __declspec(dllexport) 
struct miqt_array /* of QObject* */  QPluginLoader_StaticInstances();
extern __declspec(dllexport) 
struct miqt_array /* of QStaticPlugin* */  QPluginLoader_StaticPlugins();
extern __declspec(dllexport) 
bool QPluginLoader_Load(QPluginLoader* self);
extern __declspec(dllexport) 
bool QPluginLoader_Unload(QPluginLoader* self);
extern __declspec(dllexport) 
bool QPluginLoader_IsLoaded(const QPluginLoader* self);
extern __declspec(dllexport) 
void QPluginLoader_SetFileName(QPluginLoader* self, struct miqt_string fileName);
extern __declspec(dllexport) 
struct miqt_string QPluginLoader_FileName(const QPluginLoader* self);
extern __declspec(dllexport) 
struct miqt_string QPluginLoader_ErrorString(const QPluginLoader* self);
extern __declspec(dllexport) 
void QPluginLoader_SetLoadHints(QPluginLoader* self, int loadHints);
extern __declspec(dllexport) 
int QPluginLoader_LoadHints(const QPluginLoader* self);
extern __declspec(dllexport) 
struct miqt_string QPluginLoader_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QPluginLoader_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QPluginLoader_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QPluginLoader_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QPluginLoader_override_virtual_Metacast(void* self, intptr_t slot);
void* QPluginLoader_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QPluginLoader_Delete(QPluginLoader* self, bool isSubclass);

}
