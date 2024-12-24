#pragma once
#ifndef MIQT_QT_CBOR_GEN_QCBORSTREAMWRITER_H
#define MIQT_QT_CBOR_GEN_QCBORSTREAMWRITER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QCborStreamWriter QCborStreamWriter;
typedef struct QIODevice QIODevice;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QCborStreamWriter* QCborStreamWriter_new(QIODevice* device);
extern __declspec(dllexport) void QCborStreamWriter_SetDevice(QCborStreamWriter* self, QIODevice* device);
extern __declspec(dllexport) QIODevice* QCborStreamWriter_Device(const QCborStreamWriter* self);
extern __declspec(dllexport) void QCborStreamWriter_Append(QCborStreamWriter* self, unsigned long long u);
extern __declspec(dllexport) void QCborStreamWriter_AppendWithQint64(QCborStreamWriter* self, long long i);
extern __declspec(dllexport) void QCborStreamWriter_AppendWithQCborNegativeInteger(QCborStreamWriter* self, uint64_t n);
extern __declspec(dllexport) void QCborStreamWriter_AppendWithBa(QCborStreamWriter* self, struct miqt_string ba);
extern __declspec(dllexport) void QCborStreamWriter_AppendWithTag(QCborStreamWriter* self, uint64_t tag);
extern __declspec(dllexport) void QCborStreamWriter_Append3(QCborStreamWriter* self, int tag);
extern __declspec(dllexport) void QCborStreamWriter_AppendWithSt(QCborStreamWriter* self, uint8_t st);
extern __declspec(dllexport) void QCborStreamWriter_AppendWithFloat(QCborStreamWriter* self, float f);
extern __declspec(dllexport) void QCborStreamWriter_AppendWithDouble(QCborStreamWriter* self, double d);
extern __declspec(dllexport) void QCborStreamWriter_AppendByteString(QCborStreamWriter* self, const char* data, ptrdiff_t lenVal);
extern __declspec(dllexport) void QCborStreamWriter_AppendTextString(QCborStreamWriter* self, const char* utf8, ptrdiff_t lenVal);
extern __declspec(dllexport) void QCborStreamWriter_AppendWithBool(QCborStreamWriter* self, bool b);
extern __declspec(dllexport) void QCborStreamWriter_AppendNull(QCborStreamWriter* self);
extern __declspec(dllexport) void QCborStreamWriter_AppendUndefined(QCborStreamWriter* self);
extern __declspec(dllexport) void QCborStreamWriter_AppendWithInt(QCborStreamWriter* self, int i);
extern __declspec(dllexport) void QCborStreamWriter_AppendWithUint(QCborStreamWriter* self, unsigned int u);
extern __declspec(dllexport) void QCborStreamWriter_Append4(QCborStreamWriter* self, const char* str);
extern __declspec(dllexport) void QCborStreamWriter_StartArray(QCborStreamWriter* self);
extern __declspec(dllexport) void QCborStreamWriter_StartArrayWithCount(QCborStreamWriter* self, unsigned long long count);
extern __declspec(dllexport) bool QCborStreamWriter_EndArray(QCborStreamWriter* self);
extern __declspec(dllexport) void QCborStreamWriter_StartMap(QCborStreamWriter* self);
extern __declspec(dllexport) void QCborStreamWriter_StartMapWithCount(QCborStreamWriter* self, unsigned long long count);
extern __declspec(dllexport) bool QCborStreamWriter_EndMap(QCborStreamWriter* self);
extern __declspec(dllexport) void QCborStreamWriter_Append22(QCborStreamWriter* self, const char* str, ptrdiff_t size);
extern __declspec(dllexport) void QCborStreamWriter_Delete(QCborStreamWriter* self, bool isSubclass);

} 
