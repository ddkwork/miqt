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
	g := newQAtomicInt(QAtomicInt_new())
	g.isSubclass = true
	return g
}

// NewQAtomicInt2 constructs a new QAtomicInt object.
func NewQAtomicInt2(param1 *QAtomicInt) *QAtomicInt {
	g := newQAtomicInt(QAtomicInt_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}

// NewQAtomicInt3 constructs a new QAtomicInt object.
func NewQAtomicInt3(value int) *QAtomicInt {
	g := newQAtomicInt(QAtomicInt_new3((int)(value)))
	g.isSubclass = true
	return g
}

func (this *QAtomicInt) OperatorAssign(param1 *QAtomicInt) {
	QAtomicInt_OperatorAssign(this.h, param1.cPointer())
}
