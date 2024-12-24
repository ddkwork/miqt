#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QCAMERADEVICE_H
#define MIQT_QT_MULTIMEDIA_GEN_QCAMERADEVICE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QCameraDevice QCameraDevice;
typedef struct QCameraFormat QCameraFormat;
typedef struct QSize QSize;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QCameraFormat* QCameraFormat_new();
extern __declspec(dllexport) 
QCameraFormat* QCameraFormat_new2(QCameraFormat* other);
extern __declspec(dllexport) 
void QCameraFormat_OperatorAssign(QCameraFormat* self, QCameraFormat* other);
extern __declspec(dllexport) 
int QCameraFormat_PixelFormat(const QCameraFormat* self);
extern __declspec(dllexport) 
QSize* QCameraFormat_Resolution(const QCameraFormat* self);
extern __declspec(dllexport) 
float QCameraFormat_MinFrameRate(const QCameraFormat* self);
extern __declspec(dllexport) 
float QCameraFormat_MaxFrameRate(const QCameraFormat* self);
extern __declspec(dllexport) 
bool QCameraFormat_IsNull(const QCameraFormat* self);
extern __declspec(dllexport) 
bool QCameraFormat_OperatorEqual(const QCameraFormat* self, QCameraFormat* other);
extern __declspec(dllexport) 
bool QCameraFormat_OperatorNotEqual(const QCameraFormat* self, QCameraFormat* other);
extern __declspec(dllexport) 
void QCameraFormat_Delete(QCameraFormat* self, bool isSubclass);

extern __declspec(dllexport) 
QCameraDevice* QCameraDevice_new();
extern __declspec(dllexport) 
QCameraDevice* QCameraDevice_new2(QCameraDevice* other);
extern __declspec(dllexport) 
void QCameraDevice_OperatorAssign(QCameraDevice* self, QCameraDevice* other);
extern __declspec(dllexport) 
bool QCameraDevice_OperatorEqual(const QCameraDevice* self, QCameraDevice* other);
extern __declspec(dllexport) 
bool QCameraDevice_OperatorNotEqual(const QCameraDevice* self, QCameraDevice* other);
extern __declspec(dllexport) 
bool QCameraDevice_IsNull(const QCameraDevice* self);
extern __declspec(dllexport) 
struct miqt_string QCameraDevice_Id(const QCameraDevice* self);
extern __declspec(dllexport) 
struct miqt_string QCameraDevice_Description(const QCameraDevice* self);
extern __declspec(dllexport) 
bool QCameraDevice_IsDefault(const QCameraDevice* self);
extern __declspec(dllexport) 
Position QCameraDevice_Position(const QCameraDevice* self);
extern __declspec(dllexport) 
struct miqt_array /* of QSize* */  QCameraDevice_PhotoResolutions(const QCameraDevice* self);
extern __declspec(dllexport) 
struct miqt_array /* of QCameraFormat* */  QCameraDevice_VideoFormats(const QCameraDevice* self);
extern __declspec(dllexport) 
int QCameraDevice_CorrectionAngle(const QCameraDevice* self);
extern __declspec(dllexport) 
void QCameraDevice_Delete(QCameraDevice* self, bool isSubclass);

}
