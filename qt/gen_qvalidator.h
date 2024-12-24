#pragma once
#ifndef MIQT_QT_GEN_QVALIDATOR_H
#define MIQT_QT_GEN_QVALIDATOR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDoubleValidator QDoubleValidator;
typedef struct QIntValidator QIntValidator;
typedef struct QLocale QLocale;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QRegularExpression QRegularExpression;
typedef struct QRegularExpressionValidator QRegularExpressionValidator;
typedef struct QValidator QValidator;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QValidator* QValidator_new();
extern __declspec(dllexport) 
QValidator* QValidator_new2(QObject* parent);
extern __declspec(dllexport) 
void QValidator_virtbase(QValidator* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QValidator_MetaObject(const QValidator* self);
extern __declspec(dllexport) 
void* QValidator_Metacast(QValidator* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QValidator_Tr(const char* s);
extern __declspec(dllexport) 
void QValidator_SetLocale(QValidator* self, QLocale* locale);
extern __declspec(dllexport) 
QLocale* QValidator_Locale(const QValidator* self);
extern __declspec(dllexport) 
State QValidator_Validate(const QValidator* self, struct miqt_string param1, int* param2);
extern __declspec(dllexport) 
void QValidator_Fixup(const QValidator* self, struct miqt_string param1);
extern __declspec(dllexport) 
void QValidator_Changed(QValidator* self);
void QValidator_connect_Changed(QValidator* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QValidator_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QValidator_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QValidator_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QValidator_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QValidator_override_virtual_Metacast(void* self, intptr_t slot);
void* QValidator_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QValidator_Delete(QValidator* self, bool isSubclass);

extern __declspec(dllexport) 
QIntValidator* QIntValidator_new();
extern __declspec(dllexport) 
QIntValidator* QIntValidator_new2(int bottom, int top);
extern __declspec(dllexport) 
QIntValidator* QIntValidator_new3(QObject* parent);
extern __declspec(dllexport) 
QIntValidator* QIntValidator_new4(int bottom, int top, QObject* parent);
extern __declspec(dllexport) 
void QIntValidator_virtbase(QIntValidator* src
, QValidator** outptr_QValidator
);
extern __declspec(dllexport) 
QMetaObject* QIntValidator_MetaObject(const QIntValidator* self);
extern __declspec(dllexport) 
void* QIntValidator_Metacast(QIntValidator* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QIntValidator_Tr(const char* s);
extern __declspec(dllexport) 
int QIntValidator_Validate(const QIntValidator* self, struct miqt_string param1, int* param2);
extern __declspec(dllexport) 
void QIntValidator_Fixup(const QIntValidator* self, struct miqt_string input);
extern __declspec(dllexport) 
void QIntValidator_SetBottom(QIntValidator* self, int bottom);
extern __declspec(dllexport) 
void QIntValidator_SetTop(QIntValidator* self, int top);
extern __declspec(dllexport) 
void QIntValidator_SetRange(QIntValidator* self, int bottom, int top);
extern __declspec(dllexport) 
int QIntValidator_Bottom(const QIntValidator* self);
extern __declspec(dllexport) 
int QIntValidator_Top(const QIntValidator* self);
extern __declspec(dllexport) 
void QIntValidator_BottomChanged(QIntValidator* self, int bottom);
void QIntValidator_connect_BottomChanged(QIntValidator* self, intptr_t slot);
extern __declspec(dllexport) 
void QIntValidator_TopChanged(QIntValidator* self, int top);
void QIntValidator_connect_TopChanged(QIntValidator* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QIntValidator_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QIntValidator_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QIntValidator_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QIntValidator_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QIntValidator_override_virtual_Metacast(void* self, intptr_t slot);
void* QIntValidator_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QIntValidator_Delete(QIntValidator* self, bool isSubclass);

extern __declspec(dllexport) 
QDoubleValidator* QDoubleValidator_new();
extern __declspec(dllexport) 
QDoubleValidator* QDoubleValidator_new2(double bottom, double top, int decimals);
extern __declspec(dllexport) 
QDoubleValidator* QDoubleValidator_new3(QObject* parent);
extern __declspec(dllexport) 
QDoubleValidator* QDoubleValidator_new4(double bottom, double top, int decimals, QObject* parent);
extern __declspec(dllexport) 
void QDoubleValidator_virtbase(QDoubleValidator* src
, QValidator** outptr_QValidator
);
extern __declspec(dllexport) 
QMetaObject* QDoubleValidator_MetaObject(const QDoubleValidator* self);
extern __declspec(dllexport) 
void* QDoubleValidator_Metacast(QDoubleValidator* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QDoubleValidator_Tr(const char* s);
extern __declspec(dllexport) 
int QDoubleValidator_Validate(const QDoubleValidator* self, struct miqt_string param1, int* param2);
extern __declspec(dllexport) 
void QDoubleValidator_Fixup(const QDoubleValidator* self, struct miqt_string input);
extern __declspec(dllexport) 
void QDoubleValidator_SetRange(QDoubleValidator* self, double bottom, double top, int decimals);
extern __declspec(dllexport) 
void QDoubleValidator_SetRange2(QDoubleValidator* self, double bottom, double top);
extern __declspec(dllexport) 
void QDoubleValidator_SetBottom(QDoubleValidator* self, double bottom);
extern __declspec(dllexport) 
void QDoubleValidator_SetTop(QDoubleValidator* self, double top);
extern __declspec(dllexport) 
void QDoubleValidator_SetDecimals(QDoubleValidator* self, int decimals);
extern __declspec(dllexport) 
void QDoubleValidator_SetNotation(QDoubleValidator* self, Notation notation);
extern __declspec(dllexport) 
double QDoubleValidator_Bottom(const QDoubleValidator* self);
extern __declspec(dllexport) 
double QDoubleValidator_Top(const QDoubleValidator* self);
extern __declspec(dllexport) 
int QDoubleValidator_Decimals(const QDoubleValidator* self);
extern __declspec(dllexport) 
Notation QDoubleValidator_Notation(const QDoubleValidator* self);
extern __declspec(dllexport) 
void QDoubleValidator_BottomChanged(QDoubleValidator* self, double bottom);
void QDoubleValidator_connect_BottomChanged(QDoubleValidator* self, intptr_t slot);
extern __declspec(dllexport) 
void QDoubleValidator_TopChanged(QDoubleValidator* self, double top);
void QDoubleValidator_connect_TopChanged(QDoubleValidator* self, intptr_t slot);
extern __declspec(dllexport) 
void QDoubleValidator_DecimalsChanged(QDoubleValidator* self, int decimals);
void QDoubleValidator_connect_DecimalsChanged(QDoubleValidator* self, intptr_t slot);
extern __declspec(dllexport) 
void QDoubleValidator_NotationChanged(QDoubleValidator* self, int notation);
void QDoubleValidator_connect_NotationChanged(QDoubleValidator* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QDoubleValidator_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QDoubleValidator_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QDoubleValidator_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QDoubleValidator_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QDoubleValidator_override_virtual_Metacast(void* self, intptr_t slot);
void* QDoubleValidator_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QDoubleValidator_Delete(QDoubleValidator* self, bool isSubclass);

extern __declspec(dllexport) 
QRegularExpressionValidator* QRegularExpressionValidator_new();
extern __declspec(dllexport) 
QRegularExpressionValidator* QRegularExpressionValidator_new2(QRegularExpression* re);
extern __declspec(dllexport) 
QRegularExpressionValidator* QRegularExpressionValidator_new3(QObject* parent);
extern __declspec(dllexport) 
QRegularExpressionValidator* QRegularExpressionValidator_new4(QRegularExpression* re, QObject* parent);
extern __declspec(dllexport) 
void QRegularExpressionValidator_virtbase(QRegularExpressionValidator* src
, QValidator** outptr_QValidator
);
extern __declspec(dllexport) 
QMetaObject* QRegularExpressionValidator_MetaObject(const QRegularExpressionValidator* self);
extern __declspec(dllexport) 
void* QRegularExpressionValidator_Metacast(QRegularExpressionValidator* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QRegularExpressionValidator_Tr(const char* s);
extern __declspec(dllexport) 
int QRegularExpressionValidator_Validate(const QRegularExpressionValidator* self, struct miqt_string input, int* pos);
extern __declspec(dllexport) 
QRegularExpression* QRegularExpressionValidator_RegularExpression(const QRegularExpressionValidator* self);
extern __declspec(dllexport) 
void QRegularExpressionValidator_SetRegularExpression(QRegularExpressionValidator* self, QRegularExpression* re);
extern __declspec(dllexport) 
void QRegularExpressionValidator_RegularExpressionChanged(QRegularExpressionValidator* self, QRegularExpression* re);
void QRegularExpressionValidator_connect_RegularExpressionChanged(QRegularExpressionValidator* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QRegularExpressionValidator_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QRegularExpressionValidator_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QRegularExpressionValidator_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QRegularExpressionValidator_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QRegularExpressionValidator_override_virtual_Metacast(void* self, intptr_t slot);
void* QRegularExpressionValidator_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QRegularExpressionValidator_Delete(QRegularExpressionValidator* self, bool isSubclass);

}
