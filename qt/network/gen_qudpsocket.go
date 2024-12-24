package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QUdpSocket struct {
	h          uintptr
	isSubclass bool
}

// NewQUdpSocket constructs a new QUdpSocket object.
func NewQUdpSocket() *QUdpSocket {
	ret := newQUdpSocket(QUdpSocket_new())
	ret.isSubclass = true
	return ret
}

// NewQUdpSocket2 constructs a new QUdpSocket object.
func NewQUdpSocket2(parent *qt.QObject) *QUdpSocket {
	ret := newQUdpSocket(QUdpSocket_new2((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QUdpSocket) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QUdpSocket_MetaObject(this.h)))
}

func (this *QUdpSocket) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QUdpSocket_Metacast(this.h, param1_Cstring))
}

func QUdpSocket_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QUdpSocket_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUdpSocket) Bind(addr QHostAddress__SpecialAddress) bool {
	return (bool)(QUdpSocket_Bind(this.h, (int)(addr)))
}

func (this *QUdpSocket) JoinMulticastGroup(groupAddress *QHostAddress) bool {
	return (bool)(QUdpSocket_JoinMulticastGroup(this.h, groupAddress.cPointer()))
}

func (this *QUdpSocket) JoinMulticastGroup2(groupAddress *QHostAddress, iface *QNetworkInterface) bool {
	return (bool)(QUdpSocket_JoinMulticastGroup2(this.h, groupAddress.cPointer(), iface.cPointer()))
}

func (this *QUdpSocket) LeaveMulticastGroup(groupAddress *QHostAddress) bool {
	return (bool)(QUdpSocket_LeaveMulticastGroup(this.h, groupAddress.cPointer()))
}

func (this *QUdpSocket) LeaveMulticastGroup2(groupAddress *QHostAddress, iface *QNetworkInterface) bool {
	return (bool)(QUdpSocket_LeaveMulticastGroup2(this.h, groupAddress.cPointer(), iface.cPointer()))
}

func (this *QUdpSocket) MulticastInterface() *QNetworkInterface {
	_goptr := newQNetworkInterface(QUdpSocket_MulticastInterface(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUdpSocket) SetMulticastInterface(iface *QNetworkInterface) {
	QUdpSocket_SetMulticastInterface(this.h, iface.cPointer())
}

func (this *QUdpSocket) HasPendingDatagrams() bool {
	return (bool)(QUdpSocket_HasPendingDatagrams(this.h))
}

func (this *QUdpSocket) PendingDatagramSize() int64 {
	return (int64)(QUdpSocket_PendingDatagramSize(this.h))
}

func (this *QUdpSocket) ReceiveDatagram() *QNetworkDatagram {
	_goptr := newQNetworkDatagram(QUdpSocket_ReceiveDatagram(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUdpSocket) ReadDatagram(data string, maxlen int64) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))
	return (int64)(QUdpSocket_ReadDatagram(this.h, data_Cstring, (longlong)(maxlen)))
}

func (this *QUdpSocket) WriteDatagram(datagram *QNetworkDatagram) int64 {
	return (int64)(QUdpSocket_WriteDatagram(this.h, datagram.cPointer()))
}

func (this *QUdpSocket) WriteDatagram2(data string, lenVal int64, host *QHostAddress, port uint16) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))
	return (int64)(QUdpSocket_WriteDatagram2(this.h, data_Cstring, (longlong)(lenVal), host.cPointer(), (uint16_t)(port)))
}

func (this *QUdpSocket) WriteDatagram3(datagram []byte, host *QHostAddress, port uint16) int64 {
	datagram_alias := struct_miqt_string{}
	datagram_alias.data = (char)(unsafe.Pointer(&datagram[0]))
	datagram_alias.len = size_t(len(datagram))
	return (int64)(QUdpSocket_WriteDatagram3(this.h, datagram_alias, host.cPointer(), (uint16_t)(port)))
}

func QUdpSocket_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QUdpSocket_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QUdpSocket_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QUdpSocket_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUdpSocket) Bind2(addr QHostAddress__SpecialAddress, port uint16) bool {
	return (bool)(QUdpSocket_Bind2(this.h, (int)(addr), (uint16_t)(port)))
}

func (this *QUdpSocket) Bind3(addr QHostAddress__SpecialAddress, port uint16, mode BindMode) bool {
	return (bool)(QUdpSocket_Bind3(this.h, (int)(addr), (uint16_t)(port), mode))
}

func (this *QUdpSocket) ReceiveDatagram1(maxSize int64) *QNetworkDatagram {
	_goptr := newQNetworkDatagram(QUdpSocket_ReceiveDatagram1(this.h, (longlong)(maxSize)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUdpSocket) ReadDatagram3(data string, maxlen int64, host *QHostAddress) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))
	return (int64)(QUdpSocket_ReadDatagram3(this.h, data_Cstring, (longlong)(maxlen), host.cPointer()))
}

func (this *QUdpSocket) ReadDatagram4(data string, maxlen int64, host *QHostAddress, port *uint16) int64 {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))
	return (int64)(QUdpSocket_ReadDatagram4(this.h, data_Cstring, (longlong)(maxlen), host.cPointer(), (*uint16_t)(unsafe.Pointer(port))))
}

func (this *QUdpSocket) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QUdpSocket_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QUdpSocket) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_MetaObject
func miqt_exec_callback_QUdpSocket_MetaObject(self QUdpSocket, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QUdpSocket{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QUdpSocket) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QUdpSocket_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QUdpSocket) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUdpSocket_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUdpSocket_Metacast
func miqt_exec_callback_QUdpSocket_Metacast(self QUdpSocket, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QUdpSocket{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
