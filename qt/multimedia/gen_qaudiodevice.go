package multimedia

import (
	"unsafe"
)

type QAudioDevice__Mode int

const (
	QAudioDevice__Null   QAudioDevice__Mode = 0
	QAudioDevice__Input  QAudioDevice__Mode = 1
	QAudioDevice__Output QAudioDevice__Mode = 2
)

type QAudioDevice struct {
	h          uintptr
	isSubclass bool
}

// NewQAudioDevice constructs a new QAudioDevice object.
func NewQAudioDevice() *QAudioDevice {
	g := newQAudioDevice(QAudioDevice_new())
	g.isSubclass = true
	return g
}

// NewQAudioDevice2 constructs a new QAudioDevice object.
func NewQAudioDevice2(other *QAudioDevice) *QAudioDevice {
	g := newQAudioDevice(QAudioDevice_new2(other.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QAudioDevice) Swap(other *QAudioDevice) {
	QAudioDevice_Swap(this.h, other.cPointer())
}

func (this *QAudioDevice) OperatorAssign(other *QAudioDevice) {
	QAudioDevice_OperatorAssign(this.h, other.cPointer())
}

func (this *QAudioDevice) OperatorEqual(other *QAudioDevice) bool {
	return (bool)(QAudioDevice_OperatorEqual(this.h, other.cPointer()))
}

func (this *QAudioDevice) OperatorNotEqual(other *QAudioDevice) bool {
	return (bool)(QAudioDevice_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QAudioDevice) IsNull() bool {
	return (bool)(QAudioDevice_IsNull(this.h))
}

func (this *QAudioDevice) Id() []byte {
	var _bytearray struct_miqt_string = QAudioDevice_Id(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QAudioDevice) Description() string {
	var _ms struct_miqt_string = QAudioDevice_Description(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioDevice) IsDefault() bool {
	return (bool)(QAudioDevice_IsDefault(this.h))
}

func (this *QAudioDevice) Mode() QAudioDevice__Mode {
	return (QAudioDevice__Mode)(QAudioDevice_Mode(this.h))
}

func (this *QAudioDevice) IsFormatSupported(format *QAudioFormat) bool {
	return (bool)(QAudioDevice_IsFormatSupported(this.h, format.cPointer()))
}

func (this *QAudioDevice) PreferredFormat() *QAudioFormat {
	_goptr := newQAudioFormat(QAudioDevice_PreferredFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioDevice) MinimumSampleRate() int {
	return (int)(QAudioDevice_MinimumSampleRate(this.h))
}

func (this *QAudioDevice) MaximumSampleRate() int {
	return (int)(QAudioDevice_MaximumSampleRate(this.h))
}

func (this *QAudioDevice) MinimumChannelCount() int {
	return (int)(QAudioDevice_MinimumChannelCount(this.h))
}

func (this *QAudioDevice) MaximumChannelCount() int {
	return (int)(QAudioDevice_MaximumChannelCount(this.h))
}

func (this *QAudioDevice) SupportedSampleFormats() []QAudioFormat__SampleFormat {
	var _ma struct_miqt_array = QAudioDevice_SupportedSampleFormats(this.h)
	_ret := make([]QAudioFormat__SampleFormat, int(_ma.len))
	_outCast := (*[0xffff]uint16_t)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (QAudioFormat__SampleFormat)(_outCast[i])
	}
	return _ret
}

func (this *QAudioDevice) ChannelConfiguration() QAudioFormat__ChannelConfig {
	return (QAudioFormat__ChannelConfig)(QAudioDevice_ChannelConfiguration(this.h))
}
