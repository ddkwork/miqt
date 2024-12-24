package qt

import (
	"unsafe"
)

type QJsonValue__Type int

const (
	QJsonValue__Null      QJsonValue__Type = 0
	QJsonValue__Bool      QJsonValue__Type = 1
	QJsonValue__Double    QJsonValue__Type = 2
	QJsonValue__String    QJsonValue__Type = 3
	QJsonValue__Array     QJsonValue__Type = 4
	QJsonValue__Object    QJsonValue__Type = 5
	QJsonValue__Undefined QJsonValue__Type = 128
)

type QJsonValue__JsonFormat int

const (
	QJsonValue__Indented QJsonValue__JsonFormat = 0
	QJsonValue__Compact  QJsonValue__JsonFormat = 1
)

type QJsonValue struct {
	h          uintptr
	isSubclass bool
}

// NewQJsonValue constructs a new QJsonValue object.
func NewQJsonValue() *QJsonValue {
	g := newQJsonValue(QJsonValue_new())
	g.isSubclass = true
	return g
}

// NewQJsonValue2 constructs a new QJsonValue object.
func NewQJsonValue2(b bool) *QJsonValue {
	g := newQJsonValue(QJsonValue_new2((bool)(b)))
	g.isSubclass = true
	return g
}

// NewQJsonValue3 constructs a new QJsonValue object.
func NewQJsonValue3(n float64) *QJsonValue {
	g := newQJsonValue(QJsonValue_new3((double)(n)))
	g.isSubclass = true
	return g
}

// NewQJsonValue4 constructs a new QJsonValue object.
func NewQJsonValue4(n int) *QJsonValue {
	g := newQJsonValue(QJsonValue_new4((int)(n)))
	g.isSubclass = true
	return g
}

// NewQJsonValue5 constructs a new QJsonValue object.
func NewQJsonValue5(v int64) *QJsonValue {
	g := newQJsonValue(QJsonValue_new5((longlong)(v)))
	g.isSubclass = true
	return g
}

// NewQJsonValue6 constructs a new QJsonValue object.
func NewQJsonValue6(s string) *QJsonValue {
	s_ms := struct_miqt_string{}
	s_ms.data = CString(s)
	s_ms.len = size_t(len(s))
	defer free(unsafe.Pointer(s_ms.data))
	g := newQJsonValue(QJsonValue_new6(s_ms))
	g.isSubclass = true
	return g
}

// NewQJsonValue7 constructs a new QJsonValue object.
func NewQJsonValue7(s string) *QJsonValue {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	g := newQJsonValue(QJsonValue_new7(s_Cstring))
	g.isSubclass = true
	return g
}

// NewQJsonValue8 constructs a new QJsonValue object.
func NewQJsonValue8(a *QJsonArray) *QJsonValue {
	g := newQJsonValue(QJsonValue_new8(a.cPointer()))
	g.isSubclass = true
	return g
}

// NewQJsonValue9 constructs a new QJsonValue object.
func NewQJsonValue9(o *QJsonObject) *QJsonValue {
	g := newQJsonValue(QJsonValue_new9(o.cPointer()))
	g.isSubclass = true
	return g
}

// NewQJsonValue10 constructs a new QJsonValue object.
func NewQJsonValue10(other *QJsonValue) *QJsonValue {
	g := newQJsonValue(QJsonValue_new10(other.cPointer()))
	g.isSubclass = true
	return g
}

// NewQJsonValue11 constructs a new QJsonValue object.
func NewQJsonValue11(param1 Type) *QJsonValue {
	g := newQJsonValue(QJsonValue_new11(param1))
	g.isSubclass = true
	return g
}

func (this *QJsonValue) OperatorAssign(other *QJsonValue) {
	QJsonValue_OperatorAssign(this.h, other.cPointer())
}

func (this *QJsonValue) Swap(other *QJsonValue) {
	QJsonValue_Swap(this.h, other.cPointer())
}

func QJsonValue_FromVariant(variant *QVariant) *QJsonValue {
	_goptr := newQJsonValue(QJsonValue_FromVariant(variant.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonValue) ToVariant() *QVariant {
	_goptr := newQVariant(QJsonValue_ToVariant(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QJsonValue_FromJson(json QByteArrayView) *QJsonValue {
	_goptr := newQJsonValue(QJsonValue_FromJson(json.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonValue) ToJson() []byte {
	var _bytearray struct_miqt_string = QJsonValue_ToJson(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QJsonValue) Type() Type {
	xxxxxxxxx
}

func (this *QJsonValue) IsNull() bool {
	return (bool)(QJsonValue_IsNull(this.h))
}

func (this *QJsonValue) IsBool() bool {
	return (bool)(QJsonValue_IsBool(this.h))
}

func (this *QJsonValue) IsDouble() bool {
	return (bool)(QJsonValue_IsDouble(this.h))
}

func (this *QJsonValue) IsString() bool {
	return (bool)(QJsonValue_IsString(this.h))
}

func (this *QJsonValue) IsArray() bool {
	return (bool)(QJsonValue_IsArray(this.h))
}

func (this *QJsonValue) IsObject() bool {
	return (bool)(QJsonValue_IsObject(this.h))
}

func (this *QJsonValue) IsUndefined() bool {
	return (bool)(QJsonValue_IsUndefined(this.h))
}

func (this *QJsonValue) ToBool() bool {
	return (bool)(QJsonValue_ToBool(this.h))
}

func (this *QJsonValue) ToInt() int {
	return (int)(QJsonValue_ToInt(this.h))
}

func (this *QJsonValue) ToInteger() int64 {
	return (int64)(QJsonValue_ToInteger(this.h))
}

func (this *QJsonValue) ToDouble() float64 {
	return (float64)(QJsonValue_ToDouble(this.h))
}

func (this *QJsonValue) ToString() string {
	var _ms struct_miqt_string = QJsonValue_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QJsonValue) ToStringWithDefaultValue(defaultValue string) string {
	defaultValue_ms := struct_miqt_string{}
	defaultValue_ms.data = CString(defaultValue)
	defaultValue_ms.len = size_t(len(defaultValue))
	defer free(unsafe.Pointer(defaultValue_ms.data))
	var _ms struct_miqt_string = QJsonValue_ToStringWithDefaultValue(this.h, defaultValue_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QJsonValue) ToArray() *QJsonArray {
	_goptr := newQJsonArray(QJsonValue_ToArray(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonValue) ToArrayWithDefaultValue(defaultValue *QJsonArray) *QJsonArray {
	_goptr := newQJsonArray(QJsonValue_ToArrayWithDefaultValue(this.h, defaultValue.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonValue) ToObject() *QJsonObject {
	_goptr := newQJsonObject(QJsonValue_ToObject(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonValue) ToObjectWithDefaultValue(defaultValue *QJsonObject) *QJsonObject {
	_goptr := newQJsonObject(QJsonValue_ToObjectWithDefaultValue(this.h, defaultValue.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonValue) OperatorSubscript(key string) *QJsonValue {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQJsonValue(QJsonValue_OperatorSubscript(this.h, key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonValue) OperatorSubscriptWithQsizetype(i int64) *QJsonValue {
	_goptr := newQJsonValue(QJsonValue_OperatorSubscriptWithQsizetype(this.h, (ptrdiff_t)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QJsonValue_FromJson2(json QByteArrayView, error *QJsonParseError) *QJsonValue {
	_goptr := newQJsonValue(QJsonValue_FromJson2(json.cPointer(), error.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonValue) ToJson1(format JsonFormat) []byte {
	var _bytearray struct_miqt_string = QJsonValue_ToJson1(this.h, format)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QJsonValue) ToBool1(defaultValue bool) bool {
	return (bool)(QJsonValue_ToBool1(this.h, (bool)(defaultValue)))
}

func (this *QJsonValue) ToInt1(defaultValue int) int {
	return (int)(QJsonValue_ToInt1(this.h, (int)(defaultValue)))
}

func (this *QJsonValue) ToInteger1(defaultValue int64) int64 {
	return (int64)(QJsonValue_ToInteger1(this.h, (longlong)(defaultValue)))
}

func (this *QJsonValue) ToDouble1(defaultValue float64) float64 {
	return (float64)(QJsonValue_ToDouble1(this.h, (double)(defaultValue)))
}

type QJsonValueConstRef struct {
	h          uintptr
	isSubclass bool
}

// NewQJsonValueConstRef constructs a new QJsonValueConstRef object.
func NewQJsonValueConstRef(param1 *QJsonValueConstRef) *QJsonValueConstRef {
	g := newQJsonValueConstRef(QJsonValueConstRef_new(param1.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QJsonValueConstRef) ToVariant() *QVariant {
	_goptr := newQVariant(QJsonValueConstRef_ToVariant(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonValueConstRef) Type() QJsonValue__Type {
	return (QJsonValue__Type)(QJsonValueConstRef_Type(this.h))
}

func (this *QJsonValueConstRef) IsNull() bool {
	return (bool)(QJsonValueConstRef_IsNull(this.h))
}

func (this *QJsonValueConstRef) IsBool() bool {
	return (bool)(QJsonValueConstRef_IsBool(this.h))
}

func (this *QJsonValueConstRef) IsDouble() bool {
	return (bool)(QJsonValueConstRef_IsDouble(this.h))
}

func (this *QJsonValueConstRef) IsString() bool {
	return (bool)(QJsonValueConstRef_IsString(this.h))
}

func (this *QJsonValueConstRef) IsArray() bool {
	return (bool)(QJsonValueConstRef_IsArray(this.h))
}

func (this *QJsonValueConstRef) IsObject() bool {
	return (bool)(QJsonValueConstRef_IsObject(this.h))
}

func (this *QJsonValueConstRef) IsUndefined() bool {
	return (bool)(QJsonValueConstRef_IsUndefined(this.h))
}

func (this *QJsonValueConstRef) ToBool() bool {
	return (bool)(QJsonValueConstRef_ToBool(this.h))
}

func (this *QJsonValueConstRef) ToInt() int {
	return (int)(QJsonValueConstRef_ToInt(this.h))
}

func (this *QJsonValueConstRef) ToInteger() int64 {
	return (int64)(QJsonValueConstRef_ToInteger(this.h))
}

func (this *QJsonValueConstRef) ToDouble() float64 {
	return (float64)(QJsonValueConstRef_ToDouble(this.h))
}

func (this *QJsonValueConstRef) ToString() string {
	var _ms struct_miqt_string = QJsonValueConstRef_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QJsonValueConstRef) ToArray() *QJsonArray {
	_goptr := newQJsonArray(QJsonValueConstRef_ToArray(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonValueConstRef) ToObject() *QJsonObject {
	_goptr := newQJsonObject(QJsonValueConstRef_ToObject(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonValueConstRef) OperatorSubscriptWithQsizetype(i int64) *QJsonValue {
	_goptr := newQJsonValue(QJsonValueConstRef_OperatorSubscriptWithQsizetype(this.h, (ptrdiff_t)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonValueConstRef) ToBool1(defaultValue bool) bool {
	return (bool)(QJsonValueConstRef_ToBool1(this.h, (bool)(defaultValue)))
}

func (this *QJsonValueConstRef) ToInt1(defaultValue int) int {
	return (int)(QJsonValueConstRef_ToInt1(this.h, (int)(defaultValue)))
}

func (this *QJsonValueConstRef) ToInteger1(defaultValue int64) int64 {
	return (int64)(QJsonValueConstRef_ToInteger1(this.h, (longlong)(defaultValue)))
}

func (this *QJsonValueConstRef) ToDouble1(defaultValue float64) float64 {
	return (float64)(QJsonValueConstRef_ToDouble1(this.h, (double)(defaultValue)))
}

func (this *QJsonValueConstRef) ToString1(defaultValue string) string {
	defaultValue_ms := struct_miqt_string{}
	defaultValue_ms.data = CString(defaultValue)
	defaultValue_ms.len = size_t(len(defaultValue))
	defer free(unsafe.Pointer(defaultValue_ms.data))
	var _ms struct_miqt_string = QJsonValueConstRef_ToString1(this.h, defaultValue_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QJsonValueRef struct {
	h          uintptr
	isSubclass bool
}

// NewQJsonValueRef constructs a new QJsonValueRef object.
func NewQJsonValueRef(param1 *QJsonValueRef) *QJsonValueRef {
	g := newQJsonValueRef(QJsonValueRef_new(param1.cPointer()))
	g.isSubclass = true
	return g
}

// NewQJsonValueRef2 constructs a new QJsonValueRef object.
func NewQJsonValueRef2(array *QJsonArray, idx int64) *QJsonValueRef {
	g := newQJsonValueRef(QJsonValueRef_new2(array.cPointer(), (ptrdiff_t)(idx)))
	g.isSubclass = true
	return g
}

// NewQJsonValueRef3 constructs a new QJsonValueRef object.
func NewQJsonValueRef3(object *QJsonObject, idx int64) *QJsonValueRef {
	g := newQJsonValueRef(QJsonValueRef_new3(object.cPointer(), (ptrdiff_t)(idx)))
	g.isSubclass = true
	return g
}

func (this *QJsonValueRef) OperatorAssign(val *QJsonValue) {
	QJsonValueRef_OperatorAssign(this.h, val.cPointer())
}

func (this *QJsonValueRef) OperatorAssignWithVal(val *QJsonValueRef) {
	QJsonValueRef_OperatorAssignWithVal(this.h, val.cPointer())
}

func (this *QJsonValueRef) ToVariant() *QVariant {
	_goptr := newQVariant(QJsonValueRef_ToVariant(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonValueRef) Type() QJsonValue__Type {
	return (QJsonValue__Type)(QJsonValueRef_Type(this.h))
}

func (this *QJsonValueRef) IsNull() bool {
	return (bool)(QJsonValueRef_IsNull(this.h))
}

func (this *QJsonValueRef) IsBool() bool {
	return (bool)(QJsonValueRef_IsBool(this.h))
}

func (this *QJsonValueRef) IsDouble() bool {
	return (bool)(QJsonValueRef_IsDouble(this.h))
}

func (this *QJsonValueRef) IsString() bool {
	return (bool)(QJsonValueRef_IsString(this.h))
}

func (this *QJsonValueRef) IsArray() bool {
	return (bool)(QJsonValueRef_IsArray(this.h))
}

func (this *QJsonValueRef) IsObject() bool {
	return (bool)(QJsonValueRef_IsObject(this.h))
}

func (this *QJsonValueRef) IsUndefined() bool {
	return (bool)(QJsonValueRef_IsUndefined(this.h))
}

func (this *QJsonValueRef) ToBool() bool {
	return (bool)(QJsonValueRef_ToBool(this.h))
}

func (this *QJsonValueRef) ToInt() int {
	return (int)(QJsonValueRef_ToInt(this.h))
}

func (this *QJsonValueRef) ToInteger() int64 {
	return (int64)(QJsonValueRef_ToInteger(this.h))
}

func (this *QJsonValueRef) ToDouble() float64 {
	return (float64)(QJsonValueRef_ToDouble(this.h))
}

func (this *QJsonValueRef) ToString() string {
	var _ms struct_miqt_string = QJsonValueRef_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QJsonValueRef) ToArray() *QJsonArray {
	_goptr := newQJsonArray(QJsonValueRef_ToArray(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonValueRef) ToObject() *QJsonObject {
	_goptr := newQJsonObject(QJsonValueRef_ToObject(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonValueRef) OperatorSubscriptWithQsizetype(i int64) *QJsonValue {
	_goptr := newQJsonValue(QJsonValueRef_OperatorSubscriptWithQsizetype(this.h, (ptrdiff_t)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QJsonValueRef) ToBool1(defaultValue bool) bool {
	return (bool)(QJsonValueRef_ToBool1(this.h, (bool)(defaultValue)))
}

func (this *QJsonValueRef) ToInt1(defaultValue int) int {
	return (int)(QJsonValueRef_ToInt1(this.h, (int)(defaultValue)))
}

func (this *QJsonValueRef) ToInteger1(defaultValue int64) int64 {
	return (int64)(QJsonValueRef_ToInteger1(this.h, (longlong)(defaultValue)))
}

func (this *QJsonValueRef) ToDouble1(defaultValue float64) float64 {
	return (float64)(QJsonValueRef_ToDouble1(this.h, (double)(defaultValue)))
}

func (this *QJsonValueRef) ToString1(defaultValue string) string {
	defaultValue_ms := struct_miqt_string{}
	defaultValue_ms.data = CString(defaultValue)
	defaultValue_ms.len = size_t(len(defaultValue))
	defer free(unsafe.Pointer(defaultValue_ms.data))
	var _ms struct_miqt_string = QJsonValueRef_ToString1(this.h, defaultValue_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
