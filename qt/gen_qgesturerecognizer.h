#pragma once
#ifndef MIQT_QT_GEN_QGESTURERECOGNIZER_H
#define MIQT_QT_GEN_QGESTURERECOGNIZER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEvent QEvent;
typedef struct QGesture QGesture;
typedef struct QGestureRecognizer QGestureRecognizer;
typedef struct QObject QObject;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QGestureRecognizer* QGestureRecognizer_new();
extern __declspec(dllexport) QGesture* QGestureRecognizer_Create(QGestureRecognizer* self, QObject* target);
extern __declspec(dllexport) Result QGestureRecognizer_Recognize(QGestureRecognizer* self, QGesture* state, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QGestureRecognizer_Reset(QGestureRecognizer* self, QGesture* state);
extern __declspec(dllexport) int QGestureRecognizer_RegisterRecognizer(QGestureRecognizer* recognizer);
extern __declspec(dllexport) void QGestureRecognizer_UnregisterRecognizer(int typeVal);
extern __declspec(dllexport) void QGestureRecognizer_OperatorAssign(QGestureRecognizer* self, QGestureRecognizer* param1);
extern __declspec(dllexport) void QGestureRecognizer_override_virtual_Create(void* self, intptr_t slot);
QGesture* QGestureRecognizer_virtualbase_Create(void* self, QObject* target);
extern __declspec(dllexport) void QGestureRecognizer_override_virtual_Recognize(void* self, intptr_t slot);
Result QGestureRecognizer_virtualbase_Recognize(void* self, QGesture* state, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QGestureRecognizer_override_virtual_Reset(void* self, intptr_t slot);
void QGestureRecognizer_virtualbase_Reset(void* self, QGesture* state);
extern __declspec(dllexport) void QGestureRecognizer_Delete(QGestureRecognizer* self, bool isSubclass);

} 
