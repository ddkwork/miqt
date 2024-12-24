package qt

import (
	"unsafe"
)

type QBitArray struct {
	h          uintptr
	isSubclass bool
}

// NewQBitArray constructs a new QBitArray object.
func NewQBitArray() *QBitArray {
	g := newQBitArray(QBitArray_new())
	g.isSubclass = true
	return g
}

// NewQBitArray2 constructs a new QBitArray object.
func NewQBitArray2(size int64) *QBitArray {
	g := newQBitArray(QBitArray_new2((ptrdiff_t)(size)))
	g.isSubclass = true
	return g
}

// NewQBitArray3 constructs a new QBitArray object.
func NewQBitArray3(other *QBitArray) *QBitArray {
	g := newQBitArray(QBitArray_new3(other.cPointer()))
	g.isSubclass = true
	return g
}

// NewQBitArray4 constructs a new QBitArray object.
func NewQBitArray4(size int64, val bool) *QBitArray {
	g := newQBitArray(QBitArray_new4((ptrdiff_t)(size), (bool)(val)))
	g.isSubclass = true
	return g
}

func (this *QBitArray) OperatorAssign(other *QBitArray) {
	QBitArray_OperatorAssign(this.h, other.cPointer())
}

func (this *QBitArray) Swap(other *QBitArray) {
	QBitArray_Swap(this.h, other.cPointer())
}

func (this *QBitArray) Size() int64 {
	return (int64)(QBitArray_Size(this.h))
}

func (this *QBitArray) Count() int64 {
	return (int64)(QBitArray_Count(this.h))
}

func (this *QBitArray) CountWithOn(on bool) int64 {
	return (int64)(QBitArray_CountWithOn(this.h, (bool)(on)))
}

func (this *QBitArray) IsEmpty() bool {
	return (bool)(QBitArray_IsEmpty(this.h))
}

func (this *QBitArray) IsNull() bool {
	return (bool)(QBitArray_IsNull(this.h))
}

func (this *QBitArray) Resize(size int64) {
	QBitArray_Resize(this.h, (ptrdiff_t)(size))
}

func (this *QBitArray) Detach() {
	QBitArray_Detach(this.h)
}

func (this *QBitArray) IsDetached() bool {
	return (bool)(QBitArray_IsDetached(this.h))
}

func (this *QBitArray) Clear() {
	QBitArray_Clear(this.h)
}

func (this *QBitArray) TestBit(i int64) bool {
	return (bool)(QBitArray_TestBit(this.h, (ptrdiff_t)(i)))
}

func (this *QBitArray) SetBit(i int64) {
	QBitArray_SetBit(this.h, (ptrdiff_t)(i))
}

func (this *QBitArray) SetBit2(i int64, val bool) {
	QBitArray_SetBit2(this.h, (ptrdiff_t)(i), (bool)(val))
}

func (this *QBitArray) ClearBit(i int64) {
	QBitArray_ClearBit(this.h, (ptrdiff_t)(i))
}

func (this *QBitArray) ToggleBit(i int64) bool {
	return (bool)(QBitArray_ToggleBit(this.h, (ptrdiff_t)(i)))
}

func (this *QBitArray) At(i int64) bool {
	return (bool)(QBitArray_At(this.h, (ptrdiff_t)(i)))
}

func (this *QBitArray) OperatorSubscript(i int64) *QBitRef {
	_goptr := newQBitRef(QBitArray_OperatorSubscript(this.h, (ptrdiff_t)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QBitArray) OperatorSubscriptWithQsizetype(i int64) bool {
	return (bool)(QBitArray_OperatorSubscriptWithQsizetype(this.h, (ptrdiff_t)(i)))
}

func (this *QBitArray) OperatorBitwiseAndAssign(param1 *QBitArray) {
	QBitArray_OperatorBitwiseAndAssign(this.h, param1.cPointer())
}

func (this *QBitArray) OperatorBitwiseOrAssign(param1 *QBitArray) {
	QBitArray_OperatorBitwiseOrAssign(this.h, param1.cPointer())
}

func (this *QBitArray) OperatorBitwiseNotAssign(param1 *QBitArray) {
	QBitArray_OperatorBitwiseNotAssign(this.h, param1.cPointer())
}

func (this *QBitArray) Fill(aval bool) bool {
	return (bool)(QBitArray_Fill(this.h, (bool)(aval)))
}

func (this *QBitArray) Fill2(val bool, first int64, last int64) {
	QBitArray_Fill2(this.h, (bool)(val), (ptrdiff_t)(first), (ptrdiff_t)(last))
}

func (this *QBitArray) Truncate(pos int64) {
	QBitArray_Truncate(this.h, (ptrdiff_t)(pos))
}

func (this *QBitArray) Bits() string {
	_ret := QBitArray_Bits(this.h)
	return GoString(_ret)
}

func QBitArray_FromBits(data string, lenVal int64) *QBitArray {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))
	_goptr := newQBitArray(QBitArray_FromBits(data_Cstring, (ptrdiff_t)(lenVal)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QBitArray) ToUInt32(endianness QSysInfo__Endian) uint {
	return (uint)(QBitArray_ToUInt32(this.h, (int)(endianness)))
}

func (this *QBitArray) DataPtr() *DataPtr {
	xxxxxxxxx
}

func (this *QBitArray) DataPtr2() *DataPtr {
	xxxxxxxxx
}

func (this *QBitArray) Fill22(aval bool, asize int64) bool {
	return (bool)(QBitArray_Fill22(this.h, (bool)(aval), (ptrdiff_t)(asize)))
}

func (this *QBitArray) ToUInt322(endianness QSysInfo__Endian, ok *bool) uint {
	return (uint)(QBitArray_ToUInt322(this.h, (int)(endianness), (*bool)(unsafe.Pointer(ok))))
}

type QBitRef struct {
	h          uintptr
	isSubclass bool
}

// NewQBitRef constructs a new QBitRef object.
func NewQBitRef(param1 *QBitRef) *QBitRef {
	g := newQBitRef(QBitRef_new(param1.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QBitRef) OperatorNot() bool {
	return (bool)(QBitRef_OperatorNot(this.h))
}

func (this *QBitRef) OperatorAssign(val *QBitRef) {
	QBitRef_OperatorAssign(this.h, val.cPointer())
}

func (this *QBitRef) OperatorAssignWithVal(val bool) {
	QBitRef_OperatorAssignWithVal(this.h, (bool)(val))
}
