package qt

import (
	"unsafe"
)

type QJsonDocument__JsonFormat int

const (
	QJsonDocument__Indented QJsonDocument__JsonFormat = 0
	QJsonDocument__Compact  QJsonDocument__JsonFormat = 1
)

type QJsonDocument struct {
	h          uintptr
	isSubclass bool
}

// NewQJsonDocument constructs a new QJsonDocument object.
func NewQJsonDocument() *QJsonDocument {
	ret := newQJsonDocument(QJsonDocument_new())
	ret.isSubclass = true
	return ret
}

// NewQJsonDocument2 constructs a new QJsonDocument object.
func NewQJsonDocument2(object *QJsonObject) *QJsonDocument {
	ret := newQJsonDocument(QJsonDocument_new2(object.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQJsonDocument3 constructs a new QJsonDocument object.
func NewQJsonDocument3(array *QJsonArray) *QJsonDocument {
	ret := newQJsonDocument(QJsonDocument_new3(array.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQJsonDocument4 constructs a new QJsonDocument object.
func NewQJsonDocument4(other *QJsonDocument) *QJsonDocument {
	ret := newQJsonDocument(QJsonDocument_new4(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QJsonDocument) OperatorAssign(other *QJsonDocument) {
	QJsonDocument_OperatorAssign(this.h, other.cPointer())
}

func (this *QJsonDocument) Swap(other *QJsonDocument) {
	QJsonDocument_Swap(this.h, other.cPointer())
}

func QJsonDocument_FromVariant(variant *QVariant) *QJsonDocument {
	_goptr := newQJsonDocument(QJsonDocument_FromVariant(variant.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonDocument) ToVariant() *QVariant {
	_goptr := newQVariant(QJsonDocument_ToVariant(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QJsonDocument_FromJson(json []byte) *QJsonDocument {
	json_alias := struct_miqt_string{}
	json_alias.data = (char)(unsafe.Pointer(&json[0]))
	json_alias.len = size_t(len(json))
	_goptr := newQJsonDocument(QJsonDocument_FromJson(json_alias))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonDocument) ToJson() []byte {
	var _bytearray struct_miqt_string = QJsonDocument_ToJson(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QJsonDocument) IsEmpty() bool {
	return (bool)(QJsonDocument_IsEmpty(this.h))
}

func (this *QJsonDocument) IsArray() bool {
	return (bool)(QJsonDocument_IsArray(this.h))
}

func (this *QJsonDocument) IsObject() bool {
	return (bool)(QJsonDocument_IsObject(this.h))
}

func (this *QJsonDocument) Object() *QJsonObject {
	_goptr := newQJsonObject(QJsonDocument_Object(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonDocument) Array() *QJsonArray {
	_goptr := newQJsonArray(QJsonDocument_Array(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonDocument) SetObject(object *QJsonObject) {
	QJsonDocument_SetObject(this.h, object.cPointer())
}

func (this *QJsonDocument) SetArray(array *QJsonArray) {
	QJsonDocument_SetArray(this.h, array.cPointer())
}

func (this *QJsonDocument) OperatorSubscript(key string) *QJsonValue {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQJsonValue(QJsonDocument_OperatorSubscript(this.h, key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonDocument) OperatorSubscriptWithQsizetype(i int64) *QJsonValue {
	_goptr := newQJsonValue(QJsonDocument_OperatorSubscriptWithQsizetype(this.h, (ptrdiff_t)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonDocument) IsNull() bool {
	return (bool)(QJsonDocument_IsNull(this.h))
}

func QJsonDocument_FromJson2(json []byte, error *QJsonParseError) *QJsonDocument {
	json_alias := struct_miqt_string{}
	json_alias.data = (char)(unsafe.Pointer(&json[0]))
	json_alias.len = size_t(len(json))
	_goptr := newQJsonDocument(QJsonDocument_FromJson2(json_alias, error.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonDocument) ToJson1(format JsonFormat) []byte {
	var _bytearray struct_miqt_string = QJsonDocument_ToJson1(this.h, format)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}
