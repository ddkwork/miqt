package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QCameraDevice__Position int

const (
	QCameraDevice__UnspecifiedPosition QCameraDevice__Position = 0
	QCameraDevice__BackFace            QCameraDevice__Position = 1
	QCameraDevice__FrontFace           QCameraDevice__Position = 2
)

type QCameraFormat struct {
	h          uintptr
	isSubclass bool
}

// NewQCameraFormat constructs a new QCameraFormat object.
func NewQCameraFormat() *QCameraFormat {
	ret := newQCameraFormat(QCameraFormat_new())
	ret.isSubclass = true
	return ret
}

// NewQCameraFormat2 constructs a new QCameraFormat object.
func NewQCameraFormat2(other *QCameraFormat) *QCameraFormat {
	ret := newQCameraFormat(QCameraFormat_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QCameraFormat) OperatorAssign(other *QCameraFormat) {
	QCameraFormat_OperatorAssign(this.h, other.cPointer())
}

func (this *QCameraFormat) PixelFormat() QVideoFrameFormat__PixelFormat {
	return (QVideoFrameFormat__PixelFormat)(QCameraFormat_PixelFormat(this.h))
}

func (this *QCameraFormat) Resolution() *qt.QSize {
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QCameraFormat_Resolution(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCameraFormat) MinFrameRate() float32 {
	return (float32)(QCameraFormat_MinFrameRate(this.h))
}

func (this *QCameraFormat) MaxFrameRate() float32 {
	return (float32)(QCameraFormat_MaxFrameRate(this.h))
}

func (this *QCameraFormat) IsNull() bool {
	return (bool)(QCameraFormat_IsNull(this.h))
}

func (this *QCameraFormat) OperatorEqual(other *QCameraFormat) bool {
	return (bool)(QCameraFormat_OperatorEqual(this.h, other.cPointer()))
}

func (this *QCameraFormat) OperatorNotEqual(other *QCameraFormat) bool {
	return (bool)(QCameraFormat_OperatorNotEqual(this.h, other.cPointer()))
}

type QCameraDevice struct {
	h          uintptr
	isSubclass bool
}

// NewQCameraDevice constructs a new QCameraDevice object.
func NewQCameraDevice() *QCameraDevice {
	ret := newQCameraDevice(QCameraDevice_new())
	ret.isSubclass = true
	return ret
}

// NewQCameraDevice2 constructs a new QCameraDevice object.
func NewQCameraDevice2(other *QCameraDevice) *QCameraDevice {
	ret := newQCameraDevice(QCameraDevice_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QCameraDevice) OperatorAssign(other *QCameraDevice) {
	QCameraDevice_OperatorAssign(this.h, other.cPointer())
}

func (this *QCameraDevice) OperatorEqual(other *QCameraDevice) bool {
	return (bool)(QCameraDevice_OperatorEqual(this.h, other.cPointer()))
}

func (this *QCameraDevice) OperatorNotEqual(other *QCameraDevice) bool {
	return (bool)(QCameraDevice_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QCameraDevice) IsNull() bool {
	return (bool)(QCameraDevice_IsNull(this.h))
}

func (this *QCameraDevice) Id() []byte {
	var _bytearray struct_miqt_string = QCameraDevice_Id(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QCameraDevice) Description() string {
	var _ms struct_miqt_string = QCameraDevice_Description(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCameraDevice) IsDefault() bool {
	return (bool)(QCameraDevice_IsDefault(this.h))
}

func (this *QCameraDevice) Position() Position {
	xxxxxxxxx
}

func (this *QCameraDevice) PhotoResolutions() []qt.QSize {
	var _ma struct_miqt_array = QCameraDevice_PhotoResolutions(this.h)
	_ret := make([]qt.QSize, int(_ma.len))
	_outCast := (*[0xffff]*QSize)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := qt.UnsafeNewQSize(unsafe.Pointer(_outCast[i]))
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QCameraDevice) VideoFormats() []QCameraFormat {
	var _ma struct_miqt_array = QCameraDevice_VideoFormats(this.h)
	_ret := make([]QCameraFormat, int(_ma.len))
	_outCast := (*[0xffff]*QCameraFormat)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQCameraFormat(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QCameraDevice) CorrectionAngle() QtVideo__Rotation {
	return (QtVideo__Rotation)(QCameraDevice_CorrectionAngle(this.h))
}
