package qt

import (
	"unsafe"
)

type QStringView struct {
	h          uintptr
	isSubclass bool
}

// NewQStringView constructs a new QStringView object.
func NewQStringView() *QStringView {
	g := newQStringView(QStringView_new())
	g.isSubclass = true
	return g
}

func (this *QStringView) ToString() string {
	var _ms struct_miqt_string = QStringView_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStringView) Size() int64 {
	return (int64)(QStringView_Size(this.h))
}

func (this *QStringView) Data() const_pointer {
	xxxxxxxxx
}

func (this *QStringView) ConstData() const_pointer {
	xxxxxxxxx
}

func (this *QStringView) Utf16() *storage_type {
	xxxxxxxxx
}

func (this *QStringView) OperatorSubscript(n int64) *QChar {
	_goptr := newQChar(QStringView_OperatorSubscript(this.h, (ptrdiff_t)(n)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStringView) ToLatin1() []byte {
	var _bytearray struct_miqt_string = QStringView_ToLatin1(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QStringView) ToUtf8() []byte {
	var _bytearray struct_miqt_string = QStringView_ToUtf8(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QStringView) ToLocal8Bit() []byte {
	var _bytearray struct_miqt_string = QStringView_ToLocal8Bit(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QStringView) ToUcs4() []uint {
	var _ma struct_miqt_array = QStringView_ToUcs4(this.h)
	_ret := make([]uint, int(_ma.len))
	_outCast := (*[0xffff]uint)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (uint)(_outCast[i])
	}
	return _ret
}

func (this *QStringView) At(n int64) *QChar {
	_goptr := newQChar(QStringView_At(this.h, (ptrdiff_t)(n)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStringView) Truncate(n int64) {
	QStringView_Truncate(this.h, (ptrdiff_t)(n))
}

func (this *QStringView) Chop(n int64) {
	QStringView_Chop(this.h, (ptrdiff_t)(n))
}

func (this *QStringView) CompareWithQChar(c QChar) int {
	return (int)(QStringView_CompareWithQChar(this.h, c.cPointer()))
}

func (this *QStringView) Compare3(c QChar, cs CaseSensitivity) int {
	return (int)(QStringView_Compare3(this.h, c.cPointer(), (int)(cs)))
}

func (this *QStringView) StartsWithWithQChar(c QChar) bool {
	return (bool)(QStringView_StartsWithWithQChar(this.h, c.cPointer()))
}

func (this *QStringView) StartsWith2(c QChar, cs CaseSensitivity) bool {
	return (bool)(QStringView_StartsWith2(this.h, c.cPointer(), (int)(cs)))
}

func (this *QStringView) EndsWithWithQChar(c QChar) bool {
	return (bool)(QStringView_EndsWithWithQChar(this.h, c.cPointer()))
}

func (this *QStringView) EndsWith2(c QChar, cs CaseSensitivity) bool {
	return (bool)(QStringView_EndsWith2(this.h, c.cPointer(), (int)(cs)))
}

func (this *QStringView) IndexOf(c QChar) int64 {
	return (int64)(QStringView_IndexOf(this.h, c.cPointer()))
}

func (this *QStringView) Contains(c QChar) bool {
	return (bool)(QStringView_Contains(this.h, c.cPointer()))
}

func (this *QStringView) Count(c QChar) int64 {
	return (int64)(QStringView_Count(this.h, c.cPointer()))
}

func (this *QStringView) LastIndexOf(c QChar) int64 {
	return (int64)(QStringView_LastIndexOf(this.h, c.cPointer()))
}

func (this *QStringView) LastIndexOf2(c QChar, from int64) int64 {
	return (int64)(QStringView_LastIndexOf2(this.h, c.cPointer(), (ptrdiff_t)(from)))
}

func (this *QStringView) IndexOfWithRe(re *QRegularExpression) int64 {
	return (int64)(QStringView_IndexOfWithRe(this.h, re.cPointer()))
}

func (this *QStringView) LastIndexOf5(re *QRegularExpression, from int64) int64 {
	return (int64)(QStringView_LastIndexOf5(this.h, re.cPointer(), (ptrdiff_t)(from)))
}

func (this *QStringView) ContainsWithRe(re *QRegularExpression) bool {
	return (bool)(QStringView_ContainsWithRe(this.h, re.cPointer()))
}

func (this *QStringView) CountWithRe(re *QRegularExpression) int64 {
	return (int64)(QStringView_CountWithRe(this.h, re.cPointer()))
}

func (this *QStringView) IsRightToLeft() bool {
	return (bool)(QStringView_IsRightToLeft(this.h))
}

func (this *QStringView) IsValidUtf16() bool {
	return (bool)(QStringView_IsValidUtf16(this.h))
}

func (this *QStringView) IsUpper() bool {
	return (bool)(QStringView_IsUpper(this.h))
}

func (this *QStringView) IsLower() bool {
	return (bool)(QStringView_IsLower(this.h))
}

func (this *QStringView) ToShort() int16 {
	return (int16)(QStringView_ToShort(this.h))
}

func (this *QStringView) ToUShort() uint16 {
	return (uint16)(QStringView_ToUShort(this.h))
}

func (this *QStringView) ToInt() int {
	return (int)(QStringView_ToInt(this.h))
}

func (this *QStringView) ToUInt() uint {
	return (uint)(QStringView_ToUInt(this.h))
}

func (this *QStringView) ToLong() int32 {
	return (int32)(QStringView_ToLong(this.h))
}

func (this *QStringView) ToULong() uint32 {
	return (uint32)(QStringView_ToULong(this.h))
}

func (this *QStringView) ToLongLong() int64 {
	return (int64)(QStringView_ToLongLong(this.h))
}

func (this *QStringView) ToULongLong() uint64 {
	return (uint64)(QStringView_ToULongLong(this.h))
}

func (this *QStringView) ToFloat() float32 {
	return (float32)(QStringView_ToFloat(this.h))
}

func (this *QStringView) ToDouble() float64 {
	return (float64)(QStringView_ToDouble(this.h))
}

func (this *QStringView) Begin() const_iterator {
	xxxxxxxxx
}

func (this *QStringView) End() const_iterator {
	xxxxxxxxx
}

func (this *QStringView) Cbegin() const_iterator {
	xxxxxxxxx
}

func (this *QStringView) Cend() const_iterator {
	xxxxxxxxx
}

func (this *QStringView) Rbegin() const_reverse_iterator {
	xxxxxxxxx
}

func (this *QStringView) Rend() const_reverse_iterator {
	xxxxxxxxx
}

func (this *QStringView) Crbegin() const_reverse_iterator {
	xxxxxxxxx
}

func (this *QStringView) Crend() const_reverse_iterator {
	xxxxxxxxx
}

func (this *QStringView) Empty() bool {
	return (bool)(QStringView_Empty(this.h))
}

func (this *QStringView) Front() *QChar {
	_goptr := newQChar(QStringView_Front(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStringView) Back() *QChar {
	_goptr := newQChar(QStringView_Back(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStringView) MaxSize() int64 {
	return (int64)(QStringView_MaxSize(this.h))
}

func (this *QStringView) ConstBegin() const_iterator {
	xxxxxxxxx
}

func (this *QStringView) ConstEnd() const_iterator {
	xxxxxxxxx
}

func (this *QStringView) IsNull() bool {
	return (bool)(QStringView_IsNull(this.h))
}

func (this *QStringView) IsEmpty() bool {
	return (bool)(QStringView_IsEmpty(this.h))
}

func (this *QStringView) Length() int64 {
	return (int64)(QStringView_Length(this.h))
}

func (this *QStringView) First2() *QChar {
	_goptr := newQChar(QStringView_First2(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStringView) Last2() *QChar {
	_goptr := newQChar(QStringView_Last2(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QStringView_MaxSize2() int64 {
	return (int64)(QStringView_MaxSize2())
}

func (this *QStringView) IndexOf2(c QChar, from int64) int64 {
	return (int64)(QStringView_IndexOf2(this.h, c.cPointer(), (ptrdiff_t)(from)))
}

func (this *QStringView) IndexOf3(c QChar, from int64, cs CaseSensitivity) int64 {
	return (int64)(QStringView_IndexOf3(this.h, c.cPointer(), (ptrdiff_t)(from), (int)(cs)))
}

func (this *QStringView) Contains2(c QChar, cs CaseSensitivity) bool {
	return (bool)(QStringView_Contains2(this.h, c.cPointer(), (int)(cs)))
}

func (this *QStringView) Count2(c QChar, cs CaseSensitivity) int64 {
	return (int64)(QStringView_Count2(this.h, c.cPointer(), (int)(cs)))
}

func (this *QStringView) LastIndexOf22(c QChar, cs CaseSensitivity) int64 {
	return (int64)(QStringView_LastIndexOf22(this.h, c.cPointer(), (int)(cs)))
}

func (this *QStringView) LastIndexOf32(c QChar, from int64, cs CaseSensitivity) int64 {
	return (int64)(QStringView_LastIndexOf32(this.h, c.cPointer(), (ptrdiff_t)(from), (int)(cs)))
}

func (this *QStringView) IndexOf24(re *QRegularExpression, from int64) int64 {
	return (int64)(QStringView_IndexOf24(this.h, re.cPointer(), (ptrdiff_t)(from)))
}

func (this *QStringView) IndexOf34(re *QRegularExpression, from int64, rmatch *QRegularExpressionMatch) int64 {
	return (int64)(QStringView_IndexOf34(this.h, re.cPointer(), (ptrdiff_t)(from), rmatch.cPointer()))
}

func (this *QStringView) LastIndexOf35(re *QRegularExpression, from int64, rmatch *QRegularExpressionMatch) int64 {
	return (int64)(QStringView_LastIndexOf35(this.h, re.cPointer(), (ptrdiff_t)(from), rmatch.cPointer()))
}

func (this *QStringView) Contains24(re *QRegularExpression, rmatch *QRegularExpressionMatch) bool {
	return (bool)(QStringView_Contains24(this.h, re.cPointer(), rmatch.cPointer()))
}

func (this *QStringView) ToShort1(ok *bool) int16 {
	return (int16)(QStringView_ToShort1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QStringView) ToShort2(ok *bool, base int) int16 {
	return (int16)(QStringView_ToShort2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QStringView) ToUShort1(ok *bool) uint16 {
	return (uint16)(QStringView_ToUShort1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QStringView) ToUShort2(ok *bool, base int) uint16 {
	return (uint16)(QStringView_ToUShort2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QStringView) ToInt1(ok *bool) int {
	return (int)(QStringView_ToInt1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QStringView) ToInt2(ok *bool, base int) int {
	return (int)(QStringView_ToInt2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QStringView) ToUInt1(ok *bool) uint {
	return (uint)(QStringView_ToUInt1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QStringView) ToUInt2(ok *bool, base int) uint {
	return (uint)(QStringView_ToUInt2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QStringView) ToLong1(ok *bool) int32 {
	return (int32)(QStringView_ToLong1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QStringView) ToLong2(ok *bool, base int) int32 {
	return (int32)(QStringView_ToLong2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QStringView) ToULong1(ok *bool) uint32 {
	return (uint32)(QStringView_ToULong1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QStringView) ToULong2(ok *bool, base int) uint32 {
	return (uint32)(QStringView_ToULong2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QStringView) ToLongLong1(ok *bool) int64 {
	return (int64)(QStringView_ToLongLong1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QStringView) ToLongLong2(ok *bool, base int) int64 {
	return (int64)(QStringView_ToLongLong2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QStringView) ToULongLong1(ok *bool) uint64 {
	return (uint64)(QStringView_ToULongLong1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QStringView) ToULongLong2(ok *bool, base int) uint64 {
	return (uint64)(QStringView_ToULongLong2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QStringView) ToFloat1(ok *bool) float32 {
	return (float32)(QStringView_ToFloat1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QStringView) ToDouble1(ok *bool) float64 {
	return (float64)(QStringView_ToDouble1(this.h, (*bool)(unsafe.Pointer(ok))))
}
