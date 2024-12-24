#pragma once
#ifndef MIQT_QT_CBOR_GEN_QCBORSTREAMREADER_H
#define MIQT_QT_CBOR_GEN_QCBORSTREAMREADER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QCborError QCborError;
typedef struct QCborStreamReader QCborStreamReader;
typedef struct QIODevice QIODevice;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QCborStreamReader* QCborStreamReader_new();
extern __declspec(dllexport) QCborStreamReader* QCborStreamReader_new2(const char* data, ptrdiff_t lenVal);
extern __declspec(dllexport) QCborStreamReader* QCborStreamReader_new3(const unsigned char* data, ptrdiff_t lenVal);
extern __declspec(dllexport) QCborStreamReader* QCborStreamReader_new4(struct miqt_string data);
extern __declspec(dllexport) QCborStreamReader* QCborStreamReader_new5(QIODevice* device);
extern __declspec(dllexport) void QCborStreamReader_SetDevice(QCborStreamReader* self, QIODevice* device);
extern __declspec(dllexport) QIODevice* QCborStreamReader_Device(const QCborStreamReader* self);
extern __declspec(dllexport) void QCborStreamReader_AddData(QCborStreamReader* self, struct miqt_string data);
extern __declspec(dllexport) void QCborStreamReader_AddData2(QCborStreamReader* self, const char* data, ptrdiff_t lenVal);
extern __declspec(dllexport) void QCborStreamReader_AddData3(QCborStreamReader* self, const unsigned char* data, ptrdiff_t lenVal);
extern __declspec(dllexport) void QCborStreamReader_Reparse(QCborStreamReader* self);
extern __declspec(dllexport) void QCborStreamReader_Clear(QCborStreamReader* self);
extern __declspec(dllexport) void QCborStreamReader_Reset(QCborStreamReader* self);
extern __declspec(dllexport) QCborError* QCborStreamReader_LastError(const QCborStreamReader* self);
extern __declspec(dllexport) long long QCborStreamReader_CurrentOffset(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsValid(const QCborStreamReader* self);
extern __declspec(dllexport) int QCborStreamReader_ContainerDepth(const QCborStreamReader* self);
extern __declspec(dllexport) uint8_t QCborStreamReader_ParentContainerType(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_HasNext(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_Next(QCborStreamReader* self);
extern __declspec(dllexport) Type QCborStreamReader_Type(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsUnsignedInteger(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsNegativeInteger(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsInteger(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsByteArray(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsString(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsArray(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsMap(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsTag(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsSimpleType(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsFloat16(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsFloat(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsDouble(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsInvalid(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsSimpleTypeWithSt(const QCborStreamReader* self, uint8_t st);
extern __declspec(dllexport) bool QCborStreamReader_IsFalse(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsTrue(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsBool(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsNull(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsUndefined(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsLengthKnown(const QCborStreamReader* self);
extern __declspec(dllexport) unsigned long long QCborStreamReader_Length(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_IsContainer(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_EnterContainer(QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_LeaveContainer(QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_ReadAndAppendToString(QCborStreamReader* self, struct miqt_string dst);
extern __declspec(dllexport) bool QCborStreamReader_ReadAndAppendToUtf8String(QCborStreamReader* self, struct miqt_string dst);
extern __declspec(dllexport) bool QCborStreamReader_ReadAndAppendToByteArray(QCborStreamReader* self, struct miqt_string dst);
extern __declspec(dllexport) ptrdiff_t QCborStreamReader_CurrentStringChunkSize(const QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_ToBool(const QCborStreamReader* self);
extern __declspec(dllexport) uint64_t QCborStreamReader_ToTag(const QCborStreamReader* self);
extern __declspec(dllexport) unsigned long long QCborStreamReader_ToUnsignedInteger(const QCborStreamReader* self);
extern __declspec(dllexport) uint64_t QCborStreamReader_ToNegativeInteger(const QCborStreamReader* self);
extern __declspec(dllexport) uint8_t QCborStreamReader_ToSimpleType(const QCborStreamReader* self);
extern __declspec(dllexport) float QCborStreamReader_ToFloat(const QCborStreamReader* self);
extern __declspec(dllexport) double QCborStreamReader_ToDouble(const QCborStreamReader* self);
extern __declspec(dllexport) long long QCborStreamReader_ToInteger(const QCborStreamReader* self);
extern __declspec(dllexport) struct miqt_string QCborStreamReader_ReadAllString(QCborStreamReader* self);
extern __declspec(dllexport) struct miqt_string QCborStreamReader_ReadAllUtf8String(QCborStreamReader* self);
extern __declspec(dllexport) struct miqt_string QCborStreamReader_ReadAllByteArray(QCborStreamReader* self);
extern __declspec(dllexport) bool QCborStreamReader_Next1(QCborStreamReader* self, int maxRecursion);
extern __declspec(dllexport) void QCborStreamReader_Delete(QCborStreamReader* self, bool isSubclass);

} 
