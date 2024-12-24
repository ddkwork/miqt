#pragma once
#ifndef MIQT_QT_GEN_QLOGGINGCATEGORY_H
#define MIQT_QT_GEN_QLOGGINGCATEGORY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QLoggingCategory QLoggingCategory;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QLoggingCategory* QLoggingCategory_new(const char* category);
extern __declspec(dllexport) bool QLoggingCategory_IsDebugEnabled(const QLoggingCategory* self);
extern __declspec(dllexport) bool QLoggingCategory_IsInfoEnabled(const QLoggingCategory* self);
extern __declspec(dllexport) bool QLoggingCategory_IsWarningEnabled(const QLoggingCategory* self);
extern __declspec(dllexport) bool QLoggingCategory_IsCriticalEnabled(const QLoggingCategory* self);
extern __declspec(dllexport) const char* QLoggingCategory_CategoryName(const QLoggingCategory* self);
extern __declspec(dllexport) QLoggingCategory* QLoggingCategory_OperatorCall(QLoggingCategory* self);
extern __declspec(dllexport) QLoggingCategory* QLoggingCategory_OperatorCall2(const QLoggingCategory* self);
extern __declspec(dllexport) QLoggingCategory* QLoggingCategory_DefaultCategory();
extern __declspec(dllexport) CategoryFilter QLoggingCategory_InstallFilter(CategoryFilter param1);
extern __declspec(dllexport) void QLoggingCategory_SetFilterRules(struct miqt_string rules);
extern __declspec(dllexport) void QLoggingCategory_Delete(QLoggingCategory* self, bool isSubclass);

} 
