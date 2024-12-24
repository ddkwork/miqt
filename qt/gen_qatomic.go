package qt

import (
	"unsafe"
)

type QAtomicInt struct {
	h          uintptr
	isSubclass bool
}

// NewQAtomicInt constructs a new QAtomicInt object.
func NewQAtomicInt() *QAtomicInt {

	ret := newQAtomicInt(QAtomicInt_new())
	ret.isSubclass = true
	return ret
}

// NewQAtomicInt2 constructs a new QAtomicInt object.
func NewQAtomicInt2(param1 *QAtomicInt) *QAtomicInt {

	ret := newQAtomicInt(QAtomicInt_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQAtomicInt3 constructs a new QAtomicInt object.
func NewQAtomicInt3(value int) *QAtomicInt {

	ret := newQAtomicInt(QAtomicInt_new3((int)(value)))
	ret.isSubclass = true
	return ret
}

func (this *QAtomicInt) OperatorAssign(param1 *QAtomicInt) {
	QAtomicInt_OperatorAssign(this.h, param1.cPointer())
}
