package qt6

/*

#include "gen_qvarlengtharray.h"
#include <stdlib.h>

*/
import "C"

import (
	"unsafe"
)

type QVLABaseBase struct {
	h *C.QVLABaseBase
}

func (this *QVLABaseBase) cPointer() *C.QVLABaseBase {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QVLABaseBase) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQVLABaseBase constructs the type using only CGO pointers.
func newQVLABaseBase(h *C.QVLABaseBase) *QVLABaseBase {
	if h == nil {
		return nil
	}

	return &QVLABaseBase{h: h}
}

// UnsafeNewQVLABaseBase constructs the type using only unsafe pointers.
func UnsafeNewQVLABaseBase(h unsafe.Pointer) *QVLABaseBase {
	return newQVLABaseBase((*C.QVLABaseBase)(h))
}

func (this *QVLABaseBase) Capacity() int64 {
	return (int64)(C.QVLABaseBase_capacity(this.h))
}

func (this *QVLABaseBase) Size() int64 {
	return (int64)(C.QVLABaseBase_size(this.h))
}

func (this *QVLABaseBase) Empty() bool {
	return (bool)(C.QVLABaseBase_empty(this.h))
}
