#pragma once
#ifndef MIQT_QT_GEN_QPERMISSIONS_H
#define MIQT_QT_GEN_QPERMISSIONS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QBluetoothPermission;
class QCalendarPermission;
class QCameraPermission;
class QContactsPermission;
class QLocationPermission;
class QMetaType;
class QMicrophonePermission;
class QPermission;
class _GUID;
class type_info;
#else
typedef struct QBluetoothPermission QBluetoothPermission;
typedef struct QCalendarPermission QCalendarPermission;
typedef struct QCameraPermission QCameraPermission;
typedef struct QContactsPermission QContactsPermission;
typedef struct QLocationPermission QLocationPermission;
typedef struct QMetaType QMetaType;
typedef struct QMicrophonePermission QMicrophonePermission;
typedef struct QPermission QPermission;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QPermission* QPermission_new();
extern __declspec(dllexport) QPermission* QPermission_new2(QPermission* param1);
extern __declspec(dllexport) int QPermission_Status(const QPermission* self);
extern __declspec(dllexport) QMetaType* QPermission_Type(const QPermission* self);
extern __declspec(dllexport) void QPermission_Delete(QPermission* self, bool isSubclass);

extern __declspec(dllexport) QLocationPermission* QLocationPermission_new();
extern __declspec(dllexport) QLocationPermission* QLocationPermission_new2(QLocationPermission* other);
extern __declspec(dllexport) void QLocationPermission_SetAccuracy(QLocationPermission* self, Accuracy accuracy);
extern __declspec(dllexport) Accuracy QLocationPermission_Accuracy(const QLocationPermission* self);
extern __declspec(dllexport) void QLocationPermission_SetAvailability(QLocationPermission* self, Availability availability);
extern __declspec(dllexport) Availability QLocationPermission_Availability(const QLocationPermission* self);
extern __declspec(dllexport) void QLocationPermission_OperatorAssign(QLocationPermission* self, QLocationPermission* other);
void QLocationPermission_connect_OperatorAssign(QLocationPermission* self, intptr_t slot);
extern __declspec(dllexport) void QLocationPermission_Swap(QLocationPermission* self, QLocationPermission* other);
void QLocationPermission_connect_Swap(QLocationPermission* self, intptr_t slot);
extern __declspec(dllexport) void QLocationPermission_Delete(QLocationPermission* self, bool isSubclass);

extern __declspec(dllexport) QCalendarPermission* QCalendarPermission_new();
extern __declspec(dllexport) QCalendarPermission* QCalendarPermission_new2(QCalendarPermission* other);
extern __declspec(dllexport) void QCalendarPermission_SetAccessMode(QCalendarPermission* self, AccessMode mode);
extern __declspec(dllexport) AccessMode QCalendarPermission_AccessMode(const QCalendarPermission* self);
extern __declspec(dllexport) void QCalendarPermission_OperatorAssign(QCalendarPermission* self, QCalendarPermission* other);
void QCalendarPermission_connect_OperatorAssign(QCalendarPermission* self, intptr_t slot);
extern __declspec(dllexport) void QCalendarPermission_Swap(QCalendarPermission* self, QCalendarPermission* other);
void QCalendarPermission_connect_Swap(QCalendarPermission* self, intptr_t slot);
extern __declspec(dllexport) void QCalendarPermission_Delete(QCalendarPermission* self, bool isSubclass);

extern __declspec(dllexport) QContactsPermission* QContactsPermission_new();
extern __declspec(dllexport) QContactsPermission* QContactsPermission_new2(QContactsPermission* other);
extern __declspec(dllexport) void QContactsPermission_SetAccessMode(QContactsPermission* self, AccessMode mode);
extern __declspec(dllexport) AccessMode QContactsPermission_AccessMode(const QContactsPermission* self);
extern __declspec(dllexport) void QContactsPermission_OperatorAssign(QContactsPermission* self, QContactsPermission* other);
void QContactsPermission_connect_OperatorAssign(QContactsPermission* self, intptr_t slot);
extern __declspec(dllexport) void QContactsPermission_Swap(QContactsPermission* self, QContactsPermission* other);
void QContactsPermission_connect_Swap(QContactsPermission* self, intptr_t slot);
extern __declspec(dllexport) void QContactsPermission_Delete(QContactsPermission* self, bool isSubclass);

extern __declspec(dllexport) QBluetoothPermission* QBluetoothPermission_new();
extern __declspec(dllexport) QBluetoothPermission* QBluetoothPermission_new2(QBluetoothPermission* other);
extern __declspec(dllexport) void QBluetoothPermission_SetCommunicationModes(QBluetoothPermission* self, CommunicationModes modes);
extern __declspec(dllexport) CommunicationModes QBluetoothPermission_CommunicationModes(const QBluetoothPermission* self);
extern __declspec(dllexport) void QBluetoothPermission_OperatorAssign(QBluetoothPermission* self, QBluetoothPermission* other);
void QBluetoothPermission_connect_OperatorAssign(QBluetoothPermission* self, intptr_t slot);
extern __declspec(dllexport) void QBluetoothPermission_Swap(QBluetoothPermission* self, QBluetoothPermission* other);
void QBluetoothPermission_connect_Swap(QBluetoothPermission* self, intptr_t slot);
extern __declspec(dllexport) void QBluetoothPermission_Delete(QBluetoothPermission* self, bool isSubclass);

extern __declspec(dllexport) QCameraPermission* QCameraPermission_new();
extern __declspec(dllexport) QCameraPermission* QCameraPermission_new2(QCameraPermission* other);
extern __declspec(dllexport) void QCameraPermission_OperatorAssign(QCameraPermission* self, QCameraPermission* other);
void QCameraPermission_connect_OperatorAssign(QCameraPermission* self, intptr_t slot);
extern __declspec(dllexport) void QCameraPermission_Swap(QCameraPermission* self, QCameraPermission* other);
void QCameraPermission_connect_Swap(QCameraPermission* self, intptr_t slot);
extern __declspec(dllexport) void QCameraPermission_Delete(QCameraPermission* self, bool isSubclass);

extern __declspec(dllexport) QMicrophonePermission* QMicrophonePermission_new();
extern __declspec(dllexport) QMicrophonePermission* QMicrophonePermission_new2(QMicrophonePermission* other);
extern __declspec(dllexport) void QMicrophonePermission_OperatorAssign(QMicrophonePermission* self, QMicrophonePermission* other);
void QMicrophonePermission_connect_OperatorAssign(QMicrophonePermission* self, intptr_t slot);
extern __declspec(dllexport) void QMicrophonePermission_Swap(QMicrophonePermission* self, QMicrophonePermission* other);
void QMicrophonePermission_connect_Swap(QMicrophonePermission* self, intptr_t slot);
extern __declspec(dllexport) void QMicrophonePermission_Delete(QMicrophonePermission* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
