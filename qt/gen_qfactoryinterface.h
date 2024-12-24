#pragma once
#ifndef MIQT_QT_GEN_QFACTORYINTERFACE_H
#define MIQT_QT_GEN_QFACTORYINTERFACE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QFactoryInterface;
class _GUID;
class type_info;
#else
typedef struct QFactoryInterface QFactoryInterface;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QFactoryInterface_Keys(const QFactoryInterface* self);
extern __declspec(dllexport) void QFactoryInterface_Delete(QFactoryInterface* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
