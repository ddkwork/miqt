package cbor

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QCborMap struct {
	h          uintptr
	isSubclass bool
}

// NewQCborMap constructs a new QCborMap object.
func NewQCborMap() *QCborMap {

	ret := newQCborMap(QCborMap_new())
	ret.isSubclass = true
	return ret
}

// NewQCborMap2 constructs a new QCborMap object.
func NewQCborMap2(other *QCborMap) *QCborMap {

	ret := newQCborMap(QCborMap_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QCborMap) OperatorAssign(other *QCborMap) {
	QCborMap_OperatorAssign(this.h, other.cPointer())
}

func (this *QCborMap) Swap(other *QCborMap) {
	QCborMap_Swap(this.h, other.cPointer())
}

func (this *QCborMap) ToCborValue() *QCborValue {
	_goptr := newQCborValue(QCborMap_ToCborValue(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Size() int64 {
	return (int64)(QCborMap_Size(this.h))
}

func (this *QCborMap) IsEmpty() bool {
	return (bool)(QCborMap_IsEmpty(this.h))
}

func (this *QCborMap) Clear() {
	QCborMap_Clear(this.h)
}

func (this *QCborMap) Keys() []QCborValue {
	var _ma struct_miqt_array = QCborMap_Keys(this.h)
	_ret := make([]QCborValue, int(_ma.len))
	_outCast := (*[0xffff]*QCborValue)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQCborValue(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QCborMap) Value(key int64) *QCborValue {
	_goptr := newQCborValue(QCborMap_Value(this.h, (longlong)(key)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Value2(key string) *QCborValue {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQCborValue(QCborMap_Value2(this.h, key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Value3(key *QCborValue) *QCborValue {
	_goptr := newQCborValue(QCborMap_Value3(this.h, key.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) OperatorSubscript(key int64) *QCborValue {
	_goptr := newQCborValue(QCborMap_OperatorSubscript(this.h, (longlong)(key)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) OperatorSubscript2(key string) *QCborValue {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQCborValue(QCborMap_OperatorSubscript2(this.h, key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) OperatorSubscript3(key *QCborValue) *QCborValue {
	_goptr := newQCborValue(QCborMap_OperatorSubscript3(this.h, key.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) OperatorSubscript4(key int64) *QCborValueRef {
	_goptr := newQCborValueRef(QCborMap_OperatorSubscript4(this.h, (longlong)(key)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) OperatorSubscript6(key string) *QCborValueRef {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQCborValueRef(QCborMap_OperatorSubscript6(this.h, key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) OperatorSubscript7(key *QCborValue) *QCborValueRef {
	_goptr := newQCborValueRef(QCborMap_OperatorSubscript7(this.h, key.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Take(key int64) *QCborValue {
	_goptr := newQCborValue(QCborMap_Take(this.h, (longlong)(key)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Take2(key string) *QCborValue {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQCborValue(QCborMap_Take2(this.h, key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Take3(key *QCborValue) *QCborValue {
	_goptr := newQCborValue(QCborMap_Take3(this.h, key.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Remove(key int64) {
	QCborMap_Remove(this.h, (longlong)(key))
}

func (this *QCborMap) Remove2(key string) {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	QCborMap_Remove2(this.h, key_ms)
}

func (this *QCborMap) Remove3(key *QCborValue) {
	QCborMap_Remove3(this.h, key.cPointer())
}

func (this *QCborMap) Contains(key int64) bool {
	return (bool)(QCborMap_Contains(this.h, (longlong)(key)))
}

func (this *QCborMap) Contains2(key string) bool {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	return (bool)(QCborMap_Contains2(this.h, key_ms))
}

func (this *QCborMap) Contains3(key *QCborValue) bool {
	return (bool)(QCborMap_Contains3(this.h, key.cPointer()))
}

func (this *QCborMap) Compare(other *QCborMap) int {
	return (int)(QCborMap_Compare(this.h, other.cPointer()))
}

func (this *QCborMap) Begin() iterator {
	xxxxxxxxx
}

func (this *QCborMap) ConstBegin() const_iterator {
	xxxxxxxxx
}

func (this *QCborMap) Begin2() const_iterator {
	xxxxxxxxx
}

func (this *QCborMap) Cbegin() const_iterator {
	xxxxxxxxx
}

func (this *QCborMap) End() iterator {
	xxxxxxxxx
}

func (this *QCborMap) ConstEnd() const_iterator {
	xxxxxxxxx
}

func (this *QCborMap) End2() const_iterator {
	xxxxxxxxx
}

func (this *QCborMap) Cend() const_iterator {
	xxxxxxxxx
}

func (this *QCborMap) Erase(it iterator) iterator {
	xxxxxxxxx
}

func (this *QCborMap) EraseWithIt(it const_iterator) iterator {
	xxxxxxxxx
}

func (this *QCborMap) Extract(it iterator) *QCborValue {
	_goptr := newQCborValue(QCborMap_Extract(this.h, it))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) ExtractWithIt(it const_iterator) *QCborValue {
	_goptr := newQCborValue(QCborMap_ExtractWithIt(this.h, it))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) Empty() bool {
	return (bool)(QCborMap_Empty(this.h))
}

func (this *QCborMap) Find(key int64) iterator {
	xxxxxxxxx
}

func (this *QCborMap) Find2(key string) iterator {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	xxxxxxxxx
}

func (this *QCborMap) Find3(key *QCborValue) iterator {
	xxxxxxxxx
}

func (this *QCborMap) ConstFind(key int64) const_iterator {
	xxxxxxxxx
}

func (this *QCborMap) ConstFind2(key string) const_iterator {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	xxxxxxxxx
}

func (this *QCborMap) ConstFind3(key *QCborValue) const_iterator {
	xxxxxxxxx
}

func (this *QCborMap) Find4(key int64) const_iterator {
	xxxxxxxxx
}

func (this *QCborMap) Find6(key string) const_iterator {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	xxxxxxxxx
}

func (this *QCborMap) Find7(key *QCborValue) const_iterator {
	xxxxxxxxx
}

func (this *QCborMap) Insert(key int64, value_ *QCborValue) iterator {
	xxxxxxxxx
}

func (this *QCborMap) Insert3(key string, value_ *QCborValue) iterator {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	xxxxxxxxx
}

func (this *QCborMap) Insert4(key *QCborValue, value_ *QCborValue) iterator {
	xxxxxxxxx
}

func (this *QCborMap) InsertWithValueType(v value_type) iterator {
	xxxxxxxxx
}

func QCborMap_FromVariantMap(mapVal map[string]qt.QVariant) *QCborMap {
	mapVal_Keys_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(mapVal))))
	defer free(unsafe.Pointer(mapVal_Keys_CArray))
	mapVal_Values_CArray := (*[0xffff]*QVariant)(malloc(size_t(8 * len(mapVal))))
	defer free(unsafe.Pointer(mapVal_Values_CArray))
	mapVal_ctr := 0
	for mapVal_k, mapVal_v := range mapVal {
		mapVal_k_ms := struct_miqt_string{}
		mapVal_k_ms.data = CString(mapVal_k)
		mapVal_k_ms.len = size_t(len(mapVal_k))
		defer free(unsafe.Pointer(mapVal_k_ms.data))
		mapVal_Keys_CArray[mapVal_ctr] = mapVal_k_ms
		mapVal_Values_CArray[mapVal_ctr] = (*QVariant)(mapVal_v.UnsafePointer())
		mapVal_ctr++
	}
	mapVal_mm := struct_miqt_map{
		len:    size_t(len(mapVal)),
		keys:   unsafe.Pointer(mapVal_Keys_CArray),
		values: unsafe.Pointer(mapVal_Values_CArray),
	}
	_goptr := newQCborMap(QCborMap_FromVariantMap(mapVal_mm))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QCborMap_FromVariantHash(hash map[string]qt.QVariant) *QCborMap {
	hash_Keys_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(hash))))
	defer free(unsafe.Pointer(hash_Keys_CArray))
	hash_Values_CArray := (*[0xffff]*QVariant)(malloc(size_t(8 * len(hash))))
	defer free(unsafe.Pointer(hash_Values_CArray))
	hash_ctr := 0
	for hash_k, hash_v := range hash {
		hash_k_ms := struct_miqt_string{}
		hash_k_ms.data = CString(hash_k)
		hash_k_ms.len = size_t(len(hash_k))
		defer free(unsafe.Pointer(hash_k_ms.data))
		hash_Keys_CArray[hash_ctr] = hash_k_ms
		hash_Values_CArray[hash_ctr] = (*QVariant)(hash_v.UnsafePointer())
		hash_ctr++
	}
	hash_mm := struct_miqt_map{
		len:    size_t(len(hash)),
		keys:   unsafe.Pointer(hash_Keys_CArray),
		values: unsafe.Pointer(hash_Values_CArray),
	}
	_goptr := newQCborMap(QCborMap_FromVariantHash(hash_mm))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QCborMap_FromJsonObject(o *qt.QJsonObject) *QCborMap {
	_goptr := newQCborMap(QCborMap_FromJsonObject((*QJsonObject)(o.UnsafePointer())))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap) ToVariantMap() map[string]qt.QVariant {
	var _mm struct_miqt_map = QCborMap_ToVariantMap(this.h)
	_ret := make(map[string]qt.QVariant, int(_mm.len))
	_Keys := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_mm.keys))
	_Values := (*[0xffff]*QVariant)(unsafe.Pointer(_mm.values))
	for i := 0; i < int(_mm.len); i++ {
		var _mapkey_ms struct_miqt_string = _Keys[i]
		_mapkey_ret := GoStringN(_mapkey_ms.data, int(int64(_mapkey_ms.len)))
		free(unsafe.Pointer(_mapkey_ms.data))
		_entry_Key := _mapkey_ret
		_mapval_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(_Values[i]))
		_mapval_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_entry_Value := *_mapval_goptr

		_ret[_entry_Key] = _entry_Value
	}
	return _ret
}

func (this *QCborMap) ToVariantHash() map[string]qt.QVariant {
	var _mm struct_miqt_map = QCborMap_ToVariantHash(this.h)
	_ret := make(map[string]qt.QVariant, int(_mm.len))
	_Keys := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_mm.keys))
	_Values := (*[0xffff]*QVariant)(unsafe.Pointer(_mm.values))
	for i := 0; i < int(_mm.len); i++ {
		var _hashkey_ms struct_miqt_string = _Keys[i]
		_hashkey_ret := GoStringN(_hashkey_ms.data, int(int64(_hashkey_ms.len)))
		free(unsafe.Pointer(_hashkey_ms.data))
		_entry_Key := _hashkey_ret
		_hashval_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(_Values[i]))
		_hashval_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_entry_Value := *_hashval_goptr

		_ret[_entry_Key] = _entry_Value
	}
	return _ret
}

func (this *QCborMap) ToJsonObject() *qt.QJsonObject {
	_goptr := qt.UnsafeNewQJsonObject(unsafe.Pointer(QCborMap_ToJsonObject(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QCborMap__Iterator struct {
	h          uintptr
	isSubclass bool
}

// NewQCborMap__Iterator constructs a new QCborMap::Iterator object.
func NewQCborMap__Iterator() *QCborMap__Iterator {

	ret := newQCborMap__Iterator(QCborMap__Iterator_new())
	ret.isSubclass = true
	return ret
}

// NewQCborMap__Iterator2 constructs a new QCborMap::Iterator object.
func NewQCborMap__Iterator2(param1 *Iterator) *QCborMap__Iterator {

	ret := newQCborMap__Iterator(QCborMap__Iterator_new2(param1))
	ret.isSubclass = true
	return ret
}

func (this *QCborMap__Iterator) OperatorAssign(other *Iterator) {
	QCborMap__Iterator_OperatorAssign(this.h, other)
}

func (this *QCborMap__Iterator) OperatorMultiply() value_type {
	xxxxxxxxx
}

func (this *QCborMap__Iterator) OperatorSubscript(j int64) value_type {
	xxxxxxxxx
}

func (this *QCborMap__Iterator) OperatorMinusGreater() *QCborValueRef {
	return newQCborValueRef(QCborMap__Iterator_OperatorMinusGreater(this.h))
}

func (this *QCborMap__Iterator) OperatorMinusGreater2() *QCborValueConstRef {
	return newQCborValueConstRef(QCborMap__Iterator_OperatorMinusGreater2(this.h))
}

func (this *QCborMap__Iterator) Key() *QCborValue {
	_goptr := newQCborValue(QCborMap__Iterator_Key(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap__Iterator) Value() *QCborValueRef {
	_goptr := newQCborValueRef(QCborMap__Iterator_Value(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap__Iterator) OperatorPlusPlus() *Iterator {
	xxxxxxxxx
}

func (this *QCborMap__Iterator) OperatorPlusPlusWithInt(param1 int) Iterator {
	xxxxxxxxx
}

func (this *QCborMap__Iterator) OperatorMinusMinus() *Iterator {
	xxxxxxxxx
}

func (this *QCborMap__Iterator) OperatorMinusMinusWithInt(param1 int) Iterator {
	xxxxxxxxx
}

func (this *QCborMap__Iterator) OperatorPlusAssign(j int64) *Iterator {
	xxxxxxxxx
}

func (this *QCborMap__Iterator) OperatorMinusAssign(j int64) *Iterator {
	xxxxxxxxx
}

func (this *QCborMap__Iterator) OperatorPlus(j int64) Iterator {
	xxxxxxxxx
}

func (this *QCborMap__Iterator) OperatorMinus(j int64) Iterator {
	xxxxxxxxx
}

func (this *QCborMap__Iterator) OperatorMinusWithIterator(j Iterator) int64 {
	return (int64)(QCborMap__Iterator_OperatorMinusWithIterator(this.h, j))
}

type QCborMap__ConstIterator struct {
	h          uintptr
	isSubclass bool
}

// NewQCborMap__ConstIterator constructs a new QCborMap::ConstIterator object.
func NewQCborMap__ConstIterator() *QCborMap__ConstIterator {

	ret := newQCborMap__ConstIterator(QCborMap__ConstIterator_new())
	ret.isSubclass = true
	return ret
}

// NewQCborMap__ConstIterator2 constructs a new QCborMap::ConstIterator object.
func NewQCborMap__ConstIterator2(param1 *ConstIterator) *QCborMap__ConstIterator {

	ret := newQCborMap__ConstIterator(QCborMap__ConstIterator_new2(param1))
	ret.isSubclass = true
	return ret
}

func (this *QCborMap__ConstIterator) OperatorAssign(other *ConstIterator) {
	QCborMap__ConstIterator_OperatorAssign(this.h, other)
}

func (this *QCborMap__ConstIterator) OperatorMultiply() value_type {
	xxxxxxxxx
}

func (this *QCborMap__ConstIterator) OperatorSubscript(j int64) value_type {
	xxxxxxxxx
}

func (this *QCborMap__ConstIterator) OperatorMinusGreater() *QCborValueConstRef {
	return newQCborValueConstRef(QCborMap__ConstIterator_OperatorMinusGreater(this.h))
}

func (this *QCborMap__ConstIterator) Key() *QCborValue {
	_goptr := newQCborValue(QCborMap__ConstIterator_Key(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap__ConstIterator) Value() *QCborValueConstRef {
	_goptr := newQCborValueConstRef(QCborMap__ConstIterator_Value(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborMap__ConstIterator) OperatorPlusPlus() *ConstIterator {
	xxxxxxxxx
}

func (this *QCborMap__ConstIterator) OperatorPlusPlusWithInt(param1 int) ConstIterator {
	xxxxxxxxx
}

func (this *QCborMap__ConstIterator) OperatorMinusMinus() *ConstIterator {
	xxxxxxxxx
}

func (this *QCborMap__ConstIterator) OperatorMinusMinusWithInt(param1 int) ConstIterator {
	xxxxxxxxx
}

func (this *QCborMap__ConstIterator) OperatorPlusAssign(j int64) *ConstIterator {
	xxxxxxxxx
}

func (this *QCborMap__ConstIterator) OperatorMinusAssign(j int64) *ConstIterator {
	xxxxxxxxx
}

func (this *QCborMap__ConstIterator) OperatorPlus(j int64) ConstIterator {
	xxxxxxxxx
}

func (this *QCborMap__ConstIterator) OperatorMinus(j int64) ConstIterator {
	xxxxxxxxx
}

func (this *QCborMap__ConstIterator) OperatorMinusWithConstIterator(j ConstIterator) int64 {
	return (int64)(QCborMap__ConstIterator_OperatorMinusWithConstIterator(this.h, j))
}
