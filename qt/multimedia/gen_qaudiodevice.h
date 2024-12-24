#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QAUDIODEVICE_H
#define MIQT_QT_MULTIMEDIA_GEN_QAUDIODEVICE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAudioDevice QAudioDevice;
typedef struct QAudioFormat QAudioFormat;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAudioDevice* QAudioDevice_new();
extern __declspec(dllexport) 
QAudioDevice* QAudioDevice_new2(QAudioDevice* other);
extern __declspec(dllexport) 
void QAudioDevice_Swap(QAudioDevice* self, QAudioDevice* other);
extern __declspec(dllexport) 
void QAudioDevice_OperatorAssign(QAudioDevice* self, QAudioDevice* other);
extern __declspec(dllexport) 
bool QAudioDevice_OperatorEqual(const QAudioDevice* self, QAudioDevice* other);
extern __declspec(dllexport) 
bool QAudioDevice_OperatorNotEqual(const QAudioDevice* self, QAudioDevice* other);
extern __declspec(dllexport) 
bool QAudioDevice_IsNull(const QAudioDevice* self);
extern __declspec(dllexport) 
struct miqt_string QAudioDevice_Id(const QAudioDevice* self);
extern __declspec(dllexport) 
struct miqt_string QAudioDevice_Description(const QAudioDevice* self);
extern __declspec(dllexport) 
bool QAudioDevice_IsDefault(const QAudioDevice* self);
extern __declspec(dllexport) 
int QAudioDevice_Mode(const QAudioDevice* self);
extern __declspec(dllexport) 
bool QAudioDevice_IsFormatSupported(const QAudioDevice* self, QAudioFormat* format);
extern __declspec(dllexport) 
QAudioFormat* QAudioDevice_PreferredFormat(const QAudioDevice* self);
extern __declspec(dllexport) 
int QAudioDevice_MinimumSampleRate(const QAudioDevice* self);
extern __declspec(dllexport) 
int QAudioDevice_MaximumSampleRate(const QAudioDevice* self);
extern __declspec(dllexport) 
int QAudioDevice_MinimumChannelCount(const QAudioDevice* self);
extern __declspec(dllexport) 
int QAudioDevice_MaximumChannelCount(const QAudioDevice* self);
extern __declspec(dllexport) 
struct miqt_array /* of uint16_t */  QAudioDevice_SupportedSampleFormats(const QAudioDevice* self);
extern __declspec(dllexport) 
uint32_t QAudioDevice_ChannelConfiguration(const QAudioDevice* self);
extern __declspec(dllexport) 
void QAudioDevice_Delete(QAudioDevice* self, bool isSubclass);

}
