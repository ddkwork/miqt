#pragma once
#ifndef MIQT_QT_GEN_QPROPERTYPRIVATE_H
#define MIQT_QT_GEN_QPROPERTYPRIVATE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QPropertyProxyBindingData QPropertyProxyBindingData;
typedef struct QUntypedPropertyData QUntypedPropertyData;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QUntypedPropertyData_Delete(QUntypedPropertyData* self, bool isSubclass);

extern __declspec(dllexport) void QPropertyProxyBindingData_Delete(QPropertyProxyBindingData* self, bool isSubclass);

} 