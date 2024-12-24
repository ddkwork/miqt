package qt

import (
	"unsafe"
)

type QBindingStatus struct {
	h          uintptr
	isSubclass bool
}
type QBindingStorage struct {
	h          uintptr
	isSubclass bool
}

// NewQBindingStorage constructs a new QBindingStorage object.
func NewQBindingStorage() *QBindingStorage {
	g := newQBindingStorage(QBindingStorage_new())
	g.isSubclass = true
	return g
}

func (this *QBindingStorage) IsEmpty() bool {
	return (bool)(QBindingStorage_IsEmpty(this.h))
}

func (this *QBindingStorage) IsValid() bool {
	return (bool)(QBindingStorage_IsValid(this.h))
}

func (this *QBindingStorage) RegisterDependency(data *QUntypedPropertyData) {
	QBindingStorage_RegisterDependency(this.h, data.cPointer())
}
