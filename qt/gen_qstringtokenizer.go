package qt

import (
	"unsafe"
)

type QStringTokenizerBaseBase struct {
	h          uintptr
	isSubclass bool
}

// NewQStringTokenizerBaseBase constructs a new QStringTokenizerBaseBase object.
func NewQStringTokenizerBaseBase(param1 *QStringTokenizerBaseBase) *QStringTokenizerBaseBase {
	ret := newQStringTokenizerBaseBase(QStringTokenizerBaseBase_new(param1.cPointer()))
	ret.isSubclass = true
	return ret
}
