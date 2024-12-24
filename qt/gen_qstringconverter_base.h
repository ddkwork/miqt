#pragma once
#ifndef MIQT_QT_GEN_QSTRINGCONVERTER_BASE_H
#define MIQT_QT_GEN_QSTRINGCONVERTER_BASE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QStringConverterBase__State)
typedef QStringConverterBase::State QStringConverterBase__State;
typedef struct QStringConverter QStringConverter;
typedef struct QStringConverterBase QStringConverterBase;
typedef struct QStringConverterBase__State QStringConverterBase__State;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QStringConverterBase* QStringConverterBase_new(QStringConverterBase* param1);
extern __declspec(dllexport) 
QStringConverterBase* QStringConverterBase_new2();

extern __declspec(dllexport) 
void QStringConverter_virtbase(QStringConverter* src
, QStringConverterBase** outptr_QStringConverterBase
);
extern __declspec(dllexport) 
bool QStringConverter_IsValid(const QStringConverter* self);
extern __declspec(dllexport) 
void QStringConverter_ResetState(QStringConverter* self);
extern __declspec(dllexport) 
bool QStringConverter_HasError(const QStringConverter* self);
extern __declspec(dllexport) 
const char* QStringConverter_Name(const QStringConverter* self);
extern __declspec(dllexport) 
const char* QStringConverter_NameForEncoding(Encoding e);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QStringConverter_AvailableCodecs();

extern __declspec(dllexport) 
QStringConverterBase__State* QStringConverterBase__State_new();
extern __declspec(dllexport) 
QStringConverterBase__State* QStringConverterBase__State_new2(Flags f);
extern __declspec(dllexport) 
void QStringConverterBase__State_Clear(QStringConverterBase__State* self);
extern __declspec(dllexport) 
void QStringConverterBase__State_Reset(QStringConverterBase__State* self);
extern __declspec(dllexport) 
void QStringConverterBase__State_Delete(QStringConverterBase__State* self, bool isSubclass);

}
