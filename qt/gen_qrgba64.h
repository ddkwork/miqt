#pragma once
#ifndef MIQT_QT_GEN_QRGBA64_H
#define MIQT_QT_GEN_QRGBA64_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QRgba64 QRgba64;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QRgba64* QRgba64_new();
extern __declspec(dllexport) QRgba64* QRgba64_new2(QRgba64* param1);
extern __declspec(dllexport) QRgba64* QRgba64_FromRgba64(unsigned long long c);
extern __declspec(dllexport) QRgba64* QRgba64_FromRgba642(uint16_t red, uint16_t green, uint16_t blue, uint16_t alpha);
extern __declspec(dllexport) QRgba64* QRgba64_FromRgba(unsigned char red, unsigned char green, unsigned char blue, unsigned char alpha);
extern __declspec(dllexport) QRgba64* QRgba64_FromArgb32(unsigned int rgb);
extern __declspec(dllexport) bool QRgba64_IsOpaque(const QRgba64* self);
extern __declspec(dllexport) bool QRgba64_IsTransparent(const QRgba64* self);
extern __declspec(dllexport) uint16_t QRgba64_Red(const QRgba64* self);
extern __declspec(dllexport) uint16_t QRgba64_Green(const QRgba64* self);
extern __declspec(dllexport) uint16_t QRgba64_Blue(const QRgba64* self);
extern __declspec(dllexport) uint16_t QRgba64_Alpha(const QRgba64* self);
extern __declspec(dllexport) void QRgba64_SetRed(QRgba64* self, uint16_t _red);
extern __declspec(dllexport) void QRgba64_SetGreen(QRgba64* self, uint16_t _green);
extern __declspec(dllexport) void QRgba64_SetBlue(QRgba64* self, uint16_t _blue);
extern __declspec(dllexport) void QRgba64_SetAlpha(QRgba64* self, uint16_t _alpha);
extern __declspec(dllexport) unsigned char QRgba64_Red8(const QRgba64* self);
extern __declspec(dllexport) unsigned char QRgba64_Green8(const QRgba64* self);
extern __declspec(dllexport) unsigned char QRgba64_Blue8(const QRgba64* self);
extern __declspec(dllexport) unsigned char QRgba64_Alpha8(const QRgba64* self);
extern __declspec(dllexport) unsigned int QRgba64_ToArgb32(const QRgba64* self);
extern __declspec(dllexport) uint16_t QRgba64_ToRgb16(const QRgba64* self);
extern __declspec(dllexport) QRgba64* QRgba64_Premultiplied(const QRgba64* self);
extern __declspec(dllexport) QRgba64* QRgba64_Unpremultiplied(const QRgba64* self);
extern __declspec(dllexport) void QRgba64_OperatorAssign(QRgba64* self, unsigned long long _rgba);
extern __declspec(dllexport) void QRgba64_Delete(QRgba64* self, bool isSubclass);

} 
