#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QHTTP1CONFIGURATION_H
#define MIQT_QT_NETWORK_GEN_QHTTP1CONFIGURATION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QHttp1Configuration QHttp1Configuration;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QHttp1Configuration* QHttp1Configuration_new();
extern __declspec(dllexport) QHttp1Configuration* QHttp1Configuration_new2(QHttp1Configuration* other);
extern __declspec(dllexport) void QHttp1Configuration_OperatorAssign(QHttp1Configuration* self, QHttp1Configuration* other);
extern __declspec(dllexport) void QHttp1Configuration_SetNumberOfConnectionsPerHost(QHttp1Configuration* self, ptrdiff_t amount);
extern __declspec(dllexport) ptrdiff_t QHttp1Configuration_NumberOfConnectionsPerHost(const QHttp1Configuration* self);
extern __declspec(dllexport) void QHttp1Configuration_Swap(QHttp1Configuration* self, QHttp1Configuration* other);
extern __declspec(dllexport) void QHttp1Configuration_Delete(QHttp1Configuration* self, bool isSubclass);

} 
