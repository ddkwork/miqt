#pragma once
#ifndef MIQT_QT_GEN_QSYSINFO_H
#define MIQT_QT_GEN_QSYSINFO_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QSysInfo QSysInfo;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) struct miqt_string QSysInfo_BuildCpuArchitecture();
extern __declspec(dllexport) struct miqt_string QSysInfo_CurrentCpuArchitecture();
extern __declspec(dllexport) struct miqt_string QSysInfo_BuildAbi();
extern __declspec(dllexport) struct miqt_string QSysInfo_KernelType();
extern __declspec(dllexport) struct miqt_string QSysInfo_KernelVersion();
extern __declspec(dllexport) struct miqt_string QSysInfo_ProductType();
extern __declspec(dllexport) struct miqt_string QSysInfo_ProductVersion();
extern __declspec(dllexport) struct miqt_string QSysInfo_PrettyProductName();
extern __declspec(dllexport) struct miqt_string QSysInfo_MachineHostName();
extern __declspec(dllexport) struct miqt_string QSysInfo_MachineUniqueId();
extern __declspec(dllexport) struct miqt_string QSysInfo_BootUniqueId();
extern __declspec(dllexport) void QSysInfo_Delete(QSysInfo* self, bool isSubclass);

} 
