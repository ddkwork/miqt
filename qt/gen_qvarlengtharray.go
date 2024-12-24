package qt

import (
	"unsafe"
)

type QVLABaseBase struct {
	h          uintptr
	isSubclass bool
}

func (this *QVLABaseBase) Capacity() size_type {
	xxxxxxxxx
}

func (this *QVLABaseBase) Size() size_type {
	xxxxxxxxx
}

func (this *QVLABaseBase) Empty() bool {
	return (bool)(QVLABaseBase_Empty(this.h))
}
