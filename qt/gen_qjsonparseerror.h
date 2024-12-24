#pragma once
#ifndef MIQT_QT_GEN_QJSONPARSEERROR_H
#define MIQT_QT_GEN_QJSONPARSEERROR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QJsonParseError QJsonParseError;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) struct miqt_string QJsonParseError_ErrorString(const QJsonParseError* self);
extern __declspec(dllexport) void QJsonParseError_Delete(QJsonParseError* self, bool isSubclass);

} 
