#pragma once
#ifndef MIQT_QT_GEN_QCOMMANDLINEOPTION_H
#define MIQT_QT_GEN_QCOMMANDLINEOPTION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QCommandLineOption QCommandLineOption;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QCommandLineOption* QCommandLineOption_new(struct miqt_string name);
extern __declspec(dllexport) QCommandLineOption* QCommandLineOption_new2(struct miqt_array /* of struct miqt_string */  names);
extern __declspec(dllexport) QCommandLineOption* QCommandLineOption_new3(struct miqt_string name, struct miqt_string description);
extern __declspec(dllexport) QCommandLineOption* QCommandLineOption_new4(struct miqt_array /* of struct miqt_string */  names, struct miqt_string description);
extern __declspec(dllexport) QCommandLineOption* QCommandLineOption_new5(QCommandLineOption* other);
extern __declspec(dllexport) QCommandLineOption* QCommandLineOption_new6(struct miqt_string name, struct miqt_string description, struct miqt_string valueName);
extern __declspec(dllexport) QCommandLineOption* QCommandLineOption_new7(struct miqt_string name, struct miqt_string description, struct miqt_string valueName, struct miqt_string defaultValue);
extern __declspec(dllexport) QCommandLineOption* QCommandLineOption_new8(struct miqt_array /* of struct miqt_string */  names, struct miqt_string description, struct miqt_string valueName);
extern __declspec(dllexport) QCommandLineOption* QCommandLineOption_new9(struct miqt_array /* of struct miqt_string */  names, struct miqt_string description, struct miqt_string valueName, struct miqt_string defaultValue);
extern __declspec(dllexport) void QCommandLineOption_OperatorAssign(QCommandLineOption* self, QCommandLineOption* other);
extern __declspec(dllexport) void QCommandLineOption_Swap(QCommandLineOption* self, QCommandLineOption* other);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QCommandLineOption_Names(const QCommandLineOption* self);
extern __declspec(dllexport) void QCommandLineOption_SetValueName(QCommandLineOption* self, struct miqt_string name);
extern __declspec(dllexport) struct miqt_string QCommandLineOption_ValueName(const QCommandLineOption* self);
extern __declspec(dllexport) void QCommandLineOption_SetDescription(QCommandLineOption* self, struct miqt_string description);
extern __declspec(dllexport) struct miqt_string QCommandLineOption_Description(const QCommandLineOption* self);
extern __declspec(dllexport) void QCommandLineOption_SetDefaultValue(QCommandLineOption* self, struct miqt_string defaultValue);
extern __declspec(dllexport) void QCommandLineOption_SetDefaultValues(QCommandLineOption* self, struct miqt_array /* of struct miqt_string */  defaultValues);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QCommandLineOption_DefaultValues(const QCommandLineOption* self);
extern __declspec(dllexport) Flags QCommandLineOption_Flags(const QCommandLineOption* self);
extern __declspec(dllexport) void QCommandLineOption_SetFlags(QCommandLineOption* self, Flags aflags);
extern __declspec(dllexport) void QCommandLineOption_Delete(QCommandLineOption* self, bool isSubclass);

} 
