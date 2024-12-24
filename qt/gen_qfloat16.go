package qt

import (
	"unsafe"
)

type qfloat16 struct {
	h          uintptr
	isSubclass bool
}

// Newqfloat16 constructs a new qfloat16 object.
func Newqfloat16() *qfloat16 {
	ret := newqfloat16(qfloat16_new())
	ret.isSubclass = true
	return ret
}

// Newqfloat162 constructs a new qfloat16 object.
func Newqfloat162(param1 Initialization) *qfloat16 {
	ret := newqfloat16(qfloat16_new2((int)(param1)))
	ret.isSubclass = true
	return ret
}

// Newqfloat163 constructs a new qfloat16 object.
func Newqfloat163(f float32) *qfloat16 {
	ret := newqfloat16(qfloat16_new3((float)(f)))
	ret.isSubclass = true
	return ret
}

func (this *qfloat16) IsInf() bool {
	return (bool)(qfloat16_IsInf(this.h))
}

func (this *qfloat16) IsNaN() bool {
	return (bool)(qfloat16_IsNaN(this.h))
}

func (this *qfloat16) IsFinite() bool {
	return (bool)(qfloat16_IsFinite(this.h))
}

func (this *qfloat16) FpClassify() int {
	return (int)(qfloat16_FpClassify(this.h))
}

func (this *qfloat16) IsNormal() bool {
	return (bool)(qfloat16_IsNormal(this.h))
}
