#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QWINDOWCAPTURE_H
#define MIQT_QT_MULTIMEDIA_GEN_QWINDOWCAPTURE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QCapturableWindow QCapturableWindow;
typedef struct QMediaCaptureSession QMediaCaptureSession;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QWindowCapture QWindowCapture;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QWindowCapture* QWindowCapture_new();
extern __declspec(dllexport) 
QWindowCapture* QWindowCapture_new2(QObject* parent);
extern __declspec(dllexport) 
void QWindowCapture_virtbase(QWindowCapture* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QWindowCapture_MetaObject(const QWindowCapture* self);
extern __declspec(dllexport) 
void* QWindowCapture_Metacast(QWindowCapture* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QWindowCapture_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_array /* of QCapturableWindow* */  QWindowCapture_CapturableWindows();
extern __declspec(dllexport) 
QMediaCaptureSession* QWindowCapture_CaptureSession(const QWindowCapture* self);
extern __declspec(dllexport) 
void QWindowCapture_SetWindow(QWindowCapture* self, QCapturableWindow* window);
extern __declspec(dllexport) 
QCapturableWindow* QWindowCapture_Window(const QWindowCapture* self);
extern __declspec(dllexport) 
bool QWindowCapture_IsActive(const QWindowCapture* self);
extern __declspec(dllexport) 
Error QWindowCapture_Error(const QWindowCapture* self);
extern __declspec(dllexport) 
struct miqt_string QWindowCapture_ErrorString(const QWindowCapture* self);
extern __declspec(dllexport) 
void QWindowCapture_SetActive(QWindowCapture* self, bool active);
extern __declspec(dllexport) 
void QWindowCapture_Start(QWindowCapture* self);
extern __declspec(dllexport) 
void QWindowCapture_Stop(QWindowCapture* self);
extern __declspec(dllexport) 
void QWindowCapture_ActiveChanged(QWindowCapture* self, bool param1);
void QWindowCapture_connect_ActiveChanged(QWindowCapture* self, intptr_t slot);
extern __declspec(dllexport) 
void QWindowCapture_WindowChanged(QWindowCapture* self, QCapturableWindow* window);
void QWindowCapture_connect_WindowChanged(QWindowCapture* self, intptr_t slot);
extern __declspec(dllexport) 
void QWindowCapture_ErrorChanged(QWindowCapture* self);
void QWindowCapture_connect_ErrorChanged(QWindowCapture* self, intptr_t slot);
extern __declspec(dllexport) 
void QWindowCapture_ErrorOccurred(QWindowCapture* self, int error, struct miqt_string errorString);
void QWindowCapture_connect_ErrorOccurred(QWindowCapture* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QWindowCapture_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QWindowCapture_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QWindowCapture_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QWindowCapture_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QWindowCapture_override_virtual_Metacast(void* self, intptr_t slot);
void* QWindowCapture_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QWindowCapture_Delete(QWindowCapture* self, bool isSubclass);

}
