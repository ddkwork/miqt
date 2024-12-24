package webchannel

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QWebChannelAbstractTransport struct {
	h          uintptr
	isSubclass bool
}

// NewQWebChannelAbstractTransport constructs a new QWebChannelAbstractTransport object.
func NewQWebChannelAbstractTransport() *QWebChannelAbstractTransport {
	g := newQWebChannelAbstractTransport(QWebChannelAbstractTransport_new())
	g.isSubclass = true
	return g
}

// NewQWebChannelAbstractTransport2 constructs a new QWebChannelAbstractTransport object.
func NewQWebChannelAbstractTransport2(parent *qt.QObject) *QWebChannelAbstractTransport {
	g := newQWebChannelAbstractTransport(QWebChannelAbstractTransport_new2((*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

func (this *QWebChannelAbstractTransport) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QWebChannelAbstractTransport_MetaObject(this.h)))
}

func (this *QWebChannelAbstractTransport) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QWebChannelAbstractTransport_Metacast(this.h, param1_Cstring))
}

func QWebChannelAbstractTransport_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QWebChannelAbstractTransport_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebChannelAbstractTransport) SendMessage(message *qt.QJsonObject) {
	QWebChannelAbstractTransport_SendMessage(this.h, (*QJsonObject)(message.UnsafePointer()))
}

func (this *QWebChannelAbstractTransport) MessageReceived(message *qt.QJsonObject, transport *QWebChannelAbstractTransport) {
	QWebChannelAbstractTransport_MessageReceived(this.h, (*QJsonObject)(message.UnsafePointer()), transport.cPointer())
}

func (this *QWebChannelAbstractTransport) OnMessageReceived(slot func(message *qt.QJsonObject, transport *QWebChannelAbstractTransport)) {
	QWebChannelAbstractTransport_connect_MessageReceived(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannelAbstractTransport_MessageReceived
func miqt_exec_callback_QWebChannelAbstractTransport_MessageReceived(cb intptr_t, message *QJsonObject, transport *QWebChannelAbstractTransport) {
	gofunc, ok := cgo.Handle(cb).Value().(func(message *qt.QJsonObject, transport *QWebChannelAbstractTransport))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQJsonObject(unsafe.Pointer(message))

	slotval2 := newQWebChannelAbstractTransport(transport)

	gofunc(slotval1, slotval2)
}

func QWebChannelAbstractTransport_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWebChannelAbstractTransport_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWebChannelAbstractTransport_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWebChannelAbstractTransport_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebChannelAbstractTransport) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QWebChannelAbstractTransport_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QWebChannelAbstractTransport) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWebChannelAbstractTransport_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannelAbstractTransport_MetaObject
func miqt_exec_callback_QWebChannelAbstractTransport_MetaObject(self QWebChannelAbstractTransport, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWebChannelAbstractTransport{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QWebChannelAbstractTransport) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QWebChannelAbstractTransport_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QWebChannelAbstractTransport) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWebChannelAbstractTransport_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannelAbstractTransport_Metacast
func miqt_exec_callback_QWebChannelAbstractTransport_Metacast(self QWebChannelAbstractTransport, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QWebChannelAbstractTransport{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
