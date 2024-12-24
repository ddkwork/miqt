package network

import (
	"unsafe"
)

type QHttp1Configuration struct {
	h          uintptr
	isSubclass bool
}

// NewQHttp1Configuration constructs a new QHttp1Configuration object.
func NewQHttp1Configuration() *QHttp1Configuration {

	ret := newQHttp1Configuration(QHttp1Configuration_new())
	ret.isSubclass = true
	return ret
}

// NewQHttp1Configuration2 constructs a new QHttp1Configuration object.
func NewQHttp1Configuration2(other *QHttp1Configuration) *QHttp1Configuration {

	ret := newQHttp1Configuration(QHttp1Configuration_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QHttp1Configuration) OperatorAssign(other *QHttp1Configuration) {
	QHttp1Configuration_OperatorAssign(this.h, other.cPointer())
}

func (this *QHttp1Configuration) SetNumberOfConnectionsPerHost(amount int64) {
	QHttp1Configuration_SetNumberOfConnectionsPerHost(this.h, (ptrdiff_t)(amount))
}

func (this *QHttp1Configuration) NumberOfConnectionsPerHost() int64 {
	return (int64)(QHttp1Configuration_NumberOfConnectionsPerHost(this.h))
}

func (this *QHttp1Configuration) Swap(other *QHttp1Configuration) {
	QHttp1Configuration_Swap(this.h, other.cPointer())
}
