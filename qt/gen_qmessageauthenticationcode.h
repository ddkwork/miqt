#pragma once
#ifndef MIQT_QT_GEN_QMESSAGEAUTHENTICATIONCODE_H
#define MIQT_QT_GEN_QMESSAGEAUTHENTICATIONCODE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QByteArrayView QByteArrayView;
typedef struct QIODevice QIODevice;
typedef struct QMessageAuthenticationCode QMessageAuthenticationCode;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QMessageAuthenticationCode* QMessageAuthenticationCode_new(int method);
extern __declspec(dllexport) QMessageAuthenticationCode* QMessageAuthenticationCode_new2(int method, QByteArrayView* key);
extern __declspec(dllexport) void QMessageAuthenticationCode_Swap(QMessageAuthenticationCode* self, QMessageAuthenticationCode* other);
extern __declspec(dllexport) void QMessageAuthenticationCode_Reset(QMessageAuthenticationCode* self);
extern __declspec(dllexport) void QMessageAuthenticationCode_SetKey(QMessageAuthenticationCode* self, QByteArrayView* key);
extern __declspec(dllexport) void QMessageAuthenticationCode_AddData(QMessageAuthenticationCode* self, const char* data, ptrdiff_t length);
extern __declspec(dllexport) void QMessageAuthenticationCode_AddDataWithData(QMessageAuthenticationCode* self, QByteArrayView* data);
extern __declspec(dllexport) bool QMessageAuthenticationCode_AddDataWithDevice(QMessageAuthenticationCode* self, QIODevice* device);
extern __declspec(dllexport) QByteArrayView* QMessageAuthenticationCode_ResultView(const QMessageAuthenticationCode* self);
extern __declspec(dllexport) struct miqt_string QMessageAuthenticationCode_Result(const QMessageAuthenticationCode* self);
extern __declspec(dllexport) struct miqt_string QMessageAuthenticationCode_Hash(QByteArrayView* message, QByteArrayView* key, int method);
extern __declspec(dllexport) QByteArrayView* QMessageAuthenticationCode_HashInto(QSpan<char> buffer, QByteArrayView* message, QByteArrayView* key, int method);
extern __declspec(dllexport) QByteArrayView* QMessageAuthenticationCode_HashInto2(QSpan<uchar> buffer, QByteArrayView* message, QByteArrayView* key, int method);
extern __declspec(dllexport) QByteArrayView* QMessageAuthenticationCode_HashInto3(QSpan<std::byte> buffer, QByteArrayView* message, QByteArrayView* key, int method);
extern __declspec(dllexport) QByteArrayView* QMessageAuthenticationCode_HashInto4(QSpan<char> buffer, QSpan<const QByteArrayView> messageParts, QByteArrayView* key, int method);
extern __declspec(dllexport) QByteArrayView* QMessageAuthenticationCode_HashInto5(QSpan<uchar> buffer, QSpan<const QByteArrayView> messageParts, QByteArrayView* key, int method);
extern __declspec(dllexport) QByteArrayView* QMessageAuthenticationCode_HashInto6(QSpan<std::byte> buffer, QSpan<const QByteArrayView> message, QByteArrayView* key, int method);
extern __declspec(dllexport) void QMessageAuthenticationCode_Delete(QMessageAuthenticationCode* self, bool isSubclass);

} 
