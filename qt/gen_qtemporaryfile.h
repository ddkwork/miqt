#pragma once
#ifndef MIQT_QT_GEN_QTEMPORARYFILE_H
#define MIQT_QT_GEN_QTEMPORARYFILE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QFile QFile;
typedef struct QFileDevice QFileDevice;
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTemporaryFile QTemporaryFile;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QTemporaryFile* QTemporaryFile_new();
extern __declspec(dllexport) 
QTemporaryFile* QTemporaryFile_new2(struct miqt_string templateName);
extern __declspec(dllexport) 
QTemporaryFile* QTemporaryFile_new3(QObject* parent);
extern __declspec(dllexport) 
QTemporaryFile* QTemporaryFile_new4(struct miqt_string templateName, QObject* parent);
extern __declspec(dllexport) 
void QTemporaryFile_virtbase(QTemporaryFile* src
, QFile** outptr_QFile
);
extern __declspec(dllexport) 
QMetaObject* QTemporaryFile_MetaObject(const QTemporaryFile* self);
extern __declspec(dllexport) 
void* QTemporaryFile_Metacast(QTemporaryFile* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QTemporaryFile_Tr(const char* s);
extern __declspec(dllexport) 
bool QTemporaryFile_AutoRemove(const QTemporaryFile* self);
extern __declspec(dllexport) 
void QTemporaryFile_SetAutoRemove(QTemporaryFile* self, bool b);
extern __declspec(dllexport) 
bool QTemporaryFile_Open(QTemporaryFile* self);
extern __declspec(dllexport) 
struct miqt_string QTemporaryFile_FileName(const QTemporaryFile* self);
extern __declspec(dllexport) 
struct miqt_string QTemporaryFile_FileTemplate(const QTemporaryFile* self);
extern __declspec(dllexport) 
void QTemporaryFile_SetFileTemplate(QTemporaryFile* self, struct miqt_string name);
extern __declspec(dllexport) 
bool QTemporaryFile_Rename(QTemporaryFile* self, struct miqt_string newName);
extern __declspec(dllexport) 
QTemporaryFile* QTemporaryFile_CreateNativeFile(struct miqt_string fileName);
extern __declspec(dllexport) 
QTemporaryFile* QTemporaryFile_CreateNativeFileWithFile(QFile* file);
extern __declspec(dllexport) 
bool QTemporaryFile_OpenWithFlags(QTemporaryFile* self, OpenMode flags);
extern __declspec(dllexport) 
struct miqt_string QTemporaryFile_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QTemporaryFile_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QTemporaryFile_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QTemporaryFile_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QTemporaryFile_override_virtual_Metacast(void* self, intptr_t slot);
void* QTemporaryFile_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QTemporaryFile_Delete(QTemporaryFile* self, bool isSubclass);

}
