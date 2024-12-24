package qt

import (
	"unsafe"
)

type QFlag struct {
	h          uintptr
	isSubclass bool
}

// NewQFlag constructs a new QFlag object.
func NewQFlag(value int) *QFlag {
	g := newQFlag(QFlag_new((int)(value)))
	g.isSubclass = true
	return g
}

// NewQFlag2 constructs a new QFlag object.
func NewQFlag2(param1 *QFlag) *QFlag {
	g := newQFlag(QFlag_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}

type QIncompatibleFlag struct {
	h          uintptr
	isSubclass bool
}

// NewQIncompatibleFlag constructs a new QIncompatibleFlag object.
func NewQIncompatibleFlag(i int) *QIncompatibleFlag {
	g := newQIncompatibleFlag(QIncompatibleFlag_new((int)(i)))
	g.isSubclass = true
	return g
}

// NewQIncompatibleFlag2 constructs a new QIncompatibleFlag object.
func NewQIncompatibleFlag2(param1 *QIncompatibleFlag) *QIncompatibleFlag {
	g := newQIncompatibleFlag(QIncompatibleFlag_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}
