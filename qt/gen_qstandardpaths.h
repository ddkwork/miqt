#pragma once
#ifndef MIQT_QT_GEN_QSTANDARDPATHS_H
#define MIQT_QT_GEN_QSTANDARDPATHS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QStandardPaths QStandardPaths;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
struct miqt_string QStandardPaths_WritableLocation(StandardLocation typeVal);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QStandardPaths_StandardLocations(StandardLocation typeVal);
extern __declspec(dllexport) 
struct miqt_string QStandardPaths_Locate(StandardLocation typeVal, struct miqt_string fileName);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QStandardPaths_LocateAll(StandardLocation typeVal, struct miqt_string fileName);
extern __declspec(dllexport) 
struct miqt_string QStandardPaths_DisplayName(StandardLocation typeVal);
extern __declspec(dllexport) 
struct miqt_string QStandardPaths_FindExecutable(struct miqt_string executableName);
extern __declspec(dllexport) 
void QStandardPaths_SetTestModeEnabled(bool testMode);
extern __declspec(dllexport) 
bool QStandardPaths_IsTestModeEnabled();
extern __declspec(dllexport) 
struct miqt_string QStandardPaths_Locate3(StandardLocation typeVal, struct miqt_string fileName, LocateOptions options);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QStandardPaths_LocateAll3(StandardLocation typeVal, struct miqt_string fileName, LocateOptions options);
extern __declspec(dllexport) 
struct miqt_string QStandardPaths_FindExecutable2(struct miqt_string executableName, struct miqt_array /* of struct miqt_string */  paths);

}
