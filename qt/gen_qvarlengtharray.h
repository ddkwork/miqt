#pragma once
#ifndef MIQT_QT_GEN_QVARLENGTHARRAY_H
#define MIQT_QT_GEN_QVARLENGTHARRAY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QVLABaseBase QVLABaseBase;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) size_type QVLABaseBase_Capacity(const QVLABaseBase* self);
extern __declspec(dllexport) size_type QVLABaseBase_Size(const QVLABaseBase* self);
extern __declspec(dllexport) bool QVLABaseBase_Empty(const QVLABaseBase* self);

} 
