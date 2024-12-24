#pragma once
#ifndef MIQT_QT_GEN_QIMAGEIOHANDLER_H
#define MIQT_QT_GEN_QIMAGEIOHANDLER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QIODevice QIODevice;
typedef struct QImage QImage;
typedef struct QImageIOHandler QImageIOHandler;
typedef struct QImageIOPlugin QImageIOPlugin;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QImageIOHandler* QImageIOHandler_new();
extern __declspec(dllexport) 
void QImageIOHandler_SetDevice(QImageIOHandler* self, QIODevice* device);
extern __declspec(dllexport) 
QIODevice* QImageIOHandler_Device(const QImageIOHandler* self);
extern __declspec(dllexport) 
void QImageIOHandler_SetFormat(QImageIOHandler* self, struct miqt_string format);
extern __declspec(dllexport) 
void QImageIOHandler_SetFormatWithFormat(const QImageIOHandler* self, struct miqt_string format);
extern __declspec(dllexport) 
struct miqt_string QImageIOHandler_Format(const QImageIOHandler* self);
extern __declspec(dllexport) 
bool QImageIOHandler_CanRead(const QImageIOHandler* self);
extern __declspec(dllexport) 
bool QImageIOHandler_Read(QImageIOHandler* self, QImage* image);
extern __declspec(dllexport) 
bool QImageIOHandler_Write(QImageIOHandler* self, QImage* image);
extern __declspec(dllexport) 
QVariant* QImageIOHandler_Option(const QImageIOHandler* self, ImageOption option);
extern __declspec(dllexport) 
void QImageIOHandler_SetOption(QImageIOHandler* self, ImageOption option, QVariant* value);
extern __declspec(dllexport) 
bool QImageIOHandler_SupportsOption(const QImageIOHandler* self, ImageOption option);
extern __declspec(dllexport) 
bool QImageIOHandler_JumpToNextImage(QImageIOHandler* self);
extern __declspec(dllexport) 
bool QImageIOHandler_JumpToImage(QImageIOHandler* self, int imageNumber);
extern __declspec(dllexport) 
int QImageIOHandler_LoopCount(const QImageIOHandler* self);
extern __declspec(dllexport) 
int QImageIOHandler_ImageCount(const QImageIOHandler* self);
extern __declspec(dllexport) 
int QImageIOHandler_NextImageDelay(const QImageIOHandler* self);
extern __declspec(dllexport) 
int QImageIOHandler_CurrentImageNumber(const QImageIOHandler* self);
extern __declspec(dllexport) 
QRect* QImageIOHandler_CurrentImageRect(const QImageIOHandler* self);
extern __declspec(dllexport) 
bool QImageIOHandler_AllocateImage(QSize* size, int format, QImage* image);
extern __declspec(dllexport) 
void QImageIOHandler_Delete(QImageIOHandler* self, bool isSubclass);

extern __declspec(dllexport) 
QImageIOPlugin* QImageIOPlugin_new();
extern __declspec(dllexport) 
QImageIOPlugin* QImageIOPlugin_new2(QObject* parent);
extern __declspec(dllexport) 
void QImageIOPlugin_virtbase(QImageIOPlugin* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QImageIOPlugin_MetaObject(const QImageIOPlugin* self);
extern __declspec(dllexport) 
void* QImageIOPlugin_Metacast(QImageIOPlugin* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QImageIOPlugin_Tr(const char* s);
extern __declspec(dllexport) 
Capabilities QImageIOPlugin_Capabilities(const QImageIOPlugin* self, QIODevice* device, struct miqt_string format);
extern __declspec(dllexport) 
QImageIOHandler* QImageIOPlugin_Create(const QImageIOPlugin* self, QIODevice* device, struct miqt_string format);
extern __declspec(dllexport) 
struct miqt_string QImageIOPlugin_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QImageIOPlugin_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QImageIOPlugin_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QImageIOPlugin_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QImageIOPlugin_override_virtual_Metacast(void* self, intptr_t slot);
void* QImageIOPlugin_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QImageIOPlugin_Delete(QImageIOPlugin* self, bool isSubclass);

}
