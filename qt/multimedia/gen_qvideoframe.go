package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QVideoFrame__HandleType int

const (
	QVideoFrame__NoHandle         QVideoFrame__HandleType = 0
	QVideoFrame__RhiTextureHandle QVideoFrame__HandleType = 1
)

type QVideoFrame__MapMode int

const (
	QVideoFrame__NotMapped QVideoFrame__MapMode = 0
	QVideoFrame__ReadOnly  QVideoFrame__MapMode = 1
	QVideoFrame__WriteOnly QVideoFrame__MapMode = 2
	QVideoFrame__ReadWrite QVideoFrame__MapMode = 3
)

type QVideoFrame__RotationAngle int

const (
	QVideoFrame__Rotation0   QVideoFrame__RotationAngle = 0
	QVideoFrame__Rotation90  QVideoFrame__RotationAngle = 90
	QVideoFrame__Rotation180 QVideoFrame__RotationAngle = 180
	QVideoFrame__Rotation270 QVideoFrame__RotationAngle = 270
)

type QVideoFrame__PaintOptions__PaintFlag int

const (
	QVideoFrame__PaintOptions__DontDrawSubtitles QVideoFrame__PaintOptions__PaintFlag = 1
)

type QVideoFrame struct {
	h          uintptr
	isSubclass bool
}

// NewQVideoFrame constructs a new QVideoFrame object.
func NewQVideoFrame() *QVideoFrame {
	ret := newQVideoFrame(QVideoFrame_new())
	ret.isSubclass = true
	return ret
}

// NewQVideoFrame2 constructs a new QVideoFrame object.
func NewQVideoFrame2(format *QVideoFrameFormat) *QVideoFrame {
	ret := newQVideoFrame(QVideoFrame_new2(format.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVideoFrame3 constructs a new QVideoFrame object.
func NewQVideoFrame3(image *qt.QImage) *QVideoFrame {
	ret := newQVideoFrame(QVideoFrame_new3((*QImage)(image.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQVideoFrame4 constructs a new QVideoFrame object.
func NewQVideoFrame4(other *QVideoFrame) *QVideoFrame {
	ret := newQVideoFrame(QVideoFrame_new4(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QVideoFrame) Swap(other *QVideoFrame) {
	QVideoFrame_Swap(this.h, other.cPointer())
}

func (this *QVideoFrame) OperatorAssign(other *QVideoFrame) {
	QVideoFrame_OperatorAssign(this.h, other.cPointer())
}

func (this *QVideoFrame) OperatorEqual(other *QVideoFrame) bool {
	return (bool)(QVideoFrame_OperatorEqual(this.h, other.cPointer()))
}

func (this *QVideoFrame) OperatorNotEqual(other *QVideoFrame) bool {
	return (bool)(QVideoFrame_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QVideoFrame) IsValid() bool {
	return (bool)(QVideoFrame_IsValid(this.h))
}

func (this *QVideoFrame) PixelFormat() QVideoFrameFormat__PixelFormat {
	return (QVideoFrameFormat__PixelFormat)(QVideoFrame_PixelFormat(this.h))
}

func (this *QVideoFrame) SurfaceFormat() *QVideoFrameFormat {
	_goptr := newQVideoFrameFormat(QVideoFrame_SurfaceFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVideoFrame) HandleType() QVideoFrame__HandleType {
	return (QVideoFrame__HandleType)(QVideoFrame_HandleType(this.h))
}

func (this *QVideoFrame) Size() *qt.QSize {
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QVideoFrame_Size(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVideoFrame) Width() int {
	return (int)(QVideoFrame_Width(this.h))
}

func (this *QVideoFrame) Height() int {
	return (int)(QVideoFrame_Height(this.h))
}

func (this *QVideoFrame) IsMapped() bool {
	return (bool)(QVideoFrame_IsMapped(this.h))
}

func (this *QVideoFrame) IsReadable() bool {
	return (bool)(QVideoFrame_IsReadable(this.h))
}

func (this *QVideoFrame) IsWritable() bool {
	return (bool)(QVideoFrame_IsWritable(this.h))
}

func (this *QVideoFrame) MapMode() QVideoFrame__MapMode {
	return (QVideoFrame__MapMode)(QVideoFrame_MapMode(this.h))
}

func (this *QVideoFrame) Map(mode QVideoFrame__MapMode) bool {
	return (bool)(QVideoFrame_Map(this.h, (int)(mode)))
}

func (this *QVideoFrame) Unmap() {
	QVideoFrame_Unmap(this.h)
}

func (this *QVideoFrame) BytesPerLine(plane int) int {
	return (int)(QVideoFrame_BytesPerLine(this.h, (int)(plane)))
}

func (this *QVideoFrame) Bits(plane int) *byte {
	return (*byte)(unsafe.Pointer(QVideoFrame_Bits(this.h, (int)(plane))))
}

func (this *QVideoFrame) BitsWithPlane(plane int) *byte {
	return (*byte)(unsafe.Pointer(QVideoFrame_BitsWithPlane(this.h, (int)(plane))))
}

func (this *QVideoFrame) MappedBytes(plane int) int {
	return (int)(QVideoFrame_MappedBytes(this.h, (int)(plane)))
}

func (this *QVideoFrame) PlaneCount() int {
	return (int)(QVideoFrame_PlaneCount(this.h))
}

func (this *QVideoFrame) StartTime() int64 {
	return (int64)(QVideoFrame_StartTime(this.h))
}

func (this *QVideoFrame) SetStartTime(time int64) {
	QVideoFrame_SetStartTime(this.h, (longlong)(time))
}

func (this *QVideoFrame) EndTime() int64 {
	return (int64)(QVideoFrame_EndTime(this.h))
}

func (this *QVideoFrame) SetEndTime(time int64) {
	QVideoFrame_SetEndTime(this.h, (longlong)(time))
}

func (this *QVideoFrame) SetRotationAngle(angle RotationAngle) {
	QVideoFrame_SetRotationAngle(this.h, angle)
}

func (this *QVideoFrame) RotationAngle() RotationAngle {
	xxxxxxxxx
}

func (this *QVideoFrame) SetRotation(angle QtVideo__Rotation) {
	QVideoFrame_SetRotation(this.h, (int)(angle))
}

func (this *QVideoFrame) Rotation() QtVideo__Rotation {
	return (QtVideo__Rotation)(QVideoFrame_Rotation(this.h))
}

func (this *QVideoFrame) SetMirrored(mirrored bool) {
	QVideoFrame_SetMirrored(this.h, (bool)(mirrored))
}

func (this *QVideoFrame) Mirrored() bool {
	return (bool)(QVideoFrame_Mirrored(this.h))
}

func (this *QVideoFrame) SetStreamFrameRate(rate float64) {
	QVideoFrame_SetStreamFrameRate(this.h, (double)(rate))
}

func (this *QVideoFrame) StreamFrameRate() float64 {
	return (float64)(QVideoFrame_StreamFrameRate(this.h))
}

func (this *QVideoFrame) ToImage() *qt.QImage {
	_goptr := qt.UnsafeNewQImage(unsafe.Pointer(QVideoFrame_ToImage(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVideoFrame) SubtitleText() string {
	var _ms struct_miqt_string = QVideoFrame_SubtitleText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVideoFrame) SetSubtitleText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QVideoFrame_SetSubtitleText(this.h, text_ms)
}

func (this *QVideoFrame) Paint(painter *qt.QPainter, rect *qt.QRectF, options *PaintOptions) {
	QVideoFrame_Paint(this.h, (*QPainter)(painter.UnsafePointer()), (*QRectF)(rect.UnsafePointer()), options)
}

type QVideoFrame__PaintOptions struct {
	h          uintptr
	isSubclass bool
}
