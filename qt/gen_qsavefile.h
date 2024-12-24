#pragma once
#ifndef MIQT_QT_GEN_QSAVEFILE_H
#define MIQT_QT_GEN_QSAVEFILE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QFileDevice QFileDevice;
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSaveFile QSaveFile;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QSaveFile* QSaveFile_new(struct miqt_string name);
extern __declspec(dllexport) 
QSaveFile* QSaveFile_new2();
extern __declspec(dllexport) 
QSaveFile* QSaveFile_new3(struct miqt_string name, QObject* parent);
extern __declspec(dllexport) 
QSaveFile* QSaveFile_new4(QObject* parent);
extern __declspec(dllexport) 
void QSaveFile_virtbase(QSaveFile* src
, QFileDevice** outptr_QFileDevice
);
extern __declspec(dllexport) 
QMetaObject* QSaveFile_MetaObject(const QSaveFile* self);
extern __declspec(dllexport) 
void* QSaveFile_Metacast(QSaveFile* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QSaveFile_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_string QSaveFile_FileName(const QSaveFile* self);
extern __declspec(dllexport) 
void QSaveFile_SetFileName(QSaveFile* self, struct miqt_string name);
extern __declspec(dllexport) 
bool QSaveFile_Open(QSaveFile* self, OpenMode flags);
extern __declspec(dllexport) 
bool QSaveFile_Commit(QSaveFile* self);
extern __declspec(dllexport) 
void QSaveFile_CancelWriting(QSaveFile* self);
extern __declspec(dllexport) 
void QSaveFile_SetDirectWriteFallback(QSaveFile* self, bool enabled);
extern __declspec(dllexport) 
bool QSaveFile_DirectWriteFallback(const QSaveFile* self);
extern __declspec(dllexport) 
long long QSaveFile_WriteData(QSaveFile* self, const char* data, long long lenVal);
extern __declspec(dllexport) 
struct miqt_string QSaveFile_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QSaveFile_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QSaveFile_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QSaveFile_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QSaveFile_override_virtual_Metacast(void* self, intptr_t slot);
void* QSaveFile_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QSaveFile_Delete(QSaveFile* self, bool isSubclass);

}
