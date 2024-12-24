#pragma once
#ifndef MIQT_QT_GEN_QFILESELECTOR_H
#define MIQT_QT_GEN_QFILESELECTOR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QFileSelector QFileSelector;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QUrl QUrl;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QFileSelector* QFileSelector_new();
extern __declspec(dllexport) 
QFileSelector* QFileSelector_new2(QObject* parent);
extern __declspec(dllexport) 
void QFileSelector_virtbase(QFileSelector* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QFileSelector_MetaObject(const QFileSelector* self);
extern __declspec(dllexport) 
void* QFileSelector_Metacast(QFileSelector* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QFileSelector_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_string QFileSelector_Select(const QFileSelector* self, struct miqt_string filePath);
extern __declspec(dllexport) 
QUrl* QFileSelector_SelectWithFilePath(const QFileSelector* self, QUrl* filePath);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QFileSelector_ExtraSelectors(const QFileSelector* self);
extern __declspec(dllexport) 
void QFileSelector_SetExtraSelectors(QFileSelector* self, struct miqt_array /* of struct miqt_string */  list);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QFileSelector_AllSelectors(const QFileSelector* self);
extern __declspec(dllexport) 
struct miqt_string QFileSelector_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QFileSelector_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QFileSelector_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QFileSelector_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QFileSelector_override_virtual_Metacast(void* self, intptr_t slot);
void* QFileSelector_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QFileSelector_Delete(QFileSelector* self, bool isSubclass);

}
