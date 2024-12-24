package qt

import (
	"unsafe"
)

type QJsonObject struct {
	h          uintptr
	isSubclass bool
}

// NewQJsonObject constructs a new QJsonObject object.
func NewQJsonObject() *QJsonObject {
	ret := newQJsonObject(QJsonObject_new())
	ret.isSubclass = true
	return ret
}

// NewQJsonObject2 constructs a new QJsonObject object.
func NewQJsonObject2(other *QJsonObject) *QJsonObject {
	ret := newQJsonObject(QJsonObject_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QJsonObject) OperatorAssign(other *QJsonObject) {
	QJsonObject_OperatorAssign(this.h, other.cPointer())
}

func (this *QJsonObject) Swap(other *QJsonObject) {
	QJsonObject_Swap(this.h, other.cPointer())
}

func QJsonObject_FromVariantMap(mapVal map[string]QVariant) *QJsonObject {
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
		mapVal_Values_CArray[mapVal_ctr] = mapVal_v.cPointer()
		mapVal_ctr++
	}
	mapVal_mm := struct_miqt_map{
		len:    size_t(len(mapVal)),
		keys:   unsafe.Pointer(mapVal_Keys_CArray),
		values: unsafe.Pointer(mapVal_Values_CArray),
	}
	_goptr := newQJsonObject(QJsonObject_FromVariantMap(mapVal_mm))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonObject) ToVariantMap() map[string]QVariant {
	var _mm struct_miqt_map = QJsonObject_ToVariantMap(this.h)
	_ret := make(map[string]QVariant, int(_mm.len))
	_Keys := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_mm.keys))
	_Values := (*[0xffff]*QVariant)(unsafe.Pointer(_mm.values))
	for i := 0; i < int(_mm.len); i++ {
		var _mapkey_ms struct_miqt_string = _Keys[i]
		_mapkey_ret := GoStringN(_mapkey_ms.data, int(int64(_mapkey_ms.len)))
		free(unsafe.Pointer(_mapkey_ms.data))
		_entry_Key := _mapkey_ret
		_mapval_goptr := newQVariant(_Values[i])
		_mapval_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_entry_Value := *_mapval_goptr

		_ret[_entry_Key] = _entry_Value
	}
	return _ret
}

func QJsonObject_FromVariantHash(mapVal map[string]QVariant) *QJsonObject {
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
		mapVal_Values_CArray[mapVal_ctr] = mapVal_v.cPointer()
		mapVal_ctr++
	}
	mapVal_mm := struct_miqt_map{
		len:    size_t(len(mapVal)),
		keys:   unsafe.Pointer(mapVal_Keys_CArray),
		values: unsafe.Pointer(mapVal_Values_CArray),
	}
	_goptr := newQJsonObject(QJsonObject_FromVariantHash(mapVal_mm))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonObject) ToVariantHash() map[string]QVariant {
	var _mm struct_miqt_map = QJsonObject_ToVariantHash(this.h)
	_ret := make(map[string]QVariant, int(_mm.len))
	_Keys := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_mm.keys))
	_Values := (*[0xffff]*QVariant)(unsafe.Pointer(_mm.values))
	for i := 0; i < int(_mm.len); i++ {
		var _hashkey_ms struct_miqt_string = _Keys[i]
		_hashkey_ret := GoStringN(_hashkey_ms.data, int(int64(_hashkey_ms.len)))
		free(unsafe.Pointer(_hashkey_ms.data))
		_entry_Key := _hashkey_ret
		_hashval_goptr := newQVariant(_Values[i])
		_hashval_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_entry_Value := *_hashval_goptr

		_ret[_entry_Key] = _entry_Value
	}
	return _ret
}

func (this *QJsonObject) Keys() []string {
	var _ma struct_miqt_array = QJsonObject_Keys(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QJsonObject) Size() int64 {
	return (int64)(QJsonObject_Size(this.h))
}

func (this *QJsonObject) Count() int64 {
	return (int64)(QJsonObject_Count(this.h))
}

func (this *QJsonObject) Length() int64 {
	return (int64)(QJsonObject_Length(this.h))
}

func (this *QJsonObject) IsEmpty() bool {
	return (bool)(QJsonObject_IsEmpty(this.h))
}

func (this *QJsonObject) Value(key string) *QJsonValue {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQJsonValue(QJsonObject_Value(this.h, key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonObject) OperatorSubscript(key string) *QJsonValue {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQJsonValue(QJsonObject_OperatorSubscript(this.h, key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonObject) OperatorSubscriptWithKey(key string) *QJsonValueRef {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQJsonValueRef(QJsonObject_OperatorSubscriptWithKey(this.h, key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonObject) Remove(key string) {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	QJsonObject_Remove(this.h, key_ms)
}

func (this *QJsonObject) Take(key string) *QJsonValue {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQJsonValue(QJsonObject_Take(this.h, key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonObject) Contains(key string) bool {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	return (bool)(QJsonObject_Contains(this.h, key_ms))
}

func (this *QJsonObject) Begin() iterator {
	xxxxxxxxx
}

func (this *QJsonObject) Begin2() const_iterator {
	xxxxxxxxx
}

func (this *QJsonObject) ConstBegin() const_iterator {
	xxxxxxxxx
}

func (this *QJsonObject) End() iterator {
	xxxxxxxxx
}

func (this *QJsonObject) End2() const_iterator {
	xxxxxxxxx
}

func (this *QJsonObject) ConstEnd() const_iterator {
	xxxxxxxxx
}

func (this *QJsonObject) Erase(it iterator) iterator {
	xxxxxxxxx
}

func (this *QJsonObject) Find(key string) iterator {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	xxxxxxxxx
}

func (this *QJsonObject) FindWithKey(key string) const_iterator {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	xxxxxxxxx
}

func (this *QJsonObject) ConstFind(key string) const_iterator {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	xxxxxxxxx
}

func (this *QJsonObject) Insert(key string, value *QJsonValue) iterator {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	xxxxxxxxx
}

func (this *QJsonObject) Empty() bool {
	return (bool)(QJsonObject_Empty(this.h))
}

type QJsonObject__iterator struct {
	h          uintptr
	isSubclass bool
}

// NewQJsonObject__iterator constructs a new QJsonObject::iterator object.
func NewQJsonObject__iterator() *QJsonObject__iterator {
	ret := newQJsonObject__iterator(QJsonObject__iterator_new())
	ret.isSubclass = true
	return ret
}

// NewQJsonObject__iterator2 constructs a new QJsonObject::iterator object.
func NewQJsonObject__iterator2(obj *QJsonObject, index int64) *QJsonObject__iterator {
	ret := newQJsonObject__iterator(QJsonObject__iterator_new2(obj.cPointer(), (ptrdiff_t)(index)))
	ret.isSubclass = true
	return ret
}

// NewQJsonObject__iterator3 constructs a new QJsonObject::iterator object.
func NewQJsonObject__iterator3(other *iterator) *QJsonObject__iterator {
	ret := newQJsonObject__iterator(QJsonObject__iterator_new3(other))
	ret.isSubclass = true
	return ret
}

func (this *QJsonObject__iterator) OperatorAssign(other *iterator) {
	QJsonObject__iterator_OperatorAssign(this.h, other)
}

func (this *QJsonObject__iterator) Key() string {
	var _ms struct_miqt_string = QJsonObject__iterator_Key(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QJsonObject__iterator) Value() *QJsonValueRef {
	_goptr := newQJsonValueRef(QJsonObject__iterator_Value(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonObject__iterator) OperatorMultiply() *QJsonValueRef {
	_goptr := newQJsonValueRef(QJsonObject__iterator_OperatorMultiply(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonObject__iterator) OperatorMinusGreater() *QJsonValueConstRef {
	return newQJsonValueConstRef(QJsonObject__iterator_OperatorMinusGreater(this.h))
}

func (this *QJsonObject__iterator) OperatorMinusGreater2() *QJsonValueRef {
	return newQJsonValueRef(QJsonObject__iterator_OperatorMinusGreater2(this.h))
}

func (this *QJsonObject__iterator) OperatorSubscript(j int64) *QJsonValueRef {
	_goptr := newQJsonValueRef(QJsonObject__iterator_OperatorSubscript(this.h, (ptrdiff_t)(j)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonObject__iterator) OperatorPlusPlus() *iterator {
	xxxxxxxxx
}

func (this *QJsonObject__iterator) OperatorPlusPlusWithInt(param1 int) iterator {
	xxxxxxxxx
}

func (this *QJsonObject__iterator) OperatorMinusMinus() *iterator {
	xxxxxxxxx
}

func (this *QJsonObject__iterator) OperatorMinusMinusWithInt(param1 int) iterator {
	xxxxxxxxx
}

func (this *QJsonObject__iterator) OperatorPlus(j int64) iterator {
	xxxxxxxxx
}

func (this *QJsonObject__iterator) OperatorMinus(j int64) iterator {
	xxxxxxxxx
}

func (this *QJsonObject__iterator) OperatorPlusAssign(j int64) *iterator {
	xxxxxxxxx
}

func (this *QJsonObject__iterator) OperatorMinusAssign(j int64) *iterator {
	xxxxxxxxx
}

func (this *QJsonObject__iterator) OperatorMinusWithIterator(j iterator) int64 {
	return (int64)(QJsonObject__iterator_OperatorMinusWithIterator(this.h, j))
}

type QJsonObject__const_iterator struct {
	h          uintptr
	isSubclass bool
}

// NewQJsonObject__const_iterator constructs a new QJsonObject::const_iterator object.
func NewQJsonObject__const_iterator() *QJsonObject__const_iterator {
	ret := newQJsonObject__const_iterator(QJsonObject__const_iterator_new())
	ret.isSubclass = true
	return ret
}

// NewQJsonObject__const_iterator2 constructs a new QJsonObject::const_iterator object.
func NewQJsonObject__const_iterator2(obj *QJsonObject, index int64) *QJsonObject__const_iterator {
	ret := newQJsonObject__const_iterator(QJsonObject__const_iterator_new2(obj.cPointer(), (ptrdiff_t)(index)))
	ret.isSubclass = true
	return ret
}

// NewQJsonObject__const_iterator3 constructs a new QJsonObject::const_iterator object.
func NewQJsonObject__const_iterator3(other *iterator) *QJsonObject__const_iterator {
	ret := newQJsonObject__const_iterator(QJsonObject__const_iterator_new3(other))
	ret.isSubclass = true
	return ret
}

// NewQJsonObject__const_iterator4 constructs a new QJsonObject::const_iterator object.
func NewQJsonObject__const_iterator4(other *const_iterator) *QJsonObject__const_iterator {
	ret := newQJsonObject__const_iterator(QJsonObject__const_iterator_new4(other))
	ret.isSubclass = true
	return ret
}

func (this *QJsonObject__const_iterator) OperatorAssign(other *const_iterator) {
	QJsonObject__const_iterator_OperatorAssign(this.h, other)
}

func (this *QJsonObject__const_iterator) Key() string {
	var _ms struct_miqt_string = QJsonObject__const_iterator_Key(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QJsonObject__const_iterator) Value() *QJsonValueConstRef {
	_goptr := newQJsonValueConstRef(QJsonObject__const_iterator_Value(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonObject__const_iterator) OperatorMultiply() *QJsonValueConstRef {
	_goptr := newQJsonValueConstRef(QJsonObject__const_iterator_OperatorMultiply(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonObject__const_iterator) OperatorMinusGreater() *QJsonValueConstRef {
	return newQJsonValueConstRef(QJsonObject__const_iterator_OperatorMinusGreater(this.h))
}

func (this *QJsonObject__const_iterator) OperatorSubscript(j int64) *QJsonValueConstRef {
	_goptr := newQJsonValueConstRef(QJsonObject__const_iterator_OperatorSubscript(this.h, (ptrdiff_t)(j)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonObject__const_iterator) OperatorPlusPlus() *const_iterator {
	xxxxxxxxx
}

func (this *QJsonObject__const_iterator) OperatorPlusPlusWithInt(param1 int) const_iterator {
	xxxxxxxxx
}

func (this *QJsonObject__const_iterator) OperatorMinusMinus() *const_iterator {
	xxxxxxxxx
}

func (this *QJsonObject__const_iterator) OperatorMinusMinusWithInt(param1 int) const_iterator {
	xxxxxxxxx
}

func (this *QJsonObject__const_iterator) OperatorPlus(j int64) const_iterator {
	xxxxxxxxx
}

func (this *QJsonObject__const_iterator) OperatorMinus(j int64) const_iterator {
	xxxxxxxxx
}

func (this *QJsonObject__const_iterator) OperatorPlusAssign(j int64) *const_iterator {
	xxxxxxxxx
}

func (this *QJsonObject__const_iterator) OperatorMinusAssign(j int64) *const_iterator {
	xxxxxxxxx
}

func (this *QJsonObject__const_iterator) OperatorMinusWithConstIterator(j const_iterator) int64 {
	return (int64)(QJsonObject__const_iterator_OperatorMinusWithConstIterator(this.h, j))
}
