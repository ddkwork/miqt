package cbor

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QCborValue__EncodingOption int

const (
	QCborValue__SortKeysInMaps   QCborValue__EncodingOption = 1
	QCborValue__UseFloat         QCborValue__EncodingOption = 2
	QCborValue__UseFloat16       QCborValue__EncodingOption = 6
	QCborValue__UseIntegers      QCborValue__EncodingOption = 8
	QCborValue__NoTransformation QCborValue__EncodingOption = 0
)

type QCborValue__DiagnosticNotationOption int

const (
	QCborValue__Compact        QCborValue__DiagnosticNotationOption = 0
	QCborValue__LineWrapped    QCborValue__DiagnosticNotationOption = 1
	QCborValue__ExtendedFormat QCborValue__DiagnosticNotationOption = 2
)

type QCborValue__Type int

const (
	QCborValue__Integer           QCborValue__Type = 0
	QCborValue__ByteArray         QCborValue__Type = 64
	QCborValue__String            QCborValue__Type = 96
	QCborValue__Array             QCborValue__Type = 128
	QCborValue__Map               QCborValue__Type = 160
	QCborValue__Tag               QCborValue__Type = 192
	QCborValue__SimpleType        QCborValue__Type = 256
	QCborValue__False             QCborValue__Type = 276
	QCborValue__True              QCborValue__Type = 277
	QCborValue__Null              QCborValue__Type = 278
	QCborValue__Undefined         QCborValue__Type = 279
	QCborValue__Double            QCborValue__Type = 514
	QCborValue__DateTime          QCborValue__Type = 65536
	QCborValue__Url               QCborValue__Type = 65568
	QCborValue__RegularExpression QCborValue__Type = 65571
	QCborValue__Uuid              QCborValue__Type = 65573
	QCborValue__Invalid           QCborValue__Type = -1
)

type QCborParserError struct {
	h          uintptr
	isSubclass bool
}

func (this *QCborParserError) ErrorString() string {
	var _ms struct_miqt_string = QCborParserError_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QCborValue struct {
	h          uintptr
	isSubclass bool
}

// NewQCborValue constructs a new QCborValue object.
func NewQCborValue() *QCborValue {

	ret := newQCborValue(QCborValue_new())
	ret.isSubclass = true
	return ret
}

// NewQCborValue2 constructs a new QCborValue object.
func NewQCborValue2(t_ Type) *QCborValue {

	ret := newQCborValue(QCborValue_new2(t_))
	ret.isSubclass = true
	return ret
}

// NewQCborValue3 constructs a new QCborValue object.
func NewQCborValue3(b_ bool) *QCborValue {

	ret := newQCborValue(QCborValue_new3((bool)(b_)))
	ret.isSubclass = true
	return ret
}

// NewQCborValue4 constructs a new QCborValue object.
func NewQCborValue4(i int) *QCborValue {

	ret := newQCborValue(QCborValue_new4((int)(i)))
	ret.isSubclass = true
	return ret
}

// NewQCborValue5 constructs a new QCborValue object.
func NewQCborValue5(u uint) *QCborValue {

	ret := newQCborValue(QCborValue_new5((uint)(u)))
	ret.isSubclass = true
	return ret
}

// NewQCborValue6 constructs a new QCborValue object.
func NewQCborValue6(i int64) *QCborValue {

	ret := newQCborValue(QCborValue_new6((longlong)(i)))
	ret.isSubclass = true
	return ret
}

// NewQCborValue7 constructs a new QCborValue object.
func NewQCborValue7(v float64) *QCborValue {

	ret := newQCborValue(QCborValue_new7((double)(v)))
	ret.isSubclass = true
	return ret
}

// NewQCborValue8 constructs a new QCborValue object.
func NewQCborValue8(st QCborSimpleType) *QCborValue {

	ret := newQCborValue(QCborValue_new8((uint8_t)(st)))
	ret.isSubclass = true
	return ret
}

// NewQCborValue9 constructs a new QCborValue object.
func NewQCborValue9(ba []byte) *QCborValue {
	ba_alias := struct_miqt_string{}
	ba_alias.data = (char)(unsafe.Pointer(&ba[0]))
	ba_alias.len = size_t(len(ba))

	ret := newQCborValue(QCborValue_new9(ba_alias))
	ret.isSubclass = true
	return ret
}

// NewQCborValue10 constructs a new QCborValue object.
func NewQCborValue10(s string) *QCborValue {
	s_ms := struct_miqt_string{}
	s_ms.data = CString(s)
	s_ms.len = size_t(len(s))
	defer free(unsafe.Pointer(s_ms.data))

	ret := newQCborValue(QCborValue_new10(s_ms))
	ret.isSubclass = true
	return ret
}

// NewQCborValue11 constructs a new QCborValue object.
func NewQCborValue11(s string) *QCborValue {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))

	ret := newQCborValue(QCborValue_new11(s_Cstring))
	ret.isSubclass = true
	return ret
}

// NewQCborValue12 constructs a new QCborValue object.
func NewQCborValue12(a *QCborArray) *QCborValue {

	ret := newQCborValue(QCborValue_new12(a.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQCborValue13 constructs a new QCborValue object.
func NewQCborValue13(m *QCborMap) *QCborValue {

	ret := newQCborValue(QCborValue_new13(m.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQCborValue14 constructs a new QCborValue object.
func NewQCborValue14(tag QCborTag) *QCborValue {

	ret := newQCborValue(QCborValue_new14((uint64_t)(tag)))
	ret.isSubclass = true
	return ret
}

// NewQCborValue15 constructs a new QCborValue object.
func NewQCborValue15(t_ QCborKnownTags) *QCborValue {

	ret := newQCborValue(QCborValue_new15((int)(t_)))
	ret.isSubclass = true
	return ret
}

// NewQCborValue16 constructs a new QCborValue object.
func NewQCborValue16(dt *qt.QDateTime) *QCborValue {

	ret := newQCborValue(QCborValue_new16((*QDateTime)(dt.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQCborValue17 constructs a new QCborValue object.
func NewQCborValue17(url *qt.QUrl) *QCborValue {

	ret := newQCborValue(QCborValue_new17((*QUrl)(url.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQCborValue18 constructs a new QCborValue object.
func NewQCborValue18(rx *qt.QRegularExpression) *QCborValue {

	ret := newQCborValue(QCborValue_new18((*QRegularExpression)(rx.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQCborValue19 constructs a new QCborValue object.
func NewQCborValue19(uuid *qt.QUuid) *QCborValue {

	ret := newQCborValue(QCborValue_new19((*QUuid)(uuid.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQCborValue20 constructs a new QCborValue object.
func NewQCborValue20(other *QCborValue) *QCborValue {

	ret := newQCborValue(QCborValue_new20(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQCborValue21 constructs a new QCborValue object.
func NewQCborValue21(tag QCborTag, taggedValue *QCborValue) *QCborValue {

	ret := newQCborValue(QCborValue_new21((uint64_t)(tag), taggedValue.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQCborValue22 constructs a new QCborValue object.
func NewQCborValue22(t_ QCborKnownTags, tv *QCborValue) *QCborValue {

	ret := newQCborValue(QCborValue_new22((int)(t_), tv.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QCborValue) OperatorAssign(other *QCborValue) {
	QCborValue_OperatorAssign(this.h, other.cPointer())
}

func (this *QCborValue) Swap(other *QCborValue) {
	QCborValue_Swap(this.h, other.cPointer())
}

func (this *QCborValue) Type() Type {
	xxxxxxxxx
}

func (this *QCborValue) IsInteger() bool {
	return (bool)(QCborValue_IsInteger(this.h))
}

func (this *QCborValue) IsByteArray() bool {
	return (bool)(QCborValue_IsByteArray(this.h))
}

func (this *QCborValue) IsString() bool {
	return (bool)(QCborValue_IsString(this.h))
}

func (this *QCborValue) IsArray() bool {
	return (bool)(QCborValue_IsArray(this.h))
}

func (this *QCborValue) IsMap() bool {
	return (bool)(QCborValue_IsMap(this.h))
}

func (this *QCborValue) IsTag() bool {
	return (bool)(QCborValue_IsTag(this.h))
}

func (this *QCborValue) IsFalse() bool {
	return (bool)(QCborValue_IsFalse(this.h))
}

func (this *QCborValue) IsTrue() bool {
	return (bool)(QCborValue_IsTrue(this.h))
}

func (this *QCborValue) IsBool() bool {
	return (bool)(QCborValue_IsBool(this.h))
}

func (this *QCborValue) IsNull() bool {
	return (bool)(QCborValue_IsNull(this.h))
}

func (this *QCborValue) IsUndefined() bool {
	return (bool)(QCborValue_IsUndefined(this.h))
}

func (this *QCborValue) IsDouble() bool {
	return (bool)(QCborValue_IsDouble(this.h))
}

func (this *QCborValue) IsDateTime() bool {
	return (bool)(QCborValue_IsDateTime(this.h))
}

func (this *QCborValue) IsUrl() bool {
	return (bool)(QCborValue_IsUrl(this.h))
}

func (this *QCborValue) IsRegularExpression() bool {
	return (bool)(QCborValue_IsRegularExpression(this.h))
}

func (this *QCborValue) IsUuid() bool {
	return (bool)(QCborValue_IsUuid(this.h))
}

func (this *QCborValue) IsInvalid() bool {
	return (bool)(QCborValue_IsInvalid(this.h))
}

func (this *QCborValue) IsContainer() bool {
	return (bool)(QCborValue_IsContainer(this.h))
}

func (this *QCborValue) IsSimpleType() bool {
	return (bool)(QCborValue_IsSimpleType(this.h))
}

func (this *QCborValue) IsSimpleTypeWithSt(st QCborSimpleType) bool {
	return (bool)(QCborValue_IsSimpleTypeWithSt(this.h, (uint8_t)(st)))
}

func (this *QCborValue) ToSimpleType() QCborSimpleType {
	return (QCborSimpleType)(QCborValue_ToSimpleType(this.h))
}

func (this *QCborValue) ToInteger() int64 {
	return (int64)(QCborValue_ToInteger(this.h))
}

func (this *QCborValue) ToBool() bool {
	return (bool)(QCborValue_ToBool(this.h))
}

func (this *QCborValue) ToDouble() float64 {
	return (float64)(QCborValue_ToDouble(this.h))
}

func (this *QCborValue) Tag() QCborTag {
	return (QCborTag)(QCborValue_Tag(this.h))
}

func (this *QCborValue) TaggedValue() *QCborValue {
	_goptr := newQCborValue(QCborValue_TaggedValue(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) ToByteArray() []byte {
	var _bytearray struct_miqt_string = QCborValue_ToByteArray(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QCborValue) ToString() string {
	var _ms struct_miqt_string = QCborValue_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCborValue) ToDateTime() *qt.QDateTime {
	_goptr := qt.UnsafeNewQDateTime(unsafe.Pointer(QCborValue_ToDateTime(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) ToUrl() *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QCborValue_ToUrl(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) ToRegularExpression() *qt.QRegularExpression {
	_goptr := qt.UnsafeNewQRegularExpression(unsafe.Pointer(QCborValue_ToRegularExpression(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) ToUuid() *qt.QUuid {
	_goptr := qt.UnsafeNewQUuid(unsafe.Pointer(QCborValue_ToUuid(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) ToArray() *QCborArray {
	_goptr := newQCborArray(QCborValue_ToArray(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) ToArrayWithDefaultValue(defaultValue *QCborArray) *QCborArray {
	_goptr := newQCborArray(QCborValue_ToArrayWithDefaultValue(this.h, defaultValue.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) ToMap() *QCborMap {
	_goptr := newQCborMap(QCborValue_ToMap(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) ToMapWithDefaultValue(defaultValue *QCborMap) *QCborMap {
	_goptr := newQCborMap(QCborValue_ToMapWithDefaultValue(this.h, defaultValue.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) OperatorSubscript(key string) *QCborValue {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQCborValue(QCborValue_OperatorSubscript(this.h, key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) OperatorSubscript2(key int64) *QCborValue {
	_goptr := newQCborValue(QCborValue_OperatorSubscript2(this.h, (longlong)(key)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) OperatorSubscript3(key int64) *QCborValueRef {
	_goptr := newQCborValueRef(QCborValue_OperatorSubscript3(this.h, (longlong)(key)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) OperatorSubscript5(key string) *QCborValueRef {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQCborValueRef(QCborValue_OperatorSubscript5(this.h, key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) Compare(other *QCborValue) int {
	return (int)(QCborValue_Compare(this.h, other.cPointer()))
}

func QCborValue_FromVariant(variant *qt.QVariant) *QCborValue {
	_goptr := newQCborValue(QCborValue_FromVariant((*QVariant)(variant.UnsafePointer())))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) ToVariant() *qt.QVariant {
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QCborValue_ToVariant(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QCborValue_FromJsonValue(v *qt.QJsonValue) *QCborValue {
	_goptr := newQCborValue(QCborValue_FromJsonValue((*QJsonValue)(v.UnsafePointer())))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) ToJsonValue() *qt.QJsonValue {
	_goptr := qt.UnsafeNewQJsonValue(unsafe.Pointer(QCborValue_ToJsonValue(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QCborValue_FromCbor(reader *QCborStreamReader) *QCborValue {
	_goptr := newQCborValue(QCborValue_FromCbor(reader.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QCborValue_FromCborWithBa(ba []byte) *QCborValue {
	ba_alias := struct_miqt_string{}
	ba_alias.data = (char)(unsafe.Pointer(&ba[0]))
	ba_alias.len = size_t(len(ba))
	_goptr := newQCborValue(QCborValue_FromCborWithBa(ba_alias))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QCborValue_FromCbor2(data string, lenVal int64) *QCborValue {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))
	_goptr := newQCborValue(QCborValue_FromCbor2(data_Cstring, (ptrdiff_t)(lenVal)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QCborValue_FromCbor3(data *byte, lenVal int64) *QCborValue {
	_goptr := newQCborValue(QCborValue_FromCbor3((*uchar)(unsafe.Pointer(data)), (ptrdiff_t)(lenVal)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) ToCbor() []byte {
	var _bytearray struct_miqt_string = QCborValue_ToCbor(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QCborValue) ToCborWithWriter(writer *QCborStreamWriter) {
	QCborValue_ToCborWithWriter(this.h, writer.cPointer())
}

func (this *QCborValue) ToDiagnosticNotation() string {
	var _ms struct_miqt_string = QCborValue_ToDiagnosticNotation(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCborValue) ToSimpleType1(defaultValue QCborSimpleType) QCborSimpleType {
	return (QCborSimpleType)(QCborValue_ToSimpleType1(this.h, (uint8_t)(defaultValue)))
}

func (this *QCborValue) ToInteger1(defaultValue int64) int64 {
	return (int64)(QCborValue_ToInteger1(this.h, (longlong)(defaultValue)))
}

func (this *QCborValue) ToBool1(defaultValue bool) bool {
	return (bool)(QCborValue_ToBool1(this.h, (bool)(defaultValue)))
}

func (this *QCborValue) ToDouble1(defaultValue float64) float64 {
	return (float64)(QCborValue_ToDouble1(this.h, (double)(defaultValue)))
}

func (this *QCborValue) Tag1(defaultValue QCborTag) QCborTag {
	return (QCborTag)(QCborValue_Tag1(this.h, (uint64_t)(defaultValue)))
}

func (this *QCborValue) TaggedValue1(defaultValue *QCborValue) *QCborValue {
	_goptr := newQCborValue(QCborValue_TaggedValue1(this.h, defaultValue.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) ToByteArray1(defaultValue []byte) []byte {
	defaultValue_alias := struct_miqt_string{}
	defaultValue_alias.data = (char)(unsafe.Pointer(&defaultValue[0]))
	defaultValue_alias.len = size_t(len(defaultValue))
	var _bytearray struct_miqt_string = QCborValue_ToByteArray1(this.h, defaultValue_alias)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QCborValue) ToString1(defaultValue string) string {
	defaultValue_ms := struct_miqt_string{}
	defaultValue_ms.data = CString(defaultValue)
	defaultValue_ms.len = size_t(len(defaultValue))
	defer free(unsafe.Pointer(defaultValue_ms.data))
	var _ms struct_miqt_string = QCborValue_ToString1(this.h, defaultValue_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCborValue) ToDateTime1(defaultValue *qt.QDateTime) *qt.QDateTime {
	_goptr := qt.UnsafeNewQDateTime(unsafe.Pointer(QCborValue_ToDateTime1(this.h, (*QDateTime)(defaultValue.UnsafePointer()))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) ToUrl1(defaultValue *qt.QUrl) *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QCborValue_ToUrl1(this.h, (*QUrl)(defaultValue.UnsafePointer()))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) ToRegularExpression1(defaultValue *qt.QRegularExpression) *qt.QRegularExpression {
	_goptr := qt.UnsafeNewQRegularExpression(unsafe.Pointer(QCborValue_ToRegularExpression1(this.h, (*QRegularExpression)(defaultValue.UnsafePointer()))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) ToUuid1(defaultValue *qt.QUuid) *qt.QUuid {
	_goptr := qt.UnsafeNewQUuid(unsafe.Pointer(QCborValue_ToUuid1(this.h, (*QUuid)(defaultValue.UnsafePointer()))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QCborValue_FromCbor22(ba []byte, error *QCborParserError) *QCborValue {
	ba_alias := struct_miqt_string{}
	ba_alias.data = (char)(unsafe.Pointer(&ba[0]))
	ba_alias.len = size_t(len(ba))
	_goptr := newQCborValue(QCborValue_FromCbor22(ba_alias, error.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QCborValue_FromCbor32(data string, lenVal int64, error *QCborParserError) *QCborValue {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))
	_goptr := newQCborValue(QCborValue_FromCbor32(data_Cstring, (ptrdiff_t)(lenVal), error.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QCborValue_FromCbor33(data *byte, lenVal int64, error *QCborParserError) *QCborValue {
	_goptr := newQCborValue(QCborValue_FromCbor33((*uchar)(unsafe.Pointer(data)), (ptrdiff_t)(lenVal), error.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValue) ToCbor1(opt EncodingOptions) []byte {
	var _bytearray struct_miqt_string = QCborValue_ToCbor1(this.h, opt)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QCborValue) ToCbor2(writer *QCborStreamWriter, opt EncodingOptions) {
	QCborValue_ToCbor2(this.h, writer.cPointer(), opt)
}

func (this *QCborValue) ToDiagnosticNotation1(opts DiagnosticNotationOptions) string {
	var _ms struct_miqt_string = QCborValue_ToDiagnosticNotation1(this.h, opts)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QCborValueConstRef struct {
	h          uintptr
	isSubclass bool
}

// NewQCborValueConstRef constructs a new QCborValueConstRef object.
func NewQCborValueConstRef(param1 *QCborValueConstRef) *QCborValueConstRef {

	ret := newQCborValueConstRef(QCborValueConstRef_new(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QCborValueConstRef) Type() QCborValue__Type {
	return (QCborValue__Type)(QCborValueConstRef_Type(this.h))
}

func (this *QCborValueConstRef) IsInteger() bool {
	return (bool)(QCborValueConstRef_IsInteger(this.h))
}

func (this *QCborValueConstRef) IsByteArray() bool {
	return (bool)(QCborValueConstRef_IsByteArray(this.h))
}

func (this *QCborValueConstRef) IsString() bool {
	return (bool)(QCborValueConstRef_IsString(this.h))
}

func (this *QCborValueConstRef) IsArray() bool {
	return (bool)(QCborValueConstRef_IsArray(this.h))
}

func (this *QCborValueConstRef) IsMap() bool {
	return (bool)(QCborValueConstRef_IsMap(this.h))
}

func (this *QCborValueConstRef) IsTag() bool {
	return (bool)(QCborValueConstRef_IsTag(this.h))
}

func (this *QCborValueConstRef) IsFalse() bool {
	return (bool)(QCborValueConstRef_IsFalse(this.h))
}

func (this *QCborValueConstRef) IsTrue() bool {
	return (bool)(QCborValueConstRef_IsTrue(this.h))
}

func (this *QCborValueConstRef) IsBool() bool {
	return (bool)(QCborValueConstRef_IsBool(this.h))
}

func (this *QCborValueConstRef) IsNull() bool {
	return (bool)(QCborValueConstRef_IsNull(this.h))
}

func (this *QCborValueConstRef) IsUndefined() bool {
	return (bool)(QCborValueConstRef_IsUndefined(this.h))
}

func (this *QCborValueConstRef) IsDouble() bool {
	return (bool)(QCborValueConstRef_IsDouble(this.h))
}

func (this *QCborValueConstRef) IsDateTime() bool {
	return (bool)(QCborValueConstRef_IsDateTime(this.h))
}

func (this *QCborValueConstRef) IsUrl() bool {
	return (bool)(QCborValueConstRef_IsUrl(this.h))
}

func (this *QCborValueConstRef) IsRegularExpression() bool {
	return (bool)(QCborValueConstRef_IsRegularExpression(this.h))
}

func (this *QCborValueConstRef) IsUuid() bool {
	return (bool)(QCborValueConstRef_IsUuid(this.h))
}

func (this *QCborValueConstRef) IsInvalid() bool {
	return (bool)(QCborValueConstRef_IsInvalid(this.h))
}

func (this *QCborValueConstRef) IsContainer() bool {
	return (bool)(QCborValueConstRef_IsContainer(this.h))
}

func (this *QCborValueConstRef) IsSimpleType() bool {
	return (bool)(QCborValueConstRef_IsSimpleType(this.h))
}

func (this *QCborValueConstRef) IsSimpleTypeWithSt(st QCborSimpleType) bool {
	return (bool)(QCborValueConstRef_IsSimpleTypeWithSt(this.h, (uint8_t)(st)))
}

func (this *QCborValueConstRef) ToSimpleType() QCborSimpleType {
	return (QCborSimpleType)(QCborValueConstRef_ToSimpleType(this.h))
}

func (this *QCborValueConstRef) Tag() QCborTag {
	return (QCborTag)(QCborValueConstRef_Tag(this.h))
}

func (this *QCborValueConstRef) TaggedValue() *QCborValue {
	_goptr := newQCborValue(QCborValueConstRef_TaggedValue(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) ToInteger() int64 {
	return (int64)(QCborValueConstRef_ToInteger(this.h))
}

func (this *QCborValueConstRef) ToBool() bool {
	return (bool)(QCborValueConstRef_ToBool(this.h))
}

func (this *QCborValueConstRef) ToDouble() float64 {
	return (float64)(QCborValueConstRef_ToDouble(this.h))
}

func (this *QCborValueConstRef) ToByteArray() []byte {
	var _bytearray struct_miqt_string = QCborValueConstRef_ToByteArray(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QCborValueConstRef) ToString() string {
	var _ms struct_miqt_string = QCborValueConstRef_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCborValueConstRef) ToDateTime() *qt.QDateTime {
	_goptr := qt.UnsafeNewQDateTime(unsafe.Pointer(QCborValueConstRef_ToDateTime(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) ToUrl() *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QCborValueConstRef_ToUrl(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) ToRegularExpression() *qt.QRegularExpression {
	_goptr := qt.UnsafeNewQRegularExpression(unsafe.Pointer(QCborValueConstRef_ToRegularExpression(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) ToUuid() *qt.QUuid {
	_goptr := qt.UnsafeNewQUuid(unsafe.Pointer(QCborValueConstRef_ToUuid(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) ToArray() *QCborArray {
	_goptr := newQCborArray(QCborValueConstRef_ToArray(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) ToArrayWithQCborArray(a *QCborArray) *QCborArray {
	_goptr := newQCborArray(QCborValueConstRef_ToArrayWithQCborArray(this.h, a.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) ToMap() *QCborMap {
	_goptr := newQCborMap(QCborValueConstRef_ToMap(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) ToMapWithQCborMap(m *QCborMap) *QCborMap {
	_goptr := newQCborMap(QCborValueConstRef_ToMapWithQCborMap(this.h, m.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) OperatorSubscript(key string) *QCborValue {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQCborValue(QCborValueConstRef_OperatorSubscript(this.h, key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) OperatorSubscript2(key int64) *QCborValue {
	_goptr := newQCborValue(QCborValueConstRef_OperatorSubscript2(this.h, (longlong)(key)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) Compare(other *QCborValue) int {
	return (int)(QCborValueConstRef_Compare(this.h, other.cPointer()))
}

func (this *QCborValueConstRef) ToVariant() *qt.QVariant {
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QCborValueConstRef_ToVariant(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) ToJsonValue() *qt.QJsonValue {
	_goptr := qt.UnsafeNewQJsonValue(unsafe.Pointer(QCborValueConstRef_ToJsonValue(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) ToCbor() []byte {
	var _bytearray struct_miqt_string = QCborValueConstRef_ToCbor(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QCborValueConstRef) ToCborWithWriter(writer *QCborStreamWriter) {
	QCborValueConstRef_ToCborWithWriter(this.h, writer.cPointer())
}

func (this *QCborValueConstRef) ToDiagnosticNotation() string {
	var _ms struct_miqt_string = QCborValueConstRef_ToDiagnosticNotation(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCborValueConstRef) ToSimpleType1(defaultValue QCborSimpleType) QCborSimpleType {
	return (QCborSimpleType)(QCborValueConstRef_ToSimpleType1(this.h, (uint8_t)(defaultValue)))
}

func (this *QCborValueConstRef) Tag1(defaultValue QCborTag) QCborTag {
	return (QCborTag)(QCborValueConstRef_Tag1(this.h, (uint64_t)(defaultValue)))
}

func (this *QCborValueConstRef) TaggedValue1(defaultValue *QCborValue) *QCborValue {
	_goptr := newQCborValue(QCborValueConstRef_TaggedValue1(this.h, defaultValue.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) ToInteger1(defaultValue int64) int64 {
	return (int64)(QCborValueConstRef_ToInteger1(this.h, (longlong)(defaultValue)))
}

func (this *QCborValueConstRef) ToBool1(defaultValue bool) bool {
	return (bool)(QCborValueConstRef_ToBool1(this.h, (bool)(defaultValue)))
}

func (this *QCborValueConstRef) ToDouble1(defaultValue float64) float64 {
	return (float64)(QCborValueConstRef_ToDouble1(this.h, (double)(defaultValue)))
}

func (this *QCborValueConstRef) ToByteArray1(defaultValue []byte) []byte {
	defaultValue_alias := struct_miqt_string{}
	defaultValue_alias.data = (char)(unsafe.Pointer(&defaultValue[0]))
	defaultValue_alias.len = size_t(len(defaultValue))
	var _bytearray struct_miqt_string = QCborValueConstRef_ToByteArray1(this.h, defaultValue_alias)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QCborValueConstRef) ToString1(defaultValue string) string {
	defaultValue_ms := struct_miqt_string{}
	defaultValue_ms.data = CString(defaultValue)
	defaultValue_ms.len = size_t(len(defaultValue))
	defer free(unsafe.Pointer(defaultValue_ms.data))
	var _ms struct_miqt_string = QCborValueConstRef_ToString1(this.h, defaultValue_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCborValueConstRef) ToDateTime1(defaultValue *qt.QDateTime) *qt.QDateTime {
	_goptr := qt.UnsafeNewQDateTime(unsafe.Pointer(QCborValueConstRef_ToDateTime1(this.h, (*QDateTime)(defaultValue.UnsafePointer()))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) ToUrl1(defaultValue *qt.QUrl) *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QCborValueConstRef_ToUrl1(this.h, (*QUrl)(defaultValue.UnsafePointer()))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) ToRegularExpression1(defaultValue *qt.QRegularExpression) *qt.QRegularExpression {
	_goptr := qt.UnsafeNewQRegularExpression(unsafe.Pointer(QCborValueConstRef_ToRegularExpression1(this.h, (*QRegularExpression)(defaultValue.UnsafePointer()))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) ToUuid1(defaultValue *qt.QUuid) *qt.QUuid {
	_goptr := qt.UnsafeNewQUuid(unsafe.Pointer(QCborValueConstRef_ToUuid1(this.h, (*QUuid)(defaultValue.UnsafePointer()))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueConstRef) ToCbor1(opt EncodingOption) []byte {
	var _bytearray struct_miqt_string = QCborValueConstRef_ToCbor1(this.h, (int)(opt))
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QCborValueConstRef) ToCbor2(writer *QCborStreamWriter, opt EncodingOption) {
	QCborValueConstRef_ToCbor2(this.h, writer.cPointer(), (int)(opt))
}

func (this *QCborValueConstRef) ToDiagnosticNotation1(opt DiagnosticNotationOption) string {
	var _ms struct_miqt_string = QCborValueConstRef_ToDiagnosticNotation1(this.h, (int)(opt))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QCborValueRef struct {
	h          uintptr
	isSubclass bool
}

// NewQCborValueRef constructs a new QCborValueRef object.
func NewQCborValueRef(param1 *QCborValueRef) *QCborValueRef {

	ret := newQCborValueRef(QCborValueRef_new(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QCborValueRef) OperatorAssign(other *QCborValue) {
	QCborValueRef_OperatorAssign(this.h, other.cPointer())
}

func (this *QCborValueRef) OperatorAssignWithOther(other *QCborValueRef) {
	QCborValueRef_OperatorAssignWithOther(this.h, other.cPointer())
}

func (this *QCborValueRef) OperatorSubscript(key int64) *QCborValueRef {
	_goptr := newQCborValueRef(QCborValueRef_OperatorSubscript(this.h, (longlong)(key)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) OperatorSubscript2(key string) *QCborValueRef {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQCborValueRef(QCborValueRef_OperatorSubscript2(this.h, key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) Type() QCborValue__Type {
	return (QCborValue__Type)(QCborValueRef_Type(this.h))
}

func (this *QCborValueRef) IsInteger() bool {
	return (bool)(QCborValueRef_IsInteger(this.h))
}

func (this *QCborValueRef) IsByteArray() bool {
	return (bool)(QCborValueRef_IsByteArray(this.h))
}

func (this *QCborValueRef) IsString() bool {
	return (bool)(QCborValueRef_IsString(this.h))
}

func (this *QCborValueRef) IsArray() bool {
	return (bool)(QCborValueRef_IsArray(this.h))
}

func (this *QCborValueRef) IsMap() bool {
	return (bool)(QCborValueRef_IsMap(this.h))
}

func (this *QCborValueRef) IsTag() bool {
	return (bool)(QCborValueRef_IsTag(this.h))
}

func (this *QCborValueRef) IsFalse() bool {
	return (bool)(QCborValueRef_IsFalse(this.h))
}

func (this *QCborValueRef) IsTrue() bool {
	return (bool)(QCborValueRef_IsTrue(this.h))
}

func (this *QCborValueRef) IsBool() bool {
	return (bool)(QCborValueRef_IsBool(this.h))
}

func (this *QCborValueRef) IsNull() bool {
	return (bool)(QCborValueRef_IsNull(this.h))
}

func (this *QCborValueRef) IsUndefined() bool {
	return (bool)(QCborValueRef_IsUndefined(this.h))
}

func (this *QCborValueRef) IsDouble() bool {
	return (bool)(QCborValueRef_IsDouble(this.h))
}

func (this *QCborValueRef) IsDateTime() bool {
	return (bool)(QCborValueRef_IsDateTime(this.h))
}

func (this *QCborValueRef) IsUrl() bool {
	return (bool)(QCborValueRef_IsUrl(this.h))
}

func (this *QCborValueRef) IsRegularExpression() bool {
	return (bool)(QCborValueRef_IsRegularExpression(this.h))
}

func (this *QCborValueRef) IsUuid() bool {
	return (bool)(QCborValueRef_IsUuid(this.h))
}

func (this *QCborValueRef) IsInvalid() bool {
	return (bool)(QCborValueRef_IsInvalid(this.h))
}

func (this *QCborValueRef) IsContainer() bool {
	return (bool)(QCborValueRef_IsContainer(this.h))
}

func (this *QCborValueRef) IsSimpleType() bool {
	return (bool)(QCborValueRef_IsSimpleType(this.h))
}

func (this *QCborValueRef) IsSimpleTypeWithSt(st QCborSimpleType) bool {
	return (bool)(QCborValueRef_IsSimpleTypeWithSt(this.h, (uint8_t)(st)))
}

func (this *QCborValueRef) ToSimpleType() QCborSimpleType {
	return (QCborSimpleType)(QCborValueRef_ToSimpleType(this.h))
}

func (this *QCborValueRef) Tag() QCborTag {
	return (QCborTag)(QCborValueRef_Tag(this.h))
}

func (this *QCborValueRef) TaggedValue() *QCborValue {
	_goptr := newQCborValue(QCborValueRef_TaggedValue(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) ToInteger() int64 {
	return (int64)(QCborValueRef_ToInteger(this.h))
}

func (this *QCborValueRef) ToBool() bool {
	return (bool)(QCborValueRef_ToBool(this.h))
}

func (this *QCborValueRef) ToDouble() float64 {
	return (float64)(QCborValueRef_ToDouble(this.h))
}

func (this *QCborValueRef) ToByteArray() []byte {
	var _bytearray struct_miqt_string = QCborValueRef_ToByteArray(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QCborValueRef) ToString() string {
	var _ms struct_miqt_string = QCborValueRef_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCborValueRef) ToDateTime() *qt.QDateTime {
	_goptr := qt.UnsafeNewQDateTime(unsafe.Pointer(QCborValueRef_ToDateTime(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) ToUrl() *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QCborValueRef_ToUrl(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) ToRegularExpression() *qt.QRegularExpression {
	_goptr := qt.UnsafeNewQRegularExpression(unsafe.Pointer(QCborValueRef_ToRegularExpression(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) ToUuid() *qt.QUuid {
	_goptr := qt.UnsafeNewQUuid(unsafe.Pointer(QCborValueRef_ToUuid(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) ToArray() *QCborArray {
	_goptr := newQCborArray(QCborValueRef_ToArray(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) ToArrayWithQCborArray(a *QCborArray) *QCborArray {
	_goptr := newQCborArray(QCborValueRef_ToArrayWithQCborArray(this.h, a.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) ToMap() *QCborMap {
	_goptr := newQCborMap(QCborValueRef_ToMap(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) ToMapWithQCborMap(m *QCborMap) *QCborMap {
	_goptr := newQCborMap(QCborValueRef_ToMapWithQCborMap(this.h, m.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) OperatorSubscript3(key string) *QCborValue {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	_goptr := newQCborValue(QCborValueRef_OperatorSubscript3(this.h, key_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) OperatorSubscript5(key int64) *QCborValue {
	_goptr := newQCborValue(QCborValueRef_OperatorSubscript5(this.h, (longlong)(key)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) Compare(other *QCborValue) int {
	return (int)(QCborValueRef_Compare(this.h, other.cPointer()))
}

func (this *QCborValueRef) ToVariant() *qt.QVariant {
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QCborValueRef_ToVariant(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) ToJsonValue() *qt.QJsonValue {
	_goptr := qt.UnsafeNewQJsonValue(unsafe.Pointer(QCborValueRef_ToJsonValue(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) ToCbor() []byte {
	var _bytearray struct_miqt_string = QCborValueRef_ToCbor(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QCborValueRef) ToCborWithWriter(writer *QCborStreamWriter) {
	QCborValueRef_ToCborWithWriter(this.h, writer.cPointer())
}

func (this *QCborValueRef) ToDiagnosticNotation() string {
	var _ms struct_miqt_string = QCborValueRef_ToDiagnosticNotation(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCborValueRef) ToSimpleType1(defaultValue QCborSimpleType) QCborSimpleType {
	return (QCborSimpleType)(QCborValueRef_ToSimpleType1(this.h, (uint8_t)(defaultValue)))
}

func (this *QCborValueRef) Tag1(defaultValue QCborTag) QCborTag {
	return (QCborTag)(QCborValueRef_Tag1(this.h, (uint64_t)(defaultValue)))
}

func (this *QCborValueRef) TaggedValue1(defaultValue *QCborValue) *QCborValue {
	_goptr := newQCborValue(QCborValueRef_TaggedValue1(this.h, defaultValue.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) ToInteger1(defaultValue int64) int64 {
	return (int64)(QCborValueRef_ToInteger1(this.h, (longlong)(defaultValue)))
}

func (this *QCborValueRef) ToBool1(defaultValue bool) bool {
	return (bool)(QCborValueRef_ToBool1(this.h, (bool)(defaultValue)))
}

func (this *QCborValueRef) ToDouble1(defaultValue float64) float64 {
	return (float64)(QCborValueRef_ToDouble1(this.h, (double)(defaultValue)))
}

func (this *QCborValueRef) ToByteArray1(defaultValue []byte) []byte {
	defaultValue_alias := struct_miqt_string{}
	defaultValue_alias.data = (char)(unsafe.Pointer(&defaultValue[0]))
	defaultValue_alias.len = size_t(len(defaultValue))
	var _bytearray struct_miqt_string = QCborValueRef_ToByteArray1(this.h, defaultValue_alias)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QCborValueRef) ToString1(defaultValue string) string {
	defaultValue_ms := struct_miqt_string{}
	defaultValue_ms.data = CString(defaultValue)
	defaultValue_ms.len = size_t(len(defaultValue))
	defer free(unsafe.Pointer(defaultValue_ms.data))
	var _ms struct_miqt_string = QCborValueRef_ToString1(this.h, defaultValue_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCborValueRef) ToDateTime1(defaultValue *qt.QDateTime) *qt.QDateTime {
	_goptr := qt.UnsafeNewQDateTime(unsafe.Pointer(QCborValueRef_ToDateTime1(this.h, (*QDateTime)(defaultValue.UnsafePointer()))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) ToUrl1(defaultValue *qt.QUrl) *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(QCborValueRef_ToUrl1(this.h, (*QUrl)(defaultValue.UnsafePointer()))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) ToRegularExpression1(defaultValue *qt.QRegularExpression) *qt.QRegularExpression {
	_goptr := qt.UnsafeNewQRegularExpression(unsafe.Pointer(QCborValueRef_ToRegularExpression1(this.h, (*QRegularExpression)(defaultValue.UnsafePointer()))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) ToUuid1(defaultValue *qt.QUuid) *qt.QUuid {
	_goptr := qt.UnsafeNewQUuid(unsafe.Pointer(QCborValueRef_ToUuid1(this.h, (*QUuid)(defaultValue.UnsafePointer()))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCborValueRef) ToCbor1(opt EncodingOption) []byte {
	var _bytearray struct_miqt_string = QCborValueRef_ToCbor1(this.h, (int)(opt))
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QCborValueRef) ToCbor2(writer *QCborStreamWriter, opt EncodingOption) {
	QCborValueRef_ToCbor2(this.h, writer.cPointer(), (int)(opt))
}

func (this *QCborValueRef) ToDiagnosticNotation1(opt DiagnosticNotationOption) string {
	var _ms struct_miqt_string = QCborValueRef_ToDiagnosticNotation1(this.h, (int)(opt))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
