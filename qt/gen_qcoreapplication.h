#pragma once
#ifndef MIQT_QT_GEN_QCOREAPPLICATION_H
#define MIQT_QT_GEN_QCOREAPPLICATION_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractEventDispatcher QAbstractEventDispatcher;
typedef struct QAbstractNativeEventFilter QAbstractNativeEventFilter;
typedef struct QCoreApplication QCoreApplication;
typedef struct QDeadlineTimer QDeadlineTimer;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPermission QPermission;
typedef struct QTranslator QTranslator;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QCoreApplication* QCoreApplication_new(int* argc, char** argv);
extern __declspec(dllexport) 
QCoreApplication* QCoreApplication_new2(int* argc, char** argv, int param3);
extern __declspec(dllexport) 
void QCoreApplication_virtbase(QCoreApplication* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QCoreApplication_MetaObject(const QCoreApplication* self);
extern __declspec(dllexport) 
void* QCoreApplication_Metacast(QCoreApplication* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QCoreApplication_Tr(const char* s);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QCoreApplication_Arguments();
extern __declspec(dllexport) 
void QCoreApplication_SetAttribute(int attribute);
extern __declspec(dllexport) 
bool QCoreApplication_TestAttribute(int attribute);
extern __declspec(dllexport) 
void QCoreApplication_SetOrganizationDomain(struct miqt_string orgDomain);
extern __declspec(dllexport) 
struct miqt_string QCoreApplication_OrganizationDomain();
extern __declspec(dllexport) 
void QCoreApplication_SetOrganizationName(struct miqt_string orgName);
extern __declspec(dllexport) 
struct miqt_string QCoreApplication_OrganizationName();
extern __declspec(dllexport) 
void QCoreApplication_SetApplicationName(struct miqt_string application);
extern __declspec(dllexport) 
struct miqt_string QCoreApplication_ApplicationName();
extern __declspec(dllexport) 
void QCoreApplication_SetApplicationVersion(struct miqt_string version);
extern __declspec(dllexport) 
struct miqt_string QCoreApplication_ApplicationVersion();
extern __declspec(dllexport) 
void QCoreApplication_SetSetuidAllowed(bool allow);
extern __declspec(dllexport) 
bool QCoreApplication_IsSetuidAllowed();
extern __declspec(dllexport) 
QCoreApplication* QCoreApplication_Instance();
extern __declspec(dllexport) 
int QCoreApplication_Exec();
extern __declspec(dllexport) 
void QCoreApplication_ProcessEvents();
extern __declspec(dllexport) 
void QCoreApplication_ProcessEvents2(int flags, int maxtime);
extern __declspec(dllexport) 
void QCoreApplication_ProcessEvents3(int flags, QDeadlineTimer* deadline);
extern __declspec(dllexport) 
bool QCoreApplication_SendEvent(QObject* receiver, QEvent* event);
extern __declspec(dllexport) 
void QCoreApplication_PostEvent(QObject* receiver, QEvent* event);
extern __declspec(dllexport) 
void QCoreApplication_SendPostedEvents();
extern __declspec(dllexport) 
void QCoreApplication_RemovePostedEvents(QObject* receiver);
extern __declspec(dllexport) 
QAbstractEventDispatcher* QCoreApplication_EventDispatcher();
extern __declspec(dllexport) 
void QCoreApplication_SetEventDispatcher(QAbstractEventDispatcher* eventDispatcher);
extern __declspec(dllexport) 
bool QCoreApplication_Notify(QCoreApplication* self, QObject* param1, QEvent* param2);
extern __declspec(dllexport) 
bool QCoreApplication_StartingUp();
extern __declspec(dllexport) 
bool QCoreApplication_ClosingDown();
extern __declspec(dllexport) 
struct miqt_string QCoreApplication_ApplicationDirPath();
extern __declspec(dllexport) 
struct miqt_string QCoreApplication_ApplicationFilePath();
extern __declspec(dllexport) 
long long QCoreApplication_ApplicationPid();
extern __declspec(dllexport) 
int QCoreApplication_CheckPermission(QCoreApplication* self, QPermission* permission);
extern __declspec(dllexport) 
void QCoreApplication_SetLibraryPaths(struct miqt_array /* of struct miqt_string */  libraryPaths);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QCoreApplication_LibraryPaths();
extern __declspec(dllexport) 
void QCoreApplication_AddLibraryPath(struct miqt_string param1);
extern __declspec(dllexport) 
void QCoreApplication_RemoveLibraryPath(struct miqt_string param1);
extern __declspec(dllexport) 
bool QCoreApplication_InstallTranslator(QTranslator* messageFile);
extern __declspec(dllexport) 
bool QCoreApplication_RemoveTranslator(QTranslator* messageFile);
extern __declspec(dllexport) 
struct miqt_string QCoreApplication_Translate(const char* context, const char* key);
extern __declspec(dllexport) 
void QCoreApplication_InstallNativeEventFilter(QCoreApplication* self, QAbstractNativeEventFilter* filterObj);
void QCoreApplication_connect_InstallNativeEventFilter(QCoreApplication* self, intptr_t slot);
extern __declspec(dllexport) 
void QCoreApplication_RemoveNativeEventFilter(QCoreApplication* self, QAbstractNativeEventFilter* filterObj);
void QCoreApplication_connect_RemoveNativeEventFilter(QCoreApplication* self, intptr_t slot);
extern __declspec(dllexport) 
bool QCoreApplication_IsQuitLockEnabled();
extern __declspec(dllexport) 
void QCoreApplication_SetQuitLockEnabled(bool enabled);
extern __declspec(dllexport) 
void QCoreApplication_Quit();
extern __declspec(dllexport) 
void QCoreApplication_Exit();
extern __declspec(dllexport) 
void QCoreApplication_OrganizationNameChanged(QCoreApplication* self);
void QCoreApplication_connect_OrganizationNameChanged(QCoreApplication* self, intptr_t slot);
extern __declspec(dllexport) 
void QCoreApplication_OrganizationDomainChanged(QCoreApplication* self);
void QCoreApplication_connect_OrganizationDomainChanged(QCoreApplication* self, intptr_t slot);
extern __declspec(dllexport) 
void QCoreApplication_ApplicationNameChanged(QCoreApplication* self);
void QCoreApplication_connect_ApplicationNameChanged(QCoreApplication* self, intptr_t slot);
extern __declspec(dllexport) 
void QCoreApplication_ApplicationVersionChanged(QCoreApplication* self);
void QCoreApplication_connect_ApplicationVersionChanged(QCoreApplication* self, intptr_t slot);
extern __declspec(dllexport) 
bool QCoreApplication_Event(QCoreApplication* self, QEvent* param1);
extern __declspec(dllexport) 
struct miqt_string QCoreApplication_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QCoreApplication_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QCoreApplication_SetAttribute2(int attribute, bool on);
extern __declspec(dllexport) 
void QCoreApplication_ProcessEvents1(int flags);
extern __declspec(dllexport) 
void QCoreApplication_PostEvent3(QObject* receiver, QEvent* event, int priority);
extern __declspec(dllexport) 
void QCoreApplication_SendPostedEvents1(QObject* receiver);
extern __declspec(dllexport) 
void QCoreApplication_SendPostedEvents2(QObject* receiver, int event_type);
extern __declspec(dllexport) 
void QCoreApplication_RemovePostedEvents2(QObject* receiver, int eventType);
extern __declspec(dllexport) 
struct miqt_string QCoreApplication_Translate3(const char* context, const char* key, const char* disambiguation);
extern __declspec(dllexport) 
struct miqt_string QCoreApplication_Translate4(const char* context, const char* key, const char* disambiguation, int n);
extern __declspec(dllexport) 
void QCoreApplication_Exit1(int retcode);
extern __declspec(dllexport) 
void QCoreApplication_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QCoreApplication_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QCoreApplication_override_virtual_Metacast(void* self, intptr_t slot);
void* QCoreApplication_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QCoreApplication_Delete(QCoreApplication* self, bool isSubclass);

}
