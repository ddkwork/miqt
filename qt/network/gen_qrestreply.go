package network

import (
	"unsafe"
)

type QRestReply struct {
	h          uintptr
	isSubclass bool
}

// NewQRestReply constructs a new QRestReply object.
func NewQRestReply(reply *QNetworkReply) *QRestReply {
	g := newQRestReply(QRestReply_new(reply.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QRestReply) Swap(other *QRestReply) {
	QRestReply_Swap(this.h, other.cPointer())
}

func (this *QRestReply) NetworkReply() *QNetworkReply {
	return newQNetworkReply(QRestReply_NetworkReply(this.h))
}

func (this *QRestReply) ReadBody() []byte {
	var _bytearray struct_miqt_string = QRestReply_ReadBody(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QRestReply) ReadText() string {
	var _ms struct_miqt_string = QRestReply_ReadText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRestReply) IsSuccess() bool {
	return (bool)(QRestReply_IsSuccess(this.h))
}

func (this *QRestReply) HttpStatus() int {
	return (int)(QRestReply_HttpStatus(this.h))
}

func (this *QRestReply) IsHttpStatusSuccess() bool {
	return (bool)(QRestReply_IsHttpStatusSuccess(this.h))
}

func (this *QRestReply) HasError() bool {
	return (bool)(QRestReply_HasError(this.h))
}

func (this *QRestReply) Error() QNetworkReply__NetworkError {
	return (QNetworkReply__NetworkError)(QRestReply_Error(this.h))
}

func (this *QRestReply) ErrorString() string {
	var _ms struct_miqt_string = QRestReply_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
