#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QMEDIADEVICES_H
#define MIQT_QT_MULTIMEDIA_GEN_QMEDIADEVICES_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioDevice QAudioDevice;
typedef struct QCameraDevice QCameraDevice;
typedef struct QMediaDevices QMediaDevices;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QMediaDevices* QMediaDevices_new();
extern __declspec(dllexport) 
QMediaDevices* QMediaDevices_new2(QObject* parent);
extern __declspec(dllexport) 
void QMediaDevices_virtbase(QMediaDevices* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QMediaDevices_MetaObject(const QMediaDevices* self);
extern __declspec(dllexport) 
void* QMediaDevices_Metacast(QMediaDevices* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QMediaDevices_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_array /* of QAudioDevice* */  QMediaDevices_AudioInputs();
extern __declspec(dllexport) 
struct miqt_array /* of QAudioDevice* */  QMediaDevices_AudioOutputs();
extern __declspec(dllexport) 
struct miqt_array /* of QCameraDevice* */  QMediaDevices_VideoInputs();
extern __declspec(dllexport) 
QAudioDevice* QMediaDevices_DefaultAudioInput();
extern __declspec(dllexport) 
QAudioDevice* QMediaDevices_DefaultAudioOutput();
extern __declspec(dllexport) 
QCameraDevice* QMediaDevices_DefaultVideoInput();
extern __declspec(dllexport) 
void QMediaDevices_AudioInputsChanged(QMediaDevices* self);
void QMediaDevices_connect_AudioInputsChanged(QMediaDevices* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaDevices_AudioOutputsChanged(QMediaDevices* self);
void QMediaDevices_connect_AudioOutputsChanged(QMediaDevices* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaDevices_VideoInputsChanged(QMediaDevices* self);
void QMediaDevices_connect_VideoInputsChanged(QMediaDevices* self, intptr_t slot);
extern __declspec(dllexport) 
void QMediaDevices_ConnectNotify(QMediaDevices* self, QMetaMethod* signal);
extern __declspec(dllexport) 
struct miqt_string QMediaDevices_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QMediaDevices_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QMediaDevices_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QMediaDevices_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QMediaDevices_override_virtual_Metacast(void* self, intptr_t slot);
void* QMediaDevices_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QMediaDevices_Delete(QMediaDevices* self, bool isSubclass);

}
