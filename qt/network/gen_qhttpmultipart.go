package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
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
	g := newQHttpPart(QHttpPart_new())
	g.isSubclass = true
	return g
}

// NewQHttpPart2 constructs a new QHttpPart object.
func NewQHttpPart2(other *QHttpPart) *QHttpPart {
	g := newQHttpPart(QHttpPart_new2(other.cPointer()))
	g.isSubclass = true
	return g
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
	g := newQHttpMultiPart(QHttpMultiPart_new())
	g.isSubclass = true
	return g
}

// NewQHttpMultiPart2 constructs a new QHttpMultiPart object.
func NewQHttpMultiPart2(contentType ContentType) *QHttpMultiPart {
	g := newQHttpMultiPart(QHttpMultiPart_new2(contentType))
	g.isSubclass = true
	return g
}

// NewQHttpMultiPart3 constructs a new QHttpMultiPart object.
func NewQHttpMultiPart3(parent *qt.QObject) *QHttpMultiPart {
	g := newQHttpMultiPart(QHttpMultiPart_new3((*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

// NewQHttpMultiPart4 constructs a new QHttpMultiPart object.
func NewQHttpMultiPart4(contentType ContentType, parent *qt.QObject) *QHttpMultiPart {
	g := newQHttpMultiPart(QHttpMultiPart_new4(contentType, (*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
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

func (this *QHttpMultiPart) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QHttpMultiPart_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QHttpMultiPart) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHttpMultiPart_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHttpMultiPart_MetaObject
func miqt_exec_callback_QHttpMultiPart_MetaObject(self QHttpMultiPart, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHttpMultiPart{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QHttpMultiPart) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QHttpMultiPart_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QHttpMultiPart) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHttpMultiPart_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHttpMultiPart_Metacast
func miqt_exec_callback_QHttpMultiPart_Metacast(self QHttpMultiPart, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QHttpMultiPart{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
