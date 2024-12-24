#pragma once
#ifndef MIQT_QT_GEN_QINPUTDEVICE_H
#define MIQT_QT_GEN_QINPUTDEVICE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChildEvent;
class QEvent;
class QInputDevice;
class QMetaMethod;
class QMetaObject;
class QObject;
class QRect;
class QTimerEvent;
class _GUID;
class type_info;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QInputDevice QInputDevice;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QRect QRect;
typedef struct QTimerEvent QTimerEvent;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QInputDevice* QInputDevice_new();
extern __declspec(dllexport) QInputDevice* QInputDevice_new2(struct miqt_string name, long long systemId, DeviceType typeVal);
extern __declspec(dllexport) QInputDevice* QInputDevice_new3(QObject* parent);
extern __declspec(dllexport) QInputDevice* QInputDevice_new4(struct miqt_string name, long long systemId, DeviceType typeVal, struct miqt_string seatName);
extern __declspec(dllexport) QInputDevice* QInputDevice_new5(struct miqt_string name, long long systemId, DeviceType typeVal, struct miqt_string seatName, QObject* parent);
extern __declspec(dllexport) void QInputDevice_virtbase(QInputDevice* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QInputDevice_MetaObject(const QInputDevice* self);
extern __declspec(dllexport) void* QInputDevice_Metacast(QInputDevice* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QInputDevice_Tr(const char* s);
extern __declspec(dllexport) struct miqt_string QInputDevice_Name(const QInputDevice* self);
extern __declspec(dllexport) DeviceType QInputDevice_Type(const QInputDevice* self);
extern __declspec(dllexport) Capabilities QInputDevice_Capabilities(const QInputDevice* self);
extern __declspec(dllexport) bool QInputDevice_HasCapability(const QInputDevice* self, Capability cap);
extern __declspec(dllexport) long long QInputDevice_SystemId(const QInputDevice* self);
extern __declspec(dllexport) struct miqt_string QInputDevice_SeatName(const QInputDevice* self);
extern __declspec(dllexport) QRect* QInputDevice_AvailableVirtualGeometry(const QInputDevice* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QInputDevice_SeatNames();
extern __declspec(dllexport) struct miqt_array /* of QInputDevice* */  QInputDevice_Devices();
extern __declspec(dllexport) QInputDevice* QInputDevice_PrimaryKeyboard();
extern __declspec(dllexport) bool QInputDevice_OperatorEqual(const QInputDevice* self, QInputDevice* other);
extern __declspec(dllexport) void QInputDevice_AvailableVirtualGeometryChanged(QInputDevice* self, QRect* area);
void QInputDevice_connect_AvailableVirtualGeometryChanged(QInputDevice* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QInputDevice_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QInputDevice_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) QInputDevice* QInputDevice_PrimaryKeyboard1(struct miqt_string seatName);
extern __declspec(dllexport) void QInputDevice_override_virtual_Event(void* self, intptr_t slot);
bool QInputDevice_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QInputDevice_override_virtual_EventFilter(void* self, intptr_t slot);
bool QInputDevice_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QInputDevice_override_virtual_TimerEvent(void* self, intptr_t slot);
void QInputDevice_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QInputDevice_override_virtual_ChildEvent(void* self, intptr_t slot);
void QInputDevice_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QInputDevice_override_virtual_CustomEvent(void* self, intptr_t slot);
void QInputDevice_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QInputDevice_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QInputDevice_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QInputDevice_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QInputDevice_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QInputDevice_Delete(QInputDevice* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
