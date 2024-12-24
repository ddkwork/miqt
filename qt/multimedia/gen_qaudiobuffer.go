package multimedia

import (
	"unsafe"
)

type QAudioBuffer struct {
	h          uintptr
	isSubclass bool
}

// NewQAudioBuffer constructs a new QAudioBuffer object.
func NewQAudioBuffer() *QAudioBuffer {
	g := newQAudioBuffer(QAudioBuffer_new())
	g.isSubclass = true
	return g
}

// NewQAudioBuffer2 constructs a new QAudioBuffer object.
func NewQAudioBuffer2(other *QAudioBuffer) *QAudioBuffer {
	g := newQAudioBuffer(QAudioBuffer_new2(other.cPointer()))
	g.isSubclass = true
	return g
}

// NewQAudioBuffer3 constructs a new QAudioBuffer object.
func NewQAudioBuffer3(data []byte, format *QAudioFormat) *QAudioBuffer {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	g := newQAudioBuffer(QAudioBuffer_new3(data_alias, format.cPointer()))
	g.isSubclass = true
	return g
}

// NewQAudioBuffer4 constructs a new QAudioBuffer object.
func NewQAudioBuffer4(numFrames int, format *QAudioFormat) *QAudioBuffer {
	g := newQAudioBuffer(QAudioBuffer_new4((int)(numFrames), format.cPointer()))
	g.isSubclass = true
	return g
}

// NewQAudioBuffer5 constructs a new QAudioBuffer object.
func NewQAudioBuffer5(data []byte, format *QAudioFormat, startTime int64) *QAudioBuffer {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	g := newQAudioBuffer(QAudioBuffer_new5(data_alias, format.cPointer(), (longlong)(startTime)))
	g.isSubclass = true
	return g
}

// NewQAudioBuffer6 constructs a new QAudioBuffer object.
func NewQAudioBuffer6(numFrames int, format *QAudioFormat, startTime int64) *QAudioBuffer {
	g := newQAudioBuffer(QAudioBuffer_new6((int)(numFrames), format.cPointer(), (longlong)(startTime)))
	g.isSubclass = true
	return g
}

func (this *QAudioBuffer) OperatorAssign(other *QAudioBuffer) {
	QAudioBuffer_OperatorAssign(this.h, other.cPointer())
}

func (this *QAudioBuffer) Swap(other *QAudioBuffer) {
	QAudioBuffer_Swap(this.h, other.cPointer())
}

func (this *QAudioBuffer) IsValid() bool {
	return (bool)(QAudioBuffer_IsValid(this.h))
}

func (this *QAudioBuffer) Detach() {
	QAudioBuffer_Detach(this.h)
}

func (this *QAudioBuffer) Format() *QAudioFormat {
	_goptr := newQAudioFormat(QAudioBuffer_Format(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioBuffer) FrameCount() int64 {
	return (int64)(QAudioBuffer_FrameCount(this.h))
}

func (this *QAudioBuffer) SampleCount() int64 {
	return (int64)(QAudioBuffer_SampleCount(this.h))
}

func (this *QAudioBuffer) ByteCount() int64 {
	return (int64)(QAudioBuffer_ByteCount(this.h))
}

func (this *QAudioBuffer) Duration() int64 {
	return (int64)(QAudioBuffer_Duration(this.h))
}

func (this *QAudioBuffer) StartTime() int64 {
	return (int64)(QAudioBuffer_StartTime(this.h))
}
