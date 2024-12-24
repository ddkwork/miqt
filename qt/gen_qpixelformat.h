#pragma once
#ifndef MIQT_QT_GEN_QPIXELFORMAT_H
#define MIQT_QT_GEN_QPIXELFORMAT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QPixelFormat QPixelFormat;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QPixelFormat* QPixelFormat_new();
extern __declspec(dllexport) 
QPixelFormat* QPixelFormat_new2(ColorModel colorModel, unsigned char firstSize, unsigned char secondSize, unsigned char thirdSize, unsigned char fourthSize, unsigned char fifthSize, unsigned char alphaSize, AlphaUsage alphaUsage, AlphaPosition alphaPosition, AlphaPremultiplied premultiplied, TypeInterpretation typeInterpretation);
extern __declspec(dllexport) 
QPixelFormat* QPixelFormat_new3(QPixelFormat* param1);
extern __declspec(dllexport) 
QPixelFormat* QPixelFormat_new4(ColorModel colorModel, unsigned char firstSize, unsigned char secondSize, unsigned char thirdSize, unsigned char fourthSize, unsigned char fifthSize, unsigned char alphaSize, AlphaUsage alphaUsage, AlphaPosition alphaPosition, AlphaPremultiplied premultiplied, TypeInterpretation typeInterpretation, ByteOrder byteOrder);
extern __declspec(dllexport) 
QPixelFormat* QPixelFormat_new5(ColorModel colorModel, unsigned char firstSize, unsigned char secondSize, unsigned char thirdSize, unsigned char fourthSize, unsigned char fifthSize, unsigned char alphaSize, AlphaUsage alphaUsage, AlphaPosition alphaPosition, AlphaPremultiplied premultiplied, TypeInterpretation typeInterpretation, ByteOrder byteOrder, unsigned char subEnum);
extern __declspec(dllexport) 
ColorModel QPixelFormat_ColorModel(const QPixelFormat* self);
extern __declspec(dllexport) 
unsigned char QPixelFormat_ChannelCount(const QPixelFormat* self);
extern __declspec(dllexport) 
unsigned char QPixelFormat_RedSize(const QPixelFormat* self);
extern __declspec(dllexport) 
unsigned char QPixelFormat_GreenSize(const QPixelFormat* self);
extern __declspec(dllexport) 
unsigned char QPixelFormat_BlueSize(const QPixelFormat* self);
extern __declspec(dllexport) 
unsigned char QPixelFormat_CyanSize(const QPixelFormat* self);
extern __declspec(dllexport) 
unsigned char QPixelFormat_MagentaSize(const QPixelFormat* self);
extern __declspec(dllexport) 
unsigned char QPixelFormat_YellowSize(const QPixelFormat* self);
extern __declspec(dllexport) 
unsigned char QPixelFormat_BlackSize(const QPixelFormat* self);
extern __declspec(dllexport) 
unsigned char QPixelFormat_HueSize(const QPixelFormat* self);
extern __declspec(dllexport) 
unsigned char QPixelFormat_SaturationSize(const QPixelFormat* self);
extern __declspec(dllexport) 
unsigned char QPixelFormat_LightnessSize(const QPixelFormat* self);
extern __declspec(dllexport) 
unsigned char QPixelFormat_BrightnessSize(const QPixelFormat* self);
extern __declspec(dllexport) 
unsigned char QPixelFormat_AlphaSize(const QPixelFormat* self);
extern __declspec(dllexport) 
unsigned char QPixelFormat_BitsPerPixel(const QPixelFormat* self);
extern __declspec(dllexport) 
AlphaUsage QPixelFormat_AlphaUsage(const QPixelFormat* self);
extern __declspec(dllexport) 
AlphaPosition QPixelFormat_AlphaPosition(const QPixelFormat* self);
extern __declspec(dllexport) 
AlphaPremultiplied QPixelFormat_Premultiplied(const QPixelFormat* self);
extern __declspec(dllexport) 
TypeInterpretation QPixelFormat_TypeInterpretation(const QPixelFormat* self);
extern __declspec(dllexport) 
ByteOrder QPixelFormat_ByteOrder(const QPixelFormat* self);
extern __declspec(dllexport) 
YUVLayout QPixelFormat_YuvLayout(const QPixelFormat* self);
extern __declspec(dllexport) 
unsigned char QPixelFormat_SubEnum(const QPixelFormat* self);
extern __declspec(dllexport) 
void QPixelFormat_Delete(QPixelFormat* self, bool isSubclass);

}
