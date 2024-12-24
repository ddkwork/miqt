#pragma once
#ifndef MIQT_QT_GEN_QLOGGING_H
#define MIQT_QT_GEN_QLOGGING_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDebug QDebug;
typedef struct QLoggingCategory QLoggingCategory;
typedef struct QMessageLogContext QMessageLogContext;
typedef struct QMessageLogger QMessageLogger;
typedef struct QNoDebug QNoDebug;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QMessageLogContext* QMessageLogContext_new();
extern __declspec(dllexport) QMessageLogContext* QMessageLogContext_new2(const char* fileName, int lineNumber, const char* functionName, const char* categoryName);
extern __declspec(dllexport) void QMessageLogContext_Delete(QMessageLogContext* self, bool isSubclass);

extern __declspec(dllexport) QMessageLogger* QMessageLogger_new();
extern __declspec(dllexport) QMessageLogger* QMessageLogger_new2(const char* file, int line, const char* function);
extern __declspec(dllexport) QMessageLogger* QMessageLogger_new3(const char* file, int line, const char* function, const char* category);
extern __declspec(dllexport) void QMessageLogger_Debug(const QMessageLogger* self, const char* msg, ... );
extern __declspec(dllexport) void QMessageLogger_NoDebug(const QMessageLogger* self, const char* param1, ... );
extern __declspec(dllexport) void QMessageLogger_Info(const QMessageLogger* self, const char* msg, ... );
extern __declspec(dllexport) void QMessageLogger_Warning(const QMessageLogger* self, const char* msg, ... );
extern __declspec(dllexport) void QMessageLogger_Critical(const QMessageLogger* self, const char* msg, ... );
extern __declspec(dllexport) void QMessageLogger_Fatal(const QMessageLogger* self, const char* msg, ... );
extern __declspec(dllexport) void QMessageLogger_Debug2(const QMessageLogger* self, QLoggingCategory* cat, const char* msg, ... );
extern __declspec(dllexport) void QMessageLogger_Debug3(const QMessageLogger* self, CategoryFunction catFunc, const char* msg, ... );
extern __declspec(dllexport) void QMessageLogger_Info2(const QMessageLogger* self, QLoggingCategory* cat, const char* msg, ... );
extern __declspec(dllexport) void QMessageLogger_Info3(const QMessageLogger* self, CategoryFunction catFunc, const char* msg, ... );
extern __declspec(dllexport) void QMessageLogger_Warning2(const QMessageLogger* self, QLoggingCategory* cat, const char* msg, ... );
extern __declspec(dllexport) void QMessageLogger_Warning3(const QMessageLogger* self, CategoryFunction catFunc, const char* msg, ... );
extern __declspec(dllexport) void QMessageLogger_Critical2(const QMessageLogger* self, QLoggingCategory* cat, const char* msg, ... );
extern __declspec(dllexport) void QMessageLogger_Critical3(const QMessageLogger* self, CategoryFunction catFunc, const char* msg, ... );
extern __declspec(dllexport) void QMessageLogger_Fatal2(const QMessageLogger* self, QLoggingCategory* cat, const char* msg, ... );
extern __declspec(dllexport) void QMessageLogger_Fatal3(const QMessageLogger* self, CategoryFunction catFunc, const char* msg, ... );
extern __declspec(dllexport) QDebug* QMessageLogger_Debug4(const QMessageLogger* self);
extern __declspec(dllexport) QDebug* QMessageLogger_DebugWithCat(const QMessageLogger* self, QLoggingCategory* cat);
extern __declspec(dllexport) QDebug* QMessageLogger_DebugWithCatFunc(const QMessageLogger* self, CategoryFunction catFunc);
extern __declspec(dllexport) QDebug* QMessageLogger_Info4(const QMessageLogger* self);
extern __declspec(dllexport) QDebug* QMessageLogger_InfoWithCat(const QMessageLogger* self, QLoggingCategory* cat);
extern __declspec(dllexport) QDebug* QMessageLogger_InfoWithCatFunc(const QMessageLogger* self, CategoryFunction catFunc);
extern __declspec(dllexport) QDebug* QMessageLogger_Warning4(const QMessageLogger* self);
extern __declspec(dllexport) QDebug* QMessageLogger_WarningWithCat(const QMessageLogger* self, QLoggingCategory* cat);
extern __declspec(dllexport) QDebug* QMessageLogger_WarningWithCatFunc(const QMessageLogger* self, CategoryFunction catFunc);
extern __declspec(dllexport) QDebug* QMessageLogger_Critical4(const QMessageLogger* self);
extern __declspec(dllexport) QDebug* QMessageLogger_CriticalWithCat(const QMessageLogger* self, QLoggingCategory* cat);
extern __declspec(dllexport) QDebug* QMessageLogger_CriticalWithCatFunc(const QMessageLogger* self, CategoryFunction catFunc);
extern __declspec(dllexport) QDebug* QMessageLogger_Fatal4(const QMessageLogger* self);
extern __declspec(dllexport) QDebug* QMessageLogger_FatalWithCat(const QMessageLogger* self, QLoggingCategory* cat);
extern __declspec(dllexport) QDebug* QMessageLogger_FatalWithCatFunc(const QMessageLogger* self, CategoryFunction catFunc);
extern __declspec(dllexport) QNoDebug* QMessageLogger_NoDebug2(const QMessageLogger* self);
extern __declspec(dllexport) void QMessageLogger_Delete(QMessageLogger* self, bool isSubclass);

} 
