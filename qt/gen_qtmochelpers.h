#pragma once
#ifndef MIQT_QT_GEN_QTMOCHELPERS_H
#define MIQT_QT_GEN_QTMOCHELPERS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QtMocHelpers__NoType)
typedef QtMocHelpers::NoType QtMocHelpers__NoType;
#else
class QtMocHelpers__NoType;
#endif
class _GUID;
class type_info;
#else
typedef struct QtMocHelpers__NoType QtMocHelpers__NoType;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QtMocHelpers__NoType_Delete(QtMocHelpers__NoType* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
