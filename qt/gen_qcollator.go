package qt

import (
	"unsafe"
)

type QCollatorSortKey struct {
	h          uintptr
	isSubclass bool
}

// NewQCollatorSortKey constructs a new QCollatorSortKey object.
func NewQCollatorSortKey(other *QCollatorSortKey) *QCollatorSortKey {

	ret := newQCollatorSortKey(QCollatorSortKey_new(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QCollatorSortKey) OperatorAssign(other *QCollatorSortKey) {
	QCollatorSortKey_OperatorAssign(this.h, other.cPointer())
}

func (this *QCollatorSortKey) Swap(other *QCollatorSortKey) {
	QCollatorSortKey_Swap(this.h, other.cPointer())
}

func (this *QCollatorSortKey) Compare(key *QCollatorSortKey) int {
	return (int)(QCollatorSortKey_Compare(this.h, key.cPointer()))
}

type QCollator struct {
	h          uintptr
	isSubclass bool
}

// NewQCollator constructs a new QCollator object.
func NewQCollator() *QCollator {

	ret := newQCollator(QCollator_new())
	ret.isSubclass = true
	return ret
}

// NewQCollator2 constructs a new QCollator object.
func NewQCollator2(locale *QLocale) *QCollator {

	ret := newQCollator(QCollator_new2(locale.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQCollator3 constructs a new QCollator object.
func NewQCollator3(param1 *QCollator) *QCollator {

	ret := newQCollator(QCollator_new3(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QCollator) OperatorAssign(param1 *QCollator) {
	QCollator_OperatorAssign(this.h, param1.cPointer())
}

func (this *QCollator) Swap(other *QCollator) {
	QCollator_Swap(this.h, other.cPointer())
}

func (this *QCollator) SetLocale(locale *QLocale) {
	QCollator_SetLocale(this.h, locale.cPointer())
}

func (this *QCollator) Locale() *QLocale {
	_goptr := newQLocale(QCollator_Locale(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCollator) CaseSensitivity() CaseSensitivity {
	return (CaseSensitivity)(QCollator_CaseSensitivity(this.h))
}

func (this *QCollator) SetCaseSensitivity(cs CaseSensitivity) {
	QCollator_SetCaseSensitivity(this.h, (int)(cs))
}

func (this *QCollator) SetNumericMode(on bool) {
	QCollator_SetNumericMode(this.h, (bool)(on))
}

func (this *QCollator) NumericMode() bool {
	return (bool)(QCollator_NumericMode(this.h))
}

func (this *QCollator) SetIgnorePunctuation(on bool) {
	QCollator_SetIgnorePunctuation(this.h, (bool)(on))
}

func (this *QCollator) IgnorePunctuation() bool {
	return (bool)(QCollator_IgnorePunctuation(this.h))
}

func (this *QCollator) Compare(s1 string, s2 string) int {
	s1_ms := struct_miqt_string{}
	s1_ms.data = CString(s1)
	s1_ms.len = size_t(len(s1))
	defer free(unsafe.Pointer(s1_ms.data))
	s2_ms := struct_miqt_string{}
	s2_ms.data = CString(s2)
	s2_ms.len = size_t(len(s2))
	defer free(unsafe.Pointer(s2_ms.data))
	return (int)(QCollator_Compare(this.h, s1_ms, s2_ms))
}

func (this *QCollator) Compare2(s1 *QChar, len1 int64, s2 *QChar, len2 int64) int {
	return (int)(QCollator_Compare2(this.h, s1.cPointer(), (ptrdiff_t)(len1), s2.cPointer(), (ptrdiff_t)(len2)))
}

func (this *QCollator) OperatorCall(s1 string, s2 string) bool {
	s1_ms := struct_miqt_string{}
	s1_ms.data = CString(s1)
	s1_ms.len = size_t(len(s1))
	defer free(unsafe.Pointer(s1_ms.data))
	s2_ms := struct_miqt_string{}
	s2_ms.data = CString(s2)
	s2_ms.len = size_t(len(s2))
	defer free(unsafe.Pointer(s2_ms.data))
	return (bool)(QCollator_OperatorCall(this.h, s1_ms, s2_ms))
}

func (this *QCollator) SortKey(stringVal string) *QCollatorSortKey {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	_goptr := newQCollatorSortKey(QCollator_SortKey(this.h, stringVal_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
