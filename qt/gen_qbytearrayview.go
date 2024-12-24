package qt

import (
	"unsafe"
)

type QByteArrayView struct {
	h          uintptr
	isSubclass bool
}

// NewQByteArrayView constructs a new QByteArrayView object.
func NewQByteArrayView() *QByteArrayView {

	ret := newQByteArrayView(QByteArrayView_new())
	ret.isSubclass = true
	return ret
}

// NewQByteArrayView2 constructs a new QByteArrayView object.
func NewQByteArrayView2(param1 *QByteArrayView) *QByteArrayView {

	ret := newQByteArrayView(QByteArrayView_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QByteArrayView) ToByteArray() []byte {
	var _bytearray struct_miqt_string = QByteArrayView_ToByteArray(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QByteArrayView) Size() int64 {
	return (int64)(QByteArrayView_Size(this.h))
}

func (this *QByteArrayView) Data() const_pointer {
	xxxxxxxxx
}

func (this *QByteArrayView) ConstData() const_pointer {
	xxxxxxxxx
}

func (this *QByteArrayView) OperatorSubscript(n int64) int8 {
	return (int8)(QByteArrayView_OperatorSubscript(this.h, (ptrdiff_t)(n)))
}

func (this *QByteArrayView) At(n int64) int8 {
	return (int8)(QByteArrayView_At(this.h, (ptrdiff_t)(n)))
}

func (this *QByteArrayView) First(n int64) *QByteArrayView {
	_goptr := newQByteArrayView(QByteArrayView_First(this.h, (ptrdiff_t)(n)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QByteArrayView) Last(n int64) *QByteArrayView {
	_goptr := newQByteArrayView(QByteArrayView_Last(this.h, (ptrdiff_t)(n)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QByteArrayView) Sliced(pos int64) *QByteArrayView {
	_goptr := newQByteArrayView(QByteArrayView_Sliced(this.h, (ptrdiff_t)(pos)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QByteArrayView) Sliced2(pos int64, n int64) *QByteArrayView {
	_goptr := newQByteArrayView(QByteArrayView_Sliced2(this.h, (ptrdiff_t)(pos), (ptrdiff_t)(n)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QByteArrayView) Slice(pos int64) *QByteArrayView {
	return newQByteArrayView(QByteArrayView_Slice(this.h, (ptrdiff_t)(pos)))
}

func (this *QByteArrayView) Slice2(pos int64, n int64) *QByteArrayView {
	return newQByteArrayView(QByteArrayView_Slice2(this.h, (ptrdiff_t)(pos), (ptrdiff_t)(n)))
}

func (this *QByteArrayView) Chopped(lenVal int64) *QByteArrayView {
	_goptr := newQByteArrayView(QByteArrayView_Chopped(this.h, (ptrdiff_t)(lenVal)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QByteArrayView) Left(n int64) *QByteArrayView {
	_goptr := newQByteArrayView(QByteArrayView_Left(this.h, (ptrdiff_t)(n)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QByteArrayView) Right(n int64) *QByteArrayView {
	_goptr := newQByteArrayView(QByteArrayView_Right(this.h, (ptrdiff_t)(n)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QByteArrayView) Mid(pos int64) *QByteArrayView {
	_goptr := newQByteArrayView(QByteArrayView_Mid(this.h, (ptrdiff_t)(pos)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QByteArrayView) Truncate(n int64) {
	QByteArrayView_Truncate(this.h, (ptrdiff_t)(n))
}

func (this *QByteArrayView) Chop(n int64) {
	QByteArrayView_Chop(this.h, (ptrdiff_t)(n))
}

func (this *QByteArrayView) Trimmed() *QByteArrayView {
	_goptr := newQByteArrayView(QByteArrayView_Trimmed(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QByteArrayView) ToShort() int16 {
	return (int16)(QByteArrayView_ToShort(this.h))
}

func (this *QByteArrayView) ToUShort() uint16 {
	return (uint16)(QByteArrayView_ToUShort(this.h))
}

func (this *QByteArrayView) ToInt() int {
	return (int)(QByteArrayView_ToInt(this.h))
}

func (this *QByteArrayView) ToUInt() uint {
	return (uint)(QByteArrayView_ToUInt(this.h))
}

func (this *QByteArrayView) ToLong() int32 {
	return (int32)(QByteArrayView_ToLong(this.h))
}

func (this *QByteArrayView) ToULong() uint32 {
	return (uint32)(QByteArrayView_ToULong(this.h))
}

func (this *QByteArrayView) ToLongLong() int64 {
	return (int64)(QByteArrayView_ToLongLong(this.h))
}

func (this *QByteArrayView) ToULongLong() uint64 {
	return (uint64)(QByteArrayView_ToULongLong(this.h))
}

func (this *QByteArrayView) ToFloat() float32 {
	return (float32)(QByteArrayView_ToFloat(this.h))
}

func (this *QByteArrayView) ToDouble() float64 {
	return (float64)(QByteArrayView_ToDouble(this.h))
}

func (this *QByteArrayView) StartsWith(other QByteArrayView) bool {
	return (bool)(QByteArrayView_StartsWith(this.h, other.cPointer()))
}

func (this *QByteArrayView) StartsWithWithChar(c int8) bool {
	return (bool)(QByteArrayView_StartsWithWithChar(this.h, (char)(c)))
}

func (this *QByteArrayView) EndsWith(other QByteArrayView) bool {
	return (bool)(QByteArrayView_EndsWith(this.h, other.cPointer()))
}

func (this *QByteArrayView) EndsWithWithChar(c int8) bool {
	return (bool)(QByteArrayView_EndsWithWithChar(this.h, (char)(c)))
}

func (this *QByteArrayView) IndexOf(a QByteArrayView) int64 {
	return (int64)(QByteArrayView_IndexOf(this.h, a.cPointer()))
}

func (this *QByteArrayView) IndexOfWithCh(ch int8) int64 {
	return (int64)(QByteArrayView_IndexOfWithCh(this.h, (char)(ch)))
}

func (this *QByteArrayView) Contains(a QByteArrayView) bool {
	return (bool)(QByteArrayView_Contains(this.h, a.cPointer()))
}

func (this *QByteArrayView) ContainsWithChar(c int8) bool {
	return (bool)(QByteArrayView_ContainsWithChar(this.h, (char)(c)))
}

func (this *QByteArrayView) LastIndexOf(a QByteArrayView) int64 {
	return (int64)(QByteArrayView_LastIndexOf(this.h, a.cPointer()))
}

func (this *QByteArrayView) LastIndexOf2(a QByteArrayView, from int64) int64 {
	return (int64)(QByteArrayView_LastIndexOf2(this.h, a.cPointer(), (ptrdiff_t)(from)))
}

func (this *QByteArrayView) LastIndexOfWithCh(ch int8) int64 {
	return (int64)(QByteArrayView_LastIndexOfWithCh(this.h, (char)(ch)))
}

func (this *QByteArrayView) Count(a QByteArrayView) int64 {
	return (int64)(QByteArrayView_Count(this.h, a.cPointer()))
}

func (this *QByteArrayView) CountWithCh(ch int8) int64 {
	return (int64)(QByteArrayView_CountWithCh(this.h, (char)(ch)))
}

func (this *QByteArrayView) Compare(a QByteArrayView) int {
	return (int)(QByteArrayView_Compare(this.h, a.cPointer()))
}

func (this *QByteArrayView) IsValidUtf8() bool {
	return (bool)(QByteArrayView_IsValidUtf8(this.h))
}

func (this *QByteArrayView) Begin() const_iterator {
	xxxxxxxxx
}

func (this *QByteArrayView) End() const_iterator {
	xxxxxxxxx
}

func (this *QByteArrayView) Cbegin() const_iterator {
	xxxxxxxxx
}

func (this *QByteArrayView) Cend() const_iterator {
	xxxxxxxxx
}

func (this *QByteArrayView) Rbegin() const_reverse_iterator {
	xxxxxxxxx
}

func (this *QByteArrayView) Rend() const_reverse_iterator {
	xxxxxxxxx
}

func (this *QByteArrayView) Crbegin() const_reverse_iterator {
	xxxxxxxxx
}

func (this *QByteArrayView) Crend() const_reverse_iterator {
	xxxxxxxxx
}

func (this *QByteArrayView) Empty() bool {
	return (bool)(QByteArrayView_Empty(this.h))
}

func (this *QByteArrayView) Front() int8 {
	return (int8)(QByteArrayView_Front(this.h))
}

func (this *QByteArrayView) Back() int8 {
	return (int8)(QByteArrayView_Back(this.h))
}

func (this *QByteArrayView) MaxSize() int64 {
	return (int64)(QByteArrayView_MaxSize(this.h))
}

func (this *QByteArrayView) IsNull() bool {
	return (bool)(QByteArrayView_IsNull(this.h))
}

func (this *QByteArrayView) IsEmpty() bool {
	return (bool)(QByteArrayView_IsEmpty(this.h))
}

func (this *QByteArrayView) Length() int64 {
	return (int64)(QByteArrayView_Length(this.h))
}

func (this *QByteArrayView) First2() int8 {
	return (int8)(QByteArrayView_First2(this.h))
}

func (this *QByteArrayView) Last2() int8 {
	return (int8)(QByteArrayView_Last2(this.h))
}

func QByteArrayView_MaxSize2() int64 {
	return (int64)(QByteArrayView_MaxSize2())
}

func (this *QByteArrayView) Mid2(pos int64, n int64) *QByteArrayView {
	_goptr := newQByteArrayView(QByteArrayView_Mid2(this.h, (ptrdiff_t)(pos), (ptrdiff_t)(n)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QByteArrayView) ToShort1(ok *bool) int16 {
	return (int16)(QByteArrayView_ToShort1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QByteArrayView) ToShort2(ok *bool, base int) int16 {
	return (int16)(QByteArrayView_ToShort2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QByteArrayView) ToUShort1(ok *bool) uint16 {
	return (uint16)(QByteArrayView_ToUShort1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QByteArrayView) ToUShort2(ok *bool, base int) uint16 {
	return (uint16)(QByteArrayView_ToUShort2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QByteArrayView) ToInt1(ok *bool) int {
	return (int)(QByteArrayView_ToInt1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QByteArrayView) ToInt2(ok *bool, base int) int {
	return (int)(QByteArrayView_ToInt2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QByteArrayView) ToUInt1(ok *bool) uint {
	return (uint)(QByteArrayView_ToUInt1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QByteArrayView) ToUInt2(ok *bool, base int) uint {
	return (uint)(QByteArrayView_ToUInt2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QByteArrayView) ToLong1(ok *bool) int32 {
	return (int32)(QByteArrayView_ToLong1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QByteArrayView) ToLong2(ok *bool, base int) int32 {
	return (int32)(QByteArrayView_ToLong2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QByteArrayView) ToULong1(ok *bool) uint32 {
	return (uint32)(QByteArrayView_ToULong1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QByteArrayView) ToULong2(ok *bool, base int) uint32 {
	return (uint32)(QByteArrayView_ToULong2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QByteArrayView) ToLongLong1(ok *bool) int64 {
	return (int64)(QByteArrayView_ToLongLong1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QByteArrayView) ToLongLong2(ok *bool, base int) int64 {
	return (int64)(QByteArrayView_ToLongLong2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QByteArrayView) ToULongLong1(ok *bool) uint64 {
	return (uint64)(QByteArrayView_ToULongLong1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QByteArrayView) ToULongLong2(ok *bool, base int) uint64 {
	return (uint64)(QByteArrayView_ToULongLong2(this.h, (*bool)(unsafe.Pointer(ok)), (int)(base)))
}

func (this *QByteArrayView) ToFloat1(ok *bool) float32 {
	return (float32)(QByteArrayView_ToFloat1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QByteArrayView) ToDouble1(ok *bool) float64 {
	return (float64)(QByteArrayView_ToDouble1(this.h, (*bool)(unsafe.Pointer(ok))))
}

func (this *QByteArrayView) IndexOf2(a QByteArrayView, from int64) int64 {
	return (int64)(QByteArrayView_IndexOf2(this.h, a.cPointer(), (ptrdiff_t)(from)))
}

func (this *QByteArrayView) IndexOf22(ch int8, from int64) int64 {
	return (int64)(QByteArrayView_IndexOf22(this.h, (char)(ch), (ptrdiff_t)(from)))
}

func (this *QByteArrayView) LastIndexOf22(ch int8, from int64) int64 {
	return (int64)(QByteArrayView_LastIndexOf22(this.h, (char)(ch), (ptrdiff_t)(from)))
}

func (this *QByteArrayView) Compare2(a QByteArrayView, cs CaseSensitivity) int {
	return (int)(QByteArrayView_Compare2(this.h, a.cPointer(), (int)(cs)))
}
