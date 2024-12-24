package qt

import (
	"unsafe"
)

type QArrayData__AllocationOption int

const (
	QArrayData__Grow     QArrayData__AllocationOption = 0
	QArrayData__KeepSize QArrayData__AllocationOption = 1
)

type QArrayData__GrowthPosition int

const (
	QArrayData__GrowsAtEnd       QArrayData__GrowthPosition = 0
	QArrayData__GrowsAtBeginning QArrayData__GrowthPosition = 1
)

type QArrayData__ArrayOption int

const (
	QArrayData__ArrayOptionDefault QArrayData__ArrayOption = 0
	QArrayData__CapacityReserved   QArrayData__ArrayOption = 1
)

type QtPrivate__QContainerImplHelper__CutResult int

const (
	QtPrivate__QContainerImplHelper__Null   QtPrivate__QContainerImplHelper__CutResult = 0
	QtPrivate__QContainerImplHelper__Empty  QtPrivate__QContainerImplHelper__CutResult = 1
	QtPrivate__QContainerImplHelper__Full   QtPrivate__QContainerImplHelper__CutResult = 2
	QtPrivate__QContainerImplHelper__Subset QtPrivate__QContainerImplHelper__CutResult = 3
)

type QArrayData struct {
	h          uintptr
	isSubclass bool
}

func (this *QArrayData) AllocatedCapacity() int64 {
	return (int64)(QArrayData_AllocatedCapacity(this.h))
}

func (this *QArrayData) ConstAllocatedCapacity() int64 {
	return (int64)(QArrayData_ConstAllocatedCapacity(this.h))
}

func (this *QArrayData) Ref() bool {
	return (bool)(QArrayData_Ref(this.h))
}

func (this *QArrayData) Deref() bool {
	return (bool)(QArrayData_Deref(this.h))
}

func (this *QArrayData) IsShared() bool {
	return (bool)(QArrayData_IsShared(this.h))
}

func (this *QArrayData) NeedsDetach() bool {
	return (bool)(QArrayData_NeedsDetach(this.h))
}

func (this *QArrayData) DetachCapacity(newSize int64) int64 {
	return (int64)(QArrayData_DetachCapacity(this.h, (ptrdiff_t)(newSize)))
}

func QArrayData_Deallocate(data *QArrayData, objectSize int64, alignment int64) {
	QArrayData_Deallocate(data.cPointer(), (ptrdiff_t)(objectSize), (ptrdiff_t)(alignment))
}
