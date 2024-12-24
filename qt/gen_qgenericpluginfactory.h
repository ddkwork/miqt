#pragma once
#ifndef MIQT_QT_GEN_QGENERICPLUGINFACTORY_H
#define MIQT_QT_GEN_QGENERICPLUGINFACTORY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QGenericPluginFactory QGenericPluginFactory;
typedef struct QObject QObject;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QGenericPluginFactory_Keys();
extern __declspec(dllexport) QObject* QGenericPluginFactory_Create(struct miqt_string param1, struct miqt_string param2);
extern __declspec(dllexport) void QGenericPluginFactory_Delete(QGenericPluginFactory* self, bool isSubclass);

} 
