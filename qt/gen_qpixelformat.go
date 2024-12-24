package qt

import (
	"unsafe"
)

type QPixelFormat__ColorModel int

const (
	QPixelFormat__RGB       QPixelFormat__ColorModel = 0
	QPixelFormat__BGR       QPixelFormat__ColorModel = 1
	QPixelFormat__Indexed   QPixelFormat__ColorModel = 2
	QPixelFormat__Grayscale QPixelFormat__ColorModel = 3
	QPixelFormat__CMYK      QPixelFormat__ColorModel = 4
	QPixelFormat__HSL       QPixelFormat__ColorModel = 5
	QPixelFormat__HSV       QPixelFormat__ColorModel = 6
	QPixelFormat__YUV       QPixelFormat__ColorModel = 7
	QPixelFormat__Alpha     QPixelFormat__ColorModel = 8
)

type QPixelFormat__AlphaUsage int

const (
	QPixelFormat__UsesAlpha    QPixelFormat__AlphaUsage = 0
	QPixelFormat__IgnoresAlpha QPixelFormat__AlphaUsage = 1
)

type QPixelFormat__AlphaPosition int

const (
	QPixelFormat__AtBeginning QPixelFormat__AlphaPosition = 0
	QPixelFormat__AtEnd       QPixelFormat__AlphaPosition = 1
)

type QPixelFormat__AlphaPremultiplied int

const (
	QPixelFormat__NotPremultiplied QPixelFormat__AlphaPremultiplied = 0
	QPixelFormat__Premultiplied    QPixelFormat__AlphaPremultiplied = 1
)

type QPixelFormat__TypeInterpretation int

const (
	QPixelFormat__UnsignedInteger QPixelFormat__TypeInterpretation = 0
	QPixelFormat__UnsignedShort   QPixelFormat__TypeInterpretation = 1
	QPixelFormat__UnsignedByte    QPixelFormat__TypeInterpretation = 2
	QPixelFormat__FloatingPoint   QPixelFormat__TypeInterpretation = 3
)

type QPixelFormat__YUVLayout int

const (
	QPixelFormat__YUV444   QPixelFormat__YUVLayout = 0
	QPixelFormat__YUV422   QPixelFormat__YUVLayout = 1
	QPixelFormat__YUV411   QPixelFormat__YUVLayout = 2
	QPixelFormat__YUV420P  QPixelFormat__YUVLayout = 3
	QPixelFormat__YUV420SP QPixelFormat__YUVLayout = 4
	QPixelFormat__YV12     QPixelFormat__YUVLayout = 5
	QPixelFormat__UYVY     QPixelFormat__YUVLayout = 6
	QPixelFormat__YUYV     QPixelFormat__YUVLayout = 7
	QPixelFormat__NV12     QPixelFormat__YUVLayout = 8
	QPixelFormat__NV21     QPixelFormat__YUVLayout = 9
	QPixelFormat__IMC1     QPixelFormat__YUVLayout = 10
	QPixelFormat__IMC2     QPixelFormat__YUVLayout = 11
	QPixelFormat__IMC3     QPixelFormat__YUVLayout = 12
	QPixelFormat__IMC4     QPixelFormat__YUVLayout = 13
	QPixelFormat__Y8       QPixelFormat__YUVLayout = 14
	QPixelFormat__Y16      QPixelFormat__YUVLayout = 15
)

type QPixelFormat__ByteOrder int

const (
	QPixelFormat__LittleEndian        QPixelFormat__ByteOrder = 0
	QPixelFormat__BigEndian           QPixelFormat__ByteOrder = 1
	QPixelFormat__CurrentSystemEndian QPixelFormat__ByteOrder = 2
)

type QPixelFormat struct {
	h          uintptr
	isSubclass bool
}

// NewQPixelFormat constructs a new QPixelFormat object.
func NewQPixelFormat() *QPixelFormat {
	g := newQPixelFormat(QPixelFormat_new())
	g.isSubclass = true
	return g
}

// NewQPixelFormat2 constructs a new QPixelFormat object.
func NewQPixelFormat2(colorModel ColorModel, firstSize byte, secondSize byte, thirdSize byte, fourthSize byte, fifthSize byte, alphaSize byte, alphaUsage AlphaUsage, alphaPosition AlphaPosition, premultiplied AlphaPremultiplied, typeInterpretation TypeInterpretation) *QPixelFormat {
	g := newQPixelFormat(QPixelFormat_new2(colorModel, (uchar)(firstSize), (uchar)(secondSize), (uchar)(thirdSize), (uchar)(fourthSize), (uchar)(fifthSize), (uchar)(alphaSize), alphaUsage, alphaPosition, premultiplied, typeInterpretation))
	g.isSubclass = true
	return g
}

// NewQPixelFormat3 constructs a new QPixelFormat object.
func NewQPixelFormat3(param1 *QPixelFormat) *QPixelFormat {
	g := newQPixelFormat(QPixelFormat_new3(param1.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPixelFormat4 constructs a new QPixelFormat object.
func NewQPixelFormat4(colorModel ColorModel, firstSize byte, secondSize byte, thirdSize byte, fourthSize byte, fifthSize byte, alphaSize byte, alphaUsage AlphaUsage, alphaPosition AlphaPosition, premultiplied AlphaPremultiplied, typeInterpretation TypeInterpretation, byteOrder ByteOrder) *QPixelFormat {
	g := newQPixelFormat(QPixelFormat_new4(colorModel, (uchar)(firstSize), (uchar)(secondSize), (uchar)(thirdSize), (uchar)(fourthSize), (uchar)(fifthSize), (uchar)(alphaSize), alphaUsage, alphaPosition, premultiplied, typeInterpretation, byteOrder))
	g.isSubclass = true
	return g
}

// NewQPixelFormat5 constructs a new QPixelFormat object.
func NewQPixelFormat5(colorModel ColorModel, firstSize byte, secondSize byte, thirdSize byte, fourthSize byte, fifthSize byte, alphaSize byte, alphaUsage AlphaUsage, alphaPosition AlphaPosition, premultiplied AlphaPremultiplied, typeInterpretation TypeInterpretation, byteOrder ByteOrder, subEnum byte) *QPixelFormat {
	g := newQPixelFormat(QPixelFormat_new5(colorModel, (uchar)(firstSize), (uchar)(secondSize), (uchar)(thirdSize), (uchar)(fourthSize), (uchar)(fifthSize), (uchar)(alphaSize), alphaUsage, alphaPosition, premultiplied, typeInterpretation, byteOrder, (uchar)(subEnum)))
	g.isSubclass = true
	return g
}

func (this *QPixelFormat) ColorModel() ColorModel {
	xxxxxxxxx
}

func (this *QPixelFormat) ChannelCount() byte {
	return (byte)(QPixelFormat_ChannelCount(this.h))
}

func (this *QPixelFormat) RedSize() byte {
	return (byte)(QPixelFormat_RedSize(this.h))
}

func (this *QPixelFormat) GreenSize() byte {
	return (byte)(QPixelFormat_GreenSize(this.h))
}

func (this *QPixelFormat) BlueSize() byte {
	return (byte)(QPixelFormat_BlueSize(this.h))
}

func (this *QPixelFormat) CyanSize() byte {
	return (byte)(QPixelFormat_CyanSize(this.h))
}

func (this *QPixelFormat) MagentaSize() byte {
	return (byte)(QPixelFormat_MagentaSize(this.h))
}

func (this *QPixelFormat) YellowSize() byte {
	return (byte)(QPixelFormat_YellowSize(this.h))
}

func (this *QPixelFormat) BlackSize() byte {
	return (byte)(QPixelFormat_BlackSize(this.h))
}

func (this *QPixelFormat) HueSize() byte {
	return (byte)(QPixelFormat_HueSize(this.h))
}

func (this *QPixelFormat) SaturationSize() byte {
	return (byte)(QPixelFormat_SaturationSize(this.h))
}

func (this *QPixelFormat) LightnessSize() byte {
	return (byte)(QPixelFormat_LightnessSize(this.h))
}

func (this *QPixelFormat) BrightnessSize() byte {
	return (byte)(QPixelFormat_BrightnessSize(this.h))
}

func (this *QPixelFormat) AlphaSize() byte {
	return (byte)(QPixelFormat_AlphaSize(this.h))
}

func (this *QPixelFormat) BitsPerPixel() byte {
	return (byte)(QPixelFormat_BitsPerPixel(this.h))
}

func (this *QPixelFormat) AlphaUsage() AlphaUsage {
	xxxxxxxxx
}

func (this *QPixelFormat) AlphaPosition() AlphaPosition {
	xxxxxxxxx
}

func (this *QPixelFormat) Premultiplied() AlphaPremultiplied {
	xxxxxxxxx
}

func (this *QPixelFormat) TypeInterpretation() TypeInterpretation {
	xxxxxxxxx
}

func (this *QPixelFormat) ByteOrder() ByteOrder {
	xxxxxxxxx
}

func (this *QPixelFormat) YuvLayout() YUVLayout {
	xxxxxxxxx
}

func (this *QPixelFormat) SubEnum() byte {
	return (byte)(QPixelFormat_SubEnum(this.h))
}
