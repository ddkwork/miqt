package cbor

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QCborArray struct {
	h          uintptr
	isSubclass bool
}

// NewQCborArray constructs a new QCborArray object.
func NewQCborArray() *QCborArray {
	ret := newQCborArray(QCborArray_new())
	ret.isSubclass = true
	return ret
}

// NewQCborArray2 constructs a new QCborArray object.
func NewQCborArray2(other *QCborArray) *QCborArray {
	ret := newQCborArray(QCborArray_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QCborArray) OperatorAssign(other *QCborArray) {
	QCborArray_OperatorAssign(this.h, other.cPointer())
}

func (this *QCborArray) Swap(other *QCborArray) {
	QCborArray_Swap(this.h, other.cPointer())
}

func (this *QCborArray) ToCborValue() *QCborValue {
	_goptr := newQCborValue(QCborArray_ToCborValue(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray) Size() int64 {
	return (int64)(QCborArray_Size(this.h))
}

func (this *QCborArray) IsEmpty() bool {
	return (bool)(QCborArray_IsEmpty(this.h))
}

func (this *QCborArray) Clear() {
	QCborArray_Clear(this.h)
}

func (this *QCborArray) At(i int64) *QCborValue {
	_goptr := newQCborValue(QCborArray_At(this.h, (ptrdiff_t)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray) First() *QCborValue {
	_goptr := newQCborValue(QCborArray_First(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray) Last() *QCborValue {
	_goptr := newQCborValue(QCborArray_Last(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray) OperatorSubscript(i int64) *QCborValue {
	_goptr := newQCborValue(QCborArray_OperatorSubscript(this.h, (ptrdiff_t)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray) First2() *QCborValueRef {
	_goptr := newQCborValueRef(QCborArray_First2(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray) Last2() *QCborValueRef {
	_goptr := newQCborValueRef(QCborArray_Last2(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray) OperatorSubscriptWithQsizetype(i int64) *QCborValueRef {
	_goptr := newQCborValueRef(QCborArray_OperatorSubscriptWithQsizetype(this.h, (ptrdiff_t)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray) Insert(i int64, value *QCborValue) {
	QCborArray_Insert(this.h, (ptrdiff_t)(i), value.cPointer())
}

func (this *QCborArray) Prepend(value *QCborValue) {
	QCborArray_Prepend(this.h, value.cPointer())
}

func (this *QCborArray) Append(value *QCborValue) {
	QCborArray_Append(this.h, value.cPointer())
}

func (this *QCborArray) Extract(it ConstIterator) *QCborValue {
	_goptr := newQCborValue(QCborArray_Extract(this.h, it))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray) ExtractWithIt(it Iterator) *QCborValue {
	_goptr := newQCborValue(QCborArray_ExtractWithIt(this.h, it))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray) RemoveAt(i int64) {
	QCborArray_RemoveAt(this.h, (ptrdiff_t)(i))
}

func (this *QCborArray) TakeAt(i int64) *QCborValue {
	_goptr := newQCborValue(QCborArray_TakeAt(this.h, (ptrdiff_t)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray) RemoveFirst() {
	QCborArray_RemoveFirst(this.h)
}

func (this *QCborArray) RemoveLast() {
	QCborArray_RemoveLast(this.h)
}

func (this *QCborArray) TakeFirst() *QCborValue {
	_goptr := newQCborValue(QCborArray_TakeFirst(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray) TakeLast() *QCborValue {
	_goptr := newQCborValue(QCborArray_TakeLast(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray) Contains(value *QCborValue) bool {
	return (bool)(QCborArray_Contains(this.h, value.cPointer()))
}

func (this *QCborArray) Compare(other *QCborArray) int {
	return (int)(QCborArray_Compare(this.h, other.cPointer()))
}

func (this *QCborArray) Begin() iterator {
	xxxxxxxxx
}

func (this *QCborArray) ConstBegin() const_iterator {
	xxxxxxxxx
}

func (this *QCborArray) Begin2() const_iterator {
	xxxxxxxxx
}

func (this *QCborArray) Cbegin() const_iterator {
	xxxxxxxxx
}

func (this *QCborArray) End() iterator {
	xxxxxxxxx
}

func (this *QCborArray) ConstEnd() const_iterator {
	xxxxxxxxx
}

func (this *QCborArray) End2() const_iterator {
	xxxxxxxxx
}

func (this *QCborArray) Cend() const_iterator {
	xxxxxxxxx
}

func (this *QCborArray) Insert2(before iterator, value *QCborValue) iterator {
	xxxxxxxxx
}

func (this *QCborArray) Insert3(before const_iterator, value *QCborValue) iterator {
	xxxxxxxxx
}

func (this *QCborArray) Erase(it iterator) iterator {
	xxxxxxxxx
}

func (this *QCborArray) EraseWithIt(it const_iterator) iterator {
	xxxxxxxxx
}

func (this *QCborArray) PushBack(t *QCborValue) {
	QCborArray_PushBack(this.h, t.cPointer())
}

func (this *QCborArray) PushFront(t *QCborValue) {
	QCborArray_PushFront(this.h, t.cPointer())
}

func (this *QCborArray) PopFront() {
	QCborArray_PopFront(this.h)
}

func (this *QCborArray) PopBack() {
	QCborArray_PopBack(this.h)
}

func (this *QCborArray) Empty() bool {
	return (bool)(QCborArray_Empty(this.h))
}

func (this *QCborArray) OperatorPlus(v *QCborValue) *QCborArray {
	_goptr := newQCborArray(QCborArray_OperatorPlus(this.h, v.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray) OperatorPlusAssign(v *QCborValue) *QCborArray {
	return newQCborArray(QCborArray_OperatorPlusAssign(this.h, v.cPointer()))
}

func (this *QCborArray) OperatorShiftLeft(v *QCborValue) *QCborArray {
	return newQCborArray(QCborArray_OperatorShiftLeft(this.h, v.cPointer()))
}

func QCborArray_FromStringList(list []string) *QCborArray {
	list_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(list))))
	defer free(unsafe.Pointer(list_CArray))
	for i := range list {
		list_i_ms := struct_miqt_string{}
		list_i_ms.data = CString(list[i])
		list_i_ms.len = size_t(len(list[i]))
		defer free(unsafe.Pointer(list_i_ms.data))
		list_CArray[i] = list_i_ms
	}
	list_ma := struct_miqt_array{len: size_t(len(list)), data: unsafe.Pointer(list_CArray)}
	_goptr := newQCborArray(QCborArray_FromStringList(list_ma))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QCborArray_FromJsonArray(array *qt.QJsonArray) *QCborArray {
	_goptr := newQCborArray(QCborArray_FromJsonArray((*QJsonArray)(array.UnsafePointer())))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray) ToJsonArray() *qt.QJsonArray {
	_goptr := qt.UnsafeNewQJsonArray(unsafe.Pointer(QCborArray_ToJsonArray(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QCborArray__Iterator struct {
	h          uintptr
	isSubclass bool
}

// NewQCborArray__Iterator constructs a new QCborArray::Iterator object.
func NewQCborArray__Iterator() *QCborArray__Iterator {
	ret := newQCborArray__Iterator(QCborArray__Iterator_new())
	ret.isSubclass = true
	return ret
}

// NewQCborArray__Iterator2 constructs a new QCborArray::Iterator object.
func NewQCborArray__Iterator2(param1 *Iterator) *QCborArray__Iterator {
	ret := newQCborArray__Iterator(QCborArray__Iterator_new2(param1))
	ret.isSubclass = true
	return ret
}

func (this *QCborArray__Iterator) OperatorAssign(other *Iterator) {
	QCborArray__Iterator_OperatorAssign(this.h, other)
}

func (this *QCborArray__Iterator) OperatorMultiply() *QCborValueRef {
	_goptr := newQCborValueRef(QCborArray__Iterator_OperatorMultiply(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray__Iterator) OperatorMinusGreater() *QCborValueRef {
	return newQCborValueRef(QCborArray__Iterator_OperatorMinusGreater(this.h))
}

func (this *QCborArray__Iterator) OperatorMinusGreater2() *QCborValueConstRef {
	return newQCborValueConstRef(QCborArray__Iterator_OperatorMinusGreater2(this.h))
}

func (this *QCborArray__Iterator) OperatorSubscript(j int64) *QCborValueRef {
	_goptr := newQCborValueRef(QCborArray__Iterator_OperatorSubscript(this.h, (ptrdiff_t)(j)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray__Iterator) OperatorPlusPlus() *Iterator {
	xxxxxxxxx
}

func (this *QCborArray__Iterator) OperatorPlusPlusWithInt(param1 int) Iterator {
	xxxxxxxxx
}

func (this *QCborArray__Iterator) OperatorMinusMinus() *Iterator {
	xxxxxxxxx
}

func (this *QCborArray__Iterator) OperatorMinusMinusWithInt(param1 int) Iterator {
	xxxxxxxxx
}

func (this *QCborArray__Iterator) OperatorPlusAssign(j int64) *Iterator {
	xxxxxxxxx
}

func (this *QCborArray__Iterator) OperatorMinusAssign(j int64) *Iterator {
	xxxxxxxxx
}

func (this *QCborArray__Iterator) OperatorPlus(j int64) Iterator {
	xxxxxxxxx
}

func (this *QCborArray__Iterator) OperatorMinus(j int64) Iterator {
	xxxxxxxxx
}

func (this *QCborArray__Iterator) OperatorMinusWithIterator(j Iterator) int64 {
	return (int64)(QCborArray__Iterator_OperatorMinusWithIterator(this.h, j))
}

type QCborArray__ConstIterator struct {
	h          uintptr
	isSubclass bool
}

// NewQCborArray__ConstIterator constructs a new QCborArray::ConstIterator object.
func NewQCborArray__ConstIterator() *QCborArray__ConstIterator {
	ret := newQCborArray__ConstIterator(QCborArray__ConstIterator_new())
	ret.isSubclass = true
	return ret
}

// NewQCborArray__ConstIterator2 constructs a new QCborArray::ConstIterator object.
func NewQCborArray__ConstIterator2(param1 *ConstIterator) *QCborArray__ConstIterator {
	ret := newQCborArray__ConstIterator(QCborArray__ConstIterator_new2(param1))
	ret.isSubclass = true
	return ret
}

func (this *QCborArray__ConstIterator) OperatorAssign(other *ConstIterator) {
	QCborArray__ConstIterator_OperatorAssign(this.h, other)
}

func (this *QCborArray__ConstIterator) OperatorMultiply() *QCborValueConstRef {
	_goptr := newQCborValueConstRef(QCborArray__ConstIterator_OperatorMultiply(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray__ConstIterator) OperatorMinusGreater() *QCborValueConstRef {
	return newQCborValueConstRef(QCborArray__ConstIterator_OperatorMinusGreater(this.h))
}

func (this *QCborArray__ConstIterator) OperatorSubscript(j int64) *QCborValueConstRef {
	_goptr := newQCborValueConstRef(QCborArray__ConstIterator_OperatorSubscript(this.h, (ptrdiff_t)(j)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborArray__ConstIterator) OperatorPlusPlus() *ConstIterator {
	xxxxxxxxx
}

func (this *QCborArray__ConstIterator) OperatorPlusPlusWithInt(param1 int) ConstIterator {
	xxxxxxxxx
}

func (this *QCborArray__ConstIterator) OperatorMinusMinus() *ConstIterator {
	xxxxxxxxx
}

func (this *QCborArray__ConstIterator) OperatorMinusMinusWithInt(param1 int) ConstIterator {
	xxxxxxxxx
}

func (this *QCborArray__ConstIterator) OperatorPlusAssign(j int64) *ConstIterator {
	xxxxxxxxx
}

func (this *QCborArray__ConstIterator) OperatorMinusAssign(j int64) *ConstIterator {
	xxxxxxxxx
}

func (this *QCborArray__ConstIterator) OperatorPlus(j int64) ConstIterator {
	xxxxxxxxx
}

func (this *QCborArray__ConstIterator) OperatorMinus(j int64) ConstIterator {
	xxxxxxxxx
}

func (this *QCborArray__ConstIterator) OperatorMinusWithConstIterator(j ConstIterator) int64 {
	return (int64)(QCborArray__ConstIterator_OperatorMinusWithConstIterator(this.h, j))
}
