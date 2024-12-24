#pragma once
#ifndef MIQT_QT_GEN_QCHAR_H
#define MIQT_QT_GEN_QCHAR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChar QChar;
typedef struct QLatin1Char QLatin1Char;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QLatin1Char* QLatin1Char_new(char c);
extern __declspec(dllexport) 
QLatin1Char* QLatin1Char_new2(QLatin1Char* param1);
extern __declspec(dllexport) 
char QLatin1Char_ToLatin1(const QLatin1Char* self);
extern __declspec(dllexport) 
void QLatin1Char_Delete(QLatin1Char* self, bool isSubclass);

extern __declspec(dllexport) 
QChar* QChar_new();
extern __declspec(dllexport) 
QChar* QChar_new2(unsigned char c, unsigned char r);
extern __declspec(dllexport) 
QChar* QChar_new3(QChar* param1);
extern __declspec(dllexport) 
Category QChar_Category(const QChar* self);
extern __declspec(dllexport) 
Direction QChar_Direction(const QChar* self);
extern __declspec(dllexport) 
JoiningType QChar_JoiningType(const QChar* self);
extern __declspec(dllexport) 
unsigned char QChar_CombiningClass(const QChar* self);
extern __declspec(dllexport) 
QChar* QChar_MirroredChar(const QChar* self);
extern __declspec(dllexport) 
bool QChar_HasMirrored(const QChar* self);
extern __declspec(dllexport) 
struct miqt_string QChar_Decomposition(const QChar* self);
extern __declspec(dllexport) 
Decomposition QChar_DecompositionTag(const QChar* self);
extern __declspec(dllexport) 
int QChar_DigitValue(const QChar* self);
extern __declspec(dllexport) 
QChar* QChar_ToLower(const QChar* self);
extern __declspec(dllexport) 
QChar* QChar_ToUpper(const QChar* self);
extern __declspec(dllexport) 
QChar* QChar_ToTitleCase(const QChar* self);
extern __declspec(dllexport) 
QChar* QChar_ToCaseFolded(const QChar* self);
extern __declspec(dllexport) 
Script QChar_Script(const QChar* self);
extern __declspec(dllexport) 
UnicodeVersion QChar_UnicodeVersion(const QChar* self);
extern __declspec(dllexport) 
char QChar_ToLatin1(const QChar* self);
extern __declspec(dllexport) 
QChar* QChar_FromLatin1(char c);
extern __declspec(dllexport) 
bool QChar_IsNull(const QChar* self);
extern __declspec(dllexport) 
bool QChar_IsPrint(const QChar* self);
extern __declspec(dllexport) 
bool QChar_IsSpace(const QChar* self);
extern __declspec(dllexport) 
bool QChar_IsMark(const QChar* self);
extern __declspec(dllexport) 
bool QChar_IsPunct(const QChar* self);
extern __declspec(dllexport) 
bool QChar_IsSymbol(const QChar* self);
extern __declspec(dllexport) 
bool QChar_IsLetter(const QChar* self);
extern __declspec(dllexport) 
bool QChar_IsNumber(const QChar* self);
extern __declspec(dllexport) 
bool QChar_IsLetterOrNumber(const QChar* self);
extern __declspec(dllexport) 
bool QChar_IsDigit(const QChar* self);
extern __declspec(dllexport) 
bool QChar_IsLower(const QChar* self);
extern __declspec(dllexport) 
bool QChar_IsUpper(const QChar* self);
extern __declspec(dllexport) 
bool QChar_IsTitleCase(const QChar* self);
extern __declspec(dllexport) 
bool QChar_IsNonCharacter(const QChar* self);
extern __declspec(dllexport) 
bool QChar_IsHighSurrogate(const QChar* self);
extern __declspec(dllexport) 
bool QChar_IsLowSurrogate(const QChar* self);
extern __declspec(dllexport) 
bool QChar_IsSurrogate(const QChar* self);
extern __declspec(dllexport) 
unsigned char QChar_Cell(const QChar* self);
extern __declspec(dllexport) 
unsigned char QChar_Row(const QChar* self);
extern __declspec(dllexport) 
void QChar_SetCell(QChar* self, unsigned char acell);
extern __declspec(dllexport) 
void QChar_SetRow(QChar* self, unsigned char arow);
extern __declspec(dllexport) 
UnicodeVersion QChar_CurrentUnicodeVersion();
extern __declspec(dllexport) 
void QChar_Delete(QChar* self, bool isSubclass);

}
