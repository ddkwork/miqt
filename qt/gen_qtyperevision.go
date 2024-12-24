package qt

import (
	"unsafe"
)

type QTypeRevision struct {
	h          uintptr
	isSubclass bool
}

// NewQTypeRevision constructs a new QTypeRevision object.
func NewQTypeRevision() *QTypeRevision {

	ret := newQTypeRevision(QTypeRevision_new())
	ret.isSubclass = true
	return ret
}

// NewQTypeRevision2 constructs a new QTypeRevision object.
func NewQTypeRevision2(param1 *QTypeRevision) *QTypeRevision {

	ret := newQTypeRevision(QTypeRevision_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func QTypeRevision_Zero() *QTypeRevision {
	_goptr := newQTypeRevision(QTypeRevision_Zero())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTypeRevision) HasMajorVersion() bool {
	return (bool)(QTypeRevision_HasMajorVersion(this.h))
}

func (this *QTypeRevision) MajorVersion() byte {
	return (byte)(QTypeRevision_MajorVersion(this.h))
}

func (this *QTypeRevision) HasMinorVersion() bool {
	return (bool)(QTypeRevision_HasMinorVersion(this.h))
}

func (this *QTypeRevision) MinorVersion() byte {
	return (byte)(QTypeRevision_MinorVersion(this.h))
}

func (this *QTypeRevision) IsValid() bool {
	return (bool)(QTypeRevision_IsValid(this.h))
}
