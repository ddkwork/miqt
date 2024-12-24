#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QHTTP2CONFIGURATION_H
#define MIQT_QT_NETWORK_GEN_QHTTP2CONFIGURATION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QHttp2Configuration QHttp2Configuration;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QHttp2Configuration* QHttp2Configuration_new();
extern __declspec(dllexport) QHttp2Configuration* QHttp2Configuration_new2(QHttp2Configuration* other);
extern __declspec(dllexport) void QHttp2Configuration_OperatorAssign(QHttp2Configuration* self, QHttp2Configuration* other);
extern __declspec(dllexport) void QHttp2Configuration_SetServerPushEnabled(QHttp2Configuration* self, bool enable);
extern __declspec(dllexport) bool QHttp2Configuration_ServerPushEnabled(const QHttp2Configuration* self);
extern __declspec(dllexport) void QHttp2Configuration_SetHuffmanCompressionEnabled(QHttp2Configuration* self, bool enable);
extern __declspec(dllexport) bool QHttp2Configuration_HuffmanCompressionEnabled(const QHttp2Configuration* self);
extern __declspec(dllexport) bool QHttp2Configuration_SetSessionReceiveWindowSize(QHttp2Configuration* self, unsigned int size);
extern __declspec(dllexport) unsigned int QHttp2Configuration_SessionReceiveWindowSize(const QHttp2Configuration* self);
extern __declspec(dllexport) bool QHttp2Configuration_SetStreamReceiveWindowSize(QHttp2Configuration* self, unsigned int size);
extern __declspec(dllexport) unsigned int QHttp2Configuration_StreamReceiveWindowSize(const QHttp2Configuration* self);
extern __declspec(dllexport) bool QHttp2Configuration_SetMaxFrameSize(QHttp2Configuration* self, unsigned int size);
extern __declspec(dllexport) unsigned int QHttp2Configuration_MaxFrameSize(const QHttp2Configuration* self);
extern __declspec(dllexport) void QHttp2Configuration_SetMaxConcurrentStreams(QHttp2Configuration* self, unsigned int value);
extern __declspec(dllexport) unsigned int QHttp2Configuration_MaxConcurrentStreams(const QHttp2Configuration* self);
extern __declspec(dllexport) void QHttp2Configuration_Swap(QHttp2Configuration* self, QHttp2Configuration* other);
extern __declspec(dllexport) void QHttp2Configuration_Delete(QHttp2Configuration* self, bool isSubclass);

} 
