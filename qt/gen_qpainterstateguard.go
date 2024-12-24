package qt

import (
	"unsafe"
)

type QPainterStateGuard__InitialState int

const (
	QPainterStateGuard__Save   QPainterStateGuard__InitialState = 0
	QPainterStateGuard__NoSave QPainterStateGuard__InitialState = 1
)

type QPainterStateGuard struct {
	h          uintptr
	isSubclass bool
}

// NewQPainterStateGuard constructs a new QPainterStateGuard object.
func NewQPainterStateGuard(painter *QPainter) *QPainterStateGuard {
	g := newQPainterStateGuard(QPainterStateGuard_new(painter.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPainterStateGuard2 constructs a new QPainterStateGuard object.
func NewQPainterStateGuard2(painter *QPainter, state InitialState) *QPainterStateGuard {
	g := newQPainterStateGuard(QPainterStateGuard_new2(painter.cPointer(), state))
	g.isSubclass = true
	return g
}

func (this *QPainterStateGuard) Save() {
	QPainterStateGuard_Save(this.h)
}

func (this *QPainterStateGuard) Restore() {
	QPainterStateGuard_Restore(this.h)
}
