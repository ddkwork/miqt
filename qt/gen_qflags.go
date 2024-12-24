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
	ret := newQFlag(QFlag_new((int)(value)))
	ret.isSubclass = true
	return ret
}

// NewQFlag2 constructs a new QFlag object.
func NewQFlag2(param1 *QFlag) *QFlag {
	ret := newQFlag(QFlag_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

type QIncompatibleFlag struct {
	h          uintptr
	isSubclass bool
}

// NewQIncompatibleFlag constructs a new QIncompatibleFlag object.
func NewQIncompatibleFlag(i int) *QIncompatibleFlag {
	ret := newQIncompatibleFlag(QIncompatibleFlag_new((int)(i)))
	ret.isSubclass = true
	return ret
}

// NewQIncompatibleFlag2 constructs a new QIncompatibleFlag object.
func NewQIncompatibleFlag2(param1 *QIncompatibleFlag) *QIncompatibleFlag {
	ret := newQIncompatibleFlag(QIncompatibleFlag_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}
