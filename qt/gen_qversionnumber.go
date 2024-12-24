package qt

import (
	"unsafe"
)

type QVersionNumber struct {
	h          uintptr
	isSubclass bool
}

// NewQVersionNumber constructs a new QVersionNumber object.
func NewQVersionNumber() *QVersionNumber {

	ret := newQVersionNumber(QVersionNumber_new())
	ret.isSubclass = true
	return ret
}

// NewQVersionNumber2 constructs a new QVersionNumber object.
func NewQVersionNumber2(args QSpan<

const int>) *QVersionNumber {

ret := newQVersionNumber(QVersionNumber_new2(args))
ret.isSubclass = true
return ret
}

// NewQVersionNumber3 constructs a new QVersionNumber object.
func NewQVersionNumber3(maj int) *QVersionNumber {

	ret := newQVersionNumber(QVersionNumber_new3((int)(maj)))
	ret.isSubclass = true
	return ret
}

// NewQVersionNumber4 constructs a new QVersionNumber object.
func NewQVersionNumber4(maj int, min int) *QVersionNumber {

	ret := newQVersionNumber(QVersionNumber_new4((int)(maj), (int)(min)))
	ret.isSubclass = true
	return ret
}

// NewQVersionNumber5 constructs a new QVersionNumber object.
func NewQVersionNumber5(maj int, min int, mic int) *QVersionNumber {

	ret := newQVersionNumber(QVersionNumber_new5((int)(maj), (int)(min), (int)(mic)))
	ret.isSubclass = true
	return ret
}

func (this *QVersionNumber) IsNull() bool {
	return (bool)(QVersionNumber_IsNull(this.h))
}

func (this *QVersionNumber) IsNormalized() bool {
	return (bool)(QVersionNumber_IsNormalized(this.h))
}

func (this *QVersionNumber) MajorVersion() int {
	return (int)(QVersionNumber_MajorVersion(this.h))
}

func (this *QVersionNumber) MinorVersion() int {
	return (int)(QVersionNumber_MinorVersion(this.h))
}

func (this *QVersionNumber) MicroVersion() int {
	return (int)(QVersionNumber_MicroVersion(this.h))
}

func (this *QVersionNumber) Normalized() *QVersionNumber {
	_goptr := newQVersionNumber(QVersionNumber_Normalized(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVersionNumber) Segments() []int {
	var _ma struct_miqt_array = QVersionNumber_Segments(this.h)
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret
}

func (this *QVersionNumber) SegmentAt(index int64) int {
	return (int)(QVersionNumber_SegmentAt(this.h, (ptrdiff_t)(index)))
}

func (this *QVersionNumber) SegmentCount() int64 {
	return (int64)(QVersionNumber_SegmentCount(this.h))
}

func (this *QVersionNumber) Begin() const_iterator {
	xxxxxxxxx
}

func (this *QVersionNumber) End() const_iterator {
	xxxxxxxxx
}

func (this *QVersionNumber) Cbegin() const_iterator {
	xxxxxxxxx
}

func (this *QVersionNumber) Cend() const_iterator {
	xxxxxxxxx
}

func (this *QVersionNumber) Rbegin() const_reverse_iterator {
	xxxxxxxxx
}

func (this *QVersionNumber) Rend() const_reverse_iterator {
	xxxxxxxxx
}

func (this *QVersionNumber) Crbegin() const_reverse_iterator {
	xxxxxxxxx
}

func (this *QVersionNumber) Crend() const_reverse_iterator {
	xxxxxxxxx
}

func (this *QVersionNumber) ConstBegin() const_iterator {
	xxxxxxxxx
}

func (this *QVersionNumber) ConstEnd() const_iterator {
	xxxxxxxxx
}

func (this *QVersionNumber) IsPrefixOf(other *QVersionNumber) bool {
	return (bool)(QVersionNumber_IsPrefixOf(this.h, other.cPointer()))
}

func QVersionNumber_Compare(v1 *QVersionNumber, v2 *QVersionNumber) int {
	return (int)(QVersionNumber_Compare(v1.cPointer(), v2.cPointer()))
}

func QVersionNumber_CommonPrefix(v1 *QVersionNumber, v2 *QVersionNumber) *QVersionNumber {
	_goptr := newQVersionNumber(QVersionNumber_CommonPrefix(v1.cPointer(), v2.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVersionNumber) ToString() string {
	var _ms struct_miqt_string = QVersionNumber_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QVersionNumber_FromString(stringVal QAnyStringView) *QVersionNumber {
	_goptr := newQVersionNumber(QVersionNumber_FromString(stringVal.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QVersionNumber_FromString2(stringVal QAnyStringView, suffixIndex *int64) *QVersionNumber {
	_goptr := newQVersionNumber(QVersionNumber_FromString2(stringVal.cPointer(), (*ptrdiff_t)(unsafe.Pointer(suffixIndex))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
