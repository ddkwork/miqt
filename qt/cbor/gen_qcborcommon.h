#pragma once
#ifndef MIQT_QT_CBOR_GEN_QCBORCOMMON_H
#define MIQT_QT_CBOR_GEN_QCBORCOMMON_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QCborError QCborError;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
struct miqt_string QCborError_ToString(const QCborError* self);
extern __declspec(dllexport) 
void QCborError_Delete(QCborError* self, bool isSubclass);

}
