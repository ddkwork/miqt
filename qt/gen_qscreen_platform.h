#pragma once
#ifndef MIQT_QT_GEN_QSCREEN_PLATFORM_H
#define MIQT_QT_GEN_QSCREEN_PLATFORM_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QNativeInterface__QWindowsScreen)
typedef QNativeInterface::QWindowsScreen QNativeInterface__QWindowsScreen;
typedef struct QNativeInterface__QWindowsScreen QNativeInterface__QWindowsScreen;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QNativeInterface__QWindowsScreen* QNativeInterface__QWindowsScreen_new();
extern __declspec(dllexport) struct HMONITOR__* QNativeInterface__QWindowsScreen_Handle(const QNativeInterface__QWindowsScreen* self);
extern __declspec(dllexport) void QNativeInterface__QWindowsScreen_override_virtual_Handle(void* self, intptr_t slot);
struct HMONITOR__* QNativeInterface__QWindowsScreen_virtualbase_Handle(const void* self);

} 
