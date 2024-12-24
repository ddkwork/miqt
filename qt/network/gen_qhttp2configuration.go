package network

import (
	"unsafe"
)

type QHttp2Configuration struct {
	h          uintptr
	isSubclass bool
}

// NewQHttp2Configuration constructs a new QHttp2Configuration object.
func NewQHttp2Configuration() *QHttp2Configuration {
	g := newQHttp2Configuration(QHttp2Configuration_new())
	g.isSubclass = true
	return g
}

// NewQHttp2Configuration2 constructs a new QHttp2Configuration object.
func NewQHttp2Configuration2(other *QHttp2Configuration) *QHttp2Configuration {
	g := newQHttp2Configuration(QHttp2Configuration_new2(other.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QHttp2Configuration) OperatorAssign(other *QHttp2Configuration) {
	QHttp2Configuration_OperatorAssign(this.h, other.cPointer())
}

func (this *QHttp2Configuration) SetServerPushEnabled(enable bool) {
	QHttp2Configuration_SetServerPushEnabled(this.h, (bool)(enable))
}

func (this *QHttp2Configuration) ServerPushEnabled() bool {
	return (bool)(QHttp2Configuration_ServerPushEnabled(this.h))
}

func (this *QHttp2Configuration) SetHuffmanCompressionEnabled(enable bool) {
	QHttp2Configuration_SetHuffmanCompressionEnabled(this.h, (bool)(enable))
}

func (this *QHttp2Configuration) HuffmanCompressionEnabled() bool {
	return (bool)(QHttp2Configuration_HuffmanCompressionEnabled(this.h))
}

func (this *QHttp2Configuration) SetSessionReceiveWindowSize(size uint) bool {
	return (bool)(QHttp2Configuration_SetSessionReceiveWindowSize(this.h, (uint)(size)))
}

func (this *QHttp2Configuration) SessionReceiveWindowSize() uint {
	return (uint)(QHttp2Configuration_SessionReceiveWindowSize(this.h))
}

func (this *QHttp2Configuration) SetStreamReceiveWindowSize(size uint) bool {
	return (bool)(QHttp2Configuration_SetStreamReceiveWindowSize(this.h, (uint)(size)))
}

func (this *QHttp2Configuration) StreamReceiveWindowSize() uint {
	return (uint)(QHttp2Configuration_StreamReceiveWindowSize(this.h))
}

func (this *QHttp2Configuration) SetMaxFrameSize(size uint) bool {
	return (bool)(QHttp2Configuration_SetMaxFrameSize(this.h, (uint)(size)))
}

func (this *QHttp2Configuration) MaxFrameSize() uint {
	return (uint)(QHttp2Configuration_MaxFrameSize(this.h))
}

func (this *QHttp2Configuration) SetMaxConcurrentStreams(value uint) {
	QHttp2Configuration_SetMaxConcurrentStreams(this.h, (uint)(value))
}

func (this *QHttp2Configuration) MaxConcurrentStreams() uint {
	return (uint)(QHttp2Configuration_MaxConcurrentStreams(this.h))
}

func (this *QHttp2Configuration) Swap(other *QHttp2Configuration) {
	QHttp2Configuration_Swap(this.h, other.cPointer())
}
