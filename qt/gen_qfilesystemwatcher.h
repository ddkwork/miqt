#pragma once
#ifndef MIQT_QT_GEN_QFILESYSTEMWATCHER_H
#define MIQT_QT_GEN_QFILESYSTEMWATCHER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QFileSystemWatcher QFileSystemWatcher;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QFileSystemWatcher* QFileSystemWatcher_new();
extern __declspec(dllexport) 
QFileSystemWatcher* QFileSystemWatcher_new2(struct miqt_array /* of struct miqt_string */  paths);
extern __declspec(dllexport) 
QFileSystemWatcher* QFileSystemWatcher_new3(QObject* parent);
extern __declspec(dllexport) 
QFileSystemWatcher* QFileSystemWatcher_new4(struct miqt_array /* of struct miqt_string */  paths, QObject* parent);
extern __declspec(dllexport) 
void QFileSystemWatcher_virtbase(QFileSystemWatcher* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QFileSystemWatcher_MetaObject(const QFileSystemWatcher* self);
extern __declspec(dllexport) 
void* QFileSystemWatcher_Metacast(QFileSystemWatcher* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QFileSystemWatcher_Tr(const char* s);
extern __declspec(dllexport) 
bool QFileSystemWatcher_AddPath(QFileSystemWatcher* self, struct miqt_string file);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QFileSystemWatcher_AddPaths(QFileSystemWatcher* self, struct miqt_array /* of struct miqt_string */  files);
extern __declspec(dllexport) 
bool QFileSystemWatcher_RemovePath(QFileSystemWatcher* self, struct miqt_string file);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QFileSystemWatcher_RemovePaths(QFileSystemWatcher* self, struct miqt_array /* of struct miqt_string */  files);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QFileSystemWatcher_Files(const QFileSystemWatcher* self);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QFileSystemWatcher_Directories(const QFileSystemWatcher* self);
extern __declspec(dllexport) 
struct miqt_string QFileSystemWatcher_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QFileSystemWatcher_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QFileSystemWatcher_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QFileSystemWatcher_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QFileSystemWatcher_override_virtual_Metacast(void* self, intptr_t slot);
void* QFileSystemWatcher_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QFileSystemWatcher_Delete(QFileSystemWatcher* self, bool isSubclass);

}
