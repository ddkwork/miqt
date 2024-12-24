package qt

import (
	"unsafe"
)

type QByteArrayMatcher struct {
	h          uintptr
	isSubclass bool
}

// NewQByteArrayMatcher constructs a new QByteArrayMatcher object.
func NewQByteArrayMatcher() *QByteArrayMatcher {

	ret := newQByteArrayMatcher(QByteArrayMatcher_new())
	ret.isSubclass = true
	return ret
}

// NewQByteArrayMatcher2 constructs a new QByteArrayMatcher object.
func NewQByteArrayMatcher2(pattern []byte) *QByteArrayMatcher {
	pattern_alias := struct_miqt_string{}
	pattern_alias.data = (char)(unsafe.Pointer(&pattern[0]))
	pattern_alias.len = size_t(len(pattern))

	ret := newQByteArrayMatcher(QByteArrayMatcher_new2(pattern_alias))
	ret.isSubclass = true
	return ret
}

// NewQByteArrayMatcher3 constructs a new QByteArrayMatcher object.
func NewQByteArrayMatcher3(pattern QByteArrayView) *QByteArrayMatcher {

	ret := newQByteArrayMatcher(QByteArrayMatcher_new3(pattern.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQByteArrayMatcher4 constructs a new QByteArrayMatcher object.
func NewQByteArrayMatcher4(pattern string) *QByteArrayMatcher {
	pattern_Cstring := CString(pattern)
	defer free(unsafe.Pointer(pattern_Cstring))

	ret := newQByteArrayMatcher(QByteArrayMatcher_new4(pattern_Cstring))
	ret.isSubclass = true
	return ret
}

// NewQByteArrayMatcher5 constructs a new QByteArrayMatcher object.
func NewQByteArrayMatcher5(other *QByteArrayMatcher) *QByteArrayMatcher {

	ret := newQByteArrayMatcher(QByteArrayMatcher_new5(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQByteArrayMatcher6 constructs a new QByteArrayMatcher object.
func NewQByteArrayMatcher6(pattern string, length int64) *QByteArrayMatcher {
	pattern_Cstring := CString(pattern)
	defer free(unsafe.Pointer(pattern_Cstring))

	ret := newQByteArrayMatcher(QByteArrayMatcher_new6(pattern_Cstring, (ptrdiff_t)(length)))
	ret.isSubclass = true
	return ret
}

func (this *QByteArrayMatcher) OperatorAssign(other *QByteArrayMatcher) {
	QByteArrayMatcher_OperatorAssign(this.h, other.cPointer())
}

func (this *QByteArrayMatcher) SetPattern(pattern []byte) {
	pattern_alias := struct_miqt_string{}
	pattern_alias.data = (char)(unsafe.Pointer(&pattern[0]))
	pattern_alias.len = size_t(len(pattern))
	QByteArrayMatcher_SetPattern(this.h, pattern_alias)
}

func (this *QByteArrayMatcher) IndexIn(str string, lenVal int64) int64 {
	str_Cstring := CString(str)
	defer free(unsafe.Pointer(str_Cstring))
	return (int64)(QByteArrayMatcher_IndexIn(this.h, str_Cstring, (ptrdiff_t)(lenVal)))
}

func (this *QByteArrayMatcher) IndexInWithData(data QByteArrayView) int64 {
	return (int64)(QByteArrayMatcher_IndexInWithData(this.h, data.cPointer()))
}

func (this *QByteArrayMatcher) Pattern() []byte {
	var _bytearray struct_miqt_string = QByteArrayMatcher_Pattern(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QByteArrayMatcher) IndexIn3(str string, lenVal int64, from int64) int64 {
	str_Cstring := CString(str)
	defer free(unsafe.Pointer(str_Cstring))
	return (int64)(QByteArrayMatcher_IndexIn3(this.h, str_Cstring, (ptrdiff_t)(lenVal), (ptrdiff_t)(from)))
}

func (this *QByteArrayMatcher) IndexIn2(data QByteArrayView, from int64) int64 {
	return (int64)(QByteArrayMatcher_IndexIn2(this.h, data.cPointer(), (ptrdiff_t)(from)))
}

type QStaticByteArrayMatcherBase struct {
	h          uintptr
	isSubclass bool
}
