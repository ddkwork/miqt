#pragma once
#ifndef MIQT_QT_SPATIALAUDIO_GEN_QAUDIOLISTENER_H
#define MIQT_QT_SPATIALAUDIO_GEN_QAUDIOLISTENER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioEngine QAudioEngine;
typedef struct QAudioListener QAudioListener;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QQuaternion QQuaternion;
typedef struct QVector3D QVector3D;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAudioListener* QAudioListener_new(QAudioEngine* engine);
extern __declspec(dllexport) 
void QAudioListener_virtbase(QAudioListener* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
void QAudioListener_SetPosition(QAudioListener* self, QVector3D* pos);
extern __declspec(dllexport) 
QVector3D* QAudioListener_Position(const QAudioListener* self);
extern __declspec(dllexport) 
void QAudioListener_SetRotation(QAudioListener* self, QQuaternion* q);
extern __declspec(dllexport) 
QQuaternion* QAudioListener_Rotation(const QAudioListener* self);
extern __declspec(dllexport) 
QAudioEngine* QAudioListener_Engine(const QAudioListener* self);
extern __declspec(dllexport) 
void QAudioListener_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAudioListener_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAudioListener_override_virtual_Metacast(void* self, intptr_t slot);
void* QAudioListener_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAudioListener_Delete(QAudioListener* self, bool isSubclass);

}
