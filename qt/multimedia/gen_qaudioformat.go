package multimedia

import (
	"unsafe"
)

type QAudioFormat__SampleFormat uint16

const (
	QAudioFormat__Unknown        QAudioFormat__SampleFormat = 0
	QAudioFormat__UInt8          QAudioFormat__SampleFormat = 1
	QAudioFormat__Int16          QAudioFormat__SampleFormat = 2
	QAudioFormat__Int32          QAudioFormat__SampleFormat = 3
	QAudioFormat__Float          QAudioFormat__SampleFormat = 4
	QAudioFormat__NSampleFormats QAudioFormat__SampleFormat = 5
)

type QAudioFormat__AudioChannelPosition int

const (
	QAudioFormat__UnknownPosition    QAudioFormat__AudioChannelPosition = 0
	QAudioFormat__FrontLeft          QAudioFormat__AudioChannelPosition = 1
	QAudioFormat__FrontRight         QAudioFormat__AudioChannelPosition = 2
	QAudioFormat__FrontCenter        QAudioFormat__AudioChannelPosition = 3
	QAudioFormat__LFE                QAudioFormat__AudioChannelPosition = 4
	QAudioFormat__BackLeft           QAudioFormat__AudioChannelPosition = 5
	QAudioFormat__BackRight          QAudioFormat__AudioChannelPosition = 6
	QAudioFormat__FrontLeftOfCenter  QAudioFormat__AudioChannelPosition = 7
	QAudioFormat__FrontRightOfCenter QAudioFormat__AudioChannelPosition = 8
	QAudioFormat__BackCenter         QAudioFormat__AudioChannelPosition = 9
	QAudioFormat__SideLeft           QAudioFormat__AudioChannelPosition = 10
	QAudioFormat__SideRight          QAudioFormat__AudioChannelPosition = 11
	QAudioFormat__TopCenter          QAudioFormat__AudioChannelPosition = 12
	QAudioFormat__TopFrontLeft       QAudioFormat__AudioChannelPosition = 13
	QAudioFormat__TopFrontCenter     QAudioFormat__AudioChannelPosition = 14
	QAudioFormat__TopFrontRight      QAudioFormat__AudioChannelPosition = 15
	QAudioFormat__TopBackLeft        QAudioFormat__AudioChannelPosition = 16
	QAudioFormat__TopBackCenter      QAudioFormat__AudioChannelPosition = 17
	QAudioFormat__TopBackRight       QAudioFormat__AudioChannelPosition = 18
	QAudioFormat__LFE2               QAudioFormat__AudioChannelPosition = 19
	QAudioFormat__TopSideLeft        QAudioFormat__AudioChannelPosition = 20
	QAudioFormat__TopSideRight       QAudioFormat__AudioChannelPosition = 21
	QAudioFormat__BottomFrontCenter  QAudioFormat__AudioChannelPosition = 22
	QAudioFormat__BottomFrontLeft    QAudioFormat__AudioChannelPosition = 23
	QAudioFormat__BottomFrontRight   QAudioFormat__AudioChannelPosition = 24
)

type QAudioFormat__ChannelConfig uint

const (
	QAudioFormat__ChannelConfigUnknown       QAudioFormat__ChannelConfig = 0
	QAudioFormat__ChannelConfigMono          QAudioFormat__ChannelConfig = 8
	QAudioFormat__ChannelConfigStereo        QAudioFormat__ChannelConfig = 6
	QAudioFormat__ChannelConfig2Dot1         QAudioFormat__ChannelConfig = 22
	QAudioFormat__ChannelConfig3Dot0         QAudioFormat__ChannelConfig = 14
	QAudioFormat__ChannelConfig3Dot1         QAudioFormat__ChannelConfig = 30
	QAudioFormat__ChannelConfigSurround5Dot0 QAudioFormat__ChannelConfig = 110
	QAudioFormat__ChannelConfigSurround5Dot1 QAudioFormat__ChannelConfig = 126
	QAudioFormat__ChannelConfigSurround7Dot0 QAudioFormat__ChannelConfig = 3182
	QAudioFormat__ChannelConfigSurround7Dot1 QAudioFormat__ChannelConfig = 3198
)

type QAudioFormat struct {
	h          uintptr
	isSubclass bool
}

// NewQAudioFormat constructs a new QAudioFormat object.
func NewQAudioFormat() *QAudioFormat {
	g := newQAudioFormat(QAudioFormat_new())
	g.isSubclass = true
	return g
}

// NewQAudioFormat2 constructs a new QAudioFormat object.
func NewQAudioFormat2(param1 *QAudioFormat) *QAudioFormat {
	g := newQAudioFormat(QAudioFormat_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QAudioFormat) IsValid() bool {
	return (bool)(QAudioFormat_IsValid(this.h))
}

func (this *QAudioFormat) SetSampleRate(sampleRate int) {
	QAudioFormat_SetSampleRate(this.h, (int)(sampleRate))
}

func (this *QAudioFormat) SampleRate() int {
	return (int)(QAudioFormat_SampleRate(this.h))
}

func (this *QAudioFormat) SetChannelConfig(config ChannelConfig) {
	QAudioFormat_SetChannelConfig(this.h, config)
}

func (this *QAudioFormat) ChannelConfig() ChannelConfig {
	xxxxxxxxx
}

func (this *QAudioFormat) SetChannelCount(channelCount int) {
	QAudioFormat_SetChannelCount(this.h, (int)(channelCount))
}

func (this *QAudioFormat) ChannelCount() int {
	return (int)(QAudioFormat_ChannelCount(this.h))
}

func (this *QAudioFormat) ChannelOffset(channel AudioChannelPosition) int {
	return (int)(QAudioFormat_ChannelOffset(this.h, channel))
}

func (this *QAudioFormat) SetSampleFormat(f SampleFormat) {
	QAudioFormat_SetSampleFormat(this.h, f)
}

func (this *QAudioFormat) SampleFormat() SampleFormat {
	xxxxxxxxx
}

func (this *QAudioFormat) BytesForDuration(microseconds int64) int {
	return (int)(QAudioFormat_BytesForDuration(this.h, (longlong)(microseconds)))
}

func (this *QAudioFormat) DurationForBytes(byteCount int) int64 {
	return (int64)(QAudioFormat_DurationForBytes(this.h, (int)(byteCount)))
}

func (this *QAudioFormat) BytesForFrames(frameCount int) int {
	return (int)(QAudioFormat_BytesForFrames(this.h, (int)(frameCount)))
}

func (this *QAudioFormat) FramesForBytes(byteCount int) int {
	return (int)(QAudioFormat_FramesForBytes(this.h, (int)(byteCount)))
}

func (this *QAudioFormat) FramesForDuration(microseconds int64) int {
	return (int)(QAudioFormat_FramesForDuration(this.h, (longlong)(microseconds)))
}

func (this *QAudioFormat) DurationForFrames(frameCount int) int64 {
	return (int64)(QAudioFormat_DurationForFrames(this.h, (int)(frameCount)))
}

func (this *QAudioFormat) BytesPerFrame() int {
	return (int)(QAudioFormat_BytesPerFrame(this.h))
}

func (this *QAudioFormat) BytesPerSample() int {
	return (int)(QAudioFormat_BytesPerSample(this.h))
}

func (this *QAudioFormat) NormalizedSampleValue(sample unsafe.Pointer) float32 {
	return (float32)(QAudioFormat_NormalizedSampleValue(this.h, sample))
}

func QAudioFormat_DefaultChannelConfigForChannelCount(channelCount int) ChannelConfig {
	xxxxxxxxx
}
