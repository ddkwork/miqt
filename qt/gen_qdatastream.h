#pragma once
#ifndef MIQT_QT_GEN_QDATASTREAM_H
#define MIQT_QT_GEN_QDATASTREAM_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDataStream QDataStream;
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QDataStream* QDataStream_new();
extern __declspec(dllexport) QDataStream* QDataStream_new2(QIODevice* param1);
extern __declspec(dllexport) QDataStream* QDataStream_new3(struct miqt_string param1);
extern __declspec(dllexport) void QDataStream_virtbase(QDataStream* src, QIODeviceBase** outptr_QIODeviceBase);
extern __declspec(dllexport) QIODevice* QDataStream_Device(const QDataStream* self);
extern __declspec(dllexport) void QDataStream_SetDevice(QDataStream* self, QIODevice* device);
extern __declspec(dllexport) bool QDataStream_AtEnd(const QDataStream* self);
extern __declspec(dllexport) Status QDataStream_Status(const QDataStream* self);
extern __declspec(dllexport) void QDataStream_SetStatus(QDataStream* self, Status status);
extern __declspec(dllexport) void QDataStream_ResetStatus(QDataStream* self);
extern __declspec(dllexport) FloatingPointPrecision QDataStream_FloatingPointPrecision(const QDataStream* self);
extern __declspec(dllexport) void QDataStream_SetFloatingPointPrecision(QDataStream* self, FloatingPointPrecision precision);
extern __declspec(dllexport) ByteOrder QDataStream_ByteOrder(const QDataStream* self);
extern __declspec(dllexport) void QDataStream_SetByteOrder(QDataStream* self, ByteOrder byteOrder);
extern __declspec(dllexport) int QDataStream_Version(const QDataStream* self);
extern __declspec(dllexport) void QDataStream_SetVersion(QDataStream* self, int version);
extern __declspec(dllexport) void QDataStream_OperatorShiftRight(QDataStream* self, char* i);
extern __declspec(dllexport) void QDataStream_OperatorShiftRightWithQint8(QDataStream* self, signed char* i);
extern __declspec(dllexport) void QDataStream_OperatorShiftRightWithQuint8(QDataStream* self, unsigned char* i);
extern __declspec(dllexport) void QDataStream_OperatorShiftRightWithQint16(QDataStream* self, int16_t* i);
extern __declspec(dllexport) void QDataStream_OperatorShiftRightWithQuint16(QDataStream* self, uint16_t* i);
extern __declspec(dllexport) void QDataStream_OperatorShiftRightWithQint32(QDataStream* self, int* i);
extern __declspec(dllexport) void QDataStream_OperatorShiftRightWithQuint32(QDataStream* self, unsigned int* i);
extern __declspec(dllexport) void QDataStream_OperatorShiftRightWithQint64(QDataStream* self, long long* i);
extern __declspec(dllexport) void QDataStream_OperatorShiftRightWithQuint64(QDataStream* self, unsigned long long* i);
extern __declspec(dllexport) void QDataStream_OperatorShiftRightWithBool(QDataStream* self, bool* i);
extern __declspec(dllexport) void QDataStream_OperatorShiftRightWithFloat(QDataStream* self, float* f);
extern __declspec(dllexport) void QDataStream_OperatorShiftRightWithDouble(QDataStream* self, double* f);
extern __declspec(dllexport) void QDataStream_OperatorShiftRightWithStr(QDataStream* self, char* str);
extern __declspec(dllexport) void QDataStream_OperatorShiftLeft(QDataStream* self, char i);
extern __declspec(dllexport) void QDataStream_OperatorShiftLeftWithQint8(QDataStream* self, signed char i);
extern __declspec(dllexport) void QDataStream_OperatorShiftLeftWithQuint8(QDataStream* self, unsigned char i);
extern __declspec(dllexport) void QDataStream_OperatorShiftLeftWithQint16(QDataStream* self, int16_t i);
extern __declspec(dllexport) void QDataStream_OperatorShiftLeftWithQuint16(QDataStream* self, uint16_t i);
extern __declspec(dllexport) void QDataStream_OperatorShiftLeftWithQint32(QDataStream* self, int i);
extern __declspec(dllexport) void QDataStream_OperatorShiftLeftWithQuint32(QDataStream* self, unsigned int i);
extern __declspec(dllexport) void QDataStream_OperatorShiftLeftWithQint64(QDataStream* self, long long i);
extern __declspec(dllexport) void QDataStream_OperatorShiftLeftWithQuint64(QDataStream* self, unsigned long long i);
extern __declspec(dllexport) void QDataStream_OperatorShiftLeftWithFloat(QDataStream* self, float f);
extern __declspec(dllexport) void QDataStream_OperatorShiftLeftWithDouble(QDataStream* self, double f);
extern __declspec(dllexport) void QDataStream_OperatorShiftLeftWithStr(QDataStream* self, const char* str);
extern __declspec(dllexport) QDataStream* QDataStream_ReadBytes(QDataStream* self, char* param1, unsigned int* lenVal);
extern __declspec(dllexport) QDataStream* QDataStream_ReadBytes2(QDataStream* self, char* param1, long long* lenVal);
extern __declspec(dllexport) long long QDataStream_ReadRawData(QDataStream* self, char* param1, long long lenVal);
extern __declspec(dllexport) void QDataStream_WriteBytes(QDataStream* self, const char* param1, long long lenVal);
extern __declspec(dllexport) long long QDataStream_WriteRawData(QDataStream* self, const char* param1, long long lenVal);
extern __declspec(dllexport) long long QDataStream_SkipRawData(QDataStream* self, long long lenVal);
extern __declspec(dllexport) void QDataStream_StartTransaction(QDataStream* self);
extern __declspec(dllexport) bool QDataStream_CommitTransaction(QDataStream* self);
extern __declspec(dllexport) void QDataStream_RollbackTransaction(QDataStream* self);
extern __declspec(dllexport) void QDataStream_AbortTransaction(QDataStream* self);
extern __declspec(dllexport) bool QDataStream_IsDeviceTransactionStarted(const QDataStream* self);
extern __declspec(dllexport) void QDataStream_Delete(QDataStream* self, bool isSubclass);

} 
