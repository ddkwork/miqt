#pragma once
#ifndef MIQT_QT_GEN_QTYPEREVISION_H
#define MIQT_QT_GEN_QTYPEREVISION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QTypeRevision QTypeRevision;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTypeRevision* QTypeRevision_new();
extern __declspec(dllexport) QTypeRevision* QTypeRevision_new2(QTypeRevision* param1);
extern __declspec(dllexport) QTypeRevision* QTypeRevision_Zero();
extern __declspec(dllexport) bool QTypeRevision_HasMajorVersion(const QTypeRevision* self);
extern __declspec(dllexport) unsigned char QTypeRevision_MajorVersion(const QTypeRevision* self);
extern __declspec(dllexport) bool QTypeRevision_HasMinorVersion(const QTypeRevision* self);
extern __declspec(dllexport) unsigned char QTypeRevision_MinorVersion(const QTypeRevision* self);
extern __declspec(dllexport) bool QTypeRevision_IsValid(const QTypeRevision* self);
extern __declspec(dllexport) void QTypeRevision_Delete(QTypeRevision* self, bool isSubclass);

} 
