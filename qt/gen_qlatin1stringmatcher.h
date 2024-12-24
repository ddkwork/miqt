#pragma once
#ifndef MIQT_QT_GEN_QLATIN1STRINGMATCHER_H
#define MIQT_QT_GEN_QLATIN1STRINGMATCHER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QLatin1StringMatcher QLatin1StringMatcher;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QLatin1StringMatcher* QLatin1StringMatcher_new();
extern __declspec(dllexport) void QLatin1StringMatcher_SetCaseSensitivity(QLatin1StringMatcher* self, int cs);
extern __declspec(dllexport) int QLatin1StringMatcher_CaseSensitivity(const QLatin1StringMatcher* self);
extern __declspec(dllexport) void QLatin1StringMatcher_Delete(QLatin1StringMatcher* self, bool isSubclass);

} 
