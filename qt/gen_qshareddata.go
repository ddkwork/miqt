package qt

import (
	"unsafe"
)

type QSharedData struct {
	h          uintptr
	isSubclass bool
}

// NewQSharedData constructs a new QSharedData object.
func NewQSharedData() *QSharedData {

	ret := newQSharedData(QSharedData_new())
	ret.isSubclass = true
	return ret
}

// NewQSharedData2 constructs a new QSharedData object.
func NewQSharedData2(param1 *QSharedData) *QSharedData {

	ret := newQSharedData(QSharedData_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

type QAdoptSharedDataTag struct {
	h          uintptr
	isSubclass bool
}

// NewQAdoptSharedDataTag constructs a new QAdoptSharedDataTag object.
func NewQAdoptSharedDataTag() *QAdoptSharedDataTag {

	ret := newQAdoptSharedDataTag(QAdoptSharedDataTag_new())
	ret.isSubclass = true
	return ret
}
