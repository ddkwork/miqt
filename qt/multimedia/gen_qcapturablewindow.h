#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QCAPTURABLEWINDOW_H
#define MIQT_QT_MULTIMEDIA_GEN_QCAPTURABLEWINDOW_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QCapturableWindow QCapturableWindow;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QCapturableWindow* QCapturableWindow_new();
extern __declspec(dllexport) 
QCapturableWindow* QCapturableWindow_new2(QCapturableWindow* other);
extern __declspec(dllexport) 
void QCapturableWindow_OperatorAssign(QCapturableWindow* self, QCapturableWindow* other);
extern __declspec(dllexport) 
void QCapturableWindow_Swap(QCapturableWindow* self, QCapturableWindow* other);
extern __declspec(dllexport) 
bool QCapturableWindow_IsValid(const QCapturableWindow* self);
extern __declspec(dllexport) 
struct miqt_string QCapturableWindow_Description(const QCapturableWindow* self);
extern __declspec(dllexport) 
void QCapturableWindow_Delete(QCapturableWindow* self, bool isSubclass);

}
