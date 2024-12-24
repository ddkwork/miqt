#pragma once
#ifndef MIQT_QT_GEN_QSETTINGS_H
#define MIQT_QT_GEN_QSETTINGS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAnyStringView;
class QChildEvent;
class QEvent;
class QMetaMethod;
class QMetaObject;
class QObject;
class QSettings;
class QTimerEvent;
class QVariant;
class _GUID;
class type_info;
#else
typedef struct QAnyStringView QAnyStringView;
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QSettings QSettings;
typedef struct QTimerEvent QTimerEvent;
typedef struct QVariant QVariant;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QSettings* QSettings_new(struct miqt_string organization);
extern __declspec(dllexport) QSettings* QSettings_new2(Scope scope, struct miqt_string organization);
extern __declspec(dllexport) QSettings* QSettings_new3(Format format, Scope scope, struct miqt_string organization);
extern __declspec(dllexport) QSettings* QSettings_new4(struct miqt_string fileName, Format format);
extern __declspec(dllexport) QSettings* QSettings_new5();
extern __declspec(dllexport) QSettings* QSettings_new6(Scope scope);
extern __declspec(dllexport) QSettings* QSettings_new7(struct miqt_string organization, struct miqt_string application);
extern __declspec(dllexport) QSettings* QSettings_new8(struct miqt_string organization, struct miqt_string application, QObject* parent);
extern __declspec(dllexport) QSettings* QSettings_new9(Scope scope, struct miqt_string organization, struct miqt_string application);
extern __declspec(dllexport) QSettings* QSettings_new10(Scope scope, struct miqt_string organization, struct miqt_string application, QObject* parent);
extern __declspec(dllexport) QSettings* QSettings_new11(Format format, Scope scope, struct miqt_string organization, struct miqt_string application);
extern __declspec(dllexport) QSettings* QSettings_new12(Format format, Scope scope, struct miqt_string organization, struct miqt_string application, QObject* parent);
extern __declspec(dllexport) QSettings* QSettings_new13(struct miqt_string fileName, Format format, QObject* parent);
extern __declspec(dllexport) QSettings* QSettings_new14(QObject* parent);
extern __declspec(dllexport) QSettings* QSettings_new15(Scope scope, QObject* parent);
extern __declspec(dllexport) void QSettings_virtbase(QSettings* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QSettings_MetaObject(const QSettings* self);
extern __declspec(dllexport) void* QSettings_Metacast(QSettings* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QSettings_Tr(const char* s);
extern __declspec(dllexport) void QSettings_Clear(QSettings* self);
extern __declspec(dllexport) void QSettings_Sync(QSettings* self);
extern __declspec(dllexport) Status QSettings_Status(const QSettings* self);
extern __declspec(dllexport) bool QSettings_IsAtomicSyncRequired(const QSettings* self);
extern __declspec(dllexport) void QSettings_SetAtomicSyncRequired(QSettings* self, bool enable);
extern __declspec(dllexport) void QSettings_BeginGroup(QSettings* self, QAnyStringView* prefix);
extern __declspec(dllexport) void QSettings_EndGroup(QSettings* self);
extern __declspec(dllexport) struct miqt_string QSettings_Group(const QSettings* self);
extern __declspec(dllexport) int QSettings_BeginReadArray(QSettings* self, QAnyStringView* prefix);
extern __declspec(dllexport) void QSettings_BeginWriteArray(QSettings* self, QAnyStringView* prefix);
extern __declspec(dllexport) void QSettings_EndArray(QSettings* self);
extern __declspec(dllexport) void QSettings_SetArrayIndex(QSettings* self, int i);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QSettings_AllKeys(const QSettings* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QSettings_ChildKeys(const QSettings* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QSettings_ChildGroups(const QSettings* self);
extern __declspec(dllexport) bool QSettings_IsWritable(const QSettings* self);
extern __declspec(dllexport) void QSettings_SetValue(QSettings* self, QAnyStringView* key, QVariant* value);
extern __declspec(dllexport) QVariant* QSettings_Value(const QSettings* self, QAnyStringView* key, QVariant* defaultValue);
extern __declspec(dllexport) QVariant* QSettings_ValueWithKey(const QSettings* self, QAnyStringView* key);
extern __declspec(dllexport) void QSettings_Remove(QSettings* self, QAnyStringView* key);
extern __declspec(dllexport) bool QSettings_Contains(const QSettings* self, QAnyStringView* key);
extern __declspec(dllexport) void QSettings_SetFallbacksEnabled(QSettings* self, bool b);
extern __declspec(dllexport) bool QSettings_FallbacksEnabled(const QSettings* self);
extern __declspec(dllexport) struct miqt_string QSettings_FileName(const QSettings* self);
extern __declspec(dllexport) Format QSettings_Format(const QSettings* self);
extern __declspec(dllexport) Scope QSettings_Scope(const QSettings* self);
extern __declspec(dllexport) struct miqt_string QSettings_OrganizationName(const QSettings* self);
extern __declspec(dllexport) struct miqt_string QSettings_ApplicationName(const QSettings* self);
extern __declspec(dllexport) void QSettings_SetDefaultFormat(Format format);
extern __declspec(dllexport) Format QSettings_DefaultFormat();
extern __declspec(dllexport) void QSettings_SetPath(Format format, Scope scope, struct miqt_string path);
extern __declspec(dllexport) Format QSettings_RegisterFormat(struct miqt_string extension, ReadFunc readFunc, WriteFunc writeFunc);
extern __declspec(dllexport) bool QSettings_Event(QSettings* self, QEvent* event);
extern __declspec(dllexport) struct miqt_string QSettings_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QSettings_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QSettings_BeginWriteArray2(QSettings* self, QAnyStringView* prefix, int size);
extern __declspec(dllexport) Format QSettings_RegisterFormat4(struct miqt_string extension, ReadFunc readFunc, WriteFunc writeFunc, int caseSensitivity);
extern __declspec(dllexport) void QSettings_override_virtual_Event(void* self, intptr_t slot);
bool QSettings_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QSettings_override_virtual_EventFilter(void* self, intptr_t slot);
bool QSettings_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QSettings_override_virtual_TimerEvent(void* self, intptr_t slot);
void QSettings_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QSettings_override_virtual_ChildEvent(void* self, intptr_t slot);
void QSettings_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QSettings_override_virtual_CustomEvent(void* self, intptr_t slot);
void QSettings_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QSettings_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QSettings_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QSettings_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QSettings_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QSettings_Delete(QSettings* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
