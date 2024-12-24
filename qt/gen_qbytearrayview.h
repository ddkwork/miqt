#pragma once
#ifndef MIQT_QT_GEN_QBYTEARRAYVIEW_H
#define MIQT_QT_GEN_QBYTEARRAYVIEW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QByteArrayView QByteArrayView;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QByteArrayView* QByteArrayView_new();
extern __declspec(dllexport) 
QByteArrayView* QByteArrayView_new2(QByteArrayView* param1);
extern __declspec(dllexport) 
struct miqt_string QByteArrayView_ToByteArray(const QByteArrayView* self);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayView_Size(const QByteArrayView* self);
extern __declspec(dllexport) 
const_pointer QByteArrayView_Data(const QByteArrayView* self);
extern __declspec(dllexport) 
const_pointer QByteArrayView_ConstData(const QByteArrayView* self);
extern __declspec(dllexport) 
char QByteArrayView_OperatorSubscript(const QByteArrayView* self, ptrdiff_t n);
extern __declspec(dllexport) 
char QByteArrayView_At(const QByteArrayView* self, ptrdiff_t n);
extern __declspec(dllexport) 
QByteArrayView* QByteArrayView_First(const QByteArrayView* self, ptrdiff_t n);
extern __declspec(dllexport) 
QByteArrayView* QByteArrayView_Last(const QByteArrayView* self, ptrdiff_t n);
extern __declspec(dllexport) 
QByteArrayView* QByteArrayView_Sliced(const QByteArrayView* self, ptrdiff_t pos);
extern __declspec(dllexport) 
QByteArrayView* QByteArrayView_Sliced2(const QByteArrayView* self, ptrdiff_t pos, ptrdiff_t n);
extern __declspec(dllexport) 
QByteArrayView* QByteArrayView_Slice(QByteArrayView* self, ptrdiff_t pos);
extern __declspec(dllexport) 
QByteArrayView* QByteArrayView_Slice2(QByteArrayView* self, ptrdiff_t pos, ptrdiff_t n);
extern __declspec(dllexport) 
QByteArrayView* QByteArrayView_Chopped(const QByteArrayView* self, ptrdiff_t lenVal);
extern __declspec(dllexport) 
QByteArrayView* QByteArrayView_Left(const QByteArrayView* self, ptrdiff_t n);
extern __declspec(dllexport) 
QByteArrayView* QByteArrayView_Right(const QByteArrayView* self, ptrdiff_t n);
extern __declspec(dllexport) 
QByteArrayView* QByteArrayView_Mid(const QByteArrayView* self, ptrdiff_t pos);
extern __declspec(dllexport) 
void QByteArrayView_Truncate(QByteArrayView* self, ptrdiff_t n);
extern __declspec(dllexport) 
void QByteArrayView_Chop(QByteArrayView* self, ptrdiff_t n);
extern __declspec(dllexport) 
QByteArrayView* QByteArrayView_Trimmed(const QByteArrayView* self);
extern __declspec(dllexport) 
int16_t QByteArrayView_ToShort(const QByteArrayView* self);
extern __declspec(dllexport) 
uint16_t QByteArrayView_ToUShort(const QByteArrayView* self);
extern __declspec(dllexport) 
int QByteArrayView_ToInt(const QByteArrayView* self);
extern __declspec(dllexport) 
unsigned int QByteArrayView_ToUInt(const QByteArrayView* self);
extern __declspec(dllexport) 
long QByteArrayView_ToLong(const QByteArrayView* self);
extern __declspec(dllexport) 
unsigned long QByteArrayView_ToULong(const QByteArrayView* self);
extern __declspec(dllexport) 
long long QByteArrayView_ToLongLong(const QByteArrayView* self);
extern __declspec(dllexport) 
unsigned long long QByteArrayView_ToULongLong(const QByteArrayView* self);
extern __declspec(dllexport) 
float QByteArrayView_ToFloat(const QByteArrayView* self);
extern __declspec(dllexport) 
double QByteArrayView_ToDouble(const QByteArrayView* self);
extern __declspec(dllexport) 
bool QByteArrayView_StartsWith(const QByteArrayView* self, QByteArrayView* other);
extern __declspec(dllexport) 
bool QByteArrayView_StartsWithWithChar(const QByteArrayView* self, char c);
extern __declspec(dllexport) 
bool QByteArrayView_EndsWith(const QByteArrayView* self, QByteArrayView* other);
extern __declspec(dllexport) 
bool QByteArrayView_EndsWithWithChar(const QByteArrayView* self, char c);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayView_IndexOf(const QByteArrayView* self, QByteArrayView* a);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayView_IndexOfWithCh(const QByteArrayView* self, char ch);
extern __declspec(dllexport) 
bool QByteArrayView_Contains(const QByteArrayView* self, QByteArrayView* a);
extern __declspec(dllexport) 
bool QByteArrayView_ContainsWithChar(const QByteArrayView* self, char c);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayView_LastIndexOf(const QByteArrayView* self, QByteArrayView* a);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayView_LastIndexOf2(const QByteArrayView* self, QByteArrayView* a, ptrdiff_t from);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayView_LastIndexOfWithCh(const QByteArrayView* self, char ch);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayView_Count(const QByteArrayView* self, QByteArrayView* a);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayView_CountWithCh(const QByteArrayView* self, char ch);
extern __declspec(dllexport) 
int QByteArrayView_Compare(const QByteArrayView* self, QByteArrayView* a);
extern __declspec(dllexport) 
bool QByteArrayView_IsValidUtf8(const QByteArrayView* self);
extern __declspec(dllexport) 
const_iterator QByteArrayView_Begin(const QByteArrayView* self);
extern __declspec(dllexport) 
const_iterator QByteArrayView_End(const QByteArrayView* self);
extern __declspec(dllexport) 
const_iterator QByteArrayView_Cbegin(const QByteArrayView* self);
extern __declspec(dllexport) 
const_iterator QByteArrayView_Cend(const QByteArrayView* self);
extern __declspec(dllexport) 
const_reverse_iterator QByteArrayView_Rbegin(const QByteArrayView* self);
extern __declspec(dllexport) 
const_reverse_iterator QByteArrayView_Rend(const QByteArrayView* self);
extern __declspec(dllexport) 
const_reverse_iterator QByteArrayView_Crbegin(const QByteArrayView* self);
extern __declspec(dllexport) 
const_reverse_iterator QByteArrayView_Crend(const QByteArrayView* self);
extern __declspec(dllexport) 
bool QByteArrayView_Empty(const QByteArrayView* self);
extern __declspec(dllexport) 
char QByteArrayView_Front(const QByteArrayView* self);
extern __declspec(dllexport) 
char QByteArrayView_Back(const QByteArrayView* self);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayView_MaxSize(const QByteArrayView* self);
extern __declspec(dllexport) 
bool QByteArrayView_IsNull(const QByteArrayView* self);
extern __declspec(dllexport) 
bool QByteArrayView_IsEmpty(const QByteArrayView* self);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayView_Length(const QByteArrayView* self);
extern __declspec(dllexport) 
char QByteArrayView_First2(const QByteArrayView* self);
extern __declspec(dllexport) 
char QByteArrayView_Last2(const QByteArrayView* self);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayView_MaxSize2();
extern __declspec(dllexport) 
QByteArrayView* QByteArrayView_Mid2(const QByteArrayView* self, ptrdiff_t pos, ptrdiff_t n);
extern __declspec(dllexport) 
int16_t QByteArrayView_ToShort1(const QByteArrayView* self, bool* ok);
extern __declspec(dllexport) 
int16_t QByteArrayView_ToShort2(const QByteArrayView* self, bool* ok, int base);
extern __declspec(dllexport) 
uint16_t QByteArrayView_ToUShort1(const QByteArrayView* self, bool* ok);
extern __declspec(dllexport) 
uint16_t QByteArrayView_ToUShort2(const QByteArrayView* self, bool* ok, int base);
extern __declspec(dllexport) 
int QByteArrayView_ToInt1(const QByteArrayView* self, bool* ok);
extern __declspec(dllexport) 
int QByteArrayView_ToInt2(const QByteArrayView* self, bool* ok, int base);
extern __declspec(dllexport) 
unsigned int QByteArrayView_ToUInt1(const QByteArrayView* self, bool* ok);
extern __declspec(dllexport) 
unsigned int QByteArrayView_ToUInt2(const QByteArrayView* self, bool* ok, int base);
extern __declspec(dllexport) 
long QByteArrayView_ToLong1(const QByteArrayView* self, bool* ok);
extern __declspec(dllexport) 
long QByteArrayView_ToLong2(const QByteArrayView* self, bool* ok, int base);
extern __declspec(dllexport) 
unsigned long QByteArrayView_ToULong1(const QByteArrayView* self, bool* ok);
extern __declspec(dllexport) 
unsigned long QByteArrayView_ToULong2(const QByteArrayView* self, bool* ok, int base);
extern __declspec(dllexport) 
long long QByteArrayView_ToLongLong1(const QByteArrayView* self, bool* ok);
extern __declspec(dllexport) 
long long QByteArrayView_ToLongLong2(const QByteArrayView* self, bool* ok, int base);
extern __declspec(dllexport) 
unsigned long long QByteArrayView_ToULongLong1(const QByteArrayView* self, bool* ok);
extern __declspec(dllexport) 
unsigned long long QByteArrayView_ToULongLong2(const QByteArrayView* self, bool* ok, int base);
extern __declspec(dllexport) 
float QByteArrayView_ToFloat1(const QByteArrayView* self, bool* ok);
extern __declspec(dllexport) 
double QByteArrayView_ToDouble1(const QByteArrayView* self, bool* ok);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayView_IndexOf2(const QByteArrayView* self, QByteArrayView* a, ptrdiff_t from);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayView_IndexOf22(const QByteArrayView* self, char ch, ptrdiff_t from);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayView_LastIndexOf22(const QByteArrayView* self, char ch, ptrdiff_t from);
extern __declspec(dllexport) 
int QByteArrayView_Compare2(const QByteArrayView* self, QByteArrayView* a, int cs);
extern __declspec(dllexport) 
void QByteArrayView_Delete(QByteArrayView* self, bool isSubclass);

}
