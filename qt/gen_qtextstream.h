#pragma once
#ifndef MIQT_QT_GEN_QTEXTSTREAM_H
#define MIQT_QT_GEN_QTEXTSTREAM_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChar QChar;
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;
typedef struct QLocale QLocale;
typedef struct QTextStream QTextStream;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTextStream* QTextStream_new();
extern __declspec(dllexport) QTextStream* QTextStream_new2(QIODevice* device);
extern __declspec(dllexport) QTextStream* QTextStream_new3(struct miqt_string array);
extern __declspec(dllexport) QTextStream* QTextStream_new4(struct miqt_string array, OpenMode openMode);
extern __declspec(dllexport) void QTextStream_virtbase(QTextStream* src, QIODeviceBase** outptr_QIODeviceBase);
extern __declspec(dllexport) void QTextStream_SetEncoding(QTextStream* self, int encoding);
extern __declspec(dllexport) int QTextStream_Encoding(const QTextStream* self);
extern __declspec(dllexport) void QTextStream_SetAutoDetectUnicode(QTextStream* self, bool enabled);
extern __declspec(dllexport) bool QTextStream_AutoDetectUnicode(const QTextStream* self);
extern __declspec(dllexport) void QTextStream_SetGenerateByteOrderMark(QTextStream* self, bool generate);
extern __declspec(dllexport) bool QTextStream_GenerateByteOrderMark(const QTextStream* self);
extern __declspec(dllexport) void QTextStream_SetLocale(QTextStream* self, QLocale* locale);
extern __declspec(dllexport) QLocale* QTextStream_Locale(const QTextStream* self);
extern __declspec(dllexport) void QTextStream_SetDevice(QTextStream* self, QIODevice* device);
extern __declspec(dllexport) QIODevice* QTextStream_Device(const QTextStream* self);
extern __declspec(dllexport) struct miqt_string QTextStream_String(const QTextStream* self);
extern __declspec(dllexport) Status QTextStream_Status(const QTextStream* self);
extern __declspec(dllexport) void QTextStream_SetStatus(QTextStream* self, Status status);
extern __declspec(dllexport) void QTextStream_ResetStatus(QTextStream* self);
extern __declspec(dllexport) bool QTextStream_AtEnd(const QTextStream* self);
extern __declspec(dllexport) void QTextStream_Reset(QTextStream* self);
extern __declspec(dllexport) void QTextStream_Flush(QTextStream* self);
extern __declspec(dllexport) bool QTextStream_Seek(QTextStream* self, long long pos);
extern __declspec(dllexport) long long QTextStream_Pos(const QTextStream* self);
extern __declspec(dllexport) void QTextStream_SkipWhiteSpace(QTextStream* self);
extern __declspec(dllexport) struct miqt_string QTextStream_ReadLine(QTextStream* self);
extern __declspec(dllexport) struct miqt_string QTextStream_ReadAll(QTextStream* self);
extern __declspec(dllexport) struct miqt_string QTextStream_Read(QTextStream* self, long long maxlen);
extern __declspec(dllexport) void QTextStream_SetFieldAlignment(QTextStream* self, FieldAlignment alignment);
extern __declspec(dllexport) FieldAlignment QTextStream_FieldAlignment(const QTextStream* self);
extern __declspec(dllexport) void QTextStream_SetPadChar(QTextStream* self, QChar* ch);
extern __declspec(dllexport) QChar* QTextStream_PadChar(const QTextStream* self);
extern __declspec(dllexport) void QTextStream_SetFieldWidth(QTextStream* self, int width);
extern __declspec(dllexport) int QTextStream_FieldWidth(const QTextStream* self);
extern __declspec(dllexport) void QTextStream_SetNumberFlags(QTextStream* self, NumberFlags flags);
extern __declspec(dllexport) NumberFlags QTextStream_NumberFlags(const QTextStream* self);
extern __declspec(dllexport) void QTextStream_SetIntegerBase(QTextStream* self, int base);
extern __declspec(dllexport) int QTextStream_IntegerBase(const QTextStream* self);
extern __declspec(dllexport) void QTextStream_SetRealNumberNotation(QTextStream* self, RealNumberNotation notation);
extern __declspec(dllexport) RealNumberNotation QTextStream_RealNumberNotation(const QTextStream* self);
extern __declspec(dllexport) void QTextStream_SetRealNumberPrecision(QTextStream* self, int precision);
extern __declspec(dllexport) int QTextStream_RealNumberPrecision(const QTextStream* self);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftRight(QTextStream* self, QChar* ch);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftRightWithCh(QTextStream* self, char* ch);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftRightWithShort(QTextStream* self, int16_t* i);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftRightWithUnsignedshort(QTextStream* self, uint16_t* i);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftRightWithInt(QTextStream* self, int* i);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftRightWithUnsignedint(QTextStream* self, unsigned int* i);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftRightWithLong(QTextStream* self, long* i);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftRightWithUnsignedlong(QTextStream* self, unsigned long* i);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftRightWithQlonglong(QTextStream* self, long long* i);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftRightWithQulonglong(QTextStream* self, unsigned long long* i);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftRightWithFloat(QTextStream* self, float* f);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftRightWithDouble(QTextStream* self, double* f);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftRightWithQString(QTextStream* self, struct miqt_string s);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftRightWithArray(QTextStream* self, struct miqt_string array);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftRightWithChar(QTextStream* self, char* c);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftLeft(QTextStream* self, QChar* ch);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftLeftWithCh(QTextStream* self, char ch);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftLeftWithShort(QTextStream* self, int16_t i);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftLeftWithUnsignedshort(QTextStream* self, uint16_t i);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftLeftWithInt(QTextStream* self, int i);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftLeftWithUnsignedint(QTextStream* self, unsigned int i);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftLeftWithLong(QTextStream* self, long i);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftLeftWithUnsignedlong(QTextStream* self, unsigned long i);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftLeftWithQlonglong(QTextStream* self, long long i);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftLeftWithQulonglong(QTextStream* self, unsigned long long i);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftLeftWithFloat(QTextStream* self, float f);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftLeftWithDouble(QTextStream* self, double f);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftLeftWithQString(QTextStream* self, struct miqt_string s);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftLeftWithArray(QTextStream* self, struct miqt_string array);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftLeftWithChar(QTextStream* self, const char* c);
extern __declspec(dllexport) QTextStream* QTextStream_OperatorShiftLeftWithPtr(QTextStream* self, const void* ptr);
extern __declspec(dllexport) struct miqt_string QTextStream_ReadLine1(QTextStream* self, long long maxlen);
extern __declspec(dllexport) void QTextStream_Delete(QTextStream* self, bool isSubclass);

} 
