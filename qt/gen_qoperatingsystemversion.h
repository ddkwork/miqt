#pragma once
#ifndef MIQT_QT_GEN_QOPERATINGSYSTEMVERSION_H
#define MIQT_QT_GEN_QOPERATINGSYSTEMVERSION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QOperatingSystemVersion QOperatingSystemVersion;
typedef struct QOperatingSystemVersionBase QOperatingSystemVersionBase;
typedef struct QOperatingSystemVersionUnexported QOperatingSystemVersionUnexported;
typedef struct QVersionNumber QVersionNumber;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QOperatingSystemVersionBase* QOperatingSystemVersionBase_new(OSType osType, int vmajor);
extern __declspec(dllexport) QOperatingSystemVersionBase* QOperatingSystemVersionBase_new2(QOperatingSystemVersionBase* param1);
extern __declspec(dllexport) QOperatingSystemVersionBase* QOperatingSystemVersionBase_new3(OSType osType, int vmajor, int vminor);
extern __declspec(dllexport) QOperatingSystemVersionBase* QOperatingSystemVersionBase_new4(OSType osType, int vmajor, int vminor, int vmicro);
extern __declspec(dllexport) QOperatingSystemVersionBase* QOperatingSystemVersionBase_Current();
extern __declspec(dllexport) struct miqt_string QOperatingSystemVersionBase_Name(QOperatingSystemVersionBase* osversion);
extern __declspec(dllexport) OSType QOperatingSystemVersionBase_CurrentType();
extern __declspec(dllexport) QVersionNumber* QOperatingSystemVersionBase_Version(const QOperatingSystemVersionBase* self);
extern __declspec(dllexport) int QOperatingSystemVersionBase_MajorVersion(const QOperatingSystemVersionBase* self);
extern __declspec(dllexport) int QOperatingSystemVersionBase_MinorVersion(const QOperatingSystemVersionBase* self);
extern __declspec(dllexport) int QOperatingSystemVersionBase_MicroVersion(const QOperatingSystemVersionBase* self);
extern __declspec(dllexport) int QOperatingSystemVersionBase_SegmentCount(const QOperatingSystemVersionBase* self);
extern __declspec(dllexport) OSType QOperatingSystemVersionBase_Type(const QOperatingSystemVersionBase* self);
extern __declspec(dllexport) struct miqt_string QOperatingSystemVersionBase_Name2(const QOperatingSystemVersionBase* self);
extern __declspec(dllexport) void QOperatingSystemVersionBase_Delete(QOperatingSystemVersionBase* self, bool isSubclass);

extern __declspec(dllexport) QOperatingSystemVersionUnexported* QOperatingSystemVersionUnexported_new(QOperatingSystemVersionBase* other);
extern __declspec(dllexport) QOperatingSystemVersionUnexported* QOperatingSystemVersionUnexported_new2();
extern __declspec(dllexport) QOperatingSystemVersionUnexported* QOperatingSystemVersionUnexported_new3(QOperatingSystemVersionUnexported* param1);
extern __declspec(dllexport) void QOperatingSystemVersionUnexported_virtbase(QOperatingSystemVersionUnexported* src, QOperatingSystemVersionBase** outptr_QOperatingSystemVersionBase);
extern __declspec(dllexport) void QOperatingSystemVersionUnexported_Delete(QOperatingSystemVersionUnexported* self, bool isSubclass);

extern __declspec(dllexport) QOperatingSystemVersion* QOperatingSystemVersion_new(QOperatingSystemVersionBase* osversion);
extern __declspec(dllexport) QOperatingSystemVersion* QOperatingSystemVersion_new2(OSType osType, int vmajor);
extern __declspec(dllexport) QOperatingSystemVersion* QOperatingSystemVersion_new3(QOperatingSystemVersion* param1);
extern __declspec(dllexport) QOperatingSystemVersion* QOperatingSystemVersion_new4(OSType osType, int vmajor, int vminor);
extern __declspec(dllexport) QOperatingSystemVersion* QOperatingSystemVersion_new5(OSType osType, int vmajor, int vminor, int vmicro);
extern __declspec(dllexport) void QOperatingSystemVersion_virtbase(QOperatingSystemVersion* src, QOperatingSystemVersionUnexported** outptr_QOperatingSystemVersionUnexported);
extern __declspec(dllexport) OSType QOperatingSystemVersion_CurrentType();
extern __declspec(dllexport) OSType QOperatingSystemVersion_Type(const QOperatingSystemVersion* self);
extern __declspec(dllexport) void QOperatingSystemVersion_Delete(QOperatingSystemVersion* self, bool isSubclass);

} 
