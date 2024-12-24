package qt

import (
	"unsafe"
)

type QUrlQuery struct {
	h          uintptr
	isSubclass bool
}

// NewQUrlQuery constructs a new QUrlQuery object.
func NewQUrlQuery() *QUrlQuery {

	ret := newQUrlQuery(QUrlQuery_new())
	ret.isSubclass = true
	return ret
}

// NewQUrlQuery2 constructs a new QUrlQuery object.
func NewQUrlQuery2(url *QUrl) *QUrlQuery {

	ret := newQUrlQuery(QUrlQuery_new2(url.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQUrlQuery3 constructs a new QUrlQuery object.
func NewQUrlQuery3(queryString string) *QUrlQuery {
	queryString_ms := struct_miqt_string{}
	queryString_ms.data = CString(queryString)
	queryString_ms.len = size_t(len(queryString))
	defer free(unsafe.Pointer(queryString_ms.data))

	ret := newQUrlQuery(QUrlQuery_new3(queryString_ms))
	ret.isSubclass = true
	return ret
}

// NewQUrlQuery4 constructs a new QUrlQuery object.
func NewQUrlQuery4(other *QUrlQuery) *QUrlQuery {

	ret := newQUrlQuery(QUrlQuery_new4(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QUrlQuery) OperatorAssign(other *QUrlQuery) {
	QUrlQuery_OperatorAssign(this.h, other.cPointer())
}

func (this *QUrlQuery) Swap(other *QUrlQuery) {
	QUrlQuery_Swap(this.h, other.cPointer())
}

func (this *QUrlQuery) IsEmpty() bool {
	return (bool)(QUrlQuery_IsEmpty(this.h))
}

func (this *QUrlQuery) IsDetached() bool {
	return (bool)(QUrlQuery_IsDetached(this.h))
}

func (this *QUrlQuery) Clear() {
	QUrlQuery_Clear(this.h)
}

func (this *QUrlQuery) Query() string {
	var _ms struct_miqt_string = QUrlQuery_Query(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrlQuery) SetQuery(queryString string) {
	queryString_ms := struct_miqt_string{}
	queryString_ms.data = CString(queryString)
	queryString_ms.len = size_t(len(queryString))
	defer free(unsafe.Pointer(queryString_ms.data))
	QUrlQuery_SetQuery(this.h, queryString_ms)
}

func (this *QUrlQuery) ToString() string {
	var _ms struct_miqt_string = QUrlQuery_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrlQuery) SetQueryDelimiters(valueDelimiter QChar, pairDelimiter QChar) {
	QUrlQuery_SetQueryDelimiters(this.h, valueDelimiter.cPointer(), pairDelimiter.cPointer())
}

func (this *QUrlQuery) QueryValueDelimiter() *QChar {
	_goptr := newQChar(QUrlQuery_QueryValueDelimiter(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUrlQuery) QueryPairDelimiter() *QChar {
	_goptr := newQChar(QUrlQuery_QueryPairDelimiter(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUrlQuery) HasQueryItem(key string) bool {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	return (bool)(QUrlQuery_HasQueryItem(this.h, key_ms))
}

func (this *QUrlQuery) AddQueryItem(key string, value string) {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	value_ms := struct_miqt_string{}
	value_ms.data = CString(value)
	value_ms.len = size_t(len(value))
	defer free(unsafe.Pointer(value_ms.data))
	QUrlQuery_AddQueryItem(this.h, key_ms, value_ms)
}

func (this *QUrlQuery) RemoveQueryItem(key string) {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	QUrlQuery_RemoveQueryItem(this.h, key_ms)
}

func (this *QUrlQuery) QueryItemValue(key string) string {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	var _ms struct_miqt_string = QUrlQuery_QueryItemValue(this.h, key_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrlQuery) AllQueryItemValues(key string) []string {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	var _ma struct_miqt_array = QUrlQuery_AllQueryItemValues(this.h, key_ms)
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

func (this *QUrlQuery) RemoveAllQueryItems(key string) {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	QUrlQuery_RemoveAllQueryItems(this.h, key_ms)
}

func (this *QUrlQuery) DataPtr() *DataPtr {
	xxxxxxxxx
}

func (this *QUrlQuery) Query1(encoding ComponentFormattingOption) string {
	var _ms struct_miqt_string = QUrlQuery_Query1(this.h, (int)(encoding))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrlQuery) ToString1(encoding ComponentFormattingOption) string {
	var _ms struct_miqt_string = QUrlQuery_ToString1(this.h, (int)(encoding))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrlQuery) QueryItemValue2(key string, encoding ComponentFormattingOption) string {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	var _ms struct_miqt_string = QUrlQuery_QueryItemValue2(this.h, key_ms, (int)(encoding))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrlQuery) AllQueryItemValues2(key string, encoding ComponentFormattingOption) []string {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	var _ma struct_miqt_array = QUrlQuery_AllQueryItemValues2(this.h, key_ms, (int)(encoding))
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
