#pragma once
#ifndef MIQT_QT_GEN_QLIBRARYINFO_H
#define MIQT_QT_GEN_QLIBRARYINFO_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QLibraryInfo QLibraryInfo;
typedef struct QVersionNumber QVersionNumber;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
const char* QLibraryInfo_Build();
extern __declspec(dllexport) 
bool QLibraryInfo_IsDebugBuild();
extern __declspec(dllexport) 
bool QLibraryInfo_IsSharedBuild();
extern __declspec(dllexport) 
QVersionNumber* QLibraryInfo_Version();
extern __declspec(dllexport) 
struct miqt_string QLibraryInfo_Path(LibraryPath p);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QLibraryInfo_Paths(LibraryPath p);
extern __declspec(dllexport) 
struct miqt_string QLibraryInfo_Location(LibraryLocation location);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QLibraryInfo_PlatformPluginArguments(struct miqt_string platformName);
extern __declspec(dllexport) 
void QLibraryInfo_Delete(QLibraryInfo* self, bool isSubclass);

}
