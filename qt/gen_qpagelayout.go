package qt

import (
	"unsafe"
)

type QPageLayout__Unit int

const (
	QPageLayout__Millimeter QPageLayout__Unit = 0
	QPageLayout__Point      QPageLayout__Unit = 1
	QPageLayout__Inch       QPageLayout__Unit = 2
	QPageLayout__Pica       QPageLayout__Unit = 3
	QPageLayout__Didot      QPageLayout__Unit = 4
	QPageLayout__Cicero     QPageLayout__Unit = 5
)

type QPageLayout__Orientation int

const (
	QPageLayout__Portrait  QPageLayout__Orientation = 0
	QPageLayout__Landscape QPageLayout__Orientation = 1
)

type QPageLayout__Mode int

const (
	QPageLayout__StandardMode QPageLayout__Mode = 0
	QPageLayout__FullPageMode QPageLayout__Mode = 1
)

type QPageLayout__OutOfBoundsPolicy int

const (
	QPageLayout__Reject QPageLayout__OutOfBoundsPolicy = 0
	QPageLayout__Clamp  QPageLayout__OutOfBoundsPolicy = 1
)

type QPageLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQPageLayout constructs a new QPageLayout object.
func NewQPageLayout() *QPageLayout {
	g := newQPageLayout(QPageLayout_new())
	g.isSubclass = true
	return g
}

// NewQPageLayout2 constructs a new QPageLayout object.
func NewQPageLayout2(pageSize *QPageSize, orientation Orientation, margins *QMarginsF) *QPageLayout {
	g := newQPageLayout(QPageLayout_new2(pageSize.cPointer(), orientation, margins.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPageLayout3 constructs a new QPageLayout object.
func NewQPageLayout3(other *QPageLayout) *QPageLayout {
	g := newQPageLayout(QPageLayout_new3(other.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPageLayout4 constructs a new QPageLayout object.
func NewQPageLayout4(pageSize *QPageSize, orientation Orientation, margins *QMarginsF, units Unit) *QPageLayout {
	g := newQPageLayout(QPageLayout_new4(pageSize.cPointer(), orientation, margins.cPointer(), units))
	g.isSubclass = true
	return g
}

// NewQPageLayout5 constructs a new QPageLayout object.
func NewQPageLayout5(pageSize *QPageSize, orientation Orientation, margins *QMarginsF, units Unit, minMargins *QMarginsF) *QPageLayout {
	g := newQPageLayout(QPageLayout_new5(pageSize.cPointer(), orientation, margins.cPointer(), units, minMargins.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QPageLayout) OperatorAssign(other *QPageLayout) {
	QPageLayout_OperatorAssign(this.h, other.cPointer())
}

func (this *QPageLayout) Swap(other *QPageLayout) {
	QPageLayout_Swap(this.h, other.cPointer())
}

func (this *QPageLayout) IsEquivalentTo(other *QPageLayout) bool {
	return (bool)(QPageLayout_IsEquivalentTo(this.h, other.cPointer()))
}

func (this *QPageLayout) IsValid() bool {
	return (bool)(QPageLayout_IsValid(this.h))
}

func (this *QPageLayout) SetMode(mode Mode) {
	QPageLayout_SetMode(this.h, mode)
}

func (this *QPageLayout) Mode() Mode {
	xxxxxxxxx
}

func (this *QPageLayout) SetPageSize(pageSize *QPageSize) {
	QPageLayout_SetPageSize(this.h, pageSize.cPointer())
}

func (this *QPageLayout) PageSize() *QPageSize {
	_goptr := newQPageSize(QPageLayout_PageSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPageLayout) SetOrientation(orientation Orientation) {
	QPageLayout_SetOrientation(this.h, orientation)
}

func (this *QPageLayout) Orientation() Orientation {
	xxxxxxxxx
}

func (this *QPageLayout) SetUnits(units Unit) {
	QPageLayout_SetUnits(this.h, units)
}

func (this *QPageLayout) Units() Unit {
	xxxxxxxxx
}

func (this *QPageLayout) SetMargins(margins *QMarginsF) bool {
	return (bool)(QPageLayout_SetMargins(this.h, margins.cPointer()))
}

func (this *QPageLayout) SetLeftMargin(leftMargin float64) bool {
	return (bool)(QPageLayout_SetLeftMargin(this.h, (double)(leftMargin)))
}

func (this *QPageLayout) SetRightMargin(rightMargin float64) bool {
	return (bool)(QPageLayout_SetRightMargin(this.h, (double)(rightMargin)))
}

func (this *QPageLayout) SetTopMargin(topMargin float64) bool {
	return (bool)(QPageLayout_SetTopMargin(this.h, (double)(topMargin)))
}

func (this *QPageLayout) SetBottomMargin(bottomMargin float64) bool {
	return (bool)(QPageLayout_SetBottomMargin(this.h, (double)(bottomMargin)))
}

func (this *QPageLayout) Margins() *QMarginsF {
	_goptr := newQMarginsF(QPageLayout_Margins(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPageLayout) MarginsWithUnits(units Unit) *QMarginsF {
	_goptr := newQMarginsF(QPageLayout_MarginsWithUnits(this.h, units))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPageLayout) MarginsPoints() *QMargins {
	_goptr := newQMargins(QPageLayout_MarginsPoints(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPageLayout) MarginsPixels(resolution int) *QMargins {
	_goptr := newQMargins(QPageLayout_MarginsPixels(this.h, (int)(resolution)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPageLayout) SetMinimumMargins(minMargins *QMarginsF) {
	QPageLayout_SetMinimumMargins(this.h, minMargins.cPointer())
}

func (this *QPageLayout) MinimumMargins() *QMarginsF {
	_goptr := newQMarginsF(QPageLayout_MinimumMargins(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPageLayout) MaximumMargins() *QMarginsF {
	_goptr := newQMarginsF(QPageLayout_MaximumMargins(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPageLayout) FullRect() *QRectF {
	_goptr := newQRectF(QPageLayout_FullRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPageLayout) FullRectWithUnits(units Unit) *QRectF {
	_goptr := newQRectF(QPageLayout_FullRectWithUnits(this.h, units))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPageLayout) FullRectPoints() *QRect {
	_goptr := newQRect(QPageLayout_FullRectPoints(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPageLayout) FullRectPixels(resolution int) *QRect {
	_goptr := newQRect(QPageLayout_FullRectPixels(this.h, (int)(resolution)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPageLayout) PaintRect() *QRectF {
	_goptr := newQRectF(QPageLayout_PaintRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPageLayout) PaintRectWithUnits(units Unit) *QRectF {
	_goptr := newQRectF(QPageLayout_PaintRectWithUnits(this.h, units))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPageLayout) PaintRectPoints() *QRect {
	_goptr := newQRect(QPageLayout_PaintRectPoints(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPageLayout) PaintRectPixels(resolution int) *QRect {
	_goptr := newQRect(QPageLayout_PaintRectPixels(this.h, (int)(resolution)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPageLayout) SetPageSize2(pageSize *QPageSize, minMargins *QMarginsF) {
	QPageLayout_SetPageSize2(this.h, pageSize.cPointer(), minMargins.cPointer())
}

func (this *QPageLayout) SetMargins2(margins *QMarginsF, outOfBoundsPolicy OutOfBoundsPolicy) bool {
	return (bool)(QPageLayout_SetMargins2(this.h, margins.cPointer(), outOfBoundsPolicy))
}

func (this *QPageLayout) SetLeftMargin2(leftMargin float64, outOfBoundsPolicy OutOfBoundsPolicy) bool {
	return (bool)(QPageLayout_SetLeftMargin2(this.h, (double)(leftMargin), outOfBoundsPolicy))
}

func (this *QPageLayout) SetRightMargin2(rightMargin float64, outOfBoundsPolicy OutOfBoundsPolicy) bool {
	return (bool)(QPageLayout_SetRightMargin2(this.h, (double)(rightMargin), outOfBoundsPolicy))
}

func (this *QPageLayout) SetTopMargin2(topMargin float64, outOfBoundsPolicy OutOfBoundsPolicy) bool {
	return (bool)(QPageLayout_SetTopMargin2(this.h, (double)(topMargin), outOfBoundsPolicy))
}

func (this *QPageLayout) SetBottomMargin2(bottomMargin float64, outOfBoundsPolicy OutOfBoundsPolicy) bool {
	return (bool)(QPageLayout_SetBottomMargin2(this.h, (double)(bottomMargin), outOfBoundsPolicy))
}
