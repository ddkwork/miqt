package qt

import (
	"unsafe"
)

type QJsonArray struct {
	h          uintptr
	isSubclass bool
}

// NewQJsonArray constructs a new QJsonArray object.
func NewQJsonArray() *QJsonArray {

	ret := newQJsonArray(QJsonArray_new())
	ret.isSubclass = true
	return ret
}

// NewQJsonArray2 constructs a new QJsonArray object.
func NewQJsonArray2(other *QJsonArray) *QJsonArray {

	ret := newQJsonArray(QJsonArray_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QJsonArray) OperatorAssign(other *QJsonArray) {
	QJsonArray_OperatorAssign(this.h, other.cPointer())
}

func QJsonArray_FromStringList(list []string) *QJsonArray {
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
	_goptr := newQJsonArray(QJsonArray_FromStringList(list_ma))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonArray) Size() int64 {
	return (int64)(QJsonArray_Size(this.h))
}

func (this *QJsonArray) Count() int64 {
	return (int64)(QJsonArray_Count(this.h))
}

func (this *QJsonArray) IsEmpty() bool {
	return (bool)(QJsonArray_IsEmpty(this.h))
}

func (this *QJsonArray) At(i int64) *QJsonValue {
	_goptr := newQJsonValue(QJsonArray_At(this.h, (ptrdiff_t)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonArray) First() *QJsonValue {
	_goptr := newQJsonValue(QJsonArray_First(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonArray) Last() *QJsonValue {
	_goptr := newQJsonValue(QJsonArray_Last(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonArray) Prepend(value *QJsonValue) {
	QJsonArray_Prepend(this.h, value.cPointer())
}

func (this *QJsonArray) Append(value *QJsonValue) {
	QJsonArray_Append(this.h, value.cPointer())
}

func (this *QJsonArray) RemoveAt(i int64) {
	QJsonArray_RemoveAt(this.h, (ptrdiff_t)(i))
}

func (this *QJsonArray) TakeAt(i int64) *QJsonValue {
	_goptr := newQJsonValue(QJsonArray_TakeAt(this.h, (ptrdiff_t)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonArray) RemoveFirst() {
	QJsonArray_RemoveFirst(this.h)
}

func (this *QJsonArray) RemoveLast() {
	QJsonArray_RemoveLast(this.h)
}

func (this *QJsonArray) Insert(i int64, value *QJsonValue) {
	QJsonArray_Insert(this.h, (ptrdiff_t)(i), value.cPointer())
}

func (this *QJsonArray) Replace(i int64, value *QJsonValue) {
	QJsonArray_Replace(this.h, (ptrdiff_t)(i), value.cPointer())
}

func (this *QJsonArray) Contains(element *QJsonValue) bool {
	return (bool)(QJsonArray_Contains(this.h, element.cPointer()))
}

func (this *QJsonArray) OperatorSubscript(i int64) *QJsonValueRef {
	_goptr := newQJsonValueRef(QJsonArray_OperatorSubscript(this.h, (ptrdiff_t)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonArray) OperatorSubscriptWithQsizetype(i int64) *QJsonValue {
	_goptr := newQJsonValue(QJsonArray_OperatorSubscriptWithQsizetype(this.h, (ptrdiff_t)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonArray) Swap(other *QJsonArray) {
	QJsonArray_Swap(this.h, other.cPointer())
}

func (this *QJsonArray) Begin() iterator {
	xxxxxxxxx
}

func (this *QJsonArray) Begin2() const_iterator {
	xxxxxxxxx
}

func (this *QJsonArray) ConstBegin() const_iterator {
	xxxxxxxxx
}

func (this *QJsonArray) Cbegin() const_iterator {
	xxxxxxxxx
}

func (this *QJsonArray) End() iterator {
	xxxxxxxxx
}

func (this *QJsonArray) End2() const_iterator {
	xxxxxxxxx
}

func (this *QJsonArray) ConstEnd() const_iterator {
	xxxxxxxxx
}

func (this *QJsonArray) Cend() const_iterator {
	xxxxxxxxx
}

func (this *QJsonArray) Insert2(before iterator, value *QJsonValue) iterator {
	xxxxxxxxx
}

func (this *QJsonArray) Erase(it iterator) iterator {
	xxxxxxxxx
}

func (this *QJsonArray) OperatorPlus(v *QJsonValue) *QJsonArray {
	_goptr := newQJsonArray(QJsonArray_OperatorPlus(this.h, v.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonArray) OperatorPlusAssign(v *QJsonValue) *QJsonArray {
	return newQJsonArray(QJsonArray_OperatorPlusAssign(this.h, v.cPointer()))
}

func (this *QJsonArray) OperatorShiftLeft(v *QJsonValue) *QJsonArray {
	return newQJsonArray(QJsonArray_OperatorShiftLeft(this.h, v.cPointer()))
}

func (this *QJsonArray) PushBack(t *QJsonValue) {
	QJsonArray_PushBack(this.h, t.cPointer())
}

func (this *QJsonArray) PushFront(t *QJsonValue) {
	QJsonArray_PushFront(this.h, t.cPointer())
}

func (this *QJsonArray) PopFront() {
	QJsonArray_PopFront(this.h)
}

func (this *QJsonArray) PopBack() {
	QJsonArray_PopBack(this.h)
}

func (this *QJsonArray) Empty() bool {
	return (bool)(QJsonArray_Empty(this.h))
}

type QJsonArray__iterator struct {
	h          uintptr
	isSubclass bool
}

// NewQJsonArray__iterator constructs a new QJsonArray::iterator object.
func NewQJsonArray__iterator() *QJsonArray__iterator {

	ret := newQJsonArray__iterator(QJsonArray__iterator_new())
	ret.isSubclass = true
	return ret
}

// NewQJsonArray__iterator2 constructs a new QJsonArray::iterator object.
func NewQJsonArray__iterator2(array *QJsonArray, index int64) *QJsonArray__iterator {

	ret := newQJsonArray__iterator(QJsonArray__iterator_new2(array.cPointer(), (ptrdiff_t)(index)))
	ret.isSubclass = true
	return ret
}

// NewQJsonArray__iterator3 constructs a new QJsonArray::iterator object.
func NewQJsonArray__iterator3(other *iterator) *QJsonArray__iterator {

	ret := newQJsonArray__iterator(QJsonArray__iterator_new3(other))
	ret.isSubclass = true
	return ret
}

func (this *QJsonArray__iterator) OperatorAssign(other *iterator) {
	QJsonArray__iterator_OperatorAssign(this.h, other)
}

func (this *QJsonArray__iterator) OperatorMultiply() *QJsonValueRef {
	_goptr := newQJsonValueRef(QJsonArray__iterator_OperatorMultiply(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonArray__iterator) OperatorMinusGreater() *QJsonValueConstRef {
	return newQJsonValueConstRef(QJsonArray__iterator_OperatorMinusGreater(this.h))
}

func (this *QJsonArray__iterator) OperatorMinusGreater2() *QJsonValueRef {
	return newQJsonValueRef(QJsonArray__iterator_OperatorMinusGreater2(this.h))
}

func (this *QJsonArray__iterator) OperatorSubscript(j int64) *QJsonValueRef {
	_goptr := newQJsonValueRef(QJsonArray__iterator_OperatorSubscript(this.h, (ptrdiff_t)(j)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonArray__iterator) OperatorPlusPlus() *iterator {
	xxxxxxxxx
}

func (this *QJsonArray__iterator) OperatorPlusPlusWithInt(param1 int) iterator {
	xxxxxxxxx
}

func (this *QJsonArray__iterator) OperatorMinusMinus() *iterator {
	xxxxxxxxx
}

func (this *QJsonArray__iterator) OperatorMinusMinusWithInt(param1 int) iterator {
	xxxxxxxxx
}

func (this *QJsonArray__iterator) OperatorPlusAssign(j int64) *iterator {
	xxxxxxxxx
}

func (this *QJsonArray__iterator) OperatorMinusAssign(j int64) *iterator {
	xxxxxxxxx
}

func (this *QJsonArray__iterator) OperatorPlus(j int64) iterator {
	xxxxxxxxx
}

func (this *QJsonArray__iterator) OperatorMinus(j int64) iterator {
	xxxxxxxxx
}

func (this *QJsonArray__iterator) OperatorMinusWithIterator(j iterator) int64 {
	return (int64)(QJsonArray__iterator_OperatorMinusWithIterator(this.h, j))
}

type QJsonArray__const_iterator struct {
	h          uintptr
	isSubclass bool
}

// NewQJsonArray__const_iterator constructs a new QJsonArray::const_iterator object.
func NewQJsonArray__const_iterator() *QJsonArray__const_iterator {

	ret := newQJsonArray__const_iterator(QJsonArray__const_iterator_new())
	ret.isSubclass = true
	return ret
}

// NewQJsonArray__const_iterator2 constructs a new QJsonArray::const_iterator object.
func NewQJsonArray__const_iterator2(array *QJsonArray, index int64) *QJsonArray__const_iterator {

	ret := newQJsonArray__const_iterator(QJsonArray__const_iterator_new2(array.cPointer(), (ptrdiff_t)(index)))
	ret.isSubclass = true
	return ret
}

// NewQJsonArray__const_iterator3 constructs a new QJsonArray::const_iterator object.
func NewQJsonArray__const_iterator3(o *iterator) *QJsonArray__const_iterator {

	ret := newQJsonArray__const_iterator(QJsonArray__const_iterator_new3(o))
	ret.isSubclass = true
	return ret
}

// NewQJsonArray__const_iterator4 constructs a new QJsonArray::const_iterator object.
func NewQJsonArray__const_iterator4(other *const_iterator) *QJsonArray__const_iterator {

	ret := newQJsonArray__const_iterator(QJsonArray__const_iterator_new4(other))
	ret.isSubclass = true
	return ret
}

func (this *QJsonArray__const_iterator) OperatorAssign(other *const_iterator) {
	QJsonArray__const_iterator_OperatorAssign(this.h, other)
}

func (this *QJsonArray__const_iterator) OperatorMultiply() *QJsonValueConstRef {
	_goptr := newQJsonValueConstRef(QJsonArray__const_iterator_OperatorMultiply(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonArray__const_iterator) OperatorMinusGreater() *QJsonValueConstRef {
	return newQJsonValueConstRef(QJsonArray__const_iterator_OperatorMinusGreater(this.h))
}

func (this *QJsonArray__const_iterator) OperatorSubscript(j int64) *QJsonValueConstRef {
	_goptr := newQJsonValueConstRef(QJsonArray__const_iterator_OperatorSubscript(this.h, (ptrdiff_t)(j)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonArray__const_iterator) OperatorPlusPlus() *const_iterator {
	xxxxxxxxx
}

func (this *QJsonArray__const_iterator) OperatorPlusPlusWithInt(param1 int) const_iterator {
	xxxxxxxxx
}

func (this *QJsonArray__const_iterator) OperatorMinusMinus() *const_iterator {
	xxxxxxxxx
}

func (this *QJsonArray__const_iterator) OperatorMinusMinusWithInt(param1 int) const_iterator {
	xxxxxxxxx
}

func (this *QJsonArray__const_iterator) OperatorPlusAssign(j int64) *const_iterator {
	xxxxxxxxx
}

func (this *QJsonArray__const_iterator) OperatorMinusAssign(j int64) *const_iterator {
	xxxxxxxxx
}

func (this *QJsonArray__const_iterator) OperatorPlus(j int64) const_iterator {
	xxxxxxxxx
}

func (this *QJsonArray__const_iterator) OperatorMinus(j int64) const_iterator {
	xxxxxxxxx
}

func (this *QJsonArray__const_iterator) OperatorMinusWithConstIterator(j const_iterator) int64 {
	return (int64)(QJsonArray__const_iterator_OperatorMinusWithConstIterator(this.h, j))
}
