package qt

import (
	"unsafe"
)

type QtPrivate__Ordering CompareUnderlyingType

const (
	QtPrivate__Equal      QtPrivate__Ordering = 0
	QtPrivate__Equivalent QtPrivate__Ordering = 0
	QtPrivate__Less       QtPrivate__Ordering = -1
	QtPrivate__Greater    QtPrivate__Ordering = 1
)

type QtPrivate__Uncomparable CompareUnderlyingType

const (
	QtPrivate__Uncomparable__Unordered QtPrivate__Uncomparable = -128
)

type QtPrivate__LegacyUncomparable CompareUnderlyingType

const (
	QtPrivate__LegacyUncomparable__Unordered QtPrivate__LegacyUncomparable = -127
)

type partial_ordering struct {
	h          uintptr
	isSubclass bool
}

// Newpartial_ordering constructs a new Qt::partial_ordering object.
func Newpartial_ordering(param1 *partial_ordering) *partial_ordering {
	g := newpartial_ordering(partial_ordering_new(param1))
	g.isSubclass = true
	return g
}

type weak_ordering struct {
	h          uintptr
	isSubclass bool
}

// Newweak_ordering constructs a new Qt::weak_ordering object.
func Newweak_ordering(param1 *weak_ordering) *weak_ordering {
	g := newweak_ordering(weak_ordering_new(param1))
	g.isSubclass = true
	return g
}

type strong_ordering struct {
	h          uintptr
	isSubclass bool
}

// Newstrong_ordering constructs a new Qt::strong_ordering object.
func Newstrong_ordering(param1 *strong_ordering) *strong_ordering {
	g := newstrong_ordering(strong_ordering_new(param1))
	g.isSubclass = true
	return g
}

type QPartialOrdering struct {
	h          uintptr
	isSubclass bool
}

// NewQPartialOrdering constructs a new QPartialOrdering object.
func NewQPartialOrdering(order partial_ordering) *QPartialOrdering {
	g := newQPartialOrdering(QPartialOrdering_new(order.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPartialOrdering2 constructs a new QPartialOrdering object.
func NewQPartialOrdering2(stdorder weak_ordering) *QPartialOrdering {
	g := newQPartialOrdering(QPartialOrdering_new2(stdorder.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPartialOrdering3 constructs a new QPartialOrdering object.
func NewQPartialOrdering3(stdorder strong_ordering) *QPartialOrdering {
	g := newQPartialOrdering(QPartialOrdering_new3(stdorder.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPartialOrdering4 constructs a new QPartialOrdering object.
func NewQPartialOrdering4(param1 *QPartialOrdering) *QPartialOrdering {
	g := newQPartialOrdering(QPartialOrdering_new4(param1.cPointer()))
	g.isSubclass = true
	return g
}
