package qt

import (
	"unsafe"
)

type QTreeWidgetItemIterator__IteratorFlag int

const (
	QTreeWidgetItemIterator__All           QTreeWidgetItemIterator__IteratorFlag = 0
	QTreeWidgetItemIterator__Hidden        QTreeWidgetItemIterator__IteratorFlag = 1
	QTreeWidgetItemIterator__NotHidden     QTreeWidgetItemIterator__IteratorFlag = 2
	QTreeWidgetItemIterator__Selected      QTreeWidgetItemIterator__IteratorFlag = 4
	QTreeWidgetItemIterator__Unselected    QTreeWidgetItemIterator__IteratorFlag = 8
	QTreeWidgetItemIterator__Selectable    QTreeWidgetItemIterator__IteratorFlag = 16
	QTreeWidgetItemIterator__NotSelectable QTreeWidgetItemIterator__IteratorFlag = 32
	QTreeWidgetItemIterator__DragEnabled   QTreeWidgetItemIterator__IteratorFlag = 64
	QTreeWidgetItemIterator__DragDisabled  QTreeWidgetItemIterator__IteratorFlag = 128
	QTreeWidgetItemIterator__DropEnabled   QTreeWidgetItemIterator__IteratorFlag = 256
	QTreeWidgetItemIterator__DropDisabled  QTreeWidgetItemIterator__IteratorFlag = 512
	QTreeWidgetItemIterator__HasChildren   QTreeWidgetItemIterator__IteratorFlag = 1024
	QTreeWidgetItemIterator__NoChildren    QTreeWidgetItemIterator__IteratorFlag = 2048
	QTreeWidgetItemIterator__Checked       QTreeWidgetItemIterator__IteratorFlag = 4096
	QTreeWidgetItemIterator__NotChecked    QTreeWidgetItemIterator__IteratorFlag = 8192
	QTreeWidgetItemIterator__Enabled       QTreeWidgetItemIterator__IteratorFlag = 16384
	QTreeWidgetItemIterator__Disabled      QTreeWidgetItemIterator__IteratorFlag = 32768
	QTreeWidgetItemIterator__Editable      QTreeWidgetItemIterator__IteratorFlag = 65536
	QTreeWidgetItemIterator__NotEditable   QTreeWidgetItemIterator__IteratorFlag = 131072
	QTreeWidgetItemIterator__UserFlag      QTreeWidgetItemIterator__IteratorFlag = 16777216
)

type QTreeWidgetItemIterator struct {
	h          uintptr
	isSubclass bool
}

// NewQTreeWidgetItemIterator constructs a new QTreeWidgetItemIterator object.
func NewQTreeWidgetItemIterator(it *QTreeWidgetItemIterator) *QTreeWidgetItemIterator {
	ret := newQTreeWidgetItemIterator(QTreeWidgetItemIterator_new(it.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItemIterator2 constructs a new QTreeWidgetItemIterator object.
func NewQTreeWidgetItemIterator2(widget *QTreeWidget) *QTreeWidgetItemIterator {
	ret := newQTreeWidgetItemIterator(QTreeWidgetItemIterator_new2(widget.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItemIterator3 constructs a new QTreeWidgetItemIterator object.
func NewQTreeWidgetItemIterator3(item *QTreeWidgetItem) *QTreeWidgetItemIterator {
	ret := newQTreeWidgetItemIterator(QTreeWidgetItemIterator_new3(item.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItemIterator4 constructs a new QTreeWidgetItemIterator object.
func NewQTreeWidgetItemIterator4(widget *QTreeWidget, flags IteratorFlags) *QTreeWidgetItemIterator {
	ret := newQTreeWidgetItemIterator(QTreeWidgetItemIterator_new4(widget.cPointer(), flags))
	ret.isSubclass = true
	return ret
}

// NewQTreeWidgetItemIterator5 constructs a new QTreeWidgetItemIterator object.
func NewQTreeWidgetItemIterator5(item *QTreeWidgetItem, flags IteratorFlags) *QTreeWidgetItemIterator {
	ret := newQTreeWidgetItemIterator(QTreeWidgetItemIterator_new5(item.cPointer(), flags))
	ret.isSubclass = true
	return ret
}

func (this *QTreeWidgetItemIterator) OperatorAssign(it *QTreeWidgetItemIterator) {
	QTreeWidgetItemIterator_OperatorAssign(this.h, it.cPointer())
}

func (this *QTreeWidgetItemIterator) OperatorPlusPlus() *QTreeWidgetItemIterator {
	return newQTreeWidgetItemIterator(QTreeWidgetItemIterator_OperatorPlusPlus(this.h))
}

func (this *QTreeWidgetItemIterator) OperatorPlusPlusWithInt(param1 int) *QTreeWidgetItemIterator {
	_goptr := newQTreeWidgetItemIterator(QTreeWidgetItemIterator_OperatorPlusPlusWithInt(this.h, (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTreeWidgetItemIterator) OperatorPlusAssign(n int) *QTreeWidgetItemIterator {
	return newQTreeWidgetItemIterator(QTreeWidgetItemIterator_OperatorPlusAssign(this.h, (int)(n)))
}

func (this *QTreeWidgetItemIterator) OperatorMinusMinus() *QTreeWidgetItemIterator {
	return newQTreeWidgetItemIterator(QTreeWidgetItemIterator_OperatorMinusMinus(this.h))
}

func (this *QTreeWidgetItemIterator) OperatorMinusMinusWithInt(param1 int) *QTreeWidgetItemIterator {
	_goptr := newQTreeWidgetItemIterator(QTreeWidgetItemIterator_OperatorMinusMinusWithInt(this.h, (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTreeWidgetItemIterator) OperatorMinusAssign(n int) *QTreeWidgetItemIterator {
	return newQTreeWidgetItemIterator(QTreeWidgetItemIterator_OperatorMinusAssign(this.h, (int)(n)))
}

func (this *QTreeWidgetItemIterator) OperatorMultiply() *QTreeWidgetItem {
	return newQTreeWidgetItem(QTreeWidgetItemIterator_OperatorMultiply(this.h))
}
