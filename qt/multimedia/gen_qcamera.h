#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QCAMERA_H
#define MIQT_QT_MULTIMEDIA_GEN_QCAMERA_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QCamera QCamera;
typedef struct QCameraDevice QCameraDevice;
typedef struct QCameraFormat QCameraFormat;
typedef struct QMediaCaptureSession QMediaCaptureSession;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPointF QPointF;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QCamera* QCamera_new();
extern __declspec(dllexport) 
QCamera* QCamera_new2(QCameraDevice* cameraDevice);
extern __declspec(dllexport) 
QCamera* QCamera_new3(int position);
extern __declspec(dllexport) 
QCamera* QCamera_new4(QObject* parent);
extern __declspec(dllexport) 
QCamera* QCamera_new5(QCameraDevice* cameraDevice, QObject* parent);
extern __declspec(dllexport) 
QCamera* QCamera_new6(int position, QObject* parent);
extern __declspec(dllexport) 
void QCamera_virtbase(QCamera* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QCamera_MetaObject(const QCamera* self);
extern __declspec(dllexport) 
void* QCamera_Metacast(QCamera* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QCamera_Tr(const char* s);
extern __declspec(dllexport) 
bool QCamera_IsAvailable(const QCamera* self);
extern __declspec(dllexport) 
bool QCamera_IsActive(const QCamera* self);
extern __declspec(dllexport) 
QMediaCaptureSession* QCamera_CaptureSession(const QCamera* self);
extern __declspec(dllexport) 
QCameraDevice* QCamera_CameraDevice(const QCamera* self);
extern __declspec(dllexport) 
void QCamera_SetCameraDevice(QCamera* self, QCameraDevice* cameraDevice);
extern __declspec(dllexport) 
QCameraFormat* QCamera_CameraFormat(const QCamera* self);
extern __declspec(dllexport) 
void QCamera_SetCameraFormat(QCamera* self, QCameraFormat* format);
extern __declspec(dllexport) 
Error QCamera_Error(const QCamera* self);
extern __declspec(dllexport) 
struct miqt_string QCamera_ErrorString(const QCamera* self);
extern __declspec(dllexport) 
Features QCamera_SupportedFeatures(const QCamera* self);
extern __declspec(dllexport) 
FocusMode QCamera_FocusMode(const QCamera* self);
extern __declspec(dllexport) 
void QCamera_SetFocusMode(QCamera* self, FocusMode mode);
extern __declspec(dllexport) 
bool QCamera_IsFocusModeSupported(const QCamera* self, FocusMode mode);
extern __declspec(dllexport) 
QPointF* QCamera_FocusPoint(const QCamera* self);
extern __declspec(dllexport) 
QPointF* QCamera_CustomFocusPoint(const QCamera* self);
extern __declspec(dllexport) 
void QCamera_SetCustomFocusPoint(QCamera* self, QPointF* point);
extern __declspec(dllexport) 
void QCamera_SetFocusDistance(QCamera* self, float d);
extern __declspec(dllexport) 
float QCamera_FocusDistance(const QCamera* self);
extern __declspec(dllexport) 
float QCamera_MinimumZoomFactor(const QCamera* self);
extern __declspec(dllexport) 
float QCamera_MaximumZoomFactor(const QCamera* self);
extern __declspec(dllexport) 
float QCamera_ZoomFactor(const QCamera* self);
extern __declspec(dllexport) 
void QCamera_SetZoomFactor(QCamera* self, float factor);
extern __declspec(dllexport) 
FlashMode QCamera_FlashMode(const QCamera* self);
extern __declspec(dllexport) 
bool QCamera_IsFlashModeSupported(const QCamera* self, FlashMode mode);
extern __declspec(dllexport) 
bool QCamera_IsFlashReady(const QCamera* self);
extern __declspec(dllexport) 
TorchMode QCamera_TorchMode(const QCamera* self);
extern __declspec(dllexport) 
bool QCamera_IsTorchModeSupported(const QCamera* self, TorchMode mode);
extern __declspec(dllexport) 
ExposureMode QCamera_ExposureMode(const QCamera* self);
extern __declspec(dllexport) 
bool QCamera_IsExposureModeSupported(const QCamera* self, ExposureMode mode);
extern __declspec(dllexport) 
float QCamera_ExposureCompensation(const QCamera* self);
extern __declspec(dllexport) 
int QCamera_IsoSensitivity(const QCamera* self);
extern __declspec(dllexport) 
int QCamera_ManualIsoSensitivity(const QCamera* self);
extern __declspec(dllexport) 
float QCamera_ExposureTime(const QCamera* self);
extern __declspec(dllexport) 
float QCamera_ManualExposureTime(const QCamera* self);
extern __declspec(dllexport) 
int QCamera_MinimumIsoSensitivity(const QCamera* self);
extern __declspec(dllexport) 
int QCamera_MaximumIsoSensitivity(const QCamera* self);
extern __declspec(dllexport) 
float QCamera_MinimumExposureTime(const QCamera* self);
extern __declspec(dllexport) 
float QCamera_MaximumExposureTime(const QCamera* self);
extern __declspec(dllexport) 
WhiteBalanceMode QCamera_WhiteBalanceMode(const QCamera* self);
extern __declspec(dllexport) 
bool QCamera_IsWhiteBalanceModeSupported(const QCamera* self, WhiteBalanceMode mode);
extern __declspec(dllexport) 
int QCamera_ColorTemperature(const QCamera* self);
extern __declspec(dllexport) 
void QCamera_SetActive(QCamera* self, bool active);
extern __declspec(dllexport) 
void QCamera_Start(QCamera* self);
extern __declspec(dllexport) 
void QCamera_Stop(QCamera* self);
extern __declspec(dllexport) 
void QCamera_ZoomTo(QCamera* self, float zoom, float rate);
extern __declspec(dllexport) 
void QCamera_SetFlashMode(QCamera* self, FlashMode mode);
extern __declspec(dllexport) 
void QCamera_SetTorchMode(QCamera* self, TorchMode mode);
extern __declspec(dllexport) 
void QCamera_SetExposureMode(QCamera* self, ExposureMode mode);
extern __declspec(dllexport) 
void QCamera_SetExposureCompensation(QCamera* self, float ev);
extern __declspec(dllexport) 
void QCamera_SetManualIsoSensitivity(QCamera* self, int iso);
extern __declspec(dllexport) 
void QCamera_SetAutoIsoSensitivity(QCamera* self);
extern __declspec(dllexport) 
void QCamera_SetManualExposureTime(QCamera* self, float seconds);
extern __declspec(dllexport) 
void QCamera_SetAutoExposureTime(QCamera* self);
extern __declspec(dllexport) 
void QCamera_SetWhiteBalanceMode(QCamera* self, WhiteBalanceMode mode);
extern __declspec(dllexport) 
void QCamera_SetColorTemperature(QCamera* self, int colorTemperature);
extern __declspec(dllexport) 
void QCamera_ActiveChanged(QCamera* self, bool param1);
void QCamera_connect_ActiveChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_ErrorChanged(QCamera* self);
void QCamera_connect_ErrorChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_ErrorOccurred(QCamera* self, int error, struct miqt_string errorString);
void QCamera_connect_ErrorOccurred(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_CameraDeviceChanged(QCamera* self);
void QCamera_connect_CameraDeviceChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_CameraFormatChanged(QCamera* self);
void QCamera_connect_CameraFormatChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_SupportedFeaturesChanged(QCamera* self);
void QCamera_connect_SupportedFeaturesChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_FocusModeChanged(QCamera* self);
void QCamera_connect_FocusModeChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_ZoomFactorChanged(QCamera* self, float param1);
void QCamera_connect_ZoomFactorChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_MinimumZoomFactorChanged(QCamera* self, float param1);
void QCamera_connect_MinimumZoomFactorChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_MaximumZoomFactorChanged(QCamera* self, float param1);
void QCamera_connect_MaximumZoomFactorChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_FocusDistanceChanged(QCamera* self, float param1);
void QCamera_connect_FocusDistanceChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_FocusPointChanged(QCamera* self);
void QCamera_connect_FocusPointChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_CustomFocusPointChanged(QCamera* self);
void QCamera_connect_CustomFocusPointChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_FlashReady(QCamera* self, bool param1);
void QCamera_connect_FlashReady(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_FlashModeChanged(QCamera* self);
void QCamera_connect_FlashModeChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_TorchModeChanged(QCamera* self);
void QCamera_connect_TorchModeChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_ExposureTimeChanged(QCamera* self, float speed);
void QCamera_connect_ExposureTimeChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_ManualExposureTimeChanged(QCamera* self, float speed);
void QCamera_connect_ManualExposureTimeChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_IsoSensitivityChanged(QCamera* self, int param1);
void QCamera_connect_IsoSensitivityChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_ManualIsoSensitivityChanged(QCamera* self, int param1);
void QCamera_connect_ManualIsoSensitivityChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_ExposureCompensationChanged(QCamera* self, float param1);
void QCamera_connect_ExposureCompensationChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_ExposureModeChanged(QCamera* self);
void QCamera_connect_ExposureModeChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_WhiteBalanceModeChanged(const QCamera* self);
void QCamera_connect_WhiteBalanceModeChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_ColorTemperatureChanged(const QCamera* self);
void QCamera_connect_ColorTemperatureChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_BrightnessChanged(QCamera* self);
void QCamera_connect_BrightnessChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_ContrastChanged(QCamera* self);
void QCamera_connect_ContrastChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_SaturationChanged(QCamera* self);
void QCamera_connect_SaturationChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
void QCamera_HueChanged(QCamera* self);
void QCamera_connect_HueChanged(QCamera* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QCamera_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QCamera_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QCamera_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QCamera_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QCamera_override_virtual_Metacast(void* self, intptr_t slot);
void* QCamera_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QCamera_Delete(QCamera* self, bool isSubclass);

}
