#pragma once
#ifndef MIQT_QT_GEN_QLOCKFILE_H
#define MIQT_QT_GEN_QLOCKFILE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QLockFile;
class _GUID;
class type_info;
#else
typedef struct QLockFile QLockFile;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QLockFile* QLockFile_new(struct miqt_string fileName);
extern __declspec(dllexport) struct miqt_string QLockFile_FileName(const QLockFile* self);
extern __declspec(dllexport) bool QLockFile_Lock(QLockFile* self);
extern __declspec(dllexport) bool QLockFile_TryLock(QLockFile* self, int timeout);
extern __declspec(dllexport) void QLockFile_Unlock(QLockFile* self);
extern __declspec(dllexport) void QLockFile_SetStaleLockTime(QLockFile* self, int staleLockTime);
extern __declspec(dllexport) int QLockFile_StaleLockTime(const QLockFile* self);
extern __declspec(dllexport) bool QLockFile_TryLock2(QLockFile* self);
extern __declspec(dllexport) bool QLockFile_IsLocked(const QLockFile* self);
extern __declspec(dllexport) bool QLockFile_RemoveStaleLockFile(QLockFile* self);
extern __declspec(dllexport) LockError QLockFile_Error(const QLockFile* self);
extern __declspec(dllexport) void QLockFile_Delete(QLockFile* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
