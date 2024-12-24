#pragma once
#ifndef MIQT_QT_GEN_QFLOAT16_H
#define MIQT_QT_GEN_QFLOAT16_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct _GUID _GUID;
typedef struct qfloat16 qfloat16;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) qfloat16* qfloat16_new();
extern __declspec(dllexport) qfloat16* qfloat16_new2(int param1);
extern __declspec(dllexport) qfloat16* qfloat16_new3(float f);
extern __declspec(dllexport) bool qfloat16_IsInf(const qfloat16* self);
extern __declspec(dllexport) bool qfloat16_IsNaN(const qfloat16* self);
extern __declspec(dllexport) bool qfloat16_IsFinite(const qfloat16* self);
extern __declspec(dllexport) int qfloat16_FpClassify(const qfloat16* self);
extern __declspec(dllexport) bool qfloat16_IsNormal(const qfloat16* self);
extern __declspec(dllexport) void qfloat16_Delete(qfloat16* self, bool isSubclass);

} 
