#pragma once
#ifndef MIQT_QT_GEN_QBINDINGSTORAGE_H
#define MIQT_QT_GEN_QBINDINGSTORAGE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QBindingStatus QBindingStatus;
typedef struct QBindingStorage QBindingStorage;
typedef struct QUntypedPropertyData QUntypedPropertyData;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QBindingStatus_Delete(QBindingStatus* self, bool isSubclass);

extern __declspec(dllexport) QBindingStorage* QBindingStorage_new();
extern __declspec(dllexport) bool QBindingStorage_IsEmpty(QBindingStorage* self);
extern __declspec(dllexport) bool QBindingStorage_IsValid(const QBindingStorage* self);
extern __declspec(dllexport) void QBindingStorage_RegisterDependency(const QBindingStorage* self, QUntypedPropertyData* data);
extern __declspec(dllexport) void QBindingStorage_Delete(QBindingStorage* self, bool isSubclass);

} 
