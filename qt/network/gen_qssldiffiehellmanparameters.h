#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QSSLDIFFIEHELLMANPARAMETERS_H
#define MIQT_QT_NETWORK_GEN_QSSLDIFFIEHELLMANPARAMETERS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QIODevice QIODevice;
typedef struct QSslDiffieHellmanParameters QSslDiffieHellmanParameters;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSslDiffieHellmanParameters* QSslDiffieHellmanParameters_new();
extern __declspec(dllexport) QSslDiffieHellmanParameters* QSslDiffieHellmanParameters_new2(QSslDiffieHellmanParameters* other);
extern __declspec(dllexport) QSslDiffieHellmanParameters* QSslDiffieHellmanParameters_DefaultParameters();
extern __declspec(dllexport) void QSslDiffieHellmanParameters_OperatorAssign(QSslDiffieHellmanParameters* self, QSslDiffieHellmanParameters* other);
extern __declspec(dllexport) void QSslDiffieHellmanParameters_Swap(QSslDiffieHellmanParameters* self, QSslDiffieHellmanParameters* other);
extern __declspec(dllexport) QSslDiffieHellmanParameters* QSslDiffieHellmanParameters_FromEncoded(struct miqt_string encoded);
extern __declspec(dllexport) QSslDiffieHellmanParameters* QSslDiffieHellmanParameters_FromEncodedWithDevice(QIODevice* device);
extern __declspec(dllexport) bool QSslDiffieHellmanParameters_IsEmpty(const QSslDiffieHellmanParameters* self);
extern __declspec(dllexport) bool QSslDiffieHellmanParameters_IsValid(const QSslDiffieHellmanParameters* self);
extern __declspec(dllexport) Error QSslDiffieHellmanParameters_Error(const QSslDiffieHellmanParameters* self);
extern __declspec(dllexport) struct miqt_string QSslDiffieHellmanParameters_ErrorString(const QSslDiffieHellmanParameters* self);
extern __declspec(dllexport) QSslDiffieHellmanParameters* QSslDiffieHellmanParameters_FromEncoded2(struct miqt_string encoded, int format);
extern __declspec(dllexport) QSslDiffieHellmanParameters* QSslDiffieHellmanParameters_FromEncoded22(QIODevice* device, int format);
extern __declspec(dllexport) void QSslDiffieHellmanParameters_Delete(QSslDiffieHellmanParameters* self, bool isSubclass);

} 
