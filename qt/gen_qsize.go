package qt

import (
	"unsafe"
)

type QSize struct {
	h          uintptr
	isSubclass bool
}

// NewQSize constructs a new QSize object.
func NewQSize() *QSize {
	g := newQSize(QSize_new())
	g.isSubclass = true
	return g
}

// NewQSize2 constructs a new QSize object.
func NewQSize2(w int, h int) *QSize {
	g := newQSize(QSize_new2((int)(w), (int)(h)))
	g.isSubclass = true
	return g
}

// NewQSize3 constructs a new QSize object.
func NewQSize3(param1 *QSize) *QSize {
	g := newQSize(QSize_new3(param1.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QSize) IsNull() bool {
	return (bool)(QSize_IsNull(this.h))
}

func (this *QSize) IsEmpty() bool {
	return (bool)(QSize_IsEmpty(this.h))
}

func (this *QSize) IsValid() bool {
	return (bool)(QSize_IsValid(this.h))
}

func (this *QSize) Width() int {
	return (int)(QSize_Width(this.h))
}

func (this *QSize) Height() int {
	return (int)(QSize_Height(this.h))
}

func (this *QSize) SetWidth(w int) {
	QSize_SetWidth(this.h, (int)(w))
}

func (this *QSize) SetHeight(h int) {
	QSize_SetHeight(this.h, (int)(h))
}

func (this *QSize) Transpose() {
	QSize_Transpose(this.h)
}

func (this *QSize) Transposed() *QSize {
	_goptr := newQSize(QSize_Transposed(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSize) Scale(w int, h int, mode AspectRatioMode) {
	QSize_Scale(this.h, (int)(w), (int)(h), (int)(mode))
}

func (this *QSize) Scale2(s *QSize, mode AspectRatioMode) {
	QSize_Scale2(this.h, s.cPointer(), (int)(mode))
}

func (this *QSize) Scaled(w int, h int, mode AspectRatioMode) *QSize {
	_goptr := newQSize(QSize_Scaled(this.h, (int)(w), (int)(h), (int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSize) Scaled2(s *QSize, mode AspectRatioMode) *QSize {
	_goptr := newQSize(QSize_Scaled2(this.h, s.cPointer(), (int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSize) ExpandedTo(param1 *QSize) *QSize {
	_goptr := newQSize(QSize_ExpandedTo(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSize) BoundedTo(param1 *QSize) *QSize {
	_goptr := newQSize(QSize_BoundedTo(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSize) GrownBy(m QMargins) *QSize {
	_goptr := newQSize(QSize_GrownBy(this.h, m.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSize) ShrunkBy(m QMargins) *QSize {
	_goptr := newQSize(QSize_ShrunkBy(this.h, m.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSize) OperatorPlusAssign(param1 *QSize) *QSize {
	return newQSize(QSize_OperatorPlusAssign(this.h, param1.cPointer()))
}

func (this *QSize) OperatorMinusAssign(param1 *QSize) *QSize {
	return newQSize(QSize_OperatorMinusAssign(this.h, param1.cPointer()))
}

func (this *QSize) OperatorMultiplyAssign(c float64) *QSize {
	return newQSize(QSize_OperatorMultiplyAssign(this.h, (double)(c)))
}

func (this *QSize) OperatorDivideAssign(c float64) *QSize {
	return newQSize(QSize_OperatorDivideAssign(this.h, (double)(c)))
}

func (this *QSize) ToSizeF() *QSizeF {
	_goptr := newQSizeF(QSize_ToSizeF(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QSizeF struct {
	h          uintptr
	isSubclass bool
}

// NewQSizeF constructs a new QSizeF object.
func NewQSizeF() *QSizeF {
	g := newQSizeF(QSizeF_new())
	g.isSubclass = true
	return g
}

// NewQSizeF2 constructs a new QSizeF object.
func NewQSizeF2(sz *QSize) *QSizeF {
	g := newQSizeF(QSizeF_new2(sz.cPointer()))
	g.isSubclass = true
	return g
}

// NewQSizeF3 constructs a new QSizeF object.
func NewQSizeF3(w float64, h float64) *QSizeF {
	g := newQSizeF(QSizeF_new3((double)(w), (double)(h)))
	g.isSubclass = true
	return g
}

// NewQSizeF4 constructs a new QSizeF object.
func NewQSizeF4(param1 *QSizeF) *QSizeF {
	g := newQSizeF(QSizeF_new4(param1.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QSizeF) IsNull() bool {
	return (bool)(QSizeF_IsNull(this.h))
}

func (this *QSizeF) IsEmpty() bool {
	return (bool)(QSizeF_IsEmpty(this.h))
}

func (this *QSizeF) IsValid() bool {
	return (bool)(QSizeF_IsValid(this.h))
}

func (this *QSizeF) Width() float64 {
	return (float64)(QSizeF_Width(this.h))
}

func (this *QSizeF) Height() float64 {
	return (float64)(QSizeF_Height(this.h))
}

func (this *QSizeF) SetWidth(w float64) {
	QSizeF_SetWidth(this.h, (double)(w))
}

func (this *QSizeF) SetHeight(h float64) {
	QSizeF_SetHeight(this.h, (double)(h))
}

func (this *QSizeF) Transpose() {
	QSizeF_Transpose(this.h)
}

func (this *QSizeF) Transposed() *QSizeF {
	_goptr := newQSizeF(QSizeF_Transposed(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSizeF) Scale(w float64, h float64, mode AspectRatioMode) {
	QSizeF_Scale(this.h, (double)(w), (double)(h), (int)(mode))
}

func (this *QSizeF) Scale2(s *QSizeF, mode AspectRatioMode) {
	QSizeF_Scale2(this.h, s.cPointer(), (int)(mode))
}

func (this *QSizeF) Scaled(w float64, h float64, mode AspectRatioMode) *QSizeF {
	_goptr := newQSizeF(QSizeF_Scaled(this.h, (double)(w), (double)(h), (int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSizeF) Scaled2(s *QSizeF, mode AspectRatioMode) *QSizeF {
	_goptr := newQSizeF(QSizeF_Scaled2(this.h, s.cPointer(), (int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSizeF) ExpandedTo(param1 *QSizeF) *QSizeF {
	_goptr := newQSizeF(QSizeF_ExpandedTo(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSizeF) BoundedTo(param1 *QSizeF) *QSizeF {
	_goptr := newQSizeF(QSizeF_BoundedTo(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSizeF) GrownBy(m QMarginsF) *QSizeF {
	_goptr := newQSizeF(QSizeF_GrownBy(this.h, m.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSizeF) ShrunkBy(m QMarginsF) *QSizeF {
	_goptr := newQSizeF(QSizeF_ShrunkBy(this.h, m.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSizeF) OperatorPlusAssign(param1 *QSizeF) *QSizeF {
	return newQSizeF(QSizeF_OperatorPlusAssign(this.h, param1.cPointer()))
}

func (this *QSizeF) OperatorMinusAssign(param1 *QSizeF) *QSizeF {
	return newQSizeF(QSizeF_OperatorMinusAssign(this.h, param1.cPointer()))
}

func (this *QSizeF) OperatorMultiplyAssign(c float64) *QSizeF {
	return newQSizeF(QSizeF_OperatorMultiplyAssign(this.h, (double)(c)))
}

func (this *QSizeF) OperatorDivideAssign(c float64) *QSizeF {
	return newQSizeF(QSizeF_OperatorDivideAssign(this.h, (double)(c)))
}

func (this *QSizeF) ToSize() *QSize {
	_goptr := newQSize(QSizeF_ToSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
