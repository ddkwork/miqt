#pragma once
#ifndef MIQT_QT_GEN_QXPTYPE_TRAITS_H
#define MIQT_QT_GEN_QXPTYPE_TRAITS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class _GUID;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_qxp__nonesuch)
typedef qxp::nonesuch qxp__nonesuch;
#else
class qxp__nonesuch;
#endif
class type_info;
#else
typedef struct _GUID _GUID;
typedef struct qxp__nonesuch qxp__nonesuch;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);


#ifdef __cplusplus
} /* extern C */
#endif 

#endif
