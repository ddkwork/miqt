#pragma once
#ifndef MIQT_QT_GEN_QPROCESS_H
#define MIQT_QT_GEN_QPROCESS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QIODevice;
class QIODeviceBase;
class QMetaObject;
class QObject;
class QProcess;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QProcess__CreateProcessArguments)
typedef QProcess::CreateProcessArguments QProcess__CreateProcessArguments;
#else
class QProcess__CreateProcessArguments;
#endif
class QProcessEnvironment;
class _GUID;
class type_info;
#else
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QProcess QProcess;
typedef struct QProcess__CreateProcessArguments QProcess__CreateProcessArguments;
typedef struct QProcessEnvironment QProcessEnvironment;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QProcessEnvironment* QProcessEnvironment_new();
extern __declspec(dllexport) QProcessEnvironment* QProcessEnvironment_new2(Initialization param1);
extern __declspec(dllexport) QProcessEnvironment* QProcessEnvironment_new3(QProcessEnvironment* other);
extern __declspec(dllexport) void QProcessEnvironment_OperatorAssign(QProcessEnvironment* self, QProcessEnvironment* other);
extern __declspec(dllexport) void QProcessEnvironment_Swap(QProcessEnvironment* self, QProcessEnvironment* other);
extern __declspec(dllexport) bool QProcessEnvironment_IsEmpty(const QProcessEnvironment* self);
extern __declspec(dllexport) bool QProcessEnvironment_InheritsFromParent(const QProcessEnvironment* self);
extern __declspec(dllexport) void QProcessEnvironment_Clear(QProcessEnvironment* self);
extern __declspec(dllexport) bool QProcessEnvironment_Contains(const QProcessEnvironment* self, struct miqt_string name);
extern __declspec(dllexport) void QProcessEnvironment_Insert(QProcessEnvironment* self, struct miqt_string name, struct miqt_string value);
extern __declspec(dllexport) void QProcessEnvironment_Remove(QProcessEnvironment* self, struct miqt_string name);
extern __declspec(dllexport) struct miqt_string QProcessEnvironment_Value(const QProcessEnvironment* self, struct miqt_string name);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QProcessEnvironment_ToStringList(const QProcessEnvironment* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QProcessEnvironment_Keys(const QProcessEnvironment* self);
extern __declspec(dllexport) void QProcessEnvironment_InsertWithQProcessEnvironment(QProcessEnvironment* self, QProcessEnvironment* e);
extern __declspec(dllexport) QProcessEnvironment* QProcessEnvironment_SystemEnvironment();
extern __declspec(dllexport) struct miqt_string QProcessEnvironment_Value2(const QProcessEnvironment* self, struct miqt_string name, struct miqt_string defaultValue);
extern __declspec(dllexport) void QProcessEnvironment_Delete(QProcessEnvironment* self, bool isSubclass);

extern __declspec(dllexport) QProcess* QProcess_new();
extern __declspec(dllexport) QProcess* QProcess_new2(QObject* parent);
extern __declspec(dllexport) void QProcess_virtbase(QProcess* src, QIODevice** outptr_QIODevice);
extern __declspec(dllexport) QMetaObject* QProcess_MetaObject(const QProcess* self);
extern __declspec(dllexport) void* QProcess_Metacast(QProcess* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QProcess_Tr(const char* s);
extern __declspec(dllexport) void QProcess_Start(QProcess* self, struct miqt_string program);
extern __declspec(dllexport) void QProcess_Start2(QProcess* self);
extern __declspec(dllexport) void QProcess_StartCommand(QProcess* self, struct miqt_string command);
extern __declspec(dllexport) bool QProcess_StartDetached(QProcess* self);
extern __declspec(dllexport) bool QProcess_Open(QProcess* self, OpenMode mode);
extern __declspec(dllexport) struct miqt_string QProcess_Program(const QProcess* self);
extern __declspec(dllexport) void QProcess_SetProgram(QProcess* self, struct miqt_string program);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QProcess_Arguments(const QProcess* self);
extern __declspec(dllexport) void QProcess_SetArguments(QProcess* self, struct miqt_array /* of struct miqt_string */  arguments);
extern __declspec(dllexport) ProcessChannelMode QProcess_ProcessChannelMode(const QProcess* self);
extern __declspec(dllexport) void QProcess_SetProcessChannelMode(QProcess* self, ProcessChannelMode mode);
extern __declspec(dllexport) InputChannelMode QProcess_InputChannelMode(const QProcess* self);
extern __declspec(dllexport) void QProcess_SetInputChannelMode(QProcess* self, InputChannelMode mode);
extern __declspec(dllexport) ProcessChannel QProcess_ReadChannel(const QProcess* self);
extern __declspec(dllexport) void QProcess_SetReadChannel(QProcess* self, ProcessChannel channel);
extern __declspec(dllexport) void QProcess_CloseReadChannel(QProcess* self, ProcessChannel channel);
extern __declspec(dllexport) void QProcess_CloseWriteChannel(QProcess* self);
extern __declspec(dllexport) void QProcess_SetStandardInputFile(QProcess* self, struct miqt_string fileName);
extern __declspec(dllexport) void QProcess_SetStandardOutputFile(QProcess* self, struct miqt_string fileName);
extern __declspec(dllexport) void QProcess_SetStandardErrorFile(QProcess* self, struct miqt_string fileName);
extern __declspec(dllexport) void QProcess_SetStandardOutputProcess(QProcess* self, QProcess* destination);
extern __declspec(dllexport) struct miqt_string QProcess_NativeArguments(const QProcess* self);
extern __declspec(dllexport) void QProcess_SetNativeArguments(QProcess* self, struct miqt_string arguments);
extern __declspec(dllexport) CreateProcessArgumentModifier QProcess_CreateProcessArgumentsModifier(const QProcess* self);
extern __declspec(dllexport) void QProcess_SetCreateProcessArgumentsModifier(QProcess* self, CreateProcessArgumentModifier modifier);
extern __declspec(dllexport) struct miqt_string QProcess_WorkingDirectory(const QProcess* self);
extern __declspec(dllexport) void QProcess_SetWorkingDirectory(QProcess* self, struct miqt_string dir);
extern __declspec(dllexport) void QProcess_SetEnvironment(QProcess* self, struct miqt_array /* of struct miqt_string */  environment);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QProcess_Environment(const QProcess* self);
extern __declspec(dllexport) void QProcess_SetProcessEnvironment(QProcess* self, QProcessEnvironment* environment);
extern __declspec(dllexport) QProcessEnvironment* QProcess_ProcessEnvironment(const QProcess* self);
extern __declspec(dllexport) int QProcess_Error(const QProcess* self);
extern __declspec(dllexport) int QProcess_State(const QProcess* self);
extern __declspec(dllexport) long long QProcess_ProcessId(const QProcess* self);
extern __declspec(dllexport) bool QProcess_WaitForStarted(QProcess* self);
extern __declspec(dllexport) bool QProcess_WaitForReadyRead(QProcess* self, int msecs);
extern __declspec(dllexport) bool QProcess_WaitForBytesWritten(QProcess* self, int msecs);
extern __declspec(dllexport) bool QProcess_WaitForFinished(QProcess* self);
extern __declspec(dllexport) struct miqt_string QProcess_ReadAllStandardOutput(QProcess* self);
extern __declspec(dllexport) struct miqt_string QProcess_ReadAllStandardError(QProcess* self);
extern __declspec(dllexport) int QProcess_ExitCode(const QProcess* self);
extern __declspec(dllexport) int QProcess_ExitStatus(const QProcess* self);
extern __declspec(dllexport) long long QProcess_BytesToWrite(const QProcess* self);
extern __declspec(dllexport) bool QProcess_IsSequential(const QProcess* self);
extern __declspec(dllexport) void QProcess_Close(QProcess* self);
extern __declspec(dllexport) int QProcess_Execute(struct miqt_string program);
extern __declspec(dllexport) bool QProcess_StartDetachedWithProgram(struct miqt_string program);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QProcess_SystemEnvironment();
extern __declspec(dllexport) struct miqt_string QProcess_NullDevice();
extern __declspec(dllexport) void QProcess_Terminate(QProcess* self);
extern __declspec(dllexport) void QProcess_Kill(QProcess* self);
extern __declspec(dllexport) void QProcess_Finished(QProcess* self, int exitCode);
void QProcess_connect_Finished(QProcess* self, intptr_t slot);
extern __declspec(dllexport) void QProcess_ErrorOccurred(QProcess* self, int error);
void QProcess_connect_ErrorOccurred(QProcess* self, intptr_t slot);
extern __declspec(dllexport) long long QProcess_ReadData(QProcess* self, char* data, long long maxlen);
extern __declspec(dllexport) long long QProcess_WriteData(QProcess* self, const char* data, long long lenVal);
extern __declspec(dllexport) struct miqt_string QProcess_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QProcess_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QProcess_Start22(QProcess* self, struct miqt_string program, struct miqt_array /* of struct miqt_string */  arguments);
extern __declspec(dllexport) void QProcess_Start3(QProcess* self, struct miqt_string program, struct miqt_array /* of struct miqt_string */  arguments, OpenMode mode);
extern __declspec(dllexport) void QProcess_Start1(QProcess* self, OpenMode mode);
extern __declspec(dllexport) void QProcess_StartCommand2(QProcess* self, struct miqt_string command, OpenMode mode);
extern __declspec(dllexport) bool QProcess_StartDetached1(QProcess* self, long long* pid);
extern __declspec(dllexport) void QProcess_SetStandardOutputFile2(QProcess* self, struct miqt_string fileName, OpenMode mode);
extern __declspec(dllexport) void QProcess_SetStandardErrorFile2(QProcess* self, struct miqt_string fileName, OpenMode mode);
extern __declspec(dllexport) bool QProcess_WaitForStarted1(QProcess* self, int msecs);
extern __declspec(dllexport) bool QProcess_WaitForFinished1(QProcess* self, int msecs);
extern __declspec(dllexport) int QProcess_Execute2(struct miqt_string program, struct miqt_array /* of struct miqt_string */  arguments);
extern __declspec(dllexport) bool QProcess_StartDetached2(struct miqt_string program, struct miqt_array /* of struct miqt_string */  arguments);
extern __declspec(dllexport) bool QProcess_StartDetached3(struct miqt_string program, struct miqt_array /* of struct miqt_string */  arguments, struct miqt_string workingDirectory);
extern __declspec(dllexport) bool QProcess_StartDetached4(struct miqt_string program, struct miqt_array /* of struct miqt_string */  arguments, struct miqt_string workingDirectory, long long* pid);
extern __declspec(dllexport) void QProcess_Finished2(QProcess* self, int exitCode, int exitStatus);
void QProcess_connect_Finished2(QProcess* self, intptr_t slot);
extern __declspec(dllexport) void QProcess_override_virtual_Open(void* self, intptr_t slot);
bool QProcess_virtualbase_Open(void* self, OpenMode mode);
extern __declspec(dllexport) void QProcess_override_virtual_WaitForReadyRead(void* self, intptr_t slot);
bool QProcess_virtualbase_WaitForReadyRead(void* self, int msecs);
extern __declspec(dllexport) void QProcess_override_virtual_WaitForBytesWritten(void* self, intptr_t slot);
bool QProcess_virtualbase_WaitForBytesWritten(void* self, int msecs);
extern __declspec(dllexport) void QProcess_override_virtual_BytesToWrite(void* self, intptr_t slot);
long long QProcess_virtualbase_BytesToWrite(const void* self);
extern __declspec(dllexport) void QProcess_override_virtual_IsSequential(void* self, intptr_t slot);
bool QProcess_virtualbase_IsSequential(const void* self);
extern __declspec(dllexport) void QProcess_override_virtual_Close(void* self, intptr_t slot);
void QProcess_virtualbase_Close(void* self);
extern __declspec(dllexport) void QProcess_override_virtual_ReadData(void* self, intptr_t slot);
long long QProcess_virtualbase_ReadData(void* self, char* data, long long maxlen);
extern __declspec(dllexport) void QProcess_override_virtual_WriteData(void* self, intptr_t slot);
long long QProcess_virtualbase_WriteData(void* self, const char* data, long long lenVal);
extern __declspec(dllexport) void QProcess_override_virtual_Pos(void* self, intptr_t slot);
long long QProcess_virtualbase_Pos(const void* self);
extern __declspec(dllexport) void QProcess_override_virtual_Size(void* self, intptr_t slot);
long long QProcess_virtualbase_Size(const void* self);
extern __declspec(dllexport) void QProcess_override_virtual_Seek(void* self, intptr_t slot);
bool QProcess_virtualbase_Seek(void* self, long long pos);
extern __declspec(dllexport) void QProcess_override_virtual_AtEnd(void* self, intptr_t slot);
bool QProcess_virtualbase_AtEnd(const void* self);
extern __declspec(dllexport) void QProcess_override_virtual_Reset(void* self, intptr_t slot);
bool QProcess_virtualbase_Reset(void* self);
extern __declspec(dllexport) void QProcess_override_virtual_BytesAvailable(void* self, intptr_t slot);
long long QProcess_virtualbase_BytesAvailable(const void* self);
extern __declspec(dllexport) void QProcess_override_virtual_CanReadLine(void* self, intptr_t slot);
bool QProcess_virtualbase_CanReadLine(const void* self);
extern __declspec(dllexport) void QProcess_override_virtual_ReadLineData(void* self, intptr_t slot);
long long QProcess_virtualbase_ReadLineData(void* self, char* data, long long maxlen);
extern __declspec(dllexport) void QProcess_override_virtual_SkipData(void* self, intptr_t slot);
long long QProcess_virtualbase_SkipData(void* self, long long maxSize);
extern __declspec(dllexport) void QProcess_Delete(QProcess* self, bool isSubclass);

extern __declspec(dllexport) void QProcess__CreateProcessArguments_Delete(QProcess__CreateProcessArguments* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
