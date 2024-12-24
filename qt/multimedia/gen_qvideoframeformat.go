package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QVideoFrameFormat__PixelFormat int

const (
	QVideoFrameFormat__Format_Invalid                QVideoFrameFormat__PixelFormat = 0
	QVideoFrameFormat__Format_ARGB8888               QVideoFrameFormat__PixelFormat = 1
	QVideoFrameFormat__Format_ARGB8888_Premultiplied QVideoFrameFormat__PixelFormat = 2
	QVideoFrameFormat__Format_XRGB8888               QVideoFrameFormat__PixelFormat = 3
	QVideoFrameFormat__Format_BGRA8888               QVideoFrameFormat__PixelFormat = 4
	QVideoFrameFormat__Format_BGRA8888_Premultiplied QVideoFrameFormat__PixelFormat = 5
	QVideoFrameFormat__Format_BGRX8888               QVideoFrameFormat__PixelFormat = 6
	QVideoFrameFormat__Format_ABGR8888               QVideoFrameFormat__PixelFormat = 7
	QVideoFrameFormat__Format_XBGR8888               QVideoFrameFormat__PixelFormat = 8
	QVideoFrameFormat__Format_RGBA8888               QVideoFrameFormat__PixelFormat = 9
	QVideoFrameFormat__Format_RGBX8888               QVideoFrameFormat__PixelFormat = 10
	QVideoFrameFormat__Format_AYUV                   QVideoFrameFormat__PixelFormat = 11
	QVideoFrameFormat__Format_AYUV_Premultiplied     QVideoFrameFormat__PixelFormat = 12
	QVideoFrameFormat__Format_YUV420P                QVideoFrameFormat__PixelFormat = 13
	QVideoFrameFormat__Format_YUV422P                QVideoFrameFormat__PixelFormat = 14
	QVideoFrameFormat__Format_YV12                   QVideoFrameFormat__PixelFormat = 15
	QVideoFrameFormat__Format_UYVY                   QVideoFrameFormat__PixelFormat = 16
	QVideoFrameFormat__Format_YUYV                   QVideoFrameFormat__PixelFormat = 17
	QVideoFrameFormat__Format_NV12                   QVideoFrameFormat__PixelFormat = 18
	QVideoFrameFormat__Format_NV21                   QVideoFrameFormat__PixelFormat = 19
	QVideoFrameFormat__Format_IMC1                   QVideoFrameFormat__PixelFormat = 20
	QVideoFrameFormat__Format_IMC2                   QVideoFrameFormat__PixelFormat = 21
	QVideoFrameFormat__Format_IMC3                   QVideoFrameFormat__PixelFormat = 22
	QVideoFrameFormat__Format_IMC4                   QVideoFrameFormat__PixelFormat = 23
	QVideoFrameFormat__Format_Y8                     QVideoFrameFormat__PixelFormat = 24
	QVideoFrameFormat__Format_Y16                    QVideoFrameFormat__PixelFormat = 25
	QVideoFrameFormat__Format_P010                   QVideoFrameFormat__PixelFormat = 26
	QVideoFrameFormat__Format_P016                   QVideoFrameFormat__PixelFormat = 27
	QVideoFrameFormat__Format_SamplerExternalOES     QVideoFrameFormat__PixelFormat = 28
	QVideoFrameFormat__Format_Jpeg                   QVideoFrameFormat__PixelFormat = 29
	QVideoFrameFormat__Format_SamplerRect            QVideoFrameFormat__PixelFormat = 30
	QVideoFrameFormat__Format_YUV420P10              QVideoFrameFormat__PixelFormat = 31
)

type QVideoFrameFormat__Direction int

const (
	QVideoFrameFormat__TopToBottom QVideoFrameFormat__Direction = 0
	QVideoFrameFormat__BottomToTop QVideoFrameFormat__Direction = 1
)

type QVideoFrameFormat__YCbCrColorSpace int

const (
	QVideoFrameFormat__YCbCr_Undefined QVideoFrameFormat__YCbCrColorSpace = 0
	QVideoFrameFormat__YCbCr_BT601     QVideoFrameFormat__YCbCrColorSpace = 1
	QVideoFrameFormat__YCbCr_BT709     QVideoFrameFormat__YCbCrColorSpace = 2
	QVideoFrameFormat__YCbCr_xvYCC601  QVideoFrameFormat__YCbCrColorSpace = 3
	QVideoFrameFormat__YCbCr_xvYCC709  QVideoFrameFormat__YCbCrColorSpace = 4
	QVideoFrameFormat__YCbCr_JPEG      QVideoFrameFormat__YCbCrColorSpace = 5
	QVideoFrameFormat__YCbCr_BT2020    QVideoFrameFormat__YCbCrColorSpace = 6
)

type QVideoFrameFormat__ColorSpace int

const (
	QVideoFrameFormat__ColorSpace_Undefined QVideoFrameFormat__ColorSpace = 0
	QVideoFrameFormat__ColorSpace_BT601     QVideoFrameFormat__ColorSpace = 1
	QVideoFrameFormat__ColorSpace_BT709     QVideoFrameFormat__ColorSpace = 2
	QVideoFrameFormat__ColorSpace_AdobeRgb  QVideoFrameFormat__ColorSpace = 5
	QVideoFrameFormat__ColorSpace_BT2020    QVideoFrameFormat__ColorSpace = 6
)

type QVideoFrameFormat__ColorTransfer int

const (
	QVideoFrameFormat__ColorTransfer_Unknown QVideoFrameFormat__ColorTransfer = 0
	QVideoFrameFormat__ColorTransfer_BT709   QVideoFrameFormat__ColorTransfer = 1
	QVideoFrameFormat__ColorTransfer_BT601   QVideoFrameFormat__ColorTransfer = 2
	QVideoFrameFormat__ColorTransfer_Linear  QVideoFrameFormat__ColorTransfer = 3
	QVideoFrameFormat__ColorTransfer_Gamma22 QVideoFrameFormat__ColorTransfer = 4
	QVideoFrameFormat__ColorTransfer_Gamma28 QVideoFrameFormat__ColorTransfer = 5
	QVideoFrameFormat__ColorTransfer_ST2084  QVideoFrameFormat__ColorTransfer = 6
	QVideoFrameFormat__ColorTransfer_STD_B67 QVideoFrameFormat__ColorTransfer = 7
)

type QVideoFrameFormat__ColorRange int

const (
	QVideoFrameFormat__ColorRange_Unknown QVideoFrameFormat__ColorRange = 0
	QVideoFrameFormat__ColorRange_Video   QVideoFrameFormat__ColorRange = 1
	QVideoFrameFormat__ColorRange_Full    QVideoFrameFormat__ColorRange = 2
)

type QVideoFrameFormat struct {
	h          uintptr
	isSubclass bool
}

// NewQVideoFrameFormat constructs a new QVideoFrameFormat object.
func NewQVideoFrameFormat() *QVideoFrameFormat {
	g := newQVideoFrameFormat(QVideoFrameFormat_new())
	g.isSubclass = true
	return g
}

// NewQVideoFrameFormat2 constructs a new QVideoFrameFormat object.
func NewQVideoFrameFormat2(size *qt.QSize, pixelFormat PixelFormat) *QVideoFrameFormat {
	g := newQVideoFrameFormat(QVideoFrameFormat_new2((*QSize)(size.UnsafePointer()), pixelFormat))
	g.isSubclass = true
	return g
}

// NewQVideoFrameFormat3 constructs a new QVideoFrameFormat object.
func NewQVideoFrameFormat3(format *QVideoFrameFormat) *QVideoFrameFormat {
	g := newQVideoFrameFormat(QVideoFrameFormat_new3(format.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QVideoFrameFormat) Swap(other *QVideoFrameFormat) {
	QVideoFrameFormat_Swap(this.h, other.cPointer())
}

func (this *QVideoFrameFormat) Detach() {
	QVideoFrameFormat_Detach(this.h)
}

func (this *QVideoFrameFormat) OperatorAssign(format *QVideoFrameFormat) {
	QVideoFrameFormat_OperatorAssign(this.h, format.cPointer())
}

func (this *QVideoFrameFormat) OperatorEqual(format *QVideoFrameFormat) bool {
	return (bool)(QVideoFrameFormat_OperatorEqual(this.h, format.cPointer()))
}

func (this *QVideoFrameFormat) OperatorNotEqual(format *QVideoFrameFormat) bool {
	return (bool)(QVideoFrameFormat_OperatorNotEqual(this.h, format.cPointer()))
}

func (this *QVideoFrameFormat) IsValid() bool {
	return (bool)(QVideoFrameFormat_IsValid(this.h))
}

func (this *QVideoFrameFormat) PixelFormat() QVideoFrameFormat__PixelFormat {
	return (QVideoFrameFormat__PixelFormat)(QVideoFrameFormat_PixelFormat(this.h))
}

func (this *QVideoFrameFormat) FrameSize() *qt.QSize {
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QVideoFrameFormat_FrameSize(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVideoFrameFormat) SetFrameSize(size *qt.QSize) {
	QVideoFrameFormat_SetFrameSize(this.h, (*QSize)(size.UnsafePointer()))
}

func (this *QVideoFrameFormat) SetFrameSize2(width int, height int) {
	QVideoFrameFormat_SetFrameSize2(this.h, (int)(width), (int)(height))
}

func (this *QVideoFrameFormat) FrameWidth() int {
	return (int)(QVideoFrameFormat_FrameWidth(this.h))
}

func (this *QVideoFrameFormat) FrameHeight() int {
	return (int)(QVideoFrameFormat_FrameHeight(this.h))
}

func (this *QVideoFrameFormat) PlaneCount() int {
	return (int)(QVideoFrameFormat_PlaneCount(this.h))
}

func (this *QVideoFrameFormat) Viewport() *qt.QRect {
	_goptr := qt.UnsafeNewQRect(unsafe.Pointer(QVideoFrameFormat_Viewport(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVideoFrameFormat) SetViewport(viewport *qt.QRect) {
	QVideoFrameFormat_SetViewport(this.h, (*QRect)(viewport.UnsafePointer()))
}

func (this *QVideoFrameFormat) ScanLineDirection() Direction {
	xxxxxxxxx
}

func (this *QVideoFrameFormat) SetScanLineDirection(direction Direction) {
	QVideoFrameFormat_SetScanLineDirection(this.h, direction)
}

func (this *QVideoFrameFormat) FrameRate() float64 {
	return (float64)(QVideoFrameFormat_FrameRate(this.h))
}

func (this *QVideoFrameFormat) SetFrameRate(rate float64) {
	QVideoFrameFormat_SetFrameRate(this.h, (double)(rate))
}

func (this *QVideoFrameFormat) StreamFrameRate() float64 {
	return (float64)(QVideoFrameFormat_StreamFrameRate(this.h))
}

func (this *QVideoFrameFormat) SetStreamFrameRate(rate float64) {
	QVideoFrameFormat_SetStreamFrameRate(this.h, (double)(rate))
}

func (this *QVideoFrameFormat) YCbCrColorSpace() YCbCrColorSpace {
	xxxxxxxxx
}

func (this *QVideoFrameFormat) SetYCbCrColorSpace(colorSpace YCbCrColorSpace) {
	QVideoFrameFormat_SetYCbCrColorSpace(this.h, colorSpace)
}

func (this *QVideoFrameFormat) ColorSpace() ColorSpace {
	xxxxxxxxx
}

func (this *QVideoFrameFormat) SetColorSpace(colorSpace ColorSpace) {
	QVideoFrameFormat_SetColorSpace(this.h, colorSpace)
}

func (this *QVideoFrameFormat) ColorTransfer() ColorTransfer {
	xxxxxxxxx
}

func (this *QVideoFrameFormat) SetColorTransfer(colorTransfer ColorTransfer) {
	QVideoFrameFormat_SetColorTransfer(this.h, colorTransfer)
}

func (this *QVideoFrameFormat) ColorRange() ColorRange {
	xxxxxxxxx
}

func (this *QVideoFrameFormat) SetColorRange(rangeVal ColorRange) {
	QVideoFrameFormat_SetColorRange(this.h, rangeVal)
}

func (this *QVideoFrameFormat) IsMirrored() bool {
	return (bool)(QVideoFrameFormat_IsMirrored(this.h))
}

func (this *QVideoFrameFormat) SetMirrored(mirrored bool) {
	QVideoFrameFormat_SetMirrored(this.h, (bool)(mirrored))
}

func (this *QVideoFrameFormat) Rotation() QtVideo__Rotation {
	return (QtVideo__Rotation)(QVideoFrameFormat_Rotation(this.h))
}

func (this *QVideoFrameFormat) SetRotation(rotation QtVideo__Rotation) {
	QVideoFrameFormat_SetRotation(this.h, (int)(rotation))
}

func (this *QVideoFrameFormat) VertexShaderFileName() string {
	var _ms struct_miqt_string = QVideoFrameFormat_VertexShaderFileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVideoFrameFormat) FragmentShaderFileName() string {
	var _ms struct_miqt_string = QVideoFrameFormat_FragmentShaderFileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVideoFrameFormat) MaxLuminance() float32 {
	return (float32)(QVideoFrameFormat_MaxLuminance(this.h))
}

func (this *QVideoFrameFormat) SetMaxLuminance(lum float32) {
	QVideoFrameFormat_SetMaxLuminance(this.h, (float)(lum))
}

func QVideoFrameFormat_PixelFormatFromImageFormat(format qt.QImage__Format) PixelFormat {
	xxxxxxxxx
}

func QVideoFrameFormat_ImageFormatFromPixelFormat(format PixelFormat) qt.QImage__Format {
	return (qt.QImage__Format)(QVideoFrameFormat_ImageFormatFromPixelFormat(format))
}

func QVideoFrameFormat_PixelFormatToString(pixelFormat QVideoFrameFormat__PixelFormat) string {
	var _ms struct_miqt_string = QVideoFrameFormat_PixelFormatToString((int)(pixelFormat))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
