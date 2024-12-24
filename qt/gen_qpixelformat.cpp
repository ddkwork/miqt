// +build ignore

#include <QPixelFormat>
#include <qpixelformat.h>
#include "gen_qpixelformat.h"

#ifndef _Bool
#define _Bool bool
#endif

QPixelFormat* QPixelFormat_new() {
	return new QPixelFormat();
}

QPixelFormat* QPixelFormat_new2(ColorModel colorModel, unsigned char firstSize, unsigned char secondSize, unsigned char thirdSize, unsigned char fourthSize, unsigned char fifthSize, unsigned char alphaSize, AlphaUsage alphaUsage, AlphaPosition alphaPosition, AlphaPremultiplied premultiplied, TypeInterpretation typeInterpretation) {
	return new QPixelFormat(colorModel, static_cast<uchar>(firstSize), static_cast<uchar>(secondSize), static_cast<uchar>(thirdSize), static_cast<uchar>(fourthSize), static_cast<uchar>(fifthSize), static_cast<uchar>(alphaSize), alphaUsage, alphaPosition, premultiplied, typeInterpretation);
}

QPixelFormat* QPixelFormat_new3(QPixelFormat* param1) {
	return new QPixelFormat(*param1);
}

QPixelFormat* QPixelFormat_new4(ColorModel colorModel, unsigned char firstSize, unsigned char secondSize, unsigned char thirdSize, unsigned char fourthSize, unsigned char fifthSize, unsigned char alphaSize, AlphaUsage alphaUsage, AlphaPosition alphaPosition, AlphaPremultiplied premultiplied, TypeInterpretation typeInterpretation, ByteOrder byteOrder) {
	return new QPixelFormat(colorModel, static_cast<uchar>(firstSize), static_cast<uchar>(secondSize), static_cast<uchar>(thirdSize), static_cast<uchar>(fourthSize), static_cast<uchar>(fifthSize), static_cast<uchar>(alphaSize), alphaUsage, alphaPosition, premultiplied, typeInterpretation, byteOrder);
}

QPixelFormat* QPixelFormat_new5(ColorModel colorModel, unsigned char firstSize, unsigned char secondSize, unsigned char thirdSize, unsigned char fourthSize, unsigned char fifthSize, unsigned char alphaSize, AlphaUsage alphaUsage, AlphaPosition alphaPosition, AlphaPremultiplied premultiplied, TypeInterpretation typeInterpretation, ByteOrder byteOrder, unsigned char subEnum) {
	return new QPixelFormat(colorModel, static_cast<uchar>(firstSize), static_cast<uchar>(secondSize), static_cast<uchar>(thirdSize), static_cast<uchar>(fourthSize), static_cast<uchar>(fifthSize), static_cast<uchar>(alphaSize), alphaUsage, alphaPosition, premultiplied, typeInterpretation, byteOrder, static_cast<uchar>(subEnum));
}

ColorModel QPixelFormat_ColorModel(const QPixelFormat* self) {
	return self->colorModel();
}

unsigned char QPixelFormat_ChannelCount(const QPixelFormat* self) {
	uchar _ret = self->channelCount();
	return static_cast<unsigned char>(_ret);
}

unsigned char QPixelFormat_RedSize(const QPixelFormat* self) {
	uchar _ret = self->redSize();
	return static_cast<unsigned char>(_ret);
}

unsigned char QPixelFormat_GreenSize(const QPixelFormat* self) {
	uchar _ret = self->greenSize();
	return static_cast<unsigned char>(_ret);
}

unsigned char QPixelFormat_BlueSize(const QPixelFormat* self) {
	uchar _ret = self->blueSize();
	return static_cast<unsigned char>(_ret);
}

unsigned char QPixelFormat_CyanSize(const QPixelFormat* self) {
	uchar _ret = self->cyanSize();
	return static_cast<unsigned char>(_ret);
}

unsigned char QPixelFormat_MagentaSize(const QPixelFormat* self) {
	uchar _ret = self->magentaSize();
	return static_cast<unsigned char>(_ret);
}

unsigned char QPixelFormat_YellowSize(const QPixelFormat* self) {
	uchar _ret = self->yellowSize();
	return static_cast<unsigned char>(_ret);
}

unsigned char QPixelFormat_BlackSize(const QPixelFormat* self) {
	uchar _ret = self->blackSize();
	return static_cast<unsigned char>(_ret);
}

unsigned char QPixelFormat_HueSize(const QPixelFormat* self) {
	uchar _ret = self->hueSize();
	return static_cast<unsigned char>(_ret);
}

unsigned char QPixelFormat_SaturationSize(const QPixelFormat* self) {
	uchar _ret = self->saturationSize();
	return static_cast<unsigned char>(_ret);
}

unsigned char QPixelFormat_LightnessSize(const QPixelFormat* self) {
	uchar _ret = self->lightnessSize();
	return static_cast<unsigned char>(_ret);
}

unsigned char QPixelFormat_BrightnessSize(const QPixelFormat* self) {
	uchar _ret = self->brightnessSize();
	return static_cast<unsigned char>(_ret);
}

unsigned char QPixelFormat_AlphaSize(const QPixelFormat* self) {
	uchar _ret = self->alphaSize();
	return static_cast<unsigned char>(_ret);
}

unsigned char QPixelFormat_BitsPerPixel(const QPixelFormat* self) {
	uchar _ret = self->bitsPerPixel();
	return static_cast<unsigned char>(_ret);
}

AlphaUsage QPixelFormat_AlphaUsage(const QPixelFormat* self) {
	return self->alphaUsage();
}

AlphaPosition QPixelFormat_AlphaPosition(const QPixelFormat* self) {
	return self->alphaPosition();
}

AlphaPremultiplied QPixelFormat_Premultiplied(const QPixelFormat* self) {
	return self->premultiplied();
}

TypeInterpretation QPixelFormat_TypeInterpretation(const QPixelFormat* self) {
	return self->typeInterpretation();
}

ByteOrder QPixelFormat_ByteOrder(const QPixelFormat* self) {
	return self->byteOrder();
}

YUVLayout QPixelFormat_YuvLayout(const QPixelFormat* self) {
	return self->yuvLayout();
}

unsigned char QPixelFormat_SubEnum(const QPixelFormat* self) {
	uchar _ret = self->subEnum();
	return static_cast<unsigned char>(_ret);
}

void QPixelFormat_Delete(QPixelFormat* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QPixelFormat*>( self );
	} else {
		delete self;
	}
}

