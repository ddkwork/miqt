#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QMEDIAMETADATA_H
#define MIQT_QT_MULTIMEDIA_GEN_QMEDIAMETADATA_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMediaMetaData QMediaMetaData;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QMediaMetaData* QMediaMetaData_new(QMediaMetaData* param1);
extern __declspec(dllexport) 
QMediaMetaData* QMediaMetaData_new2();
extern __declspec(dllexport) 
QVariant* QMediaMetaData_Value(const QMediaMetaData* self, Key k);
extern __declspec(dllexport) 
void QMediaMetaData_Insert(QMediaMetaData* self, Key k, QVariant* value);
extern __declspec(dllexport) 
void QMediaMetaData_Remove(QMediaMetaData* self, Key k);
extern __declspec(dllexport) 
struct miqt_array /* of Key */  QMediaMetaData_Keys(const QMediaMetaData* self);
extern __declspec(dllexport) 
QVariant* QMediaMetaData_OperatorSubscript(QMediaMetaData* self, Key k);
extern __declspec(dllexport) 
void QMediaMetaData_Clear(QMediaMetaData* self);
extern __declspec(dllexport) 
bool QMediaMetaData_IsEmpty(const QMediaMetaData* self);
extern __declspec(dllexport) 
struct miqt_string QMediaMetaData_StringValue(const QMediaMetaData* self, Key k);
extern __declspec(dllexport) 
struct miqt_string QMediaMetaData_MetaDataKeyToString(Key k);
extern __declspec(dllexport) 
void QMediaMetaData_Delete(QMediaMetaData* self, bool isSubclass);

}
