#pragma once
#ifndef MIQT_QT_GEN_QRUNNABLE_H
#define MIQT_QT_GEN_QRUNNABLE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QGenericRunnable QGenericRunnable;
typedef struct QRunnable QRunnable;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QRunnable* QRunnable_new();
extern __declspec(dllexport) 
void QRunnable_Run(QRunnable* self);
extern __declspec(dllexport) 
bool QRunnable_AutoDelete(const QRunnable* self);
extern __declspec(dllexport) 
void QRunnable_SetAutoDelete(QRunnable* self, bool autoDelete);
extern __declspec(dllexport) 
void QRunnable_Delete(QRunnable* self, bool isSubclass);

extern __declspec(dllexport) 
void QGenericRunnable_virtbase(QGenericRunnable* src
, QRunnable** outptr_QRunnable
);
extern __declspec(dllexport) 
void QGenericRunnable_Run(QGenericRunnable* self);
extern __declspec(dllexport) 
void QGenericRunnable_Delete(QGenericRunnable* self, bool isSubclass);

}
