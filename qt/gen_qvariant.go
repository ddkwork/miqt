package qt

import (
	"unsafe"
)

type QVariant__Type int

const (
	QVariant__Invalid              QVariant__Type = 0
	QVariant__Bool                 QVariant__Type = 1
	QVariant__Int                  QVariant__Type = 2
	QVariant__UInt                 QVariant__Type = 3
	QVariant__LongLong             QVariant__Type = 4
	QVariant__ULongLong            QVariant__Type = 5
	QVariant__Double               QVariant__Type = 6
	QVariant__Char                 QVariant__Type = 7
	QVariant__Map                  QVariant__Type = 8
	QVariant__List                 QVariant__Type = 9
	QVariant__String               QVariant__Type = 10
	QVariant__StringList           QVariant__Type = 11
	QVariant__ByteArray            QVariant__Type = 12
	QVariant__BitArray             QVariant__Type = 13
	QVariant__Date                 QVariant__Type = 14
	QVariant__Time                 QVariant__Type = 15
	QVariant__DateTime             QVariant__Type = 16
	QVariant__Url                  QVariant__Type = 17
	QVariant__Locale               QVariant__Type = 18
	QVariant__Rect                 QVariant__Type = 19
	QVariant__RectF                QVariant__Type = 20
	QVariant__Size                 QVariant__Type = 21
	QVariant__SizeF                QVariant__Type = 22
	QVariant__Line                 QVariant__Type = 23
	QVariant__LineF                QVariant__Type = 24
	QVariant__Point                QVariant__Type = 25
	QVariant__PointF               QVariant__Type = 26
	QVariant__RegularExpression    QVariant__Type = 44
	QVariant__Hash                 QVariant__Type = 28
	QVariant__EasingCurve          QVariant__Type = 29
	QVariant__Uuid                 QVariant__Type = 30
	QVariant__ModelIndex           QVariant__Type = 42
	QVariant__PersistentModelIndex QVariant__Type = 50
	QVariant__LastCoreType         QVariant__Type = 63
	QVariant__Font                 QVariant__Type = 4096
	QVariant__Pixmap               QVariant__Type = 4097
	QVariant__Brush                QVariant__Type = 4098
	QVariant__Color                QVariant__Type = 4099
	QVariant__Palette              QVariant__Type = 4100
	QVariant__Image                QVariant__Type = 4102
	QVariant__Polygon              QVariant__Type = 4103
	QVariant__Region               QVariant__Type = 4104
	QVariant__Bitmap               QVariant__Type = 4105
	QVariant__Cursor               QVariant__Type = 4106
	QVariant__KeySequence          QVariant__Type = 4107
	QVariant__Pen                  QVariant__Type = 4108
	QVariant__TextLength           QVariant__Type = 4109
	QVariant__TextFormat           QVariant__Type = 4110
	QVariant__Transform            QVariant__Type = 4112
	QVariant__Matrix4x4            QVariant__Type = 4113
	QVariant__Vector2D             QVariant__Type = 4114
	QVariant__Vector3D             QVariant__Type = 4115
	QVariant__Vector4D             QVariant__Type = 4116
	QVariant__Quaternion           QVariant__Type = 4117
	QVariant__PolygonF             QVariant__Type = 4118
	QVariant__Icon                 QVariant__Type = 4101
	QVariant__LastGuiType          QVariant__Type = 4119
	QVariant__SizePolicy           QVariant__Type = 8192
	QVariant__UserType             QVariant__Type = 65536
	QVariant__LastType             QVariant__Type = 4294967295
)

type QVariant struct {
	h          uintptr
	isSubclass bool
}

// NewQVariant constructs a new QVariant object.
func NewQVariant() *QVariant {
	ret := newQVariant(QVariant_new())
	ret.isSubclass = true
	return ret
}

// NewQVariant2 constructs a new QVariant object.
func NewQVariant2(typeVal QMetaType) *QVariant {
	ret := newQVariant(QVariant_new2(typeVal.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVariant3 constructs a new QVariant object.
func NewQVariant3(other *QVariant) *QVariant {
	ret := newQVariant(QVariant_new3(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVariant4 constructs a new QVariant object.
func NewQVariant4(i int) *QVariant {
	ret := newQVariant(QVariant_new4((int)(i)))
	ret.isSubclass = true
	return ret
}

// NewQVariant5 constructs a new QVariant object.
func NewQVariant5(ui uint) *QVariant {
	ret := newQVariant(QVariant_new5((uint)(ui)))
	ret.isSubclass = true
	return ret
}

// NewQVariant6 constructs a new QVariant object.
func NewQVariant6(ll int64) *QVariant {
	ret := newQVariant(QVariant_new6((longlong)(ll)))
	ret.isSubclass = true
	return ret
}

// NewQVariant7 constructs a new QVariant object.
func NewQVariant7(ull uint64) *QVariant {
	ret := newQVariant(QVariant_new7((ulonglong)(ull)))
	ret.isSubclass = true
	return ret
}

// NewQVariant8 constructs a new QVariant object.
func NewQVariant8(b bool) *QVariant {
	ret := newQVariant(QVariant_new8((bool)(b)))
	ret.isSubclass = true
	return ret
}

// NewQVariant9 constructs a new QVariant object.
func NewQVariant9(d float64) *QVariant {
	ret := newQVariant(QVariant_new9((double)(d)))
	ret.isSubclass = true
	return ret
}

// NewQVariant10 constructs a new QVariant object.
func NewQVariant10(f float32) *QVariant {
	ret := newQVariant(QVariant_new10((float)(f)))
	ret.isSubclass = true
	return ret
}

// NewQVariant11 constructs a new QVariant object.
func NewQVariant11(qchar QChar) *QVariant {
	ret := newQVariant(QVariant_new11(qchar.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVariant12 constructs a new QVariant object.
func NewQVariant12(date QDate) *QVariant {
	ret := newQVariant(QVariant_new12(date.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVariant13 constructs a new QVariant object.
func NewQVariant13(time QTime) *QVariant {
	ret := newQVariant(QVariant_new13(time.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVariant14 constructs a new QVariant object.
func NewQVariant14(bitarray *QBitArray) *QVariant {
	ret := newQVariant(QVariant_new14(bitarray.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVariant15 constructs a new QVariant object.
func NewQVariant15(bytearray []byte) *QVariant {
	bytearray_alias := struct_miqt_string{}
	bytearray_alias.data = (char)(unsafe.Pointer(&bytearray[0]))
	bytearray_alias.len = size_t(len(bytearray))

	ret := newQVariant(QVariant_new15(bytearray_alias))
	ret.isSubclass = true
	return ret
}

// NewQVariant16 constructs a new QVariant object.
func NewQVariant16(datetime *QDateTime) *QVariant {
	ret := newQVariant(QVariant_new16(datetime.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVariant17 constructs a new QVariant object.
func NewQVariant17(hash map[string]QVariant) *QVariant {
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
		hash_Values_CArray[hash_ctr] = hash_v.cPointer()
		hash_ctr++
	}
	hash_mm := struct_miqt_map{
		len:    size_t(len(hash)),
		keys:   unsafe.Pointer(hash_Keys_CArray),
		values: unsafe.Pointer(hash_Values_CArray),
	}

	ret := newQVariant(QVariant_new17(hash_mm))
	ret.isSubclass = true
	return ret
}

// NewQVariant18 constructs a new QVariant object.
func NewQVariant18(jsonArray *QJsonArray) *QVariant {
	ret := newQVariant(QVariant_new18(jsonArray.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVariant19 constructs a new QVariant object.
func NewQVariant19(jsonObject *QJsonObject) *QVariant {
	ret := newQVariant(QVariant_new19(jsonObject.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVariant20 constructs a new QVariant object.
func NewQVariant20(locale *QLocale) *QVariant {
	ret := newQVariant(QVariant_new20(locale.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVariant21 constructs a new QVariant object.
func NewQVariant21(mapVal map[string]QVariant) *QVariant {
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

	ret := newQVariant(QVariant_new21(mapVal_mm))
	ret.isSubclass = true
	return ret
}

// NewQVariant22 constructs a new QVariant object.
func NewQVariant22(re *QRegularExpression) *QVariant {
	ret := newQVariant(QVariant_new22(re.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVariant23 constructs a new QVariant object.
func NewQVariant23(stringVal string) *QVariant {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))

	ret := newQVariant(QVariant_new23(stringVal_ms))
	ret.isSubclass = true
	return ret
}

// NewQVariant24 constructs a new QVariant object.
func NewQVariant24(stringlist []string) *QVariant {
	stringlist_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(stringlist))))
	defer free(unsafe.Pointer(stringlist_CArray))
	for i := range stringlist {
		stringlist_i_ms := struct_miqt_string{}
		stringlist_i_ms.data = CString(stringlist[i])
		stringlist_i_ms.len = size_t(len(stringlist[i]))
		defer free(unsafe.Pointer(stringlist_i_ms.data))
		stringlist_CArray[i] = stringlist_i_ms
	}
	stringlist_ma := struct_miqt_array{len: size_t(len(stringlist)), data: unsafe.Pointer(stringlist_CArray)}

	ret := newQVariant(QVariant_new24(stringlist_ma))
	ret.isSubclass = true
	return ret
}

// NewQVariant25 constructs a new QVariant object.
func NewQVariant25(url *QUrl) *QVariant {
	ret := newQVariant(QVariant_new25(url.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVariant26 constructs a new QVariant object.
func NewQVariant26(size QSize) *QVariant {
	ret := newQVariant(QVariant_new26(size.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVariant27 constructs a new QVariant object.
func NewQVariant27(pt QPoint) *QVariant {
	ret := newQVariant(QVariant_new27(pt.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVariant28 constructs a new QVariant object.
func NewQVariant28(typeVal Type) *QVariant {
	ret := newQVariant(QVariant_new28(typeVal))
	ret.isSubclass = true
	return ret
}

// NewQVariant29 constructs a new QVariant object.
func NewQVariant29(typeVal QMetaType, copyVal unsafe.Pointer) *QVariant {
	ret := newQVariant(QVariant_new29(typeVal.cPointer(), copyVal))
	ret.isSubclass = true
	return ret
}

func (this *QVariant) OperatorAssign(other *QVariant) {
	QVariant_OperatorAssign(this.h, other.cPointer())
}

func (this *QVariant) Swap(other *QVariant) {
	QVariant_Swap(this.h, other.cPointer())
}

func (this *QVariant) UserType() int {
	return (int)(QVariant_UserType(this.h))
}

func (this *QVariant) TypeId() int {
	return (int)(QVariant_TypeId(this.h))
}

func (this *QVariant) TypeName() string {
	_ret := QVariant_TypeName(this.h)
	return GoString(_ret)
}

func (this *QVariant) MetaType() *QMetaType {
	_goptr := newQMetaType(QVariant_MetaType(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) CanConvert(targetType QMetaType) bool {
	return (bool)(QVariant_CanConvert(this.h, targetType.cPointer()))
}

func (this *QVariant) Convert(typeVal QMetaType) bool {
	return (bool)(QVariant_Convert(this.h, typeVal.cPointer()))
}

func (this *QVariant) CanView(targetType QMetaType) bool {
	return (bool)(QVariant_CanView(this.h, targetType.cPointer()))
}

func (this *QVariant) CanConvertWithTargetTypeId(targetTypeId int) bool {
	return (bool)(QVariant_CanConvertWithTargetTypeId(this.h, (int)(targetTypeId)))
}

func (this *QVariant) ConvertWithTargetTypeId(targetTypeId int) bool {
	return (bool)(QVariant_ConvertWithTargetTypeId(this.h, (int)(targetTypeId)))
}

func (this *QVariant) IsValid() bool {
	return (bool)(QVariant_IsValid(this.h))
}

func (this *QVariant) IsNull() bool {
	return (bool)(QVariant_IsNull(this.h))
}

func (this *QVariant) Clear() {
	QVariant_Clear(this.h)
}

func (this *QVariant) Detach() {
	QVariant_Detach(this.h)
}

func (this *QVariant) IsDetached() bool {
	return (bool)(QVariant_IsDetached(this.h))
}

func (this *QVariant) ToInt() int {
	return (int)(QVariant_ToInt(this.h))
}

func (this *QVariant) ToUInt() uint {
	return (uint)(QVariant_ToUInt(this.h))
}

func (this *QVariant) ToLongLong() int64 {
	return (int64)(QVariant_ToLongLong(this.h))
}

func (this *QVariant) ToULongLong() uint64 {
	return (uint64)(QVariant_ToULongLong(this.h))
}

func (this *QVariant) ToBool() bool {
	return (bool)(QVariant_ToBool(this.h))
}

func (this *QVariant) ToDouble() float64 {
	return (float64)(QVariant_ToDouble(this.h))
}

func (this *QVariant) ToFloat() float32 {
	return (float32)(QVariant_ToFloat(this.h))
}

func (this *QVariant) ToReal() float64 {
	return (float64)(QVariant_ToReal(this.h))
}

func (this *QVariant) ToByteArray() []byte {
	var _bytearray struct_miqt_string = QVariant_ToByteArray(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QVariant) ToBitArray() *QBitArray {
	_goptr := newQBitArray(QVariant_ToBitArray(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToString() string {
	var _ms struct_miqt_string = QVariant_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVariant) ToStringList() []string {
	var _ma struct_miqt_array = QVariant_ToStringList(this.h)
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

func (this *QVariant) ToChar() *QChar {
	_goptr := newQChar(QVariant_ToChar(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToDate() *QDate {
	_goptr := newQDate(QVariant_ToDate(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToTime() *QTime {
	_goptr := newQTime(QVariant_ToTime(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToDateTime() *QDateTime {
	_goptr := newQDateTime(QVariant_ToDateTime(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToMap() map[string]QVariant {
	var _mm struct_miqt_map = QVariant_ToMap(this.h)
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

func (this *QVariant) ToHash() map[string]QVariant {
	var _mm struct_miqt_map = QVariant_ToHash(this.h)
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

func (this *QVariant) ToPoint() *QPoint {
	_goptr := newQPoint(QVariant_ToPoint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToPointF() *QPointF {
	_goptr := newQPointF(QVariant_ToPointF(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToRect() *QRect {
	_goptr := newQRect(QVariant_ToRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToSize() *QSize {
	_goptr := newQSize(QVariant_ToSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToSizeF() *QSizeF {
	_goptr := newQSizeF(QVariant_ToSizeF(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToLine() *QLine {
	_goptr := newQLine(QVariant_ToLine(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToLineF() *QLineF {
	_goptr := newQLineF(QVariant_ToLineF(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToRectF() *QRectF {
	_goptr := newQRectF(QVariant_ToRectF(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToLocale() *QLocale {
	_goptr := newQLocale(QVariant_ToLocale(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToRegularExpression() *QRegularExpression {
	_goptr := newQRegularExpression(QVariant_ToRegularExpression(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToEasingCurve() *QEasingCurve {
	_goptr := newQEasingCurve(QVariant_ToEasingCurve(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToUuid() *QUuid {
	_goptr := newQUuid(QVariant_ToUuid(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToUrl() *QUrl {
	_goptr := newQUrl(QVariant_ToUrl(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToJsonValue() *QJsonValue {
	_goptr := newQJsonValue(QVariant_ToJsonValue(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToJsonObject() *QJsonObject {
	_goptr := newQJsonObject(QVariant_ToJsonObject(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToJsonArray() *QJsonArray {
	_goptr := newQJsonArray(QVariant_ToJsonArray(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToJsonDocument() *QJsonDocument {
	_goptr := newQJsonDocument(QVariant_ToJsonDocument(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToModelIndex() *QModelIndex {
	_goptr := newQModelIndex(QVariant_ToModelIndex(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) ToPersistentModelIndex() *QPersistentModelIndex {
	_goptr := newQPersistentModelIndex(QVariant_ToPersistentModelIndex(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) Load(ds *QDataStream) {
	QVariant_Load(this.h, ds.cPointer())
}

func (this *QVariant) Save(ds *QDataStream) {
	QVariant_Save(this.h, ds.cPointer())
}

func (this *QVariant) Type() Type {
	xxxxxxxxx
}

func QVariant_TypeToName(typeId int) string {
	_ret := QVariant_TypeToName((int)(typeId))
	return GoString(_ret)
}

func QVariant_NameToType(name string) Type {
	name_Cstring := CString(name)
	defer free(unsafe.Pointer(name_Cstring))
	xxxxxxxxx
}

func (this *QVariant) Data() unsafe.Pointer {
	return (unsafe.Pointer)(QVariant_Data(this.h))
}

func (this *QVariant) ConstData() unsafe.Pointer {
	return (unsafe.Pointer)(QVariant_ConstData(this.h))
}

func (this *QVariant) Data2() unsafe.Pointer {
	return (unsafe.Pointer)(QVariant_Data2(this.h))
}

func (this *QVariant) SetValue(avalue *QVariant) {
	QVariant_SetValue(this.h, avalue.cPointer())
}

func QVariant_FromMetaType(typeVal QMetaType) *QVariant {
	_goptr := newQVariant(QVariant_FromMetaType(typeVal.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QVariant_Compare(lhs *QVariant, rhs *QVariant) *QPartialOrdering {
	_goptr := newQPartialOrdering(QVariant_Compare(lhs.cPointer(), rhs.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariant) DataPtr() *DataPtr {
	xxxxxxxxx
}

func (this *QVariant) DataPtr2() *DataPtr {
	xxxxxxxxx
}

func (this *QVariant) ToInt1(ok *bool) int {
	return (int)(QVariant_ToInt1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QVariant) ToUInt1(ok *bool) uint {
	return (uint)(QVariant_ToUInt1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QVariant) ToLongLong1(ok *bool) int64 {
	return (int64)(QVariant_ToLongLong1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QVariant) ToULongLong1(ok *bool) uint64 {
	return (uint64)(QVariant_ToULongLong1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QVariant) ToDouble1(ok *bool) float64 {
	return (float64)(QVariant_ToDouble1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QVariant) ToFloat1(ok *bool) float32 {
	return (float32)(QVariant_ToFloat1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QVariant) ToReal1(ok *bool) float64 {
	return (float64)(QVariant_ToReal1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func QVariant_FromMetaType2(typeVal QMetaType, copyVal unsafe.Pointer) *QVariant {
	_goptr := newQVariant(QVariant_FromMetaType2(typeVal.cPointer(), copyVal))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QVariantConstPointer struct {
	h          uintptr
	isSubclass bool
}

// NewQVariantConstPointer constructs a new QVariantConstPointer object.
func NewQVariantConstPointer(variant QVariant) *QVariantConstPointer {
	ret := newQVariantConstPointer(QVariantConstPointer_new(variant.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQVariantConstPointer2 constructs a new QVariantConstPointer object.
func NewQVariantConstPointer2(param1 *QVariantConstPointer) *QVariantConstPointer {
	ret := newQVariantConstPointer(QVariantConstPointer_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QVariantConstPointer) OperatorMultiply() *QVariant {
	_goptr := newQVariant(QVariantConstPointer_OperatorMultiply(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariantConstPointer) OperatorMinusGreater() *QVariant {
	return newQVariant(QVariantConstPointer_OperatorMinusGreater(this.h))
}

func (this *QVariantConstPointer) OperatorAssign(param1 *QVariantConstPointer) {
	QVariantConstPointer_OperatorAssign(this.h, param1.cPointer())
}
