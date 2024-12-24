#pragma once
#ifndef MIQT_QT_GEN_QFACTORYINTERFACE_H
#define MIQT_QT_GEN_QFACTORYINTERFACE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QFactoryInterface QFactoryInterface;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QFactoryInterface_Keys(const QFactoryInterface* self);
extern __declspec(dllexport) 
void QFactoryInterface_Delete(QFactoryInterface* self, bool isSubclass);

}
