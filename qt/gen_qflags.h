#pragma once
#ifndef MIQT_QT_GEN_QFLAGS_H
#define MIQT_QT_GEN_QFLAGS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QFlag QFlag;
typedef struct QIncompatibleFlag QIncompatibleFlag;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QFlag* QFlag_new(int value);
extern __declspec(dllexport) QFlag* QFlag_new2(QFlag* param1);
extern __declspec(dllexport) void QFlag_Delete(QFlag* self, bool isSubclass);

extern __declspec(dllexport) QIncompatibleFlag* QIncompatibleFlag_new(int i);
extern __declspec(dllexport) QIncompatibleFlag* QIncompatibleFlag_new2(QIncompatibleFlag* param1);
extern __declspec(dllexport) void QIncompatibleFlag_Delete(QIncompatibleFlag* self, bool isSubclass);

} 
