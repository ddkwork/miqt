package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QRestAccessManager struct {
	h          uintptr
	isSubclass bool
}

// NewQRestAccessManager constructs a new QRestAccessManager object.
func NewQRestAccessManager(manager *QNetworkAccessManager) *QRestAccessManager {
	ret := newQRestAccessManager(QRestAccessManager_new(manager.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQRestAccessManager2 constructs a new QRestAccessManager object.
func NewQRestAccessManager2(manager *QNetworkAccessManager, parent *qt.QObject) *QRestAccessManager {
	ret := newQRestAccessManager(QRestAccessManager_new2(manager.cPointer(), (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QRestAccessManager) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QRestAccessManager_MetaObject(this.h)))
}

func (this *QRestAccessManager) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QRestAccessManager_Metacast(this.h, param1_Cstring))
}

func QRestAccessManager_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QRestAccessManager_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRestAccessManager) NetworkAccessManager() *QNetworkAccessManager {
	return newQNetworkAccessManager(QRestAccessManager_NetworkAccessManager(this.h))
}

func (this *QRestAccessManager) DeleteResource(request *QNetworkRequest) *QNetworkReply {
	return newQNetworkReply(QRestAccessManager_DeleteResource(this.h, request.cPointer()))
}

func (this *QRestAccessManager) Head(request *QNetworkRequest) *QNetworkReply {
	return newQNetworkReply(QRestAccessManager_Head(this.h, request.cPointer()))
}

func (this *QRestAccessManager) Get(request *QNetworkRequest) *QNetworkReply {
	return newQNetworkReply(QRestAccessManager_Get(this.h, request.cPointer()))
}

func (this *QRestAccessManager) Get2(request *QNetworkRequest, data []byte) *QNetworkReply {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	return newQNetworkReply(QRestAccessManager_Get2(this.h, request.cPointer(), data_alias))
}

func (this *QRestAccessManager) Get3(request *QNetworkRequest, data *qt.QJsonDocument) *QNetworkReply {
	return newQNetworkReply(QRestAccessManager_Get3(this.h, request.cPointer(), (*QJsonDocument)(data.UnsafePointer())))
}

func (this *QRestAccessManager) Get4(request *QNetworkRequest, data *qt.QIODevice) *QNetworkReply {
	return newQNetworkReply(QRestAccessManager_Get4(this.h, request.cPointer(), (*QIODevice)(data.UnsafePointer())))
}

func (this *QRestAccessManager) Post(request *QNetworkRequest, data *qt.QJsonDocument) *QNetworkReply {
	return newQNetworkReply(QRestAccessManager_Post(this.h, request.cPointer(), (*QJsonDocument)(data.UnsafePointer())))
}

func (this *QRestAccessManager) Post2(request *QNetworkRequest, data map[string]qt.QVariant) *QNetworkReply {
	data_Keys_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(data))))
	defer free(unsafe.Pointer(data_Keys_CArray))
	data_Values_CArray := (*[0xffff]*QVariant)(malloc(size_t(8 * len(data))))
	defer free(unsafe.Pointer(data_Values_CArray))
	data_ctr := 0
	for data_k, data_v := range data {
		data_k_ms := struct_miqt_string{}
		data_k_ms.data = CString(data_k)
		data_k_ms.len = size_t(len(data_k))
		defer free(unsafe.Pointer(data_k_ms.data))
		data_Keys_CArray[data_ctr] = data_k_ms
		data_Values_CArray[data_ctr] = (*QVariant)(data_v.UnsafePointer())
		data_ctr++
	}
	data_mm := struct_miqt_map{
		len:    size_t(len(data)),
		keys:   unsafe.Pointer(data_Keys_CArray),
		values: unsafe.Pointer(data_Values_CArray),
	}
	return newQNetworkReply(QRestAccessManager_Post2(this.h, request.cPointer(), data_mm))
}

func (this *QRestAccessManager) Post3(request *QNetworkRequest, data []byte) *QNetworkReply {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	return newQNetworkReply(QRestAccessManager_Post3(this.h, request.cPointer(), data_alias))
}

func (this *QRestAccessManager) Post4(request *QNetworkRequest, data *QHttpMultiPart) *QNetworkReply {
	return newQNetworkReply(QRestAccessManager_Post4(this.h, request.cPointer(), data.cPointer()))
}

func (this *QRestAccessManager) Post5(request *QNetworkRequest, data *qt.QIODevice) *QNetworkReply {
	return newQNetworkReply(QRestAccessManager_Post5(this.h, request.cPointer(), (*QIODevice)(data.UnsafePointer())))
}

func (this *QRestAccessManager) Put(request *QNetworkRequest, data *qt.QJsonDocument) *QNetworkReply {
	return newQNetworkReply(QRestAccessManager_Put(this.h, request.cPointer(), (*QJsonDocument)(data.UnsafePointer())))
}

func (this *QRestAccessManager) Put2(request *QNetworkRequest, data map[string]qt.QVariant) *QNetworkReply {
	data_Keys_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(data))))
	defer free(unsafe.Pointer(data_Keys_CArray))
	data_Values_CArray := (*[0xffff]*QVariant)(malloc(size_t(8 * len(data))))
	defer free(unsafe.Pointer(data_Values_CArray))
	data_ctr := 0
	for data_k, data_v := range data {
		data_k_ms := struct_miqt_string{}
		data_k_ms.data = CString(data_k)
		data_k_ms.len = size_t(len(data_k))
		defer free(unsafe.Pointer(data_k_ms.data))
		data_Keys_CArray[data_ctr] = data_k_ms
		data_Values_CArray[data_ctr] = (*QVariant)(data_v.UnsafePointer())
		data_ctr++
	}
	data_mm := struct_miqt_map{
		len:    size_t(len(data)),
		keys:   unsafe.Pointer(data_Keys_CArray),
		values: unsafe.Pointer(data_Values_CArray),
	}
	return newQNetworkReply(QRestAccessManager_Put2(this.h, request.cPointer(), data_mm))
}

func (this *QRestAccessManager) Put3(request *QNetworkRequest, data []byte) *QNetworkReply {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	return newQNetworkReply(QRestAccessManager_Put3(this.h, request.cPointer(), data_alias))
}

func (this *QRestAccessManager) Put4(request *QNetworkRequest, data *QHttpMultiPart) *QNetworkReply {
	return newQNetworkReply(QRestAccessManager_Put4(this.h, request.cPointer(), data.cPointer()))
}

func (this *QRestAccessManager) Put5(request *QNetworkRequest, data *qt.QIODevice) *QNetworkReply {
	return newQNetworkReply(QRestAccessManager_Put5(this.h, request.cPointer(), (*QIODevice)(data.UnsafePointer())))
}

func (this *QRestAccessManager) Patch(request *QNetworkRequest, data *qt.QJsonDocument) *QNetworkReply {
	return newQNetworkReply(QRestAccessManager_Patch(this.h, request.cPointer(), (*QJsonDocument)(data.UnsafePointer())))
}

func (this *QRestAccessManager) Patch2(request *QNetworkRequest, data map[string]qt.QVariant) *QNetworkReply {
	data_Keys_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(data))))
	defer free(unsafe.Pointer(data_Keys_CArray))
	data_Values_CArray := (*[0xffff]*QVariant)(malloc(size_t(8 * len(data))))
	defer free(unsafe.Pointer(data_Values_CArray))
	data_ctr := 0
	for data_k, data_v := range data {
		data_k_ms := struct_miqt_string{}
		data_k_ms.data = CString(data_k)
		data_k_ms.len = size_t(len(data_k))
		defer free(unsafe.Pointer(data_k_ms.data))
		data_Keys_CArray[data_ctr] = data_k_ms
		data_Values_CArray[data_ctr] = (*QVariant)(data_v.UnsafePointer())
		data_ctr++
	}
	data_mm := struct_miqt_map{
		len:    size_t(len(data)),
		keys:   unsafe.Pointer(data_Keys_CArray),
		values: unsafe.Pointer(data_Values_CArray),
	}
	return newQNetworkReply(QRestAccessManager_Patch2(this.h, request.cPointer(), data_mm))
}

func (this *QRestAccessManager) Patch3(request *QNetworkRequest, data []byte) *QNetworkReply {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	return newQNetworkReply(QRestAccessManager_Patch3(this.h, request.cPointer(), data_alias))
}

func (this *QRestAccessManager) Patch4(request *QNetworkRequest, data *qt.QIODevice) *QNetworkReply {
	return newQNetworkReply(QRestAccessManager_Patch4(this.h, request.cPointer(), (*QIODevice)(data.UnsafePointer())))
}

func (this *QRestAccessManager) SendCustomRequest(request *QNetworkRequest, method []byte, data []byte) *QNetworkReply {
	method_alias := struct_miqt_string{}
	method_alias.data = (char)(unsafe.Pointer(&method[0]))
	method_alias.len = size_t(len(method))
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	return newQNetworkReply(QRestAccessManager_SendCustomRequest(this.h, request.cPointer(), method_alias, data_alias))
}

func (this *QRestAccessManager) SendCustomRequest2(request *QNetworkRequest, method []byte, data *qt.QIODevice) *QNetworkReply {
	method_alias := struct_miqt_string{}
	method_alias.data = (char)(unsafe.Pointer(&method[0]))
	method_alias.len = size_t(len(method))
	return newQNetworkReply(QRestAccessManager_SendCustomRequest2(this.h, request.cPointer(), method_alias, (*QIODevice)(data.UnsafePointer())))
}

func (this *QRestAccessManager) SendCustomRequest3(request *QNetworkRequest, method []byte, data *QHttpMultiPart) *QNetworkReply {
	method_alias := struct_miqt_string{}
	method_alias.data = (char)(unsafe.Pointer(&method[0]))
	method_alias.len = size_t(len(method))
	return newQNetworkReply(QRestAccessManager_SendCustomRequest3(this.h, request.cPointer(), method_alias, data.cPointer()))
}

func QRestAccessManager_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QRestAccessManager_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QRestAccessManager_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QRestAccessManager_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRestAccessManager) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QRestAccessManager_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QRestAccessManager) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRestAccessManager_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRestAccessManager_MetaObject
func miqt_exec_callback_QRestAccessManager_MetaObject(self QRestAccessManager, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRestAccessManager{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QRestAccessManager) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QRestAccessManager_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QRestAccessManager) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRestAccessManager_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRestAccessManager_Metacast
func miqt_exec_callback_QRestAccessManager_Metacast(self QRestAccessManager, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QRestAccessManager{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
