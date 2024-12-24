#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QSCREENCAPTURE_H
#define MIQT_QT_MULTIMEDIA_GEN_QSCREENCAPTURE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMediaCaptureSession QMediaCaptureSession;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QScreen QScreen;
typedef struct QScreenCapture QScreenCapture;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QScreenCapture* QScreenCapture_new();
extern __declspec(dllexport) 
QScreenCapture* QScreenCapture_new2(QObject* parent);
extern __declspec(dllexport) 
void QScreenCapture_virtbase(QScreenCapture* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QScreenCapture_MetaObject(const QScreenCapture* self);
extern __declspec(dllexport) 
void* QScreenCapture_Metacast(QScreenCapture* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QScreenCapture_Tr(const char* s);
extern __declspec(dllexport) 
QMediaCaptureSession* QScreenCapture_CaptureSession(const QScreenCapture* self);
extern __declspec(dllexport) 
void QScreenCapture_SetScreen(QScreenCapture* self, QScreen* screen);
extern __declspec(dllexport) 
QScreen* QScreenCapture_Screen(const QScreenCapture* self);
extern __declspec(dllexport) 
bool QScreenCapture_IsActive(const QScreenCapture* self);
extern __declspec(dllexport) 
Error QScreenCapture_Error(const QScreenCapture* self);
extern __declspec(dllexport) 
struct miqt_string QScreenCapture_ErrorString(const QScreenCapture* self);
extern __declspec(dllexport) 
void QScreenCapture_SetActive(QScreenCapture* self, bool active);
extern __declspec(dllexport) 
void QScreenCapture_Start(QScreenCapture* self);
extern __declspec(dllexport) 
void QScreenCapture_Stop(QScreenCapture* self);
extern __declspec(dllexport) 
void QScreenCapture_ActiveChanged(QScreenCapture* self, bool param1);
void QScreenCapture_connect_ActiveChanged(QScreenCapture* self, intptr_t slot);
extern __declspec(dllexport) 
void QScreenCapture_ErrorChanged(QScreenCapture* self);
void QScreenCapture_connect_ErrorChanged(QScreenCapture* self, intptr_t slot);
extern __declspec(dllexport) 
void QScreenCapture_ScreenChanged(QScreenCapture* self, QScreen* param1);
void QScreenCapture_connect_ScreenChanged(QScreenCapture* self, intptr_t slot);
extern __declspec(dllexport) 
void QScreenCapture_ErrorOccurred(QScreenCapture* self, int error, struct miqt_string errorString);
void QScreenCapture_connect_ErrorOccurred(QScreenCapture* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QScreenCapture_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QScreenCapture_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QScreenCapture_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QScreenCapture_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QScreenCapture_override_virtual_Metacast(void* self, intptr_t slot);
void* QScreenCapture_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QScreenCapture_Delete(QScreenCapture* self, bool isSubclass);

}
