#pragma once
#ifndef MIQT_QT_GEN_QTRANSLATOR_H
#define MIQT_QT_GEN_QTRANSLATOR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QLocale QLocale;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTranslator QTranslator;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QTranslator* QTranslator_new();
extern __declspec(dllexport) 
QTranslator* QTranslator_new2(QObject* parent);
extern __declspec(dllexport) 
void QTranslator_virtbase(QTranslator* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QTranslator_MetaObject(const QTranslator* self);
extern __declspec(dllexport) 
void* QTranslator_Metacast(QTranslator* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QTranslator_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_string QTranslator_Translate(const QTranslator* self, const char* context, const char* sourceText, const char* disambiguation, int n);
extern __declspec(dllexport) 
bool QTranslator_IsEmpty(const QTranslator* self);
extern __declspec(dllexport) 
struct miqt_string QTranslator_Language(const QTranslator* self);
extern __declspec(dllexport) 
struct miqt_string QTranslator_FilePath(const QTranslator* self);
extern __declspec(dllexport) 
bool QTranslator_Load(QTranslator* self, struct miqt_string filename);
extern __declspec(dllexport) 
bool QTranslator_Load2(QTranslator* self, QLocale* locale, struct miqt_string filename);
extern __declspec(dllexport) 
bool QTranslator_Load3(QTranslator* self, const unsigned char* data, int lenVal);
extern __declspec(dllexport) 
struct miqt_string QTranslator_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QTranslator_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
bool QTranslator_Load22(QTranslator* self, struct miqt_string filename, struct miqt_string directory);
extern __declspec(dllexport) 
bool QTranslator_Load32(QTranslator* self, struct miqt_string filename, struct miqt_string directory, struct miqt_string search_delimiters);
extern __declspec(dllexport) 
bool QTranslator_Load4(QTranslator* self, struct miqt_string filename, struct miqt_string directory, struct miqt_string search_delimiters, struct miqt_string suffix);
extern __declspec(dllexport) 
bool QTranslator_Load33(QTranslator* self, QLocale* locale, struct miqt_string filename, struct miqt_string prefix);
extern __declspec(dllexport) 
bool QTranslator_Load42(QTranslator* self, QLocale* locale, struct miqt_string filename, struct miqt_string prefix, struct miqt_string directory);
extern __declspec(dllexport) 
bool QTranslator_Load5(QTranslator* self, QLocale* locale, struct miqt_string filename, struct miqt_string prefix, struct miqt_string directory, struct miqt_string suffix);
extern __declspec(dllexport) 
bool QTranslator_Load34(QTranslator* self, const unsigned char* data, int lenVal, struct miqt_string directory);
extern __declspec(dllexport) 
void QTranslator_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QTranslator_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QTranslator_override_virtual_Metacast(void* self, intptr_t slot);
void* QTranslator_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QTranslator_Delete(QTranslator* self, bool isSubclass);

}
