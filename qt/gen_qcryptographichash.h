#pragma once
#ifndef MIQT_QT_GEN_QCRYPTOGRAPHICHASH_H
#define MIQT_QT_GEN_QCRYPTOGRAPHICHASH_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QByteArrayView QByteArrayView;
typedef struct QCryptographicHash QCryptographicHash;
typedef struct QIODevice QIODevice;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QCryptographicHash* QCryptographicHash_new(Algorithm method);
extern __declspec(dllexport) void QCryptographicHash_Swap(QCryptographicHash* self, QCryptographicHash* other);
extern __declspec(dllexport) void QCryptographicHash_Reset(QCryptographicHash* self);
extern __declspec(dllexport) Algorithm QCryptographicHash_Algorithm(const QCryptographicHash* self);
extern __declspec(dllexport) void QCryptographicHash_AddData(QCryptographicHash* self, const char* data, ptrdiff_t length);
extern __declspec(dllexport) void QCryptographicHash_AddDataWithData(QCryptographicHash* self, QByteArrayView* data);
extern __declspec(dllexport) bool QCryptographicHash_AddDataWithDevice(QCryptographicHash* self, QIODevice* device);
extern __declspec(dllexport) struct miqt_string QCryptographicHash_Result(const QCryptographicHash* self);
extern __declspec(dllexport) QByteArrayView* QCryptographicHash_ResultView(const QCryptographicHash* self);
extern __declspec(dllexport) struct miqt_string QCryptographicHash_Hash(QByteArrayView* data, Algorithm method);
extern __declspec(dllexport) QByteArrayView* QCryptographicHash_HashInto(QSpan<char> buffer, QByteArrayView* data, Algorithm method);
extern __declspec(dllexport) QByteArrayView* QCryptographicHash_HashInto2(QSpan<uchar> buffer, QByteArrayView* data, Algorithm method);
extern __declspec(dllexport) QByteArrayView* QCryptographicHash_HashInto3(QSpan<std::byte> buffer, QByteArrayView* data, Algorithm method);
extern __declspec(dllexport) QByteArrayView* QCryptographicHash_HashInto4(QSpan<char> buffer, QSpan<const QByteArrayView> data, Algorithm method);
extern __declspec(dllexport) QByteArrayView* QCryptographicHash_HashInto5(QSpan<uchar> buffer, QSpan<const QByteArrayView> data, Algorithm method);
extern __declspec(dllexport) QByteArrayView* QCryptographicHash_HashInto6(QSpan<std::byte> buffer, QSpan<const QByteArrayView> data, Algorithm method);
extern __declspec(dllexport) int QCryptographicHash_HashLength(Algorithm method);
extern __declspec(dllexport) bool QCryptographicHash_SupportsAlgorithm(Algorithm method);
extern __declspec(dllexport) void QCryptographicHash_Delete(QCryptographicHash* self, bool isSubclass);

} 
