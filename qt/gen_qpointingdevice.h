#pragma once
#ifndef MIQT_QT_GEN_QPOINTINGDEVICE_H
#define MIQT_QT_GEN_QPOINTINGDEVICE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QEventPoint QEventPoint;
typedef struct QInputDevice QInputDevice;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPointerEvent QPointerEvent;
typedef struct QPointingDevice QPointingDevice;
typedef struct QPointingDeviceUniqueId QPointingDeviceUniqueId;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QPointingDeviceUniqueId* QPointingDeviceUniqueId_new();
extern __declspec(dllexport) 
QPointingDeviceUniqueId* QPointingDeviceUniqueId_new2(QPointingDeviceUniqueId* param1);
extern __declspec(dllexport) 
QPointingDeviceUniqueId* QPointingDeviceUniqueId_FromNumericId(long long id);
extern __declspec(dllexport) 
bool QPointingDeviceUniqueId_IsValid(const QPointingDeviceUniqueId* self);
extern __declspec(dllexport) 
long long QPointingDeviceUniqueId_NumericId(const QPointingDeviceUniqueId* self);
extern __declspec(dllexport) 
void QPointingDeviceUniqueId_Delete(QPointingDeviceUniqueId* self, bool isSubclass);

extern __declspec(dllexport) 
QPointingDevice* QPointingDevice_new();
extern __declspec(dllexport) 
QPointingDevice* QPointingDevice_new2(struct miqt_string name, long long systemId, int devType, PointerType pType, Capabilities caps, int maxPoints, int buttonCount);
extern __declspec(dllexport) 
QPointingDevice* QPointingDevice_new3(QObject* parent);
extern __declspec(dllexport) 
QPointingDevice* QPointingDevice_new4(struct miqt_string name, long long systemId, int devType, PointerType pType, Capabilities caps, int maxPoints, int buttonCount, struct miqt_string seatName);
extern __declspec(dllexport) 
QPointingDevice* QPointingDevice_new5(struct miqt_string name, long long systemId, int devType, PointerType pType, Capabilities caps, int maxPoints, int buttonCount, struct miqt_string seatName, QPointingDeviceUniqueId* uniqueId);
extern __declspec(dllexport) 
QPointingDevice* QPointingDevice_new6(struct miqt_string name, long long systemId, int devType, PointerType pType, Capabilities caps, int maxPoints, int buttonCount, struct miqt_string seatName, QPointingDeviceUniqueId* uniqueId, QObject* parent);
extern __declspec(dllexport) 
void QPointingDevice_virtbase(QPointingDevice* src
, QInputDevice** outptr_QInputDevice
);
extern __declspec(dllexport) 
QMetaObject* QPointingDevice_MetaObject(const QPointingDevice* self);
extern __declspec(dllexport) 
void* QPointingDevice_Metacast(QPointingDevice* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QPointingDevice_Tr(const char* s);
extern __declspec(dllexport) 
void QPointingDevice_SetType(QPointingDevice* self, DeviceType devType);
extern __declspec(dllexport) 
void QPointingDevice_SetCapabilities(QPointingDevice* self, int caps);
extern __declspec(dllexport) 
void QPointingDevice_SetMaximumTouchPoints(QPointingDevice* self, int c);
extern __declspec(dllexport) 
PointerType QPointingDevice_PointerType(const QPointingDevice* self);
extern __declspec(dllexport) 
int QPointingDevice_MaximumPoints(const QPointingDevice* self);
extern __declspec(dllexport) 
int QPointingDevice_ButtonCount(const QPointingDevice* self);
extern __declspec(dllexport) 
QPointingDeviceUniqueId* QPointingDevice_UniqueId(const QPointingDevice* self);
extern __declspec(dllexport) 
QPointingDevice* QPointingDevice_PrimaryPointingDevice();
extern __declspec(dllexport) 
bool QPointingDevice_OperatorEqual(const QPointingDevice* self, QPointingDevice* other);
extern __declspec(dllexport) 
void QPointingDevice_GrabChanged(const QPointingDevice* self, QObject* grabber, GrabTransition transition, QPointerEvent* event, QEventPoint* point);
void QPointingDevice_connect_GrabChanged(QPointingDevice* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QPointingDevice_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QPointingDevice_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
QPointingDevice* QPointingDevice_PrimaryPointingDevice1(struct miqt_string seatName);
extern __declspec(dllexport) 
void QPointingDevice_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QPointingDevice_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QPointingDevice_override_virtual_Metacast(void* self, intptr_t slot);
void* QPointingDevice_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QPointingDevice_Delete(QPointingDevice* self, bool isSubclass);

}
