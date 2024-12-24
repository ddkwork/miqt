package qt

import (
	"unsafe"
)

type QLatin1String struct {
	h          uintptr
	isSubclass bool
}

// NewQLatin1String constructs a new QLatin1String object.
func NewQLatin1String() *QLatin1String {

	ret := newQLatin1String(QLatin1String_new())
	ret.isSubclass = true
	return ret
}

// NewQLatin1String2 constructs a new QLatin1String object.
func NewQLatin1String2(s string) *QLatin1String {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))

	ret := newQLatin1String(QLatin1String_new2(s_Cstring))
	ret.isSubclass = true
	return ret
}

// NewQLatin1String3 constructs a new QLatin1String object.
func NewQLatin1String3(f string, l string) *QLatin1String {
	f_Cstring := CString(f)
	defer free(unsafe.Pointer(f_Cstring))
	l_Cstring := CString(l)
	defer free(unsafe.Pointer(l_Cstring))

	ret := newQLatin1String(QLatin1String_new3(f_Cstring, l_Cstring))
	ret.isSubclass = true
	return ret
}

// NewQLatin1String4 constructs a new QLatin1String object.
func NewQLatin1String4(s string, sz int64) *QLatin1String {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))

	ret := newQLatin1String(QLatin1String_new4(s_Cstring, (ptrdiff_t)(sz)))
	ret.isSubclass = true
	return ret
}

// NewQLatin1String5 constructs a new QLatin1String object.
func NewQLatin1String5(s []byte) *QLatin1String {
	s_alias := struct_miqt_string{}
	s_alias.data = (char)(unsafe.Pointer(&s[0]))
	s_alias.len = size_t(len(s))

	ret := newQLatin1String(QLatin1String_new5(s_alias))
	ret.isSubclass = true
	return ret
}

// NewQLatin1String6 constructs a new QLatin1String object.
func NewQLatin1String6(s QByteArrayView) *QLatin1String {

	ret := newQLatin1String(QLatin1String_new6(s.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QLatin1String) ToString() string {
	var _ms struct_miqt_string = QLatin1String_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLatin1String) ToUtf8() []byte {
	var _bytearray struct_miqt_string = QLatin1String_ToUtf8(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QLatin1String) Latin1() string {
	_ret := QLatin1String_Latin1(this.h)
	return GoString(_ret)
}

func (this *QLatin1String) Size() int64 {
	return (int64)(QLatin1String_Size(this.h))
}

func (this *QLatin1String) Data() string {
	_ret := QLatin1String_Data(this.h)
	return GoString(_ret)
}

func (this *QLatin1String) ConstData() string {
	_ret := QLatin1String_ConstData(this.h)
	return GoString(_ret)
}

func (this *QLatin1String) ConstBegin() string {
	_ret := QLatin1String_ConstBegin(this.h)
	return GoString(_ret)
}

func (this *QLatin1String) ConstEnd() string {
	_ret := QLatin1String_ConstEnd(this.h)
	return GoString(_ret)
}

func (this *QLatin1String) First() *QLatin1Char {
	_goptr := newQLatin1Char(QLatin1String_First(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLatin1String) Last() *QLatin1Char {
	_goptr := newQLatin1Char(QLatin1String_Last(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLatin1String) Length() int64 {
	return (int64)(QLatin1String_Length(this.h))
}

func (this *QLatin1String) IsNull() bool {
	return (bool)(QLatin1String_IsNull(this.h))
}

func (this *QLatin1String) IsEmpty() bool {
	return (bool)(QLatin1String_IsEmpty(this.h))
}

func (this *QLatin1String) Empty() bool {
	return (bool)(QLatin1String_Empty(this.h))
}

func (this *QLatin1String) At(i int64) *QLatin1Char {
	_goptr := newQLatin1Char(QLatin1String_At(this.h, (ptrdiff_t)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLatin1String) OperatorSubscript(i int64) *QLatin1Char {
	_goptr := newQLatin1Char(QLatin1String_OperatorSubscript(this.h, (ptrdiff_t)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLatin1String) Front() *QLatin1Char {
	_goptr := newQLatin1Char(QLatin1String_Front(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLatin1String) Back() *QLatin1Char {
	_goptr := newQLatin1Char(QLatin1String_Back(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLatin1String) CompareWithQChar(c QChar) int {
	return (int)(QLatin1String_CompareWithQChar(this.h, c.cPointer()))
}

func (this *QLatin1String) Compare3(c QChar, cs CaseSensitivity) int {
	return (int)(QLatin1String_Compare3(this.h, c.cPointer(), (int)(cs)))
}

func (this *QLatin1String) StartsWithWithQChar(c QChar) bool {
	return (bool)(QLatin1String_StartsWithWithQChar(this.h, c.cPointer()))
}

func (this *QLatin1String) StartsWith2(c QChar, cs CaseSensitivity) bool {
	return (bool)(QLatin1String_StartsWith2(this.h, c.cPointer(), (int)(cs)))
}

func (this *QLatin1String) EndsWithWithQChar(c QChar) bool {
	return (bool)(QLatin1String_EndsWithWithQChar(this.h, c.cPointer()))
}

func (this *QLatin1String) EndsWith2(c QChar, cs CaseSensitivity) bool {
	return (bool)(QLatin1String_EndsWith2(this.h, c.cPointer(), (int)(cs)))
}

func (this *QLatin1String) IndexOfWithQChar(c QChar) int64 {
	return (int64)(QLatin1String_IndexOfWithQChar(this.h, c.cPointer()))
}

func (this *QLatin1String) ContainsWithQChar(c QChar) bool {
	return (bool)(QLatin1String_ContainsWithQChar(this.h, c.cPointer()))
}

func (this *QLatin1String) LastIndexOfWithQChar(c QChar) int64 {
	return (int64)(QLatin1String_LastIndexOfWithQChar(this.h, c.cPointer()))
}

func (this *QLatin1String) LastIndexOf4(c QChar, from int64) int64 {
	return (int64)(QLatin1String_LastIndexOf4(this.h, c.cPointer(), (ptrdiff_t)(from)))
}

func (this *QLatin1String) CountWithCh(ch QChar) int64 {
	return (int64)(QLatin1String_CountWithCh(this.h, ch.cPointer()))
}

func (this *QLatin1String) ToShort() int16 {
	return (int16)(QLatin1String_ToShort(this.h))
}

func (this *QLatin1String) ToUShort() uint16 {
	return (uint16)(QLatin1String_ToUShort(this.h))
}

func (this *QLatin1String) ToInt() int {
	return (int)(QLatin1String_ToInt(this.h))
}

func (this *QLatin1String) ToUInt() uint {
	return (uint)(QLatin1String_ToUInt(this.h))
}

func (this *QLatin1String) ToLong() int32 {
	return (int32)(QLatin1String_ToLong(this.h))
}

func (this *QLatin1String) ToULong() uint32 {
	return (uint32)(QLatin1String_ToULong(this.h))
}

func (this *QLatin1String) ToLongLong() int64 {
	return (int64)(QLatin1String_ToLongLong(this.h))
}

func (this *QLatin1String) ToULongLong() uint64 {
	return (uint64)(QLatin1String_ToULongLong(this.h))
}

func (this *QLatin1String) ToFloat() float32 {
	return (float32)(QLatin1String_ToFloat(this.h))
}

func (this *QLatin1String) ToDouble() float64 {
	return (float64)(QLatin1String_ToDouble(this.h))
}

func (this *QLatin1String) Begin() const_iterator {
	xxxxxxxxx
}

func (this *QLatin1String) Cbegin() const_iterator {
	xxxxxxxxx
}

func (this *QLatin1String) End() const_iterator {
	xxxxxxxxx
}

func (this *QLatin1String) Cend() const_iterator {
	xxxxxxxxx
}

func (this *QLatin1String) Rbegin() const_reverse_iterator {
	xxxxxxxxx
}

func (this *QLatin1String) Crbegin() const_reverse_iterator {
	xxxxxxxxx
}

func (this *QLatin1String) Rend() const_reverse_iterator {
	xxxxxxxxx
}

func (this *QLatin1String) Crend() const_reverse_iterator {
	xxxxxxxxx
}

func (this *QLatin1String) MaxSize() int64 {
	return (int64)(QLatin1String_MaxSize(this.h))
}

func QLatin1String_MaxSize2() int64 {
	return (int64)(QLatin1String_MaxSize2())
}

func (this *QLatin1String) Chop(n int64) {
	QLatin1String_Chop(this.h, (ptrdiff_t)(n))
}

func (this *QLatin1String) Truncate(n int64) {
	QLatin1String_Truncate(this.h, (ptrdiff_t)(n))
}

func (this *QLatin1String) IndexOf23(c QChar, from int64) int64 {
	return (int64)(QLatin1String_IndexOf23(this.h, c.cPointer(), (ptrdiff_t)(from)))
}

func (this *QLatin1String) IndexOf33(c QChar, from int64, cs CaseSensitivity) int64 {
	return (int64)(QLatin1String_IndexOf33(this.h, c.cPointer(), (ptrdiff_t)(from), (int)(cs)))
}

func (this *QLatin1String) Contains23(c QChar, cs CaseSensitivity) bool {
	return (bool)(QLatin1String_Contains23(this.h, c.cPointer(), (int)(cs)))
}

func (this *QLatin1String) LastIndexOf24(c QChar, cs CaseSensitivity) int64 {
	return (int64)(QLatin1String_LastIndexOf24(this.h, c.cPointer(), (int)(cs)))
}

func (this *QLatin1String) LastIndexOf34(c QChar, from int64, cs CaseSensitivity) int64 {
	return (int64)(QLatin1String_LastIndexOf34(this.h, c.cPointer(), (ptrdiff_t)(from), (int)(cs)))
}

func (this *QLatin1String) Count23(ch QChar, cs CaseSensitivity) int64 {
	return (int64)(QLatin1String_Count23(this.h, ch.cPointer(), (int)(cs)))
}

func (this *QLatin1String) ToShort1(ok *bool) int16 {
	return (int16)(QLatin1String_ToShort1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QLatin1String) ToShort2(ok *bool, base int) int16 {
	return (int16)(QLatin1String_ToShort2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QLatin1String) ToUShort1(ok *bool) uint16 {
	return (uint16)(QLatin1String_ToUShort1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QLatin1String) ToUShort2(ok *bool, base int) uint16 {
	return (uint16)(QLatin1String_ToUShort2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QLatin1String) ToInt1(ok *bool) int {
	return (int)(QLatin1String_ToInt1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QLatin1String) ToInt2(ok *bool, base int) int {
	return (int)(QLatin1String_ToInt2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QLatin1String) ToUInt1(ok *bool) uint {
	return (uint)(QLatin1String_ToUInt1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QLatin1String) ToUInt2(ok *bool, base int) uint {
	return (uint)(QLatin1String_ToUInt2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QLatin1String) ToLong1(ok *bool) int32 {
	return (int32)(QLatin1String_ToLong1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QLatin1String) ToLong2(ok *bool, base int) int32 {
	return (int32)(QLatin1String_ToLong2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QLatin1String) ToULong1(ok *bool) uint32 {
	return (uint32)(QLatin1String_ToULong1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QLatin1String) ToULong2(ok *bool, base int) uint32 {
	return (uint32)(QLatin1String_ToULong2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QLatin1String) ToLongLong1(ok *bool) int64 {
	return (int64)(QLatin1String_ToLongLong1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QLatin1String) ToLongLong2(ok *bool, base int) int64 {
	return (int64)(QLatin1String_ToLongLong2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QLatin1String) ToULongLong1(ok *bool) uint64 {
	return (uint64)(QLatin1String_ToULongLong1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QLatin1String) ToULongLong2(ok *bool, base int) uint64 {
	return (uint64)(QLatin1String_ToULongLong2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QLatin1String) ToFloat1(ok *bool) float32 {
	return (float32)(QLatin1String_ToFloat1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QLatin1String) ToDouble1(ok *bool) float64 {
	return (float64)(QLatin1String_ToDouble1(this.h, (*bool)(unsafe.Pointer(ok))))
}
