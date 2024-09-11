#ifndef GEN_QPIXELFORMAT_H
#define GEN_QPIXELFORMAT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QPixelFormat;
#else
typedef struct QPixelFormat QPixelFormat;
#endif

QPixelFormat* QPixelFormat_new();
QPixelFormat* QPixelFormat_new2(uintptr_t colorModel, unsigned char firstSize, unsigned char secondSize, unsigned char thirdSize, unsigned char fourthSize, unsigned char fifthSize, unsigned char alphaSize, uintptr_t alphaUsage, uintptr_t alphaPosition, uintptr_t premultiplied, uintptr_t typeInterpretation);
QPixelFormat* QPixelFormat_new3(QPixelFormat* param1);
QPixelFormat* QPixelFormat_new4(uintptr_t colorModel, unsigned char firstSize, unsigned char secondSize, unsigned char thirdSize, unsigned char fourthSize, unsigned char fifthSize, unsigned char alphaSize, uintptr_t alphaUsage, uintptr_t alphaPosition, uintptr_t premultiplied, uintptr_t typeInterpretation, uintptr_t byteOrder);
QPixelFormat* QPixelFormat_new5(uintptr_t colorModel, unsigned char firstSize, unsigned char secondSize, unsigned char thirdSize, unsigned char fourthSize, unsigned char fifthSize, unsigned char alphaSize, uintptr_t alphaUsage, uintptr_t alphaPosition, uintptr_t premultiplied, uintptr_t typeInterpretation, uintptr_t byteOrder, unsigned char subEnum);
uintptr_t QPixelFormat_ColorModel(const QPixelFormat* self);
unsigned char QPixelFormat_ChannelCount(const QPixelFormat* self);
unsigned char QPixelFormat_RedSize(const QPixelFormat* self);
unsigned char QPixelFormat_GreenSize(const QPixelFormat* self);
unsigned char QPixelFormat_BlueSize(const QPixelFormat* self);
unsigned char QPixelFormat_CyanSize(const QPixelFormat* self);
unsigned char QPixelFormat_MagentaSize(const QPixelFormat* self);
unsigned char QPixelFormat_YellowSize(const QPixelFormat* self);
unsigned char QPixelFormat_BlackSize(const QPixelFormat* self);
unsigned char QPixelFormat_HueSize(const QPixelFormat* self);
unsigned char QPixelFormat_SaturationSize(const QPixelFormat* self);
unsigned char QPixelFormat_LightnessSize(const QPixelFormat* self);
unsigned char QPixelFormat_BrightnessSize(const QPixelFormat* self);
unsigned char QPixelFormat_AlphaSize(const QPixelFormat* self);
unsigned char QPixelFormat_BitsPerPixel(const QPixelFormat* self);
uintptr_t QPixelFormat_AlphaUsage(const QPixelFormat* self);
uintptr_t QPixelFormat_AlphaPosition(const QPixelFormat* self);
uintptr_t QPixelFormat_Premultiplied(const QPixelFormat* self);
uintptr_t QPixelFormat_TypeInterpretation(const QPixelFormat* self);
uintptr_t QPixelFormat_ByteOrder(const QPixelFormat* self);
uintptr_t QPixelFormat_YuvLayout(const QPixelFormat* self);
unsigned char QPixelFormat_SubEnum(const QPixelFormat* self);
void QPixelFormat_Delete(QPixelFormat* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
