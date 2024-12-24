package network

import (
	"unsafe"
)

type QNetworkDatagram struct {
	h          uintptr
	isSubclass bool
}

// NewQNetworkDatagram constructs a new QNetworkDatagram object.
func NewQNetworkDatagram() *QNetworkDatagram {
	g := newQNetworkDatagram(QNetworkDatagram_new())
	g.isSubclass = true
	return g
}

// NewQNetworkDatagram2 constructs a new QNetworkDatagram object.
func NewQNetworkDatagram2(data []byte) *QNetworkDatagram {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	g := newQNetworkDatagram(QNetworkDatagram_new2(data_alias))
	g.isSubclass = true
	return g
}

// NewQNetworkDatagram3 constructs a new QNetworkDatagram object.
func NewQNetworkDatagram3(other *QNetworkDatagram) *QNetworkDatagram {
	g := newQNetworkDatagram(QNetworkDatagram_new3(other.cPointer()))
	g.isSubclass = true
	return g
}

// NewQNetworkDatagram4 constructs a new QNetworkDatagram object.
func NewQNetworkDatagram4(data []byte, destinationAddress *QHostAddress) *QNetworkDatagram {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	g := newQNetworkDatagram(QNetworkDatagram_new4(data_alias, destinationAddress.cPointer()))
	g.isSubclass = true
	return g
}

// NewQNetworkDatagram5 constructs a new QNetworkDatagram object.
func NewQNetworkDatagram5(data []byte, destinationAddress *QHostAddress, port uint16) *QNetworkDatagram {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	g := newQNetworkDatagram(QNetworkDatagram_new5(data_alias, destinationAddress.cPointer(), (uint16_t)(port)))
	g.isSubclass = true
	return g
}

func (this *QNetworkDatagram) OperatorAssign(other *QNetworkDatagram) {
	QNetworkDatagram_OperatorAssign(this.h, other.cPointer())
}

func (this *QNetworkDatagram) Swap(other *QNetworkDatagram) {
	QNetworkDatagram_Swap(this.h, other.cPointer())
}

func (this *QNetworkDatagram) Clear() {
	QNetworkDatagram_Clear(this.h)
}

func (this *QNetworkDatagram) IsValid() bool {
	return (bool)(QNetworkDatagram_IsValid(this.h))
}

func (this *QNetworkDatagram) IsNull() bool {
	return (bool)(QNetworkDatagram_IsNull(this.h))
}

func (this *QNetworkDatagram) InterfaceIndex() uint {
	return (uint)(QNetworkDatagram_InterfaceIndex(this.h))
}

func (this *QNetworkDatagram) SetInterfaceIndex(index uint) {
	QNetworkDatagram_SetInterfaceIndex(this.h, (uint)(index))
}

func (this *QNetworkDatagram) SenderAddress() *QHostAddress {
	_goptr := newQHostAddress(QNetworkDatagram_SenderAddress(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkDatagram) DestinationAddress() *QHostAddress {
	_goptr := newQHostAddress(QNetworkDatagram_DestinationAddress(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkDatagram) SenderPort() int {
	return (int)(QNetworkDatagram_SenderPort(this.h))
}

func (this *QNetworkDatagram) DestinationPort() int {
	return (int)(QNetworkDatagram_DestinationPort(this.h))
}

func (this *QNetworkDatagram) SetSender(address *QHostAddress) {
	QNetworkDatagram_SetSender(this.h, address.cPointer())
}

func (this *QNetworkDatagram) SetDestination(address *QHostAddress, port uint16) {
	QNetworkDatagram_SetDestination(this.h, address.cPointer(), (uint16_t)(port))
}

func (this *QNetworkDatagram) HopLimit() int {
	return (int)(QNetworkDatagram_HopLimit(this.h))
}

func (this *QNetworkDatagram) SetHopLimit(count int) {
	QNetworkDatagram_SetHopLimit(this.h, (int)(count))
}

func (this *QNetworkDatagram) Data() []byte {
	var _bytearray struct_miqt_string = QNetworkDatagram_Data(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QNetworkDatagram) SetData(data []byte) {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	QNetworkDatagram_SetData(this.h, data_alias)
}

func (this *QNetworkDatagram) MakeReply(payload []byte) *QNetworkDatagram {
	payload_alias := struct_miqt_string{}
	payload_alias.data = (char)(unsafe.Pointer(&payload[0]))
	payload_alias.len = size_t(len(payload))
	_goptr := newQNetworkDatagram(QNetworkDatagram_MakeReply(this.h, payload_alias))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkDatagram) SetSender2(address *QHostAddress, port uint16) {
	QNetworkDatagram_SetSender2(this.h, address.cPointer(), (uint16_t)(port))
}
