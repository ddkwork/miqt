package qt

import (
	"unsafe"
)

type QPdfOutputIntent struct {
	h          uintptr
	isSubclass bool
}

// NewQPdfOutputIntent constructs a new QPdfOutputIntent object.
func NewQPdfOutputIntent() *QPdfOutputIntent {

	ret := newQPdfOutputIntent(QPdfOutputIntent_new())
	ret.isSubclass = true
	return ret
}

// NewQPdfOutputIntent2 constructs a new QPdfOutputIntent object.
func NewQPdfOutputIntent2(other *QPdfOutputIntent) *QPdfOutputIntent {

	ret := newQPdfOutputIntent(QPdfOutputIntent_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QPdfOutputIntent) OperatorAssign(other *QPdfOutputIntent) {
	QPdfOutputIntent_OperatorAssign(this.h, other.cPointer())
}

func (this *QPdfOutputIntent) Swap(other *QPdfOutputIntent) {
	QPdfOutputIntent_Swap(this.h, other.cPointer())
}

func (this *QPdfOutputIntent) OutputConditionIdentifier() string {
	var _ms struct_miqt_string = QPdfOutputIntent_OutputConditionIdentifier(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPdfOutputIntent) SetOutputConditionIdentifier(identifier string) {
	identifier_ms := struct_miqt_string{}
	identifier_ms.data = CString(identifier)
	identifier_ms.len = size_t(len(identifier))
	defer free(unsafe.Pointer(identifier_ms.data))
	QPdfOutputIntent_SetOutputConditionIdentifier(this.h, identifier_ms)
}

func (this *QPdfOutputIntent) OutputCondition() string {
	var _ms struct_miqt_string = QPdfOutputIntent_OutputCondition(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPdfOutputIntent) SetOutputCondition(condition string) {
	condition_ms := struct_miqt_string{}
	condition_ms.data = CString(condition)
	condition_ms.len = size_t(len(condition))
	defer free(unsafe.Pointer(condition_ms.data))
	QPdfOutputIntent_SetOutputCondition(this.h, condition_ms)
}

func (this *QPdfOutputIntent) RegistryName() *QUrl {
	_goptr := newQUrl(QPdfOutputIntent_RegistryName(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPdfOutputIntent) SetRegistryName(name *QUrl) {
	QPdfOutputIntent_SetRegistryName(this.h, name.cPointer())
}

func (this *QPdfOutputIntent) OutputProfile() *QColorSpace {
	_goptr := newQColorSpace(QPdfOutputIntent_OutputProfile(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPdfOutputIntent) SetOutputProfile(profile *QColorSpace) {
	QPdfOutputIntent_SetOutputProfile(this.h, profile.cPointer())
}
