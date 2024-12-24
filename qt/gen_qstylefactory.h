#pragma once
#ifndef MIQT_QT_GEN_QSTYLEFACTORY_H
#define MIQT_QT_GEN_QSTYLEFACTORY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QStyle QStyle;
typedef struct QStyleFactory QStyleFactory;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QStyleFactory_Keys();
extern __declspec(dllexport) 
QStyle* QStyleFactory_Create(struct miqt_string param1);
extern __declspec(dllexport) 
void QStyleFactory_Delete(QStyleFactory* self, bool isSubclass);

}
