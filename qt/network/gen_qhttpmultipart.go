package network

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QHttpMultiPart__ContentType int

const (
	QHttpMultiPart__MixedType       QHttpMultiPart__ContentType = 0
	QHttpMultiPart__RelatedType     QHttpMultiPart__ContentType = 1
	QHttpMultiPart__FormDataType    QHttpMultiPart__ContentType = 2
	QHttpMultiPart__AlternativeType QHttpMultiPart__ContentType = 3
)

type QHttpPart struct {
	h          uintptr
	isSubclass bool
}

// NewQHttpPart constructs a new QHttpPart object.
func NewQHttpPart() *QHttpPart {

	ret := newQHttpPart(QHttpPart_new())
	ret.isSubclass = true
	return ret
}

// NewQHttpPart2 constructs a new QHttpPart object.
func NewQHttpPart2(other *QHttpPart) *QHttpPart {

	ret := newQHttpPart(QHttpPart_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QHttpPart) OperatorAssign(other *QHttpPart) {
	QHttpPart_OperatorAssign(this.h, other.cPointer())
}

func (this *QHttpPart) Swap(other *QHttpPart) {
	QHttpPart_Swap(this.h, other.cPointer())
}

func (this *QHttpPart) OperatorEqual(other *QHttpPart) bool {
	return (bool)(QHttpPart_OperatorEqual(this.h, other.cPointer()))
}

func (this *QHttpPart) OperatorNotEqual(other *QHttpPart) bool {
	return (bool)(QHttpPart_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QHttpPart) SetHeader(header QNetworkRequest__KnownHeaders, value *qt.QVariant) {
	QHttpPart_SetHeader(this.h, (int)(header), (*QVariant)(value.UnsafePointer()))
}

func (this *QHttpPart) SetRawHeader(headerName []byte, headerValue []byte) {
	headerName_alias := struct_miqt_string{}
	headerName_alias.data = (char)(unsafe.Pointer(&headerName[0]))
	headerName_alias.len = size_t(len(headerName))
	headerValue_alias := struct_miqt_string{}
	headerValue_alias.data = (char)(unsafe.Pointer(&headerValue[0]))
	headerValue_alias.len = size_t(len(headerValue))
	QHttpPart_SetRawHeader(this.h, headerName_alias, headerValue_alias)
}

func (this *QHttpPart) SetBody(body []byte) {
	body_alias := struct_miqt_string{}
	body_alias.data = (char)(unsafe.Pointer(&body[0]))
	body_alias.len = size_t(len(body))
	QHttpPart_SetBody(this.h, body_alias)
}

func (this *QHttpPart) SetBodyDevice(device *qt.QIODevice) {
	QHttpPart_SetBodyDevice(this.h, (*QIODevice)(device.UnsafePointer()))
}

type QHttpMultiPart struct {
	h          uintptr
	isSubclass bool
}

// NewQHttpMultiPart constructs a new QHttpMultiPart object.
func NewQHttpMultiPart() *QHttpMultiPart {

	ret := newQHttpMultiPart(QHttpMultiPart_new())
	ret.isSubclass = true
	return ret
}

// NewQHttpMultiPart2 constructs a new QHttpMultiPart object.
func NewQHttpMultiPart2(contentType ContentType) *QHttpMultiPart {

	ret := newQHttpMultiPart(QHttpMultiPart_new2(contentType))
	ret.isSubclass = true
	return ret
}

// NewQHttpMultiPart3 constructs a new QHttpMultiPart object.
func NewQHttpMultiPart3(parent *qt.QObject) *QHttpMultiPart {

	ret := newQHttpMultiPart(QHttpMultiPart_new3((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQHttpMultiPart4 constructs a new QHttpMultiPart object.
func NewQHttpMultiPart4(contentType ContentType, parent *qt.QObject) *QHttpMultiPart {

	ret := newQHttpMultiPart(QHttpMultiPart_new4(contentType, (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QHttpMultiPart) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QHttpMultiPart_MetaObject(this.h)))
}

func (this *QHttpMultiPart) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QHttpMultiPart_Metacast(this.h, param1_Cstring))
}

func QHttpMultiPart_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QHttpMultiPart_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QHttpMultiPart) Append(httpPart *QHttpPart) {
	QHttpMultiPart_Append(this.h, httpPart.cPointer())
}

func (this *QHttpMultiPart) SetContentType(contentType ContentType) {
	QHttpMultiPart_SetContentType(this.h, contentType)
}

func (this *QHttpMultiPart) Boundary() []byte {
	var _bytearray struct_miqt_string = QHttpMultiPart_Boundary(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QHttpMultiPart) SetBoundary(boundary []byte) {
	boundary_alias := struct_miqt_string{}
	boundary_alias.data = (char)(unsafe.Pointer(&boundary[0]))
	boundary_alias.len = size_t(len(boundary))
	QHttpMultiPart_SetBoundary(this.h, boundary_alias)
}

func QHttpMultiPart_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QHttpMultiPart_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QHttpMultiPart_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QHttpMultiPart_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QHttpMultiPart) callVirtualBase_Event(event *qt.QEvent) bool {

	return (bool)(QHttpMultiPart_virtualbase_Event(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer())))

}
func (this *QHttpMultiPart) OnEvent(slot func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHttpMultiPart_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHttpMultiPart_Event
func miqt_exec_callback_QHttpMultiPart_Event(self QHttpMultiPart, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QHttpMultiPart{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QHttpMultiPart) callVirtualBase_EventFilter(watched *qt.QObject, event *qt.QEvent) bool {

	return (bool)(QHttpMultiPart_virtualbase_EventFilter(unsafe.Pointer(this.h), (*QObject)(watched.UnsafePointer()), (*QEvent)(event.UnsafePointer())))

}
func (this *QHttpMultiPart) OnEventFilter(slot func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHttpMultiPart_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHttpMultiPart_EventFilter
func miqt_exec_callback_QHttpMultiPart_EventFilter(self QHttpMultiPart, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QHttpMultiPart{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QHttpMultiPart) callVirtualBase_TimerEvent(event *qt.QTimerEvent) {

	QHttpMultiPart_virtualbase_TimerEvent(unsafe.Pointer(this.h), (*QTimerEvent)(event.UnsafePointer()))

}
func (this *QHttpMultiPart) OnTimerEvent(slot func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHttpMultiPart_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHttpMultiPart_TimerEvent
func miqt_exec_callback_QHttpMultiPart_TimerEvent(self QHttpMultiPart, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QHttpMultiPart{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QHttpMultiPart) callVirtualBase_ChildEvent(event *qt.QChildEvent) {

	QHttpMultiPart_virtualbase_ChildEvent(unsafe.Pointer(this.h), (*QChildEvent)(event.UnsafePointer()))

}
func (this *QHttpMultiPart) OnChildEvent(slot func(super func(event *qt.QChildEvent), event *qt.QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHttpMultiPart_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHttpMultiPart_ChildEvent
func miqt_exec_callback_QHttpMultiPart_ChildEvent(self QHttpMultiPart, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QChildEvent), event *qt.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QHttpMultiPart{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QHttpMultiPart) callVirtualBase_CustomEvent(event *qt.QEvent) {

	QHttpMultiPart_virtualbase_CustomEvent(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer()))

}
func (this *QHttpMultiPart) OnCustomEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHttpMultiPart_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHttpMultiPart_CustomEvent
func miqt_exec_callback_QHttpMultiPart_CustomEvent(self QHttpMultiPart, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QHttpMultiPart{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QHttpMultiPart) callVirtualBase_ConnectNotify(signal *qt.QMetaMethod) {

	QHttpMultiPart_virtualbase_ConnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QHttpMultiPart) OnConnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHttpMultiPart_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHttpMultiPart_ConnectNotify
func miqt_exec_callback_QHttpMultiPart_ConnectNotify(self QHttpMultiPart, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QHttpMultiPart{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QHttpMultiPart) callVirtualBase_DisconnectNotify(signal *qt.QMetaMethod) {

	QHttpMultiPart_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QHttpMultiPart) OnDisconnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHttpMultiPart_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHttpMultiPart_DisconnectNotify
func miqt_exec_callback_QHttpMultiPart_DisconnectNotify(self QHttpMultiPart, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QHttpMultiPart{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
