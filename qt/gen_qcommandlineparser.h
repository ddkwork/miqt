#pragma once
#ifndef MIQT_QT_GEN_QCOMMANDLINEPARSER_H
#define MIQT_QT_GEN_QCOMMANDLINEPARSER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QCommandLineOption QCommandLineOption;
typedef struct QCommandLineParser QCommandLineParser;
typedef struct QCoreApplication QCoreApplication;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QCommandLineParser* QCommandLineParser_new();
extern __declspec(dllexport) 
struct miqt_string QCommandLineParser_Tr(const char* sourceText);
extern __declspec(dllexport) 
void QCommandLineParser_SetSingleDashWordOptionMode(QCommandLineParser* self, SingleDashWordOptionMode parsingMode);
extern __declspec(dllexport) 
void QCommandLineParser_SetOptionsAfterPositionalArgumentsMode(QCommandLineParser* self, OptionsAfterPositionalArgumentsMode mode);
extern __declspec(dllexport) 
bool QCommandLineParser_AddOption(QCommandLineParser* self, QCommandLineOption* commandLineOption);
extern __declspec(dllexport) 
bool QCommandLineParser_AddOptions(QCommandLineParser* self, struct miqt_array /* of QCommandLineOption* */  options);
extern __declspec(dllexport) 
QCommandLineOption* QCommandLineParser_AddVersionOption(QCommandLineParser* self);
extern __declspec(dllexport) 
QCommandLineOption* QCommandLineParser_AddHelpOption(QCommandLineParser* self);
extern __declspec(dllexport) 
void QCommandLineParser_SetApplicationDescription(QCommandLineParser* self, struct miqt_string description);
extern __declspec(dllexport) 
struct miqt_string QCommandLineParser_ApplicationDescription(const QCommandLineParser* self);
extern __declspec(dllexport) 
void QCommandLineParser_AddPositionalArgument(QCommandLineParser* self, struct miqt_string name, struct miqt_string description);
extern __declspec(dllexport) 
void QCommandLineParser_ClearPositionalArguments(QCommandLineParser* self);
extern __declspec(dllexport) 
void QCommandLineParser_Process(QCommandLineParser* self, struct miqt_array /* of struct miqt_string */  arguments);
extern __declspec(dllexport) 
void QCommandLineParser_ProcessWithApp(QCommandLineParser* self, QCoreApplication* app);
extern __declspec(dllexport) 
bool QCommandLineParser_Parse(QCommandLineParser* self, struct miqt_array /* of struct miqt_string */  arguments);
extern __declspec(dllexport) 
struct miqt_string QCommandLineParser_ErrorText(const QCommandLineParser* self);
extern __declspec(dllexport) 
bool QCommandLineParser_IsSet(const QCommandLineParser* self, struct miqt_string name);
extern __declspec(dllexport) 
struct miqt_string QCommandLineParser_Value(const QCommandLineParser* self, struct miqt_string name);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QCommandLineParser_Values(const QCommandLineParser* self, struct miqt_string name);
extern __declspec(dllexport) 
bool QCommandLineParser_IsSetWithOption(const QCommandLineParser* self, QCommandLineOption* option);
extern __declspec(dllexport) 
struct miqt_string QCommandLineParser_ValueWithOption(const QCommandLineParser* self, QCommandLineOption* option);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QCommandLineParser_ValuesWithOption(const QCommandLineParser* self, QCommandLineOption* option);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QCommandLineParser_PositionalArguments(const QCommandLineParser* self);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QCommandLineParser_OptionNames(const QCommandLineParser* self);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QCommandLineParser_UnknownOptionNames(const QCommandLineParser* self);
extern __declspec(dllexport) 
void QCommandLineParser_ShowVersion(QCommandLineParser* self);
extern __declspec(dllexport) 
void QCommandLineParser_ShowHelp(QCommandLineParser* self);
extern __declspec(dllexport) 
struct miqt_string QCommandLineParser_HelpText(const QCommandLineParser* self);
extern __declspec(dllexport) 
void QCommandLineParser_ShowMessageAndExit(struct miqt_string message, MessageType typeVal);
extern __declspec(dllexport) 
struct miqt_string QCommandLineParser_Tr2(const char* sourceText, const char* disambiguation);
extern __declspec(dllexport) 
struct miqt_string QCommandLineParser_Tr3(const char* sourceText, const char* disambiguation, int n);
extern __declspec(dllexport) 
void QCommandLineParser_AddPositionalArgument3(QCommandLineParser* self, struct miqt_string name, struct miqt_string description, struct miqt_string syntax);
extern __declspec(dllexport) 
void QCommandLineParser_ShowHelp1(QCommandLineParser* self, int exitCode);
extern __declspec(dllexport) 
void QCommandLineParser_ShowMessageAndExit3(struct miqt_string message, MessageType typeVal, int exitCode);
extern __declspec(dllexport) 
void QCommandLineParser_Delete(QCommandLineParser* self, bool isSubclass);

}
