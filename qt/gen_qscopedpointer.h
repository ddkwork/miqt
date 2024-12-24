#pragma once
#ifndef MIQT_QT_GEN_QSCOPEDPOINTER_H
#define MIQT_QT_GEN_QSCOPEDPOINTER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QScopedPointerPodDeleter QScopedPointerPodDeleter;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
void QScopedPointerPodDeleter_Cleanup(void* pointer);
extern __declspec(dllexport) 
void QScopedPointerPodDeleter_OperatorCall(const QScopedPointerPodDeleter* self, void* pointer);
extern __declspec(dllexport) 
void QScopedPointerPodDeleter_Delete(QScopedPointerPodDeleter* self, bool isSubclass);

}
