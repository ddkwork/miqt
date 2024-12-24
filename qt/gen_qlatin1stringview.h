#pragma once
#ifndef MIQT_QT_GEN_QLATIN1STRINGVIEW_H
#define MIQT_QT_GEN_QLATIN1STRINGVIEW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QByteArrayView QByteArrayView;
typedef struct QChar QChar;
typedef struct QLatin1Char QLatin1Char;
typedef struct QLatin1String QLatin1String;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QLatin1String* QLatin1String_new();
extern __declspec(dllexport) QLatin1String* QLatin1String_new2(const char* s);
extern __declspec(dllexport) QLatin1String* QLatin1String_new3(const char* f, const char* l);
extern __declspec(dllexport) QLatin1String* QLatin1String_new4(const char* s, ptrdiff_t sz);
extern __declspec(dllexport) QLatin1String* QLatin1String_new5(struct miqt_string s);
extern __declspec(dllexport) QLatin1String* QLatin1String_new6(QByteArrayView* s);
extern __declspec(dllexport) struct miqt_string QLatin1String_ToString(const QLatin1String* self);
extern __declspec(dllexport) struct miqt_string QLatin1String_ToUtf8(const QLatin1String* self);
extern __declspec(dllexport) const char* QLatin1String_Latin1(const QLatin1String* self);
extern __declspec(dllexport) ptrdiff_t QLatin1String_Size(const QLatin1String* self);
extern __declspec(dllexport) const char* QLatin1String_Data(const QLatin1String* self);
extern __declspec(dllexport) const char* QLatin1String_ConstData(const QLatin1String* self);
extern __declspec(dllexport) const char* QLatin1String_ConstBegin(const QLatin1String* self);
extern __declspec(dllexport) const char* QLatin1String_ConstEnd(const QLatin1String* self);
extern __declspec(dllexport) QLatin1Char* QLatin1String_First(const QLatin1String* self);
extern __declspec(dllexport) QLatin1Char* QLatin1String_Last(const QLatin1String* self);
extern __declspec(dllexport) ptrdiff_t QLatin1String_Length(const QLatin1String* self);
extern __declspec(dllexport) bool QLatin1String_IsNull(const QLatin1String* self);
extern __declspec(dllexport) bool QLatin1String_IsEmpty(const QLatin1String* self);
extern __declspec(dllexport) bool QLatin1String_Empty(const QLatin1String* self);
extern __declspec(dllexport) QLatin1Char* QLatin1String_At(const QLatin1String* self, ptrdiff_t i);
extern __declspec(dllexport) QLatin1Char* QLatin1String_OperatorSubscript(const QLatin1String* self, ptrdiff_t i);
extern __declspec(dllexport) QLatin1Char* QLatin1String_Front(const QLatin1String* self);
extern __declspec(dllexport) QLatin1Char* QLatin1String_Back(const QLatin1String* self);
extern __declspec(dllexport) int QLatin1String_CompareWithQChar(const QLatin1String* self, QChar* c);
extern __declspec(dllexport) int QLatin1String_Compare3(const QLatin1String* self, QChar* c, int cs);
extern __declspec(dllexport) bool QLatin1String_StartsWithWithQChar(const QLatin1String* self, QChar* c);
extern __declspec(dllexport) bool QLatin1String_StartsWith2(const QLatin1String* self, QChar* c, int cs);
extern __declspec(dllexport) bool QLatin1String_EndsWithWithQChar(const QLatin1String* self, QChar* c);
extern __declspec(dllexport) bool QLatin1String_EndsWith2(const QLatin1String* self, QChar* c, int cs);
extern __declspec(dllexport) ptrdiff_t QLatin1String_IndexOfWithQChar(const QLatin1String* self, QChar* c);
extern __declspec(dllexport) bool QLatin1String_ContainsWithQChar(const QLatin1String* self, QChar* c);
extern __declspec(dllexport) ptrdiff_t QLatin1String_LastIndexOfWithQChar(const QLatin1String* self, QChar* c);
extern __declspec(dllexport) ptrdiff_t QLatin1String_LastIndexOf4(const QLatin1String* self, QChar* c, ptrdiff_t from);
extern __declspec(dllexport) ptrdiff_t QLatin1String_CountWithCh(const QLatin1String* self, QChar* ch);
extern __declspec(dllexport) int16_t QLatin1String_ToShort(const QLatin1String* self);
extern __declspec(dllexport) uint16_t QLatin1String_ToUShort(const QLatin1String* self);
extern __declspec(dllexport) int QLatin1String_ToInt(const QLatin1String* self);
extern __declspec(dllexport) unsigned int QLatin1String_ToUInt(const QLatin1String* self);
extern __declspec(dllexport) long QLatin1String_ToLong(const QLatin1String* self);
extern __declspec(dllexport) unsigned long QLatin1String_ToULong(const QLatin1String* self);
extern __declspec(dllexport) long long QLatin1String_ToLongLong(const QLatin1String* self);
extern __declspec(dllexport) unsigned long long QLatin1String_ToULongLong(const QLatin1String* self);
extern __declspec(dllexport) float QLatin1String_ToFloat(const QLatin1String* self);
extern __declspec(dllexport) double QLatin1String_ToDouble(const QLatin1String* self);
extern __declspec(dllexport) const_iterator QLatin1String_Begin(const QLatin1String* self);
extern __declspec(dllexport) const_iterator QLatin1String_Cbegin(const QLatin1String* self);
extern __declspec(dllexport) const_iterator QLatin1String_End(const QLatin1String* self);
extern __declspec(dllexport) const_iterator QLatin1String_Cend(const QLatin1String* self);
extern __declspec(dllexport) const_reverse_iterator QLatin1String_Rbegin(const QLatin1String* self);
extern __declspec(dllexport) const_reverse_iterator QLatin1String_Crbegin(const QLatin1String* self);
extern __declspec(dllexport) const_reverse_iterator QLatin1String_Rend(const QLatin1String* self);
extern __declspec(dllexport) const_reverse_iterator QLatin1String_Crend(const QLatin1String* self);
extern __declspec(dllexport) ptrdiff_t QLatin1String_MaxSize(const QLatin1String* self);
extern __declspec(dllexport) ptrdiff_t QLatin1String_MaxSize2();
extern __declspec(dllexport) void QLatin1String_Chop(QLatin1String* self, ptrdiff_t n);
extern __declspec(dllexport) void QLatin1String_Truncate(QLatin1String* self, ptrdiff_t n);
extern __declspec(dllexport) ptrdiff_t QLatin1String_IndexOf23(const QLatin1String* self, QChar* c, ptrdiff_t from);
extern __declspec(dllexport) ptrdiff_t QLatin1String_IndexOf33(const QLatin1String* self, QChar* c, ptrdiff_t from, int cs);
extern __declspec(dllexport) bool QLatin1String_Contains23(const QLatin1String* self, QChar* c, int cs);
extern __declspec(dllexport) ptrdiff_t QLatin1String_LastIndexOf24(const QLatin1String* self, QChar* c, int cs);
extern __declspec(dllexport) ptrdiff_t QLatin1String_LastIndexOf34(const QLatin1String* self, QChar* c, ptrdiff_t from, int cs);
extern __declspec(dllexport) ptrdiff_t QLatin1String_Count23(const QLatin1String* self, QChar* ch, int cs);
extern __declspec(dllexport) int16_t QLatin1String_ToShort1(const QLatin1String* self, bool* ok);
extern __declspec(dllexport) int16_t QLatin1String_ToShort2(const QLatin1String* self, bool* ok, int base);
extern __declspec(dllexport) uint16_t QLatin1String_ToUShort1(const QLatin1String* self, bool* ok);
extern __declspec(dllexport) uint16_t QLatin1String_ToUShort2(const QLatin1String* self, bool* ok, int base);
extern __declspec(dllexport) int QLatin1String_ToInt1(const QLatin1String* self, bool* ok);
extern __declspec(dllexport) int QLatin1String_ToInt2(const QLatin1String* self, bool* ok, int base);
extern __declspec(dllexport) unsigned int QLatin1String_ToUInt1(const QLatin1String* self, bool* ok);
extern __declspec(dllexport) unsigned int QLatin1String_ToUInt2(const QLatin1String* self, bool* ok, int base);
extern __declspec(dllexport) long QLatin1String_ToLong1(const QLatin1String* self, bool* ok);
extern __declspec(dllexport) long QLatin1String_ToLong2(const QLatin1String* self, bool* ok, int base);
extern __declspec(dllexport) unsigned long QLatin1String_ToULong1(const QLatin1String* self, bool* ok);
extern __declspec(dllexport) unsigned long QLatin1String_ToULong2(const QLatin1String* self, bool* ok, int base);
extern __declspec(dllexport) long long QLatin1String_ToLongLong1(const QLatin1String* self, bool* ok);
extern __declspec(dllexport) long long QLatin1String_ToLongLong2(const QLatin1String* self, bool* ok, int base);
extern __declspec(dllexport) unsigned long long QLatin1String_ToULongLong1(const QLatin1String* self, bool* ok);
extern __declspec(dllexport) unsigned long long QLatin1String_ToULongLong2(const QLatin1String* self, bool* ok, int base);
extern __declspec(dllexport) float QLatin1String_ToFloat1(const QLatin1String* self, bool* ok);
extern __declspec(dllexport) double QLatin1String_ToDouble1(const QLatin1String* self, bool* ok);
extern __declspec(dllexport) void QLatin1String_Delete(QLatin1String* self, bool isSubclass);

} 
