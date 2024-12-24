#pragma once
#ifndef MIQT_QT_PRINTSUPPORT_GEN_QPRINTENGINE_H
#define MIQT_QT_PRINTSUPPORT_GEN_QPRINTENGINE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QPrintEngine QPrintEngine;
typedef struct QVariant QVariant;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QPrintEngine_SetProperty(QPrintEngine* self, PrintEnginePropertyKey key, QVariant* value);
extern __declspec(dllexport) QVariant* QPrintEngine_Property(const QPrintEngine* self, PrintEnginePropertyKey key);
extern __declspec(dllexport) bool QPrintEngine_NewPage(QPrintEngine* self);
extern __declspec(dllexport) bool QPrintEngine_Abort(QPrintEngine* self);
extern __declspec(dllexport) int QPrintEngine_Metric(const QPrintEngine* self, int param1);
extern __declspec(dllexport) int QPrintEngine_PrinterState(const QPrintEngine* self);
extern __declspec(dllexport) void QPrintEngine_OperatorAssign(QPrintEngine* self, QPrintEngine* param1);
extern __declspec(dllexport) void QPrintEngine_Delete(QPrintEngine* self, bool isSubclass);

} 
