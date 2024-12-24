package qt

import (
	"unsafe"
)

type QAnyStringView struct {
	h          uintptr
	isSubclass bool
}

// NewQAnyStringView constructs a new QAnyStringView object.
func NewQAnyStringView() *QAnyStringView {

	ret := newQAnyStringView(QAnyStringView_new())
	ret.isSubclass = true
	return ret
}

// NewQAnyStringView2 constructs a new QAnyStringView object.
func NewQAnyStringView2(str []byte) *QAnyStringView {
	str_alias := struct_miqt_string{}
	str_alias.data = (char)(unsafe.Pointer(&str[0]))
	str_alias.len = size_t(len(str))

	ret := newQAnyStringView(QAnyStringView_new2(str_alias))
	ret.isSubclass = true
	return ret
}

// NewQAnyStringView3 constructs a new QAnyStringView object.
func NewQAnyStringView3(str string) *QAnyStringView {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))

	ret := newQAnyStringView(QAnyStringView_new3(str_ms))
	ret.isSubclass = true
	return ret
}

// NewQAnyStringView4 constructs a new QAnyStringView object.
func NewQAnyStringView4(param1 *QAnyStringView) *QAnyStringView {

	ret := newQAnyStringView(QAnyStringView_new4(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QAnyStringView) Mid(pos int64) *QAnyStringView {
	_goptr := newQAnyStringView(QAnyStringView_Mid(this.h, (ptrdiff_t)(pos)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAnyStringView) Left(n int64) *QAnyStringView {
	_goptr := newQAnyStringView(QAnyStringView_Left(this.h, (ptrdiff_t)(n)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAnyStringView) Right(n int64) *QAnyStringView {
	_goptr := newQAnyStringView(QAnyStringView_Right(this.h, (ptrdiff_t)(n)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAnyStringView) Sliced(pos int64) *QAnyStringView {
	_goptr := newQAnyStringView(QAnyStringView_Sliced(this.h, (ptrdiff_t)(pos)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAnyStringView) Sliced2(pos int64, n int64) *QAnyStringView {
	_goptr := newQAnyStringView(QAnyStringView_Sliced2(this.h, (ptrdiff_t)(pos), (ptrdiff_t)(n)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAnyStringView) First(n int64) *QAnyStringView {
	_goptr := newQAnyStringView(QAnyStringView_First(this.h, (ptrdiff_t)(n)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAnyStringView) Last(n int64) *QAnyStringView {
	_goptr := newQAnyStringView(QAnyStringView_Last(this.h, (ptrdiff_t)(n)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAnyStringView) Chopped(n int64) *QAnyStringView {
	_goptr := newQAnyStringView(QAnyStringView_Chopped(this.h, (ptrdiff_t)(n)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAnyStringView) Slice(pos int64) *QAnyStringView {
	return newQAnyStringView(QAnyStringView_Slice(this.h, (ptrdiff_t)(pos)))
}

func (this *QAnyStringView) Slice2(pos int64, n int64) *QAnyStringView {
	return newQAnyStringView(QAnyStringView_Slice2(this.h, (ptrdiff_t)(pos), (ptrdiff_t)(n)))
}

func (this *QAnyStringView) Truncate(n int64) {
	QAnyStringView_Truncate(this.h, (ptrdiff_t)(n))
}

func (this *QAnyStringView) Chop(n int64) {
	QAnyStringView_Chop(this.h, (ptrdiff_t)(n))
}

func (this *QAnyStringView) ToString() string {
	var _ms struct_miqt_string = QAnyStringView_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAnyStringView) Size() int64 {
	return (int64)(QAnyStringView_Size(this.h))
}

func (this *QAnyStringView) Data() unsafe.Pointer {
	return (unsafe.Pointer)(QAnyStringView_Data(this.h))
}

func QAnyStringView_Compare(lhs QAnyStringView, rhs QAnyStringView) int {
	return (int)(QAnyStringView_Compare(lhs.cPointer(), rhs.cPointer()))
}

func QAnyStringView_Equal(lhs QAnyStringView, rhs QAnyStringView) bool {
	return (bool)(QAnyStringView_Equal(lhs.cPointer(), rhs.cPointer()))
}

func (this *QAnyStringView) Front() *QChar {
	_goptr := newQChar(QAnyStringView_Front(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAnyStringView) Back() *QChar {
	_goptr := newQChar(QAnyStringView_Back(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAnyStringView) Empty() bool {
	return (bool)(QAnyStringView_Empty(this.h))
}

func (this *QAnyStringView) SizeBytes() int64 {
	return (int64)(QAnyStringView_SizeBytes(this.h))
}

func (this *QAnyStringView) MaxSize() int64 {
	return (int64)(QAnyStringView_MaxSize(this.h))
}

func (this *QAnyStringView) IsNull() bool {
	return (bool)(QAnyStringView_IsNull(this.h))
}

func (this *QAnyStringView) IsEmpty() bool {
	return (bool)(QAnyStringView_IsEmpty(this.h))
}

func (this *QAnyStringView) Length() int64 {
	return (int64)(QAnyStringView_Length(this.h))
}

func (this *QAnyStringView) Mid2(pos int64, n int64) *QAnyStringView {
	_goptr := newQAnyStringView(QAnyStringView_Mid2(this.h, (ptrdiff_t)(pos), (ptrdiff_t)(n)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QAnyStringView_Compare3(lhs QAnyStringView, rhs QAnyStringView, cs CaseSensitivity) int {
	return (int)(QAnyStringView_Compare3(lhs.cPointer(), rhs.cPointer(), (int)(cs)))
}
