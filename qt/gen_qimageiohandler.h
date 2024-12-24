#pragma once
#ifndef MIQT_QT_GEN_QIMAGEIOHANDLER_H
#define MIQT_QT_GEN_QIMAGEIOHANDLER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QIODevice QIODevice;
typedef struct QImage QImage;
typedef struct QImageIOHandler QImageIOHandler;
typedef struct QImageIOPlugin QImageIOPlugin;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QTimerEvent QTimerEvent;
typedef struct QVariant QVariant;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QImageIOHandler* QImageIOHandler_new();
extern __declspec(dllexport) void QImageIOHandler_SetDevice(QImageIOHandler* self, QIODevice* device);
extern __declspec(dllexport) QIODevice* QImageIOHandler_Device(const QImageIOHandler* self);
extern __declspec(dllexport) void QImageIOHandler_SetFormat(QImageIOHandler* self, struct miqt_string format);
extern __declspec(dllexport) void QImageIOHandler_SetFormatWithFormat(const QImageIOHandler* self, struct miqt_string format);
extern __declspec(dllexport) struct miqt_string QImageIOHandler_Format(const QImageIOHandler* self);
extern __declspec(dllexport) bool QImageIOHandler_CanRead(const QImageIOHandler* self);
extern __declspec(dllexport) bool QImageIOHandler_Read(QImageIOHandler* self, QImage* image);
extern __declspec(dllexport) bool QImageIOHandler_Write(QImageIOHandler* self, QImage* image);
extern __declspec(dllexport) QVariant* QImageIOHandler_Option(const QImageIOHandler* self, ImageOption option);
extern __declspec(dllexport) void QImageIOHandler_SetOption(QImageIOHandler* self, ImageOption option, QVariant* value);
extern __declspec(dllexport) bool QImageIOHandler_SupportsOption(const QImageIOHandler* self, ImageOption option);
extern __declspec(dllexport) bool QImageIOHandler_JumpToNextImage(QImageIOHandler* self);
extern __declspec(dllexport) bool QImageIOHandler_JumpToImage(QImageIOHandler* self, int imageNumber);
extern __declspec(dllexport) int QImageIOHandler_LoopCount(const QImageIOHandler* self);
extern __declspec(dllexport) int QImageIOHandler_ImageCount(const QImageIOHandler* self);
extern __declspec(dllexport) int QImageIOHandler_NextImageDelay(const QImageIOHandler* self);
extern __declspec(dllexport) int QImageIOHandler_CurrentImageNumber(const QImageIOHandler* self);
extern __declspec(dllexport) QRect* QImageIOHandler_CurrentImageRect(const QImageIOHandler* self);
extern __declspec(dllexport) bool QImageIOHandler_AllocateImage(QSize* size, int format, QImage* image);
extern __declspec(dllexport) void QImageIOHandler_override_virtual_CanRead(void* self, intptr_t slot);
bool QImageIOHandler_virtualbase_CanRead(const void* self);
extern __declspec(dllexport) void QImageIOHandler_override_virtual_Read(void* self, intptr_t slot);
bool QImageIOHandler_virtualbase_Read(void* self, QImage* image);
extern __declspec(dllexport) void QImageIOHandler_override_virtual_Write(void* self, intptr_t slot);
bool QImageIOHandler_virtualbase_Write(void* self, QImage* image);
extern __declspec(dllexport) void QImageIOHandler_override_virtual_Option(void* self, intptr_t slot);
QVariant* QImageIOHandler_virtualbase_Option(const void* self, ImageOption option);
extern __declspec(dllexport) void QImageIOHandler_override_virtual_SetOption(void* self, intptr_t slot);
void QImageIOHandler_virtualbase_SetOption(void* self, ImageOption option, QVariant* value);
extern __declspec(dllexport) void QImageIOHandler_override_virtual_SupportsOption(void* self, intptr_t slot);
bool QImageIOHandler_virtualbase_SupportsOption(const void* self, ImageOption option);
extern __declspec(dllexport) void QImageIOHandler_override_virtual_JumpToNextImage(void* self, intptr_t slot);
bool QImageIOHandler_virtualbase_JumpToNextImage(void* self);
extern __declspec(dllexport) void QImageIOHandler_override_virtual_JumpToImage(void* self, intptr_t slot);
bool QImageIOHandler_virtualbase_JumpToImage(void* self, int imageNumber);
extern __declspec(dllexport) void QImageIOHandler_override_virtual_LoopCount(void* self, intptr_t slot);
int QImageIOHandler_virtualbase_LoopCount(const void* self);
extern __declspec(dllexport) void QImageIOHandler_override_virtual_ImageCount(void* self, intptr_t slot);
int QImageIOHandler_virtualbase_ImageCount(const void* self);
extern __declspec(dllexport) void QImageIOHandler_override_virtual_NextImageDelay(void* self, intptr_t slot);
int QImageIOHandler_virtualbase_NextImageDelay(const void* self);
extern __declspec(dllexport) void QImageIOHandler_override_virtual_CurrentImageNumber(void* self, intptr_t slot);
int QImageIOHandler_virtualbase_CurrentImageNumber(const void* self);
extern __declspec(dllexport) void QImageIOHandler_override_virtual_CurrentImageRect(void* self, intptr_t slot);
QRect* QImageIOHandler_virtualbase_CurrentImageRect(const void* self);
extern __declspec(dllexport) void QImageIOHandler_Delete(QImageIOHandler* self, bool isSubclass);

extern __declspec(dllexport) QImageIOPlugin* QImageIOPlugin_new();
extern __declspec(dllexport) QImageIOPlugin* QImageIOPlugin_new2(QObject* parent);
extern __declspec(dllexport) void QImageIOPlugin_virtbase(QImageIOPlugin* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QImageIOPlugin_MetaObject(const QImageIOPlugin* self);
extern __declspec(dllexport) void* QImageIOPlugin_Metacast(QImageIOPlugin* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QImageIOPlugin_Tr(const char* s);
extern __declspec(dllexport) Capabilities QImageIOPlugin_Capabilities(const QImageIOPlugin* self, QIODevice* device, struct miqt_string format);
extern __declspec(dllexport) QImageIOHandler* QImageIOPlugin_Create(const QImageIOPlugin* self, QIODevice* device, struct miqt_string format);
extern __declspec(dllexport) struct miqt_string QImageIOPlugin_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QImageIOPlugin_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QImageIOPlugin_override_virtual_Capabilities(void* self, intptr_t slot);
Capabilities QImageIOPlugin_virtualbase_Capabilities(const void* self, QIODevice* device, struct miqt_string format);
extern __declspec(dllexport) void QImageIOPlugin_override_virtual_Create(void* self, intptr_t slot);
QImageIOHandler* QImageIOPlugin_virtualbase_Create(const void* self, QIODevice* device, struct miqt_string format);
extern __declspec(dllexport) void QImageIOPlugin_override_virtual_Event(void* self, intptr_t slot);
bool QImageIOPlugin_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QImageIOPlugin_override_virtual_EventFilter(void* self, intptr_t slot);
bool QImageIOPlugin_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QImageIOPlugin_override_virtual_TimerEvent(void* self, intptr_t slot);
void QImageIOPlugin_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QImageIOPlugin_override_virtual_ChildEvent(void* self, intptr_t slot);
void QImageIOPlugin_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QImageIOPlugin_override_virtual_CustomEvent(void* self, intptr_t slot);
void QImageIOPlugin_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QImageIOPlugin_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QImageIOPlugin_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QImageIOPlugin_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QImageIOPlugin_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QImageIOPlugin_Delete(QImageIOPlugin* self, bool isSubclass);

} 
