package qt

import (
	"unsafe"
)

type QRect struct {
	h          uintptr
	isSubclass bool
}

// NewQRect constructs a new QRect object.
func NewQRect() *QRect {
	ret := newQRect(QRect_new())
	ret.isSubclass = true
	return ret
}

// NewQRect2 constructs a new QRect object.
func NewQRect2(topleft *QPoint, bottomright *QPoint) *QRect {
	ret := newQRect(QRect_new2(topleft.cPointer(), bottomright.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQRect3 constructs a new QRect object.
func NewQRect3(topleft *QPoint, size *QSize) *QRect {
	ret := newQRect(QRect_new3(topleft.cPointer(), size.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQRect4 constructs a new QRect object.
func NewQRect4(left int, top int, width int, height int) *QRect {
	ret := newQRect(QRect_new4((int)(left), (int)(top), (int)(width), (int)(height)))
	ret.isSubclass = true
	return ret
}

// NewQRect5 constructs a new QRect object.
func NewQRect5(param1 *QRect) *QRect {
	ret := newQRect(QRect_new5(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QRect) IsNull() bool {
	return (bool)(QRect_IsNull(this.h))
}

func (this *QRect) IsEmpty() bool {
	return (bool)(QRect_IsEmpty(this.h))
}

func (this *QRect) IsValid() bool {
	return (bool)(QRect_IsValid(this.h))
}

func (this *QRect) Left() int {
	return (int)(QRect_Left(this.h))
}

func (this *QRect) Top() int {
	return (int)(QRect_Top(this.h))
}

func (this *QRect) Right() int {
	return (int)(QRect_Right(this.h))
}

func (this *QRect) Bottom() int {
	return (int)(QRect_Bottom(this.h))
}

func (this *QRect) Normalized() *QRect {
	_goptr := newQRect(QRect_Normalized(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) X() int {
	return (int)(QRect_X(this.h))
}

func (this *QRect) Y() int {
	return (int)(QRect_Y(this.h))
}

func (this *QRect) SetLeft(pos int) {
	QRect_SetLeft(this.h, (int)(pos))
}

func (this *QRect) SetTop(pos int) {
	QRect_SetTop(this.h, (int)(pos))
}

func (this *QRect) SetRight(pos int) {
	QRect_SetRight(this.h, (int)(pos))
}

func (this *QRect) SetBottom(pos int) {
	QRect_SetBottom(this.h, (int)(pos))
}

func (this *QRect) SetX(x int) {
	QRect_SetX(this.h, (int)(x))
}

func (this *QRect) SetY(y int) {
	QRect_SetY(this.h, (int)(y))
}

func (this *QRect) SetTopLeft(p *QPoint) {
	QRect_SetTopLeft(this.h, p.cPointer())
}

func (this *QRect) SetBottomRight(p *QPoint) {
	QRect_SetBottomRight(this.h, p.cPointer())
}

func (this *QRect) SetTopRight(p *QPoint) {
	QRect_SetTopRight(this.h, p.cPointer())
}

func (this *QRect) SetBottomLeft(p *QPoint) {
	QRect_SetBottomLeft(this.h, p.cPointer())
}

func (this *QRect) TopLeft() *QPoint {
	_goptr := newQPoint(QRect_TopLeft(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) BottomRight() *QPoint {
	_goptr := newQPoint(QRect_BottomRight(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) TopRight() *QPoint {
	_goptr := newQPoint(QRect_TopRight(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) BottomLeft() *QPoint {
	_goptr := newQPoint(QRect_BottomLeft(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) Center() *QPoint {
	_goptr := newQPoint(QRect_Center(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) MoveLeft(pos int) {
	QRect_MoveLeft(this.h, (int)(pos))
}

func (this *QRect) MoveTop(pos int) {
	QRect_MoveTop(this.h, (int)(pos))
}

func (this *QRect) MoveRight(pos int) {
	QRect_MoveRight(this.h, (int)(pos))
}

func (this *QRect) MoveBottom(pos int) {
	QRect_MoveBottom(this.h, (int)(pos))
}

func (this *QRect) MoveTopLeft(p *QPoint) {
	QRect_MoveTopLeft(this.h, p.cPointer())
}

func (this *QRect) MoveBottomRight(p *QPoint) {
	QRect_MoveBottomRight(this.h, p.cPointer())
}

func (this *QRect) MoveTopRight(p *QPoint) {
	QRect_MoveTopRight(this.h, p.cPointer())
}

func (this *QRect) MoveBottomLeft(p *QPoint) {
	QRect_MoveBottomLeft(this.h, p.cPointer())
}

func (this *QRect) MoveCenter(p *QPoint) {
	QRect_MoveCenter(this.h, p.cPointer())
}

func (this *QRect) Translate(dx int, dy int) {
	QRect_Translate(this.h, (int)(dx), (int)(dy))
}

func (this *QRect) TranslateWithQPoint(p *QPoint) {
	QRect_TranslateWithQPoint(this.h, p.cPointer())
}

func (this *QRect) Translated(dx int, dy int) *QRect {
	_goptr := newQRect(QRect_Translated(this.h, (int)(dx), (int)(dy)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) TranslatedWithQPoint(p *QPoint) *QRect {
	_goptr := newQRect(QRect_TranslatedWithQPoint(this.h, p.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) Transposed() *QRect {
	_goptr := newQRect(QRect_Transposed(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) MoveTo(x int, t int) {
	QRect_MoveTo(this.h, (int)(x), (int)(t))
}

func (this *QRect) MoveToWithQPoint(p *QPoint) {
	QRect_MoveToWithQPoint(this.h, p.cPointer())
}

func (this *QRect) SetRect(x int, y int, w int, h int) {
	QRect_SetRect(this.h, (int)(x), (int)(y), (int)(w), (int)(h))
}

func (this *QRect) GetRect(x *int, y *int, w *int, h *int) {
	QRect_GetRect(this.h, (*int)(unsafe.Pointer(x)), (*int)(unsafe.Pointer(y)), (*int)(unsafe.Pointer(w)), (*int)(unsafe.Pointer(h)))
}

func (this *QRect) SetCoords(x1 int, y1 int, x2 int, y2 int) {
	QRect_SetCoords(this.h, (int)(x1), (int)(y1), (int)(x2), (int)(y2))
}

func (this *QRect) GetCoords(x1 *int, y1 *int, x2 *int, y2 *int) {
	QRect_GetCoords(this.h, (*int)(unsafe.Pointer(x1)), (*int)(unsafe.Pointer(y1)), (*int)(unsafe.Pointer(x2)), (*int)(unsafe.Pointer(y2)))
}

func (this *QRect) Adjust(x1 int, y1 int, x2 int, y2 int) {
	QRect_Adjust(this.h, (int)(x1), (int)(y1), (int)(x2), (int)(y2))
}

func (this *QRect) Adjusted(x1 int, y1 int, x2 int, y2 int) *QRect {
	_goptr := newQRect(QRect_Adjusted(this.h, (int)(x1), (int)(y1), (int)(x2), (int)(y2)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) Size() *QSize {
	_goptr := newQSize(QRect_Size(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) Width() int {
	return (int)(QRect_Width(this.h))
}

func (this *QRect) Height() int {
	return (int)(QRect_Height(this.h))
}

func (this *QRect) SetWidth(w int) {
	QRect_SetWidth(this.h, (int)(w))
}

func (this *QRect) SetHeight(h int) {
	QRect_SetHeight(this.h, (int)(h))
}

func (this *QRect) SetSize(s *QSize) {
	QRect_SetSize(this.h, s.cPointer())
}

func (this *QRect) OperatorBitwiseOr(r *QRect) *QRect {
	_goptr := newQRect(QRect_OperatorBitwiseOr(this.h, r.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) OperatorBitwiseAnd(r *QRect) *QRect {
	_goptr := newQRect(QRect_OperatorBitwiseAnd(this.h, r.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) OperatorBitwiseOrAssign(r *QRect) {
	QRect_OperatorBitwiseOrAssign(this.h, r.cPointer())
}

func (this *QRect) OperatorBitwiseAndAssign(r *QRect) {
	QRect_OperatorBitwiseAndAssign(this.h, r.cPointer())
}

func (this *QRect) Contains(r *QRect) bool {
	return (bool)(QRect_Contains(this.h, r.cPointer()))
}

func (this *QRect) ContainsWithQPoint(p *QPoint) bool {
	return (bool)(QRect_ContainsWithQPoint(this.h, p.cPointer()))
}

func (this *QRect) Contains2(x int, y int) bool {
	return (bool)(QRect_Contains2(this.h, (int)(x), (int)(y)))
}

func (this *QRect) Contains3(x int, y int, proper bool) bool {
	return (bool)(QRect_Contains3(this.h, (int)(x), (int)(y), (bool)(proper)))
}

func (this *QRect) United(other *QRect) *QRect {
	_goptr := newQRect(QRect_United(this.h, other.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) Intersected(other *QRect) *QRect {
	_goptr := newQRect(QRect_Intersected(this.h, other.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) Intersects(r *QRect) bool {
	return (bool)(QRect_Intersects(this.h, r.cPointer()))
}

func (this *QRect) MarginsAdded(margins *QMargins) *QRect {
	_goptr := newQRect(QRect_MarginsAdded(this.h, margins.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) MarginsRemoved(margins *QMargins) *QRect {
	_goptr := newQRect(QRect_MarginsRemoved(this.h, margins.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) OperatorPlusAssign(margins *QMargins) *QRect {
	return newQRect(QRect_OperatorPlusAssign(this.h, margins.cPointer()))
}

func (this *QRect) OperatorMinusAssign(margins *QMargins) *QRect {
	return newQRect(QRect_OperatorMinusAssign(this.h, margins.cPointer()))
}

func QRect_Span(p1 *QPoint, p2 *QPoint) *QRect {
	_goptr := newQRect(QRect_Span(p1.cPointer(), p2.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) ToRectF() *QRectF {
	_goptr := newQRectF(QRect_ToRectF(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRect) Contains22(r *QRect, proper bool) bool {
	return (bool)(QRect_Contains22(this.h, r.cPointer(), (bool)(proper)))
}

func (this *QRect) Contains23(p *QPoint, proper bool) bool {
	return (bool)(QRect_Contains23(this.h, p.cPointer(), (bool)(proper)))
}

type QRectF struct {
	h          uintptr
	isSubclass bool
}

// NewQRectF constructs a new QRectF object.
func NewQRectF() *QRectF {
	ret := newQRectF(QRectF_new())
	ret.isSubclass = true
	return ret
}

// NewQRectF2 constructs a new QRectF object.
func NewQRectF2(topleft *QPointF, size *QSizeF) *QRectF {
	ret := newQRectF(QRectF_new2(topleft.cPointer(), size.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQRectF3 constructs a new QRectF object.
func NewQRectF3(topleft *QPointF, bottomRight *QPointF) *QRectF {
	ret := newQRectF(QRectF_new3(topleft.cPointer(), bottomRight.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQRectF4 constructs a new QRectF object.
func NewQRectF4(left float64, top float64, width float64, height float64) *QRectF {
	ret := newQRectF(QRectF_new4((double)(left), (double)(top), (double)(width), (double)(height)))
	ret.isSubclass = true
	return ret
}

// NewQRectF5 constructs a new QRectF object.
func NewQRectF5(rect *QRect) *QRectF {
	ret := newQRectF(QRectF_new5(rect.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQRectF6 constructs a new QRectF object.
func NewQRectF6(param1 *QRectF) *QRectF {
	ret := newQRectF(QRectF_new6(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QRectF) IsNull() bool {
	return (bool)(QRectF_IsNull(this.h))
}

func (this *QRectF) IsEmpty() bool {
	return (bool)(QRectF_IsEmpty(this.h))
}

func (this *QRectF) IsValid() bool {
	return (bool)(QRectF_IsValid(this.h))
}

func (this *QRectF) Normalized() *QRectF {
	_goptr := newQRectF(QRectF_Normalized(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) Left() float64 {
	return (float64)(QRectF_Left(this.h))
}

func (this *QRectF) Top() float64 {
	return (float64)(QRectF_Top(this.h))
}

func (this *QRectF) Right() float64 {
	return (float64)(QRectF_Right(this.h))
}

func (this *QRectF) Bottom() float64 {
	return (float64)(QRectF_Bottom(this.h))
}

func (this *QRectF) X() float64 {
	return (float64)(QRectF_X(this.h))
}

func (this *QRectF) Y() float64 {
	return (float64)(QRectF_Y(this.h))
}

func (this *QRectF) SetLeft(pos float64) {
	QRectF_SetLeft(this.h, (double)(pos))
}

func (this *QRectF) SetTop(pos float64) {
	QRectF_SetTop(this.h, (double)(pos))
}

func (this *QRectF) SetRight(pos float64) {
	QRectF_SetRight(this.h, (double)(pos))
}

func (this *QRectF) SetBottom(pos float64) {
	QRectF_SetBottom(this.h, (double)(pos))
}

func (this *QRectF) SetX(pos float64) {
	QRectF_SetX(this.h, (double)(pos))
}

func (this *QRectF) SetY(pos float64) {
	QRectF_SetY(this.h, (double)(pos))
}

func (this *QRectF) TopLeft() *QPointF {
	_goptr := newQPointF(QRectF_TopLeft(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) BottomRight() *QPointF {
	_goptr := newQPointF(QRectF_BottomRight(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) TopRight() *QPointF {
	_goptr := newQPointF(QRectF_TopRight(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) BottomLeft() *QPointF {
	_goptr := newQPointF(QRectF_BottomLeft(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) Center() *QPointF {
	_goptr := newQPointF(QRectF_Center(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) SetTopLeft(p *QPointF) {
	QRectF_SetTopLeft(this.h, p.cPointer())
}

func (this *QRectF) SetBottomRight(p *QPointF) {
	QRectF_SetBottomRight(this.h, p.cPointer())
}

func (this *QRectF) SetTopRight(p *QPointF) {
	QRectF_SetTopRight(this.h, p.cPointer())
}

func (this *QRectF) SetBottomLeft(p *QPointF) {
	QRectF_SetBottomLeft(this.h, p.cPointer())
}

func (this *QRectF) MoveLeft(pos float64) {
	QRectF_MoveLeft(this.h, (double)(pos))
}

func (this *QRectF) MoveTop(pos float64) {
	QRectF_MoveTop(this.h, (double)(pos))
}

func (this *QRectF) MoveRight(pos float64) {
	QRectF_MoveRight(this.h, (double)(pos))
}

func (this *QRectF) MoveBottom(pos float64) {
	QRectF_MoveBottom(this.h, (double)(pos))
}

func (this *QRectF) MoveTopLeft(p *QPointF) {
	QRectF_MoveTopLeft(this.h, p.cPointer())
}

func (this *QRectF) MoveBottomRight(p *QPointF) {
	QRectF_MoveBottomRight(this.h, p.cPointer())
}

func (this *QRectF) MoveTopRight(p *QPointF) {
	QRectF_MoveTopRight(this.h, p.cPointer())
}

func (this *QRectF) MoveBottomLeft(p *QPointF) {
	QRectF_MoveBottomLeft(this.h, p.cPointer())
}

func (this *QRectF) MoveCenter(p *QPointF) {
	QRectF_MoveCenter(this.h, p.cPointer())
}

func (this *QRectF) Translate(dx float64, dy float64) {
	QRectF_Translate(this.h, (double)(dx), (double)(dy))
}

func (this *QRectF) TranslateWithQPointF(p *QPointF) {
	QRectF_TranslateWithQPointF(this.h, p.cPointer())
}

func (this *QRectF) Translated(dx float64, dy float64) *QRectF {
	_goptr := newQRectF(QRectF_Translated(this.h, (double)(dx), (double)(dy)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) TranslatedWithQPointF(p *QPointF) *QRectF {
	_goptr := newQRectF(QRectF_TranslatedWithQPointF(this.h, p.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) Transposed() *QRectF {
	_goptr := newQRectF(QRectF_Transposed(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) MoveTo(x float64, y float64) {
	QRectF_MoveTo(this.h, (double)(x), (double)(y))
}

func (this *QRectF) MoveToWithQPointF(p *QPointF) {
	QRectF_MoveToWithQPointF(this.h, p.cPointer())
}

func (this *QRectF) SetRect(x float64, y float64, w float64, h float64) {
	QRectF_SetRect(this.h, (double)(x), (double)(y), (double)(w), (double)(h))
}

func (this *QRectF) GetRect(x *float64, y *float64, w *float64, h *float64) {
	QRectF_GetRect(this.h, (*double)(unsafe.Pointer(x)), (*double)(unsafe.Pointer(y)), (*double)(unsafe.Pointer(w)), (*double)(unsafe.Pointer(h)))
}

func (this *QRectF) SetCoords(x1 float64, y1 float64, x2 float64, y2 float64) {
	QRectF_SetCoords(this.h, (double)(x1), (double)(y1), (double)(x2), (double)(y2))
}

func (this *QRectF) GetCoords(x1 *float64, y1 *float64, x2 *float64, y2 *float64) {
	QRectF_GetCoords(this.h, (*double)(unsafe.Pointer(x1)), (*double)(unsafe.Pointer(y1)), (*double)(unsafe.Pointer(x2)), (*double)(unsafe.Pointer(y2)))
}

func (this *QRectF) Adjust(x1 float64, y1 float64, x2 float64, y2 float64) {
	QRectF_Adjust(this.h, (double)(x1), (double)(y1), (double)(x2), (double)(y2))
}

func (this *QRectF) Adjusted(x1 float64, y1 float64, x2 float64, y2 float64) *QRectF {
	_goptr := newQRectF(QRectF_Adjusted(this.h, (double)(x1), (double)(y1), (double)(x2), (double)(y2)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) Size() *QSizeF {
	_goptr := newQSizeF(QRectF_Size(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) Width() float64 {
	return (float64)(QRectF_Width(this.h))
}

func (this *QRectF) Height() float64 {
	return (float64)(QRectF_Height(this.h))
}

func (this *QRectF) SetWidth(w float64) {
	QRectF_SetWidth(this.h, (double)(w))
}

func (this *QRectF) SetHeight(h float64) {
	QRectF_SetHeight(this.h, (double)(h))
}

func (this *QRectF) SetSize(s *QSizeF) {
	QRectF_SetSize(this.h, s.cPointer())
}

func (this *QRectF) OperatorBitwiseOr(r *QRectF) *QRectF {
	_goptr := newQRectF(QRectF_OperatorBitwiseOr(this.h, r.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) OperatorBitwiseAnd(r *QRectF) *QRectF {
	_goptr := newQRectF(QRectF_OperatorBitwiseAnd(this.h, r.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) OperatorBitwiseOrAssign(r *QRectF) {
	QRectF_OperatorBitwiseOrAssign(this.h, r.cPointer())
}

func (this *QRectF) OperatorBitwiseAndAssign(r *QRectF) {
	QRectF_OperatorBitwiseAndAssign(this.h, r.cPointer())
}

func (this *QRectF) Contains(r *QRectF) bool {
	return (bool)(QRectF_Contains(this.h, r.cPointer()))
}

func (this *QRectF) ContainsWithQPointF(p *QPointF) bool {
	return (bool)(QRectF_ContainsWithQPointF(this.h, p.cPointer()))
}

func (this *QRectF) Contains2(x float64, y float64) bool {
	return (bool)(QRectF_Contains2(this.h, (double)(x), (double)(y)))
}

func (this *QRectF) United(other *QRectF) *QRectF {
	_goptr := newQRectF(QRectF_United(this.h, other.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) Intersected(other *QRectF) *QRectF {
	_goptr := newQRectF(QRectF_Intersected(this.h, other.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) Intersects(r *QRectF) bool {
	return (bool)(QRectF_Intersects(this.h, r.cPointer()))
}

func (this *QRectF) MarginsAdded(margins *QMarginsF) *QRectF {
	_goptr := newQRectF(QRectF_MarginsAdded(this.h, margins.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) MarginsRemoved(margins *QMarginsF) *QRectF {
	_goptr := newQRectF(QRectF_MarginsRemoved(this.h, margins.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) OperatorPlusAssign(margins *QMarginsF) *QRectF {
	return newQRectF(QRectF_OperatorPlusAssign(this.h, margins.cPointer()))
}

func (this *QRectF) OperatorMinusAssign(margins *QMarginsF) *QRectF {
	return newQRectF(QRectF_OperatorMinusAssign(this.h, margins.cPointer()))
}

func (this *QRectF) ToRect() *QRect {
	_goptr := newQRect(QRectF_ToRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRectF) ToAlignedRect() *QRect {
	_goptr := newQRect(QRectF_ToAlignedRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
