package qt

import (
	"unsafe"
)

type QtMetaContainerPrivate__IteratorCapability byte

const (
	QtMetaContainerPrivate__InputCapability         QtMetaContainerPrivate__IteratorCapability = 1
	QtMetaContainerPrivate__ForwardCapability       QtMetaContainerPrivate__IteratorCapability = 2
	QtMetaContainerPrivate__BiDirectionalCapability QtMetaContainerPrivate__IteratorCapability = 4
	QtMetaContainerPrivate__RandomAccessCapability  QtMetaContainerPrivate__IteratorCapability = 8
)

type QtMetaContainerPrivate__AddRemoveCapability byte

const (
	QtMetaContainerPrivate__CanAddAtBegin    QtMetaContainerPrivate__AddRemoveCapability = 1
	QtMetaContainerPrivate__CanRemoveAtBegin QtMetaContainerPrivate__AddRemoveCapability = 2
	QtMetaContainerPrivate__CanAddAtEnd      QtMetaContainerPrivate__AddRemoveCapability = 4
	QtMetaContainerPrivate__CanRemoveAtEnd   QtMetaContainerPrivate__AddRemoveCapability = 8
)

type QtMetaContainerPrivate__QMetaContainerInterface__Position byte

const (
	QtMetaContainerPrivate__QMetaContainerInterface__AtBegin     QtMetaContainerPrivate__QMetaContainerInterface__Position = 0
	QtMetaContainerPrivate__QMetaContainerInterface__AtEnd       QtMetaContainerPrivate__QMetaContainerInterface__Position = 1
	QtMetaContainerPrivate__QMetaContainerInterface__Unspecified QtMetaContainerPrivate__QMetaContainerInterface__Position = 2
)

type QMetaContainer struct {
	h          uintptr
	isSubclass bool
}

// NewQMetaContainer constructs a new QMetaContainer object.
func NewQMetaContainer() *QMetaContainer {

	ret := newQMetaContainer(QMetaContainer_new())
	ret.isSubclass = true
	return ret
}

// NewQMetaContainer2 constructs a new QMetaContainer object.
func NewQMetaContainer2(param1 *QMetaContainer) *QMetaContainer {

	ret := newQMetaContainer(QMetaContainer_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QMetaContainer) HasInputIterator() bool {
	return (bool)(QMetaContainer_HasInputIterator(this.h))
}

func (this *QMetaContainer) HasForwardIterator() bool {
	return (bool)(QMetaContainer_HasForwardIterator(this.h))
}

func (this *QMetaContainer) HasBidirectionalIterator() bool {
	return (bool)(QMetaContainer_HasBidirectionalIterator(this.h))
}

func (this *QMetaContainer) HasRandomAccessIterator() bool {
	return (bool)(QMetaContainer_HasRandomAccessIterator(this.h))
}

func (this *QMetaContainer) HasSize() bool {
	return (bool)(QMetaContainer_HasSize(this.h))
}

func (this *QMetaContainer) Size(container unsafe.Pointer) int64 {
	return (int64)(QMetaContainer_Size(this.h, container))
}

func (this *QMetaContainer) CanClear() bool {
	return (bool)(QMetaContainer_CanClear(this.h))
}

func (this *QMetaContainer) Clear(container unsafe.Pointer) {
	QMetaContainer_Clear(this.h, container)
}

func (this *QMetaContainer) HasIterator() bool {
	return (bool)(QMetaContainer_HasIterator(this.h))
}

func (this *QMetaContainer) Begin(container unsafe.Pointer) unsafe.Pointer {
	return (unsafe.Pointer)(QMetaContainer_Begin(this.h, container))
}

func (this *QMetaContainer) End(container unsafe.Pointer) unsafe.Pointer {
	return (unsafe.Pointer)(QMetaContainer_End(this.h, container))
}

func (this *QMetaContainer) DestroyIterator(iterator unsafe.Pointer) {
	QMetaContainer_DestroyIterator(this.h, iterator)
}

func (this *QMetaContainer) CompareIterator(i unsafe.Pointer, j unsafe.Pointer) bool {
	return (bool)(QMetaContainer_CompareIterator(this.h, i, j))
}

func (this *QMetaContainer) CopyIterator(target unsafe.Pointer, source unsafe.Pointer) {
	QMetaContainer_CopyIterator(this.h, target, source)
}

func (this *QMetaContainer) AdvanceIterator(iterator unsafe.Pointer, step int64) {
	QMetaContainer_AdvanceIterator(this.h, iterator, (ptrdiff_t)(step))
}

func (this *QMetaContainer) DiffIterator(i unsafe.Pointer, j unsafe.Pointer) int64 {
	return (int64)(QMetaContainer_DiffIterator(this.h, i, j))
}

func (this *QMetaContainer) HasConstIterator() bool {
	return (bool)(QMetaContainer_HasConstIterator(this.h))
}

func (this *QMetaContainer) ConstBegin(container unsafe.Pointer) unsafe.Pointer {
	return (unsafe.Pointer)(QMetaContainer_ConstBegin(this.h, container))
}

func (this *QMetaContainer) ConstEnd(container unsafe.Pointer) unsafe.Pointer {
	return (unsafe.Pointer)(QMetaContainer_ConstEnd(this.h, container))
}

func (this *QMetaContainer) DestroyConstIterator(iterator unsafe.Pointer) {
	QMetaContainer_DestroyConstIterator(this.h, iterator)
}

func (this *QMetaContainer) CompareConstIterator(i unsafe.Pointer, j unsafe.Pointer) bool {
	return (bool)(QMetaContainer_CompareConstIterator(this.h, i, j))
}

func (this *QMetaContainer) CopyConstIterator(target unsafe.Pointer, source unsafe.Pointer) {
	QMetaContainer_CopyConstIterator(this.h, target, source)
}

func (this *QMetaContainer) AdvanceConstIterator(iterator unsafe.Pointer, step int64) {
	QMetaContainer_AdvanceConstIterator(this.h, iterator, (ptrdiff_t)(step))
}

func (this *QMetaContainer) DiffConstIterator(i unsafe.Pointer, j unsafe.Pointer) int64 {
	return (int64)(QMetaContainer_DiffConstIterator(this.h, i, j))
}

type QMetaSequence struct {
	h          uintptr
	isSubclass bool
}

// NewQMetaSequence constructs a new QMetaSequence object.
func NewQMetaSequence() *QMetaSequence {

	ret := newQMetaSequence(QMetaSequence_new())
	ret.isSubclass = true
	return ret
}

func (this *QMetaSequence) ValueMetaType() *QMetaType {
	_goptr := newQMetaType(QMetaSequence_ValueMetaType(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaSequence) IsSortable() bool {
	return (bool)(QMetaSequence_IsSortable(this.h))
}

func (this *QMetaSequence) CanAddValueAtBegin() bool {
	return (bool)(QMetaSequence_CanAddValueAtBegin(this.h))
}

func (this *QMetaSequence) AddValueAtBegin(container unsafe.Pointer, value unsafe.Pointer) {
	QMetaSequence_AddValueAtBegin(this.h, container, value)
}

func (this *QMetaSequence) CanAddValueAtEnd() bool {
	return (bool)(QMetaSequence_CanAddValueAtEnd(this.h))
}

func (this *QMetaSequence) AddValueAtEnd(container unsafe.Pointer, value unsafe.Pointer) {
	QMetaSequence_AddValueAtEnd(this.h, container, value)
}

func (this *QMetaSequence) CanRemoveValueAtBegin() bool {
	return (bool)(QMetaSequence_CanRemoveValueAtBegin(this.h))
}

func (this *QMetaSequence) RemoveValueAtBegin(container unsafe.Pointer) {
	QMetaSequence_RemoveValueAtBegin(this.h, container)
}

func (this *QMetaSequence) CanRemoveValueAtEnd() bool {
	return (bool)(QMetaSequence_CanRemoveValueAtEnd(this.h))
}

func (this *QMetaSequence) RemoveValueAtEnd(container unsafe.Pointer) {
	QMetaSequence_RemoveValueAtEnd(this.h, container)
}

func (this *QMetaSequence) CanGetValueAtIndex() bool {
	return (bool)(QMetaSequence_CanGetValueAtIndex(this.h))
}

func (this *QMetaSequence) ValueAtIndex(container unsafe.Pointer, index int64, result unsafe.Pointer) {
	QMetaSequence_ValueAtIndex(this.h, container, (ptrdiff_t)(index), result)
}

func (this *QMetaSequence) CanSetValueAtIndex() bool {
	return (bool)(QMetaSequence_CanSetValueAtIndex(this.h))
}

func (this *QMetaSequence) SetValueAtIndex(container unsafe.Pointer, index int64, value unsafe.Pointer) {
	QMetaSequence_SetValueAtIndex(this.h, container, (ptrdiff_t)(index), value)
}

func (this *QMetaSequence) CanAddValue() bool {
	return (bool)(QMetaSequence_CanAddValue(this.h))
}

func (this *QMetaSequence) AddValue(container unsafe.Pointer, value unsafe.Pointer) {
	QMetaSequence_AddValue(this.h, container, value)
}

func (this *QMetaSequence) CanRemoveValue() bool {
	return (bool)(QMetaSequence_CanRemoveValue(this.h))
}

func (this *QMetaSequence) RemoveValue(container unsafe.Pointer) {
	QMetaSequence_RemoveValue(this.h, container)
}

func (this *QMetaSequence) CanGetValueAtIterator() bool {
	return (bool)(QMetaSequence_CanGetValueAtIterator(this.h))
}

func (this *QMetaSequence) ValueAtIterator(iterator unsafe.Pointer, result unsafe.Pointer) {
	QMetaSequence_ValueAtIterator(this.h, iterator, result)
}

func (this *QMetaSequence) CanSetValueAtIterator() bool {
	return (bool)(QMetaSequence_CanSetValueAtIterator(this.h))
}

func (this *QMetaSequence) SetValueAtIterator(iterator unsafe.Pointer, value unsafe.Pointer) {
	QMetaSequence_SetValueAtIterator(this.h, iterator, value)
}

func (this *QMetaSequence) CanInsertValueAtIterator() bool {
	return (bool)(QMetaSequence_CanInsertValueAtIterator(this.h))
}

func (this *QMetaSequence) InsertValueAtIterator(container unsafe.Pointer, iterator unsafe.Pointer, value unsafe.Pointer) {
	QMetaSequence_InsertValueAtIterator(this.h, container, iterator, value)
}

func (this *QMetaSequence) CanEraseValueAtIterator() bool {
	return (bool)(QMetaSequence_CanEraseValueAtIterator(this.h))
}

func (this *QMetaSequence) EraseValueAtIterator(container unsafe.Pointer, iterator unsafe.Pointer) {
	QMetaSequence_EraseValueAtIterator(this.h, container, iterator)
}

func (this *QMetaSequence) CanEraseRangeAtIterator() bool {
	return (bool)(QMetaSequence_CanEraseRangeAtIterator(this.h))
}

func (this *QMetaSequence) EraseRangeAtIterator(container unsafe.Pointer, iterator1 unsafe.Pointer, iterator2 unsafe.Pointer) {
	QMetaSequence_EraseRangeAtIterator(this.h, container, iterator1, iterator2)
}

func (this *QMetaSequence) CanGetValueAtConstIterator() bool {
	return (bool)(QMetaSequence_CanGetValueAtConstIterator(this.h))
}

func (this *QMetaSequence) ValueAtConstIterator(iterator unsafe.Pointer, result unsafe.Pointer) {
	QMetaSequence_ValueAtConstIterator(this.h, iterator, result)
}

type QMetaAssociation struct {
	h          uintptr
	isSubclass bool
}

// NewQMetaAssociation constructs a new QMetaAssociation object.
func NewQMetaAssociation() *QMetaAssociation {

	ret := newQMetaAssociation(QMetaAssociation_new())
	ret.isSubclass = true
	return ret
}

func (this *QMetaAssociation) KeyMetaType() *QMetaType {
	_goptr := newQMetaType(QMetaAssociation_KeyMetaType(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaAssociation) MappedMetaType() *QMetaType {
	_goptr := newQMetaType(QMetaAssociation_MappedMetaType(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaAssociation) CanInsertKey() bool {
	return (bool)(QMetaAssociation_CanInsertKey(this.h))
}

func (this *QMetaAssociation) InsertKey(container unsafe.Pointer, key unsafe.Pointer) {
	QMetaAssociation_InsertKey(this.h, container, key)
}

func (this *QMetaAssociation) CanRemoveKey() bool {
	return (bool)(QMetaAssociation_CanRemoveKey(this.h))
}

func (this *QMetaAssociation) RemoveKey(container unsafe.Pointer, key unsafe.Pointer) {
	QMetaAssociation_RemoveKey(this.h, container, key)
}

func (this *QMetaAssociation) CanContainsKey() bool {
	return (bool)(QMetaAssociation_CanContainsKey(this.h))
}

func (this *QMetaAssociation) ContainsKey(container unsafe.Pointer, key unsafe.Pointer) bool {
	return (bool)(QMetaAssociation_ContainsKey(this.h, container, key))
}

func (this *QMetaAssociation) CanGetMappedAtKey() bool {
	return (bool)(QMetaAssociation_CanGetMappedAtKey(this.h))
}

func (this *QMetaAssociation) MappedAtKey(container unsafe.Pointer, key unsafe.Pointer, mapped unsafe.Pointer) {
	QMetaAssociation_MappedAtKey(this.h, container, key, mapped)
}

func (this *QMetaAssociation) CanSetMappedAtKey() bool {
	return (bool)(QMetaAssociation_CanSetMappedAtKey(this.h))
}

func (this *QMetaAssociation) SetMappedAtKey(container unsafe.Pointer, key unsafe.Pointer, mapped unsafe.Pointer) {
	QMetaAssociation_SetMappedAtKey(this.h, container, key, mapped)
}

func (this *QMetaAssociation) CanGetKeyAtIterator() bool {
	return (bool)(QMetaAssociation_CanGetKeyAtIterator(this.h))
}

func (this *QMetaAssociation) KeyAtIterator(iterator unsafe.Pointer, key unsafe.Pointer) {
	QMetaAssociation_KeyAtIterator(this.h, iterator, key)
}

func (this *QMetaAssociation) CanGetKeyAtConstIterator() bool {
	return (bool)(QMetaAssociation_CanGetKeyAtConstIterator(this.h))
}

func (this *QMetaAssociation) KeyAtConstIterator(iterator unsafe.Pointer, key unsafe.Pointer) {
	QMetaAssociation_KeyAtConstIterator(this.h, iterator, key)
}

func (this *QMetaAssociation) CanGetMappedAtIterator() bool {
	return (bool)(QMetaAssociation_CanGetMappedAtIterator(this.h))
}

func (this *QMetaAssociation) MappedAtIterator(iterator unsafe.Pointer, mapped unsafe.Pointer) {
	QMetaAssociation_MappedAtIterator(this.h, iterator, mapped)
}

func (this *QMetaAssociation) CanGetMappedAtConstIterator() bool {
	return (bool)(QMetaAssociation_CanGetMappedAtConstIterator(this.h))
}

func (this *QMetaAssociation) MappedAtConstIterator(iterator unsafe.Pointer, mapped unsafe.Pointer) {
	QMetaAssociation_MappedAtConstIterator(this.h, iterator, mapped)
}

func (this *QMetaAssociation) CanSetMappedAtIterator() bool {
	return (bool)(QMetaAssociation_CanSetMappedAtIterator(this.h))
}

func (this *QMetaAssociation) SetMappedAtIterator(iterator unsafe.Pointer, mapped unsafe.Pointer) {
	QMetaAssociation_SetMappedAtIterator(this.h, iterator, mapped)
}

func (this *QMetaAssociation) CanCreateIteratorAtKey() bool {
	return (bool)(QMetaAssociation_CanCreateIteratorAtKey(this.h))
}

func (this *QMetaAssociation) CreateIteratorAtKey(container unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return (unsafe.Pointer)(QMetaAssociation_CreateIteratorAtKey(this.h, container, key))
}

func (this *QMetaAssociation) CanCreateConstIteratorAtKey() bool {
	return (bool)(QMetaAssociation_CanCreateConstIteratorAtKey(this.h))
}

func (this *QMetaAssociation) CreateConstIteratorAtKey(container unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return (unsafe.Pointer)(QMetaAssociation_CreateConstIteratorAtKey(this.h, container, key))
}
