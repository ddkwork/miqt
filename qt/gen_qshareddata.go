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
	g := newQSharedData(QSharedData_new())
	g.isSubclass = true
	return g
}

// NewQSharedData2 constructs a new QSharedData object.
func NewQSharedData2(param1 *QSharedData) *QSharedData {
	g := newQSharedData(QSharedData_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}

type QAdoptSharedDataTag struct {
	h          uintptr
	isSubclass bool
}

// NewQAdoptSharedDataTag constructs a new QAdoptSharedDataTag object.
func NewQAdoptSharedDataTag() *QAdoptSharedDataTag {
	g := newQAdoptSharedDataTag(QAdoptSharedDataTag_new())
	g.isSubclass = true
	return g
}
